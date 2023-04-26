package metrics

import (
	"fmt"
	"strconv"

	"github.com/dariubs/percent"
)

const validatorsQuorum = 6667

// SuiSystemState represents the current state of the Sui blockchain system.
type (
	SuiSystemState struct {
		Epoch                                 string          `json:"epoch"`
		ProtocolVersion                       string          `json:"protocolVersion"`
		SystemStateVersion                    string          `json:"systemStateVersion"`
		StorageFundTotalObjectStorageRebates  string          `json:"storageFundTotalObjectStorageRebates"`
		StorageFundNonRefundableBalance       string          `json:"storageFundNonRefundableBalance"`
		ReferenceGasPrice                     string          `json:"referenceGasPrice"`
		SafeMode                              bool            `json:"safeMode"`
		SafeModeStorageRewards                string          `json:"safeModeStorageRewards"`
		SafeModeComputationRewards            string          `json:"safeModeComputationRewards"`
		SafeModeStorageRebates                string          `json:"safeModeStorageRebates"`
		SafeModeNonRefundableStorageFee       string          `json:"safeModeNonRefundableStorageFee"`
		EpochStartTimestampMs                 string          `json:"epochStartTimestampMs"`
		EpochDurationMs                       string          `json:"epochDurationMs"`
		StakeSubsidyStartEpoch                string          `json:"stakeSubsidyStartEpoch"`
		MaxValidatorCount                     string          `json:"maxValidatorCount"`
		MinValidatorJoiningStake              string          `json:"minValidatorJoiningStake"`
		ValidatorLowStakeThreshold            string          `json:"validatorLowStakeThreshold"`
		ValidatorVeryLowStakeThreshold        string          `json:"validatorVeryLowStakeThreshold"`
		ValidatorLowStakeGracePeriod          string          `json:"validatorLowStakeGracePeriod"`
		StakeSubsidyBalance                   string          `json:"stakeSubsidyBalance"`
		StakeSubsidyDistributionCounter       string          `json:"stakeSubsidyDistributionCounter"`
		StakeSubsidyCurrentDistributionAmount string          `json:"stakeSubsidyCurrentDistributionAmount"`
		StakeSubsidyPeriodLength              string          `json:"stakeSubsidyPeriodLength"`
		StakeSubsidyDecreaseRate              int             `json:"stakeSubsidyDecreaseRate"`
		TotalStake                            string          `json:"totalStake"`
		ActiveValidators                      Validators      `json:"activeValidators"`
		PendingActiveValidatorsID             string          `json:"pendingActiveValidatorsId"`
		PendingActiveValidatorsSize           string          `json:"pendingActiveValidatorsSize"`
		PendingRemovals                       []interface{}   `json:"pendingRemovals"`
		StakingPoolMappingsID                 string          `json:"stakingPoolMappingsId"`
		StakingPoolMappingsSize               string          `json:"stakingPoolMappingsSize"`
		InactivePoolsID                       string          `json:"inactivePoolsId"`
		InactivePoolsSize                     string          `json:"inactivePoolsSize"`
		ValidatorCandidatesID                 string          `json:"validatorCandidatesId"`
		ValidatorCandidatesSize               string          `json:"validatorCandidatesSize"`
		AtRiskValidators                      [][]interface{} `json:"atRiskValidators"`
		ValidatorReportRecords                [][]interface{} `json:"validatorReportRecords"`
		AddressToValidator                    AddressToValidator
		ValidatorsAtRiskParsed                ValidatorsAtRisk
		ValidatorReportsParsed                ValidatorsReports
	}

	// ValidatorsReports is a map where the key is the address of a validator and the value is a ValidatorReports instance
	// containing all the reports for that validator
	ValidatorsReports []ValidatorReport

	// ValidatorReport represents validator reporters
	ValidatorReport struct {
		Name               string
		SlashingPercentage float64
		Reporters          []ValidatorReporter
	}

	// ValidatorReporter contains information about a validator reporter
	ValidatorReporter struct {
		Name        string
		Address     string
		VotingPower int
	}

	ValidatorsAtRisk []ValidatorAtRisk

	// ValidatorAtRisk represents a validator node at risk on the Sui blockchain network.
	ValidatorAtRisk struct {
		Name         string
		Address      string
		EpochsAtRisk string
	}
)

// parseValidatorsAtRisk is a helper function that parses the validators at risk from the raw JSON data
func (systemState *SuiSystemState) parseValidatorsAtRisk() error {
	if len(systemState.AtRiskValidators) == 0 {
		return nil
	}

	validatorsAtRisk := make([]ValidatorAtRisk, 0, len(systemState.AtRiskValidators))

	for _, atRiskValidator := range systemState.AtRiskValidators {
		address, ok := atRiskValidator[0].(string)
		if !ok {
			return fmt.Errorf(ErrUnsupportedValidatorsAtRiskAttr, atRiskValidator)
		}

		epochCount, ok := atRiskValidator[1].(string)
		if !ok {
			return fmt.Errorf(ErrUnsupportedValidatorsAtRiskAttr, atRiskValidator)
		}

		activeValidator, ok := systemState.AddressToValidator[address]
		if !ok {
			return fmt.Errorf("failed to loookup validator by address: %s", address)
		}

		validatorAtRisk := NewValidatorAtRisk(activeValidator.Name, address, epochCount)

		validatorsAtRisk = append(validatorsAtRisk, validatorAtRisk)
	}

	systemState.ValidatorsAtRiskParsed = validatorsAtRisk

	return nil
}

// parseValidatorReports parses the validator reports and calculates the slashing
// percentage for each validator based on the number of reporter validators and the
// validators quorum. The results are stored in the ValidatorReportsParsed field of
// the SuiSystemState struct.
func (systemState *SuiSystemState) parseValidatorReports() error {
	if len(systemState.ValidatorReportRecords) == 0 {
		return nil
	}

	validatorsReports := make(ValidatorsReports, 0, len(systemState.ValidatorReportRecords))

	for _, report := range systemState.ValidatorReportRecords {
		reportedAddress, ok := report[0].(string)
		if !ok {
			return fmt.Errorf(ErrUnsupportedValidatorsReportAttr, report)
		}

		reporters, ok := report[1].([]any)
		if !ok {
			return fmt.Errorf(ErrUnsupportedValidatorsReportAttr, report)
		}

		reportedValidator, ok := systemState.AddressToValidator[reportedAddress]
		if !ok {
			return fmt.Errorf("failed to loookup validator by address: %s", reportedAddress)
		}

		validatorReporters := make([]ValidatorReporter, 0, len(reporters))

		for _, reporterAddress := range reporters {
			reporter, ok := reporterAddress.(string)
			if !ok {
				return fmt.Errorf(ErrUnsupportedSuiAddressAttr, reporterAddress)
			}

			reporterValidator, ok := systemState.AddressToValidator[reporter]
			if !ok {
				return fmt.Errorf("failed to loookup validator by address: %s", reporterAddress)
			}

			reporterVotingPower, err := strconv.Atoi(reporterValidator.VotingPower)
			if err != nil {
				return fmt.Errorf("unexpected metric value type for VotingPower: %s", reporterValidator.VotingPower)
			}

			validatorReporter := NewValidatorReporter(reporterValidator.Name, reporterValidator.SuiAddress, reporterVotingPower)

			validatorReporters = append(validatorReporters, validatorReporter)
		}

		cumulativePower := 0

		for _, reporter := range validatorReporters {
			cumulativePower += reporter.VotingPower
		}

		slashingPercentage := percent.PercentOf(cumulativePower, validatorsQuorum)
		validatorReport := NewValidatorReport(reportedValidator.Name, slashingPercentage, validatorReporters)

		validatorsReports = append(validatorsReports, validatorReport)
	}

	systemState.ValidatorReportsParsed = validatorsReports

	return nil
}

// NewValidatorReporter creates a new ValidatorReporter instance with the specified
// name, address, and voting power.
func NewValidatorReporter(name, address string, votingPower int) ValidatorReporter {
	return ValidatorReporter{
		Name:        name,
		Address:     address,
		VotingPower: votingPower,
	}
}

// NewValidatorReport creates a new ValidatorReport instance with the specified
// name, slashing percentage, and reporters.
func NewValidatorReport(name string, slashingPct float64, reporters []ValidatorReporter) ValidatorReport {
	return ValidatorReport{
		Name:               name,
		SlashingPercentage: slashingPct,
		Reporters:          reporters,
	}
}

// NewValidatorAtRisk creates a new ValidatorAtRisk instance with the specified
// name, address, and epochs at risk.
func NewValidatorAtRisk(name, address, epochsAtRisk string) ValidatorAtRisk {
	return ValidatorAtRisk{
		Name:         name,
		Address:      address,
		EpochsAtRisk: epochsAtRisk,
	}
}
