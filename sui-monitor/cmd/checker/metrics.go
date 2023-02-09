package checker

import (
	"strings"

	"github.com/bartosian/sui_helpers/sui-monitor/cmd/checker/enums"
)

type (
	MetricsMap map[enums.MetricType]string

	Metrics struct {
		Updated                 bool
		TotalTransactionNumber  string
		HighestSyncedCheckpoint string
		LatestCheckpoint        string
		SuiNetworkPeers         string
		Uptime                  string
		Version                 string
		Commit                  string
	}
)

func (metrics *Metrics) SetValue(metric enums.MetricType, value string) {
	metrics.Updated = true

	switch metric {
	case enums.MetricTypeUptime:
		metrics.Uptime = value
	case enums.MetricTypeVersion:
		metrics.Version = strings.Trim(value, "\"")
	case enums.MetricTypeCommit:
		metrics.Commit = strings.Trim(value, "\"")
	case enums.MetricTypeHighestSyncedCheckpoint:
		metrics.HighestSyncedCheckpoint = value
	case enums.MetricTypeSuiNetworkPeers:
		metrics.SuiNetworkPeers = value
	case enums.MetricTypeTotalTransactionsNumber:
		metrics.TotalTransactionNumber = value
	case enums.MetricTypeLatestCheckpoint:
		metrics.LatestCheckpoint = value
	}
}
