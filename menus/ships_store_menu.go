package menus

import (
	"fmt"
	"startrader/utils"
)

var ShipsForSale []Ship

func init() {
	ShipsForSale = []Ship{
		{utils.Generate_Combat_Ship_Name(), 100000, "Combat", 10, 10, 10, 10, 10, nil, "Idle", nil},
		{utils.Generate_Combat_Ship_Name(), 250000, "Combat", 10, 10, 10, 10, 10, nil, "Idle", nil},
		{utils.Generate_Combat_Ship_Name(), 400000, "Combat", 10, 10, 10, 10, 10, nil, "Idle", nil},
		{utils.Generate_Combat_Ship_Name(), 750000, "Combat", 10, 10, 10, 10, 10, nil, "Idle", nil},
		{utils.Generate_Combat_Ship_Name(), 1200000, "Combat", 10, 10, 10, 10, 10, nil, "Idle", nil},
	}
}

func ShipsStoreMenuIntro(m *Menu) {
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Println("\rSpaceships for Sale:")
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Printf("\r%-20s | %-10s | %-10s\n", "Name", "Type", "Price")
	fmt.Println("\r----------------------------------------------------------------------------")
	if len(ShipsForSale) == 0 {
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
		shipCopy := *selectedShip
		CompanyShips = append(CompanyShips, shipCopy)
		// Remove the purchased ship from ShipsForSale
		for i, s := range ShipsForSale {
			if s.Name == shipCopy.Name && s.Price == shipCopy.Price {
				ShipsForSale = append(ShipsForSale[:i], ShipsForSale[i+1:]...)
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
	for _, ship := range ShipsForSale {
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
	for _, ship := range ShipsForSale {
		shipCopy := ship // avoid closure capture bug
		ShipsStoreMenuOptions = append(ShipsStoreMenuOptions, MenuItem{
			Name:     fmt.Sprintf("%-20s | %-10s | $%-9d", ship.Name, ship.Type, ship.Price),
			Callback: ShipPurchasePrompt(shipCopy),
		})
	}
	ShipsStoreMenuOptions = append(ShipsStoreMenuOptions, MenuItem{Name: "Back", Callback: ShipsStoreBack})
	ShipsStoreMenu = Menu{
		Name:    "Ships Store Menu",
		Intro:   ShipsStoreMenuIntro,
		Options: ShipsStoreMenuOptions,
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
	CurrentMenu = &StoreMenu
}
	// End of file. All menu initialization is handled in the init() above.

