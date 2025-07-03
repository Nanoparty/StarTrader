package client

import (
	"fmt"
	"os"
	"startrader/globals"
	"startrader/menus"

	"golang.org/x/term"
)

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func StartRepl(cfg *globals.Config){	

	menus.CurrentMenu = &menus.MainMenu

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
		if err != nil {
			panic(err)
		}
		defer term.Restore(int(os.Stdin.Fd()), oldState)

		clearScreen()
		selected := 0

	for {
		clearScreen()

		menus.CurrentMenu.Intro(menus.CurrentMenu)

		for i, option := range menus.CurrentMenu.Options {
			if i == selected {
				// Invert colors for selected line
				fmt.Printf("\r\033[7m> %s\033[0m\n", option.Name)
			} else {
				fmt.Printf("\r  %s\n\r", option.Name)
			}
		}

		

		var b = make([]byte, 3)
		os.Stdin.Read(b)

		if b[0] == 3 { // Ctrl+C
			break
		}

		// Handle arrow keys
		if b[0] == 27 && b[1] == 91 {
			switch b[2] {
			case 65: // Up arrow
				if selected > 0 {
					selected--
				}
			case 66: // Down arrow
				if selected < len(menus.CurrentMenu.Options)-1 {
					selected++
				}
			case 67: // Right arrow
				OptionSelection(&selected)
			case 68: // Left arrow
				
				if len(menus.PreviousMenu) > 0 {
					previousMenu := menus.GetPreviousMenu()
					if previousMenu == nil {
						// Already at the main menu, do nothing
						break
					}
					if menus.CurrentMenu.Name == "Company Menu" {
						menus.PreviousMenu = menus.PreviousMenu[:0]
						menus.CurrentMenu = &menus.QuitMenu
					} else {
						menus.CurrentMenu = previousMenu
					}
					selected = 0
				} // else: at main menu, ignore left arrow

			}
		} else if b[0] == 13 || b[0] == 67 { // Enter key
			clearScreen()
			OptionSelection(&selected)
			continue
		}
	}
}

func OptionSelection(selected *int) {
	menus.AddPreviousMenu()
	menus.CurrentMenu.Options[*selected].Callback()
	*selected = 0
}


