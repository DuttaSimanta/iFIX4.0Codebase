package dao

import (
	"database/sql"
	"errors"
	"src/entities"
	Logger "src/logger"
)

func GetTicketList(db *sql.DB, srouceType string) (entities.TicketDetailsEntities, error) {

	var tzList = entities.TicketDetailsEntities{}

	getTicketDetailsQuery := "SELECT a.id, a.clientid, a.mstorgnhirarchyid,a.code FROM trnrecord a , maprecordtorecorddifferentiation b where a.id=b.recordid and  a.source =? and b.islatest=1 and b.recorddiffid in (select id from mstrecorddifferentiation where seqno=1 and recorddifftypeid=3 and activeflg=1 and deleteflg=0) and  a.deleteflg = 0 and a.activeflg =1"
	Resultset, resulsetErr := db.Query(getTicketDetailsQuery, srouceType)

	if resulsetErr != nil {
		Logger.Log.Println(resulsetErr)
		return tzList, errors.New("unable to fetch Ticketdetails")
	}
	defer Resultset.Close()
	for Resultset.Next() {

		var tz = entities.TicketDetailsEntity{}
		scanErr := Resultset.Scan(&tz.TicketID, &tz.ClientID, &tz.OrgID, &tz.TicketCode)
		if scanErr != nil {
			Logger.Log.Println("unable to scan Ticketdetails", scanErr)
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
