package menus

import "fmt"

var KnownSectors = []string{
	"Earth Sector",
}

var UnknownSectors = []string{
	"Mars Sector",
	"Venus Sector",
	"Jupiter Sector",
	"Central Belt",
	"Keiper Belt",
}

func SectorsMenuIntro(m *Menu) {
	moneyHeader := fmt.Sprintf("$%d", CompanyMoney)
	fmt.Println("\r----------------------------------------------------------------------------")
	header := "Known Sectors:"
	fmt.Printf("\r%s%*s%s\n", header, 76-len(header)-len(moneyHeader), "", moneyHeader)
	fmt.Println("\r----------------------------------------------------------------------------")
}

func SectorSelected(sector string) func() {
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
		Back:    SectorsMenuBack,
	}
}

func SectorsMenuBack() {
	CurrentMenu = &CompanyMenu
}
