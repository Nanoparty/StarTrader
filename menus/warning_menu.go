package menus

import "fmt"

// WarningMenu is a generic menu to display a warning and a confirm option, then return to a previous menu.
var WarningMenu Menu
var warningMenuText string
var warningMenuPrevMenu *Menu

// ShowWarningMenu displays the warning text and sets up the confirm option to return to the previous menu.
func ShowWarningMenu(warningText string, prevMenu *Menu) {
	warningMenuText = warningText
	warningMenuPrevMenu = prevMenu
	WarningMenu.Options = []MenuItem{
		{Name: "Confirm", Callback: func() {
			if warningMenuPrevMenu != nil {
				CurrentMenu = warningMenuPrevMenu
			}
		}},
	}
	CurrentMenu = &WarningMenu
}

func WarningMenuIntro(m *Menu) {
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Printf("\r%s\n", warningMenuText)
	fmt.Println("\r----------------------------------------------------------------------------")
}

func init() {
	WarningMenu = Menu{
		Name:  "Warning",
		Intro: WarningMenuIntro,
		Back: func() {
			if warningMenuPrevMenu != nil {
				CurrentMenu = warningMenuPrevMenu
			}
		},
	}
}
