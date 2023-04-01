package metricsparser

import (
	"context"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/expfmt"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	ioPrometheusClient "github.com/prometheus/client_model/go"
)

const (
	httpClientTimeout = 2 * time.Second
)

type (
	MetricConfig struct {
		MetricType enums.PrometheusMetricType
		Labels     prometheus.Labels
	}
	MetricResult struct {
		Value  float64
		Labels prometheus.Labels
	}
	MetricsResult          map[enums.PrometheusMetricName]MetricResult
	PrometheusMetricParser struct {
		metricsURL string
		client     *http.Client
		parser     expfmt.TextParser

		Metrics map[enums.PrometheusMetricName]MetricConfig
	}
)

func NewPrometheusMetricParser(httpClient *http.Client, url string, metrics map[enums.PrometheusMetricName]MetricConfig) *PrometheusMetricParser {
	return &PrometheusMetricParser{
		metricsURL: url,
		client:     httpClient,
		parser:     expfmt.TextParser{},

		Metrics: metrics,
	}
}

func (mp *PrometheusMetricParser) GetMetrics() (MetricsResult, error) {
	req, err := http.NewRequest("GET", mp.metricsURL, nil)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), httpClientTimeout)
	defer cancel()

	req = req.WithContext(ctx)

	resp, err := mp.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	parser := expfmt.TextParser{}
	metrics, err := parser.TextToMetricFamilies(resp.Body)
	if err != nil {
		return nil, err
	}

	metricsResult := make(MetricsResult)

	for metricType, metricConfig := range mp.Metrics {
		result := getMetricValueWithLabelFiltering(metrics, metricType.ToString(), metricConfig)
		metricsResult[metricType] = result
	}

	return metricsResult, nil
}

func getMetricValueWithLabelFiltering(metrics map[string]*ioPrometheusClient.MetricFamily, metricName string, metricConfig MetricConfig) MetricResult {
	metricType := metricConfig.MetricType
	labels := metricConfig.Labels

	var result MetricResult

	metricFamily := metrics[metricName]
	if metricFamily == nil {
		return result
	}

	// Find the metric instance with the specified labels
OUTER:
	for _, metric := range metricFamily.Metric {
	INNER:
		for key, value := range labels {
			for _, label := range metric.Label {
				if label.GetName() == key && label.GetValue() == value {
					continue INNER
				}
			}

			continue OUTER
		}

		labelsResult := make(prometheus.Labels, len(metric.Label))

		for _, label := range metric.Label {
			labelsResult[label.GetName()] = label.GetValue()
		}

		result.Labels = labelsResult

		switch metricType {
		case enums.PrometheusMetricTypeGauge:
			result.Value = metric.GetGauge().GetValue()
		case enums.PrometheusMetricTypeCounter:
			result.Value = metric.GetCounter().GetValue()
		case enums.PrometheusMetricTypeSummary:
			result.Value = metric.GetSummary().GetSampleSum()
		case enums.PrometheusMetricTypeHistogram:
			result.Value = metric.GetHistogram().GetSampleSum()
		}
	}

	return result
}
