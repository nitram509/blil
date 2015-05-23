package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nitram509/blil/blilweb/handler"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{"Index", "GET", "/", handler.Index},
	Route{"LedIndex", "GET", "/led", handler.LedIndex},
	Route{"LedSetColor", "POST", "/led/{ledNr}/{color}", handler.LedSetColor},
	Route{"LedGetColor", "GET", "/led/{ledNr}", handler.LedGetColor},
}
