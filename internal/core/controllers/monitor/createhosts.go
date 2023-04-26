package monitor

import (
	"sync"

	"github.com/hashicorp/go-multierror"

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
		address := addressInfo.Endpoint.Address
		if _, ok := processedAddresses[address]; ok {
			continue
		}

		processedAddresses[address] = struct{}{}

		wg.Add(1)

		go func(addressInfo host.AddressInfo) {
			defer wg.Done()

			var result responseWithError

			rpcUrl, err := addressInfo.GetUrlRPC()
			if err != nil {
				result.err = err
				respChan <- result

				return
			}

			rpcGateway := rpcgw.NewGateway(c.gateways.cli, rpcUrl)

			metricsUrl, err := addressInfo.GetUrlPrometheus()
			if err != nil {
				result.err = err
				respChan <- result

				return
			}

			prometheusGateway := prometheusgw.NewGateway(c.gateways.cli, metricsUrl)
			geoGateway := geogw.NewGateway(c.gateways.cli, c.selectedConfig.IPLookup.AccessToken)

			createdHost := host.NewHost(table, addressInfo, rpcGateway, geoGateway, prometheusGateway, c.gateways.cli)
			result.response = createdHost

			if c.selectedConfig.IPLookup.AccessToken != "" {
				if err := createdHost.SetIPInfo(); err != nil {
					result.err = err
					respChan <- result

					return
				}
			}

			if err := createdHost.GetMetrics(); err != nil {
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

	var mErr *multierror.Error

	for result := range respChan {
		if result.err != nil {
			mErr = multierror.Append(mErr, result.err)

			continue
		}

		hosts = append(hosts, *result.response)
	}

	if len(hosts) == 0 {
		return nil, mErr.ErrorOrNil()
	}

	return hosts, nil
}
