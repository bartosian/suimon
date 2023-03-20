package checker

import (
	"fmt"
	"math"
	"time"

	"github.com/dariubs/percent"

	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
)

const (
	transactionsPerSecondWindow     = 10
	checkpointsPerSecondWindow      = 10
	transactionsPerSecondLag        = 5
	totalTransactionsLag            = 100
	checkpointsPerSecondLag         = 10
	latestCheckpointLag             = 30
	highestSyncedCheckpointLag      = 30
	totalTransactionsSyncPercentage = 99
	totalCheckpointsSyncPercentage  = 99
)

// GetTimeTillNextEpoch returns the time remaining until the next epoch in milliseconds.
// Returns: an integer representing the time remaining until the next epoch in seconds.
func (metrics *Metrics) GetTimeTillNextEpoch() int {
	nextEpochStartMs := metrics.SystemState.EpochStartTimestampMs + metrics.SystemState.EpochDurationMs
	currentTimeMs := int(time.Now().UnixNano() / 1000000)

	return nextEpochStartMs - currentTimeMs
}

// GetEpochTimer returns an array of strings representing the epoch timer data.
// Returns: an array of strings representing the epoch timer data.
func (metrics *Metrics) GetEpochTimer() []string {
	duration := time.Duration(metrics.TimeTillNextEpochMs) * time.Millisecond
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) - (hours * 60)
	second := time.Now().Second()

	if hours < 0 {
		return []string{""}
	}

	spacer := " "
	if second%2 == 0 {
		spacer = ":"
	}

	return []string{fmt.Sprintf("%02d%s%02d", hours, spacer, minutes), "H"}
}

// GetEpochLabel returns a string representing the epoch label.
// Returns: a string representing the epoch label.
func (metrics *Metrics) GetEpochLabel() string {
	return fmt.Sprintf("EPOCH %d", metrics.SystemState.Epoch)
}

// GetEpochProgress returns an integer representing the current epoch progress.
// Returns: an integer representing the current epoch progress.
func (metrics *Metrics) GetEpochProgress() int {
	epochCurrentLength := metrics.SystemState.EpochDurationMs - metrics.TimeTillNextEpochMs

	return int(percent.PercentOf(epochCurrentLength, metrics.SystemState.EpochDurationMs))
}

type Metrics struct {
	Updated bool

	SystemState SuiSystemState

	TotalTransactions            int
	TotalTransactionCertificates int
	TotalTransactionEffects      int
	TransactionsPerSecond        int
	TransactionsHistory          []int
	LatestCheckpoint             int
	HighestKnownCheckpoint       int
	HighestSyncedCheckpoint      int
	LastExecutedCheckpoint       int
	CheckpointsPerSecond         int
	CheckpointExecBacklog        int
	CheckpointSyncBacklog        int
	CheckpointsHistory           []int
	CurrentEpoch                 int
	EpochTotalDuration           int
	EpochPercentage              int
	TimeTillNextEpochMs          int
	TxSyncPercentage             int
	CheckSyncPercentage          int
	NetworkPeers                 int
	Uptime                       string
	Version                      string
	Commit                       string
	CurrentRound                 int
	HighestProcessedRound        int
	LastCommittedRound           int
	PrimaryNetworkPeers          int
	WorkerNetworkPeers           int
	SkippedConsensusTransactions int
	TotalSignatureErrors         int
}

// NewMetrics creates and returns a new instance of the Metrics struct.
// Returns: a new instance of the Metrics struct.
func NewMetrics() Metrics {
	return Metrics{
		TransactionsHistory: make([]int, 0, transactionsPerSecondWindow),
		CheckpointsHistory:  make([]int, 0, checkpointsPerSecondWindow),
	}
}

// CalculateTPS calculates the current transaction per second (TPS) based on the number of transactions processed
// within the current period. The TPS value is then stored in the Metrics struct.
func (metrics *Metrics) CalculateTPS() {
	var (
		transactionsHistory = metrics.TransactionsHistory
		transactionsStart   int
		transactionsEnd     int
		tps                 int
	)

	transactionsHistory = append(transactionsHistory, metrics.TotalTransactions)
	if len(transactionsHistory) < transactionsPerSecondWindow {
		metrics.TransactionsHistory = transactionsHistory

		return
	}

	if len(transactionsHistory) > transactionsPerSecondWindow {
		transactionsHistory = transactionsHistory[1:]
	}

	transactionsStart = transactionsHistory[0]
	transactionsEnd = transactionsHistory[transactionsPerSecondWindow-1]
	tps = (transactionsEnd - transactionsStart) / transactionsPerSecondWindow

	metrics.TransactionsHistory = transactionsHistory
	metrics.TransactionsPerSecond = tps
}

// CalculateCPS calculates the current checkpoints per second (CPS) based on the number of checkpoints generated
// within the current period. The CPS value is then stored in the Metrics struct.
func (metrics *Metrics) CalculateCPS() {
	var (
		checkpointsHistory = metrics.CheckpointsHistory
		checkpointsStart   int
		checkpointsEnd     int
		cps                int
	)

	checkpointsHistory = append(checkpointsHistory, metrics.HighestSyncedCheckpoint)
	if len(checkpointsHistory) < checkpointsPerSecondWindow {
		metrics.CheckpointsHistory = checkpointsHistory

		return
	}

	if len(checkpointsHistory) > checkpointsPerSecondWindow {
		checkpointsHistory = checkpointsHistory[1:]
	}

	checkpointsStart = checkpointsHistory[0]
	checkpointsEnd = checkpointsHistory[checkpointsPerSecondWindow-1]
	cps = (checkpointsEnd - checkpointsStart) / checkpointsPerSecondWindow

	metrics.CheckpointsHistory = checkpointsHistory
	metrics.CheckpointsPerSecond = cps
}

// SetValue sets the value of a given metric type.
// Parameters:
// - metric: an enums.MetricType representing the metric type to set.
// - value: a value of any type representing the value to set for the given metric type.
func (metrics *Metrics) SetValue(metric enums.MetricType, value any) {
	metrics.Updated = true

	var convFToI = func(input float64) int {
		return int(math.Round(input))
	}

	switch metric {
	case enums.MetricTypeSuiSystemState:
		if valueSystemState, ok := value.(SuiSystemState); ok {
			metrics.SystemState = valueSystemState
			metrics.TimeTillNextEpochMs = metrics.GetTimeTillNextEpoch()
		}
	case enums.MetricTypeTotalTransactions:
		valueInt := value.(int)

		metrics.TotalTransactions = valueInt

		metrics.CalculateTPS()
	case enums.MetricTypeTotalTransactionCertificates:
		valueFloat := value.(float64)

		metrics.TotalTransactionCertificates = convFToI(valueFloat)
	case enums.MetricTypeTotalTransactionEffects:
		valueFloat := value.(float64)

		metrics.TotalTransactionEffects = convFToI(valueFloat)
	case enums.MetricTypeLatestCheckpoint:
		valueInt := value.(int)

		metrics.LatestCheckpoint = valueInt
	case enums.MetricTypeHighestKnownCheckpoint:
		valueFloat := value.(float64)

		metrics.HighestKnownCheckpoint = convFToI(valueFloat)
	case enums.MetricTypeHighestSyncedCheckpoint:
		valueFloat := value.(float64)

		metrics.HighestSyncedCheckpoint = convFToI(valueFloat)

		metrics.CalculateCPS()
	case enums.MetricTypeLastExecutedCheckpoint:
		valueFloat := value.(float64)

		metrics.LastExecutedCheckpoint = convFToI(valueFloat)
	case enums.MetricTypeCheckpointExecBacklog:
		valueInt := value.(int)

		metrics.CheckpointExecBacklog = valueInt
	case enums.MetricTypeCheckpointSyncBacklog:
		valueInt := value.(int)

		metrics.CheckpointSyncBacklog = valueInt
	case enums.MetricTypeCurrentEpoch:
		valueFloat := value.(float64)

		metrics.CurrentEpoch = convFToI(valueFloat)
	case enums.MetricTypeEpochTotalDuration:
		valueFloat := value.(float64)

		metrics.EpochTotalDuration = convFToI(valueFloat)
	case enums.MetricTypeSuiNetworkPeers:
		valueFloat := value.(float64)

		metrics.NetworkPeers = convFToI(valueFloat)
	case enums.MetricTypeUptime:
		valueFloat := value.(float64)

		metrics.Uptime = fmt.Sprintf("%.2f", valueFloat)
	case enums.MetricTypeVersion:
		valueString := value.(string)

		metrics.Version = valueString
	case enums.MetricTypeCommit:
		valueString := value.(string)

		metrics.Commit = valueString
	case enums.MetricTypeCurrentRound:
		valueFloat := value.(float64)

		metrics.CurrentRound = convFToI(valueFloat)
	case enums.MetricTypeHighestProcessedRound:
		valueFloat := value.(float64)

		metrics.HighestProcessedRound = convFToI(valueFloat)
	case enums.MetricTypeLastCommittedRound:
		valueFloat := value.(float64)

		metrics.LastCommittedRound = convFToI(valueFloat)
	case enums.MetricTypePrimaryNetworkPeers:
		valueFloat := value.(float64)

		metrics.PrimaryNetworkPeers = convFToI(valueFloat)
	case enums.MetricTypeWorkerNetworkPeers:
		valueFloat := value.(float64)

		metrics.WorkerNetworkPeers = convFToI(valueFloat)
	case enums.MetricTypeSkippedConsensusTransactions:
		valueFloat := value.(float64)

		metrics.SkippedConsensusTransactions = convFToI(valueFloat)
	case enums.MetricTypeTotalSignatureErrors:
		valueFloat := value.(float64)

		metrics.TotalSignatureErrors = convFToI(valueFloat)
	case enums.MetricTypeTxSyncPercentage:
		valueInt := value.(int)

		metrics.TxSyncPercentage = valueInt
	case enums.MetricTypeCheckSyncPercentage:
		valueInt := value.(int)

		metrics.CheckSyncPercentage = valueInt
	}
}

// GetValue returns the current value of the specified metric.
// Parameters:
// - metric: an enums.MetricType representing the metric whose value to retrieve.
// - rpc: a boolean indicating whether to retrieve the value for RPC host.
// Returns: the current value of the specified metric, which can be of any type.
func (metrics *Metrics) GetValue(metric enums.MetricType, rpc bool) any {
	switch metric {
	case enums.MetricTypeUptime:
		return metrics.Uptime
	case enums.MetricTypeVersion:
		return metrics.Version
	case enums.MetricTypeCommit:
		return metrics.Commit
	case enums.MetricTypeHighestSyncedCheckpoint:
		return metrics.HighestSyncedCheckpoint
	case enums.MetricTypeSuiNetworkPeers:
		return metrics.NetworkPeers
	case enums.MetricTypeTotalTransactions:
		return metrics.TotalTransactions
	case enums.MetricTypeTransactionsPerSecond:
		return metrics.TransactionsPerSecond
	case enums.MetricTypeCheckpointsPerSecond:
		return metrics.CheckpointsPerSecond
	case enums.MetricTypeTxSyncPercentage:
		return metrics.TotalTransactions
	case enums.MetricTypeCheckSyncPercentage:
		if rpc {
			return metrics.LatestCheckpoint
		}

		return metrics.HighestSyncedCheckpoint
	case enums.MetricTypeLatestCheckpoint:
		return metrics.LatestCheckpoint
	default:
		return nil
	}
}

// IsHealthy checks if a given metric is healthy based on its current value.
// It returns a boolean value indicating whether the metric is healthy or not.
// A metric is considered healthy if it meets a certain condition or threshold,
// otherwise, it is considered unhealthy.
// Parameters:
//   - metric: the metric to check for healthiness.
//   - valueRPC: the current value of the metric, retrieved via an RPC call.
//
// Returns:
//   - a boolean value indicating whether the metric is healthy or not.
func (metrics *Metrics) IsHealthy(metric enums.MetricType, valueRPC any) bool {
	switch metric {
	case enums.MetricTypeTotalTransactions:
		return metrics.TxSyncPercentage >= totalTransactionsSyncPercentage
	case enums.MetricTypeTransactionsPerSecond:
		valueRPCInt := valueRPC.(int)

		return metrics.TransactionsPerSecond >= valueRPCInt-transactionsPerSecondLag
	case enums.MetricTypeLatestCheckpoint:
		valueRPCInt := valueRPC.(int)

		return metrics.CheckSyncPercentage >= totalCheckpointsSyncPercentage || metrics.LatestCheckpoint >= valueRPCInt-latestCheckpointLag
	case enums.MetricTypeHighestSyncedCheckpoint:
		valueRPCInt := valueRPC.(int)

		return metrics.CheckSyncPercentage >= totalCheckpointsSyncPercentage || metrics.HighestSyncedCheckpoint >= valueRPCInt-highestSyncedCheckpointLag
	case enums.MetricTypeCheckpointsPerSecond:
		valueRPCInt := valueRPC.(int)

		return metrics.CheckpointsPerSecond >= valueRPCInt-checkpointsPerSecondLag
	case enums.MetricTypeVersion:
		return metrics.Version == valueRPC
	}

	return true
}

func (metrics *Metrics) IsUnhealthy(metric enums.MetricType, valueRPC any) bool {
	return !metrics.IsHealthy(metric, valueRPC)
}
