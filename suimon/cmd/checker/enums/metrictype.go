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
	MetricTypeTxSyncProgress          MetricType = "TX_SYNC_PROGRESS"
	MetricTypeCheckSyncProgress       MetricType = "CHECK_SYNC_PROGRESS"
	MetricTypeCurrentEpoch            MetricType = "CURRENT_EPOCH"
	MetricTypeTimeTillNextEpoch       MetricType = "TIME_TILL_NEXT_EPOCH"
	MetricTypeCheckpointsPerSecond    MetricType = "CHECKPOINTS_PER_SECOND"
	MetricTypeSuiSystemState          MetricType = "SYSTEM_STATE"
)

func (e MetricType) ToString() string {
	return string(e)
}

func MetricTypeFromString(value string) (MetricType, error) {
	value = strings.ToUpper(strings.TrimSpace(value))

	if strings.HasPrefix(value, MetricTypeUptime.ToString()) {
		return MetricTypeUptime, nil
	}

	result, ok := map[string]MetricType{
		"UNDEFINED":                 MetricTypeUndefined,
		"UPTIME":                    MetricTypeUptime,
		"VERSION":                   MetricTypeVersion,
		"COMMIT":                    MetricTypeCommit,
		"HIGHEST_SYNCED_CHECKPOINT": MetricTypeHighestSyncedCheckpoint,
		"SUI_NETWORK_PEERS":         MetricTypeSuiNetworkPeers,
		"TRANSACTIONS_PER_SECOND":   MetricTypeTransactionsPerSecond,
		"TOTAL_TRANSACTIONS_NUMBER": MetricTypeTotalTransactionsNumber,
		"LATEST_CHECKPOINT":         MetricTypeLatestCheckpoint,
		"TX_SYNC_PROGRESS":          MetricTypeTxSyncProgress,
		"CHECK_SYNC_PROGRESS":       MetricTypeCheckSyncProgress,
		"CURRENT_EPOCH":             MetricTypeCurrentEpoch,
		"TIME_TILL_NEXT_EPOCH":      MetricTypeTimeTillNextEpoch,
		"CHECKPOINTS_PER_SECOND":    MetricTypeCheckpointsPerSecond,
		"SYSTEM_STATE":              MetricTypeSuiSystemState,
	}[value]

	if ok {
		return result, nil
	}

	return MetricTypeUndefined, fmt.Errorf("unsupported metric type enum string: %s", value)
}
