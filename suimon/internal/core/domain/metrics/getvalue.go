package metrics

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dariubs/percent"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/pkg/utility"
)

type (
	MetricValue interface{}
)

// GetValue returns the metric value for the given metric type.
func (metrics *Metrics) GetValue(metric enums.MetricType) MetricValue {
	switch metric {
	case enums.MetricTypeSuiSystemState:
		return metrics.SystemState
	case enums.MetricTypeTotalTransactionBlocks:
		return metrics.TotalTransactionsBlocks
	case enums.MetricTypeTotalTransactionCertificates:
		return metrics.TotalTransactionCertificates
	case enums.MetricTypeTotalTransactionEffects:
		return metrics.TotalTransactionEffects
	case enums.MetricTypeTransactionsPerSecond:
		return metrics.TransactionsPerSecond
	case enums.MetricTypeLatestCheckpoint:
		return metrics.LatestCheckpoint
	case enums.MetricTypeHighestKnownCheckpoint:
		return metrics.HighestKnownCheckpoint
	case enums.MetricTypeHighestSyncedCheckpoint:
		return metrics.HighestSyncedCheckpoint
	case enums.MetricTypeLastExecutedCheckpoint:
		return metrics.LastExecutedCheckpoint
	case enums.MetricTypeCheckpointExecBacklog:
		return metrics.CheckpointExecBacklog
	case enums.MetricTypeCheckpointSyncBacklog:
		return metrics.CheckpointSyncBacklog
	case enums.MetricTypeCheckpointsPerSecond:
		return metrics.CheckpointsPerSecond
	case enums.MetricTypeCurrentEpoch:
		return metrics.CurrentEpoch
	case enums.MetricTypeEpochTotalDuration:
		return metrics.EpochTotalDuration
	case enums.MetricTypeTimeTillNextEpoch:
		return metrics.TimeTillNextEpoch
	case enums.MetricTypeTxSyncPercentage:
		return metrics.TotalTransactionsBlocks
	case enums.MetricTypeCheckSyncPercentage:
		return metrics.HighestSyncedCheckpoint
	case enums.MetricTypeSuiNetworkPeers:
		return metrics.NetworkPeers
	case enums.MetricTypeUptime:
		return metrics.Uptime
	case enums.MetricTypeVersion:
		return metrics.Version
	case enums.MetricTypeCommit:
		return metrics.Commit
	case enums.MetricTypeCurrentRound:
		return metrics.CurrentRound
	case enums.MetricTypeHighestProcessedRound:
		return metrics.HighestProcessedRound
	case enums.MetricTypeLastCommittedRound:
		return metrics.LastCommittedRound
	case enums.MetricTypeCertificatesCreated:
		return metrics.CertificatesCreated
	case enums.MetricTypePrimaryNetworkPeers:
		return metrics.PrimaryNetworkPeers
	case enums.MetricTypeWorkerNetworkPeers:
		return metrics.WorkerNetworkPeers
	case enums.MetricTypeSkippedConsensusTransactions:
		return metrics.SkippedConsensusTransactions
	case enums.MetricTypeTotalSignatureErrors:
		return metrics.TotalSignatureErrors
	default:
		return nil
	}
}

// GetMillisecondsTillNextEpoch returns the milliseconds till the next epoch.
func (metrics *Metrics) GetMillisecondsTillNextEpoch() (int64, error) {
	epochStartMs, err := strconv.ParseInt(metrics.SystemState.EpochStartTimestampMs, 10, 64)
	if err != nil {
		return 0, err
	}

	epochDurationMs, err := strconv.ParseInt(metrics.SystemState.EpochDurationMs, 10, 64)
	if err != nil {
		return 0, err
	}

	nextEpochStartMs := epochStartMs + epochDurationMs
	currentTimeMs := time.Now().UnixNano() / int64(time.Millisecond)

	return nextEpochStartMs - currentTimeMs, nil
}

// GetTimeUntilNextEpochDisplay returns the remaining time till the next epoch in human-readable format.
func (metrics *Metrics) GetTimeUntilNextEpochDisplay() []string {
	duration := time.Duration(metrics.TimeTillNextEpoch) * time.Millisecond
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) - (hours * 60)
	second := time.Now().Second()

	if hours < 0 {
		return nil
	}

	spacer := " "
	if second%2 == 0 {
		spacer = ":"
	}

	return []string{fmt.Sprintf("%02d%s%02d", hours, spacer, minutes), "H"}
}

// GetEpochLabel returns a string representing the current epoch number.
func (metrics *Metrics) GetEpochLabel() string {
	return fmt.Sprintf("EPOCH %d", metrics.SystemState.Epoch)
}

// GetEpochProgress calculates and returns the percentage of current epoch progress.
func (metrics *Metrics) GetEpochProgress() (int, error) {
	epochDurationMs, err := strconv.ParseInt(metrics.SystemState.EpochDurationMs, 10, 64)
	if err != nil {
		return 0, err
	}

	epochCurrentLength := epochDurationMs - metrics.TimeTillNextEpoch
	progressPercent := percent.PercentOf(int(epochCurrentLength), int(epochDurationMs))

	return int(progressPercent), nil
}

// GetUsageDataForDonutChart returns a label and percentage for displaying data in a donut chart.
// It takes in a metric unit and a function that returns a usage data and error.
func GetUsageDataForDonutChart(unit enums.MetricUnit, option func() (*utility.UsageData, error)) (string, int) {
	var (
		usageLabel      = "LOADING..."
		usagePercentage = 1
	)

	if usageDataResult, err := option(); err == nil {
		usageLabel = fmt.Sprintf("TOTAL/USED: %d/%d%s", usageDataResult.Total, usageDataResult.Used, unit)
		usagePercentage = usageDataResult.PercentageUsed

		if usagePercentage == 0 {
			usagePercentage = 1
		}
	}

	return usageLabel, usagePercentage
}
