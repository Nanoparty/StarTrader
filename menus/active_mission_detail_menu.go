package menus

import (
	"fmt"
	"startrader/globals"
	"startrader/types"
)

var selectedActiveMission *types.Mission
var selectedActiveMissionShip *types.Ship
var ActiveMissionDetailMenu types.Menu

func ActiveMissionDetailMenuIntro(m *types.Menu) {
	if selectedActiveMission == nil {
		fmt.Println("\rNo mission selected.")
		return
	}
	fmt.Println("\r----------------------------------------------------------------------------")
		header := "Mission Details:"
	moneyHeader := fmt.Sprintf("$%d", globals.Company.Money)
	fmt.Printf("\r%s%*s%s\n\r", header, 76 - len(header) - len(moneyHeader), "", moneyHeader)
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Printf("\rName: %s\n", selectedActiveMission.ShortName)
	fmt.Printf("\rType: %s\n", selectedActiveMission.Type)
	fmt.Printf("\rStatus: %s\n", selectedActiveMission.Status)
	fmt.Printf("\rDuration: %d min %d sec\n", selectedActiveMission.Minutes, selectedActiveMission.Seconds)
	fmt.Printf("\rPayout: $%d\n", selectedActiveMission.Payout)
	fmt.Printf("\rDescription: %s\n", selectedActiveMission.Description)
	if selectedActiveMissionShip != nil {
		fmt.Printf("\rAssigned Ship: %s\n", selectedActiveMissionShip.Name)
		if selectedActiveMissionShip.AssignedPilot != nil {
			fmt.Printf("\rAssigned Pilot: %s\n", selectedActiveMissionShip.AssignedPilot.Name)
		}
	}
	fmt.Println("\r----------------------------------------------------------------------------")
}

func BuildActiveMissionDetailMenuOptions() []types.MenuItem {
	options := []types.MenuItem{}
	if selectedActiveMission != nil {
		if selectedActiveMission.Status == "Complete" {
			options = append(options, types.MenuItem{
				Name:     "Complete Mission",
				Callback: CompleteActiveMission,
			})
		}
		if selectedActiveMission.Status == "In Progress" {
			options = append(options, types.MenuItem{
				Name:     "Cancel Mission",
				Callback: func() { globals.CurrentMenu = &CancelMissionConfirmMenu },
			})
		}
	}
	options = append(options, types.MenuItem{Name: "Back", Callback: func() { BuildActiveMissionsMenuOptions(); globals.CurrentMenu = &ActiveMissionsMenu }})
	return options
}

func CancelActiveMission() {
	if selectedActiveMission != nil && selectedActiveMissionShip != nil {
		selectedActiveMissionShip.AssignedPilot.AssignedMission = nil
		selectedActiveMissionShip.AssignedPilot.Status = "Idle"
		selectedActiveMissionShip.AssignedMission = nil
		selectedActiveMissionShip.Status = "Idle"
		selectedActiveMission.Status = "Cancelled"
	}
	BuildActiveMissionsMenuOptions()
	globals.CurrentMenu = &ActiveMissionsMenu
}

func CompleteActiveMission() {
	if selectedActiveMission != nil && selectedActiveMissionShip != nil && selectedActiveMission.Status == "Complete" {
		// Only show the MissionCompleteMenu, do not update state yet
		globals.CurrentMenu = &MissionCompleteMenu
	}
}

func init() {
	ActiveMissionDetailMenu = types.Menu{
		Name:    "Mission Detail",
		Intro:   ActiveMissionDetailMenuIntro,
		Options: nil, // Set dynamically
		Back:    func() { BuildActiveMissionsMenuOptions(); globals.CurrentMenu = &ActiveMissionsMenu },
	}
}
