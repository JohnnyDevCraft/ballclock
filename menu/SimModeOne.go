package menu

import (
	"ballclock/models"
	"fmt"
)

func SimModeOneStep1() {

	for {
		ClearScreen()
		RenderTitle("Sim Mode 1 : Run Simulation")
		bc, _ := GetIntValue("Enter ball count for simulation.")
		fmt.Println(bc)

		if bc > 27 && bc < 127 {
			clock := models.Clock{}

			clock.PopulateBalls(bc)
			clock.SetupLevels(5)

			clock.CalRunTimes()

			rv := ReturnValue{}
			rv.InitMin(clock.Levels[0].Balls)
			rv.InitFiveMin(clock.Levels[1].Balls)
			rv.InitHour(clock.Levels[2].Balls)
			rv.InitMain(clock.PickUp.ToBallStack())
			fmt.Println(rv)
			str := fmt.Sprintf("%d balls cycle after %d days.\n\nThats a total of %d minutes.", bc, clock.RtDays+(clock.RtYears*365), clock.RoundTime)
			SimModeOneResults(string(str))
			return
		}

	}

}

func SimModeOneResults(json string) {
	ClearScreen()
	RenderTitle("Sim Mode 1: Results")
	result := fmt.Sprintf("%s \n\nPress return to continue.", json)
	GetIntValue(result)
}
