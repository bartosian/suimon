package dashboards

import (
	"fmt"

	"github.com/mum4k/termdash/cell"

	"github.com/bartosian/suimon/internal/core/domain/enums"
	domainhost "github.com/bartosian/suimon/internal/core/domain/host"
)

var (
	ColumnsConfigNode = ColumnsConfig{
		// Overview section
		enums.ColumnNameCurrentEpoch:          ColumnWidth33,
		enums.ColumnNameNetworkPeers:          ColumnWidth15,
		enums.ColumnNameUptime:                ColumnWidth25,
		enums.ColumnNameVersion:               ColumnWidth25,
		enums.ColumnNameCommit:                ColumnWidth25,
		enums.ColumnNameCheckpointExecBacklog: ColumnWidth33,
		enums.ColumnNameCheckpointSyncBacklog: ColumnWidth33,

		// Transactions section
		enums.ColumnNameTotalTransactionBlocks:       ColumnWidth33,
		enums.ColumnNameTotalTransactionCertificates: ColumnWidth33,
		enums.ColumnNameTotalTransactionEffects:      ColumnWidth33,
		enums.ColumnNameTXSyncPercentage:             ColumnWidth49,
		enums.ColumnNameTransactionsPerSecond:        ColumnWidth49,

		// Checkpoints section
		enums.ColumnNameLatestCheckpoint:        ColumnWidth24,
		enums.ColumnNameHighestKnownCheckpoint:  ColumnWidth24,
		enums.ColumnNameHighestSyncedCheckpoint: ColumnWidth24,
		enums.ColumnNameLastExecutedCheckpoint:  ColumnWidth24,
		enums.ColumnNameCheckSyncPercentage:     ColumnWidth49,
		enums.ColumnNameCheckpointsPerSecond:    ColumnWidth49,
	}

	RowsConfigNode = RowsConfig{
		0: {
			Height: RowHeight14,
			Columns: []enums.ColumnName{
				enums.ColumnNameNetworkPeers,
				enums.ColumnNameUptime,
				enums.ColumnNameVersion,
				enums.ColumnNameCommit,
			},
		},
		1: {
			Height: RowHeight14,
			Columns: []enums.ColumnName{
				enums.ColumnNameCurrentEpoch,
				enums.ColumnNameCheckpointExecBacklog,
				enums.ColumnNameCheckpointSyncBacklog,
			},
		},
		2: {
			Height: RowHeight14,
			Columns: []enums.ColumnName{
				enums.ColumnNameLatestCheckpoint,
				enums.ColumnNameHighestKnownCheckpoint,
				enums.ColumnNameHighestSyncedCheckpoint,
				enums.ColumnNameLastExecutedCheckpoint,
			},
		},
		3: {
			Height: RowHeight14,
			Columns: []enums.ColumnName{
				enums.ColumnNameCheckSyncPercentage,
				enums.ColumnNameCheckpointsPerSecond,
			},
		},
		4: {
			Height: RowHeight14,
			Columns: []enums.ColumnName{
				enums.ColumnNameTotalTransactionBlocks,
				enums.ColumnNameTotalTransactionCertificates,
				enums.ColumnNameTotalTransactionEffects,
			},
		},
		5: {
			Height: RowHeight14,
			Columns: []enums.ColumnName{
				enums.ColumnNameTXSyncPercentage,
				enums.ColumnNameTransactionsPerSecond,
			},
		},
	}

	CellsConfigNode = CellsConfig{
		enums.ColumnNameNetworkPeers:                 {"NETWORK PEERS", cell.ColorGreen},
		enums.ColumnNameUptime:                       {"UPTIME", cell.ColorGreen},
		enums.ColumnNameVersion:                      {"VERSION", cell.ColorGreen},
		enums.ColumnNameCommit:                       {"COMMIT", cell.ColorGreen},
		enums.ColumnNameCurrentEpoch:                 {"CURRENT EPOCH", cell.ColorGreen},
		enums.ColumnNameCheckpointExecBacklog:        {"CHECKPOINT EXEC BACKLOG", cell.ColorGreen},
		enums.ColumnNameCheckpointSyncBacklog:        {"CHECKPOINT SYNC BACKLOG", cell.ColorGreen},
		enums.ColumnNameHighestKnownCheckpoint:       {"HIGHEST KNOWN CHECKPOINT", cell.ColorBlue},
		enums.ColumnNameHighestSyncedCheckpoint:      {"HIGHEST SYNCED CHECKPOINT", cell.ColorBlue},
		enums.ColumnNameLastExecutedCheckpoint:       {"LAST EXECUTED CHECKPOINT", cell.ColorBlue},
		enums.ColumnNameLatestCheckpoint:             {"LATEST CHECKPOINT", cell.ColorBlue},
		enums.ColumnNameCheckSyncPercentage:          {"CHECKPOINTS SYNC PERCENTAGE", cell.ColorBlue},
		enums.ColumnNameCheckpointsPerSecond:         {"CHECKPOINTS VOLUME", cell.ColorBlue},
		enums.ColumnNameTotalTransactionBlocks:       {"TOTAL TRANSACTION BLOCKS", cell.ColorYellow},
		enums.ColumnNameTotalTransactionCertificates: {"TOTAL TRANSACTION CERTIFICATES", cell.ColorYellow},
		enums.ColumnNameTotalTransactionEffects:      {"TOTAL TRANSACTION EFFECTS", cell.ColorYellow},
		enums.ColumnNameTransactionsPerSecond:        {"TRANSACTIONS VOLUME", cell.ColorYellow},
		enums.ColumnNameTXSyncPercentage:             {"TRANSACTIONS SYNC PERCENTAGE", cell.ColorYellow},
	}
)

// GetNodeColumnValues returns a map of ColumnName values to corresponding values for a node at the specified index on the specified host.
// The function retrieves information about the node from the host's internal state and formats it into a map of NodeColumnName keys and corresponding values.
// The function also includes emoji values in the map if the specified flag is true.
func GetNodeColumnValues(host *domainhost.Host) (ColumnValues, error) {
	return ColumnValues{
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
		enums.ColumnNameTXSyncPercentage:             fmt.Sprintf("%v%%", host.Metrics.TxSyncPercentage),
		enums.ColumnNameCheckSyncPercentage:          fmt.Sprintf("%v%%", host.Metrics.CheckSyncPercentage),
		enums.ColumnNameCheckpointsPerSecond:         host.Metrics.CheckpointsPerSecond,
		enums.ColumnNameNetworkPeers:                 host.Metrics.NetworkPeers,
		enums.ColumnNameUptime:                       host.Metrics.Uptime,
		enums.ColumnNameVersion:                      host.Metrics.Version,
		enums.ColumnNameCommit:                       host.Metrics.Commit,
	}, nil
}
