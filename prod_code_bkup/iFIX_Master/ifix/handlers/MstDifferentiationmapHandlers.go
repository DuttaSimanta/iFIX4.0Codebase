package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"net/http"
)

// ThrowClientUserRoleIntResponse function is used to throw success response of integer data in JSON format
func ThrowDifferentiationMapResponse(successMessage string, w http.ResponseWriter, success bool) {
	var response = entities.MapClientUserRoleUserEntityResponseInt{}
	response.Success = success
	response.Message = successMessage
	//response.Details = responseData
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		logger.Log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func ThrowDifferentiationMapAllResponse(successMessage string, responseData entities.MstDifferentiationmaEntities, w http.ResponseWriter, success bool) {
	var response = entities.MstResponse{}
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

//Createdifferentiationmap method is used for insert user data
func Createdifferentiationmap(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstDifferentiationmapEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		success, _, msg := models.InsertDifferentiationModelsMap(&data)
		ThrowDifferentiationMapResponse(msg, w, success)
	}
}

//Deletedifferentiationmap method is used for insert user data
func Deletedifferentiationmap(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstDifferentiationmapEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		success, _, msg := models.DeleteDifferentiationModelsMap(&data)
		ThrowDifferentiationMapResponse(msg, w, success)
	}
}

func GetAllDifferentiationDtls(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstDifferentiationmapEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllDifferentiationMapDtls(&data)
		ThrowDifferentiationMapAllResponse(msg, data, w, success)
	}
}

//Createrecordsterms method is used for insert user data
func Createrecordtermsmap(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstRecordtermsmapEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		success, _, msg := models.InsertRecordtermsMap(&data)
		ThrowDifferentiationMapResponse(msg, w, success)
	}
}
