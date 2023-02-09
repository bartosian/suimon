package checker

import (
	"strings"
	"sync"

	"github.com/bartosian/sui_helpers/suimon/pkg/validation"
)

func (checker *Checker) parsePeers() error {
	var (
		wg             sync.WaitGroup
		peerCH         = make(chan Peer)
		processedPeers = make(map[string]struct{})
		cfgPeers       = checker.nodeYaml.Config.SeedPeers
		peers          = make([]Peer, 0, len(checker.nodeYaml.Config.SeedPeers))
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
			peer := newPeer(checker.geoDbClient, checker.httpClient, peerInfo[2], peerInfo[4])
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

	checker.peers = peers

	return nil
}
