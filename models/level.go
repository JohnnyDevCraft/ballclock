package models

//Level Object
type Level struct {
	Name         string
	MaxBalls     int
	Increment    int
	ShowPosition bool
	Balls        BallStack
}

//Add adds a ball to the level
func (l Level) Add(ball *Ball) bool {
	length := l.Balls.Length()

	if length < l.MaxBalls {
		l.Balls.Insert(ball)
		return true
	}

	return false
}

//IsFull returns true if the Level has reached capacity.
func (l Level) IsFull() bool {
	if l.Balls.Length() >= l.MaxBalls {
		return true
	}
	return false
}

//Empty Empties the level when it is full
func (l Level) Empty() (BallStack, *Ball) {
	var layover *Ball
	layover = l.Balls.Pop()
	l.Balls.Reverse()
	return l.Balls, layover
}
