package tablebuilder

import (
	"errors"

	"github.com/bartosian/suimon/internal/core/domain/enums"
	domainhost "github.com/bartosian/suimon/internal/core/domain/host"
)

const utcTimeZone = "America/New_York"

// Init initializes the table configuration based on the given table type and host data.
// It processes the host data and calls the appropriate handler function for the specified table type.
func (tb *Builder) Init() error {
	if len(tb.hosts) == 0 && tb.tableType != enums.TableTypeReleases {
		return errors.New("hosts are not initialized")
	}

	handlerMap := map[enums.TableType]func([]domainhost.Host) error{
		enums.TableTypeNode:               tb.handleNodeTable,
		enums.TableTypeRPC:                tb.handleRPCTable,
		enums.TableTypeValidator:          tb.handleValidatorTable,
		enums.TableTypeGasPriceAndSubsidy: tb.handleSystemStateTableWrapper,
		enums.TableTypeValidatorsParams:   tb.handleValidatorParamsTableWrapper,
		enums.TableTypeValidatorsAtRisk:   tb.handleValidatorsAtRiskTableWrapper,
		enums.TableTypeValidatorReports:   tb.handleValidatorReportsTableWrapper,
		enums.TableTypeActiveValidators:   tb.handleActiveValidatorsTableWrapper,
		enums.TableTypeReleases:           tb.handleReleasesTableWrapper,
	}

	if handler, ok := handlerMap[tb.tableType]; ok {
		return handler(tb.hosts)
	}

	return nil
}

func (tb *Builder) handleSystemStateTableWrapper(hosts []domainhost.Host) error {
	metrics := hosts[0].Metrics
	return tb.handleSystemStateTable(&metrics)
}

func (tb *Builder) handleValidatorParamsTableWrapper(hosts []domainhost.Host) error {
	systemState := hosts[0].Metrics.SystemState
	return tb.handleValidatorParamsTable(&systemState)
}

func (tb *Builder) handleValidatorsAtRiskTableWrapper(hosts []domainhost.Host) error {
	systemState := hosts[0].Metrics.SystemState
	return tb.handleValidatorsAtRiskTable(&systemState)
}

func (tb *Builder) handleValidatorReportsTableWrapper(hosts []domainhost.Host) error {
	systemState := hosts[0].Metrics.SystemState
	return tb.handleValidatorReportsTable(&systemState)
}

func (tb *Builder) handleActiveValidatorsTableWrapper(hosts []domainhost.Host) error {
	metrics := hosts[0].Metrics
	return tb.handleActiveValidatorsTable(&metrics)
}

func (tb *Builder) handleReleasesTableWrapper(hosts []domainhost.Host) error {
	return tb.handleReleasesTable(tb.Releases)
}
