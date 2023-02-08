package tables

import (
	"fmt"
	"github.com/bartosian/sui_helpers/sui-peer-checker/cmd/checker/enums"
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

func GetTableTitleSUI(network enums.NetworkType, table enums.TableType) string {
	switch network {
	case enums.NetworkTypeTestnet:
		return fmt.Sprintf("💧 SUI PEERS CHECKER %sv0.1.0%s - %s%s %s%s", enums.ColorGreen, enums.ColorReset, enums.ColorRed, table, network, enums.ColorReset)
	case enums.NetworkTypeDevnet:
		fallthrough
	default:
		return fmt.Sprintf("💧 SUI PEERS CHECKER %sv0.1.0%s - %s%s %s%s", enums.ColorGreen, enums.ColorReset, enums.ColorRed, table, network, enums.ColorReset)
	}
}

var (
	TableStyleSUI      = table.StyleLight
	TableTagSUI        = "BartestneT 2023"
	TableSortConfigSUI = []table.SortBy{
		{Name: "COUNTRY", Mode: table.Asc},
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
