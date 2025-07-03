package menus

type Ship struct {
	Name string
	Price int
	Type string
	Storage int
	Speed int
	CurrentHealth int
	MaxHealth int
	Damage int
	AssignedPilot *Pilot
	Status string // "Idle", "In Progress", or "Complete"
}

type Pilot struct {
	Name string
	Price int
	TransportSkill int
	CombatSkill int
	MiningSkill int
	AssignedShip *Ship
}
