package menus

import (
	"fmt"
	"os"
	"startrader/globals"
	"startrader/types"
)

func QuitYes() {
	fmt.Println("\rQuitting game...\n\r")
	fmt.Println("\rThe Solar System will be waiting for your return...\n\r")
	os.Exit(0)
}

func QuitNo() {
	globals.CurrentMenu = &CompanyMenu
}

func QuitMenuIntro(m *types.Menu) {
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Println("\rAre you sure you want to quit the game?")
	fmt.Println("\r----------------------------------------------------------------------------")
}

var QuitMenuOptions []types.MenuItem
var QuitMenu types.Menu

func init() {
	QuitMenuOptions = []types.MenuItem{
		{Name: "Yes", Callback: QuitYes},
		{Name: "No", Callback: QuitNo},
	}
	QuitMenu = types.Menu{
		Name:    "Quit Game?",
		Intro:   QuitMenuIntro,
		Options: QuitMenuOptions,
		Back:    QuitNo,
	}
}
