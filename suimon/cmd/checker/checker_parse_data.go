package checker

import (
	"errors"
	"fmt"
	"sort"
	"time"

	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/pkg/address"
	"github.com/bartosian/sui_helpers/suimon/internal/pkg/progress"
)

const (
	httpClientTimeout = 3 * time.Second
)

// getAddressInfoByTableType returns a slice of "AddressInfo" values for a specific table of the "Checker"
// struct instance passed as a pointer receiver, based on the provided "TableType".
// Parameters:
// - tableType: an enums.TableType representing the type of table to retrieve the address information for.
// Returns:
// - a slice of "AddressInfo" values for the specified table.
// - an error, if any occurred during retrieval of the address information.
func (checker *Checker) getAddressInfoByTableType(tableType enums.TableType) ([]AddressInfo, error) {
	var (
		addresses []AddressInfo
		err       error
	)

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

		addressInfo := AddressInfo{HostPort: *hostPortRPC, Ports: make(map[enums.PortType]string)}
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

		addressInfo := AddressInfo{HostPort: *hostPortMetrics, Ports: make(map[enums.PortType]string)}

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

			addressInfo := AddressInfo{HostPort: *hostPort, Ports: make(map[enums.PortType]string)}
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

			addressInfo := AddressInfo{HostPort: *hostPort, Ports: make(map[enums.PortType]string)}
			if hostPort.Port != nil {
				addressInfo.Ports[enums.PortTypeRPC] = *hostPort.Port
			}

			addresses = append(addresses, addressInfo)
		}
	}

	return addresses, nil
}

// Init initializes the "Checker" struct instance passed as a pointer receiver.
// Parameters: None.
// Returns: an error, if any occurred during initialization.
func (checker *Checker) Init() error {
	var errChan = make(chan error, 4)

	defer close(errChan)

	if monitorsConfig := checker.suimonConfig.MonitorsConfig; !monitorsConfig.PeersTable.Display && !monitorsConfig.RPCTable.Display &&
		!monitorsConfig.NodeTable.Display && !monitorsConfig.ActiveValidatorsTable.Display && !monitorsConfig.ValidatorTable.Display {
		return errors.New("all tables disabled in suimon.yaml")
	}

	// parse data for the RPC servers
	go checker.getHostsData(enums.TableTypeRPC, progress.ColorBlue, errChan)

	// parse data for the full node
	go checker.getHostsData(enums.TableTypeNode, progress.ColorRed, errChan)

	// parse data for the validator
	go checker.getHostsData(enums.TableTypeValidator, progress.ColorRed, errChan)

	// parse data for the peers servers
	go checker.getHostsData(enums.TableTypePeers, progress.ColorGreen, errChan)

	for i := 0; i < cap(errChan); i++ {
		if err := <-errChan; err != nil {
			return err
		}
	}

	checker.setHostsHealth(enums.TableTypeRPC)
	checker.setHostsHealth(enums.TableTypeNode)
	checker.setHostsHealth(enums.TableTypeValidator)
	checker.setHostsHealth(enums.TableTypePeers)

	return nil
}

func (checker *Checker) getHostsData(tableType enums.TableType, progressColor progress.Color, errChan chan<- error) {
	var (
		progressChan   = progress.NewProgressBar("PARSING DATA FOR "+string(tableType), progressColor)
		monitorsConfig = checker.suimonConfig.MonitorsConfig
		addresses      []AddressInfo
		hosts          []Host
		err            error
	)

	defer func() {
		progressChan <- struct{}{}
		errChan <- err
	}()

	var parseHosts = func() {
		if addresses, err = checker.getAddressInfoByTableType(tableType); err != nil {
			return
		}

		if hosts, err = checker.createHosts(tableType, addresses); err != nil {
			return
		}

		checker.setHostsByTableType(tableType, hosts)
	}

	switch tableType {
	case enums.TableTypeRPC:
		parseHosts()

		checker.sortHosts(enums.TableTypeRPC)

		if monitorsConfig.ActiveValidatorsTable.Display || monitorsConfig.SystemTable.Display {
			checker.rpc[0].GetLatestSuiSystemState()
		}
	case enums.TableTypePeers:
		if monitorsConfig.PeersTable.Display {
			parseHosts()
		}
	case enums.TableTypeValidator:
		if monitorsConfig.ValidatorTable.Display {
			parseHosts()
		}
	case enums.TableTypeNode:
		if monitorsConfig.NodeTable.Display {
			parseHosts()
		}
	}
}

func (checker *Checker) sortHosts(tableType enums.TableType) {
	hosts := checker.getHostsByTableType(tableType)

	if len(hosts) > 1 {
		sort.Slice(hosts, func(left, right int) bool {
			return hosts[left].Metrics.TotalTransactions > hosts[right].Metrics.TotalTransactions
		})

		sort.SliceStable(hosts, func(left, right int) bool {
			return hosts[left].Metrics.LatestCheckpoint > hosts[right].Metrics.LatestCheckpoint
		})
	}

	checker.setHostsByTableType(tableType, hosts)
}

func (checker *Checker) setHostsHealth(tableType enums.TableType) {
	hosts := checker.getHostsByTableType(tableType)
	primaryRPC := checker.rpc[0]

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
