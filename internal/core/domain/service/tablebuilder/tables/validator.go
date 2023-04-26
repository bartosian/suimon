package tables

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/host"
)

var (
	ColumnsConfigValidator = ColumnsConfig{
		enums.ColumnNameIndex:                                   NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameHealth:                                  NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameAddress:                                 NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameTotalTransactionCertificates:            NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameCertificatesCreated:                     NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameTotalTransactionEffects:                 NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameHighestKnownCheckpoint:                  NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameHighestSyncedCheckpoint:                 NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameLastExecutedCheckpoint:                  NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameCheckpointExecBacklog:                   NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameCheckpointSyncBacklog:                   NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameCurrentEpoch:                            NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameCheckSyncPercentage:                     NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameNetworkPeers:                            NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameUptime:                                  NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameVersion:                                 NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameCommit:                                  NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameCountry:                                 NewDefaultColumnConfig(text.AlignLeft, text.AlignCenter, false),
		enums.ColumnNameCurrentRound:                            NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameHighestProcessedRound:                   NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameLastCommittedRound:                      NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNamePrimaryNetworkPeers:                     NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameWorkerNetworkPeers:                      NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSkippedConsensusTransactions:            NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameTotalSignatureErrors:                    NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameHandleCertificateNonConsensusLatencySum: NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
	}

	RowsConfigValidator = RowsConfig{
		0: {
			enums.ColumnNameIndex,
			enums.ColumnNameHealth,
			enums.ColumnNameAddress,
			enums.ColumnNameCurrentEpoch,
			enums.ColumnNameTotalTransactionCertificates,
			enums.ColumnNameCertificatesCreated,
			enums.ColumnNameSkippedConsensusTransactions,
			enums.ColumnNameTotalTransactionEffects,
			enums.ColumnNameHighestKnownCheckpoint,
			enums.ColumnNameLastExecutedCheckpoint,
			enums.ColumnNameCheckpointExecBacklog,
			enums.ColumnNameHighestSyncedCheckpoint,
			enums.ColumnNameCheckpointSyncBacklog,
		},
		1: {
			enums.ColumnNameUptime,
			enums.ColumnNameVersion,
			enums.ColumnNameCommit,
			enums.ColumnNameCurrentRound,
			enums.ColumnNameHighestProcessedRound,
			enums.ColumnNameLastCommittedRound,
			enums.ColumnNameNetworkPeers,
			enums.ColumnNamePrimaryNetworkPeers,
			enums.ColumnNameWorkerNetworkPeers,
			enums.ColumnNameTotalSignatureErrors,
		},
	}
)

// GetValidatorColumnValues returns a map of ValidatorColumnName values to corresponding values for a validator at the specified index on the specified host.
// The function retrieves information about the validator from the host's internal state and formats it into a map of ValidatorColumnName keys and corresponding values.
// The function also includes emoji values in the map if the specified flag is true.
// Returns a map of ValidatorColumnName keys to corresponding values.
func GetValidatorColumnValues(idx int, host host.Host) ColumnValues {
	status := host.Status.StatusToPlaceholder()

	var country string
	if host.IPInfo != nil {
		country = host.IPInfo.CountryName
	}

	address := host.Endpoint.Address

	columnValues := ColumnValues{
		enums.ColumnNameIndex:                                   idx + 1,
		enums.ColumnNameHealth:                                  status,
		enums.ColumnNameAddress:                                 address,
		enums.ColumnNameTotalTransactionCertificates:            host.Metrics.TotalTransactionCertificates,
		enums.ColumnNameTotalTransactionEffects:                 host.Metrics.TotalTransactionEffects,
		enums.ColumnNameHighestKnownCheckpoint:                  host.Metrics.HighestKnownCheckpoint,
		enums.ColumnNameHighestSyncedCheckpoint:                 host.Metrics.HighestSyncedCheckpoint,
		enums.ColumnNameLastExecutedCheckpoint:                  host.Metrics.LastExecutedCheckpoint,
		enums.ColumnNameCheckpointExecBacklog:                   host.Metrics.CheckpointExecBacklog,
		enums.ColumnNameCheckpointSyncBacklog:                   host.Metrics.CheckpointSyncBacklog,
		enums.ColumnNameCurrentEpoch:                            host.Metrics.CurrentEpoch,
		enums.ColumnNameCheckSyncPercentage:                     fmt.Sprintf("%v%%", host.Metrics.CheckSyncPercentage),
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
		enums.ColumnNameCountry:                                 country,
	}

	return columnValues
}
