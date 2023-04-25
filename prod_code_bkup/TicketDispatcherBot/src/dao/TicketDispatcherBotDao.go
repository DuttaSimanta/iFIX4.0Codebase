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
