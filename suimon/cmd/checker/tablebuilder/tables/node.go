package tables

import (
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums/columnnames"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

var (
	TableStyleNode      = table.StyleLight
	TableTagNode        = ""
	TableSortConfigNode = []table.SortBy{
		{Name: "HEALTH", Mode: table.Dsc},
		{Name: "TOTAL\nTRANSACTIONS", Mode: table.Dsc},
		{Name: "LATEST\nCHECKPOINT", Mode: table.Dsc},
	}
	ColumnsConfigNode = []table.ColumnConfig{
		columnnames.NodeColumnNameIndex: {
			Name:         "IDX",
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
			Align:        text.AlignCenter,
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
		columnnames.NodeColumnNameTotalTransactionCertificates: {
			Name:         "TOTAL TRANSACTION\nCERTIFICATES",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.NodeColumnNameTotalTransactionEffects: {
			Name:         "TOTAL TRANSACTION\nEFFECTS",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.NodeColumnNameHighestKnownCheckpoint: {
			Name:         "HIGHEST KNOWN\nCHECKPOINT",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.NodeColumnNameHighestSyncedCheckpoint: {
			Name:         "HIGHEST SYNCED\nCHECKPOINT",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.NodeColumnNameLastExecutedCheckpoint: {
			Name:         "LAST EXECUTED\nCHECKPOINT",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.NodeColumnNameCheckpointExecBacklog: {
			Name:         "CHECKPOINT EXEC\nBACKLOG",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.NodeColumnNameCheckpointSyncBacklog: {
			Name:         "CHECKPOINT SYNC\nBACKLOG",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.NodeColumnNameCurrentEpoch: {
			Name:         "CURRENT\nEPOCH",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.NodeColumnNameTXSyncPercentage: {
			Name:         "TRANSACTIONS\nSYNC PERCENTAGE",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.NodeColumnNameCheckSyncPercentage: {
			Name:         "CHECKPOINTS\nSYNC PERCENTAGE",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.NodeColumnNameNetworkPeers: {
			Name:         "NETWORK\nPEERS",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.NodeColumnNameUptime: {
			Name:         "UPTIME DAYS",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.NodeColumnNameVersion: {
			Name:         "VERSION",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.NodeColumnNameCommit: {
			Name:         "COMMIT",
			Align:        text.AlignCenter,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
		columnnames.NodeColumnNameCountry: {
			Name:         "COUNTRY",
			Align:        text.AlignLeft,
			AlignHeader:  text.AlignCenter,
			VAlign:       text.VAlignMiddle,
			VAlignHeader: text.VAlignMiddle,
			Hidden:       false,
		},
	}
	RowsNode = [][]int{
		0: {
			int(columnnames.NodeColumnNameIndex),
			int(columnnames.NodeColumnNameHealth),
			int(columnnames.NodeColumnNameAddress),
			int(columnnames.NodeColumnNamePortRPC),
			int(columnnames.NodeColumnNameTotalTransactions),
			int(columnnames.NodeColumnNameLatestCheckpoint),
			int(columnnames.NodeColumnNameTotalTransactionCertificates),
			int(columnnames.NodeColumnNameTotalTransactionEffects),
			int(columnnames.NodeColumnNameHighestKnownCheckpoint),
			int(columnnames.NodeColumnNameLastExecutedCheckpoint),
			int(columnnames.NodeColumnNameCheckpointExecBacklog),
			int(columnnames.NodeColumnNameHighestSyncedCheckpoint),
			int(columnnames.NodeColumnNameCheckpointSyncBacklog),
		},
		1: {
			int(columnnames.NodeColumnNameCurrentEpoch),
			int(columnnames.NodeColumnNameTXSyncPercentage),
			int(columnnames.NodeColumnNameCheckSyncPercentage),
			int(columnnames.NodeColumnNameNetworkPeers),
			int(columnnames.NodeColumnNameUptime),
			int(columnnames.NodeColumnNameVersion),
			int(columnnames.NodeColumnNameCommit),
			int(columnnames.NodeColumnNameCountry),
		},
	}
)
