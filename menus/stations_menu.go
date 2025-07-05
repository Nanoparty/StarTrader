package menus

import "fmt"

// Map of sector name to known stations
var KnownStations = map[string][]string{
	"Earth Sector": {"Lunar Station", "Terra Station"},
}

var UnknownStations = map[string][]string{
	"Mars Sector": {"Mars Station", "Deimos Station", "Phobos Station"},
	"Venus Sector": {"Venus Station"},
	"Jupiter Sector": {"Jupitor Station", "Ganymede Station", "Europa Station", "Io Station"},
	"Central Belt": {"Ceres Station", "Pallas Station", "Vesta Station", "Hygiea Station"},
	"Keiper Belt": {"Pluto Outpost", "Voyager Station"},
}

var SelectedSector string

func StationsMenuIntro(m *Menu) {
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Printf("\rStations in %s:\n", SelectedSector)
	fmt.Println("\r----------------------------------------------------------------------------")
}

func StationSelected(station *Station) func() {
	return func() {
		selectedDetailStation = station
		CurrentMenu = &StationDetailMenu
	}
}

var StationsMenuOptions []MenuItem
var StationsMenu Menu
var StationsByName = make(map[string]*Station)

func ShowStationsMenu(sector string) {
	SelectedSector = sector
	StationsMenuOptions = []MenuItem{}
	for _, stationName := range KnownStations[sector] {
		var stationPtr *Station
		if existing, ok := StationsByName[stationName]; ok {
			stationPtr = existing
		} else {
			station := NewStation(stationName)
			StationsByName[stationName] = station
			stationPtr = station
		}
		StationsMenuOptions = append(StationsMenuOptions, MenuItem{
			Name:     stationName,
			Callback: StationSelected(stationPtr),
		})
	}
	StationsMenuOptions = append(StationsMenuOptions, MenuItem{Name: "Back", Callback: StationsMenuBack})
	StationsMenu = Menu{
		Name:    "Stations Menu",
		Intro:   StationsMenuIntro,
		Options: StationsMenuOptions,
		Back:    StationsMenuBack,
	}
	CurrentMenu = &StationsMenu
}

func StationsMenuBack() {
	CurrentMenu = &SectorsMenu
}
