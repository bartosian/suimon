package dashboards

import (
	"github.com/mum4k/termdash/cell"

	"github.com/bartosian/suimon/internal/core/domain/enums"
	domainhost "github.com/bartosian/suimon/internal/core/domain/host"
	domainmetrics "github.com/bartosian/suimon/internal/core/domain/metrics"
)

const (
	ColumnWidth14 = 14
	ColumnWidth15 = 15
	ColumnWidth20 = 20
	ColumnWidth24 = 24
	ColumnWidth25 = 25
	ColumnWidth30 = 30
	ColumnWidth33 = 33
	ColumnWidth49 = 49
	ColumnWidth99 = 20
	ColumnWidth19 = 19
	RowHeight14   = 14
)

var ColumnsConfigSystemState = ColumnsConfig{
	// Overview section
	enums.ColumnNameSystemEpochStartTimestamp: ColumnWidth99,
	enums.ColumnNameCurrentEpoch:              ColumnWidth20,
	enums.ColumnNameSystemEpochDuration:       ColumnWidth20,
	enums.ColumnNameSystemTimeTillNextEpoch:   ColumnWidth30,

	// Gas Price section
	enums.ColumnNameSystemReferenceGasPrice:                  ColumnWidth20,
	enums.ColumnNameSystemMinReferenceGasPrice:               ColumnWidth20,
	enums.ColumnNameSystemMaxReferenceGasPrice:               ColumnWidth20,
	enums.ColumnNameSystemMeanReferenceGasPrice:              ColumnWidth20,
	enums.ColumnNameSystemStakeWeightedMeanReferenceGasPrice: ColumnWidth19,
	enums.ColumnNameSystemMedianReferenceGasPrice:            ColumnWidth20,
	enums.ColumnNameSystemEstimatedReferenceGasPrice:         ColumnWidth20,

	// Subsidy section
	enums.ColumnNameSystemStakeSubsidyStartEpoch:                ColumnWidth20,
	enums.ColumnNameSystemStakeSubsidyBalance:                   ColumnWidth49,
	enums.ColumnNameSystemStakeSubsidyDistributionCounter:       ColumnWidth33,
	enums.ColumnNameSystemStakeSubsidyCurrentDistributionAmount: ColumnWidth49,
	enums.ColumnNameSystemStakeSubsidyPeriodLength:              ColumnWidth20,
	enums.ColumnNameSystemStakeSubsidyDecreaseRate:              ColumnWidth20,

	// Stake section
	enums.ColumnNameSystemTotalStake:                           ColumnWidth33,
	enums.ColumnNameSystemStorageFundTotalObjectStorageRebates: ColumnWidth33,
	enums.ColumnNameSystemStorageFundNonRefundableBalance:      ColumnWidth33,
}

var RowsConfigSystemState = RowsConfig{
	0: {
		Height: RowHeight14,
		Columns: []enums.ColumnName{
			enums.ColumnNameSystemEpochStartTimestamp,
		},
	},
	1: {
		Height: RowHeight14,
		Columns: []enums.ColumnName{
			enums.ColumnNameCurrentEpoch,
			enums.ColumnNameSystemEpochDuration,
			enums.ColumnNameSystemTimeTillNextEpoch,
		},
	},
	2: {
		Height: RowHeight14,
		Columns: []enums.ColumnName{
			enums.ColumnNameSystemTotalStake,
			enums.ColumnNameSystemStorageFundTotalObjectStorageRebates,
			enums.ColumnNameSystemStorageFundNonRefundableBalance,
		},
	},
	3: {
		Height: RowHeight14,
		Columns: []enums.ColumnName{
			enums.ColumnNameSystemReferenceGasPrice,
			enums.ColumnNameSystemMinReferenceGasPrice,
			enums.ColumnNameSystemMaxReferenceGasPrice,
			enums.ColumnNameSystemMeanReferenceGasPrice,
			enums.ColumnNameSystemStakeWeightedMeanReferenceGasPrice,
		},
	},
	4: {
		Height: RowHeight14,
		Columns: []enums.ColumnName{
			enums.ColumnNameSystemMedianReferenceGasPrice,
			enums.ColumnNameSystemEstimatedReferenceGasPrice,
			enums.ColumnNameSystemStakeSubsidyPeriodLength,
			enums.ColumnNameSystemStakeSubsidyDecreaseRate,
			enums.ColumnNameSystemStakeSubsidyStartEpoch,
		},
	},
	5: {
		Height: RowHeight14,
		Columns: []enums.ColumnName{
			enums.ColumnNameSystemStakeSubsidyCurrentDistributionAmount,
			enums.ColumnNameSystemStakeSubsidyBalance,
		},
	},
}

var CellsConfigSystemState = CellsConfig{
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

// GeSystemStateColumnValues returns a map of ColumnName values to corresponding values for a node at the specified index on the specified host.
// The function retrieves information about the node from the host's internal state and formats it into a map of NodeColumnName keys and corresponding values.
// The function also includes emoji values in the map if the specified flag is true.
func GeSystemStateColumnValues(host *domainhost.Host) (ColumnValues, error) {
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
