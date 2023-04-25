package entities

import (
	"encoding/json"
	"iFIX/ifix/logger"
	"io"
	"log"
	"net/http"
)

type UserWithGroupAndCategoryEntity struct {
	Id                  int64   `json:"id"`
	Clientid            int64   `json:"clientid"`
	Mstorgnhirarchyid   int64   `json:"mstorgnhirarchyid"`
	Recorddifftypeid    int64   `json:"recorddifftypeid"`
	Recorddiffid        int64   `json:"recorddiffid"`
	WorkingCategories   []int64 `json:"workingcategories"`
	Categoryid          int64   `json:"categoryid"`
	Groupid             int64   `json:"groupid"`
	Userid              int64   `json:"refuserid"`
	Activeflg           int64   `json:"activeflg"`
	Audittransactionid  int64   `json:"audittransactionid"`
	Offset              int64   `json:"offset"`
	Limit               int64   `json:"limit"`
	Clientname          string  `json:"clientname"`
	Mstorgnhirarchyname string  `json:"mstorgnhirarchyname"`
	Recorddifftypename  string  `json:"recorddifftypename"`
	Recorddiffname      string  `json:"recorddiffname"`
	Categoryname        string  `json:"categoryname"`
	Groupname           string  `json:"groupname"`
	Username            string  `json:"username"`
	Userloginname       string  `json:"userloginname"`
}

type UserWithGroupAndCategoryForBulkUploadEntity struct {
	Clientid          int64 `json:"clientid"`
	Mstorgnhirarchyid int64 `json:"mstorgnhirarchyid"`
	Recorddifftypeid  int64 `json:"recorddifftypeid"`
	Recorddiffid      int64 `json:"recorddiffid"`
	// WorkingCategories []int64 `json:"workingcategories"`
	Categoryid       int64  `json:"categoryid"`
	Groupid          int64  `json:"groupid"`
	Userid           int64  `json:"userid"`
	OriginalFileName string `json:"originalfilename"`
	UploadedFileName string `json:"uploadedfilename"`
}

type UserWithGroupAndCategoryForBulkUploadEntities struct {
	Total  int64                                         `json:"total"`
	Values []UserWithGroupAndCategoryForBulkUploadEntity `json:"values"`
}

type UserWithGroupAndCategoryEntities struct {
	Total  int64                            `json:"total"`
	Values []UserWithGroupAndCategoryEntity `json:"values"`
}

type UserWithGroupAndCategoryResponse struct {
	Success bool                             `json:"success"`
	Message string                           `json:"message"`
	Details UserWithGroupAndCategoryEntities `json:"details"`
}

type UserWithGroupAndCategoryResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

type BulkUserWithGroupAndCategoryUploadResponse struct {
	Status   bool   `json:"success"`
	Message  string `json:"message"`
	Response string `json:"response"`
}

type BulkUserWithGroupAndCategoryUploadErrorResponse struct {
	Status  bool   `json:"success"`
	Message string `json:"message"`
}

type BulkUserWithGroupAndCategoryDownload struct {
	Status           bool   `json:"success"`
	Message          string `json:"message"`
	UploadedFileName string `json:"uploadedfilename"`
	OriginalFileName string `json:"originalfilename"`
}

func ThrowBulkUserWithGroupAndCategoryUploadJSONResponse(response BulkUserWithGroupAndCategoryUploadResponse, w http.ResponseWriter) {
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		logger.Log.Println("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func ThrowBulkUserWithGroupAndCategoryUploadJSONErrorResponse(response BulkUserWithGroupAndCategoryUploadErrorResponse, w http.ResponseWriter) {
	//log.Println(response)
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		logger.Log.Println("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
func ThrowBulkUserWithGroupAndCategoryDownloadResponse(response BulkUserWithGroupAndCategoryDownload, w http.ResponseWriter) {
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func (w *UserWithGroupAndCategoryEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}

func (w *UserWithGroupAndCategoryForBulkUploadEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
