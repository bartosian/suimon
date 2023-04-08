package prometheusgw

import (
	"context"
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	ioPrometheusClient "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/ports"
)

type MetricsData map[string]*ioPrometheusClient.MetricFamily

func (gateway *Gateway) CallFor(metrics ports.Metrics) (result ports.MetricsResult, err error) {
	if len(metrics) == 0 {
		return nil, fmt.Errorf("no metrics provided")
	}

	req, err := http.NewRequest("GET", gateway.url, nil)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), httpClientTimeout)
	defer cancel()

	req = req.WithContext(ctx)

	resp, err := gateway.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			err = closeErr
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	parser := expfmt.TextParser{}
	data, err := parser.TextToMetricFamilies(resp.Body)
	if err != nil {
		return nil, err
	}

	metricsResult := make(ports.MetricsResult)

	for metricType, metricConfig := range metrics {
		result, err := getMetricValueWithLabelFiltering(data, metricType.ToString(), metricConfig)
		if err != nil {
			return nil, err
		}

		metricsResult[metricType] = result
	}

	return metricsResult, nil
}

func getMetricValueWithLabelFiltering(metrics MetricsData, metricName string, metricConfig ports.MetricConfig) (result ports.MetricResult, err error) {
	if metrics == nil {
		return result, fmt.Errorf("no metrics provided")
	}

	metricType := metricConfig.MetricType
	labels := metricConfig.Labels

	metricFamily := metrics[metricName]
	if metricFamily == nil {
		return result, fmt.Errorf("no metric family found")
	}

	extractMetricValue := map[enums.PrometheusMetricType]func(metric *ioPrometheusClient.Metric) float64{
		enums.PrometheusMetricTypeGauge:     func(metric *ioPrometheusClient.Metric) float64 { return metric.GetGauge().GetValue() },
		enums.PrometheusMetricTypeCounter:   func(metric *ioPrometheusClient.Metric) float64 { return metric.GetCounter().GetValue() },
		enums.PrometheusMetricTypeSummary:   func(metric *ioPrometheusClient.Metric) float64 { return metric.GetSummary().GetSampleSum() },
		enums.PrometheusMetricTypeHistogram: func(metric *ioPrometheusClient.Metric) float64 { return metric.GetHistogram().GetSampleSum() },
	}

	foundMetric := false
	for _, metric := range metricFamily.Metric {
		matchLabels := true

		for key, value := range labels {
			if !labelMatches(key, value, metric.Label) {
				matchLabels = false

				break
			}
		}

		if !matchLabels {
			continue
		}

		foundMetric = true

		labelsResult := make(prometheus.Labels, len(metric.Label))

		for _, label := range metric.Label {
			labelsResult[label.GetName()] = label.GetValue()
		}

		result.Labels = labelsResult

		if extractValue := extractMetricValue[metricType]; extractValue != nil {
			result.Value = extractValue(metric)
		}
	}

	if !foundMetric {
		return result, fmt.Errorf("no metric found matching labels")
	}

	return result, nil
}

func labelMatches(key string, value string, labels []*ioPrometheusClient.LabelPair) bool {
	for _, label := range labels {
		if label.GetName() == key && label.GetValue() == value {
			return true
		}
	}
	return false
}
