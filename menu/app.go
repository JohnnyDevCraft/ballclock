package menu

import (
	"ballclock/models"
	"bufio"
)

//App stores app setings
type App struct {
	Interval      int
	BallCount     int
	Reader        *bufio.Reader
	MinBalls      int
	Clock         *models.Clock
	ExitApp       bool
	SetupComplete bool
}
