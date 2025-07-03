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
	base := 2 + rand.Intn(3) + level // base 2-4, +level
	domBonus := 4 + rand.Intn(3) + (level-1)*2 // dominant skill bonus: 4-6 + scaling
	otherSpread := 2 + rand.Intn(3) + (level-1) // 2-4 + scaling

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
	price := 20000 + totalSkill*4000 + level*3000 + rand.Intn(5000) // price scales with skills and level
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
	basePrice := 50000

	switch shipType {
	case "Combat":
		name = utils.Generate_Combat_Ship_Name()
		storage = (10 + rand.Intn(6)) + (level-1)*2   // base 10-15, +2 per level above 1
		speed = (10 + rand.Intn(6)) + (level-1)*2     // base 10-15, +2 per level
		damage = (15 + rand.Intn(11)) + (level-1)*4   // base 15-25, +4 per level
		maxHealth = (10 + rand.Intn(11)) + (level-1)*3 // base 10-20, +3 per level
		price = basePrice + 50000 + damage*2000 + level*8000
	case "Transport":
		name = utils.Generate_Transport_Ship_Name()
		storage = (20 + rand.Intn(11)) + (level-1)*3  // base 20-30, +3 per level
		speed = (18 + rand.Intn(8)) + (level-1)*3     // base 18-25, +3 per level
		damage = (5 + rand.Intn(6)) + (level-1)*2     // base 5-10, +2 per level
		maxHealth = (10 + rand.Intn(11)) + (level-1)*3
		price = basePrice + 30000 + speed*1500 + level*6000
	case "Mining":
		name = utils.Generate_Mining_Ship_Name()
		storage = (30 + rand.Intn(16)) + (level-1)*4  // base 30-45, +4 per level
		speed = (8 + rand.Intn(5)) + (level-1)*2      // base 8-12, +2 per level
		damage = (7 + rand.Intn(5)) + (level-1)*2     // base 7-11, +2 per level
		maxHealth = (10 + rand.Intn(11)) + (level-1)*3
		price = basePrice + 20000 + storage*1200 + level*5000
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

