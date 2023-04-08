package tables

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/service/tablebuilder"
)

var (
	SortConfigPeer = tablebuilder.SortConfig{
		{Name: enums.ColumnNameHealth.ToString(), Mode: table.Dsc},
		{Name: enums.ColumnNameTotalTransactionBlocks.ToString(), Mode: table.Dsc},
		{Name: enums.ColumnNameLatestCheckpoint.ToString(), Mode: table.Dsc},
	}
	ColumnsConfigPeer = tablebuilder.ColumnsConfig{
		enums.ColumnNameIndex:                        tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameHealth:                       tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameAddress:                      tablebuilder.NewDefaultColumnConfig(text.AlignLeft, text.AlignCenter, false),
		enums.ColumnNamePortRPC:                      tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameTotalTransactionBlocks:       tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameLatestCheckpoint:             tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameTotalTransactionCertificates: tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameTotalTransactionEffects:      tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameHighestKnownCheckpoint:       tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameHighestSyncedCheckpoint:      tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameLastExecutedCheckpoint:       tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameCheckpointExecBacklog:        tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameCheckpointSyncBacklog:        tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameCurrentEpoch:                 tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameTXSyncPercentage:             tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameCheckSyncPercentage:          tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameNetworkPeers:                 tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameUptime:                       tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameVersion:                      tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameCommit:                       tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameCountry:                      tablebuilder.NewDefaultColumnConfig(text.AlignLeft, text.AlignCenter, false),
	}
	RowsConfigPeer = tablebuilder.RowsConfig{
		0: {
			enums.ColumnNameIndex,
			enums.ColumnNameHealth,
			enums.ColumnNameAddress,
			enums.ColumnNamePortRPC,
			enums.ColumnNameTotalTransactionBlocks,
			enums.ColumnNameLatestCheckpoint,
			enums.ColumnNameTotalTransactionCertificates,
			enums.ColumnNameTotalTransactionEffects,
			enums.ColumnNameHighestKnownCheckpoint,
			enums.ColumnNameLastExecutedCheckpoint,
			enums.ColumnNameCheckpointExecBacklog,
			enums.ColumnNameHighestSyncedCheckpoint,
			enums.ColumnNameCheckpointSyncBacklog,
		},
		1: {
			enums.ColumnNameCurrentEpoch,
			enums.ColumnNameTXSyncPercentage,
			enums.ColumnNameCheckSyncPercentage,
			enums.ColumnNameNetworkPeers,
			enums.ColumnNameUptime,
			enums.ColumnNameVersion,
			enums.ColumnNameCommit,
			enums.ColumnNameCountry,
		},
	}
)
