package checker

import (
	"errors"
	"net/http"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/oschwald/geoip2-golang"

	"github.com/bartosian/sui_helpers/peer_checker/pkg/validation"
)

const pathToGeoDB = "./vendors/geodb/GeoLite2-Country.mmdb"

type P2PConfig struct {
	SeedPeers []PeerData `yaml:"seed-peers"`
}

func (config *P2PConfig) parsePeers() ([]Peer, error) {
	filePath, _ := filepath.Abs(pathToGeoDB)

	db, err := geoip2.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	httpClient := &http.Client{
		Timeout: 2 * time.Second,
	}

	configPeers, checkerPeers := config.SeedPeers, make([]Peer, 0, len(config.SeedPeers))
	if len(config.SeedPeers) == 0 {
		return nil, errors.New("no peers found in config file")
	}

	peerCH := make(chan Peer, len(checkerPeers))

	var wg sync.WaitGroup

	for _, peer := range configPeers {
		wg.Add(3)

		go func(peer PeerData) {
			defer wg.Done()

			if !validation.IsValidCharCount(peer.Address, peerSeparator, peerCount) {
				return
			}

			peerInfo := strings.Split(peer.Address, peerSeparator)

			checkerPeer := newPeer(db, httpClient, peerInfo[2], peerInfo[4])

			err := checkerPeer.Parse()
			if err != nil {
				return
			}

			doneCH := make(chan struct{}, 2)

			go func() {
				defer wg.Done()

				checkerPeer.GetTotalTransactionNumber()

				doneCH <- struct{}{}
			}()

			go func() {
				defer wg.Done()

				checkerPeer.GetMetrics()

				doneCH <- struct{}{}
			}()

			for i := 0; i < 2; i++ {
				<-doneCH
			}

			peerCH <- *checkerPeer
		}(peer)
	}

	go func() {
		wg.Wait()
		close(peerCH)
	}()

	for peer := range peerCH {
		checkerPeers = append(checkerPeers, peer)
	}

	if len(checkerPeers) == 0 {
		return nil, errors.New("no peers found in config file")
	}

	return checkerPeers, nil
}
