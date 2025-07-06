package menus

import (
	"fmt"
	"startrader/globals"
	"startrader/types"
	"startrader/utils"
)

var selectedDetailStation *types.Station

func StationDetailMenuIntro(m *types.Menu) {
	if selectedDetailStation == nil {
		fmt.Println("\rNo station selected.")
		return
	}
	// Generate contracts if this is the first visit
	if len(selectedDetailStation.Contracts) == 0 {
		selectedDetailStation.Contracts = utils.GenerateRandomContractList(6, selectedDetailStation.RelationshipLevel, selectedDetailStation)
	}
	// Generate ships for sale if this is the first visit
	if len(selectedDetailStation.ShipsForSale) == 0 {
		selectedDetailStation.ShipsForSale = utils.GenerateRandomShipList(5, selectedDetailStation.RelationshipLevel)
	}
	// Generate pilots for sale if this is the first visit
	if len(selectedDetailStation.PilotsForSale) == 0 {
		selectedDetailStation.PilotsForSale = utils.GenerateRandomPilotList(5, selectedDetailStation.RelationshipLevel)
	}
	moneyHeader := fmt.Sprintf("Credits: $%d", globals.Company.Money)
	fmt.Println("\r----------------------------------------------------------------------------")
	header := fmt.Sprintf("Station: %s", selectedDetailStation.Name)
	fmt.Printf("\r%s%*s%s\n", header, 76-len(header)-len(moneyHeader), "", moneyHeader)
	fmt.Println("\r----------------------------------------------------------------------------")
}

func StationDetailShips() {
	ShowStationShipsStoreMenu()
}

func StationDetailPilots() {
	ShowStationPilotsStoreMenu()
}

func StationDetailContracts() {
	ShowStationContractMenu()
}

func StationDetailInfo() {
	// TODO: Implement station information view
}

var StationDetailMenuOptions = []types.MenuItem{
	{Name: "Ships", Callback: StationDetailShips},
	{Name: "Pilots", Callback: StationDetailPilots},
	{Name: "Contracts", Callback: StationDetailContracts},
	{Name: "Information", Callback: func() { globals.CurrentMenu = &StationInformationMenu }},
	{Name: "Back", Callback: func() { globals.CurrentMenu = &StationsMenu }},
}

var StationDetailMenu types.Menu

func init() {
	StationDetailMenu = types.Menu{
		Name:    "Station Detail",
		Intro:   StationDetailMenuIntro,
		Options: StationDetailMenuOptions,
		Back:    func() { globals.CurrentMenu = &StationsMenu },
	}
}
