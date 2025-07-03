package menus

import "fmt"

var selectedDetailStation *Station

func StationDetailMenuIntro(m *Menu) {
	if selectedDetailStation == nil {
		fmt.Println("\rNo station selected.")
		return
	}
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Printf("\rStation: %s\n", selectedDetailStation.Name)
	fmt.Println("\r----------------------------------------------------------------------------")
}

func StationDetailShips() {
	// TODO: Implement station ships view
}

func StationDetailPilots() {
	// TODO: Implement station pilots view
}

func StationDetailMissions() {
	ShowStationMissionMenu()
}

func StationDetailInfo() {
	// TODO: Implement station information view
}

var StationDetailMenuOptions = []MenuItem{
	{Name: "Ships", Callback: StationDetailShips},
	{Name: "Pilots", Callback: StationDetailPilots},
	{Name: "Missions", Callback: StationDetailMissions},
	{Name: "Information", Callback: func() { CurrentMenu = &StationInformationMenu }},
	{Name: "Back", Callback: func() { CurrentMenu = &StationsMenu }},
}

var StationDetailMenu Menu

func init() {
	StationDetailMenu = Menu{
		Name:    "Station Detail",
		Intro:   StationDetailMenuIntro,
		Options: StationDetailMenuOptions,
		Back:    func() { CurrentMenu = &StationsMenu },
	}
}
