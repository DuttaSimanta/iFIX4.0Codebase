package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"io/ioutil"
	"log"
	"net/http"
)

// ThrowLoginAllResponse function is used to throw success response of All data in JSON format
func ThrowLoginAllResponse(successMessage string, responseData []entities.LoginEntityResp, w http.ResponseWriter, success bool) {
	var response = entities.LoginResponse{}
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

// ThrowUserResponse function is used to throw success response of All data in JSON format
func ThrowUserResponse(successMessage string, responseData []entities.UserEntity, w http.ResponseWriter, success bool) {
	var response = entities.UserResponse{}
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
func ThrowLdapAttrResponse(successMessage string, responseData entities.LdapAttrEntityResponse, w http.ResponseWriter, success bool) {
	var response = entities.LdapAttrEntityResponse{}
	response.Success = success
	response.Message = successMessage
	response.Details = responseData.Details
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		logger.Log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
func VerifyTOTP(w http.ResponseWriter, req *http.Request) {
	var data = entities.LoginEntityReq{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.VerifyTOTP(&data)
		ThrowLoginAllResponse(msg, data, w, success)
	}
}

func Login(w http.ResponseWriter, req *http.Request) {
	var data = entities.LoginEntityReq{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.Loginchecking(&data)
		logger.Log.Println(success, msg)
		ThrowLoginAllResponse(msg, data, w, success)
	}
}
func SendOTP(w http.ResponseWriter, req *http.Request) {
	var data = entities.LoginEntityReq{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.SendOTP(&data)
		logger.Log.Println(success, msg)
		ThrowLoginAllResponse(msg, data, w, success)
	}
}
func Adfslogin(w http.ResponseWriter, req *http.Request) {
	var result map[string]interface{}
	body, jsonError := ioutil.ReadAll(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		json.Unmarshal(body, &result)
			data, success, _, msg := models.Adfslogin(result)
			logger.Log.Println(success,msg)
			ThrowLoginAllResponse(msg, data, w, success)
	}
	//var data = entities.LoginEntityReq{}
	//jsonError := data.FromJSON(req.Body)
	//
	//if jsonError != nil {
	//	log.Print(jsonError)
	//	logger.Log.Println(jsonError)
	//	entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	//} else {
	//	//responseError = validators.ValidateAddMstClient(data)
	//	//if(len(responseError)==0){
	//	data, success, _, msg := models.Loginchecking(&data)
	//	logger.Log.Println(success,msg)
	//	ThrowLoginAllResponse(msg, data, w, success)
	//}
}
func Getorgname(w http.ResponseWriter, req *http.Request) {
	var data = entities.LoginEntityReq{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.Getorgname(&data)
		ThrowLoginAllResponse(msg, data, w, success)
	}
}
func Inpersonatelogin(w http.ResponseWriter, req *http.Request) {
	var data = entities.LoginEntityReq{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.Inpersonatelogin(&data)
		ThrowLoginAllResponse(msg, data, w, success)
	}
}
func GetUserDetailsById(w http.ResponseWriter, req *http.Request) {
	var data = entities.UserEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//var token = req.Header.Get("Authorization")
		//success := utility.CheckToken(token, data.Userid)
		//if !success {
		//	w.WriteHeader(http.StatusUnauthorized)
		//	return
		//}
		data, success, _, msg := models.GetUserDetailsById(&data)
		ThrowUserResponse(msg, data, w, success)
	}
}
func Logout(w http.ResponseWriter, req *http.Request) {
	var data = entities.UserEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.Logout(&data)
		ThrowUserResponse(msg, data, w, success)
	}
}
func Generatetoken(w http.ResponseWriter, req *http.Request) {
	var data = entities.LoginEntityReq{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.Generatetoken(&data)
		ThrowLoginAllResponse(msg, data, w, success)
	}
}
func Changepassword(w http.ResponseWriter, req *http.Request) {
	var data = entities.LoginEntityReq{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.Changepassword(&data)
		ThrowClientUsersIntResponse(msg, data, w, success)
	}
}
func ForgetPassword(w http.ResponseWriter, req *http.Request) {
	var data = entities.LoginEntityReq{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.ForgetPassword(&data)
		ThrowClientUsersIntResponse(msg, data, w, success)
	}
}
func Getldapattributes(w http.ResponseWriter, req *http.Request) {
	var data = entities.LoginEntityReq{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.Getldapattributes(&data)
		ThrowLdapAttrResponse(msg, data, w, success)
	}
}
func Validateusertoken(w http.ResponseWriter, req *http.Request) {
	var data = entities.LoginEntityResp{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.Validateusertoken(&data)
		ThrowClientUsersIntResponse(msg, data, w, success)
	}
}
