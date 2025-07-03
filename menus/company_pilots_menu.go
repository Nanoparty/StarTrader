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
	fmt.Printf("\r%-20s | %-10s | %-10s | %-10s | %-20s\n", "Name", "Transport", "Combat", "Mining", "Assigned Ship")
	fmt.Println("\r----------------------------------------------------------------------------")
}

func PilotSelected(pilot Pilot) func() {
	return func() {
		selectedDetailPilot = &pilot
		ShowPilotDetailMenu()
	}
}

func CompanyPilotsBack() {
	CurrentMenu = &CompanyMenu
}

func BuildCompanyPilotsMenuOptions() {
	CompanyPilotOptions = []MenuItem{}
	for _, pilot := range CompanyPilots {
		pilotCopy := pilot // avoid closure capture bug
		shipName := "None"
		if pilotCopy.AssignedShip != nil {
			shipName = pilotCopy.AssignedShip.Name
		}
		menuName := fmt.Sprintf("%-20s | %-10d | %-10d | %-10d | %-20s", pilotCopy.Name, pilotCopy.TransportSkill, pilotCopy.CombatSkill, pilotCopy.MiningSkill, shipName)
		CompanyPilotOptions = append(CompanyPilotOptions, MenuItem{
			Name:     menuName,
			Callback: PilotSelected(pilotCopy),
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
