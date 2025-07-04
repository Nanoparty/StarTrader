package menus

import (
	"fmt"
	"startrader/globals"
)

var CompanyPilots []Pilot
var CompanyPilotOptions []MenuItem
var CompanyPilotsMenu Menu

func CompanyPilotsIntro(m *Menu) {
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Println("\r" + globals.CompanyName + " Pilots:")
	fmt.Println("\r----------------------------------------------------------------------------")
	if len(CompanyPilots) == 0 {
		fmt.Println("\rThere are currently no pilots.\n\r")
		return
	}
	fmt.Printf("\r%-20s | %-10s | %-20s | %-12s\n", "Name", "Skills (T/C/M)", "Assigned Ship", "Status")
	fmt.Println("\r----------------------------------------------------------------------------")
}

func PilotSelected(pilot *Pilot) func() {
	return func() {
		selectedDetailPilot = pilot
		ShowPilotDetailMenu()
	}
}

func CompanyPilotsBack() {
	CurrentMenu = &CompanyMenu
}

func BuildCompanyPilotsMenuOptions() {
	CompanyPilotOptions = []MenuItem{}
	for i := range CompanyPilots {
		pilotPtr := &CompanyPilots[i]
		shipName := "None"
		if pilotPtr.AssignedShip != nil {
			shipName = pilotPtr.AssignedShip.Name
		}
		skills := fmt.Sprintf("%d/%d/%d", pilotPtr.TransportSkill, pilotPtr.CombatSkill, pilotPtr.MiningSkill)
		menuName := fmt.Sprintf("%-20s | %-14s | %-20s | %-12s", pilotPtr.Name, skills, shipName, pilotPtr.Status)
		CompanyPilotOptions = append(CompanyPilotOptions, MenuItem{
			Name:     menuName,
			Callback: PilotSelected(pilotPtr),
		})
	}
	CompanyPilotOptions = append(CompanyPilotOptions, MenuItem{Name: "Back", Callback: CompanyPilotsBack})
	CompanyPilotsMenu.Options = CompanyPilotOptions
}

func init() {
	CompanyPilotsMenu = Menu{
		Name:    "Company Pilots",
		Intro:   CompanyPilotsIntro,
		Options: nil, // Will be set by BuildCompanyPilotsMenuOptions
		Back:    CompanyPilotsBack,
	}
}
