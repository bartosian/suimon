package controller

import (
	"errors"
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/host"
	"github.com/bartosian/sui_helpers/suimon/internal/pkg/address"
	"github.com/bartosian/sui_helpers/suimon/internal/pkg/progress"
)

const (
	httpClientTimeout = 2 * time.Second
)

// getAddressInfoByTableType retrieves the list of addresses for hosts that support the specified table type from the CheckerController's internal state.
// The function returns an error if the specified table type is invalid or if there are no hosts that support the specified table type.
// Returns a slice of AddressInfo structs and an error if the specified table type is invalid or if there are no hosts that support the specified table type.
func (checker CheckerController) getAddressInfoByTableType(tableType enums.TableType) (addresses []host.AddressInfo, err error) {
	switch tableType {
	case enums.TableTypeNode:
		nodeConfig := checker.suimonConfig.FullNode
		addressRPC, addressMetrics := nodeConfig.JSONRPCAddress, nodeConfig.MetricsAddress

		if addressRPC == "" && addressMetrics == "" {
			return nil, errors.New("full-node config not found in suimon.yaml")
		}

		var (
			hostPortRPC     *address.HostPort
			hostPortMetrics *address.HostPort
		)

		if addressRPC != "" {
			if hostPortRPC, err = address.ParseURL(addressRPC); err != nil {
				return nil, fmt.Errorf("invalid full-node json-rpc-address in fullnode.yaml: %s", err)
			}
		}

		if addressMetrics != "" {
			if hostPortMetrics, err = address.ParseURL(addressMetrics); err != nil {
				return nil, fmt.Errorf("invalid full-node metrics-address in fullnode.yaml: %s", err)
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
		validatorConfig := checker.suimonConfig.Validator
		addressMetrics := validatorConfig.MetricsAddress

		if addressMetrics == "" {
			return nil, errors.New("validator config not found in suimon.yaml")
		}

		var hostPortMetrics *address.HostPort

		if hostPortMetrics, err = address.ParseURL(addressMetrics); err != nil {
			return nil, fmt.Errorf("invalid validator metrics-address in fullnode.yaml: %s", err)
		}

		addressInfo := host.AddressInfo{HostPort: *hostPortMetrics, Ports: make(map[enums.PortType]string)}

		if hostPortMetrics.Port != nil {
			addressInfo.Ports[enums.PortTypeMetrics] = *hostPortMetrics.Port
		}

		addresses = append(addresses, addressInfo)
	case enums.TableTypePeers:
		peersConfig := checker.suimonConfig.SeedPeers
		if len(peersConfig) == 0 {
			return nil, errors.New("seed-peers config not found in suimon.yaml")
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
		rpcConfig := checker.suimonConfig.PublicRPC
		if len(rpcConfig) == 0 {
			return nil, errors.New("public-rpc config not found in suimon.yaml")
		}

		for _, rpc := range rpcConfig {
			hostPort, err := address.ParseURL(rpc)
			if err != nil {
				return nil, fmt.Errorf("invalid rpc url in suimon.yaml: %s", rpc)
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

// ParseData retrieves the latest data from all active hosts and updates the CheckerController's internal state with the new data.
// The function parses the data for each table type and sets the corresponding metrics and dashboard options accordingly.
// Returns an error if the data cannot be retrieved from any of the active hosts or if there is an issue parsing the data for any table type.
func (checker *CheckerController) ParseData() error {
	monitorsConfig := checker.suimonConfig.MonitorsConfig

	tableMap := map[enums.TableType]bool{
		enums.TableTypeRPC:       monitorsConfig.RPCTable.Display,
		enums.TableTypeNode:      monitorsConfig.NodeTable.Display,
		enums.TableTypeValidator: monitorsConfig.ValidatorTable.Display,
		enums.TableTypePeers:     monitorsConfig.PeersTable.Display,
	}

	errChan := make(chan error, len(tableMap))
	defer close(errChan)

	var wg sync.WaitGroup

	for tableType, isEnabled := range tableMap {
		if !isEnabled {
			continue
		}

		wg.Add(1)

		go func(tt enums.TableType) {
			defer wg.Done()

			if err := checker.getHostsData(tt, progress.ColorBlue); err != nil {
				errChan <- err
			}

			checker.setHostsHealth(tt)
		}(tableType)
	}

	wg.Wait()

	select {
	case err := <-errChan:
		return err
	default:
		if len(errChan) == len(tableMap) {
			return errors.New("all tables disabled in suimon.yaml")
		}
		return nil
	}
}

// getHostsData retrieves the latest data for the specified table type from all active hosts and updates the CheckerController's internal state with the new data.
// The function retrieves data for each host in parallel and displays a progress bar indicating the progress of the data retrieval process.
// Returns an error if the data cannot be retrieved from any of the active hosts or if there is an issue updating the CheckerController's internal state.
func (checker *CheckerController) getHostsData(tableType enums.TableType, progressColor progress.Color) error {
	monitorsConfig := checker.suimonConfig.MonitorsConfig

	progressChan := progress.NewProgressBar("PARSING DATA FOR "+string(tableType), progressColor)
	defer func() { progressChan <- struct{}{} }()

	var (
		addresses []host.AddressInfo
		hosts     []host.Host
		err       error
	)

	parseHosts := func() error {
		if addresses, err = checker.getAddressInfoByTableType(tableType); err != nil {
			return err
		}

		if hosts, err = checker.createHosts(tableType, addresses); err != nil {
			return err
		}

		checker.setHostsByTableType(tableType, hosts)

		return nil
	}

	if err := parseHosts(); err != nil {
		return err
	}

	if tableType == enums.TableTypeRPC {
		checker.sortHosts(enums.TableTypeRPC)

		if monitorsConfig.ActiveValidatorsTable.Display || monitorsConfig.SystemTable.Display {
			checker.hosts.rpc[0].GetMetricRPC(enums.RPCMethodGetSuiSystemState, enums.MetricTypeSuiSystemState)
		}
	}

	return nil
}

// sortHosts sorts the active hosts for the specified table type based on their corresponding metric values.
// The function retrieves the relevant metric for each host, sorts the hosts by their metric values, and updates the CheckerController's internal state accordingly.
// Returns an error if the specified table type is invalid or if there is an issue sorting the hosts based on their corresponding metric values.
func (checker *CheckerController) sortHosts(tableType enums.TableType) {
	hosts := checker.getHostsByTableType(tableType)

	if len(hosts) > 1 {
		sort.Slice(hosts, func(left, right int) bool {
			return hosts[left].Metrics.TotalTransactionsBlocks > hosts[right].Metrics.TotalTransactionsBlocks
		})

		sort.SliceStable(hosts, func(left, right int) bool {
			return hosts[left].Metrics.LatestCheckpoint > hosts[right].Metrics.LatestCheckpoint
		})
	}

	checker.setHostsByTableType(tableType, hosts)
}

// setHostsHealth retrieves the latest health information for all active hosts and updates the CheckerController's internal state with the new information.
// The function retrieves health information for each host in parallel and sets the corresponding health status in the internal state.
// Returns an error if the health information cannot be retrieved from any of the active hosts or if there is an issue updating the CheckerController's internal state.
func (checker *CheckerController) setHostsHealth(tableType enums.TableType) {
	hosts := checker.getHostsByTableType(tableType)
	primaryRPC := checker.hosts.rpc[0]

	for idx := range hosts {
		metrics := hosts[idx].Metrics

		checkpointExecBacklog := metrics.HighestKnownCheckpoint - metrics.LastExecutedCheckpoint
		checkpointSyncBacklog := metrics.HighestKnownCheckpoint - metrics.HighestSyncedCheckpoint

		hosts[idx].SetPctProgress(enums.MetricTypeTxSyncPercentage, primaryRPC)
		hosts[idx].SetPctProgress(enums.MetricTypeCheckSyncPercentage, primaryRPC)
		hosts[idx].Metrics.SetValue(enums.MetricTypeCheckpointExecBacklog, checkpointExecBacklog)
		hosts[idx].Metrics.SetValue(enums.MetricTypeCheckpointSyncBacklog, checkpointSyncBacklog)

		hosts[idx].SetStatus(primaryRPC)
	}
}
