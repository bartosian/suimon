package checker

import (
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/dariubs/percent"
	"github.com/ipinfo/go/v2/ipinfo"
	"github.com/ybbus/jsonrpc/v3"

	"github.com/bartosian/sui_helpers/suimon/cmd/checker/dashboardbuilder"
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/dashboardbuilder/dashboards"
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
	"github.com/bartosian/sui_helpers/suimon/pkg/address"
	"github.com/bartosian/sui_helpers/suimon/pkg/log"
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
		stateMutex sync.RWMutex
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

func (host *Host) SetStatus(tableType enums.TableType, rpc Host) {
	status := enums.StatusGreen
	metricsHost := host.Metrics
	metricsRPC := rpc.Metrics

	switch metricsHost.Updated {
	case false:
		status = enums.StatusRed

		if tableType == enums.TableTypePeers {
			status = enums.StatusYellow
		}
	case true:
		if metricsHost.TotalTransactionNumber == 0 || metricsHost.LatestCheckpoint == 0 {
			status = enums.StatusRed

			if tableType == enums.TableTypePeers {
				status = enums.StatusYellow
			}

			break
		}

		if metricsHost.TransactionsPerSecond == 0 && rpc.Metrics.TransactionsPerSecond != 0 {
			status = enums.StatusRed

			break
		}

		if metricsHost.TransactionsPerSecond != 0 &&
			metricsHost.IsUnhealthy(enums.MetricTypeTransactionsPerSecond, metricsRPC.TransactionsPerSecond) {
			status = enums.StatusYellow
		}

		if metricsHost.TotalTransactionNumber != 0 &&
			metricsHost.IsUnhealthy(enums.MetricTypeTotalTransactionsNumber, metricsRPC.TotalTransactionNumber) {
			status = enums.StatusYellow

			break
		}

		if metricsHost.LatestCheckpoint != 0 &&
			metricsHost.IsUnhealthy(enums.MetricTypeLatestCheckpoint, metricsRPC.LatestCheckpoint) {
			status = enums.StatusYellow
		}
	}

	host.Status = status
}

func (host *Host) getMetricByDashboardCell(cellName dashboards.CellName) any {
	switch cellName {
	case dashboards.CellNameStatus:
		return host.Status.DashboardStatus()
	case dashboards.CellNameAddress:
		return host.AddressInfo.HostPort.Address
	case dashboards.CellNameTransactionsPerSecond:
		return host.Metrics.TransactionsPerSecond
	case dashboards.CellNameTotalTransactions:
		return host.Metrics.TotalTransactionNumber
	case dashboards.CellNameLatestCheckpoint:
		return host.Metrics.LatestCheckpoint
	case dashboards.CellNameHighestCheckpoint:
		return host.Metrics.HighestSyncedCheckpoint
	case dashboards.CellNameConnectedPeers:
		return host.Metrics.SuiNetworkPeers
	case dashboards.CellNameTXSyncProgress:
		return host.Metrics.TxSyncPercentage
	case dashboards.CellNameCheckSyncProgress:
		return host.Metrics.CheckSyncPercentage
	case dashboards.CellNameUptime:
		return strings.Split(host.Metrics.Uptime, " ")[0]
	case dashboards.CellNameVersion:
		return host.Metrics.Version
	case dashboards.CellNameCommit:
		return host.Metrics.Commit
	case dashboards.CellNameCompany:
		return host.Location.Provider
	case dashboards.CellNameCountry:
		return host.Location.String()
	case dashboards.CellNameEpoch:
		epochLabel := host.Metrics.GetEpochLabel()
		epochPercentage := host.Metrics.GetEpochProgress()

		return dashboardbuilder.NewDonutInput(epochLabel, epochPercentage)
	case dashboards.CellNameEpochEnd:
		return host.Metrics.GetEpochTimer()
	}

	return "no data"
}
