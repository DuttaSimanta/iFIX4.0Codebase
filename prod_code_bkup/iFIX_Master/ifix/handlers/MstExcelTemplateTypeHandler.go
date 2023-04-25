package handlers


import (
  "iFIX/ifix/entities"
  "iFIX/ifix/models"
  "iFIX/ifix/logger"
//  "log"
  "net/http"
  "encoding/json"
  )
func ThrowMstExcelTemplateTypeAllResponse(successMessage string, responseData []entities.MstExcelTemplateTypeEntity, w http.ResponseWriter, success bool) {
    var response = entities.MstExcelTemplateTypeResponse{}
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

func GetAllMstExcelTemplateType(w http.ResponseWriter, req *http.Request) {
    /*var data = entities.MstExcelTemplateTypeEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {*/
        data, success, _, msg := models.GetAllMstExcelTemplateType()
        ThrowMstExcelTemplateTypeAllResponse(msg, data, w, success)
   // }
}