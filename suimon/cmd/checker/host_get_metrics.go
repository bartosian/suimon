package checker

import (
	"bufio"
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/widgets/gauge"
	"github.com/mum4k/termdash/widgets/segmentdisplay"
	"github.com/mum4k/termdash/widgets/sparkline"
	"github.com/ybbus/jsonrpc/v3"

	"github.com/bartosian/sui_helpers/suimon/cmd/checker/dashboardbuilder/dashboards"
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
	"github.com/bartosian/sui_helpers/suimon/pkg/utility"
)

// GetMetrics returns a "Metrics" object representing the current state of network checks for the "Host" object
// passed as a pointer receiver. This object contains the results of each metric check performed on the host,
// including the number of successful checks, the total number of checks performed, and the percentage progress
// for each metric.
// Parameters: None.
// Returns: - a "Metrics" object representing the current state of network checks for the "Host" object.
func (host *Host) GetMetrics() {
	metricsURL := host.getUrl(requestTypeMetrics, false)

	result, err := host.httpClient.Get(metricsURL)
	if err != nil {
		return
	}

	defer result.Body.Close()

	reader := bufio.NewReader(result.Body)
	for {
		line, err := reader.ReadString('\n')
		if len(line) == 0 && err != nil {
			break
		}

		if strings.HasPrefix(line, "#") {
			continue
		}

		metric := strings.Split(line, " ")
		if len(metric) != 2 {
			continue
		}

		key, value := strings.TrimSpace(metric[0]), strings.TrimSpace(metric[1])

		metricName, err := enums.MetricTypeFromString(key)
		if err != nil {
			continue
		}

		if metricName == enums.MetricTypeUptime {
			versionMetric := versionRegex.FindStringSubmatch(key)
			version := strings.Split(versionMetric[1], "=")

			uptimeSeconds, err := strconv.Atoi(value)
			if err != nil {
				continue
			}

			value = fmt.Sprintf("%.1f days", float64(uptimeSeconds)/(60*60*24))

			versionInfo := strings.Split(version[1], "-")
			if len(versionInfo) != 2 {
				continue
			}

			host.Metrics.SetValue(enums.MetricTypeVersion, versionInfo[0])
			host.Metrics.SetValue(enums.MetricTypeCommit, versionInfo[1])
		}

		host.Metrics.SetValue(metricName, value)
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

	host.Metrics.SetValue(enums.MetricTypeTotalTransactionsNumber, result)
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

// GetSUISystemState returns a "SUISystemState" object representing the current system state for the "Host"
// object passed as a pointer receiver. This object contains information about the status of various components
// of the SUISystem software running on the host.
// Parameters: None.
// Returns: a "SUISystemState" object representing the current system state for the "Host" object.
func (host *Host) GetSUISystemState() {
	var result any

	if result = getFromRPC(host.rpcHttpClient, enums.RPCMethodGetSuiSystemState); result == nil {
		if result = getFromRPC(host.rpcHttpsClient, enums.RPCMethodGetSuiSystemState); result == nil {
			return
		}
	}

	host.Metrics.SetValue(enums.MetricTypeCurrentEpoch, result)
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
		host.GetSUISystemState()

		doneCH <- struct{}{}
	}()

	go func() {
		host.GetLatestCheckpoint()

		doneCH <- struct{}{}
	}()

	go func() {
		host.GetMetrics()

		doneCH <- struct{}{}
	}()

	for i := 0; i < 4; i++ {
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
		} else if hostPort.Host == nil {
			hostUrl.Host = hostUrl.Hostname() + ":" + rpcPortDefault
		}
	case requestTypeMetrics:
		fallthrough
	default:
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
func (checker *Checker) getMetricForDashboardCell(cellName dashboards.CellName) any {
	var (
		node = checker.node[0]
		rpc  = checker.rpc[0]
	)

	switch cellName {
	case dashboards.CellNameNodeStatus:
		return node.Status.DashboardStatus()
	case dashboards.CellNameNetworkStatus:
		return rpc.Status.DashboardStatus()
	case dashboards.CellNameTransactionsPerSecond:
		if len(node.Metrics.TransactionsHistory) < transactionsPerSecondTimeout {
			return dashboards.DashboardLoadingBlinkValue()
		}

		fallthrough
	case dashboards.CellNameTPSTracker:
		return node.Metrics.TransactionsPerSecond
	case dashboards.CellNameCheckpointsPerSecond:
		if len(node.Metrics.CheckpointsHistory) < checkpointsPerSecondTimeout {
			return dashboards.DashboardLoadingBlinkValue()
		}

		fallthrough
	case dashboards.CellNameCPSTracker:
		return node.Metrics.CheckpointsPerSecond
	case dashboards.CellNameTotalTransactions:
		return node.Metrics.TotalTransactionNumber
	case dashboards.CellNameLatestCheckpoint:
		return node.Metrics.LatestCheckpoint
	case dashboards.CellNameHighestCheckpoint:
		return node.Metrics.HighestSyncedCheckpoint
	case dashboards.CellNameConnectedPeers:
		return node.Metrics.SuiNetworkPeers
	case dashboards.CellNameTXSyncProgress:
		return node.Metrics.TxSyncPercentage
	case dashboards.CellNameCheckSyncProgress:
		return node.Metrics.CheckSyncPercentage
	case dashboards.CellNameUptime:
		return []string{strings.Split(node.Metrics.Uptime, " ")[0], "D"}
	case dashboards.CellNameVersion:
		return node.Metrics.Version
	case dashboards.CellNameCommit:
		return node.Metrics.Commit
	case dashboards.CellNameEpochProgress:
		epochLabel := rpc.Metrics.GetEpochLabel()
		epochPercentage := rpc.Metrics.GetEpochProgress()

		return dashboards.NewDonutInput(epochLabel, epochPercentage)
	case dashboards.CellNameCurrentEpoch:
		return rpc.Metrics.SystemState.Epoch
	case dashboards.CellNameEpochEnd:
		return rpc.Metrics.GetEpochTimer()
	case dashboards.CellNameDiskUsage:
		usageLabel, usagePercentage := getDonutUsageMetric("GB", utility.GetDiskUsage)

		return dashboards.NewDonutInput(usageLabel, usagePercentage)
	case dashboards.CellNameDatabaseSize:
		dbSize := getDirectorySize(checker.nodeConfig.DbPath)

		return dbSize
	case dashboards.CellNameBytesSent:
		bytesSent := getNetworkUsageMetric(dashboards.CellNameBytesSent)

		return bytesSent
	case dashboards.CellNameBytesReceived:
		bytesReceived := getNetworkUsageMetric(dashboards.CellNameBytesReceived)

		return bytesReceived
	case dashboards.CellNameMemoryUsage:
		usageLabel, usagePercentage := getDonutUsageMetric("%", utility.GetMemoryUsage)

		return dashboards.NewDonutInput(usageLabel, usagePercentage)
	case dashboards.CellNameCpuUsage:
		usageLabel, usagePercentage := getDonutUsageMetric("%", utility.GetCPUUsage)

		return dashboards.NewDonutInput(usageLabel, usagePercentage)
	default:
		return ""
	}
}

// getOptionsForDashboardCell retrieves the options data for a specific dashboard cell and returns it in the appropriate format for display in the dashboard.
// Parameters: cellName: a dashboards.CellName representing the name of the dashboard cell for which to retrieve options data.
// Returns: the options data for the specified dashboard cell in the appropriate format for display in the dashboard.
func (checker *Checker) getOptionsForDashboardCell(cellName dashboards.CellName) any {
	var (
		node    = checker.node[0]
		rpc     = checker.rpc[0]
		options []cell.Option
	)

	switch cellName {
	case dashboards.CellNameNodeStatus:
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
	case dashboards.CellNameNetworkStatus:
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
	case dashboards.CellNameTotalTransactions:
		var (
			transactionsNode     = node.Metrics.TotalTransactionNumber
			txSyncPercentageNode = node.Metrics.TxSyncPercentage
			transactionsRpc      = rpc.Metrics.TotalTransactionNumber
			color                = cell.ColorWhite
		)

		switch {
		case transactionsNode == 0:
			color = cell.ColorRed
		case txSyncPercentageNode < totalTransactionsSyncPercentage || transactionsNode < transactionsRpc-totalTransactionsLag:
			color = cell.ColorYellow
		}

		return []segmentdisplay.WriteOption{segmentdisplay.WriteCellOpts(cell.FgColor(color))}
	case dashboards.CellNameLatestCheckpoint:
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
	case dashboards.CellNameHighestCheckpoint:
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
	case dashboards.CellNameTransactionsPerSecond:
		var (
			tpsNode       = node.Metrics.TransactionsPerSecond
			txHistoryNode = node.Metrics.TransactionsHistory
			tpsRpc        = rpc.Metrics.TransactionsPerSecond
			color         = cell.ColorWhite
		)

		switch {
		case len(txHistoryNode) != transactionsPerSecondTimeout:
		case tpsNode == 0:
			color = cell.ColorRed
		case tpsNode < tpsRpc-transactionsPerSecondLag:
			color = cell.ColorYellow
		}

		return []segmentdisplay.WriteOption{segmentdisplay.WriteCellOpts(cell.FgColor(color))}
	case dashboards.CellNameCheckpointsPerSecond:
		var (
			checkNode        = node.Metrics.CheckpointsPerSecond
			checkHistoryNode = node.Metrics.CheckpointsHistory
			checkRpc         = rpc.Metrics.CheckpointsPerSecond
			color            = cell.ColorWhite
		)

		switch {
		case len(checkHistoryNode) != checkpointsPerSecondTimeout:
		case checkNode == 0:
			color = cell.ColorRed
		case checkNode < checkRpc-checkpointsPerSecondLag:
			color = cell.ColorYellow
		}

		return []segmentdisplay.WriteOption{segmentdisplay.WriteCellOpts(cell.FgColor(color))}
	case dashboards.CellNameTXSyncProgress:
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
	case dashboards.CellNameCheckSyncProgress:
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
	case dashboards.CellNameEpochProgress, dashboards.CellNameDiskUsage, dashboards.CellNameCpuUsage, dashboards.CellNameMemoryUsage:
		options = append(options, cell.Bold())
	case dashboards.CellNameEpochEnd, dashboards.CellNameDatabaseSize, dashboards.CellNameBytesReceived, dashboards.CellNameBytesSent:
		return []segmentdisplay.WriteOption{segmentdisplay.WriteCellOpts(cell.FgColor(cell.ColorWhite)), segmentdisplay.WriteCellOpts(cell.FgColor(cell.ColorGreen))}
	case dashboards.CellNameUptime:
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
	case dashboards.CellNameConnectedPeers:
		var (
			peers = node.Metrics.SuiNetworkPeers
			color = cell.ColorWhite
		)

		switch {
		case peers == 0:
			color = cell.ColorRed
		case peers < 3:
			color = cell.ColorYellow
		}

		return []segmentdisplay.WriteOption{segmentdisplay.WriteCellOpts(cell.FgColor(color))}
	case dashboards.CellNameNodeLogs:
		options = append(options, cell.FgColor(cell.ColorWhite), cell.Bold())
	case dashboards.CellNameTPSTracker:
		var (
			tpsNode = node.Metrics.TransactionsPerSecond
			tpsRpc  = rpc.Metrics.TransactionsPerSecond
			color   = cell.ColorGreen
		)

		if tpsNode < tpsRpc-transactionsPerSecondLag {
			color = cell.ColorYellow
		}

		return []sparkline.Option{sparkline.Color(color)}
	case dashboards.CellNameCPSTracker:
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
