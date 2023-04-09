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
	SortConfigNode = tablebuilder.SortConfig{
		{Name: enums.ColumnNameHealth.ToString(), Mode: table.Dsc},
		{Name: enums.ColumnNameTotalTransactionBlocks.ToString(), Mode: table.Dsc},
		{Name: enums.ColumnNameLatestCheckpoint.ToString(), Mode: table.Dsc},
	}
	ColumnsConfigNode = tablebuilder.ColumnsConfig{
		enums.ColumnNameIndex:                        tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameHealth:                       tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameAddress:                      tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNamePortRPC:                      tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameTotalTransactionBlocks:       tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameLatestCheckpoint:             tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameTotalTransactionCertificates: tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameTotalTransactionEffects:      tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameHighestKnownCheckpoint:       tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameHighestSyncedCheckpoint:      tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameLastExecutedCheckpoint:       tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameCheckpointExecBacklog:        tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameCheckpointSyncBacklog:        tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameCurrentEpoch:                 tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameTXSyncPercentage:             tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameCheckSyncPercentage:          tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameNetworkPeers:                 tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameUptime:                       tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameVersion:                      tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameCommit:                       tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameCountry:                      tablebuilder.NewDefaultColumnConfig(text.AlignCenter, text.AlignLeft, false),
	}
	RowsConfigNode = tablebuilder.RowsConfig{
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
func GetNodeColumnValues(idx int, host host.Host) tablebuilder.ColumnValues {
	status := host.Status.StatusToPlaceholder()
	country := ""
	if host.IPInfo != nil {
		country = host.IPInfo.Country
	}
	port := host.Ports[enums.PortTypeRPC]
	if port == "" {
		port = tablebuilder.RpcPortDefault
	}
	address := host.HostPort.Address

	columnValues := tablebuilder.ColumnValues{
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
