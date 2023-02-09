package tables

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

type ColumnNameSUISystem int

const (
	ColumnNameSUISystemStatus ColumnNameSUI = iota
	ColumnNameSUISystemRPC
	ColumnNameSUISystemTotalTransactions
	ColumnNameSUISystemLatestCheckpoint
)

var (
	TableStyleSystemSUI      = table.StyleLight
	TableTagSystemSUI        = ""
	TableSortConfigSystemSUI = []table.SortBy{
		{Name: "STATUS", Mode: table.Dsc},
		{Name: "HIGHEST\nCHECKPOINTS", Mode: table.Asc},
	}
	ColumnConfigSystemSUI = [...]table.ColumnConfig{
		ColumnNameSUISystemStatus: {
			Name:         "STATUS",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		ColumnNameSUISystemRPC: {
			Name:         "RPC ADDRESS",
			Align:        text.AlignLeft,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		ColumnNameSUISystemTotalTransactions: {
			Name:         "TOTAL\nTRANSACTIONS",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		ColumnNameSUISystemLatestCheckpoint: {
			Name:         "LATEST\nCHECKPOINT",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
	}
)
