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
		"/locations",
		ThrowBlankResponse,
	},
	Route{
		"updatelocation",
		"POST",
		"/locations/updatelocation",
		handlers.UpdateLocation,
	},
	Route{
		"deletelocation",
		"POST",
		"/locations/deletelocation",
		handlers.DeleteLocation,
	},
	Route{
		"prioritydownload",
		"POST",
		"/locations/prioritydownload",
		handlers.BulkLocationWisePriorityDownload,
	},
	Route{
		"priorityupload",
		"POST",
		"/locations/priorityupload",
		handlers.LocationPriorityUpload,
	},
	Route{
		"getlocation",
		"POST",
		"/locations/getlocation",
		handlers.GetAllLocation,
	},
	Route{
		"addlocation",
		"POST",
		"/locations/addlocation",
		handlers.AddLocation,
	},
	Route{
		"selectlocation",
		"POST",
		"/locations/selectlocation",
		handlers.SelectLocation,
	},
	Route{
		"searchlocation",
		"POST",
		"/locations/searchlocation",
		handlers.SearchLocation,
	},
	// Route{
	// 	"bulkcategoryUpload",
	// 	"POST",
	// 	"/categories/bulkcategoryupload",
	// 	handlers.BulkCategoryUpload,
	// },
	// Route{
	// 	"bulkcategoryDataDownload",
	// 	"POST",
	// 	"/categories/bulkcategorydownload",
	// 	handlers.BulkCategoryDownload,
	// },
}
