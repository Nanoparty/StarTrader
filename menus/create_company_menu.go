package menus

import (
	"fmt"
	"startrader/globals"
	"startrader/types"
)

var companyName string

func GetCompanyName() string {
	return companyName
}

func AppendToCompanyName(char string) {
	companyName += char
}

func BackspaceCompanyName() {
	if len(companyName) > 0 {
		companyName = companyName[:len(companyName)-1]
	}
}

func CreateCompanyMenuIntro(m *types.Menu) {
	fmt.Println("\r", "----------------------------------------------------------------------------")
	fmt.Println("\r", "What would you like to name your company?")
	fmt.Println("\r", "----------------------------------------------------------------------------")
	fmt.Printf("\r> %s", companyName)
}

func CreateCompany() {
	globals.Company.Name = companyName
	globals.CurrentMenu = &CompanyMenu
}

var CreateCompanyMenu = types.Menu{
	Name:    "Create Company",
	Intro:   CreateCompanyMenuIntro,
	Options: []types.MenuItem{},
}

func init() {
	CreateCompanyMenu.Back = func() { globals.CurrentMenu = &MainMenu }
}
