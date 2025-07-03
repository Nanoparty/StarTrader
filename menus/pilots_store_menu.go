package menus

import (
	"fmt"
	"startrader/utils"
)

var PilotsForSale []Pilot
var PilotsStoreMenuOptions []MenuItem
var PilotsStoreMenu Menu
var selectedPilot *Pilot

func init() {
	PilotsForSale = []Pilot{
		{utils.Generate_Pilot_Name(), 50000, 8, 9, 2, nil, nil, "Idle"},
		{utils.Generate_Pilot_Name(), 40000, 7, 6, 5, nil, nil, "Idle"},
		{utils.Generate_Pilot_Name(), 30000, 5, 4, 7, nil, nil, "Idle"},
		{utils.Generate_Pilot_Name(), 60000, 10, 4, 3, nil, nil, "Idle"},
		{utils.Generate_Pilot_Name(), 55000, 6, 8, 6, nil, nil, "Idle"},
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
	if len(PilotsForSale) == 0 {
		fmt.Println("\rThere are no more pilots available for hire at this point in time.")
		fmt.Println("\r----------------------------------------------------------------------------")
		return
	}
	fmt.Printf("\r%-20s | %-8s | %-8s | %-8s | %-8s\n", "Name", "Price", "Transport", "Combat", "Mining")
	fmt.Println("\r----------------------------------------------------------------------------")
}

func PilotPurchasePrompt(pilot Pilot) func() {
	return func() {
		selectedPilot = &pilot
		CurrentMenu = &PilotPurchaseMenu
	}
}

func PilotPurchaseYes() {
	if selectedPilot != nil {
		pilotCopy := *selectedPilot
		CompanyPilots = append(CompanyPilots, pilotCopy)
		// Remove the purchased pilot from PilotsForSale
		for i, p := range PilotsForSale {
			if p.Name == pilotCopy.Name && p.Price == pilotCopy.Price {
				PilotsForSale = append(PilotsForSale[:i], PilotsForSale[i+1:]...)
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
	for _, pilot := range PilotsForSale {
		pilotCopy := pilot // avoid closure capture bug
		menuName := fmt.Sprintf("%-20s | $%-7d | %-8d | %-8d | %-8d", pilotCopy.Name, pilotCopy.Price, pilotCopy.TransportSkill, pilotCopy.CombatSkill, pilotCopy.MiningSkill)
		PilotsStoreMenuOptions = append(PilotsStoreMenuOptions, MenuItem{
			Name:     menuName,
			Callback: PilotPurchasePrompt(pilotCopy),
		})
	}
	PilotsStoreMenuOptions = append(PilotsStoreMenuOptions, MenuItem{Name: "Back", Callback: PilotsStoreBack})
	PilotsStoreMenu.Options = PilotsStoreMenuOptions
}

func PilotsStoreBack() {
	CurrentMenu = &StoreMenu
}

func init() {
	BuildPilotsStoreMenuOptions()
	PilotsStoreMenu = Menu{
		Name:    "Pilots Store Menu",
		Intro:   PilotsStoreMenuIntro,
		Options: PilotsStoreMenuOptions,
		Back:    PilotsStoreBack,
	}
}
