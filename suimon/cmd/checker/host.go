package checker

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/dariubs/percent"
	"github.com/ipinfo/go/v2/ipinfo"
	"github.com/ybbus/jsonrpc/v3"

	"github.com/bartosian/sui_helpers/suimon/cmd/checker/dashboardbuilder"
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/dashboardbuilder/dashboards"
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
	"github.com/bartosian/sui_helpers/suimon/pkg/address"
	"github.com/bartosian/sui_helpers/suimon/pkg/log"
	"github.com/bartosian/sui_helpers/suimon/pkg/utility"
)

type requestType int

const (
	rpcPortDefault      = "9000"
	metricsPortDefault  = "9184"
	rpcClientTimeout    = 3 * time.Second
	metricVersionRegexp = `\{(.*?)\}`

	requestTypeRPC requestType = iota
	requestTypeMetrics
)

var versionRegex = regexp.MustCompile(metricVersionRegexp)

type (
	AddressInfo struct {
		HostPort address.HostPort
		Ports    map[enums.PortType]string
	}
	Host struct {
		AddressInfo

		Status   enums.Status
		Location *Location
		Metrics  Metrics

		rpcHttpClient  jsonrpc.RPCClient
		rpcHttpsClient jsonrpc.RPCClient
		httpClient     *http.Client
		ipClient       *ipinfo.Client

		logger log.Logger
	}
)

func newHost(addressInfo AddressInfo, ipClient *ipinfo.Client, httpClient *http.Client) *Host {
	host := &Host{
		AddressInfo: addressInfo,
		ipClient:    ipClient,
		httpClient:  httpClient,
		Metrics:     NewMetrics(),
		logger:      log.NewLogger(),
	}

	host.rpcHttpClient = jsonrpc.NewClient(host.getUrl(requestTypeRPC, false))
	host.rpcHttpsClient = jsonrpc.NewClient(host.getUrl(requestTypeRPC, true))

	return host
}

func (host *Host) SetPctProgress(metricType enums.MetricType, rpc Host) {
	hostMetric := host.Metrics.GetValue(metricType, false)
	rpcMetric := rpc.Metrics.GetValue(metricType, true)
	hostMetricInt, rpcMetricInt := hostMetric.(int), rpcMetric.(int)

	percentage := int(percent.PercentOf(hostMetricInt, rpcMetricInt))

	host.Metrics.SetValue(metricType, percentage)
}

func (host *Host) SetStatus(rpc Host) {
	status := enums.StatusGreen
	metricsHost := host.Metrics
	metricsRPC := rpc.Metrics

	switch metricsHost.Updated {
	case false:
		status = enums.StatusRed
	case true:
		if metricsHost.TotalTransactionNumber == 0 && metricsRPC.TotalTransactionNumber != 0 ||
			metricsHost.LatestCheckpoint == 0 && metricsRPC.LatestCheckpoint != 0 ||
			metricsHost.TransactionsPerSecond == 0 && metricsRPC.TransactionsPerSecond != 0 ||
			metricsHost.TxSyncPercentage > 100 || metricsHost.CheckSyncPercentage > 100 {
			status = enums.StatusRed

			break
		}

		if metricsHost.IsUnhealthy(enums.MetricTypeTransactionsPerSecond, metricsRPC.TransactionsPerSecond) ||
			metricsHost.IsUnhealthy(enums.MetricTypeTotalTransactionsNumber, metricsRPC.TotalTransactionNumber) ||
			metricsHost.IsUnhealthy(enums.MetricTypeLatestCheckpoint, metricsRPC.LatestCheckpoint) {
			status = enums.StatusYellow
		}
	}

	host.Status = status
}

func (checker *Checker) getMetricByDashboardCell(cellName dashboards.CellName) any {
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

		return dashboardbuilder.NewDonutInput(epochLabel, epochPercentage)
	case dashboards.CellNameEpochEnd:
		return node.Metrics.GetEpochTimer()
	case dashboards.CellNameDiskUsage:
		var (
			diskUsage *utility.DiskUsage
			err       error
		)

		if diskUsage, err = utility.GetDiskUsage(); err != nil {
			panic(err)
		}

		usageLabel := fmt.Sprintf("TOTAL/USED: %d/%dGB", diskUsage.Total, diskUsage.Used)
		usagePercentage := diskUsage.PercentageUsed

		if usagePercentage == 0 {
			usagePercentage = 1
		}

		return dashboardbuilder.NewDonutInput(usageLabel, usagePercentage)
	case dashboards.CellNameDatabaseSize:
		var (
			dbPath = checker.nodeConfig.DbPath
			dbSize float64
			err    error
		)

		if dbSize, err = utility.GetDirSize(dbPath); err != nil {
			if dbSize, err = utility.GetVolumeSize("suidb"); err != nil {
				return 0
			}
		}

		if dbSize >= 100 {
			dbSize = dbSize / 100

			return fmt.Sprintf("%.02fTB", dbSize)
		}

		return fmt.Sprintf("%.01fGB", dbSize)
	case dashboards.CellNameBytesSent:
		var (
			networkUsage *utility.NetworkUsage
			err          error
		)

		if networkUsage, err = utility.GetNetworkUsage(); err != nil {
			panic(err)
		}

		sent := networkUsage.Sent

		if sent >= 100 {
			sent = sent / 100

			return fmt.Sprintf("%.02fTB", sent)
		}

		return fmt.Sprintf("%.01fGB", sent)
	case dashboards.CellNameBytesReceived:
		var (
			networkUsage *utility.NetworkUsage
			err          error
		)

		if networkUsage, err = utility.GetNetworkUsage(); err != nil {
			panic(err)
		}

		recv := networkUsage.Recv

		if recv >= 100 {
			recv = recv / 100

			return fmt.Sprintf("%.02fTB", recv)
		}

		return fmt.Sprintf("%.01fGB", recv)
	case dashboards.CellNameMemoryUsage:
		var (
			memoryUsage *utility.MemoryUsage
			err         error
		)

		if memoryUsage, err = utility.GetMemoryUsage(); err != nil {
			panic(err)
		}

		usageLabel := fmt.Sprintf("TOTAL/USED: %d/%dGB", memoryUsage.Total, memoryUsage.Used)
		usagePercentage := memoryUsage.PercentageUsed

		if usagePercentage == 0 {
			usagePercentage = 1
		}

		return dashboardbuilder.NewDonutInput(usageLabel, usagePercentage)
	case dashboards.CellNameCpuUsage:
		var (
			cpuUsage *utility.CPUUsage
			err      error
		)

		if cpuUsage, err = utility.GetCPUUsage(); err != nil {
			return dashboardbuilder.NewDonutInput("", 0)
		}

		usageLabel := fmt.Sprintf("TOTAL/USED: %d/%d%%", cpuUsage.Total, cpuUsage.Used)
		usagePercentage := cpuUsage.PercentageUsed

		if usagePercentage == 0 {
			usagePercentage = 1
		}

		return dashboardbuilder.NewDonutInput(usageLabel, usagePercentage)
	default:
		return ""
	}
}
