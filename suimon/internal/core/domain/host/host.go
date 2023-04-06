package host

import (
	"net/http"
	"time"

	"github.com/dariubs/percent"
	"github.com/ipinfo/go/v2/ipinfo"
	"github.com/ybbus/jsonrpc/v3"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/location"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/metrics"
	"github.com/bartosian/sui_helpers/suimon/internal/pkg/address"
	"github.com/bartosian/sui_helpers/suimon/internal/pkg/log"
)

type requestType int

const (
	rpcPortDefault     = "9000"
	metricsPortDefault = "9184"
	rpcClientTimeout   = 4 * time.Second

	requestTypeRPC requestType = iota
	requestTypeMetrics
)

type (
	AddressInfo struct {
		HostPort address.HostPort
		Ports    map[enums.PortType]string
	}

	Clients struct {
		rpcClient  jsonrpc.RPCClient
		httpClient *http.Client
		ipClient   *ipinfo.Client
	}

	Host struct {
		AddressInfo

		TableType enums.TableType

		Status   enums.Status
		Location *location.Location
		Metrics  metrics.Metrics

		clients Clients

		logger log.Logger
	}
)

// NewHost creates a new Host instance with the given table type and address information.
// The function initializes a new Metrics instance and logger, and creates a new JSON-RPC client for both secure and non-secure connections using the Host's URL obtained from the address information.
// The function also initializes an HTTP client and IP info client for the Host.
// Returns a pointer to the Host instance.
func NewHost(tableType enums.TableType, addressInfo AddressInfo, ipClient *ipinfo.Client, httpClient *http.Client) *Host {
	host := &Host{
		TableType:   tableType,
		AddressInfo: addressInfo,
		logger:      log.NewLogger(),
	}

	secureURL := addressInfo.HostPort.SSL
	rpcClient := jsonrpc.NewClient(host.getUrl(requestTypeRPC, secureURL))

	host.clients = Clients{
		rpcClient:  rpcClient,
		httpClient: httpClient,
		ipClient:   ipClient,
	}

	return host
}

// SetPctProgress updates the value of the specified metric type for the Host instance with a percentage that reflects the Host's progress relative to the progress of the RPC Host.
// The function obtains the current metric value for the Host and RPC Host, calculates the percentage using the percent.PercentOf function, and sets the new percentage value for the Host's Metrics instance for the specified metric type.
// The second argument is the RPC Host to compare the progress against.
func (host *Host) SetPctProgress(metricType enums.MetricType, rpc Host) error {
	hostMetric := host.Metrics.GetValue(metricType, false)
	rpcMetric := rpc.Metrics.GetValue(metricType, true)
	hostMetricInt, rpcMetricInt := hostMetric.(int), rpcMetric.(int)

	percentage := int(percent.PercentOf(hostMetricInt, rpcMetricInt))

	return host.Metrics.SetValue(metricType, percentage)
}

func (host *Host) SetStatus(rpc Host) {
	metricsHost := host.Metrics
	metricsRPC := rpc.Metrics

	switch host.TableType {
	case enums.TableTypeValidator:
		if !metricsHost.Updated || metricsHost.Uptime == "" {
			host.Status = enums.StatusRed

			return
		}
	case enums.TableTypeNode, enums.TableTypeRPC, enums.TableTypePeers:
		if !metricsHost.Updated || metricsHost.TotalTransactionsBlocks == 0 || metricsHost.LatestCheckpoint == 0 ||
			metricsHost.TransactionsPerSecond == 0 && len(metricsHost.TransactionsHistory) == metrics.TransactionsPerSecondWindow ||
			metricsHost.TxSyncPercentage == 0 || metricsHost.TxSyncPercentage > 110 || metricsHost.CheckSyncPercentage > 110 {

			host.Status = enums.StatusRed

			return
		}

		if metricsHost.IsUnhealthy(enums.MetricTypeTransactionsPerSecond, metricsRPC.TransactionsPerSecond) ||
			metricsHost.IsUnhealthy(enums.MetricTypeTotalTransactionBlocks, metricsRPC.TotalTransactionsBlocks) ||
			metricsHost.IsUnhealthy(enums.MetricTypeLatestCheckpoint, metricsRPC.LatestCheckpoint) {

			host.Status = enums.StatusYellow

			return
		}
	}

	host.Status = enums.StatusGreen
}
