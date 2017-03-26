package routing

import (
	"ballclock/menu"
	"ballclock/models"
	"encoding/json"
	"net/http"
)

//MapSimRoutes maps the routes in this file
func MapSimRoutes() {
	http.HandleFunc("/RunSim/", RunSimRoute)
}

//RunSimRoute runs a simgle simulations and return result
func RunSimRoute(w http.ResponseWriter, r *http.Request) {
	menu.ClearScreen()
	//fmt.Printf("Request for resource recieived: %s\n\n", r.URL.Path[1:])

	decoder := json.NewDecoder(r.Body)
	var simRun models.SimRun

	err := decoder.Decode(&simRun)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("Data passed to request: \n%s\n\n\n", r.Body)
	//simRunVal, _ := json.Marshal(simRun)
	//fmt.Printf("Converted Data passed to request: \n%s\n\n\n", simRunVal)

	w.Header().Set("Content-Type", "application/json")

	clock := models.Clock{}
	clock.PopulateBalls(simRun.BallCount)

	minutes := (simRun.Years * 365 * 24 * 60) + (((simRun.Years - (simRun.Years % 4)) / 4) * 24 * 60) + (simRun.Days * 24 * 60) + (simRun.Hours * 60) + simRun.Minutes

	clock.SetupLevels(menu.StringToInt(simRun.Interval))

	for i := 0; i < minutes; i++ {
		//fmt.Printf("Tick [%d] of [%d]\n", i+1, minutes)
		//fmt.Printf("Level 1 [%d] Level 2 [%d] Level 3 [%d] Queue [%d]\n", clock.Levels[0].Balls.Length(), clock.Levels[1].Balls.Length(), clock.Levels[2].Balls.Length(), clock.PickUp.Length())
		//fmt.Printf("**************************************************************\n")
		clock.Progress()
		//fmt.Printf("**************************************************************\n\n\n")
	}
	//val, _ := json.Marshal(clock)
	//fmt.Printf("Data passed to ResponseWriter: \n%s\n\n\n", val)
	json.NewEncoder(w).Encode(clock)
	return
}
