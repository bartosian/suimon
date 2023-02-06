package checker

import (
	"github.com/bartosian/sui_helpers/peer_checker/domain/enums"
	"strings"
)

type (
	MetricsMap map[enums.MetricName]string

	Metrics struct {
		HighestSyncedCheckpoint string
		SuiNetworkPeers         string
		Uptime                  string
		Version                 string
	}
)

func NewMetrics(input MetricsMap) *Metrics {
	metrics := new(Metrics)

	for metric, value := range input {
		switch metric {
		case enums.MetricNameUptime:
			metrics.Uptime = value
		case enums.MetricNameVersion:
			metrics.Version = strings.Trim(value, "\"")
		case enums.MetricNameHighestSyncedCheckpoint:
			metrics.HighestSyncedCheckpoint = value
		case enums.MetricNameSuiNetworkPeers:
			metrics.SuiNetworkPeers = value
		}
	}

	return metrics
}
