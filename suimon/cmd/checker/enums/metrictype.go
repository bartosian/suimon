package enums

import (
	"fmt"
	"strings"
)

type MetricType string

const (
	MetricTypeUndefined                    MetricType = "UNDEFINED"
	MetricTypeSuiSystemState               MetricType = "SYSTEM_STATE"
	MetricTypeTotalTransactions            MetricType = "TOTAL_TRANSACTIONS"
	MetricTypeTotalTransactionCertificates MetricType = "TOTAL_TRANSACTION_CERTIFICATES"
	MetricTypeTotalTransactionEffects      MetricType = "TOTAL_TRANSACTION_EFFECTS"
	MetricTypeTransactionsPerSecond        MetricType = "TRANSACTIONS_PER_SECOND"
	MetricTypeLatestCheckpoint             MetricType = "LATEST_CHECKPOINT"
	MetricTypeHighestKnownCheckpoint       MetricType = "HIGHEST_KNOWN_CHECKPOINT"
	MetricTypeHighestSyncedCheckpoint      MetricType = "HIGHEST_SYNCED_CHECKPOINT"
	MetricTypeCheckpointsPerSecond         MetricType = "CHECKPOINTS_PER_SECOND"
	MetricTypeCurrentEpoch                 MetricType = "CURRENT_EPOCH"
	MetricTypeEpochTotalDuration           MetricType = "EPOCH_TOTAL_DURATION"
	MetricTypeTimeTillNextEpoch            MetricType = "TIME_TILL_NEXT_EPOCH"
	MetricTypeTxSyncProgress               MetricType = "TX_SYNC_PROGRESS"
	MetricTypeCheckSyncProgress            MetricType = "CHECK_SYNC_PROGRESS"
	MetricTypeSuiNetworkPeers              MetricType = "SUI_NETWORK_PEERS"
	MetricTypeUptime                       MetricType = "UPTIME"
	MetricTypeVersion                      MetricType = "VERSION"
	MetricTypeCommit                       MetricType = "COMMIT"
	MetricTypeCurrentRound                 MetricType = "CURRENT_ROUND"
	MetricTypeHighestProcessedRound        MetricType = "HIGHEST_PROCESSED_ROUND"
	MetricTypeLastCommittedRound           MetricType = "LAST_COMMITTED_ROUND"
	MetricTypePrimaryNetworkPeers          MetricType = "PRIMARY_NETWORK_PEERS"
	MetricTypeWorkerNetworkPeers           MetricType = "WORKER_NETWORK_PEERS"
	MetricTypeSkippedConsensusTransactions MetricType = "SKIPPED_CONSENSUS_TRANSACTIONS"
	MetricTypeTotalSignatureErrors         MetricType = "TOTAL_SIGNATURE_ERRORS"
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
		"UNDEFINED":                      MetricTypeUndefined,
		"SYSTEM_STATE":                   MetricTypeSuiSystemState,
		"TOTAL_TRANSACTIONS":             MetricTypeTotalTransactions,
		"TOTAL_TRANSACTION_CERTIFICATES": MetricTypeTotalTransactionCertificates,
		"TOTAL_TRANSACTION_EFFECTS":      MetricTypeTotalTransactionEffects,
		"TRANSACTIONS_PER_SECOND":        MetricTypeTransactionsPerSecond,
		"LATEST_CHECKPOINT":              MetricTypeLatestCheckpoint,
		"HIGHEST_KNOWN_CHECKPOINT":       MetricTypeHighestKnownCheckpoint,
		"HIGHEST_SYNCED_CHECKPOINT":      MetricTypeHighestSyncedCheckpoint,
		"CHECKPOINTS_PER_SECOND":         MetricTypeCheckpointsPerSecond,
		"CURRENT_EPOCH":                  MetricTypeCurrentEpoch,
		"EPOCH_TOTAL_DURATION":           MetricTypeEpochTotalDuration,
		"TIME_TILL_NEXT_EPOCH":           MetricTypeTimeTillNextEpoch,
		"TX_SYNC_PROGRESS":               MetricTypeTxSyncProgress,
		"CHECK_SYNC_PROGRESS":            MetricTypeCheckSyncProgress,
		"SUI_NETWORK_PEERS":              MetricTypeSuiNetworkPeers,
		"UPTIME":                         MetricTypeUptime,
		"VERSION":                        MetricTypeVersion,
		"COMMIT":                         MetricTypeCommit,
		"CURRENT_ROUND":                  MetricTypeCurrentRound,
		"HIGHEST_PROCESSED_ROUND":        MetricTypeHighestProcessedRound,
		"LAST_COMMITTED_ROUND":           MetricTypeLastCommittedRound,
		"PRIMARY_NETWORK_PEERS":          MetricTypePrimaryNetworkPeers,
		"WORKER_NETWORK_PEERS":           MetricTypeWorkerNetworkPeers,
		"SKIPPED_CONSENSUS_TRANSACTIONS": MetricTypeSkippedConsensusTransactions,
		"TOTAL_SIGNATURE_ERRORS":         MetricTypeTotalSignatureErrors,
	}[value]

	if ok {
		return result, nil
	}

	return MetricTypeUndefined, fmt.Errorf("unsupported metric type enum string: %s", value)
}
