package menus

import (
	"fmt"
	"startrader/globals"
	"startrader/types"
)

var StationContractMenu types.Menu

func BuildStationContractMenuOptions() []types.MenuItem {
	options := []types.MenuItem{}
	if selectedDetailStation != nil {
		for i, contract := range selectedDetailStation.Contracts {
			contractCopy := contract // avoid closure capture bug
			label := fmt.Sprintf("%-3d | %-22s | %-9s | %d min %d sec | $%-7d", i+1, contractCopy.ShortName, contractCopy.Type, contractCopy.Minutes, contractCopy.Seconds, contractCopy.Payout)
			options = append(options, types.MenuItem{
				Name: label,
				Callback: func(m types.Contract) func() {
					return func() {
						selectedStationContract = &m
						globals.CurrentMenu = &StationContractDetailMenu
					}
				}(contractCopy),
			})
		}
	}
	options = append(options, types.MenuItem{Name: "Back", Callback: func() { globals.CurrentMenu = &StationDetailMenu }})
	return options
}

func StationContractMenuIntro(m *types.Menu) {
	moneyHeader := fmt.Sprintf("Credits: $%d", globals.Company.Money)
	fmt.Println("\r----------------------------------------------------------------------------")
	header := "Available Contracts:"
	fmt.Printf("\r%s%*s%s\n", header, 76-len(header)-len(moneyHeader), "", moneyHeader)
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Printf("\r%-3s | %-22s | %-9s | %-8s | %-8s\n", "#", "Name", "Type", "Duration", "Payout")
	fmt.Println("\r----------------------------------------------------------------------------")
}

func ShowStationContractMenu() {
	StationContractMenu.Options = BuildStationContractMenuOptions()
	globals.CurrentMenu = &StationContractMenu
}

func init() {
	StationContractMenu = types.Menu{
		Name:    "Station Contracts",
		Intro:   StationContractMenuIntro,
		Options: nil, // Set dynamically
		Back:    func() { globals.CurrentMenu = &StationDetailMenu },
	}
}
