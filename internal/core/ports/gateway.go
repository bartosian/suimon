package ports

import (
	"net"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/bartosian/suimon/internal/core/domain/enums"
)

type RPCGateway interface {
	CallFor(method enums.RPCMethod, params ...interface{}) (result any, err error)
}

type PrometheusGateway interface {
	CallFor(metrics Metrics) (result MetricsResult, err error)
}

type GeoGateway interface {
	CallFor(ip net.IP) (result *IPResult, err error)
}

type MetricResult struct {
	Value  float64
	Labels prometheus.Labels
}

type MetricsResult map[enums.PrometheusMetricName]MetricResult

type MetricConfig struct {
	MetricType enums.PrometheusMetricType
	Labels     prometheus.Labels
}

type Metrics map[enums.PrometheusMetricName]MetricConfig

type Company struct {
	Name   string
	Domain string
	Type   string
}

type IPResult struct {
	IP           net.IP
	Hostname     string
	City         string
	Region       string
	Country      string
	CountryName  string
	CountryEmoji string
	Location     string
	Company      *Company
}
