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
	for i := range CompanyShips {
		if CompanyShips[i].AssignedPilot == nil {
			shipCopy := CompanyShips[i] // avoid closure capture bug
			unassignedShipsMenuOptions = append(unassignedShipsMenuOptions, types.MenuItem{
				Name:     shipCopy.Name,
				Callback: AssignShipPrompt(shipCopy),
			})
		}
	}
	unassignedShipsMenuOptions = append(unassignedShipsMenuOptions, types.MenuItem{Name: "Back", Callback: BackToPilotDetailMenu})
}

func AssignShipPrompt(ship types.Ship) func() {
	return func() {
		selectedAssignShip = &ship
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
		for i := range CompanyShips {
			if CompanyShips[i].Name == selectedAssignShip.Name {
				CompanyShips[i].AssignedPilot = selectedDetailPilot
				selectedAssignShip = &CompanyShips[i] // update pointer
				break
			}
		}
		for i := range CompanyPilots {
			if CompanyPilots[i].Name == selectedDetailPilot.Name {
				CompanyPilots[i].AssignedShip = selectedAssignShip
				selectedDetailPilot = &CompanyPilots[i] // update pointer
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
