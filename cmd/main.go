package main

import (
	"fmt"
	"os"

	"github.com/bartosian/sui_helpers/suimon/internal/core/controllers"
	"github.com/bartosian/sui_helpers/suimon/internal/core/controllers/monitor"
	domainconfig "github.com/bartosian/sui_helpers/suimon/internal/core/domain/config"
	"github.com/bartosian/sui_helpers/suimon/internal/core/gateways/cligw"
	"github.com/bartosian/sui_helpers/suimon/internal/core/handlers/commands"
)

func main() {
	cliGateway := cligw.NewGateway()

	defer handlePanic(cliGateway)

	config, err := domainconfig.NewConfig()
	if err != nil {
		// If an error occurs during initialization of the tables object, log the error and exit the program.
		cliGateway.Error(err.Error())

		return
	}

	// Instantiate controllers
	rootController := controllers.NewRootController(cliGateway)
	versionController := controllers.NewVersionController(cliGateway)
	monitorController := monitor.NewController(config, cliGateway)

	// Instantiate Handlers - Root
	rootCmdHandler := cmdhandlers.NewRootHandler(rootController)

	// Instantiate Handlers - second level
	versionCmdHandler := cmdhandlers.NewVersionHandler(versionController)
	monitorCmdHandler := cmdhandlers.NewMonitorHandler(monitorController)

	// Add subcommands to the root command handler
	rootCmdHandler.AddSubCommands(versionCmdHandler, monitorCmdHandler)

	// Start the root command handler
	rootCmdHandler.Start()
}

func handlePanic(cliGateway *cligw.Gateway) {
	if r := recover(); r != nil {
		// Handle the panic by logging the error and exiting the program
		cliGateway.Error(fmt.Sprintf("panic: %v", r))

		os.Exit(1)
	}
}
