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

type (
	Clients struct {
		rpcClient  jsonrpc.RPCClient
		httpClient *http.Client
		ipClient   *ipinfo.Client
	}

	Builders struct {
		peer             ports.Builder
		node             ports.Builder
		validator        ports.Builder
		rpc              ports.Builder
		system           ports.Builder
		activeValidators ports.Builder
	}

	Hosts struct {
		rpc       []host.Host
		node      []host.Host
		validator []host.Host
		peers     []host.Host
	}

	CheckerController struct {
		suimonConfig config.SuimonConfig

		logger log.Logger

		hosts            Hosts
		clients          Clients
		tableBuilders    Builders
		DashboardBuilder *dashboardbuilder.DashboardBuilder
	}
)

// NewCheckerController creates a new CheckerController object based on the specified SuimonConfig object.
// The function initializes the CheckerController's internal state and creates a new Host object for each host in the SuimonConfig.
// Returns a pointer to the new CheckerController object and an error value if the creation process fails for any reason.
func NewCheckerController(suimonConfig config.SuimonConfig) (*CheckerController, error) {
	rpcClient := jsonrpc.NewClient(suimonConfig.PublicRPC[0])
	httpClient := &http.Client{Timeout: httpClientTimeout}
	ipClient := ipinfo.NewClient(httpClient, ipinfo.NewCache(cache.NewInMemory().WithExpiration(ipInfoCacheExp)), suimonConfig.IPLookup.AccessToken)

	return &CheckerController{
		clients: Clients{
			rpcClient:  rpcClient,
			httpClient: httpClient,
			ipClient:   ipClient,
		},

		suimonConfig: suimonConfig,
		logger:       log.NewLogger(),
	}, nil
}

// getHostsByTableType returns a list of Host objects that correspond to the specified table type.
// The function searches through the CheckerController's internal state to find all hosts that have the specified table type, and returns a list of those hosts.
// Returns a slice of Host objects that correspond to the specified table type.
func (checker CheckerController) getHostsByTableType(tableType enums.TableType) []host.Host {
	switch tableType {
	case enums.TableTypeNode:
		return checker.hosts.node
	case enums.TableTypeValidator:
		return checker.hosts.validator
	case enums.TableTypeActiveValidators:
		return checker.hosts.rpc[:1]
	case enums.TableTypeSystemState:
		return checker.hosts.rpc[:1]
	case enums.TableTypePeers:
		return checker.hosts.peers
	case enums.TableTypeRPC:
		return checker.hosts.rpc
	}

	return nil
}

// setHostsByTableType sets the list of Host objects that correspond to the specified table type.
// The function updates the CheckerController's internal state to include the specified hosts for the specified table type.
// Returns nothing.
func (checker *CheckerController) setHostsByTableType(tableType enums.TableType, hosts []host.Host) {
	switch tableType {
	case enums.TableTypeNode:
		checker.hosts.node = hosts
	case enums.TableTypeValidator:
		checker.hosts.validator = hosts
	case enums.TableTypePeers:
		checker.hosts.peers = hosts
	case enums.TableTypeRPC:
		checker.hosts.rpc = hosts
	default:
		checker.logger.Error("Unknown table type:", tableType)
	}
}

// setBuilderTableType sets the TableConfig object for the specified table type in the CheckerController's internal state.
// The function updates the CheckerController's internal state to include the specified TableConfig object for the specified table type.
// Returns nothing.
func (checker *CheckerController) setBuilderTableType(tableType enums.TableType, tableConfig tablebuilder.TableConfig) {
	tableBuilder := tablebuilder.NewTableBuilder(tableConfig)

	switch tableType {
	case enums.TableTypeNode:
		checker.tableBuilders.node = tableBuilder
	case enums.TableTypeValidator:
		checker.tableBuilders.validator = tableBuilder
	case enums.TableTypePeers:
		checker.tableBuilders.peer = tableBuilder
	case enums.TableTypeRPC:
		checker.tableBuilders.rpc = tableBuilder
	case enums.TableTypeSystemState:
		checker.tableBuilders.system = tableBuilder
	case enums.TableTypeActiveValidators:
		checker.tableBuilders.activeValidators = tableBuilder
	default:
		checker.logger.Error("Unknown table type:", tableType)
	}
}

// RenderTables draws tables for each type of table in the Checker struct.
// This function does not accept any parameters.
// This function does not return anything.
func (checker *CheckerController) RenderTables() error {
	if checker.suimonConfig.MonitorsConfig.RPCTable.Display && len(checker.hosts.rpc) > 0 {
		if err := checker.tableBuilders.rpc.Render(); err != nil {
			return err
		}
	}

	if checker.suimonConfig.MonitorsConfig.NodeTable.Display && len(checker.hosts.node) > 0 {
		if err := checker.tableBuilders.node.Render(); err != nil {
			return err
		}
	}

	if checker.suimonConfig.MonitorsConfig.ValidatorTable.Display && len(checker.hosts.validator) > 0 {
		if err := checker.tableBuilders.validator.Render(); err != nil {
			return err
		}
	}

	if checker.suimonConfig.MonitorsConfig.SystemTable.Display && len(checker.hosts.rpc) > 0 {
		if err := checker.tableBuilders.system.Render(); err != nil {
			return nil
		}
	}

	if checker.suimonConfig.MonitorsConfig.PeersTable.Display && len(checker.hosts.peers) > 0 {
		if err := checker.tableBuilders.peer.Render(); err != nil {
			return err
		}
	}

	if checker.suimonConfig.MonitorsConfig.ActiveValidatorsTable.Display && len(checker.hosts.rpc) > 0 {
		systemState := checker.hosts.rpc[0].Metrics.SystemState

		if len(systemState.ActiveValidators) == 0 {
			return nil
		}

		if err := checker.tableBuilders.activeValidators.Render(); err != nil {
			return err
		}
	}

	return nil
}
