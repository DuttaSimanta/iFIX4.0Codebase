package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowRecordfulldetailsToExcelResponse(successMessage string, OriginalFileName string, UploadedFileName string, w http.ResponseWriter, success bool) {
	var response = entities.APIResponseDownload{}
	response.Status = success
	response.Message = successMessage
	response.OriginalFileName = OriginalFileName
	response.UploadedFileName = UploadedFileName
	// response.RequestResultsetData = v
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		logger.Log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
func ThrowRecordfulldetailsResponse(successMessage string, result entities.ResultEntity, w http.ResponseWriter, success bool) {
	var response = entities.RecordfulldetailsResponse{}
	response.Success = success
	response.Message = successMessage
	response.Details = result
	// response.UploadedFileName = UploadedFileName
	// response.RequestResultsetData = v
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		logger.Log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
func ThrowGetRecordByDiffTypeOfMultiOrgResponse(successMessage string, result []entities.RecordDiffOfMultiOrgEntity, w http.ResponseWriter, success bool) {
	var response = entities.RecordDiffOfMultiOrgResponse{}
	response.Success = success
	response.Message = successMessage
	response.Details = result
	// response.UploadedFileName = UploadedFileName
	// response.RequestResultsetData = v
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		logger.Log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
func RecordfulldetailsToExcelConverter(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecordfulldetailsRequestEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		OriginalFileName, UploadedFileName, success, _, msg := models.RecordfulldetailsToExcelConverter(&data)
		ThrowRecordfulldetailsToExcelResponse(msg, OriginalFileName, UploadedFileName, w, success)
	}
}
func Getrecordfulldetails(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecordfulldetailsRequestEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		result, success, _, msg := models.Getrecordfulldetailss(&data)
		ThrowRecordfulldetailsResponse(msg, result, w, success)
	}
}
func GetRecordByDiffTypeOfMultiOrg(w http.ResponseWriter, req *http.Request) {

	var data = entities.RecordfulldetailsRequestEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print("data", jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.GetRecordByDiffTypeOfMultiOrg(&data)
		ThrowGetRecordByDiffTypeOfMultiOrgResponse(msg, data, w, success)
	}
}
