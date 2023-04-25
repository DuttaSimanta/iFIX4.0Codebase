package handlers


import (
  "iFIX/ifix/entities"
  "iFIX/ifix/models"
  "iFIX/ifix/logger"
  "log"
  "net/http"
  "encoding/json"
  )

func ThrowClientsupportgroupAllResponse(successMessage string, responseData entities.ClientsupportgroupEntities, w http.ResponseWriter, success bool) {
    var response = entities.ClientsupportgroupResponse{}
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

// ThrowGroupSingleResponse function is used to throw success response of All data in JSON format
func ThrowGroupSingleResponse(successMessage string, responseData []entities.ClientsupportgroupsingleEntity, w http.ResponseWriter, success bool) {
    var response = entities.ClientsupportgroupsingleResponse{}
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
func ThrowClientsupportgroupIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
    var response = entities.ClientsupportgroupResponseInt{}
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


func Getgroupbyorgid(w http.ResponseWriter, req *http.Request) {
    var data = entities.ClientsupportgroupEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.Getgroupbyorgid(&data)
        ThrowGroupSingleResponse(msg, data, w, success)
    }
}
func Getprocessgroupbyorgid(w http.ResponseWriter, req *http.Request) {
    var data = entities.ClientsupportgroupEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.Getprocessgroupbyorgid(&data)
        ThrowGroupSingleResponse(msg, data, w, success)
    }
}
func Getprocessgroupbyorgids(w http.ResponseWriter, req *http.Request) {
    var data = entities.ClientsupportgroupEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.Getprocessgroupbyorgids(&data)
        ThrowGroupSingleResponse(msg, data, w, success)
    }
}
func InsertClientsupportgroup(w http.ResponseWriter, req *http.Request) {
    var data = entities.ClientsupportgroupEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.InsertClientsupportgroup(&data)
        ThrowClientsupportgroupIntResponse(msg, data, w, success)
    }
}

func GetAllClientsupportgroup(w http.ResponseWriter, req *http.Request) {
    var data = entities.ClientsupportgroupEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.GetAllClientsupportgroup(&data)
        ThrowClientsupportgroupAllResponse(msg, data, w, success)
    }
}

func DeleteClientsupportgroup(w http.ResponseWriter, req *http.Request) {
    var data = entities.ClientsupportgroupEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.DeleteClientsupportgroup(&data)
        ThrowClientsupportgroupIntResponse(msg, 0, w, success)
    }
}


func UpdateClientsupportgroup(w http.ResponseWriter, req *http.Request) {
    var data = entities.ClientsupportgroupEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.UpdateClientsupportgroup(&data)
        ThrowClientsupportgroupIntResponse(msg, 0, w, success)
    }
}


