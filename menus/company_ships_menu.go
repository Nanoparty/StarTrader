package menus

import (
	"fmt"
	"startrader/globals"
	"startrader/types"
)

var CompanyShipOptions []types.MenuItem
var CompanyShipsMenu types.Menu

func CompanyShipsIntro(m *types.Menu) {
	header := globals.Company.Name + " Ships:"
	moneyHeader := fmt.Sprintf("$%d", globals.Company.Money)
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Printf("\r%s%*s%s\n\r", header, 76 - len(header) - len(moneyHeader), "", moneyHeader)
	fmt.Println("\r----------------------------------------------------------------------------")
	if len(globals.Company.Ships) == 0 {
		fmt.Println("\rThere are currently no ships.\n\r")
		return
	}
	
	fmt.Printf("\r%-20s | %-10s | %-15s | %-10s\n", "Name", "Type", "Assigned Pilot", "Status")
	fmt.Println("\r----------------------------------------------------------------------------")
}

func ShipSelected(ship types.Ship) func() {
	return func() {
		selectedDetailShip = &ship
		ShowShipDetailMenu()
	}
}


func CompanyShipsBack() {
	globals.CurrentMenu = &CompanyMenu
}

func BuildCompanyShipsMenuOptions() {
	CompanyShipOptions = []types.MenuItem{}
	for _, ship := range globals.Company.Ships {
		shipCopy := ship // avoid closure capture bug

		assignedPilot := "None"
		if shipCopy.AssignedPilot != nil {
			assignedPilot = shipCopy.AssignedPilot.Name
		}
		status := shipCopy.Status
		menuName := fmt.Sprintf("%-20s | %-10s | %-15s | %-10s", shipCopy.Name, shipCopy.Type, assignedPilot, status)
		CompanyShipOptions = append(CompanyShipOptions, types.MenuItem{
			Name:     menuName,
			Callback: ShipSelected(shipCopy),
		})
	}
	CompanyShipOptions = append(CompanyShipOptions, types.MenuItem{Name: "Back\n\r", Callback: CompanyShipsBack})
	CompanyShipsMenu.Options = CompanyShipOptions
}

func init() {
	CompanyShipsMenu = types.Menu{
		Name:    "Company Ships",
		Intro:   CompanyShipsIntro,
		Options: nil, // Will be set by BuildCompanyShipsMenuOptions
		Back:    CompanyShipsBack,
	}
}
