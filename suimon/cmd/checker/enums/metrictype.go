package enums

import (
	"fmt"
	"strings"
)

type MetricType string

const (
	MetricTypeUndefined               MetricType = "UNDEFINED"
	MetricTypeUptime                  MetricType = "UPTIME"
	MetricTypeVersion                 MetricType = "VERSION"
	MetricTypeCommit                  MetricType = "COMMIT"
	MetricTypeHighestSyncedCheckpoint MetricType = "HIGHEST_SYNCED_CHECKPOINT"
	MetricTypeSuiNetworkPeers         MetricType = "SUI_NETWORK_PEERS"
	MetricTypeTransactionsPerSecond   MetricType = "TRANSACTIONS_PER_SECOND"
	MetricTypeTotalTransactionsNumber MetricType = "TOTAL_TRANSACTIONS_NUMBER"
	MetricTypeLatestCheckpoint        MetricType = "LATEST_CHECKPOINT"
)

func (e MetricType) String() string {
	return string(e)
}

func MetricTypeFromString(value string) (MetricType, error) {
	value = strings.ToUpper(strings.TrimSpace(value))

	if strings.HasPrefix(value, MetricTypeUptime.String()) {
		return MetricTypeUptime, nil
	}

	result, ok := map[string]MetricType{
		"UPTIME":                    MetricTypeUptime,
		"VERSION":                   MetricTypeVersion,
		"COMMIT":                    MetricTypeCommit,
		"HIGHEST_SYNCED_CHECKPOINT": MetricTypeHighestSyncedCheckpoint,
		"SUI_NETWORK_PEERS":         MetricTypeSuiNetworkPeers,
		"TRANSACTIONS_PER_SECOND":   MetricTypeTransactionsPerSecond,
		"TOTAL_TRANSACTIONS_NUMBER": MetricTypeTotalTransactionsNumber,
		"LATEST_CHECKPOINT":         MetricTypeLatestCheckpoint,
	}[value]

	if ok {
		return result, nil
	}

	return MetricTypeUndefined, fmt.Errorf("unsupported metric type enum string: %s", value)
}
