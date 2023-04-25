 package handlers


import (
  "iFIX/ifix/entities"
  "iFIX/ifix/models"
  "iFIX/ifix/logger"
  "log"
  "net/http"
  "encoding/json"
  )

func ThrowMstClientCredentialAllResponse(successMessage string, responseData entities.MstClientCredentialEntities, w http.ResponseWriter, success bool) {
    var response = entities.MstClientCredentialResponse{}
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


func ThrowMstClientCredentialIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
    var response = entities.MstClientCredentialResponseInt{}
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


func InsertMstClientCredential(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstClientCredentialEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.InsertMstClientCredential(&data)
        ThrowMstClientCredentialIntResponse(msg, data, w, success)
    }
}


func GetAllMstClientCredential(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstClientCredentialEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.GetAllMstClientCredential(&data)
        ThrowMstClientCredentialAllResponse(msg, data, w, success)
    }
}


func DeleteMstClientCredential(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstClientCredentialEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.DeleteMstClientCredential(&data)
        ThrowMstClientCredentialIntResponse(msg, 0, w, success)
    }
}


func UpdateMstClientCredential(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstClientCredentialEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.UpdateMstClientCredential(&data)
        ThrowMstClientCredentialIntResponse(msg, 0, w, success)
    }
}




