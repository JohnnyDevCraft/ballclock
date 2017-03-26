package menu

import (
	"fmt"
)

//settings Menu Methods

func (m *App) setupMenu() {

	moveOn := false

	for !moveOn {
		ClearScreen()
		m.renderTitle("Settings Menu")
		m.renderLine(1, "Change Interval for Clock, and reinitialize")
		m.renderLine(2, "Change ball count, and reinitialize.")
		m.renderLine(3, "Return to main menu.")
		val, input := m.getIntValue("Enter Selection: ")

		switch val {
		case 1:
			m.setupMenuChangeInterval()
		case 2:
			m.setupMenuChangeBallCount()
		case 3:
			moveOn = true
		default:
			m.showError(input)
		}

	}

}

func (m *App) setupMenuChangeInterval() {
	moveOn := false

	for !moveOn {
		ClearScreen()
		m.renderTitle("Settings Menu >>> Change Interval")
		m.renderLine(1, "5 Minute Interval")
		m.renderLine(2, "10 Minute Interval")
		m.renderLine(3, "Cancel and Go Back")
		val, input := m.getIntValue("Enter Selection: ")

		switch val {
		case 1:
			m.Interval = 5
			m.MinBalls = 29
			moveOn = true
		case 2:
			m.Interval = 10
			m.MinBalls = 28
			moveOn = true
		case 3:
			moveOn = true
		default:
			m.showError(input)

		}
	}
}

func (m *App) setupMenuChangeBallCount() {
	moveOn := false

	for !moveOn {
		ClearScreen()
		m.renderTitle("Settings Menu >>> Change Ball Count")

		val, input := m.getIntValue(fmt.Sprintf("Enter a new ball count between %d and 90.  Enter c or C to cancel: ", m.MinBalls))

		switch {
		case val < m.MinBalls && val != 0:
			m.showError(input)
		case val > 90:
			m.showError(input)
		default:
			if input == "c" || input == "C" {
				moveOn = true
			} else {

				if val != 0 {
					m.BallCount = val
				}

				m.showError(input)
			}

		}
	}
}
