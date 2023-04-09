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
			hostPortRPC     *address.HostPort
			hostPortMetrics *address.HostPort
		)

		if addressRPC != "" {
			if hostPortRPC, err = address.ParseURL(addressRPC); err != nil {
				return nil, fmt.Errorf("invalid full-node json-rpcgw-address in fullnode.yaml: %s", err)
			}
		}

		if addressMetrics != "" {
			if hostPortMetrics, err = address.ParseURL(addressMetrics); err != nil {
				return nil, fmt.Errorf("invalid full-node rpcgw-address in fullnode.yaml: %s", err)
			}
		}

		addressInfo := host.AddressInfo{HostPort: *hostPortRPC, Ports: make(map[enums.PortType]string)}
		if hostPortRPC.Port != nil {
			addressInfo.Ports[enums.PortTypeRPC] = *hostPortRPC.Port
		}

		if hostPortMetrics.Port != nil {
			addressInfo.Ports[enums.PortTypeMetrics] = *hostPortMetrics.Port
		}

		addresses = append(addresses, addressInfo)
	case enums.TableTypeValidator:
		validatorConfig := c.config.Validator
		addressMetrics := validatorConfig.MetricsAddress

		if addressMetrics == "" {
			return nil, errors.New("validator tables not found in suimon.yaml")
		}

		var hostPortMetrics *address.HostPort

		if hostPortMetrics, err = address.ParseURL(addressMetrics); err != nil {
			return nil, fmt.Errorf("invalid validator rpcgw-address in fullnode.yaml: %s", err)
		}

		addressInfo := host.AddressInfo{HostPort: *hostPortMetrics, Ports: make(map[enums.PortType]string)}

		if hostPortMetrics.Port != nil {
			addressInfo.Ports[enums.PortTypeMetrics] = *hostPortMetrics.Port
		}

		addresses = append(addresses, addressInfo)
	case enums.TableTypePeers:
		peersConfig := c.config.SeedPeers
		if len(peersConfig) == 0 {
			return nil, errors.New("seed-peers tables not found in suimon.yaml")
		}

		for _, peer := range peersConfig {
			hostPort, err := address.ParsePeer(peer)
			if err != nil {
				return nil, fmt.Errorf("invalid peer in suimon.yaml: %s", peer)
			}

			addressInfo := host.AddressInfo{HostPort: *hostPort, Ports: make(map[enums.PortType]string)}
			if hostPort.Port != nil {
				addressInfo.Ports[enums.PortTypePeer] = *hostPort.Port
			}

			addresses = append(addresses, addressInfo)
		}
	case enums.TableTypeRPC:
		rpcConfig := c.config.PublicRPC
		if len(rpcConfig) == 0 {
			return nil, errors.New("public-rpcgw tables not found in suimon.yaml")
		}

		for _, rpc := range rpcConfig {
			hostPort, err := address.ParseURL(rpc)
			if err != nil {
				return nil, fmt.Errorf("invalid rpcgw url in suimon.yaml: %s", rpc)
			}

			addressInfo := host.AddressInfo{HostPort: *hostPort, Ports: make(map[enums.PortType]string)}
			if hostPort.Port != nil {
				addressInfo.Ports[enums.PortTypeRPC] = *hostPort.Port
			}

			addresses = append(addresses, addressInfo)
		}
	}

	return addresses, nil
}
