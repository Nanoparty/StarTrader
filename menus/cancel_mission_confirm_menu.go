package menus

import (
	"fmt"
	"startrader/types"
	"startrader/globals"
)

var CancelMissionConfirmMenu types.Menu

func CancelMissionConfirmMenuIntro(m *types.Menu) {
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
	globals.CurrentMenu = &ActiveMissionDetailMenu
}

func CancelMissionConfirmMenuOptions() []types.MenuItem {
	return []types.MenuItem{
		{Name: "Yes", Callback: CancelMissionConfirmYes},
		{Name: "No", Callback: CancelMissionConfirmNo},
	}
}

func init() {
	CancelMissionConfirmMenu = types.Menu{
		Name:    "Cancel Mission?",
		Intro:   CancelMissionConfirmMenuIntro,
		Options: CancelMissionConfirmMenuOptions(),
		Back:    CancelMissionConfirmNo,
	}
}
