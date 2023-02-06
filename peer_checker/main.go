package main

import (
	"flag"
	"fmt"

	"github.com/bartosian/sui_helpers/peer_checker/domain/checker"
	"github.com/bartosian/sui_helpers/peer_checker/domain/enums"
)

var (
	filePath = flag.String("f", "", "path to node config file")
	network  = flag.String("n", "devnet", "network name")
)

func main() {
	flag.Parse()

	if filePath == nil {
		colorPrint(enums.ColorRed, "provide path to the config file by using -f option")

		return
	}

	network, err := enums.NetworkTypeFromString(*network)
	if err != nil {
		colorPrint(enums.ColorRed, "provide supported network type by using -n option")

		return
	}

	checker, err := checker.NewChecker(*filePath, network)
	if err != nil {
		colorPrint(enums.ColorRed, "failed to create peers checker: ", err.Error())

		return
	}

	checker.DrawTable()
}

func colorPrint(color enums.Color, messages ...any) {
	fmt.Println(color, messages, enums.ColorReset)
}
