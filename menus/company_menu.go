package menus

import (
	"fmt"
	"startrader/globals"
)


func Ships() {
	CurrentMenu = &CompanyShipsMenu
}

func Pilots() {
	fmt.Println("Pilots\n\r")
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
		{Name: "Information", Callback: Information},
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
