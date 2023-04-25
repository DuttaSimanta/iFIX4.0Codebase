package handlers


import (
  "iFIX/ifix/entities"
  "iFIX/ifix/models"
  "iFIX/ifix/logger"
  "log"
  "net/http"
  "encoding/json"
  )

func ThrowMstdocumentdtlsAllResponse(successMessage string, responseData entities.MstdocumentdtlsEntities, w http.ResponseWriter, success bool) {
    var response = entities.MstdocumentdtlsResponse{}
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


func ThrowMstdocumentdtlsIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
    var response = entities.MstdocumentdtlsResponseInt{}
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


func InsertMstdocumentdtls(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstdocumentdtlsEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.InsertMstdocumentdtls(&data)
        ThrowMstdocumentdtlsIntResponse(msg, data, w, success)
    }
}


func GetAllMstdocumentdtls(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstdocumentdtlsEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.GetAllMstdocumentdtls(&data)
        ThrowMstdocumentdtlsAllResponse(msg, data, w, success)
    }
}


func DeleteMstdocumentdtls(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstdocumentdtlsEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.DeleteMstdocumentdtls(&data)
        ThrowMstdocumentdtlsIntResponse(msg, 0, w, success)
    }
}


func UpdateMstdocumentdtls(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstdocumentdtlsEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.UpdateMstdocumentdtls(&data)
        ThrowMstdocumentdtlsIntResponse(msg, 0, w, success)
    }
}


