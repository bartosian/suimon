package dashboards

import (
	"fmt"

	"github.com/mum4k/termdash/cell"

	"github.com/bartosian/suimon/internal/core/domain/enums"
	domainhost "github.com/bartosian/suimon/internal/core/domain/host"
)

var (
	ColumnsConfigValidator = ColumnsConfig{
		// Overview section
		enums.ColumnNameCurrentEpoch: ColumnWidth20,
		enums.ColumnNameUptime:       ColumnWidth25,
		enums.ColumnNameVersion:      ColumnWidth25,
		enums.ColumnNameCommit:       ColumnWidth25,

		// Transactions section
		enums.ColumnNameTotalTransactionCertificates: ColumnWidth33,
		enums.ColumnNameTotalTransactionEffects:      ColumnWidth33,
		enums.ColumnNameCertificatesCreated:          ColumnWidth33,
		enums.ColumnNameCertificatesPerSecond:        ColumnWidth49,

		// Checkpoints section
		enums.ColumnNameLastExecutedCheckpoint:  ColumnWidth33,
		enums.ColumnNameHighestKnownCheckpoint:  ColumnWidth33,
		enums.ColumnNameHighestSyncedCheckpoint: ColumnWidth33,
		enums.ColumnNameCheckSyncPercentage:     ColumnWidth49,
		enums.ColumnNameCheckpointsPerSecond:    ColumnWidth49,

		// Rounds section
		enums.ColumnNameCurrentRound:          ColumnWidth33,
		enums.ColumnNameHighestProcessedRound: ColumnWidth33,
		enums.ColumnNameLastCommittedRound:    ColumnWidth33,
		enums.ColumnNameRoundsPerSecond:       ColumnWidth49,

		// Peers section
		enums.ColumnNameNetworkPeers:        ColumnWidth19,
		enums.ColumnNamePrimaryNetworkPeers: ColumnWidth19,
		enums.ColumnNameWorkerNetworkPeers:  ColumnWidth19,

		// Performance section
		enums.ColumnNameSkippedConsensusTransactions: ColumnWidth19,
		enums.ColumnNameTotalSignatureErrors:         ColumnWidth19,
	}

	RowsConfigValidator = RowsConfig{
		0: {
			Height: ColumnWidth14,
			Columns: []enums.ColumnName{
				enums.ColumnNameCurrentEpoch,
				enums.ColumnNameUptime,
				enums.ColumnNameVersion,
				enums.ColumnNameCommit,
			},
		},
		1: {
			Height: ColumnWidth14,
			Columns: []enums.ColumnName{
				enums.ColumnNameNetworkPeers,
				enums.ColumnNamePrimaryNetworkPeers,
				enums.ColumnNameWorkerNetworkPeers,
				enums.ColumnNameSkippedConsensusTransactions,
				enums.ColumnNameTotalSignatureErrors,
			},
		},
		2: {
			Height: ColumnWidth14,
			Columns: []enums.ColumnName{
				enums.ColumnNameLastExecutedCheckpoint,
				enums.ColumnNameHighestKnownCheckpoint,
				enums.ColumnNameHighestSyncedCheckpoint,
			},
		},
		3: {
			Height: ColumnWidth14,
			Columns: []enums.ColumnName{
				enums.ColumnNameCheckSyncPercentage,
				enums.ColumnNameCheckpointsPerSecond,
			},
		},
		4: {
			Height: ColumnWidth14,
			Columns: []enums.ColumnName{
				enums.ColumnNameTotalTransactionCertificates,
				enums.ColumnNameTotalTransactionEffects,
				enums.ColumnNameCertificatesCreated,
			},
		},
		5: {
			Height: ColumnWidth14,
			Columns: []enums.ColumnName{
				enums.ColumnNameCurrentRound,
				enums.ColumnNameHighestProcessedRound,
				enums.ColumnNameLastCommittedRound,
			},
		},
		6: {
			Height: ColumnWidth14,
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
func GetValidatorColumnValues(host *domainhost.Host) (ColumnValues, error) {
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
