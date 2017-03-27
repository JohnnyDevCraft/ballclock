package routing

import (
	"ballclock/menu"
	"fmt"
	"net/http"
	"strings"
)

//MapAllRoutes maps the routes in this namespace
func MapAllRoutes() {
	MapDefaultRoutes()
	MapSimRoutes()

}

//MapDefaultRoutes maps the routes in this file
func MapDefaultRoutes() {
	http.HandleFunc("/", DefaultRoute)
}

//DefaultRoute routs to index.html in the static folder
func DefaultRoute(w http.ResponseWriter, r *http.Request) {
	menu.ClearScreen()
	fmt.Printf("Request for resource recieived: %s\n\n", r.URL.Path[1:])
	var path string

	if r.URL.Path[1:] != "" {
		path = r.URL.Path[1:]
	} else {
		path = "index.html"
	}

	join := []string{"static/", path}
	stg := strings.Join(join, "")
	http.ServeFile(w, r, stg)
	fmt.Printf("Response sent.\n\n")
}
