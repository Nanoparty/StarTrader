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
	for i := range globals.Company.Pilots {
		if globals.Company.Pilots[i].AssignedShip == nil {
			pilotPtr := &globals.Company.Pilots[i]
			unassignedPilotsMenuOptions = append(unassignedPilotsMenuOptions, types.MenuItem{
				Name:     pilotPtr.Name,
				Callback: AssignPilotPrompt(pilotPtr),
			})
		}
	}
	unassignedPilotsMenuOptions = append(unassignedPilotsMenuOptions, types.MenuItem{Name: "Back", Callback: BackToShipDetailMenu})
}

func AssignPilotPrompt(pilot *types.Pilot) func() {
	return func() {
		selectedAssignPilot = pilot
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
		for i := range globals.Company.Pilots {
			if globals.Company.Pilots[i].Name == selectedAssignPilot.Name {
				globals.Company.Pilots[i].AssignedShip = selectedDetailShip
				selectedAssignPilot = &globals.Company.Pilots[i] // update pointer
				break
			}
		}
		for i := range globals.Company.Ships {
			if globals.Company.Ships[i].Name == selectedDetailShip.Name {
				globals.Company.Ships[i].AssignedPilot = selectedAssignPilot
				selectedDetailShip = &globals.Company.Ships[i] // update pointer
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
