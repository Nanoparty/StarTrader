package menus

import (
	"fmt"
	"startrader/globals"
)

var CompanyShips []Ship
var CompanyShipOptions []MenuItem
var CompanyShipsMenu Menu

func CompanyShipsIntro(m *Menu) {
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Println("\r" + globals.CompanyName + " Ships:")
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Printf("\r%-20s | %-15s | %-10s\n", "Name", "Assigned Pilot", "Status")
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
		pilotName := "None"
		if shipCopy.AssignedPilot != nil {
			pilotName = shipCopy.AssignedPilot.Name
		}
		status := shipCopy.Status
		menuName := fmt.Sprintf("%-20s | %-15s | %-10s", shipCopy.Name, pilotName, status)
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
