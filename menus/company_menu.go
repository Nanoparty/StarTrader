package menus

import (
	"fmt"
	"startrader/globals"
	"startrader/types"
)


func Ships() {
	BuildCompanyShipsMenuOptions()
	globals.CurrentMenu = &CompanyShipsMenu
}

func Pilots() {
	BuildCompanyPilotsMenuOptions()
	globals.CurrentMenu = &CompanyPilotsMenu
}

func Information() {
	fmt.Println("Information\n\r")
}


func Locations() {
	globals.CurrentMenu = &SectorsMenu
}

func QuitGame() {
	globals.CurrentMenu = &QuitMenu
}

func CompanyMenuIntro(m *types.Menu) {
	header := "Company Menu: " + globals.Company.Name
	moneyHeader := fmt.Sprintf("Credits: $%d", globals.Company.Money)
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Printf("\r%s%*s%s\n\r", header, 76 - len(header) - len(moneyHeader), "", moneyHeader)
	fmt.Println("\r----------------------------------------------------------------------------")
}

var CompanyMenuOptions []types.MenuItem
var CompanyMenu types.Menu

func init() {
	CompanyMenuOptions = []types.MenuItem{
		{Name: "Ships", Callback: Ships},
		{Name: "Pilots", Callback: Pilots},
		{Name: "Active Contracts", Callback: func() { BuildActiveContractsMenuOptions(); globals.CurrentMenu = &ActiveContractsMenu }},
		{Name: "Information", Callback: func() { globals.CurrentMenu = &CompanyInformationMenu }},

		{Name: "Locations", Callback: Locations},
			{Name: "Quit Game", Callback: QuitGame},
	}
	CompanyMenu = types.Menu{
		Name:    "Company Menu",
		Intro:   CompanyMenuIntro,
		Options: CompanyMenuOptions,
		Back:    func() {},
	}
}
