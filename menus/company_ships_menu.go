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
}

func ShipSelected(ship Ship) func() {
	return func() {
		
	}
}

func CompanyShipsBack() {
	CurrentMenu = &CompanyMenu
}

func init() {
	CompanyShipOptions = []MenuItem{}
	for _, ship := range CompanyShips {
		shipCopy := ship // avoid closure capture bug
		CompanyShipOptions = append(CompanyShipOptions, MenuItem{
			Name:     shipCopy.Name,
			Callback: ShipSelected(shipCopy),
		})
	}
	CompanyShipOptions = append(CompanyShipOptions, MenuItem{Name: "Back\n\r", Callback: CompanyShipsBack})

	CompanyShipsMenu = Menu{
		Name:    "Company Ships",
		Intro:   CompanyShipsIntro,
		Options: CompanyShipOptions,
		Back:    CompanyShipsBack,
	}
}
