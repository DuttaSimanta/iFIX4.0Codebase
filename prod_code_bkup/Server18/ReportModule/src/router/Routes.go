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
		"/reports",
		ThrowBlankResponse,
	},
Route{
		"heartbeat",
		"POST",
		"/reports/heartbeat",
		handlers.GetReportDockerStatus,
	},
	Route{
		"generatereport",
		"POST",
		"/reports/generatereport",
		handlers.JsonToExcelConverter,
	},
	{
		"reporttodownloadlist",
		"POST",
		"/reports/reportgeneratedlist",
		handlers.GetDownloadList,
	},
	{
		"getqueryresult",
		"POST",
		"/reports/getqueryresult",
		handlers.RecordGridResultOnly,
	},
}
