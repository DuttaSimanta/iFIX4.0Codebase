package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowCatalogwithcategoryAllResponse(successMessage string, responseData entities.CatalogwithcategoryEntities, w http.ResponseWriter, success bool) {
	var response = entities.CatalogwithcategoryResponse{}
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
func ThrowCatalogwithcategorySingleResponse(successMessage string, responseData []entities.CatalogwithsingleEntity, w http.ResponseWriter, success bool) {
	var response = entities.CatalogwithcategorysingleResponse{}
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
func ThrowCatalogwithcategoryIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
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

func InsertCatalogwithcategory(w http.ResponseWriter, req *http.Request) {
	var data = entities.CatalogwithcategoryEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.InsertCatalogwithcategory(&data)
		ThrowCatalogwithcategoryIntResponse(msg, data, w, success)
	}
}

func Getcategorybycatalog(w http.ResponseWriter, req *http.Request) {
	var data = entities.CatalogwithcategoryEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.Getcategorybycatalog(&data)
		ThrowCatalogwithcategorySingleResponse(msg, data, w, success)
	}
}
func Getcategorybyparentname(w http.ResponseWriter, req *http.Request) {
	var data = entities.CatalogwithcategoryEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.Getcategorybyparentname(&data)
		ThrowCatalogwithcategorySingleResponse(msg, data, w, success)
	}
}
func Getfromtypebydiffname(w http.ResponseWriter, req *http.Request) {
	var data = entities.CatalogwithcategoryEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.Getfromtypebydiffname(&data)
		ThrowCatalogwithcategorySingleResponse(msg, data, w, success)
	}
}
func GetAllCatalogwithcategory(w http.ResponseWriter, req *http.Request) {
	var data = entities.CatalogwithcategoryEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllCatalogwithcategory(&data)
		ThrowCatalogwithcategoryAllResponse(msg, data, w, success)
	}
}

func DeleteCatalogwithcategory(w http.ResponseWriter, req *http.Request) {
	var data = entities.CatalogwithcategoryEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteCatalogwithcategory(&data)
		ThrowCatalogwithcategoryIntResponse(msg, 0, w, success)
	}
}

func UpdateCatalogwithcategory(w http.ResponseWriter, req *http.Request) {
	var data = entities.CatalogwithcategoryEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.UpdateCatalogwithcategory(&data)
		ThrowCatalogwithcategoryIntResponse(msg, 0, w, success)
	}
}
