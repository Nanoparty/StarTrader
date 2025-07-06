package menus

import (
	"fmt"
	"startrader/globals"
	"startrader/types"
)

var ContractCompleteMenu types.Menu

func ContractCompleteMenuIntro(m *types.Menu) {
	if selectedActiveContract == nil {
		fmt.Println("\rNo contract selected.")
		return
	}
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Println("\rContract Complete!")
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Printf("\rContract: %s\n", selectedActiveContract.ShortName)
	fmt.Printf("\rPayout: $%d\n", selectedActiveContract.Payout)
	fmt.Println("\r----------------------------------------------------------------------------")
}

func ContractCompleteConfirm() {
	if selectedActiveContract == nil {
		globals.CurrentMenu = &ActiveContractsMenu
		return
	}

	// Process contract completion and payment
	if selectedActiveContractShip != nil && selectedActiveContract.Status == "Complete" {
		globals.Company.Money += selectedActiveContract.Payout
		if selectedActiveContractShip.AssignedPilot != nil {
			selectedActiveContractShip.AssignedPilot.AssignedContract = nil
			selectedActiveContractShip.AssignedPilot.Status = "Idle"
		}
		selectedActiveContractShip.AssignedContract = nil
		selectedActiveContractShip.Status = "Idle"
		selectedActiveContract.Status = "Redeemed"
	}

	// Handle relationship experience and level-ups
	if selectedActiveContract.OriginStation != nil {
		duration := float64(selectedActiveContract.Minutes*60 + selectedActiveContract.Seconds)
		experienceGained := (duration / 600) * 100 // 10 minutes = 100 exp

		selectedActiveContract.OriginStation.Experience += experienceGained
		selectedActiveContract.OriginStation.ContractsCompleted++

		if selectedActiveContract.OriginStation.Experience >= selectedActiveContract.OriginStation.ExpToNextLevel {
			selectedActiveContract.OriginStation.RelationshipLevel++
			selectedActiveContract.OriginStation.Experience = 0
			selectedActiveContract.OriginStation.ExpToNextLevel *= 1.5

			stationName := selectedActiveContract.OriginStation.Name
			newLevel := selectedActiveContract.OriginStation.RelationshipLevel
			message := fmt.Sprintf("Congratulations! Your relationship with %s has reached level %d!\n\rBetter Ships, Pilots, and Contracts are now available there.", stationName, newLevel)
			
			BuildActiveContractsMenuOptions()
			ShowWarningMenu(message, &ActiveContractsMenu)
			return
		}
	}

	BuildActiveContractsMenuOptions()
	globals.CurrentMenu = &ActiveContractsMenu
}

func ContractCompleteMenuOptions() []types.MenuItem {
	return []types.MenuItem{
		{Name: "Confirm", Callback: ContractCompleteConfirm},
	}
}

func init() {
	ContractCompleteMenu = types.Menu{
		Name:    "Contract Complete",
		Intro:   ContractCompleteMenuIntro,
		Options: ContractCompleteMenuOptions(),
		Back:    ContractCompleteConfirm,
	}
}
