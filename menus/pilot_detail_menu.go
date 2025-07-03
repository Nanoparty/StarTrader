package menus

import "fmt"

var selectedDetailPilot *Pilot

func PilotDetailMenuIntro(m *Menu) {
	if selectedDetailPilot == nil {
		fmt.Println("\rNo pilot selected.")
		return
	}
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Printf("\rPilot: %s\n", selectedDetailPilot.Name)
	fmt.Printf("\rPrice: $%d\n", selectedDetailPilot.Price)
	fmt.Printf("\rTransport Skill: %d\n", selectedDetailPilot.TransportSkill)
	fmt.Printf("\rCombat Skill: %d\n", selectedDetailPilot.CombatSkill)
	fmt.Printf("\rMining Skill: %d\n", selectedDetailPilot.MiningSkill)
	if selectedDetailPilot.AssignedShip != nil {
		fmt.Printf("\rAssigned Ship: %s\n", selectedDetailPilot.AssignedShip.Name)
	} else {
		fmt.Println("\rAssigned Ship: None")
	}
	fmt.Println("\r----------------------------------------------------------------------------")
}

func PilotDetailBack() {
	BuildCompanyPilotsMenuOptions()
	CurrentMenu = &CompanyPilotsMenu
}

func ShowPilotDetailMenu() {
	PilotDetailMenu.Options = BuildPilotDetailMenuOptions()
	CurrentMenu = &PilotDetailMenu
}

func BuildPilotDetailMenuOptions() []MenuItem {
	options := []MenuItem{}
	if selectedDetailPilot != nil {
		if selectedDetailPilot.AssignedShip == nil {
			options = append(options, MenuItem{Name: "Assign to Ship", Callback: ShowUnassignedShipsMenu})
		} else {
			options = append(options, MenuItem{Name: "Unassign from Ship", Callback: UnassignPilotFromShip})
		}
	}
	options = append(options, MenuItem{Name: "Back", Callback: PilotDetailBack})
	return options
}

func UnassignPilotFromShip() {
	if selectedDetailPilot != nil && selectedDetailPilot.AssignedShip != nil {
		// Unlink both sides
		ship := selectedDetailPilot.AssignedShip
		selectedDetailPilot.AssignedShip = nil
		for i := range CompanyShips {
			if &CompanyShips[i] == ship {
				CompanyShips[i].AssignedPilot = nil
				break
			}
		}
	}
	ShowPilotDetailMenu()
}

var PilotDetailMenu Menu

func init() {
	PilotDetailMenu = Menu{
		Name:    "Pilot Detail",
		Intro:   PilotDetailMenuIntro,
		Options: nil, // Set dynamically
		Back:    PilotDetailBack,
	}
}
