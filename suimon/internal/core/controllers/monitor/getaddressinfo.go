package monitor

import (
	"errors"
	"fmt"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/host"
	"github.com/bartosian/sui_helpers/suimon/internal/pkg/address"
)

// getAddressInfoByTableType retrieves the list of addresses for hosts that support the specified table type from the CheckerController's internal state.
// The function returns an error if the specified table type is invalid or if there are no hosts that support the specified table type.
// Returns a slice of AddressInfo structs and an error if the specified table type is invalid or if there are no hosts that support the specified table type.
func (c *Controller) getAddressInfoByTableType(table enums.TableType) (addresses []host.AddressInfo, err error) {
	switch table {
	case enums.TableTypeNode:
		nodeConfig := c.config.FullNode
		addressRPC, addressMetrics := nodeConfig.JSONRPCAddress, nodeConfig.MetricsAddress

		if addressRPC == "" && addressMetrics == "" {
			return nil, errors.New("full-node tables not found in suimon.yaml")
		}

		var (
			endpointRPC     *address.Endpoint
			endpointMetrics *address.Endpoint
		)

		if addressRPC != "" {
			if endpointRPC, err = address.ParseURL(addressRPC); err != nil {
				return nil, fmt.Errorf("invalid full-node json-rpc-address in fullnode.yaml: %s", err)
			}
		}

		if addressMetrics != "" {
			if endpointMetrics, err = address.ParseURL(addressMetrics); err != nil {
				return nil, fmt.Errorf("invalid full-node rpc-address in fullnode.yaml: %s", err)
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
	case enums.TableTypeValidator:
		validatorConfig := c.config.Validator
		addressMetrics := validatorConfig.MetricsAddress

		if addressMetrics == "" {
			return nil, errors.New("validator tables not found in suimon.yaml")
		}

		var endpointMetrics *address.Endpoint

		if endpointMetrics, err = address.ParseURL(addressMetrics); err != nil {
			return nil, fmt.Errorf("invalid validator rpc-address in fullnode.yaml: %s", err)
		}

		addressInfo := host.AddressInfo{Endpoint: *endpointMetrics, Ports: make(map[enums.PortType]string)}

		if endpointMetrics.Port != nil {
			addressInfo.Ports[enums.PortTypeMetrics] = *endpointMetrics.Port
		}

		addresses = append(addresses, addressInfo)
	case enums.TableTypePeers:
		peersConfig := c.config.SeedPeers
		if len(peersConfig) == 0 {
			return nil, errors.New("seed-peers tables not found in suimon.yaml")
		}

		for _, peer := range peersConfig {
			endpoint, err := address.ParsePeer(peer)
			if err != nil {
				return nil, fmt.Errorf("invalid peer in suimon.yaml: %s", peer)
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
			return nil, errors.New("public-rpc tables not found in suimon.yaml")
		}

		for _, rpc := range rpcConfig {
			endpoint, err := address.ParseURL(rpc)
			if err != nil {
				return nil, fmt.Errorf("invalid rpc url in suimon.yaml: %s", rpc)
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
