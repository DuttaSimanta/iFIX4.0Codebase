package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"

	//"strconv"
	"errors"
	"net/http"
	"src/entities"
	Logger "src/logger"
	InstantMessagingService "src/models"
)

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func ThrowBlankResponse(w http.ResponseWriter, req *http.Request) {
	entities.ThrowJSONResponse(entities.BlankPathCheckResponse(), w)
}

func InstantMessaging(w http.ResponseWriter, req *http.Request) {
	var successResponse entities.APIResponse
	var errResponse entities.ErrorResponse
	//var payload map[string]interface{}
	var instantMessagingEntity entities.InstantMessagingEntity
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		Logger.Log.Println(err)
		errResponse.Status = false
		errResponse.Message = errors.New("ERROR: Not able to fetch Request Data").Error()
		entities.ThrowJSONErrorResponse(errResponse, w)
		return
	}
	jsonErr := json.Unmarshal(body, &instantMessagingEntity)
	if jsonErr != nil {
		Logger.Log.Println(jsonErr)
		errResponse.Status = false
		errResponse.Message = errors.New("ERROR: Json Unmarshal error").Error()
		entities.ThrowJSONErrorResponse(errResponse, w)
		return
	}

	Logger.Log.Println("instantMessagingEntity====>", instantMessagingEntity)

	notificationErr := InstantMessagingService.RecordInstantMessaging(&instantMessagingEntity)

	if notificationErr != nil {
		log.Println(notificationErr)
		errResponse.Status = false
		errResponse.Message = notificationErr.Error()
		//log.Println(errResponse.Message )
		entities.ThrowJSONErrorResponse(errResponse, w)
		return
	} else {
		successResponse.Status = true
		successResponse.Message = "Email Sent Successfully"
		entities.ThrowJSONResponse(successResponse, w)
		return
	}

}
