package utils

import "startrader/types"

func NewStation(name string, isKnown bool) *types.Station {
	return &types.Station{
		Name: name,
		IsKnown: isKnown,
		RelationshipLevel: 1,
		Experience: 0.0,
		ExpToNextLevel: 100.0,
		Contracts:           []types.Contract{},
		ShipsForSale:       []types.Ship{},
		PilotsForSale:      []types.Pilot{},
		ContractsCompleted:  0,
		MoneySpent:         0,
	}
}