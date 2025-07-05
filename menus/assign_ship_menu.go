package menus

import (
	"fmt"
	"startrader/types"
	"startrader/globals"
)

var unassignedShipsMenuOptions []types.MenuItem
var UnassignedShipsMenu types.Menu
var selectedAssignShip *types.Ship

func ShowUnassignedShipsMenu() {
	BuildUnassignedShipsMenuOptions()
	UnassignedShipsMenu.Options = unassignedShipsMenuOptions
	globals.CurrentMenu = &UnassignedShipsMenu
}

func BuildUnassignedShipsMenuOptions() {
	unassignedShipsMenuOptions = []types.MenuItem{}
	for i := range globals.Company.Ships {
		if globals.Company.Ships[i].AssignedPilot == nil {
			shipPtr := &globals.Company.Ships[i]
			unassignedShipsMenuOptions = append(unassignedShipsMenuOptions, types.MenuItem{
				Name:     shipPtr.Name,
				Callback: AssignShipPrompt(shipPtr),
			})
		}
	}
	unassignedShipsMenuOptions = append(unassignedShipsMenuOptions, types.MenuItem{Name: "Back", Callback: BackToPilotDetailMenu})
}

func AssignShipPrompt(ship *types.Ship) func() {
	return func() {
		selectedAssignShip = ship
		globals.CurrentMenu = &AssignShipConfirmMenu
	}
}

func BackToPilotDetailMenu() {
	PilotDetailMenu.Options = BuildPilotDetailMenuOptions()
	globals.CurrentMenu = &PilotDetailMenu
}

func AssignShipConfirmIntro(m *types.Menu) {
	if selectedDetailPilot != nil && selectedAssignShip != nil {
		fmt.Printf("\rAssign %s to %s?\n", selectedDetailPilot.Name, selectedAssignShip.Name)
	} else {
		fmt.Println("\rNo pilot or ship selected.")
	}
}

func AssignShipYes() {
	if selectedDetailPilot != nil && selectedAssignShip != nil {
		// Update both objects to point to each other
		for i := range globals.Company.Ships {
			if globals.Company.Ships[i].Name == selectedAssignShip.Name {
				globals.Company.Ships[i].AssignedPilot = selectedDetailPilot
				selectedAssignShip = &globals.Company.Ships[i] // update pointer
				break
			}
		}
		for i := range globals.Company.Pilots {
			if globals.Company.Pilots[i].Name == selectedDetailPilot.Name {
				globals.Company.Pilots[i].AssignedShip = selectedAssignShip
				selectedDetailPilot = &globals.Company.Pilots[i] // update pointer
				break
			}
		}
	}
	PilotDetailMenu.Options = BuildPilotDetailMenuOptions()
	globals.CurrentMenu = &PilotDetailMenu
}

func AssignShipNo() {
	ShowUnassignedShipsMenu()
}

func AssignShipConfirmMenuOptions() []types.MenuItem {
	return []types.MenuItem{
		{Name: "Yes", Callback: AssignShipYes},
		{Name: "No", Callback: AssignShipNo},
	}
}

var AssignShipConfirmMenu types.Menu

func init() {
	UnassignedShipsMenu = types.Menu{
		Name:    "Unassigned Ships",
		Intro:   func(m *types.Menu) { fmt.Println("\rSelect a ship to assign:") },
		Options: nil, // set dynamically
		Back:    BackToPilotDetailMenu,
	}

	AssignShipConfirmMenu = types.Menu{
		Name:    "Assign Ship?",
		Intro:   AssignShipConfirmIntro,
		Options: AssignShipConfirmMenuOptions(),
		Back:    AssignShipNo,
	}
}
