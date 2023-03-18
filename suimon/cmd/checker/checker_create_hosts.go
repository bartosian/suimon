package checker

import (
	"sync"
)

// createHosts creates a list of "Host" values based on the provided "AddressInfo" values
// and returns the list along with an error, if any. This method belongs to the "Checker"
// struct and operates on an instance of the struct passed as a pointer receiver.
// Parameters:
// - addresses: a slice of "AddressInfo" values representing the addresses to create hosts for.
// Returns:
// - a slice of "Host" values created from the provided "AddressInfo" values.
// - an error, if any occurred during host creation.
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

			host := newHost(addressInfo, checker.suimonConfig.Network.EpochLengthSeconds, checker.ipClient, checker.httpClient)

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
