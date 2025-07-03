package menus

import (
	"fmt"
)

var StationMissionMenu Menu

func BuildStationMissionMenuOptions() []MenuItem {
	options := []MenuItem{}
	if selectedDetailStation != nil {
		for i, mission := range selectedDetailStation.Missions {
			missionCopy := mission // avoid closure capture bug
			label := fmt.Sprintf("%-3d | %-22s | %-9s | %d min %d sec | $%-7d", i+1, missionCopy.ShortName, missionCopy.Type, missionCopy.Minutes, missionCopy.Seconds, missionCopy.Payout)
			options = append(options, MenuItem{
				Name: label,
				Callback: func(m Mission) func() {
					return func() {
						selectedStationMission = &m
						CurrentMenu = &StationMissionDetailMenu
					}
				}(missionCopy),
			})
		}
	}
	options = append(options, MenuItem{Name: "Back", Callback: func() { CurrentMenu = &StationDetailMenu }})
	return options
}

func StationMissionMenuIntro(m *Menu) {
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Println("\rAvailable Missions:")
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Printf("\r%-3s | %-22s | %-9s | %-8s | %-8s\n", "#", "Name", "Type", "Duration", "Payout")
	fmt.Println("\r----------------------------------------------------------------------------")
}

func ShowStationMissionMenu() {
	StationMissionMenu.Options = BuildStationMissionMenuOptions()
	CurrentMenu = &StationMissionMenu
}

func init() {
	StationMissionMenu = Menu{
		Name:    "Station Missions",
		Intro:   StationMissionMenuIntro,
		Options: nil, // Set dynamically
		Back:    func() { CurrentMenu = &StationDetailMenu },
	}
}
