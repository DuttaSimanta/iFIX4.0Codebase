package fileutils

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	model "src/entities"
	Logger "src/logger"
	"strconv"
	"strings"
	"time"
)

func FileUploadAPICall(ClientID int64, orgID int64, url string, path string) (string, string, error) {
	Logger.Log.Println("FileUploadAPICall")
	client := &http.Client{
		Timeout: time.Second * 30,
	}
	var originalFilename string
	var newFileName string
	file, err := os.Open(path)
	if err != nil {
		Logger.Log.Println("Unable to Open File======>", err)
		return originalFilename, newFileName, err
	}
	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("myFile", filepath.Base(path))
	if err != nil {
		Logger.Log.Println("Unable to write File in stream======>", err)
		return originalFilename, newFileName, err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		Logger.Log.Println("Unable to copy File in write buffer======>", err)
		return originalFilename, newFileName, err
	}
	fw, err := writer.CreateFormField("clientid")
	if err != nil {
		Logger.Log.Println(err)
		return originalFilename, newFileName, err
	}
	_, err = io.Copy(fw, strings.NewReader(strconv.FormatInt(ClientID, 10)))
	if err != nil {
		Logger.Log.Println(err)
		return originalFilename, newFileName, err
	}
	fw, err = writer.CreateFormField("mstorgnhirarchyid")
	if err != nil {
		Logger.Log.Println(err)
		return originalFilename, newFileName, err
	}
	_, err = io.Copy(fw, strings.NewReader(strconv.FormatInt(orgID, 10)))
	if err != nil {
		Logger.Log.Println(err)
		return originalFilename, newFileName, err
	}
	err = writer.Close()
	if err != nil {
		Logger.Log.Println(err)
		return originalFilename, newFileName, err
	}
	req, _ := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, _ := client.Do(req)
	var fileuploadResponse model.FileuploadResponse
	respBody, err := ioutil.ReadAll(resp.Body)

	Logger.Log.Println("respbody===>", string(respBody))
	if err != nil {
		Logger.Log.Println(err)
		return originalFilename, newFileName, err
	}
	err = json.Unmarshal(respBody, &fileuploadResponse)
	Logger.Log.Println("fileuploadResponse========>", fileuploadResponse)
	if err != nil {
		Logger.Log.Println(err)
		return originalFilename, newFileName, err
	}
	//Logger.Log.Println("Response====>", resp.Body)
	originalFilename = fileuploadResponse.Details.Originalfile
	newFileName = fileuploadResponse.Details.Filename
	return originalFilename, newFileName, nil
}
