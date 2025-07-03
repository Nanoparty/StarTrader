package menus

import "fmt"

var unassignedShipsMenuOptions []MenuItem
var UnassignedShipsMenu Menu
var selectedAssignShip *Ship

func ShowUnassignedShipsMenu() {
	BuildUnassignedShipsMenuOptions()
	UnassignedShipsMenu.Options = unassignedShipsMenuOptions
	CurrentMenu = &UnassignedShipsMenu
}

func BuildUnassignedShipsMenuOptions() {
	unassignedShipsMenuOptions = []MenuItem{}
	for i := range CompanyShips {
		if CompanyShips[i].AssignedPilot == nil {
			shipCopy := CompanyShips[i] // avoid closure capture bug
			unassignedShipsMenuOptions = append(unassignedShipsMenuOptions, MenuItem{
				Name:     shipCopy.Name,
				Callback: AssignShipPrompt(shipCopy),
			})
		}
	}
	unassignedShipsMenuOptions = append(unassignedShipsMenuOptions, MenuItem{Name: "Back", Callback: BackToPilotDetailMenu})
}

func AssignShipPrompt(ship Ship) func() {
	return func() {
		selectedAssignShip = &ship
		CurrentMenu = &AssignShipConfirmMenu
	}
}

func BackToPilotDetailMenu() {
	PilotDetailMenu.Options = BuildPilotDetailMenuOptions()
	CurrentMenu = &PilotDetailMenu
}

func AssignShipConfirmIntro(m *Menu) {
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
	CurrentMenu = &PilotDetailMenu
}

func AssignShipNo() {
	ShowUnassignedShipsMenu()
}

func AssignShipConfirmMenuOptions() []MenuItem {
	return []MenuItem{
		{Name: "Yes", Callback: AssignShipYes},
		{Name: "No", Callback: AssignShipNo},
	}
}

var AssignShipConfirmMenu Menu

func init() {
	UnassignedShipsMenu = Menu{
		Name:    "Unassigned Ships",
		Intro:   func(m *Menu) { fmt.Println("\rSelect a ship to assign:") },
		Options: nil, // set dynamically
		Back:    BackToPilotDetailMenu,
	}

	AssignShipConfirmMenu = Menu{
		Name:    "Assign Ship?",
		Intro:   AssignShipConfirmIntro,
		Options: AssignShipConfirmMenuOptions(),
		Back:    AssignShipNo,
	}
}
