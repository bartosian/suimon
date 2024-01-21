package prometheusgw

import (
	"context"
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	ioPrometheusClient "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"

	"github.com/bartosian/suimon/internal/core/domain/enums"
	"github.com/bartosian/suimon/internal/core/ports"
)

type MetricsData map[string]*ioPrometheusClient.MetricFamily

type responseWithError struct {
	response *http.Response
	err      error
}

var extractMetricValue = map[enums.PrometheusMetricType]func(metric *ioPrometheusClient.Metric) float64{
	enums.PrometheusMetricTypeGauge:     func(metric *ioPrometheusClient.Metric) float64 { return metric.GetGauge().GetValue() },
	enums.PrometheusMetricTypeCounter:   func(metric *ioPrometheusClient.Metric) float64 { return metric.GetCounter().GetValue() },
	enums.PrometheusMetricTypeSummary:   func(metric *ioPrometheusClient.Metric) float64 { return metric.GetSummary().GetSampleSum() },
	enums.PrometheusMetricTypeHistogram: func(metric *ioPrometheusClient.Metric) float64 { return metric.GetHistogram().GetSampleSum() },
	enums.PrometheusMetricTypeUntyped:   func(metric *ioPrometheusClient.Metric) float64 { return metric.GetUntyped().GetValue() },
}

// CallFor makes an HTTP request to the specified gateway URL to fetch metrics.
// It returns the metrics result or an error if something goes wrong.
func (gateway *Gateway) CallFor(metrics ports.Metrics) (result ports.MetricsResult, err error) {
	if len(metrics) == 0 {
		return nil, fmt.Errorf("no metrics provided")
	}

	req, err := http.NewRequest("GET", gateway.url, http.NoBody)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(gateway.ctx, httpClientTimeout)
	defer cancel()

	req = req.WithContext(ctx)

	respChan := make(chan responseWithError, 1)

	go func() {
		//nolint:bodyclose // The response body is closed below to handle the response properly.
		resp, err := gateway.client.Do(req)

		respChan <- responseWithError{response: resp, err: err}
	}()

	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("http call timed out: %w", ctx.Err())
	case result := <-respChan:
		if result.err != nil {
			return nil, fmt.Errorf("failed to get response from http client: %w", result.err)
		}

		response := result.response
		if response != nil {
			defer func() {
				if closeErr := response.Body.Close(); closeErr != nil {
					err = fmt.Errorf("failed to close response body: %w", closeErr)
				}
			}()
		}

		if response.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
		}

		parser := expfmt.TextParser{}

		data, err := parser.TextToMetricFamilies(response.Body)
		if err != nil {
			return nil, err
		}

		metricsResult := make(ports.MetricsResult)

		for metricName, metricConfig := range metrics {
			result, err := getMetricValueWithLabelFiltering(data, metricName.ToString(), metricConfig)
			if err != nil {
				return nil, err
			}

			metricsResult[metricName] = result
		}

		return metricsResult, nil
	}
}

// getMetricValueWithLabelFiltering searches for a specific metric in the provided MetricsData
// by the given metricName and metricConfig. If a matching metric is found, it returns
// the metric result containing its value and labels. If no matching metric is found, it returns an error.
func getMetricValueWithLabelFiltering(metrics MetricsData, metricName string, metricConfig ports.MetricConfig) (result ports.MetricResult, err error) {
	if metrics == nil {
		return result, fmt.Errorf("no metrics provided")
	}

	metricType, labels := metricConfig.MetricType, metricConfig.Labels

	if _, validType := extractMetricValue[metricType]; !validType {
		return result, fmt.Errorf("invalid metric type: %v", metricType)
	}

	metricFamily, found := metrics[metricName]
	if !found {
		return result, fmt.Errorf("no metric family found for metric: %s with type: %v", metricName, metricType)
	}

	for _, metric := range metricFamily.Metric {
		if matchLabels(labels, metric.Label) {
			result.Labels = extractLabels(metric.Label)
			result.Value = extractMetricValue[metricType](metric)

			return result, nil
		}
	}

	return result, fmt.Errorf("no metric found matching labels: %v", labels)
}

// matchLabels checks if all labels match.
func matchLabels(expected prometheus.Labels, actual []*ioPrometheusClient.LabelPair) bool {
	for key, value := range expected {
		if !labelMatches(key, value, actual) {
			return false
		}
	}

	return true
}

// extractLabels extracts labels from the metric.
func extractLabels(labels []*ioPrometheusClient.LabelPair) prometheus.Labels {
	labelsResult := make(prometheus.Labels, len(labels))
	for _, label := range labels {
		labelsResult[label.GetName()] = label.GetValue()
	}

	return labelsResult
}

// labelMatches checks if the key-value pair exists in the given labels
func labelMatches(key, value string, labels []*ioPrometheusClient.LabelPair) bool {
	for _, label := range labels {
		if label.GetName() == key && label.GetValue() == value {
			return true
		}
	}

	return false
}
