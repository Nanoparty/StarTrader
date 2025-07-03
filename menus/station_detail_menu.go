package menus

import (
	"fmt"
	"startrader/utils"
)

var selectedDetailStation *Station

func StationDetailMenuIntro(m *Menu) {
	if selectedDetailStation == nil {
		fmt.Println("\rNo station selected.")
		return
	}
	// Generate missions if this is the first visit
	if len(selectedDetailStation.Missions) == 0 {
		for i := 0; i < 6; i++ {
			selectedDetailStation.Missions = append(selectedDetailStation.Missions, GenerateRandomMission())
		}
	}
	// Generate ships for sale if this is the first visit
	if len(selectedDetailStation.ShipsForSale) == 0 {
		selectedDetailStation.ShipsForSale = []Ship{
			{utils.Generate_Combat_Ship_Name(), 100000, "Combat", 10, 10, 10, 10, 10, nil, "Idle", nil},
			{utils.Generate_Combat_Ship_Name(), 250000, "Combat", 10, 10, 10, 10, 10, nil, "Idle", nil},
			{utils.Generate_Combat_Ship_Name(), 400000, "Combat", 10, 10, 10, 10, 10, nil, "Idle", nil},
			{utils.Generate_Combat_Ship_Name(), 750000, "Combat", 10, 10, 10, 10, 10, nil, "Idle", nil},
			{utils.Generate_Combat_Ship_Name(), 1200000, "Combat", 10, 10, 10, 10, 10, nil, "Idle", nil},
		}
	}
	// Generate pilots for sale if this is the first visit
	if len(selectedDetailStation.PilotsForSale) == 0 {
		selectedDetailStation.PilotsForSale = []Pilot{
			{utils.Generate_Pilot_Name(), 50000, 8, 9, 2, nil, nil, "Idle"},
			{utils.Generate_Pilot_Name(), 40000, 7, 6, 5, nil, nil, "Idle"},
			{utils.Generate_Pilot_Name(), 30000, 5, 4, 7, nil, nil, "Idle"},
			{utils.Generate_Pilot_Name(), 60000, 10, 4, 3, nil, nil, "Idle"},
			{utils.Generate_Pilot_Name(), 55000, 6, 8, 6, nil, nil, "Idle"},
		}
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
