package menus

import (
	"fmt"
	"startrader/globals"
	"startrader/types"
)

func ShowStationPilotsStoreMenu() {
	if selectedDetailStation == nil {
		fmt.Println("\rNo station selected.")
		return
	}
	BuildStationPilotsStoreMenuOptions()
	globals.CurrentMenu = &PilotsStoreMenu
}

func BuildStationPilotsStoreMenuOptions() {
	PilotsStoreMenuOptions = []types.MenuItem{}
	for i, pilot := range selectedDetailStation.PilotsForSale {
		pilotCopy := pilot // avoid closure capture bug
		menuName := fmt.Sprintf("%-20s | $%-7d | %-8d | %-8d | %-8d", pilotCopy.Name, pilotCopy.Price, pilotCopy.TransportSkill, pilotCopy.CombatSkill, pilotCopy.MiningSkill)
		PilotsStoreMenuOptions = append(PilotsStoreMenuOptions, types.MenuItem{
			Name:     menuName,
			Callback: PilotPurchasePrompt(pilotCopy, i),
		})
	}
	PilotsStoreMenuOptions = append(PilotsStoreMenuOptions, types.MenuItem{Name: "Back", Callback: PilotsStoreBack})
	PilotsStoreMenu.Options = PilotsStoreMenuOptions
}

var PilotsStoreMenuOptions []types.MenuItem
var PilotsStoreMenu types.Menu
var selectedPilot *types.Pilot

func init() {
	PilotsStoreMenu = types.Menu{
		Name:  "Pilots Store Menu",
		Intro: PilotsStoreMenuIntro,
	}

	PilotPurchaseMenu = types.Menu{
		Name:    "Hire Pilot?",
		Intro:   PilotPurchaseMenuIntro,
		Options: PilotPurchaseMenuOptions(),
		Back:    PilotPurchaseNo,
	}
}

func PilotsStoreMenuIntro(m *types.Menu) {
	moneyHeader := fmt.Sprintf("$%d", globals.Company.Money)
	fmt.Println("\r----------------------------------------------------------------------------")
	header := "Pilots for Hire:"
	fmt.Printf("\r%s%*s%s\n", header, 76-len(header)-len(moneyHeader), "", moneyHeader)
	fmt.Println("\r----------------------------------------------------------------------------")
	if selectedDetailStation == nil || len(selectedDetailStation.PilotsForSale) == 0 {
		fmt.Println("\rThere are no more pilots available for hire at this point in time.")
		fmt.Println("\r----------------------------------------------------------------------------")
		return
	}
	fmt.Printf("\r%-20s | %-8s | %-8s | %-8s | %-8s\n", "Name", "Price", "Transport", "Combat", "Mining")
	fmt.Println("\r----------------------------------------------------------------------------")
}

func PilotPurchasePrompt(pilot types.Pilot, i int) func() {
	return func() {
		selectedPilot = &selectedDetailStation.PilotsForSale[i]
		globals.CurrentMenu = &PilotPurchaseMenu
	}
}

func PilotPurchaseYes() {
	if selectedPilot != nil {
		if globals.Company.Money < selectedPilot.Price {
			ShowWarningMenu("Insufficient funds to hire this pilot.", &PilotPurchaseMenu)
			return
		}
		globals.Company.Money -= selectedPilot.Price
		// --- Station Relationship Logic ---
		if selectedDetailStation != nil {
			selectedDetailStation.MoneySpent += selectedPilot.Price
			// Add experience as a float for every $1000 spent (partial allowed)
			expGained := float64(selectedPilot.Price) / 1000.0
			selectedDetailStation.Experience += expGained
			// Level up if experience threshold reached
			for selectedDetailStation.Experience >= selectedDetailStation.ExpToNextLevel {
				selectedDetailStation.RelationshipLevel++
				selectedDetailStation.Experience -= selectedDetailStation.ExpToNextLevel
				// Increase required experience for next level by 50%
				selectedDetailStation.ExpToNextLevel *= 1.5
				if selectedDetailStation.ExpToNextLevel < 1.0 {
					selectedDetailStation.ExpToNextLevel = 1.0
				}
			}
		}
		// --- End Station Relationship Logic ---
		pilotCopy := *selectedPilot
		globals.Company.Pilots = append(globals.Company.Pilots, pilotCopy)
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
	globals.CurrentMenu = &PilotsStoreMenu
}

func PilotPurchaseNo() {
	selectedPilot = nil
	globals.CurrentMenu = &PilotsStoreMenu
}

func PilotPurchaseMenuIntro(m *types.Menu) {
	if selectedPilot != nil {
		moneyHeader := fmt.Sprintf("$%d", globals.Company.Money)
		fmt.Println("\r----------------------------------------------------------------------------")
		header := "Pilot Details:"
		fmt.Printf("\r%s%*s%s\n", header, 76-len(header)-len(moneyHeader), "", moneyHeader)
		fmt.Println("\r----------------------------------------------------------------------------")
		fmt.Printf("\rName:      %s\n", selectedPilot.Name)
		fmt.Printf("\rPrice:     $%d\n", selectedPilot.Price)
		fmt.Printf("\rTransport: %d\n", selectedPilot.TransportSkill)
		fmt.Printf("\rCombat:    %d\n", selectedPilot.CombatSkill)
		fmt.Printf("\rMining:    %d\n", selectedPilot.MiningSkill)
		fmt.Println("\r----------------------------------------------------------------------------")
		fmt.Printf("\rHire %s for $%d?\n", selectedPilot.Name, selectedPilot.Price)
	} else {
		fmt.Println("\rNo pilot selected.")
	}
}

func PilotPurchaseMenuOptions() []types.MenuItem {
	return []types.MenuItem{
		{Name: "Yes", Callback: PilotPurchaseYes},
		{Name: "No", Callback: PilotPurchaseNo},
	}
}

var PilotPurchaseMenu types.Menu

func BuildPilotsStoreMenuOptions() {
	PilotsStoreMenuOptions = []types.MenuItem{}
	for i, pilot := range selectedDetailStation.PilotsForSale {
		pilotCopy := pilot // avoid closure capture bug
		menuName := fmt.Sprintf("%-20s | $%-7d | %-8d | %-8d | %-8d", pilotCopy.Name, pilotCopy.Price, pilotCopy.TransportSkill, pilotCopy.CombatSkill, pilotCopy.MiningSkill)
		PilotsStoreMenuOptions = append(PilotsStoreMenuOptions, types.MenuItem{
			Name:     menuName,
			Callback: PilotPurchasePrompt(pilotCopy, i),
		})
	}
	PilotsStoreMenuOptions = append(PilotsStoreMenuOptions, types.MenuItem{Name: "Back", Callback: PilotsStoreBack})
	PilotsStoreMenu.Options = PilotsStoreMenuOptions
}

func PilotsStoreBack() {
	globals.CurrentMenu = &StationDetailMenu
}

func init() {
	PilotsStoreMenuOptions = []types.MenuItem{}
	PilotsStoreMenu = types.Menu{
		Name:    "Pilots Store Menu",
		Intro:   PilotsStoreMenuIntro,
		Options: PilotsStoreMenuOptions, // will be set dynamically
		Back:    PilotsStoreBack,
	}
}
