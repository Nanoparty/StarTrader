package menus

import "fmt"

var KnownSectors = []string{
	"Earth Sector",
	"Mars Sector",
	"Venus Sector",
	"Jupiter Sector",
	"Central Belt",
	"Keiper Belt",
}

func SectorsMenuIntro(m *Menu) {
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Println("\rKnown Sectors:")
	fmt.Println("\r----------------------------------------------------------------------------")
}

func SectorSelected(sector string) func() {
	AddPreviousMenu()
	return func() {
		ShowStationsMenu(sector)
	}
}

var SectorsMenuOptions []MenuItem
var SectorsMenu Menu

func init() {
	SectorsMenuOptions = []MenuItem{}
	for _, sector := range KnownSectors {
		sectorCopy := sector // avoid closure capture bug
		SectorsMenuOptions = append(SectorsMenuOptions, MenuItem{
			Name:     sectorCopy,
			Callback: SectorSelected(sectorCopy),
		})
	}
	SectorsMenuOptions = append(SectorsMenuOptions, MenuItem{Name: "Back", Callback: SectorsMenuBack})
	SectorsMenu = Menu{
		Name:    "Sectors Menu",
		Intro:   SectorsMenuIntro,
		Options: SectorsMenuOptions,
	}
}

func SectorsMenuBack() {
	CurrentMenu = GetPreviousMenu()
}
