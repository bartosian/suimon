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

const utcTimeZone = "America/New_York"

// Init initializes the table configuration based on the given table type and host data.
// It processes the host data and calls the appropriate handler function for the specified table type.
func (tb *Builder) Init() error {
	if len(tb.hosts) == 0 {
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

// handleNodeTable handles the configuration for the Node table.
func (tb *Builder) handleNodeTable(hosts []domainhost.Host) error {
	tableConfig := tables.NewDefaultTableConfig(enums.TableTypeNode)

	sort.SliceStable(hosts, func(i, j int) bool {
		left, right := hosts[i], hosts[j]
		if left.Status != right.Status {
			return left.Status > right.Status
		}
		if left.Metrics.TotalTransactionsBlocks != right.Metrics.TotalTransactionsBlocks {
			return left.Metrics.TotalTransactionsBlocks > right.Metrics.TotalTransactionsBlocks
		}
		return left.Metrics.HighestSyncedCheckpoint > right.Metrics.HighestSyncedCheckpoint
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

	return nil
}

// handleRPCTable handles the configuration for the RPC table.
func (tb *Builder) handleRPCTable(hosts []domainhost.Host) error {
	tableConfig := tables.NewDefaultTableConfig(enums.TableTypeRPC)

	sort.SliceStable(hosts, func(i, j int) bool {
		left, right := hosts[i], hosts[j]
		if left.Status != right.Status {
			return left.Status > right.Status
		}
		return left.Metrics.TotalTransactionsBlocks > right.Metrics.TotalTransactionsBlocks
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

	return nil
}

// handleValidatorTable handles the configuration for the Validator table.
func (tb *Builder) handleValidatorTable(hosts []domainhost.Host) error {
	tableConfig := tables.NewDefaultTableConfig(enums.TableTypeValidator)

	sort.SliceStable(hosts, func(i, j int) bool {
		left, right := hosts[i], hosts[j]
		if left.Status != right.Status {
			return left.Status > right.Status
		}
		if left.Metrics.CurrentRound != right.Metrics.CurrentRound {
			return left.Metrics.CurrentRound > right.Metrics.CurrentRound
		}
		return left.Metrics.HighestSyncedCheckpoint > right.Metrics.HighestSyncedCheckpoint
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

	return nil
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
	const base = 10

	// Optimized sorting logic
	sort.SliceStable(validatorsAtRisk, func(i, j int) bool {
		leftEpochs, leftErr := strconv.ParseInt(validatorsAtRisk[i].EpochsAtRisk, base, 64)
		rightEpochs, rightErr := strconv.ParseInt(validatorsAtRisk[j].EpochsAtRisk, base, 64)

		if leftErr != nil || rightErr != nil {
			return leftErr == nil
		}

		if leftEpochs != rightEpochs {
			return leftEpochs > rightEpochs
		}

		return validatorsAtRisk[i].Name < validatorsAtRisk[j].Name
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

	const base = 10

	sort.SliceStable(activeValidators, func(i, j int) bool {
		leftVotingPower, leftErr := strconv.ParseInt(activeValidators[i].VotingPower, base, 64)
		rightVotingPower, rightErr := strconv.ParseInt(activeValidators[j].VotingPower, base, 64)

		if leftErr != nil {
			return false
		}
		if rightErr != nil {
			return true
		}

		leftNextEpochStake, leftStakeErr := strconv.ParseInt(activeValidators[i].NextEpochStake, base, 64)
		rightNextEpochStake, rightStakeErr := strconv.ParseInt(activeValidators[j].NextEpochStake, base, 64)

		if leftStakeErr != nil {
			return false
		}
		if rightStakeErr != nil {
			return true
		}

		if leftVotingPower != rightVotingPower {
			return leftVotingPower > rightVotingPower
		}

		if leftNextEpochStake != rightNextEpochStake {
			return leftNextEpochStake > rightNextEpochStake
		}

		return activeValidators[i].Name < activeValidators[j].Name
	})

	for idx, validator := range activeValidators {
		validatorApy, ok := validatorsApy[validator.SuiAddress]
		if !ok {
			return fmt.Errorf("failed to lookup validator APY by address: %s", validator.SuiAddress)
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
