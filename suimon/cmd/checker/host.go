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

// newHost creates and returns a new "Host" object, based on the provided "AddressInfo" value.
// Parameters:
// - addressInfo: an "AddressInfo" value representing the address information for the new host.
// - ipClient: a pointer to an "ipinfo.Client" instance for retrieving additional information about the host's IP.
// - httpClient: a pointer to an "http.Client" instance for performing HTTP requests to the host.
// Returns: a pointer to the newly created "Host" object.
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

// SetPctProgress sets the percentage progress for the specified "MetricType" on the "Host" object passed
// as a pointer receiver. It calculates the progress percentage based on the number of successful checks
// and the total number of checks performed for that metric, and updates the corresponding progress value
// on the "Host" object.
// Parameters:
// - metricType: an enums.MetricType representing the metric type to set the percentage progress for.
// - rpc: a "Host" object representing the RPC host to update with the new progress percentage.
// Returns: None.
func (host *Host) SetPctProgress(metricType enums.MetricType, rpc Host) {
	hostMetric := host.Metrics.GetValue(metricType, false)
	rpcMetric := rpc.Metrics.GetValue(metricType, true)
	hostMetricInt, rpcMetricInt := hostMetric.(int), rpcMetric.(int)

	percentage := int(percent.PercentOf(hostMetricInt, rpcMetricInt))

	host.Metrics.SetValue(metricType, percentage)
}

// SetStatus updates the status of the "Host" object passed as a pointer receiver, based on the results of
// the most recent network checks performed on the host. It sets the "Status" field of the "Host" object
// to reflect whether the host is up or down, and updates the corresponding status value on the "rpc" host.
// Parameters: rpc: a "Host" object representing the RPC host to update with the new status.
// Returns: None.
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
