package handlers

import (
	"encoding/json"
	"net/http"
	"src/entities"
	"src/logger"
	"src/models"
)

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
func UpdateIfixsysid(w http.ResponseWriter, req *http.Request) {
	// var data = entities.TransporttableEntity{}
	// jsonError := data.FromJSON(req.Body)
	// if jsonError != nil {
	// 	log.Print(jsonError)
	// 	logger.Log.Println(jsonError)
	// 	entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	// } else {
	success, _, msg := models.UpdateIfixsysid()
	logger.Log.Println(success)
	ThrowTransporttableIntResponse(msg, 0, w, success)
	// }
}
