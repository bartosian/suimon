package host

import (
	"context"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/metrics"
	"net/url"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/ybbus/jsonrpc/v3"

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

// GetPrometheusMetrics returns a "Metrics" object representing the current state of network checks for the "Host" object
// passed as a pointer receiver. This object contains the results of each metric check performed on the host,
// including the number of successful checks, the total number of checks performed, and the percentage progress
// for each metric.
// Parameters: None.
// Returns: - a "Metrics" object representing the current state of network checks for the "Host" object.
func (host *Host) GetPrometheusMetrics() {
	metricsURL := host.getUrl(requestTypeMetrics, false)
	parser := metricsparser.NewPrometheusMetricParser(host.httpClient, metricsURL, prometheusMetrics)

	result, err := parser.GetMetrics()
	if err != nil {
		return
	}

	for metricName, metricValue := range result {
		switch metricName {
		case enums.PrometheusMetricNameTotalTransactionCertificates:
			host.Metrics.SetValue(enums.MetricTypeTotalTransactionCertificates, metricValue.Value)
		case enums.PrometheusMetricNameTotalTransactionEffects:
			host.Metrics.SetValue(enums.MetricTypeTotalTransactionEffects, metricValue.Value)
		case enums.PrometheusMetricNameHighestKnownCheckpoint:
			host.Metrics.SetValue(enums.MetricTypeHighestKnownCheckpoint, metricValue.Value)
		case enums.PrometheusMetricNameHighestSyncedCheckpoint:
			host.Metrics.SetValue(enums.MetricTypeHighestSyncedCheckpoint, metricValue.Value)
		case enums.PrometheusMetricNameLastExecutedCheckpoint:
			host.Metrics.SetValue(enums.MetricTypeLastExecutedCheckpoint, metricValue.Value)
		case enums.PrometheusMetricNameCurrentEpoch:
			host.Metrics.SetValue(enums.MetricTypeCurrentEpoch, metricValue.Value)
		case enums.PrometheusMetricNameEpochTotalDuration:
			host.Metrics.SetValue(enums.MetricTypeEpochTotalDuration, metricValue.Value)
		case enums.PrometheusMetricNameCurrentRound:
			host.Metrics.SetValue(enums.MetricTypeCurrentRound, metricValue.Value)
		case enums.PrometheusMetricNameHighestProcessedRound:
			host.Metrics.SetValue(enums.MetricTypeHighestProcessedRound, metricValue.Value)
		case enums.PrometheusMetricNameLastCommittedRound:
			host.Metrics.SetValue(enums.MetricTypeLastCommittedRound, metricValue.Value)
		case enums.PrometheusMetricNamePrimaryNetworkPeers:
			host.Metrics.SetValue(enums.MetricTypePrimaryNetworkPeers, metricValue.Value)
		case enums.PrometheusMetricNameWorkerNetworkPeers:
			host.Metrics.SetValue(enums.MetricTypeWorkerNetworkPeers, metricValue.Value)
		case enums.PrometheusMetricNameSuiNetworkPeers:
			host.Metrics.SetValue(enums.MetricTypeSuiNetworkPeers, metricValue.Value)
		case enums.PrometheusMetricNameSkippedConsensusTransactions:
			host.Metrics.SetValue(enums.MetricTypeSkippedConsensusTransactions, metricValue.Value)
		case enums.PrometheusMetricNameTotalSignatureErrors:
			host.Metrics.SetValue(enums.MetricTypeTotalSignatureErrors, metricValue.Value)
		case enums.PrometheusMetricNameUptime:
			host.Metrics.SetValue(enums.MetricTypeUptime, metricValue.Value)

			if value, ok := metricValue.Labels["version"]; ok {
				versionInfo := strings.Split(value, "-")

				host.Metrics.SetValue(enums.MetricTypeVersion, versionInfo[0])

				if len(versionInfo) == 2 {
					host.Metrics.SetValue(enums.MetricTypeCommit, versionInfo[1])
				}
			}
		}
	}
}

// GetTotalTransactionNumber returns the total number of transactions performed on the "Host" object passed
// as a pointer receiver. This method retrieves the "Metrics" object for the host and calculates the total
// number of transactions performed across all metric types.
// Parameters: None.
// Returns: an integer representing the total number of transactions performed on the "Host" object.
func (host *Host) GetTotalTransactionNumber() {
	var result any

	if result = getFromRPC(host.rpcHttpClient, enums.RPCMethodGetTotalTransactionNumber); result == nil {
		if result = getFromRPC(host.rpcHttpsClient, enums.RPCMethodGetTotalTransactionNumber); result == nil {
			return
		}
	}

	host.Metrics.SetValue(enums.MetricTypeTotalTransactions, result)
}

// GetLatestCheckpoint returns a "Checkpoint" object representing the most recent checkpoint for the "Host"
// object passed as a pointer receiver. This object contains information about the time and status of the
// most recent checkpoint performed on the host.
// Parameters: None.
// Returns: a "Checkpoint" object representing the most recent checkpoint for the "Host" object.
func (host *Host) GetLatestCheckpoint() {
	var result any

	if result = getFromRPC(host.rpcHttpClient, enums.RPCMethodGetLatestCheckpointSequenceNumber); result == nil {
		if result = getFromRPC(host.rpcHttpsClient, enums.RPCMethodGetLatestCheckpointSequenceNumber); result == nil {
			return
		}
	}

	host.Metrics.SetValue(enums.MetricTypeLatestCheckpoint, result)
}

// GetLatestSuiSystemState returns a "SUISystemState" object representing the current system state for the "Host"
// object passed as a pointer receiver. This object contains information about the status of various components
// of the SUISystem software running on the host.
// Parameters: None.
// Returns: a "SUISystemState" object representing the current system state for the "Host" object.
func (host *Host) GetLatestSuiSystemState() {
	var result any

	if result = getFromRPC(host.rpcHttpClient, enums.RPCMethodGetSuiSystemState); result == nil {
		if result = getFromRPC(host.rpcHttpsClient, enums.RPCMethodGetSuiSystemState); result == nil {
			return
		}
	}

	host.Metrics.SetValue(enums.MetricTypeSuiSystemState, result)
}

// getFromRPC makes a JSON-RPC call to the specified method using the provided RPC client, and returns
// the result of the call. The returned type will depend on the specific method called and its response.
// Parameters:
// - rpcClient: a jsonrpc.RPCClient representing the client to use for the JSON-RPC call.
// - method: an enums.RPCMethod representing the name of the JSON-RPC method to call.
// Returns:
//   - the result of the JSON-RPC call. The specific type of the returned value will depend on the method called
//     and its response.
func getFromRPC(rpcClient jsonrpc.RPCClient, method enums.RPCMethod) any {
	var (
		respChan = make(chan any)
		timeout  = time.After(rpcClientTimeout)
	)

	switch method {
	case enums.RPCMethodGetSuiSystemState:
		var response metrics.SuiSystemState

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

// GetData returns a "HostData" object representing the current data for the "Host" object passed as a pointer receiver.
// This object contains various metrics and status information about the host.
// Parameters: None.
// Returns: a "HostData" object representing the current data for the "Host" object.
func (host *Host) GetData() {
	doneCH := make(chan struct{})

	defer close(doneCH)

	go func() {
		host.GetTotalTransactionNumber()

		doneCH <- struct{}{}
	}()

	go func() {
		host.GetLatestCheckpoint()

		doneCH <- struct{}{}
	}()

	go func() {
		host.GetPrometheusMetrics()

		doneCH <- struct{}{}
	}()

	for i := 0; i < 3; i++ {
		<-doneCH
	}
}

// getUrl returns the URL for a given request type and security setting.
// Parameters:
// - request: a requestType indicating the type of request for which to generate a URL.
// - secure: a boolean indicating whether the generated URL should use HTTPS (true) or HTTP (false).
// Returns: a string representing the URL for the specified request type and security setting.
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
		} else {
			hostUrl.Host = hostUrl.Hostname() + ":" + rpcPortDefault
		}
	case requestTypeMetrics:
		hostUrl.Path = "/metrics"

		if port, ok := host.Ports[enums.PortTypeMetrics]; ok {
			hostUrl.Host = hostUrl.Hostname() + ":" + port
		} else {
			hostUrl.Host = hostUrl.Hostname() + ":" + metricsPortDefault
		}
	}

	return hostUrl.String()
}
