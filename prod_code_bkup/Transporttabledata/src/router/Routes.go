package router

import (
	"net/http"
	"src/handlers"
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
		"/ifixtransport",
		ThrowBlankResponse,
	},
	Route{
		"updateifixsysid",
		"POST",
		"/ifixtransport/updateifixsysid",
		handlers.UpdateIfixsysid,
	},
	Route{
		"uploadmasterdata",
		"POST",
		"/ifixtransport/uploadmasterdata",
		handlers.UploadTablesdata,
	},
	Route{
		"downloadmasterdata",
		"POST",
		"/ifixtransport/downloadmasterdata",
		handlers.DownloadTablesdata,
	},
}
