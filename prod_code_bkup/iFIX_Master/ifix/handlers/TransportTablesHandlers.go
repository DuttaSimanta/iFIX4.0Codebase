package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowtablelResponse(successMessage string, responseData []entities.TableEntity, w http.ResponseWriter, success bool) {
	var response = entities.TableResponse{}
	response.Success = success
	response.Message = successMessage
	response.Details = responseData
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		logger.Log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
func ThrowTransporttableAllResponse(successMessage string, responseData entities.TransporttableEntities, w http.ResponseWriter, success bool) {
	var response = entities.TransporttableResponse{}
	response.Success = success
	response.Message = successMessage
	response.Details = responseData
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		logger.Log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
func ThrowTransporttabletypeAllResponse(successMessage string, responseData []entities.GettableEntity, w http.ResponseWriter, success bool) {
	var response = entities.TransporttabletypeResponse{}
	response.Success = success
	response.Message = successMessage
	response.Details = responseData
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		logger.Log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
func ThrowTransporttableIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.TransporttableResponseInt{}
	response.Success = success
	response.Message = successMessage
	response.Details = responseData
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		logger.Log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func InsertTransporttable(w http.ResponseWriter, req *http.Request) {
	var data = entities.TransporttableEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.InsertTransporttable(&data)
		ThrowTransporttableIntResponse(msg, data, w, success)
	}
}

func GetAllTransporttable(w http.ResponseWriter, req *http.Request) {
	var data = entities.TransporttableEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllTransporttable(&data)
		ThrowTransporttableAllResponse(msg, data, w, success)
	}
}

func DeleteTransporttable(w http.ResponseWriter, req *http.Request) {
	var data = entities.TransporttableEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteTransporttable(&data)
		ThrowTransporttableIntResponse(msg, 0, w, success)
	}
}

func UpdateTransporttable(w http.ResponseWriter, req *http.Request) {
	var data = entities.TransporttableEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.UpdateTransporttable(&data)
		ThrowTransporttableIntResponse(msg, 0, w, success)
	}
}
func Gettypedescription(w http.ResponseWriter, req *http.Request) {
	var data = entities.TransporttableEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data1, success, _, msg := models.Gettypedescription(&data)
		ThrowTransporttabletypeAllResponse(msg, data1, w, success)
	}
}
func Gettable(w http.ResponseWriter, req *http.Request) {
	var data = entities.TransporttableEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data1, success, _, msg := models.Gettable(&data)
		ThrowtablelResponse(msg, data1, w, success)
	}
}
func Gettypefortransport(w http.ResponseWriter, req *http.Request) {
	var data = entities.TransporttableEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data1, success, _, msg := models.Gettypefortransport(&data)
		ThrowTransporttabletypeAllResponse(msg, data1, w, success)
	}
}
