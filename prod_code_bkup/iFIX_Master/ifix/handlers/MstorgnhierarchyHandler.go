package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"net/http"
)

// ThrowOrganizationAllResponse function is used to throw success response of All data in JSON format
func ThrowOrganizationAllResponse(successMessage string, responseData entities.MstorgnhierarchyEntities, w http.ResponseWriter, success bool) {
	var response = entities.MstorgnhierarchyEntityResponse{}
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

// ThrowOrganizationSingleResponse function is used to throw success response of All data in JSON format
func ThrowOrganizationSingleResponse(successMessage string, responseData []entities.MstorgnhierarchyEntityResp, w http.ResponseWriter, success bool) {
	var response = entities.MstorgnhierarchyResponse{}
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

// ThrowOrganizationIntResponse function is used to throw success response of integer data in JSON format
func ThrowOrganizationIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.MstorgnhierarchyEntityResponseInt{}
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

func ThrowLogintypeResponse(successMessage string, responseData []entities.LogintypeEntity, w http.ResponseWriter, success bool) {
	var response = entities.LogintypeEntityResponse{}
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

func ThrowOrganizationFieldAllResponse(successMessage string, responseData entities.MstorgnhierarchyFieldEntities, w http.ResponseWriter, success bool) {
	var response = entities.MstorgnhierarchyFieldEntityResponse{}
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
func UpdateOrganizationField(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstorgnhierarchyFieldEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		success, _, msg := models.UpdateOrganizationField(&data)
		ThrowOrganizationIntResponse(msg, 0, w, success)
	}
}
func GetAllOrganizationField(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstorgnhierarchyFieldEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data1, success, _, msg := models.GetAllOrganizationField(&data)
		ThrowOrganizationFieldAllResponse(msg, data1, w, success)
	}
}
//AddOrganization method is used for insert client data
func AddOrganization(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstorgnhierarchyEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data, success, _, msg := models.AddOrganizations(&data)
		ThrowOrganizationIntResponse(msg, data, w, success)
	}
}

//UpdateOrganization  method is used for update client data
func UpdateOrganization(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstorgnhierarchyEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		success, _, msg := models.UpdateOrganizations(&data)
		ThrowOrganizationIntResponse(msg, 0, w, success)
	}
}

//GetAllOrganization  method is used for get clients data
func GetAllOrganization(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstorgnhierarchyEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data1, success, _, msg := models.GetAllOrganizations(&data)
		ThrowOrganizationAllResponse(msg, data1, w, success)
	}
}

//GetAllOrganizationClientWise  method is used for get clients data
func GetAllOrganizationClientWise(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstorgnhierarchyEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data1, success, _, msg := models.GetAllOrganizationsClientWise(&data)
		ThrowOrganizationSingleResponse(msg, data1, w, success)
	}
}

func GetAllOrganizationClientWisenew(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstorgnhierarchyEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data1, success, _, msg := models.GetAllOrganizationsClientWisenew(&data)
		ThrowOrganizationSingleResponse(msg, data1, w, success)
	}
}

func Gettimeformat(w http.ResponseWriter, req *http.Request) {
	data1, success, _, msg := models.Gettimeformat()
	ThrowOrganizationSingleResponse(msg, data1, w, success)
}

func GetLogintype(w http.ResponseWriter, req *http.Request) {
	data1, success, _, msg := models.GetLogintype()
	ThrowLogintypeResponse(msg, data1, w, success)
}
