package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

// ThrowRecordDiffTypeResponse function is used to throw success response of All data in JSON format
func ThrowUseractionnameResponse(successMessage string, responseData []int64, w http.ResponseWriter, success bool) {
	var response = entities.UserroleactionnameEntityResponse{}
	response.Success = success
	response.Message = successMessage
	response.Values = responseData
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		logger.Log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func Getuserwiseaction(w http.ResponseWriter, req *http.Request) {
	var data = entities.UserroleactionnameEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetUserActionnamewithapi(&data)
		ThrowUseractionnameResponse(msg, data, w, success)
	}
}
