package ports

import "github.com/spf13/cobra"

type Command interface {
	Start()
	AddSubCommands(subcommands ...Command)
	Command() *cobra.Command
}
