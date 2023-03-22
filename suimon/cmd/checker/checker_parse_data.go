package checker

import (
	"errors"
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
	var addresses []AddressInfo

	switch tableType {
	case enums.TableTypeNode:
		addressRPC, addressMetrics := checker.nodeConfig.JSONRPCAddress, checker.nodeConfig.MetricsAddress

		var (
			hostPortRPC     *address.HostPort
			hostPortMetrics *address.HostPort
		)

		if addressRPC != "" {
			hostPort, err := address.ParseIpPort(addressRPC)
			if err != nil {
				return nil, errors.New("invalid json-rpc-address in fullnode.yaml")
			}

			hostPortRPC = hostPort
		}

		if addressMetrics != "" {
			hostPort, err := address.ParseIpPort(addressMetrics)
			if err != nil {
				return nil, errors.New("invalid metrics-address in fullnode.yaml")
			}

			hostPortMetrics = hostPort
		}

		addressInfo := AddressInfo{HostPort: *hostPortRPC, Ports: make(map[enums.PortType]string)}
		if hostPortRPC.Port != nil {
			addressInfo.Ports[enums.PortTypeRPC] = *hostPortRPC.Port
		}

		if hostPortMetrics.Port != nil {
			addressInfo.Ports[enums.PortTypeMetrics] = *hostPortMetrics.Port
		}

		addresses = append(addresses, addressInfo)
	case enums.TableTypePeers:
		peers := checker.nodeConfig.P2PConfig.SeedPeers
		if len(peers) == 0 {
			return nil, errors.New("peers not found in fullnode.yaml")
		}

		for _, peer := range peers {
			hostPort, err := address.ParsePeer(peer.Address)
			if err != nil {
				return nil, errors.New("invalid peer in fullnode.yaml")
			}

			addressInfo := AddressInfo{HostPort: *hostPort, Ports: make(map[enums.PortType]string)}
			if hostPort.Port != nil {
				addressInfo.Ports[enums.PortTypePeer] = *hostPort.Port
			}

			addresses = append(addresses, addressInfo)
		}
	case enums.TableTypeRPC:
		rpc := checker.suimonConfig.GetRPCByNetwork()
		if len(rpc) == 0 {
			return nil, errors.New("rpc not found in suimon.yaml")
		}

		for _, url := range rpc {
			hostPort, err := address.ParseURL(url)
			if err != nil {
				return nil, errors.New("invalid rpc url in suimon.yaml")
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
	var errChan = make(chan error, 3)

	defer close(errChan)

	// parse data for the RPC servers
	go checker.getHostsData(enums.TableTypeRPC, progress.ColorBlue, errChan)

	// parse data for the user servers
	go checker.getHostsData(enums.TableTypeNode, progress.ColorRed, errChan)

	// parse data for the peers servers
	go checker.getHostsData(enums.TableTypePeers, progress.ColorGreen, errChan)

	for i := 0; i < cap(errChan); i++ {
		if err := <-errChan; err != nil {
			return err
		}
	}

	checker.sortHosts(enums.TableTypeRPC)

	checker.setHostsHealth(enums.TableTypeRPC)
	checker.setHostsHealth(enums.TableTypeNode)
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

		if hosts, err = checker.createHosts(addresses); err != nil {
			return
		}

		checker.setHostsByTableType(tableType, hosts)
	}

	switch tableType {
	case enums.TableTypeRPC:
		if monitorsConfig.RPCTable.Display != true {
			return
		}

		parseHosts()
	case enums.TableTypePeers:
		if monitorsConfig.NodeTable.Display != true {
			return
		}

		parseHosts()
	case enums.TableTypeNode:
		if monitorsConfig.NodeTable.Display != true {
			return
		}

		parseHosts()

		fallthrough
	case enums.TableTypeSystemState, enums.TableTypeValidators:
		if monitorsConfig.ValidatorsTable.Display != true {
			return
		}

		checker.node[0].GetLatestSuiSystemState()
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
