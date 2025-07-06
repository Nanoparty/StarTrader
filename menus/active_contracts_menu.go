package menus

import (
	"fmt"
	"startrader/globals"
	"startrader/types"
)

var ActiveContractsMenu types.Menu
var ActiveContractsOptions []types.MenuItem

func BuildActiveContractsMenuOptions() {
	ActiveContractsOptions = []types.MenuItem{}
	for i := range globals.Company.Ships {
		ship := &globals.Company.Ships[i]
		if ship.AssignedContract != nil && (ship.AssignedContract.Status == "In Progress" || ship.AssignedContract.Status == "Complete") {
			contract := ship.AssignedContract
			label := fmt.Sprintf("%-20s | %-18s | %2d min %2d sec | %-10s", contract.ShortName, ship.Name, contract.Minutes, contract.Seconds, contract.Status)
			ActiveContractsOptions = append(ActiveContractsOptions, types.MenuItem{
				Name: label,
				Callback: func() {
					selectedActiveContract = contract
					selectedActiveContractShip = ship
					ActiveContractDetailMenu.Options = BuildActiveContractDetailMenuOptions()
					globals.CurrentMenu = &ActiveContractDetailMenu
				},
			})
		}
	}
	ActiveContractsOptions = append(ActiveContractsOptions, types.MenuItem{Name: "Back", Callback: func() { BuildActiveContractsMenuOptions(); globals.CurrentMenu = &CompanyMenu }})
	ActiveContractsMenu.Options = ActiveContractsOptions
}

func ActiveContractsMenuIntro(m *types.Menu) {
	header := "Active Contracts: " + globals.Company.Name
	moneyHeader := fmt.Sprintf("$%d", globals.Company.Money)
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Printf("\r%s%*s%s\n\r", header, 76 - len(header) - len(moneyHeader), "", moneyHeader)
	fmt.Println("\r----------------------------------------------------------------------------")

	activeCount := 0
	for _, ship := range globals.Company.Ships {
		if ship.AssignedContract != nil && (ship.AssignedContract.Status == "In Progress" || ship.AssignedContract.Status == "Complete") {
			activeCount++
		}
	}
	if activeCount == 0 {
		fmt.Println("\rThere are currently no active contracts.\n\r")
		return
	}
	
	fmt.Printf("\r%-20s | %-18s | %-12s | %-10s\n", "Contract", "Ship", "Duration", "Status")
	fmt.Println("\r----------------------------------------------------------------------------")
}

func init() {
	ActiveContractsMenu = types.Menu{
		Name:    "Active Contracts",
		Intro:   ActiveContractsMenuIntro,
		Options: nil, // set by BuildActiveContractsMenuOptions
		Back:    func() { globals.CurrentMenu = &CompanyMenu },
	}
}
