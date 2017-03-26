package models

//BallStack Object
type BallStack struct {
	Balls []*Ball
}

//Insert adds an element to the top of the stack.
func (bs *BallStack) Insert(b *Ball) {
	bs.Balls = append(bs.Balls, b)
}

//InsertRange adds a range of Balls to the stack
func (bs *BallStack) InsertRange(balls []*Ball) {
	for _, v := range balls {
		bs.Balls = append(bs.Balls, v)
	}
}

//Pop returns the last item in the slice, and a slice without the item.
func (bs *BallStack) Pop() *Ball {
	length := len(bs.Balls)
	ball := bs.Balls[length-1]
	bs.Balls = bs.Balls[:length-1]
	return ball
}

//Peek Returns top object without removing it from the stack.
func (bs *BallStack) Peek() *Ball {
	return bs.Balls[bs.Length()]
}

//ToBallQueue Returns a BallQueue with same elements
func (bs *BallStack) ToBallQueue() BallQueue {
	return BallQueue{bs.Balls}
}

//Reverse Reverses the order of the Stack
func (bs *BallStack) Reverse() {
	var new []*Ball

	for i := bs.Length() - 1; i >= 0; i-- {
		new = append(new, bs.Balls[i])
	}

	bs.Balls = new

}

//Length returns the length of the slice
func (bs *BallStack) Length() int {
	return len(bs.Balls)
}

func (bs *BallStack) Clear() {
	bs.Balls = nil
}
