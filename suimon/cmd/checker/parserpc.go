package checker

import (
	"sync"
)

func (checker *Checker) parseRPCHosts() error {
	var (
		wg      sync.WaitGroup
		hostCH  = make(chan RPCHost)
		rpcList []RPCHost
	)

	hosts := checker.suimonConfig.GetRPCByNetwork()
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

	checker.rpcList = rpcList

	return nil
}
