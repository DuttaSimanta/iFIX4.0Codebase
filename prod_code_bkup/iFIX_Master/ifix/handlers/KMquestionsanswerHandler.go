package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowQuestionsanswerAllResponse(successMessage string, responseData entities.KMquestionsanswerEntities, w http.ResponseWriter, success bool) {
	var response = entities.KMquestionsanswerResponse{}
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

func ThrowQuestionsanswerIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.KMquestionsanswerResponseInt{}
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

func InsertQuestionanswer(w http.ResponseWriter, req *http.Request) {
	var data = entities.KMquestionsanswerEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.InsertQuestionanswer(&data)
		ThrowQuestionsanswerIntResponse(msg, data, w, success)
	}
}

func GetAllquestionsanswer(w http.ResponseWriter, req *http.Request) {
	var data = entities.KMquestionsanswerEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllQuestionAnswers(&data)
		ThrowQuestionsanswerAllResponse(msg, data, w, success)
	}
}

func Deletequestionsanswer(w http.ResponseWriter, req *http.Request) {
	var data = entities.KMquestionsanswerEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteQuetionanswer(&data)
		ThrowQuestionsanswerIntResponse(msg, 0, w, success)
	}
}
