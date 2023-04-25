package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"net/http"
)

func ThrowJsonToExcelResponse(successMessage string, OriginalFileName string, UploadedFileName string, w http.ResponseWriter, success bool) {
	var response = entities.APIResponseDownload{}
	response.Status = success
	response.Message = successMessage
	response.OriginalFileName = OriginalFileName
	response.UploadedFileName = UploadedFileName
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		logger.Log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
func JsonToExcelConverter(w http.ResponseWriter, req *http.Request) {
	var data = entities.ResultSetRequestEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		OriginalFileName, UploadedFileName, success, _, msg := models.JsonToExcelConverter(&data)
		ThrowJsonToExcelResponse(msg, OriginalFileName, UploadedFileName, w, success)
	}
}
