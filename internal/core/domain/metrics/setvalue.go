package metrics

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/big"
	"strconv"

	"github.com/bartosian/suimon/internal/core/domain/enums"
	"github.com/bartosian/suimon/internal/pkg/utility"
)

const (
	ErrUnexpectedMetricValueType       = "unexpected value type for %s: %T"
	ErrUnsupportedValidatorsAtRiskAttr = "unsupported validatorsAtRisk attribute type: %v"
	ErrUnsupportedValidatorsReportAttr = "unsupported validatorsReport attribute type: %v"
	ErrUnsupportedSuiAddressAttr       = "unsupported suiAddress attribute type: %v"
	utcTimeZone                        = "America/New_York"
	suiRate                            = 1_000_000_000
)

// SetValue updates a metric with the given value, parsing it if necessary.
// It returns an error if the value type is not supported for the given metric.
func (metrics *Metrics) SetValue(metric enums.MetricType, value any) error {
	metrics.Updated = true

	var convFToI = func(input float64) int {
		return int(math.Round(input))
	}

	switch metric {
	case enums.MetricTypeSuiSystemState:
		return metrics.SetSystemStateValue(value)
	case enums.MetricTypeValidatorsApy:
		return metrics.SetValidatorsApyValue(value)
	case enums.MetricTypeEpochsHistory:
		return metrics.SetEpochsHistoryValue(value)
	case enums.MetricTypeTotalTransactionBlocks:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf(ErrUnexpectedMetricValueType, metric, value)
		}

		valueInt, err := strconv.Atoi(v)
		if err != nil {
			return err
		}

		metrics.TotalTransactionsBlocks = valueInt
		metrics.CalculateTransactionsRatio()
	case enums.MetricTypeTotalTransactionCertificates:
		valueFloat, ok := value.(float64)
		if !ok {
			return fmt.Errorf(ErrUnexpectedMetricValueType, metric, value)
		}

		metrics.TotalTransactionCertificates = convFToI(valueFloat)
	case enums.MetricTypeTotalTransactionEffects:
		valueFloat, ok := value.(float64)
		if !ok {
			return fmt.Errorf(ErrUnexpectedMetricValueType, metric, value)
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
			return fmt.Errorf(ErrUnexpectedMetricValueType, metric, value)
		}
	case enums.MetricTypeHighestKnownCheckpoint:
		valueFloat, ok := value.(float64)
		if !ok {
			return fmt.Errorf(ErrUnexpectedMetricValueType, metric, value)
		}

		metrics.HighestKnownCheckpoint = convFToI(valueFloat)
	case enums.MetricTypeHighestSyncedCheckpoint:
		valueFloat, ok := value.(float64)
		if !ok {
			return fmt.Errorf(ErrUnexpectedMetricValueType, metric, value)
		}

		metrics.HighestSyncedCheckpoint = convFToI(valueFloat)

		metrics.CalculateCheckpointsRatio()
	case enums.MetricTypeLastExecutedCheckpoint:
		valueFloat, ok := value.(float64)
		if !ok {
			return fmt.Errorf(ErrUnexpectedMetricValueType, metric, value)
		}

		metrics.LastExecutedCheckpoint = convFToI(valueFloat)
	case enums.MetricTypeCheckpointExecBacklog:
		valueInt, ok := value.(int)
		if !ok {
			return fmt.Errorf(ErrUnexpectedMetricValueType, metric, value)
		}

		if valueInt < 0 {
			valueInt = 0
		}

		metrics.CheckpointExecBacklog = valueInt
	case enums.MetricTypeCheckpointSyncBacklog:
		valueInt, ok := value.(int)
		if !ok {
			return fmt.Errorf(ErrUnexpectedMetricValueType, metric, value)
		}

		if valueInt < 0 {
			valueInt = 0
		}

		metrics.CheckpointSyncBacklog = valueInt
	case enums.MetricTypeCurrentEpoch:
		valueFloat, ok := value.(float64)
		if !ok {
			return fmt.Errorf(ErrUnexpectedMetricValueType, metric, value)
		}

		metrics.CurrentEpoch = convFToI(valueFloat)
	case enums.MetricTypeEpochTotalDuration:
		valueFloat, ok := value.(float64)
		if !ok {
			return fmt.Errorf(ErrUnexpectedMetricValueType, metric, value)
		}

		metrics.EpochTotalDuration = convFToI(valueFloat)
	case enums.MetricTypeSuiNetworkPeers:
		valueFloat, ok := value.(float64)
		if !ok {
			return fmt.Errorf(ErrUnexpectedMetricValueType, metric, value)
		}

		metrics.NetworkPeers = convFToI(valueFloat)
	case enums.MetricTypeUptime:
		valueFloat, ok := value.(float64)
		if !ok {
			return fmt.Errorf(ErrUnexpectedMetricValueType, metric, value)
		}

		metrics.Uptime = fmt.Sprintf("%.2f", valueFloat/86400)
	case enums.MetricTypeVersion:
		valueString, ok := value.(string)
		if !ok {
			return fmt.Errorf(ErrUnexpectedMetricValueType, metric, value)
		}

		metrics.Version = valueString
	case enums.MetricTypeCommit:
		valueString, ok := value.(string)
		if !ok {
			return fmt.Errorf(ErrUnexpectedMetricValueType, metric, value)
		}

		metrics.Commit = valueString
	case enums.MetricTypeCurrentRound:
		valueFloat, ok := value.(float64)
		if !ok {
			return fmt.Errorf(ErrUnexpectedMetricValueType, metric, value)
		}

		metrics.CurrentRound = convFToI(valueFloat)
	case enums.MetricTypeHighestProcessedRound:
		valueFloat, ok := value.(float64)
		if !ok {
			return fmt.Errorf(ErrUnexpectedMetricValueType, metric, value)
		}

		metrics.HighestProcessedRound = convFToI(valueFloat)
		metrics.CalculateRoundsRatio()
	case enums.MetricTypeLastCommittedRound:
		valueFloat, ok := value.(float64)
		if !ok {
			return fmt.Errorf(ErrUnexpectedMetricValueType, metric, value)
		}

		metrics.LastCommittedRound = convFToI(valueFloat)
	case enums.MetricTypePrimaryNetworkPeers:
		valueFloat, ok := value.(float64)
		if !ok {
			return fmt.Errorf(ErrUnexpectedMetricValueType, metric, value)
		}

		metrics.PrimaryNetworkPeers = convFToI(valueFloat)
	case enums.MetricTypeWorkerNetworkPeers:
		valueFloat, ok := value.(float64)
		if !ok {
			return fmt.Errorf(ErrUnexpectedMetricValueType, metric, value)
		}

		metrics.WorkerNetworkPeers = convFToI(valueFloat)
	case enums.MetricTypeSkippedConsensusTransactions:
		valueFloat, ok := value.(float64)
		if !ok {
			return fmt.Errorf(ErrUnexpectedMetricValueType, metric, value)
		}

		metrics.SkippedConsensusTransactions = convFToI(valueFloat)
	case enums.MetricTypeCertificatesCreated:
		valueFloat, ok := value.(float64)
		if !ok {
			return fmt.Errorf(ErrUnexpectedMetricValueType, metric, value)
		}

		metrics.CertificatesCreated = convFToI(valueFloat)
		metrics.CalculateCertificatesRatio()
	case enums.MetricTypeTotalSignatureErrors:
		valueFloat, ok := value.(float64)
		if !ok {
			return fmt.Errorf(ErrUnexpectedMetricValueType, metric, value)
		}

		metrics.TotalSignatureErrors = convFToI(valueFloat)
	case enums.MetricTypeNonConsensusLatencySum:
		valueFloat, ok := value.(float64)
		if !ok {
			return fmt.Errorf(ErrUnexpectedMetricValueType, metric, value)
		}

		metrics.NonConsensusLatency = convFToI(valueFloat)
	case enums.MetricTypeTxSyncPercentage:
		valueInt, ok := value.(int)
		if !ok {
			return fmt.Errorf(ErrUnexpectedMetricValueType, metric, value)
		}

		metrics.TxSyncPercentage = valueInt
	case enums.MetricTypeCheckSyncPercentage:
		valueInt, ok := value.(int)
		if !ok {
			return fmt.Errorf(ErrUnexpectedMetricValueType, metric, value)
		}

		metrics.CheckSyncPercentage = valueInt
	}

	return nil
}

// SetSystemStateValue sets the SUI system state metrics based on the parsed data.
func (metrics *Metrics) SetSystemStateValue(value any) error {
	// Parse the JSON data of the SystemState object.
	dataBytes, err := json.Marshal(value.(map[string]interface{}))
	if err != nil {
		return fmt.Errorf(ErrUnexpectedMetricValueType, enums.MetricTypeSuiSystemState, value)
	}

	// Unmarshal the JSON data into a SuiSystemState struct.
	var valueSystemState SuiSystemState
	if err = json.Unmarshal(dataBytes, &valueSystemState); err != nil {
		return fmt.Errorf(ErrUnexpectedMetricValueType, enums.MetricTypeSuiSystemState, value)
	}

	// Create a mapping between validator addresses and their corresponding values.
	addressToValidator := make(AddressToValidator, len(valueSystemState.ActiveValidators))
	for _, activeValidator := range valueSystemState.ActiveValidators {
		addressToValidator[activeValidator.SuiAddress] = activeValidator
	}

	valueSystemState.AddressToValidator = addressToValidator

	// Parse the validators at risk from the raw JSON data.
	if err := valueSystemState.parseValidatorsAtRisk(); err != nil {
		return err
	}

	// Parse the validator reports from the raw JSON data.
	if err := valueSystemState.parseValidatorReports(); err != nil {
		return err
	}

	// Update the SystemState property of the Metrics struct with the parsed data.
	metrics.SystemState = valueSystemState

	// Calculate the epoch metrics.
	if err := metrics.setEpochMetrics(); err != nil {
		return err
	}

	// Calculate the reference gas price metrics.
	return metrics.setRefGasPriceMetrics()
}

// SetValidatorsApyValue sets the validators apy metrics based on the parsed data.
func (metrics *Metrics) SetValidatorsApyValue(value any) error {
	// Parse the JSON data of the ValidatorsAPY object.
	dataBytes, err := json.Marshal(value.(map[string]interface{}))
	if err != nil {
		return fmt.Errorf(ErrUnexpectedMetricValueType, enums.MetricTypeValidatorsApy, value)
	}

	// Unmarshal the JSON data into a ValidatorsApy struct.
	var validatorsApy ValidatorsApy
	if err = json.Unmarshal(dataBytes, &validatorsApy); err != nil {
		return fmt.Errorf(ErrUnexpectedMetricValueType, enums.MetricTypeValidatorsApy, value)
	}

	validatorsApyParsed := make(map[string]float64, len(validatorsApy.Apys))

	for _, validatorApy := range validatorsApy.Apys {
		validatorsApyParsed[validatorApy.Address] = validatorApy.Apy
	}

	metrics.ValidatorsApyParsed = validatorsApyParsed

	return nil
}

// SetEpochsHistoryValue sets the epochs history based on the parsed data.
func (metrics *Metrics) SetEpochsHistoryValue(value any) error {
	// Parse the JSON data of the EpochsHistory object.
	dataBytes, err := json.Marshal(value.(map[string]interface{}))
	if err != nil {
		return fmt.Errorf(ErrUnexpectedMetricValueType, enums.MetricTypeEpochsHistory, value)
	}

	// Unmarshal the JSON data into a EpochsHistory struct.
	var epochsHistory EpochsHistory
	if err = json.Unmarshal(dataBytes, &epochsHistory); err != nil {
		return fmt.Errorf(ErrUnexpectedMetricValueType, enums.MetricTypeEpochsHistory, value)
	}

	metrics.EpochsHistory = epochsHistory.Data[1:]

	return nil
}

// setEpochMetrics is a helper function that sets the epoch-related metrics based on the parsed data
func (metrics *Metrics) setEpochMetrics() error {
	systemState := metrics.SystemState

	epochStart, err := utility.ParseEpochTime(systemState.EpochStartTimestampMs)
	if err != nil {
		return err
	}

	metrics.EpochStartTimeUTC = utility.FormatDate(*epochStart, utcTimeZone)

	epochDuration, err := utility.StringMsToDuration(systemState.EpochDurationMs)
	if err != nil {
		return err
	}

	metrics.EpochDurationHHMM = utility.DurationToHoursAndMinutes(epochDuration)

	durationTillEpochEnd, err := utility.GetDurationTillTime(*epochStart, epochDuration)
	if err == nil {
		metrics.DurationTillEpochEndHHMM = utility.DurationToHoursAndMinutes(durationTillEpochEnd)
	}

	return nil
}

// setRefGasPriceMetrics is a helper function that sets the reference gas price metrics based on the parsed data
func (metrics *Metrics) setRefGasPriceMetrics() error {
	activeValidators := metrics.SystemState.ActiveValidators

	minRefGasPrice, err := activeValidators.GetMinRefGasPrice()
	if err != nil {
		return err
	}

	maxRefGasPrice, err := activeValidators.GetMaxRefGasPrice()
	if err != nil {
		return err
	}

	meanRefGasPrice, err := activeValidators.GetMeanRefGasPrice()
	if err != nil {
		return err
	}

	stakeWeightedMeanReferenceGasPrice, err := activeValidators.GetWeightedMeanRefGasPrice()
	if err != nil {
		return err
	}

	medianReferenceGasPrice, err := activeValidators.GetMedianRefGasPrice()
	if err != nil {
		return err
	}

	estimatedNextReferenceGasPrice, err := activeValidators.GetNextRefGasPrice()
	if err != nil {
		return err
	}

	metrics.MinReferenceGasPrice = minRefGasPrice
	metrics.MaxReferenceGasPrice = maxRefGasPrice
	metrics.MeanReferenceGasPrice = meanRefGasPrice
	metrics.StakeWeightedMeanReferenceGasPrice = stakeWeightedMeanReferenceGasPrice
	metrics.MedianReferenceGasPrice = medianReferenceGasPrice
	metrics.EstimatedNextReferenceGasPrice = estimatedNextReferenceGasPrice

	return nil
}

// CalculateTransactionsRatio calculates the current transaction per second (TPS) based on the number of transactions processed
// within the current period. The TPS value is then stored in the Metrics struct.
func (metrics *Metrics) CalculateTransactionsRatio() {
	var (
		transactionsHistory = metrics.TransactionsHistory
		transactionsStart   int
		transactionsEnd     int
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

	metrics.TransactionsHistory = transactionsHistory
	metrics.TransactionsPerSecond = transactionsEnd - transactionsStart
}

// CalculateCheckpointsRatio calculates the current checkpoints per second (CPS) based on the number of checkpoints generated
// within the current period. The CPS value is then stored in the Metrics struct.
func (metrics *Metrics) CalculateCheckpointsRatio() {
	var (
		checkpointsHistory = metrics.CheckpointsHistory
		checkpointsStart   int
		checkpointsEnd     int
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

	metrics.CheckpointsHistory = checkpointsHistory
	metrics.CheckpointsPerSecond = checkpointsEnd - checkpointsStart
}

// CalculateRoundsRatio calculates the current rounds per second (RPS) based on the number of rounds processed
// within the current period. The RPS value is then stored in the Metrics struct.
func (metrics *Metrics) CalculateRoundsRatio() {
	var (
		roundsHistory = metrics.RoundsHistory
		roundsStart   int
		roundsEnd     int
	)

	roundsHistory = append(roundsHistory, metrics.HighestProcessedRound)
	if len(roundsHistory) < RoundsPerSecondWindow {
		metrics.RoundsHistory = roundsHistory

		return
	}

	if len(roundsHistory) > RoundsPerSecondWindow {
		roundsHistory = roundsHistory[1:]
	}

	roundsStart = roundsHistory[0]
	roundsEnd = roundsHistory[RoundsPerSecondWindow-1]

	metrics.RoundsHistory = roundsHistory
	metrics.RoundsPerSecond = roundsEnd - roundsStart
}

// CalculateCertificatesRatio calculates the current certificates per second (CPS) based on the number of certificates created
// within the current period. The CPS value is then stored in the Metrics struct.
func (metrics *Metrics) CalculateCertificatesRatio() {
	var (
		certificatesHistory = metrics.CertificatesHistory
		certificatesStart   int
		certificatesEnd     int
	)

	certificatesHistory = append(certificatesHistory, metrics.CertificatesCreated)
	if len(certificatesHistory) < CertificatesPerSecondWindow {
		metrics.CertificatesHistory = certificatesHistory

		return
	}

	if len(certificatesHistory) > CertificatesPerSecondWindow {
		certificatesHistory = certificatesHistory[1:]
	}

	certificatesStart = certificatesHistory[0]
	certificatesEnd = certificatesHistory[CertificatesPerSecondWindow-1]

	metrics.CertificatesHistory = certificatesHistory
	metrics.CertificatesPerSecond = certificatesEnd - certificatesStart
}

// IsHealthy checks if the given metric's value satisfies the threshold defined for it.
// If the metric type is not recognized, returns true.
// The valueRPC argument is the value retrieved from the Sui RPC endpoint for the corresponding metric.
// Returns true if the metric value is healthy, false otherwise.
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

// GetMinRefGasPrice returns the minimum reference gas price among all validators.
// If there are no validators or if all validators have an invalid gas price, it returns an error.
func (validators Validators) GetMinRefGasPrice() (int, error) {
	if len(validators) == 0 {
		return 0, nil
	}

	var minRefGasPrice int = math.MaxInt

	for _, validator := range validators {
		validatorGasPrice, err := strconv.Atoi(validator.NextEpochGasPrice)
		if err != nil {
			return 0, fmt.Errorf("unexpected metric value type for NextEpochGasPrice: %s", validator.NextEpochGasPrice)
		}

		if validatorGasPrice < minRefGasPrice {
			minRefGasPrice = validatorGasPrice
		}
	}

	if minRefGasPrice == math.MaxInt {
		return 0, errors.New("no validators with valid gas price found")
	}

	return minRefGasPrice, nil
}

// MistToSui converts a string representing a value in "mist" units to its
// equivalent value in "sui" units.
//
// mist: a string representing a value in "mist" units
//
// Returns: the equivalent value in "sui" units, and an error if the input
// value cannot be converted to an integer.
func MistToSui(mist string) (int64, error) {
	// Convert the input string to an int64
	mistInt, ok := new(big.Int).SetString(mist, 10)
	if !ok {
		// Return an error if the input value cannot be converted to an integer
		return 0, fmt.Errorf("unexpected metric value type: %s", mist)
	}

	suiInt := new(big.Int).Div(mistInt, big.NewInt(suiRate))

	return suiInt.Int64(), nil
}
