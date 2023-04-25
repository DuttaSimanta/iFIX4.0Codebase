package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"src/entities"
	"src/logger"
	"src/models"
)

func ThrowGetDynamicQueryCountResultResponse(successMessage string, recordResponseData map[string]interface{}, w http.ResponseWriter, success bool) {
	var response = entities.QueryCountResultResponeData{}
	response.Status = success
	response.Message = successMessage
	response.Response = recordResponseData
	jsonResponse, jsonError := json.Marshal(response)
	//var result map[string]interface{}
	//json.Unmarshal(jsonResponse, &result)
	if jsonError != nil {
		logger.Log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonResponse)
}

func RecordGridResultOnly(w http.ResponseWriter, req *http.Request) {
	var reqData map[string]interface{}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.Log.Println(err)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		json.Unmarshal(body, &reqData)
		response, success, _, msg := models.RecordGridResultOnly(reqData)
		ThrowGetDynamicQueryCountResultResponse(msg, response, w, success)
	}
}
