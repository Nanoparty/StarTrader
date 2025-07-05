package menus

import (
	"math/rand"
	"startrader/utils"
	"time"
)

var CompanyMoney int = 10000

type Ship struct {
	Name string
	Price int
	Type string
	Storage int
	Speed int
	CurrentHealth int
	MaxHealth int
	Damage int
	Level int // Starts at 1, max 10
	AssignedPilot *Pilot
	Status string // "Idle", "In Progress", or "Complete"
	AssignedMission *Mission // nil if not assigned
}

type Pilot struct {
	Name string
	Price int
	TransportSkill int
	CombatSkill int
	MiningSkill int
	Level int // 1 to 10
	AssignedShip *Ship
	AssignedMission *Mission // nil if not assigned
	Status string // "Idle", "In Progress", or "Complete"
}

// GenerateRandomPilot returns a random Pilot at the given level (1-10).
// Each pilot leans toward one skill (Transport, Combat, or Mining), which is significantly higher than the others.
// All skills scale with level, and price scales with total skill.
func GenerateRandomPilot(level int) Pilot {
	if level < 1 {
		level = 1
	} else if level > 10 {
		level = 10
	}
	name := utils.Generate_Pilot_Name()
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
	return Pilot{
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

// GenerateRandomMissionList returns a slice of n random missions at the given level.
// This function is now defined in mission.go to avoid dependency issues.
// Please use the GenerateRandomMissionList from mission.go.

// GenerateRandomShipList returns a slice of n random ships, all at the given level (e.g., the station's relationship level).
func GenerateRandomShipList(n int, level int) []Ship {
	ships := make([]Ship, 0, n)
	for i := 0; i < n; i++ {
		ships = append(ships, GenerateRandomShip(level))
	}
	return ships
}

// GenerateRandomPilotList returns a slice of n random pilots at the given level.
func GenerateRandomPilotList(n int, level int) []Pilot {
	pilots := make([]Pilot, 0, n)
	for i := 0; i < n; i++ {
		pilots = append(pilots, GenerateRandomPilot(level))
	}
	return pilots
}

// GenerateRandomShip returns a random Ship of type Combat, Transport, or Mining at the given level (1-10).
// Higher level ships have proportionally better stats.
// Combat ships have higher damage, transport ships have higher speed, mining ships have higher storage.
func GenerateRandomShip(level int) Ship {
	if level < 1 {
		level = 1
	} else if level > 10 {
		level = 10
	}
	rand.Seed(time.Now().UnixNano())
	types := []string{"Combat", "Transport", "Mining"}
	shipType := types[rand.Intn(len(types))]
	var name string
	var price, storage, speed, damage int
	var maxHealth int
	// Ship stats
	var avgStat float64
	switch shipType {
	case "Combat":
		name = utils.Generate_Combat_Ship_Name()
		storage = (10 + rand.Intn(6)) + (level-1)*2
		speed = (10 + rand.Intn(6)) + (level-1)*2
		damage = (15 + rand.Intn(11)) + (level-1)*4
		maxHealth = (10 + rand.Intn(11)) + (level-1)*3
		avgStat = float64(storage+speed+damage+maxHealth) / 4.0
	case "Transport":
		name = utils.Generate_Transport_Ship_Name()
		storage = (20 + rand.Intn(11)) + (level-1)*3
		speed = (18 + rand.Intn(8)) + (level-1)*3
		damage = (5 + rand.Intn(6)) + (level-1)*2
		maxHealth = (10 + rand.Intn(11)) + (level-1)*3
		avgStat = float64(storage+speed+damage+maxHealth) / 4.0
	case "Mining":
		name = utils.Generate_Mining_Ship_Name()
		storage = (30 + rand.Intn(16)) + (level-1)*4
		speed = (8 + rand.Intn(5)) + (level-1)*2
		damage = (7 + rand.Intn(5)) + (level-1)*2
		maxHealth = (10 + rand.Intn(11)) + (level-1)*3
		avgStat = float64(storage+speed+damage+maxHealth) / 4.0
	}

	// Price scaling: level 1 ships 5k-20k based on avgStat, higher levels scale up
	if level == 1 {
		minPrice := 5000.0
		maxPrice := 20000.0
		// Normalize avgStat for level 1 ships (roughly between 10 and 30)
		statNorm := (avgStat - 10.0) / 20.0
		if statNorm < 0.0 { statNorm = 0.0 }
		if statNorm > 1.0 { statNorm = 1.0 }
		price = int(minPrice + statNorm*(maxPrice-minPrice)) + rand.Intn(1001)
	} else {
		// For higher levels, scale price up with level and avgStat
		base := 20000 + (level-1)*12000
		price = base + int(avgStat*float64(level)*400) + rand.Intn(2001)
	}

	return Ship{
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
		AssignedMission: nil,
	}
}

