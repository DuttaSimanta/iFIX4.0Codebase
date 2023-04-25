package handlers


import (
  "iFIX/ifix/entities"
  "iFIX/ifix/models"
  "iFIX/ifix/logger"
  "log"
  "net/http"
  "encoding/json"
  )

func ThrowDashboarddtlsAllResponse(successMessage string, responseData entities.DashboarddtlsEntities, w http.ResponseWriter, success bool) {
    var response = entities.DashboarddtlsResponse{}
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


func ThrowDashboarddtlsIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
    var response = entities.DashboarddtlsResponseInt{}
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


func InsertDashboarddtls(w http.ResponseWriter, req *http.Request) {
    var data = entities.DashboarddtlsEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.InsertDashboarddtls(&data)
        ThrowDashboarddtlsIntResponse(msg, data, w, success)
    }
}


func GetAllDashboarddtls(w http.ResponseWriter, req *http.Request) {
    var data = entities.DashboarddtlsEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.GetAllDashboarddtls(&data)
        ThrowDashboarddtlsAllResponse(msg, data, w, success)
    }
}


func DeleteDashboarddtls(w http.ResponseWriter, req *http.Request) {
    var data = entities.DashboarddtlsEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.DeleteDashboarddtls(&data)
        ThrowDashboarddtlsIntResponse(msg, 0, w, success)
    }
}


func UpdateDashboarddtls(w http.ResponseWriter, req *http.Request) {
    var data = entities.DashboarddtlsEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.UpdateDashboarddtls(&data)
        ThrowDashboarddtlsIntResponse(msg, 0, w, success)
    }
}


