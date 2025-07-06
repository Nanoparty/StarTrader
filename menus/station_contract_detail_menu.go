package menus

import (
	"fmt"
	"startrader/globals"
	"startrader/types"
)

var selectedStationContract *types.Contract
var StationContractDetailMenu types.Menu

func StationContractDetailMenuIntro(m *types.Menu) {
	if selectedStationContract == nil {
		fmt.Println("\rNo contract selected.")
		return
	}
	moneyHeader := fmt.Sprintf("Credits: $%d", globals.Company.Money)
	fmt.Println("\r----------------------------------------------------------------------------")
	header := "Contract Details:"
	fmt.Printf("\r%s%*s%s\n", header, 76-len(header)-len(moneyHeader), "", moneyHeader)
	fmt.Println("\r----------------------------------------------------------------------------")
	fmt.Printf("\rName:        %s\n", selectedStationContract.ShortName)
	fmt.Printf("\rType:        %s\n", selectedStationContract.Type)
	fmt.Printf("\rDuration:    %d min %d sec\n", selectedStationContract.Minutes, selectedStationContract.Seconds)
	fmt.Printf("\rPayout:      $%d\n", selectedStationContract.Payout)
	fmt.Printf("\rDescription: %s\n", selectedStationContract.Description)
	fmt.Println("\r----------------------------------------------------------------------------")
}

func AcceptStationContract() {
	ShowAssignContractShipMenu()
}

func BackToStationContractMenu() {
	globals.CurrentMenu = &StationContractMenu
}

var StationContractDetailMenuOptions = []types.MenuItem{
	{Name: "Accept Contract", Callback: AcceptStationContract},
	{Name: "Back", Callback: BackToStationContractMenu},
}

func init() {
	StationContractDetailMenu = types.Menu{
		Name:    "Contract Detail",
		Intro:   StationContractDetailMenuIntro,
		Options: StationContractDetailMenuOptions,
		Back:    BackToStationContractMenu,
	}
}
