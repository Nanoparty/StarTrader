package globals

import "startrader/types"

var CurrentMenu *types.Menu

type Config struct {
	SaveFile string
	
}

type CompanyConfig struct {
	Name string
	Money int
	Ships []types.Ship
	Pilots []types.Pilot
}

var Company = CompanyConfig {
	Name: "Default Company",
	Money: 20000000,
	Ships: []types.Ship{},
	Pilots: []types.Pilot{},
}
