package menu

import (
	"fmt"
)

//topMenu Methods

func (m *App) topMenu() {
	ClearScreen()
	RenderTitle("Main Menu System")
	fmt.Println("1. Mode 1 Simulation (Return Days until Reset)")
	fmt.Println("2. Mode 2 Simulation (Return positions after run.)")
	fmt.Println("3. Exit Application")
	val, input := GetIntValue("Enter Choice: ")

	switch val {
	case 1:
		SimModeOneStep1()
	case 2:
		SimModeTwoStep1()
	case 3:
		m.ExitApp = true
	default:
		ShowError(input)
	}

}
