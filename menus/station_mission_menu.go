package menus

import (
	"fmt"
	"startrader/globals"
	"startrader/types"
)

var StationMissionMenu types.Menu

func BuildStationMissionMenuOptions() []types.MenuItem {
	options := []types.MenuItem{}
	if selectedDetailStation != nil {
		for i, mission := range selectedDetailStation.Missions {
			missionCopy := mission // avoid closure capture bug
			label := fmt.Sprintf("%-3d | %-22s | %-9s | %d min %d sec | $%-7d", i+1, missionCopy.ShortName, missionCopy.Type, missionCopy.Minutes, missionCopy.Seconds, missionCopy.Payout)
			options = append(options, types.MenuItem{
				Name: label,
				Callback: func(m types.Mission) func() {
					return func() {
						selectedStationMission = &m
						globals.CurrentMenu = &StationMissionDetailMenu
					}
				}(missionCopy),
			})
		}
	}
	options = append(options, types.MenuItem{Name: "Back", Callback: func() { globals.CurrentMenu = &StationDetailMenu }})
	return options
}

func StationMissionMenuIntro(m *types.Menu) {
	moneyHeader := fmt.Sprintf("$%d", globals.Company.Money)
	fmt.Println("\r----------------------------------------------------------------------------")
	header := "Available Missions:"
	fmt.Printf("\r%s%*s%s\n", header, 76-len(header)-len(moneyHeader), "", moneyHeader)
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Printf("\r%-3s | %-22s | %-9s | %-8s | %-8s\n", "#", "Name", "Type", "Duration", "Payout")
	fmt.Println("\r----------------------------------------------------------------------------")
}

func ShowStationMissionMenu() {
	StationMissionMenu.Options = BuildStationMissionMenuOptions()
	globals.CurrentMenu = &StationMissionMenu
}

func init() {
	StationMissionMenu = types.Menu{
		Name:    "Station Missions",
		Intro:   StationMissionMenuIntro,
		Options: nil, // Set dynamically
		Back:    func() { globals.CurrentMenu = &StationDetailMenu },
	}
}
