package controller

import (
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/host"
)

// createHosts creates a list of "Host" values based on the provided "AddressInfo" values
// and returns the list along with an error, if any. This method belongs to the "Checker"
// struct and operates on an instance of the struct passed as a pointer receiver.
// Parameters:
// - addresses: a slice of "AddressInfo" values representing the addresses to create hosts for.
// Returns:
// - a slice of "Host" values created from the provided "AddressInfo" values.
// - an error, if any occurred during host creation.
func (checker CheckerController) createHosts(tableType enums.TableType, addresses []host.AddressInfo) ([]host.Host, error) {
	var (
		hostCH             = make(chan host.Host, len(addresses))
		hosts              = make([]host.Host, 0, len(addresses))
		processedAddresses = make(map[string]struct{})
	)

	defer close(hostCH)

	for _, addressInfo := range addresses {
		address := addressInfo.HostPort.Address
		if _, ok := processedAddresses[address]; ok {
			continue
		}

		processedAddresses[address] = struct{}{}

		go func(addressInfo host.AddressInfo) {
			host := host.NewHost(tableType, addressInfo, checker.ipClient, checker.httpClient)

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
