package menus

import (
	"fmt"
	"startrader/globals"
	"startrader/types"
)

var CompanyPilotOptions []types.MenuItem
var CompanyPilotsMenu types.Menu

func CompanyPilotsIntro(m *types.Menu) {
	header := globals.Company.Name + " Pilots:"
	moneyHeader := fmt.Sprintf("Credits: $%d", globals.Company.Money)
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Printf("\r%s%*s%s\n\r", header, 76 - len(header) - len(moneyHeader), "", moneyHeader)
	fmt.Println("\r----------------------------------------------------------------------------")
	if len(globals.Company.Pilots) == 0 {
		fmt.Println("\rThere are currently no pilots.\n\r")
		return
	}
	fmt.Printf("\r%-20s | %-10s | %-20s | %-12s\n", "Name", "Skills (T/C/M)", "Assigned Ship", "Status")
	fmt.Println("\r----------------------------------------------------------------------------")
}

func PilotSelected(pilot *types.Pilot) func() {
	return func() {
		selectedDetailPilot = pilot
		ShowPilotDetailMenu()
	}
}

func CompanyPilotsBack() {
	globals.CurrentMenu = &CompanyMenu
}

func BuildCompanyPilotsMenuOptions() {
	CompanyPilotOptions = []types.MenuItem{}
	for i := range globals.Company.Pilots {
		pilotPtr := &globals.Company.Pilots[i]
		shipName := "None"
		if pilotPtr.AssignedShip != nil {
			shipName = pilotPtr.AssignedShip.Name
		}
		skills := fmt.Sprintf("%d/%d/%d", pilotPtr.TransportSkill, pilotPtr.CombatSkill, pilotPtr.MiningSkill)
		menuName := fmt.Sprintf("%-20s | %-14s | %-20s | %-12s", pilotPtr.Name, skills, shipName, pilotPtr.Status)
		CompanyPilotOptions = append(CompanyPilotOptions, types.MenuItem{
			Name:     menuName,
			Callback: PilotSelected(pilotPtr),
		})
	}
	CompanyPilotOptions = append(CompanyPilotOptions, types.MenuItem{Name: "Back", Callback: CompanyPilotsBack})
	CompanyPilotsMenu.Options = CompanyPilotOptions
}

func init() {
	CompanyPilotsMenu = types.Menu{
		Name:    "Company Pilots",
		Intro:   CompanyPilotsIntro,
		Options: nil, // Will be set by BuildCompanyPilotsMenuOptions
		Back:    CompanyPilotsBack,
	}
}
