package handlers

import (
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"

	//"log"
	"encoding/json"
	"net/http"
)

// ThrowRecordDiffTypeResponse function is used to throw success response of All data in JSON format
func ThrowRecordDiffTypeResponse(successMessage string, responseData []entities.MstrecorddifftypeEntity, w http.ResponseWriter, success bool) {
	var response = entities.RecordDiffTypeResponse{}
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

// ThrowRecordDiffResponse function is used to throw success response of All data in JSON format
func ThrowRecordDiffResponse(successMessage string, responseData entities.RecordDiffEntities, w http.ResponseWriter, success bool) {
	var response = entities.RecordDiffResponse{}
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

// ThrowRecordDiffIntResponse function is used to throw success response of integer data in JSON format
func ThrowRecordDiffIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.ApiResponseInt{}
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

func GetRecordDiffType(w http.ResponseWriter, req *http.Request) {
	data, success, _, msg := models.GetRecordDiffType()
	ThrowRecordDiffTypeResponse(msg, data, w, success)
}
func GetRecordByDiffType(w http.ResponseWriter, req *http.Request) {

	var data = entities.RecordDiffEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print("data", jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.GetRecordByDiffType(&data)
		ThrowRecordDiffTypeResponse(msg, data, w, success)
	}
}
func InsertRecordDiff(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecordDiffEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data, success, _, msg := models.InsertRecordDiff(&data)
		ThrowRecordDiffIntResponse(msg, data, w, success)
	}
}
func GetAllRecordDiff(w http.ResponseWriter, req *http.Request) {

	var data = entities.RecordDiffEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print("data", jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.GetAllRecordDiff(&data)
		ThrowRecordDiffResponse(msg, data, w, success)
	}
}
func GetRecordDiff(w http.ResponseWriter, req *http.Request) {

	var data = entities.RecordDiffEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print("data", jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.GetRecordDiff(&data)
		ThrowRecordDiffResponse(msg, data, w, success)
	}
}
func GetRecordDiffByOrg(w http.ResponseWriter, req *http.Request) {

	var data = entities.RecordDiffEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print("data", jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.GetRecordDiffByOrg(&data)
		ThrowRecordDiffResponse(msg, data, w, success)
	}
}

func GetAssetRecordDiffByOrg(w http.ResponseWriter, req *http.Request) {

	var data = entities.RecordDiffEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print("data", jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.GetAssetRecordDiffByOrg(&data)
		ThrowRecordDiffResponse(msg, data, w, success)
	}
}

func DeleteRecordDiff(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecordDiffEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		success, _, msg := models.DeleteRecordDiff(&data)
		ThrowRecordDiffIntResponse(msg, 0, w, success)
	}
}
func UpdateRecordDiff(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecordDiffEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		success, _, msg := models.UpdateRecordDiff(&data)
		ThrowRecordDiffIntResponse(msg, 0, w, success)
	}
}

func UpdateAssetRecordDiff(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecordDiffEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		success, _, msg := models.UpdateAssetRecordDiff(&data)
		ThrowRecordDiffIntResponse(msg, 0, w, success)
	}
}

func GetCategoryLevel(w http.ResponseWriter, req *http.Request) {

	var data = entities.RecordDiffEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print("data", jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.GetCategoryLevel(&data)
		ThrowRecordDiffTypeResponse(msg, data, w, success)
	}
}
func GetCategoriesLevel(w http.ResponseWriter, req *http.Request) {

	var data = entities.RecordDiffEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print("data", jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.GetCategoriesLevel(&data)
		ThrowRecordDiffTypeResponse(msg, data, w, success)
	}
}
func GetAllCategoryLevel(w http.ResponseWriter, req *http.Request) {

	var data = entities.RecordDiffEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print("data", jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.GetAllCategoryLevel(&data)
		ThrowRecordDiffTypeResponse(msg, data, w, success)
	}
}
func GetAllRecordDiffTypeByClient(w http.ResponseWriter, req *http.Request) {

	var data = entities.RecordDiffEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print("data", jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.GetAllRecordDiffTypeByClient(&data)
		ThrowRecordDiffTypeResponse(msg, data, w, success)
	}
}
