package cmdhandlers

import (
	"github.com/spf13/cobra"

	"github.com/bartosian/sui_helpers/suimon/internal/core/ports"
)

type StaticHandler struct {
	command    *cobra.Command
	controller ports.VersionController
}

func NewStaticHandler(
	controller ports.VersionController,
) *StaticHandler {
	handler := &StaticHandler{
		controller: controller,
	}

	handler.command = handler.newCommand()

	return handler
}

func (h *StaticHandler) Start() {
	_ = h.command.Execute()
}

func (h *StaticHandler) AddSubCommands(subcommands ...ports.Command) {
	for _, subcommand := range subcommands {
		h.command.AddCommand(subcommand.Command())
	}
}

func (h *StaticHandler) Command() *cobra.Command {
	return h.command
}

func (h *StaticHandler) newCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "static",
		Aliases: []string{"s"},
		Short:   "suimon static",
		Long:    "suimon static",
		Run:     h.handleCommand,
	}

	return cmd
}

func (h *StaticHandler) handleCommand(_ *cobra.Command, _ []string) {
	h.controller.PrintVersion()
}
