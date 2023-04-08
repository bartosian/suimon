package cmdhandlers

import (
	"github.com/spf13/cobra"

	"github.com/bartosian/sui_helpers/suimon/internal/core/ports"
)

type VersionHandler struct {
	command    *cobra.Command
	controller ports.VersionController
}

func NewVersionHandler(
	controller ports.VersionController,
) *VersionHandler {
	handler := &VersionHandler{
		controller: controller,
	}

	handler.command = handler.newCommand()

	return handler
}

func (h *VersionHandler) Start() {
	_ = h.command.Execute()
}

func (h *VersionHandler) AddSubCommands(subcommands ...ports.Command) {
	for _, subcommand := range subcommands {
		h.command.AddCommand(subcommand.Command())
	}
}

func (h *VersionHandler) Command() *cobra.Command {
	return h.command
}

func (h *VersionHandler) newCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "suimon version",
		Long:    "suimon version",
		Run:     h.handleCommand,
	}

	return cmd
}

func (h *VersionHandler) handleCommand(_ *cobra.Command, _ []string) {
	h.controller.PrintVersion()
}
