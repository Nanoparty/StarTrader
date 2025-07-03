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
