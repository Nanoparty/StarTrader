package menus


type Station struct {
	Name               string
	RelationshipLevel  int
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
	}
}
