  package handlers


import (
  "iFIX/ifix/entities"
  "iFIX/ifix/models"
  "iFIX/ifix/logger"
 // "log"
  "net/http"
  "encoding/json"
  )

  
func ThrowMstClientCredentialTypeAllResponse(successMessage string, responseData []entities.MstClientCredentialTypeEntity, w http.ResponseWriter, success bool) {
    var response = entities.MstClientCredentialTypeResponse{}
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

func GetAllMstClientCredentialType(w http.ResponseWriter, req *http.Request) {
    /*var data = entities.MstExcelTemplateTypeEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {*/
        data, success, _, msg := models.GetAllMstClientCredentialType()
        ThrowMstClientCredentialTypeAllResponse(msg, data, w, success)
   // }
}