package enums

type PrometheusMetricName string

const (
	PrometheusMetricNameTotalTransactionCertificates         PrometheusMetricName = "total_transaction_certificates"
	PrometheusMetricNameTotalTransactionEffects              PrometheusMetricName = "total_transaction_effects"
	PrometheusMetricNameHighestKnownCheckpoint               PrometheusMetricName = "highest_known_checkpoint"
	PrometheusMetricNameHighestSyncedCheckpoint              PrometheusMetricName = "highest_synced_checkpoint"
	PrometheusMetricNameLastExecutedCheckpoint               PrometheusMetricName = "last_executed_checkpoint"
	PrometheusMetricNameCurrentEpoch                         PrometheusMetricName = "current_epoch"
	PrometheusMetricNameEpochTotalDuration                   PrometheusMetricName = "epoch_total_duration"
	PrometheusMetricNameConsensusRoundProberCurrentRoundGaps PrometheusMetricName = "consensus_round_prober_current_round_gaps"
	PrometheusMetricNameSuiNetworkPeers                      PrometheusMetricName = "sui_network_peers"
	PrometheusMetricNameSkippedConsensusTransactions         PrometheusMetricName = "skipped_consensus_txns"
	PrometheusMetricNameTotalSignatureErrors                 PrometheusMetricName = "total_signature_errors"
	PrometheusMetricNameTotalTransactionCertificatesCreated  PrometheusMetricName = "total_tx_certificates_created"
	PrometheusMetricNameNonConsensusLatencySum               PrometheusMetricName = "validator_service_handle_certificate_non_consensus_latency_sum"
	PrometheusMetricNameUptime                               PrometheusMetricName = "uptime"
	PrometheusMetricNameCurrentVotingRight                   PrometheusMetricName = "current_voting_right"
	PrometheusMetricNameNumberSharedObjectTransactions       PrometheusMetricName = "num_shared_obj_tx"
	PrometheusMetricNameConsensusLastCommittedLeaderRound    PrometheusMetricName = "consensus_last_committed_leader_round"
	PrometheusMetricNameConsensusCommittedMessages           PrometheusMetricName = "consensus_committed_messages"
	PrometheusMetricNameConsensusProposedBlocks              PrometheusMetricName = "consensus_proposed_blocks"
	PrometheusMetricNameConsensusHighestAcceptedRound        PrometheusMetricName = "consensus_highest_accepted_round"
)

func (e PrometheusMetricName) ToString() string {
	return string(e)
}
