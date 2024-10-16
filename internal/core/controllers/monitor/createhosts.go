package monitor

import (
	"sync"

	"github.com/hashicorp/go-multierror"

	"github.com/bartosian/suimon/internal/core/domain/enums"
	"github.com/bartosian/suimon/internal/core/domain/host"
	"github.com/bartosian/suimon/internal/core/gateways/geogw"
	"github.com/bartosian/suimon/internal/core/gateways/prometheusgw"
	"github.com/bartosian/suimon/internal/core/gateways/rpcgw"
)

type responseWithError struct {
	response *host.Host
	err      error
}

// createHosts creates hosts based on the provided table type and addresses.
// It initializes the hosts, processes the addresses, and sets up the necessary gateways for each host.
// It returns the created hosts and any error encountered during the process.
func (c *Controller) createHosts(table enums.TableType, addresses []host.AddressInfo) ([]host.Host, error) {
	hosts := make([]host.Host, 0, len(addresses))
	processedAddresses := make(map[string]struct{})

	respChan := make(chan responseWithError, len(addresses))

	var wg sync.WaitGroup

	// Helper function to send error responses
	sendErrorResponse := func(result responseWithError, err error) {
		result.err = err
		respChan <- result
	}

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

			rpcURL, err := addressInfo.GetURLRPC()
			if err != nil {
				sendErrorResponse(result, err)
				return
			}

			metricsURL, err := addressInfo.GetURLPrometheus()
			if err != nil {
				sendErrorResponse(result, err)
				return
			}

			rpcGateway := rpcgw.NewGateway(c.gateways.cli, rpcURL)
			prometheusGateway := prometheusgw.NewGateway(c.gateways.cli, metricsURL)
			geoGateway := geogw.NewGateway(c.gateways.cli, c.selectedConfig.IPLookup.AccessToken)

			createdHost := host.NewHost(table, addressInfo, rpcGateway, geoGateway, prometheusGateway, c.gateways.cli)
			result.response = createdHost

			// If IP lookup token exists, set IP info
			if c.selectedConfig.IPLookup.AccessToken != "" {
				if createErr := createdHost.SetIPInfo(); createErr != nil {
					sendErrorResponse(result, createErr)
					return
				}
			}

			// Fetch metrics
			if getMetricsErr := createdHost.GetMetrics(); getMetricsErr != nil {
				sendErrorResponse(result, err)
				return
			}

			respChan <- result
		}(addressInfo)
	}

	// Close the response channel once all goroutines finish
	go func() {
		wg.Wait()
		close(respChan)
	}()

	// Collect results and errors
	var mErr *multierror.Error

	for result := range respChan {
		if result.err != nil {
			mErr = multierror.Append(mErr, result.err)
		} else {
			hosts = append(hosts, *result.response)
		}
	}

	// Return errors if no hosts are created
	if len(hosts) == 0 {
		return nil, mErr.ErrorOrNil()
	}

	return hosts, nil
}
