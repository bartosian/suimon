package monitor

import (
	"errors"
	"fmt"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/host"
	"github.com/bartosian/sui_helpers/suimon/internal/pkg/address"
)

type addressParser func(string) (*address.Endpoint, error)

var parserMap = map[enums.TableType]addressParser{
	enums.TableTypeNode:      address.ParseURL,
	enums.TableTypeValidator: address.ParseURL,
	enums.TableTypePeers:     address.ParsePeer,
	enums.TableTypeRPC:       address.ParseURL,
}

// getAddressInfoByTableType retrieves the list of addresses for hosts that support the specified table type from the CheckerController's internal state.
// The function returns an error if the specified table type is invalid or if there are no hosts that support the specified table type.
// Returns a slice of AddressInfo structs and an error if the specified table type is invalid or if there are no hosts that support the specified table type.
func (c *Controller) getAddressInfoByTableType(table enums.TableType) (addresses []host.AddressInfo, err error) {
	parser, ok := parserMap[table]
	if !ok {
		return nil, fmt.Errorf("invalid table type: %v", table)
	}

	switch table {
	case enums.TableTypeNode:
		nodesConfig := c.config.FullNodes
		if len(nodesConfig) == 0 {
			return nil, errors.New("full-nodes not provided in config file")
		}

		for _, node := range nodesConfig {
			addressRPC, addressMetrics := node.JSONRPCAddress, node.MetricsAddress

			if addressRPC == "" && addressMetrics == "" {
				return nil, errors.New("invalid format for full-nodes in config file")
			}

			var (
				endpointRPC     *address.Endpoint
				endpointMetrics *address.Endpoint
			)

			if addressRPC != "" {
				if endpointRPC, err = parser(addressRPC); err != nil {
					return nil, fmt.Errorf("invalid full-node json-rpc-address in config file: %s", err)
				}
			}

			if addressMetrics != "" {
				if endpointMetrics, err = parser(addressMetrics); err != nil {
					return nil, fmt.Errorf("invalid full-node rpc-address in config file: %s", err)
				}
			}

			addressInfo := host.AddressInfo{Endpoint: *endpointRPC, Ports: make(map[enums.PortType]string)}
			if endpointRPC.Port != nil {
				addressInfo.Ports[enums.PortTypeRPC] = *endpointRPC.Port
			}

			if endpointMetrics.Port != nil {
				addressInfo.Ports[enums.PortTypeMetrics] = *endpointMetrics.Port
			}

			addresses = append(addresses, addressInfo)
		}
	case enums.TableTypeValidator:
		validatorsConfig := c.config.Validators
		if len(validatorsConfig) == 0 {
			return nil, errors.New("validators not provided in config file")
		}

		for _, validator := range validatorsConfig {
			addressMetrics := validator.MetricsAddress

			if addressMetrics == "" {
				return nil, errors.New("invalid format for validators in config file")
			}

			var endpointMetrics *address.Endpoint

			if endpointMetrics, err = parser(addressMetrics); err != nil {
				return nil, fmt.Errorf("invalid validator rpc-address in config file: %s", err)
			}

			addressInfo := host.AddressInfo{Endpoint: *endpointMetrics, Ports: make(map[enums.PortType]string)}

			if endpointMetrics.Port != nil {
				addressInfo.Ports[enums.PortTypeMetrics] = *endpointMetrics.Port
			}

			addresses = append(addresses, addressInfo)
		}
	case enums.TableTypePeers:
		peersConfig := c.config.SeedPeers
		if len(peersConfig) == 0 {
			return nil, errors.New("seed-peers not provided in config file")
		}

		for _, peer := range peersConfig {
			endpoint, err := parser(peer)
			if err != nil {
				return nil, fmt.Errorf("invalid peer in config file: %s", peer)
			}

			addressInfo := host.AddressInfo{Endpoint: *endpoint, Ports: make(map[enums.PortType]string)}
			if endpoint.Port != nil {
				addressInfo.Ports[enums.PortTypePeer] = *endpoint.Port
			}

			addresses = append(addresses, addressInfo)
		}
	case enums.TableTypeRPC:
		rpcConfig := c.config.PublicRPC
		if len(rpcConfig) == 0 {
			return nil, errors.New("public-rpc not found in config file")
		}

		for _, rpc := range rpcConfig {
			endpoint, err := parser(rpc)
			if err != nil {
				return nil, fmt.Errorf("invalid rpc url in config file: %s", rpc)
			}

			addressInfo := host.AddressInfo{Endpoint: *endpoint, Ports: make(map[enums.PortType]string)}
			if endpoint.Port != nil {
				addressInfo.Ports[enums.PortTypeRPC] = *endpoint.Port
			}

			addresses = append(addresses, addressInfo)
		}
	}

	return addresses, nil
}
