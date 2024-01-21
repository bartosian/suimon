package tablebuilder

import (
	"errors"

	"github.com/bartosian/suimon/internal/core/domain/enums"
	domainhost "github.com/bartosian/suimon/internal/core/domain/host"
	domainmetrics "github.com/bartosian/suimon/internal/core/domain/metrics"
)

const utcTimeZone = "America/New_York"

// Init initializes the table configuration based on the given table type and host data.
// It processes the host data and calls the appropriate handler function for the specified table type.
func (tb *Builder) Init() error {
	if len(tb.hosts) == 0 && tb.tableType != enums.TableTypeReleases {
		return errors.New("hosts are not initialized")
	}

	handlerMap := map[enums.TableType]func() error{
		enums.TableTypeNode:               func() error { return tb.handleNodeTable(tb.hosts) },
		enums.TableTypeRPC:                func() error { return tb.handleRPCTable(tb.hosts) },
		enums.TableTypeValidator:          func() error { return tb.handleValidatorTable(tb.hosts) },
		enums.TableTypeGasPriceAndSubsidy: func() error { return tb.handleTableWithMetrics(tb.hosts, tb.handleSystemStateTable) },
		enums.TableTypeProtocol:           func() error { return tb.handleTableWithMetrics(tb.hosts, tb.handleProtocolTable) },
		enums.TableTypeValidatorsParams:   func() error { return tb.handleTableWithSystemState(tb.hosts, tb.handleValidatorParamsTable) },
		enums.TableTypeValidatorsAtRisk:   func() error { return tb.handleTableWithSystemState(tb.hosts, tb.handleValidatorsAtRiskTable) },
		enums.TableTypeValidatorReports:   func() error { return tb.handleTableWithSystemState(tb.hosts, tb.handleValidatorReportsTable) },
		enums.TableTypeActiveValidators:   func() error { return tb.handleTableWithMetrics(tb.hosts, tb.handleActiveValidatorsTable) },
		enums.TableTypeReleases:           func() error { return tb.handleReleasesTable(tb.Releases) },
	}

	if handler, ok := handlerMap[tb.tableType]; ok {
		return handler()
	}

	return nil
}

// handleTableWithMetrics is a generic wrapper for handling tables that require metrics.
func (tb *Builder) handleTableWithMetrics(hosts []domainhost.Host, handlerFunc func(*domainmetrics.Metrics) error) error {
	if len(hosts) == 0 {
		return errors.New("no hosts available")
	}
	return handlerFunc(&hosts[0].Metrics)
}

// handleTableWithSystemState is a generic wrapper for handling tables that require system state.
func (tb *Builder) handleTableWithSystemState(hosts []domainhost.Host, handlerFunc func(*domainmetrics.SuiSystemState) error) error {
	if len(hosts) == 0 {
		return errors.New("no hosts available")
	}
	return handlerFunc(&hosts[0].Metrics.SystemState)
}
