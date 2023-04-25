package models

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"src/config"
	"src/dao"
	"src/entities"
	"src/fileutils"
	"src/logger"
	Logger "src/logger"
	"strings"
)

var db *sql.DB

func ChangeStatusFromStartToNew() {

	wd, err := os.Getwd() // to get working directory
	if err != nil {
		Logger.Log.Println(err)
	}
	//log.Println(wd)
	contextPath := strings.ReplaceAll(wd, "\\", "/") // replacing backslash by  forwardslash
	//log.Println(contextPath)
	props, err := fileutils.ReadPropertiesFile(contextPath + "/resource/application.properties")
	if err != nil {
		Logger.Log.Println(err)
	}

	if db == nil {
		dbcon, err := config.GetDB()
		if err != nil {
			logger.Log.Println("Error in DBConnection")
			// return false, err, "Something Went Wrong"
		}
		db = dbcon
	}
	dataAccess := dao.DbConn{DB: db}

	clientIds, orgnIds, recordIds, err := dataAccess.GetRecordInStartStatus()
	if err != nil {
		logger.Log.Println(err)
		// return false, err, "Something Went Wrong"
	}

	for i := 0; i < len(recordIds); i++ {
		logger.Log.Println("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&& STARTED FRO RECORD &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&", recordIds[i])
		clientID := clientIds[i]
		orgnID := orgnIds[i]
		id := recordIds[i]
		typeID, err := dataAccess.GetRecordTypeID(clientID, orgnID, id)
		if err != nil {
			logger.Log.Println(err)
			// return false, err, "Something Went Wrong"
		}
		if typeID == 0 {
			// return false, err, "Something Went Wrong"
			logger.Log.Println("Not Found Typeid..........For Recordid", typeID,recordIds[i])

			continue
		}
		currentgrp, currentuID, err := dataAccess.GetRecordCurrentGrpID(clientID, orgnID, id)
		if err != nil {
			logger.Log.Println(err)
			// return false, err, "Something Went Wrong"
		}
		logger.Log.Println("grout,user:", currentgrp, currentuID)
		if currentgrp == 0 || currentuID == 0 {
			logger.Log.Println("Not Found Current Support Group or user..........For Recordid", currentgrp, currentuID,recordIds[i])
			// return false, err, "Something Went Wrong"
			continue
		}
		workingtypeID, workingcatID, err := dataAccess.GetWokinglabel(clientID, orgnID, id)
		if err != nil {
			logger.Log.Println(err)
			// return false, err, "Something Went Wrong"
		}
		if workingtypeID == 0 || workingcatID == 0 {
			logger.Log.Println("Not Found workingtypeID or workingcatID..........For recordid", workingtypeID, workingcatID,recordIds[i])
			// return false, err, "Something Went Wrong"
			continue
		}
		reqbd := &entities.RequestBody{}

		reqbd.ClientID = clientID
		reqbd.MstorgnhirarchyID = orgnID
		reqbd.RecorddifftypeID = workingtypeID
		reqbd.RecorddiffID = workingcatID
		reqbd.PreviousstateID = -1
		reqbd.CurrentstateID = 0
		reqbd.TransactionID = id
		reqbd.CreatedgroupID = currentgrp
		reqbd.MstgroupID = currentgrp
		reqbd.MstuserID = currentuID
		reqbd.UserID = currentuID
		reqbd.Manualstateselection = 0
		//Record status request body --> &{2 2 2 4 2 4 1065 1 1 4 false 0}

		postBody, _ := json.Marshal(reqbd)
		logger.Log.Println("Record status request body -->", reqbd)
		responseBody := bytes.NewBuffer(postBody)
		resp, err := http.Post(props["MoveWorkflowURL"], "application/json", responseBody)
		if err != nil {
			logger.Log.Println("Error is ---111111111111111111------>", err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logger.Log.Println("Error is --22222222222222222222------->", err)
		}
		sb := string(body)
		wfres := entities.WorkflowResponse{}
		json.Unmarshal([]byte(sb), &wfres)
		var workflowflag = wfres.Success
		//var errormsg = wfres.Message
		logger.Log.Println("Error is --333333333333333333333333333333333------->", workflowflag)
		if workflowflag == true {
			logger.Log.Println("Success:", wfres)
		}
		logger.Log.Println("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&& ENDED &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&", recordIds[i])
	}
	// Get TicketType

}
