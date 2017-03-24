package models

//Clock Object
type Clock struct {
	Levels []Level
	Balls  BallStack
	PickUp BallQueue
}

//SetClock fastforwards throught the simulation until time matches system time.
func (c Clock) SetClock() {

}

//PopulateBalls Itializes the Clock with specified Ball Count
func (c Clock) PopulateBalls(count int) {
	c.Balls = BallStack{nil}

	for i := 1; i <= count; i++ {
		ball := Ball{i, 0}
		c.Balls.Insert(&ball)
	}

	c.PickUp = c.Balls.ToBallQueue()

}

//SetupLevels initializes the Levels Collection
func (c Clock) SetupLevels(interval int) {
	c.Levels = nil

	level1 := Level{"Minutes", interval, 1, false, BallStack{nil}}

	level2 := Level{"Minute Segments", 60 / interval, interval, false, BallStack{nil}}

	level3 := Level{"Hours", 12, 1, true, BallStack{nil}}

	c.Levels = append(c.Levels, level1, level2, level3)
}

//Progress progresses the clock forward one increment.
func (c Clock) Progress() {
	pickup := c.PickUp.Pop()
	pickup.PickupCount++
	c.addToLevel(pickup, 0)
}

//addToLevel adds a ball to a level
func (c Clock) addToLevel(b *Ball, level int) {
	switch {
	case level == 0:
		if c.Levels[0].Add(b) {
			if c.Levels[0].IsFull() {
				stack, ball := c.Levels[0].Empty()
				stack.Reverse()
				c.PickUp.InsertRange(stack.balls)
				c.addToLevel(ball, 1)
			}
		}
	case level == 1:
		if c.Levels[1].Add(b) {
			if c.Levels[1].IsFull() {
				stack, ball := c.Levels[1].Empty()
				stack.Reverse()
				c.PickUp.InsertRange(stack.balls)
				c.addToLevel(ball, 2)
			}
		}
	case level == 2:
		if c.Levels[2].Add(b) {
			if c.Levels[2].IsFull() {
				stack, ball := c.Levels[2].Empty()
				stack.Reverse()
				c.PickUp.Insert(ball)
				c.PickUp.InsertRange(stack.balls)
			}
		}
	}
}
