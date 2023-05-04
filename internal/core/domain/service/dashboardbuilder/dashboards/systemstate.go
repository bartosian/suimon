package dashboards

import (
	"github.com/mum4k/termdash/cell"

	"github.com/bartosian/suimon/internal/core/domain/enums"
	"github.com/bartosian/suimon/internal/core/domain/host"
	domainmetrics "github.com/bartosian/suimon/internal/core/domain/metrics"
)

var (
	ColumnsConfigSystemState = ColumnsConfig{
		// Overview section
		enums.ColumnNameSystemEpochStartTimestamp: 99,
		enums.ColumnNameCurrentEpoch:              20,
		enums.ColumnNameSystemEpochDuration:       20,
		enums.ColumnNameSystemTimeTillNextEpoch:   30,

		// Gas Price section
		enums.ColumnNameSystemReferenceGasPrice:                  20,
		enums.ColumnNameSystemMinReferenceGasPrice:               20,
		enums.ColumnNameSystemMaxReferenceGasPrice:               20,
		enums.ColumnNameSystemMeanReferenceGasPrice:              20,
		enums.ColumnNameSystemStakeWeightedMeanReferenceGasPrice: 19,
		enums.ColumnNameSystemMedianReferenceGasPrice:            20,
		enums.ColumnNameSystemEstimatedReferenceGasPrice:         20,

		// Subsidy section
		enums.ColumnNameSystemStakeSubsidyStartEpoch:                20,
		enums.ColumnNameSystemStakeSubsidyBalance:                   49,
		enums.ColumnNameSystemStakeSubsidyDistributionCounter:       33,
		enums.ColumnNameSystemStakeSubsidyCurrentDistributionAmount: 49,
		enums.ColumnNameSystemStakeSubsidyPeriodLength:              20,
		enums.ColumnNameSystemStakeSubsidyDecreaseRate:              20,

		// Stake section
		enums.ColumnNameSystemTotalStake:                           33,
		enums.ColumnNameSystemStorageFundTotalObjectStorageRebates: 33,
		enums.ColumnNameSystemStorageFundNonRefundableBalance:      33,
	}

	RowsConfigSystemState = RowsConfig{
		0: {
			Height: 14,
			Columns: []enums.ColumnName{
				enums.ColumnNameSystemEpochStartTimestamp,
			},
		},
		1: {
			Height: 14,
			Columns: []enums.ColumnName{
				enums.ColumnNameCurrentEpoch,
				enums.ColumnNameSystemEpochDuration,
				enums.ColumnNameSystemTimeTillNextEpoch,
			},
		},
		2: {
			Height: 14,
			Columns: []enums.ColumnName{
				enums.ColumnNameSystemTotalStake,
				enums.ColumnNameSystemStorageFundTotalObjectStorageRebates,
				enums.ColumnNameSystemStorageFundNonRefundableBalance,
			},
		},
		3: {
			Height: 14,
			Columns: []enums.ColumnName{
				enums.ColumnNameSystemReferenceGasPrice,
				enums.ColumnNameSystemMinReferenceGasPrice,
				enums.ColumnNameSystemMaxReferenceGasPrice,
				enums.ColumnNameSystemMeanReferenceGasPrice,
				enums.ColumnNameSystemStakeWeightedMeanReferenceGasPrice,
			},
		},
		4: {
			Height: 14,
			Columns: []enums.ColumnName{
				enums.ColumnNameSystemMedianReferenceGasPrice,
				enums.ColumnNameSystemEstimatedReferenceGasPrice,
				enums.ColumnNameSystemStakeSubsidyPeriodLength,
				enums.ColumnNameSystemStakeSubsidyDecreaseRate,
				enums.ColumnNameSystemStakeSubsidyStartEpoch,
			},
		},
		5: {
			Height: 14,
			Columns: []enums.ColumnName{
				enums.ColumnNameSystemStakeSubsidyCurrentDistributionAmount,
				enums.ColumnNameSystemStakeSubsidyBalance,
			},
		},
	}

	CellsConfigSystemState = CellsConfig{
		enums.ColumnNameCurrentEpoch:                                {"CURRENT EPOCH", cell.ColorGreen},
		enums.ColumnNameSystemEpochStartTimestamp:                   {"EPOCH START TIME UTC", cell.ColorGreen},
		enums.ColumnNameSystemEpochDuration:                         {"EPOCH DURATION", cell.ColorGreen},
		enums.ColumnNameSystemTimeTillNextEpoch:                     {"TIME TILL NEXT EPOCH", cell.ColorGreen},
		enums.ColumnNameSystemTotalStake:                            {"TOTAL STAKE, SUI", cell.ColorBlue},
		enums.ColumnNameSystemStorageFundTotalObjectStorageRebates:  {"STORAGE FUND TOTAL OBJECT REBATES, SUI", cell.ColorBlue},
		enums.ColumnNameSystemStorageFundNonRefundableBalance:       {"STORAGE FUND REFUNDABLE BALANCE, SUI", cell.ColorBlue},
		enums.ColumnNameSystemReferenceGasPrice:                     {"REFERENCE GAS PRICE", cell.ColorYellow},
		enums.ColumnNameSystemMinReferenceGasPrice:                  {"MIN REFERENCE GAS PRICE", cell.ColorYellow},
		enums.ColumnNameSystemMaxReferenceGasPrice:                  {"MAX REFERENCE GAS PRICE", cell.ColorYellow},
		enums.ColumnNameSystemMeanReferenceGasPrice:                 {"MEAN REFERENCE GAS PRICE", cell.ColorYellow},
		enums.ColumnNameSystemStakeWeightedMeanReferenceGasPrice:    {"STAKE WEIGHTED MEAN REFERENCE GAS PRICE", cell.ColorYellow},
		enums.ColumnNameSystemMedianReferenceGasPrice:               {"MEDIAN REFERENCE GAS PRICE", cell.ColorYellow},
		enums.ColumnNameSystemEstimatedReferenceGasPrice:            {"ESTIMATED REFERENCE GAS PRICE", cell.ColorYellow},
		enums.ColumnNameSystemStakeSubsidyBalance:                   {"STAKE SUBSIDY BALANCE, SUI", cell.ColorRed},
		enums.ColumnNameSystemStakeSubsidyStartEpoch:                {"STAKE SUBSIDY START EPOCH", cell.ColorRed},
		enums.ColumnNameSystemStakeSubsidyDistributionCounter:       {"STAKE SUBSIDY DISTRIBUTION COUNTER", cell.ColorRed},
		enums.ColumnNameSystemStakeSubsidyCurrentDistributionAmount: {"STAKE SUBSIDY DISTRIBUTION AMOUNT, SUI", cell.ColorRed},
		enums.ColumnNameSystemStakeSubsidyPeriodLength:              {"STAKE SUBSIDY PERIOD LENGTH", cell.ColorRed},
		enums.ColumnNameSystemStakeSubsidyDecreaseRate:              {"STAKE SUBSIDY DECREASE RATE", cell.ColorRed},
	}
)

// GeSystemStateColumnValues returns a map of ColumnName values to corresponding values for a node at the specified index on the specified host.
// The function retrieves information about the node from the host's internal state and formats it into a map of NodeColumnName keys and corresponding values.
// The function also includes emoji values in the map if the specified flag is true.
func GeSystemStateColumnValues(host host.Host) (ColumnValues, error) {
	result := ColumnValues{
		enums.ColumnNameCurrentEpoch:                             host.Metrics.SystemState.Epoch,
		enums.ColumnNameSystemEpochStartTimestamp:                host.Metrics.EpochStartTimeUTC,
		enums.ColumnNameSystemEpochDuration:                      host.Metrics.EpochDurationHHMM,
		enums.ColumnNameSystemTimeTillNextEpoch:                  host.Metrics.DurationTillEpochEndHHMM,
		enums.ColumnNameSystemStakeSubsidyStartEpoch:             host.Metrics.SystemState.StakeSubsidyStartEpoch,
		enums.ColumnNameSystemStakeSubsidyDistributionCounter:    host.Metrics.SystemState.StakeSubsidyDistributionCounter,
		enums.ColumnNameSystemStakeSubsidyPeriodLength:           host.Metrics.SystemState.StakeSubsidyPeriodLength,
		enums.ColumnNameSystemStakeSubsidyDecreaseRate:           host.Metrics.SystemState.StakeSubsidyDecreaseRate,
		enums.ColumnNameSystemReferenceGasPrice:                  host.Metrics.SystemState.ReferenceGasPrice,
		enums.ColumnNameSystemMinReferenceGasPrice:               host.Metrics.MinReferenceGasPrice,
		enums.ColumnNameSystemMaxReferenceGasPrice:               host.Metrics.MaxReferenceGasPrice,
		enums.ColumnNameSystemMeanReferenceGasPrice:              host.Metrics.MeanReferenceGasPrice,
		enums.ColumnNameSystemStakeWeightedMeanReferenceGasPrice: host.Metrics.StakeWeightedMeanReferenceGasPrice,
		enums.ColumnNameSystemMedianReferenceGasPrice:            host.Metrics.MedianReferenceGasPrice,
		enums.ColumnNameSystemEstimatedReferenceGasPrice:         host.Metrics.EstimatedNextReferenceGasPrice,
	}

	mistValues := map[enums.ColumnName]string{
		enums.ColumnNameSystemTotalStake:                            host.Metrics.SystemState.TotalStake,
		enums.ColumnNameSystemStorageFundTotalObjectStorageRebates:  host.Metrics.SystemState.StorageFundTotalObjectStorageRebates,
		enums.ColumnNameSystemStorageFundNonRefundableBalance:       host.Metrics.SystemState.StorageFundNonRefundableBalance,
		enums.ColumnNameSystemStakeSubsidyBalance:                   host.Metrics.SystemState.StakeSubsidyBalance,
		enums.ColumnNameSystemStakeSubsidyCurrentDistributionAmount: host.Metrics.SystemState.StakeSubsidyCurrentDistributionAmount,
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
