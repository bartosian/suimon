package controller

import (
	"sync"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/host"
)

// createHosts creates a list of Host objects based on the specified table type and address information.
// The function creates a new Host object for each address in the specified list and sets the Host's internal state based on the specified table type.
// Returns a slice of Host objects and an error value if the creation process fails for any reason.
func (checker *CheckerController) createHosts(tableType enums.TableType, addresses []host.AddressInfo) ([]host.Host, error) {
	hosts := make([]host.Host, 0, len(addresses))
	processedAddresses := make(map[string]struct{})
	hostCH := make(chan host.Host, len(addresses))

	var wg sync.WaitGroup

	for _, addressInfo := range addresses {
		address := addressInfo.HostPort.Address
		if _, ok := processedAddresses[address]; ok {
			continue
		}

		processedAddresses[address] = struct{}{}

		wg.Add(1)

		go func(addressInfo host.AddressInfo) {
			defer wg.Done()

			createdHost := host.NewHost(tableType, addressInfo, checker.clients.ipClient, checker.clients.httpClient)

			if checker.suimonConfig.IPLookup.AccessToken != "" {
				createdHost.SetLocation()
			}

			if err := createdHost.GetData(); err != nil && tableType != enums.TableTypePeers {
				checker.logger.Error("failed to get host data: ", err)

				return
			}

			hostCH <- *createdHost
		}(addressInfo)
	}

	go func() {
		wg.Wait()
		close(hostCH)
	}()

	for h := range hostCH {
		hosts = append(hosts, h)
	}

	return hosts, nil
}
