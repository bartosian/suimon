package tables

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

type ColumnNameSUI int

const (
	ColumnNameSUIPeer ColumnNameSUI = iota
	ColumnNameSUIPort
	ColumnNameSUITotalTransactions
	ColumnNameSUIHighestCheckpoints
	ColumnNameSUIConnectedPeers
	ColumnNameSUIUptime
	ColumnNameSUIVersion
	ColumnNameSUICommit
	ColumnNameSUICountry
)

var nameTransformer = text.Transformer(func(val interface{}) string {
	return text.Bold.Sprint(val)
})

var (
	TableTitleSUI      = "💧 SUI PEERS CHECKER v0.1.0"
	TableStyleSUI      = table.StyleLight
	TableSortConfigSUI = []table.SortBy{
		{Name: "PEER", Mode: table.Asc},
		{Name: "UPTIME", Mode: table.Asc},
	}
	ColumnConfigSUI = [...]table.ColumnConfig{
		ColumnNameSUIPeer: {
			Name:         "PEER",
			Align:        text.AlignLeft,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		ColumnNameSUIPort: {
			Name:         "PORT",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		ColumnNameSUITotalTransactions: {
			Name:         "TOTAL\nTRANSACTIONS",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		ColumnNameSUIHighestCheckpoints: {
			Name:         "HIGHEST\nCHECKPOINTS",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		ColumnNameSUIConnectedPeers: {
			Name:         "CONNECTED\nPEERS",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		ColumnNameSUIUptime: {
			Name:         "UPTIME",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		ColumnNameSUIVersion: {
			Name:         "VERSION",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		ColumnNameSUICommit: {
			Name:         "COMMIT",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		ColumnNameSUICountry: {
			Name:         "COUNTRY",
			Align:        text.AlignLeft,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
	}
)
