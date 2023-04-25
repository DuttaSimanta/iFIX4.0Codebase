//***************************//
// Package handlers
// Date Of Creation: 16/12/2020
// Authour Name: Moloy Mondal
// History: N/A
// Synopsis: This file is used to do crud operation of mstcountry table. It is used as Controller
// Functions: ValidateAddMstCountry,AddMstCountry, GetMstCountryById
// ThrowMstCountryResponse() Parameter:  (<string>,<Structure entities.MstCountry>,<http.ResponseWriter>)
// AddMstCountry() Parameter:  (<http.ResponseWriter>, <*http.Request>)
// GetMstCountryById() Parameter:  {<http.ResponseWriter>, <*http.Request>}
// Global Variable: N/A
// Version: 1.0.0
//***************************//

package handlers

import (
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"iFIX/ifix/validators"
	"net/http"
	"encoding/json"
)

// ThrowMstCountryResponse function is used to throw success response in JSON format
func ThrowMstCountryResponse(successMessage string,mstCountryResponseData entities.MstCountry,w http.ResponseWriter){
	var response = entities.MstCountryResponse{}
	response.Status	= true
	response.Message = successMessage
	response.Response = mstCountryResponseData
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		logger.Log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// ThrowMstCountryAllResponse function is used to throw success response of All data in JSON format
func ThrowMstCountryAllResponse(successMessage string,mstCountryResponseData []entities.MstCountry,w http.ResponseWriter){
	var response = entities.MstCountryAllResponse{}
	response.Status	= true
	response.Message = successMessage
	response.Response = mstCountryResponseData
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		logger.Log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// AddMstCountry function is used to handle mstcountry insert operation
func AddMstCountry(w http.ResponseWriter, req *http.Request){
	var mstCountryData = entities.MstCountry{}
	var responseError []string
	jsonError := mstCountryData.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(),w)	
	} else {		
		responseError = validators.ValidateAddMstCountry(mstCountryData)

		if(len(responseError)==0){
			
			mstCountryResponseData,modelResponseError:= models.InsertMstCountry(&mstCountryData)
			if modelResponseError!=nil{
				entities.ThrowJSONResponse(entities.DbErrorResponse(),w)	
			} else {
				ThrowMstCountryResponse("mstcountry data inserted successfully.",mstCountryResponseData,w)
			}
			
		} else {
			entities.ThrowErrorResponse(responseError,w)
		}
	}
}


// GetMstCountryByID function is used to fetch mstcountry data by id
func GetMstCountryByID(w http.ResponseWriter, req *http.Request){
	var mstCountryData = entities.MstCountry{}
	jsonError := mstCountryData.FromJSON(req.Body)

	var responseError []string
	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(),w)	
	} else {
		responseError = validators.ValidateMstCountryId(mstCountryData)

		if(len(responseError)==0){
			
			mstCountryResponseData,err := models.GetMstCountryByID(&mstCountryData)
			if err != nil {
				logger.Log.Println(err)
				entities.ThrowJSONResponse(entities.DbErrorResponse(),w)
			} else {
				ThrowMstCountryResponse("mstcountry data fetched successfully.",mstCountryResponseData,w)
			}

		} else {
			entities.ThrowErrorResponse(responseError,w)
		}
	}
}

// GetMstCountryAll function is used to fetch mstcountry data 
func GetMstCountryAll(w http.ResponseWriter, req *http.Request){
	
	mstCountryResponseData,err := models.GetMstCountryAll()
	if err != nil {
		logger.Log.Println(err)
		entities.ThrowJSONResponse(entities.DbErrorResponse(),w)
	} else {
		ThrowMstCountryAllResponse("mstcountry data fetched successfully.",mstCountryResponseData,w)
	}
	
}

// DelMstCountryByID function is used to delete mstcountry data by id
func DelMstCountryByID(w http.ResponseWriter, req *http.Request){
	var mstCountryData = entities.MstCountry{}
	jsonError := mstCountryData.FromJSON(req.Body)

	var responseError []string
	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(),w)	
	} else {
		responseError = validators.ValidateMstCountryId(mstCountryData)

		if(len(responseError)==0){
			mstCountryResponseData,err := models.DelMstCountryByID(&mstCountryData)
			if err != nil {
				logger.Log.Println(err)
				entities.ThrowJSONResponse(entities.DbErrorResponse(),w)
			} else {
				ThrowMstCountryResponse("mstcountry data deleted successfully.",mstCountryResponseData,w)
			}
		} else {
			entities.ThrowErrorResponse(responseError,w)
		}
	}
}

// UpdateMstCountryByID function is used to handle mstcountry update operation by id
func UpdateMstCountryByID(w http.ResponseWriter, req *http.Request){
	var mstCountryData = entities.MstCountry{}
	var responseError []string
	jsonError := mstCountryData.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(),w)	
	} else {		
		responseError = validators.ValidateUpdateMstCountry(mstCountryData)

		if(len(responseError)==0){

			mstCountryResponseData,modelResponseError:= models.UpdateMstCountryByID(&mstCountryData)
			if modelResponseError!=nil{
				entities.ThrowJSONResponse(entities.DbErrorResponse(),w)	
			} else {
				ThrowMstCountryResponse("mstcountry data updated successfully.",mstCountryResponseData,w)
			}
			
		} else {
			entities.ThrowErrorResponse(responseError,w)
		}
	}
}


