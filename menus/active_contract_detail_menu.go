package menus

import (
	"fmt"
	"startrader/globals"
	"startrader/types"
)

var selectedActiveContract *types.Contract
var selectedActiveContractShip *types.Ship
var ActiveContractDetailMenu types.Menu

func ActiveContractDetailMenuIntro(m *types.Menu) {
	if selectedActiveContract == nil {
		fmt.Println("\rNo contract selected.")
		return
	}
	fmt.Println("\r----------------------------------------------------------------------------")
		header := "Contract Details:"
	moneyHeader := fmt.Sprintf("Credits: $%d", globals.Company.Money)
	fmt.Printf("\r%s%*s%s\n\r", header, 76 - len(header) - len(moneyHeader), "", moneyHeader)
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Printf("\rName: %s\n", selectedActiveContract.ShortName)
	fmt.Printf("\rType: %s\n", selectedActiveContract.Type)
	fmt.Printf("\rStatus: %s\n", selectedActiveContract.Status)
	fmt.Printf("\rDuration: %d min %d sec\n", selectedActiveContract.Minutes, selectedActiveContract.Seconds)
	fmt.Printf("\rPayout: $%d\n", selectedActiveContract.Payout)
	fmt.Printf("\rDescription: %s\n", selectedActiveContract.Description)
	if selectedActiveContractShip != nil {
		fmt.Printf("\rAssigned Ship: %s\n", selectedActiveContractShip.Name)
		if selectedActiveContractShip.AssignedPilot != nil {
			fmt.Printf("\rAssigned Pilot: %s\n", selectedActiveContractShip.AssignedPilot.Name)
		}
	}
	fmt.Println("\r----------------------------------------------------------------------------")
}

func BuildActiveContractDetailMenuOptions() []types.MenuItem {
	options := []types.MenuItem{}
	if selectedActiveContract != nil {
		if selectedActiveContract.Status == "Complete" {
			options = append(options, types.MenuItem{
				Name:     "Complete Contract",
				Callback: CompleteActiveContract,
			})
		}
		if selectedActiveContract.Status == "In Progress" {
			options = append(options, types.MenuItem{
				Name:     "Cancel Contract",
				Callback: func() { globals.CurrentMenu = &CancelContractConfirmMenu },
			})
		}
	}
	options = append(options, types.MenuItem{Name: "Back", Callback: func() { BuildActiveContractsMenuOptions(); globals.CurrentMenu = &ActiveContractsMenu }})
	return options
}

func CancelActiveContract() {
	if selectedActiveContract != nil && selectedActiveContractShip != nil {
		selectedActiveContractShip.AssignedPilot.AssignedContract = nil
		selectedActiveContractShip.AssignedPilot.Status = "Idle"
		selectedActiveContractShip.AssignedContract = nil
		selectedActiveContractShip.Status = "Idle"
		selectedActiveContract.Status = "Cancelled"
	}
	BuildActiveContractsMenuOptions()
	globals.CurrentMenu = &ActiveContractsMenu
}

func CompleteActiveContract() {
	if selectedActiveContract != nil && selectedActiveContractShip != nil && selectedActiveContract.Status == "Complete" {
		// Only show the ContractCompleteMenu, do not update state yet
		globals.CurrentMenu = &ContractCompleteMenu
	}
}

func init() {
	ActiveContractDetailMenu = types.Menu{
		Name:    "Contract Detail",
		Intro:   ActiveContractDetailMenuIntro,
		Options: nil, // Set dynamically
		Back:    func() { BuildActiveContractsMenuOptions(); globals.CurrentMenu = &ActiveContractsMenu },
	}
}
