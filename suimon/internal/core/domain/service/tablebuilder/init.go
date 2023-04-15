package tablebuilder

import (
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	domainhost "github.com/bartosian/sui_helpers/suimon/internal/core/domain/host"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/metrics"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/service/tablebuilder/tables"
)

// Init initializes the table configuration based on the given table type and host data.
// It processes the host data and calls the appropriate handler function for the specified table type.
func (tb *Builder) Init() error {
	hosts := tb.hosts

	switch tb.tableType {
	case enums.TableTypeNode:
		tb.handleNodeTable(hosts)
	case enums.TableTypeRPC:
		tb.handleRPCTable(hosts)
	case enums.TableTypePeers:
		tb.handlePeersTable(hosts)
	case enums.TableTypeValidator:
		tb.handleValidatorTable(hosts)
	case enums.TableTypeSystemState:
		systemState := hosts[0].Metrics.SystemState

		return tb.handleSystemStateTable(&systemState)
	case enums.TableTypeValidatorsCounts:
		systemState := hosts[0].Metrics.SystemState

		tb.handleValidatorCountsTable(&systemState)
	case enums.TableTypeValidatorsAtRisk:
		systemState := hosts[0].Metrics.SystemState

		if err := tb.handleValidatorsAtRiskTable(&systemState); err != nil {
			return err
		}
	case enums.TableTypeValidatorReports:
		systemState := hosts[0].Metrics.SystemState

		if err := tb.handleValidatorReportsTable(&systemState); err != nil {
			return err
		}
	case enums.TableTypeActiveValidators:
		systemState := hosts[0].Metrics.SystemState

		tb.handleActiveValidatorsTable(&systemState)
	}

	return nil
}

// handleNodeTable handles the configuration for the Node table.
func (tb *Builder) handleNodeTable(hosts []domainhost.Host) {
	tableConfig := tables.NewDefaultTableConfig(enums.TableTypeNode)

	for idx, host := range hosts {
		if !host.Metrics.Updated {
			continue
		}

		columnValues := tables.GetNodeColumnValues(idx, host)

		tableConfig.Columns.SetColumnValues(columnValues)

		tableConfig.RowsCount++
	}

	tb.config = tableConfig
}

// handleRPCTable handles the configuration for the RPC table.
func (tb *Builder) handleRPCTable(hosts []domainhost.Host) {
	tableConfig := tables.NewDefaultTableConfig(enums.TableTypeRPC)

	for idx, host := range hosts {
		if !host.Metrics.Updated {
			continue
		}

		columnValues := tables.GetRPCColumnValues(idx, host)

		tableConfig.Columns.SetColumnValues(columnValues)

		tableConfig.RowsCount++
	}

	tb.config = tableConfig
}

// handlePeersTable handles the configuration for the Peers table.
func (tb *Builder) handlePeersTable(hosts []domainhost.Host) {
	tableConfig := tables.NewDefaultTableConfig(enums.TableTypePeers)

	for idx, host := range hosts {
		if !host.Metrics.Updated {
			continue
		}

		columnValues := tables.GetNodeColumnValues(idx, host)

		tableConfig.Columns.SetColumnValues(columnValues)

		tableConfig.RowsCount++
	}

	tb.config = tableConfig
}

// handleValidatorTable handles the configuration for the Validator table.
func (tb *Builder) handleValidatorTable(hosts []domainhost.Host) {
	tableConfig := tables.NewDefaultTableConfig(enums.TableTypeValidator)

	for idx, host := range hosts {
		if !host.Metrics.Updated {
			continue
		}

		columnValues := tables.GetValidatorColumnValues(idx, host)

		tableConfig.Columns.SetColumnValues(columnValues)

		tableConfig.RowsCount++
	}

	tb.config = tableConfig
}

// handleSystemStateTable handles the configuration for the System State table.
func (tb *Builder) handleSystemStateTable(systemState *metrics.SuiSystemState) error {
	tableConfig := tables.NewDefaultTableConfig(enums.TableTypeSystemState)

	columnValues, err := tables.GetSystemStateColumnValues(systemState)
	if err != nil {
		return err
	}

	tableConfig.Columns.SetColumnValues(columnValues)

	tableConfig.RowsCount++

	tb.config = tableConfig

	return nil
}

// handleValidatorCountsTable handles the configuration for the Validator Counts table.
func (tb *Builder) handleValidatorCountsTable(systemState *metrics.SuiSystemState) {
	tableConfig := tables.NewDefaultTableConfig(enums.TableTypeValidatorsCounts)

	columnValues := tables.GetValidatorCountsColumnValues(systemState)

	tableConfig.Columns.SetColumnValues(columnValues)

	tableConfig.RowsCount++

	tb.config = tableConfig
}

// handleValidatorsAtRiskTable handles the configuration for the Validators At Risk table.
// It takes the system state, extracts the necessary data, and updates the table configuration.
func (tb *Builder) handleValidatorsAtRiskTable(systemState *metrics.SuiSystemState) error {
	tableConfig := tables.NewDefaultTableConfig(enums.TableTypeValidatorsAtRisk)

	validatorsAtRisk := systemState.ValidatorsAtRiskParsed

	for idx, validator := range validatorsAtRisk {
		columnValues := tables.GetValidatorAtRiskColumnValues(idx, validator)

		tableConfig.Columns.SetColumnValues(columnValues)

		tableConfig.RowsCount++
	}

	tb.config = tableConfig

	return nil
}

// handleValidatorReportsTable handles the configuration for the Validator Reports table.
// It takes the system state, extracts the necessary data, and updates the table configuration.
func (tb *Builder) handleValidatorReportsTable(systemState *metrics.SuiSystemState) error {
	tableConfig := tables.NewDefaultTableConfig(enums.TableTypeValidatorReports)

	validatorReports := systemState.ValidatorReportsParsed

	for idx, report := range validatorReports {
		columnValues := tables.GetValidatorReportColumnValues(idx, report)

		tableConfig.Columns.SetColumnValues(columnValues)

		tableConfig.RowsCount++
	}

	tb.config = tableConfig

	return nil
}

// handleActiveValidatorsTable handles the configuration for the Active Validators table.
// It takes the system state, extracts the necessary data, and updates the table configuration.
func (tb *Builder) handleActiveValidatorsTable(systemState *metrics.SuiSystemState) {
	tableConfig := tables.NewDefaultTableConfig(enums.TableTypeActiveValidators)

	activeValidators := systemState.ActiveValidators

	for idx, validator := range activeValidators {
		columnValues := tables.GetActiveValidatorColumnValues(idx, validator)

		tableConfig.Columns.SetColumnValues(columnValues)

		tableConfig.RowsCount++
	}

	tb.config = tableConfig
}
