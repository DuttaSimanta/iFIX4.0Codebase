package models

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	dbconfig "src/config"
	"src/dao"
	"src/entities"
	"src/logger"
	"time"
)

//var lock1 = &sync.Mutex{}

func RecordStatusUpdation(tz *entities.RecordstatusEntity) (int64, bool, error, string) {

	// if mutexutility.MutexLocked(lock) == false {
	// 	lock.Lock()
	// 	defer lock.Unlock()
	// }

	// db, err := dbconfig.ConnectMySqlDb()
	// if err != nil {
	// 	log.Println("database connection failure", err)
	// 	return 0, false, err, "Something Went Wrong"
	// }
	//defer db.Close()
	if db == nil {
		dbcon, err := ConnectMySqlDb()
		if err != nil {
			logger.Log.Println("Error in DBConnection")
			return 0, false, err, "Something Went Wrong"
		}
		db = dbcon
	}
	dataAccess := dao.DbConn{DB: db}
	recordtypeSeq, err := dataAccess.Getrecordtypeseq(tz.ClientID, tz.Mstorgnhirarchyid, tz.RecordID)
	if err != nil {
		logger.Log.Println(err)
		return 0, false, err, "Something Went Wrong"
	}
	if recordtypeSeq == 1 {
		data, success, err, msg := Updaterecordstatus(tz, db)
		if err != nil {
			logger.Log.Println("Error is ----------->", err)
			return 0, false, err, "Something Went Wrong"
		}
		logger.Log.Println("Error is ----------->", data, success, err, msg)
		return data, success, err, msg
	} else if recordtypeSeq == 2 || recordtypeSeq == 4 {
		data, success, err, msg := UpdateSRrecordstatus(tz, db, recordtypeSeq)
		if err != nil {
			logger.Log.Println("Error is ----------->", err)
			return 0, false, err, "Something Went Wrong"
		}
		logger.Log.Println("Error is ----------->", data, success, err, msg)
		return data, success, err, msg
	} else if recordtypeSeq == 3 || recordtypeSeq == 5 {
		data, success, err, msg := UpdateStaskrecordstatus(tz, db, recordtypeSeq)
		if err != nil {
			logger.Log.Println("Error is ----------->", err)
			return 0, false, err, "Something Went Wrong"
		}
		logger.Log.Println("Error is ----------->", data, success, err, msg)
		return data, success, err, msg
	}

	return 0, false, err, "Something Went Wrong"
}

func Updaterecordstatus(tz *entities.RecordstatusEntity, db *sql.DB) (int64, bool, error, string) {
	logger.Log.Println("In side Updaterecordstatus----------000000000000000000000000000---------------------------->", tz)
	tx, err := db.Begin()
	if err != nil {
		logger.Log.Println("Transaction creation error in Updaterecordstatus", err)
		return 0, false, err, "Something Went Wrong"
	}

	dataAccess := dao.DbConn{DB: db}
	diffid, err := dataAccess.Getrecordtypediffid(tz.RecordID, tz.ClientID, tz.Mstorgnhirarchyid)
	if err != nil {
		log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	childids, err := dataAccess.Getchildrecordids(tz.RecordID, tz.ClientID, tz.Mstorgnhirarchyid, 2, diffid)
	if err != nil {
		log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}

	ID, err := Parentrecordstatusupdation(tx, tz, db)
	if err != nil {
		log.Println("Parent Record status updation failed", err)
		tx.Rollback()
		//db.Close()
		return 0, false, err, "Something Went Wrong"
	}
	logger.Log.Println("Child records length is ----------->", len(childids))
	for i := 0; i < len(childids); i++ {
		stateID, err := dataAccess.GetrecordlateststateID(tz.RecordID, tz.ClientID, tz.Mstorgnhirarchyid)
		if err != nil {
			log.Println("Find child record status error", err)
			tx.Rollback()
			//db.Close()
			return 0, false, err, "Something Went Wrong"
		}
		_, err = Childrecordstatusupdation(tx, tz.ClientID, tz.Mstorgnhirarchyid, childids[i], stateID, tz.UserID, tz.Usergroupid, db)
		if err != nil {
			log.Println("Child Record status updation failed", err)
			tx.Rollback()
			//db.Close()
			return 0, false, err, "Something Went Wrong"
		}
	}

	var workflowflag bool
	var errormsg string
	if len(childids) > 0 {
		reqbd := &entities.ParentchildEntity{}
		reqbd.Parentid = tz.RecordID
		reqbd.Childids = childids
		reqbd.Userid = tz.UserID
		reqbd.Createdgroupid = tz.Usergroupid
		postBody, _ := json.Marshal(reqbd)

		logger.Log.Println("Record status request body -->", reqbd)

		responseBody := bytes.NewBuffer(postBody)
		resp, err := http.Post(dbconfig.MASTER_URL+"/updatechildstatus", "application/json", responseBody)
		if err != nil {
			logger.Log.Println("An Error Occured --->", err)
			return 0, false, err, "Something Went Wrong"
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logger.Log.Println("response body ------> ", err)
			return 0, false, err, "Something Went Wrong"
		}
		sb := string(body)
		wfres := entities.WorkflowResponse{}
		json.Unmarshal([]byte(sb), &wfres)
		workflowflag = wfres.Success
		errormsg = wfres.Message
		logger.Log.Println("Record status response message -->", workflowflag)
		logger.Log.Println("Record status response error message -->", errormsg)
	}
	logger.Log.Println("ID value is -------- -->", ID)
	if ID > 0 {
		err = tx.Commit()
		if err != nil {
			log.Println("DB commit is failed", err)
			tx.Rollback()
			//db.Close()
			return 0, false, err, "Something Went Wrong"
		}

		//Email Notification Start Here
		statusID, _, _, _ := dataAccess.Getcurrentsatusid(tz.ClientID, tz.Mstorgnhirarchyid, tz.RecordID)
		logger.Log.Println(statusID)
		go StatusChangeEmail(tz.ClientID, tz.Mstorgnhirarchyid, tz.RecordID, statusID)
		for i := 0; i < len(childids); i++ {
			go StatusChangeEmail(tz.ClientID, tz.Mstorgnhirarchyid, childids[i], statusID)
		}
		//Email Notification End Here
		//db.Close()
		return ID, true, err, ""
	}

	return 0, false, err, "Something Went Wrong"
}

func Parentrecordstatusupdation(tx *sql.Tx, tz *entities.RecordstatusEntity, db *sql.DB) (int64, error) {
	logger.Log.Println("In side Mstslastatemodel")
	dataAccess := dao.DbConn{DB: db}
	recordtypeSeq, err := dataAccess.Getrecordtypeseq(tz.ClientID, tz.Mstorgnhirarchyid, tz.RecordID)
	if err != nil {
		logger.Log.Println(err)
		return 0, err
	}
	recorddiffid, recorddiffseq, currentstatusname, err := dataAccess.Getrecorddiffidbystateid(tz.ClientID, tz.Mstorgnhirarchyid, tz.ReordstatusID)
	if err != nil {
		logger.Log.Println(err)
		return 0, err
	}
	if recorddiffid > 0 {
		laststageID, err := dataAccess.GetMaxstageID(tz.ClientID, tz.Mstorgnhirarchyid, tz.RecordID)
		if err != nil {
			logger.Log.Println(err)
			return 0, err
		}
		previousstatus, previousseq, previousstatusname, err := dataAccess.Getcurrentsatusid(tz.ClientID, tz.Mstorgnhirarchyid, tz.RecordID)
		if err != nil {
			logger.Log.Println(err)
			return 0, err
		}

		err = dao.Updatepreviousstatus(tx, tz.RecordID, tz.ClientID, tz.Mstorgnhirarchyid)
		if err != nil {
			logger.Log.Println(err)
			return 0, err
		}
		id, err := dao.Updaterecordstatus(tx, tz.ClientID, tz.Mstorgnhirarchyid, tz.RecordID, recorddiffid, laststageID)
		if err != nil {
			logger.Log.Println(err)
			return 0, err
		}

		res, err := dataAccess.Getrecorddetails(tz.RecordID)
		if err != nil {
			logger.Log.Println(err)
			return 0, err
		}
		returnValue, _, _, _ := SLACriteriaRespResl(tz.ClientID, tz.Mstorgnhirarchyid, res.RecordtypeID, res.WorkingcatID, res.PriorityID)

		if id > 0 {

			//activity log entry here
			if previousstatus != recorddiffid {
				err = dao.InsertActivityLogs(tx, tz.ClientID, tz.Mstorgnhirarchyid, tz.RecordID, 4, "From "+previousstatusname+" To "+currentstatusname, tz.UserID, tz.Usergroupid)
				if err != nil {
					log.Println("error is ----->", err)
					return 0, err
				}
				//Update Stage TBL For Status

				err = dao.UpdateStageStatus(tx, tz.ClientID, tz.Mstorgnhirarchyid, tz.RecordID, recorddiffid, currentstatusname)
				if err != nil {
					log.Println("error is ----->", err)
					tx.Rollback()
					//	db.Close()
					return 0, err
				}

				//Update Stage TBL For Status
			} else {
				prename, err := dataAccess.Getprestausname(tz.ClientID, tz.Mstorgnhirarchyid, tz.RecordID)
				if err != nil {
					log.Println("error is ----->", err)
					return 0, err
				}
				err = dao.InsertActivityLogs(tx, tz.ClientID, tz.Mstorgnhirarchyid, tz.RecordID, 4, "From "+prename+" To "+currentstatusname, tz.UserID, tz.Usergroupid)
				if err != nil {
					log.Println("error is ----->", err)
					return 0, err
				}

				//Update Stage TBL For Status

				err = dao.UpdateStageStatus(tx, tz.ClientID, tz.Mstorgnhirarchyid, tz.RecordID, recorddiffid, currentstatusname)
				if err != nil {
					log.Println("error is ----->", err)
					tx.Rollback()
					//db.Close()
					return 0, err
				}

				//Update Stage TBL For Status
			}

			if recorddiffseq == 2 {
				err := UpdateResponseValueinStagetbl(tx, db, tz.ClientID, tz.Mstorgnhirarchyid, tz.RecordID, recorddiffseq, tz.Usergroupid, returnValue.Supportgroupspecific, recordtypeSeq)
				if err != nil {
					logger.Log.Println("error is ----->", err)
					tx.Rollback()
					return 0, err
				}

			}

			if recorddiffseq == 3 {
				err := UpdateResolutionValueinStagetbl(tx, db, tz.ClientID, tz.Mstorgnhirarchyid, tz.RecordID, recorddiffseq, tz.Usergroupid, returnValue.Supportgroupspecific, recordtypeSeq)
				if err != nil {
					logger.Log.Println("error is ----->", err)
					tx.Rollback()
					return 0, err
				}

				err = UpdateStageResolverInfo(tx, db, tz.ClientID, tz.Mstorgnhirarchyid, tz.RecordID, tz.UserID, tz.Usergroupid, returnValue.Supportgroupspecific, recorddiffseq, recordtypeSeq)
				if err != nil {
					logger.Log.Println(err)
					tx.Rollback()
					return 0, err
				}

			}

			if recorddiffseq == 10 {
				err = dao.UpdateReopenCount(tx, tz.ClientID, tz.Mstorgnhirarchyid, tz.RecordID)
				if err != nil {
					log.Println("error is ----->", err)
					tx.Rollback()
					//db.Close()
					return 0, err
				}
			}

			if recorddiffseq == 9 {
				err := UpdateUserreplytimetakenValueinStagetbl(tx, db, tz.ClientID, tz.Mstorgnhirarchyid, tz.RecordID, recorddiffseq, tz.Usergroupid, returnValue.Supportgroupspecific, recordtypeSeq)
				if err != nil {
					logger.Log.Println("error is ----->", err)
					tx.Rollback()
					return 0, err
				}
			}

			if recorddiffseq == 5 {
				err = dao.UpdatePendinguserAction(tx, tz.ClientID, tz.Mstorgnhirarchyid, tz.RecordID)
				if err != nil {
					log.Println("error is ----->", err)
					tx.Rollback()
					//db.Close()
					return 0, err
				}
			}

			if recorddiffseq == 4 {

				err = dao.UpdateFollowupcount(tx, tz.ClientID, tz.Mstorgnhirarchyid, tz.RecordID)
				if err != nil {
					log.Println("error is ----->", err)
					tx.Rollback()
					//db.Close()
					return 0, err
				}
			}

			if recorddiffseq == 8 {
				err := UpdateCloseValueinStagetbl(tx, db, tz.ClientID, tz.Mstorgnhirarchyid, tz.RecordID, recorddiffseq, tz.Usergroupid, returnValue.Supportgroupspecific, recordtypeSeq)
				if err != nil {
					logger.Log.Println("error is ----->", err)
					tx.Rollback()
					return 0, err
				}
			}

			if recorddiffseq == 2 && previousseq == 4 {
				err := UpdateFollowuptimetakenValueinStagetbl(tx, db, tz.ClientID, tz.Mstorgnhirarchyid, tz.RecordID, tz.Usergroupid, returnValue.Supportgroupspecific, recordtypeSeq)
				if err != nil {
					logger.Log.Println("error is ----->", err)
					tx.Rollback()
					return 0, err
				}

			}

			Username, err := dataAccess.GetUsername(tz.UserID)
			if err != nil {
				logger.Log.Println(err)
				tx.Rollback()
				return 0, err
			}

			err = dao.UpdateUserInfo(tx, tz.ClientID, tz.Mstorgnhirarchyid, tz.RecordID, tz.UserID, Username)
			if err != nil {
				logger.Log.Println(err)
				tx.Rollback()
				return 0, err
			}

			//activity log entry end
			//Condition For SLA
			if recordtypeSeq == 1 {
				currentTime := time.Now()
				zonediff, _, _, _ := Getutcdiff(tz.ClientID, tz.Mstorgnhirarchyid)
				datetime := AddSubSecondsToDate(currentTime, zonediff.UTCdiff)
				//For Response meter status checking
				seq, err := dataAccess.Getemeterseqno(tz.ClientID, tz.Mstorgnhirarchyid, recorddiffid, 1)
				if err != nil {
					logger.Log.Println(err)
					return 0, err
				}
				logger.Log.Println("Responsemeter sequance no ---->", seq)
				if seq > 0 {
					res, err := dataAccess.Getrecorddetails(tz.RecordID)
					if err != nil {
						logger.Log.Println(err)
						return 0, err
					}
					res.SupportgroupId = tz.Usergroupid
					GetSLAResolution(&res)
					if seq == 4 {
						flag, err := dataAccess.UpdateResponseEndFlag(tz.ClientID, tz.Mstorgnhirarchyid, tz.RecordID, currentTime.Format("2006-01-02 15:04:05"))
						logger.Log.Println(flag)
						if err != nil {
							logger.Log.Println(err)
							return 0, err
						}
					}

				}

				//For Resolution meter status checking
				seq, err1 := dataAccess.Getemeterseqno(tz.ClientID, tz.Mstorgnhirarchyid, recorddiffid, 2)
				if err1 != nil {
					logger.Log.Println(err1)
					return 0, err
				}
				logger.Log.Println("Resolutionmeter sequance no ---->", seq)
				if seq > 0 {
					res, err := dataAccess.Getrecorddetails(tz.RecordID)
					if err != nil {
						logger.Log.Println(err)
						return 0, err
					}
					res.SupportgroupId = tz.Usergroupid
					historyrecord, err := dataAccess.GetLatesttrnhistory(tz.ClientID, tz.Mstorgnhirarchyid, tz.RecordID)
					if err != nil {
						logger.Log.Println(err)
						return 0, err
					}
					if seq != 1 {
						// Change in 15.05.2021 -----------------------------

						histrn := entities.TrnslaentityhistoryEntity{}
						histrn.Clientid = tz.ClientID
						histrn.Mstorgnhirarchyid = tz.Mstorgnhirarchyid
						histrn.Therecordid = tz.RecordID
						if historyrecord.Slastartstopindicator == 2 && seq == 2 {
							var dt = historyrecord.Recorddatetime
							var dttime = historyrecord.Recorddatetoint
							histrn.Recorddatetime = dt
							histrn.Recorddatetoint = dttime
						} else {
							histrn.Recorddatetime = datetime.Format("2006-01-02 15:04:05")
							histrn.Recorddatetoint = TimeParse(datetime.Format("2006-01-02 15:04:05"), "").Unix()
						}
						histrn.Slastartstopindicator = seq

						trnid, err := dataAccess.InsertTrnslaentityhistory(&histrn)
						if err != nil {
							logger.Log.Println(err)
							return 0, err
						}
						logger.Log.Println("history table id---->", trnid)

					}

					if seq == 2 {

						_, _, err, _ = GetSLAResolution(&res)
						if err != nil {
							logger.Log.Println(err)
							return 0, err
						}
					}
					if seq == 4 {
						_, err, _ := UpdateRessolutionEndFlag(tz.ClientID, tz.Mstorgnhirarchyid, tz.RecordID)
						if err != nil {
							logger.Log.Println(err)
							return 0, err
						}
						_, _, err, _ = GetSLAResolution(&res)
						if err != nil {
							logger.Log.Println(err)
							return 0, err
						}
					}
					if seq == 1 || seq == 5 || seq == 3 {
						grpID, err := dataAccess.FetchCurrentGrpID(tz.RecordID)
						if err != nil {
							return 0, err
						}
						returnValue, _, _, _ := SLACriteriaRespResl(tz.ClientID, tz.Mstorgnhirarchyid, res.RecordtypeID, res.WorkingcatID, res.PriorityID)
						if returnValue.Supportgroupspecific == 1 {
							count, err := dataAccess.GetSupportgrpdayofweekcount(tz.ClientID, tz.Mstorgnhirarchyid, grpID)
							if err != nil {
								return 0, err
							}
							if count < 7 {
								return 0, errors.New("Day Of Week Not Properly Configured.Please Check.")
							}
						} else {
							count, err := dataAccess.GetOrganizationdayofweekcount(tz.ClientID, tz.Mstorgnhirarchyid)
							if err != nil {
								return 0, err
							}
							if count < 7 {
								return 0, errors.New("Day Of Week Not Properly Configured.Please Check.")
							}
						}
						SLADueTimeCalculation(tz.RecordID, 0, 1, 3, datetime.Format("2006-01-02 15:04:05"), tz.ClientID, tz.Mstorgnhirarchyid, res.RecordtypeID, res.WorkingcatID, res.PriorityID, "", grpID)
					}

					t := entities.SLATabEntity{}
					t.ClientID = tz.ClientID
					t.Mstorgnhirarchyid = tz.Mstorgnhirarchyid
					t.RecordID = tz.RecordID
					logger.Log.Println("111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111")
					sladtls, _, err, _ := GetSLATabvalues(&t)
					if err != nil {
						logger.Log.Println(err)
					}
					logger.Log.Println(sladtls)
					err = dao.UpdateSLAFields(tx, tz.ClientID, tz.Mstorgnhirarchyid, tz.RecordID, sladtls.Responsedetails.Responseduetime, sladtls.Responsedetails.Responseclockstatus, sladtls.Resolutionetails.Resolutionduetime, sladtls.Resolutionetails.Resolutionclockstatus)
					if err != nil {
						logger.Log.Println(err)
					}
				}
			}
			//End
			return id, err
		} else {
			return 0, err
		}
	} else {
		return 0, err
	}

}

func Childrecordstatusupdation(tx *sql.Tx, ClientID int64, Mstorgnhirarchyid int64, Recordid int64, StateID int64, UserID int64, Usergroupid int64, db *sql.DB) (int64, error) {
	logger.Log.Println("In side Mstslastatemodel")
	//Condition For SLA
	currentTime := time.Now()
	zonediff, _, _, _ := Getutcdiff(ClientID, Mstorgnhirarchyid)
	datetime := AddSubSecondsToDate(currentTime, zonediff.UTCdiff)
	// db, err := dbconfig.ConnectMySqlDb()
	// if err != nil {
	// 	log.Println("database connection failure", err)
	// 	return 0, err
	// }
	// defer db.Close()
	dataAccess := dao.DbConn{DB: db}
	recordtypeSeq, err := dataAccess.Getrecordtypeseq(ClientID, Mstorgnhirarchyid, Recordid)
	if err != nil {
		logger.Log.Println(err)
		return 0, err
	}
	recorddiffid, seqno, currentstatusname, err := dataAccess.Getrecorddiffidbystateid(ClientID, Mstorgnhirarchyid, StateID)
	if err != nil {
		logger.Log.Println(err)
		return 0, err
	}
	if recorddiffid > 0 {

		previousstatus, currentseqno, previousstatusname, err := dataAccess.Getcurrentsatusid(ClientID, Mstorgnhirarchyid, Recordid)
		if err != nil {
			logger.Log.Println(err)
			return 0, err
		}

		res, err := dataAccess.Getrecorddetails(Recordid)
		if err != nil {
			logger.Log.Println(err)
			return 0, err
		}
		returnValue, _, _, _ := SLACriteriaRespResl(ClientID, Mstorgnhirarchyid, res.RecordtypeID, res.WorkingcatID, res.PriorityID)

		if seqno == 3 { // this is the ticket status sequance number
			err = dao.Updatechildrecord(tx, Recordid, ClientID, Mstorgnhirarchyid)
			if err != nil {
				logger.Log.Println(err)
				return 0, err
			}
			// new logic 25.08.2021
			err = UpdateStageResolverInfo(tx, db, ClientID, Mstorgnhirarchyid, Recordid, UserID, Usergroupid, returnValue.Supportgroupspecific, seqno, recordtypeSeq)
			if err != nil {
				logger.Log.Println(err)
				tx.Rollback()
				return 0, err
			}
			// new logic 25.08.2021
		}

		laststageID, err := dataAccess.GetMaxstageID(ClientID, Mstorgnhirarchyid, Recordid)
		if err != nil {
			logger.Log.Println(err)
			return 0, err
		}
		err = dao.Updatepreviousstatus(tx, Recordid, ClientID, Mstorgnhirarchyid)
		if err != nil {
			logger.Log.Println(err)
			return 0, err
		}
		// New logic added in 14.06.2021
		if currentseqno == 1 || currentseqno == 10 {
			//get Active status ID
			activeID, err := dataAccess.GetActivestatusID(ClientID, Mstorgnhirarchyid)
			if err != nil {
				logger.Log.Println(err)
				return 0, err
			}
			_, err = dao.UpdaterecordActivestatus(tx, ClientID, Mstorgnhirarchyid, Recordid, activeID, laststageID)
			if err != nil {
				logger.Log.Println(err)
				return 0, err
			}
			flag, err := dataAccess.UpdateResponseEndFlag(ClientID, Mstorgnhirarchyid, Recordid, currentTime.Format("2006-01-02 15:04:05"))
			logger.Log.Println(flag)
			if err != nil {
				logger.Log.Println(err)
				return 0, err
			}
		}
		// End New logic added in 14.06.2021
		id, err := dao.Updaterecordstatus(tx, ClientID, Mstorgnhirarchyid, Recordid, recorddiffid, laststageID)
		if err != nil {
			logger.Log.Println(err)
			return 0, err
		}

		if id > 0 {
			//activity log entry here
			if previousstatus != recorddiffid {
				err = dao.InsertActivityLogs(tx, ClientID, Mstorgnhirarchyid, Recordid, 4, "From "+previousstatusname+" To "+currentstatusname, UserID, Usergroupid)
				if err != nil {
					log.Println("error is ----->", err)
					return 0, err
				}

				//Update Stage TBL For Status

				err = dao.UpdateStageStatus(tx, ClientID, Mstorgnhirarchyid, Recordid, recorddiffid, currentstatusname)
				if err != nil {
					log.Println("error is ----->", err)
					tx.Rollback()
					//db.Close()
					return 0, err
				}

			} else {
				prename, err := dataAccess.Getprestausname(ClientID, Mstorgnhirarchyid, Recordid)
				if err != nil {
					log.Println("error is ----->", err)
					return 0, err
				}
				err = dao.InsertActivityLogs(tx, ClientID, Mstorgnhirarchyid, Recordid, 4, "From "+prename+" To "+currentstatusname, UserID, Usergroupid)
				if err != nil {
					log.Println("error is ----->", err)
					return 0, err
				}

				//Update Stage TBL For Status

				err = dao.UpdateStageStatus(tx, ClientID, Mstorgnhirarchyid, Recordid, recorddiffid, currentstatusname)
				if err != nil {
					log.Println("error is ----->", err)
					tx.Rollback()
					//db.Close()
					return 0, err
				}

				//Update Stage TBL For Status
			}

			if seqno == 2 {
				err := UpdateResponseValueinStagetbl(tx, db, ClientID, Mstorgnhirarchyid, Recordid, seqno, Usergroupid, returnValue.Supportgroupspecific, recordtypeSeq)
				if err != nil {
					logger.Log.Println("error is ----->", err)
					tx.Rollback()
					return 0, err
				}

			}

			if seqno == 3 {
				err := UpdateResolutionValueinStagetbl(tx, db, ClientID, Mstorgnhirarchyid, Recordid, seqno, Usergroupid, returnValue.Supportgroupspecific, recordtypeSeq)
				if err != nil {
					logger.Log.Println("error is ----->", err)
					tx.Rollback()
					return 0, err
				}

			}

			if seqno == 8 {

				err := UpdateCloseValueinStagetbl(tx, db, ClientID, Mstorgnhirarchyid, Recordid, seqno, Usergroupid, returnValue.Supportgroupspecific, recordtypeSeq)
				if err != nil {
					logger.Log.Println("error is ----->", err)
					tx.Rollback()
					return 0, err
				}
			}

			if seqno == 10 {
				err = dao.UpdateReopenCount(tx, ClientID, Mstorgnhirarchyid, Recordid)
				if err != nil {
					log.Println("error is ----->", err)
					tx.Rollback()
					//db.Close()
					return 0, err
				}
			}

			if seqno == 9 {
				err := UpdateUserreplytimetakenValueinStagetbl(tx, db, ClientID, Mstorgnhirarchyid, Recordid, seqno, Usergroupid, returnValue.Supportgroupspecific, recordtypeSeq)
				if err != nil {
					logger.Log.Println("error is ----->", err)
					tx.Rollback()
					return 0, err
				}
			}

			if seqno == 5 {
				err = dao.UpdatePendinguserAction(tx, ClientID, Mstorgnhirarchyid, Recordid)
				if err != nil {
					log.Println("error is ----->", err)
					tx.Rollback()
					//db.Close()
					return 0, err
				}
			}

			if seqno == 4 {
				err = dao.UpdateFollowupcount(tx, ClientID, Mstorgnhirarchyid, Recordid)
				if err != nil {
					log.Println("error is ----->", err)
					tx.Rollback()
					//db.Close()
					return 0, err
				}
			}

			if seqno == 2 && currentseqno == 4 {
				err := UpdateFollowuptimetakenValueinStagetbl(tx, db, ClientID, Mstorgnhirarchyid, Recordid, Usergroupid, returnValue.Supportgroupspecific, recordtypeSeq)
				if err != nil {
					logger.Log.Println("error is ----->", err)
					tx.Rollback()
					return 0, err
				}
			}

			Username, err := dataAccess.GetUsername(UserID)
			if err != nil {
				logger.Log.Println(err)
				tx.Rollback()
				return 0, err
			}

			err = dao.UpdateUserInfo(tx, ClientID, Mstorgnhirarchyid, Recordid, UserID, Username)
			if err != nil {
				logger.Log.Println(err)
				tx.Rollback()
				return 0, err
			}

			//activity log entry end

			//For Response meter status checking
			seq, err := dataAccess.Getemeterseqno(ClientID, Mstorgnhirarchyid, recorddiffid, 1)
			if err != nil {
				logger.Log.Println(err)
				return 0, err
			}
			logger.Log.Println("Responsemeter sequance no ---->", seq)
			if seq > 0 {
				res, err := dataAccess.Getrecorddetails(Recordid)
				res.SupportgroupId = Usergroupid
				if err != nil {
					logger.Log.Println(err)
					return 0, err
				}
				_, _, err, _ = GetSLAResolution(&res)
				if err != nil {
					logger.Log.Println(err)
					return 0, err
				}
				if seq == 4 {
					flag, err := dataAccess.UpdateResponseEndFlag(ClientID, Mstorgnhirarchyid, Recordid, currentTime.Format("2006-01-02 15:04:05"))
					logger.Log.Println(flag)
					if err != nil {
						logger.Log.Println(err)
						return 0, err
					}
				}

			}
			if recordtypeSeq == 1 {
				//For Resolution meter status checking
				seq, err1 := dataAccess.Getemeterseqno(ClientID, Mstorgnhirarchyid, recorddiffid, 2)
				if err1 != nil {
					logger.Log.Println(err1)
					return 0, err
				}
				logger.Log.Println("Resolutionmeter sequance no ---->", seq)
				if seq > 0 {
					res, err := dataAccess.Getrecorddetails(Recordid)
					res.SupportgroupId = Usergroupid
					if err != nil {
						logger.Log.Println(err)
						return 0, err
					}
					historyrecord, err := dataAccess.GetLatesttrnhistory(ClientID, Mstorgnhirarchyid, Recordid)
					if err != nil {
						logger.Log.Println(err)
						return 0, err
					}
					if seq != 1 {
						// Change in 15.05.2021 -----------------------------

						histrn := entities.TrnslaentityhistoryEntity{}
						histrn.Clientid = ClientID
						histrn.Mstorgnhirarchyid = Mstorgnhirarchyid
						histrn.Therecordid = Recordid
						if historyrecord.Slastartstopindicator == 2 && seq == 2 {
							var dt = historyrecord.Recorddatetime
							var dttime = historyrecord.Recorddatetoint
							histrn.Recorddatetime = dt
							histrn.Recorddatetoint = dttime
						} else {
							histrn.Recorddatetime = datetime.Format("2006-01-02 15:04:05")
							histrn.Recorddatetoint = TimeParse(datetime.Format("2006-01-02 15:04:05"), "").Unix()
						}
						histrn.Slastartstopindicator = seq

						trnid, err := dataAccess.InsertTrnslaentityhistory(&histrn)
						if err != nil {
							logger.Log.Println(err)
							return 0, err
						}
						logger.Log.Println("history table id---->", trnid)

					}
					if seq == 2 {
						_, _, err, _ = GetSLAResolution(&res)
						if err != nil {
							logger.Log.Println(err)
							return 0, err
						}
					}
					if seq == 4 {
						_, err, _ := UpdateRessolutionEndFlag(ClientID, Mstorgnhirarchyid, Recordid)
						if err != nil {
							logger.Log.Println(err)
							return 0, err
						}
						_, _, err, _ = GetSLAResolution(&res)
						if err != nil {
							logger.Log.Println(err)
							return 0, err
						}
					}
					logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
					if seq == 1 || seq == 5 || seq == 3 {
						grpID, err := dataAccess.FetchCurrentGrpID(Recordid)
						if err != nil {
							return 0, err
						}
						returnValue, _, _, _ := SLACriteriaRespResl(ClientID, Mstorgnhirarchyid, res.RecordtypeID, res.WorkingcatID, res.PriorityID)
						if returnValue.Supportgroupspecific == 1 {
							count, err := dataAccess.GetSupportgrpdayofweekcount(ClientID, Mstorgnhirarchyid, grpID)
							if err != nil {
								return 0, err
							}
							if count < 7 {
								return 0, errors.New("Day Of Week Not Properly Configured.Please Check.")
							}
						} else {
							count, err := dataAccess.GetOrganizationdayofweekcount(ClientID, Mstorgnhirarchyid)
							if err != nil {
								return 0, err
							}
							if count < 7 {
								return 0, errors.New("Day Of Week Not Properly Configured.Please Check.")
							}
						}
						SLADueTimeCalculation(Recordid, 0, 1, 3, datetime.Format("2006-01-02 15:04:05"), ClientID, Mstorgnhirarchyid, res.RecordtypeID, res.WorkingcatID, res.PriorityID, "", grpID)
					}

					t := entities.SLATabEntity{}
					t.ClientID = ClientID
					t.Mstorgnhirarchyid = Mstorgnhirarchyid
					t.RecordID = Recordid
					sladtls, _, err, _ := GetSLATabvalues(&t)
					if err != nil {
						logger.Log.Println(err)
					}
					err = dao.UpdateSLAFields(tx, ClientID, Mstorgnhirarchyid, Recordid, sladtls.Responsedetails.Responseduetime, sladtls.Responsedetails.Responseclockstatus, sladtls.Resolutionetails.Resolutionduetime, sladtls.Resolutionetails.Resolutionclockstatus)
					if err != nil {
						logger.Log.Println(err)
					}

				}
			}
			//End
			return id, err
		} else {
			return 0, err
		}
	} else {
		return 0, err
	}

}
