package product

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var controller = &Controller{Repository: Repository{}}

// Route defines a route
type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

// Routes defines the list of routes
type Routes []Route

var routes = Routes{
	Route{
		"Authentication",
		"POST",
		"/get-token",
		controller.GetToken,
	},

	Route{
		"Index",
		"GET",
		"/",
		controller.Index,
	},
}

// NewRouter configures a new router to th app API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler
		log.Println(route.Name)
		handler = route.HandleFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
