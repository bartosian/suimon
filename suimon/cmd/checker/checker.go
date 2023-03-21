package checker

import (
	"net/http"
	"sync"
	"time"

	"github.com/ipinfo/go/v2/ipinfo"
	"github.com/ipinfo/go/v2/ipinfo/cache"
	"github.com/mum4k/termdash"
	"github.com/ybbus/jsonrpc/v3"

	"github.com/bartosian/sui_helpers/suimon/cmd/checker/config"
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/dashboardbuilder"
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/tablebuilder"
	"github.com/bartosian/sui_helpers/suimon/internal/pkg/log"
)

const ipInfoCacheExp = 5 * time.Minute

type (
	Checker struct {
		suimonConfig config.SuimonConfig
		nodeConfig   config.NodeConfig

		rpc   []Host
		node  []Host
		peers []Host

		rpcClient  jsonrpc.RPCClient
		httpClient *http.Client
		ipClient   *ipinfo.Client

		tableBuilderPeer       *tablebuilder.TableBuilder
		tableBuilderNode       *tablebuilder.TableBuilder
		tableBuilderRPC        *tablebuilder.TableBuilder
		tableBuilderSystem     *tablebuilder.TableBuilder
		tableBuilderValidators *tablebuilder.TableBuilder
		tableConfig            tablebuilder.TableConfig

		DashboardBuilder *dashboardbuilder.DashboardBuilder

		logger log.Logger
	}
)

// NewChecker creates a new Checker object for the given Suimon and node configurations and network type.
// This function accepts the following parameters:
// - suimonConfig: a config.SuimonConfig struct containing the Suimon configuration data.
// - nodeConfig: a config.NodeConfig struct containing the node configuration data.
// - networkType: an enums.NetworkType representing the type of network to configure the Checker object for.
// The function returns a pointer to a Checker struct, and an error if there was an issue creating the Checker object.
func NewChecker(suimonConfig config.SuimonConfig, nodeConfig config.NodeConfig, networkType enums.NetworkType) (*Checker, error) {
	suimonConfig.SetNetworkConfig(networkType)

	rpcClient := jsonrpc.NewClient(suimonConfig.Network.NetworkType.ToRPC())
	httpClient := &http.Client{Timeout: httpClientTimeout}
	ipClient := ipinfo.NewClient(httpClient, ipinfo.NewCache(cache.NewInMemory().WithExpiration(ipInfoCacheExp)), suimonConfig.IPLookup.AccessToken)

	return &Checker{
		rpcClient:    rpcClient,
		httpClient:   httpClient,
		ipClient:     ipClient,
		suimonConfig: suimonConfig,
		nodeConfig:   nodeConfig,
		logger:       log.NewLogger(),
	}, nil
}

// getHostsByTableType returns a list of Host objects associated with the given table type in the Checker struct.
// This function accepts the following parameter:
// - tableType: an enums.TableType representing the type of table to retrieve hosts for.
// The function returns a slice of Host objects representing the hosts associated with the given table type.
func (checker *Checker) getHostsByTableType(tableType enums.TableType) []Host {
	var hosts []Host

	switch tableType {
	case enums.TableTypeNode:
		hosts = checker.node
	case enums.TableTypeValidators:
		hosts = checker.node
	case enums.TableTypeSystemState:
		hosts = checker.node
	case enums.TableTypePeers:
		hosts = checker.peers
	case enums.TableTypeRPC:
		hosts = checker.rpc
	}

	return hosts
}

// setHostsByTableType sets the list of Host objects associated with the given table type in the Checker struct.
// This function accepts the following parameters:
// - tableType: an enums.TableType representing the type of table to set hosts for.
// - hosts: a slice of Host objects representing the hosts to associate with the given table type.
// This function does not return anything.
func (checker *Checker) setHostsByTableType(tableType enums.TableType, hosts []Host) {
	switch tableType {
	case enums.TableTypeNode:
		checker.node = hosts
	case enums.TableTypePeers:
		checker.peers = hosts
	case enums.TableTypeRPC:
		checker.rpc = hosts
	}
}

// setBuilderTableType sets the table configuration for the given table type in the Checker struct using the provided TableConfig.
// This function accepts the following parameters:
// - tableType: an enums.TableType representing the type of table to set the configuration for.
// - tableConfig: a tablebuilder.TableConfig struct containing the configuration data for the table.
// This function does not return anything.
func (checker *Checker) setBuilderTableType(tableType enums.TableType, tableConfig tablebuilder.TableConfig) {
	tableBuilder := tablebuilder.NewTableBuilder(tableConfig)

	switch tableType {
	case enums.TableTypeNode:
		checker.tableBuilderNode = tableBuilder
	case enums.TableTypePeers:
		checker.tableBuilderPeer = tableBuilder
	case enums.TableTypeRPC:
		checker.tableBuilderRPC = tableBuilder
	case enums.TableTypeSystemState:
		checker.tableBuilderSystem = tableBuilder
	case enums.TableTypeValidators:
		checker.tableBuilderValidators = tableBuilder
	}
}

// DrawTables draws tables for each type of table in the Checker struct.
// This function does not accept any parameters.
// This function does not return anything.
func (checker *Checker) DrawTables() {
	if checker.suimonConfig.MonitorsConfig.RPCTable.Display && len(checker.rpc) > 0 {
		checker.tableBuilderRPC.Build()
	}

	if checker.suimonConfig.MonitorsConfig.NodeTable.Display && len(checker.node) > 0 {
		checker.tableBuilderNode.Build()
	}

	if checker.suimonConfig.MonitorsConfig.PeersTable.Display && len(checker.peers) > 0 {
		checker.tableBuilderPeer.Build()
	}

	if checker.suimonConfig.MonitorsConfig.SystemTable.Display && len(checker.node) > 0 {
		checker.tableBuilderSystem.Build()
	}

	if checker.suimonConfig.MonitorsConfig.ValidatorsTable.Display && len(checker.node) > 0 {
		systemState := checker.node[0].Metrics.SystemState

		if len(systemState.ActiveValidators) == 0 {
			return
		}

		checker.tableBuilderValidators.Build()
	}
}

// DrawDashboards draws dashboards for each type of table in the Checker struct.
// This function does not accept any parameters.
// This function does not return anything.
func (checker *Checker) DrawDashboards() {
	var (
		monitorsConfig    = checker.suimonConfig.MonitorsConfig
		processLaunchType = checker.suimonConfig.ProcessLaunchType
		dashboardBuilder  = checker.DashboardBuilder
		dashCells         = dashboardBuilder.Cells
		ticker            = time.NewTicker(watchHostsTimeout)
		logsCH            = make(chan string)
		wg                sync.WaitGroup
	)

	defer ticker.Stop()

	var stream = func() {
		defer wg.Done()

		var err error

		if processLaunchType.ServiceName != "" {
			if err = checker.logger.StreamFromService(processLaunchType.ServiceName, logsCH); err == nil {
				return
			}
		}

		if processLaunchType.DockerImageName != "" {
			if err = checker.logger.StreamFromContainer(processLaunchType.DockerImageName, logsCH); err == nil {
				return
			}
		}

		if processLaunchType.ScreenSessionName != "" {
			if err = checker.logger.StreamFromScreen(processLaunchType.ScreenSessionName, logsCH); err == nil {
				return
			}
		}
	}

	var draw = func(hosts []Host) {
		defer wg.Done()

		doneCH := make(chan struct{}, len(hosts))

		for {
			select {
			case <-ticker.C:
				for _, host := range hosts {
					go func(host Host) {
						for idx, dashCell := range dashCells {
							cellName := enums.CellName(idx)

							metric := checker.getMetricForDashboardCell(cellName)
							options := checker.getOptionsForDashboardCell(cellName)

							dashCell.Write(metric, options)
						}

						doneCH <- struct{}{}
					}(host)
				}

				for i := 0; i < len(hosts); i++ {
					<-doneCH
				}
			case log := <-logsCH:
				dashCell := dashCells[enums.CellNameNodeLogs]
				options := checker.getOptionsForDashboardCell(enums.CellNameNodeLogs)

				dashCell.Write(log+"\n", options)
			case <-dashboardBuilder.Ctx.Done():
				close(doneCH)
				close(logsCH)

				return
			}
		}
	}

	wg.Add(1)

	go func() {
		defer wg.Done()

		checker.WatchHosts()
	}()

	if monitorsConfig.NodeTable.Display && len(checker.node) > 0 {
		wg.Add(2)

		go stream()
		go draw(checker.node)
	}

	if err := termdash.Run(dashboardBuilder.Ctx, dashboardBuilder.Terminal, dashboardBuilder.Dashboard, termdash.KeyboardSubscriber(dashboardBuilder.Quitter)); err != nil {
		panic(err)
	}

	wg.Wait()
}
