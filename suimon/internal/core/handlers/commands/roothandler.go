package cmdhandlers

import (
	"github.com/spf13/cobra"

	"github.com/bartosian/sui_helpers/suimon/internal/core/ports"
)

type RootHandler struct {
	command    *cobra.Command
	controller ports.RootController
}

func NewRootHandler(
	controller ports.RootController,
) *RootHandler {
	handler := &RootHandler{
		controller: controller,
	}

	handler.command = newCommand()

	return handler
}

func (h *RootHandler) Start() {
	_ = h.command.Execute()
}

func (h *RootHandler) AddSubCommands(subcommands ...ports.Command) {
	for _, subcommand := range subcommands {
		h.command.AddCommand(subcommand.Command())
	}
}

func newCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "suimon",
		Short: "Get real-time insights of SUI nodes and network performance.",
		Long:  "A comprehensive monitoring tool designed to provide real-time performance for SUI nodes and networks.\nWith an easy-to-install and user-friendly YAML configuration file, users can easily monitor network traffic, checkpoints, transactions, uptime, network status, peers, remote RPC, and more.\nFor help, use 'suimon --help'",
	}
}
