package handlers

import (
	"encoding/json"
	"errors"
	"iFIX/ifix/entities"
	Logger "iFIX/ifix/logger"
	service "iFIX/ifix/models"
	"io/ioutil"
	"log"
	"net/http"
)

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}
func ThrowBlankResponse(w http.ResponseWriter, req *http.Request) {
	entities.EmailThrowJSONResponse(entities.EmailBlankPathCheckResponse(), w)
}
func GetDelimiter(w http.ResponseWriter, req *http.Request) {
	var successResponse entities.EmailAPIResponse
	var errResponse entities.EmailErrorResponse

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
	delimeter, getdelemiterErr := service.GetDelimiter(result)
	if getdelemiterErr != nil {
		log.Println(getdelemiterErr)
		errResponse.Status = false
		errResponse.Message = getdelemiterErr.Error()
		//log.Println(errResponse.Message )
		entities.ThrowJSONErrorResponse(errResponse, w)
		return
	} else {
		successResponse.Status = true
		successResponse.Message = "Delemeter has been fetched  Successfuly"
		successResponse.Delimeter = delimeter
		entities.EmailThrowJSONResponse(successResponse, w)
		return
	}
}

func GetLastCategoryList(w http.ResponseWriter, req *http.Request) {

	var errResponse entities.EmailErrorResponse
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
	lastCategoryList, getdelemiterErr := service.GetLastCategoryList(result)
	if getdelemiterErr != nil {
		log.Println(getdelemiterErr)
		errResponse.Status = false
		errResponse.Message = getdelemiterErr.Error()
		//log.Println(errResponse.Message )
		entities.ThrowJSONErrorResponse(errResponse, w)
		return
	} else {
		status := true
		message := "Succefully Fetched LastCategory List"
		entities.ThrowCategoryListResponse(message, lastCategoryList, w, status)
		return
	}

}
func GetServiceUser(w http.ResponseWriter, req *http.Request) {
	var errResponse entities.EmailErrorResponse
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
	userList, getServiceUserErr := service.GetServiceUser(result)
	if getServiceUserErr != nil {
		log.Println(getServiceUserErr)
		errResponse.Status = false
		errResponse.Message = getServiceUserErr.Error()
		//log.Println(errResponse.Message )
		entities.ThrowJSONErrorResponse(errResponse, w)
		return
	} else {
		status := true
		message := "Succefully Fetched UserList"
		entities.ThrowServiceUsersResponse(message, userList, w, status)
		return
	}

}
func SaveEmailTicketConfiguration(w http.ResponseWriter, req *http.Request) {
	var successResponse entities.EmailAPIResponse
	var errResponse entities.EmailErrorResponse

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
	saveError := service.SaveEmailTicketConfiguration(result)
	if saveError != nil {
		log.Println(saveError)
		errResponse.Status = false
		errResponse.Message = saveError.Error()
		//log.Println(errResponse.Message )
		entities.ThrowJSONErrorResponse(errResponse, w)
		return
	} else {
		successResponse.Status = true
		successResponse.Message = "Saved Successfully"
		//successResponse.Delimeter = delimeter
		entities.EmailThrowJSONResponse(successResponse, w)
		return
	}
}
func GetEmailTicketConfigurations(w http.ResponseWriter, req *http.Request) {
	var errResponse entities.EmailErrorResponse

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
	emailTicketViewList, emailticketViewError := service.GetEmailTicketConfigurations(result)
	if emailticketViewError != nil {
		log.Println(emailticketViewError)
		errResponse.Status = false
		errResponse.Message = emailticketViewError.Error()
		//log.Println(errResponse.Message )
		entities.ThrowJSONErrorResponse(errResponse, w)
		return
	} else {
		status := true
		message := "Email Ticket List Fetched Successfully"
		//successResponse.Delimeter = delimeter
		entities.ThrowMstEmailTicketConfigtResponse(message, emailTicketViewList, w, status)
		return
	}
}

func DeleteEmailTicketConfiguration(w http.ResponseWriter, req *http.Request) {
	var successResponse entities.EmailAPIResponse
	var errResponse entities.EmailErrorResponse

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
	deleteError := service.DeleteEmailTicketConfiguration(result)
	if deleteError != nil {
		log.Println(deleteError)
		errResponse.Status = false
		errResponse.Message = deleteError.Error()
		//log.Println(errResponse.Message )
		entities.ThrowJSONErrorResponse(errResponse, w)
		return
	} else {
		successResponse.Status = true
		successResponse.Message = "Deleted Successfully"
		//successResponse.Delimeter = delimeter
		entities.EmailThrowJSONResponse(successResponse, w)
		return
	}
}

func DeleteEmailTicketConfigu(w http.ResponseWriter, req *http.Request) {
	var successResponse entities.EmailAPIResponse
	var errResponse entities.EmailErrorResponse

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
	deleteError := service.DeleteEmailTicketConfigu(result)
	if deleteError != nil {
		log.Println(deleteError)
		errResponse.Status = false
		errResponse.Message = deleteError.Error()
		//log.Println(errResponse.Message )
		entities.ThrowJSONErrorResponse(errResponse, w)
		return
	} else {
		successResponse.Status = true
		successResponse.Message = "Deleted Successfully"
		//successResponse.Delimeter = delimeter
		entities.EmailThrowJSONResponse(successResponse, w)
		return
	}
}

func UpdateEmailTicketConfiguration(w http.ResponseWriter, req *http.Request) {
	var successResponse entities.EmailAPIResponse
	var errResponse entities.EmailErrorResponse

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
	updateError := service.UpdateEmailTicketConfiguration(result)
	if updateError != nil {
		log.Println(updateError)
		errResponse.Status = false
		errResponse.Message = updateError.Error()
		//log.Println(errResponse.Message )
		entities.ThrowJSONErrorResponse(errResponse, w)
		return
	} else {
		successResponse.Status = true
		successResponse.Message = "Updated Successfully"
		//successResponse.Delimeter = delimeter
		entities.EmailThrowJSONResponse(successResponse, w)
		return
	}
}
func AddEmailBaseConfig(w http.ResponseWriter, req *http.Request) {
	var successResponse entities.EmailAPIResponse
	var errResponse entities.EmailErrorResponse

	var result = entities.EmailBaseConfig{}
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
	updateError := service.AddEmailBaseConfig(&result)
	if updateError != nil {
		log.Println(updateError)
		errResponse.Status = false
		errResponse.Message = updateError.Error()
		//log.Println(errResponse.Message )
		entities.ThrowJSONErrorResponse(errResponse, w)
		return
	} else {
		successResponse.Status = true
		successResponse.Message = "Added Successfully"
		//successResponse.Delimeter = delimeter
		entities.EmailThrowJSONResponse(successResponse, w)
		return
	}
}

func GetDelimiterForAllClient(w http.ResponseWriter, req *http.Request) {
	var errResponse entities.EmailErrorResponse

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
	emailTicketViewList, emailticketViewError := service.GetDelimiterForAllClient(result)
	if emailticketViewError != nil {
		log.Println(emailticketViewError)
		errResponse.Status = false
		errResponse.Message = emailticketViewError.Error()
		//log.Println(errResponse.Message )
		entities.ThrowJSONErrorResponse(errResponse, w)
		return
	} else {
		status := true
		message := "Email Ticket List Fetched Successfully"
		//successResponse.Delimeter = delimeter
		entities.ThrowMstEmailTicketBaseConfigtResponse(message, emailTicketViewList, w, status)
		return
	}
}
