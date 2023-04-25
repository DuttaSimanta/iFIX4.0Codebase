package handlers


import (
  "iFIX/ifix/entities"
  "iFIX/ifix/models"
  "iFIX/ifix/logger"
  "log"
  "net/http"
  "encoding/json"
  )
//   DashboardQuery
// func ThrowDashboardQueryAllResponse(successMessage string, responseData entities.DashboardQueryEntities, w http.ResponseWriter, success bool) {
//     var response = entities.DashboardQueryResponse{}
//     response.Success = success
//     response.Message = successMessage
//     response.Details = responseData
//     jsonResponse, jsonError := json.Marshal(response)
//     if jsonError != nil {
//         logger.Log.Fatal("Internel Server Error")
//     }
//     w.Header().Set("Content-Type", "application/json")
//     w.Write(jsonResponse)
// }


func ThrowDashboardQueryIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
    var response = entities.DashboardQueryResponseInt{}
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


func AddDashboardQuery(w http.ResponseWriter, req *http.Request) {
    var data = entities.DashboardQueryEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.AddDashboardQuery(&data)
        ThrowDashboardQueryIntResponse(msg, data, w, success)
    }
}


// func GetAllDashboardQuery(w http.ResponseWriter, req *http.Request) {
//     var data = entities.DashboardQueryEntity{}
//     jsonError := data.FromJSON(req.Body)
//     if jsonError != nil {
//         log.Print(jsonError)
//         logger.Log.Println(jsonError)
//         entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
//     } else {
//         data, success, _, msg := models.GetAllDashboardQuery(&data)
//         ThrowDashboardQueryAllResponse(msg, data, w, success)
//     }
// }


// func DeleteDashboardQuery(w http.ResponseWriter, req *http.Request) {
//     var data = entities.DashboardQueryEntity{}
//     jsonError := data.FromJSON(req.Body)
//     if jsonError != nil {
//         log.Print(jsonError)
//         logger.Log.Println(jsonError)
//         entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
//     } else {
//         success, _, msg := models.DeleteDashboardQuery(&data)
//         ThrowDashboardQueryIntResponse(msg, 0, w, success)
//     }
// }


func UpdateDashboardQuery(w http.ResponseWriter, req *http.Request) {
    var data = entities.DashboardQueryEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.UpdateDashboardQuery(&data)
        ThrowDashboardQueryIntResponse(msg, 0, w, success)
    }
}
