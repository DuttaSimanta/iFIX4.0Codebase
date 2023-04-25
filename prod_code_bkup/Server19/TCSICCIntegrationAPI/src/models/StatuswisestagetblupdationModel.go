package models

import (
	"database/sql"
	"src/dao"
	"src/entities"
	"src/logger"
	"log"
	"time"
)

func UpdateResponseValueinStagetbl(tx *sql.Tx, db *sql.DB, ClientID int64, Mstorgnhirarchyid int64, RecordID int64, recorddiffseq int64, Usergroupid int64, Supportgroupspecific int64, recordtypeSeq int64) error {

	logger.Log.Println("In side UpdateResponseValueinStagetbl")

	if recordtypeSeq == 1 || recordtypeSeq == 3 {
		dataAccess := dao.DbConn{DB: db}
		value, err := dataAccess.GetFirstResponseValue(ClientID, Mstorgnhirarchyid, RecordID)
		if err != nil {
			logger.Log.Println("error is ----->", err)
			return err
		}
		if value == "NA" {
			createdt, err := dataAccess.GetRecordcreatedate(ClientID, Mstorgnhirarchyid, RecordID)
			if err != nil {
				logger.Log.Println("error is ----->", err)
				return err
			}
			currentTime := time.Now().UTC()
			today := TimeParse(currentTime.Format("2006-01-02 15:04:05"), "")
			strtDay := AddSubSecondsToDate(TimeParse(createdt, ""), 0)
			responsetiimetaken := CalculateWorkingHourBetweenTwoDates(ClientID, Mstorgnhirarchyid, strtDay, today, 0, Supportgroupspecific, Usergroupid)
			log.Println("responsetiimetaken  value is ----->", responsetiimetaken)

			err = dao.UpdateFirstResponse(tx, ClientID, Mstorgnhirarchyid, RecordID, responsetiimetaken)
			if err != nil {
				logger.Log.Println("error is ----->", err)
				return err
			}
		} else {
			if recorddiffseq == 1 || recorddiffseq == 10 || recorddiffseq == 2 {
				reopendt, err := dataAccess.GetReopendate(ClientID, Mstorgnhirarchyid, RecordID, recorddiffseq)
				if err != nil {
					logger.Log.Println("error is ----->", err)
					return err
				}
				if len(reopendt) > 0 {
					currentTime := time.Now().UTC()
					today := TimeParse(currentTime.Format("2006-01-02 15:04:05"), "")
					strtDay := AddSubSecondsToDate(TimeParse(reopendt, ""), 0)
					responsetiimetaken := CalculateWorkingHourBetweenTwoDates(ClientID, Mstorgnhirarchyid, strtDay, today, 0, Supportgroupspecific, Usergroupid)
					log.Println("responsetiimetaken  value is ----->", responsetiimetaken)

					err = dao.UpdateLatestResponse(tx, ClientID, Mstorgnhirarchyid, RecordID, responsetiimetaken)
					if err != nil {
						logger.Log.Println("error is ----->", err)
						return err
					}
				}

			}

		}
	}
	return nil

}

func UpdateResolutionValueinStagetbl(tx *sql.Tx, db *sql.DB, ClientID int64, Mstorgnhirarchyid int64, RecordID int64, recorddiffseq int64, Usergroupid int64, Supportgroupspecific int64, recordtypeSeq int64) error {
	logger.Log.Println("In side UpdateResolutionValueinStagetbl")

	if recordtypeSeq == 1 || recordtypeSeq == 3 {
		dataAccess := dao.DbConn{DB: db}
		value, err := dataAccess.GetFirstResolutionValue(ClientID, Mstorgnhirarchyid, RecordID)
		if err != nil {
			logger.Log.Println("error is ----->", err)
			return err
		}

		if value == "NA" {
			createdt, err := dataAccess.GetRecordcreatedate(ClientID, Mstorgnhirarchyid, RecordID)
			if err != nil {
				logger.Log.Println("error is ----->", err)
				return err
			}
			currentTime := time.Now().UTC()
			today := TimeParse(currentTime.Format("2006-01-02 15:04:05"), "")
			strtDay := AddSubSecondsToDate(TimeParse(createdt, ""), 0)
			resolutiontiimetaken := CalculateWorkingHourBetweenTwoDates(ClientID, Mstorgnhirarchyid, strtDay, today, 0, Supportgroupspecific, Usergroupid)
			log.Println("responsetiimetaken  value is ----->", resolutiontiimetaken)

			err = dao.UpdateFirstResolution(tx, ClientID, Mstorgnhirarchyid, RecordID, resolutiontiimetaken)
			if err != nil {
				logger.Log.Println("error is ----->", err)
				return err
			}
		} else {
			if recorddiffseq == 3 {
				var seqno int64
				if recordtypeSeq == 1 {
					seqno = 10
				} else {
					seqno = 1
				}

				reopendt, err := dataAccess.GetReopendate(ClientID, Mstorgnhirarchyid, RecordID, seqno)
				if err != nil {
					logger.Log.Println("error is ----->", err)
					return err
				}
				if len(reopendt) > 0 {
					currentTime := time.Now().UTC()
					today := TimeParse(currentTime.Format("2006-01-02 15:04:05"), "")
					strtDay := AddSubSecondsToDate(TimeParse(reopendt, ""), 0)
					resolutiontiimetaken := CalculateWorkingHourBetweenTwoDates(ClientID, Mstorgnhirarchyid, strtDay, today, 0, Supportgroupspecific, Usergroupid)
					log.Println("responsetiimetaken  value is ----->", resolutiontiimetaken)

					err = dao.UpdateLatestResolution(tx, ClientID, Mstorgnhirarchyid, RecordID, resolutiontiimetaken)
					if err != nil {
						logger.Log.Println("error is ----->", err)
						return err
					}
				}
			}
		}
	}
	return nil
}

func UpdateUserreplytimetakenValueinStagetbl(tx *sql.Tx, db *sql.DB, ClientID int64, Mstorgnhirarchyid int64, RecordID int64, recorddiffseq int64, Usergroupid int64, Supportgroupspecific int64, recordtypeSeq int64) error {
	logger.Log.Println("In side UpdateUserreplytimetakenValueinStagetbl")
	if recordtypeSeq == 1 || recordtypeSeq == 3 {
		dataAccess := dao.DbConn{DB: db}
		previousdt, err := dataAccess.GetPreviousstatusdate(ClientID, Mstorgnhirarchyid, RecordID)
		logger.Log.Println("previousdt  -----------previousdt previousdt previousdt previousdt previousdt previousdt previousdt previousdt-- ----->", previousdt)
		if err != nil {
			logger.Log.Println("error is ----->", err)
			return err
		}
		if len(previousdt) > 0 {
			currentTime := time.Now().UTC()
			today := TimeParse(currentTime.Format("2006-01-02 15:04:05"), "")
			strtDay := AddSubSecondsToDate(TimeParse(previousdt, ""), 0)
			replytimetaken := CalculateWorkingHourBetweenTwoDates(ClientID, Mstorgnhirarchyid, strtDay, today, 0, Supportgroupspecific, Usergroupid)
			err = dao.UpdateUserreplydatetime(tx, ClientID, Mstorgnhirarchyid, RecordID, replytimetaken)
			if err != nil {
				logger.Log.Println("error is ----->", err)
				return err
			}
		}

	}
	return nil
}

func UpdateFollowuptimetakenValueinStagetbl(tx *sql.Tx, db *sql.DB, ClientID int64, Mstorgnhirarchyid int64, RecordID int64, Usergroupid int64, Supportgroupspecific int64, recordtypeSeq int64) error {
	logger.Log.Println("In side UpdateFollowuptimetakenValueinStagetbl")
	if recordtypeSeq == 1 || recordtypeSeq == 3 {
		dataAccess := dao.DbConn{DB: db}
		previousdt, err := dataAccess.GetPreviousstatusdate(ClientID, Mstorgnhirarchyid, RecordID)
		if err != nil {
			log.Println("error is ----->", err)
			return err
		}
		if len(previousdt) > 0 {
			currentTime := time.Now().UTC()
			today := TimeParse(currentTime.Format("2006-01-02 15:04:05"), "")
			strtDay := AddSubSecondsToDate(TimeParse(previousdt, ""), 0)
			followuptimetaken := CalculateWorkingHourBetweenTwoDates(ClientID, Mstorgnhirarchyid, strtDay, today, 0, Supportgroupspecific, Usergroupid)
			log.Println("followuptimetaken  value is ----->", followuptimetaken)

			err = dao.UpdateFollowuptimetaken(tx, ClientID, Mstorgnhirarchyid, RecordID, followuptimetaken)
			if err != nil {
				log.Println("error is ----->", err)
				return err
			}
		}
	}
	return nil
}

func UpdateCloseValueinStagetbl(tx *sql.Tx, db *sql.DB, ClientID int64, Mstorgnhirarchyid int64, RecordID int64, recorddiffseq int64, Usergroupid int64, Supportgroupspecific int64, recordtypeSeq int64) error {

	// For close status update response resolution meter for converted ticket
	logger.Log.Println("In side UpdateCloseValueinStagetbl ---------------->", recorddiffseq, recordtypeSeq)

	if recordtypeSeq == 1 || recordtypeSeq == 3 {
		dataAccess := dao.DbConn{DB: db}

		issladue, err := dataAccess.FetchSLADueRow(ClientID, Mstorgnhirarchyid, RecordID)
		if err != nil {
			logger.Log.Println("error is ----->", err)

			return err
		}

		isresponsecomplete, isresolutioncomplete, err := dataAccess.FetchResponseResolutionCompleteValue(ClientID, Mstorgnhirarchyid, RecordID)
		if err != nil {
			logger.Log.Println("error is ----->", err)

			return err
		}
		logger.Log.Println("In side UpdateCloseValueinStagetbl -------isresolutioncomplete  --------->", isresponsecomplete, isresolutioncomplete)
		value, err := dataAccess.GetFirstResponseValue(ClientID, Mstorgnhirarchyid, RecordID)
		if err != nil {
			logger.Log.Println("error is ----->", err)

			return err
		}
		if value == "NA" {
			createdt, err := dataAccess.GetRecordcreatedate(ClientID, Mstorgnhirarchyid, RecordID)
			if err != nil {
				logger.Log.Println("error is ----->", err)
				return err
			}
			currentTime := time.Now().UTC()
			today := TimeParse(currentTime.Format("2006-01-02 15:04:05"), "")
			strtDay := AddSubSecondsToDate(TimeParse(createdt, ""), 0)
			responsetiimetaken := CalculateWorkingHourBetweenTwoDates(ClientID, Mstorgnhirarchyid, strtDay, today, 0, Supportgroupspecific, Usergroupid)
			log.Println("responsetiimetaken  value is ----->", responsetiimetaken)

			err = dao.UpdateFirstResponse(tx, ClientID, Mstorgnhirarchyid, RecordID, responsetiimetaken)
			if err != nil {
				logger.Log.Println("error is ----->", err)
				return err
			}
		} else {
			if recorddiffseq == 1 || recorddiffseq == 10 {
				reopendt, err := dataAccess.GetReopendate(ClientID, Mstorgnhirarchyid, RecordID, recorddiffseq)
				if err != nil {
					logger.Log.Println("error is ----->", err)

					return err
				}
				if len(reopendt) > 0 {
					currentTime := time.Now().UTC()
					today := TimeParse(currentTime.Format("2006-01-02 15:04:05"), "")
					strtDay := AddSubSecondsToDate(TimeParse(reopendt, ""), 0)
					responsetiimetaken := CalculateWorkingHourBetweenTwoDates(ClientID, Mstorgnhirarchyid, strtDay, today, 0, Supportgroupspecific, Usergroupid)
					log.Println("responsetiimetaken  value is ----->", responsetiimetaken)

					err = dao.UpdateLatestResponse(tx, ClientID, Mstorgnhirarchyid, RecordID, responsetiimetaken)
					if err != nil {
						log.Println("error is ----->", err)

						return err
					}
				}
			}

		}
		logger.Log.Println("In side UpdateCloseValueinStagetbl -------issladue  --------->", issladue)
		if isresponsecomplete == 0 && issladue > 0 {
			currentTime := time.Now()
			flag, err := dataAccess.UpdateResponseEndFlag(ClientID, Mstorgnhirarchyid, RecordID, currentTime.Format("2006-01-02 15:04:05"))
			logger.Log.Println(flag)
			if err != nil {
				logger.Log.Println(err)
				return err
			}

		}

		if isresolutioncomplete == 0 && issladue > 0 {
			_, err, _ := UpdateRessolutionEndFlag(ClientID, Mstorgnhirarchyid, RecordID)
			if err != nil {
				logger.Log.Println(err)
				return err
			}

			currentTime := time.Now()
			zonediff, _, _, _ := Getutcdiff(ClientID, Mstorgnhirarchyid)
			datetime := AddSubSecondsToDate(currentTime, zonediff.UTCdiff)
			histrn := entities.TrnslaentityhistoryEntity{}
			histrn.Clientid = ClientID
			histrn.Mstorgnhirarchyid = Mstorgnhirarchyid
			histrn.Therecordid = RecordID
			histrn.Recorddatetime = datetime.Format("2006-01-02 15:04:05")
			histrn.Recorddatetoint = TimeParse(datetime.Format("2006-01-02 15:04:05"), "").Unix()
			histrn.Slastartstopindicator = 4

			trnid, err := dataAccess.InsertTrnslaentityhistory(&histrn)
			if err != nil {
				logger.Log.Println(err)
				return err
			}
			logger.Log.Println("history table id---->", trnid)

		}

		value, err = dataAccess.GetFirstResolutionValue(ClientID, Mstorgnhirarchyid, RecordID)
		if err != nil {
			logger.Log.Println("error is ----->", err)

			return err
		}

		if value == "NA" {
			createdt, err := dataAccess.GetRecordcreatedate(ClientID, Mstorgnhirarchyid, RecordID)
			if err != nil {
				logger.Log.Println("error is ----->", err)
				return err
			}
			currentTime := time.Now().UTC()
			today := TimeParse(currentTime.Format("2006-01-02 15:04:05"), "")
			strtDay := AddSubSecondsToDate(TimeParse(createdt, ""), 0)
			resolutiontiimetaken := CalculateWorkingHourBetweenTwoDates(ClientID, Mstorgnhirarchyid, strtDay, today, 0, Supportgroupspecific, Usergroupid)
			log.Println("responsetiimetaken  value is ----->", resolutiontiimetaken)

			err = dao.UpdateFirstResolution(tx, ClientID, Mstorgnhirarchyid, RecordID, resolutiontiimetaken)
			if err != nil {
				logger.Log.Println("error is ----->", err)
				return err
			}
		} else {
			reopendt, err := dataAccess.GetReopendate(ClientID, Mstorgnhirarchyid, RecordID, recorddiffseq)
			if err != nil {
				logger.Log.Println("error is ----->", err)
				return err
			}
			if len(reopendt) > 0 {
				currentTime := time.Now().UTC()
				today := TimeParse(currentTime.Format("2006-01-02 15:04:05"), "")
				strtDay := AddSubSecondsToDate(TimeParse(reopendt, ""), 0)
				resolutiontiimetaken := CalculateWorkingHourBetweenTwoDates(ClientID, Mstorgnhirarchyid, strtDay, today, 0, Supportgroupspecific, Usergroupid)
				log.Println("responsetiimetaken  value is ----->", resolutiontiimetaken)

				err = dao.UpdateLatestResolution(tx, ClientID, Mstorgnhirarchyid, RecordID, resolutiontiimetaken)
				if err != nil {
					logger.Log.Println("error is ----->", err)
					return err
				}
			}
		}
	}
	// For close status update response resolution meter for converted ticket

	err := dao.UpdateCloseddate(tx, ClientID, Mstorgnhirarchyid, RecordID)
	if err != nil {
		logger.Log.Println("error is ----->", err)

		return err
	}
	return nil
}

func UpdateStageResolverInfo(tx *sql.Tx, db *sql.DB, ClientID int64, Mstorgnhirarchyid int64, RecordID int64, UserID int64, Usergroupid int64, Supportgroupspecific int64, recorddiffseq int64, recordtypeSeq int64) error {
	logger.Log.Println("In side UpdateStageResolverInfo")
	dataAccess := dao.DbConn{DB: db}

	if recordtypeSeq == 1 || recordtypeSeq == 3 {
		currentTime := time.Now()
		zonediff, _, _, _ := Getutcdiff(ClientID, Mstorgnhirarchyid)
		datetime := AddSubSecondsToDate(currentTime, zonediff.UTCdiff)
		logger.Log.Println("In side UpdateStageResolverInfo==================================================================>", ClientID, Mstorgnhirarchyid, datetime, Supportgroupspecific, Usergroupid)
		time, _ := GetSLAEndTimeForClient(ClientID, Mstorgnhirarchyid, datetime, 259200, Supportgroupspecific, Usergroupid)

		count, err := dataAccess.FetchAutoCloseRecordCount(RecordID)
		if err != nil {
			logger.Log.Println(err)
			return err
		}
		if count > 0 {
			err = dataAccess.DeleteFromClosureTable(RecordID)
			if err != nil {
				logger.Log.Println(err)
				return err
			}
		}

		err = dao.InsertRecordClosure(tx, ClientID, Mstorgnhirarchyid, RecordID, recorddiffseq, time.Format("2006-01-02 15:04:05"))
		if err != nil {
			logger.Log.Println(err)
			return err
		}
	}
	originalInfo, err := dataAccess.GetOriginalInfo(UserID)
	if err != nil {
		logger.Log.Println(err)
		return err
	}

	grpname, err := dataAccess.GetGrpname(Usergroupid)

	if err != nil {
		logger.Log.Println(err)
		return err
	}

	err = dao.UpdateStageResolver(tx, ClientID, Mstorgnhirarchyid, RecordID, UserID, originalInfo.Orgcreatorname, Usergroupid, grpname)
	if err != nil {
		logger.Log.Println("error is ----->", err)
		return err
	}
	return nil
}
