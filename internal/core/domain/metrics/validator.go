package metrics

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"sort"
	"strconv"
)

// Validators represents a validator nodes on the Sui blockchain network.
type (
	Validators         []*Validator
	AddressToValidator map[string]*Validator
	Validator          struct {
		SuiAddress                   string      `json:"suiAddress"`
		ProtocolPubkeyBytes          string      `json:"protocolPubkeyBytes"`
		NetworkPubkeyBytes           string      `json:"networkPubkeyBytes"`
		WorkerPubkeyBytes            string      `json:"workerPubkeyBytes"`
		ProofOfPossessionBytes       string      `json:"proofOfPossessionBytes"`
		Name                         string      `json:"name"`
		Description                  string      `json:"description"`
		ImageURL                     string      `json:"imageUrl"`
		ProjectURL                   string      `json:"projectUrl"`
		NetAddress                   string      `json:"netAddress"`
		P2PAddress                   string      `json:"p2pAddress"`
		PrimaryAddress               string      `json:"primaryAddress"`
		WorkerAddress                string      `json:"workerAddress"`
		NextEpochProtocolPubkeyBytes interface{} `json:"nextEpochProtocolPubkeyBytes"`
		NextEpochProofOfPossession   interface{} `json:"nextEpochProofOfPossession"`
		NextEpochNetworkPubkeyBytes  interface{} `json:"nextEpochNetworkPubkeyBytes"`
		NextEpochWorkerPubkeyBytes   interface{} `json:"nextEpochWorkerPubkeyBytes"`
		NextEpochNetAddress          interface{} `json:"nextEpochNetAddress"`
		NextEpochP2PAddress          interface{} `json:"nextEpochP2pAddress"`
		NextEpochPrimaryAddress      interface{} `json:"nextEpochPrimaryAddress"`
		NextEpochWorkerAddress       interface{} `json:"nextEpochWorkerAddress"`
		VotingPower                  string      `json:"votingPower"`
		OperationCapID               string      `json:"operationCapId"`
		GasPrice                     string      `json:"gasPrice"`
		CommissionRate               string      `json:"commissionRate"`
		NextEpochStake               string      `json:"nextEpochStake"`
		NextEpochGasPrice            string      `json:"nextEpochGasPrice"`
		NextEpochCommissionRate      string      `json:"nextEpochCommissionRate"`
		StakingPoolID                string      `json:"stakingPoolId"`
		StakingPoolActivationEpoch   string      `json:"stakingPoolActivationEpoch"`
		StakingPoolDeactivationEpoch interface{} `json:"stakingPoolDeactivationEpoch"`
		StakingPoolSuiBalance        string      `json:"stakingPoolSuiBalance"`
		RewardsPool                  string      `json:"rewardsPool"`
		PoolTokenBalance             string      `json:"poolTokenBalance"`
		PendingStake                 string      `json:"pendingStake"`
		PendingTotalSuiWithdraw      string      `json:"pendingTotalSuiWithdraw"`
		PendingPoolTokenWithdraw     string      `json:"pendingPoolTokenWithdraw"`
		ExchangeRatesID              string      `json:"exchangeRatesId"`
		ExchangeRatesSize            string      `json:"exchangeRatesSize"`
		APY                          string
	}
)

// GetMaxRefGasPrice returns the maximum reference gas price among all validators.
// If there are no validators or if all validators have an invalid gas price, it returns an error.
func (validators Validators) GetMaxRefGasPrice() (int, error) {
	if len(validators) == 0 {
		return 0, nil
	}

	var maxRefGasPrice int = math.MinInt

	for _, validator := range validators {
		validatorGasPrice, err := strconv.Atoi(validator.NextEpochGasPrice)
		if err != nil {
			return 0, fmt.Errorf("unexpected metric value type for NextEpochGasPrice: %s", validator.NextEpochGasPrice)
		}

		if validatorGasPrice > maxRefGasPrice {
			maxRefGasPrice = validatorGasPrice
		}
	}

	if maxRefGasPrice == math.MinInt {
		return 0, errors.New("no validators with valid gas price found")
	}

	return maxRefGasPrice, nil
}

// GetMeanRefGasPrice calculates the mean reference gas price among all validators.
// If there are no validators or if all validators have an invalid gas price, it returns an error.
func (validators Validators) GetMeanRefGasPrice() (int, error) {
	if len(validators) == 0 {
		return 0, nil
	}

	var sumRefGasPrice int

	for _, validator := range validators {
		validatorGasPrice, err := strconv.Atoi(validator.NextEpochGasPrice)
		if err != nil {
			return 0, fmt.Errorf("unexpected metric value type for NextEpochGasPrice: %s", validator.NextEpochGasPrice)
		}

		sumRefGasPrice += validatorGasPrice
	}

	return sumRefGasPrice / len(validators), nil
}

// GetMedianRefGasPrice calculates the median reference gas price among all validators.
// If there are no validators or if all validators have an invalid gas price, it returns an error.
func (validators Validators) GetMedianRefGasPrice() (int, error) {
	if len(validators) == 0 {
		return 0, nil
	}

	gasPrices := make([]int, 0, len(validators))

	for _, validator := range validators {
		validatorGasPrice, err := strconv.Atoi(validator.NextEpochGasPrice)
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

func intToBigInt(x int) *big.Int {
	result := big.NewInt(int64(x))

	return result
}

func bigIntToInt(x *big.Int) int {
	result := x.Int64()

	return int(result)
}

// GetWeightedMeanRefGasPrice calculates the stake-weighted mean reference gas price among all validators.
// If there are no validators or if all validators have an invalid gas price or stake, it returns an error.
func (validators Validators) GetWeightedMeanRefGasPrice() (int, error) {
	if len(validators) == 0 {
		return 0, nil
	}

	var totalDelegation int

	gasMultiples := new(big.Int)

	for _, validator := range validators {
		validatorGasPrice, err := strconv.Atoi(validator.NextEpochGasPrice)
		if err != nil {
			return 0, fmt.Errorf("unexpected metric value type for NextEpochGasPrice: %s", validator.NextEpochGasPrice)
		}

		validatorTotalStake, err := strconv.Atoi(validator.NextEpochStake)
		if err != nil {
			return 0, fmt.Errorf("unexpected metric value type for NextEpochStake: %s", validator.NextEpochStake)
		}

		totalDelegation += validatorTotalStake

		mulResult := new(big.Int)
		mulResult.Mul(intToBigInt(validatorGasPrice), intToBigInt(validatorTotalStake))

		gasMultiples = gasMultiples.Add(gasMultiples, mulResult)
	}

	mean := gasMultiples.Div(gasMultiples, intToBigInt(totalDelegation))

	return bigIntToInt(mean), nil
}

// GetNextRefGasPrice calculates the next reference gas price for the Sui blockchain network.
// The reference gas price is determined by sorting the validators by their gas prices and
// selecting the gas price for which the cumulative voting power of validators exceeds
// two-thirds of the total voting power. If there are no validators or if all validators
// have an invalid gas price or voting power, it returns an error.
func (validators Validators) GetNextRefGasPrice() (int, error) {
	if len(validators) == 0 {
		return 0, nil
	}

	cumulativePower := 0
	referenceGasPrice := 0

	sort.SliceStable(validators, func(left, right int) bool {
		validatorLeftGasPrice, err := strconv.Atoi(validators[left].NextEpochGasPrice)
		if err != nil {
			return true
		}

		validatorRightGasPrice, err := strconv.Atoi(validators[right].NextEpochGasPrice)
		if err != nil {
			return true
		}

		return validatorLeftGasPrice < validatorRightGasPrice
	})

	for _, validator := range validators {
		if cumulativePower < validatorsQuorum {
			validatorGasPrice, err := strconv.Atoi(validator.NextEpochGasPrice)
			if err != nil {
				return 0, fmt.Errorf("unexpected metric value type for NextEpochGasPrice: %s", validator.NextEpochGasPrice)
			}

			referenceGasPrice = validatorGasPrice

			validatorVotingPower, err := strconv.Atoi(validator.VotingPower)
			if err != nil {
				return 0, fmt.Errorf("unexpected metric value type for VotingPower: %s", validator.VotingPower)
			}

			cumulativePower += validatorVotingPower
		}
	}

	return referenceGasPrice, nil
}
