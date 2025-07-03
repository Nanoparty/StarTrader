package menus

import (
	"fmt"
)

var CancelMissionConfirmMenu Menu

func CancelMissionConfirmMenuIntro(m *Menu) {
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Println("\rAre you sure you want to cancel this mission?")
	fmt.Println("\r----------------------------------------------------------------------------")
	if selectedActiveMission != nil {
		fmt.Printf("\rMission: %s\n", selectedActiveMission.ShortName)
		fmt.Printf("\rAssigned Ship: %s\n", selectedActiveMissionShip.Name)
	}
	fmt.Println("\r----------------------------------------------------------------------------")
}

func CancelMissionConfirmYes() {
	CancelActiveMission()
}

func CancelMissionConfirmNo() {
	ActiveMissionDetailMenu.Options = BuildActiveMissionDetailMenuOptions()
	CurrentMenu = &ActiveMissionDetailMenu
}

func CancelMissionConfirmMenuOptions() []MenuItem {
	return []MenuItem{
		{Name: "Yes", Callback: CancelMissionConfirmYes},
		{Name: "No", Callback: CancelMissionConfirmNo},
	}
}

func init() {
	CancelMissionConfirmMenu = Menu{
		Name:    "Cancel Mission?",
		Intro:   CancelMissionConfirmMenuIntro,
		Options: CancelMissionConfirmMenuOptions(),
		Back:    CancelMissionConfirmNo,
	}
}
