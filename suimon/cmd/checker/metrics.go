package checker

import (
	"strconv"
	"strings"

	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
)

type (
	MetricsMap map[enums.MetricType]string

	Metrics struct {
		Updated bool

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

func (metrics *Metrics) IsHealthy(metric enums.MetricType, valueRPC string) bool {
	var convertToInt = func(values ...string) []int {
		valueAInt, err := strconv.Atoi(values[0])
		if err != nil {
			return nil
		}
		valueBInt, err := strconv.Atoi(values[1])
		if err != nil {
			return nil
		}

		return []int{valueAInt, valueBInt}
	}

	switch metric {
	case enums.MetricTypeTotalTransactionsNumber:
		values := convertToInt(metrics.TotalTransactionNumber, valueRPC)

		return values[0] >= values[1]-100
	case enums.MetricTypeLatestCheckpoint:
		values := convertToInt(metrics.LatestCheckpoint, valueRPC)

		return values[0] >= values[1]-10
	case enums.MetricTypeHighestSyncedCheckpoint:
		values := convertToInt(metrics.HighestSyncedCheckpoint, valueRPC)

		return values[0] >= values[1]-10
	case enums.MetricTypeVersion:
		return metrics.Version == valueRPC
	}

	return true
}

func (metrics *Metrics) IsUnhealthy(metric enums.MetricType, valueRPC string) bool {
	return !metrics.IsHealthy(metric, valueRPC)
}
