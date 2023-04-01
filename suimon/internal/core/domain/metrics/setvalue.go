package metrics

import (
	"fmt"
	"math"
	"strconv"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
)

// SetValue sets the value of a given metric type.
// Parameters:
// - metric: an enums.MetricType representing the metric type to set.
// - value: a value of any type representing the value to set for the given metric type.
func (metrics *Metrics) SetValue(metric enums.MetricType, value any) error {
	metrics.Updated = true

	var convFToI = func(input float64) int {
		return int(math.Round(input))
	}

	switch metric {
	case enums.MetricTypeSuiSystemState:
		valueSystemState, ok := value.(SuiSystemState)
		if !ok {
			return fmt.Errorf("unexpected value type for MetricTypeSuiSystemState: %T", value)
		}

		metrics.SystemState = valueSystemState
		metrics.TimeTillNextEpochMs = metrics.GetMillisecondsTillNextEpoch()
	case enums.MetricTypeTotalTransactionBlocks:
		switch v := value.(type) {
		case string:
			valueInt, err := strconv.Atoi(v)
			if err != nil {
				return err
			}
			metrics.TotalTransactionsBlocks = valueInt

			metrics.CalculateTPS()
		default:
			return fmt.Errorf("unexpected value type for MetricTypeTotalTransactionBlocks: %T", value)
		}
	case enums.MetricTypeTotalTransactionCertificates:
		valueFloat, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected value type for MetricTypeTotalTransactionCertificates: %T", value)
		}

		metrics.TotalTransactionCertificates = convFToI(valueFloat)
	case enums.MetricTypeTotalTransactionEffects:
		valueFloat, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected value type for MetricTypeTotalTransactionEffects: %T", value)
		}

		metrics.TotalTransactionEffects = convFToI(valueFloat)
	case enums.MetricTypeLatestCheckpoint:
		switch v := value.(type) {
		case string:
			valueInt, err := strconv.Atoi(v)
			if err != nil {
				return err
			}
			metrics.LatestCheckpoint = valueInt
		default:
			return fmt.Errorf("unexpected value type for MetricTypeLatestCheckpoint: %T", value)
		}
	case enums.MetricTypeHighestKnownCheckpoint:
		valueFloat, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected value type for MetricTypeHighestKnownCheckpoint: %T", value)
		}

		metrics.HighestKnownCheckpoint = convFToI(valueFloat)
	case enums.MetricTypeHighestSyncedCheckpoint:
		valueFloat, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected value type for MetricTypeHighestSyncedCheckpoint: %T", value)
		}

		metrics.HighestSyncedCheckpoint = convFToI(valueFloat)

		metrics.CalculateCPS()
	case enums.MetricTypeLastExecutedCheckpoint:
		valueFloat, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected value type for MetricTypeLastExecutedCheckpoint: %T", value)
		}

		metrics.LastExecutedCheckpoint = convFToI(valueFloat)
	case enums.MetricTypeCheckpointExecBacklog:
		valueInt, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected value type for MetricTypeCheckpointExecBacklog: %T", value)
		}

		metrics.CheckpointExecBacklog = valueInt
	case enums.MetricTypeCheckpointSyncBacklog:
		valueInt, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected value type for MetricTypeCheckpointSyncBacklog: %T", value)
		}

		metrics.CheckpointSyncBacklog = valueInt
	case enums.MetricTypeCurrentEpoch:
		valueFloat, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected value type for MetricTypeCurrentEpoch: %T", value)
		}

		metrics.CurrentEpoch = convFToI(valueFloat)
	case enums.MetricTypeEpochTotalDuration:
		valueFloat, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected value type for MetricTypeEpochTotalDuration: %T", value)
		}

		metrics.EpochTotalDuration = convFToI(valueFloat)
	case enums.MetricTypeSuiNetworkPeers:
		valueFloat, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected value type for MetricTypeSuiNetworkPeers: %T", value)
		}

		metrics.NetworkPeers = convFToI(valueFloat)
	case enums.MetricTypeUptime:
		valueFloat, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected value type for MetricTypeUptime: %T", value)
		}

		metrics.Uptime = fmt.Sprintf("%.2f", valueFloat/86400)
	case enums.MetricTypeVersion:
		valueString, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected value type for MetricTypeVersion: %T", value)
		}

		metrics.Version = valueString
	case enums.MetricTypeCommit:
		valueString, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected value type for MetricTypeCommit: %T", value)
		}

		metrics.Commit = valueString
	case enums.MetricTypeCurrentRound:
		valueFloat, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected value type for MetricTypeCurrentRound: %T", value)
		}

		metrics.CurrentRound = convFToI(valueFloat)
	case enums.MetricTypeHighestProcessedRound:
		valueFloat, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected value type for MetricTypeHighestProcessedRound: %T", value)
		}

		metrics.HighestProcessedRound = convFToI(valueFloat)
	case enums.MetricTypeLastCommittedRound:
		valueFloat, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected value type for MetricTypeLastCommittedRound: %T", value)
		}

		metrics.LastCommittedRound = convFToI(valueFloat)
	case enums.MetricTypePrimaryNetworkPeers:
		valueFloat, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected value type for MetricTypePrimaryNetworkPeers: %T", value)
		}

		metrics.PrimaryNetworkPeers = convFToI(valueFloat)
	case enums.MetricTypeWorkerNetworkPeers:
		valueFloat, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected value type for MetricTypeWorkerNetworkPeers: %T", value)
		}

		metrics.WorkerNetworkPeers = convFToI(valueFloat)
	case enums.MetricTypeSkippedConsensusTransactions:
		valueFloat, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected value type for MetricTypeSkippedConsensusTransactions: %T", value)
		}

		metrics.SkippedConsensusTransactions = convFToI(valueFloat)
	case enums.MetricTypeTotalSignatureErrors:
		valueFloat, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected value type for MetricTypeTotalSignatureErrors: %T", value)
		}

		metrics.TotalSignatureErrors = convFToI(valueFloat)
	case enums.MetricTypeTxSyncPercentage:
		valueInt, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected value type for MetricTypeTxSyncPercentage: %T", value)
		}

		metrics.TxSyncPercentage = valueInt
	case enums.MetricTypeCheckSyncPercentage:
		valueInt, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected value type for MetricTypeCheckSyncPercentage: %T", value)
		}

		metrics.CheckSyncPercentage = valueInt
	}

	return nil
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

	transactionsHistory = append(transactionsHistory, metrics.TotalTransactionsBlocks)
	if len(transactionsHistory) < TransactionsPerSecondWindow {
		metrics.TransactionsHistory = transactionsHistory

		return
	}

	if len(transactionsHistory) > TransactionsPerSecondWindow {
		transactionsHistory = transactionsHistory[1:]
	}

	transactionsStart = transactionsHistory[0]
	transactionsEnd = transactionsHistory[TransactionsPerSecondWindow-1]
	tps = (transactionsEnd - transactionsStart) / TransactionsPerSecondWindow

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
	if len(checkpointsHistory) < CheckpointsPerSecondWindow {
		metrics.CheckpointsHistory = checkpointsHistory

		return
	}

	if len(checkpointsHistory) > CheckpointsPerSecondWindow {
		checkpointsHistory = checkpointsHistory[1:]
	}

	checkpointsStart = checkpointsHistory[0]
	checkpointsEnd = checkpointsHistory[CheckpointsPerSecondWindow-1]
	cps = (checkpointsEnd - checkpointsStart) / CheckpointsPerSecondWindow

	metrics.CheckpointsHistory = checkpointsHistory
	metrics.CheckpointsPerSecond = cps
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
	case enums.MetricTypeTotalTransactionBlocks:
		return metrics.TxSyncPercentage >= TotalTransactionsSyncPercentage
	case enums.MetricTypeTransactionsPerSecond:
		valueRPCInt := valueRPC.(int)

		return metrics.TransactionsPerSecond >= valueRPCInt-TransactionsPerSecondLag
	case enums.MetricTypeLatestCheckpoint:
		valueRPCInt := valueRPC.(int)

		return metrics.CheckSyncPercentage >= TotalCheckpointsSyncPercentage || metrics.LatestCheckpoint >= valueRPCInt-LatestCheckpointLag
	case enums.MetricTypeHighestSyncedCheckpoint:
		valueRPCInt := valueRPC.(int)

		return metrics.CheckSyncPercentage >= TotalCheckpointsSyncPercentage || metrics.HighestSyncedCheckpoint >= valueRPCInt-HighestSyncedCheckpointLag
	case enums.MetricTypeCheckpointsPerSecond:
		valueRPCInt := valueRPC.(int)

		return metrics.CheckpointsPerSecond >= valueRPCInt-CheckpointsPerSecondLag
	case enums.MetricTypeVersion:
		return metrics.Version == valueRPC
	}

	return true
}

func (metrics *Metrics) IsUnhealthy(metric enums.MetricType, valueRPC any) bool {
	return !metrics.IsHealthy(metric, valueRPC)
}
