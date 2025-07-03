package menus

import "fmt"

var selectedDetailShip *Ship

func ShipDetailMenuIntro(m *Menu) {
	if selectedDetailShip == nil {
		fmt.Println("\rNo ship selected.")
		return
	}
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Println("\rShip Details:")
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Printf("\rShip: %s\n", selectedDetailShip.Name)
	fmt.Printf("\rType: %s\n", selectedDetailShip.Type)
	fmt.Printf("\rPrice: $%d\n", selectedDetailShip.Price)
	fmt.Printf("\rStorage: %d\n", selectedDetailShip.Storage)
	fmt.Printf("\rSpeed: %d\n", selectedDetailShip.Speed)
	fmt.Printf("\rHealth: %d/%d\n", selectedDetailShip.CurrentHealth, selectedDetailShip.MaxHealth)
	fmt.Printf("\rDamage: %d\n", selectedDetailShip.Damage)
	if selectedDetailShip.AssignedPilot != nil {
		fmt.Printf("\rAssigned Pilot: %s\n", selectedDetailShip.AssignedPilot.Name)
	} else {
		fmt.Println("\rAssigned Pilot: None")
	}
	if selectedDetailShip.AssignedMission != nil {
		fmt.Printf("\rAssigned Mission: %s (%s)\n", selectedDetailShip.AssignedMission.ShortName, selectedDetailShip.AssignedMission.Type)
	} else {
		fmt.Println("\rAssigned Mission: None")
	}
	fmt.Printf("\rStatus: %s\n", selectedDetailShip.Status)
	fmt.Println("\r----------------------------------------------------------------------------")
}

func ShipDetailBack() {
	BuildCompanyShipsMenuOptions()
	CurrentMenu = &CompanyShipsMenu
}

func BuildShipDetailMenuOptions() []MenuItem {
	options := []MenuItem{}
	if selectedDetailShip != nil {
		if selectedDetailShip.AssignedPilot == nil {
			options = append(options, MenuItem{Name: "Assign Pilot", Callback: ShowUnassignedPilotsMenu})
		} else {
			options = append(options, MenuItem{Name: "Unassign Pilot", Callback: UnassignPilotFromShipInShipDetail})
		}
	}
	options = append(options, MenuItem{Name: "Back", Callback: ShipDetailBack})
	return options
}

func UnassignPilotFromShipInShipDetail() {
	if selectedDetailShip != nil && selectedDetailShip.AssignedPilot != nil {
		pilot := selectedDetailShip.AssignedPilot
		selectedDetailShip.AssignedPilot = nil
		for i := range CompanyPilots {
			if &CompanyPilots[i] == pilot {
				CompanyPilots[i].AssignedShip = nil
				break
			}
		}
	}
	ShowShipDetailMenu()
}

var ShipDetailMenu Menu

func ShowShipDetailMenu() {
	ShipDetailMenu.Options = BuildShipDetailMenuOptions()
	CurrentMenu = &ShipDetailMenu
}

func init() {
	ShipDetailMenu = Menu{
		Name:    "Ship Detail",
		Intro:   ShipDetailMenuIntro,
		Options: nil, // Set dynamically
		Back:    ShipDetailBack,
	}
}

