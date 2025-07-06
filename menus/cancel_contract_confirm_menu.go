package menus

import (
	"fmt"
	"startrader/types"
	"startrader/globals"
)

var CancelContractConfirmMenu types.Menu

func CancelContractConfirmMenuIntro(m *types.Menu) {
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Println("\rAre you sure you want to cancel this contract?")
	fmt.Println("\r----------------------------------------------------------------------------")
	if selectedActiveContract != nil {
		fmt.Printf("\rContract: %s\n", selectedActiveContract.ShortName)
		fmt.Printf("\rAssigned Ship: %s\n", selectedActiveContractShip.Name)
	}
	fmt.Println("\r----------------------------------------------------------------------------")
}

func CancelContractConfirmYes() {
	CancelActiveContract()
}

func CancelContractConfirmNo() {
	ActiveContractDetailMenu.Options = BuildActiveContractDetailMenuOptions()
	globals.CurrentMenu = &ActiveContractDetailMenu
}

func CancelContractConfirmMenuOptions() []types.MenuItem {
	return []types.MenuItem{
		{Name: "Yes", Callback: CancelContractConfirmYes},
		{Name: "No", Callback: CancelContractConfirmNo},
	}
}

func init() {
	CancelContractConfirmMenu = types.Menu{
		Name:    "Cancel Contract?",
		Intro:   CancelContractConfirmMenuIntro,
		Options: CancelContractConfirmMenuOptions(),
		Back:    CancelContractConfirmNo,
	}
}
