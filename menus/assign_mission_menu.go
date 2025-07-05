package menus

import (
	"fmt"
	"startrader/types"
	"startrader/globals"
)

var AssignMissionShipMenu types.Menu
var assignMissionShipOptions []types.MenuItem
var selectedAssignMissionShip *types.Ship

// Step 1: Show ships eligible for assignment
func BuildAssignMissionShipOptions() []types.MenuItem {
	assignMissionShipOptions = []types.MenuItem{}
	for i := range globals.Company.Ships {
		ship := &globals.Company.Ships[i]
		if ship.AssignedMission == nil && ship.AssignedPilot != nil {
			shipCopy := ship // capture pointer
			label := fmt.Sprintf("%s (Pilot: %s)", shipCopy.Name, shipCopy.AssignedPilot.Name)
			assignMissionShipOptions = append(assignMissionShipOptions, types.MenuItem{
				Name:     label,
				Callback: func() { selectedAssignMissionShip = shipCopy; ShowAssignMissionConfirmMenu() },
			})
		}
	}
	assignMissionShipOptions = append(assignMissionShipOptions, types.MenuItem{Name: "Back", Callback: BackToMissionDetailMenu})
	return assignMissionShipOptions
}

func AssignMissionShipMenuIntro(m *types.Menu) {
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Println("\rSelect a ship to assign this mission:")
	fmt.Println("\r----------------------------------------------------------------------------")
}

func ShowAssignMissionShipMenu() {
	AssignMissionShipMenu.Options = BuildAssignMissionShipOptions()
	globals.CurrentMenu = &AssignMissionShipMenu
}

func BackToMissionDetailMenu() {
	globals.CurrentMenu = &StationMissionDetailMenu
}

func init() {
	AssignMissionShipMenu = types.Menu{
		Name:    "Assign Mission - Select Ship",
		Intro:   AssignMissionShipMenuIntro,
		Options: nil, // set dynamically
		Back:    BackToMissionDetailMenu,
	}
}

// Step 2: Yes/No confirmation prompt
var AssignMissionConfirmMenu types.Menu

func AssignMissionConfirmMenuIntro(m *types.Menu) {
	if selectedAssignMissionShip != nil && selectedStationMission != nil {
		fmt.Printf("\rAssign mission '%s' to ship '%s' (Pilot: %s)?\n", selectedStationMission.ShortName, selectedAssignMissionShip.Name, selectedAssignMissionShip.AssignedPilot.Name)
	} else {
		fmt.Println("\rNo ship or mission selected.")
	}
}

func AssignMissionYes() {
	if selectedAssignMissionShip != nil && selectedStationMission != nil {
		selectedAssignMissionShip.AssignedMission = selectedStationMission
		selectedAssignMissionShip.Status = "In Progress"
		// Always update the pilot in globals.Company.Pilots whose AssignedShip matches this ship
		for i := range globals.Company.Pilots {
			if globals.Company.Pilots[i].AssignedShip == selectedAssignMissionShip {
				globals.Company.Pilots[i].AssignedMission = selectedStationMission
				globals.Company.Pilots[i].Status = "In Progress"
			}
		}
		selectedStationMission.Status = "In Progress"
		// Remove mission from the current station's Missions
		if selectedDetailStation != nil {
			for i, m := range selectedDetailStation.Missions {
				if m.ShortName == selectedStationMission.ShortName && m.Type == selectedStationMission.Type && fmt.Sprintf("%d min %d sec", m.Minutes, m.Seconds) == fmt.Sprintf("%d min %d sec", selectedStationMission.Minutes, selectedStationMission.Seconds) && m.Payout == selectedStationMission.Payout {
					selectedDetailStation.Missions = append(selectedDetailStation.Missions[:i], selectedDetailStation.Missions[i+1:]...)
					break
				}
			}
		}
	}
	// Go back to station mission menu with updated list
	StationMissionMenu.Options = BuildStationMissionMenuOptions()
	globals.CurrentMenu = &StationMissionMenu
}

func AssignMissionNo() {
	ShowAssignMissionShipMenu()
}

func AssignMissionConfirmMenuOptions() []types.MenuItem {
	return []types.MenuItem{
		{Name: "Yes", Callback: AssignMissionYes},
		{Name: "No", Callback: AssignMissionNo},
	}
}

func ShowAssignMissionConfirmMenu() {
	AssignMissionConfirmMenu.Options = AssignMissionConfirmMenuOptions()
	globals.CurrentMenu = &AssignMissionConfirmMenu
}

func init() {
	AssignMissionConfirmMenu = types.Menu{
		Name:    "Confirm Mission Assignment",
		Intro:   AssignMissionConfirmMenuIntro,
		Options: nil, // set dynamically
		Back:    AssignMissionNo,
	}
}
