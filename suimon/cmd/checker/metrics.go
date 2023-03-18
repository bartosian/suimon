package checker

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dariubs/percent"

	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
)

const (
	transactionsPerSecondTimeout    = 10
	checkpointsPerSecondTimeout     = 10
	transactionsPerSecondLag        = 5
	totalTransactionsLag            = 100
	checkpointsPerSecondLag         = 10
	latestCheckpointLag             = 30
	highestSyncedCheckpointLag      = 30
	totalTransactionsSyncPercentage = 99
	totalCheckpointsSyncPercentage  = 99
)

type SuiSystemState struct {
	Epoch                 int `json:"epoch"`
	EpochStartTimestampMs int `json:"epoch_start_timestamp_ms"`
}

// GetTimeTillNextEpoch returns the time remaining until the next epoch in milliseconds.
// Returns: an integer representing the time remaining until the next epoch in seconds.
func (metrics *Metrics) GetTimeTillNextEpoch() int {
	nextEpochStartMs := metrics.SystemState.EpochStartTimestampMs + metrics.EpochLength
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
	epochCurrentLength := metrics.EpochLength - metrics.TimeTillNextEpochMs

	return int(percent.PercentOf(epochCurrentLength, metrics.EpochLength))
}

type Metrics struct {
	Updated bool

	SystemState SuiSystemState

	TxSyncPercentage        int
	EpochPercentage         int
	EpochLength             int
	TimeTillNextEpochMs     int
	CheckSyncPercentage     int
	TransactionsPerSecond   int
	TransactionsHistory     []int
	TotalTransactionNumber  int
	LatestCheckpoint        int
	CheckpointsPerSecond    int
	CheckpointsHistory      []int
	HighestSyncedCheckpoint int
	SuiNetworkPeers         int
	Uptime                  string
	Version                 string
	Commit                  string
}

// NewMetrics creates and returns a new instance of the Metrics struct, with the epoch length set to the given value.
// Parameters: epochLength: an integer representing the length of each epoch, in seconds.
// Returns: a new instance of the Metrics struct.
func NewMetrics(epochLength int) Metrics {
	return Metrics{
		EpochLength:         epochLength * 1000,
		TransactionsHistory: make([]int, 0, transactionsPerSecondTimeout),
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

	transactionsHistory = append(transactionsHistory, metrics.TotalTransactionNumber)
	if len(transactionsHistory) < transactionsPerSecondTimeout {
		metrics.TransactionsHistory = transactionsHistory

		return
	}

	if len(transactionsHistory) > transactionsPerSecondTimeout {
		transactionsHistory = transactionsHistory[1:]
	}

	transactionsStart = transactionsHistory[0]
	transactionsEnd = transactionsHistory[transactionsPerSecondTimeout-1]
	tps = (transactionsEnd - transactionsStart) / transactionsPerSecondTimeout

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
	if len(checkpointsHistory) < checkpointsPerSecondTimeout {
		metrics.CheckpointsHistory = checkpointsHistory

		return
	}

	if len(checkpointsHistory) > checkpointsPerSecondTimeout {
		checkpointsHistory = checkpointsHistory[1:]
	}

	checkpointsStart = checkpointsHistory[0]
	checkpointsEnd = checkpointsHistory[checkpointsPerSecondTimeout-1]
	cps = (checkpointsEnd - checkpointsStart) / checkpointsPerSecondTimeout

	metrics.CheckpointsHistory = checkpointsHistory
	metrics.CheckpointsPerSecond = cps
}

// SetValue sets the value of a given metric type.
// Parameters:
// - metric: an enums.MetricType representing the metric type to set.
// - value: a value of any type representing the value to set for the given metric type.
func (metrics *Metrics) SetValue(metric enums.MetricType, value any) {
	metrics.Updated = true

	switch metric {
	case enums.MetricTypeUptime:
		valueString := value.(string)

		metrics.Uptime = valueString
	case enums.MetricTypeVersion:
		valueString := value.(string)

		metrics.Version = strings.Trim(valueString, "\"")
	case enums.MetricTypeCommit:
		valueString := value.(string)

		metrics.Commit = strings.Trim(valueString, "\"")
	case enums.MetricTypeHighestSyncedCheckpoint:
		var (
			valueString = value.(string)
			valueInt    int
			err         error
		)

		if valueInt, err = strconv.Atoi(valueString); err != nil {
			return
		}

		metrics.HighestSyncedCheckpoint = valueInt

		metrics.CalculateCPS()
	case enums.MetricTypeSuiNetworkPeers:
		var (
			valueString = value.(string)
			valueInt    int
			err         error
		)

		if valueInt, err = strconv.Atoi(valueString); err != nil {
			return
		}

		metrics.SuiNetworkPeers = valueInt
	case enums.MetricTypeTotalTransactionsNumber:
		valueInt := value.(int)

		metrics.TotalTransactionNumber = valueInt

		metrics.CalculateTPS()
	case enums.MetricTypeTxSyncProgress:
		valueInt := value.(int)

		metrics.TxSyncPercentage = valueInt
	case enums.MetricTypeCheckSyncProgress:
		valueInt := value.(int)

		metrics.CheckSyncPercentage = valueInt
	case enums.MetricTypeLatestCheckpoint:
		valueInt := value.(int)

		metrics.LatestCheckpoint = valueInt
	case enums.MetricTypeCurrentEpoch:
		if valueSystemState, ok := value.(SuiSystemState); ok {
			metrics.SystemState = valueSystemState
			metrics.TimeTillNextEpochMs = metrics.GetTimeTillNextEpoch()
		}
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
		return metrics.SuiNetworkPeers
	case enums.MetricTypeTotalTransactionsNumber:
		return metrics.TotalTransactionNumber
	case enums.MetricTypeTransactionsPerSecond:
		return metrics.TransactionsPerSecond
	case enums.MetricTypeCheckpointsPerSecond:
		return metrics.CheckpointsPerSecond
	case enums.MetricTypeTxSyncProgress:
		return metrics.TotalTransactionNumber
	case enums.MetricTypeCheckSyncProgress:
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
	case enums.MetricTypeTotalTransactionsNumber:
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
