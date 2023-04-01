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

type (
	DashboardMetric  interface{}
	DashboardOptions interface{}
)

// getMetricForDashboardCell returns the DashboardMetric object for the specified dashboard cell name.
// The function searches through the CheckerController's internal state to find the DashboardMetric object that corresponds to the specified cell name, and returns that object.
// Returns a DashboardMetric object.
func (checker *CheckerController) getMetricForDashboardCell(cellName enums.CellName) DashboardMetric {
	var (
		node = checker.hosts.node[0]
		rpc  = checker.hosts.rpc[0]
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

		return node.Metrics.TransactionsPerSecond
	case enums.CellNameTPSTracker:
		return node.Metrics.TransactionsPerSecond
	case enums.CellNameCheckpointsPerSecond:
		if len(node.Metrics.CheckpointsHistory) < metrics.CheckpointsPerSecondWindow {
			return dashboards.DashboardLoadingBlinkValue()
		}

		return node.Metrics.CheckpointsPerSecond
	case enums.CellNameCPSTracker:
		return node.Metrics.CheckpointsPerSecond
	case enums.CellNameTotalTransactions:
		return node.Metrics.TotalTransactionsBlocks
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
		return rpc.Metrics.GetTimeUntilNextEpochDisplay()
	case enums.CellNameDiskUsage:
		usageLabel, usagePercentage := metrics.GetUsageDataForDonutChart(enums.MetricUnitGB, utility.GetDiskUsage)

		return dashboards.NewDonutInput(usageLabel, usagePercentage)
	case enums.CellNameDatabaseSize:
		dbSize := metrics.GetFileSize(checker.suimonConfig.DbPath)

		return dbSize
	case enums.CellNameBytesSent:
		bytesSent := metrics.GetFormattedNetworkUsage(enums.CellNameBytesSent)

		return bytesSent
	case enums.CellNameBytesReceived:
		bytesReceived := metrics.GetFormattedNetworkUsage(enums.CellNameBytesReceived)

		return bytesReceived
	case enums.CellNameMemoryUsage:
		usageLabel, usagePercentage := metrics.GetUsageDataForDonutChart(enums.MetricUnitPercentage, utility.GetMemoryUsage)

		return dashboards.NewDonutInput(usageLabel, usagePercentage)
	case enums.CellNameCpuUsage:
		usageLabel, usagePercentage := metrics.GetUsageDataForDonutChart(enums.MetricUnitPercentage, utility.GetCPUUsage)

		return dashboards.NewDonutInput(usageLabel, usagePercentage)
	default:
		return ""
	}
}

// getOptionsForDashboardCell returns the DashboardOptions object for the specified dashboard cell name.
// The function searches through the CheckerController's internal state to find the DashboardOptions object that corresponds to the specified cell name, and returns that object.
// Returns a DashboardOptions object.
func (checker *CheckerController) getOptionsForDashboardCell(cellName enums.CellName) DashboardOptions {
	var (
		node    = checker.hosts.node[0]
		rpc     = checker.hosts.rpc[0]
		options []cell.Option
	)

	switch cellName {
	case enums.CellNameNodeHealth:
		status := node.Status
		color := cell.ColorGreen

		switch status {
		case enums.StatusYellow:
			color = cell.ColorYellow
		case enums.StatusRed:
			color = cell.ColorRed
		}

		options = append(options, cell.BgColor(color), cell.FgColor(color))
	case enums.CellNameNetworkHealth:
		status := rpc.Status
		color := cell.ColorGreen

		switch status {
		case enums.StatusYellow:
			color = cell.ColorYellow
		case enums.StatusRed:
			color = cell.ColorRed
		}

		options = append(options, cell.BgColor(color), cell.FgColor(color))
	case enums.CellNameTotalTransactions:
		transactionsNode := node.Metrics.TotalTransactionsBlocks
		txSyncPercentageNode := node.Metrics.TxSyncPercentage
		transactionsRpc := rpc.Metrics.TotalTransactionsBlocks
		color := cell.ColorWhite

		switch {
		case transactionsNode == 0:
			color = cell.ColorRed
		case txSyncPercentageNode < metrics.TotalTransactionsSyncPercentage || transactionsNode < transactionsRpc-metrics.TotalTransactionsLag:
			color = cell.ColorYellow
		}

		return []segmentdisplay.WriteOption{segmentdisplay.WriteCellOpts(cell.FgColor(color))}
	case enums.CellNameLatestCheckpoint:
		latestCheckpointNode := node.Metrics.LatestCheckpoint
		latestCheckpointRpc := rpc.Metrics.LatestCheckpoint
		color := cell.ColorWhite

		switch {
		case latestCheckpointNode == 0:
			color = cell.ColorRed
		case latestCheckpointNode < latestCheckpointRpc-metrics.LatestCheckpointLag:
			color = cell.ColorYellow
		}

		return []segmentdisplay.WriteOption{segmentdisplay.WriteCellOpts(cell.FgColor(color))}
	case enums.CellNameHighestSyncedCheckpoint:
		highestCheckpointNode := node.Metrics.HighestSyncedCheckpoint
		highestCheckpointRpc := rpc.Metrics.HighestSyncedCheckpoint
		latestCheckpointRpc := rpc.Metrics.LatestCheckpoint
		color := cell.ColorWhite

		switch {
		case highestCheckpointNode == 0:
			color = cell.ColorRed
		case highestCheckpointNode < highestCheckpointRpc-metrics.HighestSyncedCheckpointLag || highestCheckpointNode < latestCheckpointRpc-metrics.LatestCheckpointLag:
			color = cell.ColorYellow
		}

		return []segmentdisplay.WriteOption{segmentdisplay.WriteCellOpts(cell.FgColor(color))}
	case enums.CellNameTransactionsPerSecond:
		tpsNode := node.Metrics.TransactionsPerSecond
		txHistoryNode := node.Metrics.TransactionsHistory
		tpsRpc := rpc.Metrics.TransactionsPerSecond
		color := cell.ColorWhite

		switch {
		case len(txHistoryNode) != metrics.TransactionsPerSecondWindow:
		case tpsNode == 0:
			color = cell.ColorRed
		case tpsNode < tpsRpc-metrics.TransactionsPerSecondLag:
			color = cell.ColorYellow
		}

		return []segmentdisplay.WriteOption{segmentdisplay.WriteCellOpts(cell.FgColor(color))}
	case enums.CellNameCheckpointsPerSecond:
		checkNode := node.Metrics.CheckpointsPerSecond
		checkHistoryNode := node.Metrics.CheckpointsHistory
		checkRpc := rpc.Metrics.CheckpointsPerSecond
		color := cell.ColorWhite

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
		syncProgress := node.Metrics.CheckSyncPercentage
		color := cell.ColorGreen

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
		uptime := node.Metrics.Uptime
		color := cell.ColorWhite

		switch {
		case uptime == "":
		case uptime == "0.0":
			color = cell.ColorRed
		case uptime < "1.0":
			color = cell.ColorYellow
		}

		return []segmentdisplay.WriteOption{segmentdisplay.WriteCellOpts(cell.FgColor(color)), segmentdisplay.WriteCellOpts(cell.FgColor(cell.ColorGreen))}
	case enums.CellNameConnectedPeers:
		peers := node.Metrics.NetworkPeers
		color := cell.ColorWhite

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
		tpsNode := node.Metrics.TransactionsPerSecond
		tpsRpc := rpc.Metrics.TransactionsPerSecond
		color := cell.ColorGreen

		if tpsNode < tpsRpc-metrics.TransactionsPerSecondLag {
			color = cell.ColorYellow
		}

		return []sparkline.Option{sparkline.Color(color)}
	case enums.CellNameCPSTracker:
		checkNode := node.Metrics.CheckpointsPerSecond
		checkRpc := rpc.Metrics.CheckpointsPerSecond
		color := cell.ColorBlue

		if checkNode < checkRpc-metrics.CheckpointsPerSecondLag {
			color = cell.ColorYellow
		}

		return []sparkline.Option{sparkline.Color(color)}
	default:
		return []segmentdisplay.WriteOption{segmentdisplay.WriteCellOpts(cell.FgColor(cell.ColorWhite))}
	}

	return options
}
