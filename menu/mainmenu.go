package menu

import (
	"ballclock/models"
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Reader used for reading from consoled
var Reader *bufio.Reader

//InitReader initializes a reader.
func InitReader() {
	if Reader == nil {
		Reader = bufio.NewReader(os.Stdin)
	}
}

//Run starts the main menu
func (m *App) Run(clock *models.Clock) {
	InitReader()
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

	m.topMenu()

}

//ShowError renders an error in the Console
func ShowError(s string) {
	ClearScreen()
	RenderTitle("Error with Input! ****************")
	fmt.Println("I'm sorry, Im sure you had the best intentions, but I really need you")
	fmt.Printf("to stick to the options presented and [%s] just isn't a valid option.  Lets\n", s)
	fmt.Println("take another crack at it.")
	fmt.Println()
	fmt.Println("Press any key to continue.")
	Reader.ReadByte()
}

//RenderTitle is used to render the Title Section for Console.
func RenderTitle(title string) {
	fmt.Printf("_________________________________________________________________________________\n")
	fmt.Printf("BallClock Simulation 1.0                                       By:Johnathon Smith\n")
	fmt.Printf(title)
	fmt.Printf("\n_________________________________________________________________________________")
	fmt.Println()
	fmt.Println()
}

//GetInput returns a string from the console.
func GetInput() string {
	val, _ := Reader.ReadString('\n')
	val = strings.Replace(val, "\n", "", -1)
	return val
}

//RenderLine renders a choice line in the console.
func RenderLine(line int, val string) {
	fmt.Printf("%d. %s\n", line, val)
}

//GetIntValue returns an in and string from stdin.
func GetIntValue(promt string) (int, string) {
	fmt.Println(promt)
	val := GetInput()
	if StringIsInt(val) {
		return StringToInt(val), val
	}

	return 0, val
}
