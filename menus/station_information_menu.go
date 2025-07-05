package menus

import "fmt"

func StationInformationMenuIntro(m *Menu) {
	if selectedDetailStation == nil {
		fmt.Println("\rNo station selected.")
		return
	}
	moneyHeader := fmt.Sprintf("$%d", CompanyMoney)
	fmt.Println("\r----------------------------------------------------------------------------")
	header := "Station Details:"
	fmt.Printf("\r%s%*s%s\n", header, 76-len(header)-len(moneyHeader), "", moneyHeader)
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Printf("\rStation: %s\n", selectedDetailStation.Name)
	fmt.Printf("\rRelationship Level: %d\n", selectedDetailStation.RelationshipLevel)
	fmt.Printf("\rMissions Completed: %d\n", selectedDetailStation.MissionsCompleted)
	fmt.Printf("\rMoney Spent: $%d\n", selectedDetailStation.MoneySpent)
	fmt.Printf("\rExperience: %.1f / %.1f\n", selectedDetailStation.Experience, selectedDetailStation.ExpToNextLevel)
	fmt.Println("\r----------------------------------------------------------------------------")
}

func StationInformationBack() {
	CurrentMenu = &StationDetailMenu
}

var StationInformationMenuOptions = []MenuItem{
	{Name: "Back", Callback: StationInformationBack},
}

var StationInformationMenu Menu

func init() {
	StationInformationMenu = Menu{
		Name:    "Station Information",
		Intro:   StationInformationMenuIntro,
		Options: StationInformationMenuOptions,
		Back:    StationInformationBack,
	}
}
