package models

//Ball Object
type Ball struct {
	ID          int
	PickupCount int
}

//Pickup Increments PickupCount
func (b Ball) Pickup() {
	b.PickupCount++
}
