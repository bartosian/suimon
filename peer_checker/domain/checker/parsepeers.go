package checker

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/oschwald/geoip2-golang"
)

func (config *P2PConfig) parsePeers() ([]Peer, error) {
	filePath, _ := filepath.Abs("./vendors/geodb/GeoLite2-Country.mmdb")

	db, err := geoip2.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	configPeers := config.SeedPeers
	checkerPeers := make([]Peer, 0, len(configPeers))

	for _, peer := range configPeers {
		if isValidCharCount(peer.Address, peerSeparator, peerCount) {
			peerInfo := strings.Split(peer.Address, peerSeparator)

			checkerPeer, err := newPeer(db, peerInfo[2], peerInfo[4])
			if err != nil {
				return nil, err
			}

			checkerPeers = append(checkerPeers, *checkerPeer)

			continue
		}

		return nil, fmt.Errorf("invalid peer address value provided %s", peer.Address)
	}

	if len(checkerPeers) == 0 {
		return nil, errors.New("no peers found in config file")
	}

	return checkerPeers, nil
}
