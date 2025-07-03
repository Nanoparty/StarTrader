package menus

type Menu struct {
	Name		string
	Intro		func(m *Menu)
	Options		[]MenuItem
}

var CurrentMenu *Menu
var PreviousMenu []*Menu

//Add Current Menu to Previous Menu List
func AddPreviousMenu() {
	PreviousMenu = append(PreviousMenu, CurrentMenu)
}

//Pop Previous Menu
func GetPreviousMenu () *Menu {
	last := PreviousMenu[len(PreviousMenu)-1]
	PreviousMenu = PreviousMenu[:len(PreviousMenu)-1]
	return last
}

type MenuItem struct {
	Name		string
	Callback	func()
	Navigation	func()
}