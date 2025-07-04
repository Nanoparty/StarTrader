package menus

import "fmt"

var AssignMissionShipMenu Menu
var assignMissionShipOptions []MenuItem
var selectedAssignMissionShip *Ship

// Step 1: Show ships eligible for assignment
func BuildAssignMissionShipOptions() []MenuItem {
	assignMissionShipOptions = []MenuItem{}
	for i := range CompanyShips {
		ship := &CompanyShips[i]
		if ship.AssignedMission == nil && ship.AssignedPilot != nil {
			shipCopy := ship // capture pointer
			label := fmt.Sprintf("%s (Pilot: %s)", shipCopy.Name, shipCopy.AssignedPilot.Name)
			assignMissionShipOptions = append(assignMissionShipOptions, MenuItem{
				Name:     label,
				Callback: func() { selectedAssignMissionShip = shipCopy; ShowAssignMissionConfirmMenu() },
			})
		}
	}
	assignMissionShipOptions = append(assignMissionShipOptions, MenuItem{Name: "Back", Callback: BackToMissionDetailMenu})
	return assignMissionShipOptions
}

func AssignMissionShipMenuIntro(m *Menu) {
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Println("\rSelect a ship to assign this mission:")
	fmt.Println("\r----------------------------------------------------------------------------")
}

func ShowAssignMissionShipMenu() {
	AssignMissionShipMenu.Options = BuildAssignMissionShipOptions()
	CurrentMenu = &AssignMissionShipMenu
}

func BackToMissionDetailMenu() {
	CurrentMenu = &StationMissionDetailMenu
}

func init() {
	AssignMissionShipMenu = Menu{
		Name:    "Assign Mission - Select Ship",
		Intro:   AssignMissionShipMenuIntro,
		Options: nil, // set dynamically
		Back:    BackToMissionDetailMenu,
	}
}

// Step 2: Yes/No confirmation prompt
var AssignMissionConfirmMenu Menu

func AssignMissionConfirmMenuIntro(m *Menu) {
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
		// Always update the pilot in CompanyPilots whose AssignedShip matches this ship
		for i := range CompanyPilots {
			if CompanyPilots[i].AssignedShip == selectedAssignMissionShip {
				CompanyPilots[i].AssignedMission = selectedStationMission
				CompanyPilots[i].Status = "In Progress"
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
	CurrentMenu = &StationMissionMenu
}

func AssignMissionNo() {
	ShowAssignMissionShipMenu()
}

func AssignMissionConfirmMenuOptions() []MenuItem {
	return []MenuItem{
		{Name: "Yes", Callback: AssignMissionYes},
		{Name: "No", Callback: AssignMissionNo},
	}
}

func ShowAssignMissionConfirmMenu() {
	AssignMissionConfirmMenu.Options = AssignMissionConfirmMenuOptions()
	CurrentMenu = &AssignMissionConfirmMenu
}

func init() {
	AssignMissionConfirmMenu = Menu{
		Name:    "Confirm Mission Assignment",
		Intro:   AssignMissionConfirmMenuIntro,
		Options: nil, // set dynamically
		Back:    AssignMissionNo,
	}
}
