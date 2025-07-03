package menus

import "fmt"

func StoreMenuIntro(m *Menu) {
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Println("\rStore Menu:")
	fmt.Println("\r----------------------------------------------------------------------------")
}

func StoreShips() {
	CurrentMenu = &ShipsStoreMenu
}

func StorePilots() {
	fmt.Println("Store Pilots\n\r")
}

func StoreBack() {
	CurrentMenu = &CompanyMenu
}

var StoreMenuOptions []MenuItem
var StoreMenu Menu

func init() {
	StoreMenuOptions = []MenuItem{
		{Name: "Ships", Callback: StoreShips},
		{Name: "Pilots", Callback: StorePilots},
		{Name: "Back", Callback: StoreBack},
	}
	StoreMenu = Menu{
		Name:    "Store Menu",
		Intro:   StoreMenuIntro,
		Options: StoreMenuOptions,
		Back:    StoreBack,
	}
}
