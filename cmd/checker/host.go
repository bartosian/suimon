package checker

import (
	"net/http"
	"regexp"
	"time"

	"github.com/dariubs/percent"
	"github.com/ipinfo/go/v2/ipinfo"
	"github.com/ybbus/jsonrpc/v3"

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
	metricsHost := host.Metrics
	metricsRPC := rpc.Metrics

	if !metricsHost.Updated || metricsHost.TotalTransactionNumber == 0 || metricsHost.LatestCheckpoint == 0 ||
		metricsHost.TransactionsPerSecond == 0 && len(metricsHost.TransactionsHistory) == transactionsPerSecondTimeout ||
		metricsHost.TxSyncPercentage == 0 {

		host.Status = enums.StatusRed

		return
	}

	if metricsHost.IsUnhealthy(enums.MetricTypeTransactionsPerSecond, metricsRPC.TransactionsPerSecond) ||
		metricsHost.IsUnhealthy(enums.MetricTypeTotalTransactionsNumber, metricsRPC.TotalTransactionNumber) ||
		metricsHost.IsUnhealthy(enums.MetricTypeLatestCheckpoint, metricsRPC.LatestCheckpoint) {
		host.Status = enums.StatusYellow

		return
	}

	host.Status = enums.StatusGreen

	return
}
