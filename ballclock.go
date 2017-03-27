package main

import (
	"ballclock/menu"
	"ballclock/models"
	"ballclock/routing"
	"fmt"
	"log"
	"net/http"
	"time"
)

var clock models.Clock
var run bool

func main() {
	menu.ClearScreen()
	app := menu.App{}
	clock = models.Clock{}
	menu.InitReader()
	switch renderFirstChoice() {
	case 1:
		app.Run(&clock)
	case 2:
		routing.MapAllRoutes()
		renderHttpServer()
		log.Fatal(http.ListenAndServe(":8081", nil))
	}
}

func renderHttpServer() {
	menu.ClearScreen()
	menu.RenderTitle("Running HTTP Server")
	fmt.Println("Server Running at http://127.0.0.1:8081")
}

func renderFirstChoice() int {

	for {
		menu.ClearScreen()
		menu.RenderTitle("Select Run Type")
		menu.RenderLine(1, "Console App")
		menu.RenderLine(2, "Run Web Application")
		val, _ := menu.GetIntValue("Enter Selection")

		if val != 0 && val < 3 {
			return val
		}

	}
}

//InitClockToTime initializes the running clock with 45 balls.
func InitClockToTime() {
	rClock := models.Clock{}
	models.RunningClock = &rClock
	models.RunningClock.PopulateBalls(45)
	models.RunningClock.SetClock()
	go RunClock()
}

//RunClock runs the clock progressor on a one minute interval.
func RunClock() {
	for run {
		time.Sleep(100 * time.Millisecond * 60)
		models.RunningClock.Progress()
	}
}

//StopClock stops the running clock.
func StopClock() {
	run = false
}
