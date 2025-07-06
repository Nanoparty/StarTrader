package client

import (
	"fmt"
	"os"
	"startrader/globals"
	"startrader/menus"
	"time"

	"golang.org/x/term"
)

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func StartRepl(cfg *globals.Config){	

	globals.CurrentMenu = &menus.MainMenu

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
		if err != nil {
			panic(err)
		}
		defer term.Restore(int(os.Stdin.Fd()), oldState)

		clearScreen()

	refreshTicker := (*time.Ticker)(nil)
refreshStop := make(chan struct{})

for {
	// If we're in the Active Contracts menu, start a ticker to refresh every second
	if globals.CurrentMenu == &menus.ActiveContractsMenu {
		if refreshTicker == nil {
			refreshTicker = time.NewTicker(time.Second)
			go func() {
				for {
					select {
					case <-refreshTicker.C:
						menus.BuildActiveContractsMenuOptions()
						clearScreen()
						globals.CurrentMenu.Intro(globals.CurrentMenu)
						for i, option := range globals.CurrentMenu.Options {
							if i == globals.CurrentMenu.Selected {
								fmt.Printf("\r\033[7m> %s\033[0m\n\r", option.Name)
							} else {
								fmt.Printf("\r  %s\n\r", option.Name)
							}
						}
					case <-refreshStop:
						refreshTicker.Stop()
						refreshTicker = nil
						return
					}
				}
			}()
		}
		// Always refresh options before showing
		menus.BuildActiveContractsMenuOptions()
	}

	clearScreen()
	globals.CurrentMenu.Intro(globals.CurrentMenu)
	for i, option := range globals.CurrentMenu.Options {
		if i == globals.CurrentMenu.Selected {
			fmt.Printf("\r\033[7m> %s\033[0m\n\r", option.Name)
		} else {
			fmt.Printf("\r  %s\n\r", option.Name)
		}
	}

	var b = make([]byte, 3)
	os.Stdin.Read(b)

	// If we leave the Active Contracts menu, stop the ticker
	if refreshTicker != nil && globals.CurrentMenu != &menus.ActiveContractsMenu {
		refreshStop <- struct{}{}
	}

		if b[0] == 3 { // Ctrl+C
			break
		}

		// Handle arrow keys
		if b[0] == 27 && b[1] == 91 {
			switch b[2] {
			case 65: // Up arrow
				if globals.CurrentMenu.Selected > 0 {
					globals.CurrentMenu.Selected--
				}
			case 66: // Down arrow
				if globals.CurrentMenu.Selected < len(globals.CurrentMenu.Options)-1 {
					globals.CurrentMenu.Selected++
				}
			case 67: // Right arrow
				OptionSelection(&globals.CurrentMenu.Selected)
			case 68: // Left arrow
				globals.CurrentMenu.Back()
			}
		} else if b[0] == 13 || b[0] == 67 { // Enter key
			clearScreen()
			OptionSelection(&globals.CurrentMenu.Selected)
			continue
		}
	}
}

func OptionSelection(selected *int) {
	globals.CurrentMenu.Options[*selected].Callback()
}


