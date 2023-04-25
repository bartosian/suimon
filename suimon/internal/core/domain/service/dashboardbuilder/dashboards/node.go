package dashboards

import (
	"fmt"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/host"
)

var (
	ColumnsConfigNode = ColumnsConfig{
		// Overview section
		enums.ColumnNameCurrentEpoch:            19,
		enums.ColumnNameSystemTimeTillNextEpoch: 25,
		enums.ColumnNameNetworkPeers:            15,
		enums.ColumnNameUptime:                  25,
		enums.ColumnNameVersion:                 25,
		enums.ColumnNameCommit:                  25,
		enums.ColumnNameHealth:                  5,
		enums.ColumnNameCheckpointExecBacklog:   24,
		enums.ColumnNameCheckpointSyncBacklog:   24,

		// Transactions section
		enums.ColumnNameTotalTransactionBlocks:       33,
		enums.ColumnNameTotalTransactionCertificates: 33,
		enums.ColumnNameTotalTransactionEffects:      33,
		enums.ColumnNameTXSyncPercentage:             49,
		enums.ColumnNameTransactionsPerSecond:        49,

		// Checkpoints section
		enums.ColumnNameLatestCheckpoint:        24,
		enums.ColumnNameHighestKnownCheckpoint:  24,
		enums.ColumnNameHighestSyncedCheckpoint: 24,
		enums.ColumnNameLastExecutedCheckpoint:  24,
		enums.ColumnNameCheckSyncPercentage:     49,
		enums.ColumnNameCheckpointsPerSecond:    49,
	}

	RowsConfigNode = RowsConfig{
		0: {
			Height: 14,
			Columns: []enums.ColumnName{
				enums.ColumnNameNetworkPeers,
				enums.ColumnNameUptime,
				enums.ColumnNameVersion,
				enums.ColumnNameCommit,
			},
		},
		1: {
			Height: 14,
			Columns: []enums.ColumnName{
				enums.ColumnNameCurrentEpoch,
				enums.ColumnNameSystemTimeTillNextEpoch,
				enums.ColumnNameCheckpointExecBacklog,
				enums.ColumnNameCheckpointSyncBacklog,
			},
		},
		2: {
			Height: 14,
			Columns: []enums.ColumnName{
				enums.ColumnNameTotalTransactionBlocks,
				enums.ColumnNameTotalTransactionCertificates,
				enums.ColumnNameTotalTransactionEffects,
			},
		},
		3: {
			Height: 14,
			Columns: []enums.ColumnName{
				enums.ColumnNameTXSyncPercentage,
				enums.ColumnNameTransactionsPerSecond,
			},
		},
		4: {
			Height: 14,
			Columns: []enums.ColumnName{
				enums.ColumnNameLatestCheckpoint,
				enums.ColumnNameHighestKnownCheckpoint,
				enums.ColumnNameHighestSyncedCheckpoint,
				enums.ColumnNameLastExecutedCheckpoint,
			},
		},
		5: {
			Height: 14,
			Columns: []enums.ColumnName{
				enums.ColumnNameCheckSyncPercentage,
				enums.ColumnNameCheckpointsPerSecond,
			},
		},
	}

	CellsConfigNode = CellsConfig{
		enums.ColumnNameHealth:                       "HEALTH",
		enums.ColumnNameTotalTransactionBlocks:       "TOTAL TRANSACTION BLOCKS",
		enums.ColumnNameLatestCheckpoint:             "LATEST CHECKPOINT",
		enums.ColumnNameTotalTransactionCertificates: "TOTAL TRANSACTION CERTIFICATES",
		enums.ColumnNameTotalTransactionEffects:      "TOTAL TRANSACTION EFFECTS",
		enums.ColumnNameTransactionsPerSecond:        "TRANSACTIONS PER SECOND",
		enums.ColumnNameHighestKnownCheckpoint:       "HIGHEST KNOWN CHECKPOINT",
		enums.ColumnNameHighestSyncedCheckpoint:      "HIGHEST SYNCED CHECKPOINT",
		enums.ColumnNameLastExecutedCheckpoint:       "LAST EXECUTED CHECKPOINT",
		enums.ColumnNameCheckpointExecBacklog:        "CHECKPOINT EXEC BACKLOG",
		enums.ColumnNameCheckpointSyncBacklog:        "CHECKPOINT SYNC BACKLOG",
		enums.ColumnNameCurrentEpoch:                 "CURRENT EPOCH",
		enums.ColumnNameSystemTimeTillNextEpoch:      "TIME TILL NEXT EPOCH",
		enums.ColumnNameTXSyncPercentage:             "TX SYNC PERCENTAGE",
		enums.ColumnNameCheckSyncPercentage:          "CHECKPOINTS SYNC PERCENTAGE",
		enums.ColumnNameCheckpointsPerSecond:         "CHECKPOINTS PER SECOND",
		enums.ColumnNameNetworkPeers:                 "NETWORK PEERS",
		enums.ColumnNameUptime:                       "UPTIME",
		enums.ColumnNameVersion:                      "VERSION",
		enums.ColumnNameCommit:                       "COMMIT",
	}
)

// GetNodeColumnValues returns a map of NodeColumnName values to corresponding values for a node at the specified index on the specified host.
// The function retrieves information about the node from the host's internal state and formats it into a map of NodeColumnName keys and corresponding values.
// The function also includes emoji values in the map if the specified flag is true.
// Returns a map of NodeColumnName keys to corresponding values.
func GetNodeColumnValues(host host.Host) ColumnValues {
	status := host.Status.StatusToPlaceholder()

	columnValues := ColumnValues{
		enums.ColumnNameHealth:                       status,
		enums.ColumnNameTotalTransactionBlocks:       host.Metrics.TotalTransactionsBlocks,
		enums.ColumnNameTotalTransactionCertificates: host.Metrics.TotalTransactionCertificates,
		enums.ColumnNameTotalTransactionEffects:      host.Metrics.TotalTransactionEffects,
		enums.ColumnNameTransactionsPerSecond:        host.Metrics.TransactionsPerSecond,
		enums.ColumnNameLatestCheckpoint:             host.Metrics.LatestCheckpoint,
		enums.ColumnNameHighestKnownCheckpoint:       host.Metrics.HighestKnownCheckpoint,
		enums.ColumnNameHighestSyncedCheckpoint:      host.Metrics.HighestSyncedCheckpoint,
		enums.ColumnNameLastExecutedCheckpoint:       host.Metrics.LastExecutedCheckpoint,
		enums.ColumnNameCheckpointExecBacklog:        host.Metrics.CheckpointExecBacklog,
		enums.ColumnNameCheckpointSyncBacklog:        host.Metrics.CheckpointSyncBacklog,
		enums.ColumnNameCurrentEpoch:                 host.Metrics.CurrentEpoch,
		enums.ColumnNameSystemTimeTillNextEpoch:      host.Metrics.DurationTillEpochEndHHMM,
		enums.ColumnNameTXSyncPercentage:             fmt.Sprintf("%v%%", host.Metrics.TxSyncPercentage),
		enums.ColumnNameCheckSyncPercentage:          fmt.Sprintf("%v%%", host.Metrics.CheckSyncPercentage),
		enums.ColumnNameCheckpointsPerSecond:         host.Metrics.CheckpointsPerSecond,
		enums.ColumnNameNetworkPeers:                 host.Metrics.NetworkPeers,
		enums.ColumnNameUptime:                       host.Metrics.Uptime,
		enums.ColumnNameVersion:                      host.Metrics.Version,
		enums.ColumnNameCommit:                       host.Metrics.Commit,
	}

	return columnValues
}
