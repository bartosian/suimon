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
		return c.getNodeAddresses(parser)
	case enums.TableTypeValidator:
		return c.getValidatorAddresses(parser)
	case enums.TableTypePeers:
		return c.getPeerAddresses(parser)
	case enums.TableTypeRPC:
		return c.getRPCAddresses(parser)
	}

	return addresses, nil
}

// getNodeAddresses returns the list of addresses of full nodes.
func (c *Controller) getNodeAddresses(parser addressParser) (addresses []host.AddressInfo, err error) {
	nodesConfig := c.selectedConfig.FullNodes
	if len(nodesConfig) == 0 {
		return
	}

	for _, node := range nodesConfig {
		addressRPC, addressMetrics := node.JSONRPCAddress, node.MetricsAddress

		if addressRPC == "" && addressMetrics == "" {
			return nil, errors.New("invalid format for full-node in dashboards file: at least one of json-rpc-address or metrics-address is required")
		}

		var (
			endpointRPC     *address.Endpoint
			endpointMetrics *address.Endpoint
		)

		if addressRPC != "" {
			if endpointRPC, err = parser(addressRPC); err != nil {
				return nil, fmt.Errorf("invalid format for full-node json-rpc-address in dashboards file: %w", err)
			}
		}

		if addressMetrics != "" {
			if endpointMetrics, err = parser(addressMetrics); err != nil {
				return nil, fmt.Errorf("invalid format for full-node metrics-address in dashboards file: %w", err)
			}
		}

		// Check if both endpoints are nil and return an error if so.
		if endpointRPC == nil && endpointMetrics == nil {
			return nil, errors.New("invalid format for full-node in dashboards file: at least one of json-rpc-address or metrics-address is required")
		}

		addressInfo := host.AddressInfo{Endpoint: *endpointRPC, Ports: make(map[enums.PortType]string)}

		if endpointRPC != nil && endpointRPC.Port != nil {
			addressInfo.Ports[enums.PortTypeRPC] = *endpointRPC.Port
		}

		if endpointMetrics != nil && endpointMetrics.Port != nil {
			addressInfo.Ports[enums.PortTypeMetrics] = *endpointMetrics.Port
		}

		addresses = append(addresses, addressInfo)
	}

	return addresses, nil
}

// getValidatorAddresses returns the list of addresses of validators.
func (c *Controller) getValidatorAddresses(parser addressParser) (addresses []host.AddressInfo, err error) {
	validatorsConfig := c.selectedConfig.Validators
	if len(validatorsConfig) == 0 {
		return
	}

	for _, validator := range validatorsConfig {
		addressMetrics := validator.MetricsAddress

		if addressMetrics == "" {
			return nil, errors.New("invalid format for validator in dashboards file: metrics-address is required")
		}

		endpointMetrics, err := parser(addressMetrics)
		if err != nil {
			return nil, fmt.Errorf("invalid format for validator metrics-address in dashboards file: %w", err)
		}

		addressInfo := host.AddressInfo{Endpoint: *endpointMetrics, Ports: make(map[enums.PortType]string)}

		if endpointMetrics.Port != nil {
			addressInfo.Ports[enums.PortTypeMetrics] = *endpointMetrics.Port
		}

		addresses = append(addresses, addressInfo)
	}

	return addresses, nil
}

// getPeerAddresses returns the list of seed peer addresses.
func (c *Controller) getPeerAddresses(parser addressParser) (addresses []host.AddressInfo, err error) {
	peersConfig := c.selectedConfig.SeedPeers
	if len(peersConfig) == 0 {
		return
	}

	for _, peer := range peersConfig {
		endpoint, err := parser(peer)
		if err != nil {
			return nil, fmt.Errorf("invalid format for seed-peer in dashboards file: %w", err)
		}

		addressInfo := host.AddressInfo{Endpoint: *endpoint, Ports: make(map[enums.PortType]string)}
		if endpoint.Port != nil {
			addressInfo.Ports[enums.PortTypePeer] = *endpoint.Port
		}

		addresses = append(addresses, addressInfo)
	}

	return addresses, nil
}

// getRPCAddresses returns the list of public RPC addresses.
func (c *Controller) getRPCAddresses(parser addressParser) (addresses []host.AddressInfo, err error) {
	rpcConfig := c.selectedConfig.PublicRPC
	if len(rpcConfig) == 0 {
		return nil, errors.New("public-rpc not provided in dashboards file")
	}

	for _, rpc := range rpcConfig {
		endpoint, err := parser(rpc)
		if err != nil {
			return nil, fmt.Errorf("invalid format for public-rpc in dashboards file: %w", err)
		}

		addressInfo := host.AddressInfo{Endpoint: *endpoint, Ports: make(map[enums.PortType]string)}
		if endpoint.Port != nil {
			addressInfo.Ports[enums.PortTypeRPC] = *endpoint.Port
		}

		addresses = append(addresses, addressInfo)
	}

	return addresses, nil
}
