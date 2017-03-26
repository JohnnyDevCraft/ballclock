package main

import (
	"ballclock/models"
	"ballclock/routing"
	"log"
	"net/http"
)

var clock models.Clock

func main() {
	clock = models.Clock{}
	routing.MapAllRoutes()
	log.Fatal(http.ListenAndServe(":8081", nil))
}
