package checker

import (
	"bufio"
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/ybbus/jsonrpc/v3"

	"github.com/bartosian/sui_helpers/suimon/cmd/checker/dashboardbuilder/dashboards"
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
	"github.com/bartosian/sui_helpers/suimon/pkg/utility"
	"github.com/mum4k/termdash/cell"
)

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

func (host *Host) GetTotalTransactionNumber() {
	var result any

	if result = getFromRPC(host.rpcHttpClient, enums.RPCMethodGetTotalTransactionNumber); result == nil {
		if result = getFromRPC(host.rpcHttpsClient, enums.RPCMethodGetTotalTransactionNumber); result == nil {
			return
		}
	}

	host.Metrics.SetValue(enums.MetricTypeTotalTransactionsNumber, result)
}

func (host *Host) GetLatestCheckpoint() {
	var result any

	if result = getFromRPC(host.rpcHttpClient, enums.RPCMethodGetLatestCheckpointSequenceNumber); result == nil {
		if result = getFromRPC(host.rpcHttpsClient, enums.RPCMethodGetLatestCheckpointSequenceNumber); result == nil {
			return
		}
	}

	host.Metrics.SetValue(enums.MetricTypeLatestCheckpoint, result)
}

func (host *Host) GetSUISystemState() {
	var result any

	if result = getFromRPC(host.rpcHttpClient, enums.RPCMethodGetSuiSystemState); result == nil {
		if result = getFromRPC(host.rpcHttpsClient, enums.RPCMethodGetSuiSystemState); result == nil {
			return
		}
	}

	host.Metrics.SetValue(enums.MetricTypeCurrentEpoch, result)
}

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

func (checker *Checker) getMetricForDashboardCell(cellName dashboards.CellName) any {
	node, rpc := checker.node[0], checker.rpc[0]

	switch cellName {
	case dashboards.CellNameNodeStatus:
		return node.Status.DashboardStatus()
	case dashboards.CellNameNetworkStatus:
		return rpc.Status.DashboardStatus()
	case dashboards.CellNameAddress:
		return node.AddressInfo.HostPort.Address
	case dashboards.CellNameTransactionsPerSecond:
		return node.Metrics.TransactionsPerSecond
	case dashboards.CellNameCheckpointsPerSecond:
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
		return strings.Split(node.Metrics.Uptime, " ")[0]
	case dashboards.CellNameVersion:
		return node.Metrics.Version
	case dashboards.CellNameCommit:
		return node.Metrics.Commit
	case dashboards.CellNameCompany:
		return node.Location.Provider
	case dashboards.CellNameCountry:
		return node.Location.String()
	case dashboards.CellNameEpoch:
		epochLabel := node.Metrics.GetEpochLabel()
		epochPercentage := node.Metrics.GetEpochProgress()

		return dashboards.NewDonutInput(epochLabel, epochPercentage)
	case dashboards.CellNameEpochEnd:
		return node.Metrics.GetEpochTimer()
	case dashboards.CellNameDiskUsage:
		usageLabel, usagePercentage := getDonutUsageMetric(utility.GetDiskUsage)

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
		usageLabel, usagePercentage := getDonutUsageMetric(utility.GetMemoryUsage)

		return dashboards.NewDonutInput(usageLabel, usagePercentage)
	case dashboards.CellNameCpuUsage:
		usageLabel, usagePercentage := getDonutUsageMetric(utility.GetCPUUsage)

		return dashboards.NewDonutInput(usageLabel, usagePercentage)
	default:
		return ""
	}
}

func getDonutUsageMetric(option func() (*utility.UsageData, error)) (string, int) {
	var (
		usageLabel      = "LOADING..."
		usagePercentage = 1
		usageData       *utility.UsageData
		err             error
	)

	if usageData, err = option(); err == nil {
		usageLabel = fmt.Sprintf("TOTAL/USED: %d/%d%%", usageData.Total, usageData.Used)
		usagePercentage = usageData.PercentageUsed

		if usagePercentage == 0 {
			usagePercentage = 1
		}
	}

	return usageLabel, usagePercentage
}

func getNetworkUsageMetric(networkMetric dashboards.CellName) string {
	var (
		usageData    = ""
		networkUsage *utility.NetworkUsage
		err          error
	)

	if networkUsage, err = utility.GetNetworkUsage(); err == nil {
		metric := networkUsage.Sent
		formatString := "%.02f%s"
		unit := "GB"

		if networkMetric == dashboards.CellNameBytesReceived {
			metric = networkUsage.Recv
		}

		if metric >= 100 {
			metric = metric / 100
			unit = "TB"

			if metric >= 100 {
				formatString = "%.01f%s"
			}
		}

		usageData = fmt.Sprintf(formatString, metric, unit)
	}

	return usageData
}

func getDirectorySize(dirPath string) string {
	var (
		usageData = ""
		dirSize   float64
		err       error
	)

	var processSize = func() {
		formatString := "%.02f%s"
		unit := "GB"

		if dirSize >= 100 {
			dirSize = dirSize / 100
			unit = "TB"

			if dirSize >= 100 {
				formatString = "%.01f%s"
			}
		}

		usageData = fmt.Sprintf(formatString, dirSize, unit)
	}

	if dirSize, err = utility.GetDirSize(dirPath); err == nil {
		processSize()

		return usageData
	}

	if dirSize, err = utility.GetVolumeSize("suidb"); err == nil {
		processSize()
	}

	return usageData
}

func (checker *Checker) getOptionsForDashboardCell(cellName dashboards.CellName) []cell.Option {
	var (
		options []cell.Option
		node    = checker.node[0]
		rpc     = checker.rpc[0]
	)

	var getColorOptions = func(status enums.Status) cell.Color {
		color := cell.ColorGreen

		switch status {
		case enums.StatusYellow:
			color = cell.ColorYellow
		case enums.StatusRed:
			color = cell.ColorRed
		}

		return color
	}

	switch cellName {
	case dashboards.CellNameNodeStatus:
		color := getColorOptions(node.Status)

		options = append(options, cell.BgColor(color), cell.FgColor(color))
	case dashboards.CellNameNetworkStatus:
		color := getColorOptions(rpc.Status)

		options = append(options, cell.BgColor(color), cell.FgColor(color))
	case dashboards.CellNameEpoch, dashboards.CellNameDiskUsage, dashboards.CellNameCpuUsage, dashboards.CellNameMemoryUsage:
		options = append(options, cell.Bold())
	default:
		options = append(options, cell.FgColor(cell.ColorWhite))
	}

	return options
}
