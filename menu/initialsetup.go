package menu

import (
	"fmt"
	"strings"
)

//InitialSetup Methods

func (m *App) setup() {
	m.setupIntro()
	ClearScreen()
	m.renderTitle("Welcome to John Smith's virtual BallClock Emulator>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Println("Well then, see, now that wasn't so bad.  Thanks for taking the")
	fmt.Println("time to start this simulation off on the right foot.  Use the ")
	fmt.Println("following menu to proceed. ")
	fmt.Println("Please press any key to continue.")
	m.Reader.ReadByte()
	m.SetupComplete = true
}

func (m *App) setupIntro() {
	m.renderTitle("Welcome to John Smith's virtual BallClock Emulator>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Println("This applicatoin is designed to run a BallClock simulation")
	fmt.Println("that allows you (the user) to track balls through the system")
	fmt.Println("over a specified amount of time.  It also provides a live ")
	fmt.Println("time mode that updates a command line version of the clock")
	fmt.Println("in real time.  In order to use this simulation I need to ")
	fmt.Println("collect some general information from you.  Please press any")
	fmt.Println("key to continue.")
	m.Reader.ReadByte()
	m.setupInterval()
}

func (m *App) setupInterval() {
	moveOn := false
	for !moveOn {
		ClearScreen()
		m.renderTitle("Welcome to John Smith's virtual BallClock Emulator>>>>>>>>>>>>>>>>>>>>>>>>")
		fmt.Println("Ah yes, you see, a BallClock uses levels to seperate ")
		fmt.Println("diffrent spans of time.  The top level usaully measures")
		fmt.Println("individual minutes.  The second level measures an  ")
		fmt.Println("interval of minutes. This simulator supports either")
		fmt.Println("5 minute intervals, or 10 minute intervals.  Please   ")
		fmt.Println("select the interval you wish to use.")
		fmt.Println()
		fmt.Println("1. 5 Minute Intervals")
		fmt.Println("2. 10 Minute intervals")
		val, _ := m.Reader.ReadString('\n')
		val = strings.Replace(val, "\n", "", -1)
		switch val {
		case "1":
			m.Interval = 5
			m.MinBalls = 29
			moveOn = true
		case "2":
			m.Interval = 10
			m.MinBalls = 28
			moveOn = true
		default:
			m.showError(val)
		}

	}
	m.setupBalls()
}

func (m *App) setupBalls() {
	moveOn := false
	for !moveOn {
		ClearScreen()
		m.renderTitle("Welcome to John Smith's virtual BallClock Emulator>>>>>>>>>>>>>>>>>>>>>>>>")
		fmt.Println("Wonderful choice ;). Now we need to make another important")
		fmt.Println("decision.  Because BallClocks run on balls, we need to")
		fmt.Println("add some balls to our clock to make it work.  Each ball")
		fmt.Println("we add will be given an id to allow us to track it through")
		fmt.Println("the system.  In order for the system to work, we need a ")
		fmt.Printf("minimum of %d balls.  You can use up to 90 balls in this", m.MinBalls)
		fmt.Println("this simulation.")
		fmt.Println()
		fmt.Println("How many balls should we use?: ")
		val, _ := m.Reader.ReadString('\n')
		val = strings.Replace(val, "\n", "", -1)

		if StringIsInt(val) {
			switch {
			case StringToInt(val) < m.MinBalls:
				m.showError(val)
			case StringToInt(val) > 90:
				m.showError(val)
			default:
				m.BallCount = StringToInt(val)
				moveOn = true
			}
		} else {
			m.showError(val)
		}

	}
}
