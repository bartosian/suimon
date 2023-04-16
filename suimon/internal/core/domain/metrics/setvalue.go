package metrics

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"sort"
	"strconv"

	"encoding/json"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/pkg/utility"
)

const (
	ErrUnexpectedMetricValueType       = "unexpected value type for %s: %T"
	ErrUnsupportedValidatorsAtRiskAttr = "unsupported validatorsAtRisk attribute type: %v"
	ErrUnsupportedValidatorsReportAttr = "unsupported validatorsReport attribute type: %v"
	ErrUnsupportedSuiAddressAttr       = "unsupported suiAddress attribute type: %v"
	utcTimeZone                        = "America/New_York"
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
			return fmt.Errorf(ErrUnexpectedMetricValueType, metric, value)
		}
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

		metrics.CalculateCPS()
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

	// Create a mapping between validator addresses and their corresponding names.
	addressToValidatorName := make(map[string]string, len(valueSystemState.ActiveValidators))
	for _, activeValidator := range valueSystemState.ActiveValidators {
		addressToValidatorName[activeValidator.SuiAddress] = activeValidator.Name
	}
	valueSystemState.AddressToValidatorName = addressToValidatorName

	// Parse the validators at risk from the raw JSON data.
	if err := parseValidatorsAtRisk(valueSystemState, addressToValidatorName); err != nil {
		return err
	}

	// Parse the validator reports from the raw JSON data.
	if err := parseValidatorReports(valueSystemState, addressToValidatorName); err != nil {
		return err
	}

	// Update the SystemState property of the Metrics struct with the parsed data.
	metrics.SystemState = valueSystemState

	// Calculate the epoch metrics.
	if err := setEpochMetrics(metrics); err != nil {
		return err
	}

	// Calculate the reference gas price metrics.
	return setRefGasPriceMetrics(metrics)
}

// parseValidatorsAtRisk is a helper function that parses the validators at risk from the raw JSON data
func parseValidatorsAtRisk(systemState SuiSystemState, addressToValidatorName map[string]string) error {
	if len(systemState.AtRiskValidators) == 0 {
		return nil
	}

	validatorsAtRisk := make([]ValidatorAtRisk, 0, len(systemState.AtRiskValidators))

	for _, validator := range systemState.AtRiskValidators {
		address, ok := validator[0].(string)
		if !ok {
			return fmt.Errorf(ErrUnsupportedValidatorsAtRiskAttr, validator)
		}
		epochCount, ok := validator[1].(string)
		if !ok {
			return fmt.Errorf(ErrUnsupportedValidatorsAtRiskAttr, validator)
		}

		validatorName := addressToValidatorName[address]
		validatorAtRisk := NewValidatorAtRisk(validatorName, address, epochCount)

		validatorsAtRisk = append(validatorsAtRisk, validatorAtRisk)
	}

	systemState.ValidatorsAtRiskParsed = validatorsAtRisk

	return nil
}

// parseValidatorReports is a helper function that parses the validator reports from the raw JSON data
func parseValidatorReports(valueSystemState SuiSystemState, addressToValidatorName map[string]string) error {
	validatorReports := make([]ValidatorReport, 0, len(valueSystemState.ValidatorReportRecords))

	for _, report := range valueSystemState.ValidatorReportRecords {
		reportedAddress, ok := report[0].(string)
		if !ok {
			return fmt.Errorf(ErrUnsupportedValidatorsReportAttr, report)
		}
		reporters, ok := report[1].([]any)
		if !ok {
			return fmt.Errorf(ErrUnsupportedValidatorsReportAttr, report)
		}

		reportedName := addressToValidatorName[reportedAddress]

		for _, reporterAddress := range reporters {
			reporter, ok := reporterAddress.(string)
			if !ok {
				return fmt.Errorf(ErrUnsupportedSuiAddressAttr, reporterAddress)
			}

			reporterName := addressToValidatorName[reporter]
			validatorReport := NewValidatorReport(reportedName, reportedAddress, reporterName, reporter)

			validatorReports = append(validatorReports, validatorReport)
		}
	}

	valueSystemState.ValidatorReportsParsed = validatorReports

	return nil
}

// setEpochMetrics is a helper function that sets the epoch-related metrics based on the parsed data
func setEpochMetrics(metrics *Metrics) error {
	systemState := metrics.SystemState

	epochStart, err := utility.ParseEpochTime(systemState.EpochStartTimestampMs)
	if err != nil {
		return err
	}

	epochDuration, err := utility.StringMsToDuration(systemState.EpochDurationMs)
	if err != nil {
		return err
	}

	durationTillEpochEnd, err := utility.GetDurationTillTime(*epochStart, epochDuration)
	if err != nil {
		return err
	}

	metrics.EpochStartTimeUTC = utility.FormatDate(*epochStart, utcTimeZone)
	metrics.EpochDurationHHMM = utility.DurationToHoursAndMinutes(epochDuration)
	metrics.DurationTillEpochEndHHMM = utility.DurationToHoursAndMinutes(durationTillEpochEnd)

	return nil
}

// setRefGasPriceMetrics is a helper function that sets the reference gas price metrics based on the parsed data
func setRefGasPriceMetrics(metrics *Metrics) error {
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

	metrics.MinReferenceGasPrice = int(minRefGasPrice)
	metrics.MaxReferenceGasPrice = int(maxRefGasPrice)
	metrics.MeanReferenceGasPrice = int(meanRefGasPrice)
	metrics.StakeWeightedMeanReferenceGasPrice = int(stakeWeightedMeanReferenceGasPrice)
	metrics.MedianReferenceGasPrice = int(medianReferenceGasPrice)
	metrics.EstimatedNextReferenceGasPrice = int(estimatedNextReferenceGasPrice)

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
func (validators Validators) GetMinRefGasPrice() (int64, error) {
	if len(validators) == 0 {
		return 0, nil
	}

	var minRefGasPrice int64 = math.MaxInt64

	for _, validator := range validators {
		validatorGasPrice, err := strconv.ParseInt(validator.NextEpochGasPrice, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("unexpected metric value type for NextEpochGasPrice: %s", validator.NextEpochGasPrice)
		}

		if validatorGasPrice < minRefGasPrice {
			minRefGasPrice = validatorGasPrice
		}
	}

	if minRefGasPrice == math.MaxInt64 {
		return 0, errors.New("no validators with valid gas price found")
	}

	return minRefGasPrice, nil
}

// GetMaxRefGasPrice returns the maximum reference gas price among all validators.
// If there are no validators or if all validators have an invalid gas price, it returns an error.
func (validators Validators) GetMaxRefGasPrice() (int64, error) {
	if len(validators) == 0 {
		return 0, nil
	}

	var maxRefGasPrice int64 = math.MinInt64

	for _, validator := range validators {
		validatorGasPrice, err := strconv.ParseInt(validator.NextEpochGasPrice, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("unexpected metric value type for NextEpochGasPrice: %s", validator.NextEpochGasPrice)
		}

		if validatorGasPrice > maxRefGasPrice {
			maxRefGasPrice = validatorGasPrice
		}
	}

	if maxRefGasPrice == math.MinInt64 {
		return 0, errors.New("no validators with valid gas price found")
	}

	return maxRefGasPrice, nil
}

// GetMeanRefGasPrice calculates the mean reference gas price among all validators.
// If there are no validators or if all validators have an invalid gas price, it returns an error.
func (validators Validators) GetMeanRefGasPrice() (int64, error) {
	if len(validators) == 0 {
		return 0, nil
	}

	var sumRefGasPrice int64

	for _, validator := range validators {
		validatorGasPrice, err := strconv.ParseInt(validator.NextEpochGasPrice, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("unexpected metric value type for NextEpochGasPrice: %s", validator.NextEpochGasPrice)
		}

		sumRefGasPrice += validatorGasPrice
	}

	return sumRefGasPrice / int64(len(validators)), nil
}

// GetMedianRefGasPrice calculates the median reference gas price among all validators.
// If there are no validators or if all validators have an invalid gas price, it returns an error.
func (validators Validators) GetMedianRefGasPrice() (int64, error) {
	if len(validators) == 0 {
		return 0, nil
	}

	gasPrices := make([]int64, 0, len(validators))

	for _, validator := range validators {
		validatorGasPrice, err := strconv.ParseInt(validator.NextEpochGasPrice, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("unexpected metric value type for NextEpochGasPrice: %s", validator.NextEpochGasPrice)
		}

		gasPrices = append(gasPrices, validatorGasPrice)
	}

	sort.Slice(gasPrices, func(left, right int) bool {
		return gasPrices[left] < gasPrices[right]
	})

	if len(validators)%2 == 0 {
		return (gasPrices[len(validators)/2-1] + gasPrices[len(validators)/2]) / 2, nil
	}

	return gasPrices[len(validators)/2], nil
}

func int64ToBigInt(x int64) *big.Int {
	result := big.NewInt(x)
	return result
}

// GetWeightedMeanRefGasPrice calculates the stake-weighted mean reference gas price among all validators.
// If there are no validators or if all validators have an invalid gas price or stake, it returns an error.
func (validators Validators) GetWeightedMeanRefGasPrice() (int64, error) {
	if len(validators) == 0 {
		return 0, nil
	}

	var totalDelegation int64

	gasMultiples := new(big.Int)

	for _, validator := range validators {
		validatorGasPrice, err := strconv.ParseInt(validator.NextEpochGasPrice, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("unexpected metric value type for NextEpochGasPrice: %s", validator.NextEpochGasPrice)
		}

		validatorTotalStake, err := strconv.ParseInt(validator.NextEpochStake, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("unexpected metric value type for NextEpochStake: %s", validator.NextEpochStake)
		}

		totalDelegation += validatorTotalStake

		mulResult := new(big.Int)
		mulResult.Mul(int64ToBigInt(validatorGasPrice), int64ToBigInt(validatorTotalStake))

		gasMultiples = gasMultiples.Add(gasMultiples, mulResult)
	}

	return gasMultiples.Div(gasMultiples, int64ToBigInt(totalDelegation)).Int64(), nil
}

// GetNextRefGasPrice calculates the next reference gas price for the Sui blockchain network.
// The reference gas price is determined by sorting the validators by their gas prices and
// selecting the gas price for which the cumulative voting power of validators exceeds
// two-thirds of the total voting power. If there are no validators or if all validators
// have an invalid gas price or voting power, it returns an error.
func (validators Validators) GetNextRefGasPrice() (int64, error) {
	if len(validators) == 0 {
		return 0, nil
	}

	var (
		quorum            int64 = 6667
		cumulativePower   int64 = 0
		referenceGasPrice int64 = 0
	)

	sort.SliceStable(validators, func(left, right int) bool {
		validatorLeftGasPrice, err := strconv.ParseInt(validators[left].NextEpochGasPrice, 10, 64)
		if err != nil {
			return true
		}

		validatorRightGasPrice, err := strconv.ParseInt(validators[right].NextEpochGasPrice, 10, 64)
		if err != nil {
			return true
		}

		return validatorLeftGasPrice < validatorRightGasPrice
	})

	for _, validator := range validators {
		if cumulativePower < quorum {
			validatorGasPrice, err := strconv.ParseInt(validator.NextEpochGasPrice, 10, 64)
			if err != nil {
				return 0, fmt.Errorf("unexpected metric value type for NextEpochGasPrice: %s", validator.NextEpochGasPrice)
			}

			referenceGasPrice = validatorGasPrice

			validatorVotingPower, err := strconv.ParseInt(validator.VotingPower, 10, 64)
			if err != nil {
				return 0, fmt.Errorf("unexpected metric value type for VotingPower: %s", validator.VotingPower)
			}

			cumulativePower += validatorVotingPower
		}
	}

	return referenceGasPrice, nil
}
