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

	byteData, _ := io.ReadAll(jsonFile)

	var validators []Validator

	err = json.Unmarshal(byteData, &validators)
	if err != nil {
		fmt.Printf("could not unmarshal validators json: %s\n", err)

		return
	}

	sort.Slice(validators, func(a, b int) bool {
		return validators[a].NextEpochGasPrice < validators[b].NextEpochGasPrice
	})

	var totalStake int

	for _, validator := range validators {
		totalStake += validator.NextEpochDelegation + validator.NextEpochStake
	}

	totalStakeTwoThird := 2.0 / 3.0 * float64(totalStake)

	var (
		partialStake int
		refGasPrice  int
	)

	for _, validator := range validators {
		refGasPrice = validator.NextEpochGasPrice
		
		if partialStake += validator.NextEpochDelegation + validator.NextEpochStake; float64(partialStake) >= totalStakeTwoThird {			

			break
		}
	}

	fmt.Printf("\nNEXT EPOCH REFERENCE GAS PRICE: %d\n", refGasPrice)
}
