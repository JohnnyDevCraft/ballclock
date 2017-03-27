package menu

import (
	"ballclock/models"
	"encoding/json"
	"fmt"
)

func SimModeTwoStep1() {

	for {
		ClearScreen()
		RenderTitle("Sim Mode 1 : Run Simulation")
		bc, _ := GetIntValue("Enter ball count for simulation.")
		min, _ := GetIntValue("Enter the number of minutes for simulation.")

		if bc > 27 && bc < 127 {
			if min > 0 {
				clock := models.Clock{}

				clock.PopulateBalls(bc)
				clock.SetupLevels(5)

				for i := 0; i < min; i++ {
					clock.Progress()
				}

				rv := ReturnValue{}
				rv.InitMin(clock.Levels[0].Balls)
				rv.InitFiveMin(clock.Levels[1].Balls)
				rv.InitHour(clock.Levels[2].Balls)
				rv.InitMain(clock.PickUp.ToBallStack())

				js, _ := json.Marshal(rv)
				str := fmt.Sprintf("%s \n\nPress any key to continue.", string(js))
				SimModeTwoResults(string(str))
				Reader.ReadByte()
				return
			}
		}

	}

}

func SimModeTwoResults(json string) {
	ClearScreen()
	RenderTitle("Sim Mode 1: Results")
	fmt.Println(json)
}
