package tables

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

type ColumnNameSUINode int

const (
	ColumnNameSUINodeAddress ColumnNameSUINode = iota
	ColumnNameSUINodePortRPC
	ColumnNameSUINodePortMetrics
	ColumnNameSUINodeTotalTransactions
	ColumnNameSUINodeHighestCheckpoints
	ColumnNameSUINodeConnectedPeers
	ColumnNameSUINodeUptime
	ColumnNameSUINodeVersion
	ColumnNameSUINodeCommit
	ColumnNameSUINodeCountry
)

var (
	TableStyleSUINode      = table.StyleLight
	TableTagSUINode        = ""
	TableSortConfigSUINode = []table.SortBy{
		{Name: "COUNTRY", Mode: table.Asc},
		{Name: "UPTIME", Mode: table.Asc},
	}
	ColumnConfigSUINode = [...]table.ColumnConfig{
		ColumnNameSUINodeAddress: {
			Name:         "ADDRESS",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		ColumnNameSUINodePortRPC: {
			Name:         "RPC",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		ColumnNameSUINodePortMetrics: {
			Name:         "Metrics",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		ColumnNameSUINodeTotalTransactions: {
			Name:         "TOTAL\nTRANSACTIONS",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		ColumnNameSUINodeHighestCheckpoints: {
			Name:         "HIGHEST\nCHECKPOINTS",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		ColumnNameSUINodeConnectedPeers: {
			Name:         "CONNECTED\nPEERS",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		ColumnNameSUINodeUptime: {
			Name:         "UPTIME",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		ColumnNameSUINodeVersion: {
			Name:         "VERSION",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		ColumnNameSUINodeCommit: {
			Name:         "COMMIT",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		ColumnNameSUINodeCountry: {
			Name:         "COUNTRY",
			Align:        text.AlignLeft,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
	}
)
