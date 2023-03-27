package checker

import (
	"context"
	"net/url"
	"strings"
	"time"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/widgets/gauge"
	"github.com/mum4k/termdash/widgets/segmentdisplay"
	"github.com/mum4k/termdash/widgets/sparkline"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/ybbus/jsonrpc/v3"

	"github.com/bartosian/sui_helpers/suimon/cmd/checker/dashboardbuilder/dashboards"
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/pkg/metricsparser"
	"github.com/bartosian/sui_helpers/suimon/internal/pkg/utility"
)

var prometheusMetrics = map[enums.PrometheusMetricName]metricsparser.MetricConfig{
	enums.PrometheusMetricNameTotalTransactionCertificates: {
		MetricType: enums.PrometheusMetricTypeCounter,
	},
	enums.PrometheusMetricNameTotalTransactionEffects: {
		MetricType: enums.PrometheusMetricTypeCounter,
	},
	enums.PrometheusMetricNameHighestKnownCheckpoint: {
		MetricType: enums.PrometheusMetricTypeGauge,
	},
	enums.PrometheusMetricNameHighestSyncedCheckpoint: {
		MetricType: enums.PrometheusMetricTypeGauge,
	},
	enums.PrometheusMetricNameLastExecutedCheckpoint: {
		MetricType: enums.PrometheusMetricTypeGauge,
	},
	enums.PrometheusMetricNameCurrentEpoch: {
		MetricType: enums.PrometheusMetricTypeGauge,
	},
	enums.PrometheusMetricNameEpochTotalDuration: {
		MetricType: enums.PrometheusMetricTypeGauge,
	},
	enums.PrometheusMetricNameCurrentRound: {
		MetricType: enums.PrometheusMetricTypeGauge,
	},
	enums.PrometheusMetricNameHighestProcessedRound: {
		MetricType: enums.PrometheusMetricTypeGauge,
		Labels: prometheus.Labels{
			"source": "own",
		},
	},
	enums.PrometheusMetricNameLastCommittedRound: {
		MetricType: enums.PrometheusMetricTypeGauge,
	},
	enums.PrometheusMetricNamePrimaryNetworkPeers: {
		MetricType: enums.PrometheusMetricTypeGauge,
	},
	enums.PrometheusMetricNameWorkerNetworkPeers: {
		MetricType: enums.PrometheusMetricTypeGauge,
	},
	enums.PrometheusMetricNameSuiNetworkPeers: {
		MetricType: enums.PrometheusMetricTypeGauge,
	},
	enums.PrometheusMetricNameSkippedConsensusTransactions: {
		MetricType: enums.PrometheusMetricTypeCounter,
	},
	enums.PrometheusMetricNameTotalSignatureErrors: {
		MetricType: enums.PrometheusMetricTypeCounter,
	},
	enums.PrometheusMetricNameUptime: {
		MetricType: enums.PrometheusMetricTypeCounter,
	},
}

// GetPrometheusMetrics returns a "Metrics" object representing the current state of network checks for the "Host" object
// passed as a pointer receiver. This object contains the results of each metric check performed on the host,
// including the number of successful checks, the total number of checks performed, and the percentage progress
// for each metric.
// Parameters: None.
// Returns: - a "Metrics" object representing the current state of network checks for the "Host" object.
func (host *Host) GetPrometheusMetrics() {
	metricsURL := host.getUrl(requestTypeMetrics, false)
	parser := metricsparser.NewPrometheusMetricParser(host.httpClient, metricsURL, prometheusMetrics)

	result, err := parser.GetMetrics()
	if err != nil {
		return
	}

	for metricName, metricValue := range result {
		switch metricName {
		case enums.PrometheusMetricNameTotalTransactionCertificates:
			host.Metrics.SetValue(enums.MetricTypeTotalTransactionCertificates, metricValue.Value)
		case enums.PrometheusMetricNameTotalTransactionEffects:
			host.Metrics.SetValue(enums.MetricTypeTotalTransactionEffects, metricValue.Value)
		case enums.PrometheusMetricNameHighestKnownCheckpoint:
			host.Metrics.SetValue(enums.MetricTypeHighestKnownCheckpoint, metricValue.Value)
		case enums.PrometheusMetricNameHighestSyncedCheckpoint:
			host.Metrics.SetValue(enums.MetricTypeHighestSyncedCheckpoint, metricValue.Value)
		case enums.PrometheusMetricNameLastExecutedCheckpoint:
			host.Metrics.SetValue(enums.MetricTypeLastExecutedCheckpoint, metricValue.Value)
		case enums.PrometheusMetricNameCurrentEpoch:
			host.Metrics.SetValue(enums.MetricTypeCurrentEpoch, metricValue.Value)
		case enums.PrometheusMetricNameEpochTotalDuration:
			host.Metrics.SetValue(enums.MetricTypeEpochTotalDuration, metricValue.Value)
		case enums.PrometheusMetricNameCurrentRound:
			host.Metrics.SetValue(enums.MetricTypeCurrentRound, metricValue.Value)
		case enums.PrometheusMetricNameHighestProcessedRound:
			host.Metrics.SetValue(enums.MetricTypeHighestProcessedRound, metricValue.Value)
		case enums.PrometheusMetricNameLastCommittedRound:
			host.Metrics.SetValue(enums.MetricTypeLastCommittedRound, metricValue.Value)
		case enums.PrometheusMetricNamePrimaryNetworkPeers:
			host.Metrics.SetValue(enums.MetricTypePrimaryNetworkPeers, metricValue.Value)
		case enums.PrometheusMetricNameWorkerNetworkPeers:
			host.Metrics.SetValue(enums.MetricTypeWorkerNetworkPeers, metricValue.Value)
		case enums.PrometheusMetricNameSuiNetworkPeers:
			host.Metrics.SetValue(enums.MetricTypeSuiNetworkPeers, metricValue.Value)
		case enums.PrometheusMetricNameSkippedConsensusTransactions:
			host.Metrics.SetValue(enums.MetricTypeSkippedConsensusTransactions, metricValue.Value)
		case enums.PrometheusMetricNameTotalSignatureErrors:
			host.Metrics.SetValue(enums.MetricTypeTotalSignatureErrors, metricValue.Value)
		case enums.PrometheusMetricNameUptime:
			host.Metrics.SetValue(enums.MetricTypeUptime, metricValue.Value)

			if value, ok := metricValue.Labels["version"]; ok {
				versionInfo := strings.Split(value, "-")

				host.Metrics.SetValue(enums.MetricTypeVersion, versionInfo[0])

				if len(versionInfo) == 2 {
					host.Metrics.SetValue(enums.MetricTypeCommit, versionInfo[1])
				}
			}
		}
	}
}

// GetTotalTransactionNumber returns the total number of transactions performed on the "Host" object passed
// as a pointer receiver. This method retrieves the "Metrics" object for the host and calculates the total
// number of transactions performed across all metric types.
// Parameters: None.
// Returns: an integer representing the total number of transactions performed on the "Host" object.
func (host *Host) GetTotalTransactionNumber() {
	var result any

	if result = getFromRPC(host.rpcHttpClient, enums.RPCMethodGetTotalTransactionNumber); result == nil {
		if result = getFromRPC(host.rpcHttpsClient, enums.RPCMethodGetTotalTransactionNumber); result == nil {
			return
		}
	}

	host.Metrics.SetValue(enums.MetricTypeTotalTransactions, result)
}

// GetLatestCheckpoint returns a "Checkpoint" object representing the most recent checkpoint for the "Host"
// object passed as a pointer receiver. This object contains information about the time and status of the
// most recent checkpoint performed on the host.
// Parameters: None.
// Returns: a "Checkpoint" object representing the most recent checkpoint for the "Host" object.
func (host *Host) GetLatestCheckpoint() {
	var result any

	if result = getFromRPC(host.rpcHttpClient, enums.RPCMethodGetLatestCheckpointSequenceNumber); result == nil {
		if result = getFromRPC(host.rpcHttpsClient, enums.RPCMethodGetLatestCheckpointSequenceNumber); result == nil {
			return
		}
	}

	host.Metrics.SetValue(enums.MetricTypeLatestCheckpoint, result)
}

// GetLatestSuiSystemState returns a "SUISystemState" object representing the current system state for the "Host"
// object passed as a pointer receiver. This object contains information about the status of various components
// of the SUISystem software running on the host.
// Parameters: None.
// Returns: a "SUISystemState" object representing the current system state for the "Host" object.
func (host *Host) GetLatestSuiSystemState() {
	var result any

	if result = getFromRPC(host.rpcHttpClient, enums.RPCMethodGetSuiSystemState); result == nil {
		if result = getFromRPC(host.rpcHttpsClient, enums.RPCMethodGetSuiSystemState); result == nil {
			return
		}
	}

	host.Metrics.SetValue(enums.MetricTypeSuiSystemState, result)
}

// getFromRPC makes a JSON-RPC call to the specified method using the provided RPC client, and returns
// the result of the call. The returned type will depend on the specific method called and its response.
// Parameters:
// - rpcClient: a jsonrpc.RPCClient representing the client to use for the JSON-RPC call.
// - method: an enums.RPCMethod representing the name of the JSON-RPC method to call.
// Returns:
//   - the result of the JSON-RPC call. The specific type of the returned value will depend on the method called
//     and its response.
func getFromRPC(rpcClient jsonrpc.RPCClient, method enums.RPCMethod) any {
	var (
		respChan = make(chan any)
		timeout  = time.After(rpcClientTimeout)
	)

	switch method {
	case enums.RPCMethodGetSuiSystemState:
		var response SuiSystemState

		go func() {
			if err := rpcClient.CallFor(context.Background(), &response, method.String()); err != nil {
				return
			}

			respChan <- response
		}()
	default:
		var response int

		go func() {
			if err := rpcClient.CallFor(context.Background(), &response, method.String()); err != nil {
				return
			}

			respChan <- response
		}()
	}

	select {
	case response := <-respChan:
		return response
	case <-timeout:
		return nil
	}
}

// GetData returns a "HostData" object representing the current data for the "Host" object passed as a pointer receiver.
// This object contains various metrics and status information about the host.
// Parameters: None.
// Returns: a "HostData" object representing the current data for the "Host" object.
func (host *Host) GetData() {
	doneCH := make(chan struct{})

	defer close(doneCH)

	go func() {
		host.GetTotalTransactionNumber()

		doneCH <- struct{}{}
	}()

	go func() {
		host.GetLatestCheckpoint()

		doneCH <- struct{}{}
	}()

	go func() {
		host.GetPrometheusMetrics()

		doneCH <- struct{}{}
	}()

	for i := 0; i < 3; i++ {
		<-doneCH
	}
}

// getUrl returns the URL for a given request type and security setting.
// Parameters:
// - request: a requestType indicating the type of request for which to generate a URL.
// - secure: a boolean indicating whether the generated URL should use HTTPS (true) or HTTP (false).
// Returns: a string representing the URL for the specified request type and security setting.
func (host *Host) getUrl(request requestType, secure bool) string {
	var (
		protocol = "http"
		hostPort = host.HostPort
		hostUrl  = new(url.URL)
	)

	if hostPort.Host != nil {
		hostUrl.Host = *hostPort.Host
	} else {
		hostUrl.Host = *hostPort.IP
	}

	if hostPort.Path != nil {
		hostUrl.Path = *hostPort.Path
	}

	if secure {
		protocol = protocol + "s"
	}

	hostUrl.Scheme = protocol

	switch request {
	case requestTypeRPC:
		if port, ok := host.Ports[enums.PortTypeRPC]; ok {
			hostUrl.Host = hostUrl.Hostname() + ":" + port
		} else {
			hostUrl.Host = hostUrl.Hostname() + ":" + rpcPortDefault
		}
	case requestTypeMetrics:
		hostUrl.Path = "/metrics"

		if port, ok := host.Ports[enums.PortTypeMetrics]; ok {
			hostUrl.Host = hostUrl.Hostname() + ":" + port
		} else {
			hostUrl.Host = hostUrl.Hostname() + ":" + metricsPortDefault
		}
	}

	return hostUrl.String()
}

// getMetricForDashboardCell retrieves the current metric data for the specified dashboard cell.
// Parameters: cellName: a dashboards.CellName representing the name of the dashboard cell for which to retrieve data.
// Returns: a value of any type representing the current metric data for the specified dashboard cell.
func (checker *Checker) getMetricForDashboardCell(cellName enums.CellName) any {
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
		if len(node.Metrics.TransactionsHistory) < transactionsPerSecondWindow {
			return dashboards.DashboardLoadingBlinkValue()
		}

		fallthrough
	case enums.CellNameTPSTracker:
		return node.Metrics.TransactionsPerSecond
	case enums.CellNameCheckpointsPerSecond:
		if len(node.Metrics.CheckpointsHistory) < checkpointsPerSecondWindow {
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
		usageLabel, usagePercentage := getDonutUsageMetric("GB", utility.GetDiskUsage)

		return dashboards.NewDonutInput(usageLabel, usagePercentage)
	case enums.CellNameDatabaseSize:
		dbSize := getDirectorySize(checker.suimonConfig.DbPath)

		return dbSize
	case enums.CellNameBytesSent:
		bytesSent := getNetworkUsageMetric(enums.CellNameBytesSent)

		return bytesSent
	case enums.CellNameBytesReceived:
		bytesReceived := getNetworkUsageMetric(enums.CellNameBytesReceived)

		return bytesReceived
	case enums.CellNameMemoryUsage:
		usageLabel, usagePercentage := getDonutUsageMetric("%", utility.GetMemoryUsage)

		return dashboards.NewDonutInput(usageLabel, usagePercentage)
	case enums.CellNameCpuUsage:
		usageLabel, usagePercentage := getDonutUsageMetric("%", utility.GetCPUUsage)

		return dashboards.NewDonutInput(usageLabel, usagePercentage)
	default:
		return ""
	}
}

// getOptionsForDashboardCell retrieves the options data for a specific dashboard cell and returns it in the appropriate format for display in the dashboard.
// Parameters: cellName: a dashboards.CellName representing the name of the dashboard cell for which to retrieve options data.
// Returns: the options data for the specified dashboard cell in the appropriate format for display in the dashboard.
func (checker *Checker) getOptionsForDashboardCell(cellName enums.CellName) any {
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
		case txSyncPercentageNode < totalTransactionsSyncPercentage || transactionsNode < transactionsRpc-totalTransactionsLag:
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
		case latestCheckpointNode < latestCheckpointRpc-latestCheckpointLag:
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
		case highestCheckpointNode < highestCheckpointRpc-highestSyncedCheckpointLag || highestCheckpointNode < latestCheckpointRpc-latestCheckpointLag:
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
		case len(txHistoryNode) != transactionsPerSecondWindow:
		case tpsNode == 0:
			color = cell.ColorRed
		case tpsNode < tpsRpc-transactionsPerSecondLag:
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
		case len(checkHistoryNode) != checkpointsPerSecondWindow:
		case checkNode == 0:
			color = cell.ColorRed
		case checkNode < checkRpc-checkpointsPerSecondLag:
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
		case syncProgress < totalTransactionsSyncPercentage:
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
		case syncProgress < totalCheckpointsSyncPercentage:
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

		if tpsNode < tpsRpc-transactionsPerSecondLag {
			color = cell.ColorYellow
		}

		return []sparkline.Option{sparkline.Color(color)}
	case enums.CellNameCPSTracker:
		var (
			checkNode = node.Metrics.CheckpointsPerSecond
			checkRpc  = rpc.Metrics.CheckpointsPerSecond
			color     = cell.ColorBlue
		)

		if checkNode < checkRpc-checkpointsPerSecondLag {
			color = cell.ColorYellow
		}

		return []sparkline.Option{sparkline.Color(color)}
	default:
		return []segmentdisplay.WriteOption{segmentdisplay.WriteCellOpts(cell.FgColor(cell.ColorWhite))}
	}

	return options
}
