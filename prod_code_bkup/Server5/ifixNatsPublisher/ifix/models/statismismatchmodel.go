package models

import (
	// "fmt"

	"ifixNatsPublisher/ifix/dao"
	"ifixNatsPublisher/ifix/dbconfig"
	"ifixNatsPublisher/ifix/entities"
	"ifixNatsPublisher/ifix/logger"
	"sync"
)

var lock = &sync.Mutex{}
func GetSRBugTickets() []entities.RecordDetailsEntity {
	p := logger.Log.Println
	p("inside ExecuteRemainingUpdate &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&  Started &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&")
	lock.Lock()
	defer lock.Unlock()
	db, err := dbconfig.ConnectMySqlDb()
	if err != nil {
		logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>", err)
		panic(err)
	}
	dataccess := dao.DbConn{DB: db}
	logger.Log.Println(dataccess)
	// var bugTickets []entities.RecordDetailsEntity
	// temp := entities.RecordDetailsEntity{ClientID: 2, Mstorgnhirarchyid: 27, RecordID: 103494, Requestid: 103207}
	// bugTickets = append(bugTickets, temp)

	bugTickets, err := dataccess.GetSRBugRecords()
	return bugTickets
}
func GetBugTickets() []entities.RecordDetailsEntity {
	p := logger.Log.Println
	p("inside ExecuteRemainingUpdate &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&  Started &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&")
	lock.Lock()
	defer lock.Unlock()
	db, err := dbconfig.ConnectMySqlDb()
	if err != nil {
		logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>", err)
		panic(err)
	}
	dataccess := dao.DbConn{DB: db}
//	logger.Log.Println(dataccess)
//	 var bugTickets []entities.RecordDetailsEntity
//	temp := entities.RecordDetailsEntity{ClientID: 2, Mstorgnhirarchyid: 27, RecordID: 103494, Requestid: 103207}
//	bugTickets = append(bugTickets, temp)

	 bugTickets, err := dataccess.GetBugRecords()
	return bugTickets
}
func DeleteDefectTkt(){
	p := logger.Log.Println
	p("inside  DeleteDefectTkt &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&  Started &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&")
	lock.Lock()
	defer lock.Unlock()
	db, err := dbconfig.ConnectMySqlDb()
	if err != nil {
		logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>", err)
		panic(err)
	}
	dataccess := dao.DbConn{DB: db}
	rowsaffacted,err:=dataccess.DeleteDefectTkt()
	logger.Log.Println(rowsaffacted,err)
}
// func Changestatus(tz entities.Workflowentity) (error, bool, string) {
// 	postBody, _ := json.Marshal(map[string]int64{"clientid": tz.Clientid, "mstorgnhirarchyid": tz.Mstorgnhirarchyid, "recordid": tz.Transactionid, "reordstatusid": tz.Currentstateid, "userid": tz.Userid, "usergroupid": tz.Createdgroupid, "changestatus": tz.Changestatus, "issrrequestor": tz.Issrrequestor})
// 	responseBody := bytes.NewBuffer(postBody)
// 	logger.Log.Println("changestatus,", tz.Changestatus, tz.Transactionid)
// 	log.Println("changestatus,", tz.Changestatus, tz.Transactionid)
// 	logger.Log.Println("postBody       --->", responseBody)
// 	log.Println("postBody       --->", responseBody)
// 	resp, err := http.Post(dbconfig.RECORD_URL+"/updaterecordstatus", "application/json", responseBody)
// 	if err != nil {
// 		logger.Log.Println("An Error Occured --->", err)
// 		log.Println("An Error Occured --->", err)
// 		return err, false, ""
// 	}
// 	defer resp.Body.Close()
// 	//Read the response body
// 	response := entities.RecordstatusResponeData{}
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		logger.Log.Println("response body ------> ", err)
// 		log.Println("response body ------> ", err)
// 		return err, false, ""
// 	}
// 	sb := string(body)
// 	errr := json.Unmarshal(body, &response)
// 	if errr != nil {
// 		logger.Log.Println(errr)
// 	}
// 	logger.Log.Println("sb body value is --->", sb)
// 	log.Println("sb body value is --->", sb)
// 	return err, response.Success, sb

// }
// func ExecuteRemainingUpdate() {

// 	p := logger.Log.Println
// 	p("inside ExecuteRemainingUpdate &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&  Started &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&")
// 	lock.Lock()
// 	defer lock.Unlock()
// 	db, err := dbconfig.ConnectMySqlDb()
// 	if err != nil {
// 		logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>", err)
// 		panic(err)
// 	}
// 	dataccess := dao.DbConn{DB: db}
// 	// var bugTickets []entities.RecordDetailsEntity
// 	// temp := entities.RecordDetailsEntity{ClientID: 2, Mstorgnhirarchyid: 15, RecordID: 7173, Requestid: 7052}
// 	bugTickets, err := dataccess.GetBugRecords()
// 	// bugTickets = append(bugTickets, temp)
// 	for i := 0; i < len(bugTickets); i++ {
// 		logger.Log.Println("Recordid:", bugTickets[i].RecordID)
// 		workflowHistory, _ := dataccess.GetWorkflowState(bugTickets[i])
// 		recordHistory, _ := dataccess.GetRecordState(bugTickets[i])
// 		workflowHistorySorted := specialsort(workflowHistory)
// 		// for i:=0;i<len(workflowHistorySorted);i++{
// 		//   if(workflowHistorySorted[i]!=recordHistory[i+1])
// 		// }
// 		if len(workflowHistorySorted) == 0 || len(recordHistory) == 0 {
// 			errentity := entities.ErrorEntity{}

// 			errentity.ClientID = bugTickets[i].ClientID
// 			errentity.Mstorgnhirarchyid = bugTickets[i].Mstorgnhirarchyid
// 			errentity.RecordID = bugTickets[i].RecordID
// 			errentity.Requestjson, _ = json.Marshal()
// 			errentity.Responsejson = string([]byte("{}"))

// 			errentity.Comment = "Length Is Zero : workflow history length is " + string(len(workflowHistorySorted)) + " and record length is " + string(len(recordHistory))
// 			dataccess.InsertCommentOrErr(&errentity)
// 			continue
// 		}
// 		if recordHistory[len(recordHistory)-1].StatusSeq == 0 {
// 			///Code for Start-State
// 			errentity := entities.ErrorEntity{}
// 			statuschangeentity := entities.Workflowentity{}
// 			statuschangeentity.Clientid = bugTickets[i].ClientID
// 			statuschangeentity.Mstorgnhirarchyid = bugTickets[i].Mstorgnhirarchyid
// 			statuschangeentity.Transactionid = bugTickets[i].RecordID
// 			statuschangeentity.Currentstateid = workflowHistorySorted[0].Stateid
// 			statuschangeentity.Createdgroupid = workflowHistorySorted[0].Createdgroupid
// 			statuschangeentity.Userid = workflowHistorySorted[0].Userid

// 			err, success, resp := Changestatus(statuschangeentity)
// 			req, _ := json.Marshal(statuschangeentity)
// 			errentity.ClientID = bugTickets[i].ClientID
// 			errentity.Mstorgnhirarchyid = bugTickets[i].Mstorgnhirarchyid
// 			errentity.RecordID = bugTickets[i].RecordID
// 			errentity.Stateid = workflowHistorySorted[0].Stateid
// 			errentity.Requestjson = string(req)
// 			errentity.Responsejson = resp
// 			if success {
// 				errentity.Comment = "From Start" + " To " + workflowHistorySorted[0].Statusname
// 			} else {
// 				// errentity.Comment = "Error To Change Status From" + workflowHistorySorted[j-1].Statusname + " To " + workflowHistorySorted[j].Statusname + string(err)
// 				errentity.Comment = fmt.Sprintf("%s:%s", "Error To Change Status From Start"+" To "+workflowHistorySorted[0].Statusname, err)
// 				logger.Log.Println(errentity.Comment)
// 			}
// 			_, err1 := dataccess.InsertCommentOrErr(&errentity)
// 			if err1 != nil {
// 				continue
// 			}

// 			recordHistory, _ = dataccess.GetRecordState(bugTickets[i])
// 		}
// 		//index := Findstatusindex(workflowHistorySorted, recordHistory[len(recordHistory)-1])
// 		length := len(recordHistory) - 1
// 		// count1 := countlaststatus(recordHistory)
// 		// index := Findstatusindex(workflowHistorySorted, count1, recordHistory[len(recordHistory)-1])
// 		logger.Log.Println("WorkflowSort:", workflowHistorySorted)
// 		logger.Log.Println("RecordSort:", recordHistory)
// 		logger.Log.Println("Length:", length)

// 		for j := 0; j < len(workflowHistorySorted); j++ {
// 			// call api
// 			// logger.Log.Println("Hiii", length)

// 			errentity := entities.ErrorEntity{}
// 			if length > 0 {
// 				logger.Log.Println(j+1, length, "Time", workflowHistorySorted[j].Statusid, recordHistory[j+1].Statusid)

// 				if workflowHistorySorted[j].Statusid != recordHistory[j+1].Statusid {
// 					errentity.ClientID = bugTickets[i].ClientID
// 					errentity.Mstorgnhirarchyid = bugTickets[i].Mstorgnhirarchyid
// 					errentity.RecordID = bugTickets[i].RecordID
// 					errentity.Stateid = workflowHistorySorted[j].Stateid
// 					errentity.Requestjson = ""
// 					errentity.Responsejson = ""
// 					errentity.Comment = fmt.Sprintf("%s:%s,%s:%s", "Not Matching State From Both side: Workflow Status:", workflowHistorySorted[j].Statusid, " Record Status", recordHistory[j+1].Statusid)
// 					dataccess.InsertCommentOrErr(&errentity)
// 					j = len(workflowHistorySorted)
// 					break
// 				}
// 				if workflowHistorySorted[j].Statusid == recordHistory[j+1].Statusid {
// 					length--
// 					continue
// 				}
// 			} else {

// 				statuschangeentity := entities.Workflowentity{}
// 				statuschangeentity.Clientid = bugTickets[i].ClientID
// 				statuschangeentity.Mstorgnhirarchyid = bugTickets[i].Mstorgnhirarchyid
// 				statuschangeentity.Transactionid = bugTickets[i].RecordID
// 				statuschangeentity.Currentstateid = workflowHistorySorted[j].Stateid
// 				statuschangeentity.Createdgroupid = workflowHistorySorted[j-1].Createdgroupid
// 				statuschangeentity.Userid = workflowHistorySorted[j].Userid

// 				err, success, resp := Changestatus(statuschangeentity)

// 				req, _ := json.Marshal(statuschangeentity)
// 				errentity.ClientID = bugTickets[i].ClientID
// 				errentity.Mstorgnhirarchyid = bugTickets[i].Mstorgnhirarchyid
// 				errentity.RecordID = bugTickets[i].RecordID
// 				errentity.Stateid = workflowHistorySorted[j].Stateid
// 				errentity.Requestjson = string(req)
// 				errentity.Responsejson = resp
// 				if success {
// 					errentity.Comment = "From" + workflowHistorySorted[j-1].Statusname + " To " + workflowHistorySorted[j].Statusname
// 					_, err1 := dataccess.InsertCommentOrErr(&errentity)
// 					if err1 != nil {
// 						continue
// 					}
// 				} else {
// 					// errentity.Comment = "Error To Change Status From" + workflowHistorySorted[j-1].Statusname + " To " + workflowHistorySorted[j].Statusname + string(err)
// 					errentity.Comment = fmt.Sprintf("%s:%s", "Error To Change Status From"+workflowHistorySorted[j-1].Statusname+" To "+workflowHistorySorted[j].Statusname, err)
// 					logger.Log.Println(errentity.Comment)
// 					_, err1 := dataccess.InsertCommentOrErr(&errentity)
// 					if err1 != nil {
// 						logger.Log.Println(err)
// 					}
// 					break
// 				}

// 			}

// 		}

// 	}
// 	p("inside ExecuteRemainingUpdate &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&  Ended &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&")
// }
// func specialsort(arr []entities.StatedetailEntity) []entities.StatedetailEntity {
// 	temp := []entities.StatedetailEntity{}
// 	for i := 0; i < len(arr); i++ {
// 		if i == 0 {
// 			temp = append(temp, arr[i])
// 		} else if arr[i-1].Statusid != arr[i].Statusid {
// 			temp = append(temp, arr[i])
// 		}
// 	}
// 	return temp
// }
// func Findstatusindex(status []entities.StatedetailEntity, checkstatus entities.StatedetailEntity) int {

// 	for i := 0; i < len(status); i++ {
// 		if status[i].Statusid == checkstatus.Statusid {
// 			return i
// 		}
// 	}
// 	return 0
// }

// // func countlaststatus(arr []entities.StatedetailEntity) int {
// // 	k := 0
// // 	for i := 0; i < len(arr); i++ {
// // 		if arr[i].Statusid == arr[len(arr)-1].Statusid {
// // 			k++
// // 		}
// // 	}
// // 	return k
// // }
// // func Findstatusindex(status []entities.StatedetailEntity, count int, checkstatus entities.StatedetailEntity) int {

// // 	for i := 0; i < len(status); i++ {
// // 		if status[i].Statusid == checkstatus.Statusid {
// // 			count--
// // 			if count == 0 {
// // 				return i
// // 			}
// // 		}
// // 	}
// // 	return 0
// // }
