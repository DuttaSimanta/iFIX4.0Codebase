package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowOpenTickeAllResponse(successMessage string, responseData entities.OpenTicketEntities, w http.ResponseWriter, success bool) {
	var response = entities.OpenTicketResponse{}
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

func ThrowOpenIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.OpenTicketResponseInt{}
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

func GetOpenTicket(w http.ResponseWriter, req *http.Request) {
	var data = entities.OpenTicketEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetOpenTicket(&data)
		ThrowOpenTickeAllResponse(msg, data, w, success)
	}
}

func DeleteOpenTicket(w http.ResponseWriter, req *http.Request) {
	var data = entities.OpenTicketEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteOpenTicket(&data)
		ThrowOpenIntResponse(msg, 0, w, success)
	}
}
