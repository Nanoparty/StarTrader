package menus


type Station struct {
	Name               string
	Ships              []*Ship
	Employees          []*Pilot
	RelationshipLevel  int
	MissionsCompleted  int
	MoneySpent         int
	// Add other properties as needed (location, owner, etc)
}
