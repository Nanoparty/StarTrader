package menus

import (
	"fmt"
	"startrader/globals"
	"startrader/types"
)

var Sectors = []*types.Sector{
	{
		Name: "Earth Sector",
		IsKnown: true,
		Stations: []types.Station{
			{Name: "Lunar Station", IsKnown: true},
			{Name: "Terra Station", IsKnown: true},
		},
	},
	{
		Name: "Mars Sector",
		IsKnown: false,
		Stations: []types.Station{
			{Name: "Mars Station", IsKnown: false},
			{Name: "Deimos Station", IsKnown: false},
			{Name: "Phobos Station", IsKnown: false},
		},
	},
	{
		Name: "Venus Sector",
		IsKnown: false,
		Stations: []types.Station{
			{Name: "Venus Station", IsKnown: false},
		},
	},
	{
		Name: "Jupiter Sector",
		IsKnown: false,
		Stations: []types.Station{
			{Name: "Jupitor Station", IsKnown: false},
			{Name: "Ganymede Station", IsKnown: false},
			{Name: "Europa Station", IsKnown: false},
			{Name: "Io Station", IsKnown: false},
		},
	},
	{
		Name: "Central Belt",
		IsKnown: false,
		Stations: []types.Station{
			{Name: "Ceres Station", IsKnown: false},
			{Name: "Pallas Station", IsKnown: false},
			{Name: "Vesta Station", IsKnown: false},
			{Name: "Hygiea Station", IsKnown: false},
		},
	},
	{
		Name: "Keiper Belt",
		IsKnown: false,
		Stations: []types.Station{
			{Name: "Pluto Outpost", IsKnown: false},
			{Name: "Voyager Station", IsKnown: false},
		},
	},
}

func SectorsMenuIntro(m *types.Menu) {
	moneyHeader := fmt.Sprintf("$%d", globals.Company.Money)
	fmt.Println("\r----------------------------------------------------------------------------")
	header := "Known Sectors:"
	fmt.Printf("\r%s%*s%s\n", header, 76-len(header)-len(moneyHeader), "", moneyHeader)
	fmt.Println("\r----------------------------------------------------------------------------")
}



var SectorsMenuOptions []types.MenuItem
var SectorsMenu types.Menu

func init() {
	SectorsMenuOptions = []types.MenuItem{}
	for _, sector := range Sectors {
		if sector.IsKnown {
			// capture sector in loop variable
			sec := sector
			SectorsMenuOptions = append(SectorsMenuOptions, types.MenuItem{
				Name:     sec.Name,
				Callback: func() { ShowStationsMenu(sec) },
			})
		}
	}
	SectorsMenuOptions = append(SectorsMenuOptions, types.MenuItem{Name: "Back", Callback: SectorsMenuBack})
	SectorsMenu = types.Menu{
		Name:    "Sectors Menu",
		Intro:   SectorsMenuIntro,
		Options: SectorsMenuOptions,
		Back:    SectorsMenuBack,
	}
}

func SectorsMenuBack() {
	globals.CurrentMenu = &CompanyMenu
}
