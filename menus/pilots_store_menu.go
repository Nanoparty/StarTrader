package menus

import (
	"fmt"
)

func ShowStationPilotsStoreMenu() {
	if selectedDetailStation == nil {
		fmt.Println("\rNo station selected.")
		return
	}
	BuildStationPilotsStoreMenuOptions()
	CurrentMenu = &PilotsStoreMenu
}

func BuildStationPilotsStoreMenuOptions() {
	PilotsStoreMenuOptions = []MenuItem{}
	for i, pilot := range selectedDetailStation.PilotsForSale {
		pilotCopy := pilot // avoid closure capture bug
		menuName := fmt.Sprintf("%-20s | $%-7d | %-8d | %-8d | %-8d", pilotCopy.Name, pilotCopy.Price, pilotCopy.TransportSkill, pilotCopy.CombatSkill, pilotCopy.MiningSkill)
		PilotsStoreMenuOptions = append(PilotsStoreMenuOptions, MenuItem{
			Name:     menuName,
			Callback: PilotPurchasePrompt(pilotCopy, i),
		})
	}
	PilotsStoreMenuOptions = append(PilotsStoreMenuOptions, MenuItem{Name: "Back", Callback: PilotsStoreBack})
	PilotsStoreMenu.Options = PilotsStoreMenuOptions
}

var PilotsStoreMenuOptions []MenuItem
var PilotsStoreMenu Menu
var selectedPilot *Pilot

func init() {
	PilotsStoreMenu = Menu{
		Name:  "Pilots Store Menu",
		Intro: PilotsStoreMenuIntro,
	}

	PilotPurchaseMenu = Menu{
		Name:    "Hire Pilot?",
		Intro:   PilotPurchaseMenuIntro,
		Options: PilotPurchaseMenuOptions(),
		Back:    PilotPurchaseNo,
	}
}

func PilotsStoreMenuIntro(m *Menu) {
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Println("\rPilots for Hire:")
	fmt.Println("\r----------------------------------------------------------------------------")
	if selectedDetailStation == nil || len(selectedDetailStation.PilotsForSale) == 0 {
		fmt.Println("\rThere are no more pilots available for hire at this point in time.")
		fmt.Println("\r----------------------------------------------------------------------------")
		return
	}
	fmt.Printf("\r%-20s | %-8s | %-8s | %-8s | %-8s\n", "Name", "Price", "Transport", "Combat", "Mining")
	fmt.Println("\r----------------------------------------------------------------------------")
}

func PilotPurchasePrompt(pilot Pilot, i int) func() {
	return func() {
		selectedPilot = &selectedDetailStation.PilotsForSale[i]
		CurrentMenu = &PilotPurchaseMenu
	}
}

func PilotPurchaseYes() {
	if selectedPilot != nil {
		if CompanyMoney < selectedPilot.Price {
			ShowWarningMenu("Insufficient funds to hire this pilot.", &PilotPurchaseMenu)
			return
		}
		CompanyMoney -= selectedPilot.Price
		pilotCopy := *selectedPilot
		CompanyPilots = append(CompanyPilots, pilotCopy)
		// Remove the purchased pilot from PilotsForSale
		for i, p := range selectedDetailStation.PilotsForSale {
			if p.Name == pilotCopy.Name && p.Price == pilotCopy.Price {
				selectedDetailStation.PilotsForSale = append(selectedDetailStation.PilotsForSale[:i], selectedDetailStation.PilotsForSale[i+1:]...)
				break
			}
		}
	}
	selectedPilot = nil
	BuildPilotsStoreMenuOptions()
	CurrentMenu = &PilotsStoreMenu
}

func PilotPurchaseNo() {
	selectedPilot = nil
	CurrentMenu = &PilotsStoreMenu
}

func PilotPurchaseMenuIntro(m *Menu) {
	if selectedPilot != nil {
		fmt.Printf("\rHire %s for $%d?\n", selectedPilot.Name, selectedPilot.Price)
	} else {
		fmt.Println("\rNo pilot selected.")
	}
}

func PilotPurchaseMenuOptions() []MenuItem {
	return []MenuItem{
		{Name: "Yes", Callback: PilotPurchaseYes},
		{Name: "No", Callback: PilotPurchaseNo},
	}
}

var PilotPurchaseMenu Menu

func BuildPilotsStoreMenuOptions() {
	PilotsStoreMenuOptions = []MenuItem{}
	for i, pilot := range selectedDetailStation.PilotsForSale {
		pilotCopy := pilot // avoid closure capture bug
		menuName := fmt.Sprintf("%-20s | $%-7d | %-8d | %-8d | %-8d", pilotCopy.Name, pilotCopy.Price, pilotCopy.TransportSkill, pilotCopy.CombatSkill, pilotCopy.MiningSkill)
		PilotsStoreMenuOptions = append(PilotsStoreMenuOptions, MenuItem{
			Name:     menuName,
			Callback: PilotPurchasePrompt(pilotCopy, i),
		})
	}
	PilotsStoreMenuOptions = append(PilotsStoreMenuOptions, MenuItem{Name: "Back", Callback: PilotsStoreBack})
	PilotsStoreMenu.Options = PilotsStoreMenuOptions
}

func PilotsStoreBack() {
	CurrentMenu = &StationDetailMenu
}

func init() {
	PilotsStoreMenuOptions = []MenuItem{}
	PilotsStoreMenu = Menu{
		Name:    "Pilots Store Menu",
		Intro:   PilotsStoreMenuIntro,
		Options: PilotsStoreMenuOptions, // will be set dynamically
		Back:    PilotsStoreBack,
	}
}
