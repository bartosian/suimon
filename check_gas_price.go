package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

type Validator struct {
	Name                    string `json:"name"`
	NextEpochStake          int    `json:"next_epoch_stake,string"`
	NextEpochGasPrice       int    `json:"next_epoch_gas_price,string"`
	NextEpochDelegation     int    `json:"next_epoch_delegation,string"`
	NextEpochCommissionRate int    `json:"next_epoch_commission_rate,string"`
}

func main() {
	jsonFile, err := os.Open("validators.json")
	if err != nil {
		fmt.Printf("could not open validators.json: %s\n", err)
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var validators []Validator

	err = json.Unmarshal(byteValue, &validators)
	if err != nil {
		fmt.Printf("could not unmarshal validators json: %s\n", err)

		return
	}

	sort.Slice(validators, func(valA, valB int) bool {
		return validators[valA].NextEpochGasPrice < validators[valB].NextEpochGasPrice
	})

	var totalStake int

	for _, validator := range validators {
		totalStake += validator.NextEpochDelegation
	}

	totalStakeTwoThird := totalStake / 3 * 2

	var (
		partialStake int
		refGasPrice  int
	)

	for _, validator := range validators {
		if partialStake += validator.NextEpochDelegation; partialStake <= totalStakeTwoThird {
			refGasPrice = validator.NextEpochGasPrice
		}
	}

	fmt.Printf("\nNEXT EPOCH REFERENCE GAS PRICE: %d\n", refGasPrice)
}
