package tables

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums/columnnames"
)

var (
	TableStyleRPC      = table.StyleLight
	TableTagRPC        = ""
	TableSortConfigRPC = []table.SortBy{
		{Name: "HEALTH", Mode: table.Dsc},
		{Name: "TOTAL\nTRANSACTIONS", Mode: table.Dsc},
		{Name: "LATEST\nCHECKPOINT", Mode: table.Dsc},
	}
	ColumnsConfigRPC = []table.ColumnConfig{
		columnnames.NodeColumnNameIndex: {
			Name:         "",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.NodeColumnNameHealth: {
			Name:         "HEALTH",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.NodeColumnNameAddress: {
			Name:         "ADDRESS",
			Align:        text.AlignLeft,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.NodeColumnNamePortRPC: {
			Name:         "RPC",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.NodeColumnNameTotalTransactions: {
			Name:         "TOTAL\nTRANSACTIONS",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.NodeColumnNameLatestCheckpoint: {
			Name:         "LATEST\nCHECKPOINT",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
	}
	RowsRPC = [][]int{
		0: {
			int(columnnames.NodeColumnNameIndex),
			int(columnnames.NodeColumnNameHealth),
			int(columnnames.NodeColumnNameAddress),
			int(columnnames.NodeColumnNamePortRPC),
			int(columnnames.NodeColumnNameTotalTransactions),
			int(columnnames.NodeColumnNameLatestCheckpoint),
		},
	}
)
