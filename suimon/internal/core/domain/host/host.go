package host

import (
	"github.com/dariubs/percent"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/metrics"
	"github.com/bartosian/sui_helpers/suimon/internal/core/gateways/cligw"
	"github.com/bartosian/sui_helpers/suimon/internal/core/ports"
	"github.com/bartosian/sui_helpers/suimon/internal/pkg/address"
	"github.com/bartosian/sui_helpers/suimon/internal/pkg/log"
)

type requestType int

type (
	AddressInfo struct {
		HostPort address.HostPort
		Ports    map[enums.PortType]string
	}

	Gateways struct {
		rpc        ports.RPCGateway
		geo        ports.GeoGateway
		prometheus ports.PrometheusGateway
		cli        *cligw.Gateway
	}

	Host struct {
		AddressInfo

		TableType enums.TableType

		Status  enums.Status
		IPInfo  *ports.IPResult
		Metrics metrics.Metrics

		gateways Gateways

		logger log.Logger
	}
)

func NewHost(
	logger log.Logger,
	tableType enums.TableType,
	addressInfo AddressInfo,
	rpcGW ports.RPCGateway,
	geoGW ports.GeoGateway,
	prometheusGW ports.PrometheusGateway,
	cliGW *cligw.Gateway,
) *Host {
	host := &Host{
		TableType:   tableType,
		AddressInfo: addressInfo,
		logger:      logger,
		gateways: Gateways{
			rpc:        rpcGW,
			geo:        geoGW,
			prometheus: prometheusGW,
			cli:        cliGW,
		},
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
