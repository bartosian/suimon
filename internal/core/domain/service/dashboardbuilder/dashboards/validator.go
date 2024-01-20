package dashboards

import (
	"fmt"

	"github.com/mum4k/termdash/cell"

	"github.com/bartosian/suimon/internal/core/domain/enums"
	"github.com/bartosian/suimon/internal/core/domain/host"
)

var (
	ColumnsConfigValidator = ColumnsConfig{
		// Overview section
		enums.ColumnNameCurrentEpoch: 20,
		enums.ColumnNameUptime:       25,
		enums.ColumnNameVersion:      25,
		enums.ColumnNameCommit:       25,

		// Transactions section
		enums.ColumnNameTotalTransactionCertificates: 33,
		enums.ColumnNameTotalTransactionEffects:      33,
		enums.ColumnNameCertificatesCreated:          33,
		enums.ColumnNameCertificatesPerSecond:        49,

		// Checkpoints section
		enums.ColumnNameLastExecutedCheckpoint:  33,
		enums.ColumnNameHighestKnownCheckpoint:  33,
		enums.ColumnNameHighestSyncedCheckpoint: 33,
		enums.ColumnNameCheckSyncPercentage:     49,
		enums.ColumnNameCheckpointsPerSecond:    49,

		// Rounds section
		enums.ColumnNameCurrentRound:          33,
		enums.ColumnNameHighestProcessedRound: 33,
		enums.ColumnNameLastCommittedRound:    33,
		enums.ColumnNameRoundsPerSecond:       49,

		// Peers section
		enums.ColumnNameNetworkPeers:        19,
		enums.ColumnNamePrimaryNetworkPeers: 19,
		enums.ColumnNameWorkerNetworkPeers:  19,

		// Performance section
		enums.ColumnNameSkippedConsensusTransactions: 19,
		enums.ColumnNameTotalSignatureErrors:         19,
	}

	RowsConfigValidator = RowsConfig{
		0: {
			Height: 14,
			Columns: []enums.ColumnName{
				enums.ColumnNameCurrentEpoch,
				enums.ColumnNameUptime,
				enums.ColumnNameVersion,
				enums.ColumnNameCommit,
			},
		},
		1: {
			Height: 14,
			Columns: []enums.ColumnName{
				enums.ColumnNameNetworkPeers,
				enums.ColumnNamePrimaryNetworkPeers,
				enums.ColumnNameWorkerNetworkPeers,
				enums.ColumnNameSkippedConsensusTransactions,
				enums.ColumnNameTotalSignatureErrors,
			},
		},
		2: {
			Height: 14,
			Columns: []enums.ColumnName{
				enums.ColumnNameLastExecutedCheckpoint,
				enums.ColumnNameHighestKnownCheckpoint,
				enums.ColumnNameHighestSyncedCheckpoint,
			},
		},
		3: {
			Height: 14,
			Columns: []enums.ColumnName{
				enums.ColumnNameCheckSyncPercentage,
				enums.ColumnNameCheckpointsPerSecond,
			},
		},
		4: {
			Height: 14,
			Columns: []enums.ColumnName{
				enums.ColumnNameTotalTransactionCertificates,
				enums.ColumnNameTotalTransactionEffects,
				enums.ColumnNameCertificatesCreated,
			},
		},
		5: {
			Height: 14,
			Columns: []enums.ColumnName{
				enums.ColumnNameCurrentRound,
				enums.ColumnNameHighestProcessedRound,
				enums.ColumnNameLastCommittedRound,
			},
		},
		6: {
			Height: 14,
			Columns: []enums.ColumnName{
				enums.ColumnNameCertificatesPerSecond,
				enums.ColumnNameRoundsPerSecond,
			},
		},
	}

	CellsConfigValidator = CellsConfig{
		enums.ColumnNameCurrentEpoch:                            {"CURRENT EPOCH", cell.ColorGreen},
		enums.ColumnNameUptime:                                  {"UPTIME", cell.ColorGreen},
		enums.ColumnNameVersion:                                 {"VERSION", cell.ColorGreen},
		enums.ColumnNameCommit:                                  {"COMMIT", cell.ColorGreen},
		enums.ColumnNameNetworkPeers:                            {"SUI NETWORK PEERS", cell.ColorGreen},
		enums.ColumnNamePrimaryNetworkPeers:                     {"PRIMARY NETWORK PEERS", cell.ColorGreen},
		enums.ColumnNameWorkerNetworkPeers:                      {"WORKER NETWORK PEERS", cell.ColorGreen},
		enums.ColumnNameSkippedConsensusTransactions:            {"SKIPPED CONSENSUS TRANSACTIONS", cell.ColorGreen},
		enums.ColumnNameTotalSignatureErrors:                    {"TOTAL SIGNATURE ERRORS", cell.ColorGreen},
		enums.ColumnNameHighestKnownCheckpoint:                  {"HIGHEST KNOWN CHECKPOINT", cell.ColorBlue},
		enums.ColumnNameHighestSyncedCheckpoint:                 {"HIGHEST SYNCED CHECKPOINT", cell.ColorBlue},
		enums.ColumnNameLastExecutedCheckpoint:                  {"LAST EXECUTED CHECKPOINT", cell.ColorBlue},
		enums.ColumnNameCheckpointExecBacklog:                   {"CHECKPOINT EXEC BACKLOG", cell.ColorBlue},
		enums.ColumnNameCheckpointSyncBacklog:                   {"CHECKPOINT SYNC BACKLOG", cell.ColorBlue},
		enums.ColumnNameCheckSyncPercentage:                     {"CHECKPOINTS SYNC PERCENTAGE", cell.ColorBlue},
		enums.ColumnNameCheckpointsPerSecond:                    {"CHECKPOINTS VOLUME", cell.ColorBlue},
		enums.ColumnNameTotalTransactionCertificates:            {"TOTAL TRANSACTION CERTIFICATES", cell.ColorYellow},
		enums.ColumnNameTotalTransactionEffects:                 {"TOTAL TRANSACTION EFFECTS", cell.ColorYellow},
		enums.ColumnNameCertificatesCreated:                     {"CERTIFICATES CREATED", cell.ColorYellow},
		enums.ColumnNameCurrentRound:                            {"CURRENT ROUND", cell.ColorRed},
		enums.ColumnNameHighestProcessedRound:                   {"HIGHEST PROCESSED ROUND", cell.ColorRed},
		enums.ColumnNameLastCommittedRound:                      {"LAST COMMITTED ROUND", cell.ColorRed},
		enums.ColumnNameRoundsPerSecond:                         {"ROUNDS RATIO", cell.ColorRed},
		enums.ColumnNameCertificatesPerSecond:                   {"CERTIFICATES RATIO", cell.ColorYellow},
		enums.ColumnNameHandleCertificateNonConsensusLatencySum: {"CERTIFICATE NON CONSENSUS LATENCY", cell.ColorRed},
	}
)

// GetValidatorColumnValues returns a map of ColumnName values to corresponding values for a node at the specified index on the specified host.
// The function retrieves information about the node from the host's internal state and formats it into a map of NodeColumnName keys and corresponding values.
// The function also includes emoji values in the map if the specified flag is true.
func GetValidatorColumnValues(host host.Host) (ColumnValues, error) {
	return ColumnValues{
		enums.ColumnNameTotalTransactionCertificates:            host.Metrics.TotalTransactionCertificates,
		enums.ColumnNameTotalTransactionEffects:                 host.Metrics.TotalTransactionEffects,
		enums.ColumnNameHighestKnownCheckpoint:                  host.Metrics.HighestKnownCheckpoint,
		enums.ColumnNameHighestSyncedCheckpoint:                 host.Metrics.HighestSyncedCheckpoint,
		enums.ColumnNameLastExecutedCheckpoint:                  host.Metrics.LastExecutedCheckpoint,
		enums.ColumnNameCheckpointExecBacklog:                   host.Metrics.CheckpointExecBacklog,
		enums.ColumnNameCheckpointSyncBacklog:                   host.Metrics.CheckpointSyncBacklog,
		enums.ColumnNameCurrentEpoch:                            host.Metrics.CurrentEpoch,
		enums.ColumnNameCheckSyncPercentage:                     fmt.Sprintf("%v%%", host.Metrics.CheckSyncPercentage),
		enums.ColumnNameCheckpointsPerSecond:                    host.Metrics.CheckpointsPerSecond,
		enums.ColumnNameNetworkPeers:                            host.Metrics.NetworkPeers,
		enums.ColumnNameUptime:                                  host.Metrics.Uptime,
		enums.ColumnNameVersion:                                 host.Metrics.Version,
		enums.ColumnNameCommit:                                  host.Metrics.Commit,
		enums.ColumnNameCurrentRound:                            host.Metrics.CurrentRound,
		enums.ColumnNameHighestProcessedRound:                   host.Metrics.HighestProcessedRound,
		enums.ColumnNameLastCommittedRound:                      host.Metrics.LastCommittedRound,
		enums.ColumnNameRoundsPerSecond:                         host.Metrics.RoundsPerSecond,
		enums.ColumnNamePrimaryNetworkPeers:                     host.Metrics.PrimaryNetworkPeers,
		enums.ColumnNameWorkerNetworkPeers:                      host.Metrics.WorkerNetworkPeers,
		enums.ColumnNameSkippedConsensusTransactions:            host.Metrics.SkippedConsensusTransactions,
		enums.ColumnNameTotalSignatureErrors:                    host.Metrics.TotalSignatureErrors,
		enums.ColumnNameCertificatesCreated:                     host.Metrics.CertificatesCreated,
		enums.ColumnNameCertificatesPerSecond:                   host.Metrics.CertificatesPerSecond,
		enums.ColumnNameHandleCertificateNonConsensusLatencySum: host.Metrics.NonConsensusLatency,
	}, nil
}
