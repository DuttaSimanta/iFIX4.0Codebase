package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func ThrowFileUploadResponse(successMessage string, responseData entities.FileuploadEntity, w http.ResponseWriter, success bool) {
	var response = entities.FileuploadResponse{}
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

func UploadFile(w http.ResponseWriter, r *http.Request) {
	logger.Log.Printf("Uploaded File: start")
	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		logger.Log.Println("Error Retrieving the File")
		logger.Log.Println(err)
		return
	}
	defer file.Close()
	logger.Log.Printf("Uploaded File: %+v\n", handler.Filename)
	logger.Log.Printf("File Size: %+v\n", handler.Size)
	logger.Log.Printf("MIME Header: %+v\n", handler.Header)

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)

	var data = entities.FileuploadEntity{}
	a, err := strconv.ParseInt(r.Form["clientid"][0], 10, 64)

	b, err := strconv.ParseInt(r.Form["mstorgnhirarchyid"][0], 10, 64)
	data.Clientid = a
	data.Mstorgnhirarchyid = b
	data, success, _, msg := models.UploadFile(&data, fileBytes, handler.Filename, handler.Header["Content-Type"][0])
	ThrowFileUploadResponse(msg, data, w, success)

}

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	logger.Log.Printf("Download File: start")
	var data = entities.FileuploadEntity{}
	jsonError := data.FromJSON(r.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	}
	data1, success, _, msg := models.DownloadFile(&data)
	logger.Log.Printf("Download File: End")
	if !success {
		ThrowFileUploadResponse(msg, data, w, success)
	} else {
		_, err := data1.WriteTo(w)
		if err != nil {
			ThrowFileUploadResponse(msg, data, w, success)
		}
	}

}
