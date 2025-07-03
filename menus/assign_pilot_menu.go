package menus

import "fmt"

var unassignedPilotsMenuOptions []MenuItem
var UnassignedPilotsMenu Menu
var selectedAssignPilot *Pilot

func ShowUnassignedPilotsMenu() {
	BuildUnassignedPilotsMenuOptions()
	UnassignedPilotsMenu.Options = unassignedPilotsMenuOptions
	CurrentMenu = &UnassignedPilotsMenu
}

func BuildUnassignedPilotsMenuOptions() {
	unassignedPilotsMenuOptions = []MenuItem{}
	for i := range CompanyPilots {
		if CompanyPilots[i].AssignedShip == nil {
			pilotCopy := CompanyPilots[i] // avoid closure capture bug
			unassignedPilotsMenuOptions = append(unassignedPilotsMenuOptions, MenuItem{
				Name:     pilotCopy.Name,
				Callback: AssignPilotPrompt(pilotCopy),
			})
		}
	}
	unassignedPilotsMenuOptions = append(unassignedPilotsMenuOptions, MenuItem{Name: "Back", Callback: BackToShipDetailMenu})
}

func AssignPilotPrompt(pilot Pilot) func() {
	return func() {
		selectedAssignPilot = &pilot
		CurrentMenu = &AssignPilotConfirmMenu
	}
}

func BackToShipDetailMenu() {
	ShipDetailMenu.Options = BuildShipDetailMenuOptions()
	CurrentMenu = &ShipDetailMenu
}

func AssignPilotConfirmIntro(m *Menu) {
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
	CurrentMenu = &ShipDetailMenu
}

func AssignPilotNo() {
	ShowUnassignedPilotsMenu()
}

func AssignPilotConfirmMenuOptions() []MenuItem {
	return []MenuItem{
		{Name: "Yes", Callback: AssignPilotYes},
		{Name: "No", Callback: AssignPilotNo},
	}
}

var AssignPilotConfirmMenu Menu

func init() {
	UnassignedPilotsMenu = Menu{
		Name:    "Unassigned Pilots",
		Intro:   func(m *Menu) { fmt.Println("\rSelect a pilot to assign:") },
		Options: nil, // set dynamically
		Back:    BackToShipDetailMenu,
	}

	AssignPilotConfirmMenu = Menu{
		Name:    "Assign Pilot?",
		Intro:   AssignPilotConfirmIntro,
		Options: AssignPilotConfirmMenuOptions(),
		Back:    AssignPilotNo,
	}
}
