package menus

import (
	"fmt"
)

var selectedStationMission *Mission
var StationMissionDetailMenu Menu

func StationMissionDetailMenuIntro(m *Menu) {
	if selectedStationMission == nil {
		fmt.Println("\rNo mission selected.")
		return
	}
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Println("\rMission Details:")
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Printf("\rName:        %s\n", selectedStationMission.ShortName)
	fmt.Printf("\rType:        %s\n", selectedStationMission.Type)
	fmt.Printf("\rDuration:    %d min %d sec\n", selectedStationMission.Minutes, selectedStationMission.Seconds)
	fmt.Printf("\rPayout:      $%d\n", selectedStationMission.Payout)
	fmt.Printf("\rDescription: %s\n", selectedStationMission.Description)
	fmt.Println("\r----------------------------------------------------------------------------")
}

func AcceptStationMission() {
	ShowAssignMissionShipMenu()
}

func BackToStationMissionMenu() {
	CurrentMenu = &StationMissionMenu
}

var StationMissionDetailMenuOptions = []MenuItem{
	{Name: "Accept Mission", Callback: AcceptStationMission},
	{Name: "Back", Callback: BackToStationMissionMenu},
}

func init() {
	StationMissionDetailMenu = Menu{
		Name:    "Mission Detail",
		Intro:   StationMissionDetailMenuIntro,
		Options: StationMissionDetailMenuOptions,
		Back:    BackToStationMissionMenu,
	}
}
