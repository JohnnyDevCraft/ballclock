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

type ReturnValue struct {
	Min     []int
	FiveMin []int
	Hour    []int
	Main    []int
}

func (rv *ReturnValue) InitMin(bs models.BallStack) {
	for _, v := range bs.Balls {
		rv.Min = append(rv.Min, v.ID)
	}
}

func (rv *ReturnValue) InitFiveMin(bs models.BallStack) {
	for _, v := range bs.Balls {
		rv.FiveMin = append(rv.FiveMin, v.ID)
	}
}

func (rv *ReturnValue) InitHour(bs models.BallStack) {
	for _, v := range bs.Balls {
		rv.Hour = append(rv.Hour, v.ID)
	}
}

func (rv *ReturnValue) InitMain(bs models.BallStack) {
	for _, v := range bs.Balls {
		rv.Main = append(rv.Main, v.ID)
	}
}
