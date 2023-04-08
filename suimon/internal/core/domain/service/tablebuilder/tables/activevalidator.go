package tables

import (
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/metrics"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/service/tablebuilder"
)

var (
	SortConfigActiveValidator = tablebuilder.SortConfig{
		{Name: string(enums.ColumnNameValidatorNextEpochStake), Mode: table.Asc},
		{Name: string(enums.ColumnNameValidatorVotingPower), Mode: table.Asc},
	}
	ColumnsConfigActiveValidator = tablebuilder.ColumnsConfig{
		enums.ColumnNameIndex:                             tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameValidatorName:                     tablebuilder.NewDefaultColumnConfig(text.AlignLeft, text.AlignCenter, false),
		enums.ColumnNameValidatorNetAddress:               tablebuilder.NewDefaultColumnConfig(text.AlignLeft, text.AlignCenter, false),
		enums.ColumnNameValidatorVotingPower:              tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameValidatorGasPrice:                 tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameValidatorCommissionRate:           tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameValidatorNextEpochStake:           tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameValidatorNextEpochGasPrice:        tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameValidatorNextEpochCommissionRate:  tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameValidatorStakingPoolSuiBalance:    tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameValidatorRewardsPool:              tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameValidatorPoolTokenBalance:         tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameValidatorPendingStake:             tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameValidatorPendingTotalSuiWithdraw:  tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameValidatorPendingPoolTokenWithdraw: tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
	}
	RowsActiveValidator = tablebuilder.RowsConfig{
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
func GetActiveValidatorColumnValues(idx int, validator metrics.Validator) tablebuilder.ColumnValues {
	return tablebuilder.ColumnValues{
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
