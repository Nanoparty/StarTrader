package menus

import (
	"fmt"
	"os"
)

func QuitYes() {
	fmt.Println("\rQuitting game...\n\r")
	os.Exit(0)
}

func QuitNo() {
	CurrentMenu = &CompanyMenu
}

func QuitMenuIntro(m *Menu) {
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Println("\rAre you sure you want to quit the game?")
	fmt.Println("\r----------------------------------------------------------------------------")
}

var QuitMenuOptions []MenuItem
var QuitMenu Menu

func init() {
	QuitMenuOptions = []MenuItem{
		{Name: "Yes", Callback: QuitYes},
		{Name: "No", Callback: QuitNo},
	}
	QuitMenu = Menu{
		Name:    "Quit Game?",
		Intro:   QuitMenuIntro,
		Options: QuitMenuOptions,
		Back:    QuitNo,
	}
}
