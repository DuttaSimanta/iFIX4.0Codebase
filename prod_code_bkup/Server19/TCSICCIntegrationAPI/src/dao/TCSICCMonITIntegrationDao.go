package dao

import (
	"database/sql"
	"errors"
	"src/entities"
	Logger "src/logger"
	"strings"
)

func GetClientID(db *sql.DB, clientCode string) (int64, error) {
	//Logger.Log.Println(" in GetClientID")

	var clientID int64
	var getClientIdByClientCodeQuery string = "select id from mstclient where code=?"
	clientRow, getClientIdByClientCodeResulsetErr := db.Query(getClientIdByClientCodeQuery, clientCode)
	if getClientIdByClientCodeResulsetErr != nil {
		Logger.Log.Println(getClientIdByClientCodeResulsetErr)
		return clientID, errors.New("Invalid Client Code")
	}
	defer clientRow.Close()
	for clientRow.Next() {
		scanerr := clientRow.Scan(&clientID)
		if scanerr != nil {
			Logger.Log.Println(scanerr)
			return clientID, errors.New("Invalid Client Code")
		}
	}
	//Logger.Log.Println("ClientID=====>", clientID)
	if clientID == 0 {
		return clientID, errors.New("Invalid Client Code")
	}
	return clientID, nil
}
func GetOrgID(db *sql.DB, clientID int64, orgCode string) (int64, error) {
	var orgID int64
	var getMstorgIdByOrgCodeQuery string = "select mstorgnhirarchyid from maporgcodewithtools where clientid=? and toolcode='MONIT' and orgcode=? and  deleteflg=0 and activeflg=1"
	orgRow, getMstorgIdByOrgCodeResulsetErr := db.Query(getMstorgIdByOrgCodeQuery, clientID, orgCode)
	if getMstorgIdByOrgCodeResulsetErr != nil {
		Logger.Log.Println(getMstorgIdByOrgCodeResulsetErr)
		return orgID, errors.New("Invalid Org Code")
	}
	defer orgRow.Close()
	for orgRow.Next() {
		scanerr := orgRow.Scan(&orgID)
		if scanerr != nil {
			Logger.Log.Println(scanerr)
			return orgID, errors.New("Invalid Org Code")
		}
	}
	//Logger.Log.Println("OrgID=====>", orgID)

	if orgID == 0 {

		return orgID, errors.New("Invalid Org Code")
	}

	return orgID, nil
}
func GetUserDetails(db *sql.DB, clientID int64, orgID int64, callerID string) (int64, string, string, string, error) {
	var Createduserid int64
	var Requestername string
	var Requesteremail string
	var Requestermobile string

	var getUserDetailsQuery string = "select id,name,useremail,usermobileno from mstclientuser where clientid=? and mstorgnhirarchyid=? and loginname=?  and deleteflag = 0 and activeflag =1"
	userRow, userResulsetErr := db.Query(getUserDetailsQuery, clientID, orgID, callerID)
	if userResulsetErr != nil {
		Logger.Log.Println(userResulsetErr)
		return Createduserid, Requestername, Requesteremail, Requestermobile, errors.New("Invalid CallerId")
	}
	defer userRow.Close()
	for userRow.Next() {
		scanerr := userRow.Scan(&Createduserid, &Requestername, &Requesteremail, &Requestermobile)
		if scanerr != nil {
			Logger.Log.Println(scanerr)
			return Createduserid, Requestername, Requesteremail, Requestermobile, errors.New("Invalid CallerId")
		}
	}
	//Logger.Log.Println("Createduserid=====>", Createduserid)
	//Logger.Log.Println("Requestername=====>", Requestername)
	//Logger.Log.Println("Requesteremail=====>", Requesteremail)
	//Logger.Log.Println("Requestermobile=====>", Requestermobile)

	if Createduserid == 0 {
		return Createduserid, Requestername, Requesteremail, Requestermobile, errors.New("Invalid CallerId")
	}

	return Createduserid, Requestername, Requesteremail, Requestermobile, nil
}

func GetUserGrpID(db *sql.DB, userID int64) (int64, error) {
	var grpID int64
	var getGroupDetailsQuery string = "select groupid from mstgroupmember where userid=? and deleteflg = 0 and activeflg =1"
	grpRow, groupResulsetErr := db.Query(getGroupDetailsQuery, userID)
	if groupResulsetErr != nil {
		Logger.Log.Println(groupResulsetErr)
		return grpID, errors.New("Invalid CallerId")
	}
	defer grpRow.Close()
	for grpRow.Next() {
		scanerr := grpRow.Scan(&grpID)
		if scanerr != nil {
			Logger.Log.Println(scanerr)
			return grpID, errors.New("Invalid CallerId")
		}
	}
	//Logger.Log.Println("grpID=====>", grpID)

	if grpID == 0 {
		Logger.Log.Println(groupResulsetErr)
		return grpID, errors.New("Invalid CallerId")
	}

	return grpID, nil
}
func GetCatDetails(db *sql.DB, clientID int64, orgID int64, catflag int64, parentID int64, catVal string) (entities.RecordData, error) {

	var catType entities.RecordData

	if catflag == 0 {
		var getCatLvl1Query string = "select  recorddifftypeid,id from mstrecorddifferentiation where  recorddifftypeid in( SELECT id FROM mstrecorddifferentiationtype where parentid=1 and  id in (SELECT torecorddifftypeid FROM mstrecordtype where clientid=? and mstorgnhirarchyid=? and torecorddiffid=0 and fromrecorddiffid in(select id from mstrecorddifferentiation where clientid=? and mstorgnhirarchyid=? and recorddifftypeid=2 and deleteflg=0 and activeflg=1 and seqno=1  ) and activeflg=1 and deleteflg=0) and deleteflg=0 and activeflg=1) and activeflg=1 and deleteflg=0 and seqno=1"
		cat1Row, cat1ResulsetErr := db.Query(getCatLvl1Query, clientID, orgID, clientID, orgID)
		if cat1ResulsetErr != nil {
			Logger.Log.Println(cat1ResulsetErr)
			return catType, errors.New("Invalid category details")
		}
		defer cat1Row.Close()
		for cat1Row.Next() {
			scanerr := cat1Row.Scan(&catType.ID, &catType.Val)
			if scanerr != nil {
				Logger.Log.Println(scanerr)
				return catType, errors.New("Invalid category details")
			}
		}

	} else {
		var getCatLvl2Query string = "select recorddifftypeid,id from mstrecorddifferentiation where parentid=? and name=? and deleteflg = 0 and activeflg =1"
		catRow, cat2ResulsetErr := db.Query(getCatLvl2Query, parentID, catVal)
		if cat2ResulsetErr != nil {
			Logger.Log.Println(cat2ResulsetErr)
			return catType, errors.New("Invalid category1 details")
		}
		defer catRow.Close()
		for catRow.Next() {
			scanerr := catRow.Scan(&catType.ID, &catType.Val)
			if scanerr != nil {
				Logger.Log.Println(scanerr)
				return catType, errors.New("Invalid category1 details")
			}
		}

	}

	return catType, nil
}
func GetIssueTypeValue(db *sql.DB, clientID int64, orgID int64, shortDesc string) (string, error) {
	var foundIssueTypeValue string

	var getIssueTypeFromSearchQuery string = "select keyword,categoryvalue from mapcategorywithkeyword where clientid=? and mstorgnhirarchyid=? and deleteflg = 0 and activeflg =1"
	getIssueTypeResultset, IssueTypeFromSearchResulsetErr := db.Query(getIssueTypeFromSearchQuery, clientID, orgID)
	if IssueTypeFromSearchResulsetErr != nil {
		Logger.Log.Println(IssueTypeFromSearchResulsetErr)
		return foundIssueTypeValue, errors.New("Invalid Issue Type")
	}
	defer getIssueTypeResultset.Close()
	for getIssueTypeResultset.Next() {
		var keyword string
		var catValue string
		scanErr := getIssueTypeResultset.Scan(&keyword, &catValue)
		if scanErr != nil {
			Logger.Log.Println(scanErr)
			return foundIssueTypeValue, errors.New("Invalid Issue tYPEs")
		}
		if strings.Contains(strings.ToLower(shortDesc), strings.ToLower(keyword)) {
			foundIssueTypeValue = catValue
			break

		}

	}

	return foundIssueTypeValue, nil
}
func GetTicketTypeDetails(db *sql.DB, clientID int64, orgID int64) (entities.RecordSet, error) {
	var recordSet entities.RecordSet

	var getTicketTypeQuery string = "select recorddifftypeid,id from mstrecorddifferentiation where clientid=? and mstorgnhirarchyid=? and seqno=1 AND recorddifftypeid in (select id from mstrecorddifferentiationtype where seqno = 1)  and deleteflg=0 and activeflg=1"
	typeRow, ticketTypeResulsetErr := db.Query(getTicketTypeQuery, clientID, orgID)
	if ticketTypeResulsetErr != nil {
		Logger.Log.Println(ticketTypeResulsetErr)
		return recordSet, errors.New("Invalid Ticket details")
	}
	defer typeRow.Close()
	for typeRow.Next() {
		scanerr := typeRow.Scan(&recordSet.ID, &recordSet.Val)
		if scanerr != nil {
			Logger.Log.Println(scanerr)
			return recordSet, errors.New("Invalid Ticket details")
		}
	}

	return recordSet, nil
}

func GetPriorityDetails(db *sql.DB, clientID int64, orgID int64, priorityDetails string) (entities.RecordSet, error) {
	var recordSet entities.RecordSet
	var getPriorityChangeDetails string = "select recorddifftypeid,id from mstrecorddifferentiation where clientid=? and mstorgnhirarchyid =? and recorddifftypeid = 5 and name=?"
	priorityRow, PriorityChangeResulsetErr := db.Query(getPriorityChangeDetails, clientID, orgID, priorityDetails)
	if PriorityChangeResulsetErr != nil {
		Logger.Log.Println(PriorityChangeResulsetErr)
		return recordSet, errors.New("Invalid Priority details")
	}
	defer priorityRow.Close()
	for priorityRow.Next() {
		scanerr := priorityRow.Scan(&recordSet.ID, &recordSet.Val)
		if scanerr != nil {
			Logger.Log.Println(scanerr)
			return recordSet, errors.New("Invalid Priority details")
		}
	}

	return recordSet, nil
}
func GetStatusDetails(db *sql.DB, clientID int64, orgID int64) (entities.RecordSet, error) {
	var recordSet entities.RecordSet
	var getStatusDetails string = "select recorddifftypeid,id from mstrecorddifferentiation where recorddifftypeid in (select id from mstrecorddifferentiationtype where seqno = 2) and clientid=? and mstorgnhirarchyid=? and seqno=0"
	statusRow, statusResulsetErr := db.Query(getStatusDetails, clientID, orgID)
	if statusResulsetErr != nil {
		Logger.Log.Println(statusResulsetErr)
		return recordSet, errors.New("Invalid Status details")
	}
	defer statusRow.Close()
	for statusRow.Next() {
		scanerr := statusRow.Scan(&recordSet.ID, &recordSet.Val)
		if scanerr != nil {
			Logger.Log.Println(scanerr)
			return recordSet, errors.New("Invalid Priority details")
		}
	}

	return recordSet, nil
}

func GetAdditionalFieldDetails(db *sql.DB, clientID int64, orgID int64, termName string) (entities.RecordAdditional, error) {
	var recordAdditional entities.RecordAdditional
	var getEventidTermIdquery string = "select id from mstrecordterms where clientid =? and  mstorgnhirarchyid =? and termname=? and deleteflg=0 and activeflg=1"
	additionalFieldRow, eventidTermIdResulsetErr := db.Query(getEventidTermIdquery, clientID, orgID, termName)
	if eventidTermIdResulsetErr != nil {
		Logger.Log.Println(eventidTermIdResulsetErr)
		return recordAdditional, errors.New("Invalid Aditional field")
	}
	defer additionalFieldRow.Close()
	for additionalFieldRow.Next() {
		scanerr := additionalFieldRow.Scan(&recordAdditional.Termsid)
		if scanerr != nil {
			Logger.Log.Println(scanerr)
			return recordAdditional, errors.New("Invalid Aditional field")
		}
	}

	var getEventidIdquery string = "select id from mstrecordfield where clientid =? and mstorgnhirarchyid =? and recordtermid = ? and deleteflg=0 and activeflg=1"
	additionalFieldRowID, eventidIdResulsetErr := db.Query(getEventidIdquery, clientID, orgID, recordAdditional.Termsid)
	if eventidIdResulsetErr != nil {
		Logger.Log.Println(eventidIdResulsetErr)
		return recordAdditional, errors.New("Invalid Aditional field")
	}
	//	defer additionalFieldRow.Close()
	//time.Sleep(10* time.Millisecond)
	defer additionalFieldRowID.Close()
	for additionalFieldRowID.Next() {
		scanerr := additionalFieldRowID.Scan(&recordAdditional.ID)
		if scanerr != nil {
			Logger.Log.Println(scanerr)
			return recordAdditional, errors.New("Invalid Aditional field")
		}
	}

	return recordAdditional, nil
}
