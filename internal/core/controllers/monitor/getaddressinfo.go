package monitor

import (
	"errors"
	"fmt"

	"github.com/bartosian/suimon/internal/core/domain/enums"
	"github.com/bartosian/suimon/internal/core/domain/host"
	"github.com/bartosian/suimon/internal/pkg/address"
)

type addressParser func(string) (*address.Endpoint, error)

var parserMap = map[enums.TableType]addressParser{
	enums.TableTypeNode:      address.ParseURL,
	enums.TableTypeValidator: address.ParseURL,
	enums.TableTypeRPC:       address.ParseURL,
}

// getAddressInfoByTableType retrieves the list of addresses for hosts that support the specified table type from the CheckerController's internal state.
// The function returns an error if the specified table type is invalid or if there are no hosts that support the specified table type.
// Returns a slice of AddressInfo structs and an error if the specified table type is invalid or if there are no hosts that support the specified table type.
func (c *Controller) getAddressInfoByTableType(table enums.TableType) ([]host.AddressInfo, error) {
	addressFuncMap := map[enums.TableType]func(parser addressParser) ([]host.AddressInfo, error){
		enums.TableTypeNode:      c.getNodeAddresses,
		enums.TableTypeValidator: c.getValidatorAddresses,
		enums.TableTypeRPC:       c.getRPCAddresses,
	}

	parser, parserExists := parserMap[table]
	if !parserExists {
		return nil, fmt.Errorf("invalid table type: %v", table)
	}

	if addressFunc, existsFunc := addressFuncMap[table]; existsFunc {
		addresses, err := addressFunc(parser)
		if err != nil {
			return nil, err
		}

		return addresses, nil
	}

	return nil, fmt.Errorf("address function not found for table type: %v", table)
}

// getNodeAddresses extracts the JSON-RPC and metrics addresses from the selected config's full nodes and
// returns an array of host.AddressInfo structs that include the endpoints and port numbers.
// The parser argument is a function used to parse the address strings.
// Returns an error if there is an invalid address format or if there is no JSON-RPC or metrics address provided for a full node.
func (c *Controller) getNodeAddresses(parser addressParser) (addresses []host.AddressInfo, err error) {
	nodesConfig := c.selectedConfig.FullNodes
	if len(nodesConfig) == 0 {
		return []host.AddressInfo{}, nil
	}

	for _, node := range nodesConfig {
		addressRPC, addressMetrics := node.JSONRPCAddress, node.MetricsAddress

		if addressRPC == "" && addressMetrics == "" {
			return nil, errors.New("invalid format for full-node in dashboards file: at least one of json-rpc-address or metrics-address is required")
		}

		var addressInfo host.AddressInfo

		if addressRPC != "" {
			endpointRPC, parseErr := parser(addressRPC)
			if parseErr != nil {
				return nil, fmt.Errorf("invalid format for full-node json-rpc-address in config file: %w", parseErr)
			}

			addressInfo = host.AddressInfo{
				Endpoint: *endpointRPC,
				Ports:    map[enums.PortType]string{},
			}

			if endpointRPC.Port != nil {
				addressInfo.Ports[enums.PortTypeRPC] = *endpointRPC.Port
			}
		}

		if addressMetrics != "" {
			endpointMetrics, parseErr := parser(addressMetrics)
			if parseErr != nil {
				return nil, fmt.Errorf("invalid format for full-node metrics-address in config file: %w", parseErr)
			}

			// If addressInfo is still empty, initialize it with endpointMetrics
			if addressInfo.Endpoint.Address == "" {
				addressInfo.Endpoint = *endpointMetrics
				addressInfo.Ports = map[enums.PortType]string{}
			}

			if endpointMetrics.Port != nil {
				addressInfo.Ports[enums.PortTypeMetrics] = *endpointMetrics.Port
			}
		}

		addresses = append(addresses, addressInfo)
	}

	return addresses, nil
}

// getValidatorAddresses returns the list of validator addresses.
// It takes an addressParser as input and returns a list of host.AddressInfo and an error.
// It processes the validator addresses and initializes the hosts.
// If the validatorsConfig is empty, it returns an empty list.
// If the metrics-address is missing for any validator, it returns an error.
// If there is an error in parsing the validator metrics-address, it returns an error.
// The function appends the processed addresses to the list and returns it along with any encountered error.
func (c *Controller) getValidatorAddresses(parser addressParser) (addresses []host.AddressInfo, err error) {
	validatorsConfig := c.selectedConfig.Validators
	if len(validatorsConfig) == 0 {
		return []host.AddressInfo{}, nil
	}

	for _, validator := range validatorsConfig {
		addressMetrics := validator.MetricsAddress

		if addressMetrics == "" {
			return nil, errors.New("invalid format for validator in dashboards file: metrics-address is required")
		}

		endpointMetrics, parseErr := parser(addressMetrics)
		if parseErr != nil {
			return nil, fmt.Errorf("invalid format for validator metrics-address in config file: %w", parseErr)
		}

		addressInfo := host.AddressInfo{Endpoint: *endpointMetrics, Ports: make(map[enums.PortType]string)}

		if endpointMetrics.Port != nil {
			addressInfo.Ports[enums.PortTypeMetrics] = *endpointMetrics.Port
		}

		addresses = append(addresses, addressInfo)
	}

	return addresses, nil
}

// getRPCAddresses returns the list of reference RPC addresses.
// It takes an addressParser as input and returns a list of host.AddressInfo and an error.
// It processes the RPC addresses and initializes the hosts.
// If the referenceRPCConfig is empty, it returns an error.
// If the rpc-address is missing for any reference RPC, it returns an error.
// If there is an error in parsing the reference RPC address, it returns an error.
// The function appends the processed addresses to the list and returns it along with any encountered error.
// This function is part of the Controller struct.
func (c *Controller) getRPCAddresses(parser addressParser) (addresses []host.AddressInfo, err error) {
	rpcConfig := c.selectedConfig.ReferenceRPC
	if len(rpcConfig) == 0 {
		return nil, errors.New("reference-rpc not provided in config file")
	}

	for _, rpc := range rpcConfig {
		endpoint, parseErr := parser(rpc)
		if parseErr != nil {
			return nil, fmt.Errorf("invalid format for reference-rpc in config file: %w", parseErr)
		}

		addressInfo := host.AddressInfo{Endpoint: *endpoint, Ports: make(map[enums.PortType]string)}
		if endpoint.Port != nil {
			addressInfo.Ports[enums.PortTypeRPC] = *endpoint.Port
		}

		addresses = append(addresses, addressInfo)
	}

	return addresses, nil
}
