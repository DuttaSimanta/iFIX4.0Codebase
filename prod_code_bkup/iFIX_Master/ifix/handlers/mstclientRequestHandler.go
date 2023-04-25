//***************************//
// Package handlers
// Date Of Creation: 16/12/2020
// Authour Name: Moloy Mondal
// History: N/A
// Synopsis: This file is used to do crud operation of mstclient table. It is used as Controller
// Functions: AddMstClient, GetMstClientById
// ThrowMstClientResponse() Parameter:  (<string>,<Structure entities.MstClient>,<http.ResponseWriter>)
// AddMstClient() Parameter:  (<http.ResponseWriter>, <*http.Request>)
// GetMstClientById() Parameter:  {<http.ResponseWriter>, <*http.Request>}
// Global Variable: N/A
// Version: 1.0.0
//***************************//

package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"iFIX/ifix/validators"
	"net/http"
)

// ThrowMstClientResponse function is used to throw success response in JSON format
func ThrowMstClientResponse(successMessage string, mstClientResponseData entities.MstClient, w http.ResponseWriter) {
	var response = entities.MstClientResponse{}
	response.Status = true
	response.Message = successMessage
	response.Response = mstClientResponseData
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		logger.Log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// ThrowMstClientAllResponse function is used to throw success response of All data in JSON format
func ThrowMstClientAllResponse(successMessage string, mstClientResponseData []entities.MstClient, w http.ResponseWriter) {
	var response = entities.MstClientAllResponse{}
	response.Status = true
	response.Message = successMessage
	response.Response = mstClientResponseData
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		logger.Log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// AddMstClient function is used to handle mstclient insert operation
func AddMstClient(w http.ResponseWriter, req *http.Request) {
	var mstClientData = entities.MstClient{}
	var responseError []string
	jsonError := mstClientData.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		responseError = validators.ValidateAddMstClient(mstClientData)

		if len(responseError) == 0 {
			mstClientResponseData, modelResponseError := models.InsertMstClient(&mstClientData)
			if modelResponseError != nil {
				entities.ThrowJSONResponse(entities.DbErrorResponse(), w)
			} else {
				ThrowMstClientResponse("mstclient data inserted successfully.", mstClientResponseData, w)
			}
		} else {
			entities.ThrowErrorResponse(responseError, w)
		}
	}
}

// GetMstClientByID function is used to fetch mstclient data by id
func GetMstClientByID(w http.ResponseWriter, req *http.Request) {
	var mstClientData = entities.MstClient{}
	jsonError := mstClientData.FromJSON(req.Body)

	var responseError []string
	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		responseError = validators.ValidateMstClientId(mstClientData)

		if len(responseError) == 0 {
			mstClientResponseData, err := models.GetMstClientByID(&mstClientData)
			if err != nil {
				logger.Log.Println(err)
				entities.ThrowJSONResponse(entities.DbErrorResponse(), w)
			} else {
				ThrowMstClientResponse("mstclient data fetched successfully.", mstClientResponseData, w)
			}
		} else {
			entities.ThrowErrorResponse(responseError, w)
		}
	}
}

// GetMstClientAll function is used to fetch mstclient data
func GetMstClientAll(w http.ResponseWriter, req *http.Request) {

	mstClientResponseData, err := models.GetMstClientAll()
	if err != nil {
		logger.Log.Println(err)
		entities.ThrowJSONResponse(entities.DbErrorResponse(), w)
	} else {
		ThrowMstClientAllResponse("mstclient data fetched successfully.", mstClientResponseData, w)
	}
}

// DelMstClientByID function is used to delete mstclient data by id
func DelMstClientByID(w http.ResponseWriter, req *http.Request) {
	var mstClientData = entities.MstClient{}
	jsonError := mstClientData.FromJSON(req.Body)

	var responseError []string
	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		responseError = validators.ValidateMstClientId(mstClientData)

		if len(responseError) == 0 {
			mstClientResponseData, err := models.DelMstClientByID(&mstClientData)
			if err != nil {
				logger.Log.Println(err)
				entities.ThrowJSONResponse(entities.DbErrorResponse(), w)
			} else {
				ThrowMstClientResponse("mstclient data deleted successfully.", mstClientResponseData, w)
			}
		} else {
			entities.ThrowErrorResponse(responseError, w)
		}
	}
}

// UpdateMstClientByID function is used to handle mstclient update operation by id
func UpdateMstClientByID(w http.ResponseWriter, req *http.Request) {
	var mstClientData = entities.MstClient{}
	var responseError []string
	jsonError := mstClientData.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		responseError = validators.ValidateUpdateMstClient(mstClientData)

		if len(responseError) == 0 {
			mstClientResponseData, modelResponseError := models.UpdateMstClientByID(&mstClientData)
			if modelResponseError != nil {
				entities.ThrowJSONResponse(entities.DbErrorResponse(), w)
			} else {
				ThrowMstClientResponse("mstclient data updated successfully.", mstClientResponseData, w)
			}

		} else {
			entities.ThrowErrorResponse(responseError, w)
		}
	}
}
