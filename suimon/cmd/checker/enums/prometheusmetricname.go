package enums

type PrometheusMetricName string

const (
	PrometheusMetricNameTotalTransactionCertificates PrometheusMetricName = "total_transaction_certificates"
	PrometheusMetricNameTotalTransactionEffects      PrometheusMetricName = "total_transaction_effects"
	PrometheusMetricNameHighestKnownCheckpoint       PrometheusMetricName = "highest_known_checkpoint"
	PrometheusMetricNameHighestSyncedCheckpoint      PrometheusMetricName = "highest_synced_checkpoint"
	PrometheusMetricNameCurrentEpoch                 PrometheusMetricName = "current_epoch"
	PrometheusMetricNameEpochTotalDuration           PrometheusMetricName = "epoch_total_duration"
	PrometheusMetricNameCurrentRound                 PrometheusMetricName = "current_round"
	PrometheusMetricNameHighestProcessedRound        PrometheusMetricName = "highest_processed_round"
	PrometheusMetricNameLastCommittedRound           PrometheusMetricName = "last_committed_round"
	PrometheusMetricNamePrimaryNetworkPeers          PrometheusMetricName = "primary_network_peers"
	PrometheusMetricNameWorkerNetworkPeers           PrometheusMetricName = "worker_network_peers"
	PrometheusMetricNameSuiNetworkPeers              PrometheusMetricName = "sui_network_peers"
	PrometheusMetricNameSkippedConsensusTransactions PrometheusMetricName = "skipped_consensus_txns"
	PrometheusMetricNameTotalSignatureErrors         PrometheusMetricName = "total_signature_errors"
	PrometheusMetricNameUptime                       PrometheusMetricName = "uptime"
)

func (e PrometheusMetricName) ToString() string {
	return string(e)
}
