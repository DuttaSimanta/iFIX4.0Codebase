package router

import (
	"net/http"
	handlers "src/handlers"
)

//Route is a basic sturct
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"404 Not Found",
		"POST",
		"/iFIXMessaging",
		handlers.ThrowBlankResponse,
	},
	Route{
		"InstantMesseging",
		"POST",
		"/iFIXMessaging/instantmessaging",
		handlers.InstantMessaging,
	},
}
