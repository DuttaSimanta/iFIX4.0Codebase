package handlers


import (
  "iFIX/ifix/entities"
  "iFIX/ifix/models"
  "iFIX/ifix/logger"
  "log"
  "net/http"
  "encoding/json"
  )

func ThrowmapldapgrouproleAllResponse(successMessage string, responseData entities.MapldapgrouproleEntities, w http.ResponseWriter, success bool) {
    var response = entities.MapldapgrouproleResponse{}
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


func ThrowmapldapgrouproleIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
    var response = entities.MapldapgrouproleResponseInt{}
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


func Insertmapldapgrouprole(w http.ResponseWriter, req *http.Request) {
    var data = entities.MapldapgrouproleEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.Insertmapldapgrouprole(&data)
        ThrowmapldapgrouproleIntResponse(msg, data, w, success)
    }
}


func GetAllmapldapgrouprole(w http.ResponseWriter, req *http.Request) {
    var data = entities.MapldapgrouproleEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.GetAllmapldapgrouprole(&data)
        ThrowmapldapgrouproleAllResponse(msg, data, w, success)
    }
}


func Deletemapldapgrouprole(w http.ResponseWriter, req *http.Request) {
    var data = entities.MapldapgrouproleEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.Deletemapldapgrouprole(&data)
        ThrowmapldapgrouproleIntResponse(msg, 0, w, success)
    }
}


func Updatemapldapgrouprole(w http.ResponseWriter, req *http.Request) {
    var data = entities.MapldapgrouproleEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.Updatemapldapgrouprole(&data)
        ThrowmapldapgrouproleIntResponse(msg, 0, w, success)
    }
}


