package main

import (
	"flag"

	"github.com/bartosian/sui_helpers/sui-peer-checker/cmd/checker"
	"github.com/bartosian/sui_helpers/sui-peer-checker/cmd/checker/enums"
	"github.com/bartosian/sui_helpers/sui-peer-checker/pkg/log"
)

var (
	filePath = flag.String("f", "", "path to node config file")
	network  = flag.String("n", "devnet", "network name")
)

func main() {
	flag.Parse()

	logger := log.NewLogger()

	if *filePath == "" {
		logger.Error("provide path to the config file by using -f option", nil)

		return
	}

	network, err := enums.NetworkTypeFromString(*network)
	if err != nil {
		logger.Error("provide valid network type by using -n option", nil)

		return
	}

	checker, err := checker.NewChecker(*filePath, network)
	if err != nil {
		logger.Error("failed to create peers checker: ", err)

		return
	}

	checker.DrawTable()
}
