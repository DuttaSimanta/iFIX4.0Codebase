package handler

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"strings"

	//"strconv"
	"errors"
	"net/http"
	model "src/entity"
	Logger "src/logger"
	MonITIntegrationApiService "src/models"
)

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func ThrowBlankResponse(w http.ResponseWriter, req *http.Request) {
	model.ThrowJSONResponse(model.BlankPathCheckResponse(), w)
}

func TCSICCPAMGetValidation(w http.ResponseWriter, req *http.Request) {
	Logger.Log.Println("TCSICCArcosGetValidation")

	var successResponse model.SuccessResponseXML
	var errResponse model.ErrorResponseXML
	// arcosUserId, password, ok := req.BasicAuth()
	// Logger.Log.Printf("username==> %s , Password====> %s", arcosUserId, password)
	// Logger.Log.Println("Basic Auth Success-->", ok)
	// if ok == false {
	// 	Logger.Log.Println("Auth not present")
	// 	errResponse.Status = false
	// 	errResponse.Message = "Basic Authentication Fails. Access Denied!!!"
	// 	model.ThrowErrorResponseXML(errResponse, w)
	// 	return
	// }

	body, err := ioutil.ReadAll(req.Body)
	//Logger.Log.Println("BODY======>", string(body))
	if err != nil {
		//Logger.Log.Println(err)
		Logger.Log.Println(err)
		errResponse.Status = "Failure"
		errResponse.Message = "Invalid XML Body"
		model.ThrowErrorResponseXML(errResponse, w)
		return
	}

	//Logger.Log.Printf("username==> %s , Password====> %s", arcosUserId, password)

	//soap := []byte(body)

	reqData := &model.Envelope{}
	err1 := xml.Unmarshal([]byte(body), reqData)
	if err1 != nil {

		Logger.Log.Println(err1)
		errResponse.Status = "Failure"
		errResponse.Message = "Invalid XML Body"
		model.ThrowErrorResponseXML(errResponse, w)
		return
	}
	//Logger.Log.Println("clientcode===>", reqData.Body.ValidateARCOSUser.ClientCode)
	//Logger.Log.Println("OrgCode====>", reqData.Body.ValidateARCOSUser.OrgCode)
	//mapping requestvalues to variables
	var apiKey string = reqData.Body.ValidateARCOSUser.ApiKey
	var clientCode string = reqData.Body.ValidateARCOSUser.ClientCode
	var orgCode string = reqData.Body.ValidateARCOSUser.OrgCode
	var ticketId string = reqData.Body.ValidateARCOSUser.TicketNo
	if strings.EqualFold(ticketId, "") {
		ticketId = reqData.Body.ValidateARCOSUser.SrTaskNumber
		if strings.EqualFold(ticketId, "") {
			ticketId = reqData.Body.ValidateARCOSUser.CRNo
			if strings.EqualFold(ticketId, "") {
				ticketId = reqData.Body.ValidateARCOSUser.TaskNo

			}

		}

	}
	userGroup := ""
	//var userGroup string = reqData.Body.ValidateARCOSUser.UserGroup
	var assignedTo string = reqData.Body.ValidateARCOSUser.UserId

	//notificationErr := MonITIntegrationApiService.GetArcosValidate(arcosUserId, password, apiKey, clientCode, orgCode, ticketId, userGroup, assignedTo)
	notificationErr := MonITIntegrationApiService.GetPAMValidate(apiKey, clientCode, orgCode, ticketId, userGroup, assignedTo)

	if notificationErr != nil {
		Logger.Log.Println(notificationErr)
		errResponse.Status = "Failure"
		errResponse.Message = notificationErr.Error()
		model.ThrowErrorResponseXML(errResponse, w)
		return
	} else {
		Logger.Log.Println("Success")
		successResponse.Status = "Success"
		model.ThrowSuccessResponseXML(successResponse, w)
		return
	}

}

func TCSICCArcosGetValidation(w http.ResponseWriter, req *http.Request) {
	Logger.Log.Println("TCSICCArcosGetValidation")

	var successResponse model.SuccessResponseXML
	var errResponse model.ErrorResponseXML
	// arcosUserId, password, ok := req.BasicAuth()
	// Logger.Log.Printf("username==> %s , Password====> %s", arcosUserId, password)
	// Logger.Log.Println("Basic Auth Success-->", ok)
	// if ok == false {
	// 	Logger.Log.Println("Auth not present")
	// 	errResponse.Status = false
	// 	errResponse.Message = "Basic Authentication Fails. Access Denied!!!"
	// 	model.ThrowErrorResponseXML(errResponse, w)
	// 	return
	// }

	body, err := ioutil.ReadAll(req.Body)
	//Logger.Log.Println("BODY======>", string(body))
	if err != nil {
		//Logger.Log.Println(err)
		Logger.Log.Println(err)
		errResponse.Status = "Failure"
		errResponse.Message = "Invalid XML Body"
		model.ThrowErrorResponseXML(errResponse, w)
		return
	}

	//Logger.Log.Printf("username==> %s , Password====> %s", arcosUserId, password)

	//soap := []byte(body)

	reqData := &model.Envelope{}
	err1 := xml.Unmarshal([]byte(body), reqData)
	if err1 != nil {

		Logger.Log.Println(err1)
		errResponse.Status = "Failure"
		errResponse.Message = "Invalid XML Body"
		model.ThrowErrorResponseXML(errResponse, w)
		return
	}
	//Logger.Log.Println("clientcode===>", reqData.Body.ValidateARCOSUser.ClientCode)
	//Logger.Log.Println("OrgCode====>", reqData.Body.ValidateARCOSUser.OrgCode)
	//mapping requestvalues to variables
	var apiKey string = reqData.Body.ValidateARCOSUser.ApiKey
	var clientCode string = reqData.Body.ValidateARCOSUser.ClientCode
	var orgCode string = reqData.Body.ValidateARCOSUser.OrgCode
	var ticketId string = reqData.Body.ValidateARCOSUser.TicketNo
	if strings.EqualFold(ticketId, "") {
		ticketId = reqData.Body.ValidateARCOSUser.SrTaskNumber
		if strings.EqualFold(ticketId, "") {
			ticketId = reqData.Body.ValidateARCOSUser.CRNo
			if strings.EqualFold(ticketId, "") {
				ticketId = reqData.Body.ValidateARCOSUser.TaskNo

			}

		}

	}
	userGroup := ""
	//var userGroup string = reqData.Body.ValidateARCOSUser.UserGroup
	var assignedTo string = reqData.Body.ValidateARCOSUser.UserId

	//notificationErr := MonITIntegrationApiService.GetArcosValidate(arcosUserId, password, apiKey, clientCode, orgCode, ticketId, userGroup, assignedTo)
	notificationErr := MonITIntegrationApiService.GetArcosValidate(apiKey, clientCode, orgCode, ticketId, userGroup, assignedTo)

	if notificationErr != nil {
		Logger.Log.Println(notificationErr)
		errResponse.Status = "Failure"
		errResponse.Message = notificationErr.Error()
		model.ThrowErrorResponseXML(errResponse, w)
		return
	} else {
		Logger.Log.Println("Success")
		successResponse.Status = "Success"
		model.ThrowSuccessResponseXML(successResponse, w)
		return
	}

}

func TCSICCMonITCreateTicket(w http.ResponseWriter, req *http.Request) {
	Logger.Log.Println("TCSICCMonITCreateTicket")

	var successResponse model.APIResponse
	var errResponse model.ErrorResponse

	var requestData map[string]interface{}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		Logger.Log.Println(err)
		errResponse.Status = false
		errResponse.Message = errors.New("ERROR: Not able to fetch Request Data").Error()
		model.ThrowJSONErrorResponse(errResponse, w)
		return
	}
	jsonErr := json.Unmarshal(body, &requestData)
	if jsonErr != nil {
		Logger.Log.Println(jsonErr)
		errResponse.Status = false
		errResponse.Message = errors.New("ERROR: Json Unmarshal error").Error()
		model.ThrowJSONErrorResponse(errResponse, w)
		return
	}

	//Logger.Log.Println("Payload====>", requestData)

	ticketNo, integrationAPIErr := MonITIntegrationApiService.MonITIntegrationApiServiceMethod(requestData)

	if integrationAPIErr != nil {
		Logger.Log.Println(integrationAPIErr)
		errResponse.Status = false
		errResponse.Message = integrationAPIErr.Error()
		//log.Println(errResponse.Message )
		model.ThrowJSONErrorResponse(errResponse, w)
		return
	} else {
		Logger.Log.Println("success=true")
		successResponse.Status = true
		successResponse.Message = "Ticket has been created successfully"
		successResponse.Response.Ticketid = ticketNo

		model.ThrowJSONResponse(successResponse, w)
		return
	}

}
func GetTicketStatus(w http.ResponseWriter, req *http.Request) {
	Logger.Log.Println("GetTicketStatus")

	var successResponse model.StatusResponse
	var errResponse model.ErrorResponse

	var requestData map[string]interface{}
	body, err := ioutil.ReadAll(req.Body)
	Logger.Log.Println("Handlers @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@  body------->",body)
	if err != nil {
		Logger.Log.Println(err)
		errResponse.Status = false
		errResponse.Message = errors.New("ERROR: Not able to fetch Request Data").Error()
		model.ThrowJSONErrorResponse(errResponse, w)
		return
	}
	jsonErr := json.Unmarshal(body, &requestData)
	if jsonErr != nil {
		Logger.Log.Println(jsonErr)
		errResponse.Status = false
		errResponse.Message = errors.New("ERROR: Json Unmarshal error").Error()
		model.ThrowJSONErrorResponse(errResponse, w)
		return
	}

	//Logger.Log.Println("Payload====>", requestData)

	presentStatus, integrationAPIErr := MonITIntegrationApiService.GetTicketStatus(requestData)
	Logger.Log.Println("Handlers @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@  presentStatus,------->",presentStatus)

	if integrationAPIErr != nil {
		Logger.Log.Println(integrationAPIErr)
		errResponse.Status = false
		errResponse.Message = integrationAPIErr.Error()
		//log.Println(errResponse.Message )
		model.ThrowJSONErrorResponse(errResponse, w)
		return
	} else {
		Logger.Log.Println("success=true")
		successResponse.Status = true
		successResponse.Message = "Status Has been retrieved successfully"
		successResponse.TicketStatus = presentStatus

		model.ThrowJSONStatusResponse(successResponse, w)
		return
	}

}
