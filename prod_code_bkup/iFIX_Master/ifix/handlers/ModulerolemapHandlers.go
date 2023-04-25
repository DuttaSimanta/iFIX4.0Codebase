package handlers


import (
  "iFIX/ifix/entities"
  "iFIX/ifix/models"
  "iFIX/ifix/logger"
  "log"
  "net/http"
  "encoding/json"
  )

func ThrowModulerolemapAllResponse(successMessage string, responseData entities.ModulerolemapEntities, w http.ResponseWriter, success bool) {
    var response = entities.ModulerolemapResponse{}
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


func ThrowModulerolemapIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
    var response = entities.ModulerolemapResponseInt{}
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


func InsertModulerolemap(w http.ResponseWriter, req *http.Request) {
    var data = entities.ModulerolemapEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.InsertModulerolemap(&data)
        ThrowModulerolemapIntResponse(msg, data, w, success)
    }
}


func GetAllModulerolemap(w http.ResponseWriter, req *http.Request) {
    var data = entities.ModulerolemapEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.GetAllModulerolemap(&data)
        ThrowModulerolemapAllResponse(msg, data, w, success)
    }
}


func DeleteModulerolemap(w http.ResponseWriter, req *http.Request) {
    var data = entities.ModulerolemapEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.DeleteModulerolemap(&data)
        ThrowModulerolemapIntResponse(msg, 0, w, success)
    }
}


func UpdateModulerolemap(w http.ResponseWriter, req *http.Request) {
    var data = entities.ModulerolemapEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.UpdateModulerolemap(&data)
        ThrowModulerolemapIntResponse(msg, 0, w, success)
    }
}


