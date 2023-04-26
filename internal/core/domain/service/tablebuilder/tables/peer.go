package tables

import (
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
)

var (
	ColumnsConfigPeer = ColumnsConfig{
		enums.ColumnNameIndex:                        NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameHealth:                       NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameAddress:                      NewDefaultColumnConfig(text.AlignLeft, text.AlignCenter, false),
		enums.ColumnNamePortRPC:                      NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameTotalTransactionBlocks:       NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameLatestCheckpoint:             NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameTotalTransactionCertificates: NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameTotalTransactionEffects:      NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameHighestKnownCheckpoint:       NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameHighestSyncedCheckpoint:      NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameLastExecutedCheckpoint:       NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameCheckpointExecBacklog:        NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameCheckpointSyncBacklog:        NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameCurrentEpoch:                 NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameTXSyncPercentage:             NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameCheckSyncPercentage:          NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameNetworkPeers:                 NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameUptime:                       NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameVersion:                      NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameCommit:                       NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameCountry:                      NewDefaultColumnConfig(text.AlignLeft, text.AlignCenter, false),
	}
	
	RowsConfigPeer = RowsConfig{
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
