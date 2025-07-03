package commands

import (
	"fmt"
	"os"
	"startrader/globals"
)

var ExitCmdConfig CliCommand

func init() {
	ExitCmdConfig = CliCommand{
		Name: "exit",
		Description: "Exit the game.",
		Callback: CommandExit,
	}
}

func CommandExit(cfg *globals.Config, args ...string) error {
	fmt.Println()
	fmt.Println("--------------------------------------")
	fmt.Println("Thank for playing Star Trader!")
	fmt.Println("The Solar System will be waiting... 👋")
	fmt.Println("--------------------------------------")
	fmt.Println()
	os.Exit(0)
	return nil
}
