package tables

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/host"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/service/tablebuilder"
)

var (
	SortConfigValidator = tablebuilder.SortConfig{
		{Name: string(enums.ColumnNameHealth), Mode: table.Dsc},
		{Name: string(enums.ColumnNameTotalTransactionBlocks), Mode: table.Dsc},
		{Name: string(enums.ColumnNameLatestCheckpoint), Mode: table.Dsc},
	}
	ColumnsConfigValidator = tablebuilder.ColumnsConfig{
		enums.ColumnNameIndex:                        tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameHealth:                       tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameAddress:                      tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameTotalTransactionCertificates: tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameTotalTransactionEffects:      tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameHighestKnownCheckpoint:       tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameHighestSyncedCheckpoint:      tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameLastExecutedCheckpoint:       tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameCheckpointExecBacklog:        tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameCheckpointSyncBacklog:        tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameCurrentEpoch:                 tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameCheckSyncPercentage:          tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameNetworkPeers:                 tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameUptime:                       tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameVersion:                      tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameCommit:                       tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameCountry:                      tablebuilder.NewDefaultColumnConfig(text.AlignLeft, text.AlignCenter, false),
		enums.ColumnNameCurrentRound:                 tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameHighestProcessedRound:        tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameLastCommittedRound:           tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNamePrimaryNetworkPeers:          tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameWorkerNetworkPeers:           tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameSkippedConsensusTransactions: tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameTotalSignatureErrors:         tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
	}
	RowsConfigValidator = tablebuilder.RowsConfig{
		0: {
			enums.ColumnNameIndex,
			enums.ColumnNameHealth,
			enums.ColumnNameAddress,
			enums.ColumnNameCurrentEpoch,
			enums.ColumnNameTotalTransactionCertificates,
			enums.ColumnNameTotalTransactionEffects,
			enums.ColumnNameHighestKnownCheckpoint,
			enums.ColumnNameLastExecutedCheckpoint,
			enums.ColumnNameCheckpointExecBacklog,
			enums.ColumnNameHighestSyncedCheckpoint,
			enums.ColumnNameCheckpointSyncBacklog,
			enums.ColumnNameNetworkPeers,
		},
		1: {
			enums.ColumnNameUptime,
			enums.ColumnNameVersion,
			enums.ColumnNameCommit,
			enums.ColumnNameCurrentRound,
			enums.ColumnNameHighestProcessedRound,
			enums.ColumnNameLastCommittedRound,
			enums.ColumnNamePrimaryNetworkPeers,
			enums.ColumnNameWorkerNetworkPeers,
			enums.ColumnNameSkippedConsensusTransactions,
			enums.ColumnNameTotalSignatureErrors,
		},
	}
)

// GetValidatorColumnValues returns a map of ValidatorColumnName values to corresponding values for a validator at the specified index on the specified host.
// The function retrieves information about the validator from the host's internal state and formats it into a map of ValidatorColumnName keys and corresponding values.
// The function also includes emoji values in the map if the specified flag is true.
// Returns a map of ValidatorColumnName keys to corresponding values.
func GetValidatorColumnValues(idx int, host host.Host) tablebuilder.ColumnValues {
	status := host.Status.StatusToPlaceholder()
	country := ""
	if host.Location != nil {
		country = host.Location.String()
	}

	address := host.HostPort.Address

	columnValues := tablebuilder.ColumnValues{
		enums.ColumnNameIndex:                        idx + 1,
		enums.ColumnNameHealth:                       status,
		enums.ColumnNameAddress:                      address,
		enums.ColumnNameTotalTransactionCertificates: host.Metrics.TotalTransactionCertificates,
		enums.ColumnNameTotalTransactionEffects:      host.Metrics.TotalTransactionEffects,
		enums.ColumnNameHighestKnownCheckpoint:       host.Metrics.HighestKnownCheckpoint,
		enums.ColumnNameHighestSyncedCheckpoint:      host.Metrics.HighestSyncedCheckpoint,
		enums.ColumnNameLastExecutedCheckpoint:       host.Metrics.LastExecutedCheckpoint,
		enums.ColumnNameCheckpointExecBacklog:        host.Metrics.CheckpointExecBacklog,
		enums.ColumnNameCheckpointSyncBacklog:        host.Metrics.CheckpointSyncBacklog,
		enums.ColumnNameCurrentEpoch:                 host.Metrics.CurrentEpoch,
		enums.ColumnNameCheckSyncPercentage:          fmt.Sprintf("%v%%", host.Metrics.CheckSyncPercentage),
		enums.ColumnNameNetworkPeers:                 host.Metrics.NetworkPeers,
		enums.ColumnNameUptime:                       host.Metrics.Uptime,
		enums.ColumnNameVersion:                      host.Metrics.Version,
		enums.ColumnNameCommit:                       host.Metrics.Commit,
		enums.ColumnNameCurrentRound:                 host.Metrics.CurrentRound,
		enums.ColumnNameHighestProcessedRound:        host.Metrics.HighestProcessedRound,
		enums.ColumnNameLastCommittedRound:           host.Metrics.LastCommittedRound,
		enums.ColumnNamePrimaryNetworkPeers:          host.Metrics.PrimaryNetworkPeers,
		enums.ColumnNameWorkerNetworkPeers:           host.Metrics.WorkerNetworkPeers,
		enums.ColumnNameSkippedConsensusTransactions: host.Metrics.SkippedConsensusTransactions,
		enums.ColumnNameTotalSignatureErrors:         host.Metrics.TotalSignatureErrors,
		enums.ColumnNameCountry:                      country,
	}

	return columnValues
}
