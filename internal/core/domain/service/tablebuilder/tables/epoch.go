package tables

import (
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/bartosian/suimon/internal/core/domain/enums"
	domainmetrics "github.com/bartosian/suimon/internal/core/domain/metrics"
)

var (
	ColumnsConfigEpoch = ColumnsConfig{
		enums.ColumnNameEpoch:                             NewDefaultColumnConfig(text.AlignLeft, text.AlignCenter, false),
		enums.ColumnNameEpochTotalTransactions:            NewDefaultColumnConfig(text.AlignLeft, text.AlignCenter, false),
		enums.ColumnNameEpochStartTimestamp:               NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameEpochEndTimestamp:                 NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameEpochFirstCheckpointId:            NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameEpochLastCheckpointId:             NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameEpochProtocolVersion:              NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameEpochReferenceGasPrice:            NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameEpochTotalStake:                   NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameEpochStorageFundReinvestment:      NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameEpochStorageCharge:                NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameEpochStorageRebate:                NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameEpochStorageFundBalance:           NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameEpochStakeSubsidyAmount:           NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameEpochTotalGasFees:                 NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameEpochTotalStakeRewardsDistributed: NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameEpochLeftoverStorageFundInflow:    NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
	}
	RowsConfigEpoch = RowsConfig{
		0: {
			enums.ColumnNameEpoch,
			enums.ColumnNameEpochTotalTransactions,
			enums.ColumnNameEpochStartTimestamp,
			enums.ColumnNameEpochEndTimestamp,
			enums.ColumnNameEpochFirstCheckpointId,
			enums.ColumnNameEpochLastCheckpointId,
			enums.ColumnNameEpochProtocolVersion,
			enums.ColumnNameEpochReferenceGasPrice,
			enums.ColumnNameEpochTotalStake,
			enums.ColumnNameEpochStorageCharge,
			enums.ColumnNameEpochStorageRebate,
			enums.ColumnNameEpochStorageFundBalance,
			enums.ColumnNameEpochStakeSubsidyAmount,
			enums.ColumnNameEpochTotalGasFees,
			enums.ColumnNameEpochTotalStakeRewardsDistributed,
		},
	}
)

func GetEpochColumnValues(idx int, epoch *domainmetrics.EpochInfo) (ColumnValues, error) {
	result := ColumnValues{
		enums.ColumnNameEpoch:                  epoch.Epoch,
		enums.ColumnNameEpochTotalTransactions: epoch.EpochTotalTransactions,
		enums.ColumnNameEpochStartTimestamp:    epoch.EpochStartTimestamp,
		enums.ColumnNameEpochEndTimestamp:      epoch.EndOfEpochInfo.EpochEndTimestamp,
		enums.ColumnNameEpochFirstCheckpointId: epoch.FirstCheckpointID,
		enums.ColumnNameEpochLastCheckpointId:  epoch.EndOfEpochInfo.LastCheckpointID,
		enums.ColumnNameEpochProtocolVersion:   epoch.EndOfEpochInfo.ProtocolVersion,
		enums.ColumnNameEpochReferenceGasPrice: epoch.EndOfEpochInfo.ReferenceGasPrice,
	}

	mistValues := map[enums.ColumnName]string{
		enums.ColumnNameEpochTotalStake:                   epoch.EndOfEpochInfo.TotalStake,
		enums.ColumnNameEpochStorageFundReinvestment:      epoch.EndOfEpochInfo.StorageFundReinvestment,
		enums.ColumnNameEpochStorageCharge:                epoch.EndOfEpochInfo.StorageCharge,
		enums.ColumnNameEpochStorageRebate:                epoch.EndOfEpochInfo.StorageRebate,
		enums.ColumnNameEpochStorageFundBalance:           epoch.EndOfEpochInfo.StorageFundBalance,
		enums.ColumnNameEpochStakeSubsidyAmount:           epoch.EndOfEpochInfo.StakeSubsidyAmount,
		enums.ColumnNameEpochTotalGasFees:                 epoch.EndOfEpochInfo.TotalGasFees,
		enums.ColumnNameEpochTotalStakeRewardsDistributed: epoch.EndOfEpochInfo.TotalStakeRewardsDistributed,
		enums.ColumnNameEpochLeftoverStorageFundInflow:    epoch.EndOfEpochInfo.LeftoverStorageFundInflow,
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
