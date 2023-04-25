//***************************//
// Package models
// Date Of Creation: 18/12/2020
// Authour Name: Subham Chatterjee
// History: N/A
// Synopsis: This file is used for workflow related works. It is used as model. All the business logic is written here.
// Functions:

// Global Variable: N/A
// Version: 1.0.0
//***************************//
package models

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/utility"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func InsertProcessDelegateUser(tz *entities.Workflowentity) (int64, bool, error, string) {
	log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	id, err := dataAccess.InsertProcessDelegateUser(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	return id, true, err, ""
}
func Bulkapprovalfortickets(tz *entities.Workflowentity) (int64, bool, error, string) {
	//lock.Lock()
	//defer lock.Unlock()
	//db, err := config.ConnectMySqlDbSingleton()
	//if err != nil {
	//	log.Println("database connection failure", err)
	//	return 0, false, err, "Something Went Wrong"
	//}
	//dataAccess := dao.DbConn{DB: db}
	log.Println(tz)
	logger.Log.Println(tz)
	var wg sync.WaitGroup
	for _, transactionid := range tz.Transactionids {
		wg.Add(1)
		go Bulkapprovalsingleticket(transactionid, tz, &wg)
	}
	wg.Wait()
	return 0, true, nil, ""
}
func Bulkapprovalsingleticket(ticketid int64, tz *entities.Workflowentity, wg *sync.WaitGroup) {
	log.Print("\n\n Bulkapprovalsingleticket Transaction Id :", ticketid)
	logger.Log.Print("\n\n Bulkapprovalsingleticket Transaction Id :", ticketid)
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		//return 0, false, err, "Something Went Wrong"
	}
	defer wg.Done()
	dataAccess := dao.DbConn{DB: db}
	catdetails, err := dataAccess.Getcategorybyticketid(ticketid)
	if err != nil {
		return
		//return 0, false, err, "Something Went Wrong"
	}
	if len(catdetails) > 0 {
		tz1 := entities.Workflowentity{}
		wue := entities.WorkflowUtilityEntity{}
		wue.Clientid = tz.Clientid
		wue.Mstorgnhirarchyid = catdetails[0].Mstorgnhirarchyid
		wue.Typeseqno = config.STATUS_SEQ
		wue.Seqno = tz.Nextstateseq
		currentstate, err := dataAccess.Getstatebyseq(&wue)
		if err != nil {
			return
		}
		if len(currentstate) > 0 {
			tz1.Currentstateid = currentstate[0].Mststateid
		} else {
			return
		}
		wue.Seqno=config.PENDING_APPROVAL_STATUS_SEQ
		prevstate, err := dataAccess.Getstatebyseq(&wue)
		if err != nil {
			return
		}
		if len(prevstate) > 0 {
			tz1.Previousstateid = prevstate[0].Mststateid
		} else {
			return
		}

		tz1.Clientid = tz.Clientid
		tz1.Mstorgnhirarchyid = catdetails[0].Mstorgnhirarchyid
		tz1.Recorddiffid = catdetails[0].Recorddiffid
		tz1.Recorddifftypeid = catdetails[0].Recorddifftypeid
		tz1.Transactionid = ticketid
		//tz1.Previousstateid = tz.Previousstateid
		//tz1.Currentstateid = tz.Currentstateid
		tz1.Manualstateselection = tz.Manualstateselection
		tz1.Mstgroupid = tz.Mstgroupid
		tz1.Mstuserid = tz.Mstuserid
		tz1.Userid = tz.Userid
		tz1.Createdgroupid = tz.Createdgroupid
		//tz1.Transitionid = tz.Transitionid
		log.Print("\n\nTransaction Id sending:", tz1.Transactionid)
		logger.Log.Print("\n\nTransaction Id sending:", tz1.Transactionid)
		_, success, _, msg := MoveWorkflow(&tz1, db)
		log.Println("After Workflow:", success, msg)
		logger.Log.Print("After Workflow:", success, msg)
	}
}
func Checkworkflow(tz *entities.Workflowentity) (int64, bool, error, string) {
	logger.Log.Println("\n\nIn side Checkworkflow::")
	log.Println("\n\nIn side Checkworkflow::")
	logger.Log.Println(tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Previousstateid,
		tz.Manualstateselection, tz.Transactionid, tz.Mstgroupid, tz.Mstuserid)
	log.Println(tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Previousstateid,
		tz.Manualstateselection, tz.Transactionid, tz.Mstgroupid, tz.Mstuserid)
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	/**
	Get Process by using record category value
	*/
	processValue, err := dataAccess.GetProcessByCategory(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if processValue.Processid == 0 {
		return 0, false, nil, "No Process is mapped with category"
	} else {
		tz.Processid = processValue.Processid
		/**
		This process checks whether the process is defined and if defined,
		whether it is completed or not
		*/
		processDetails, err := dataAccess.Checkprocesscomplete(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		if len(processDetails) == 0 {
			return 0, false, err, "Process is not defined yet."
		} else {
			if processDetails[0].Iscomplete == 0 {
				return 0, false, err, "Process is  defined but not completed yet.Please complete it first."
			} else {
				return 1, true, err, ""
			}
		}
	}
}
func Checkworkflowstate(tz *entities.Workflowentity) (int64, bool, error, string) {
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()

	if err != nil {
		log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	//defer db.Close()
	dataAccess := dao.DbConn{DB: db}

	/**
	Get Process by using record category value
	*/
	processValue, err := dataAccess.GetProcessByCategory(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if processValue.Processid == 0 {
		return 0, false, nil, "No Process is mapped with category"
	} else {
		tz.Processid = processValue.Processid
		transitionState, err := dataAccess.GetTransitionState(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		log.Print("\n\nTransition state:", transitionState)
		if len(transitionState) > 0 {
			return 0, true, nil, ""
		} else {
			return 0, false, err, "No more State is defined in Workflow"
		}
	}
}
func MoveWorkflowwithapi(tz *entities.Workflowentity) (int64, bool, error, string) {
	logger.Log.Println("\n\nIn side MoveWorkflowwithapi::")
	log.Println("\n\nIn side MoveWorkflowwithapi::")
	logger.Log.Println(tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Previousstateid, tz.Currentstateid,
		tz.Manualstateselection, tz.Transactionid, tz.Mstgroupid, tz.Mstuserid, tz.Userid, tz.Createdgroupid, tz.Changestatus, tz.Transitionid)
	log.Println(tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Previousstateid, tz.Currentstateid,
		tz.Manualstateselection, tz.Transactionid, tz.Mstgroupid, tz.Mstuserid, tz.Userid, tz.Createdgroupid, tz.Changestatus, tz.Transitionid)
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	if tz.Previousstateid > -1 {
		dataAccess := dao.DbConn{DB: db}
		err, requestIds := dataAccess.GetRequestIdbyRecordId(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		var id int64
		if len(requestIds) > 0 {
			id = requestIds[0].Requestid
		}
		checkCurrentState, err := dataAccess.CheckCurrentState(id)
		if err != nil {
			logger.Log.Println("Current State Check error", err)
			log.Println("Current State Check error", err)
			return 0, false, err, "Current State Check error."
		}
		if len(checkCurrentState) > 0 && checkCurrentState[0].Currentstateid != tz.Previousstateid {
			return 0, false, err, "Current State not matched."
		}
	}
	_, success, err, msg := MoveWorkflow(tz, db)
	return 0, success, err, msg
}
func MoveWorkflow(tz *entities.Workflowentity, db *sql.DB) (int64, bool, error, string) {
	logger.Log.Println("\n\nIn side MoveWorkflow::")
	log.Println("\n\nIn side MoveWorkflow::")
	logger.Log.Println(tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Previousstateid, tz.Currentstateid,
		tz.Manualstateselection, tz.Transactionid, tz.Mstgroupid, tz.Mstuserid, tz.Userid, tz.Createdgroupid, tz.Changestatus, tz.Transitionid)
	log.Println(tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Previousstateid, tz.Currentstateid,
		tz.Manualstateselection, tz.Transactionid, tz.Mstgroupid, tz.Mstuserid, tz.Userid, tz.Createdgroupid, tz.Changestatus, tz.Transitionid)

	dataAccess := dao.DbConn{DB: db}
	/**
	Get Process by using record category value
	*/
	processValue, err := dataAccess.GetProcessByCategory(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	//log.Print("\n\nTransaction Id 111 :", tz.Transactionid)
	if processValue.Processid == 0 {
		return 0, false, nil, "No Process is mapped with category"
	} else {
		//logger.Log.Print("\n\nTransaction Id 11 :", tz.Transactionid)
		//log.Print("\n\nTransaction Id 11 :", tz.Transactionid)
		tz.Processid = processValue.Processid
		/**
		This process checks whether the process is defined and if defined,
		whether it is completed or not
		*/
		processDetails, err := dataAccess.Checkprocesscomplete(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}

		if len(processDetails) == 0 {
			return 0, false, err, "Process is not defined yet."
		} else {
			if processDetails[0].Iscomplete == 0 {
				return 0, false, err, "Process is  defined but not completed yet.Please complete it first."
			} else {
				//logger.Log.Print("\n\nTransaction Id 22 :", tz.Transactionid)
				/**
				This method is used to getting the
				record table name by using the process id
				*/
				tableName, err := dataAccess.GetTableByProcess(tz)
				if err != nil {
					return 0, false, err, "Something Went Wrong"
				}
				//logger.Log.Print("\n\nTransaction Id 33 :", tz.Transactionid)
				if tableName.Tablename == "" {
					return 0, false, nil, "Workflow table name not found"
				}
				/**
				If current state id is not given i.e straight workflow is defined,
				we are fetching current state id from msttransition table using previous state id
				*/
				if tz.Currentstateid == 0 && tz.Transitionid == 0 {
					currentstate, err := dataAccess.Getcurrentstateid(tz)
					if err != nil {
						return 0, false, err, "Something Went Wrong"
					}
					if len(currentstate) == 0 {
						return 0, false, nil, "Workflow not defined properly"
					} else {
						tz.Currentstateid = currentstate[0].Currentstateid
					}
				}
				//logger.Log.Print("\n\nTransaction Id 44 :", tz.Transactionid)
				//log.Print("\n\nTransaction Id 44 :", tz.Transactionid)
				/**
				Fetching record details from the table (getting by the 'GetTableByProcess()')
				and TransactionId
				*/
				recordDetails, trerr := dataAccess.GetRecordDetailsById(tz, tableName.Tablename)
				if trerr != nil {
					return 0, false, trerr, "Something Went Wrong"
				}
				if len(recordDetails) > 0 {
					//logger.Log.Print("\n\nTransaction Id 55 :", tz.Transactionid)
					err, requestIds := dataAccess.GetRequestIdbyRecordId(tz)
					if err != nil {
						return 0, false, err, "Something Went Wrong"
					}
					var id int64
					var isFirstStep bool
					if len(requestIds) > 0 {
						id = requestIds[0].Requestid
						isFirstStep = false
					} else {
						isFirstStep = true
					}
					logger.Log.Println("Before transition id ", tz.Transitionid)
					log.Println("Before transition id ", tz.Transitionid)
					//logger.Log.Print("\n\nTransaction Id 66 :", tz.Transactionid)
					tz.Createduserid = recordDetails[0].Userid
					if tz.Manualstateselection == 0 {
						if tz.Transitionid == 0 {
							/**
							Using the process id and previousstate id , fetching the latest transition id
							according with the currentstate id of the record.
							*/
							transitionState, err := dataAccess.GetTransitionState(tz)
							if err != nil {
								return 0, false, err, "Something Went Wrong"
							}
							log.Print("\n\nTransition state if:", transitionState)
							logger.Log.Print("\n\nTransition state if:", transitionState)
							if len(transitionState) == 0 {
								return 0, false, err, "No more State is defined in Workflow"
							}
							tz.Transitionid = transitionState[0].Transitionid
						} else {
							/**
							Using the transition id , fetching the previous state id and
							 current state id of the record.
							*/
							transitionState, err := dataAccess.Getstatebytranid(tz.Transitionid)
							if err != nil {
								return 0, false, err, "Something Went Wrong"
							}
							log.Print("\n\nTransition state:", transitionState)
							logger.Log.Print("\n\nTransition state:", transitionState)
							if len(transitionState) == 0 {
								return 0, false, err, "Next State is not defined in Workflow"
							}
							tz.Currentstateid = transitionState[0].Currentstateid
							tz.Previousstateid = transitionState[0].Previousstateid
						}
						log.Print("\n\nTransition Id:", tz.Transitionid)
						logger.Log.Print("\n\nTransition Id:", tz.Transitionid)
						//logger.Log.Print("\n\nTransaction Id 77 :", tz.Transactionid)
						/**
						Check whether there is any delegated user in this transition state.
						If any then we'll use this delegated user as assigned user for this state
						*/
						delegateUser, err := dataAccess.GetDelegateUser(tz)
						if err != nil {
							return 0, false, err, "Something Went Wrong"
						}
						if len(delegateUser) > 0 {
							tz.Mstgroupid = delegateUser[0].Mstgroupid
							tz.Mstuserid = delegateUser[0].Mstuserid
						} else {
							/**
							If no delegated user is mapped for this state,
							We'll use the user and group that are mapped in the workflow state.
							*/
							stateuser, err := dataAccess.GetStateUserByTransitionId(tz)
							if err != nil {
								return 0, false, err, "Something Went Wrong"
							}
							if len(stateuser) > 0 {
								if stateuser[0].Mstgroupid == 0 && stateuser[0].Mstuserid == -1 {
									log.Print("Go TO CREATOR: ", recordDetails[0].Groupid, recordDetails[0].Userid)
									logger.Log.Print("Go TO CREATOR: ", recordDetails[0].Groupid, recordDetails[0].Userid)
									/**
									  Assigning the userid and groupid from record table for `Go to creator` option
									*/
									tz.Mstgroupid = recordDetails[0].Groupid
									tz.Mstuserid = recordDetails[0].Userid
								} else if stateuser[0].Mstgroupid == 0 && stateuser[0].Mstuserid == 0 {
									/**
									fetching userdetails for `Self Assign` option.
									userid and groupid coming from api(sender userid and groupid).So no need to fetch
									*/
									logger.Log.Print("SELF ASSIGN: ", tz.Mstgroupid, tz.Mstuserid)
									log.Print("SELF ASSIGN: ", tz.Mstgroupid, tz.Mstuserid)
								} else if stateuser[0].Mstgroupid == 0 && stateuser[0].Mstuserid == -2 {
									/**
									fetching details for 'Back to sender (User)' option
									fetching previous state user details(group and user both) from `mstrequesthistory` table
									*/
									prevuser, err := dataAccess.Getprevioussenderdetails(id)
									if err != nil {
										return 0, false, err, "Something Went Wrong"
									}
									logger.Log.Print("Go TO Prev: ", len(prevuser))
									log.Print("Go TO Prev: ", len(prevuser))
									if len(prevuser) > 1 {
										var currentstate = prevuser[0].Currentstateid
										for i := 1; i < len(prevuser); i++ {
											if prevuser[i].Currentstateid != currentstate {
												tz.Mstgroupid = prevuser[i].Mstgroupid
												tz.Mstuserid = prevuser[i].Mstuserid
												break
											}
										}
										//tz.Mstgroupid = prevuser[1].Mstgroupid
										//tz.Mstuserid = prevuser[1].Mstuserid
										logger.Log.Print("Go TO Prev details user: ", tz.Mstgroupid, tz.Mstuserid)
										log.Print("Go TO Prev details user: ", tz.Mstgroupid, tz.Mstuserid)
									} else {
										return 0, false, nil, "No previous state details found."
									}
								} else if stateuser[0].Mstgroupid == 0 && stateuser[0].Mstuserid == -3 {
									/**
									fetching details for 'Back to sender (Group)' option
									fetching previous state user details(group only) from `mstrequesthistory` table
									*/
									prevuser, err := dataAccess.Getprevioussenderdetails(id)
									if err != nil {
										return 0, false, err, "Something Went Wrong"
									}
									logger.Log.Print("Go TO Prev: ", len(prevuser))
									log.Print("Go TO Prev: ", len(prevuser))
									if len(prevuser) > 1 {
										var currentstate = prevuser[0].Currentstateid
										logger.Log.Println(" Current state : ", currentstate)
										log.Println(" Current state : ", currentstate)
										for i := 1; i < len(prevuser); i++ {
											if prevuser[i].Currentstateid != currentstate {
												logger.Log.Println(" Previous state : ", prevuser[i].Currentstateid)
												log.Println(" Previous state : ", prevuser[i].Currentstateid)
												tz.Mstgroupid = prevuser[i].Mstgroupid
												break
											}
										}
										tz.Mstuserid = 0
										logger.Log.Print("Go TO Prev details group: ", tz.Mstgroupid)
										log.Print("Go TO Prev details group: ", tz.Mstgroupid)
									} else {
										return 0, false, nil, "No previous state details found."
									}
								} else if stateuser[0].Mstgroupid == 0 && stateuser[0].Mstuserid == -4 {

									/**
									fetching details for 'Back to Manager' option
									fetching created user's manager id from `mstclientuser` table
									*/

									relmanagers, err := dataAccess.Getuserrelmanager(tz.Createduserid)
									if err != nil {
										return 0, false, err, "Something Went Wrong"
									}
									logger.Log.Print("Go TO Rel Manager: ", len(relmanagers))
									log.Print("Go TO Rel Manager: ", len(relmanagers))
									if len(relmanagers) > 0 {
										tz.Mstgroupid = relmanagers[0].Mstgroupid
										tz.Mstuserid = relmanagers[0].Mstuserid
										logger.Log.Print("Go TO Rel Manager: ", tz.Mstgroupid, tz.Mstuserid)
										log.Print("Go TO Rel Manager: ", tz.Mstgroupid, tz.Mstuserid)
									} else {
										return 0, false, nil, "No Rel Manager mapped for the creator."
									}
								} else {
									/**
										Assigning the first user and groupid from `maprecorddifferentiongroup`
									table for `Manual selection`
									*/
									tz.Mstgroupid = stateuser[0].Mstgroupid
									tz.Mstuserid = stateuser[0].Mstuserid
									logger.Log.Print("Manual Selection: ", tz.Mstgroupid, tz.Mstuserid)
									log.Print("Manual Selection: ", tz.Mstgroupid, tz.Mstuserid)
								}
							} else {
								return 0, false, nil, "No Group / User is mapped with State"
							}
						}
					} else {
						/**
						This is a manual selection.Setting the transitionid to -1
						Also fetching the user details against the currentstateid
						*/
						tz.Transitionid = -1
					}
					/**
					Fetching the latest id of record history table by Transaction Id
					*/
					stageDetails, trerr := dataAccess.GetLatestTransactionStageDetails(tz)
					if trerr != nil {
						return 0, false, trerr, "Something Went Wrong"
					}
					if len(stageDetails) > 0 {
						recordDetails[0].Recordstageid = stageDetails[0].Recordstageid
						/**
						Actual Workflow moving is done here
						1. Storing the latest record state in 'mstrequest' table
						2. Mapping the request state,transaction record details and transaction history
						in 'maprequestorecord'
						3. Workflow moving history is stored in mstrequesthistory.
						*/

						tx, err := db.Begin()
						if err != nil {
							logger.Log.Println("Transaction creation error.", err)
							log.Println("Transaction creation error.", err)
							return 0, false, err, "Something Went Wrong"
						}
						logger.Log.Println("Before sending ", tz.Previousstateid, tz.Currentstateid)
						log.Println("Before sending ", tz.Previousstateid, tz.Currentstateid)
						recordVal, recerr := dao.UpsertProcessDetails(tx, tz, recordDetails[0], isFirstStep, id, "N")
						if recerr != nil {
							logger.Log.Println("Role back error.")
							log.Println("Role back error.")
							tx.Rollback()
							return 0, false, recerr, "Something Went Wrong"
						}
						err = tx.Commit()
						if err != nil {
							log.Print("MoveWorkflow  Statement Commit error", err)
							logger.Log.Print("MoveWorkflow  Statement Commit error", err)
							return 0, false, err, ""
						}
						status, serr := dataAccess.Getstatusbystateid(tz.Clientid, tz.Mstorgnhirarchyid, tz.Currentstateid)
						if serr != nil {
							return 0, false, serr, "Something Went Wrong"
						}
						var updaterequired = true
						if len(status) > 0 {
							if status[0].Seqno == config.CLOSE_SEQ || status[0].Seqno == config.RESOLVE_SEQ || status[0].Seqno == config.PENDING_USER_STATUS_SEQ || status[0].Seqno == config.PENDING_REQUESTER_INPUT || status[0].Seqno == config.CANCEL_SEQ || status[0].Seqno == config.REJECTED_STATUS_SEQ {
								updaterequired = false
							}
						}
						if updaterequired {

							count, success, _, msg := CalculateHopCount(db, tz.Clientid, tz.Transactionid, recordDetails[0].Groupid)
							if !success {
								logger.Log.Print("\n Error in hop count:: ", msg)
								log.Print("\n Error in hop count:: ", msg)
							}
							loggedusername := ""
							uerr, users := dataAccess.Getusername(tz.Userid)
							if uerr != nil {
								return 0, false, uerr, "Something Went Wrong"
							}
							if len(users) > 0 {
								loggedusername = users[0].Username
							}
							var grpname string
							var username string
							var name string
							grperr, group := dataAccess.Getgroupname(tz.Mstgroupid)
							if grperr != nil {
								return 0, false, grperr, "Something Went Wrong"
							}
							if len(group) > 0 {
								grpname = group[0].Groupname
							}
							grperr, user := dataAccess.Getusername(tz.Mstuserid)
							if grperr != nil {
								return 0, false, grperr, "Something Went Wrong"
							}
							if len(user) > 0 {
								username = user[0].Loginname
								name = user[0].Username
							}

							stagEntity := entities.StagingUtilityEntity{}
							stagEntity.Clientid = tz.Clientid
							stagEntity.Mstorgnhirarchyid = tz.Mstorgnhirarchyid
							stagEntity.Assignedgroupid = tz.Mstgroupid
							stagEntity.Assignedgroupid = tz.Mstgroupid
							stagEntity.Assignedgroup = grpname
							stagEntity.Assigneduser = name
							stagEntity.Assignedloginname = username
							stagEntity.Assigneduserid = tz.Mstuserid
							stagEntity.Lastuser = loggedusername
							stagEntity.Lastuserid = tz.Userid
							stagEntity.Reassigncount = count
							stagEntity.Recordid = tz.Transactionid
							stgerr := dataAccess.Updatestagingdetailswithouttran(&stagEntity)
							if stgerr != nil {
								return 0, false, stgerr, "Something Went Wrong"
							}
						}
						//err = tx.Commit()
						//if err != nil {
						//	log.Print("MoveWorkflow  Statement Commit error", err)
						//	logger.Log.Print("MoveWorkflow  Statement Commit error", err)
						//	return 0, false, err, ""
						//}
						postBody, _ := json.Marshal(map[string]int64{"clientid": tz.Clientid, "mstorgnhirarchyid": tz.Mstorgnhirarchyid, "recordid": tz.Transactionid, "reordstatusid": tz.Currentstateid, "userid": tz.Userid, "usergroupid": tz.Createdgroupid, "changestatus": tz.Changestatus, "issrrequestor": tz.Issrrequestor})

						responseBody := bytes.NewBuffer(postBody)
						logger.Log.Println("changestatus,", tz.Changestatus, tz.Transactionid)
						log.Println("changestatus,", tz.Changestatus, tz.Transactionid)
						logger.Log.Println("postBody       --->", responseBody)
						log.Println("postBody       --->", responseBody)
						resp, err := http.Post(config.RECORD_URL+"/updaterecordstatus", "application/json", responseBody)
						if err != nil {
							logger.Log.Println("An Error Occured --->", err)
							log.Println("An Error Occured --->", err)
							return 0, false, err, "Something went wrong"
						}
						defer resp.Body.Close()
						//Read the response body
						body, err := ioutil.ReadAll(resp.Body)
						if err != nil {
							logger.Log.Println("response body ------> ", err)
							log.Println("response body ------> ", err)
							return 0, false, err, "Something went wrong"
						}
						sb := string(body)
						logger.Log.Println("sb body value is --->", sb)
						log.Println("sb body value is --->", sb)
						return recordVal, true, nil, ""
						//return 0, true, nil, ""
					} else {
						return 0, false, trerr, "No Staging Record is mapped for this record id."
					}

				} else {
					return 0, false, trerr, "No Record is mapped for this record id."
				}
			}
		}
	}
}
func Insertprocess(tz *entities.Workflowentity) (int64, bool, error, string) {
	log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	//defer db.Close()
	savemst := dao.DbConn{DB: db}
	values, err1 := savemst.Getprocessdetails(tz)
	if err1 != nil {
		return 0, false, err1, "Something Went Wrong"
	}
	if len(values) == 0 {
		id, err := savemst.Insertprocess(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		return id, true, err, ""
	} else {
		tz.Id = values[0].Id
		err = savemst.Updateprocessdetails(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		return 0, true, err, ""
	}
}
func Getprocessdetails(tz *entities.Workflowentity) ([]entities.WorkflowResponseEntity, bool, error, string) {
	log.Println("In side model")
	t := []entities.WorkflowResponseEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()

	if err != nil {
		logger.Log.Println("database connection failure", err)
		log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	//defer db.Close()
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Getprocessdetails(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}
func Gettransitionstatedetails(tz *entities.Workflowentity) (entities.WorkflowStateResponseEntity, bool, error, string) {
	log.Println("In side model")
	t := entities.WorkflowStateResponseEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	//db, err := config.ConnectMySqlDb()
	//defer db.Close()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	if len(tz.Transitionids) > 0 {
		tz.Transitionid = tz.Transitionids[0]
		details, err2 := dataAccess.Gettransitionstatedetails(tz)
		if err2 != nil {
			return t, false, err2, "Something Went Wrong"
		}
		/*activities, err3 := dataAccess.Getactivitybytransition(tz)
		if err3 != nil {
			return t, false, err3, "Something Went Wrong"
		}
		for _, activity := range activities {
			t.Activityids = append(t.Activityids, activity.Id)
		}*/
		if len(details) > 0 && details[0].Mstuserid > -1 {
			groupdetails, err2 := dataAccess.Gettransitiongroup(tz)
			if err2 != nil {
				return t, false, err2, "Something Went Wrong"
			}
			t.Groups = groupdetails
			return t, true, nil, ""
		} else {
			t.Groups = details
			return t, true, nil, ""
		}
	} else {
		return t, false, nil, "Problem with this state.Please remove and recreate this state"
	}
}
func Gettransitiongroupdetails(tz *entities.Workflowentity) ([]entities.WorkflowResponseEntity, bool, error, string) {
	log.Println("In side model")
	t := []entities.WorkflowResponseEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	//db, err := config.ConnectMySqlDb()
	//defer db.Close()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	groupdetails, err2 := dataAccess.Gettransitiongroup(tz)
	if err2 != nil {
		return t, false, err2, "Something Went Wrong"
	}

	return groupdetails, true, nil, ""
}
func Checkprocessdelete(tz *entities.Workflowentity) ([]entities.WorkflowResponseEntity, bool, error, string) {
	log.Println("In side model")
	t := []entities.WorkflowResponseEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	//db, err := config.ConnectMySqlDb()
	//defer db.Close()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Checkprocessdelete(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if len(values) > 0 {
		return values, true, nil, ""
	} else {
		return values, true, nil, ""
	}
}
func Createtransition(tw *entities.Workflowentity) (int64, bool, error, string) {
	log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	dbcon, err := config.ConnectMySqlDbSingleton()
	//dbcon, err := config.ConnectMySqlDb()
	if err != nil {
		//dbcon.Close()
		logger.Log.Println("Database connection failure", err)
		log.Println("Database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	//defer dbcon.Close()
	tx, err := dbcon.Begin()
	if err != nil {
		logger.Log.Println("Transaction creation error.", err)
		log.Println("Transaction creation error.", err)
		return 0, false, err, "Something Went Wrong"
	}

	dataAccess := dao.DbConn{DB: dbcon}
	//states, err := dataAccess.Checkduplicatestate(tw)
	//if len(states) == 0 {
	workentity := entities.MapprocesstemplateEntity{}
	workentity.Clientid = tw.Clientid
	workentity.Mstorgnhirarchyid = tw.Mstorgnhirarchyid

	prevstateseq, err := dataAccess.Getseqbystateid(tw.Previousstateid)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	prevs, err := dataAccess.Getstateidbyseq(&workentity, prevstateseq)
	if err != nil {
		tx.Rollback()
		return 0, false, err, "Something Went Wrong"
	}
	if len(prevs) > 0 {
		tw.Previousstateid = prevs[0].Id
	}
	currstateseq, err := dataAccess.Getseqbystateid(tw.Currentstateid)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	currs, err := dataAccess.Getstateidbyseq(&workentity, currstateseq)
	if err != nil {
		tx.Rollback()
		return 0, false, err, "Something Went Wrong"
	}
	if len(currs) > 0 {
		tw.Currentstateid = currs[0].Id
	}
	id, err := dao.Createtransition(tw, tx)
	if err != nil {
		tx.Rollback()
		return 0, false, err, "Something Went Wrong"
	}
	if len(tw.Transitionids) > 0 {
		tw.Transitionid = tw.Transitionids[0]
		details, err2 := dataAccess.Getalltransitionstatedetails(tw)
		if err2 != nil {
			return 0, false, err2, "Something Went Wrong"
		}
		if len(details) > 0 {
			for _, val := range details {
				tw.Recorddifftypeid = val.Recorddifftypeid
				tw.Recorddiffid = val.Recorddiffid
				tw.Mstgroupid = val.Mstgroupid
				tw.Mstuserid = val.Mstuserid
				tw.Transitionid = id
				_, err := dao.Inserttransitiondetails(tw, tx)
				if err != nil {
					tx.Rollback()
					return 0, false, err, "Something Went Wrong"
				}
			}
			err = tx.Commit()
			if err != nil {
				log.Print("Createtransition  Statement Commit error", err)
				logger.Log.Print("Createtransition  Statement Commit error", err)
				return 0, false, err, ""
			}

			return id, true, nil, ""
		} else {
			err = tx.Commit()
			if err != nil {
				log.Print("Createtransition  Statement Commit error", err)
				logger.Log.Print("Createtransition  Statement Commit error", err)
				return 0, false, err, ""
			}
			return id, true, nil, ""
		}
	} else {
		err = tx.Commit()
		if err != nil {
			log.Print("Createtransition  Statement Commit error", err)
			logger.Log.Print("Createtransition  Statement Commit error", err)
			return 0, false, err, ""
		}
		return id, true, nil, ""
	}
	//} else {
	//	return 0, false, nil, "Transition path already exist"
	//}
}
func Upserttransitiondetails(tw *entities.Workflowentity) (int64, bool, error, string) {
	log.Println("In side model")
	//dbcon, err := config.ConnectMySqlDb()
	lock.Lock()
	defer lock.Unlock()
	dbcon, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		//dbcon.Close()
		logger.Log.Println("Database connection failure", err)
		log.Println("Database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	//defer dbcon.Close()
	tx, err := dbcon.Begin()
	if err != nil {
		logger.Log.Println("Transaction creation error.", err)
		return 0, false, err, "Something Went Wrong"
	}
	var ids string = ""
	for i, transition := range tw.Transitionids {
		if i > 0 {
			ids += ","
		}
		ids += strconv.Itoa(int(transition))
	}
	err = dao.Deletetransitiondetails(tw, tx, ids)
	if err != nil {
		tx.Rollback()
		return 0, false, err, "Something Went Wrong"
	}
	/*err = dao.Deleteactivitydetails(tw, tx, ids)
	if err != nil {
		tx.Rollback()
		return 0, false, err, "Something Went Wrong"
	}*/
	for _, tid := range tw.Transitionids {
		tw.Transitionid = tid
		for _, user := range tw.Users {
			tw.Mstuserid = user.Mstuserid
			tw.Mstgroupid = user.Mstgroupid
			_, err := dao.Inserttransitiondetails(tw, tx)
			if err != nil {
				tx.Rollback()
				return 0, false, err, "Something Went Wrong"
			}
		}
		/*for _, activity := range tw.Activities {
			tw.Activity = activity
			_, err := dao.Insertstateactivity(tw, tx)
			if err != nil {
				tx.Rollback()
				return 0, false, err, "Something Went Wrong"
			}
		}*/
	}
	err = tx.Commit()
	if err != nil {
		log.Print("Upserttransitiondetails  Statement Commit error", err)
		logger.Log.Print("Upserttransitiondetails  Statement Commit error", err)
		return 0, false, err, ""
	}
	return 0, true, nil, ""
}
func Deletetransitionstate(tw *entities.Workflowentity) (int64, bool, error, string) {
	//dbcon, err := config.ConnectMySqlDb()
	lock.Lock()
	defer lock.Unlock()
	dbcon, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		//dbcon.Close()
		logger.Log.Println("Database connection failure", err)
		log.Println("Database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	//defer dbcon.Close()
	tx, err := dbcon.Begin()
	if err != nil {
		logger.Log.Println("Transaction creation error.", err)
		return 0, false, err, "Something Went Wrong"
	}
	var ids string = ""
	for i, transition := range tw.Transitionids {
		if i > 0 {
			ids += ","
		}
		ids += strconv.Itoa(int(transition))
	}
	err = dao.Deletetransitiondetails(tw, tx, ids)
	if err != nil {
		tx.Rollback()
		return 0, false, err, "Something Went Wrong"
	}
	err = dao.Deleteactivitydetails(tw, tx, ids)
	if err != nil {
		tx.Rollback()
		return 0, false, err, "Something Went Wrong"
	}
	err = dao.Deletetransition(tw, tx, ids)
	if err != nil {
		tx.Rollback()
		return 0, false, err, "Something Went Wrong"
	}
	err = tx.Commit()
	if err != nil {
		log.Print("Deletetransitionstate  Statement Commit error", err)
		logger.Log.Print("Deletetransitionstate  Statement Commit error", err)
		return 0, false, err, ""
	}
	return 0, true, nil, ""
}
func Getstatedetails(tz *entities.Workflowentity) ([]entities.TransactionRespEntity, bool, error, string) {
	log.Println("In side model")
	t := []entities.TransactionRespEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	//db, err := config.ConnectMySqlDb()
	//defer db.Close()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Getstatedetails(tz)
	if err1 != nil {
		return nil, false, err1, "Something Went Wrong"
	}
	logger.Log.Println(values)
	log.Println(values)
	if len(values) > 0 {
		parents, err1 := dataAccess.Getinparentdetails(tz)
		if err1 != nil {
			return nil, false, err1, "Something Went Wrong"
		}
		if len(parents) > 0 {
			tz.Transactionid = parents[0].Id
		} else {
			tz.Transactionid = tz.Recordid
		}

		_, lastgroupname, lastusername, lastgroupid, _, _, err := Getlastactionerdetails(tz.Transactionid, values[0].Groupid, db)
		if err != nil {
			return nil, false, err, "Something Went Wrong"
		}
		values[0].Lastgroupname = lastgroupname
		values[0].Lastusername = lastusername
		values[0].Lastgroupid = lastgroupid
		return values, true, err, ""
	} else {
		return t, false, nil, "State details not found."
	}
}
func Getlastactionerdetails(ticketid int64, groupid int64, db *sql.DB) (string, string, string, int64, int64, bool, error) {
	dataAccess := dao.DbConn{DB: db}
	tz := entities.Workflowentity{}
	tz.Transactionid = ticketid
	err, requestIds := dataAccess.GetRequestIdbyRecordId(&tz)
	if err != nil {
		return "", "", "", 0, 0, false, err
	}
	var requestid int64
	if len(requestIds) > 0 {
		requestid = requestIds[0].Requestid
	}
	logger.Log.Println("----Params-----> ", requestid, groupid)
	log.Println("----Params-----> ", requestid, groupid)
	details, err := dataAccess.Getlastactionerwithgroup(requestid)
	if err != nil {
		return "", "", "", 0, 0, false, err
	}
	lastgroupname := ""
	lastusername := ""
	lastuserloginname := ""
	var lastgroupid int64 = 0
	var lastuserid int64 = 0
	logger.Log.Println("details : ", details, len(details))
	log.Println("details : ", details, len(details))
	if len(details) > 1 {
		lastgroupname = details[1].Lastgroupname
		lastusername = details[1].Lastusername
		lastgroupid = details[1].Lastgroupid
		lastuserloginname = details[1].Lastuserloginname
		lastuserid = details[1].Userid
	} /*else {
		details1, err := dataAccess.Getlastactioneruser(requestid)
		if err != nil {
			return "", "", "", 0, 0, false, err
		}
		if len(details1) > 1 {
			logger.Log.Println("details1 : ", details1, len(details1))
			log.Println("details1 : ", details1, len(details1))
			lastgroupname = details1[1].Lastgroupname
			lastusername = details1[1].Lastusername
			lastgroupid = details1[1].Lastgroupid
			lastuserloginname = details1[1].Lastuserloginname
			lastuserid = details1[1].Userid

		}
	}*/
	logger.Log.Println("------> ", lastuserloginname, lastgroupname, lastusername, lastgroupid, lastuserid)
	log.Println("------> ", lastuserloginname, lastgroupname, lastusername, lastgroupid, lastuserid)
	return lastuserloginname, lastgroupname, lastusername, lastgroupid, lastuserid, true, nil
}
func Getnextstatedetails(tz *entities.Workflowentity) ([]entities.TransactionRespEntity, bool, error, string) {
	log.Println("In side model")
	t := []entities.TransactionRespEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	//db, err := config.ConnectMySqlDb()
	//defer db.Close()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}

	/*err, requestIds := dataAccess.GetRequestIdbyRecordId(tz)
	if err != nil {
		return t, false, err, "Something Went Wrong"
	}
	var requestid int64
	if len(requestIds) > 0 {
		requestid = requestIds[0].Requestid
	}
	ismanual, err := dataAccess.Ismanualselection(requestid)
	if err != nil {
		return t, false, err, "Something Went Wrong"
	}
	if !ismanual {*/
	details, err1 := dataAccess.Getprocessdetails(tz)
	if err1 != nil {
		return nil, false, err1, "Something Went Wrong"
	}
	if len(details) == 0 {
		return nil, false, nil, "Template details not defined"
	}
	log.Println(details[0].Detailsjson)
	in := []byte(details[0].Detailsjson)
	var detailsEntity []entities.ProcessdetailsEntity
	err = json.Unmarshal(in, &detailsEntity)
	if err != nil {
		logger.Log.Println(err)
		log.Print(err)
		return nil, false, err1, "Something Went Wrong"
	}
	var outstates []int64
	matched := false
	for i, ent := range detailsEntity {
		for _, in := range ent.Instate {
			if in == tz.Transitionid {
				log.Print(detailsEntity[i].OutState)
				outstates = detailsEntity[i].OutState
				matched = true
				break
			}
		}
		if matched {
			break
		}
	}

	log.Println(outstates)
	logger.Log.Println("outstates ", outstates)
	log.Println("outstates ", outstates)
	if len(outstates) > 0 {
		var ids string = ""
		for i, state := range outstates {
			if i > 0 {
				ids += ","
			}
			ids += strconv.Itoa(int(state))
		}
		values, err1 := dataAccess.Getnextstatedetails(tz, ids)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		return values, true, err, ""
	} else {
		tz.Recordid = tz.Transactionid
		parents, err1 := dataAccess.Getinparentdetails(tz)
		if err1 != nil {
			return nil, false, err1, "Something Went Wrong"
		}
		if len(parents) == 0 {
			return nil, false, nil, "No more state defined."
		} else {
			return nil, false, nil, ""
		}

	}
	/*} else {
		return t, false, err, ""
	}*/
}
func Gettransitionbyprocess(tz *entities.Workflowentity) ([]entities.WorkflowTransitionEntity, bool, error, string) {
	log.Println("In side model")
	t := []entities.WorkflowTransitionEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	//db, err := config.ConnectMySqlDb()
	//defer db.Close()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Gettransitionbyprocess(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}
func Changerecordgroupwithapi(tw *entities.Workflowentity) (int64, bool, error, string) {
	log.Println("In side model")
	logger.Log.Println("Changerecordgroup:", tw.Transactionid, tw.Mstgroupid, tw.Mstuserid, tw.Createdgroupid, tw.Samegroup)
	log.Println("Changerecordgroup:", tw.Transactionid, tw.Mstgroupid, tw.Mstuserid, tw.Createdgroupid, tw.Samegroup)
	//dbcon, err := config.ConnectMySqlDb()
	lock.Lock()
	defer lock.Unlock()
	dbcon, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		//dbcon.Close()
		logger.Log.Println("Database connection failure", err)
		log.Println("Database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	_, success, err, msg := Changerecordgroup(tw, dbcon)
	return 0, success, err, msg
}
func Changerecordgroup(tw *entities.Workflowentity,dbcon *sql.DB) (int64, bool, error, string) {

	//defer dbcon.Close()
	dataAccess := dao.DbConn{DB: dbcon}
	err, requestIds := dataAccess.GetRequestIdbyRecordId(tw)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	var requestid int64
	if len(requestIds) > 0 {
		requestid = requestIds[0].Requestid
	}
	requestHistory, err := dataAccess.Fetchhistorybyrequestid(requestid)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	tw.Clientid = requestHistory[0].Clientid
	tw.Mstorgnhirarchyid = requestHistory[0].Mstorgnhirarchyid
	tw.Processid = requestHistory[0].Processid
	tw.Createduserid = requestHistory[0].Createduserid
	tw.Currentstateid = requestHistory[0].Currentstateid
	tw.Transitionid = requestHistory[0].Transitionid
	tw.Manualstateselection = requestHistory[0].Manualstateselection
	prevgroup := requestHistory[0].Groupname
	prevuser := requestHistory[0].Username
	grperr, group := dataAccess.Getgroupname(tw.Mstgroupid)
	if grperr != nil {
		return 0, false, grperr, "Something Went Wrong"
	}
	var grpname string
	var username string
	var name string
	var logstring string
	if len(group) > 0 {
		grpname = group[0].Groupname
	}
	if tw.Mstuserid > 0 {
		grperr, user := dataAccess.Getusername(tw.Mstuserid)
		if grperr != nil {
			return 0, false, grperr, "Something Went Wrong"
		}
		if len(user) > 0 {
			username = user[0].Loginname
			name = user[0].Username
		}
		if prevuser == "" {
			logstring = "From Group: " + prevgroup + " To Group: " + grpname + " User:" + name
		} else {
			logstring = "From Group: " + prevgroup + " User: " + prevuser + " To Group: " + grpname + " User:" + name
		}
	} else {
		if prevuser == "" {
			logstring = "From Group: " + prevgroup + " To Group: " + grpname
		} else {
			logstring = "From Group: " + prevgroup + " User: " + prevuser + " To Group: " + grpname
		}
	}
	tx, err := dbcon.Begin()
	if err != nil {
		logger.Log.Println("Transaction creation error.", err)
		log.Println("Transaction creation error.", err)
		return 0, false, err, "Something Went Wrong"
	}

	err = dao.Updaterequestgroup(tw, tx, requestid)
	if err != nil {
		tx.Rollback()
		return 0, false, err, "Something Went Wrong"
	}
	latestTime := time.Now().Unix()
	rec := entities.TransactionEntity{}
	histerr := dao.InsertProcessHistoryRequest(tw, tx, latestTime, rec, requestid, "N")
	if histerr != nil {
		tx.Rollback()
		return 0, false, histerr, "Something Went Wrong"
	}
	log.Print("---->messsage:", logstring)
	logger.Log.Print("---->messsage:", logstring)
	var activityseq int64
	var eventnotificationid int64
	if tw.Samegroup {
		activityseq = 12
		eventnotificationid = 10
	} else {
		activityseq = 2
		eventnotificationid = 11
	}
	err = utility.InsertActivityLogs(tx, tw.Clientid, tw.Mstorgnhirarchyid, tw.Transactionid, activityseq, logstring, tw.Userid, tw.Createdgroupid)
	if err != nil {
		tx.Rollback()
		return 0, false, err, "Something Went Wrong"
	}
	err = tx.Commit()
	if err != nil {
		logger.Log.Print("Deletetransitionstate  Statement Commit error", err)
		log.Print("Deletetransitionstate  Statement Commit error", err)
		return 0, false, err, ""
	}

	if tw.Creatorgroupid == 0 {
		recordDetails, trerr := dataAccess.GetRecordDetailsById(tw, "trnrecord")
		if trerr != nil {
			return 0, false, trerr, "Something Went Wrong"
		}
		if len(recordDetails) > 0 {
			tw.Creatorgroupid = recordDetails[0].Groupid
		}
	}
	count, success, _, msg := CalculateHopCount(dbcon, tw.Clientid, tw.Transactionid, tw.Creatorgroupid)
	if !success {
		logger.Log.Print("\n Error in hop count:: ", msg)
		log.Print("\n Error in hop count:: ", msg)
	}
	loggedusername := ""
	uerr, users := dataAccess.Getusername(tw.Userid)
	if uerr != nil {
		return 0, false, uerr, "Something Went Wrong"
	}
	if len(users) > 0 {
		loggedusername = users[0].Username
	}
	stagEntity := entities.StagingUtilityEntity{}
	stagEntity.Clientid = tw.Clientid
	stagEntity.Mstorgnhirarchyid = tw.Mstorgnhirarchyid
	stagEntity.Assignedgroupid = tw.Mstgroupid
	stagEntity.Assignedgroup = grpname
	stagEntity.Assigneduser = name
	stagEntity.Assignedloginname = username
	stagEntity.Assigneduserid = tw.Mstuserid
	stagEntity.Lastuser = loggedusername
	stagEntity.Lastuserid = tw.Userid
	stagEntity.Reassigncount = count
	stagEntity.Recordid = tw.Transactionid
	stgerr := dataAccess.Updatestagingdetailswithouttran(&stagEntity)
	if stgerr != nil {
		//tx.Rollback()
		return 0, false, stgerr, "Something Went Wrong"
	}
	//err = tx.Commit()
	//if err != nil {
	//	logger.Log.Print("Deletetransitionstate  Statement Commit error", err)
	//	log.Print("Deletetransitionstate  Statement Commit error", err)
	//	return 0, false, err, ""
	//}

	/**
	Sending mail for group or user change
	*/
	postBody, _ := json.Marshal(map[string]int64{"clientid": tw.Clientid, "mstorgnhirarchyid": tw.Mstorgnhirarchyid, "recordid": tw.Transactionid, "eventnotificationid": eventnotificationid, "channeltype": 1})
	responseBody := bytes.NewBuffer(postBody)
	logger.Log.Println("postBody  change group     --->", responseBody)
	log.Println("postBody  change group     --->", responseBody)

	go utility.Sendnotification(responseBody)
	if !tw.Samegroup {
		/**
		Change SLA for group change
		*/
		postBodysla, _ := json.Marshal(map[string]int64{"clientid": tw.Clientid, "mstorgnhirarchyid": tw.Mstorgnhirarchyid, "recordid": tw.Transactionid})
		responseBodysla := bytes.NewBuffer(postBodysla)
		logger.Log.Println("postBody  change group     --->", responseBodysla)
		log.Println("postBody  change group     --->", responseBodysla)
		respsla, err := http.Post(config.RECORD_URL+"/getsladuetimecalculate", "application/json", responseBodysla)
		if err != nil {
			logger.Log.Println("An Error Occured --->", err)
			log.Println("An Error Occured --->", err)
		}
		defer respsla.Body.Close()
		bodysla, err := ioutil.ReadAll(respsla.Body)
		if err != nil {
			logger.Log.Println("response body ------> ", err)
			log.Println("response body ------> ", err)
		}
		sbsla := string(bodysla)
		logger.Log.Println("sb change group body value is --->", sbsla)
		log.Println("sb change group body value is --->", sbsla)

		/**
		Sending mail for hop count change
		*/
		postBody, _ := json.Marshal(map[string]interface{}{"clientid": tw.Clientid, "mstorgnhirarchyid": tw.Mstorgnhirarchyid, "recordid": tw.Transactionid, "eventnotificationid": 6, "channeltype": 1, "hopcount": count, "lastgroupname": grpname, "lasttolastgroupname": prevgroup})
		responseBody := bytes.NewBuffer(postBody)
		logger.Log.Println("postBody  hop count change  --->", responseBody)
		log.Println("postBody  hop count change  --->", responseBody)
		go utility.Sendnotification(responseBody)
	}
	seq, _, err := dataAccess.Getdiffseqno(tw.Transactionid, 1)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if seq == config.INCIDENT_SEQ {
		err2, childids := dataAccess.Getchildticket(tw)
		if err2 != nil {
			return 0, false, err2, "Something Went Wrong"
		} else {
			if len(childids) > 0 {
				tw.Parentid = tw.Transactionid
				tw.Childids = childids

				_, success, _, msg := Updatechildstatus(tw, dbcon)
				logger.Log.Print("\n Update child status after group change:", success, msg)
				log.Print("\n Update child status after group change:", success, msg)
			}
		}
	}else if  seq == config.STASK_SEQ{
		err3, parentids := dataAccess.Getparentticket(tw)
		if err3 != nil {
			return 0, false, err3, "Something Went Wrong"
		} else {
			if len(parentids) > 0 {
				tw.Transactionid = parentids[0]
				_, success, _, msg := Changerecordgroup(tw,dbcon)
				logger.Log.Print("\n Update Parent status after group change:", success, msg)
				log.Print("\n Update Parent status after group change:", success, msg)
			}
		}
	}
	return 0, true, nil, ""
}

func Updatechildstatus(tz *entities.Workflowentity, db *sql.DB) (int64, bool, error, string) {
	logger.Log.Print("Updatechildstatus:", tz.Parentid, tz.Childids, tz.IsAttaching)
	log.Print("Updatechildstatus:", tz.Parentid, tz.Childids, tz.IsAttaching)

	//db, err := config.ConnectMySqlDb()
	//
	//if err != nil {
	//	logger.Log.Println("database connection failure", err)
	//	log.Println("database connection failure", err)
	//	return 0, false, err, "Something Went Wrong"
	//}
	//defer db.Close()
	dataAccess := dao.DbConn{DB: db}
	tz.Transactionid = tz.Parentid
	err, requestIds := dataAccess.GetRequestIdbyRecordId(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	var parentid int64
	if len(requestIds) > 0 {
		parentid = requestIds[0].Requestid
	} else {
		return 0, false, nil, "Parent Ticket is not mapped with process"
	}

	err, requestDetails := dataAccess.Getprocessrequestdetails(parentid)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if len(requestDetails) > 0 {
		tx, err := db.Begin()
		if err != nil {
			logger.Log.Println("Transaction creation error.", err)
			log.Println("Transaction creation error.", err)
			return 0, false, err, "Something Went Wrong"
		}
		for _, childid := range tz.Childids {
			tz.Transactionid = childid
			err, childRequestId := dataAccess.GetRequestIdbyRecordId(tz)
			if err != nil {
				return 0, false, err, "Something Went Wrong"
			}
			var id int64
			var isFirstStep bool
			if len(childRequestId) > 0 {
				id = childRequestId[0].Requestid
				isFirstStep = false
			} else {
				isFirstStep = true
			}
			childRequesthistid, err := dataAccess.Getlatesthistoryid(id)
			if err != nil {
				return 0, false, err, "Something Went Wrong"
			}
			if len(childRequesthistid) == 0 {
				return 0, false, err, "Something Went Wrong"
			}
			err, childRequestDetails := dataAccess.Getprocessrequestdetails(id)
			if err != nil {
				return 0, false, err, "Child ticket history not found"
			}
			if len(childRequestDetails) > 0 {
				/**
				Mark the ticket's condition(first time attaching to ticket)
				to fetch the details(before attaching) of the ticket.
				*/
				if tz.IsAttaching == 1 {
					err = dao.Updateattachdetails(tx, childRequesthistid[0].Id)
					if err != nil {
						tx.Rollback()
						return 0, false, err, "Something Went Wrong"
					}
				}
				requestDetails[0].Processid = childRequestDetails[0].Processid
				requestDetails[0].Createduserid = childRequestDetails[0].Createduserid
				requestDetails[0].Transactionid = childid
				requestDetails[0].Manualstateselection = 0
				/** Activity Log **/
				prevgroup := childRequestDetails[0].Groupname
				prevuser := childRequestDetails[0].Loginname
				grperr, group := dataAccess.Getgroupname(requestDetails[0].Mstgroupid)
				if grperr != nil {
					return 0, false, grperr, "Something Went Wrong"
				}
				var grpname string
				var username string
				var logstring string
				var name string

				if len(group) > 0 {
					grpname = group[0].Groupname
				}
				if childRequestDetails[0].Mstuserid > 0 {
					grperr, user := dataAccess.Getusername(requestDetails[0].Mstuserid)
					if grperr != nil {
						return 0, false, grperr, "Something Went Wrong"
					}
					if len(user) > 0 {
						username = user[0].Loginname
						name = user[0].Username
					}
					logstring = "From Group: " + prevgroup + " User: " + prevuser + " To Group: " + grpname + " User:" + username
				} else {
					logstring = "From Group: " + prevgroup + " User: " + prevuser + " To Group: " + grpname
				}
				logger.Log.Print("\n\nlogstring:->" + logstring)
				log.Print("\n\nlogstring:->" + logstring)
				stageDetails, trerr := dataAccess.GetLatestTransactionStageDetails(&requestDetails[0])
				if trerr != nil {
					return 0, false, trerr, "Something Went Wrong"
				}
				if len(stageDetails) > 0 {
					tran := entities.TransactionEntity{}
					tran.Recordstageid = stageDetails[0].Recordstageid
					tran.Recordtitle = childRequestDetails[0].Recordtitle

					_, recerr := dao.UpsertProcessDetails(tx, &requestDetails[0], tran, isFirstStep, id, "C")
					if recerr != nil {
						logger.Log.Println("Role back error.")
						log.Println("Role back error.")
						tx.Rollback()
						return 0, false, recerr, "Something Went Wrong"
					}
					parentgroupid := requestDetails[0].Mstgroupid
					parentuserid := requestDetails[0].Mstuserid
					status, serr := dataAccess.Getstatusbystateid(requestDetails[0].Clientid, requestDetails[0].Mstorgnhirarchyid, requestDetails[0].Currentstateid)
					if serr != nil {
						return 0, false, serr, "Something Went Wrong"
					}
					//var updaterequired = true
					if len(status) > 0 {
						logger.Log.Println("Parent Status seq: ", status[0].Seqno)
						log.Println("Parent Status seq: ", status[0].Seqno)
						if status[0].Seqno == config.CLOSE_SEQ || status[0].Seqno == config.RESOLVE_SEQ || status[0].Seqno == config.PENDING_USER_STATUS_SEQ || status[0].Seqno == config.PENDING_REQUESTER_INPUT || status[0].Seqno == config.CANCEL_SEQ || status[0].Seqno == config.REJECTED_STATUS_SEQ {
							//updaterequired = false
							lastloginname, lastgroupname, lastusername, lastgroupid, lastuserid, _, err := Getlastactionerdetails(tz.Parentid, requestDetails[0].Mstgroupid, db)
							if err != nil {
								//return nil, false, err, "Something Went Wrong"
							}
							grpname = lastgroupname
							parentgroupid = lastgroupid
							name = lastusername
							username = lastloginname
							parentuserid = lastuserid
						}
					}

					recordDetails, trerr := dataAccess.GetRecordDetailsById(tz, "trnrecord")
					if trerr != nil {
						return 0, false, trerr, "Something Went Wrong"
					}
					if len(recordDetails) > 0 {
						tz.Creatorgroupid = recordDetails[0].Groupid
					}
					count, success, _, msg := CalculateHopCount(db, requestDetails[0].Clientid, tz.Transactionid, tz.Creatorgroupid)
					if !success {
						logger.Log.Print("\n Error in hop count:: ", msg)
						log.Print("\n Error in hop count:: ", msg)
					}
					loggedusername := ""
					uerr, users := dataAccess.Getusername(tz.Userid)
					if uerr != nil {
						return 0, false, uerr, "Something Went Wrong"
					}
					if len(users) > 0 {
						loggedusername = users[0].Username
					}

					stagEntity := entities.StagingUtilityEntity{}
					stagEntity.Clientid = requestDetails[0].Clientid
					stagEntity.Mstorgnhirarchyid = requestDetails[0].Mstorgnhirarchyid
					stagEntity.Assignedgroupid = parentgroupid
					stagEntity.Assignedgroup = grpname
					stagEntity.Assigneduser = name
					stagEntity.Assignedloginname = username
					stagEntity.Assigneduserid = parentuserid
					stagEntity.Lastuser = loggedusername
					stagEntity.Lastuserid = tz.Userid
					stagEntity.Reassigncount = count
					stagEntity.Recordid = tz.Transactionid
					stgerr := dataAccess.Updatestagingdetailswithouttran(&stagEntity)
					if stgerr != nil {
						//tx.Rollback()
						return 0, false, stgerr, "Something Went Wrong"
					}

					//}
					var activityseq int64
					if tz.Samegroup {
						activityseq = 12
					} else {
						activityseq = 2
					}
					logger.Log.Print("\n\n activity log:->", requestDetails[0].Clientid, requestDetails[0].Mstorgnhirarchyid, tz.Transactionid, activityseq, logstring, tz.Userid, tz.Createdgroupid)
					log.Print("\n\n activity log:->", requestDetails[0].Clientid, requestDetails[0].Mstorgnhirarchyid, tz.Transactionid, activityseq, logstring, tz.Userid, tz.Createdgroupid)
					err = utility.InsertActivityLogs(tx, requestDetails[0].Clientid, requestDetails[0].Mstorgnhirarchyid, tz.Transactionid, activityseq, logstring, tz.Userid, tz.Createdgroupid)
					if err != nil {

						tx.Rollback()
						return 0, false, err, "Something Went Wrong"
					}
				} else {
					return 0, false, nil, "No Staging Record is mapped for this record id."
				}
			} else {
				return 0, false, nil, "No process details mapped with child ticket."
			}
		}
		err = tx.Commit()
		if err != nil {
			log.Print("Updatechildstatus  Statement Commit error", err)
			logger.Log.Print("Updatechildstatus  Statement Commit error", err)
			return 0, false, err, ""
		}
		return 0, true, nil, ""
	} else {
		return 0, false, nil, "Parent Process Details Not Found"
	}
}
func Updatechildstatuswithapi(tz *entities.Workflowentity) (int64, bool, error, string) {
	if utility.MutexLocked(lock) == false {
		lock.Lock()
		defer lock.Unlock()
	}
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	_, success, err, msg := Updatechildstatus(tz, db)
	return 0, success, err, msg
}
func Detachchildticket(tz *entities.Workflowentity) (int64, bool, error, string) {
	if utility.MutexLocked(lock) == false {
		lock.Lock()
		defer lock.Unlock()
	}
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err, parRequestId := dataAccess.GetRequestIdbyRecordId(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	var requestid int64
	if len(parRequestId) > 0 {
		requestid = parRequestId[0].Requestid
	} else {
		return 0, false, err, "Ticket details not mapped with process table"
	}
	childhistoryid, err := dataAccess.Getattachchildfirsthistoryid(requestid)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if len(childhistoryid) == 0 {
		return 0, false, err, "Ticket not attached to parent properly."
	}
	logger.Log.Println(" Child Attach Id:", childhistoryid[0].Id)
	log.Println(" Child Attach Id:", childhistoryid[0].Id)
	childhistory, err := dataAccess.Getattachchildhistory(childhistoryid[0].Id)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if len(childhistory) > 0 {
		tx, err := db.Begin()
		if err != nil {
			logger.Log.Println("Transaction creation error.", err)
			log.Println("Transaction creation error.", err)
			return 0, false, err, "Something Went Wrong"
		}
		latestTime := time.Now().Unix()
		err = dao.UpdateProcessRequest(&childhistory[0], tx, latestTime, childhistory[0].Mstrequestid)
		if err != nil {
			tx.Rollback()
			return 0, false, err, "Something went wrong"
		}
		rec := entities.TransactionEntity{}
		rec.Recordtitle = childhistory[0].Recordtitle
		histerr := dao.InsertProcessHistoryRequest(&childhistory[0], tx, latestTime, rec, childhistory[0].Mstrequestid, "N")
		if histerr != nil {
			tx.Rollback()
			return 0, false, err, "Something went wrong"
		}
		err = tx.Commit()
		if err != nil {
			log.Print("Detachchildticket  Statement Commit error", err)
			logger.Log.Print("Detachchildticket  Statement Commit error", err)
			return 0, false, err, ""
		}
		/**
		Getting hop count of the ticket
		Hop count is number of support group ,the ticket is moving except end user
		*/
		var count int64 = 0
		err, S := dataAccess.Gethopcount(requestid)
		if err != nil {
			logger.Log.Print("\n Error in hop count for mail sending")
			log.Print("\n Error in hop count for mail sending")
		} else {
			for i := 0; i < len(S)-1; i++ {
				if S[i] != S[i+1] {
					count = count + 1
				}
			}
		}
		loggedusername := ""
		uerr, users := dataAccess.Getusername(tz.Userid)
		if uerr != nil {
			return 0, false, uerr, "Something Went Wrong"
		}
		if len(users) > 0 {
			loggedusername = users[0].Username
		}
		var grpname string
		var username string
		var name string
		grperr, group := dataAccess.Getgroupname(childhistory[0].Mstgroupid)
		if grperr != nil {
			return 0, false, grperr, "Something Went Wrong"
		}
		if len(group) > 0 {
			grpname = group[0].Groupname
		}
		grperr, user := dataAccess.Getusername(childhistory[0].Mstuserid)
		if grperr != nil {
			return 0, false, grperr, "Something Went Wrong"
		}
		if len(user) > 0 {
			username = user[0].Loginname
			name = user[0].Username
		}

		stagEntity := entities.StagingUtilityEntity{}
		stagEntity.Clientid = childhistory[0].Clientid
		stagEntity.Mstorgnhirarchyid = childhistory[0].Mstorgnhirarchyid
		stagEntity.Assignedgroupid = childhistory[0].Mstgroupid
		stagEntity.Assignedgroup = grpname
		stagEntity.Assigneduser = name
		stagEntity.Assignedloginname = username
		stagEntity.Assigneduserid = childhistory[0].Mstuserid
		stagEntity.Lastuser = loggedusername
		stagEntity.Lastuserid = tz.Userid
		stagEntity.Reassigncount = count
		stagEntity.Recordid = tz.Transactionid
		stgerr := dataAccess.Updatestagingdetailswithouttran(&stagEntity)
		if stgerr != nil {
			//tx.Rollback()
			return 0, false, stgerr, "Something Went Wrong"
		}

		postBody, _ := json.Marshal(map[string]int64{"clientid": childhistory[0].Clientid, "mstorgnhirarchyid": childhistory[0].Mstorgnhirarchyid, "recordid": tz.Transactionid, "reordstatusid": childhistory[0].Currentstateid, "userid": tz.Userid, "usergroupid": tz.Createdgroupid, "changestatus": 0, "issrrequestor": 0})
		responseBody := bytes.NewBuffer(postBody)
		logger.Log.Println("postBody       --->", responseBody)
		log.Println("postBody       --->", responseBody)
		resp, err := http.Post(config.RECORD_URL+"/updaterecordstatus", "application/json", responseBody)
		if err != nil {
			logger.Log.Println("An Error Occured --->", err)
			log.Println("An Error Occured --->", err)
			return 0, false, err, "Something went wrong"
		}
		defer resp.Body.Close()
		//Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logger.Log.Println("response body ------> ", err)
			log.Println("response body ------> ", err)
			return 0, false, err, "Something went wrong"
		}
		sb := string(body)
		logger.Log.Println("sb body value is --->", sb)
		log.Println("sb body value is --->", sb)
		return 0, true, nil, ""
	} else {
		return 0, false, err, "Child history not found."
	}
}
func Updatetaskstatus(tz *entities.Workflowentity) (int64, bool, error, string) {
	logger.Log.Print("Updatetaskstatus:", tz.Parentid, tz.Childids, tz.Createdgroupid, tz.Userid, tz.Isupdate)
	log.Print("Updatetaskstatus:", tz.Parentid, tz.Childids, tz.Createdgroupid, tz.Userid, tz.Isupdate)
	if utility.MutexLocked(lock) == false {
		lock.Lock()
		defer lock.Unlock()
	}
	//lock.Lock()
	//defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()

	if err != nil {
		logger.Log.Println("database connection failure", err)
		log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	//defer db.Close()
	dataAccess := dao.DbConn{DB: db}
	tz.Transactionid = tz.Parentid
	typeseq, typeid, err := dataAccess.Getdiffseqno(tz.Parentid, 1)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	var statusid int64 = 0
	var statusseq int64 = 0
	var parclientid int64 = 0
	var parorgid int64 = 0
	err, parRequestId := dataAccess.GetRequestIdbyRecordId(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	var parid int64
	if len(parRequestId) > 0 {
		parid = parRequestId[0].Requestid
	} else {
		return 0, false, err, "Parent ticket details not mapped with process table"
	}
	err, parrequestDetails := dataAccess.Getprocessrequestdetails(parid)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if len(parrequestDetails) > 0 {
		parclientid = parrequestDetails[0].Clientid
		parorgid = parrequestDetails[0].Mstorgnhirarchyid
	} else {
		return 0, false, err, "Parent ticket details not found"
	}
	if typeseq == config.STASK_SEQ || typeseq == config.CTASK_SEQ {

		heighesttaskprios, err := dataAccess.Gethighestchildpriority(parclientid, parorgid, tz.Childids[0], typeid)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		if len(heighesttaskprios) > 0 {
			logger.Log.Println("heighesttaskprio status ", heighesttaskprios[0].Seqno)
			log.Println("heighesttaskprio status ", heighesttaskprios[0].Seqno)
			statusid = heighesttaskprios[0].Seqno
			logger.Log.Print("statusid :", statusid)
			log.Print("statusid :", statusid)
		} else {
			return 0, false, err, "Task status priority not mapped"
		}

	} else {
		innerstatusseq, innerstatusid, err := dataAccess.Getdiffseqno(tz.Parentid, 2)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		statusid = innerstatusid
		statusseq = innerstatusseq
		logger.Log.Print("statusid :", statusid, statusseq)
		log.Print("statusid :", statusid, statusseq)
	}

	//canchangestatus := true
	for _, childid := range tz.Childids {
		canProceed := true
		var ismanual int = 0
		var currentstateid int64 = 0
		_, childtypeid, err := dataAccess.Getdiffseqno(childid, 1)
		if err != nil {
			canProceed = false
		} else {
			/*if statusseq == config.CLOSE_SEQ || statusseq == config.CANCEL_SEQ || tz.Isupdate {
				ismanual = 1
			}*/
			if tz.Isupdate {
				currentstate, err := dataAccess.Getstatebystatusseq(parclientid, parorgid, config.RESOLVE_SEQ)
				if err != nil {
					canProceed = false
				} else {
					if len(currentstate) > 0 {
						currentstateid = currentstate[0].ID
						logger.Log.Println("after update currentstateid:", currentstateid)
						log.Println("after update currentstateid:", currentstateid)
					} else {
						canProceed = false
						logger.Log.Println("No status  mapped with State")
						log.Println("No status  mapped with State")
					}
				}
			} else {
				if statusseq == config.CANCEL_SEQ {
					logger.Log.Println("Parent is canceled:")
					log.Println("Parent is canceled:")
					childseq, _, err := dataAccess.Getdiffseqno(childid, 2)
					if err != nil {
						canProceed = false
					} else {
						currentstate, err := dataAccess.Getstatebystatusseq(parclientid, parorgid, childseq)
						if err != nil {
							canProceed = false
						} else {
							if len(currentstate) > 0 {
								currentstateid = currentstate[0].ID
								logger.Log.Println("cancel / close currentstateid:", currentstateid)
								log.Println("cancel / close currentstateid:", currentstateid)
							} else {
								canProceed = false
								logger.Log.Println("No status  mapped with State")
								log.Println("No status  mapped with State")
							}
						}
					}
				} else {
					taskstates, err := dataAccess.Getstateidbyfromdiff(statusid, typeid, childtypeid)
					if err != nil {
						return 0, false, err, "Something Went Wrong"
					}
					if len(taskstates) == 0 {
						canProceed = false
						logger.Log.Println("Child status not mapped with Parent status")
						log.Println("Child status not mapped with Parent status")
					} else {
						currentstateid = taskstates[0].Id
					}
				}
			}
		}
		if canProceed {
			tz.Transactionid = childid
			err, childRequestId := dataAccess.GetRequestIdbyRecordId(tz)
			if err != nil {
				return 0, false, err, "Something Went Wrong"
			}
			var id int64
			if len(childRequestId) > 0 {
				id = childRequestId[0].Requestid
			} else {
			}
			err, requestDetails := dataAccess.Getprocessrequestdetails(id)
			if err != nil {
				return 0, false, err, "Something Went Wrong"
			}
			if len(requestDetails) > 0 {
				if typeseq == config.STASK_SEQ || typeseq == config.CTASK_SEQ {
					requestDetails[0].Changestatus = 1
				}
				logger.Log.Println("Previous:", requestDetails[0].Currentstateid)
				logger.Log.Println("Current:", currentstateid)
				log.Println("Previous:", requestDetails[0].Currentstateid)
				log.Println("Current:", currentstateid)
				if requestDetails[0].Currentstateid != currentstateid {
					requestDetails[0].Previousstateid = requestDetails[0].Currentstateid
					requestDetails[0].Currentstateid = currentstateid
					requestDetails[0].Transactionid = childid
					//requestDetails[0].Manualstateselection = ismanual
					workingdifftype, workingid, err := dataAccess.Getworkingdiffbytid(childid)
					if err != nil {
						return 0, false, err, "Something Went Wrong"
					}
					requestDetails[0].Recorddifftypeid = workingdifftype
					requestDetails[0].Recorddiffid = workingid
					requestDetails[0].Createdgroupid = tz.Createdgroupid
					requestDetails[0].Mstgroupid = tz.Createdgroupid
					requestDetails[0].Mstuserid = tz.Userid
					requestDetails[0].Userid = tz.Userid
					requestDetails[0].Transitionid = 0
					if tz.Isupdate {
						ismanual = 1
					} else if statusseq == config.CLOSE_SEQ || statusseq == config.CANCEL_SEQ {
						transitionState, err := dataAccess.GetTransitionState(&requestDetails[0])
						if err != nil {
							return 0, false, err, "Something Went Wrong"
						}
						log.Print("\n\nTransition state task:", transitionState)
						logger.Log.Print("\n\nTransition state task:", transitionState)
						if len(transitionState) > 0 {
							ismanual = 0
						} else {
							ismanual = 1
						}
					}
					requestDetails[0].Manualstateselection = ismanual
					_, success, err, msg := MoveWorkflow(&requestDetails[0], db)
					if err != nil {
						log.Print("Error in task status change ", childid)
						logger.Log.Print("Error in task status change ", childid)
					} else {
						log.Print("\n\n--------", success, msg)
						logger.Log.Print("\n\n--------", success, msg)
					}
					//return 0, success, err, msg
				} else {
					return 0, false, nil, "Current and Previous status are same"
				}

			} else {
				return 0, false, nil, "No process details mapped with child ticket."
			}
		} else {
			return 0, false, nil, ""
		}
	}
	return 0, true, nil, ""
}

func Gethopcount(tz *entities.Workflowentity) (int64, bool, error, string) {
	logger.Log.Print("Gethopcount:", tz.Transitionid)
	log.Print("Gethopcount:", tz.Transitionid)
	//var count int64 = 0
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	//db, err := config.ConnectMySqlDb()

	if err != nil {
		logger.Log.Println("database connection failure", err)
		log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	//defer db.Close()
	count, success, err, msg := CalculateHopCount(db, tz.Clientid, tz.Transactionid, tz.Createdgroupid)
	if !success {
		logger.Log.Print("\n Error in hop count:: ", msg)
		log.Print("\n Error in hop count:: ", msg)
	}
	return count, success, err, msg
	/*dataAccess := dao.DbConn{DB: db}
	dispatcher, err := dataAccess.Searchgroupbyname(tz.Clientid, "Dispatcher")
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	var dispatcherid int64 = 0
	if len(dispatcher) > 0 {
		dispatcherid = dispatcher[0].Id
	}
	err, requestIds := dataAccess.GetRequestIdbyRecordId(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if len(requestIds) > 0 {
		err, S := dataAccess.Gethopcount(requestIds[0].Requestid)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		} else {
			for i := 0; i < len(S)-1; i++ {
				//logger.Log.Println(S[i], S[i+1], tz.Createdgroupid)

				if S[i] != tz.Createdgroupid && S[i] != dispatcherid && S[i+1] != tz.Createdgroupid {
					if S[i] != S[i+1] {
						count = count + 1
					}
				}
			}
			return count, true, nil, ""
		}
	} else {
		return 0, false, nil, "Ticket is not mapped with process"
	}*/
}

func CalculateHopCount(db *sql.DB, clientid int64, transactionid int64, creatorgroupid int64) (int64, bool, error, string) {
	logger.Log.Println("Inside Hop Count :: ", clientid, transactionid, creatorgroupid)
	log.Println("Inside Hop Count :: ", clientid, transactionid, creatorgroupid)
	var count int64 = 0
	dataAccess := dao.DbConn{DB: db}
	dispatcher, err := dataAccess.Searchgroupbyname(clientid, "Dispatcher")
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	var dispatcherid int64 = 0
	if len(dispatcher) > 0 {
		dispatcherid = dispatcher[0].Id
	}
	tz := entities.Workflowentity{}
	tz.Transactionid = transactionid
	err, requestIds := dataAccess.GetRequestIdbyRecordId(&tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if len(requestIds) > 0 {
		err, S := dataAccess.Gethopcount(requestIds[0].Requestid)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		} else {
			for i := 0; i < len(S)-1; i++ {
				//logger.Log.Println(S[i], S[i+1], tz.Createdgroupid)

				if S[i] != creatorgroupid && S[i] != dispatcherid && S[i+1] != creatorgroupid {
					if S[i] != S[i+1] {
						count = count + 1
					}
				}
			}
			return count, true, nil, ""
		}
	} else {
		return 0, false, nil, "Ticket is not mapped with process"
	}

}
