package checker

import (
	"sync"

	"github.com/bartosian/sui_helpers/sui-peer-checker/cmd/checker/enums"
)

type RpcList map[enums.NetworkType][]string

func parseRPCHosts(hosts []string) ([]RPCHost, error) {
	var (
		wg      sync.WaitGroup
		hostCH  = make(chan RPCHost)
		rpcList = make([]RPCHost, 0, len(hosts))
	)

	for _, host := range hosts {
		wg.Add(1)

		go func(host string) {
			defer wg.Done()

			newRPC := newRPCHost(host)
			doneCH := make(chan struct{})

			go func() {
				newRPC.GetTotalTransactionNumber()

				doneCH <- struct{}{}
			}()

			go func() {
				newRPC.GetLatestCheckpoint()

				doneCH <- struct{}{}
			}()

			for i := 0; i < 2; i++ {
				<-doneCH
			}

			defer close(doneCH)

			hostCH <- *newRPC
		}(host)
	}

	go func() {
		wg.Wait()
		close(hostCH)
	}()

	for rpc := range hostCH {
		rpc.SetStatus()

		rpcList = append(rpcList, rpc)
	}

	return rpcList, nil
}
