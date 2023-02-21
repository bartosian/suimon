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
	transactionsPerSecondTimeout   = 10
	checkpointsPerSecondTimeout    = 10
	transactionsPerSecondLag       = 5
	checkpointsPerSecondLag        = 5
	latestCheckpointLag            = 30
	highestSyncedCheckpointLag     = 30
	totalTransactionsNumberHealthy = 98
	epochLength                    = 24 * time.Hour
)

type SuiSystemState struct {
	Epoch                 int `json:"epoch"`
	EpochStartTimestampMs int `json:"epoch_start_timestamp_ms"`
}

func (metrics *Metrics) GetTimeTillNextEpoch() int {
	nextEpochStartMs := metrics.SystemState.EpochStartTimestampMs + int(epochLength.Milliseconds())
	currentTimeMs := int(time.Now().UnixNano() / 1000000)

	return nextEpochStartMs - currentTimeMs
}

func (metrics *Metrics) GetEpochTimer() string {
	duration := time.Duration(metrics.TimeTillNextEpochMs) * time.Millisecond
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) - (hours * 60)
	second := time.Now().Second()

	spacer := " "
	if second%2 == 0 {
		spacer = ":"
	}

	return fmt.Sprintf("%02d%s%02d", hours, spacer, minutes)
}

func (metrics *Metrics) GetEpochLabel() string {
	return fmt.Sprintf("EPOCH %d", metrics.SystemState.Epoch)
}

func (metrics *Metrics) GetEpochProgress() int {
	epochLength := int(epochLength.Milliseconds())
	epochCurrentLength := epochLength - metrics.TimeTillNextEpochMs

	return int(percent.PercentOf(epochCurrentLength, epochLength))
}

type Metrics struct {
	Updated bool

	SystemState SuiSystemState

	TxSyncPercentage        int
	EpochPercentage         int
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

func NewMetrics() Metrics {
	return Metrics{
		TransactionsHistory: make([]int, 0, transactionsPerSecondTimeout),
	}
}

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

func (metrics *Metrics) SetValue(metric enums.MetricType, value any) {
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

	metrics.Updated = true
}

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

func (metrics *Metrics) IsHealthy(metric enums.MetricType, valueRPC any) bool {
	switch metric {
	case enums.MetricTypeTotalTransactionsNumber:
		return metrics.TxSyncPercentage > totalTransactionsNumberHealthy
	case enums.MetricTypeTransactionsPerSecond:
		valueRPCInt := valueRPC.(int)

		return metrics.TransactionsPerSecond >= valueRPCInt-transactionsPerSecondLag
	case enums.MetricTypeLatestCheckpoint:
		valueRPCInt := valueRPC.(int)

		return metrics.LatestCheckpoint >= valueRPCInt-latestCheckpointLag
	case enums.MetricTypeHighestSyncedCheckpoint:
		valueRPCInt := valueRPC.(int)

		return metrics.HighestSyncedCheckpoint >= valueRPCInt-highestSyncedCheckpointLag
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
