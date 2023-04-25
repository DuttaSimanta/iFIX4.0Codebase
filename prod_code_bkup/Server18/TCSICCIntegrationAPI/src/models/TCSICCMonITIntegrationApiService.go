package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	//"strings"
	"errors"
	"src/dao"
	"src/entities"
	"src/logger"
	Logger "src/logger"

	//"log"

	"os"

	//"src/resource"
	ReadProperties "src/fileutils"
	"strings"
	//SendMailUtils "src/fileutils"
)

//var lock = &sync.Mutex{}

func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		Logger.Log.Println(err)
		return false
	}
	return true
}

// func GetArcosValidate(arcosUserId string, password string, apiKey string, clientCode string,
// 	orgCode string, ticketId string, userGroup string, assignedTo string) error {
func GetPAMValidate(apiKey string, clientCode string,
	orgCode string, ticketId string, userGroup string, assignedTo string) error {

	//var aliasOrgCode = orgCode
	var clientID int64
	var orgID int64
	//var supportGroupID int64
	var recordID int64
	var assignedToID int64
	wd, err := os.Getwd() // to get working directory
	if err != nil {
		Logger.Log.Println(err)
		return errors.New("unable to Connect!!!")
	}
	//log.Println(wd)
	contextPath := strings.ReplaceAll(wd, "\\", "/") // replacing backslash by  forwardslash
	//log.Println(contextPath)
	props, err := ReadProperties.ReadPropertiesFile(contextPath + "/resource/application.properties")
	if err != nil {
		Logger.Log.Println(err)
		return errors.New("unable to Connect!!!")
	}
	//Logger.Log.Println("ARCOS ORG COde====>", orgCode)
	//orgCode = props["OrgCode"]
	// if strings.EqualFold(props["AliasOrgCode"], aliasOrgCode) {

	// 	orgCode = props["OrgCode"]
	// 	////Logger.Log.Println("Invalid Authorization. Access Denied!!!")
	// 	//return errors.New("Invalid Authorization. Access Denied!!!")
	// }
	// if !strings.EqualFold(props["AliasOrgCode"], aliasOrgCode) {
	// 	Logger.Log.Println("Invalid Authorization. Access Denied!!!")
	// 	//return errors.New("Invalid Authorization. Access Denied!!!")
	// }
	// if !strings.EqualFold(props["ArcosUserId"], arcosUserId) || !strings.EqualFold(props["ArcosPassword"], password) {
	// 	Logger.Log.Println("Invalid Authorization. Access Denied!!!")
	// 	return errors.New("Invalid Authorization. Access Denied!!!")
	// }

	if !strings.EqualFold(props["AUTH_TCS_ICC_PAM_ValidateUser"], apiKey) || !strings.EqualFold(props["ClientCode"], clientCode) {
		//Logger.Log.Println("Invalid Authorization. Access Denied!!!")
		return errors.New("Invalid Authorization. Access Denied!!!")
	}
	if db == nil {
		dbcon, err := ConnectMySqlDb()
		if err != nil {
			logger.Log.Println("Error in DBConnection")
			return errors.New("Something Went Wrong")
		}
		db = dbcon
	}
	//defer db.Close()
	var getClientIdByClientCodeQuery string = "select id from mstclient where code=?"
	getClientIdByClientCodeResulsetErr := db.QueryRow(getClientIdByClientCodeQuery, clientCode).Scan(&clientID)
	if getClientIdByClientCodeResulsetErr != nil {
		Logger.Log.Println(getClientIdByClientCodeResulsetErr)
		return errors.New("Invalid Client Code")
	}
	var getMstorgIdByOrgCodeQuery string = "select mstorgnhirarchyid from maporgcodewithtools where  toolcode='ARCOS' and orgcode=? and clientid=? and deleteflg=0 and activeflg=1"
	getMstorgIdByOrgCodeResulsetErr := db.QueryRow(getMstorgIdByOrgCodeQuery, orgCode, clientID).Scan(&orgID)
	if getMstorgIdByOrgCodeResulsetErr != nil {
		Logger.Log.Println(getMstorgIdByOrgCodeResulsetErr)
		return errors.New("Invalid Org Code")
	}
	//Logger.Log.Println("clientid===>", clientID)
	//Logger.Log.Println("orgId===>", orgID)
	//Logger.Log.Println("Tkt No===>", ticketId)

	var getTicketIDQuery string = "select id from trnrecord where clientid=? and mstorgnhirarchyid=? and code=?"
	ticketIDResultSetErr := db.QueryRow(getTicketIDQuery, clientID, orgID, ticketId).Scan(&recordID)
	if ticketIDResultSetErr != nil {
		Logger.Log.Println(ticketIDResultSetErr)
		return errors.New("Ticket No. does not exist")
	}

	var getAssignedToIDQuery string = "SELECT userid FROM mstgroupmember where userid=(select id from mstclientuser where id in(select mstuserid from mstrequest where id" +
		" in ( select mstrequestid from maprequestorecord where recordid=? and deleteflg=0 and activeflg=1)" +
		" and deleteflg=0 and activeflg=1) and loginname=?) and clientid=? and mstorgnhirarchyid=? and activeflg=1 and deleteflg=0"
	assignedToIDResultSetErr := db.QueryRow(getAssignedToIDQuery, recordID, assignedTo, clientID, orgID).Scan(&assignedToID)
	if assignedToIDResultSetErr != nil {
		Logger.Log.Println(assignedToIDResultSetErr)
		return errors.New("Ticket is not assigned to Requested Caller")
	}

	// var getSupportGroupIDQuery string = "select id from mstclientsupportgroup where id in( select mstgroupid from mstrequest where id " +
	// 	"in ( select mstrequestid from maprequestorecord where recordid=? and deleteflg=0 and activeflg=1)" +
	// 	"and deleteflg=0 and activeflg=1) and supportgroupname= ? and clientid=? and mstorgnhirarchyid=? and deleteflg=0 and activeflg=1 "
	// supportGroupResulsetErr := db.QueryRow(getSupportGroupIDQuery, recordID, userGroup, clientID, orgID).Scan(&supportGroupID)
	// if supportGroupResulsetErr != nil {
	// 	Logger.Log.Println(supportGroupResulsetErr)
	// 	return errors.New("Ticket is not assigned to the Requested Support Group ")
	// }
	var StatusName string
	var getStatusQuery string = "select name from mstrecorddifferentiation where name='Active' and id in (select recorddiffid from maprecordtorecorddifferentiation where recorddifftypeid=3  and recordid=? and clientid=? and mstorgnhirarchyid=? and islatest=1)"
	StatusGroupResulsetErr := db.QueryRow(getStatusQuery, recordID, clientID, orgID).Scan(&StatusName)
	if StatusGroupResulsetErr != nil {
		Logger.Log.Println(StatusGroupResulsetErr)
		return errors.New("Ticket is not Active")
	}
	return nil
}
func GetArcosValidate(apiKey string, clientCode string,
	orgCode string, ticketId string, userGroup string, assignedTo string) error {

	//var aliasOrgCode = orgCode
	var clientID int64
	var orgID int64
	//var supportGroupID int64
	var recordID int64
	var assignedToID int64
	wd, err := os.Getwd() // to get working directory
	if err != nil {
		Logger.Log.Println(err)
		return errors.New("unable to Connect!!!")
	}
	//log.Println(wd)
	contextPath := strings.ReplaceAll(wd, "\\", "/") // replacing backslash by  forwardslash
	//log.Println(contextPath)
	props, err := ReadProperties.ReadPropertiesFile(contextPath + "/resource/application.properties")
	if err != nil {
		Logger.Log.Println(err)
		return errors.New("unable to Connect!!!")
	}
	//Logger.Log.Println("ARCOS ORG COde====>", orgCode)
	//orgCode = props["OrgCode"]
	// if strings.EqualFold(props["AliasOrgCode"], aliasOrgCode) {

	// 	orgCode = props["OrgCode"]
	// 	//Logger.Log.Println("Invalid Authorization. Access Denied!!!")
	// 	//return errors.New("Invalid Authorization. Access Denied!!!")
	// }
	// if !strings.EqualFold(props["AliasOrgCode"], aliasOrgCode) {
	// 	Logger.Log.Println("Invalid Authorization. Access Denied!!!")
	// 	//return errors.New("Invalid Authorization. Access Denied!!!")
	// }
	// if !strings.EqualFold(props["ArcosUserId"], arcosUserId) || !strings.EqualFold(props["ArcosPassword"], password) {
	// 	Logger.Log.Println("Invalid Authorization. Access Denied!!!")
	// 	return errors.New("Invalid Authorization. Access Denied!!!")
	// }

	if !strings.EqualFold(props["AUTH_TCS_ICC_Arcos_ValidateUser"], apiKey) || !strings.EqualFold(props["ClientCode"], clientCode) {
		Logger.Log.Println("Invalid Authorization. Access Denied!!!")
		return errors.New("Invalid Authorization. Access Denied!!!")
	}
	if db == nil {
		dbcon, err := ConnectMySqlDb()
		if err != nil {
			logger.Log.Println("Error in DBConnection")
			return errors.New("Something Went Wrong")
		}
		db = dbcon
	}
	//defer db.Close()
	var getClientIdByClientCodeQuery string = "select id from mstclient where code=?"
	clientRow, getClientIdByClientCodeResulsetErr := db.Query(getClientIdByClientCodeQuery, clientCode)
	if getClientIdByClientCodeResulsetErr != nil {
		Logger.Log.Println(getClientIdByClientCodeResulsetErr)
		return errors.New("Invalid Client Code")
	}
	defer clientRow.Close()
	for clientRow.Next() {
		scanerr := clientRow.Scan(&clientID)
		if scanerr != nil {
			Logger.Log.Println(scanerr)
			return errors.New("Invalid Client Code")
		}
	}
	var getMstorgIdByOrgCodeQuery string = "select mstorgnhirarchyid from maporgcodewithtools where  toolcode='ARCOS' and orgcode=? and clientid=? and deleteflg=0 and activeflg=1"
	orgRow, getMstorgIdByOrgCodeResulsetErr := db.Query(getMstorgIdByOrgCodeQuery, orgCode, clientID)
	if getMstorgIdByOrgCodeResulsetErr != nil {
		Logger.Log.Println(getMstorgIdByOrgCodeResulsetErr)
		return errors.New("Invalid Org Code")
	}

	defer orgRow.Close()
	for orgRow.Next() {
		scanerr := orgRow.Scan(&orgID)
		if scanerr != nil {
			Logger.Log.Println(scanerr)
			return errors.New("Invalid Org Code")
		}
	}
	//Logger.Log.Println("clientid===>", clientID)
	//Logger.Log.Println("orgId===>", orgID)
	//Logger.Log.Println("Tkt No===>", ticketId)

	var getTicketIDQuery string = "select id from trnrecord where clientid=? and mstorgnhirarchyid=? and code=?"
	ticketIDResultSetErr := db.QueryRow(getTicketIDQuery, clientID, orgID, ticketId).Scan(&recordID)
	if ticketIDResultSetErr != nil {
		Logger.Log.Println(ticketIDResultSetErr)
		return errors.New("Ticket No. does not exist")
	}

	var getAssignedToIDQuery string = "SELECT userid FROM mstgroupmember where userid=(select id from mstclientuser where id in(select mstuserid from mstrequest where id" +
		" in ( select mstrequestid from maprequestorecord where recordid=? and deleteflg=0 and activeflg=1)" +
		" and deleteflg=0 and activeflg=1) and loginname=?) and clientid=? and mstorgnhirarchyid=? and activeflg=1 and deleteflg=0"
	assignedToIDResultSetErr := db.QueryRow(getAssignedToIDQuery, recordID, assignedTo, clientID, orgID).Scan(&assignedToID)
	if assignedToIDResultSetErr != nil {
		Logger.Log.Println(assignedToIDResultSetErr)
		return errors.New("Ticket is not assigned to Requested Caller")
	}

	// var getSupportGroupIDQuery string = "select id from mstclientsupportgroup where id in( select mstgroupid from mstrequest where id " +
	// 	"in ( select mstrequestid from maprequestorecord where recordid=? and deleteflg=0 and activeflg=1)" +
	// 	"and deleteflg=0 and activeflg=1) and supportgroupname= ? and clientid=? and mstorgnhirarchyid=? and deleteflg=0 and activeflg=1 "
	// supportGroupResulsetErr := db.QueryRow(getSupportGroupIDQuery, recordID, userGroup, clientID, orgID).Scan(&supportGroupID)
	// if supportGroupResulsetErr != nil {
	// 	Logger.Log.Println(supportGroupResulsetErr)
	// 	return errors.New("Ticket is not assigned to the Requested Support Group ")
	// }
	var StatusName string
	var getStatusQuery string = "select name from mstrecorddifferentiation where name='Active' and id in (select recorddiffid from maprecordtorecorddifferentiation where recorddifftypeid=3  and recordid=? and clientid=? and mstorgnhirarchyid=? and islatest=1)"
	StatusGroupResulsetErr := db.QueryRow(getStatusQuery, recordID, clientID, orgID).Scan(&StatusName)
	if StatusGroupResulsetErr != nil {
		Logger.Log.Println(StatusGroupResulsetErr)
		return errors.New("Ticket is not Active")
	}
	return nil
}
func GetTicketStatus(requestData map[string]interface{}) (string, error) {
	Logger.Log.Println("Inside Get ticket status model.................................................")
	var status string
	var clientCode = requestData["client_code"].(string)
	var orgCode string = requestData["org_code"].(string)
	var callerID string = requestData["caller_id"].(string)
	var ticketType string = requestData["ticket_type"].(string)
	var ticketNo string = requestData["ticket_no"].(string)
	//var aliasOrgCode = orgCode

	var clientID int64
	var orgID int64
	var ticketTypeID int64
	//for generic url setings
	wd, err := os.Getwd() // to get working directory
	if err != nil {
		Logger.Log.Println(err)
		return status, errors.New("unable to Connect!!!")
	}
	//log.Println(wd)
	contextPath := strings.ReplaceAll(wd, "\\", "/") // replacing backslash by  forwardslash
	//log.Println(contextPath)
	props, err := ReadProperties.ReadPropertiesFile(contextPath + "/resource/application.properties")
	if err != nil {
		Logger.Log.Println(err)
		return status, errors.New("unable to Connect!!!")
	}
	// if strings.EqualFold(props["AliasOrgCode"], aliasOrgCode) {

	// 	orgCode = props["OrgCode"]
	// 	//Logger.Log.Println("Invalid Authorization. Access Denied!!!")
	// 	//return errors.New("Invalid Authorization. Access Denied!!!")
	// }
	// orgCode = props["OrgCode"]
	// if !strings.EqualFold(props["AliasOrgCode"], aliasOrgCode) {
	// 	Logger.Log.Println("Invalid Authorization. Access Denied!!!")
	// 	return status, errors.New("Invalid Authorization. Access Denied!!!")
	// }
	if !strings.EqualFold(props["AUTH_TCS_ICC_GetStatus"], requestData["authkey"].(string)) || !strings.EqualFold(props["ClientCode"], clientCode) {
		return status, errors.New("Invalid Authorization. Access Denied!!!")
	}

	//var Status string
	if db == nil {
		dbcon, err := ConnectMySqlDb()
		if err != nil {
			logger.Log.Println("Error in DBConnection")
			return status, errors.New("Something Went Wrong")
		}
		db = dbcon
	}
	defer Logger.Log.Println("=============================OPEN CONN==================================>", db.Stats().OpenConnections)
	//defer db.Close()
	var getClientIdByClientCodeQuery string = "select id from mstclient where code=?"
	getClientIdByClientCodeResulsetErr := db.QueryRow(getClientIdByClientCodeQuery, clientCode).Scan(&clientID)
	if getClientIdByClientCodeResulsetErr != nil {
		Logger.Log.Println(getClientIdByClientCodeResulsetErr)
		return status, errors.New("Invalid Client Code")
	}
	//Logger.Log.Println("Ticket STatus ORG COde====>", orgCode)
	var getMstorgIdByOrgCodeQuery string = "select mstorgnhirarchyid from maporgcodewithtools where  toolcode='MONIT' and orgcode=? and clientid=? and deleteflg=0 and activeflg=1"
	getMstorgIdByOrgCodeResulsetErr := db.QueryRow(getMstorgIdByOrgCodeQuery, orgCode, clientID).Scan(&orgID)
	if getMstorgIdByOrgCodeResulsetErr != nil {
		Logger.Log.Println(getMstorgIdByOrgCodeResulsetErr)
		return status, errors.New("Invalid Org Code")
	}
	var userId int64
	//fetching userid email and mobileno from caller id
	var getUserDetailsQuery string = "select id from mstclientuser where loginname=? and  clientid=? and mstorgnhirarchyid=?"
	userResulsetErr := db.QueryRow(getUserDetailsQuery, callerID, clientID, orgID).Scan(&userId)
	if userResulsetErr != nil {
		Logger.Log.Println(userResulsetErr)
		return status, errors.New("Invalid CallerId")
	}
	var getTicketTypeIDByTicketTypeQuery string = "select id from mstrecorddifferentiation where name=? and  recorddifftypeid in (select id from mstrecorddifferentiationtype where seqno = 1 ) and clientid=? and mstorgnhirarchyid=?"
	getrecordIDByTicketNoResultSetErr := db.QueryRow(getTicketTypeIDByTicketTypeQuery, ticketType, clientID, orgID).Scan(&ticketTypeID)
	if getrecordIDByTicketNoResultSetErr != nil {
		Logger.Log.Println(getrecordIDByTicketNoResultSetErr)
		return status, errors.New("Invalid Ticket Type")
	}
	//var getStatusByTicketNoQuery string = "select recorddiffname as statusname from recordstatus where recordid in " +
	//	" (select id from trnrecord where code=? and clientid=? and mstorgnhirarchyid=? and activeflg=1 and deleteflg=0)"
	 var getStatusByTicketNoQuery string="select status as statusname from recordfulldetails where ticketid=? and  clientid=? and mstorgnhirarchyid=? "
	getStatusByTicketNoResultSetErr := db.QueryRow(getStatusByTicketNoQuery, ticketNo, clientID, orgID).Scan(&status)
	if getStatusByTicketNoResultSetErr != nil {
		Logger.Log.Println(getStatusByTicketNoResultSetErr)
		return status, errors.New("Ticket No. does not exist")
	}
	return status, nil
}

func MonITIntegrationApiServiceMethod(requestData map[string]interface{}) (string, error) {

	currentTime := time.Now()
	Logger.Log.Println("Current Time in milis : ", currentTime.Format("2006.01.02 15:04:05 .999"))

	//Logger.Log.Println("MonITIntegrationApiServiceMethod")
	//var clientName string
	var clientCode = requestData["client_code"].(string)
	var orgCode string = requestData["org_code"].(string)
	var callerId string = requestData["caller_id"].(string)
	var recordEntity entities.RecordEntity
	var ticketNo string
	//var aliasOrgCode string = orgCode

	//for generic url setings
	wd, err := os.Getwd() // to get working directory
	if err != nil {
		Logger.Log.Println(err)
	}
	//log.Println(wd)
	contextPath := strings.ReplaceAll(wd, "\\", "/") // replacing backslash by  forwardslash
	//log.Println(contextPath)
	props, err := ReadProperties.ReadPropertiesFile(contextPath + "/resource/application.properties")
	if err != nil {
		Logger.Log.Println(err)
	}
	// if strings.EqualFold(props["AliasOrgCode"], aliasOrgCode) {

	// 	orgCode = props["OrgCode"]
	// 	//Logger.Log.Println("Invalid Authorization. Access Denied!!!")
	// 	//return errors.New("Invalid Authorization. Access Denied!!!")
	// }
	// orgCode = props["OrgCode"]
	// if !strings.EqualFold(props["AliasOrgCode"], aliasOrgCode) {
	// 	Logger.Log.Println("Invalid Authorization. Access Denied!!!")
	// 	return ticketID, errors.New("Invalid Authorization. Access Denied!!!")
	// }
	if !strings.EqualFold(props["AUTH_TCS_ICC_MonIT_CreateTicket"], requestData["authkey"].(string)) || !strings.EqualFold(props["ClientCode"], requestData["client_code"].(string)) {
		Logger.Log.Println("Invalid Authorization. Access Denied!!!")

		return ticketNo, errors.New("Invalid Authorization. Access Denied!!!")
	}

	if db == nil {
		dbcon, err := ConnectMySqlDb()
		if err != nil {
			logger.Log.Println("Error in DBConnection")
			return ticketNo, errors.New("Something Went Wrong")
		}
		db = dbcon
	}
	Logger.Log.Println("<<<<<<<<<<<<<<<<<<<<<<<=========================OPEN CONN======================>>>>>>>>>>>>>", db.Stats().OpenConnections)
	//defer db.Close()

	recordEntity.ClientID, err = dao.GetClientID(db, clientCode)
	if err != nil {
		Logger.Log.Println(err)
		return ticketNo, errors.New("Invalid Client Code")
	}

	//Logger.Log.Println("MonIT ORG COde====>", orgCode)
	//get clientId and mstorgnhirarchyid by name of requestData["orgname"]
	recordEntity.Mstorgnhirarchyid, err = dao.GetOrgID(db, recordEntity.ClientID, orgCode)
	if err != nil {
		Logger.Log.Println(err)
		return ticketNo, errors.New("Invalid Org Code")
	}
	//Logger.Log.Println("MonIT ORG ID====>", recordEntity.Mstorgnhirarchyid)

	recordEntity.Recordname = requestData["short_description"].(string)
	recordEntity.Recordesc = requestData["long_description"].(string)

	//fetching userid email and mobileno from caller id
	recordEntity.CreateduserID, recordEntity.Requestername, recordEntity.Requesteremail, recordEntity.Requestermobile, err = dao.GetUserDetails(db, recordEntity.ClientID, recordEntity.Mstorgnhirarchyid, callerId)
	if err != nil {
		Logger.Log.Println(err)
		return ticketNo, errors.New("Invalid CallerId")
	}
	recordEntity.Originaluserid = recordEntity.CreateduserID
	// fetching group details

	recordEntity.CreatedusergroupID, err = dao.GetUserGrpID(db, recordEntity.CreateduserID)
	if err != nil {
		Logger.Log.Println(err)
		return ticketNo, errors.New("Invalid CallerId")
	}

	recordEntity.Originalusergroupid = recordEntity.CreatedusergroupID
	recordEntity.Requesterlocation = "NA"
	var recordSets []entities.RecordSet
	var recordSet entities.RecordSet
	var catType entities.RecordData
	var categories []entities.RecordData
	//Logger.Log.Println("MonITIntegrationApiServiceMethod1")

	//get cat lvl 1 details
	var catflag int64 = 0
	var parentID int64 = 0
	var catVal string
	//time.Sleep(10 * time.Millisecond)
	catType, err = dao.GetCatDetails(db, recordEntity.ClientID, recordEntity.Mstorgnhirarchyid, catflag, parentID, catVal)
	if err != nil {
		Logger.Log.Println(err)
		return ticketNo, errors.New("Invalid category1 details")
	}
	categories = append(categories, catType)
	catflag = 1
	//Logger.Log.Println("MonITIntegrationApiServiceMethod2")
	//get cat lvl 2 details
	var cat1value string
	if strings.EqualFold(requestData["category1"].(string), "") {
		cat1value = props["Category1ForOthers"]
	} else {
		cat1value = requestData["category1"].(string)
	}
	catType, err = dao.GetCatDetails(db, recordEntity.ClientID, recordEntity.Mstorgnhirarchyid, catflag, categories[0].Val, cat1value)
	if err != nil {
		Logger.Log.Println(err)
		return ticketNo, errors.New("Invalid category1 details")
	}
	categories = append(categories, catType)

	//get cat lvl 3 details
	var cat2value string
	if strings.EqualFold(requestData["category2"].(string), "") {
		cat2value = props["KeywordNotMatched"]
	} else {
		cat2value = requestData["category2"].(string)
	}
	catType, err = dao.GetCatDetails(db, recordEntity.ClientID, recordEntity.Mstorgnhirarchyid, catflag, categories[1].Val, cat2value)
	if err != nil {
		Logger.Log.Println(err)
		return ticketNo, errors.New("Invalid category1 details")
	}
	categories = append(categories, catType)
	recordEntity.Workingcatlabelid = categories[2].ID
	//get cat lvl 4 details
	var cat3value string
	if strings.EqualFold(requestData["category3"].(string), "") {
		cat3value = props["KeywordNotMatched"]
	} else {
		cat3value = requestData["category3"].(string)
	}

	catType, err = dao.GetCatDetails(db, recordEntity.ClientID, recordEntity.Mstorgnhirarchyid, catflag, categories[2].Val, cat3value)
	if err != nil {
		Logger.Log.Println(err)
		return ticketNo, errors.New("Invalid category3 details")
	}

	categories = append(categories, catType)

	//get cat lvl 5 details
	var cat4value string
	if strings.EqualFold(requestData["category4"].(string), "") {
		cat4value = props["KeywordNotMatched"]
	} else {
		cat4value = requestData["category4"].(string)
	}

	catType, err = dao.GetCatDetails(db, recordEntity.ClientID, recordEntity.Mstorgnhirarchyid, catflag, categories[3].Val, cat4value)
	if err != nil {
		Logger.Log.Println(err)
		return ticketNo, errors.New("Invalid category4 details")
	}
	categories = append(categories, catType)

	//For Issue Types Additional Field
	var foundIssueTypeValue string

	foundIssueTypeValue, err = dao.GetIssueTypeValue(db, recordEntity.ClientID, recordEntity.Mstorgnhirarchyid, recordEntity.Recordname)
	if err != nil {
		Logger.Log.Println(err)
		return ticketNo, errors.New("Invalid Issue Type")
	}
	if len(foundIssueTypeValue) == 0 {
		foundIssueTypeValue = props["KeywordNotMatched"]
	}
	recordSet.ID = 1
	recordSet.Type = categories
	recordSets = append(recordSets, recordSet)
	recordSet.Type = nil

	//get ticketType details
	recordSet, err = dao.GetTicketTypeDetails(db, recordEntity.ClientID, recordEntity.Mstorgnhirarchyid)
	if err != nil {
		Logger.Log.Println(err)
		return ticketNo, errors.New("Invalid Ticket details")
	}
	recordSets = append(recordSets, recordSet)

	//get priority details
	var priorityDetails string
	if strings.EqualFold(requestData["priority"].(string), "") {
		priorityDetails = props["PriorityForOthers"]

	} else {
		priorityDetails = requestData["priority"].(string)
	}

	recordSet, err = dao.GetPriorityDetails(db, recordEntity.ClientID, recordEntity.Mstorgnhirarchyid, priorityDetails)
	if err != nil {
		Logger.Log.Println(err)
		return ticketNo, errors.New("Invalid Ticket details")
	}
	recordSets = append(recordSets, recordSet)

	//get status details
	recordSet, err = dao.GetStatusDetails(db, recordEntity.ClientID, recordEntity.Mstorgnhirarchyid)
	if err != nil {
		Logger.Log.Println(err)
		return ticketNo, errors.New("Invalid Ticket details")
	}
	recordSets = append(recordSets, recordSet)
	recordEntity.RecordSets = recordSets

	var additionalFields []entities.RecordAdditional
	var recordAdditional entities.RecordAdditional
	// get Eventid details
	recordAdditional, err = dao.GetAdditionalFieldDetails(db, recordEntity.ClientID, recordEntity.Mstorgnhirarchyid, "Event ID")
	if err != nil {
		Logger.Log.Println(err)
		return ticketNo, errors.New("Invalid Aditional field")
	}
	recordAdditional.Val = requestData["event_id"].(string)
	additionalFields = append(additionalFields, recordAdditional)
	// get Host details
	recordAdditional, err = dao.GetAdditionalFieldDetails(db, recordEntity.ClientID, recordEntity.Mstorgnhirarchyid, "Host")
	if err != nil {
		Logger.Log.Println(err)
		return ticketNo, errors.New("Invalid Aditional field")
	}
	if strings.EqualFold(requestData["host"].(string), "") {
		return ticketNo, errors.New("Host is mandatory")
	}
	recordAdditional.Val = requestData["host"].(string)
	additionalFields = append(additionalFields, recordAdditional)

	// get Location details

	recordAdditional, err = dao.GetAdditionalFieldDetails(db, recordEntity.ClientID, recordEntity.Mstorgnhirarchyid, "Location")
	if err != nil {
		Logger.Log.Println(err)
		return ticketNo, errors.New("Invalid Aditional field")
	}
	recordAdditional.Val = requestData["location"].(string)
	additionalFields = append(additionalFields, recordAdditional)

	// get Device Type details
	recordAdditional, err = dao.GetAdditionalFieldDetails(db, recordEntity.ClientID, recordEntity.Mstorgnhirarchyid, "Device Type")
	if err != nil {
		Logger.Log.Println(err)
		return ticketNo, errors.New("Invalid Aditional field")
	}
	recordAdditional.Val = requestData["device_type"].(string)
	additionalFields = append(additionalFields, recordAdditional)

	// get Trigger ID details
	recordAdditional, err = dao.GetAdditionalFieldDetails(db, recordEntity.ClientID, recordEntity.Mstorgnhirarchyid, "Trigger ID")
	if err != nil {
		Logger.Log.Println(err)
		return ticketNo, errors.New("Invalid Aditional field")
	}
	recordAdditional.Val = requestData["trigger_id"].(string)
	additionalFields = append(additionalFields, recordAdditional)

	// get Issue Type details
	recordAdditional, err = dao.GetAdditionalFieldDetails(db, recordEntity.ClientID, recordEntity.Mstorgnhirarchyid, "Issue Type")
	if err != nil {
		Logger.Log.Println(err)
		return ticketNo, errors.New("Invalid Aditional field")
	}
	recordAdditional.Val = foundIssueTypeValue
	additionalFields = append(additionalFields, recordAdditional)
	recordEntity.Additionalfields = additionalFields
	recordEntity.Source = props["SOURCE"]
	//Logger.Log.Println("recordEntity===========>", recordEntity)
	// sendData, err := json.Marshal(recordEntity)
	// if err != nil {
	// 	Logger.Log.Println(err)
	// 	return ticketNo, errors.New("Unable to marshal data")
	// }
	// Logger.Log.Println(string(sendData))

	// resp, err := http.Post(props["URL"], "application/json", bytes.NewBuffer(sendData))
	// Logger.Log.Println("Request Sent To creat record===>", resp)
	// var result map[string]interface{}
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	Logger.Log.Println(err)
	// 	return ticketNo, errors.New("Unable to read response data")
	// }
	// err1 := json.Unmarshal(body, &result)
	// if err1 != nil {
	// 	Logger.Log.Println(err1)
	// 	return ticketNo, errors.New("Unable to Unmarchal data")
	// }
	// resp.Body.Close()
	//time.Sleep(10 * time.Millisecond)
	//  json.NewDecoder(resp.Body).Decode(&res)
	currentTime = time.Now()
	Logger.Log.Println("API  Before Current Time in milis : ", currentTime.Format("2006.01.02 15:04:05 .999"))
	ticketNo, ticketID, err := CreateRecordModel(&recordEntity, db)
	Logger.Log.Println(ticketID)
	Logger.Log.Println("Created Ticket No====================>", ticketNo)
	currentTime = time.Now()
	Logger.Log.Println("API  After Current Time in milis : ", currentTime.Format("2006.01.02 15:04:05 .999"))
	// if result["success"].(bool) == false {
	// 	Logger.Log.Println("getting False")
	// 	return ticketNo, errors.New("Ticket creation failed intermittently. Please try again later")
	// } else {
	// Logger.Log.Println(result["response"].(string))
	// ticketNo = result["response"].(string)

	/*ticketID := int64(result["id"].(float64))

		Logger.Log.Println("==========================>Ticket No:=======================>", ticketNo)

		Logger.Log.Println("==========================>Self Assignment Process starts:=======================>")

		var assignedUserID int64
		var assignedSupportgroupID int64
		var getUserIDQuery string = "select id from mstclientuser where loginname=? and activeflag=1 and deleteflag=0"
		Row, userResulsetErr := db.Query(getUserIDQuery, props["AutomationUser"])
		if userResulsetErr != nil {
			Logger.Log.Println(userResulsetErr)
			return ticketNo, errors.New("User Not Found For self Assignment")
		}
		for Row.Next() {
			scanerr := Row.Scan(&assignedUserID)
			if scanerr != nil {
				Logger.Log.Println(scanerr)
				return ticketNo, errors.New("User Not Found For self Assignment")
			}
		}
		Row.Close()
		time.Sleep(10 * time.Millisecond)
		var getAssignedSupportgroupID string = "select a.groupid from mstgroupmember a where a.clientid=? and a.mstorgnhirarchyid=? and" +
			" userid=? and groupid in( select mstgroupid from mstrequest where id in(select mstrequestid from maprequestorecord where" +
			" clientid=a.clientid and mstorgnhirarchyid=a.mstorgnhirarchyid and  recordid=? and activeflg=1 and deleteflg=0)) and" +
			" a.activeflg=1 and a.deleteflg=0"
		Row, getAssignedSupportResulsetErr := db.Query(getAssignedSupportgroupID, recordEntity.ClientID, recordEntity.Mstorgnhirarchyid, assignedUserID, ticketID)
		if getAssignedSupportResulsetErr != nil {
			Logger.Log.Println(getAssignedSupportResulsetErr)
			return ticketNo, errors.New("User is not mapped with supportgroup ")
		}
		for Row.Next() {
			scanerr := Row.Scan(&assignedSupportgroupID)
			if scanerr != nil {
				Logger.Log.Println(scanerr)
				return ticketNo, errors.New("User Not Found For self Assignment")
			}
		}
		Row.Close()
		time.Sleep(10 * time.Millisecond)
		// First API get Record Details
		Logger.Log.Println("===================================First API get Record Details===============================")

		var recordDetails model.GetRecordDetailsRequest
		recordDetails.ClientID = recordEntity.ClientID
		recordDetails.Mstorgnhirarchyid = recordEntity.Mstorgnhirarchyid
		recordDetails.RecordID = ticketID
		sendData, err := json.Marshal(recordDetails)
		if err != nil {
			Logger.Log.Println(err)
			return ticketNo, errors.New("Unable to marshal data")
		}
		Logger.Log.Println(" Get Record Details Request", string(sendData))

		resp, err := http.Post(props["URLGetRecordDetails"], "application/json", bytes.NewBuffer(sendData))
		Logger.Log.Println("Request Sent To creat record===>", resp)
		var recordDetailsResponeData model.RecordDetailsResponeData
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			Logger.Log.Println(err)
			return ticketNo, errors.New("Unable to read response data")
		}
		err1 := json.Unmarshal(body, &recordDetailsResponeData)
		if err1 != nil {
			Logger.Log.Println(err1)
			return ticketNo, errors.New("Unable to Unmarchal data")
		}
		resp.Body.Close()
		time.Sleep(10 * time.Millisecond)
		//  json.NewDecoder(resp.Body).Decode(&res)
		if recordDetailsResponeData.Status == false {
			Logger.Log.Println("getting False")
			return ticketNo, errors.New("Ticket Details Fetching failed intermittently. Please try again later")
		} else {

			Logger.Log.Println(recordDetailsResponeData)
			workflowID := recordDetailsResponeData.Details[0].WorkFlowDetails.WorkFlowID
			//recordStageID := recordDetailsResponeData.Details[0].RecordStageID
			clientID := recordDetailsResponeData.Details[0].Clientid
			mstorgnhirarchyID := recordDetailsResponeData.Details[0].Mstorgnhirarchyid
			cattypeID := recordDetailsResponeData.Details[0].WorkFlowDetails.CatTypeID
			catID := recordDetailsResponeData.Details[0].WorkFlowDetails.CatID

			// Second API get State Details

			Logger.Log.Println("===================================Second API get State Details===============================")

			// var stateDetailsRequest model.GetStateDetailsRequest
			// stateDetailsRequest.ClientID = clientID
			// stateDetailsRequest.Mstorgnhirarchyid = mstorgnhirarchyID
			// stateDetailsRequest.RecordID = ticketID
			// stateDetailsRequest.RecordStagedID = recordStageID
			// sendData, err := json.Marshal(stateDetailsRequest)
			// if err != nil {
			// 	Logger.Log.Println(err)
			// 	return ticketNo, errors.New("Unable to marshal data")
			// }
			// Logger.Log.Println(" Get State Details Request", string(sendData))

			// resp, err := http.Post(props["URLGetStateDetails"], "application/json", bytes.NewBuffer(sendData))
			// Logger.Log.Println("Request Sent To creat record===>", resp)
			// var stateDetailsResult map[string]interface{}
			// body, err := ioutil.ReadAll(resp.Body)
			// if err != nil {
			// 	Logger.Log.Println(err)
			// 	return ticketNo, errors.New("Unable to read response data")
			// }
			// err1 := json.Unmarshal(body, &stateDetailsResult)
			// if err1 != nil {
			// 	Logger.Log.Println(err1)
			// 	return ticketNo, errors.New("Unable to Unmarchal data")
			// }

			// //  json.NewDecoder(resp.Body).Decode(&res)
			// if stateDetailsResult["success"].(bool) == false {
			// 	Logger.Log.Println("getting False")
			// 	return ticketNo, errors.New("State Details Fetching failed intermittently. Please try again later")
			// } else {
			//Logger.Log.Println(stateDetailsResult)
			//transitionID := int64(stateDetailsResult["transitionid"].(float64))
			//currentstateID := int64(stateDetailsResult["currentstateid"].(float64))
			var currentStateID int64
			getRecordCurrentStateID := "select currentstateid from mstrequest where id in( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=?) "

			Row, getRecordCurrentStateIDError := db.Query(getRecordCurrentStateID, clientID, mstorgnhirarchyID, ticketID)
			if getRecordCurrentStateIDError != nil {
				Logger.Log.Println(getRecordCurrentStateIDError)
				return ticketNo, errors.New("Unable to fetch current state id")
			}
			for Row.Next() {
				scanerr := Row.Scan(&currentStateID)
				if scanerr != nil {
					Logger.Log.Println(scanerr)
					return ticketNo, errors.New("Unable to fetch current state id")
				}
			}
			Row.Close()
			time.Sleep(10 * time.Millisecond)
			//Third API get State By Seq
			Logger.Log.Println("===================================Third API get State By Seq===============================")
			var stateBySeqRequest model.GetStateBySeqRequest

			stateBySeqRequest.ClientID = clientID
			stateBySeqRequest.MstorgnhirarchyID = mstorgnhirarchyID
			stateBySeqRequest.Typeseqno = 2
			stateBySeqRequest.SeqNo = 2
			stateBySeqRequest.TransitionID = 0
			stateBySeqRequest.ProcessID = workflowID
			stateBySeqRequest.UserID = assignedUserID

			sendData, err := json.Marshal(stateBySeqRequest)
			if err != nil {
				Logger.Log.Println(err)
				return ticketNo, errors.New("Unable to marshal data")
			}
			Logger.Log.Println(" Get State By Seq Request", string(sendData))

			resp, err := http.Post(props["URLGetStateSeq"], "application/json", bytes.NewBuffer(sendData))
			Logger.Log.Println("Request Sent To creat record===>", resp)
			var stateSeqResponse = model.StateSeqResponse{}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				Logger.Log.Println(err)
				return ticketNo, errors.New("Unable to read response data")
			}
			err1 := json.Unmarshal(body, &stateSeqResponse)
			if err1 != nil {
				Logger.Log.Println(err1)
				return ticketNo, errors.New("Unable to Unmarchal data")
			}
			resp.Body.Close()
			time.Sleep(10 * time.Millisecond)
			//  json.NewDecoder(resp.Body).Decode(&res)
			if stateSeqResponse.Success == false {
				Logger.Log.Println("getting False")
				return ticketNo, errors.New("State Sequence Fetching failed intermittently. Please try again later")
			} else {
				Logger.Log.Println(stateSeqResponse.Details[0])
				mststateID := stateSeqResponse.Details[0].Mststateid

				//Forth API Move workFlow
				Logger.Log.Println("===================================Forth API Move workFlow===============================")

				var moveWorkFlowRequest model.MoveWorkFlowRequest
				moveWorkFlowRequest.ClientID = clientID
				moveWorkFlowRequest.MstorgnhirarchyID = mstorgnhirarchyID
				moveWorkFlowRequest.RecorddifftypeID = cattypeID
				moveWorkFlowRequest.RecordDiffID = catID
				moveWorkFlowRequest.TransitionID = 0
				moveWorkFlowRequest.PreviousstateID = currentStateID
				moveWorkFlowRequest.CurrentstateID = mststateID
				moveWorkFlowRequest.Manualstateselection = 0
				moveWorkFlowRequest.TransactionID = ticketID
				moveWorkFlowRequest.CreatedgroupID = assignedSupportgroupID
				moveWorkFlowRequest.Issrrequestor = 0
				moveWorkFlowRequest.UserID = assignedUserID

				sendData, err := json.Marshal(moveWorkFlowRequest)
				if err != nil {
					Logger.Log.Println(err)
					return ticketNo, errors.New("Unable to marshal data")
				}
				Logger.Log.Println(" Get Moveworkflow Request", string(sendData))

				resp, err := http.Post(props["URLMoveWorkFlow"], "application/json", bytes.NewBuffer(sendData))
				Logger.Log.Println("Request Sent To creat record===>", resp)
				var moveWorkFlowResult map[string]interface{}
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					Logger.Log.Println(err)
					return ticketNo, errors.New("Unable to read response data")
				}
				err1 := json.Unmarshal(body, &moveWorkFlowResult)
				if err1 != nil {
					Logger.Log.Println(err1)
					return ticketNo, errors.New("Unable to Unmarchal data")
				}
				resp.Body.Close()
				time.Sleep(10 * time.Millisecond)
				//  json.NewDecoder(resp.Body).Decode(&res)
				if moveWorkFlowResult["success"].(bool) == false {
					Logger.Log.Println("getting False")
					return ticketNo, errors.New("Move WorkFlow failed intermittently. Please try again later")
				} else {
					Logger.Log.Println(moveWorkFlowResult)

					//Fifth API Move ChangeSupportGroup
					Logger.Log.Println("===================================Fifth API Move ChangeSupportGroup===============================")
					var changeRecordGroupRequest model.ChangeRecordGroupRequest

					changeRecordGroupRequest.CreatedgroupID = assignedSupportgroupID
					changeRecordGroupRequest.MstgroupID = assignedSupportgroupID
					changeRecordGroupRequest.MstuserID = assignedUserID
					changeRecordGroupRequest.Samegroup = true
					changeRecordGroupRequest.TransactionID = ticketID
					changeRecordGroupRequest.UserID = assignedUserID
					sendData, err := json.Marshal(changeRecordGroupRequest)
					if err != nil {
						Logger.Log.Println(err)
						return ticketNo, errors.New("Unable to marshal data")
					}
					Logger.Log.Println(" changeRecordGroup Request", string(sendData))

					resp, err := http.Post(props["URLChangeRecordGroup"], "application/json", bytes.NewBuffer(sendData))
					Logger.Log.Println("Request Sent To creat record===>", resp)
					var changeRecordGroupResult map[string]interface{}
					body, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						Logger.Log.Println(err)
						return ticketNo, errors.New("Unable to read response data")
					}
					err1 := json.Unmarshal(body, &changeRecordGroupResult)
					if err1 != nil {
						Logger.Log.Println(err1)
						return ticketNo, errors.New("Unable to Unmarchal data")
					}
					resp.Body.Close()
					time.Sleep(10 * time.Millisecond)
					//  json.NewDecoder(resp.Body).Decode(&res)
					if changeRecordGroupResult["success"].(bool) == false {
						Logger.Log.Println("getting False")
						return ticketNo, errors.New("changeRecordGroup failed intermittently. Please try again later")
					} else {
						Logger.Log.Println(changeRecordGroupResult)
						Logger.Log.Println("======================Self Assignment Done Successfully=============>")

					}

				}

			}

		}


	}*/
	Logger.Log.Println()

	return ticketNo, nil
}
