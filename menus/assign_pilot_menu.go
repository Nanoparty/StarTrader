package menus

import (
	"fmt"
	"startrader/types"
	"startrader/globals"
)

var unassignedPilotsMenuOptions []types.MenuItem
var UnassignedPilotsMenu types.Menu
var selectedAssignPilot *types.Pilot

func ShowUnassignedPilotsMenu() {
	BuildUnassignedPilotsMenuOptions()
	UnassignedPilotsMenu.Options = unassignedPilotsMenuOptions
	globals.CurrentMenu = &UnassignedPilotsMenu
}

func BuildUnassignedPilotsMenuOptions() {
	unassignedPilotsMenuOptions = []types.MenuItem{}
	for i := range CompanyPilots {
		if CompanyPilots[i].AssignedShip == nil {
			pilotCopy := CompanyPilots[i] // avoid closure capture bug
			unassignedPilotsMenuOptions = append(unassignedPilotsMenuOptions, types.MenuItem{
				Name:     pilotCopy.Name,
				Callback: AssignPilotPrompt(pilotCopy),
			})
		}
	}
	unassignedPilotsMenuOptions = append(unassignedPilotsMenuOptions, types.MenuItem{Name: "Back", Callback: BackToShipDetailMenu})
}

func AssignPilotPrompt(pilot types.Pilot) func() {
	return func() {
		selectedAssignPilot = &pilot
		globals.CurrentMenu = &AssignPilotConfirmMenu
	}
}

func BackToShipDetailMenu() {
	ShipDetailMenu.Options = BuildShipDetailMenuOptions()
	globals.CurrentMenu = &ShipDetailMenu
}

func AssignPilotConfirmIntro(m *types.Menu) {
	if selectedDetailShip != nil && selectedAssignPilot != nil {
		fmt.Printf("\rAssign %s to %s?\n", selectedAssignPilot.Name, selectedDetailShip.Name)
	} else {
		fmt.Println("\rNo pilot or ship selected.")
	}
}

func AssignPilotYes() {
	if selectedDetailShip != nil && selectedAssignPilot != nil {
		// Update both objects to point to each other
		for i := range CompanyPilots {
			if CompanyPilots[i].Name == selectedAssignPilot.Name {
				CompanyPilots[i].AssignedShip = selectedDetailShip
				selectedAssignPilot = &CompanyPilots[i] // update pointer
				break
			}
		}
		for i := range CompanyShips {
			if CompanyShips[i].Name == selectedDetailShip.Name {
				CompanyShips[i].AssignedPilot = selectedAssignPilot
				selectedDetailShip = &CompanyShips[i] // update pointer
				break
			}
		}
	}
	ShipDetailMenu.Options = BuildShipDetailMenuOptions()
	globals.CurrentMenu = &ShipDetailMenu
}

func AssignPilotNo() {
	ShowUnassignedPilotsMenu()
}

func AssignPilotConfirmMenuOptions() []types.MenuItem {
	return []types.MenuItem{
		{Name: "Yes", Callback: AssignPilotYes},
		{Name: "No", Callback: AssignPilotNo},
	}
}

var AssignPilotConfirmMenu types.Menu

func init() {
	UnassignedPilotsMenu = types.Menu{
		Name:    "Unassigned Pilots",
		Intro:   func(m *types.Menu) { fmt.Println("\rSelect a pilot to assign:") },
		Options: nil, // set dynamically
		Back:    BackToShipDetailMenu,
	}

	AssignPilotConfirmMenu = types.Menu{
		Name:    "Assign Pilot?",
		Intro:   AssignPilotConfirmIntro,
		Options: AssignPilotConfirmMenuOptions(),
		Back:    AssignPilotNo,
	}
}
