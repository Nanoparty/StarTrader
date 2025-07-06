package menus

import (
	"fmt"
	"startrader/globals"
	"startrader/types"
)

func StationInformationMenuIntro(m *types.Menu) {
	if selectedDetailStation == nil {
		fmt.Println("\rNo station selected.")
		return
	}
	moneyHeader := fmt.Sprintf("Credits: $%d", globals.Company.Money)
	fmt.Println("\r----------------------------------------------------------------------------")
	header := "Station Details:"
	fmt.Printf("\r%s%*s%s\n", header, 76-len(header)-len(moneyHeader), "", moneyHeader)
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Printf("\rStation: %s\n", selectedDetailStation.Name)
	fmt.Printf("\rRelationship Level: %d\n", selectedDetailStation.RelationshipLevel)
	fmt.Printf("\rContracts Completed: %d\n", selectedDetailStation.ContractsCompleted)
	fmt.Printf("\rMoney Spent: $%d\n", selectedDetailStation.MoneySpent)
	fmt.Printf("\rExperience: %.1f / %.1f\n", selectedDetailStation.Experience, selectedDetailStation.ExpToNextLevel)
	fmt.Println("\r----------------------------------------------------------------------------")
}

func StationInformationBack() {
	globals.CurrentMenu = &StationDetailMenu
}

var StationInformationMenuOptions = []types.MenuItem{
	{Name: "Back", Callback: StationInformationBack},
}

var StationInformationMenu types.Menu

func init() {
	StationInformationMenu = types.Menu{
		Name:    "Station Information",
		Intro:   StationInformationMenuIntro,
		Options: StationInformationMenuOptions,
		Back:    StationInformationBack,
	}
}
