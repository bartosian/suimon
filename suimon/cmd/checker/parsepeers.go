package checker

import (
	"strings"
	"sync"

	"github.com/bartosian/sui_helpers/suimon/pkg/validation"
)

func (checker *Checker) parsePeers() error {
	var (
		wg              sync.WaitGroup
		peerCH          = make(chan Peer)
		processedPeers  = make(map[string]struct{})
		nodeConfigPeers = checker.nodeConfig.P2PConfig.SeedPeers
		peers           = make([]Peer, 0, len(nodeConfigPeers))
	)

	if len(nodeConfigPeers) == 0 {
		return nil
	}

	for _, nodePeer := range nodeConfigPeers {
		if _, ok := processedPeers[nodePeer.Address]; ok {
			continue
		}

		processedPeers[nodePeer.Address] = struct{}{}

		wg.Add(1)

		go func(address string) {
			defer wg.Done()

			if !validation.IsValidCharCount(address, peerSeparator, peerCount) {
				return
			}

			peerInfo := strings.Split(address, peerSeparator)
			peer := newPeer(checker.ipClient, checker.httpClient, peerInfo[2], peerInfo[4])

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
		}(nodePeer.Address)
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
