package menus

import (
	"fmt"
	"startrader/globals"
	"startrader/types"
)

var CompanyMissionsCompleted int = 0

var CompanyInformationMenu types.Menu

func CompanyInformationIntro(m *types.Menu) {
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Println("\rCompany Information:")
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Printf("\rCompany Name: %s\n", globals.Company.Name)
	fmt.Printf("\rTotal Money: $%d\n", globals.Company.Money)
	fmt.Printf("\rTotal Missions Completed: %d\n", CompanyMissionsCompleted)
	fmt.Printf("\rNumber of Ships Owned: %d\n", len(globals.Company.Ships))
	fmt.Printf("\rNumber of Pilots Owned: %d\n", len(globals.Company.Pilots))
	fmt.Println("\r----------------------------------------------------------------------------")
}

func CompanyInformationBack() {
	globals.CurrentMenu = &CompanyMenu
}

func init() {
	CompanyInformationMenu = types.Menu{
		Name:    "Company Information",
		Intro:   CompanyInformationIntro,
		Options: []types.MenuItem{{Name: "Back", Callback: CompanyInformationBack}},
		Back:    CompanyInformationBack,
	}
}
