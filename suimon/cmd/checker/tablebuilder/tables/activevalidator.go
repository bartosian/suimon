package tables

import (
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums/columnnames"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

var (
	TableStyleActiveValidator      = table.StyleLight
	TableTagActiveValidator        = ""
	TableSortConfigActiveValidator = []table.SortBy{
		{Name: "NEXT EPOCH STAKE", Mode: table.Dsc},
		{Name: "NEXT EPOCH\nGAS PRICE", Mode: table.Dsc},
	}
	ColumnConfigActiveValidator = []table.ColumnConfig{
		columnnames.ActiveValidatorColumnNameIndex: {
			Name:         "IDX",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.ActiveValidatorColumnNameName: {
			Name:         "NAME",
			Align:        text.AlignLeft,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.ActiveValidatorColumnNameNetAddress: {
			Name:         "NET ADDRESS",
			Align:        text.AlignLeft,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.ActiveValidatorColumnNameVotingPower: {
			Name:         "VOTING POWER",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.ActiveValidatorColumnNameGasPrice: {
			Name:         "GAS PRIZE",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.ActiveValidatorColumnNameCommissionRate: {
			Name:         "COMMISSION RATE",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.ActiveValidatorColumnNameNextEpochStake: {
			Name:         "NEXT EPOCH STAKE",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.ActiveValidatorColumnNameNextEpochGasPrice: {
			Name:         "NEXT EPOCH\nGAS PRICE",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.ActiveValidatorColumnNameNextEpochCommissionRate: {
			Name:         "NEXT EPOCH\nCOMMISSION RATE",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.ActiveValidatorColumnNameStakingPoolSuiBalance: {
			Name:         "STAKING POOL\nSUI BALANCE",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.ActiveValidatorColumnNameRewardsPool: {
			Name:         "REWARDS POOL",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.ActiveValidatorColumnNamePoolTokenBalance: {
			Name:         "POOL TOKEN BALANCE",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.ActiveValidatorColumnNamePendingStake: {
			Name:         "POOL PENDING STAKE",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
	}
	RowsActiveValidator = [][]int{
		0: {
			int(columnnames.ActiveValidatorColumnNameIndex),
			int(columnnames.ActiveValidatorColumnNameName),
			int(columnnames.ActiveValidatorColumnNameNetAddress),
			int(columnnames.ActiveValidatorColumnNameVotingPower),
			int(columnnames.ActiveValidatorColumnNameGasPrice),
			int(columnnames.ActiveValidatorColumnNameCommissionRate),
			int(columnnames.ActiveValidatorColumnNameNextEpochStake),
			int(columnnames.ActiveValidatorColumnNameNextEpochGasPrice),
			int(columnnames.ActiveValidatorColumnNameNextEpochCommissionRate),
		},
		1: {
			int(columnnames.ActiveValidatorColumnNameStakingPoolSuiBalance),
			int(columnnames.ActiveValidatorColumnNameRewardsPool),
			int(columnnames.ActiveValidatorColumnNamePoolTokenBalance),
			int(columnnames.ActiveValidatorColumnNamePendingStake),
		},
	}
)
