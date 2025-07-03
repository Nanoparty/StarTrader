package commands

import (
	"fmt"
	"startrader/globals"
)

var HelpCmdConfig CliCommand

func init() {
	HelpCmdConfig = CliCommand{
		Name:        "help",
		Description: "Shows available commands.",
		Callback:    CommandHelp,
	}
}

func CommandHelp(cfg *globals.Config, args ...string) error {
	fmt.Println()
	fmt.Println("-------------------------------")
	fmt.Println("Available Star Trader Commands:")
	fmt.Println("-------------------------------")
	for _, cmd := range GetCommands() {
		fmt.Printf("• %s: %s\n", cmd.Name, cmd.Description)
	}
	fmt.Println()
	return nil
}