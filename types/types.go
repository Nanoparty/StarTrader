package types

type Menu struct {
	Name		string
	Intro		func(m *Menu)
	Options		[]MenuItem
	Back		func()
	Selected	int
}

type MenuItem struct {
	Name		string
	Callback	func()
}

type Sector struct {
	Name     string
	Stations []Station
	IsKnown  bool
}

type Station struct {
	Sector             *Sector
	IsKnown            bool
	Name               string
	RelationshipLevel  int
	Experience         float64
	ExpToNextLevel     float64
	Missions           []Mission
	ShipsForSale       []Ship
	PilotsForSale      []Pilot
	MissionsCompleted  int
	MoneySpent         int
}

type Mission struct {
	ShortName   string
	Description string
	Type        string // "Mining", "Transport", or "Combat"
	Level       int    // 1-10, affects payout
	Payout      int    // Money earned for completion
	Minutes     int    // Duration minutes (0-4)
	Seconds     int    // Duration seconds (0-59)
	Status      string // "Idle", "In Progress", or "Complete"
	// Location   Location // Uncomment if Location is needed and defined
}

type Location struct {
	Sector  Sector
	Station Station
}

type Ship struct {
	Name            string
	Price           int
	Type            string
	Storage         int
	Speed           int
	CurrentHealth   int
	MaxHealth       int
	Damage          int
	Level           int // Starts at 1, max 10
	AssignedPilot   *Pilot
	Status          string // "Idle", "In Progress", or "Complete"
	AssignedMission *Mission // nil if not assigned
}

type Pilot struct {
	Name            string
	Price           int
	TransportSkill  int
	CombatSkill     int
	MiningSkill     int
	Level           int // 1 to 10
	AssignedShip    *Ship
	AssignedMission *Mission // nil if not assigned
	Status          string // "Idle", "In Progress", or "Complete"
}
