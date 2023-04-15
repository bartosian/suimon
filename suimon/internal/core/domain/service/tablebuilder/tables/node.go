package tables

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/host"
)

var (
	ColumnsConfigNode = ColumnsConfig{
		enums.ColumnNameIndex:                        NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameHealth:                       NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameAddress:                      NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
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
		enums.ColumnNameCountry:                      NewDefaultColumnConfig(text.AlignCenter, text.AlignLeft, false),
	}
	RowsConfigNode = RowsConfig{
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

// GetNodeColumnValues returns a map of NodeColumnName values to corresponding values for a node at the specified index on the specified host.
// The function retrieves information about the node from the host's internal state and formats it into a map of NodeColumnName keys and corresponding values.
// The function also includes emoji values in the map if the specified flag is true.
// Returns a map of NodeColumnName keys to corresponding values.
func GetNodeColumnValues(idx int, host host.Host) ColumnValues {
	status := host.Status.StatusToPlaceholder()

	var country string
	if host.IPInfo != nil {
		country = host.IPInfo.CountryName
	}

	port := host.Ports[enums.PortTypeRPC]
	if port == "" {
		port = RpcPortDefault
	}

	address := host.Endpoint.Address

	columnValues := ColumnValues{
		enums.ColumnNameIndex:                        idx + 1,
		enums.ColumnNameHealth:                       status,
		enums.ColumnNameAddress:                      address,
		enums.ColumnNamePortRPC:                      port,
		enums.ColumnNameTotalTransactionBlocks:       host.Metrics.TotalTransactionsBlocks,
		enums.ColumnNameTotalTransactionCertificates: host.Metrics.TotalTransactionCertificates,
		enums.ColumnNameTotalTransactionEffects:      host.Metrics.TotalTransactionEffects,
		enums.ColumnNameLatestCheckpoint:             host.Metrics.LatestCheckpoint,
		enums.ColumnNameHighestKnownCheckpoint:       host.Metrics.HighestKnownCheckpoint,
		enums.ColumnNameHighestSyncedCheckpoint:      host.Metrics.HighestSyncedCheckpoint,
		enums.ColumnNameLastExecutedCheckpoint:       host.Metrics.LastExecutedCheckpoint,
		enums.ColumnNameCheckpointExecBacklog:        host.Metrics.CheckpointExecBacklog,
		enums.ColumnNameCheckpointSyncBacklog:        host.Metrics.CheckpointSyncBacklog,
		enums.ColumnNameCurrentEpoch:                 host.Metrics.CurrentEpoch,
		enums.ColumnNameTXSyncPercentage:             fmt.Sprintf("%v%%", host.Metrics.TxSyncPercentage),
		enums.ColumnNameCheckSyncPercentage:          fmt.Sprintf("%v%%", host.Metrics.CheckSyncPercentage),
		enums.ColumnNameNetworkPeers:                 host.Metrics.NetworkPeers,
		enums.ColumnNameUptime:                       host.Metrics.Uptime,
		enums.ColumnNameVersion:                      host.Metrics.Version,
		enums.ColumnNameCommit:                       host.Metrics.Commit,
		enums.ColumnNameCountry:                      country,
	}

	return columnValues
}
