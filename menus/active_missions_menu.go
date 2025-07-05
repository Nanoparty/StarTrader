package menus

import (
	"fmt"
	"startrader/globals"
	"startrader/types"
)

var ActiveMissionsMenu types.Menu
var ActiveMissionsOptions []types.MenuItem

func BuildActiveMissionsMenuOptions() {
	ActiveMissionsOptions = []types.MenuItem{}
	for i := range CompanyShips {
		ship := &CompanyShips[i]
		if ship.AssignedMission != nil && (ship.AssignedMission.Status == "In Progress" || ship.AssignedMission.Status == "Complete") {
			mission := ship.AssignedMission
			label := fmt.Sprintf("%-20s | %-18s | %2d min %2d sec | %-10s", mission.ShortName, ship.Name, mission.Minutes, mission.Seconds, mission.Status)
			ActiveMissionsOptions = append(ActiveMissionsOptions, types.MenuItem{
				Name: label,
				Callback: func() {
					selectedActiveMission = mission
					selectedActiveMissionShip = ship
					ActiveMissionDetailMenu.Options = BuildActiveMissionDetailMenuOptions()
					globals.CurrentMenu = &ActiveMissionDetailMenu
				},
			})
		}
	}
	ActiveMissionsOptions = append(ActiveMissionsOptions, types.MenuItem{Name: "Back", Callback: func() { BuildActiveMissionsMenuOptions(); globals.CurrentMenu = &CompanyMenu }})
	ActiveMissionsMenu.Options = ActiveMissionsOptions
}

func ActiveMissionsMenuIntro(m *types.Menu) {
	header := "Active Missions: " + globals.Company.Name
	moneyHeader := fmt.Sprintf("$%d", globals.Company.Money)
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Printf("\r%s%*s%s\n\r", header, 76 - len(header) - len(moneyHeader), "", moneyHeader)
	fmt.Println("\r----------------------------------------------------------------------------")

	activeCount := 0
	for _, ship := range CompanyShips {
		if ship.AssignedMission != nil && (ship.AssignedMission.Status == "In Progress" || ship.AssignedMission.Status == "Complete") {
			activeCount++
		}
	}
	if activeCount == 0 {
		fmt.Println("\rThere are currently no active missions.\n\r")
		return
	}
	
	fmt.Printf("\r%-20s | %-18s | %-12s | %-10s\n", "Mission", "Ship", "Duration", "Status")
	fmt.Println("\r----------------------------------------------------------------------------")
}

func init() {
	ActiveMissionsMenu = types.Menu{
		Name:    "Active Missions",
		Intro:   ActiveMissionsMenuIntro,
		Options: nil, // set by BuildActiveMissionsMenuOptions
		Back:    func() { globals.CurrentMenu = &CompanyMenu },
	}
}
