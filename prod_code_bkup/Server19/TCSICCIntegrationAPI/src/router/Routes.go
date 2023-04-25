package router

import (
	"net/http"
	controller "src/handler"
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
		"/iFIXapi",
		controller.ThrowBlankResponse,
	},
	Route{
		"monittcsicccreateticket",
		"POST",
		"/iFIXapi/monit_tcs_icc_createticket",
		controller.TCSICCMonITCreateTicket,
	},
	Route{
		"getticketstatus",
		"POST",
		"/iFIXapi/get_ticketstatus",
		controller.GetTicketStatus,
	},
	Route{
		"getticketstatus",
		"POST",
		"/iFIXapi/arcos_tcs_icc_validateuser",
		controller.TCSICCArcosGetValidation,
	},
	Route{
		"getticketstatus",
		"POST",
		"/iFIXapi/pam_tcs_icc_validateuser",
		controller.TCSICCPAMGetValidation,
	},
}
