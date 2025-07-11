package menus

import (
	"fmt"
	"startrader/globals"
	"startrader/types"
)

var SelectedSector *types.Sector

func StationsMenuIntro(m *types.Menu) {
	moneyHeader := fmt.Sprintf("Credits: $%d", globals.Company.Money)
	fmt.Println("\r----------------------------------------------------------------------------")
	header := fmt.Sprintf("Stations in %s:", SelectedSector.Name)
	fmt.Printf("\r%s%*s%s\n", header, 76-len(header)-len(moneyHeader), "", moneyHeader)
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Printf("\r%-30s | %s\n", "Station", "Relationship Level")
	fmt.Println("\r----------------------------------------------------------------------------")
}

func StationSelected(station *types.Station) func() {
	return func() {
		selectedDetailStation = station
		globals.CurrentMenu = &StationDetailMenu
	}
}

var StationsMenuOptions []types.MenuItem
var StationsMenu types.Menu
var StationsByName = make(map[string]*types.Station)

func ShowStationsMenu(sector *types.Sector) {
	SelectedSector = sector
	StationsMenuOptions = []types.MenuItem{}
	for i := range sector.Stations {
		station := &sector.Stations[i]
		if station.IsKnown {
			StationsMenuOptions = append(StationsMenuOptions, types.MenuItem{
				Name:     fmt.Sprintf("%-30s | %d", station.Name, station.RelationshipLevel),
				Callback: StationSelected(station),
			})
		}
	}
	StationsMenuOptions = append(StationsMenuOptions, types.MenuItem{Name: "Back", Callback: StationsMenuBack})
	StationsMenu = types.Menu{
		Name:    "types.Stations Menu",
		Intro:   StationsMenuIntro,
		Options: StationsMenuOptions,
		Back:    StationsMenuBack,
	}
	globals.CurrentMenu = &StationsMenu
}

func StationsMenuBack() {
	globals.CurrentMenu = &SectorsMenu
}
