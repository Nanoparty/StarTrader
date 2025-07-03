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

func Store() {
	CurrentMenu = &StoreMenu
}

func Locations() {
	CurrentMenu = &SectorsMenu
}

func QuitGame() {
	CurrentMenu = &QuitMenu
}

func CompanyMenuIntro(m *Menu) {
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Println("\rCompany Menu: ", globals.CompanyName)
	fmt.Println("\r----------------------------------------------------------------------------")
}

var CompanyMenuOptions []MenuItem
var CompanyMenu Menu

func init() {
	CompanyMenuOptions = []MenuItem{
		{Name: "Ships", Callback: Ships},
		{Name: "Pilots", Callback: Pilots},
		{Name: "Information", Callback: func() { CurrentMenu = &CompanyInformationMenu }},

		{Name: "Locations", Callback: Locations},
		{Name: "Store", Callback: Store},
		{Name: "Quit Game", Callback: QuitGame},
	}
	CompanyMenu = Menu{
		Name:    "Company Menu",
		Intro:   CompanyMenuIntro,
		Options: CompanyMenuOptions,
		Back:    QuitGame,
	}
}
