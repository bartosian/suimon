package cmdhandlers

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/bartosian/sui_helpers/suimon/internal/core/ports"
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
		Short:   "Render monitoring dashboards",
		Long:    "Render monitoring dashboards",
		Run:     h.handleCommand,
	}

	return cmd
}

func (h *MonitorHandler) handleCommand(_ *cobra.Command, _ []string) {
	if err := h.controller.Monitor(); err != nil {
		fmt.Printf("Failed to run! %s\n", err)
	}
}