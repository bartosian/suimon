package enums

type PrometheusMetricType int

const (
	PrometheusMetricTypeUntyped PrometheusMetricType = iota
	PrometheusMetricTypeCounter
	PrometheusMetricTypeGauge
	PrometheusMetricTypeHistogram
	PrometheusMetricTypeSummary
)
