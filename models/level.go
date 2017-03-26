package models

import ()

//Level Object
type Level struct {
	Name         string
	MaxBalls     int
	Increment    int
	ShowPosition bool
	Balls        BallStack
}

//Add adds a ball to the level
func (l *Level) Add(ball *Ball) bool {
	//fmt.Printf("Adding Ball[%d] it level [%s]\n", ball.ID, l.Name)
	length := l.Balls.Length()

	//fmt.Printf("Level Count at [%d] for level [%s]: Max is [%d]\n", l.Balls.Length(), l.Name, l.MaxBalls)
	if length < l.MaxBalls {
		l.Balls.Insert(ball)
		//fmt.Printf("Level Count at [%d] for level [%s]\n", l.Balls.Length(), l.Name)
		return true
	}

	return false
}

//IsFull returns true if the Level has reached capacity.
func (l *Level) IsFull() bool {
	if l.Balls.Length() >= l.MaxBalls {
		return true
	}
	return false
}

//Empty Empties the level when it is full
func (l *Level) Empty() (BallStack, *Ball) {
	//fmt.Printf("Level full and emptying with Count at [%d] for level [%s]\n", l.Balls.Length(), l.Name)
	var layover *Ball
	layover = l.Balls.Pop()
	//fmt.Printf("Rolling Ball[%d] pver to nextlevel from level [%s]: [%d] balls still to empty.\n", layover.ID, l.Name, l.Balls.Length())
	return l.Balls, layover
}
