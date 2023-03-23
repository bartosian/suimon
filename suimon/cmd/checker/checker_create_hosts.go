package checker

import (
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
)

// createHosts creates a list of "Host" values based on the provided "AddressInfo" values
// and returns the list along with an error, if any. This method belongs to the "Checker"
// struct and operates on an instance of the struct passed as a pointer receiver.
// Parameters:
// - addresses: a slice of "AddressInfo" values representing the addresses to create hosts for.
// Returns:
// - a slice of "Host" values created from the provided "AddressInfo" values.
// - an error, if any occurred during host creation.
func (checker *Checker) createHosts(tableType enums.TableType, addresses []AddressInfo) ([]Host, error) {
	var (
		hostCH             = make(chan Host, len(addresses))
		hosts              = make([]Host, 0, len(addresses))
		processedAddresses = make(map[string]struct{})
	)

	defer close(hostCH)

	for _, addressInfo := range addresses {
		address := addressInfo.HostPort.Address
		if _, ok := processedAddresses[address]; ok {
			continue
		}

		processedAddresses[address] = struct{}{}

		go func(addressInfo AddressInfo) {
			host := newHost(tableType, addressInfo, checker.ipClient, checker.httpClient)

			if checker.suimonConfig.IPLookup.AccessToken != "" {
				host.SetLocation()
			}

			host.GetData()

			hostCH <- *host
		}(addressInfo)
	}

	for i := 0; i < len(addresses); i++ {
		hosts = append(hosts, <-hostCH)
	}

	return hosts, nil
}
