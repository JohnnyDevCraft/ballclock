package menu

import (
	"ballclock/models"
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Run starts the main menu
func (m *App) Run(clock *models.Clock) {
	m.Reader = bufio.NewReader(os.Stdin)
	m.Clock = clock

	ClearScreen()

	m.ExitApp = false
	m.SetupComplete = false

	for !m.ExitApp {
		ClearScreen()

		m.mainMenu()
	}
}

func (m *App) mainMenu() {
	ClearScreen()

	if !m.SetupComplete {
		m.setup()
	}

	m.topMenu()

}

func (m *App) showError(s string) {
	ClearScreen()
	m.renderTitle("Error with Input! ****************")
	fmt.Println("I'm sorry, Im sure you had the best intentions, but I really need you")
	fmt.Printf("to stick to the options presented and [%s] just isn't a valid option.  Lets\n", s)
	fmt.Println("take another crack at it.")
	fmt.Println()
	fmt.Println("Press any key to continue.")
	m.Reader.ReadByte()
}

func (m *App) renderTitle(title string) {
	fmt.Printf("_________________________________________________________________________________\n")
	fmt.Printf("BallClock Simulation 1.0                                       By:Johnathon Smith\n")
	fmt.Printf(title)
	fmt.Printf("\n_________________________________________________________________________________")
	fmt.Println()
	fmt.Println()
}

func (m *App) getInput() string {
	val, _ := m.Reader.ReadString('\n')
	val = strings.Replace(val, "\n", "", -1)
	return val
}

func (m *App) renderLine(line int, val string) {
	fmt.Printf("%d. %s\n", line, val)
}

func (m *App) getIntValue(promt string) (int, string) {
	fmt.Println(promt)
	val := m.getInput()
	if StringIsInt(val) {
		return StringToInt(val), val
	} else {
		return 0, val
	}
}
