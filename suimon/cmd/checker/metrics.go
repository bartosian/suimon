package checker

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
)

const (
	transactionsPerSecondTimeout   = 10
	transactionsPerSecondLag       = 5
	latestCheckpointLag            = 30
	highestSyncedCheckpointLag     = 30
	totalTransactionsNumberHealthy = "098"
	totalTransactionsNumberLimit   = "100"
)

type (
	Metrics struct {
		Updated bool

		TransactionsHistory     []string
		TransactionsPerSecond   string
		TotalTransactionNumber  string
		HighestSyncedCheckpoint string
		LatestCheckpoint        string
		SuiNetworkPeers         string
		TxSyncPercentage        string
		CheckSyncPercentage     string
		Uptime                  string
		Version                 string
		Commit                  string
	}
)

func NewMetrics() Metrics {
	return Metrics{
		TransactionsHistory: make([]string, 0, transactionsPerSecondTimeout),
	}
}

func (metrics *Metrics) CalculateTPS() {
	var (
		transactionsHistory = metrics.TransactionsHistory
		transactionsStart   int
		transactionsEnd     int
		err                 error
	)

	transactionsHistory = append(transactionsHistory, metrics.TotalTransactionNumber)
	if len(transactionsHistory) < transactionsPerSecondTimeout {
		metrics.TransactionsHistory = transactionsHistory

		return
	}

	if len(transactionsHistory) > transactionsPerSecondTimeout {
		transactionsHistory = transactionsHistory[1:]
	}

	if transactionsStart, err = strconv.Atoi(transactionsHistory[0]); err != nil {
		return
	}

	if transactionsEnd, err = strconv.Atoi(transactionsHistory[transactionsPerSecondTimeout-1]); err != nil {
		return
	}

	tps := (transactionsEnd - transactionsStart) / transactionsPerSecondTimeout

	metrics.TransactionsHistory = transactionsHistory
	metrics.TransactionsPerSecond = strconv.Itoa(tps)
}

func (metrics *Metrics) SetValue(metric enums.MetricType, value string) {
	metrics.Updated = true

	switch metric {
	case enums.MetricTypeUptime:
		metrics.Uptime = value
	case enums.MetricTypeVersion:
		metrics.Version = strings.Trim(value, "\"")
	case enums.MetricTypeCommit:
		metrics.Commit = strings.Trim(value, "\"")
	case enums.MetricTypeHighestSyncedCheckpoint:
		metrics.HighestSyncedCheckpoint = value
	case enums.MetricTypeSuiNetworkPeers:
		metrics.SuiNetworkPeers = value
	case enums.MetricTypeTotalTransactionsNumber:
		metrics.TotalTransactionNumber = value

		metrics.CalculateTPS()
	case enums.MetricTypeTxSyncProgress:
		metrics.TxSyncPercentage = value
	case enums.MetricTypeCheckSyncProgress:
		metrics.CheckSyncPercentage = value
	case enums.MetricTypeLatestCheckpoint:
		metrics.LatestCheckpoint = value
	}
}

func (metrics *Metrics) GetValue(metric enums.MetricType, rpc bool) string {
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
	}

	return ""
}

var convertToInt = func(values ...string) []int {
	result := make([]int, len(values))

	for idx, value := range values {
		if valueInt, err := strconv.Atoi(value); err != nil {
			return nil
		} else {
			result[idx] = valueInt
		}
	}

	return result
}

func (metrics *Metrics) IsHealthy(metric enums.MetricType, valueRPC string) bool {
	switch metric {
	case enums.MetricTypeTotalTransactionsNumber:
		paddedPercentage := fmt.Sprintf("%03s", metrics.TxSyncPercentage)
		return paddedPercentage > totalTransactionsNumberHealthy && paddedPercentage <= totalTransactionsNumberLimit
	case enums.MetricTypeTransactionsPerSecond:
		values := convertToInt(metrics.TransactionsPerSecond, valueRPC)

		return len(values) == 2 && values[0] >= values[1]-transactionsPerSecondLag
	case enums.MetricTypeLatestCheckpoint:
		values := convertToInt(metrics.LatestCheckpoint, valueRPC)

		return len(values) == 2 && values[0] >= values[1]-latestCheckpointLag
	case enums.MetricTypeHighestSyncedCheckpoint:
		values := convertToInt(metrics.HighestSyncedCheckpoint, valueRPC)

		return len(values) == 2 && values[0] >= values[1]-highestSyncedCheckpointLag
	case enums.MetricTypeVersion:
		return metrics.Version == valueRPC
	}

	return true
}

func (metrics *Metrics) IsUnhealthy(metric enums.MetricType, valueRPC string) bool {
	return !metrics.IsHealthy(metric, valueRPC)
}
