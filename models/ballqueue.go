package models

//BallQueue Object
type BallQueue struct {
	balls []*Ball
}

//Insert adds a ball to the Queue
func (bq BallQueue) Insert(b *Ball) {
	bq.balls = append(bq.balls, b)
}

//Pop gets the next item from the Queue
func (bq BallQueue) Pop() *Ball {
	ball := bq.balls[0]
	length := bq.Length()
	bq.balls = bq.balls[1:length]
	return ball
}

//ToBallStack Returns a BallStack with same elements
func (bq BallQueue) ToBallStack() BallStack {
	return BallStack{bq.balls}
}

//Length gets the length of the Queue
func (bq BallQueue) Length() int {
	return len(bq.balls)
}

//InsertRange adds a range of Balls to the stack
func (bq BallQueue) InsertRange(balls []*Ball) {
	for _, v := range balls {
		bq.balls = append(bq.balls, v)
	}
}
