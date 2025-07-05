package menus

import (
	"fmt"
	"startrader/globals"
)


func Ships() {
	BuildCompanyShipsMenuOptions()
	CurrentMenu = &CompanyShipsMenu
}

func Pilots() {
	BuildCompanyPilotsMenuOptions()
	CurrentMenu = &CompanyPilotsMenu
}

func Information() {
	fmt.Println("Information\n\r")
}


func Locations() {
	CurrentMenu = &SectorsMenu
}

func QuitGame() {
	CurrentMenu = &QuitMenu
}

func CompanyMenuIntro(m *Menu) {
	header := "Company Menu: " + globals.CompanyName
	moneyHeader := fmt.Sprintf("%s%d", "$", CompanyMoney)
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Printf("\r%s%*s%s\n\r", header, 76 - len(header) - len(moneyHeader), "", moneyHeader)
	fmt.Println("\r----------------------------------------------------------------------------")
}

var CompanyMenuOptions []MenuItem
var CompanyMenu Menu

func init() {
	CompanyMenuOptions = []MenuItem{
		{Name: "Ships", Callback: Ships},
		{Name: "Pilots", Callback: Pilots},
		{Name: "Active Missions", Callback: func() { BuildActiveMissionsMenuOptions(); CurrentMenu = &ActiveMissionsMenu }},
		{Name: "Information", Callback: func() { CurrentMenu = &CompanyInformationMenu }},

		{Name: "Locations", Callback: Locations},
			{Name: "Quit Game", Callback: QuitGame},
	}
	CompanyMenu = Menu{
		Name:    "Company Menu",
		Intro:   CompanyMenuIntro,
		Options: CompanyMenuOptions,
		Back:    QuitGame,
	}
}
