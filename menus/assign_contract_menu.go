package menus

import (
	"fmt"
	"startrader/types"
	"startrader/globals"
)

var AssignContractShipMenu types.Menu
var assignContractShipOptions []types.MenuItem
var selectedAssignContractShip *types.Ship

// Step 1: Show ships eligible for assignment
func BuildAssignContractShipOptions() []types.MenuItem {
	assignContractShipOptions = []types.MenuItem{}
	for i := range globals.Company.Ships {
		ship := &globals.Company.Ships[i]
		if ship.AssignedContract == nil && ship.AssignedPilot != nil {
			shipCopy := ship // capture pointer
			label := fmt.Sprintf("%s (Pilot: %s)", shipCopy.Name, shipCopy.AssignedPilot.Name)
			assignContractShipOptions = append(assignContractShipOptions, types.MenuItem{
				Name:     label,
				Callback: func() { selectedAssignContractShip = shipCopy; ShowAssignContractConfirmMenu() },
			})
		}
	}
	assignContractShipOptions = append(assignContractShipOptions, types.MenuItem{Name: "Back", Callback: BackToContractDetailMenu})
	return assignContractShipOptions
}

func AssignContractShipMenuIntro(m *types.Menu) {
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Println("\rSelect a ship to assign this contract:")
	fmt.Println("\r----------------------------------------------------------------------------")
}

func ShowAssignContractShipMenu() {
	AssignContractShipMenu.Options = BuildAssignContractShipOptions()
	globals.CurrentMenu = &AssignContractShipMenu
}

func BackToContractDetailMenu() {
	globals.CurrentMenu = &StationContractDetailMenu
}

func init() {
	AssignContractShipMenu = types.Menu{
		Name:    "Assign Contract - Select Ship",
		Intro:   AssignContractShipMenuIntro,
		Options: nil, // set dynamically
		Back:    BackToContractDetailMenu,
	}
}

// Step 2: Yes/No confirmation prompt
var AssignContractConfirmMenu types.Menu

func AssignContractConfirmMenuIntro(m *types.Menu) {
	if selectedAssignContractShip != nil && selectedStationContract != nil {
		fmt.Printf("\rAssign contract '%s' to ship '%s' (Pilot: %s)?\n", selectedStationContract.ShortName, selectedAssignContractShip.Name, selectedAssignContractShip.AssignedPilot.Name)
	} else {
		fmt.Println("\rNo ship or contract selected.")
	}
}

func AssignContractYes() {
	if selectedAssignContractShip != nil && selectedStationContract != nil {
		selectedAssignContractShip.AssignedContract = selectedStationContract
		selectedAssignContractShip.Status = "In Progress"
		// Always update the pilot in globals.Company.Pilots whose AssignedShip matches this ship
		for i := range globals.Company.Pilots {
			if globals.Company.Pilots[i].AssignedShip == selectedAssignContractShip {
				globals.Company.Pilots[i].AssignedContract = selectedStationContract
				globals.Company.Pilots[i].Status = "In Progress"
			}
		}
		selectedStationContract.Status = "In Progress"
		// Remove contract from the current station's Contracts
		if selectedDetailStation != nil {
			for i, m := range selectedDetailStation.Contracts {
				if m.ShortName == selectedStationContract.ShortName && m.Type == selectedStationContract.Type && fmt.Sprintf("%d min %d sec", m.Minutes, m.Seconds) == fmt.Sprintf("%d min %d sec", selectedStationContract.Minutes, selectedStationContract.Seconds) && m.Payout == selectedStationContract.Payout {
					selectedDetailStation.Contracts = append(selectedDetailStation.Contracts[:i], selectedDetailStation.Contracts[i+1:]...)
					break
				}
			}
		}
	}
	// Go back to station contract menu with updated list
	StationContractMenu.Options = BuildStationContractMenuOptions()
	globals.CurrentMenu = &StationContractMenu
}

func AssignContractNo() {
	ShowAssignContractShipMenu()
}

func AssignContractConfirmMenuOptions() []types.MenuItem {
	return []types.MenuItem{
		{Name: "Yes", Callback: AssignContractYes},
		{Name: "No", Callback: AssignContractNo},
	}
}

func ShowAssignContractConfirmMenu() {
	AssignContractConfirmMenu.Options = AssignContractConfirmMenuOptions()
	globals.CurrentMenu = &AssignContractConfirmMenu
}

func init() {
	AssignContractConfirmMenu = types.Menu{
		Name:    "Confirm Contract Assignment",
		Intro:   AssignContractConfirmMenuIntro,
		Options: nil, // set dynamically
		Back:    AssignContractNo,
	}
}
