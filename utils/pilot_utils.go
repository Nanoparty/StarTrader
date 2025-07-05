package utils

import (
	"math/rand"
	"startrader/types"
)

// GenerateRandomPilot returns a random Pilot at the given level (1-10).
// Each pilot leans toward one skill (Transport, Combat, or Mining), which is significantly higher than the others.
// All skills scale with level, and price scales with total skill.
func GenerateRandomPilot(level int) types.Pilot {
	if level < 1 {
		level = 1
	} else if level > 10 {
		level = 10
	}
	name := Generate_Pilot_Name()
	dominant := rand.Intn(3) // 0 = Transport, 1 = Combat, 2 = Mining
	// Skill scaling: level 1 pilots max skill 5, higher levels scale up
	base := 1 + rand.Intn(2) + (level-1) // 1-2 for level 1, +1 per level
	domBonus := 1 + rand.Intn(2) + (level-1) // 1-2 for level 1, +1 per level
	otherSpread := 1 + rand.Intn(2) + (level-1) // 1-2 for level 1, +1 per level

	var transport, combat, mining int
	switch dominant {
	case 0: // Transport
		transport = base + domBonus
		combat = base + rand.Intn(otherSpread)
		mining = base + rand.Intn(otherSpread)
	case 1: // Combat
		combat = base + domBonus
		transport = base + rand.Intn(otherSpread)
		mining = base + rand.Intn(otherSpread)
	case 2: // Mining
		mining = base + domBonus
		transport = base + rand.Intn(otherSpread)
		combat = base + rand.Intn(otherSpread)
	}
	totalSkill := transport + combat + mining
	// Price scaling: level 1 pilots 2-5k, higher levels scale up
	var price int
	if level == 1 {
		price = 2000 + rand.Intn(3001) // 2000-5000
	} else {
		price = 2000 + (level-1)*3000 + totalSkill*1000 + rand.Intn(2001)
	}
	return types.Pilot{
		Name: name,
		Price: price,
		TransportSkill: transport,
		CombatSkill: combat,
		MiningSkill: mining,
		Level: level,
		AssignedShip: nil,
		AssignedMission: nil,
		Status: "Idle",
	}
}

// GenerateRandomPilotList returns a slice of n random pilots at the given level.
func GenerateRandomPilotList(n int, level int) []types.Pilot {
	pilots := make([]types.Pilot, 0, n)
	for i := 0; i < n; i++ {
		pilots = append(pilots, GenerateRandomPilot(level))
	}
	return pilots
}
