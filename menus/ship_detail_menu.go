package menus

import "fmt"

var selectedDetailShip *Ship

func ShipDetailMenuIntro(m *Menu) {
	if selectedDetailShip == nil {
		fmt.Println("\rNo ship selected.")
		return
	}
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Printf("\rShip: %s\n", selectedDetailShip.Name)
	fmt.Printf("\rType: %s\n", selectedDetailShip.Type)
	fmt.Printf("\rPrice: $%d\n", selectedDetailShip.Price)
	fmt.Printf("\rStorage: %d\n", selectedDetailShip.Storage)
	fmt.Printf("\rSpeed: %d\n", selectedDetailShip.Speed)
	fmt.Printf("\rHealth: %d/%d\n", selectedDetailShip.CurrentHealth, selectedDetailShip.MaxHealth)
	fmt.Printf("\rDamage: %d\n", selectedDetailShip.Damage)
	if selectedDetailShip.AssignedPilot != nil {
		fmt.Printf("\rAssigned Pilot: %s\n", selectedDetailShip.AssignedPilot.Name)
	} else {
		fmt.Println("\rAssigned Pilot: None")
	}
	fmt.Println("\r----------------------------------------------------------------------------")
}

func ShipDetailBack() {
	CurrentMenu = &CompanyShipsMenu
}

var ShipDetailMenuOptions = []MenuItem{
	{Name: "Back", Callback: ShipDetailBack},
}

var ShipDetailMenu = Menu{
	Name:    "Ship Detail",
	Intro:   ShipDetailMenuIntro,
	Options: ShipDetailMenuOptions,
	Back:    ShipDetailBack,
}
