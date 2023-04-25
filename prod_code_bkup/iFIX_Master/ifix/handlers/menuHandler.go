package handlers

import (
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"

	//"log"
	"net/http"
	"encoding/json"
)

// ThrowMenuSingleResponse function is used to throw success response of All data in JSON format
func ThrowMenuSingleResponse(successMessage string, responseData []entities.MenuSingleEntity, w http.ResponseWriter, success bool) {
	var response = entities.MenuEntitySingleResponse{}
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

// ThrowMenuResponse function is used to throw success response of All data in JSON format
func ThrowMenuResponse(successMessage string, responseData entities.MenuEntities, w http.ResponseWriter, success bool) {
	var response = entities.MenuEntityResponse{}
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

// ThrowMenuIntResponse function is used to throw success response of integer data in JSON format
func ThrowMenuIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.ApiResponseInt{}
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
// ThrowMenuByUserResponse function is used to throw success response of multiple MenuHierarchyEntity data in JSON format
func ThrowMenuByUserResponse(successMessage string, responseData []entities.MenuHierarchyEntity, w http.ResponseWriter, success bool) {
	var response = entities.MenuHierarchyResponse{}
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


func GetMenuByUser(w http.ResponseWriter, req *http.Request) {

	var data = entities.MenuByUserRequest{}

	jsonError := data.FromJSONMenuByUserRequest(req.Body)

	if jsonError != nil {
		log.Print("data",jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	}else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		var data1 = []entities.MenuHierarchyEntity{}
		data1, success, _, msg := models.GetMenuByUser(&data)
		ThrowMenuByUserResponse(msg, data1, w, success)
	}
}
func SearchMenuByUser(w http.ResponseWriter, req *http.Request) {

	var data = entities.MenuByUserRequest{}

	jsonError := data.FromJSONMenuByUserRequest(req.Body)

	if jsonError != nil {
		log.Print("data",jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	}else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		var data1 = []entities.MenuHierarchyEntity{}
		data1, success, _, msg := models.SearchMenuByUser(&data)
		ThrowMenuByUserResponse(msg, data1, w, success)
	}
}
func InsertMenu(w http.ResponseWriter, req *http.Request) {
	var data = entities.MenuEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data, success, _, msg := models.InsertMenuwithapi(&data)
		ThrowMenuIntResponse(msg, data, w, success)
	}
}
func Getparentmenu(w http.ResponseWriter, req *http.Request) {

	var data = entities.MenuEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print("data",jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	}else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.Getparentmenu(&data)
		ThrowMenuSingleResponse(msg, data, w, success)
	}
}
func Getmenubymodule(w http.ResponseWriter, req *http.Request) {

	var data = entities.MenuEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print("data",jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	}else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.Getmenubymodule(&data)
		ThrowMenuSingleResponse(msg, data, w, success)
	}
}
func Getmenudetails(w http.ResponseWriter, req *http.Request) {

	var data = entities.PaginationEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print("data",jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	}else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.Getmenudetails(&data)
		ThrowMenuResponse(msg, data, w, success)
	}
}
func UpdateMenu(w http.ResponseWriter, req *http.Request) {
	var data = entities.MenuEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		success, _, msg := models.UpdateMenu(&data)
		ThrowMenuIntResponse(msg, 0, w, success)
	}
}
func DeleteUrlFromMenu(w http.ResponseWriter, req *http.Request) {
	var data = entities.MenuEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		success, _, msg := models.DeleteUrlFromMenu(&data)
		ThrowMenuIntResponse(msg, 0, w, success)
	}
}
func DeleteMenu(w http.ResponseWriter, req *http.Request) {
	var data = entities.MenuEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		success, _, msg := models.DeleteMenu(&data)
		ThrowMenuIntResponse(msg, 0, w, success)
	}
}

func Geturlmenudetails(w http.ResponseWriter, req *http.Request) {

	var data = entities.PaginationEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print("data",jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	}else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.Geturlmenudetails(&data)
		ThrowMenuResponse(msg, data, w, success)
	}
}
