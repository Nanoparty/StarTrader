package menus

import (
	"fmt"
)

var ActiveMissionsMenu Menu
var ActiveMissionsOptions []MenuItem

func BuildActiveMissionsMenuOptions() {
	ActiveMissionsOptions = []MenuItem{}
	for _, ship := range CompanyShips {
		if ship.AssignedMission != nil && (ship.AssignedMission.Status == "In Progress" || ship.AssignedMission.Status == "Complete") {
			mission := ship.AssignedMission
			label := fmt.Sprintf("%-20s | %-18s | %2d min %2d sec | %-10s", mission.ShortName, ship.Name, mission.Minutes, mission.Seconds, mission.Status)
			missionPtr := mission
			shipPtr := ship
			ActiveMissionsOptions = append(ActiveMissionsOptions, MenuItem{
				Name: label,
				Callback: func() {
					selectedActiveMission = missionPtr
					selectedActiveMissionShip = &shipPtr
					ActiveMissionDetailMenu.Options = BuildActiveMissionDetailMenuOptions()
					CurrentMenu = &ActiveMissionDetailMenu
				},
			})
		}
	}
	ActiveMissionsOptions = append(ActiveMissionsOptions, MenuItem{Name: "Back", Callback: func() { CurrentMenu = &CompanyMenu }})
	ActiveMissionsMenu.Options = ActiveMissionsOptions
}

func ActiveMissionsMenuIntro(m *Menu) {
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Println("\rActive Missions:")
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Printf("\r%-20s | %-18s | %-12s | %-10s\n", "Mission", "Ship", "Duration", "Status")
	fmt.Println("\r----------------------------------------------------------------------------")
}

func init() {
	ActiveMissionsMenu = Menu{
		Name:    "Active Missions",
		Intro:   ActiveMissionsMenuIntro,
		Options: nil, // set by BuildActiveMissionsMenuOptions
		Back:    func() { CurrentMenu = &CompanyMenu },
	}
}
