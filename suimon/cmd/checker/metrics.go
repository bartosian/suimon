package checker

import (
	"strconv"
	"strings"

	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
)

const (
	transactionsPerSecondTimeout   = 10
	transactionsPerSecondLag       = 5
	latestCheckpointLag            = 30
	highestSyncedCheckpointLag     = 30
	totalTransactionsNumberHealthy = 98
	totalTransactionsNumberLimit   = 100
)

type SuiSystemState struct {
	Epoch                 int `json:"epoch"`
	EpochStartTimestampMs int `json:"epoch_start_timestamp_ms"`
}

type Metrics struct {
	Updated bool

	SystemState SuiSystemState

	TxSyncPercentage        int
	CheckSyncPercentage     int
	TransactionsPerSecond   int
	TransactionsHistory     []int
	TotalTransactionNumber  int
	LatestCheckpoint        int
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
	case enums.MetricTypeCheckSystemState:
		valueSystemState := value.(SuiSystemState)

		metrics.SystemState = valueSystemState
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
		return metrics.TxSyncPercentage > totalTransactionsNumberHealthy && metrics.TxSyncPercentage <= totalTransactionsNumberLimit
	case enums.MetricTypeTransactionsPerSecond:
		valueRPCInt := valueRPC.(int)

		return metrics.TransactionsPerSecond >= valueRPCInt-transactionsPerSecondLag
	case enums.MetricTypeLatestCheckpoint:
		valueRPCInt := valueRPC.(int)

		return metrics.LatestCheckpoint >= valueRPCInt-latestCheckpointLag
	case enums.MetricTypeHighestSyncedCheckpoint:
		valueRPCInt := valueRPC.(int)

		return metrics.HighestSyncedCheckpoint >= valueRPCInt-highestSyncedCheckpointLag
	case enums.MetricTypeVersion:
		return metrics.Version == valueRPC
	}

	return true
}

func (metrics *Metrics) IsUnhealthy(metric enums.MetricType, valueRPC any) bool {
	return !metrics.IsHealthy(metric, valueRPC)
}
