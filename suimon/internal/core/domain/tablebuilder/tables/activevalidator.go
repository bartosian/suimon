package tables

import (
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

var (
	TableStyleActiveValidator      = table.StyleLight
	TableColorsActiveValidator     = text.Colors{text.BgHiBlue, text.FgBlack}
	TableTagActiveValidator        = ""
	TableSortConfigActiveValidator = tableSortConfig{
		{Name: "NEXT EPOCH STAKE", Mode: table.Dsc},
		{Name: "NEXT EPOCH\nGAS PRICE", Mode: table.Dsc},
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
			enums.ColumnNameIndex,
			enums.ColumnNameValidatorName,
			enums.ColumnNameValidatorNetAddress,
			enums.ColumnNameValidatorVotingPower,
			enums.ColumnNameValidatorGasPrice,
			enums.ColumnNameValidatorCommissionRate,
			enums.ColumnNameValidatorNextEpochStake,
			enums.ColumnNameValidatorNextEpochGasPrice,
			enums.ColumnNameValidatorNextEpochCommissionRate,
		},
		1: {
			enums.ColumnNameValidatorStakingPoolSuiBalance,
			enums.ColumnNameValidatorRewardsPool,
			enums.ColumnNameValidatorPoolTokenBalance,
			enums.ColumnNameValidatorPendingStake,
			enums.ColumnNameValidatorPendingTotalSuiWithdraw,
			enums.ColumnNameValidatorPendingPoolTokenWithdraw,
		},
	}
)
