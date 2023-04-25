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
		"/categories",
		handlers.ThrowBlankResponse,
	},
	Route{
		"bulkcategoryUpload",
		"POST",
		"/categories/bulkcategoryupload",
		handlers.BulkCategoryUpload,
	},
	Route{
		"bulkcategoryDataDownload",
		"POST",
		"/categories/bulkcategorydownload",
		handlers.BulkCategoryDownload,
	},
}
