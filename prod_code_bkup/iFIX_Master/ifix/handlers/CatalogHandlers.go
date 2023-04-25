package handlers


import (
  "iFIX/ifix/entities"
  "iFIX/ifix/models"
  "iFIX/ifix/logger"
  "log"
  "net/http"
  "encoding/json"
  )

func ThrowCatalogAllResponse(successMessage string, responseData entities.CatalogEntities, w http.ResponseWriter, success bool) {
    var response = entities.CatalogResponse{}
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
func ThrowCatalogResponse(successMessage string, responseData []entities.CatalogEntitySingle, w http.ResponseWriter, success bool) {
    var response = entities.CatalogSingleResponse{}
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

func ThrowCatalogIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
    var response = entities.CatalogResponseInt{}
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
func ThrowSplitAllResponse(successMessage string, responseData entities.RecordEntity, w http.ResponseWriter, success bool) {
    var response = entities.CatalogRecordResponse{}
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
func Getcatelogrecord(w http.ResponseWriter, req *http.Request) {
    log.Print("inside catelog")
    var data = entities.CatalogEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        //entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        //var data1 = entities.SplitOutputEntity{}
        data, success,_, msg := models.Getcatelogrecordmodel(&data)
        ThrowSplitAllResponse(msg, data, w, success)
    }
}
func InsertCatalog(w http.ResponseWriter, req *http.Request) {
    var data = entities.CatalogEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.InsertCatalog(&data)
        ThrowCatalogIntResponse(msg, data, w, success)
    }
}
func GetAllCatalog(w http.ResponseWriter, req *http.Request) {
    var data = entities.CatalogEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.GetAllCatalog(&data)
        ThrowCatalogAllResponse(msg, data, w, success)
    }
}
func Getcatalogorgwise(w http.ResponseWriter, req *http.Request) {
    var data = entities.CatalogEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.Getcatalogorgwise(&data)
        ThrowCatalogResponse(msg, data, w, success)
    }
}


func DeleteCatalog(w http.ResponseWriter, req *http.Request) {
    var data = entities.CatalogEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.DeleteCatalog(&data)
        ThrowCatalogIntResponse(msg, 0, w, success)
    }
}


func UpdateCatalog(w http.ResponseWriter, req *http.Request) {
    var data = entities.CatalogEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.UpdateCatalog(&data)
        ThrowCatalogIntResponse(msg, 0, w, success)
    }
}


