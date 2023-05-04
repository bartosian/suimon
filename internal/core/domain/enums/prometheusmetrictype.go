package enums

type PrometheusMetricType int

const (
	PrometheusMetricTypeCounter PrometheusMetricType = iota
	PrometheusMetricTypeGauge
	PrometheusMetricTypeHistogram
	PrometheusMetricTypeSummary
	PrometheusMetricTypeUntyped
)
