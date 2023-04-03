package host

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/ybbus/jsonrpc/v3"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/pkg/metricsparser"
)

var prometheusMetrics = map[enums.PrometheusMetricName]metricsparser.MetricConfig{
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

// GetPrometheusMetrics fetches Prometheus metrics for the host and sends the metric values to the host's Metrics field.
// The function constructs a URL using the host's HostPort and Ports fields, sends an HTTP GET request to the URL, and parses the response as a Prometheus text format.
// Returns an error if the HTTP request fails or if the response cannot be parsed as Prometheus text format.
func (host *Host) GetPrometheusMetrics() error {
	metricsURL := host.getUrl(requestTypeMetrics, false)
	parser := metricsparser.NewPrometheusMetricParser(host.clients.httpClient, metricsURL, prometheusMetrics)

	result, err := parser.GetMetrics()
	if err != nil {
		return err
	}

	metricMap := map[enums.PrometheusMetricName]enums.MetricType{
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

	for metricName, metricValue := range result {
		metricType, ok := metricMap[metricName]
		if !ok {
			delete(result, metricName)
			continue
		}

		if err := host.Metrics.SetValue(metricType, metricValue.Value); err != nil {
			continue
		}

		if metricType != enums.MetricTypeUptime {
			continue
		}

		if value, ok := metricValue.Labels["version"]; ok {
			versionInfo := strings.SplitN(value, "-", 2)

			if err := host.Metrics.SetValue(enums.MetricTypeVersion, versionInfo[0]); err != nil {
				continue
			}

			if len(versionInfo) == 2 {
				if err := host.Metrics.SetValue(enums.MetricTypeCommit, versionInfo[1]); err != nil {
					continue
				}
			}
		}
	}

	return nil
}

// GetMetricRPC fetches a specific metric for the host from an RPC server.
// The method parameter specifies the RPC method to call, and the metricType parameter specifies the type of metric to retrieve.
// The function sends the metric value to the host's Metrics field if successful.
// Returns an error if the RPC call fails or if the metric value cannot be retrieved.
func (host *Host) GetMetricRPC(method enums.RPCMethod, metricType enums.MetricType) error {
	var (
		result any
		err    error
	)

	if result, err = getFromRPC(host.clients.rpcClient, method); err != nil {
		return err
	}

	return host.Metrics.SetValue(metricType, result)
}

// getFromRPC sends an RPC call to an RPC client and returns the response. The type parameter T specifies the expected type of the response.
// The function uses the method parameter to determine which RPC method to call.
// Returns the response value of type T and an error value if the RPC call fails or if the response cannot be converted to type T.
func getFromRPC(rpcClient jsonrpc.RPCClient, method enums.RPCMethod) (any, error) {
	respChan := make(chan any)
	timeout := time.After(rpcClientTimeout)

	go func() {
		var response any

		if err := rpcClient.CallFor(context.Background(), &response, method.String()); err != nil {
			respChan <- nil

			return
		}

		respChan <- response
	}()

	select {
	case response := <-respChan:
		if response == nil {
			return nil, fmt.Errorf("failed to get response from RPC client")
		}

		return response, nil
	case <-timeout:
		return nil, fmt.Errorf("timeout while waiting for RPC response")
	}
}

// GetData fetches data from the host by calling three different methods asynchronously: GetTotalTransactionNumber, GetLatestCheckpoint, and GetPrometheusMetrics.
// The function waits for all three methods to complete before returning.
// Returns an error if any of the three methods fail or return an error.
func (host *Host) GetData() error {
	methodsMapRPC := map[enums.RPCMethod]enums.MetricType{
		enums.RPCMethodGetTotalTransactionBlocks:         enums.MetricTypeTotalTransactionBlocks,
		enums.RPCMethodGetLatestCheckpointSequenceNumber: enums.MetricTypeLatestCheckpoint,
	}

	doneCH := make(chan error, len(methodsMapRPC)+1)

	var wg sync.WaitGroup

	switch host.TableType {
	case enums.TableTypeNode, enums.TableTypePeers:
		for method, metric := range methodsMapRPC {
			wg.Add(1)

			go func(method enums.RPCMethod, metric enums.MetricType) {
				defer wg.Done()

				err := host.GetMetricRPC(method, metric)
				doneCH <- err
			}(method, metric)
		}

		wg.Add(1)

		go func() {
			defer wg.Done()

			err := host.GetPrometheusMetrics()
			doneCH <- err
		}()
	case enums.TableTypeValidator:
		wg.Add(1)

		go func() {
			defer wg.Done()

			err := host.GetPrometheusMetrics()
			doneCH <- err
		}()
	case enums.TableTypeRPC:
		for method, metric := range methodsMapRPC {
			wg.Add(1)

			go func(method enums.RPCMethod, metric enums.MetricType) {
				defer wg.Done()

				err := host.GetMetricRPC(method, metric)
				doneCH <- err
			}(method, metric)
		}
	}

	go func() {
		wg.Wait()
		close(doneCH)
	}()

	var errors []error

	for err := range doneCH {
		if err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("errors occurred while fetching data: %v", errors)
	}

	return nil
}

// getUrl returns the URL for the specified request type and security status. The request parameter specifies the type of request to be made, and the secure parameter specifies whether to use HTTPS or HTTP for the URL.
// The function constructs the URL using the host's HostPort and Ports fields, and sets the appropriate scheme and port based on the request type and security status.
// Returns the constructed URL as a string.
func (host *Host) getUrl(request requestType, secure bool) string {
	var (
		protocol    = "http"
		hostAddress = host.HostPort
		hostUrl     = new(url.URL)
	)

	if hostAddress.Host != nil {
		hostUrl.Host = *hostAddress.Host
	} else {
		hostUrl.Host = *hostAddress.IP
	}

	if hostAddress.Path != nil {
		hostUrl.Path = *hostAddress.Path
	}

	if secure {
		protocol = protocol + "s"
	}

	hostUrl.Scheme = protocol

	switch request {
	case requestTypeRPC:
		if port, ok := host.Ports[enums.PortTypeRPC]; ok {
			hostUrl.Host = net.JoinHostPort(hostUrl.Hostname(), port)
		} else {
			hostUrl.Host = net.JoinHostPort(hostUrl.Hostname(), rpcPortDefault)
		}
	case requestTypeMetrics:
		hostUrl.Path = "/metrics"

		if port, ok := host.Ports[enums.PortTypeMetrics]; ok {
			hostUrl.Host = net.JoinHostPort(hostUrl.Hostname(), port)
		} else {
			hostUrl.Host = net.JoinHostPort(hostUrl.Hostname(), metricsPortDefault)
		}
	}

	return hostUrl.String()
}
