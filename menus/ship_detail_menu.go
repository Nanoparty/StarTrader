package menus

import (
	"fmt"
	"startrader/globals"
	"startrader/types"
)

var selectedDetailShip *types.Ship

func ShipDetailMenuIntro(m *types.Menu) {
	if selectedDetailShip == nil {
		fmt.Println("\rNo ship selected.")
		return
	}
	fmt.Println("\r----------------------------------------------------------------------------")
		header := "Ship Details:"
	moneyHeader := fmt.Sprintf("$%d", globals.Company.Money)
	fmt.Printf("\r%s%*s%s\n\r", header, 76 - len(header) - len(moneyHeader), "", moneyHeader)
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
	if selectedDetailShip.AssignedContract != nil {
		fmt.Printf("\rAssigned Contract: %s (%s)\n", selectedDetailShip.AssignedContract.ShortName, selectedDetailShip.AssignedContract.Type)
	} else {
		fmt.Println("\rAssigned Contract: None")
	}
	fmt.Printf("\rStatus: %s\n", selectedDetailShip.Status)
	fmt.Println("\r----------------------------------------------------------------------------")
}

func ShipDetailBack() {
	BuildCompanyShipsMenuOptions()
	globals.CurrentMenu = &CompanyShipsMenu
}

func BuildShipDetailMenuOptions() []types.MenuItem {
	options := []types.MenuItem{}
	if selectedDetailShip != nil {
		if selectedDetailShip.AssignedPilot == nil {
			options = append(options, types.MenuItem{Name: "Assign Pilot", Callback: ShowUnassignedPilotsMenu})
		} else {
			options = append(options, types.MenuItem{Name: "Unassign Pilot", Callback: UnassignPilotFromShipInShipDetail}) // Prevents unassign if ship or pilot has a contract
		}
	}
	options = append(options, types.MenuItem{Name: "Back", Callback: ShipDetailBack})
	return options
}

func UnassignPilotFromShipInShipDetail() {
	if selectedDetailShip != nil && selectedDetailShip.AssignedPilot != nil {
		// Prevent unassign if ship or pilot has an assigned contract
		if selectedDetailShip.AssignedContract != nil {
			ShowWarningMenu("Cannot unassign: Ship is on a contract.", &ShipDetailMenu)
			return
		}
		if selectedDetailShip.AssignedPilot.AssignedContract != nil {
			ShowWarningMenu("Cannot unassign: Pilot is on a contract.", &ShipDetailMenu)
			return
		}
		pilot := selectedDetailShip.AssignedPilot
		selectedDetailShip.AssignedPilot = nil
		for i := range globals.Company.Pilots {
			if &globals.Company.Pilots[i] == pilot {
				globals.Company.Pilots[i].AssignedShip = nil
				break
			}
		}
	}
	ShowShipDetailMenu()
}

var ShipDetailMenu types.Menu

func ShowShipDetailMenu() {
	ShipDetailMenu.Options = BuildShipDetailMenuOptions()
	globals.CurrentMenu = &ShipDetailMenu
}

func init() {
	ShipDetailMenu = types.Menu{
		Name:    "Ship Detail",
		Intro:   ShipDetailMenuIntro,
		Options: nil, // Set dynamically
		Back:    ShipDetailBack,
	}
}

