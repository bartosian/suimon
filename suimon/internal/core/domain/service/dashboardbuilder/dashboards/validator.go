package dashboards

import (
	"fmt"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/host"
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
				enums.ColumnNameTotalTransactionCertificates,
				enums.ColumnNameTotalTransactionEffects,
				enums.ColumnNameCertificatesCreated,
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
				enums.ColumnNameCurrentRound,
				enums.ColumnNameHighestProcessedRound,
				enums.ColumnNameLastCommittedRound,
			},
		},
		5: {
			Height: 14,
			Columns: []enums.ColumnName{
				enums.ColumnNameNetworkPeers,
				enums.ColumnNamePrimaryNetworkPeers,
				enums.ColumnNameWorkerNetworkPeers,
				enums.ColumnNameSkippedConsensusTransactions,
				enums.ColumnNameTotalSignatureErrors,
			},
		},
	}

	CellsConfigValidator = CellsConfig{
		enums.ColumnNameTotalTransactionCertificates:            "TOTAL TRANSACTION CERTIFICATES",
		enums.ColumnNameTotalTransactionEffects:                 "TOTAL TRANSACTION EFFECTS",
		enums.ColumnNameHighestKnownCheckpoint:                  "HIGHEST KNOWN CHECKPOINT",
		enums.ColumnNameHighestSyncedCheckpoint:                 "HIGHEST SYNCED CHECKPOINT",
		enums.ColumnNameLastExecutedCheckpoint:                  "LAST EXECUTED CHECKPOINT",
		enums.ColumnNameCheckpointExecBacklog:                   "CHECKPOINT EXEC BACKLOG",
		enums.ColumnNameCheckpointSyncBacklog:                   "CHECKPOINT SYNC BACKLOG",
		enums.ColumnNameCurrentEpoch:                            "CURRENT EPOCH",
		enums.ColumnNameCheckSyncPercentage:                     "CHECKPOINTS SYNC PERCENTAGE",
		enums.ColumnNameCheckpointsPerSecond:                    "CHECKPOINTS RATE",
		enums.ColumnNameUptime:                                  "UPTIME",
		enums.ColumnNameVersion:                                 "VERSION",
		enums.ColumnNameCommit:                                  "COMMIT",
		enums.ColumnNameCurrentRound:                            "CURRENT ROUND",
		enums.ColumnNameHighestProcessedRound:                   "HIGHEST PROCESSED ROUND",
		enums.ColumnNameLastCommittedRound:                      "LAST COMMITTED ROUND",
		enums.ColumnNameNetworkPeers:                            "SUI NETWORK PEERS",
		enums.ColumnNamePrimaryNetworkPeers:                     "PRIMARY NETWORK PEERS",
		enums.ColumnNameWorkerNetworkPeers:                      "WORKER NETWORK PEERS",
		enums.ColumnNameSkippedConsensusTransactions:            "SKIPPED CONSENSUS TRANSACTIONS",
		enums.ColumnNameTotalSignatureErrors:                    "TOTAL SIGNATURE ERRORS",
		enums.ColumnNameCertificatesCreated:                     "CERTIFICATES CREATED",
		enums.ColumnNameHandleCertificateNonConsensusLatencySum: "CERTIFICATE NON CONSENSUS LATENCY",
	}
)

// GetValidatorColumnValues returns a map of ColumnName values to corresponding values for a node at the specified index on the specified host.
// The function retrieves information about the node from the host's internal state and formats it into a map of NodeColumnName keys and corresponding values.
// The function also includes emoji values in the map if the specified flag is true.
func GetValidatorColumnValues(host host.Host) ColumnValues {
	columnValues := ColumnValues{
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
		enums.ColumnNamePrimaryNetworkPeers:                     host.Metrics.PrimaryNetworkPeers,
		enums.ColumnNameWorkerNetworkPeers:                      host.Metrics.WorkerNetworkPeers,
		enums.ColumnNameSkippedConsensusTransactions:            host.Metrics.SkippedConsensusTransactions,
		enums.ColumnNameTotalSignatureErrors:                    host.Metrics.TotalSignatureErrors,
		enums.ColumnNameCertificatesCreated:                     host.Metrics.CertificatesCreated,
		enums.ColumnNameHandleCertificateNonConsensusLatencySum: host.Metrics.NonConsensusLatency,
	}

	return columnValues
}
