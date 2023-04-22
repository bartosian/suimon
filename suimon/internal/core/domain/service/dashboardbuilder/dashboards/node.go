package dashboards

import (
	"fmt"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/host"
)

var (
	ColumnsConfigNode = ColumnsConfig{
		enums.ColumnNameHealth:                       5,
		enums.ColumnNameAddress:                      5,
		enums.ColumnNamePortRPC:                      20,
		enums.ColumnNameTotalTransactionBlocks:       20,
		enums.ColumnNameLatestCheckpoint:             20,
		enums.ColumnNameTotalTransactionCertificates: 20,
		enums.ColumnNameTotalTransactionEffects:      20,
		enums.ColumnNameHighestKnownCheckpoint:       30,
		enums.ColumnNameHighestSyncedCheckpoint:      30,
		enums.ColumnNameLastExecutedCheckpoint:       30,
		enums.ColumnNameCheckpointExecBacklog:        25,
		enums.ColumnNameCheckpointSyncBacklog:        25,
		enums.ColumnNameCurrentEpoch:                 25,
		enums.ColumnNameTXSyncPercentage:             25,
		enums.ColumnNameCheckSyncPercentage:          15,
		enums.ColumnNameNetworkPeers:                 15,
		enums.ColumnNameUptime:                       15,
		enums.ColumnNameVersion:                      15,
		enums.ColumnNameCommit:                       10,
		enums.ColumnNameCountry:                      50,
	}

	RowsConfigNode = RowsConfig{
		0: {
			Height: 15,
			Columns: []enums.ColumnName{
				enums.ColumnNameHealth,
				enums.ColumnNameAddress,
				enums.ColumnNamePortRPC,
				enums.ColumnNameTotalTransactionBlocks,
				enums.ColumnNameLatestCheckpoint,
				enums.ColumnNameTotalTransactionCertificates,
				enums.ColumnNameTotalTransactionEffects,
				enums.ColumnNameHighestKnownCheckpoint,
			},
		},
		1: {
			Height: 15,
			Columns: []enums.ColumnName{
				enums.ColumnNameHighestSyncedCheckpoint,
				enums.ColumnNameLastExecutedCheckpoint,
				enums.ColumnNameCheckpointExecBacklog,
				enums.ColumnNameCheckpointSyncBacklog,
				enums.ColumnNameCurrentEpoch,
				enums.ColumnNameTXSyncPercentage,
				enums.ColumnNameCheckSyncPercentage,
				enums.ColumnNameNetworkPeers,
			},
		},
		2: {
			Height: 15,
			Columns: []enums.ColumnName{
				enums.ColumnNameUptime,
				enums.ColumnNameVersion,
				enums.ColumnNameCommit,
				enums.ColumnNameCountry,
			},
		},
	}

	CellsConfigNode = CellsConfig{
		enums.ColumnNameHealth:                       "HEALTH",
		enums.ColumnNameAddress:                      "ADDRESS",
		enums.ColumnNamePortRPC:                      "RPC PORT",
		enums.ColumnNameTotalTransactionBlocks:       "TOTAL TRANSACTION BLOCKS",
		enums.ColumnNameLatestCheckpoint:             "LATEST CHECKPOINT",
		enums.ColumnNameTotalTransactionCertificates: "TOTAL TRANSACTION CERTIFICATES",
		enums.ColumnNameTotalTransactionEffects:      "TOTAL TRANSACTION EFFECTS",
		enums.ColumnNameHighestKnownCheckpoint:       "HIGHEST KNOWN CHECKPOINT",
		enums.ColumnNameHighestSyncedCheckpoint:      "HIGHEST SYNCED CHECKPOINT",
		enums.ColumnNameLastExecutedCheckpoint:       "LAST EXECUTED CHECKPOINT",
		enums.ColumnNameCheckpointExecBacklog:        "CHECKPOINT EXEC BACKLOG",
		enums.ColumnNameCheckpointSyncBacklog:        "CHECKPOINT SYNC BACKLOG",
		enums.ColumnNameCurrentEpoch:                 "CURRENT EPOCH",
		enums.ColumnNameTXSyncPercentage:             "TX SYNC PERCENTAGE",
		enums.ColumnNameCheckSyncPercentage:          "CHECKPOINTS SYNC PERCENTAGE",
		enums.ColumnNameNetworkPeers:                 "NETWORK PEERS",
		enums.ColumnNameUptime:                       "UPTIME",
		enums.ColumnNameVersion:                      "VERSION",
		enums.ColumnNameCommit:                       "COMMIT",
		enums.ColumnNameCountry:                      "COUNTRY",
	}
)

// GetNodeColumnValues returns a map of NodeColumnName values to corresponding values for a node at the specified index on the specified host.
// The function retrieves information about the node from the host's internal state and formats it into a map of NodeColumnName keys and corresponding values.
// The function also includes emoji values in the map if the specified flag is true.
// Returns a map of NodeColumnName keys to corresponding values.
func GetNodeColumnValues(host host.Host) ColumnValues {
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

func GetNodeColumnOptions(host host.Host) ColumnOptions {
	//var options []cell.Option

	//switch columnName {
	//case enums.ColumnNameHealth:
	//	status := host.Status
	//	color := cell.ColorGreen
	//
	//	switch status {
	//	case enums.StatusYellow:
	//		color = cell.ColorYellow
	//	case enums.StatusRed:
	//		color = cell.ColorRed
	//	}
	//
	//	options = append(options, cell.BgColor(color), cell.FgColor(color))
	//case enums.ColumnNameTotalTransactionBlocks:
	//	transactionsNode := host.Metrics.TotalTransactionsBlocks
	//	txSyncPercentageNode := host.Metrics.TxSyncPercentage
	//	transactionsRpc := rpc.Metrics.TotalTransactionsBlocks
	//	color := cell.ColorWhite
	//
	//	switch {
	//	case transactionsNode == 0:
	//		color = cell.ColorRed
	//	case txSyncPercentageNode < metrics.TotalTransactionsSyncPercentage || transactionsNode < transactionsRpc-metrics.TotalTransactionsLag:
	//		color = cell.ColorYellow
	//	}
	//
	//	return []segmentdisplay.WriteOption{segmentdisplay.WriteCellOpts(cell.FgColor(color))}
	//case enums.CellNameLatestCheckpoint:
	//	latestCheckpointNode := node.Metrics.LatestCheckpoint
	//	latestCheckpointRpc := rpc.Metrics.LatestCheckpoint
	//	color := cell.ColorWhite
	//
	//	switch {
	//	case latestCheckpointNode == 0:
	//		color = cell.ColorRed
	//	case latestCheckpointNode < latestCheckpointRpc-metrics.LatestCheckpointLag:
	//		color = cell.ColorYellow
	//	}
	//
	//	return []segmentdisplay.WriteOption{segmentdisplay.WriteCellOpts(cell.FgColor(color))}
	//case enums.CellNameHighestSyncedCheckpoint:
	//	highestCheckpointNode := node.Metrics.HighestSyncedCheckpoint
	//	highestCheckpointRpc := rpc.Metrics.HighestSyncedCheckpoint
	//	latestCheckpointRpc := rpc.Metrics.LatestCheckpoint
	//	color := cell.ColorWhite
	//
	//	switch {
	//	case highestCheckpointNode == 0:
	//		color = cell.ColorRed
	//	case highestCheckpointNode < highestCheckpointRpc-metrics.HighestSyncedCheckpointLag || highestCheckpointNode < latestCheckpointRpc-metrics.LatestCheckpointLag:
	//		color = cell.ColorYellow
	//	}
	//
	//	return []segmentdisplay.WriteOption{segmentdisplay.WriteCellOpts(cell.FgColor(color))}
	//case enums.CellNameTransactionsPerSecond:
	//	tpsNode := node.Metrics.TransactionsPerSecond
	//	txHistoryNode := node.Metrics.TransactionsHistory
	//	tpsRpc := rpc.Metrics.TransactionsPerSecond
	//	color := cell.ColorWhite
	//
	//	switch {
	//	case len(txHistoryNode) != metrics.TransactionsPerSecondWindow:
	//	case tpsNode == 0:
	//		color = cell.ColorRed
	//	case tpsNode < tpsRpc-metrics.TransactionsPerSecondLag:
	//		color = cell.ColorYellow
	//	}
	//
	//	return []segmentdisplay.WriteOption{segmentdisplay.WriteCellOpts(cell.FgColor(color))}
	//case enums.CellNameCheckpointsPerSecond:
	//	checkNode := node.Metrics.CheckpointsPerSecond
	//	checkHistoryNode := node.Metrics.CheckpointsHistory
	//	checkRpc := rpc.Metrics.CheckpointsPerSecond
	//	color := cell.ColorWhite
	//
	//	switch {
	//	case len(checkHistoryNode) != metrics.CheckpointsPerSecondWindow:
	//	case checkNode == 0:
	//		color = cell.ColorRed
	//	case checkNode < checkRpc-metrics.CheckpointsPerSecondLag:
	//		color = cell.ColorYellow
	//	}
	//
	//	return []segmentdisplay.WriteOption{segmentdisplay.WriteCellOpts(cell.FgColor(color))}
	//case enums.CellNameTXSyncProgress:
	//	var (
	//		syncProgress = node.Metrics.TxSyncPercentage
	//		color        = cell.ColorGreen
	//	)
	//
	//	switch {
	//	case syncProgress == 0:
	//		color = cell.ColorRed
	//	case syncProgress < metrics.TotalTransactionsSyncPercentage:
	//		color = cell.ColorYellow
	//	}
	//
	//	return []gauge.Option{gauge.Color(color), gauge.Border(linestyle.Light, cell.FgColor(color))}
	//case enums.CellNameCheckSyncProgress:
	//	syncProgress := node.Metrics.CheckSyncPercentage
	//	color := cell.ColorGreen
	//
	//	switch {
	//	case syncProgress == 0:
	//		color = cell.ColorRed
	//	case syncProgress < metrics.TotalCheckpointsSyncPercentage:
	//		color = cell.ColorYellow
	//	}
	//
	//	return []gauge.Option{gauge.Color(color), gauge.Border(linestyle.Light, cell.FgColor(color))}
	//case enums.CellNameEpochProgress, enums.CellNameDiskUsage, enums.CellNameCpuUsage, enums.CellNameMemoryUsage:
	//	options = append(options, cell.Bold())
	//case enums.CellNameEpochTimeTillTheEnd, enums.CellNameDatabaseSize, enums.CellNameBytesReceived, enums.CellNameBytesSent:
	//	return []segmentdisplay.WriteOption{segmentdisplay.WriteCellOpts(cell.FgColor(cell.ColorWhite)), segmentdisplay.WriteCellOpts(cell.FgColor(cell.ColorGreen))}
	//case enums.CellNameUptime:
	//	uptime := node.Metrics.Uptime
	//	color := cell.ColorWhite
	//
	//	switch {
	//	case uptime == "":
	//	case uptime == "0.0":
	//		color = cell.ColorRed
	//	case uptime < "1.0":
	//		color = cell.ColorYellow
	//	}
	//
	//	return []segmentdisplay.WriteOption{segmentdisplay.WriteCellOpts(cell.FgColor(color)), segmentdisplay.WriteCellOpts(cell.FgColor(cell.ColorGreen))}
	//case enums.CellNameConnectedPeers:
	//	peers := node.Metrics.NetworkPeers
	//	color := cell.ColorWhite
	//
	//	switch {
	//	case peers == 0:
	//		color = cell.ColorRed
	//	case peers < 3:
	//		color = cell.ColorYellow
	//	}
	//
	//	return []segmentdisplay.WriteOption{segmentdisplay.WriteCellOpts(cell.FgColor(color))}
	//case enums.CellNameNodeLogs:
	//	options = append(options, cell.FgColor(cell.ColorWhite), cell.Bold())
	//case enums.CellNameTPSTracker:
	//	tpsNode := node.Metrics.TransactionsPerSecond
	//	tpsRpc := rpc.Metrics.TransactionsPerSecond
	//	color := cell.ColorGreen
	//
	//	if tpsNode < tpsRpc-metrics.TransactionsPerSecondLag {
	//		color = cell.ColorYellow
	//	}
	//
	//	return []sparkline.Option{sparkline.Color(color)}
	//case enums.CellNameCPSTracker:
	//	checkNode := node.Metrics.CheckpointsPerSecond
	//	checkRpc := rpc.Metrics.CheckpointsPerSecond
	//	color := cell.ColorBlue
	//
	//	if checkNode < checkRpc-metrics.CheckpointsPerSecondLag {
	//		color = cell.ColorYellow
	//	}
	//
	//	return []sparkline.Option{sparkline.Color(color)}
	//default:
	//	return []segmentdisplay.WriteOption{segmentdisplay.WriteCellOpts(cell.FgColor(cell.ColorWhite))}
	//}
	//
	//return options

	return nil
}
