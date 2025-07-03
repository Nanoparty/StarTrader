package menus

import "fmt"

// Map of sector name to known stations
var KnownStations = map[string][]string{
	"Earth Sector": {"Lunar Station", "Terra Station"},
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

func StationSelected(station Station) func() {
	return func() {
		selectedDetailStation = &station
		CurrentMenu = &StationDetailMenu
	}
}

var StationsMenuOptions []MenuItem
var StationsMenu Menu

func ShowStationsMenu(sector string) {
	SelectedSector = sector
	StationsMenuOptions = []MenuItem{}
	for _, station := range KnownStations[sector] {
		stationCopy := station // avoid closure capture bug
		stationObj := Station{Name: stationCopy}
		StationsMenuOptions = append(StationsMenuOptions, MenuItem{
			Name:     stationCopy,
			Callback: StationSelected(stationObj),
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
