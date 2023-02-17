package checker

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
)

const transactionsPerSecondTimeout = 10

type (
	Metrics struct {
		Updated bool

		TransactionsHistory     []string
		TransactionsPerSecond   string
		TotalTransactionNumber  string
		HighestSyncedCheckpoint string
		LatestCheckpoint        string
		SuiNetworkPeers         string
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

	if len(transactionsHistory) != transactionsPerSecondTimeout {
		return
	}

	transactionsHistory = append(transactionsHistory[1:], metrics.TotalTransactionNumber)

	if transactionsStart, err = strconv.Atoi(transactionsHistory[0]); err != nil {
		return
	}

	if transactionsEnd, err = strconv.Atoi(transactionsHistory[transactionsPerSecondTimeout-1]); err != nil {
		return
	}

	tps := (transactionsEnd - transactionsStart) / transactionsPerSecondTimeout

	metrics.TransactionsPerSecond = fmt.Sprintf("%d", tps)
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
	case enums.MetricTypeLatestCheckpoint:
		metrics.LatestCheckpoint = value
	}
}

func (metrics *Metrics) IsHealthy(metric enums.MetricType, valueRPC string) bool {
	var convertToInt = func(values ...string) []int {
		valueAInt, err := strconv.Atoi(values[0])
		if err != nil {
			return nil
		}
		valueBInt, err := strconv.Atoi(values[1])
		if err != nil {
			return nil
		}

		return []int{valueAInt, valueBInt}
	}

	switch metric {
	case enums.MetricTypeTotalTransactionsNumber:
		values := convertToInt(metrics.TotalTransactionNumber, valueRPC)

		return len(values) == 2 && values[0] >= values[1]-100
	case enums.MetricTypeTransactionsPerSecond:
		values := convertToInt(metrics.TransactionsPerSecond, valueRPC)

		return len(values) == 2 && values[0] >= values[1]-5
	case enums.MetricTypeLatestCheckpoint:
		values := convertToInt(metrics.LatestCheckpoint, valueRPC)

		return len(values) == 2 && values[0] >= values[1]-10
	case enums.MetricTypeHighestSyncedCheckpoint:
		values := convertToInt(metrics.HighestSyncedCheckpoint, valueRPC)

		return len(values) == 2 && values[0] >= values[1]-10
	case enums.MetricTypeVersion:
		return metrics.Version == valueRPC
	}

	return true
}

func (metrics *Metrics) IsUnhealthy(metric enums.MetricType, valueRPC string) bool {
	return !metrics.IsHealthy(metric, valueRPC)
}
