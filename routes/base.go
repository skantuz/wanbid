package routes

import (
	"net/http"
	c "wanbid/controllers"

	"github.com/gorilla/mux"
)

//Routes estructira de cada ruta
type Route struct {
	Name       string
	Metrhod    string
	Pattern    string
	HandleFunc http.HandlerFunc
}

func NewRouter() {
	router := mux.NewRouter().strictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Handler(router.HandleFunc)
	}

}

var routes = []Route{
	//check run API
	Route{
		"Index",
		"GET",
		"/",
		c.Index,
	},
}
