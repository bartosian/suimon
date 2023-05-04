package tables

import (
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/bartosian/suimon/internal/core/domain/enums"
	domainmetrics "github.com/bartosian/suimon/internal/core/domain/metrics"
)

var (
	ColumnsConfigActiveValidator = ColumnsConfig{
		enums.ColumnNameIndex:                             NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameValidatorName:                     NewDefaultColumnConfig(text.AlignLeft, text.AlignCenter, false),
		enums.ColumnNameValidatorNetAddress:               NewDefaultColumnConfig(text.AlignLeft, text.AlignCenter, false),
		enums.ColumnNameValidatorVotingPower:              NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameValidatorGasPrice:                 NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameValidatorCommissionRate:           NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameValidatorApy:                      NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
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
			enums.ColumnNameIndex,
			enums.ColumnNameValidatorName,
			enums.ColumnNameValidatorVotingPower,
			enums.ColumnNameValidatorGasPrice,
			enums.ColumnNameValidatorCommissionRate,
			enums.ColumnNameValidatorApy,
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
func GetActiveValidatorColumnValues(idx int, validator *domainmetrics.Validator) (ColumnValues, error) {
	result := ColumnValues{
		enums.ColumnNameIndex:                             idx + 1,
		enums.ColumnNameValidatorName:                     validator.Name,
		enums.ColumnNameValidatorNetAddress:               validator.NetAddress,
		enums.ColumnNameValidatorVotingPower:              validator.VotingPower,
		enums.ColumnNameValidatorGasPrice:                 validator.GasPrice,
		enums.ColumnNameValidatorCommissionRate:           validator.CommissionRate,
		enums.ColumnNameValidatorApy:                      validator.APY,
		enums.ColumnNameValidatorNextEpochGasPrice:        validator.NextEpochGasPrice,
		enums.ColumnNameValidatorNextEpochCommissionRate:  validator.NextEpochCommissionRate,
		enums.ColumnNameValidatorPendingTotalSuiWithdraw:  validator.PendingTotalSuiWithdraw,
		enums.ColumnNameValidatorPendingPoolTokenWithdraw: validator.PendingPoolTokenWithdraw,
	}

	mistValues := map[enums.ColumnName]string{
		enums.ColumnNameValidatorNextEpochStake:        validator.NextEpochStake,
		enums.ColumnNameValidatorStakingPoolSuiBalance: validator.StakingPoolSuiBalance,
		enums.ColumnNameValidatorRewardsPool:           validator.RewardsPool,
		enums.ColumnNameValidatorPoolTokenBalance:      validator.PoolTokenBalance,
		enums.ColumnNameValidatorPendingStake:          validator.PendingStake,
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
