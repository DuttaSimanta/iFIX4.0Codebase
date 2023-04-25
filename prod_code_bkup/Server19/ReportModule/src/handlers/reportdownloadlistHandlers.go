package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"src/entities"
	"src/logger"
	"src/models"
)

func ThrowGetDownloadListResponse(successMessage string, responseData entities.ReportDownloadListEntities, w http.ResponseWriter, success bool) {
	var response = entities.ReportDownloadListResponse{}
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

func GetDownloadList(w http.ResponseWriter, req *http.Request) {
	var data = entities.ReportDownloadListEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetDownloadList(&data)
		ThrowGetDownloadListResponse(msg, data, w, success)
	}
}



func GetReportDockerStatus(w http.ResponseWriter, req *http.Request) {
	// var data = entities.ReportDownloadListEntity{}
	// jsonError := data.FromJSON(req.Body)
	// if jsonError != nil {
	// 	log.Print(jsonError)
	// 	logger.Log.Println(jsonError)
	// 	entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	// } else {
	// 	data, success, _, msg := models.GetDownloadList(&data)
	// 	ThrowGetDownloadListResponse(msg, data, w, success)
	// }
	type response struct {
		Success bool `json:"success"`
	}
	var res response
	res.Success = true
	jsonResponse, jsonError := json.Marshal(res)
	if jsonError != nil {
		logger.Log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
