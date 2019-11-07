package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter creates a mux router
func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		log.Printf("Registering handler %s %s %s", route.Name, route.Method, route.Pattern)
		r.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)
	}
	return r
}
