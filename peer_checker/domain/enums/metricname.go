package enums

import (
	"fmt"
	"strings"
)

type MetricName string

const (
	MetricNameUndefined               MetricName = "UNDEFINED"
	MetricNameUptime                  MetricName = "UPTIME"
	MetricNameVersion                 MetricName = "VERSION"
	MetricNameHighestSyncedCheckpoint MetricName = "HIGHEST_SYNCED_CHECKPOINT"
	MetricNameSuiNetworkPeers         MetricName = "SUI_NETWORK_PEERS"
)

func (e MetricName) String() string {
	return string(e)
}

func MetricNameFromString(value string) (MetricName, error) {
	value = strings.ToUpper(strings.TrimSpace(value))

	if strings.HasPrefix(value, MetricNameUptime.String()) {
		return MetricNameUptime, nil
	}

	result, ok := map[string]MetricName{
		"UPTIME":                    MetricNameUptime,
		"HIGHEST_SYNCED_CHECKPOINT": MetricNameHighestSyncedCheckpoint,
		"SUI_NETWORK_PEERS":         MetricNameSuiNetworkPeers,
	}[value]

	if ok {
		return result, nil
	}

	return MetricNameUndefined, fmt.Errorf("unsupported metric name enum string: %s", value)
}
