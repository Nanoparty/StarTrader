package menus

import (
	"fmt"
)

var MissionCompleteMenu Menu

func MissionCompleteMenuIntro(m *Menu) {
	if selectedActiveMission == nil {
		fmt.Println("\rNo mission selected.")
		return
	}
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Println("\rMission Complete!")
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Printf("\rMission: %s\n", selectedActiveMission.ShortName)
	fmt.Printf("\rPayout: $%d\n", selectedActiveMission.Payout)
	fmt.Println("\r----------------------------------------------------------------------------")
}

func MissionCompleteConfirm() {
	if selectedActiveMission != nil && selectedActiveMissionShip != nil && selectedActiveMission.Status == "Complete" {
		CompanyMoney += selectedActiveMission.Payout
		if selectedActiveMissionShip.AssignedPilot != nil {
			selectedActiveMissionShip.AssignedPilot.AssignedMission = nil
			selectedActiveMissionShip.AssignedPilot.Status = "Idle"
		}
		selectedActiveMissionShip.AssignedMission = nil
		selectedActiveMissionShip.Status = "Idle"
		selectedActiveMission.Status = "Redeemed"
	}
	BuildActiveMissionsMenuOptions()
	CurrentMenu = &ActiveMissionsMenu
}

func MissionCompleteMenuOptions() []MenuItem {
	return []MenuItem{
		{Name: "Confirm", Callback: MissionCompleteConfirm},
	}
}

func init() {
	MissionCompleteMenu = Menu{
		Name:    "Mission Complete",
		Intro:   MissionCompleteMenuIntro,
		Options: MissionCompleteMenuOptions(),
		Back:    MissionCompleteConfirm,
	}
}
