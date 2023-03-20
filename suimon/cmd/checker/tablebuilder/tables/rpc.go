package tables

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
)

var (
	TableStyleRPC      = table.StyleLight
	TableTagRPC        = ""
	TableSortConfigRPC = []table.SortBy{
		{Name: "HEALTH", Mode: table.Dsc},
		{Name: "TOTAL\nTRANSACTIONS", Mode: table.Dsc},
		{Name: "LATEST\nCHECKPOINT", Mode: table.Dsc},
	}
	ColumnConfigRPC = []table.ColumnConfig{
		enums.ColumnNameHealth: {
			Name:         "HEALTH",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		enums.ColumnNameAddress: {
			Name:         "ADDRESS",
			Align:        text.AlignLeft,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		enums.ColumnNamePortRPC: {
			Name:         "RPC",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		enums.ColumnNameTotalTransactions: {
			Name:         "TOTAL\nTRANSACTIONS",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		enums.ColumnNameLatestCheckpoint: {
			Name:         "LATEST\nCHECKPOINT",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
	}
)
