package menus

import (
	"fmt"
)

var selectedDetailStation *Station

func StationDetailMenuIntro(m *Menu) {
	if selectedDetailStation == nil {
		fmt.Println("\rNo station selected.")
		return
	}
	// Generate missions if this is the first visit
	if len(selectedDetailStation.Missions) == 0 {
		selectedDetailStation.Missions = GenerateRandomMissionList(6, selectedDetailStation.RelationshipLevel)
	}
	// Generate ships for sale if this is the first visit
	if len(selectedDetailStation.ShipsForSale) == 0 {
		selectedDetailStation.ShipsForSale = GenerateRandomShipList(5, selectedDetailStation.RelationshipLevel)
	}
	// Generate pilots for sale if this is the first visit
	if len(selectedDetailStation.PilotsForSale) == 0 {
		selectedDetailStation.PilotsForSale = GenerateRandomPilotList(5, selectedDetailStation.RelationshipLevel)
	}
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Printf("\rStation: %s\n", selectedDetailStation.Name)
	fmt.Println("\r----------------------------------------------------------------------------")
}

func StationDetailShips() {
	ShowStationShipsStoreMenu()
}

func StationDetailPilots() {
	ShowStationPilotsStoreMenu()
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
