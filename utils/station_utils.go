package utils

import "startrader/types"

func NewStation(name string) *types.Station {
	return &types.Station{
		Name: name,
		IsKnown: false,
		RelationshipLevel: 1,
		Experience: 0.0,
		ExpToNextLevel: 100.0,
		Missions:           []types.Mission{},
		ShipsForSale:       []types.Ship{},
		PilotsForSale:      []types.Pilot{},
		MissionsCompleted:  0,
		MoneySpent:         0,
	}
}