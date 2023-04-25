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
		"/iFIXNotification",
		handlers.ThrowBlankResponse,
	},
	Route{
		"sendemailnotification",
		"POST",
		"/iFIXNotification/sendemailnotification",
		handlers.SendEmailNotification,
	}, Route{
		"sendsmsnotification",
		"POST",
		"/iFIXNotification/sendsmsnotification",
		handlers.SendSMSNotification,
	}, Route{
		"sendsmsnotification",
		"POST",
		"/iFIXNotification/sendnotification",
		handlers.SendNotification,
	},
}
