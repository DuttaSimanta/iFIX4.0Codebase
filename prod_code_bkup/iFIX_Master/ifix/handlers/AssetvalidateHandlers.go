package handlers


import (
  "iFIX/ifix/entities"
  "iFIX/ifix/models"
  "iFIX/ifix/logger"
  "log"
  "net/http"
  "encoding/json"
  )

func ThrowAssetvalidateAllResponse(successMessage string, responseData entities.AssetvalidateEntities, w http.ResponseWriter, success bool) {
    var response = entities.AssetvalidateResponse{}
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


func ThrowAssetvalidateIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
    var response = entities.AssetvalidateResponseInt{}
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


func InsertAssetvalidate(w http.ResponseWriter, req *http.Request) {
    var data = entities.AssetvalidateEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.InsertAssetvalidate(&data)
        ThrowAssetvalidateIntResponse(msg, data, w, success)
    }
}


func GetAllAssetvalidate(w http.ResponseWriter, req *http.Request) {
    var data = entities.AssetvalidateEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.GetAllAssetvalidate(&data)
        ThrowAssetvalidateAllResponse(msg, data, w, success)
    }
}


func DeleteAssetvalidate(w http.ResponseWriter, req *http.Request) {
    var data = entities.AssetvalidateEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.DeleteAssetvalidate(&data)
        ThrowAssetvalidateIntResponse(msg, 0, w, success)
    }
}


func UpdateAssetvalidate(w http.ResponseWriter, req *http.Request) {
    var data = entities.AssetvalidateEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.UpdateAssetvalidate(&data)
        ThrowAssetvalidateIntResponse(msg, 0, w, success)
    }
}


