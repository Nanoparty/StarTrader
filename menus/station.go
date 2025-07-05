package menus


type Station struct {
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

func NewStation(name string) *Station {
	return &Station{
		Name: name,
		RelationshipLevel: 1,
		Experience: 0.0,
		ExpToNextLevel: 100.0,
	}
}
