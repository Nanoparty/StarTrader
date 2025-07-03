package menus

import (
	"fmt"
	"startrader/globals"
)

var CompanyShips []string

func CompanyShipsIntro(m *Menu) {
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Println("\r" + globals.CompanyName + " Ships:")
	fmt.Println("\r----------------------------------------------------------------------------")

	if len(CompanyShips) == 0 {
		fmt.Println("\r  (No ships owned)")
	} else {
		fmt.Println("\r  Owned Ships:")
		for _, ship := range CompanyShips {
			fmt.Printf("\r    - %s\n", ship)
		}
	}
}

var CompanyShipsMenu = Menu{
	Name:    "Company Ships",
	Intro:   CompanyShipsIntro,
	Options: []MenuItem{},
}
