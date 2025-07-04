package menus

import (
	"fmt"
)

var selectedActiveMission *Mission
var selectedActiveMissionShip *Ship
var ActiveMissionDetailMenu Menu

func ActiveMissionDetailMenuIntro(m *Menu) {
	if selectedActiveMission == nil {
		fmt.Println("\rNo mission selected.")
		return
	}
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Println("\rMission Details:")
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

func BuildActiveMissionDetailMenuOptions() []MenuItem {
	options := []MenuItem{}
	if selectedActiveMission != nil {
		if selectedActiveMission.Status == "Complete" {
			options = append(options, MenuItem{
				Name:     "Complete Mission",
				Callback: CompleteActiveMission,
			})
		}
		if selectedActiveMission.Status == "In Progress" {
			options = append(options, MenuItem{
				Name:     "Cancel Mission",
				Callback: func() { CurrentMenu = &CancelMissionConfirmMenu },
			})
		}
	}
	options = append(options, MenuItem{Name: "Back", Callback: func() { BuildActiveMissionsMenuOptions(); CurrentMenu = &ActiveMissionsMenu }})
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
	CurrentMenu = &ActiveMissionsMenu
}

func CompleteActiveMission() {
	if selectedActiveMission != nil && selectedActiveMissionShip != nil && selectedActiveMission.Status == "Complete" {
		// Only show the MissionCompleteMenu, do not update state yet
		CurrentMenu = &MissionCompleteMenu
	}
}

func init() {
	ActiveMissionDetailMenu = Menu{
		Name:    "Mission Detail",
		Intro:   ActiveMissionDetailMenuIntro,
		Options: nil, // Set dynamically
		Back:    func() { BuildActiveMissionsMenuOptions(); CurrentMenu = &ActiveMissionsMenu },
	}
}
