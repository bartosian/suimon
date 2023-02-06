package checker

import (
	"errors"
	"path/filepath"
	"strings"
	"sync"

	"github.com/oschwald/geoip2-golang"
)

type P2PConfig struct {
	SeedPeers []PeerData `yaml:"seed-peers"`
}

func (config *P2PConfig) parsePeers() ([]Peer, error) {
	filePath, _ := filepath.Abs("./vendors/geodb/GeoLite2-Country.mmdb")

	db, err := geoip2.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	configPeers := config.SeedPeers
	checkerPeers := make([]Peer, 0, len(configPeers))
	peerCH := make(chan Peer, len(checkerPeers))

	var wg sync.WaitGroup

	for _, peer := range configPeers {
		wg.Add(1)

		go func(peer PeerData) {
			defer wg.Done()

			if !isValidCharCount(peer.Address, peerSeparator, peerCount) {
				return
			}

			peerInfo := strings.Split(peer.Address, peerSeparator)

			checkerPeer := newPeer(db, peerInfo[2], peerInfo[4])

			err := checkerPeer.Parse()
			if err != nil {
				return
			}

			checkerPeer.GetTotalTransactionNumber()

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
