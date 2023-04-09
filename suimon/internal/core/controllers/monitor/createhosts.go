package monitor

import (
	"sync"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/host"
	"github.com/bartosian/sui_helpers/suimon/internal/core/gateways/geogw"
	"github.com/bartosian/sui_helpers/suimon/internal/core/gateways/prometheusgw"
	"github.com/bartosian/sui_helpers/suimon/internal/core/gateways/rpcgw"
)

type responseWithError struct {
	response *host.Host
	err      error
}

// createHosts creates a list of Host objects based on the specified table type and address information.
// The function creates a new Host object for each address in the specified list and sets the Host's internal state based on the specified table type.
// Returns a slice of Host objects and an error value if the creation process fails for any reason.
func (c *Controller) createHosts(table enums.TableType, addresses []host.AddressInfo) ([]host.Host, error) {
	hosts := make([]host.Host, 0, len(addresses))
	processedAddresses := make(map[string]struct{})

	respChan := make(chan responseWithError, len(addresses))

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

			rpcGateway := rpcgw.NewGateway(c.logger, "")
			geoGateway := geogw.NewGateway(c.logger, "", "")
			prometheusGateway := prometheusgw.NewGateway(c.logger, "")

			createdHost := host.NewHost(c.logger, table, addressInfo, rpcGateway, geoGateway, prometheusGateway, c.gateways.cli)

			result := responseWithError{response: createdHost}

			if c.config.IPLookup.AccessToken != "" {
				if err := createdHost.SetIPInfo(); err != nil {
					result.err = err
					respChan <- result

					return
				}
			}

			if err := createdHost.GetMetrics(); err != nil && table != enums.TableTypePeers {
				result.err = err
				respChan <- result

				return
			}

			respChan <- result
		}(addressInfo)
	}

	go func() {
		wg.Wait()
		close(respChan)
	}()

	for result := range respChan {
		if result.err != nil {
			return nil, result.err
		}

		hosts = append(hosts, *result.response)
	}

	return hosts, nil
}
