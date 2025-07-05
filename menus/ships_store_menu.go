package menus

import (
	"fmt"
)

// ShipsForSale is now per-station, see selectedDetailStation.ShipsForSale

func ShowStationShipsStoreMenu() {
	if selectedDetailStation == nil {
		fmt.Println("\rNo station selected.")
		return
	}
	BuildStationShipsStoreMenuOptions()
	CurrentMenu = &ShipsStoreMenu
}

func BuildStationShipsStoreMenuOptions() {
	ShipsStoreMenuOptions = []MenuItem{}
	for _, ship := range selectedDetailStation.ShipsForSale {
		shipCopy := ship // avoid closure capture bug
		ShipsStoreMenuOptions = append(ShipsStoreMenuOptions, MenuItem{
			Name:     fmt.Sprintf("%-20s | %-10s | $%-9d", ship.Name, ship.Type, ship.Price),
			Callback: ShipPurchasePrompt(shipCopy),
		})
	}
	ShipsStoreMenuOptions = append(ShipsStoreMenuOptions, MenuItem{Name: "Back", Callback: ShipsStoreBack})
	ShipsStoreMenu.Options = ShipsStoreMenuOptions
}

func ShipsStoreMenuIntro(m *Menu) {
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Println("\rSpaceships for Sale:")
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Printf("\r%-20s | %-10s | %-10s\n", "Name", "Type", "Price")
	fmt.Println("\r----------------------------------------------------------------------------")
	if selectedDetailStation == nil || len(selectedDetailStation.ShipsForSale) == 0 {
		fmt.Println("\rThere are no more ships available for purchase at this point in time.")
		fmt.Println("\r----------------------------------------------------------------------------")
		return
	}
}

var selectedShip *Ship

func ShipPurchasePrompt(ship Ship) func() {
	return func() {
		selectedShip = &ship
		CurrentMenu = &ShipPurchaseMenu
	}
}

func ShipPurchaseYes() {
	if selectedShip != nil {
		if CompanyMoney < selectedShip.Price {
			ShowWarningMenu("Insufficient funds to purchase this ship.", &ShipPurchaseMenu)
			return
		}
		CompanyMoney -= selectedShip.Price
		shipCopy := *selectedShip
		CompanyShips = append(CompanyShips, shipCopy)
		// Remove the purchased ship from ShipsForSale
		for i, s := range selectedDetailStation.ShipsForSale {
			if s.Name == shipCopy.Name && s.Price == shipCopy.Price {
				selectedDetailStation.ShipsForSale = append(selectedDetailStation.ShipsForSale[:i], selectedDetailStation.ShipsForSale[i+1:]...)
				break
			}
		}
	}
	selectedShip = nil
	BuildShipsStoreMenuOptions()
	CurrentMenu = &ShipsStoreMenu
}

func BuildShipsStoreMenuOptions() {
	ShipsStoreMenuOptions = []MenuItem{}
	for _, ship := range selectedDetailStation.ShipsForSale {
		shipCopy := ship // avoid closure capture bug
		ShipsStoreMenuOptions = append(ShipsStoreMenuOptions, MenuItem{
			Name:     fmt.Sprintf("%-20s | %-10s | $%-9d", ship.Name, ship.Type, ship.Price),
			Callback: ShipPurchasePrompt(shipCopy),
		})
	}
	ShipsStoreMenuOptions = append(ShipsStoreMenuOptions, MenuItem{Name: "Back", Callback: ShipsStoreBack})
	ShipsStoreMenu.Options = ShipsStoreMenuOptions
}

func ShipPurchaseNo() {
	selectedShip = nil
	CurrentMenu = &ShipsStoreMenu
}

func ShipPurchaseMenuIntro(m *Menu) {
	if selectedShip != nil {
		fmt.Println("\r----------------------------------------------------------------------------")
		fmt.Println("\rShip Details:")
		fmt.Println("\r----------------------------------------------------------------------------")
		fmt.Printf("\rName: %s\n", selectedShip.Name)
		fmt.Printf("\rType: %s\n", selectedShip.Type)
		fmt.Printf("\rPrice: $%d\n", selectedShip.Price)
		fmt.Printf("\rStorage: %d\n", selectedShip.Storage)
		fmt.Printf("\rSpeed: %d\n", selectedShip.Speed)
		fmt.Printf("\rHealth: %d/%d\n", selectedShip.CurrentHealth, selectedShip.MaxHealth)
		fmt.Printf("\rDamage: %d\n", selectedShip.Damage)
		fmt.Println("\r----------------------------------------------------------------------------")
		fmt.Printf("\rPurchase this ship for $%d?\n", selectedShip.Price)
	} else {
		fmt.Println("\rNo ship selected.")
	}
}

func ShipPurchaseMenuOptions() []MenuItem {
	return []MenuItem{
		{Name: "Yes", Callback: ShipPurchaseYes},
		{Name: "No", Callback: ShipPurchaseNo},
	}
}

var ShipPurchaseMenu Menu

var ShipsStoreMenuOptions []MenuItem
var ShipsStoreMenu Menu

func init() {
	ShipsStoreMenuOptions = []MenuItem{}
	ShipsStoreMenu = Menu{
		Name:    "Ships Store Menu",
		Intro:   ShipsStoreMenuIntro,
		Options: ShipsStoreMenuOptions, // will be set dynamically
		Back:    ShipsStoreBack,
	}

	ShipPurchaseMenu = Menu{
		Name:    "Purchase Ship?",
		Intro:   ShipPurchaseMenuIntro,
		Options: ShipPurchaseMenuOptions(),
		Back:    ShipPurchaseNo,
	}
}

func ShipsStoreBack() {
	CurrentMenu = &StationDetailMenu
}
	// End of file. All menu initialization is handled in the init() above.

