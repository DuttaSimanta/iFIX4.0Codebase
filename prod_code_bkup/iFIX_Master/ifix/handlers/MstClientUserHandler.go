package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"net/http"
)

// ThrowSearchUsersResponse function is used to throw success response of All data in JSON format
func ThrowSearchUsersResponse(successMessage string, responseData []entities.MstUserSearchEntity, w http.ResponseWriter, success bool) {
	var response = entities.MstSearchUserEntityResponse{}
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

// ThrowClientUsersAllResponse function is used to throw success response of All data in JSON format
func ThrowClientUsersAllResponse(successMessage string, responseData entities.MstClientUserEntities, w http.ResponseWriter, success bool) {
	var response = entities.MstClientUserEntityResponse{}
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

// ThrowClientUsersIntResponse function is used to throw success response of integer data in JSON format
func ThrowClientUsersIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.MstClientUserRoleResponseInt{}
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

// ThrowRecordwiseuserinfoResponse function is used to throw success response of All data in JSON format
func ThrowRecordwiseuserinfoResponse(successMessage string, responseData []entities.MstClientUserEntity, w http.ResponseWriter, success bool) {
	var response = entities.MstRecordwiseuserinfoEntityResponse{}
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

func ThrowSearchLoginNameResponse(successMessage string, responseData []entities.LoginNameSearchEntity, w http.ResponseWriter, success bool) {
	var response = entities.LoginNameSearchEntityResponse{}
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
func ThrowSearchNameResponse(successMessage string, responseData []entities.NameSearchEntity, w http.ResponseWriter, success bool) {
	var response = entities.NameSearchEntityResponse{}
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
func ThrowSearchBranchResponse(successMessage string, responseData []entities.BranchSearchEntity, w http.ResponseWriter, success bool) {
	var response = entities.BranchSearchEntityResponse{}
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
func ThrowSearchLOginnameAndNameResponse(successMessage string, responseData []entities.LoginnameAndNameEntity, w http.ResponseWriter, success bool) {
	var response = entities.LoginnameAndNameSearchEntityResponse{}
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

//AddClientUser method is used for insert user data
func AddClientUser(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstClientUserEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data, success, _, msg := models.AddClientUsers(&data)
		ThrowClientUsersIntResponse(msg, data, w, success)
	}
}

//DeleteClientUser method is used for delete user data
func DeleteClientUser(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstClientUserEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		success, _, msg := models.DeleteUsers(&data)
		ThrowClientUsersIntResponse(msg, 0, w, success)
	}
}

//UpdateClientUser  method is used for update user data
func UpdateClientUser(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstClientUserEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		success, _, msg := models.UpdateUsers(&data)
		ThrowClientUsersIntResponse(msg, 0, w, success)
	}
}
func Updateusercolor(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstClientUserEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		success, _, msg := models.Updateusercolor(&data)
		ThrowClientUsersIntResponse(msg, 0, w, success)
	}
}

//GetAllClientUsers  method is used for get client user data
func GetAllClientUsers(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstClientUserEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data1, success, _, msg := models.GetAllUsers(&data)
		ThrowClientUsersAllResponse(msg, data1, w, success)
	}
}

//SearchUser  method is used for search  user data
func SearchUser(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstClientUserEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data1, success, _, msg := models.SearchUser(&data)
		ThrowSearchUsersResponse(msg, data1, w, success)
	}
}

//SearchUserByOrgnId  method is used for search  user data
func SearchUserByOrgnId(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstClientUserEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data1, success, _, msg := models.SearchUserByOrgnId(&data)
		ThrowSearchUsersResponse(msg, data1, w, success)
	}
}
func Searchuserbyclientid(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstClientUserEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data1, success, _, msg := models.Searchuserbyclientid(&data)
		ThrowSearchUsersResponse(msg, data1, w, success)
	}
}

//GetRecordwiseuserinfo  method is used for get client user data
func GetRecordwiseuserinfo(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstGetUserByRecordidEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data1, success, _, msg := models.Recordwiseuserinfo(&data)
		ThrowRecordwiseuserinfoResponse(msg, data1, w, success)
	}
}

//GetIDwiseuserinfo  method is used for get client user data
func GetIDwiseuserinfo(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstGetUserByRecordidEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data1, success, _, msg := models.IDwiseuserinfo(&data)
		ThrowRecordwiseuserinfoResponse(msg, data1, w, success)
	}
}

func SearchLoginName(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstClientUserEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data1, success, _, msg := models.SearchLoginName(&data)
		ThrowSearchLoginNameResponse(msg, data1, w, success)
	}
}

//GetRe
func SearchName(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstClientUserEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data1, success, _, msg := models.SearchName(&data)
		ThrowSearchNameResponse(msg, data1, w, success)
	}
}

//GetRe
func SearchBranch(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstClientUserEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data1, success, _, msg := models.SearchBranch(&data)
		ThrowSearchBranchResponse(msg, data1, w, success)
	}
}

//GetRe
func SearchLoginamebyGroupids(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstClientUserEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data1, success, _, msg := models.SearchLoginamebyGroupids(&data)
		ThrowSearchLOginnameAndNameResponse(msg, data1, w, success)
	}
}
func RemoveWorkflowUsers(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstClientUserEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		success,  _ := models.RemoveWorkflowUsersWithapi(&data)
		ThrowClientUsersIntResponse("", 0, w, success)
	}
}
