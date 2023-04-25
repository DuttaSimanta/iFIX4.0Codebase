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
	SendNotificationModels "src/models"
)

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

//var mutex = &sync.Mutex{}

func ThrowBlankResponse(w http.ResponseWriter, req *http.Request) {
	entities.ThrowJSONResponse(entities.BlankPathCheckResponse(), w)
}

func SendEmailNotification(w http.ResponseWriter, req *http.Request) {
	var successResponse entities.APIResponse
	var errResponse entities.ErrorResponse

	var result map[string]interface{}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		Logger.Log.Println(err)
		errResponse.Status = false
		errResponse.Message = errors.New("ERROR: Not able to fetch Request Data").Error()
		entities.ThrowJSONErrorResponse(errResponse, w)
		return
	}
	jsonErr := json.Unmarshal(body, &result)
	Logger.Log.Println("Payload====>", result)
	if jsonErr != nil {
		Logger.Log.Println(jsonErr)
		errResponse.Status = false
		errResponse.Message = errors.New("ERROR: Json Unmarshal error").Error()
		entities.ThrowJSONErrorResponse(errResponse, w)
		return
	}

	// mutex.Lock()
	// defer mutex.Unlock()

	// db, dBerr := config.GetDB()
	// if dBerr != nil {
	// 	Logger.Log.Println(dBerr)
	// 	return
	// }
	//db.Ping()
	//defer db.Close()
	notificationErr := SendNotificationModels.MailBodyFormationFromTemplate(result)

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
func SendSMSNotification(w http.ResponseWriter, req *http.Request) {
	var successResponse entities.APIResponse
	var errResponse entities.ErrorResponse

	var result map[string]interface{}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		Logger.Log.Println(err)
		errResponse.Status = false
		errResponse.Message = errors.New("ERROR: Not able to fetch Request Data").Error()
		entities.ThrowJSONErrorResponse(errResponse, w)
		return
	}
	jsonErr := json.Unmarshal(body, &result)
	Logger.Log.Println("Payload====>", result)
	if jsonErr != nil {
		Logger.Log.Println(jsonErr)
		errResponse.Status = false
		errResponse.Message = errors.New("ERROR: Json Unmarshal error").Error()
		entities.ThrowJSONErrorResponse(errResponse, w)
		return
	}

	// mutex.Lock()
	// defer mutex.Unlock()

	// db, dBerr := config.GetDB()
	// if dBerr != nil {
	// 	Logger.Log.Println(dBerr)
	// 	return
	// }
	//db.Ping()
	//defer db.Close()
	notificationErr := SendNotificationModels.SMSFormationFromTemplate(result)

	if notificationErr != nil {
		log.Println(notificationErr)
		errResponse.Status = false
		errResponse.Message = notificationErr.Error()
		//log.Println(errResponse.Message )
		entities.ThrowJSONErrorResponse(errResponse, w)
		return
	} else {
		successResponse.Status = true
		successResponse.Message = "SMS Sent Successfully"
		entities.ThrowJSONResponse(successResponse, w)
		return
	}

}
func SendNotification(w http.ResponseWriter, req *http.Request) {
	var successResponse entities.APIResponse
	var errResponse entities.ErrorResponse

	var result map[string]interface{}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		Logger.Log.Println(err)
		errResponse.Status = false
		errResponse.Message = errors.New("ERROR: Not able to fetch Request Data").Error()
		entities.ThrowJSONErrorResponse(errResponse, w)
		return
	}
	jsonErr := json.Unmarshal(body, &result)
	Logger.Log.Println("Payload====>", result)
	if jsonErr != nil {
		Logger.Log.Println(jsonErr)
		errResponse.Status = false
		errResponse.Message = errors.New("ERROR: Json Unmarshal error").Error()
		entities.ThrowJSONErrorResponse(errResponse, w)
		return
	}

	// mutex.Lock()
	// defer mutex.Unlock()

	// db, dBerr := config.GetDB()
	// if dBerr != nil {
	// 	Logger.Log.Println(dBerr)
	// 	return
	// }
	//db.Ping()
	//defer db.Close()
	notificationErr := SendNotificationModels.SendNotifications(result)

	if notificationErr != nil {
		log.Println(notificationErr)
		errResponse.Status = false
		errResponse.Message = notificationErr.Error()
		//log.Println(errResponse.Message )
		entities.ThrowJSONErrorResponse(errResponse, w)
		return
	} else {
		successResponse.Status = true
		successResponse.Message = "NotiFications Sent Successfully"
		entities.ThrowJSONResponse(successResponse, w)
		return
	}

}
