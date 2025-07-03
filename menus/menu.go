package menus

type Menu struct {
	Name		string
	Intro		func(m *Menu)
	Options		[]MenuItem
	Back		func()
	Selected	int
}

var CurrentMenu *Menu



type MenuItem struct {
	Name		string
	Callback	func()
}