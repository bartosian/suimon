package tables

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"strconv"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/metrics"
	"github.com/bartosian/sui_helpers/suimon/internal/pkg/utility"
)

var (
	SortConfigSystem = SortConfig{
		{Name: enums.ColumnNameSystemEpoch.ToString(), Mode: table.Dsc},
	}
	SortConfigValidatorsAtRisk = SortConfig{
		{Name: enums.ColumnNameSystemAtRiskValidatorNumberOfEpochs.ToString(), Mode: table.Dsc},
	}
	ColumnsConfigSystem = ColumnsConfig{
		enums.ColumnNameIndex:                                       NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemEpoch:                                 NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemEpochStartTimestamp:                   NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemEpochDuration:                         NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemTotalStake:                            NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemStorageFundTotalObjectStorageRebates:  NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemStorageFundNonRefundableBalance:       NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemReferenceGasPrice:                     NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemStakeSubsidyStartEpoch:                NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemMaxValidatorCount:                     NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemMinValidatorJoiningStake:              NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemValidatorLowStakeThreshold:            NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemValidatorVeryLowStakeThreshold:        NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemValidatorLowStakeGracePeriod:          NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemStakeSubsidyBalance:                   NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemStakeSubsidyDistributionCounter:       NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemStakeSubsidyCurrentDistributionAmount: NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemStakeSubsidyPeriodLength:              NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemStakeSubsidyDecreaseRate:              NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemActiveValidatorCount:                  NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemPendingActiveValidatorCount:           NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemPendingRemovalsCount:                  NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemValidatorCandidateCount:               NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemAtRiskValidatorCount:                  NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemAtRiskValidatorName:                   NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemAtRiskValidatorAddress:                NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemAtRiskValidatorNumberOfEpochs:         NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemValidatorReportedName:                 NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemValidatorReportedAddress:              NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
	}
	RowsConfigSystemState = RowsConfig{
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
	RowsConfigValidatorCounts = RowsConfig{
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
	RowsConfigValidatorsAtRisk = RowsConfig{
		0: {
			enums.ColumnNameIndex,
			enums.ColumnNameSystemAtRiskValidatorName,
			enums.ColumnNameSystemAtRiskValidatorAddress,
			enums.ColumnNameSystemAtRiskValidatorNumberOfEpochs,
		},
	}
	RowsConfigValidatorReports = RowsConfig{
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
	epochStartMs, err := strconv.ParseInt(systemState.EpochStartTimestampMs, 10, 64)
	if err != nil {
		return nil
	}

	epochDurationMs, err := strconv.ParseInt(systemState.EpochDurationMs, 10, 64)
	if err != nil {
		return nil
	}

	return map[enums.ColumnName]any{
		enums.ColumnNameIndex:                                       1,
		enums.ColumnNameSystemEpoch:                                 systemState.Epoch,
		enums.ColumnNameSystemEpochStartTimestamp:                   utility.EpochToUTCDate(epochStartMs),
		enums.ColumnNameSystemEpochDuration:                         utility.MSToHoursAndMinutes(epochDurationMs),
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

// GetValidatorReportColumnValues returns a map of ColumnName values to corresponding values for the system state validator.
// The function retrieves information about the system state from the host's internal state and formats it into a map of ColumnName keys and corresponding values.
// Returns a map of ColumnName keys to corresponding values.
func GetValidatorReportColumnValues(idx int, report metrics.ValidatorReport) ColumnValues {
	var indexValue any = idx + 1

	if report.ReportedAddress == EmptyValue {
		indexValue = EmptyValue
	}

	return ColumnValues{
		enums.ColumnNameIndex:                          indexValue,
		enums.ColumnNameSystemValidatorReportedName:    report.ReportedName,
		enums.ColumnNameSystemValidatorReportedAddress: report.ReportedAddress,
		enums.ColumnNameSystemValidatorReporterName:    report.ReporterName,
		enums.ColumnNameSystemValidatorReporterAddress: report.ReporterAddress,
	}
}

// GetValidatorAtRiskColumnValues returns a map of ColumnName values to corresponding values for the system state validators at risk.
// The function retrieves information about the system state from the host's internal state and formats it into a map of ColumnName keys and corresponding values.
// Returns a map of ColumnName keys to corresponding values.
func GetValidatorAtRiskColumnValues(idx int, validator metrics.ValidatorAtRisk) ColumnValues {
	return ColumnValues{
		enums.ColumnNameIndex:                               idx + 1,
		enums.ColumnNameSystemAtRiskValidatorName:           validator.Name,
		enums.ColumnNameSystemAtRiskValidatorAddress:        validator.Address,
		enums.ColumnNameSystemAtRiskValidatorNumberOfEpochs: validator.EpochsAtRisk,
	}
}
