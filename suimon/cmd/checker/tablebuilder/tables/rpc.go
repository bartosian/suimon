package tables

import (
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums/columnnames"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
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
)
