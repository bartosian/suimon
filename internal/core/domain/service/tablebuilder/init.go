package tablebuilder

import (
	"errors"
	"fmt"
	"sort"
	"strconv"

	"github.com/bartosian/suimon/internal/core/domain/enums"
	domainhost "github.com/bartosian/suimon/internal/core/domain/host"
	domainmetrics "github.com/bartosian/suimon/internal/core/domain/metrics"
	"github.com/bartosian/suimon/internal/core/domain/service/tablebuilder/tables"
)

// Init initializes the table configuration based on the given table type and host data.
// It processes the host data and calls the appropriate handler function for the specified table type.
func (tb *Builder) Init() error {
	hosts := tb.hosts

	if len(hosts) == 0 {
		return errors.New("hosts are not initialized")
	}

	switch tb.tableType {
	case enums.TableTypeNode:
		tb.handleNodeTable(hosts)
	case enums.TableTypeRPC:
		tb.handleRPCTable(hosts)
	case enums.TableTypeValidator:
		tb.handleValidatorTable(hosts)
	case enums.TableTypeGasPriceAndSubsidy:
		metrics := hosts[0].Metrics

		return tb.handleSystemStateTable(&metrics)
	case enums.TableTypeValidatorsParams:
		systemState := hosts[0].Metrics.SystemState

		return tb.handleValidatorParamsTable(&systemState)
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
		metrics := hosts[0].Metrics

		return tb.handleActiveValidatorsTable(&metrics)
	}

	return nil
}

// handleNodeTable handles the configuration for the Node table.
func (tb *Builder) handleNodeTable(hosts []domainhost.Host) {
	tableConfig := tables.NewDefaultTableConfig(enums.TableTypeNode)

	sort.SliceStable(hosts, func(left, right int) bool {
		if hosts[left].Status != hosts[right].Status {
			return hosts[left].Status > hosts[right].Status
		}

		if hosts[left].Metrics.TotalTransactionsBlocks != hosts[right].Metrics.TotalTransactionsBlocks {
			return hosts[left].Metrics.TotalTransactionsBlocks > hosts[right].Metrics.TotalTransactionsBlocks
		}

		return hosts[left].Metrics.HighestSyncedCheckpoint != hosts[right].Metrics.HighestSyncedCheckpoint
	})

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

	sort.SliceStable(hosts, func(left, right int) bool {
		if hosts[left].Status != hosts[right].Status {
			return hosts[left].Status > hosts[right].Status
		}

		return hosts[left].Metrics.TotalTransactionsBlocks > hosts[right].Metrics.TotalTransactionsBlocks
	})

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

// handleValidatorTable handles the configuration for the Validator table.
func (tb *Builder) handleValidatorTable(hosts []domainhost.Host) {
	tableConfig := tables.NewDefaultTableConfig(enums.TableTypeValidator)

	sort.SliceStable(hosts, func(left, right int) bool {
		if hosts[left].Status != hosts[right].Status {
			return hosts[left].Status > hosts[right].Status
		}

		if hosts[left].Metrics.CurrentRound != hosts[right].Metrics.CurrentRound {
			return hosts[left].Metrics.CurrentRound > hosts[right].Metrics.CurrentRound
		}

		return hosts[left].Metrics.HighestSyncedCheckpoint > hosts[right].Metrics.HighestSyncedCheckpoint
	})

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
func (tb *Builder) handleSystemStateTable(metrics *domainmetrics.Metrics) error {
	tableConfig := tables.NewDefaultTableConfig(enums.TableTypeGasPriceAndSubsidy)

	columnValues, err := tables.GetSystemStateColumnValues(metrics)
	if err != nil {
		return err
	}

	tableConfig.Columns.SetColumnValues(columnValues)

	tableConfig.RowsCount++

	tb.config = tableConfig

	return nil
}

// handleValidatorParamsTable handles the configuration for the Validator Counts table.
func (tb *Builder) handleValidatorParamsTable(systemState *domainmetrics.SuiSystemState) error {
	tableConfig := tables.NewDefaultTableConfig(enums.TableTypeValidatorsParams)

	columnValues, err := tables.GetValidatorParamsColumnValues(systemState)
	if err != nil {
		return err
	}

	tableConfig.Columns.SetColumnValues(columnValues)

	tableConfig.RowsCount++

	tb.config = tableConfig

	return nil
}

// handleValidatorsAtRiskTable handles the configuration for the Validators At Risk table.
// It takes the system state, extracts the necessary data, and updates the table configuration.
func (tb *Builder) handleValidatorsAtRiskTable(systemState *domainmetrics.SuiSystemState) error {
	tableConfig := tables.NewDefaultTableConfig(enums.TableTypeValidatorsAtRisk)

	validatorsAtRisk := systemState.ValidatorsAtRiskParsed

	const base = 10 // for strconv.ParseInt

	sort.SliceStable(validatorsAtRisk, func(left, right int) bool {
		epochsAtRiskLeft, err := strconv.ParseInt(validatorsAtRisk[left].EpochsAtRisk, base, 64)
		if err != nil {
			return true
		}

		epochsAtRiskRight, err := strconv.ParseInt(validatorsAtRisk[right].EpochsAtRisk, base, 64)
		if err != nil {
			return true
		}

		if epochsAtRiskLeft != epochsAtRiskRight {
			return epochsAtRiskLeft > epochsAtRiskRight
		}

		return validatorsAtRisk[left].Name < validatorsAtRisk[right].Name
	})

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
func (tb *Builder) handleValidatorReportsTable(systemState *domainmetrics.SuiSystemState) error {
	tableConfig := tables.NewDefaultTableConfig(enums.TableTypeValidatorReports)

	validatorReports := systemState.ValidatorReportsParsed

	for _, report := range validatorReports {
		for j, reporter := range report.Reporters {
			reportedName := report.Name
			slashingPct := fmt.Sprintf("%.2f", report.SlashingPercentage)

			if j > 0 {
				reportedName = " "
				slashingPct = " "
			}

			columnValues := tables.GetValidatorReportColumnValues(reportedName, slashingPct, reporter)

			tableConfig.Columns.SetColumnValues(columnValues)

			tableConfig.RowsCount++
		}
	}

	tb.config = tableConfig

	return nil
}

// handleActiveValidatorsTable handles the configuration for the Active Validators table.
// It takes the system state, extracts the necessary data, and updates the table configuration.
func (tb *Builder) handleActiveValidatorsTable(metrics *domainmetrics.Metrics) error {
	tableConfig := tables.NewDefaultTableConfig(enums.TableTypeActiveValidators)

	activeValidators := metrics.SystemState.ActiveValidators
	validatorsApy := metrics.ValidatorsApyParsed

	const base = 10 // for strconv.ParseInt

	sort.SliceStable(activeValidators, func(left, right int) bool {
		votingPowerLeft, err := strconv.ParseInt(activeValidators[left].VotingPower, base, 64)
		if err != nil {
			return true
		}

		votingPowerRight, err := strconv.ParseInt(activeValidators[right].VotingPower, base, 64)
		if err != nil {
			return false // right is considered greater
		}

		nextEpochStakeLeft, err := strconv.ParseInt(activeValidators[left].NextEpochStake, base, 64)
		if err != nil {
			return true
		}

		nextEpochStakeRight, err := strconv.ParseInt(activeValidators[right].NextEpochStake, base, 64)
		if err != nil {
			return false // right is considered greater
		}

		if votingPowerLeft != votingPowerRight {
			return votingPowerLeft > votingPowerRight
		}

		if nextEpochStakeLeft != nextEpochStakeRight {
			return nextEpochStakeLeft > nextEpochStakeRight
		}

		return activeValidators[left].Name < activeValidators[right].Name
	})

	for idx, validator := range activeValidators {
		validatorApy, ok := validatorsApy[validator.SuiAddress]
		if !ok {
			return fmt.Errorf("failed to loookup validator apy by address: %s", validator.SuiAddress)
		}

		validator.APY = strconv.FormatFloat(validatorApy*100, 'f', 3, 64)

		columnValues, err := tables.GetActiveValidatorColumnValues(idx, validator)
		if err != nil {
			return err
		}

		tableConfig.Columns.SetColumnValues(columnValues)

		tableConfig.RowsCount++
	}

	tb.config = tableConfig

	return nil
}
