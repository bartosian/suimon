package tables

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
)

var (
	TableSortConfigRPC = tableSortConfig{
		{Name: string(enums.ColumnNameHealth), Mode: table.Dsc},
		{Name: string(enums.ColumnNameTotalTransactionBlocks), Mode: table.Dsc},
		{Name: string(enums.ColumnNameLatestCheckpoint), Mode: table.Dsc},
	}
	ColumnsConfigRPC = tableColumnConfig{
		enums.ColumnNameIndex: {
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		enums.ColumnNameHealth: {
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		enums.ColumnNameAddress: {
			Align:        text.AlignLeft,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		enums.ColumnNamePortRPC: {
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		enums.ColumnNameTotalTransactionBlocks: {
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		enums.ColumnNameLatestCheckpoint: {
			Align:        text.AlignLeft,
			AlignHeader:  text.AlignLeft,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
	}
	RowsRPC = tableRows{
		0: {
			enums.ColumnNameHealth,
			enums.ColumnNameAddress,
			enums.ColumnNamePortRPC,
			enums.ColumnNameTotalTransactionBlocks,
			enums.ColumnNameLatestCheckpoint,
		},
	}
)
