package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"time"

	"github.com/dariubs/percent"
	"github.com/ybbus/jsonrpc/v3"
)

type Color string

const (
	ColorRed    = "\u001b[31m"
	ColorGreen  = "\u001b[32m"
	ColorYellow = "\u001b[33m"
	ColorReset  = "\u001b[0m"
)

const valVoteStakeLimit = 10

var (
	rpcURL         = flag.String("rpcURL", "https://fullnode.testnet.sui.io", "SUI RPC URL")
	setGas         = flag.Bool("setGas", false, "set ref gas price")
	frequency      = flag.Int("frequency", 30, "frequency of ref gas checks")
	previousRefGas uint64
)

func main() {
	flag.Parse()

	rpcClient := jsonrpc.NewClient(*rpcURL)
	ticker := time.NewTicker(time.Duration(*frequency) * time.Second)

	for {
		refGas, err := calculateRefGasPrice(rpcClient)
		if err != nil {
			colorize(ColorRed, fmt.Sprintf("\nError calculating reference gas price: %v\n\n", err))

			os.Exit(1)
		}

		colorize(ColorGreen, fmt.Sprintf("[-=-=-=-=- NEXT EPOCH REF GAS %d -=-=-=-=-]", *refGas))

		if *setGas && *refGas != previousRefGas {
			err = setGasPrice(*refGas - 1)
			if err != nil {
				colorize(ColorRed, fmt.Sprintf("\nError setting reference gas price: %v\n\n", err))

				os.Exit(1)
			}

			previousRefGas = *refGas
		}

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

func colorize(color Color, message string) {
	fmt.Println(string(color), message, string(ColorReset))
}

func setGasPrice(gas uint64) error {
	colorize(ColorYellow, fmt.Sprintf("[-=-=-=-=- SETTING REF GAS TO: %d -=-=-=-=-]", gas))

	gasCommand := `sui client call --package 0x2 --module sui_system --function request_set_gas_price --args 0x5 "425" --gas-budget 11000`

	_, err := exec.Command("bash", "-c", gasCommand).Output()
	if err != nil {
		return err
	}

	colorize(ColorYellow, fmt.Sprintf("[-=-=-=-=- UPDATED REF GAS TO: %d -=-=-=-=-]", gas))

	return nil
}
