package menus

import (
	"fmt"
	"os"
	"startrader/globals"
	"startrader/types"
	"startrader/utils"
)



func MainMenuIntro(m *types.Menu) {
	fmt.Println("\r", "----------------------------------------------------------------------------")
	fmt.Println("\r", "  _________ __                 ___________                  .___            ")
	fmt.Println("\r", " /   _____//  |______ _______  \\__    ___/___________     __| _/___________ ")
	fmt.Println("\r", " \\_____  \\\\   __\\__  \\\\_  __ \\   |    |  \\_  __ \\__  \\   / __ |/ __ \\_  __ \\")
	fmt.Println("\r", " /        \\|  |  / __ \\|  | \\/   |    |   |  | \\// __ \\_/ /_/ |  ___/|  | \\/")
	fmt.Println("\r", "/_______  /|__| (____  /__|      |____|   |__|  (____  /\\____ |\\___  >__|   ")
	fmt.Println("\r", "        \\/           \\/                              \\/      \\/    \\/       ")
	fmt.Println("\r", "----------------------------------------------------------------------------")
	fmt.Println("\r", m.Name, ":")
}


func NewGame() {
	utils.StartMissionTimers()
	globals.CurrentMenu = &CompanyMenu
}

func LoadGame() {
	fmt.Println("Load Game\n\r")
}

func Options() {
	fmt.Println("Options\n\r")
}

func Exit() {
	os.Exit(0)
}

var MainMenuOptions = []types.MenuItem {
	{
		Name: "New Game",
		Callback: NewGame,
	},
	{
		Name: "Load Game",
		Callback: LoadGame,
	},
	{
		Name: "Options",
		Callback: Options,
	},
	{
		Name: "Exit",
		Callback: Exit,
	},
}

var MainMenu = types.Menu {
	Name:		"Main Menu",
	Intro:		MainMenuIntro,
	Options: 	MainMenuOptions,
	Back:		Exit,
}