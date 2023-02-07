package checker

import (
	"strings"

	"github.com/bartosian/sui_helpers/sui-peer-checker/cmd/checker/enums"
)

type (
	MetricsMap map[enums.MetricType]string

	Metrics struct {
		TotalTransactionNumber  string
		HighestSyncedCheckpoint string
		SuiNetworkPeers         string
		Uptime                  string
		Version                 string
	}
)

func NewMetrics(input MetricsMap) Metrics {
	var metrics Metrics

	for metric, value := range input {
		switch metric {
		case enums.MetricTypeUptime:
			metrics.Uptime = value
		case enums.MetricTypeVersion:
			metrics.Version = strings.Trim(value, "\"")
		case enums.MetricTypeHighestSyncedCheckpoint:
			metrics.HighestSyncedCheckpoint = value
		case enums.MetricTypeSuiNetworkPeers:
			metrics.SuiNetworkPeers = value
		case enums.MetricTypeTotalTransactionsNumber:
			metrics.TotalTransactionNumber = value
		}
	}

	return metrics
}
