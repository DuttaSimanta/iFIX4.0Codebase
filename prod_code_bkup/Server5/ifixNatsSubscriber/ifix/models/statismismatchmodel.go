package models

import (
	// "fmt"

	"bytes"
	"encoding/json"
	"fmt"
	"ifixNatsSubscriber/ifix/dao"
	"ifixNatsSubscriber/ifix/dbconfig"
	"ifixNatsSubscriber/ifix/entities"
	"ifixNatsSubscriber/ifix/logger"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var lock = &sync.Mutex{}

func Changestatus(tz entities.Workflowentity) (error, bool, string) {
	postBody, _ := json.Marshal(map[string]int64{"clientid": tz.Clientid, "mstorgnhirarchyid": tz.Mstorgnhirarchyid, "recordid": tz.Transactionid, "reordstatusid": tz.Currentstateid, "userid": tz.Userid, "usergroupid": tz.Createdgroupid, "changestatus": tz.Changestatus, "issrrequestor": tz.Issrrequestor})
	responseBody := bytes.NewBuffer(postBody)
	logger.Log.Println("changestatus,", tz.Changestatus, tz.Transactionid)
	log.Println("changestatus,", tz.Changestatus, tz.Transactionid)
	logger.Log.Println("postBody       --->", responseBody)
	log.Println("postBody       --->", responseBody)
	resp, err := http.Post(dbconfig.RECORD_URL+"/updaterecordstatus", "application/json", responseBody)
	if err != nil {
		logger.Log.Println("An Error Occured --->", err)
		log.Println("An Error Occured --->", err)
		return err, false, ""
	}
	defer resp.Body.Close()
	//Read the response body
	response := entities.RecordstatusResponeData{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Log.Println("response body ------> ", err)
		log.Println("response body ------> ", err)
		return err, false, ""
	}
	sb := string(body)
	errr := json.Unmarshal(body, &response)
	if errr != nil {
		logger.Log.Println(errr)
	}
	logger.Log.Println("sb body value is --->", sb)
	log.Println("sb body value is --->", sb)
	return err, response.Success, sb

}
func ExecuteRemainingUpdate(tz entities.RecordDetailsEntity) {

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
	// var bugTickets []entities.RecordDetailsEntity
	// temp := entities.RecordDetailsEntity{ClientID: 2, Mstorgnhirarchyid: 15, RecordID: 7173, Requestid: 7052}
	// bugTickets, err := dataccess.GetBugRecords()
	bugTickets := []entities.RecordDetailsEntity{}
	bugTickets = append(bugTickets, tz)
	// bugTickets = append(bugTickets, temp)
	for i := 0; i < len(bugTickets); i++ {
		logger.Log.Println("Recordid:", bugTickets[i].RecordID)
		workflowHistory, _ := dataccess.GetWorkflowState(bugTickets[i])
		recordHistory, _ := dataccess.GetRecordState(bugTickets[i])
		workflowHistorySorted := specialsort(workflowHistory)
		recordHistory = specialsort(recordHistory)

		// for i:=0;i<len(workflowHistorySorted);i++{
		//   if(workflowHistorySorted[i]!=recordHistory[i+1])
		// }
		if len(workflowHistorySorted) == 0 || len(recordHistory) == 0 {
			errentity := entities.ErrorEntity{}

			errentity.ClientID = bugTickets[i].ClientID
			errentity.Mstorgnhirarchyid = bugTickets[i].Mstorgnhirarchyid
			errentity.RecordID = bugTickets[i].RecordID
			errentity.Requestjson = "{}"
			errentity.Responsejson = "{}"
			errentity.Comment = fmt.Sprintf("%v:%v,%v:%v", "Length Zero : workflow history length is ", strconv.Itoa(len(workflowHistorySorted)), " and record length is ", strconv.Itoa(len(recordHistory)))
			errentity.Isdefect = "Y"
			dataccess.InsertCommentOrErr(&errentity)
			continue
		}
		if recordHistory[len(recordHistory)-1].StatusSeq == 0 {
			///Code for Start-State
			errentity := entities.ErrorEntity{}
			statuschangeentity := entities.Workflowentity{}
			statuschangeentity.Clientid = bugTickets[i].ClientID
			statuschangeentity.Mstorgnhirarchyid = bugTickets[i].Mstorgnhirarchyid
			statuschangeentity.Transactionid = bugTickets[i].RecordID
			statuschangeentity.Currentstateid = workflowHistorySorted[0].Stateid
			statuschangeentity.Createdgroupid = workflowHistorySorted[0].Createdgroupid
			statuschangeentity.Userid = workflowHistorySorted[0].Userid

			err, success, resp := Changestatus(statuschangeentity)
			req, _ := json.Marshal(statuschangeentity)
			errentity.ClientID = bugTickets[i].ClientID
			errentity.Mstorgnhirarchyid = bugTickets[i].Mstorgnhirarchyid
			errentity.RecordID = bugTickets[i].RecordID
			errentity.Stateid = workflowHistorySorted[0].Stateid
			errentity.Requestjson = string(req)
			errentity.Responsejson = resp
			if success {
				errentity.Isdefect = "N"
				errentity.Comment = "From Start" + " To " + workflowHistorySorted[0].Statusname
				logger.Log.Println("Success:" + errentity.Comment)
				dataccess.InsertCommentOrErr(&errentity)
			} else {
				errentity.Isdefect = "Y"
				// errentity.Comment = "Error To Change Status From" + workflowHistorySorted[j-1].Statusname + " To " + workflowHistorySorted[j].Statusname + string(err)
				errentity.Comment = fmt.Sprintf("%s:%s", "Error To Change Status From Start"+" To "+workflowHistorySorted[0].Statusname, err)
				logger.Log.Println(errentity.Comment)
				dataccess.InsertCommentOrErr(&errentity)
				continue
			}
			dataccess.InsertCommentOrErr(&errentity)
			if err != nil {
				logger.Log.Println("in side Continue")
				continue
			}

			recordHistory, _ = dataccess.GetRecordState(bugTickets[i])
		}
		//index := Findstatusindex(workflowHistorySorted, recordHistory[len(recordHistory)-1])
		length := len(recordHistory) - 1
		// count1 := countlaststatus(recordHistory)
		// index := Findstatusindex(workflowHistorySorted, count1, recordHistory[len(recordHistory)-1])
		logger.Log.Println("WorkflowSort:", workflowHistorySorted)
		logger.Log.Println("RecordSort:", recordHistory)
		logger.Log.Println("Length:", length)
		if len(workflowHistorySorted) < len(recordHistory) {
			errentity := entities.ErrorEntity{}

			errentity.ClientID = bugTickets[i].ClientID
			errentity.Mstorgnhirarchyid = bugTickets[i].Mstorgnhirarchyid
			errentity.RecordID = bugTickets[i].RecordID
			errentity.Requestjson = "{}"
			errentity.Responsejson = "{}"
			errentity.Comment = fmt.Sprintf("%v:%v,%v:%v", "Record length is greater than workflow length : workflow history length is ", strconv.Itoa(len(workflowHistorySorted)), " and record length is ", strconv.Itoa(len(recordHistory)))
			errentity.Isdefect = "Y"
			dataccess.InsertCommentOrErr(&errentity)
			continue
		}
		for j := 0; j < len(workflowHistorySorted); j++ {
			// call api
			// logger.Log.Println("Hiii", length)

			errentity := entities.ErrorEntity{}
			if length > 0 {
				logger.Log.Println(j+1, length, "Time", workflowHistorySorted[j].Statusid, recordHistory[j+1].Statusid)

				if workflowHistorySorted[j].Statusid != recordHistory[j+1].Statusid {
					if len(workflowHistorySorted)-1 == len(recordHistory) && workflowHistorySorted[len(workflowHistorySorted)-2].Statusid == recordHistory[len(recordHistory)-1].Statusid {
						logger.Log.Println("============================================================================================================")
						statuschangeentity := entities.Workflowentity{}
						statuschangeentity.Clientid = bugTickets[i].ClientID
						statuschangeentity.Mstorgnhirarchyid = bugTickets[i].Mstorgnhirarchyid
						statuschangeentity.Transactionid = bugTickets[i].RecordID
						statuschangeentity.Currentstateid = workflowHistorySorted[len(workflowHistorySorted)-1].Stateid
						statuschangeentity.Createdgroupid = workflowHistorySorted[len(workflowHistorySorted)-2].Createdgroupid
						// userid := finduserid(workflowHistory, workflowHistorySorted[j-1].Index)
						statuschangeentity.Userid = workflowHistory[(workflowHistorySorted[len(workflowHistorySorted)-2].Index)+1].Userid
						err, success, resp := Changestatus(statuschangeentity)

						req, _ := json.Marshal(statuschangeentity)
						errentity.ClientID = bugTickets[i].ClientID
						errentity.Mstorgnhirarchyid = bugTickets[i].Mstorgnhirarchyid
						errentity.RecordID = bugTickets[i].RecordID
						errentity.Stateid = workflowHistorySorted[len(workflowHistorySorted)-1].Stateid
						errentity.Requestjson = string(req)
						errentity.Responsejson = resp
						logger.Log.Println("============================================================================================================")

						if success {
							errentity.Isdefect = "N"
							errentity.Comment = "Defect Ticket:From" + workflowHistorySorted[len(workflowHistorySorted)-2].Statusname + " To " + workflowHistorySorted[len(workflowHistorySorted)-1].Statusname
							logger.Log.Println("Success:Defect Ticket"+errentity.Comment, bugTickets[i].RecordID)
							_, err1 := dataccess.InsertCommentOrErr(&errentity)
							if err1 != nil {
								// continue
								logger.Log.Println(err)
								break
							}
							break
						} else {
							errentity.Isdefect = "Y"
							// errentity.Comment = "Error To Change Status From" + workflowHistorySorted[j-1].Statusname + " To " + workflowHistorySorted[j].Statusname + string(err)
							errentity.Comment = fmt.Sprintf("%s:%s", "Defect Ticket : Error To Change Status From"+workflowHistorySorted[len(workflowHistorySorted)-2].Statusname+" To "+workflowHistorySorted[len(workflowHistorySorted)-1].Statusname, err)
							logger.Log.Println(errentity.Comment)
							_, err1 := dataccess.InsertCommentOrErr(&errentity)
							if err1 != nil {
								logger.Log.Println(err)
							}
							break
						}

					} else {
						errentity.ClientID = bugTickets[i].ClientID
						errentity.Mstorgnhirarchyid = bugTickets[i].Mstorgnhirarchyid
						errentity.RecordID = bugTickets[i].RecordID
						errentity.Stateid = workflowHistorySorted[j].Stateid
						errentity.Requestjson = ""
						errentity.Responsejson = ""
						//					errentity.Comment = "Not Matching State From Both side: Workflow Status:" + string(workflowHistorySorted[j].Statusid) + string(recordHistory[j+1].Statusid)
						errentity.Comment = fmt.Sprintf("%s:%d,%s:%d", "Not Matching State From Both side: Workflow Status:", workflowHistorySorted[j].Statusid, " Record Status", recordHistory[j+1].Statusid)
						errentity.Isdefect = "Y"
						dataccess.InsertCommentOrErr(&errentity)
						j = len(workflowHistorySorted)
						break
					}
				}
				if workflowHistorySorted[j].Statusid == recordHistory[j+1].Statusid {
					length--
					continue
				}
			} else {

				statuschangeentity := entities.Workflowentity{}
				statuschangeentity.Clientid = bugTickets[i].ClientID
				statuschangeentity.Mstorgnhirarchyid = bugTickets[i].Mstorgnhirarchyid
				statuschangeentity.Transactionid = bugTickets[i].RecordID
				statuschangeentity.Currentstateid = workflowHistorySorted[j].Stateid
				statuschangeentity.Createdgroupid = workflowHistorySorted[j-1].Createdgroupid
				//	userid := finduserid(workflowHistory, workflowHistorySorted[j-1].Index)
				//	statuschangeentity.Userid = userid
				statuschangeentity.Userid = workflowHistory[(workflowHistorySorted[j-1].Index)+1].Userid

				err, success, resp := Changestatus(statuschangeentity)

				req, _ := json.Marshal(statuschangeentity)
				errentity.ClientID = bugTickets[i].ClientID
				errentity.Mstorgnhirarchyid = bugTickets[i].Mstorgnhirarchyid
				errentity.RecordID = bugTickets[i].RecordID
				errentity.Stateid = workflowHistorySorted[j].Stateid
				errentity.Requestjson = string(req)
				errentity.Responsejson = resp
				if success {
					errentity.Isdefect = "N"
					errentity.Comment = "From" + workflowHistorySorted[j-1].Statusname + " To " + workflowHistorySorted[j].Statusname
					logger.Log.Println("Success:"+errentity.Comment, bugTickets[i].RecordID)
					_, err1 := dataccess.InsertCommentOrErr(&errentity)
					if err1 != nil {
						continue
					}
				} else {
					errentity.Isdefect = "Y"
					// errentity.Comment = "Error To Change Status From" + workflowHistorySorted[j-1].Statusname + " To " + workflowHistorySorted[j].Statusname + string(err)
					errentity.Comment = fmt.Sprintf("%s:%s", "Error To Change Status From"+workflowHistorySorted[j-1].Statusname+" To "+workflowHistorySorted[j].Statusname, err)
					logger.Log.Println(errentity.Comment)
					_, err1 := dataccess.InsertCommentOrErr(&errentity)
					if err1 != nil {
						logger.Log.Println(err)
					}
					break
				}

			}

		}

	}
	p("inside ExecuteRemainingUpdate &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&  Ended &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&")
}
func specialsort(arr []entities.StatedetailEntity) []entities.StatedetailEntity {
	temp := []entities.StatedetailEntity{}
	for i := 0; i < len(arr); i++ {
		if i == len(arr)-1 {
			temp = append(temp, arr[i])
		} else if arr[i].Statusid != arr[i+1].Statusid {
			temp = append(temp, arr[i])
		}
	}
	return temp
}
func finduserid(arr []entities.StatedetailEntity, index int) int64 {
	for i := index; i >= 0; i-- {
		if arr[i].Mstuserid != 0 {
			return arr[i].Mstuserid
		}
	}
	return arr[index].Userid
}
func ExecuteSRRemainingUpdate(tz entities.RecordDetailsEntity) {

	p := logger.Log.Println
	p("inside ExecuteRemainingUpdate &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&& SR ticket Started &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&")
	lock.Lock()
	defer lock.Unlock()
	db, err := dbconfig.ConnectMySqlDb()
	if err != nil {
		logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>", err)
		panic(err)
	}
	dataccess := dao.DbConn{DB: db}
	// var bugTickets []entities.RecordDetailsEntity
	// temp := entities.RecordDetailsEntity{ClientID: 2, Mstorgnhirarchyid: 15, RecordID: 7173, Requestid: 7052}
	// bugTickets, err := dataccess.GetBugRecords()
	bugTickets := []entities.RecordDetailsEntity{}
	bugTickets = append(bugTickets, tz)
	// bugTickets = append(bugTickets, temp)
	for i := 0; i < len(bugTickets); i++ {
		logger.Log.Println("Recordid:", bugTickets[i].RecordID)
		workflowHistory, _ := dataccess.GetWorkflowState(bugTickets[i])
		recordHistory, _ := dataccess.GetRecordState(bugTickets[i])
		workflowHistorySorted := specialsort(workflowHistory)
		recordHistory = specialsort(recordHistory)

		// for i:=0;i<len(workflowHistorySorted);i++{
		//   if(workflowHistorySorted[i]!=recordHistory[i+1])
		// }
		// var flag int64
		if workflowHistorySorted[len(workflowHistorySorted)-1].StatusSeq == 15 && recordHistory[len(recordHistory)-1].StatusSeq == 12 {
			// taskTickets, _ := dataccess.GetTaskTicket(bugTickets[i])
			// if len(taskTickets)>0{
			// 	for j:=0;j<len(taskTickets);j++{
			// 		taskTicketsentity := []entities.RecordDetailsEntity{}
			//         taskTicketsentity=append(taskTicketsentity, entities.RecordDetailsEntity{RecordID: taskTickets[j]})
			// 		workflowHistoryOfTask, _ := dataccess.GetWorkflowState(taskTicketsentity[i])
			//         if workflowHistoryOfTask[len(workflowHistoryOfTask)-1].Statusname!="Inactive"{
			// 			flag=1
			// 		}
			// 	}
			// }
			// if flag==0{
			statuschangeentity := entities.Workflowentity{}
			statuschangeentity.Clientid = bugTickets[i].ClientID
			statuschangeentity.Mstorgnhirarchyid = bugTickets[i].Mstorgnhirarchyid
			statuschangeentity.Transactionid = bugTickets[i].RecordID
			statuschangeentity.Currentstateid = workflowHistorySorted[len(workflowHistorySorted)-1].Stateid
			statuschangeentity.Createdgroupid = workflowHistorySorted[len(workflowHistorySorted)-2].Createdgroupid
			// userid := finduserid(workflowHistory, workflowHistorySorted[j-1].Index)
			statuschangeentity.Userid = workflowHistory[(workflowHistorySorted[len(workflowHistorySorted)-2].Index)+1].Userid
			err, success, resp := Changestatus(statuschangeentity)

			req, _ := json.Marshal(statuschangeentity)

			errentity := entities.ErrorEntity{}
			errentity.ClientID = bugTickets[i].ClientID
			errentity.Mstorgnhirarchyid = bugTickets[i].Mstorgnhirarchyid
			errentity.RecordID = bugTickets[i].RecordID
			errentity.Stateid = workflowHistorySorted[len(workflowHistorySorted)-1].Stateid
			errentity.Requestjson = string(req)
			errentity.Responsejson = resp
			logger.Log.Println("============================================================================================================")

			if success {
				errentity.Isdefect = "N"
				errentity.Comment = "SR:Defect Ticket:From" + workflowHistorySorted[len(workflowHistorySorted)-2].Statusname + " To " + workflowHistorySorted[len(workflowHistorySorted)-1].Statusname
				logger.Log.Println("Success:Defect Ticket"+errentity.Comment, bugTickets[i].RecordID)
				_, err1 := dataccess.InsertCommentOrErr(&errentity)
				if err1 != nil {
					// continue
					logger.Log.Println(err)
					break
				}
				break
			} else {
				errentity.Isdefect = "Y"
				// errentity.Comment = "Error To Change Status From" + workflowHistorySorted[j-1].Statusname + " To " + workflowHistorySorted[j].Statusname + string(err)
				errentity.Comment = fmt.Sprintf("%s:%s", "SR:Defect Ticket : Error To Change Status From"+workflowHistorySorted[len(workflowHistorySorted)-2].Statusname+" To "+workflowHistorySorted[len(workflowHistorySorted)-1].Statusname, err)
				logger.Log.Println(errentity.Comment)
				_, err1 := dataccess.InsertCommentOrErr(&errentity)
				if err1 != nil {
					logger.Log.Println(err)
				}
				break
			}
			// }
		} else if workflowHistorySorted[len(workflowHistorySorted)-1].StatusSeq == 17 && (recordHistory[len(recordHistory)-1].StatusSeq == 3 || recordHistory[len(recordHistory)-1].StatusSeq == 4 || recordHistory[len(recordHistory)-1].StatusSeq == 5 || recordHistory[len(recordHistory)-1].StatusSeq == 7) {

			childId, err := dataccess.Getchildrecordids(tz.RecordID)
			if err != nil {
				log.Println("database connection failure", err)
				// return 0, false, err, "Something Went Wrong"
			}
			paarentid := []int64{tz.RecordID}
			if len(childId) == 1 {
				userid, grpid, _ := dataccess.Getuserandgroup(tz.RecordID)
				reqbd := entities.ParentchildEntity{}
				reqbd.Parentid = childId[0]
				reqbd.Childids = paarentid
				reqbd.Userid = userid
				reqbd.Isupdate = false
				reqbd.Createdgroupid = grpid
				postBody, _ := json.Marshal(reqbd)
				success, msg := updateSr(reqbd)
				errentity := entities.ErrorEntity{}
				errentity.ClientID = bugTickets[i].ClientID
				errentity.Mstorgnhirarchyid = bugTickets[i].Mstorgnhirarchyid
				errentity.RecordID = bugTickets[i].RecordID
				errentity.Stateid = workflowHistorySorted[len(workflowHistorySorted)-1].Stateid
				errentity.Requestjson = string(postBody)
				errentity.Responsejson = msg
				logger.Log.Println("============================================================================================================")

				if success {
					errentity.Isdefect = "N"
					errentity.Comment = "SR:Defect Ticket:From" + recordHistory[len(recordHistory)-1].Statusname + " To " + workflowHistorySorted[len(workflowHistorySorted)-1].Statusname
					logger.Log.Println("Success:Defect Ticket"+errentity.Comment, bugTickets[i].RecordID)
					_, err1 := dataccess.InsertCommentOrErr(&errentity)
					if err1 != nil {
						// continue
						logger.Log.Println(err)
						break
					}
					break
				} else {
					errentity.Isdefect = "Y"
					// errentity.Comment = "Error To Change Status From" + workflowHistorySorted[j-1].Statusname + " To " + workflowHistorySorted[j].Statusname + string(err)
					errentity.Comment = fmt.Sprintf("%s:%s", "SR:Defect Ticket : Error To Change Status From"+recordHistory[len(recordHistory)-1].Statusname+" To "+workflowHistorySorted[len(workflowHistorySorted)-1].Statusname, err)
					logger.Log.Println(errentity.Comment)
					_, err1 := dataccess.InsertCommentOrErr(&errentity)
					if err1 != nil {
						logger.Log.Println(err)
					}
					break
				}
			}
		} else if workflowHistorySorted[len(workflowHistorySorted)-1].StatusSeq == 17 && recordHistory[len(recordHistory)-1].StatusSeq == 0 {
			// taskTickets, _ := dataccess.GetTaskTicket(bugTickets[i])
			// if len(taskTickets)>0{
			// 	for j:=0;j<len(taskTickets);j++{
			// 		taskTicketsentity := []entities.RecordDetailsEntity{}
			//         taskTicketsentity=append(taskTicketsentity, entities.RecordDetailsEntity{RecordID: taskTickets[j]})
			// 		workflowHistoryOfTask, _ := dataccess.GetWorkflowState(taskTicketsentity[i])
			//         if workflowHistoryOfTask[len(workflowHistoryOfTask)-1].Statusname!="Inactive"{
			// 			flag=1
			// 		}
			// 	}
			// }
			// if flag==0{
			statuschangeentity := entities.Workflowentity{}
			statuschangeentity.Clientid = bugTickets[i].ClientID
			statuschangeentity.Mstorgnhirarchyid = bugTickets[i].Mstorgnhirarchyid
			statuschangeentity.Transactionid = bugTickets[i].RecordID
			statuschangeentity.Currentstateid = workflowHistorySorted[len(workflowHistorySorted)-1].Stateid
			statuschangeentity.Createdgroupid = workflowHistorySorted[len(workflowHistorySorted)-2].Createdgroupid
			// userid := finduserid(workflowHistory, workflowHistorySorted[j-1].Index)
			statuschangeentity.Userid = workflowHistory[(workflowHistorySorted[len(workflowHistorySorted)-2].Index)+1].Userid
			err, success, resp := Changestatus(statuschangeentity)

			req, _ := json.Marshal(statuschangeentity)

			errentity := entities.ErrorEntity{}
			errentity.ClientID = bugTickets[i].ClientID
			errentity.Mstorgnhirarchyid = bugTickets[i].Mstorgnhirarchyid
			errentity.RecordID = bugTickets[i].RecordID
			errentity.Stateid = workflowHistorySorted[len(workflowHistorySorted)-1].Stateid
			errentity.Requestjson = string(req)
			errentity.Responsejson = resp
			logger.Log.Println("============================================================================================================")

			if success {
				errentity.Isdefect = "N"
				errentity.Comment = "SR:Defect Ticket:From" + workflowHistorySorted[len(workflowHistorySorted)-2].Statusname + " To " + workflowHistorySorted[len(workflowHistorySorted)-1].Statusname
				logger.Log.Println("Success:Defect Ticket"+errentity.Comment, bugTickets[i].RecordID)
				_, err1 := dataccess.InsertCommentOrErr(&errentity)
				if err1 != nil {
					// continue
					logger.Log.Println(err)
					break
				}
				break
			} else {
				errentity.Isdefect = "Y"
				// errentity.Comment = "Error To Change Status From" + workflowHistorySorted[j-1].Statusname + " To " + workflowHistorySorted[j].Statusname + string(err)
				errentity.Comment = fmt.Sprintf("%s:%s", "SR:Defect Ticket : Error To Change Status From"+workflowHistorySorted[len(workflowHistorySorted)-2].Statusname+" To "+workflowHistorySorted[len(workflowHistorySorted)-1].Statusname, err)
				logger.Log.Println(errentity.Comment)
				_, err1 := dataccess.InsertCommentOrErr(&errentity)
				if err1 != nil {
					logger.Log.Println(err)
				}
				break
			}
		}  else {
			errentity := entities.ErrorEntity{}

			errentity.ClientID = bugTickets[i].ClientID
			errentity.Mstorgnhirarchyid = bugTickets[i].Mstorgnhirarchyid
			errentity.RecordID = bugTickets[i].RecordID
			errentity.Requestjson = "{}"
			errentity.Responsejson = "{}"
			errentity.Comment = fmt.Sprintf("%v", "SR:This Type Of Mismatch is Not Considered")
			errentity.Isdefect = "Y"
			dataccess.InsertCommentOrErr(&errentity)
			continue
		}

	}
	p("inside ExecuteRemainingUpdate &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&& SR ticket Ended &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&")
}
func updateSr(reqbd entities.ParentchildEntity) (bool, string) {
	// reqbd := &entities.ParentchildEntity{}
	// reqbd.Parentid = RecordID
	// reqbd.Childids = parentID
	// reqbd.Userid = UserID
	// reqbd.Isupdate = false
	// reqbd.Createdgroupid = Usergroupid
	postBody, _ := json.Marshal(reqbd)

	logger.Log.Println("Record status request body -->", reqbd, string(postBody))

	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(dbconfig.MASTER_URL+"/updatetaskstatus", "application/json", responseBody)
	if err != nil {
		logger.Log.Println("An Error Occured --->", err)
		return false, "Something Went Wrong"
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Log.Println("response body ------> ", err)
		return false, "Something Went Wrong"
	}
	sb := string(body)
	wfres := entities.WorkflowResponse{}
	json.Unmarshal([]byte(sb), &wfres)
	workflowflag := wfres.Success
	errormsg := wfres.Message
	logger.Log.Println("Record status response message -->", workflowflag)
	logger.Log.Println("Record status response error message -->", errormsg)
	return workflowflag, errormsg
}
