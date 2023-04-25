package handlers

import (
	"ifixRecord/ifix/entities"
	"ifixRecord/ifix/logger"
	"ifixRecord/ifix/models"
	"net/http"
)

func UpdateRecordfulTicketIdDeleteFlag(w http.ResponseWriter, req *http.Request) {
	logger.Log.Println("Inside UpdateRecordfulTicketIdDeleteFlag Handler function")
	var data = entities.RecordfulTicketIdDeleteFlagUpdateEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.UpdateRecordfulTicketIdDeleteFlag(&data)
		ThrowRecordstatusIntResponse(msg, 0, w, success)
	}
}

func UpdateTrnRecordCodeDeleteFlgById(w http.ResponseWriter, req *http.Request) {
	logger.Log.Println("Inside UpdateTrnrecordCodeDeleteFlgById Handler Function")
	var data = entities.TrnRecordCodeDeleteFlgUpdateByIdEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.UpdateTrnRecordCodeDeleteFlgById(&data)
		ThrowRecordstatusIntResponse(msg, 0, w, success)
	}

}
