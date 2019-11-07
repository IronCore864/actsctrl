package router

import (
	"net/http"

	"gitlab.com/ironcore864/actsctrl/handler"
)

// Route is the struct for an entry of route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is the list of all routes
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"POST",
		"/",
		handler.IndexPostHandler,
	},
}
