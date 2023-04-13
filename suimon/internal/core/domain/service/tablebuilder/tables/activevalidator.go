package tables

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/metrics"
)

var (
	SortConfigActiveValidator = SortConfig{
		{Name: string(enums.ColumnNameValidatorNextEpochStake), Mode: table.Asc},
		{Name: string(enums.ColumnNameValidatorVotingPower), Mode: table.Asc},
	}
	ColumnsConfigActiveValidator = ColumnsConfig{
		enums.ColumnNameIndex:                             NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameValidatorName:                     NewDefaultColumnConfig(text.AlignLeft, text.AlignCenter, false),
		enums.ColumnNameValidatorNetAddress:               NewDefaultColumnConfig(text.AlignLeft, text.AlignCenter, false),
		enums.ColumnNameValidatorVotingPower:              NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameValidatorGasPrice:                 NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameValidatorCommissionRate:           NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameValidatorNextEpochStake:           NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameValidatorNextEpochGasPrice:        NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameValidatorNextEpochCommissionRate:  NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameValidatorStakingPoolSuiBalance:    NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameValidatorRewardsPool:              NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameValidatorPoolTokenBalance:         NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameValidatorPendingStake:             NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameValidatorPendingTotalSuiWithdraw:  NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameValidatorPendingPoolTokenWithdraw: NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
	}
	RowsActiveValidator = RowsConfig{
		0: {
			enums.ColumnNameValidatorName,
			enums.ColumnNameValidatorVotingPower,
			enums.ColumnNameValidatorGasPrice,
			enums.ColumnNameValidatorCommissionRate,
			enums.ColumnNameValidatorNextEpochStake,
			enums.ColumnNameValidatorNextEpochGasPrice,
			enums.ColumnNameValidatorNextEpochCommissionRate,
			enums.ColumnNameValidatorStakingPoolSuiBalance,
			enums.ColumnNameValidatorRewardsPool,
			enums.ColumnNameValidatorPoolTokenBalance,
			enums.ColumnNameValidatorPendingStake,
		},
	}
)

// GetActiveValidatorColumnValues returns a map of ActiveValidatorColumnName values to corresponding values for the specified active validator.
// The function retrieves information about the active validator from the provided metrics.Validator object and formats it into a map of ActiveValidatorColumnName keys and corresponding values.
// Returns a map of ActiveValidatorColumnName keys to corresponding values.
func GetActiveValidatorColumnValues(idx int, validator metrics.Validator) ColumnValues {
	return ColumnValues{
		enums.ColumnNameIndex:                             idx + 1,
		enums.ColumnNameValidatorName:                     validator.Name,
		enums.ColumnNameValidatorNetAddress:               validator.NetAddress,
		enums.ColumnNameValidatorVotingPower:              validator.VotingPower,
		enums.ColumnNameValidatorGasPrice:                 validator.GasPrice,
		enums.ColumnNameValidatorCommissionRate:           validator.CommissionRate,
		enums.ColumnNameValidatorNextEpochStake:           validator.NextEpochStake,
		enums.ColumnNameValidatorNextEpochGasPrice:        validator.NextEpochGasPrice,
		enums.ColumnNameValidatorNextEpochCommissionRate:  validator.NextEpochCommissionRate,
		enums.ColumnNameValidatorStakingPoolSuiBalance:    validator.StakingPoolSuiBalance,
		enums.ColumnNameValidatorRewardsPool:              validator.RewardsPool,
		enums.ColumnNameValidatorPoolTokenBalance:         validator.PoolTokenBalance,
		enums.ColumnNameValidatorPendingStake:             validator.PendingStake,
		enums.ColumnNameValidatorPendingTotalSuiWithdraw:  validator.PendingTotalSuiWithdraw,
		enums.ColumnNameValidatorPendingPoolTokenWithdraw: validator.PendingPoolTokenWithdraw,
	}
}
