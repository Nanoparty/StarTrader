package menus

import (
	"fmt"
	"startrader/globals"
	"startrader/types"
	"startrader/utils"
)

var Sectors = []*types.Sector{
	{
		Name: "Earth Sector",
		IsKnown: true,
		Stations: []types.Station{
			*utils.NewStation("Lunar Station", true),
			*utils.NewStation("Terra Station", true),
		},
	},
	{
		Name: "Mars Sector",
		IsKnown: false,
		Stations: []types.Station{
			*utils.NewStation("Mars Station", false),
			*utils.NewStation("Deimos Station", false),
			*utils.NewStation("Phobos Station", false),
		},
	},
	{
		Name: "Venus Sector",
		IsKnown: false,
		Stations: []types.Station{
			*utils.NewStation("Venus Station", false),
		},
	},
	{
		Name: "Jupiter Sector",
		IsKnown: false,
		Stations: []types.Station{
			*utils.NewStation("Jupitor Station", false),
			*utils.NewStation("Ganymede Station", false),
			*utils.NewStation("Europa Station", false),
			*utils.NewStation("Io Station", false),
		},
	},
	{
		Name: "Central Belt",
		IsKnown: false,
		Stations: []types.Station{
			*utils.NewStation("Ceres Station", false),
			*utils.NewStation("Pallas Station", false),
			*utils.NewStation("Vesta Station", false),
			*utils.NewStation("Hygiea Station", false),
		},
	},
	{
		Name: "Keiper Belt",
		IsKnown: false,
		Stations: []types.Station{
			*utils.NewStation("Pluto Outpost", false),
			*utils.NewStation("Voyager Station", false),
		},
	},
}

func SectorsMenuIntro(m *types.Menu) {
	moneyHeader := fmt.Sprintf("Credits: $%d", globals.Company.Money)
	fmt.Println("\r----------------------------------------------------------------------------")
	header := "Known Sectors:"
	fmt.Printf("\r%s%*s%s\n", header, 76-len(header)-len(moneyHeader), "", moneyHeader)
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Printf("\r%-30s | %s\n", "Sector", "Known Stations")
	fmt.Println("\r----------------------------------------------------------------------------")
}



var SectorsMenuOptions []types.MenuItem
var SectorsMenu types.Menu

func init() {
	SectorsMenuOptions = []types.MenuItem{}
	for _, sector := range Sectors {
		if sector.IsKnown {
			knownStations := 0
			for _, station := range sector.Stations {
				if station.IsKnown {
					knownStations++
				}
			}
			// capture sector in loop variable
			sec := sector
			SectorsMenuOptions = append(SectorsMenuOptions, types.MenuItem{
				Name:     fmt.Sprintf("%-30s | %d/%d", sec.Name, knownStations, len(sec.Stations)),
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
