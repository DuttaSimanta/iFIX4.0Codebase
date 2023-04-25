package entities

import (
	"encoding/json"
	//"io"
	"iFIX/ifix/logger"
	Logger "iFIX/ifix/logger"
	"log"
	"net/http"
	//"Errors"
)

type EmailAPIResponse struct {
	Status    bool     `json:"success"`
	Message   string   `json:"message"`
	Delimeter []string `json:"delimeter"`
}

// ErrorResponse Structure used to handle error  response using json
type EmailErrorResponse struct {
	Status  bool   `json:"success"`
	Message string `json:"message"`
	//Response []string `json:"response"`
}

//Service User Fetch Entities
type ServiceUserEntity struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	LoginName string `json:"loginname"`
	GroupID   int64  `json:"serviceusergroupid"`
}

type ServiceUserEntities struct {
	Values []ServiceUserEntity `json:"values"`
}

type ServiceUserResponse struct {
	Success bool                `json:"success"`
	Message string              `json:"message"`
	Details ServiceUserEntities `json:"details"`
}

//categoryFetch Entities

type Category struct {
	ID                   int64  `json:"id"`
	CategoryName         string `json:"name"`
	CategoryParentIds    string `json:"parentcategoryids"`
	CategoryParentNames  string `json:"parentcategorynames"`
	CategoryNameWithPath string `json:"categorywithpath"`
}
type CategoryList struct {
	Values []Category `json:"values"`
}
type CategoryListResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Details CategoryList `json:"details"`
}
type MstEmailTicket struct {
	ID                 int64  `json:"id"`
	ClientID           int64  `json:"clientid"`
	OrgID              int64  `json:"mstorgnhirarchyid"`
	TicketDiffTypeID   int64  `json:"mstrecorddifftypeid"`
	TicketDiffID       int64  `json:"mstrecorddiffid"`
	CategoryDiffTypeID int64  `json:"categorydifftypeid"`
	CategoryLevelID    int64  `json:"categorylevelid"`
	LastCategoryID     int64  `json:"lastcategoryid"`
	LastCategoryName   string `json:"lastcategoryname"`
	CategoryIDList     string `json:"categoryidlist"`
	CategoryNameList   string `json:"categorynamelist"`
	CategoryWithPath   string `json:"categorywithpath"`
	ServiceUserID      int64  `json:"serviceuserid"`
	ServiceUserGroupID int64  `json:"serviceusergroupid"`
	SenderEmail        string `json:"senderemail"`
	SenderDomain       string `json:"senderdomain"`
	EmailSubKeyword    string `json:"emailsubkeyword"`
	PriorityID         int64  `json:"priorityid"`
	SenderTypeID       int64  `json:"sendertypeid"`
	SenderTypeSeq      int64  `json:"sendertypeseq"`
	Delimiter          string `json:"delimiter"`
	CreatedByID        int64  `json:"createdbyid"`
	DefaultSeq         int64  `json:"defaultseq"`
}
type MstEmailTicketList struct {
	Values []MstEmailTicket `json:"values"`
}
type MstEmailTicketListResponse struct {
	Success bool               `json:"success"`
	Message string             `json:"message"`
	Details MstEmailTicketList `json:"details"`
}

func ThrowMstEmailTicketResponse(successMessage string, responseData MstEmailTicketList, w http.ResponseWriter, success bool) {
	var response = MstEmailTicketListResponse{}
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

type EmailBaseConfig struct {
	ID              int64  `json:"id"`
	ClientID        int64  `json:"clientid"`
	OrgID           int64  `json:"mstorgnhirarchyid"`
	OrgTypeID       int64  `json:"mstorgnhirarchytypeid"`
	DiffTypeSeq     int64  `json:"seqno"`
	DelimiterHeader string `json:"delimiterheader"`
	DelimiterVal    string `json:"delimiterval"`
	// SenderType              string   `json:"sendertypeheader"`
	// SenderTypeAndDiffSeq    []string `json:"sendertypeanddiffseq"`
	// SenderTypeVal           []string `json:"sendertypeval"`
}
type EmailTicketBaseConfig struct {
	ID         int64  `json:"id"`
	ClientID   int64  `json:"clientid"`
	OrgID      int64  `json:"mstorgnhirarchyid"`
	ClientName string `json:"clientname"`
	OrgName    string `json:"orgname"`
	TypeName   string `json:"typename"`
	Name       string `json:"name"`
}

type EmailTicketBaseConfigtList struct {
	Total  int64                   `json:"total"`
	Values []EmailTicketBaseConfig `json:"values"`
}
type EmailTicketBaseConfigtListResponse struct {
	Success bool                       `json:"success"`
	Message string                     `json:"message"`
	Details EmailTicketBaseConfigtList `json:"details"`
}
type MstEmailTicketConfig struct {
	ID                 int64  `json:"id"`
	ClientID           int64  `json:"clientid"`
	OrgID              int64  `json:"mstorgnhirarchyid"`
	ClientName         string `json:"clientname"`
	OrgName            string `json:"orgname"`
	TicketTypename     string `json:"tickettypename"`
	CategoryWithPath   string `json:"categorywithpath"`
	ServiceUserName    string `json:"serviceusername"`
	SenderEmail        string `json:"senderemail"`
	SenderDomain       string `json:"senderdomain"`
	EmailSubKeyword    string `json:"emailsubkeyword"`
	Delimiter          string `json:"delimiter"`
	CreatedBynName     string `json:"createdByname"`
	SenderType         string `json:"sendertype"`
	TicketDiffTypeID   int64  `json:"mstrecorddifftypeid"`
	TicketDiffID       int64  `json:"mstrecorddiffid"`
	CategoryDiffTypeID int64  `json:"categorydifftypeid"`
	CategoryLevelID    int64  `json:"categorylevelid"`
	LastCategoryID     int64  `json:"lastcategoryid"`
	LastCategoryName   string `json:"lastcategoryname"`
	CategoryIDList     string `json:"categoryidlist"`
	CategoryNameList   string `json:"categorynamelist"`
	ServiceUserID      int64  `json:"serviceuserid"`
	ServiceUserGroupID int64  `json:"serviceusergroupid"`
	DefaultSeq         int64  `json:"defaultseq"`
	SenderTypeSeq      int64  `json:"sendertypeseq"`
}
type MstEmailTicketConfigtList struct {
	Total  int64                  `json:"total"`
	Values []MstEmailTicketConfig `json:"values"`
}
type MstEmailTicketConfigtListResponse struct {
	Success bool                      `json:"success"`
	Message string                    `json:"message"`
	Details MstEmailTicketConfigtList `json:"details"`
}

func ThrowMstEmailTicketConfigtResponse(successMessage string, responseData MstEmailTicketConfigtList, w http.ResponseWriter, success bool) {
	var response = MstEmailTicketConfigtListResponse{}
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
func ThrowMstEmailTicketBaseConfigtResponse(successMessage string, responseData EmailTicketBaseConfigtList, w http.ResponseWriter, success bool) {
	var response = EmailTicketBaseConfigtListResponse{}
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

func ThrowServiceUsersResponse(successMessage string, responseData ServiceUserEntities, w http.ResponseWriter, success bool) {
	var response = ServiceUserResponse{}
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

func ThrowCategoryListResponse(successMessage string, responseData CategoryList, w http.ResponseWriter, success bool) {
	var response = CategoryListResponse{}
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

func EmailBlankPathCheckResponse() EmailAPIResponse {
	var response = EmailAPIResponse{}
	response.Status = false
	response.Message = "404 not found."
	log.Println("Blank request called")
	Logger.Log.Println("ERROR: Blank request called")
	return response
}
func EmailThrowJSONResponse(response EmailAPIResponse, w http.ResponseWriter) {
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
func ThrowJSONErrorResponse(response EmailErrorResponse, w http.ResponseWriter) {
	//log.Println(response)
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// NotPostMethodResponse function is used to return not post method response
func EmailNotPostMethodResponse() EmailAPIResponse {
	var response = EmailAPIResponse{}
	response.Status = false
	response.Message = "405 method not allowed."
	return response
}
