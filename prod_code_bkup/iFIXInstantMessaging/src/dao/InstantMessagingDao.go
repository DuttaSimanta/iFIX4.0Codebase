package dao

import (
	"database/sql"
	"errors"
	Logger "src/logger"
)

func GetInstantMessagingSMTPDetails(db *sql.DB, clientID int64, orgID int64) (string, string, string, string, error) {

	var smtpHostForNotification string
	var emailUserName string
	var emailPassword string
	var smtpPort string
	getSMTPDetailsQuery := "select credentialkey,credentialendpoint,credentialaccount,credentialpassword from mstclientcredential where clientid=? and mstorgnhirarchyid=? and credentialtypeid=5 and activeflg=1 and deleteflg=0 "
	smtpError := db.QueryRow(getSMTPDetailsQuery, clientID, orgID).Scan(&smtpHostForNotification, &smtpPort, &emailUserName, &emailPassword)
	if smtpError != nil {
		Logger.Log.Println(smtpError)
		return smtpHostForNotification, smtpPort, emailUserName, emailPassword, errors.New("ERROR: SMTP Confu=iguration Not Found!!!")
	}
	return smtpHostForNotification, smtpPort, emailUserName, emailPassword, nil
}

func GetSRIDFromStask(db *sql.DB, clientID int64, orgID int64, recordID int64) (int64, error) {
	var sRTicketID int64

	getsTaskListQuery := "SELECT parentrecordid FROM mstparentchildmap where clientid=? and mstorgnhirarchyid=? and childrecordid=? and activeflg=1 and deleteflg=0"
	resultsetErr := db.QueryRow(getsTaskListQuery, clientID, orgID, recordID).Scan(&sRTicketID)
	if resultsetErr != nil {
		Logger.Log.Println(resultsetErr)
		return sRTicketID, errors.New("ERROR: Unable To get SR Ticket")
	}

	return sRTicketID, nil
}
func GetSTaskIDListOfSR(db *sql.DB, clientID int64, orgID int64, recordID int64) ([]int64, error) {
	var sTaskIDList []int64

	getsTaskListQuery := "SELECT childrecordid FROM mstparentchildmap where clientid=? and mstorgnhirarchyid=? and parentrecordid=? and activeflg=1 and deleteflg=0"
	resultset, resultsetErr := db.Query(getsTaskListQuery, clientID, orgID, recordID)
	if resultsetErr != nil {
		Logger.Log.Println(resultsetErr)
		return sTaskIDList, errors.New("ERROR: Unable To get STask")
	}
	defer resultset.Close()
	for resultset.Next() {
		var sTaskID int64
		scanErr := resultset.Scan(&sTaskID)
		if scanErr != nil {
			Logger.Log.Println(resultsetErr)
			return sTaskIDList, errors.New("ERROR: Unable To Scan STask")
		}
		sTaskIDList = append(sTaskIDList, sTaskID)
	}
	return sTaskIDList, nil
}
func GetTicketTypeSequenceNo(db *sql.DB, clientID int64, orgID int64, recordID int64) (int64, int64, int64, int64, error) {
	var ticketTypeSeqNo int64
	var recordDiffTypeID int64
	var recordDiffID int64
	var recordeStagedID int64
	var sql = "SELECT b.seqno,a.recorddifftypeid,a.recorddiffid,a.recordstageid  FROM maprecordtorecorddifferentiation a,mstrecorddifferentiation b WHERE a.clientid=? AND a.mstorgnhirarchyid=? AND a.recordid=? AND a.recorddifftypeid=2 AND a.islatest=1 AND a.recorddiffid=b.id"
	rowsErr := db.QueryRow(sql, clientID, orgID, recordID).Scan(&ticketTypeSeqNo, &recordDiffTypeID, &recordDiffID, &recordeStagedID)
	if rowsErr != nil {
		Logger.Log.Println(rowsErr)
		return ticketTypeSeqNo, recordDiffTypeID, recordDiffID, recordeStagedID, errors.New("ERROR: Unable to Fetch Ticket Seq No")
	}
	return ticketTypeSeqNo, recordDiffTypeID, recordDiffID, recordeStagedID, nil
}
func GetActivitySeqNo(db *sql.DB, clientID int64, orgID int64, activityName string) (int64, error) {
	var activitySeqNo int64
	var getActivitySeqNoQuery string = "select seqno from mstrecordactivitymst where clientid=? and mstorgnhirarchyid=? and activitydesc=?"
	activitySeqNoResultSetErr := db.QueryRow(getActivitySeqNoQuery, clientID, orgID, activityName).Scan(&activitySeqNo)
	if activitySeqNoResultSetErr != nil {
		Logger.Log.Println(activitySeqNoResultSetErr)
		return activitySeqNo, errors.New("ERROR: Unable to Fetch Activity Seq No")
	}
	return activitySeqNo, nil
}

func InsertMstInstantMessagingRecord(db *sql.DB, tx *sql.Tx, clientID int64, orgID int64, recordID int64, emailTo string, emailCc string, emailSub string, emailBody string) error {
	var insertMstInstantMessagingQuery = "INSERT INTO `mstrecordinstantmessaging`(`clientid`,`mstorgnhirarchyid`,`recordid`,`emailto`,`emailcc`,`emailsub`,`emailbody`,`activeflg`,`deleteflg`) VALUES(?,?,?,?,?,?,?,?,?)"

	stmt, err := tx.Prepare(insertMstInstantMessagingQuery)
	if err != nil {
		Logger.Log.Println(err)
		return errors.New("ERROR: SQL Prepare Error in mstrecordinstantmesseging")
	}
	defer stmt.Close()
	_, er := stmt.Exec(clientID, orgID, recordID, emailTo, emailCc, emailSub, emailBody, 1, 0)
	if er != nil {
		Logger.Log.Println(err)
		tx.Rollback()
		return errors.New("ERROR: SQL Execution Error in insertMstInstantMessagingQuery")
	}
	return nil
}
func InsertMstRecordActivityLogs(db *sql.DB, tx *sql.Tx, clientID int64, orgID int64, recordID int64, activitySeqNo int64, logValue string, createdID int64, createdGrpID int64) error {

	var insertlogs = "INSERT INTO mstrecordactivitylogs(clientid,mstorgnhirarchyid,recordid,activityseqno,logValue,createdid,createddate,createdgrpid) VALUES (?,?,?,?,?,?,round(UNIX_TIMESTAMP(now())),?)"
	stmt1, err := tx.Prepare(insertlogs)
	if err != nil {
		Logger.Log.Println(err)
		tx.Rollback()
		return errors.New("ERROR: SQL Prepare Error in InsertLog")
	}
	defer stmt1.Close()
	_, insertActivityLogErr := stmt1.Exec(clientID, orgID, recordID, activitySeqNo, logValue, createdID, createdGrpID)
	if insertActivityLogErr != nil {
		Logger.Log.Println(insertActivityLogErr)
		tx.Rollback()
		return errors.New("ERROR: SQL Exec Error in Insert Activity Log")
	}

	return nil
}
