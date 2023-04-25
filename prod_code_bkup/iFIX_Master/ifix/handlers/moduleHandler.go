package handlers


import (
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
	"encoding/json"
)


// ThrowModuleAllResponse function is used to throw success response of All data in JSON format
func ThrowModuleAllResponse(successMessage string, responseData entities.ModuleEntities, w http.ResponseWriter, success bool) {
	var response = entities.ModuleResponse{}
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
// ThrowModuleIntResponse function is used to throw success response of integer data in JSON format
func ThrowModuleIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.ModuleResponseInt{}
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

func InsertModule(w http.ResponseWriter, req *http.Request) {
	var data = entities.ModuleEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data, success, _, msg := models.InsertModule(&data)
		ThrowModuleIntResponse(msg, data, w, success)
	}
}

func GetAllModules(w http.ResponseWriter, req *http.Request) {
	var data = entities.PaginationEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data, success, _, msg := models.GetAllModules(&data)
		ThrowModuleAllResponse(msg, data, w, success)
	}
}
func DeleteModule(w http.ResponseWriter, req *http.Request) {
	var data = entities.ModuleEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		success, _, msg := models.DeleteModule(&data)
		ThrowModuleIntResponse(msg, 0, w, success)
	}
}
func UpdateModule(w http.ResponseWriter, req *http.Request) {
	var data = entities.ModuleEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		success, _, msg := models.UpdateModule(&data)
		ThrowModuleIntResponse(msg, 0, w, success)
	}
}
