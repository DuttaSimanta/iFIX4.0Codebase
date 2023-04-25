package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

// ThrowModuleAllResponse function is used to throw success response of All data in JSON format
func ThrowTicketCustomerAllResponse(successMessage string, responseData entities.TicketCustomerEntities, w http.ResponseWriter, success bool) {
	var response = entities.TicketCustomerResponse{}
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

func GetAllTicketCustomer(w http.ResponseWriter, req *http.Request) {
	var data = entities.TicketCustomerEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data, success, _, msg := models.GetAllTicketCustomer(&data)
		ThrowTicketCustomerAllResponse(msg, data, w, success)
	}
}
