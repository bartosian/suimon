package enums

type MetricType string

const (
	MetricTypeSuiSystemState               MetricType = "SYSTEM_STATE"
	MetricTypeTotalTransactionBlocks       MetricType = "TOTAL_TRANSACTION_BLOCKS"
	MetricTypeTotalTransactionCertificates MetricType = "TOTAL_TRANSACTION_CERTIFICATES"
	MetricTypeTotalTransactionEffects      MetricType = "TOTAL_TRANSACTION_EFFECTS"
	MetricTypeTransactionsPerSecond        MetricType = "TRANSACTIONS_PER_SECOND"
	MetricTypeLatestCheckpoint             MetricType = "LATEST_CHECKPOINT"
	MetricTypeHighestKnownCheckpoint       MetricType = "HIGHEST_KNOWN_CHECKPOINT"
	MetricTypeHighestSyncedCheckpoint      MetricType = "HIGHEST_SYNCED_CHECKPOINT"
	MetricTypeLastExecutedCheckpoint       MetricType = "LAST_EXECUTED_CHECKPOINT"
	MetricTypeCheckpointExecBacklog        MetricType = "CHECKPOINT_EXECUTION_BACKLOG"
	MetricTypeCheckpointSyncBacklog        MetricType = "CHECKPOINT_SYNC_BACKLOG"
	MetricTypeCheckpointsPerSecond         MetricType = "CHECKPOINTS_PER_SECOND"
	MetricTypeCurrentEpoch                 MetricType = "CURRENT_EPOCH"
	MetricTypeEpochTotalDuration           MetricType = "EPOCH_TOTAL_DURATION"
	MetricTypeTimeTillNextEpoch            MetricType = "TIME_TILL_NEXT_EPOCH"
	MetricTypeTxSyncPercentage             MetricType = "TX_SYNC_PERCENTAGE"
	MetricTypeCheckSyncPercentage          MetricType = "CHECK_SYNC_PERCENTAGE"
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
