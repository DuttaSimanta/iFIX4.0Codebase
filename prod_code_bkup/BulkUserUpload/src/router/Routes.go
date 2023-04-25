package router

import (
	"net/http"
	controller "src/handlers"
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
		"/users",
		controller.ThrowBlankResponse,
	},
	Route{
		"bulkuserUpload",
		"POST",
		"/users/bulkuserupload",
		controller.BulkUserUpload,
	},
	{
		"bulkuserdownload",
		"POST",
		"/users/bulkuserdownload",
		controller.BulkUserDownload,
	},
}
