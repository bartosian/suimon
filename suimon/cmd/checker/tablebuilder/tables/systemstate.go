package tables

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums/columnnames"
)

var (
	TableStyleSystem      = table.StyleLight
	TableTagSystem        = ""
	TableSortConfigSystem = make([]table.SortBy, 0)
	ColumnConfigSystem    = []table.ColumnConfig{
		columnnames.SystemColumnNameEpoch: {
			Name:         "EPOCH",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.SystemColumnNameEpochDurationMs: {
			Name:         "EPOCH DURATION",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.SystemColumnNameStorageFund: {
			Name:         "STORAGE FUND",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.SystemColumnNameReferenceGasPrice: {
			Name:         "REFERENCE GAS\nPRICE",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.SystemColumnNameStakeSubsidyCounter: {
			Name:         "STAKE SUBSIDY\nCOUNTER",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.SystemColumnNameStakeSubsidyBalance: {
			Name:         "STAKE SUBSIDY BALANCE",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.SystemColumnNameStakeSubsidyCurrentEpochAmount: {
			Name:         "STAKE SUBSIDY\nCURRENT EPOCH AMOUNT",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.SystemColumnNameTotalStake: {
			Name:         "TOTAL STAKE",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.SystemColumnNameValidatorsCount: {
			Name:         "VALIDATORS\nCOUNT",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.SystemColumnNamePendingActiveValidatorsSize: {
			Name:         "PENDING VALIDATORS\nCOUNT",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.SystemColumnNamePendingRemovals: {
			Name:         "PENDING\nREMOVALS",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.SystemColumnNameValidatorsCandidateSize: {
			Name:         "VALIDATORS CANDIDATE\nCOUNT",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.SystemColumnNameValidatorsAtRiskCount: {
			Name:         "VALIDATORS AT RISK\nCOUNT",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
	}
)
