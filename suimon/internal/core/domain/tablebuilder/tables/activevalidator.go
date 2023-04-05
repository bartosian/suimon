package tables

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
)

var (
	TableSortConfigActiveValidator = tableSortConfig{
		{Name: string(enums.ColumnNameValidatorNextEpochStake), Mode: table.Asc},
		{Name: string(enums.ColumnNameValidatorVotingPower), Mode: table.Asc},
	}
	ColumnConfigActiveValidator = tableColumnConfig{
		enums.ColumnNameIndex: {
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		enums.ColumnNameValidatorName: {
			Align:        text.AlignLeft,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		enums.ColumnNameValidatorNetAddress: {
			Align:        text.AlignLeft,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		enums.ColumnNameValidatorVotingPower: {
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		enums.ColumnNameValidatorGasPrice: {
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		enums.ColumnNameValidatorCommissionRate: {
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		enums.ColumnNameValidatorNextEpochStake: {
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		enums.ColumnNameValidatorNextEpochGasPrice: {
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		enums.ColumnNameValidatorNextEpochCommissionRate: {
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		enums.ColumnNameValidatorStakingPoolSuiBalance: {
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		enums.ColumnNameValidatorRewardsPool: {
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		enums.ColumnNameValidatorPoolTokenBalance: {
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		enums.ColumnNameValidatorPendingStake: {
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		enums.ColumnNameValidatorPendingTotalSuiWithdraw: {
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		enums.ColumnNameValidatorPendingPoolTokenWithdraw: {
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
	}
	RowsActiveValidator = tableRows{
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
