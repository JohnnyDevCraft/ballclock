package models

//BallStack Object
type BallStack struct {
	balls []*Ball
}

//Insert adds an element to the top of the stack.
func (bs BallStack) Insert(b *Ball) {
	bs.balls = append(bs.balls, b)
}

//InsertRange adds a range of Balls to the stack
func (bs BallStack) InsertRange(balls []*Ball) {
	for _, v := range balls {
		bs.balls = append(bs.balls, v)
	}
}

//Pop returns the last item in the slice, and a slice without the item.
func (bs BallStack) Pop() *Ball {
	length := len(bs.balls)
	ball := bs.balls[length]
	bs.balls = bs.balls[:length-1]
	return ball
}

//Peek Returns top object without removing it from the stack.
func (bs BallStack) Peek() *Ball {
	return bs.balls[bs.Length()]
}

//ToBallQueue Returns a BallQueue with same elements
func (bs BallStack) ToBallQueue() BallQueue {
	return BallQueue{bs.balls}
}

//Reverse Reverses the order of the Stack
func (bs BallStack) Reverse() {
	var new []*Ball

	for i := bs.Length() - 1; i >= 0; i-- {
		new = append(new, bs.balls[i])
	}

	bs.balls = new

}

//Length returns the length of the slice
func (bs BallStack) Length() int {
	return len(bs.balls)
}
