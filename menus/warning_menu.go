package menus

import (
	"fmt"
	"startrader/types"
	"startrader/globals"
)

// WarningMenu is a generic menu to display a warning and a confirm option, then return to a previous menu.
var WarningMenu types.Menu
var warningMenuText string
var warningMenuPrevMenu *types.Menu

// ShowWarningMenu displays the warning text and sets up the confirm option to return to the previous menu.
func ShowWarningMenu(warningText string, prevMenu *types.Menu) {
	warningMenuText = warningText
	warningMenuPrevMenu = prevMenu
	WarningMenu.Options = []types.MenuItem{
		{Name: "Confirm", Callback: func() {
			if warningMenuPrevMenu != nil {
				globals.CurrentMenu = warningMenuPrevMenu
			}
		}},
	}
	globals.CurrentMenu = &WarningMenu
}

func WarningMenuIntro(m *types.Menu) {
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Printf("\r%s\n", warningMenuText)
	fmt.Println("\r----------------------------------------------------------------------------")
}

func init() {
	WarningMenu = types.Menu{
		Name:  "Warning",
		Intro: WarningMenuIntro,
		Back: func() {
			if warningMenuPrevMenu != nil {
				globals.CurrentMenu = warningMenuPrevMenu
			}
		},
	}
}
