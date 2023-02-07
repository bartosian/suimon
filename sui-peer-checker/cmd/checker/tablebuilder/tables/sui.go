package tables

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

type ColumnNameSUI int

const (
	ColumnNameSUIIDX ColumnNameSUI = iota
	ColumnNameSUIPeer
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
	TableTitleSUI   = "💧 SUI PEERS CHECKER v0.1.0"
	TableStyleSUI   = table.StyleLight
	ColumnConfigSUI = [...]table.ColumnConfig{
		ColumnNameSUIIDX: {
			Name:         "#",
			Align:        text.AlignCenter,
			Colors:       text.Colors{text.FgHiRed, text.Bold},
			ColorsHeader: text.Colors{text.FgHiRed},
			ColorsFooter: text.Colors{text.FgHiRed},
			Hidden:       false,
			Transformer:  nameTransformer,
		},
		ColumnNameSUIPeer: {
			Name:        "PEER",
			Align:       text.AlignLeft,
			AlignHeader: text.AlignCenter,
			Hidden:      false,
		},
		ColumnNameSUIPort: {
			Name:        "PORT",
			Align:       text.AlignCenter,
			AlignHeader: text.AlignCenter,
			Hidden:      false,
		},
		ColumnNameSUITotalTransactions: {
			Name:        "TOTAL\nTRANSACTIONS",
			Align:       text.AlignCenter,
			AlignHeader: text.AlignCenter,
			Hidden:      false,
		},
		ColumnNameSUIHighestCheckpoints: {
			Name:        "HIGHEST\nCHECKPOINTS",
			Align:       text.AlignCenter,
			AlignHeader: text.AlignCenter,
			Hidden:      false,
		},
		ColumnNameSUIConnectedPeers: {
			Name:        "CONNECTED\nPEERS",
			Align:       text.AlignCenter,
			AlignHeader: text.AlignCenter,
			Hidden:      false,
		},
		ColumnNameSUIUptime: {
			Name:        "UPTIME",
			Align:       text.AlignCenter,
			AlignHeader: text.AlignCenter,
			Hidden:      false,
		},
		ColumnNameSUIVersion: {
			Name:        "VERSION",
			Align:       text.AlignCenter,
			AlignHeader: text.AlignCenter,
			Hidden:      false,
		},
		ColumnNameSUICommit: {
			Name:        "COMMIT",
			Align:       text.AlignCenter,
			AlignHeader: text.AlignCenter,
			Hidden:      false,
		},
		ColumnNameSUICountry: {
			Name:        "COUNTRY",
			Align:       text.AlignLeft,
			AlignHeader: text.AlignCenter,
			Hidden:      false,
		},
	}
)
