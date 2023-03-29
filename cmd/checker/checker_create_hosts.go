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

			host := newHost(addressInfo, checker.ipClient, checker.httpClient)

			if checker.suimonConfig.IPLookup.AccessToken != "" {
				host.SetLocation()
			}

			host.GetData()

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
