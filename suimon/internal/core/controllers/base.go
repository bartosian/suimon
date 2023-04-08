package controllers

import (
	"fmt"
	"github.com/bartosian/sui_helpers/suimon/internal/core/ports"
	"net/http"
	"time"

	"github.com/ipinfo/go/v2/ipinfo"
	"github.com/ipinfo/go/v2/ipinfo/cache"
	"github.com/ybbus/jsonrpc/v3"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/config"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/host"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/service/dashboardbuilder"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/service/tablebuilder"
	"github.com/bartosian/sui_helpers/suimon/internal/pkg/log"
)

const ipInfoCacheExp = 5 * time.Minute

type (
	Clients struct {
		rpcClient  jsonrpc.RPCClient
		httpClient *http.Client
		ipClient   *ipinfo.Client
	}

	Hosts struct {
		rpc       []host.Host
		node      []host.Host
		validator []host.Host
		peers     []host.Host
	}

	CheckerController struct {
		suimonConfig config.Config

		logger log.Logger

		hosts            Hosts
		clients          Clients
		tableBuilders    tablebuilder.Builders
		DashboardBuilder dashboardbuilder.Builder
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
	case enums.TableTypeSystemState,
		enums.TableTypeValidatorsCounts,
		enums.TableTypeValidatorsAtRisk,
		enums.TableTypeValidatorReports:
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
	case enums.TableTypeValidatorsCounts:
		checker.tableBuilders.validatorCounts = tableBuilder
	case enums.TableTypeValidatorsAtRisk:
		checker.tableBuilders.atRiskValidators = tableBuilder
	case enums.TableTypeValidatorReports:
		checker.tableBuilders.validatorReports = tableBuilder
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
	monitorsConfig := checker.suimonConfig.MonitorsConfig

	rpcProvided := len(checker.hosts.rpc) > 0
	nodeProvided := len(checker.hosts.node) > 0
	peersProvided := len(checker.hosts.peers) > 0
	validatorProvided := len(checker.hosts.validator) > 0

	tableTypeToBuilder := map[enums.TableType]struct {
		builder ports.Builder
		enabled bool
	}{
		enums.TableTypeRPC: {
			builder: checker.tableBuilders.rpc,
			enabled: monitorsConfig.RPCTable.Display && rpcProvided,
		},
		enums.TableTypeNode: {
			builder: checker.tableBuilders.node,
			enabled: monitorsConfig.NodeTable.Display && nodeProvided,
		},
		enums.TableTypeValidator: {
			builder: checker.tableBuilders.validator,
			enabled: monitorsConfig.ValidatorTable.Display && validatorProvided,
		},
		enums.TableTypePeers: {
			builder: checker.tableBuilders.peer,
			enabled: monitorsConfig.PeersTable.Display && peersProvided,
		},
		enums.TableTypeSystemState: {
			builder: checker.tableBuilders.system,
			enabled: monitorsConfig.SystemStateTable.Display && rpcProvided,
		},
		enums.TableTypeValidatorsCounts: {
			builder: checker.tableBuilders.validatorCounts,
			enabled: monitorsConfig.ValidatorsCountsTable.Display && rpcProvided,
		},
		enums.TableTypeValidatorsAtRisk: {
			builder: checker.tableBuilders.atRiskValidators,
			enabled: monitorsConfig.ValidatorsAtRiskTable.Display && rpcProvided,
		},
		enums.TableTypeValidatorReports: {
			builder: checker.tableBuilders.validatorReports,
			enabled: monitorsConfig.ValidatorReportsTable.Display && rpcProvided,
		},
		enums.TableTypeActiveValidators: {
			builder: checker.tableBuilders.activeValidators,
			enabled: monitorsConfig.ActiveValidatorsTable.Display && rpcProvided,
		},
	}

	for tableType, builderConfig := range tableTypeToBuilder {
		if !builderConfig.enabled {
			continue
		}

		if err := builderConfig.builder.Render(); err != nil {
			return fmt.Errorf("error rendering table %s: %w", tableType, err)
		}
	}

	return nil
}
