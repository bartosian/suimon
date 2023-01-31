package main

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/dariubs/percent"
	"github.com/ybbus/jsonrpc/v3"
)

type GasQuote struct {
	Quote uint64
	Stake uint64
}

const (
	execFrequency     = 20
	rpcURL            = "https://fullnode.testnet.sui.io"
	valVoteStakeLimit = 10
)

func main() {
	rpcClient := jsonrpc.NewClient(rpcURL)
	ticker := time.NewTicker(execFrequency * time.Second)

	for {
		refGas, err := calculateRefGasPrice(rpcClient)
		if err != nil {
			fmt.Printf("Error calculating reference gas price: %v", err)
		}

		fmt.Printf("[-=-=-=-=- NEXT EPOCH REF GAS %d -=-=-=-=-]\n", *refGas)

		<-ticker.C
	}
}

func calculateRefGasPrice(client jsonrpc.RPCClient) (gas *uint64, err error) {
	state := new(SuiSystemState)

	err = client.CallFor(context.Background(), &state, "sui_getSuiSystemState")
	if err != nil {
		return nil, err
	}

	var (
		gasQuotes           []GasQuote
		totalStake          uint64
		totalStake10Percent float64
		cumulativeStake     uint64
		referenceGasPrice   uint64
	)

	for _, validator := range state.Validators.ActiveValidators {
		nextEpochStake := validator.Metadata.NextEpochStake + validator.Metadata.NextEpochDelegation
		gasQuote := GasQuote{
			Quote: validator.Metadata.NextEpochGasPrice,
			Stake: nextEpochStake,
		}

		gasQuotes = append(gasQuotes, gasQuote)
		totalStake += nextEpochStake
	}

	totalStake10Percent = percent.Percent(valVoteStakeLimit, int(totalStake))
	countedStake := 2.0 / 3.0 * float64(totalStake)

	sort.Slice(gasQuotes, func(i, j int) bool { return gasQuotes[i].Quote < gasQuotes[j].Quote })

	for _, quote := range gasQuotes {
		if float64(quote.Stake) > totalStake10Percent {
			quote.Stake = uint64(totalStake10Percent)
		}

		cumulativeStake += quote.Stake
		referenceGasPrice = quote.Quote

		if float64(cumulativeStake) >= countedStake {
			break
		}
	}

	return &referenceGasPrice, nil
}
