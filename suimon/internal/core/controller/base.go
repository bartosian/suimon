package controller

import (
	"net/http"
	"time"

	"github.com/ipinfo/go/v2/ipinfo"
	"github.com/ipinfo/go/v2/ipinfo/cache"
	"github.com/ybbus/jsonrpc/v3"

	"github.com/bartosian/sui_helpers/suimon/cmd/config"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/dashboardbuilder"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/host"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/ports"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/tablebuilder"
	"github.com/bartosian/sui_helpers/suimon/internal/pkg/log"
)

const ipInfoCacheExp = 5 * time.Minute

type CheckerController struct {
	suimonConfig config.SuimonConfig

	rpc       []host.Host
	node      []host.Host
	validator []host.Host
	peers     []host.Host

	rpcClient  jsonrpc.RPCClient
	httpClient *http.Client
	ipClient   *ipinfo.Client

	tableConfig tablebuilder.TableConfig

	tableBuilderPeer             ports.Builder
	tableBuilderNode             ports.Builder
	tableBuilderValidator        ports.Builder
	tableBuilderRPC              ports.Builder
	tableBuilderSystem           ports.Builder
	tableBuilderActiveValidators ports.Builder

	DashboardBuilder *dashboardbuilder.DashboardBuilder

	logger log.Logger
}

// NewCheckerController creates a new Checker object for the given Suimon and node configurations and network type.
// This function accepts the following parameters:
// - suimonConfig: a config.SuimonConfig struct containing the Suimon configuration data.
// - nodeConfig: a config.NodeConfig struct containing the node configuration data.
// - networkType: an enums.NetworkType representing the type of network to configure the Checker object for.
// The function returns a pointer to a Checker struct, and an error if there was an issue creating the Checker object.
func NewCheckerController(suimonConfig config.SuimonConfig) (*CheckerController, error) {
	rpcClient := jsonrpc.NewClient(suimonConfig.PublicRPC[0])
	httpClient := &http.Client{Timeout: httpClientTimeout}
	ipClient := ipinfo.NewClient(httpClient, ipinfo.NewCache(cache.NewInMemory().WithExpiration(ipInfoCacheExp)), suimonConfig.IPLookup.AccessToken)

	return &CheckerController{
		rpcClient:    rpcClient,
		httpClient:   httpClient,
		ipClient:     ipClient,
		suimonConfig: suimonConfig,
		logger:       log.NewLogger(),
	}, nil
}

// getHostsByTableType returns a list of Host objects associated with the given table type in the Checker struct.
// This function accepts the following parameter:
// - tableType: an enums.TableType representing the type of table to retrieve hosts for.
// The function returns a slice of Host objects representing the hosts associated with the given table type.
func (checker CheckerController) getHostsByTableType(tableType enums.TableType) []host.Host {
	var hosts []host.Host

	switch tableType {
	case enums.TableTypeNode:
		hosts = checker.node
	case enums.TableTypeValidator:
		hosts = checker.validator
	case enums.TableTypeActiveValidators:
		hosts = checker.rpc[:1]
	case enums.TableTypeSystemState:
		hosts = checker.rpc[:1]
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
func (checker *CheckerController) setHostsByTableType(tableType enums.TableType, hosts []host.Host) {
	switch tableType {
	case enums.TableTypeNode:
		checker.node = hosts
	case enums.TableTypeValidator:
		checker.validator = hosts
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
func (checker *CheckerController) setBuilderTableType(tableType enums.TableType, tableConfig tablebuilder.TableConfig) {
	tableBuilder := tablebuilder.NewTableBuilder(tableConfig)

	switch tableType {
	case enums.TableTypeNode:
		checker.tableBuilderNode = tableBuilder
	case enums.TableTypeValidator:
		checker.tableBuilderValidator = tableBuilder
	case enums.TableTypePeers:
		checker.tableBuilderPeer = tableBuilder
	case enums.TableTypeRPC:
		checker.tableBuilderRPC = tableBuilder
	case enums.TableTypeSystemState:
		checker.tableBuilderSystem = tableBuilder
	case enums.TableTypeActiveValidators:
		checker.tableBuilderActiveValidators = tableBuilder
	}
}

// RenderTables draws tables for each type of table in the Checker struct.
// This function does not accept any parameters.
// This function does not return anything.
func (checker CheckerController) RenderTables() error {
	if checker.suimonConfig.MonitorsConfig.RPCTable.Display && len(checker.rpc) > 0 {
		checker.tableBuilderRPC.Render()
	}

	if checker.suimonConfig.MonitorsConfig.NodeTable.Display && len(checker.node) > 0 {
		checker.tableBuilderNode.Render()
	}

	if checker.suimonConfig.MonitorsConfig.ValidatorTable.Display && len(checker.validator) > 0 {
		checker.tableBuilderValidator.Render()
	}

	if checker.suimonConfig.MonitorsConfig.SystemTable.Display && len(checker.rpc) > 0 {
		checker.tableBuilderSystem.Render()
	}

	if checker.suimonConfig.MonitorsConfig.PeersTable.Display && len(checker.peers) > 0 {
		checker.tableBuilderPeer.Render()
	}

	if checker.suimonConfig.MonitorsConfig.ActiveValidatorsTable.Display && len(checker.rpc) > 0 {
		systemState := checker.rpc[0].Metrics.SystemState

		if len(systemState.ActiveValidators) == 0 {
			return nil
		}

		checker.tableBuilderActiveValidators.Render()
	}

	return nil
}
