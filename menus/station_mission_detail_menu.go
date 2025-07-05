package menus

import (
	"fmt"
	"startrader/globals"
	"startrader/types"
)

var selectedStationMission *types.Mission
var StationMissionDetailMenu types.Menu

func StationMissionDetailMenuIntro(m *types.Menu) {
	if selectedStationMission == nil {
		fmt.Println("\rNo mission selected.")
		return
	}
	moneyHeader := fmt.Sprintf("$%d", globals.Company.Money)
	fmt.Println("\r----------------------------------------------------------------------------")
	header := "Mission Details:"
	fmt.Printf("\r%s%*s%s\n", header, 76-len(header)-len(moneyHeader), "", moneyHeader)
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
	globals.CurrentMenu = &StationMissionMenu
}

var StationMissionDetailMenuOptions = []types.MenuItem{
	{Name: "Accept Mission", Callback: AcceptStationMission},
	{Name: "Back", Callback: BackToStationMissionMenu},
}

func init() {
	StationMissionDetailMenu = types.Menu{
		Name:    "Mission Detail",
		Intro:   StationMissionDetailMenuIntro,
		Options: StationMissionDetailMenuOptions,
		Back:    BackToStationMissionMenu,
	}
}
