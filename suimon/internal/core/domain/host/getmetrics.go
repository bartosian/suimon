package host

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"strings"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/ports"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	rpcClientTimeout   = 4 * time.Second
	rpcPortDefault     = "9000"
	metricsPortDefault = "9184"

	requestTypeRPC requestType = iota
	requestTypeMetrics
)

var (
	rpcMethodToMetricMap = map[enums.RPCMethod]enums.MetricType{
		enums.RPCMethodGetTotalTransactionBlocks:         enums.MetricTypeTotalTransactionBlocks,
		enums.RPCMethodGetLatestCheckpointSequenceNumber: enums.MetricTypeLatestCheckpoint,
		enums.RPCMethodGetSuiSystemState:                 enums.MetricTypeSuiSystemState,
	}

	prometheusToMetricMap = map[enums.PrometheusMetricName]enums.MetricType{
		enums.PrometheusMetricNameTotalTransactionCertificates: enums.MetricTypeTotalTransactionCertificates,
		enums.PrometheusMetricNameTotalTransactionEffects:      enums.MetricTypeTotalTransactionEffects,
		enums.PrometheusMetricNameHighestKnownCheckpoint:       enums.MetricTypeHighestKnownCheckpoint,
		enums.PrometheusMetricNameHighestSyncedCheckpoint:      enums.MetricTypeHighestSyncedCheckpoint,
		enums.PrometheusMetricNameLastExecutedCheckpoint:       enums.MetricTypeLastExecutedCheckpoint,
		enums.PrometheusMetricNameCurrentEpoch:                 enums.MetricTypeCurrentEpoch,
		enums.PrometheusMetricNameEpochTotalDuration:           enums.MetricTypeEpochTotalDuration,
		enums.PrometheusMetricNameCurrentRound:                 enums.MetricTypeCurrentRound,
		enums.PrometheusMetricNameHighestProcessedRound:        enums.MetricTypeHighestProcessedRound,
		enums.PrometheusMetricNameLastCommittedRound:           enums.MetricTypeLastCommittedRound,
		enums.PrometheusMetricNamePrimaryNetworkPeers:          enums.MetricTypePrimaryNetworkPeers,
		enums.PrometheusMetricNameWorkerNetworkPeers:           enums.MetricTypeWorkerNetworkPeers,
		enums.PrometheusMetricNameSuiNetworkPeers:              enums.MetricTypeSuiNetworkPeers,
		enums.PrometheusMetricNameSkippedConsensusTransactions: enums.MetricTypeSkippedConsensusTransactions,
		enums.PrometheusMetricNameTotalSignatureErrors:         enums.MetricTypeTotalSignatureErrors,
		enums.PrometheusMetricNameUptime:                       enums.MetricTypeUptime,
	}
)

// getPrometheusMetricsForHostType returns a list of Prometheus metrics that should be collected for a host.
func getPrometheusMetricsForHostType() ports.Metrics {
	return ports.Metrics{
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
}

// GetPrometheusMetrics fetches Prometheus rpcgw for the host and sends the metric values to the host's Metrics field.
// The function constructs a URL using the host's HostPort and Ports fields, sends an HTTP GET request to the URL, and parses the response as a Prometheus text format.
// Returns an error if the HTTP request fails or if the response cannot be parsed as Prometheus text format.
func (host *Host) GetPrometheusMetrics() error {
	metricsDef := getPrometheusMetricsForHostType()

	result, err := host.gateways.prometheus.CallFor(metricsDef)
	if err != nil {
		return err
	}

	if result == nil {
		return errors.New("failed to get metrics from Prometheus")
	}

	for metricName, metricValue := range result {
		metricType, ok := prometheusToMetricMap[metricName]
		if !ok {
			// ignore unused metric
			delete(result, metricName)
			continue
		}

		if err := host.Metrics.SetValue(metricType, metricValue.Value); err != nil {
			return err
		}

		// delete processed metric from result map
		delete(result, metricName)

		if metricType != enums.MetricTypeUptime {
			continue
		}

		if value, ok := metricValue.Labels["version"]; ok {
			versionInfo := strings.SplitN(value, "-", 2)

			if err := host.Metrics.SetValue(enums.MetricTypeVersion, versionInfo[0]); err != nil {
				return err
			}

			if len(versionInfo) == 2 {
				if err := host.Metrics.SetValue(enums.MetricTypeCommit, versionInfo[1]); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// GetMetrics fetches data from the host by calling three different methods asynchronously: GetTotalTransactionNumber, GetLatestCheckpoint, and GetPrometheusMetrics.
// The function waits for all three methods to complete before returning.
// Returns an error if any of the three methods fail or return an error.
func (host *Host) GetMetrics() error {
	var errGroup errgroup.Group

	rpcMethods := []enums.RPCMethod{enums.RPCMethodGetTotalTransactionBlocks, enums.RPCMethodGetLatestCheckpointSequenceNumber}

	switch host.TableType {
	case enums.TableTypeNode, enums.TableTypePeers:
		for _, method := range rpcMethods {
			method := method

			errGroup.Go(func() error {
				return host.GetDataByMetric(method)
			})
		}

		errGroup.Go(func() error {
			return host.GetPrometheusMetrics()
		})
	case enums.TableTypeValidator:
		errGroup.Go(func() error {
			return host.GetPrometheusMetrics()
		})
	case enums.TableTypeRPC:
		for _, method := range rpcMethods {
			method := method

			errGroup.Go(func() error {
				return host.GetDataByMetric(method)
			})
		}
	}

	if err := errGroup.Wait(); err != nil {
		return err
	}

	return nil
}

// GetDataByMetric is a method of the Host struct that retrieves data for a given RPC method
// and stores it as a metric in the Metrics struct. It takes an RPCMethod input parameter and
// returns an error if the method is not supported.
func (host *Host) GetDataByMetric(method enums.RPCMethod) error {
	metric, ok := rpcMethodToMetricMap[method]
	if !ok {
		return fmt.Errorf("unsupported RPC method: %v", method)
	}

	result, err := host.gateways.rpc.CallFor(method)
	if err != nil {
		return err
	}

	return host.Metrics.SetValue(metric, result)
}

// getUrl returns the URL for the specified request type and security status. The request parameter specifies the type of request to be made, and the secure parameter specifies whether to use HTTPS or HTTP for the URL.
// The function constructs the URL using the host's HostPort and Ports fields, and sets the appropriate scheme and port based on the request type and security status.
// Returns the constructed URL as a string.
func (host *Host) getUrl(request requestType, secure bool) string {
	hostUrl := new(url.URL)

	protocol := "http"
	if secure {
		protocol = protocol + "s"
	}

	hostAddress := host.HostPort
	if hostAddress.Host != nil {
		hostUrl.Host = *hostAddress.Host
	} else {
		hostUrl.Host = *hostAddress.IP
	}

	if hostAddress.Path != nil {
		hostUrl.Path = *hostAddress.Path
	}

	hostUrl.Scheme = protocol

	switch request {
	case requestTypeRPC:
		port := host.Ports[enums.PortTypeRPC]
		if port == "" {
			port = rpcPortDefault
		}
		hostUrl.Host = net.JoinHostPort(hostUrl.Hostname(), port)
	case requestTypeMetrics:
		hostUrl.Path = "/rpcgw"

		port := host.Ports[enums.PortTypeMetrics]
		if port == "" {
			port = metricsPortDefault
		}
		hostUrl.Host = net.JoinHostPort(hostUrl.Hostname(), port)
	}

	return hostUrl.String()
}
