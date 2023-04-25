package handlers


import (
  "iFIX/ifix/entities"
  "iFIX/ifix/models"
  "iFIX/ifix/logger"
  "log"
  "net/http"
  "encoding/json"
  )

func ThrowMapexternalattributesAllResponse(successMessage string, responseData entities.MapexternalattributesEntities, w http.ResponseWriter, success bool) {
    var response = entities.MapexternalattributesResponse{}
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
func ThrowMappedattributesResponse(successMessage string, responseData []entities.Attr, w http.ResponseWriter, success bool) {
    var response = entities.MappedattributesResponse{}
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

func ThrowMapexternalattributesIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
    var response = entities.MapexternalattributesResponseInt{}
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


func InsertMapexternalattributes(w http.ResponseWriter, req *http.Request) {
    var data = entities.MapexternalattributesEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.InsertMapexternalattributes(&data)
        ThrowMapexternalattributesIntResponse(msg, data, w, success)
    }
}


func GetAllMapexternalattributes(w http.ResponseWriter, req *http.Request) {
    var data = entities.MapexternalattributesEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.GetAllMapexternalattributes(&data)
        ThrowMapexternalattributesAllResponse(msg, data, w, success)
    }
}
func GetMappedattributes(w http.ResponseWriter, req *http.Request) {
    var data = entities.MapexternalattributesEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.GetMappedattributes(&data)
        ThrowMappedattributesResponse(msg, data, w, success)
    }
}


func DeleteMapexternalattributes(w http.ResponseWriter, req *http.Request) {
    var data = entities.MapexternalattributesEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.DeleteMapexternalattributes(&data)
        ThrowMapexternalattributesIntResponse(msg, 0, w, success)
    }
}

