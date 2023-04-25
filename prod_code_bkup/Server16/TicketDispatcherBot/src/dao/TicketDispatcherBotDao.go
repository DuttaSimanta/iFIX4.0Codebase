package dao

import (
	"database/sql"
	"errors"
	"src/entities"
	Logger "src/logger"
)

func GetBotUserID(db *sql.DB, username string) (int64, error) {
	var botUserID int64

	fetchUserIDQuery := "select id from mstclientuser where loginname=? and deleteflag=0 and activeflag=1"
	fetchUserIDResultseterr := db.QueryRow(fetchUserIDQuery, username).Scan(&botUserID)
	if fetchUserIDResultseterr != nil {
		Logger.Log.Println(fetchUserIDResultseterr)
		//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Assignee Supportgroup")
		return botUserID, errors.New("Error: Fetching Support groupError")
	}

	return botUserID, nil
}
func GetDefaultGroupID(db *sql.DB, groupName string) (int64, error) {
	var groupID int64

	fetchUserIDQuery := "select id from mstsupportgrp where name=? and activeflg=1 and deleteflg=0"
	fetchUserIDResultseterr := db.QueryRow(fetchUserIDQuery, groupName).Scan(&groupID)
	if fetchUserIDResultseterr != nil {
		Logger.Log.Println(fetchUserIDResultseterr)
		//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Assignee Supportgroup")
		return groupID, errors.New("Error: Fetching Support groupError")
	}

	return groupID, nil
}

func GetForwardingSupportGrp(db *sql.DB, tz *entities.TicketDispatchEntity) (int64, error) {
	var forwardingSupportGrpID int64
	forwardingSupportGrpIDQuery := "select groupid from mapuserwithgroupandcategory where clientid=? and mstorgnhirarchyid=? and userid=? and categoryid=? and recorddiffid=? and activeflg=1 and deleteflg=0"
	forwardingSupportGrpIDResultseterr := db.QueryRow(forwardingSupportGrpIDQuery, tz.ClientID, tz.OrgID, tz.MstUserID, tz.WorkingID, tz.TicketTypeID).Scan(&forwardingSupportGrpID)
	if forwardingSupportGrpIDResultseterr != nil {
		Logger.Log.Println(forwardingSupportGrpIDResultseterr)
		//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Assignee Supportgroup")
		return forwardingSupportGrpID, errors.New("Error: Fetching Support groupError")
	}

	return forwardingSupportGrpID, nil
}

func GetTicketTypeSeqNo(db *sql.DB, recorddiffid int64) (int64, error) {
	var ticketTypeSeq int64

	getTicketTypeSeqQuery := "select seqno from mstrecorddifferentiation where id=? and deleteflg=0 and activeflg=1"
	getTicketTypeSeqResultseterr := db.QueryRow(getTicketTypeSeqQuery, recorddiffid).Scan(&ticketTypeSeq)
	if getTicketTypeSeqResultseterr != nil {
		Logger.Log.Println(getTicketTypeSeqResultseterr)
		//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Assignee Supportgroup")
		return ticketTypeSeq, errors.New("Error: Fetching Support groupError")
	}
	return ticketTypeSeq, nil
}
func GetStaskIDList(db *sql.DB, tz *entities.TicketDispatchEntity) ([]int64, error) {

	var sTaskIDList []int64

	getsTaskListQuery := "SELECT childrecordid FROM mstparentchildmap where clientid=? and mstorgnhirarchyid=? and parentrecordid=? and activeflg=1 and deleteflg=0"
	resultset, resultsetErr := db.Query(getsTaskListQuery, tz.ClientID, tz.OrgID, tz.TicketID)
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

func GetTicketList(db *sql.DB, supportGroupName string) (entities.TicketDispatchEntities, error) {

	var tzList = entities.TicketDispatchEntities{}

	getTicketDetailsQuery := "select a.clientid ,a.mstorgnhirarchyid,a.recordid,b.mstgroupid,e.userid,c.recorddiffid,d.recorddiffid from maprequestorecord a, mstrequest b, maprecordtorecorddifferentiation c,maprecordtorecorddifferentiation d,trnrecord e where a.mstrequestid=b.id and b.mstgroupid in(select id from mstsupportgrp where name=? and activeflg=1 and deleteflg=0) and a.recordid=c.recordid and c.isworking=1 and c.islatest=1 and a.recordid=d.recordid  and d.recorddifftypeid=2 and d.islatest=1 and b.activeflg=1 and b.deleteflg=0 and c.activeflg=1 and c.deleteflg=0 and d.activeflg=1 and d.deleteflg=0 and a.recordid=e.id"
	Resultset, resulsetErr := db.Query(getTicketDetailsQuery, supportGroupName)

	if resulsetErr != nil {
		Logger.Log.Println(resulsetErr)
		return tzList, errors.New("unable to fetch Ticketdetails")
	}
	defer Resultset.Close()
	for Resultset.Next() {

		var tz = entities.TicketDispatchEntity{}
		scanErr := Resultset.Scan(&tz.ClientID, &tz.OrgID, &tz.TicketID, &tz.MstGroupID, &tz.MstUserID, &tz.WorkingID, &tz.TicketTypeID)
		if scanErr != nil {
			Logger.Log.Println("Unable to Scan Records for Priority", scanErr)
			return tzList, errors.New("unable to scan Ticketdetails")
		}
		tzList.Values = append(tzList.Values, tz)

	}

	return tzList, nil
}

func GetAssignedUserID(db *sql.DB, clientID int64, orgID int64, assignedUserName string) (int64, error) {
	var assignedUserID int64

	var getUserIDQuery string = "select id from mstclientuser where clientid=? and  mstorgnhirarchyid=? and  loginname=? and activeflag=1 and deleteflag=0"
	Row, userResulsetErr := db.Query(getUserIDQuery, clientID, orgID, assignedUserName)
	if userResulsetErr != nil {
		Logger.Log.Println(userResulsetErr)
		return assignedUserID, errors.New("User Not Found For self Assignment")
	}
	defer Row.Close()
	for Row.Next() {
		scanerr := Row.Scan(&assignedUserID)
		if scanerr != nil {
			Logger.Log.Println(scanerr)
			return assignedUserID, errors.New("User Not Found For self Assignment")
		}
	}

	return assignedUserID, nil
}

func GetAssignedUserGrpID(db *sql.DB, clientID int64, orgID int64, assignedUserID int64, TicketID int64) (int64, error) {
	var groupID int64

	var getAssignedSupportgroupID string = "select a.groupid from mstgroupmember a where a.clientid=? and a.mstorgnhirarchyid=? and" +
		" userid=? and groupid in( select mstgroupid from mstrequest where id in(select mstrequestid from maprequestorecord where" +
		" clientid=a.clientid and mstorgnhirarchyid=a.mstorgnhirarchyid and  recordid=? and activeflg=1 and deleteflg=0)) and" +
		" a.activeflg=1 and a.deleteflg=0"
	Row, getAssignedSupportResulsetErr := db.Query(getAssignedSupportgroupID, clientID, orgID, assignedUserID, TicketID)
	if getAssignedSupportResulsetErr != nil {
		Logger.Log.Println(getAssignedSupportResulsetErr)
		return groupID, errors.New("User is not mapped with supportgroup ")
	}
	defer Row.Close()
	for Row.Next() {
		scanerr := Row.Scan(&groupID)
		if scanerr != nil {
			Logger.Log.Println(scanerr)
			return groupID, errors.New("User Not Found For self Assignment")
		}
	}
	//Row.Close()

	return groupID, nil
}

func GetRecordCurrentStateID(db *sql.DB, clientID int64, orgID int64, TicketID int64) (int64, error) {
	var currentStateID int64
	getRecordCurrentStateID := "select currentstateid from mstrequest where id in( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=?) "

	Row, getRecordCurrentStateIDError := db.Query(getRecordCurrentStateID, clientID, orgID, TicketID)
	if getRecordCurrentStateIDError != nil {
		Logger.Log.Println(getRecordCurrentStateIDError)
		return currentStateID, errors.New("Unable to fetch current state id")
	}
	defer Row.Close()
	for Row.Next() {
		scanerr := Row.Scan(&currentStateID)
		if scanerr != nil {
			Logger.Log.Println(scanerr)
			return currentStateID, errors.New("Unable to fetch current state id")
		}
	}
	return currentStateID, nil
}
