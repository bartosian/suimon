package checker

import (
	"sync"
)

func (checker *Checker) createHosts(addresses []AddressInfo) ([]Host, error) {
	var (
		wg                 sync.WaitGroup
		hostCH             = make(chan Host)
		processedAddresses = make(map[string]struct{})
		hosts              = make([]Host, 0, len(addresses))
	)

	for _, addressInfo := range addresses {
		address := addressInfo.HostPort.Address
		if _, ok := processedAddresses[address]; ok {
			continue
		}

		processedAddresses[address] = struct{}{}

		wg.Add(1)

		go func(addressInfo AddressInfo) {
			defer wg.Done()

			host := newHost(addressInfo, checker.ipClient)

			if checker.suimonConfig.IPLookup.AccessToken != "" {
				host.SetLocation()
			}

			doneCH := make(chan struct{})

			go func() {
				host.GetTotalTransactionNumber()

				doneCH <- struct{}{}
			}()

			go func() {
				host.GetLatestCheckpoint()

				doneCH <- struct{}{}
			}()

			go func() {
				host.GetMetrics(checker.httpClient)

				doneCH <- struct{}{}
			}()

			for i := 0; i < 3; i++ {
				<-doneCH
			}

			defer close(doneCH)

			hostCH <- *host
		}(addressInfo)
	}

	go func() {
		wg.Wait()
		close(hostCH)
	}()

	for host := range hostCH {
		hosts = append(hosts, host)
	}

	return hosts, nil
}
