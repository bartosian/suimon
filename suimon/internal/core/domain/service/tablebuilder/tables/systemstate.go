package tables

import (
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/metrics"
	"github.com/bartosian/sui_helpers/suimon/internal/pkg/utility"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/service/tablebuilder"
)

var (
	SortConfigSystem = tablebuilder.SortConfig{
		{Name: enums.ColumnNameSystemEpoch.ToString(), Mode: table.Dsc},
	}
	SortConfigValidatorsAtRisk = tablebuilder.SortConfig{
		{Name: enums.ColumnNameSystemAtRiskValidatorNumberOfEpochs.ToString(), Mode: table.Dsc},
	}
	ColumnsConfigSystem = tablebuilder.ColumnsConfig{
		enums.ColumnNameIndex:                                       tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemEpoch:                                 tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemEpochStartTimestamp:                   tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemEpochDuration:                         tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemTotalStake:                            tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemStorageFundTotalObjectStorageRebates:  tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemStorageFundNonRefundableBalance:       tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemReferenceGasPrice:                     tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemStakeSubsidyStartEpoch:                tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemMaxValidatorCount:                     tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemMinValidatorJoiningStake:              tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemValidatorLowStakeThreshold:            tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemValidatorVeryLowStakeThreshold:        tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemValidatorLowStakeGracePeriod:          tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemStakeSubsidyBalance:                   tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemStakeSubsidyDistributionCounter:       tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemStakeSubsidyCurrentDistributionAmount: tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemStakeSubsidyPeriodLength:              tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemStakeSubsidyDecreaseRate:              tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemActiveValidatorCount:                  tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemPendingActiveValidatorCount:           tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemPendingRemovalsCount:                  tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemValidatorCandidateCount:               tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemAtRiskValidatorCount:                  tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemAtRiskValidatorName:                   tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemAtRiskValidatorAddress:                tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemAtRiskValidatorNumberOfEpochs:         tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemValidatorReportedName:                 tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemValidatorReportedAddress:              tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
	}
	RowsConfigSystemState = tablebuilder.RowsConfig{
		0: {
			enums.ColumnNameIndex,
			enums.ColumnNameSystemEpoch,
			enums.ColumnNameSystemEpochStartTimestamp,
			enums.ColumnNameSystemEpochDuration,
			enums.ColumnNameSystemTotalStake,
			enums.ColumnNameSystemStorageFundTotalObjectStorageRebates,
			enums.ColumnNameSystemStorageFundNonRefundableBalance,
			enums.ColumnNameSystemReferenceGasPrice,
			enums.ColumnNameSystemStakeSubsidyStartEpoch,
			enums.ColumnNameSystemStakeSubsidyBalance,
			enums.ColumnNameSystemStakeSubsidyDistributionCounter,
		},
		1: {
			enums.ColumnNameSystemStakeSubsidyCurrentDistributionAmount,
			enums.ColumnNameSystemStakeSubsidyPeriodLength,
			enums.ColumnNameSystemStakeSubsidyDecreaseRate,
		},
	}
	RowsConfigValidatorCounts = tablebuilder.RowsConfig{
		0: {
			enums.ColumnNameSystemMaxValidatorCount,
			enums.ColumnNameSystemActiveValidatorCount,
			enums.ColumnNameSystemPendingActiveValidatorCount,
			enums.ColumnNameSystemValidatorCandidateCount,
			enums.ColumnNameSystemPendingRemovalsCount,
			enums.ColumnNameSystemAtRiskValidatorCount,
			enums.ColumnNameSystemMinValidatorJoiningStake,
			enums.ColumnNameSystemValidatorLowStakeThreshold,
			enums.ColumnNameSystemValidatorVeryLowStakeThreshold,
			enums.ColumnNameSystemValidatorLowStakeGracePeriod,
		},
	}
	RowsConfigValidatorsAtRisk = tablebuilder.RowsConfig{
		0: {
			enums.ColumnNameIndex,
			enums.ColumnNameSystemAtRiskValidatorName,
			enums.ColumnNameSystemAtRiskValidatorAddress,
			enums.ColumnNameSystemAtRiskValidatorNumberOfEpochs,
		},
	}
	RowsConfigValidatorReports = tablebuilder.RowsConfig{
		0: {
			enums.ColumnNameIndex,
			enums.ColumnNameSystemValidatorReportedName,
			enums.ColumnNameSystemValidatorReportedAddress,
			enums.ColumnNameSystemValidatorReporterName,
			enums.ColumnNameSystemValidatorReporterAddress,
		},
	}
)

// GetSystemStateColumnValues returns a map of SystemColumnName values to corresponding values for the system state.
// The function retrieves information about the system state from the host's internal state and formats it into a map of SystemColumnName keys and corresponding values.
// Returns a map of SystemColumnName keys to corresponding values.
func GetSystemStateColumnValues(systemState *metrics.SuiSystemState) map[enums.ColumnName]any {
	return map[enums.ColumnName]any{
		enums.ColumnNameIndex:                                       1,
		enums.ColumnNameSystemEpoch:                                 systemState.Epoch,
		enums.ColumnNameSystemEpochStartTimestamp:                   utility.EpochToUTCDate(systemState.EpochStartTimestampMs),
		enums.ColumnNameSystemEpochDuration:                         utility.MSToHoursAndMinutes(systemState.EpochDurationMs),
		enums.ColumnNameSystemTotalStake:                            systemState.TotalStake,
		enums.ColumnNameSystemStorageFundTotalObjectStorageRebates:  systemState.StorageFundTotalObjectStorageRebates,
		enums.ColumnNameSystemStorageFundNonRefundableBalance:       systemState.StorageFundNonRefundableBalance,
		enums.ColumnNameSystemReferenceGasPrice:                     systemState.ReferenceGasPrice,
		enums.ColumnNameSystemStakeSubsidyStartEpoch:                systemState.StakeSubsidyStartEpoch,
		enums.ColumnNameSystemStakeSubsidyBalance:                   systemState.StakeSubsidyBalance,
		enums.ColumnNameSystemStakeSubsidyDistributionCounter:       systemState.StakeSubsidyDistributionCounter,
		enums.ColumnNameSystemStakeSubsidyCurrentDistributionAmount: systemState.StakeSubsidyCurrentDistributionAmount,
		enums.ColumnNameSystemStakeSubsidyPeriodLength:              systemState.StakeSubsidyPeriodLength,
		enums.ColumnNameSystemStakeSubsidyDecreaseRate:              systemState.StakeSubsidyDecreaseRate,
	}
}

// GetValidatorCountsColumnValues returns a map of ColumnName values to corresponding values for the system state validators.
// The function retrieves information about the system state from the host's internal state and formats it into a map of ColumnName keys and corresponding values.
// Returns a map of ColumnName keys to corresponding values.
func GetValidatorCountsColumnValues(systemState *metrics.SuiSystemState) map[enums.ColumnName]any {
	return map[enums.ColumnName]any{
		enums.ColumnNameIndex:                                1,
		enums.ColumnNameSystemMaxValidatorCount:              systemState.MaxValidatorCount,
		enums.ColumnNameSystemActiveValidatorCount:           len(systemState.ActiveValidators),
		enums.ColumnNameSystemPendingActiveValidatorCount:    systemState.PendingActiveValidatorsSize,
		enums.ColumnNameSystemValidatorCandidateCount:        systemState.ValidatorCandidatesSize,
		enums.ColumnNameSystemPendingRemovalsCount:           len(systemState.PendingRemovals),
		enums.ColumnNameSystemAtRiskValidatorCount:           len(systemState.AtRiskValidators),
		enums.ColumnNameSystemMinValidatorJoiningStake:       systemState.MinValidatorJoiningStake,
		enums.ColumnNameSystemValidatorLowStakeThreshold:     systemState.ValidatorLowStakeThreshold,
		enums.ColumnNameSystemValidatorVeryLowStakeThreshold: systemState.ValidatorVeryLowStakeThreshold,
		enums.ColumnNameSystemValidatorLowStakeGracePeriod:   systemState.ValidatorLowStakeGracePeriod,
	}
}

// GetValidatorReportColumnValues returns a map of ColumnName values to corresponding values for the system state validators rpcgw.
// The function retrieves information about the system state from the host's internal state and formats it into a map of ColumnName keys and corresponding values.
// Returns a map of ColumnName keys to corresponding values.
func GetValidatorReportColumnValues(idx int, report metrics.ValidatorReport) tablebuilder.ColumnValues {
	var indexValue any = idx + 1

	if report.ReportedAddress == tablebuilder.EmptyValue {
		indexValue = tablebuilder.EmptyValue
	}

	return tablebuilder.ColumnValues{
		enums.ColumnNameIndex:                          indexValue,
		enums.ColumnNameSystemValidatorReportedName:    report.ReportedName,
		enums.ColumnNameSystemValidatorReportedAddress: report.ReportedAddress,
		enums.ColumnNameSystemValidatorReporterName:    report.ReporterName,
		enums.ColumnNameSystemValidatorReporterAddress: report.ReporterAddress,
	}
}

// GetValidatorAtRiskColumnValues returns a map of ColumnName values to corresponding values for the system state validators rpcgw.
// The function retrieves information about the system state from the host's internal state and formats it into a map of ColumnName keys and corresponding values.
// Returns a map of ColumnName keys to corresponding values.
func GetValidatorAtRiskColumnValues(idx int, validator metrics.ValidatorAtRisk) tablebuilder.ColumnValues {
	return tablebuilder.ColumnValues{
		enums.ColumnNameIndex:                               idx + 1,
		enums.ColumnNameSystemAtRiskValidatorName:           validator.Name,
		enums.ColumnNameSystemAtRiskValidatorAddress:        validator.Address,
		enums.ColumnNameSystemAtRiskValidatorNumberOfEpochs: validator.EpochsAtRisk,
	}
}
