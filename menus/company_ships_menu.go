package menus

import (
	"fmt"
	"startrader/globals"
)

var CompanyShips []Ship
var CompanyShipOptions []MenuItem
var CompanyShipsMenu Menu

func CompanyShipsIntro(m *Menu) {
	header := globals.CompanyName + " Ships:"
	moneyHeader := fmt.Sprintf("$%d", CompanyMoney)
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Printf("\r%s%*s%s\n\r", header, 76 - len(header) - len(moneyHeader), "", moneyHeader)
	fmt.Println("\r----------------------------------------------------------------------------")
	if len(CompanyShips) == 0 {
		fmt.Println("\rThere are currently no ships.\n\r")
		return
	}
	
	fmt.Printf("\r%-20s | %-10s | %-15s | %-10s\n", "Name", "Type", "Assigned Pilot", "Status")
	fmt.Println("\r----------------------------------------------------------------------------")
}

func ShipSelected(ship Ship) func() {
	return func() {
		selectedDetailShip = &ship
		ShowShipDetailMenu()
	}
}


func CompanyShipsBack() {
	CurrentMenu = &CompanyMenu
}

func BuildCompanyShipsMenuOptions() {
	CompanyShipOptions = []MenuItem{}
	for _, ship := range CompanyShips {
		shipCopy := ship // avoid closure capture bug

		assignedPilot := "None"
		if shipCopy.AssignedPilot != nil {
			assignedPilot = shipCopy.AssignedPilot.Name
		}
		status := shipCopy.Status
		menuName := fmt.Sprintf("%-20s | %-10s | %-15s | %-10s", shipCopy.Name, shipCopy.Type, assignedPilot, status)
		CompanyShipOptions = append(CompanyShipOptions, MenuItem{
			Name:     menuName,
			Callback: ShipSelected(shipCopy),
		})
	}
	CompanyShipOptions = append(CompanyShipOptions, MenuItem{Name: "Back\n\r", Callback: CompanyShipsBack})
	CompanyShipsMenu.Options = CompanyShipOptions
}

func init() {
	CompanyShipsMenu = Menu{
		Name:    "Company Ships",
		Intro:   CompanyShipsIntro,
		Options: nil, // Will be set by BuildCompanyShipsMenuOptions
		Back:    CompanyShipsBack,
	}
}
