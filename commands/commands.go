package commands

import (
	"startrader/globals"
)

type CliCommand struct {
	Name		string
	Description	string
	Callback	func(*globals.Config, ...string) error
}

func GetCommands() map[string]CliCommand{
	return map[string]CliCommand {
		"help": HelpCmdConfig,
		"exit": ExitCmdConfig,
	}
}