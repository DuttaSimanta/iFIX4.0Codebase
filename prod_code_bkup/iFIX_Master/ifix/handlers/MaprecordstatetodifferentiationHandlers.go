package handlers


import (
  "iFIX/ifix/entities"
  "iFIX/ifix/models"
  "iFIX/ifix/logger"
  "log"
  "net/http"
  "encoding/json"
  )

func ThrowMaprecordstatetodifferentiationAllResponse(successMessage string, responseData entities.MaprecordstatetodifferentiationEntities, w http.ResponseWriter, success bool) {
    var response = entities.MaprecordstatetodifferentiationResponse{}
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


func ThrowMaprecordstatetodifferentiationIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
    var response = entities.MaprecordstatetodifferentiationResponseInt{}
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


func InsertMaprecordstatetodifferentiation(w http.ResponseWriter, req *http.Request) {
    var data = entities.MaprecordstatetodifferentiationEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.InsertMaprecordstatetodifferentiation(&data)
        ThrowMaprecordstatetodifferentiationIntResponse(msg, data, w, success)
    }
}


func GetAllMaprecordstatetodifferentiation(w http.ResponseWriter, req *http.Request) {
    var data = entities.MaprecordstatetodifferentiationEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.GetAllMaprecordstatetodifferentiation(&data)
        ThrowMaprecordstatetodifferentiationAllResponse(msg, data, w, success)
    }
}


func DeleteMaprecordstatetodifferentiation(w http.ResponseWriter, req *http.Request) {
    var data = entities.MaprecordstatetodifferentiationEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.DeleteMaprecordstatetodifferentiation(&data)
        ThrowMaprecordstatetodifferentiationIntResponse(msg, 0, w, success)
    }
}


func UpdateMaprecordstatetodifferentiation(w http.ResponseWriter, req *http.Request) {
    var data = entities.MaprecordstatetodifferentiationEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.UpdateMaprecordstatetodifferentiation(&data)
        ThrowMaprecordstatetodifferentiationIntResponse(msg, 0, w, success)
    }
}


