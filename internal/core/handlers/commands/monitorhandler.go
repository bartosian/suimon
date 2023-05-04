package cmdhandlers

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/bartosian/suimon/internal/core/ports"
)

type MonitorHandler struct {
	command    *cobra.Command
	controller ports.MonitorController
}

func NewMonitorHandler(
	controller ports.MonitorController,
) *MonitorHandler {
	handler := &MonitorHandler{
		controller: controller,
	}

	handler.command = handler.newCommand()

	return handler
}

func (h *MonitorHandler) Start() {
	_ = h.command.Execute()
}

func (h *MonitorHandler) AddSubCommands(subcommands ...ports.Command) {
	for _, subcommand := range subcommands {
		h.command.AddCommand(subcommand.Command())
	}
}

func (h *MonitorHandler) Command() *cobra.Command {
	return h.command
}

func (h *MonitorHandler) newCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "monitor",
		Aliases: []string{"m"},
		Short:   "Monitor the running network with the suimon monitoring tool.",
		Long:    "The suimon monitor subcommand allows you to monitor the running network with the suimon monitoring tool. This command provides options to render both static and dynamic dashboards. Static dashboards display various statistics related to the running network, such as the number of validators, peers, and gas prices. Dynamic dashboards provide real-time information about the network, such as block times and transaction throughput. You can select which dashboards to render using the command line interface. Use this command to keep an eye on the health and performance of your running network.",
		Run:     h.handleCommand,
	}

	return cmd
}

func (h *MonitorHandler) handleCommand(_ *cobra.Command, _ []string) {
	if err := h.controller.Monitor(); err != nil {
		fmt.Printf("Failed to run! %s\n", err)
	}
}
