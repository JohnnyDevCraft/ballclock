package models

//BallQueue Object
type BallQueue struct {
	Balls []*Ball
}

//Insert adds a ball to the Queue
func (bq *BallQueue) Insert(b *Ball) {
	bq.Balls = append(bq.Balls, b)
}

//Pop gets the next item from the Queue
func (bq *BallQueue) Pop() *Ball {
	ball := bq.Balls[0]
	length := bq.Length()
	bq.Balls = bq.Balls[1:length]
	return ball
}

//ToBallStack Returns a BallStack with same elements
func (bq *BallQueue) ToBallStack() BallStack {
	return BallStack{bq.Balls}
}

//Length gets the length of the Queue
func (bq *BallQueue) Length() int {
	return len(bq.Balls)
}

//InsertRange adds a range of Balls to the stack
func (bq *BallQueue) InsertRange(balls []*Ball) {
	for _, v := range balls {
		bq.Balls = append(bq.Balls, v)
	}
}
