package controller

import (
	"strings"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/widgets/gauge"
	"github.com/mum4k/termdash/widgets/segmentdisplay"
	"github.com/mum4k/termdash/widgets/sparkline"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/dashboardbuilder/dashboards"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/metrics"
	"github.com/bartosian/sui_helpers/suimon/internal/pkg/utility"
)

// getMetricForDashboardCell retrieves the current metric data for the specified dashboard cell.
// Parameters: cellName: a dashboards.CellName representing the name of the dashboard cell for which to retrieve data.
// Returns: a value of any type representing the current metric data for the specified dashboard cell.
func (checker *CheckerController) getMetricForDashboardCell(cellName enums.CellName) any {
	var (
		node = checker.node[0]
		rpc  = checker.rpc[0]
	)

	switch cellName {
	case enums.CellNameNodeHealth:
		return node.Status.DashboardStatus()
	case enums.CellNameNetworkHealth:
		return rpc.Status.DashboardStatus()
	case enums.CellNameTransactionsPerSecond:
		if len(node.Metrics.TransactionsHistory) < metrics.TransactionsPerSecondWindow {
			return dashboards.DashboardLoadingBlinkValue()
		}

		fallthrough
	case enums.CellNameTPSTracker:
		return node.Metrics.TransactionsPerSecond
	case enums.CellNameCheckpointsPerSecond:
		if len(node.Metrics.CheckpointsHistory) < metrics.CheckpointsPerSecondWindow {
			return dashboards.DashboardLoadingBlinkValue()
		}

		fallthrough
	case enums.CellNameCPSTracker:
		return node.Metrics.CheckpointsPerSecond
	case enums.CellNameTotalTransactions:
		return node.Metrics.TotalTransactions
	case enums.CellNameTotalTransactionCertificates:
		return node.Metrics.TotalTransactionCertificates
	case enums.CellNameTotalTransactionEffects:
		return node.Metrics.TotalTransactionEffects
	case enums.CellNameLatestCheckpoint:
		return node.Metrics.LatestCheckpoint
	case enums.CellNameHighestKnownCheckpoint:
		return node.Metrics.HighestKnownCheckpoint
	case enums.CellNameLastExecutedCheckpoint:
		return node.Metrics.LastExecutedCheckpoint
	case enums.CellNameHighestSyncedCheckpoint:
		return node.Metrics.HighestSyncedCheckpoint
	case enums.CellNameCheckpointSyncBacklog:
		return node.Metrics.CheckpointSyncBacklog
	case enums.CellNameCheckpointExecBacklog:
		return node.Metrics.CheckpointExecBacklog
	case enums.CellNameConnectedPeers:
		return node.Metrics.NetworkPeers
	case enums.CellNameTXSyncProgress:
		return node.Metrics.TxSyncPercentage
	case enums.CellNameCheckSyncProgress:
		return node.Metrics.CheckSyncPercentage
	case enums.CellNameUptime:
		return []string{strings.Split(node.Metrics.Uptime, " ")[0]}
	case enums.CellNameVersion:
		return node.Metrics.Version
	case enums.CellNameCommit:
		return node.Metrics.Commit
	case enums.CellNameEpochProgress:
		epochLabel := rpc.Metrics.GetEpochLabel()
		epochPercentage := rpc.Metrics.GetEpochProgress()

		return dashboards.NewDonutInput(epochLabel, epochPercentage)
	case enums.CellNameCurrentEpoch:
		return rpc.Metrics.SystemState.Epoch
	case enums.CellNameEpochTimeTillTheEnd:
		return rpc.Metrics.GetEpochTimer()
	case enums.CellNameDiskUsage:
		usageLabel, usagePercentage := metrics.GetDonutUsageMetric("GB", utility.GetDiskUsage)

		return dashboards.NewDonutInput(usageLabel, usagePercentage)
	case enums.CellNameDatabaseSize:
		dbSize := metrics.GetDirectorySize(checker.suimonConfig.DbPath)

		return dbSize
	case enums.CellNameBytesSent:
		bytesSent := metrics.GetNetworkUsageMetric(enums.CellNameBytesSent)

		return bytesSent
	case enums.CellNameBytesReceived:
		bytesReceived := metrics.GetNetworkUsageMetric(enums.CellNameBytesReceived)

		return bytesReceived
	case enums.CellNameMemoryUsage:
		usageLabel, usagePercentage := metrics.GetDonutUsageMetric("%", utility.GetMemoryUsage)

		return dashboards.NewDonutInput(usageLabel, usagePercentage)
	case enums.CellNameCpuUsage:
		usageLabel, usagePercentage := metrics.GetDonutUsageMetric("%", utility.GetCPUUsage)

		return dashboards.NewDonutInput(usageLabel, usagePercentage)
	default:
		return ""
	}
}

// getOptionsForDashboardCell retrieves the options data for a specific dashboard cell and returns it in the appropriate format for display in the dashboard.
// Parameters: cellName: a dashboards.CellName representing the name of the dashboard cell for which to retrieve options data.
// Returns: the options data for the specified dashboard cell in the appropriate format for display in the dashboard.
func (checker *CheckerController) getOptionsForDashboardCell(cellName enums.CellName) any {
	var (
		node    = checker.node[0]
		rpc     = checker.rpc[0]
		options []cell.Option
	)

	switch cellName {
	case enums.CellNameNodeHealth:
		var (
			status = node.Status
			color  = cell.ColorGreen
		)

		switch status {
		case enums.StatusYellow:
			color = cell.ColorYellow
		case enums.StatusRed:
			color = cell.ColorRed
		}

		options = append(options, cell.BgColor(color), cell.FgColor(color))
	case enums.CellNameNetworkHealth:
		var (
			status = rpc.Status
			color  = cell.ColorGreen
		)

		switch status {
		case enums.StatusYellow:
			color = cell.ColorYellow
		case enums.StatusRed:
			color = cell.ColorRed
		}

		options = append(options, cell.BgColor(color), cell.FgColor(color))
	case enums.CellNameTotalTransactions:
		var (
			transactionsNode     = node.Metrics.TotalTransactions
			txSyncPercentageNode = node.Metrics.TxSyncPercentage
			transactionsRpc      = rpc.Metrics.TotalTransactions
			color                = cell.ColorWhite
		)

		switch {
		case transactionsNode == 0:
			color = cell.ColorRed
		case txSyncPercentageNode < metrics.TotalTransactionsSyncPercentage || transactionsNode < transactionsRpc-metrics.TotalTransactionsLag:
			color = cell.ColorYellow
		}

		return []segmentdisplay.WriteOption{segmentdisplay.WriteCellOpts(cell.FgColor(color))}
	case enums.CellNameLatestCheckpoint:
		var (
			latestCheckpointNode = node.Metrics.LatestCheckpoint
			latestCheckpointRpc  = rpc.Metrics.LatestCheckpoint
			color                = cell.ColorWhite
		)

		switch {
		case latestCheckpointNode == 0:
			color = cell.ColorRed
		case latestCheckpointNode < latestCheckpointRpc-metrics.LatestCheckpointLag:
			color = cell.ColorYellow
		}

		return []segmentdisplay.WriteOption{segmentdisplay.WriteCellOpts(cell.FgColor(color))}
	case enums.CellNameHighestSyncedCheckpoint:
		var (
			highestCheckpointNode = node.Metrics.HighestSyncedCheckpoint
			highestCheckpointRpc  = rpc.Metrics.HighestSyncedCheckpoint
			latestCheckpointRpc   = rpc.Metrics.LatestCheckpoint
			color                 = cell.ColorWhite
		)

		switch {
		case highestCheckpointNode == 0:
			color = cell.ColorRed
		case highestCheckpointNode < highestCheckpointRpc-metrics.HighestSyncedCheckpointLag || highestCheckpointNode < latestCheckpointRpc-metrics.LatestCheckpointLag:
			color = cell.ColorYellow
		}

		return []segmentdisplay.WriteOption{segmentdisplay.WriteCellOpts(cell.FgColor(color))}
	case enums.CellNameTransactionsPerSecond:
		var (
			tpsNode       = node.Metrics.TransactionsPerSecond
			txHistoryNode = node.Metrics.TransactionsHistory
			tpsRpc        = rpc.Metrics.TransactionsPerSecond
			color         = cell.ColorWhite
		)

		switch {
		case len(txHistoryNode) != metrics.TransactionsPerSecondWindow:
		case tpsNode == 0:
			color = cell.ColorRed
		case tpsNode < tpsRpc-metrics.TransactionsPerSecondLag:
			color = cell.ColorYellow
		}

		return []segmentdisplay.WriteOption{segmentdisplay.WriteCellOpts(cell.FgColor(color))}
	case enums.CellNameCheckpointsPerSecond:
		var (
			checkNode        = node.Metrics.CheckpointsPerSecond
			checkHistoryNode = node.Metrics.CheckpointsHistory
			checkRpc         = rpc.Metrics.CheckpointsPerSecond
			color            = cell.ColorWhite
		)

		switch {
		case len(checkHistoryNode) != metrics.CheckpointsPerSecondWindow:
		case checkNode == 0:
			color = cell.ColorRed
		case checkNode < checkRpc-metrics.CheckpointsPerSecondLag:
			color = cell.ColorYellow
		}

		return []segmentdisplay.WriteOption{segmentdisplay.WriteCellOpts(cell.FgColor(color))}
	case enums.CellNameTXSyncProgress:
		var (
			syncProgress = node.Metrics.TxSyncPercentage
			color        = cell.ColorGreen
		)

		switch {
		case syncProgress == 0:
			color = cell.ColorRed
		case syncProgress < metrics.TotalTransactionsSyncPercentage:
			color = cell.ColorYellow
		}

		return []gauge.Option{gauge.Color(color), gauge.Border(linestyle.Light, cell.FgColor(color))}
	case enums.CellNameCheckSyncProgress:
		var (
			syncProgress = node.Metrics.CheckSyncPercentage
			color        = cell.ColorGreen
		)

		switch {
		case syncProgress == 0:
			color = cell.ColorRed
		case syncProgress < metrics.TotalCheckpointsSyncPercentage:
			color = cell.ColorYellow
		}

		return []gauge.Option{gauge.Color(color), gauge.Border(linestyle.Light, cell.FgColor(color))}
	case enums.CellNameEpochProgress, enums.CellNameDiskUsage, enums.CellNameCpuUsage, enums.CellNameMemoryUsage:
		options = append(options, cell.Bold())
	case enums.CellNameEpochTimeTillTheEnd, enums.CellNameDatabaseSize, enums.CellNameBytesReceived, enums.CellNameBytesSent:
		return []segmentdisplay.WriteOption{segmentdisplay.WriteCellOpts(cell.FgColor(cell.ColorWhite)), segmentdisplay.WriteCellOpts(cell.FgColor(cell.ColorGreen))}
	case enums.CellNameUptime:
		var (
			uptime = node.Metrics.Uptime
			color  = cell.ColorWhite
		)

		switch {
		case uptime == "":
		case uptime == "0.0":
			color = cell.ColorRed
		case uptime < "1.0":
			color = cell.ColorYellow
		}

		return []segmentdisplay.WriteOption{segmentdisplay.WriteCellOpts(cell.FgColor(color)), segmentdisplay.WriteCellOpts(cell.FgColor(cell.ColorGreen))}
	case enums.CellNameConnectedPeers:
		var (
			peers = node.Metrics.NetworkPeers
			color = cell.ColorWhite
		)

		switch {
		case peers == 0:
			color = cell.ColorRed
		case peers < 3:
			color = cell.ColorYellow
		}

		return []segmentdisplay.WriteOption{segmentdisplay.WriteCellOpts(cell.FgColor(color))}
	case enums.CellNameNodeLogs:
		options = append(options, cell.FgColor(cell.ColorWhite), cell.Bold())
	case enums.CellNameTPSTracker:
		var (
			tpsNode = node.Metrics.TransactionsPerSecond
			tpsRpc  = rpc.Metrics.TransactionsPerSecond
			color   = cell.ColorGreen
		)

		if tpsNode < tpsRpc-metrics.TransactionsPerSecondLag {
			color = cell.ColorYellow
		}

		return []sparkline.Option{sparkline.Color(color)}
	case enums.CellNameCPSTracker:
		var (
			checkNode = node.Metrics.CheckpointsPerSecond
			checkRpc  = rpc.Metrics.CheckpointsPerSecond
			color     = cell.ColorBlue
		)

		if checkNode < checkRpc-metrics.CheckpointsPerSecondLag {
			color = cell.ColorYellow
		}

		return []sparkline.Option{sparkline.Color(color)}
	default:
		return []segmentdisplay.WriteOption{segmentdisplay.WriteCellOpts(cell.FgColor(cell.ColorWhite))}
	}

	return options
}
