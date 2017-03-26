package models

//SimRun is used for requesting simulation runs
type SimRun struct {
	Interval  string
	BallCount int
	Years     int
	Months    int
	Days      int
	Hours     int
	Minutes   int
}
