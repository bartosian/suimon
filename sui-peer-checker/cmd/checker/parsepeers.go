package checker

import (
	"errors"
	"net/http"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/oschwald/geoip2-golang"

	"github.com/bartosian/sui_helpers/sui-peer-checker/pkg/validation"
)

const (
	httpClientTimeout = 2 * time.Second
	pathToGeoDB       = "./vendors/geodb/GeoLite2-Country.mmdb"
)

type Config struct {
	SeedPeers []PeerData `yaml:"seed-peers"`
}

func parsePeers(seedPeers []PeerData) ([]Peer, error) {
	filePath, err := filepath.Abs(pathToGeoDB)
	if err != nil {
		return nil, err
	}

	db, err := geoip2.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	httpClient := &http.Client{
		Timeout: httpClientTimeout,
	}

	cfgPeers, peers := seedPeers, make([]Peer, 0, len(seedPeers))
	if len(seedPeers) == 0 {
		return nil, errors.New("no peers found in config file")
	}

	var (
		wg             sync.WaitGroup
		peerCH         = make(chan Peer)
		processedPeers = make(map[string]struct{})
	)

	for _, cfgPeer := range cfgPeers {
		if _, ok := processedPeers[cfgPeer.Address]; ok {
			continue
		}

		processedPeers[cfgPeer.Address] = struct{}{}

		wg.Add(1)

		go func(cfgPeer PeerData) {
			defer wg.Done()

			if !validation.IsValidCharCount(cfgPeer.Address, peerSeparator, peerCount) {
				return
			}

			peerInfo := strings.Split(cfgPeer.Address, peerSeparator)

			peer := newPeer(db, httpClient, peerInfo[2], peerInfo[4])
			err := peer.Parse()
			if err != nil {
				return
			}

			doneCH := make(chan struct{})

			go func() {
				peer.GetTotalTransactionNumber()

				doneCH <- struct{}{}
			}()

			go func() {
				peer.GetMetrics()

				doneCH <- struct{}{}
			}()

			for i := 0; i < 2; i++ {
				<-doneCH
			}

			defer close(doneCH)

			peerCH <- *peer
		}(cfgPeer)
	}

	go func() {
		wg.Wait()
		close(peerCH)
	}()

	for peer := range peerCH {
		peers = append(peers, peer)
	}

	return peers, nil
}
