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
	if selectedActiveContract != nil && selectedActiveContractShip != nil && selectedActiveContract.Status == "Complete" {
		globals.Company.Money += selectedActiveContract.Payout
		if selectedActiveContractShip.AssignedPilot != nil {
			selectedActiveContractShip.AssignedPilot.AssignedContract = nil
			selectedActiveContractShip.AssignedPilot.Status = "Idle"
		}
		selectedActiveContractShip.AssignedContract = nil
		selectedActiveContractShip.Status = "Idle"
		selectedActiveContract.Status = "Redeemed"
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
