package dashboards

import (
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
)

var (
	ColumnsConfigNode = ColumnsConfig{
		enums.ColumnNameHealth:                       5,
		enums.ColumnNameAddress:                      5,
		enums.ColumnNamePortRPC:                      20,
		enums.ColumnNameTotalTransactionBlocks:       20,
		enums.ColumnNameLatestCheckpoint:             20,
		enums.ColumnNameTotalTransactionCertificates: 20,
		enums.ColumnNameTotalTransactionEffects:      20,
		enums.ColumnNameHighestKnownCheckpoint:       30,
		enums.ColumnNameHighestSyncedCheckpoint:      30,
		enums.ColumnNameLastExecutedCheckpoint:       30,
		enums.ColumnNameCheckpointExecBacklog:        25,
		enums.ColumnNameCheckpointSyncBacklog:        25,
		enums.ColumnNameCurrentEpoch:                 25,
		enums.ColumnNameTXSyncPercentage:             25,
		enums.ColumnNameCheckSyncPercentage:          15,
		enums.ColumnNameNetworkPeers:                 15,
		enums.ColumnNameUptime:                       15,
		enums.ColumnNameVersion:                      15,
		enums.ColumnNameCommit:                       10,
		enums.ColumnNameCountry:                      50,
	}

	RowsConfigNode = RowsConfig{
		0: {
			Height: 15,
			Columns: []enums.ColumnName{
				enums.ColumnNameHealth,
				enums.ColumnNameAddress,
				enums.ColumnNamePortRPC,
				enums.ColumnNameTotalTransactionBlocks,
				enums.ColumnNameLatestCheckpoint,
				enums.ColumnNameTotalTransactionCertificates,
				enums.ColumnNameTotalTransactionEffects,
				enums.ColumnNameHighestKnownCheckpoint,
			},
		},
		1: {
			Height: 15,
			Columns: []enums.ColumnName{
				enums.ColumnNameHighestSyncedCheckpoint,
				enums.ColumnNameLastExecutedCheckpoint,
				enums.ColumnNameCheckpointExecBacklog,
				enums.ColumnNameCheckpointSyncBacklog,
				enums.ColumnNameCurrentEpoch,
				enums.ColumnNameTXSyncPercentage,
				enums.ColumnNameCheckSyncPercentage,
				enums.ColumnNameNetworkPeers,
			},
		},
		2: {
			Height: 15,
			Columns: []enums.ColumnName{
				enums.ColumnNameUptime,
				enums.ColumnNameVersion,
				enums.ColumnNameCommit,
				enums.ColumnNameCountry,
			},
		},
	}

	CellsConfigNode = CellsConfig{
		enums.ColumnNameHealth:                       "HEALTH",
		enums.ColumnNameAddress:                      "ADDRESS",
		enums.ColumnNamePortRPC:                      "RPC PORT",
		enums.ColumnNameTotalTransactionBlocks:       "TOTAL TRANSACTION BLOCKS",
		enums.ColumnNameLatestCheckpoint:             "LATEST CHECKPOINT",
		enums.ColumnNameTotalTransactionCertificates: "TOTAL TRANSACTION CERTIFICATES",
		enums.ColumnNameTotalTransactionEffects:      "TOTAL TRANSACTION EFFECTS",
		enums.ColumnNameHighestKnownCheckpoint:       "HIGHEST KNOWN CHECKPOINT",
		enums.ColumnNameHighestSyncedCheckpoint:      "HIGHEST SYNCED CHECKPOINT",
		enums.ColumnNameLastExecutedCheckpoint:       "LAST EXECUTED CHECKPOINT",
		enums.ColumnNameCheckpointExecBacklog:        "CHECKPOINT EXEC BACKLOG",
		enums.ColumnNameCheckpointSyncBacklog:        "CHECKPOINT SYNC BACKLOG",
		enums.ColumnNameCurrentEpoch:                 "CURRENT EPOCH",
		enums.ColumnNameTXSyncPercentage:             "TX SYNC PERCENTAGE",
		enums.ColumnNameCheckSyncPercentage:          "CHECKPOINTS SYNC PERCENTAGE",
		enums.ColumnNameNetworkPeers:                 "NETWORK PEERS",
		enums.ColumnNameUptime:                       "UPTIME",
		enums.ColumnNameVersion:                      "VERSION",
		enums.ColumnNameCommit:                       "COMMIT",
		enums.ColumnNameCountry:                      "COUNTRY",
	}
)
