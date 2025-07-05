package menus

import (
	"fmt"
	"startrader/globals"
	"startrader/types"
)

var selectedDetailPilot *types.Pilot

func PilotDetailMenuIntro(m *types.Menu) {
	if selectedDetailPilot == nil {
		fmt.Println("\rNo pilot selected.")
		return
	}
	fmt.Println("\r----------------------------------------------------------------------------")
	header := "Pilot Details:"
	moneyHeader := fmt.Sprintf("$%d", globals.Company.Money)
	fmt.Printf("\r%s%*s%s\n\r", header, 76 - len(header) - len(moneyHeader), "", moneyHeader)
fmt.Printf("\rPilot: %s\n", selectedDetailPilot.Name)
	fmt.Printf("\rPrice: $%d\n", selectedDetailPilot.Price)
	fmt.Printf("\rSkills (T/C/M): %d/%d/%d\n", selectedDetailPilot.TransportSkill, selectedDetailPilot.CombatSkill, selectedDetailPilot.MiningSkill)
	if selectedDetailPilot.AssignedShip != nil {
		fmt.Printf("\rAssigned Ship: %s\n", selectedDetailPilot.AssignedShip.Name)
	} else {
		fmt.Println("\rAssigned Ship: None")
	}
	if selectedDetailPilot.AssignedMission != nil {
		fmt.Printf("\rAssigned Mission: %s (%s)\n", selectedDetailPilot.AssignedMission.ShortName, selectedDetailPilot.AssignedMission.Type)
	} else {
		fmt.Println("\rAssigned Mission: None")
	}
	fmt.Printf("\rStatus: %s\n", selectedDetailPilot.Status)
	fmt.Println("\r----------------------------------------------------------------------------")
}

func PilotDetailBack() {
	BuildCompanyPilotsMenuOptions()
	globals.CurrentMenu = &CompanyPilotsMenu
}

func ShowPilotDetailMenu() {
	PilotDetailMenu.Options = BuildPilotDetailMenuOptions()
	globals.CurrentMenu = &PilotDetailMenu
}

func BuildPilotDetailMenuOptions() []types.MenuItem {
	options := []types.MenuItem{}
	if selectedDetailPilot != nil {
		if selectedDetailPilot.AssignedShip == nil {
			options = append(options, types.MenuItem{Name: "Assign to Ship", Callback: ShowUnassignedShipsMenu})
		} else {
			options = append(options, types.MenuItem{Name: "Unassign from Ship", Callback: UnassignPilotFromShip})
		}
	}
	options = append(options, types.MenuItem{Name: "Back", Callback: PilotDetailBack})
	return options
}

func UnassignPilotFromShip() {
	if selectedDetailPilot != nil && selectedDetailPilot.AssignedShip != nil {
		// Prevent unassign if pilot or ship has an assigned mission
		if selectedDetailPilot.AssignedMission != nil {
			ShowWarningMenu("Cannot unassign: Pilot is on a mission.", &PilotDetailMenu)
			return
		}
		if selectedDetailPilot.AssignedShip.AssignedMission != nil {
			ShowWarningMenu("Cannot unassign: Ship is on a mission.", &PilotDetailMenu)
			return
		}
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

var PilotDetailMenu types.Menu

func init() {
	PilotDetailMenu = types.Menu{
		Name:    "Pilot Detail",
		Intro:   PilotDetailMenuIntro,
		Options: nil, // Set dynamically
		Back:    PilotDetailBack,
	}
}
