package tables

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums/columnnames"
)

var (
	TableStyleSystem      = table.StyleLight
	TableTagSystem        = ""
	TableSortConfigSystem = make([]table.SortBy, 0)
	ColumnsConfigSystem   = []table.ColumnConfig{
		columnnames.SystemColumnNameIndex: {
			Name:         "IDX",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
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
			Name:         "REFERENCE\nGAS PRICE",
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
			Name:         "PENDING\nVALIDATORS COUNT",
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
			Name:         "VALIDATORS\nCANDIDATE COUNT",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.SystemColumnNameValidatorsAtRiskCount: {
			Name:         "VALIDATORS\nAT RISK COUNT",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
	}
	RowsSystemState = [][]int{
		0: {
			int(columnnames.SystemColumnNameIndex),
			int(columnnames.SystemColumnNameEpoch),
			int(columnnames.SystemColumnNameEpochDurationMs),
			int(columnnames.SystemColumnNameStorageFund),
			int(columnnames.SystemColumnNameReferenceGasPrice),
			int(columnnames.SystemColumnNameStakeSubsidyCounter),
			int(columnnames.SystemColumnNameStakeSubsidyBalance),
			int(columnnames.SystemColumnNameStakeSubsidyCurrentEpochAmount),
			int(columnnames.SystemColumnNameTotalStake),
			int(columnnames.SystemColumnNameValidatorsCount),
			int(columnnames.SystemColumnNamePendingActiveValidatorsSize),
			int(columnnames.SystemColumnNamePendingRemovals),
			int(columnnames.SystemColumnNameValidatorsCandidateSize),
			int(columnnames.SystemColumnNameValidatorsAtRiskCount),
		},
	}
)
