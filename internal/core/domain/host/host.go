package host

import (
	"fmt"

	"github.com/dariubs/percent"

	"github.com/bartosian/suimon/internal/core/domain/enums"
	"github.com/bartosian/suimon/internal/core/domain/metrics"
	"github.com/bartosian/suimon/internal/core/gateways/cligw"
	"github.com/bartosian/suimon/internal/core/ports"
)

const percentage100 = 100

type Gateways struct {
	rpc        ports.RPCGateway
	geo        ports.GeoGateway
	prometheus ports.PrometheusGateway
	cli        *cligw.Gateway
}

type Host struct {
	AddressInfo

	gateways Gateways
	IPInfo   *ports.IPResult

	TableType enums.TableType

	Status  enums.Status
	Metrics metrics.Metrics
}

func NewHost(
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
func (host *Host) SetPctProgress(metricType enums.MetricType, rpc *Host) error {
	hostMetric := host.Metrics.GetValue(metricType)
	rpcMetric := rpc.Metrics.GetValue(metricType)

	hostMetricInt, ok := hostMetric.(int)
	if !ok {
		return fmt.Errorf("failed to convert metric to INT")
	}

	rpcMetricInt, ok := rpcMetric.(int)
	if !ok {
		return fmt.Errorf("failed to convert metric to INT")
	}

	percentage := int(percent.PercentOf(hostMetricInt, rpcMetricInt))
	if percentage > percentage100 {
		percentage = percentage100
	}

	return host.Metrics.SetValue(metricType, percentage)
}

// SetStatus updates the status of the Host based on the provided RPC Host.
// It compares the metrics of the Host and RPC Host and sets the status to Red, Yellow, or Green based on specific conditions.
func (host *Host) SetStatus(rpc *Host) {
	metricsHost := host.Metrics
	metricsRPC := rpc.Metrics

	if host.TableType == enums.TableTypeValidator {
		if !metricsHost.Updated || metricsHost.Uptime == "" {
			host.Status = enums.StatusRed
			return
		}
	}

	if host.TableType == enums.TableTypeNode || host.TableType == enums.TableTypeRPC {
		if !metricsHost.Updated {
			host.Status = enums.StatusRed
			return
		}

		if metricsHost.TotalTransactionsBlocks == 0 ||
			metricsHost.LatestCheckpoint == 0 ||
			(metricsHost.TransactionsPerSecond == 0 && len(metricsHost.TransactionsHistory) == metrics.TransactionsPerSecondWindow) ||
			metricsHost.TxSyncPercentage == 0 ||
			metricsHost.TxSyncPercentage > 110 ||
			metricsHost.CheckSyncPercentage > 110 {
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
