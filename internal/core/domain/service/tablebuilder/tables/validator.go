package tables

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/bartosian/suimon/internal/core/domain/enums"
	domainhost "github.com/bartosian/suimon/internal/core/domain/host"
)

var (
	ColumnsConfigValidator = ColumnsConfig{
		enums.ColumnNameIndex:                                   NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameHealth:                                  NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameAddress:                                 NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameTotalTransactionCertificates:            NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameTotalTransactionCertificatesCreated:     NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
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
		enums.ColumnNameLastCommittedLeaderRound:                NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameHighestAcceptedRound:                    NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameConsensusRoundProberCurrentRoundGaps:    NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSkippedConsensusTransactions:            NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameTotalSignatureErrors:                    NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameHandleCertificateNonConsensusLatencySum: NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameValidatorCurrentVotingRight:             NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameValidatorTotalTransactionCertificates:   NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameNumberSharedObjectTransactions:          NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
	}

	RowsConfigValidator = RowsConfig{
		0: {
			enums.ColumnNameIndex,
			enums.ColumnNameHealth,
			enums.ColumnNameAddress,
			enums.ColumnNameCurrentEpoch,
			enums.ColumnNameValidatorCurrentVotingRight,
			enums.ColumnNameTotalTransactionCertificates,
			enums.ColumnNameTotalTransactionCertificatesCreated,
			enums.ColumnNameSkippedConsensusTransactions,
			enums.ColumnNameTotalTransactionEffects,
		},
		1: {
			enums.ColumnNameHighestKnownCheckpoint,
			enums.ColumnNameLastExecutedCheckpoint,
			enums.ColumnNameCheckpointExecBacklog,
			enums.ColumnNameHighestSyncedCheckpoint,
			enums.ColumnNameCheckpointSyncBacklog,
			enums.ColumnNameUptime,
			enums.ColumnNameVersion,
			enums.ColumnNameCommit,
			enums.ColumnNameLastCommittedLeaderRound,
		},
		2: {
			enums.ColumnNameHighestAcceptedRound,
			enums.ColumnNameConsensusRoundProberCurrentRoundGaps,
			enums.ColumnNameNetworkPeers,
			enums.ColumnNameTotalSignatureErrors,
			enums.ColumnNameNumberSharedObjectTransactions,
		},
	}
)

// GetValidatorColumnValues returns a map of ValidatorColumnName values to corresponding values for a validator at the specified index on the specified host.
// The function retrieves information about the validator from the host's internal state and formats it into a map of ValidatorColumnName keys and corresponding values.
// The function also includes emoji values in the map if the specified flag is true.
// Returns a map of ValidatorColumnName keys to corresponding values.
func GetValidatorColumnValues(idx int, host *domainhost.Host) ColumnValues {
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
		enums.ColumnNameLastCommittedLeaderRound:                host.Metrics.LastCommittedLeaderRound,
		enums.ColumnNameHighestAcceptedRound:                    host.Metrics.HighestAcceptedRound,
		enums.ColumnNameConsensusRoundProberCurrentRoundGaps:    host.Metrics.ConsensusRoundProberCurrentRoundGaps,
		enums.ColumnNameSkippedConsensusTransactions:            host.Metrics.SkippedConsensusTransactions,
		enums.ColumnNameTotalSignatureErrors:                    host.Metrics.TotalSignatureErrors,
		enums.ColumnNameTotalTransactionCertificatesCreated:     host.Metrics.TotalTransactionCertificatesCreated,
		enums.ColumnNameHandleCertificateNonConsensusLatencySum: host.Metrics.NonConsensusLatency,
		enums.ColumnNameCountry:                                 country,
		enums.ColumnNameValidatorCurrentVotingRight:             fmt.Sprintf("%v%%", host.Metrics.CurrentVotingRight),
		enums.ColumnNameValidatorTotalTransactionCertificates:   host.Metrics.TotalTransactionCertificates,
		enums.ColumnNameNumberSharedObjectTransactions:          host.Metrics.NumberSharedObjectTransactions,
	}

	return columnValues
}
