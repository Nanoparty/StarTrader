package utils

import (
	"math/rand"
	"time"
	"startrader/types"
)

// GenerateRandomShip returns a random Ship of type Combat, Transport, or Mining at the given level (1-10).
// Higher level ships have proportionally better stats.
// Combat ships have higher damage, transport ships have higher speed, mining ships have higher storage.
func GenerateRandomShip(level int) types.Ship {
	if level < 1 {
		level = 1
	} else if level > 10 {
		level = 10
	}
	rand.Seed(time.Now().UnixNano())
	shipTypes := []string{"Combat", "Transport", "Mining"}
	shipType := shipTypes[rand.Intn(len(shipTypes))]
	var name string
	var price, storage, speed, damage int
	var maxHealth int
	// Ship stats
	if shipType == "Combat" {
		name = Generate_Combat_Ship_Name()
		storage = 10 + rand.Intn(11) + level*2
		speed = 10 + rand.Intn(6) + level*2
		damage = 20 + rand.Intn(11) + level*3
		maxHealth = 100 + rand.Intn(51) + level*10
	} else if shipType == "Transport" {
		name = Generate_Transport_Ship_Name()
		storage = 40 + rand.Intn(21) + level*4
		speed = 18 + rand.Intn(8) + level*2
		damage = 8 + rand.Intn(7) + level*1
		maxHealth = 80 + rand.Intn(31) + level*8
	} else { // Mining
		name = Generate_Mining_Ship_Name()
		storage = 60 + rand.Intn(21) + level*5
		speed = 8 + rand.Intn(6) + level*1
		damage = 5 + rand.Intn(6) + level*1
		maxHealth = 70 + rand.Intn(21) + level*7
	}
	// Price scaling: level 1 ships 10-20k, higher levels scale up
	avgStat := float64(storage+speed+damage+maxHealth) / 4.0
	if level == 1 {
		minPrice := 10000.0
		maxPrice := 20000.0
		statNorm := (avgStat - 40.0) / 60.0
		if statNorm < 0.0 { statNorm = 0.0 }
		if statNorm > 1.0 { statNorm = 1.0 }
		price = int(minPrice + statNorm*(maxPrice-minPrice)) + rand.Intn(1001)
	} else {
		base := 20000 + (level-1)*12000
		price = base + int(avgStat*float64(level)*400) + rand.Intn(2001)
	}
	return types.Ship{
		Name:          name,
		Price:         price,
		Type:          shipType,
		Storage:       storage,
		Speed:         speed,
		CurrentHealth: maxHealth,
		MaxHealth:     maxHealth,
		Damage:        damage,
		Level:         level,
		AssignedPilot: nil,
		Status:        "Idle",
		AssignedContract: nil,
	}
}

// GenerateRandomShipList returns a slice of n random ships, all at the given level (e.g., the station's relationship level).
func GenerateRandomShipList(n int, level int) []types.Ship {
	ships := make([]types.Ship, 0, n)
	for i := 0; i < n; i++ {
		ships = append(ships, GenerateRandomShip(level))
	}
	return ships
}

