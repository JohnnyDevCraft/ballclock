package menu

import (
	"fmt"
)

//topMenu Methods

func (m *App) topMenu() {
	ClearScreen()
	m.renderTitle("Main Menu System")
	fmt.Println("1. Simulate Run and Review Results")
	fmt.Println("2. Change Settings")
	fmt.Println("3. Exit Application")
	val, input := m.getIntValue("Enter Choice: ")

	switch val {
	case 1:

	case 2:
		m.setupMenu()
	case 3:
		m.ExitApp = true
	default:
		m.showError(input)
	}

}
