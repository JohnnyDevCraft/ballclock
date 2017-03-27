package models

import ()

//Clock Object
type Clock struct {
	Levels    []Level
	Balls     BallStack
	PickUp    BallQueue
	RoundTime int
	RtYears   int
	RtDays    int
	RtHours   int
	RtMinutes int
}

//SetClock fastforwards throught the simulation until time matches system time.
func (c *Clock) SetClock() {

}

//UpdateRoundTimes Updates RoundTimes
func (c *Clock) UpdateRoundTimes(val int) {
	c.RoundTime = val
	c.RtMinutes = val % 60
	hours := val / 60
	c.RtHours = hours % 24
	days := hours / 24
	c.RtDays = days % 365
	c.RtYears = days / 365
}

//CalRunTimes calculates run times by simulating a run
func (c *Clock) CalRunTimes() {
	found := false
	count := 1
	for !found {
		c.Progress()
		if c.IsOrig() {
			c.UpdateRoundTimes(count)
			found = true
		}
		count++
	}
}

//PopulateBalls Itializes the Clock with specified Ball Count
func (c *Clock) PopulateBalls(count int) {
	//fmt.Printf("Populating balls on clock")
	c.Balls = BallStack{nil}

	for i := 1; i <= count; i++ {
		ball := Ball{i, 0}
		c.Balls.Insert(&ball)
	}

	c.PickUp = c.Balls.ToBallQueue()

	//fmt.Printf("%d total balls poulated\n", c.Balls.Length())

}

//SetupLevels initializes the Levels Collection
func (c *Clock) SetupLevels(interval int) {
	c.Levels = nil

	level1 := Level{"Minutes", interval, 1, false, BallStack{nil}}

	level2 := Level{"Minute Segments", 60 / interval, interval, false, BallStack{nil}}

	level3 := Level{"Hours", 12, 1, true, BallStack{nil}}

	c.Levels = append(c.Levels, level1, level2, level3)
}

//Progress progresses the clock forward one increment.
func (c *Clock) Progress() {
	pickup := c.PickUp.Pop()
	pickup.PickupCount++
	//fmt.Printf("Ball[%d] is being moved to top level.\n", pickup.ID)
	c.addToLevel(pickup, 0)

}

//IsOrig determins if the queue is in original order
func (c *Clock) IsOrig() bool {

	if c.PickUp.Length() == c.Balls.Length() {

		pLen := c.PickUp.Length()
		for i := 0; i < pLen; i++ {
			if c.PickUp.Balls[i].ID != i+1 {
				return false
			}
		}
		return true
	}
	return false
}

//addToLevel adds a ball to a level
func (c *Clock) addToLevel(b *Ball, level int) {
	switch {
	case level == 0:
		if c.Levels[0].Add(b) {
			if c.Levels[0].IsFull() {
				stack, ball := c.Levels[0].Empty()
				stack.Reverse()
				c.PickUp.InsertRange(stack.Balls)
				c.addToLevel(ball, 1)
				c.Levels[0].Balls.Clear()
			}
		}
	case level == 1:
		if c.Levels[1].Add(b) {
			if c.Levels[1].IsFull() {
				stack, ball := c.Levels[1].Empty()
				stack.Reverse()
				c.PickUp.InsertRange(stack.Balls)
				c.addToLevel(ball, 2)
				c.Levels[1].Balls.Clear()
			}
		}
	case level == 2:
		if c.Levels[2].Add(b) {
			if c.Levels[2].IsFull() {
				stack, ball := c.Levels[2].Empty()
				stack.Reverse()
				c.PickUp.Insert(ball)
				c.PickUp.InsertRange(stack.Balls)
				c.Levels[2].Balls.Clear()
			}
		}
	}
}
