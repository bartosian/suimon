package tables

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	domainmetrics "github.com/bartosian/sui_helpers/suimon/internal/core/domain/metrics"
)

var (
	ColumnsConfigSystem = ColumnsConfig{
		enums.ColumnNameIndex:                                       NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemEpoch:                                 NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemEpochStartTimestamp:                   NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemEpochDuration:                         NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemTimeTillNextEpoch:                     NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
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
		enums.ColumnNameSystemValidatorSlashingPercentage:           NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemValidatorReporterName:                 NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemValidatorReporterAddress:              NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemMinReferenceGasPrice:                  NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemMaxReferenceGasPrice:                  NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemMeanReferenceGasPrice:                 NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemStakeWeightedMeanReferenceGasPrice:    NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemMedianReferenceGasPrice:               NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSystemEstimatedReferenceGasPrice:            NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
	}

	RowsConfigSystemState = RowsConfig{
		0: {
			enums.ColumnNameSystemEpoch,
			enums.ColumnNameSystemEpochDuration,
			enums.ColumnNameSystemEpochStartTimestamp,
			enums.ColumnNameSystemTimeTillNextEpoch,
			enums.ColumnNameSystemTotalStake,
			enums.ColumnNameSystemStorageFundTotalObjectStorageRebates,
			enums.ColumnNameSystemStorageFundNonRefundableBalance,
			enums.ColumnNameSystemStakeSubsidyStartEpoch,
			enums.ColumnNameSystemStakeSubsidyBalance,
			enums.ColumnNameSystemStakeSubsidyDistributionCounter,
			enums.ColumnNameSystemStakeSubsidyCurrentDistributionAmount,
		},
		1: {
			enums.ColumnNameSystemStakeSubsidyPeriodLength,
			enums.ColumnNameSystemStakeSubsidyDecreaseRate,
			enums.ColumnNameSystemReferenceGasPrice,
			enums.ColumnNameSystemMinReferenceGasPrice,
			enums.ColumnNameSystemMaxReferenceGasPrice,
			enums.ColumnNameSystemMeanReferenceGasPrice,
			enums.ColumnNameSystemStakeWeightedMeanReferenceGasPrice,
			enums.ColumnNameSystemMedianReferenceGasPrice,
			enums.ColumnNameSystemEstimatedReferenceGasPrice,
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
			enums.ColumnNameSystemValidatorReportedName,
			enums.ColumnNameSystemValidatorSlashingPercentage,
			enums.ColumnNameSystemValidatorReporterName,
			enums.ColumnNameSystemValidatorReporterAddress,
		},
	}
)

// GetSystemStateColumnValues returns a map of SystemColumnName values to corresponding values for the system state.
// The function retrieves information about the system state from the host's internal state and formats it into a map of SystemColumnName keys and corresponding values.
// Returns a map of SystemColumnName keys to corresponding values.
func GetSystemStateColumnValues(metrics *domainmetrics.Metrics) (ColumnValues, error) {
	result := ColumnValues{
		enums.ColumnNameIndex:                                       1,
		enums.ColumnNameSystemEpoch:                                 metrics.SystemState.Epoch,
		enums.ColumnNameSystemEpochStartTimestamp:                   metrics.EpochStartTimeUTC,
		enums.ColumnNameSystemEpochDuration:                         metrics.EpochDurationHHMM,
		enums.ColumnNameSystemTimeTillNextEpoch:                     fmt.Sprintf("%s HH:MM", metrics.DurationTillEpochEndHHMM),
		enums.ColumnNameSystemTotalStake:                            metrics.SystemState.TotalStake,
		enums.ColumnNameSystemStorageFundTotalObjectStorageRebates:  metrics.SystemState.StorageFundTotalObjectStorageRebates,
		enums.ColumnNameSystemStorageFundNonRefundableBalance:       metrics.SystemState.StorageFundNonRefundableBalance,
		enums.ColumnNameSystemStakeSubsidyStartEpoch:                metrics.SystemState.StakeSubsidyStartEpoch,
		enums.ColumnNameSystemStakeSubsidyBalance:                   metrics.SystemState.StakeSubsidyBalance,
		enums.ColumnNameSystemStakeSubsidyDistributionCounter:       metrics.SystemState.StakeSubsidyDistributionCounter,
		enums.ColumnNameSystemStakeSubsidyCurrentDistributionAmount: metrics.SystemState.StakeSubsidyCurrentDistributionAmount,
		enums.ColumnNameSystemStakeSubsidyPeriodLength:              metrics.SystemState.StakeSubsidyPeriodLength,
		enums.ColumnNameSystemStakeSubsidyDecreaseRate:              metrics.SystemState.StakeSubsidyDecreaseRate,
		enums.ColumnNameSystemReferenceGasPrice:                     metrics.SystemState.ReferenceGasPrice,
		enums.ColumnNameSystemMinReferenceGasPrice:                  metrics.MinReferenceGasPrice,
		enums.ColumnNameSystemMaxReferenceGasPrice:                  metrics.MaxReferenceGasPrice,
		enums.ColumnNameSystemMeanReferenceGasPrice:                 metrics.MeanReferenceGasPrice,
		enums.ColumnNameSystemStakeWeightedMeanReferenceGasPrice:    metrics.StakeWeightedMeanReferenceGasPrice,
		enums.ColumnNameSystemMedianReferenceGasPrice:               metrics.MedianReferenceGasPrice,
		enums.ColumnNameSystemEstimatedReferenceGasPrice:            metrics.EstimatedNextReferenceGasPrice,
	}

	mistValues := map[enums.ColumnName]string{
		enums.ColumnNameSystemTotalStake:                            metrics.SystemState.TotalStake,
		enums.ColumnNameSystemStorageFundTotalObjectStorageRebates:  metrics.SystemState.StorageFundTotalObjectStorageRebates,
		enums.ColumnNameSystemStorageFundNonRefundableBalance:       metrics.SystemState.StorageFundNonRefundableBalance,
		enums.ColumnNameSystemStakeSubsidyBalance:                   metrics.SystemState.StakeSubsidyBalance,
		enums.ColumnNameSystemStakeSubsidyCurrentDistributionAmount: metrics.SystemState.StakeSubsidyCurrentDistributionAmount,
	}

	for columnName, mistValue := range mistValues {
		intValue, err := domainmetrics.MistToSui(mistValue)
		if err != nil {
			return nil, err
		}

		result[columnName] = intValue
	}

	return result, nil
}

// GetValidatorCountsColumnValues returns a map of ColumnName values to corresponding values for the system state validators.
// The function retrieves information about the system state from the host's internal state and formats it into a map of ColumnName keys and corresponding values.
// Returns a map of ColumnName keys to corresponding values.
func GetValidatorCountsColumnValues(systemState *domainmetrics.SuiSystemState) (ColumnValues, error) {
	result := ColumnValues{
		enums.ColumnNameIndex:                              1,
		enums.ColumnNameSystemMaxValidatorCount:            systemState.MaxValidatorCount,
		enums.ColumnNameSystemActiveValidatorCount:         len(systemState.ActiveValidators),
		enums.ColumnNameSystemPendingActiveValidatorCount:  systemState.PendingActiveValidatorsSize,
		enums.ColumnNameSystemValidatorCandidateCount:      systemState.ValidatorCandidatesSize,
		enums.ColumnNameSystemPendingRemovalsCount:         len(systemState.PendingRemovals),
		enums.ColumnNameSystemAtRiskValidatorCount:         len(systemState.AtRiskValidators),
		enums.ColumnNameSystemValidatorLowStakeGracePeriod: systemState.ValidatorLowStakeGracePeriod,
	}

	mistValues := map[enums.ColumnName]string{
		enums.ColumnNameSystemMinValidatorJoiningStake:       systemState.MinValidatorJoiningStake,
		enums.ColumnNameSystemValidatorLowStakeThreshold:     systemState.ValidatorLowStakeThreshold,
		enums.ColumnNameSystemValidatorVeryLowStakeThreshold: systemState.ValidatorVeryLowStakeThreshold,
	}

	for columnName, mistValue := range mistValues {
		intValue, err := domainmetrics.MistToSui(mistValue)
		if err != nil {
			return nil, err
		}

		result[columnName] = intValue
	}

	return result, nil
}

// GetValidatorReportColumnValues returns a map of ColumnName values to corresponding values for the system state validator.
// The function retrieves information about the system state from the host's internal state and formats it into a map of ColumnName keys and corresponding values.
// Returns a map of ColumnName keys to corresponding values.
func GetValidatorReportColumnValues(reportedName string, slashingPct string, reporter domainmetrics.ValidatorReporter) ColumnValues {
	return ColumnValues{
		enums.ColumnNameSystemValidatorReportedName:       reportedName,
		enums.ColumnNameSystemValidatorSlashingPercentage: slashingPct,
		enums.ColumnNameSystemValidatorReporterName:       reporter.Name,
		enums.ColumnNameSystemValidatorReporterAddress:    reporter.Address,
	}
}

// GetValidatorAtRiskColumnValues returns a map of ColumnName values to corresponding values for the system state validators at risk.
// The function retrieves information about the system state from the host's internal state and formats it into a map of ColumnName keys and corresponding values.
// Returns a map of ColumnName keys to corresponding values.
func GetValidatorAtRiskColumnValues(idx int, validator domainmetrics.ValidatorAtRisk) ColumnValues {
	return ColumnValues{
		enums.ColumnNameIndex:                               idx + 1,
		enums.ColumnNameSystemAtRiskValidatorName:           validator.Name,
		enums.ColumnNameSystemAtRiskValidatorAddress:        validator.Address,
		enums.ColumnNameSystemAtRiskValidatorNumberOfEpochs: validator.EpochsAtRisk,
	}
}
