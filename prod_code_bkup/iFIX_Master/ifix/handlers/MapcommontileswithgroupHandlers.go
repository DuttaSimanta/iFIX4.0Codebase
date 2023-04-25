package handlers


import (
  "iFIX/ifix/entities"
  "iFIX/ifix/models"
  "iFIX/ifix/logger"
  "log"
  "net/http"
  "encoding/json"
  )

func ThrowMapcommontileswithgroupAllResponse(successMessage string, responseData entities.MapcommontileswithgroupEntities, w http.ResponseWriter, success bool) {
    var response = entities.MapcommontileswithgroupResponse{}
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


func ThrowMapcommontileswithgroupIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
    var response = entities.MapcommontileswithgroupResponseInt{}
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


func InsertMapcommontileswithgroup(w http.ResponseWriter, req *http.Request) {
    var data = entities.MapcommontileswithgroupEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.InsertMapcommontileswithgroup(&data)
        ThrowMapcommontileswithgroupIntResponse(msg, data, w, success)
    }
}


func GetAllMapcommontileswithgroup(w http.ResponseWriter, req *http.Request) {
    var data = entities.MapcommontileswithgroupEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.GetAllMapcommontileswithgroup(&data)
        ThrowMapcommontileswithgroupAllResponse(msg, data, w, success)
    }
}


func DeleteMapcommontileswithgroup(w http.ResponseWriter, req *http.Request) {
    var data = entities.MapcommontileswithgroupEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.DeleteMapcommontileswithgroup(&data)
        ThrowMapcommontileswithgroupIntResponse(msg, 0, w, success)
    }
}


func UpdateMapcommontileswithgroup(w http.ResponseWriter, req *http.Request) {
    var data = entities.MapcommontileswithgroupEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.UpdateMapcommontileswithgroup(&data)
        ThrowMapcommontileswithgroupIntResponse(msg, 0, w, success)
    }
}


