package handlers


import (
  "iFIX/ifix/entities"
  "iFIX/ifix/models"
  "iFIX/ifix/logger"
  "log"
  "net/http"
  "encoding/json"
  )

func ThrowGroupmemberAllResponse(successMessage string, responseData entities.GroupmemberEntities, w http.ResponseWriter, success bool) {
    var response = entities.GroupmemberResponse{}
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


func ThrowGroupmemberIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
    var response = entities.GroupmemberResponseInt{}
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


func InsertGroupmember(w http.ResponseWriter, req *http.Request) {
    var data = entities.GroupmemberEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.InsertGroupmember(&data)
        ThrowGroupmemberIntResponse(msg, data, w, success)
    }
}


func GetAllGroupmember(w http.ResponseWriter, req *http.Request) {
    var data = entities.GroupmemberEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.GetAllGroupmember(&data)
        ThrowGroupmemberAllResponse(msg, data, w, success)
    }
}


func DeleteGroupmember(w http.ResponseWriter, req *http.Request) {
    var data = entities.GroupmemberEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.DeleteGroupmember(&data)
        ThrowGroupmemberIntResponse(msg, 0, w, success)
    }
}


func UpdateGroupmember(w http.ResponseWriter, req *http.Request) {
    var data = entities.GroupmemberEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.UpdateGroupmember(&data)
        ThrowGroupmemberIntResponse(msg, 0, w, success)
    }
}

//SearchUserByGroupId  method is used for search  user data
func SearchUserByGroupId(w http.ResponseWriter, req *http.Request) {
    var data = entities.GroupmemberEntity{}
    jsonError := data.FromJSON(req.Body)

    if jsonError != nil {
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        //responseError = validators.ValidateAddMstClient(data)
        //if(len(responseError)==0){
        data1, success, _, msg := models.SearchUserByGroupId(&data)
        ThrowSearchUsersResponse(msg, data1, w, success)
    }
}
func Searchuserdetailsbygroupid(w http.ResponseWriter, req *http.Request) {
    var data = entities.GroupmemberEntity{}
    jsonError := data.FromJSON(req.Body)

    if jsonError != nil {
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        //responseError = validators.ValidateAddMstClient(data)
        //if(len(responseError)==0){
        data1, success, _, msg := models.Searchuserdetailsbygroupid(&data)
        ThrowSearchUsersResponse(msg, data1, w, success)
    }
}
func Groupbyuserwise(w http.ResponseWriter, req *http.Request) {
    var data = entities.GroupmemberEntity{}
    jsonError := data.FromJSON(req.Body)

    if jsonError != nil {
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        //responseError = validators.ValidateAddMstClient(data)
        //if(len(responseError)==0){
        data1, success, _, msg := models.Groupbyuserwise(&data)
        ThrowGroupSingleResponse(msg, data1, w, success)
    }
}
func Workflowgroupbyuserwise(w http.ResponseWriter, req *http.Request) {
    var data = entities.GroupmemberEntity{}
    jsonError := data.FromJSON(req.Body)

    if jsonError != nil {
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        //responseError = validators.ValidateAddMstClient(data)
        //if(len(responseError)==0){
        data1, success, _, msg := models.Workflowgroupbyuserwise(&data)
        ThrowGroupSingleResponse(msg, data1, w, success)
    }
}
func AddGroupmember(w http.ResponseWriter, req *http.Request) {
    var data = entities.GroupmemberEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.AddGroupmember(&data)
        ThrowGroupmemberIntResponse(msg, data, w, success)
    }
}
func GetAllGrpmember(w http.ResponseWriter, req *http.Request) {
    var data = entities.GroupmemberEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.GetAllGrpmember(&data)
        ThrowGroupmemberAllResponse(msg, data, w, success)
    }
}
//GetUserByGroupId  method is used for search  user data
func GetUserByGroupId(w http.ResponseWriter, req *http.Request) {
    var data = entities.GroupmemberEntity{}
    jsonError := data.FromJSON(req.Body)

    if jsonError != nil {
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        //responseError = validators.ValidateAddMstClient(data)
        //if(len(responseError)==0){
        data1, success, _, msg := models.GetUserByGroupId(&data)
        ThrowSearchUsersResponse(msg, data1, w, success)
    }
}
func SearchAnalystOrgWise(w http.ResponseWriter, req *http.Request) {
    var data = entities.GroupmemberEntity{}
    jsonError := data.FromJSON(req.Body)

    if jsonError != nil {
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        //responseError = validators.ValidateAddMstClient(data)
        //if(len(responseError)==0){
        data1, success, _, msg := models.SearchAnalystOrgWise(&data)
        ThrowSearchUsersResponse(msg, data1, w, success)
    }
}
