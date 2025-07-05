package menus

import (
	"fmt"
	"startrader/globals"
	"startrader/types"
)

// ShipsForSale is now per-station, see selectedDetailStation.ShipsForSale

func ShowStationShipsStoreMenu() {
	if selectedDetailStation == nil {
		fmt.Println("\rNo station selected.")
		return
	}
	BuildStationShipsStoreMenuOptions()
	globals.CurrentMenu = &ShipsStoreMenu
}

func BuildStationShipsStoreMenuOptions() {
	ShipsStoreMenuOptions = []types.MenuItem{}
	for _, ship := range selectedDetailStation.ShipsForSale {
		shipCopy := ship // avoid closure capture bug
		ShipsStoreMenuOptions = append(ShipsStoreMenuOptions, types.MenuItem{
			Name:     fmt.Sprintf("%-20s | %-10s | $%-9d", ship.Name, ship.Type, ship.Price),
			Callback: ShipPurchasePrompt(shipCopy),
		})
	}
	ShipsStoreMenuOptions = append(ShipsStoreMenuOptions, types.MenuItem{Name: "Back", Callback: ShipsStoreBack})
	ShipsStoreMenu.Options = ShipsStoreMenuOptions
}

func ShipsStoreMenuIntro(m *types.Menu) {
	moneyHeader := fmt.Sprintf("$%d", globals.Company.Money)
	fmt.Println("\r----------------------------------------------------------------------------")
	header := "Spaceships for Sale:"
	fmt.Printf("\r%s%*s%s\n", header, 76-len(header)-len(moneyHeader), "", moneyHeader)
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Printf("\r%-20s | %-10s | %-10s\n", "Name", "Type", "Price")
	fmt.Println("\r----------------------------------------------------------------------------")
	if selectedDetailStation == nil || len(selectedDetailStation.ShipsForSale) == 0 {
		fmt.Println("\rThere are no more ships available for purchase at this point in time.")
		fmt.Println("\r----------------------------------------------------------------------------")
		return
	}
}

var selectedShip *types.Ship

func ShipPurchasePrompt(ship types.Ship) func() {
	return func() {
		selectedShip = &ship
		globals.CurrentMenu = &ShipPurchaseMenu
	}
}

func ShipPurchaseYes() {
	if selectedShip != nil {
		if globals.Company.Money < selectedShip.Price {
			ShowWarningMenu("Insufficient funds to purchase this ship.", &ShipPurchaseMenu)
			return
		}
		globals.Company.Money -= selectedShip.Price
		// --- Station Relationship Logic ---
		if selectedDetailStation != nil {
			selectedDetailStation.MoneySpent += selectedShip.Price
			// Add experience as a float for every $1000 spent (partial allowed)
			expGained := float64(selectedShip.Price) / 1000.0
			selectedDetailStation.Experience += expGained
			// Level up if experience threshold reached
			for selectedDetailStation.Experience >= selectedDetailStation.ExpToNextLevel {
				selectedDetailStation.RelationshipLevel++
				selectedDetailStation.Experience -= selectedDetailStation.ExpToNextLevel
				// Increase required experience for next level by 50%
				selectedDetailStation.ExpToNextLevel *= 1.5
				if selectedDetailStation.ExpToNextLevel < 1.0 {
					selectedDetailStation.ExpToNextLevel = 1.0
				}
			}
		}
		// --- End Station Relationship Logic ---
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
	globals.CurrentMenu = &ShipsStoreMenu
}

func BuildShipsStoreMenuOptions() {
	ShipsStoreMenuOptions = []types.MenuItem{}
	for _, ship := range selectedDetailStation.ShipsForSale {
		shipCopy := ship // avoid closure capture bug
		ShipsStoreMenuOptions = append(ShipsStoreMenuOptions, types.MenuItem{
			Name:     fmt.Sprintf("%-20s | %-10s | $%-9d", ship.Name, ship.Type, ship.Price),
			Callback: ShipPurchasePrompt(shipCopy),
		})
	}
	ShipsStoreMenuOptions = append(ShipsStoreMenuOptions, types.MenuItem{Name: "Back", Callback: ShipsStoreBack})
	ShipsStoreMenu.Options = ShipsStoreMenuOptions
}

func ShipPurchaseNo() {
	selectedShip = nil
	globals.CurrentMenu = &ShipsStoreMenu
}

func ShipPurchaseMenuIntro(m *types.Menu) {
	if selectedShip != nil {
		moneyHeader := fmt.Sprintf("$%d", globals.Company.Money)
		fmt.Println("\r----------------------------------------------------------------------------")
		header := "Ship Details:"
		fmt.Printf("\r%s%*s%s\n", header, 76-len(header)-len(moneyHeader), "", moneyHeader)
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

func ShipPurchaseMenuOptions() []types.MenuItem {
	return []types.MenuItem{
		{Name: "Yes", Callback: ShipPurchaseYes},
		{Name: "No", Callback: ShipPurchaseNo},
	}
}

var ShipPurchaseMenu types.Menu

var ShipsStoreMenuOptions []types.MenuItem
var ShipsStoreMenu types.Menu

func init() {
	ShipsStoreMenuOptions = []types.MenuItem{}
	ShipsStoreMenu = types.Menu{
		Name:    "Ships Store Menu",
		Intro:   ShipsStoreMenuIntro,
		Options: ShipsStoreMenuOptions, // will be set dynamically
		Back:    ShipsStoreBack,
	}

	ShipPurchaseMenu = types.Menu{
		Name:    "Purchase Ship?",
		Intro:   ShipPurchaseMenuIntro,
		Options: ShipPurchaseMenuOptions(),
		Back:    ShipPurchaseNo,
	}
}

func ShipsStoreBack() {
	globals.CurrentMenu = &StationDetailMenu
}
	// End of file. All menu initialization is handled in the init() above.

