package handlers

import (
	"encoding/json"
	"net/http"
	"src/entities"
	"src/logger"
	"src/models"
)

func ThrowDownloadTablesdataResponse(successMessage string, originalfilename string, uploadedfilename string, w http.ResponseWriter, success bool) {
	var response = entities.Downloadresponse{}
	response.Status = success
	response.Message = successMessage
	response.OriginalFileName = originalfilename
	response.UploadedFileName = uploadedfilename
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		logger.Log.Fatal("Internel Server Error")
	}
	// err := ioutil.WriteFile("orgname_masterdata_systemdatetime.json", jsonResponse, 0644)
	// if err != nil {
	// 	logger.Log.Fatal("Json file writing error")
	// }
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
func ThrowUploadTablesdataResponse(successMessage string, w http.ResponseWriter, success bool) {
	var response = entities.InsertresponceEntity{}
	response.Status = success
	response.Message = successMessage
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		logger.Log.Fatal("Internel Server Error")
	}
	// err := ioutil.WriteFile("userfile.json", jsonResponse, 0644)
	// if err != nil {
	// 	logger.Log.Fatal("Json file writing error")
	// }
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
func DownloadTablesdata(w http.ResponseWriter, req *http.Request) {
	var data entities.Transporttable
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		/*UploadedFileName, success, _, msg := */
		originalfilename, uploadedfilename, success, _, msg := models.DownloadTablesdata(&data)
		ThrowDownloadTablesdataResponse(msg, originalfilename, uploadedfilename, w, success)
	}
}
func UploadTablesdata(w http.ResponseWriter, req *http.Request) {
	var data entities.Uploadentity
	jsonError := data.FromJSON(req.Body)
	// content, err := ioutil.ReadFile("userfile.json")
	// if err != nil {
	// 	logger.Log.Println(err)
	// }
	// data := entities.ResultofalltableEntity{}
	//err = json.Unmarshal(content, &data)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		/*UploadedFileName, success, _, msg := */
		success, _, msg := models.UploadTablesdata(&data)
		ThrowUploadTablesdataResponse(msg, w, success)
	}
}
