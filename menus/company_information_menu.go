package menus

import (
	"fmt"
	"startrader/globals"
)

var CompanyMissionsCompleted int = 0

var CompanyInformationMenu Menu

func CompanyInformationIntro(m *Menu) {
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Println("\rCompany Information:")
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Printf("\rCompany Name: %s\n", globals.CompanyName)
	fmt.Printf("\rTotal Money: $%d\n", CompanyMoney)
	fmt.Printf("\rTotal Missions Completed: %d\n", CompanyMissionsCompleted)
	fmt.Printf("\rNumber of Ships Owned: %d\n", len(CompanyShips))
	fmt.Printf("\rNumber of Pilots Owned: %d\n", len(CompanyPilots))
	fmt.Println("\r----------------------------------------------------------------------------")
}

func CompanyInformationBack() {
	CurrentMenu = &CompanyMenu
}

func init() {
	CompanyInformationMenu = Menu{
		Name:    "Company Information",
		Intro:   CompanyInformationIntro,
		Options: []MenuItem{{Name: "Back", Callback: CompanyInformationBack}},
		Back:    CompanyInformationBack,
	}
}
