package utils

import (
	"math/rand"
	"startrader/types"
	"time"
)

func GenerateMiningContract(level int) types.Contract {
	if level < 1 {
		level = 1
	} else if level > 10 {
		level = 10
	}
	miningContracts := []struct{ Name, Desc string }{
		{"Asteroid Extraction", "Extract valuable minerals from a dense asteroid field."},
		{"Survey and Sample", "Survey a new asteroid and collect mineral samples."},
		{"Deep Core Drill", "Drill deep into a large asteroid for rare ores."},
		{"Ice Mining", "Harvest water ice from the outer belt."},
		{"Crystal Hunt", "Search for rare crystals in a hazardous zone."},
		{"Salvage Operation", "Recover minerals from a derelict mining ship."},
		{"Ore Transport", "Move freshly mined ore to a refinery."},
		{"Mine Expansion", "Clear debris for a new mining tunnel."},
		{"Explosive Charge", "Place charges for controlled asteroid fragmentation."},
		{"Hazardous Gas Extraction", "Extract volatile gases from a comet."},
		{"Rare Metal Rush", "Compete to mine a newly discovered rare metal vein."},
		{"Magnetite Survey", "Locate magnetite deposits for industrial use."},
		{"Drone Deployment", "Deploy mining drones to a remote field."},
		{"Cave Mapping", "Map out a complex asteroid cave system."},
		{"Radiation Sweep", "Mine in a high-radiation environment for bonuses."},
		{"Refinery Delivery", "Deliver refined minerals to a station."},
		{"Seismic Test", "Conduct seismic tests to locate ore pockets."},
		{"Microgravity Mining", "Operate mining equipment in microgravity."},
		{"Toxic Waste Cleanup", "Remove toxic mining byproducts safely."},
		{"Experimental Mining", "Test new mining tech in the field."},
	}
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(miningContracts))
	durationSec := rand.Intn(300) // 0 to 299 seconds
	minutes := durationSec / 60
	seconds := durationSec % 60
	payout := (minutes*60 + seconds) * 20 * level // payout scaled to seconds and level
	selected := miningContracts[idx]
	return types.Contract{
		ShortName:   selected.Name,
		Description: selected.Desc,
		Type:        "Mining",
		Level:       level,
		Payout:      payout,
		Minutes:     minutes,
		Seconds:     seconds,
		Status:      "Idle",
	}
}

func GenerateCombatContract(level int) types.Contract {
	if level < 1 {
		level = 1
	} else if level > 10 {
		level = 10
	}
	combatContracts := []struct{ Name, Desc string }{
		{"Pirate Hunt", "Track down and eliminate a pirate ship in the sector."},
		{"Convoy Escort", "Defend a civilian convoy from raiders."},
		{"Station Defense", "Help defend a station under attack."},
		{"Bounty Target", "Capture or destroy a wanted criminal vessel."},
		{"Border Skirmish", "Engage hostile forces at the sector border."},
		{"Hostage Rescue", "Rescue hostages from a hijacked ship."},
		{"Minefield Clearance", "Clear a dangerous minefield for safe passage."},
		{"Enemy Recon", "Scout enemy positions and report back."},
		{"Sabotage Run", "Sabotage enemy supplies on a covert contract."},
		{"Fleet Action", "Join a fleet battle against a rival faction."},
		{"Smuggler Intercept", "Intercept and disable a smuggler's ship."},
		{"Blockade Runner", "Break through an enemy blockade."},
		{"VIP Protection", "Protect a high-value target from assassination."},
		{"Prison Break", "Prevent a prison break at a detention facility."},
		{"Weapons Test", "Test new combat systems in live fire."},
		{"Rebel Suppression", "Suppress a rebel uprising in the system."},
		{"Ambush Response", "Respond to an ambush and protect allies."},
		{"Drone Swarm", "Defend against an attacking drone swarm."},
		{"Artifact Recovery", "Recover a stolen artifact from hostile forces."},
		{"Security Sweep", "Conduct a security sweep of a dangerous area."},
	}
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(combatContracts))
	durationSec := rand.Intn(300) // 0 to 299 seconds
	minutes := durationSec / 60
	seconds := durationSec % 60
	payout := (minutes*60 + seconds) * 20 * level // payout scaled to seconds and level
	selected := combatContracts[idx]
	return types.Contract{
		ShortName:   selected.Name,
		Description: selected.Desc,
		Type:        "Combat",
		Level:       level,
		Payout:      payout,
		Minutes:     minutes,
		Seconds:     seconds,
		Status:      "Idle",
	}
}

func GenerateRandomContract(level int) types.Contract {
	if level < 1 {
		level = 1
	} else if level > 10 {
		level = 10
	}
	typeIdx := rand.Intn(3)
	switch typeIdx {
	case 0:
		return GenerateMiningContract(level)
	case 1:
		return GenerateCombatContract(level)
	default:
		return GenerateTransportContract(level)
	}
}

func GenerateTransportContract(level int) types.Contract {
	if level < 1 {
		level = 1
	} else if level > 10 {
		level = 10
	}
	transportContracts := []struct{ Name, Desc string }{
		{"Medical Supply Run", "Deliver urgent medical supplies to a nearby outpost."},
		{"VIP Transport", "Escort a VIP passenger to another station safely."},
		{"Cargo Delivery", "Transport valuable cargo across the sector."},
		{"Diplomatic Pouch", "Deliver confidential diplomatic documents."},
		{"Food types.Shipment", "Deliver perishable food goods to a colony in need."},
		{"Rescue Pickup", "Pick up stranded miners and bring them to safety."},
		{"Mail Run", "Deliver important mail to several stations."},
		{"Science Equipment", "Transport sensitive science equipment for research."},
		{"Prisoner Transfer", "Safely transfer prisoners between stations."},
		{"Tourist Group", "Take a group of tourists on a scenic tour."},
		{"Weapon types.Shipment", "Deliver a shipment of weapons to a security post."},
		{"Rare Artifact", "Transport a rare artifact for a museum exhibit."},
		{"Fuel Run", "Deliver fuel to a stranded ship in deep space."},
		{"Medical Evacuation", "Evacuate injured personnel to a hospital station."},
		{"Mining Tools", "Deliver mining tools to a remote asteroid base."},
		{"VIP Entertainer", "Transport a famous entertainer to a show venue."},
		{"Data Courier", "Carry encrypted data to a secure facility."},
		{"Livestock types.Shipment", "Transport live animals for agricultural purposes."},
		{"Construction Materials", "Deliver building materials to a new colony."},
		{"Emergency Parts", "Rush critical ship parts to a repair crew."},
	}
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(transportContracts))
	durationSec := rand.Intn(300) // 0 to 299 seconds
	minutes := durationSec / 60
	seconds := durationSec % 60
	payout := (minutes*60 + seconds) * 20 * level // payout scaled to total seconds and level
	selected := transportContracts[idx]
	return types.Contract{
		ShortName:   selected.Name,
		Description: selected.Desc,
		Type:        "Transport",
		Level:       level,
		Payout:      payout,
		Minutes:     minutes,
		Seconds:     seconds,
		Status:      "Idle",
	}
}

// GenerateRandomContractList returns a slice of n random contracts at the given level.
func GenerateRandomContractList(n int, level int) []types.Contract {
	contracts := make([]types.Contract, 0, n)
	for i := 0; i < n; i++ {
		contracts = append(contracts, GenerateRandomContract(level))
	}
	return contracts
}
