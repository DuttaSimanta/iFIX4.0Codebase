package dao

import (
	"ifixNatsSubscriber/ifix/entities"
	"ifixNatsSubscriber/ifix/logger"
	//"database/sql"
)

// func (mdao DbConn) GetBugRecords() ([]entities.RecordDetailsEntity, error) {
// 	logger.Log.Println("In side GetTermvalueagainsttermid")
// 	values := []entities.RecordDetailsEntity{}
// 	latestrecordcomment := "SELECT a.clientid,a.mstorgnhirarchyid,a.recordid,b.id as requestid FROM iFIX.maprequestorecord a,mstrequest b,maprecordstatetodifferentiation c,maprecordtorecorddifferentiation d,trnrecord e,mstrecorddifferentiation f,mstrecorddifferentiation g,maprecordtorecorddifferentiation h,mstrecorddifferentiation i where a.mstrequestid=b.id and b.currentstateid=c.mststateid and b.clientid=c.clientid and b.mstorgnhirarchyid=c.mstorgnhirarchyid and c.recorddifftypeid=3 and a.recordid=d.recordid and d.clientid=a.clientid and d.mstorgnhirarchyid=a.mstorgnhirarchyid and  d.recorddifftypeid=3 and d.islatest=1 and c.recorddiffid<>d.recorddiffid and d.recordstageid=a.recordstageid and a.recordid=e.id and e.deleteflg=0 and a.deleteflg=0 and b.deleteflg=0 and c.deleteflg=0 and d.deleteflg=0 and g.id=c.recorddiffid and f.id=d.recorddiffid and a.recordid=h.recordid and a.clientid=h.clientid and a.mstorgnhirarchyid=h.mstorgnhirarchyid and h.recorddifftypeid=2 and h.islatest=1 and h.recorddiffid=i.id  and i.name='Incident' limit 1"
// 	rows, err := mdao.DB.Query(latestrecordcomment)
// 	defer rows.Close()
// 	if err != nil {
// 		logger.Log.Println("GetTermvalueagainsttermid Get Statement Prepare Error", err)
// 		return values, err
// 	}
// 	for rows.Next() {
// 		value := entities.RecordDetailsEntity{}
// 		rows.Scan(&value.ClientID, &value.Mstorgnhirarchyid, &value.RecordID, &value.Requestid)
// 		values = append(values, value)
// 	}
// 	logger.Log.Println(values)
// 	return values, nil
// }
func (mdao DbConn) Getuserandgroup(RecordID int64) (int64, int64, error) {
	logger.Log.Println("Parameter is ------>", RecordID)

	var userid int64
	var grpid int64

	var getparentids = "SELECT createdid,createdgrpid FROM mstrecordactivitylogs where recordid=? AND activityseqno=4 order by id desc limit 1;	"
	rows, err := mdao.DB.Query(getparentids, RecordID)
	logger.Log.Println("Rows error is  ----+++++++++++++++++++++++++++++++++++++++++-->", err)

	defer rows.Close()
	if err != nil {
		logger.Log.Println("Getchildrecordids Get Statement Prepare Error", err)
		return userid, grpid, err
	}
	for rows.Next() {
		// var ID int64
		err = rows.Scan(&userid, &grpid)
		// parentids = append(parentids, ID)
		logger.Log.Println("Getchildrecordids rows.next() Error", err)
	}
	return userid, grpid, nil
}
func (mdao DbConn) Getchildrecordids(RecordID int64) ([]int64, error) {
	logger.Log.Println("Parameter is ------>", RecordID)

	var parentids []int64
	var getparentids = "select childrecordid  from mstparentchildmap where parentrecordid=? and deleteflg=0 and activeflg=1 and parentrecordid !=0"
	rows, err := mdao.DB.Query(getparentids, RecordID)
	logger.Log.Println("Rows error is  ----+++++++++++++++++++++++++++++++++++++++++-->", err)

	defer rows.Close()
	if err != nil {
		logger.Log.Println("Getchildrecordids Get Statement Prepare Error", err)
		return parentids, err
	}
	for rows.Next() {
		var ID int64
		err = rows.Scan(&ID)
		parentids = append(parentids, ID)
		logger.Log.Println("Getchildrecordids rows.next() Error", err)
	}
	return parentids, nil
}
func (mdao DbConn) GetWorkflowState(page entities.RecordDetailsEntity) ([]entities.StatedetailEntity, error) {
	logger.Log.Println("In side GetTermvalueagainsttermid")
	values := []entities.StatedetailEntity{}
	latestrecordcomment := "select distinct b.id,b.currentstateid,b.userid,b.mstgroupid,c.recorddiffid,d.name,d.seqno,b.mstuserid from maprequestorecord a,mstrequesthistory b,maprecordstatetodifferentiation c,mstrecorddifferentiation d where a.recordid=? and a.mstrequestid=b.mainrequestid and c.mststateid=b.currentstateid and c.recorddifftypeid=3 and c.deleteflg=0 and b.mstorgnhirarchyid=c.mstorgnhirarchyid and c.recorddiffid=d.id order by b.id;"
	rows, err := mdao.DB.Query(latestrecordcomment, page.RecordID)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetTermvalueagainsttermid Get Statement Prepare Error", err)
		return values, err
	}
	count := 0
	for rows.Next() {
		value := entities.StatedetailEntity{}
		rows.Scan(&value.Id, &value.Stateid, &value.Userid, &value.Createdgroupid, &value.Statusid, &value.Statusname, &value.StatusSeq, &value.Mstuserid)
		value.Index = count
		count++
		values = append(values, value)
	}
	logger.Log.Println(values)
	return values, nil
}
func (mdao DbConn) GetRecordState(page entities.RecordDetailsEntity) ([]entities.StatedetailEntity, error) {
	logger.Log.Println("In side GetTermvalueagainsttermid")
	values := []entities.StatedetailEntity{}
	latestrecordcomment := " select b.id,b.recorddiffid,a.name,a.seqno from maprecordtorecorddifferentiation b,mstrecorddifferentiation a where b.recorddifftypeid=3   and b.recordid=? and a.id=b.recorddiffid order by b.id;	"
	rows, err := mdao.DB.Query(latestrecordcomment, page.RecordID)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetTermvalueagainsttermid Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.StatedetailEntity{}
		rows.Scan(&value.Id, &value.Statusid, &value.Statusname, &value.StatusSeq)
		values = append(values, value)
	}
	logger.Log.Println(values)
	return values, nil
}
func (mdao DbConn) GetTaskTicket(page entities.RecordDetailsEntity) ([]int64, error) {
	logger.Log.Println("In side GetTermvalueagainsttermid")
	values := []int64{}
	latestrecordcomment := "select  distinct childrecordid from mstparentchildmap where parentrecordid=?	"
	rows, err := mdao.DB.Query(latestrecordcomment, page.RecordID)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetTermvalueagainsttermid Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		var value int64
		rows.Scan(&value)
		values = append(values, value)
	}
	logger.Log.Println(values)
	return values, nil
}
func (dbc DbConn) InsertCommentOrErr(tz *entities.ErrorEntity) (int64, error) {
	// logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>history table id----------22222222222222222222222222222222------------------------------------------>", tz)
	// logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>history table id----------22222222222222222222222222222222------------------------------------------>", tz.Slastartstopindicator)
	sql := "INSERT INTO defectticketlog (clientid, mstorgnhirarchyid, recordid, stateid, requestjson, responsejson, comment, isdefect) VALUES (?,?,?,?,?,?,?,?)"
	stmt, err := dbc.DB.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("Prepare Statement>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>history table id----------22222222222222222222222222222222------------>", err)
		return 0, err
	}
	res, err := stmt.Exec(tz.ClientID, tz.Mstorgnhirarchyid, tz.RecordID, tz.Stateid, tz.Requestjson, tz.Responsejson, tz.Comment, tz.Isdefect)
	if err != nil {
		logger.Log.Println("Exec Error>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>history table id----------22222222222222222222222222222222------------>", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

// var getMstslafullfillmentcriteria = `SELECT id as Id, clientid as Clientid, mstorgnhirarchyid as Mstorgnhirarchyid,
// slaid as Slaid, mstrecorddifferentiationtickettypeid as Mstrecorddifferentiationtickettypeid,
// mstrecorddifferentiationpriorityid as Mstrecorddifferentiationpriorityid, mstrecorddifferentiationworkingcatid as Mstrecorddifferentiationworkingcatid,
// responsetimeinhr as Responsetimeinhr, responsetimeinmin as Responsetimeinmin, responsetimeinsec as Responsetimeinsec,
// resolutiontimeinhr as Resolutiontimeinhr, resolutiontimeinmin as Resolutiontimeinmin, resolutiontimeinsec as Resolutiontimeinsec,
// supportgroupspecific as Supportgroupspecific, activeflg as Activeflg
// FROM mstslafullfillmentcriteria
// WHERE clientid = ? AND mstorgnhirarchyid = ?  AND mstrecorddifferentiationtickettypeid =?
// AND mstrecorddifferentiationpriorityid =? AND mstrecorddifferentiationworkingcatid =?
// AND deleteflg =0 and activeflg=1`

// var getsupportGroupid = `SELECT mstclientsupportgroupid FROM mstslaresponsiblesupportgroup WHERE clientid = ? AND mstorgnhirarchyid = ? AND mstslafullfillmentcriteriaid = ? AND activeflg = 1 AND deleteflg = 0`

// var getclientsupportgroupholyday = `SELECT starttimeinteger, endtimeinteger, FROM_UNIXTIME(dateofholiday) as dateofholiday FROM mstclientsupportgroupholiday WHERE clientid = ? AND mstorgnhirarchyid = ? AND mstclientsupportgroupid = ? AND dateofholiday = round(UNIX_TIMESTAMP(?)) AND deleteflg = 0 AND activeflg = 1`

// var getclientholyday = `SELECT starttimeinteger, endtimeinteger, FROM_UNIXTIME(dateofholiday) as dateofholiday FROM mstclientholiday WHERE clientid = ? AND mstorgnhirarchyid = ? AND dateofholiday = round(UNIX_TIMESTAMP(?)) AND deleteflg = 0 AND activeflg = 1`

// var getclientWeekDay = `SELECT starttimeinteger, endtimeinteger FROM mstclientdayofweek WHERE clientid = ? AND mstorgnhirarchyid = ?  AND dayofweekid = ? AND deleteflg = 0 AND activeflg = 1`

// var getsupportgroupWeekDay = `SELECT starttimeinteger, endtimeinteger, nextdayforward FROM mstclientsupportgroupdayofweek WHERE clientid = ? AND mstorgnhirarchyid = ? AND mstclientsupportgroupid = ?  AND dayofweekid = ? AND deleteflg = 0 AND activeflg = 1`

// var insertTrnslaentityhistory = "INSERT INTO trnslaentityhistory (clientid, mstorgnhirarchyid, mstslaentityid, therecordid, recorddatetime, recorddatetoint, donotupdatesladue, recordtimetoint, mstslastateid, commentonrecord, slastartstopindicator, fromclientuserid) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)"

// var getTrnslaentityhistory = `SELECT id, clientid, mstorgnhirarchyid, mstslaentityid, therecordid, recorddatetime, recorddatetoint, donotupdatesladue, recordtimetoint, mstslastateid, commentonrecord, slastartstopindicator, fromclientuserid FROM trnslaentityhistory WHERE clientid = ? AND mstorgnhirarchyid = ? AND therecordid = ? AND activeflg = 1 AND deleteflg = 0 ORDER BY id DESC LIMIT 1`

// //var updateRemainingPercent = `UPDATE mstsladue SET remainingtime = ?, completepercent = ? WHERE therecordid = ? AND clientid = ? AND mstorgnhirarchyid = ?`

// //var updateMstsladue = `UPDATE mstsladue SET duedatetimeresponse = ?, duedatetimeresolution = ?, duedatetimeresolutionint = ?, duedatetimeresponseint = ?,pushtime=? WHERE therecordid = ? AND clientid = ? AND mstorgnhirarchyid = ?`
// var updateMstsladue = "UPDATE mstsladue SET duedatetimeresponse = ?, duedatetimeresolution = ?, duedatetimeresolutionint = ?, duedatetimeresponseint = ?,pushtime=? WHERE therecordid = ? AND clientid = ? AND mstorgnhirarchyid = ? order by id desc limit 1"
// var insertMstsladue = "INSERT INTO mstsladue (clientid, mstorgnhirarchyid, mstslaentityid, therecordid, latestone, startdatetimeresponse, startdatetimeresolution, duedatetimeresponse, duedatetimeresolution, duedatetimetominute, resoltiondone, resolutiondatetime, lastupdatedattime, trnslaentityhistoryid, duedatetimeresolutionint, duedatetimeresponseint,remainingtime, completepercent, responseremainingtime, responsepercentage) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

// var updateRemainingPercent = `UPDATE mstsladue SET remainingtime = ?, completepercent = ?, responseremainingtime = ?, responsepercentage = ? WHERE therecordid = ? AND clientid = ? AND mstorgnhirarchyid = ?`

// //var getMstsladue = `SELECT id, clientid, mstorgnhirarchyid, mstslaentityid, therecordid, latestone, startdatetimeresponse, startdatetimeresolution, duedatetimeresponse, duedatetimeresolution, duedatetimeresolutionint, duedatetimeresponseint, resoltiondone, resolutiondatetime, trnslaentityhistoryid, remainingtime, completepercent, responseremainingtime, responsepercentage, isresponsecomplete FROM mstsladue WHERE therecordid = ? AND clientid = ? AND mstorgnhirarchyid = ? AND activeflg = 1 AND deleteflg = 0 ORDER BY id DESC LIMIT 1`

// var updateResponseEndFlag = `UPDATE mstsladue SET isresponsecomplete = 1, responseCompleteTime = ? WHERE therecordid = ? AND clientid = ? AND mstorgnhirarchyid = ?`

// var getutcdiff = `SELECT b.utcdiff FROM mstslatimezone as a, zone as b WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.msttimezoneid = b.zone_id AND a.deleteflg = 0 AND a.activeflg = 1`

// //var getMstsladue = `SELECT id, clientid, mstorgnhirarchyid, mstslaentityid, therecordid, latestone, startdatetimeresponse, startdatetimeresolution, duedatetimeresponse, duedatetimeresolution, duedatetimeresolutionint, duedatetimeresponseint, resoltiondone, resolutiondatetime, trnslaentityhistoryid, remainingtime, completepercent, responseremainingtime, responsepercentage, isresponsecomplete, isresolutioncomplete FROM mstsladue WHERE therecordid = ? AND clientid = ? AND mstorgnhirarchyid = ? AND activeflg = 1 AND deleteflg = 0 ORDER BY id DESC LIMIT 1`

// //var updateRespViolateFlag = `UPDATE mstsladue SET isresponsecomplete = 2 WHERE therecordid = ? AND clientid = ? AND mstorgnhirarchyid = ?`

// var updateResolViolateFlagFullFillDtl = `UPDATE recordfulldetails SET resolslabreachstatus = 'yes' WHERE recordid = ? AND clientid = ? AND mstorgnhirarchyid = ?`
// var updateResplViolateFlagFullFillDtl = `UPDATE recordfulldetails SET respslabreachstatus = 'yes' WHERE recordid = ? AND clientid = ? AND mstorgnhirarchyid = ?`

// var updateRespViolateFlag = `UPDATE mstsladue SET isresponseviolation = 1 WHERE therecordid = ? AND clientid = ? AND mstorgnhirarchyid = ?`

// var updateResolViolateFlag = `UPDATE mstsladue SET isresolutionviolation = 1 WHERE therecordid = ? AND clientid = ? AND mstorgnhirarchyid = ?`

// var updateRessolutionEndFlag = `UPDATE mstsladue SET isresolutioncomplete = 1, resolutionCompleteTime = ? WHERE therecordid = ? AND clientid = ? AND mstorgnhirarchyid = ?`

// var getMstslafullfillmentcriteriawithoutworkingcategory = `SELECT id as Id, clientid as Clientid, mstorgnhirarchyid as Mstorgnhirarchyid,
// slaid as Slaid, mstrecorddifferentiationtickettypeid as Mstrecorddifferentiationtickettypeid,
// mstrecorddifferentiationpriorityid as Mstrecorddifferentiationpriorityid, mstrecorddifferentiationworkingcatid as Mstrecorddifferentiationworkingcatid,
// responsetimeinhr as Responsetimeinhr, responsetimeinmin as Responsetimeinmin, responsetimeinsec as Responsetimeinsec,
// resolutiontimeinhr as Resolutiontimeinhr, resolutiontimeinmin as Resolutiontimeinmin, resolutiontimeinsec as Resolutiontimeinsec,
// supportgroupspecific as Supportgroupspecific, activeflg as Activeflg
// FROM mstslafullfillmentcriteria
// WHERE clientid = ? AND mstorgnhirarchyid = ?  AND mstrecorddifferentiationtickettypeid =?
// AND mstrecorddifferentiationpriorityid =? AND deleteflg =0 and activeflg=1`

// var getMstsladue = `SELECT id, clientid, mstorgnhirarchyid, mstslaentityid, therecordid, latestone, startdatetimeresponse, startdatetimeresolution, duedatetimeresponse, duedatetimeresolution, duedatetimeresolutionint, duedatetimeresponseint, resoltiondone, resolutiondatetime, trnslaentityhistoryid, remainingtime, completepercent, responseremainingtime, responsepercentage, isresponsecomplete, isresolutioncomplete, pushtime FROM mstsladue WHERE therecordid = ? AND clientid = ? AND mstorgnhirarchyid = ? AND activeflg = 1 AND deleteflg = 0 ORDER BY id DESC LIMIT 1`

// //var getTrnslaentityhistorytype2 = `SELECT id, clientid, mstorgnhirarchyid, mstslaentityid, therecordid, recorddatetime, recorddatetoint, donotupdatesladue, recordtimetoint, mstslastateid, commentonrecord, slastartstopindicator, fromclientuserid FROM trnslaentityhistory WHERE clientid = ? AND mstorgnhirarchyid = ? AND therecordid = ? AND activeflg = 1 AND deleteflg = 0 AND slastartstopindicator = 2 ORDER BY id DESC LIMIT 1`

// var updatePushTimeInHistory = `UPDATE trnslaentityhistory SET pushtime = ? WHERE id = ?`

// //var getTrnslaentityhistorytype2 = `SELECT id, clientid, mstorgnhirarchyid, mstslaentityid, therecordid, recorddatetime, recorddatetoint, donotupdatesladue, recordtimetoint, mstslastateid, commentonrecord, slastartstopindicator, fromclientuserid FROM trnslaentityhistory WHERE clientid = ? AND mstorgnhirarchyid = ? AND therecordid = ? AND activeflg = 1 AND deleteflg = 0 AND slastartstopindicator = 2 AND id <= ?   ORDER BY id DESC`

// var getTrnslaentityhistorytype2 = `SELECT id, clientid, mstorgnhirarchyid, mstslaentityid, therecordid, recorddatetime, recorddatetoint, donotupdatesladue, recordtimetoint, mstslastateid, commentonrecord, slastartstopindicator, fromclientuserid FROM trnslaentityhistory  WHERE clientid = ? AND mstorgnhirarchyid = ? AND therecordid = ? AND activeflg = 1 AND deleteflg = 0 AND slastartstopindicator = 2 AND id <= ? AND id > (SELECT id FROM trnslaentityhistory b WHERE id < ? and slastartstopindicator=1 and b.therecordid=trnslaentityhistory.therecordid order BY id DESC LIMIT 1) ORDER BY id DESC`

// // ****************************************** Scheduler Query Start ***************************************************************
// // b.status_id,
// var getCategoryIds = `SELECT
// a.clientid,
// a.mstorgnhirarchyid,
//  a.recordid,
//  a.ticket_type_id,
//  c.priority_id,
//  d.Working_category

// FROM
//  (SELECT
// 	 clientid,
// 		 mstorgnhirarchyid,
// 		 recordid,
// 		 recorddiffid AS ticket_type_id
//  FROM
// 	 maprecordtorecorddifferentiation m
//  WHERE
// 	 recorddifftypeid = 2 AND recordid = ?
// 		 AND islatest = 1) a,
//  (SELECT
// 		n.clientid,
// 		 n.mstorgnhirarchyid,
// 		 n.recordid,
// 		 n.recorddiffid AS status_id

//  FROM
// 	 maprecordtorecorddifferentiation n
//  WHERE
// 		 n.recorddifftypeid = 3
// 		 AND n.recordid = ?
// 		 AND n.islatest = 1) b,
//  (SELECT
// 	 o.clientid,
// 		 o.mstorgnhirarchyid,
// 		 o.recordid,
// 		 o.recorddiffid AS priority_id
//  FROM
// 	 maprecordtorecorddifferentiation o
//  WHERE
// 		 o.recorddifftypeid = 5
// 		 AND o.recordid = ?
// 		 AND o.islatest = 1) c,
// (SELECT
// 	 clientid,
// 		 mstorgnhirarchyid,
// 		 recordid,
// 		 recorddiffid AS Working_category
//  FROM
// 	 maprecordtorecorddifferentiation m
//  WHERE
// 	 isworking = 1 AND recordid = ?
// 		 AND islatest = 1) d

// WHERE
// 	 a.recordid = b.recordid
// 	 AND a.recordid = c.recordid
// 	 AND a.recordid = d.recordid
// 	 AND a.clientid = ?
// 	 AND a.mstorgnhirarchyid = ?
// 	 AND a.recordid = ?`

// var getrecordFromsladue = `SELECT id, clientid, mstorgnhirarchyid, therecordid FROM mstsladue WHERE isresolutioncomplete = 0 AND isexecute = 0  ORDER BY id DESC`

// // var getrecordFromsladue = `SELECT id, clientid, mstorgnhirarchyid, therecordid FROM mstsladue WHERE therecordid =51542`

// var updateExecuteFlag = `UPDATE mstsladue SET isexecute = ? WHERE id IN(%s)`

// // var getEventTemplateData = `SELECT id, eventparams, eventtype FROM mstnotificationtemplate WHERE clientid = ? AND mstorgnhirarchyid = ? AND recordtypeid = ? AND workingcategoryid = ? AND channeltype = 1`

// // AND JSON_CONTAINS(eventparams, ?, '$.processid')
// var checkEventProcessFlag = `SELECT processflag FROM mstnotificationlog WHERE clientid = ? AND mstorgnhirarchyid = ? AND recordid = ? AND notificationeventid = ? AND notificationtemplateid = ? AND deleteflg = 0 AND activeflg = 1`

// //===========================================================================================
// var latestrecordcomment = "SELECT a.clientid,a.mstorgnhirarchyid,a.recordid,a.recordtrackvalue,FROM_UNIXTIME(a.createddate) createddate FROM trnreordtracking a,mstrecordterms b WHERE a.clientid=? AND a.mstorgnhirarchyid=? AND a.recordid=? AND a.recordtermid =b.id AND b.seq=11 order by a.createddate desc limit 1"

// func (mdao DbConn) GetLastRecordcomment(page *entities.RecordcommonEntity) ([]entities.Customervisiblecomment, error) {
// 	logger.Log.Println("In side GetTermvalueagainsttermid")
// 	values := []entities.Customervisiblecomment{}
// 	rows, err := mdao.DB.Query(latestrecordcomment, page.ClientID, page.Mstorgnhirarchyid, page.RecordID)
// 	defer rows.Close()
// 	if err != nil {
// 		logger.Log.Println("GetTermvalueagainsttermid Get Statement Prepare Error", err)
// 		return values, err
// 	}
// 	for rows.Next() {
// 		value := entities.Customervisiblecomment{}
// 		rows.Scan(&value.ClientID, &value.Mstorgnhirarchyid, &value.RecordID, &value.Recordtrackvalue, &value.Createddate)
// 		values = append(values, value)
// 	}
// 	logger.Log.Println(values)
// 	return values, nil
// }

// // var latestsatusid = "SELECT a.recorddiffid,b.seqno,b.name FROM maprecordtorecorddifferentiation a,mstrecorddifferentiation b WHERE a.clientid=? AND a.mstorgnhirarchyid=? AND a.recorddifftypeid=3 AND a.recordid=? AND a.islatest=1 AND a.recorddiffid=b.id"
// // func (mdao DbConn) Getcurrentsatusid(ClientID int64, Mstorgnhirarchyid int64, RecordID int64) (int64, int64, string, error) {
// // logger.Log.Println("In side Getrecorddiffidbystateid")
// // var recorddiffid int64
// // var name string
// // var seqno int64
// // rows, err := mdao.DB.Query(latestsatusid, ClientID, Mstorgnhirarchyid, RecordID)
// // defer rows.Close()
// // if err != nil {
// // logger.Log.Println("Getcurrentsatusid Get Statement Prepare Error", err)
// // return recorddiffid, seqno, name, err
// // }
// // for rows.Next() {
// // rows.Scan(&recorddiffid, &seqno, &name)

// // }
// // return recorddiffid, seqno, name, nil
// // }

// var recordcreatedt = "SELECT FROM_UNIXTIME(createdatetime) recordcreatedate FROM trnrecord WHERE clientid=? AND mstorgnhirarchyid=? AND id=?"

// func (mdao DbConn) GetRecordcreatedate(ClientID int64, Mstorgnhirarchyid int64, Recordid int64) (string, error) {
// 	logger.Log.Println("In side GetRecordcreatedate")
// 	var createdate string
// 	rows, err := mdao.DB.Query(recordcreatedt, ClientID, Mstorgnhirarchyid, Recordid)
// 	defer rows.Close()
// 	if err != nil {
// 		logger.Log.Println("GetRecordcreatedate Get Statement Prepare Error", err)
// 		return createdate, err
// 	}
// 	for rows.Next() {
// 		err = rows.Scan(&createdate)
// 		logger.Log.Println("GetRecordcreatedate rows.next() Error", err)
// 	}
// 	return createdate, nil
// }

// func (mdao DbConn) GetReopendate(ClientID int64, Mstorgnhirarchyid int64, RecordID int64, recorddiffseq int64) (string, error) {
// 	logger.Log.Println("In side GetPreviousstatusdate")
// 	var name string
// 	var sql = "SELECT FROM_UNIXTIME(a.createddate) FROM maprecordtorecorddifferentiation a,mstrecorddifferentiation b WHERE a.clientid=? AND a.mstorgnhirarchyid=? AND a.recorddifftypeid=3 AND a.recordid=? AND b.clientid=? AND b.mstorgnhirarchyid=? AND b.recorddifftypeid=3 AND b.seqno=? AND b.activeflg=1 AND b.deleteflg=0 AND a.recorddiffid=b.id order by a.id desc limit 1"
// 	logger.Log.Println("Reopen query is --------------->", sql)
// 	logger.Log.Println("Reopen query is --------------->", ClientID, Mstorgnhirarchyid, RecordID, ClientID, Mstorgnhirarchyid, recorddiffseq)
// 	rows, err := mdao.DB.Query(sql, ClientID, Mstorgnhirarchyid, RecordID, ClientID, Mstorgnhirarchyid, recorddiffseq)
// 	defer rows.Close()
// 	if err != nil {
// 		logger.Log.Println("GetPreviousstatusdate Get Statement Prepare Error", err)
// 		return name, err
// 	}
// 	for rows.Next() {
// 		rows.Scan(&name)

// 	}
// 	return name, nil
// }

// func (mdao DbConn) UpdateWorknotedaycount(ClientID int64, OrgnID int64, RecordID int64, Daycount int64) error {
// 	var sql = "UPDATE recordfulldetails SET worknotenotupdated=?  WHERE clientid=? AND mstorgnhirarchyid=? AND recordid=?"
// 	stmt, err := mdao.DB.Prepare(sql)
// 	if err != nil {
// 		logger.Log.Println(err)
// 		return errors.New("SQL Prepare Error")
// 	}
// 	defer stmt.Close()
// 	_, err = stmt.Exec(Daycount, ClientID, OrgnID, RecordID)
// 	if err != nil {
// 		logger.Log.Println(err)
// 		return errors.New("SQL Execution Error")
// 	}

// 	return nil
// }

// //===========================================================================================

// func (mdao DbConn) GetMstSLADuePrevPushTime(recordID int64, id int64) (int64, error) {
// 	// logger.Log.Println("In side GetMstSLADueRowcount")
// 	logger.Log.Println("In side GetMstSLADueRowcount----->", recordID)
// 	var pushtime int64
// 	// var grpname string
// 	var sql = "select sum(pushtime) from mstsladue where therecordid=? and id<? and deleteflg=0"
// 	rows, err := mdao.DB.Query(sql, recordID, id)
// 	defer rows.Close()
// 	if err != nil {
// 		logger.Log.Println("getStageID Get Statement Prepare Error", err)
// 		return pushtime, err
// 	}
// 	for rows.Next() {
// 		err = rows.Scan(&pushtime)
// 		logger.Log.Println("getStageID rows.next() Error", err)
// 	}
// 	return pushtime, nil
// }
// func (mdao DbConn) GetMstSLADueRowcount(recordID int64) (int64, error) {
// 	// logger.Log.Println("In side GetMstSLADueRowcount")
// 	logger.Log.Println("In side GetMstSLADueRowcount----->", recordID)
// 	var count int64
// 	// var grpname string
// 	var sql = "select count(id) from mstsladue where therecordid=? and deleteflg=0"
// 	rows, err := mdao.DB.Query(sql, recordID)
// 	defer rows.Close()
// 	if err != nil {
// 		logger.Log.Println("getStageID Get Statement Prepare Error", err)
// 		return count, err
// 	}
// 	for rows.Next() {
// 		err = rows.Scan(&count)
// 		logger.Log.Println("getStageID rows.next() Error", err)
// 	}
// 	return count, nil
// }
// func (mdao DbConn) Getcurrentsatusid(ClientID int64, Mstorgnhirarchyid int64, RecordID int64) (int64, int64, string, error) {
// 	logger.Log.Println("In side Getrecorddiffidbystateid")
// 	var recorddiffid int64
// 	var name string
// 	var seqno int64
// 	var latestsatusid = "SELECT a.recorddiffid,b.seqno,b.name FROM maprecordtorecorddifferentiation a,mstrecorddifferentiation b WHERE a.clientid=? AND a.mstorgnhirarchyid=? AND a.recorddifftypeid=3 AND a.recordid=? AND a.islatest=1 AND a.recorddiffid=b.id and b.seqno in(3,13,5);"
// 	rows, err := mdao.DB.Query(latestsatusid, ClientID, Mstorgnhirarchyid, RecordID)
// 	defer rows.Close()
// 	if err != nil {
// 		logger.Log.Println("Getcurrentsatusid Get Statement Prepare Error", err)
// 		return recorddiffid, seqno, name, err
// 	}
// 	for rows.Next() {
// 		rows.Scan(&recorddiffid, &seqno, &name)

// 	}
// 	return recorddiffid, seqno, name, nil
// }
// func (mdao DbConn) GetRecordCurrentGrpID(ClientID int64, OrgnIDTypeID int64, recordID int64) (int64, string, error) {
// 	logger.Log.Println("In side getStageID")
// 	logger.Log.Println("In side getStageID----->", recordID)
// 	var grpID int64
// 	var grpname string
// 	var sql = "select assignedgroupid,assignedgroup from recordfulldetails where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0"
// 	rows, err := mdao.DB.Query(sql, ClientID, OrgnIDTypeID, recordID)
// 	defer rows.Close()
// 	if err != nil {
// 		logger.Log.Println("getStageID Get Statement Prepare Error", err)
// 		return grpID, grpname, err
// 	}
// 	for rows.Next() {
// 		err = rows.Scan(&grpID, &grpname)
// 		logger.Log.Println("getStageID rows.next() Error", err)
// 	}
// 	return grpID, grpname, nil
// }

// //===========================================

// func (dbc DbConn) CheckEventProcessFlag(clientid int64, mstorgnhirarchyid int64, therecordid int64, notificationeventid int64, notificationtemplateid int64) (string, error) {
// 	var value string
// 	// fmt.Println(getCategoryIds)
// 	rows, err := dbc.DB.Query(checkEventProcessFlag, clientid, mstorgnhirarchyid, therecordid, notificationeventid, notificationtemplateid)
// 	defer rows.Close()
// 	if err != nil {
// 		logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>CheckEventProcessFlag Get Statement Prepare Error", err)
// 		return value, err
// 	}
// 	// p("dao 2222")

// 	for rows.Next() {
// 		err := rows.Scan(&value)
// 		if err != nil {

// 			logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>CheckEventProcessFlag Get Statement Prepare Error", err)
// 		}
// 	}
// 	return value, nil
// }

// func (dbc DbConn) GetEventTemplateData(clientid int64, mstorgnhirarchyid int64, recordtypeid int64, workingcategoryid int64, priorityid int64, EventTypeid int64) ([]entities.EventTepmplateData, error) {
// 	// fmt.Println("yyyyyyyyyyyyyyyyyyyyyyyyyyyyy ")
// 	value := entities.EventTepmplateData{}
// 	allValue := []entities.EventTepmplateData{}

// 	var getEventTemplateData = `SELECT id, eventparams, eventtype FROM mstnotificationtemplate WHERE clientid = ? AND mstorgnhirarchyid = ? AND recordtypeid = ? AND workingcategoryid = ? AND eventtype=? AND channeltype = 1 AND deleteflg = 0 AND activeflg = 1 AND JSON_CONTAINS(eventparams, '` + strconv.FormatInt(priorityid, 10) + `', '$.processid')`
// 	// fmt.Println(getEventTemplateData)
// 	rows, err := dbc.DB.Query(getEventTemplateData, clientid, mstorgnhirarchyid, recordtypeid, workingcategoryid, EventTypeid)
// 	defer rows.Close()
// 	if err != nil {
// 		logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>getEventTemplateData Get Statement Prepare Error", err)
// 		return allValue, err
// 	}
// 	// fmt.Println("dao 2222")

// 	for rows.Next() {
// 		err := rows.Scan(&value.Id, &value.EventParams, &value.EventType)
// 		if err != nil {

// 			logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>getEventTemplateData Get Statement Prepare Error", err)
// 		}
// 		// fmt.Println(value)
// 		allValue = append(allValue, entities.EventTepmplateData{Id: value.Id, EventParams: value.EventParams, EventType: value.EventType})
// 		// p("dao 4444")
// 	}
// 	return allValue, nil
// }

// func (dbc DbConn) GetrecordFromsladue() ([]entities.SLASchedulerEntity, error) {
// 	value := entities.SLASchedulerEntity{}
// 	allValue := []entities.SLASchedulerEntity{}
// 	// fmt.Println(">>>>>>>>>>>>>>>>>", getrecordFromsladue)
// 	rows, err := dbc.DB.Query(getrecordFromsladue)
// 	defer rows.Close()
// 	if err != nil {
// 		logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>GetrecordFromsladue Get Statement Prepare Error", err)
// 		return allValue, err
// 	}
// 	// p("dao 2222")

// 	for rows.Next() {
// 		err := rows.Scan(&value.ID, &value.ClientID, &value.Mstorgnhirarchyid, &value.RecordID)
// 		if err != nil {

// 			logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>GetCategoryIds Get Statement Prepare Error", err)
// 		}

// 		allValue = append(allValue, entities.SLASchedulerEntity{ID: value.ID, ClientID: value.ClientID, Mstorgnhirarchyid: value.Mstorgnhirarchyid, RecordID: value.RecordID})
// 		// p("dao 4444")
// 	}
// 	return allValue, nil
// }

// func (dbc DbConn) GetCategoryIds(clientid int64, mstorgnhirarchyid int64, therecordid int64) (entities.SLATabEntity, error) {
// 	value := entities.SLATabEntity{}
// 	// fmt.Println(getCategoryIds)
// 	rows, err := dbc.DB.Query(getCategoryIds, therecordid, therecordid, therecordid, therecordid, clientid, mstorgnhirarchyid, therecordid)
// 	defer rows.Close()
// 	if err != nil {
// 		logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>GetMstsladue Get Statement Prepare Error", err)
// 		return value, err
// 	}
// 	// p("dao 2222")

// 	for rows.Next() {
// 		err := rows.Scan(&value.ClientID, &value.Mstorgnhirarchyid, &value.RecordID, &value.RecordtypeID, &value.PriorityID, &value.WorkingcatID)
// 		if err != nil {

// 			logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>GetCategoryIds Get Statement Prepare Error", err)
// 		}
// 	}
// 	return value, nil
// }

// func (dbc DbConn) UpdateExecuteFlag(ids []int64, flag int64) (bool, error) {
// 	// p := logger.Log.Println
// 	stringedIDs := fmt.Sprintf("%v", ids)
// 	stringedIDs = stringedIDs[1 : len(stringedIDs)-1]
// 	stringedIDs = strings.ReplaceAll(stringedIDs, " ", ",")
// 	if stringedIDs != "" {
// 		sqlRaw := fmt.Sprintf(updateExecuteFlag, stringedIDs)
// 		// fmt.Println(sqlRaw)
// 		rows, err := dbc.DB.Query(sqlRaw, flag)
// 		defer rows.Close()
// 		// p(rows)
// 		if err != nil {
// 			logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>UpdateExecuteFlag Get Statement Prepare Error", err)
// 			return false, err
// 		}
// 	}
// 	return true, nil
// }
// func (dbc DbConn) UpdateViolateFlag1(clientid int64, mstorgnhirarchyid int64, therecordid int64, flag int) (bool, error) {
// 	p := logger.Log.Println
// 	sql := ""
// 	if flag == 0 {
// 		sql = updateRespViolateFlag
// 		// sql2 = updateResplViolateFlagFullFillDtl
// 	} else {
// 		sql = updateResolViolateFlag
// 		// sql2 = updateResolViolateFlagFullFillDtl
// 	}
// 	stmt, err1 := dbc.DB.Prepare(sql)
// 	if err1 != nil {
// 		p("UpdateViolateFlag ------------>", err1)
// 		return false, err1
// 	}
// 	defer stmt.Close()
// 	_, err := stmt.Exec(therecordid, clientid, mstorgnhirarchyid)
// 	if err != nil {
// 		p("UpdateViolateFlag ------------>", err1)
// 		return false, err1
// 	}
// 	return true, nil
// }

// //
// func (dbc DbConn) UpdateViolateFlag2(clientid int64, mstorgnhirarchyid int64, therecordid int64, flag int) (bool, error) {
// 	p := logger.Log.Println
// 	sql2 := ""
// 	if flag == 0 {
// 		sql2 = updateResplViolateFlagFullFillDtl
// 	} else {
// 		// sql = updateResolViolateFlag
// 		sql2 = updateResolViolateFlagFullFillDtl
// 	}
// 	stmt, err1 := dbc.DB.Prepare(sql2)
// 	if err1 != nil {
// 		p("UpdateViolateFlag ------------>", err1)
// 		return false, err1
// 	}
// 	defer stmt.Close()
// 	_, err := stmt.Exec(therecordid, clientid, mstorgnhirarchyid)
// 	if err != nil {
// 		p("UpdateViolateFlag ------------>", err1)
// 		return false, err1
// 	}
// 	return true, nil
// }

// // ****************************************** Scheduler Query END *****************************************************************

// func (dbc DbConn) GetSpecificMstslafullfillmentcriteriaWithoutWCat(page *entities.MstslafullfillmentcriteriaEntity) (entities.MstslafullfillmentcriteriaEntity, error) {
// 	value := entities.MstslafullfillmentcriteriaEntity{}
// 	rows, err := dbc.DB.Query(getMstslafullfillmentcriteriawithoutworkingcategory, page.Clientid, page.Mstorgnhirarchyid, page.Mstrecorddifferentiationtickettypeid,
// 		page.Mstrecorddifferentiationpriorityid)
// 	defer rows.Close()
// 	if err != nil {
// 		logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>GetAllMstslafullfillmentcriteria Get Statement Prepare Error", err)
// 		return value, err
// 	}
// 	for rows.Next() {
// 		err := rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Slaid,
// 			&value.Mstrecorddifferentiationtickettypeid, &value.Mstrecorddifferentiationpriorityid,
// 			&value.Mstrecorddifferentiationworkingcatid, &value.Responsetimeinhr, &value.Responsetimeinmin,
// 			&value.Responsetimeinsec, &value.Resolutiontimeinhr, &value.Resolutiontimeinmin, &value.Resolutiontimeinsec,
// 			&value.Supportgroupspecific, &value.Activeflg)
// 		if err != nil {
// 			logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>GetAllMstslafullfillmentcriteria Get Statement Prepare Error", err)
// 		}
// 		// p("dao 4444")
// 	}
// 	// logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>", value)
// 	return value, nil
// }

// func (dbc DbConn) UpdateRessolutionEndFlag(clientid int64, mstorgnhirarchyid int64, therecordid int64, resolutionCompleteTime string) (bool, error) {
// 	// p := logger.Log.Println

// 	rows, err := dbc.DB.Query(updateRessolutionEndFlag, resolutionCompleteTime, therecordid, clientid, mstorgnhirarchyid)
// 	defer rows.Close()
// 	// p(rows)
// 	if err != nil {
// 		logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>GetAllMstslafullfillmentcriteria Get Statement Prepare Error", err)
// 		return false, err
// 	}
// 	// p("dao 2222")
// 	return true, nil
// }

// func (dbc DbConn) UpdateViolateFlag(clientid int64, mstorgnhirarchyid int64, therecordid int64, flag int) (bool, error) {
// 	p := logger.Log.Println
// 	sql := ""
// 	sql2 := ""
// 	if flag == 0 {
// 		sql = updateRespViolateFlag
// 		sql2 = updateResplViolateFlagFullFillDtl
// 	} else {
// 		sql = updateResolViolateFlag
// 		sql2 = updateResolViolateFlagFullFillDtl
// 	}
// 	rows, err := dbc.DB.Query(sql, therecordid, clientid, mstorgnhirarchyid)
// 	defer rows.Close()
// 	rows1, _ := dbc.DB.Query(sql2, therecordid, clientid, mstorgnhirarchyid)
// 	defer rows1.Close()
// 	if err != nil {
// 		p("UpdateViolateFlag ------------>", err)
// 		return false, err
// 	}
// 	return true, nil
// }

// func (dbc DbConn) GetMstsladue(clientid int64, mstorgnhirarchyid int64, therecordid int64) (entities.MstsladueEntity, error) {
// 	value := entities.MstsladueEntity{}
// 	rows, err := dbc.DB.Query(getMstsladue, therecordid, clientid, mstorgnhirarchyid)
// 	defer rows.Close()
// 	if err != nil {
// 		logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>GetMstsladue Get Statement Prepare Error", err)
// 		return value, err
// 	}
// 	// p("dao 2222")

// 	for rows.Next() {
// 		err := rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Mstslaentityid,
// 			&value.Therecordid, &value.Latestone,
// 			&value.Startdatetimeresponse, &value.Startdatetimeresolution, &value.Duedatetimeresponse,
// 			&value.Duedatetimeresolution, &value.DuedatetimeresolutionInt, &value.Duedatetimeresponseint, &value.Resoltiondone, &value.Resolutiondatetime,
// 			&value.Trnslaentityhistoryid, &value.Remainingtime, &value.Completepercent, &value.Responseremainingtime, &value.Responsepercentage, &value.Isresponsecomplete, &value.Isresolutioncomplete, &value.PushTime)
// 		if err != nil {

// 			logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>GetMstsladue Get Statement Prepare Error", err)
// 		}
// 	}
// 	return value, nil
// }

// func (dbc DbConn) Getutcdiff(clientid int64, mstorgnhirarchyid int64) (entities.ZoneEntity, error) {
// 	// p := fmt.Println

// 	value := entities.ZoneEntity{}
// 	// p("Client ID ")

// 	rows, err := dbc.DB.Query(getutcdiff, clientid, mstorgnhirarchyid)
// 	defer rows.Close()
// 	// p("dao 1111")
// 	// p(rows)
// 	if err != nil {
// 		//logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>GetAllMstslafullfillmentcriteria Get Statement Prepare Error", err)
// 		return value, err
// 	}
// 	// p("dao 2222")

// 	for rows.Next() {
// 		err := rows.Scan(&value.UTCdiff)
// 		// p("dao 3333")
// 		if err != nil {
// 			logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>Scan error in Getutcdiff", err) // Error related to the iteration of rows
// 		}
// 		// p("dao 4444")
// 	}
// 	return value, nil
// }

// func (dbc DbConn) UpdateResponseEndFlag(clientid int64, mstorgnhirarchyid int64, therecordid int64, responseCompleteTime string) (bool, error) {
// 	// p := fmt.Println

// 	rows, err := dbc.DB.Query(updateResponseEndFlag, responseCompleteTime, therecordid, clientid, mstorgnhirarchyid)
// 	defer rows.Close()
// 	// p("dao 1111")
// 	// p(rows)
// 	if err != nil {
// 		logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>UpdateResponseEndFlag Get Statement Prepare Error", err)
// 		return false, err
// 	}
// 	// p("dao 2222")
// 	return true, nil
// }

// func (dbc DbConn) UpdateRemainingPercent(clientid int64, mstorgnhirarchyid int64, therecordid int64, remainingtime int64, completepercent float64, responseremainingtime int64, responsepercentage float64) (bool, error) {
// 	// logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>UpdateRemainingPercent query is ------------->", updateRemainingPercent)
// 	// logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>UpdateRemainingPercent parameter is ------------->", remainingtime, completepercent, responseremainingtime, responsepercentage, therecordid, clientid, mstorgnhirarchyid)
// 	rows, err := dbc.DB.Query(updateRemainingPercent, remainingtime, completepercent, responseremainingtime, responsepercentage, therecordid, clientid, mstorgnhirarchyid)
// 	defer rows.Close()
// 	// p("dao 1111")
// 	//p(rows)
// 	if err != nil {
// 		//logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>GetAllMstslafullfillmentcriteria Get Statement Prepare Error", err)
// 		return false, err
// 	}
// 	// p("dao 2222")
// 	return true, nil
// }

// func (dbc DbConn) UpdateMstsladue(clientid int64, mstorgnhirarchyid int64, therecordid int64, duedatetimeresponse string, duedatetimeresolution string, duedatetimeresolutionint int64, duedatetimeresponseint int64, totalpushTime int64) (bool, error) {
// 	// p := fmt.Println
// 	// logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>.", duedatetimeresolutionint)
// 	rows, err := dbc.DB.Query(updateMstsladue, duedatetimeresponse, duedatetimeresolution, duedatetimeresolutionint, duedatetimeresponseint, totalpushTime, therecordid, clientid, mstorgnhirarchyid)
// 	defer rows.Close()
// 	// p("dao 1111")
// 	// p(rows)
// 	if err != nil {
// 		logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>UpdateMstsladue Get Statement Prepare Error", err)
// 		return false, err
// 	}
// 	// p("dao 2222")
// 	return true, nil
// }

// // func (dbc DbConn) GetMstsladue(clientid int64, mstorgnhirarchyid int64, therecordid int64) (entities.MstsladueEntity, error) {
// // 	value := entities.MstsladueEntity{}
// // 	rows, err := dbc.DB.Query(getMstsladue, therecordid, clientid, mstorgnhirarchyid)
// // 	defer rows.Close()
// // 	if err != nil {
// // 		logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>GetMstsladue Get Statement Prepare Error", err)
// // 		return value, err
// // 	}
// // 	// p("dao 2222")

// // 	for rows.Next() {
// // 		err := rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Mstslaentityid,
// // 			&value.Therecordid, &value.Latestone,
// // 			&value.Startdatetimeresponse, &value.Startdatetimeresolution, &value.Duedatetimeresponse,
// // 			&value.Duedatetimeresolution, &value.DuedatetimeresolutionInt, &value.Duedatetimeresponseint, &value.Resoltiondone, &value.Resolutiondatetime,
// // 			&value.Trnslaentityhistoryid, &value.Remainingtime, &value.Completepercent, &value.Responseremainingtime, &value.Responsepercentage, &value.Isresponsecomplete)
// // 		logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>GetMstsladue Get Statement Prepare Error", err)
// // 	}
// // 	return value, nil
// // }

// func (dbc DbConn) GetTrnslaentityhistory(clientid int64, mstorgnhirarchyid int64, therecordid int64) (entities.TrnslaentityhistoryEntity, error) {
// 	// p := fmt.Println

// 	value := entities.TrnslaentityhistoryEntity{}
// 	// p("Client ID ")

// 	rows, err := dbc.DB.Query(getTrnslaentityhistory, clientid, mstorgnhirarchyid, therecordid)
// 	defer rows.Close()
// 	// p("dao 1111")
// 	// p(rows)
// 	if err != nil {
// 		logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>GetTrnslaentityhistory Get Statement Prepare Error", err)
// 		return value, err
// 	}
// 	// p("dao 2222")

// 	for rows.Next() {
// 		err := rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Mstslaentityid,
// 			&value.Therecordid, &value.Recorddatetime,
// 			&value.Recorddatetoint, &value.Donotupdatesladue, &value.Recordtimetoint,
// 			&value.Mstslastateid, &value.Commentonrecord, &value.Slastartstopindicator, &value.Fromclientuserid)
// 		// p("dao 3333")
// 		if err != nil {
// 			logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>Scan error in GetTrnslaentityhistory", err) // Error related to the iteration of rows
// 		}
// 		// p("dao 4444")
// 	}
// 	return value, nil
// }

// func (dbc DbConn) InsertTrnslaentityhistory(tz *entities.TrnslaentityhistoryEntity) (int64, error) {
// 	// logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>history table id----------22222222222222222222222222222222------------------------------------------>", tz)
// 	// logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>history table id----------22222222222222222222222222222222------------------------------------------>", tz.Slastartstopindicator)
// 	stmt, err := dbc.DB.Prepare(insertTrnslaentityhistory)
// 	defer stmt.Close()
// 	if err != nil {
// 		logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>history table id----------22222222222222222222222222222222------------>", err)
// 		return 0, err
// 	}
// 	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Mstslaentityid, tz.Therecordid, tz.Recorddatetime, tz.Recorddatetoint, tz.Donotupdatesladue, tz.Recordtimetoint, tz.Mstslastateid, tz.Commentonrecord, tz.Slastartstopindicator, tz.Fromclientuserid)
// 	if err != nil {
// 		logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>history table id----------22222222222222222222222222222222------------>", err)
// 		return 0, err
// 	}
// 	lastInsertedId, err := res.LastInsertId()
// 	return lastInsertedId, nil
// }

// // func (dbc DbConn) InsertMstsladue(tz *entities.MstsladueEntity) (int64, error) {
// // 	// fmt.Println(tz.Clientid)
// // 	stmt, err := dbc.DB.Prepare(insertMstsladue)
// // 	defer stmt.Close()
// // 	if err != nil {
// // 		return 0, err
// // 	}
// // 	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Mstslaentityid, tz.Therecordid, tz.Latestone, tz.Startdatetimeresponse, tz.Startdatetimeresolution, tz.Duedatetimeresponse, tz.Duedatetimeresolution, tz.Duedatetimetominute, tz.Resoltiondone, tz.Resolutiondatetime, tz.Lastupdatedattime, tz.Trnslaentityhistoryid, tz.DuedatetimeresolutionInt, tz.Duedatetimeresponseint, tz.Remainingtime, tz.Completepercent)
// // 	if err != nil {
// // 		// fmt.Println("-----------------------------------------")
// // 		fmt.Println(err)
// // 		return 0, err
// // 	}
// // 	lastInsertedId, err := res.LastInsertId()
// // 	return lastInsertedId, nil
// // }

// func (dbc DbConn) InsertMstsladue(tz *entities.MstsladueEntity) (int64, error) {
// 	// logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>Insert into InsertMstsladue------------->", insertMstsladue)
// 	// logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>Insert into InsertMstsladue------------->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Mstslaentityid, tz.Therecordid, tz.Latestone, tz.Startdatetimeresponse, tz.Startdatetimeresolution, tz.Duedatetimeresponse, tz.Duedatetimeresolution, tz.Duedatetimetominute, tz.Resoltiondone, tz.Resolutiondatetime, tz.Lastupdatedattime, tz.Trnslaentityhistoryid, tz.DuedatetimeresolutionInt, tz.Duedatetimeresponseint, tz.Remainingtime, tz.Completepercent, tz.Responseremainingtime, tz.Responsepercentage)
// 	stmt, err := dbc.DB.Prepare(insertMstsladue)
// 	defer stmt.Close()
// 	if err != nil {
// 		logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>statement error ------------->", err)
// 		return 0, err
// 	}
// 	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Mstslaentityid, tz.Therecordid, tz.Latestone, tz.Startdatetimeresponse, tz.Startdatetimeresolution, tz.Duedatetimeresponse, tz.Duedatetimeresolution, tz.Duedatetimetominute, tz.Resoltiondone, tz.Resolutiondatetime, tz.Lastupdatedattime, tz.Trnslaentityhistoryid, tz.DuedatetimeresolutionInt, tz.Duedatetimeresponseint, tz.Remainingtime, tz.Completepercent, tz.Responseremainingtime, tz.Responsepercentage)
// 	if err != nil {
// 		// fmt.Println("-----------------------------------------")
// 		logger.Log.Println(err)
// 		return 0, err
// 	}
// 	lastInsertedId, err := res.LastInsertId()
// 	if err != nil {
// 		// fmt.Println("-----------------------------------------")
// 		logger.Log.Println(err)
// 		return 0, err
// 	}
// 	return lastInsertedId, nil
// }

// func (dbc DbConn) GetSpecificMstslafullfillmentcriteria(page *entities.MstslafullfillmentcriteriaEntity) (entities.MstslafullfillmentcriteriaEntity, error) {
// 	value := entities.MstslafullfillmentcriteriaEntity{}
// 	rows, err := dbc.DB.Query(getMstslafullfillmentcriteria, page.Clientid, page.Mstorgnhirarchyid, page.Mstrecorddifferentiationtickettypeid,
// 		page.Mstrecorddifferentiationpriorityid, page.Mstrecorddifferentiationworkingcatid)
// 	defer rows.Close()
// 	if err != nil {
// 		logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>GetAllMstslafullfillmentcriteria Get Statement Prepare Error", err)
// 		return value, err
// 	}
// 	for rows.Next() {
// 		err := rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Slaid,
// 			&value.Mstrecorddifferentiationtickettypeid, &value.Mstrecorddifferentiationpriorityid,
// 			&value.Mstrecorddifferentiationworkingcatid, &value.Responsetimeinhr, &value.Responsetimeinmin,
// 			&value.Responsetimeinsec, &value.Resolutiontimeinhr, &value.Resolutiontimeinmin, &value.Resolutiontimeinsec,
// 			&value.Supportgroupspecific, &value.Activeflg)
// 		if err != nil {
// 			logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>GetSpecificMstslafullfillmentcriteria ------------------------>", err)

// 		}
// 	}
// 	return value, nil
// }

// func (dbc DbConn) GetSupportGroupId(clientid int64, mstorgnhirarchyid int64, mstslafullfillmentcriteriaid int64) (int64, error) {
// 	// p := fmt.Println
// 	supportGroupId := int64(0)
// 	// p("clientid", clientid)
// 	// p("mstorgnhirarchyid", mstorgnhirarchyid)
// 	// p("mstslafullfillmentcriteriaid", mstslafullfillmentcriteriaid)
// 	rows, err := dbc.DB.Query(getsupportGroupid, clientid, mstorgnhirarchyid, mstslafullfillmentcriteriaid)
// 	defer rows.Close()
// 	// p("dao 1111")
// 	// p(rows)
// 	if err != nil {
// 		logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>GetAllMstslafullfillmentcriteria Get Statement Prepare Error", err)
// 		return supportGroupId, err
// 	}
// 	// p("dao 2222")
// 	// p(">>>>>>", rows)
// 	for rows.Next() {
// 		err := rows.Scan(&supportGroupId)
// 		// p("dao 3333")
// 		if err != nil {
// 			// panic(err) // Error related to the iteration of rows
// 			logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>Scan error in GetSupportGroupId", err)
// 		}
// 		// p("dao 4444")
// 	}
// 	return supportGroupId, nil
// }

// func (dbc DbConn) GetSupportGroupHoliday(clientid int64, mstorgnhirarchyid int64, supportGroupId int64, today string) (int64, int64, string, error) {
// 	// p := fmt.Println
// 	// p("************** GetSupportGroupHoliday *******************")
// 	starttime := int64(0)
// 	endtime := int64(0)
// 	var dateofholiday string
// 	rows, err := dbc.DB.Query(getclientsupportgroupholyday, clientid, mstorgnhirarchyid, supportGroupId, today)
// 	defer rows.Close()
// 	if err != nil {
// 		logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>GetAllMstslafullfillmentcriteria Get Statement Prepare Error", err)
// 		return 0, 0, "", err
// 	}
// 	for rows.Next() {
// 		err := rows.Scan(&starttime, &endtime, &dateofholiday)
// 		if err != nil {
// 			// panic(err) // Error related to the iteration of rows
// 			logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>Scan error in GetSupportGroupHoliday", err)
// 		}
// 	}
// 	return starttime, endtime, dateofholiday, nil
// }

// func (dbc DbConn) GetClientHoliday(clientid int64, mstorgnhirarchyid int64, today string) (int64, int64, string, error) {
// 	// p := fmt.Println
// 	// p("************** GetClientHoliday *******************")
// 	starttime := int64(0)
// 	endtime := int64(0)
// 	var dateofholiday string
// 	rows, err := dbc.DB.Query(getclientholyday, clientid, mstorgnhirarchyid, today)
// 	defer rows.Close()
// 	if err != nil {
// 		logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>GetAllMstslafullfillmentcriteria Get Statement Prepare Error", err)
// 		return 0, 0, "", err
// 	}
// 	for rows.Next() {
// 		err := rows.Scan(&starttime, &endtime, &dateofholiday)
// 		if err != nil {
// 			// panic(err) // Error related to the iteration of rows
// 			logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>Scan error in GetClientHoliday", err)
// 		}
// 	}
// 	return starttime, endtime, dateofholiday, nil
// }

// // This method will return the starttime and endtime of the week day
// func (dbc DbConn) GetClientDayOfWeek(clientid int64, mstorgnhirarchyid int64, dayofweekid int64) (int64, int64, error) {
// 	//p := fmt.Println
// 	// logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>************** GetClientDayOfWeek *******************", dayofweekid)
// 	starttime := int64(0)
// 	endtime := int64(0)
// 	// print(dayofweekid)
// 	rows, err := dbc.DB.Query(getclientWeekDay, clientid, mstorgnhirarchyid, dayofweekid)
// 	defer rows.Close()
// 	if err != nil {
// 		logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>************** GetClientDayOfWeek *******************", err)
// 		return 0, 0, err
// 	}
// 	for rows.Next() {
// 		err := rows.Scan(&starttime, &endtime)
// 		if err != nil {
// 			logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>************** GetClientDayOfWeek *******************", err)
// 		}
// 	}
// 	return starttime, endtime, nil
// }

// // This method will return the starttime and endtime of the week day for support group
// func (dbc DbConn) GetSupportGroupDayOfWeek(clientid int64, mstorgnhirarchyid int64, supportGroupId int64, dayofweekid int64) (int64, int64, int64, error) {
// 	// p := fmt.Println
// 	// p("************** GetSupportGroupDayOfWeek *******************")
// 	starttime := int64(0)
// 	endtime := int64(0)
// 	nextdayforward := int64(0)
// 	// print(dayofweekid)
// 	rows, err := dbc.DB.Query(getsupportgroupWeekDay, clientid, mstorgnhirarchyid, supportGroupId, dayofweekid)
// 	defer rows.Close()
// 	if err != nil {
// 		//logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>GetAllMstslafullfillmentcriteria Get Statement Prepare Error", err)
// 		return 0, 0, 0, err
// 	}
// 	for rows.Next() {
// 		err := rows.Scan(&starttime, &endtime, &nextdayforward)
// 		if err != nil {
// 			// panic(err) // Error related to the iteration of rows
// 			logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>Scan error on GetSupportGroupDayOfWeek", err)
// 		}
// 	}
// 	return starttime, endtime, nextdayforward, nil
// }

// // func (dbc DbConn) GetTrnslaentityhistorytype2(clientid int64, mstorgnhirarchyid int64, therecordid int64) (entities.TrnslaentityhistoryEntity, error) {
// // 	// p := fmt.Println

// // 	value := entities.TrnslaentityhistoryEntity{}
// // 	// p("Client ID ")

// // 	rows, err := dbc.DB.Query(getTrnslaentityhistorytype2, clientid, mstorgnhirarchyid, therecordid)
// // 	defer rows.Close()
// // 	// p("dao 1111")
// // 	// p(rows)
// // 	if err != nil {
// // 		logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>GetTrnslaentityhistory Get Statement Prepare Error", err)
// // 		return value, err
// // 	}
// // 	// p("dao 2222")

// // 	for rows.Next() {
// // 		err := rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Mstslaentityid,
// // 			&value.Therecordid, &value.Recorddatetime,
// // 			&value.Recorddatetoint, &value.Donotupdatesladue, &value.Recordtimetoint,
// // 			&value.Mstslastateid, &value.Commentonrecord, &value.Slastartstopindicator, &value.Fromclientuserid)
// // 		// p("dao 3333")
// // 		if err != nil {
// // 			panic(err) // Error related to the iteration of rows
// // 		}
// // 		// p("dao 4444")
// // 	}
// // 	return value, nil
// // }

// /*func (dbc DbConn) GetTrnslaentityhistorytype2(clientid int64, mstorgnhirarchyid int64, therecordid int64, trnId int64) ([]entities.TrnslaentityhistoryEntity, error) {
// 	// p := fmt.Println

// 	value := entities.TrnslaentityhistoryEntity{}
// 	allValue := []entities.TrnslaentityhistoryEntity{}
// 	// p("Client ID ")

// 	rows, err := dbc.DB.Query(getTrnslaentityhistorytype2, clientid, mstorgnhirarchyid, therecordid, trnId)
// 	defer rows.Close()
// 	// p("dao 1111")
// 	// p(rows)
// 	if err != nil {
// 		logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>GetTrnslaentityhistory Get Statement Prepare Error", err)
// 		return allValue, err
// 	}
// 	// p("dao 2222")

// 	for rows.Next() {
// 		err := rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Mstslaentityid,
// 			&value.Therecordid, &value.Recorddatetime,
// 			&value.Recorddatetoint, &value.Donotupdatesladue, &value.Recordtimetoint,
// 			&value.Mstslastateid, &value.Commentonrecord, &value.Slastartstopindicator, &value.Fromclientuserid)
// 		// p("dao 3333")
// 		if err != nil {
// 			// panic(err) // Error related to the iteration of rows
// 			logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>Scan error in GetTrnslaentityhistorytype2", err)
// 		}
// 		allValue = append(allValue, entities.TrnslaentityhistoryEntity{Id: value.Id, Clientid: value.Clientid, Mstorgnhirarchyid: value.Mstorgnhirarchyid, Mstslaentityid: value.Mstslaentityid,
// 			Therecordid: value.Therecordid, Recorddatetime: value.Recorddatetime,
// 			Recorddatetoint: value.Recorddatetoint, Donotupdatesladue: value.Donotupdatesladue, Recordtimetoint: value.Recordtimetoint,
// 			Mstslastateid: value.Mstslastateid, Commentonrecord: value.Commentonrecord, Slastartstopindicator: value.Slastartstopindicator, Fromclientuserid: value.Fromclientuserid})
// 		// p("dao 4444")
// 	}
// 	return allValue, nil
// }*/

// func (dbc DbConn) GetTrnslaentityhistorytype2(clientid int64, mstorgnhirarchyid int64, therecordid int64, trnId int64) ([]entities.TrnslaentityhistoryEntity, error) {
// 	// p := fmt.Println

// 	value := entities.TrnslaentityhistoryEntity{}
// 	allValue := []entities.TrnslaentityhistoryEntity{}
// 	// p("Client ID ")

// 	rows, err := dbc.DB.Query(getTrnslaentityhistorytype2, clientid, mstorgnhirarchyid, therecordid, trnId, trnId)
// 	defer rows.Close()
// 	// p("dao 1111")
// 	// p(rows)
// 	if err != nil {
// 		logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>GetTrnslaentityhistory Get Statement Prepare Error", err)
// 		return allValue, err
// 	}
// 	// p("dao 2222")

// 	for rows.Next() {
// 		err := rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Mstslaentityid,
// 			&value.Therecordid, &value.Recorddatetime,
// 			&value.Recorddatetoint, &value.Donotupdatesladue, &value.Recordtimetoint,
// 			&value.Mstslastateid, &value.Commentonrecord, &value.Slastartstopindicator, &value.Fromclientuserid)
// 		// p("dao 3333")
// 		if err != nil {
// 			// panic(err) // Error related to the iteration of rows
// 			logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>Scan error in GetTrnslaentityhistorytype2", err)
// 		}
// 		allValue = append(allValue, entities.TrnslaentityhistoryEntity{Id: value.Id, Clientid: value.Clientid, Mstorgnhirarchyid: value.Mstorgnhirarchyid, Mstslaentityid: value.Mstslaentityid,
// 			Therecordid: value.Therecordid, Recorddatetime: value.Recorddatetime,
// 			Recorddatetoint: value.Recorddatetoint, Donotupdatesladue: value.Donotupdatesladue, Recordtimetoint: value.Recordtimetoint,
// 			Mstslastateid: value.Mstslastateid, Commentonrecord: value.Commentonrecord, Slastartstopindicator: value.Slastartstopindicator, Fromclientuserid: value.Fromclientuserid})
// 		// p("dao 4444")
// 	}
// 	return allValue, nil
// }

// func (dbc DbConn) UpdatePushTimeInHistory(historyId int64, pushtime int64) (bool, error) {
// 	// p := fmt.Println
// 	rows, err := dbc.DB.Query(updatePushTimeInHistory, pushtime, historyId)
// 	defer rows.Close()
// 	// p(rows)
// 	if err != nil {
// 		return false, err
// 	}
// 	return true, nil
// }

// /*func (dbc DbConn) UpdateRecordFullFillDetails(details entities.SLARcordFullfillUpdate) (bool, error) {
// 	// fmt.Println("UpdateRecordFullFillDetails >>>>>>>>>>>>>> ", details)
// 	// // p := logger.Log.Println
// 	//logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>Updaterecordfulldetails >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")

// 	var sql = `UPDATE recordfulldetails SET responseslameterpercentage = ?, resolutionslameterpercentage = ?, businessaging = ?, calendaraging = ?, actualeffort = ?, slaidletime = ?, respoverduetime = ?, respoverdueperc = ?, resooverduetime = ?, resooverdueperc = ? WHERE clientid = ? AND mstorgnhirarchyid = ? AND recordid = ?`
// 	// // fmt.Println(sql)
// 	// // fmt.Println(details.Responseslameterpercentage, details.Resolutionslameterpercentage, details.Businessaging, details.Calendaraging, details.Actualeffort, details.Slaidletime, details.Respoverduetime, details.Respoverdueperc, details.Resooverduetime, details.Resooverdueperc, details.ClientID, details.Mstorgnhirarchyid)
// 	// rows, err := dbc.DB.Query(sql, details.Responseslameterpercentage, details.Resolutionslameterpercentage, details.Businessaging, details.Calendaraging, details.Actualeffort, details.Slaidletime, details.Respoverduetime, details.Respoverdueperc, details.Resooverduetime, details.Resooverdueperc, details.ClientID, details.Mstorgnhirarchyid, details.RecordID)
// 	// defer rows.Close()
// 	// // p(rows)
// 	// if err != nil {
// 	// 	logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>recordfulldetails Update Statement Prepare Error", err)
// 	// 	return false, err
// 	// }
// 	// // p("dao 2222")
// 	stmt, err := dbc.DB.Prepare(sql)
// 	if err != nil {
// 		logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>Updaterecordfulldetails Prepare Statement  Error", err)
// 		return false, err
// 	}
// 	defer stmt.Close()
// 	_, err = stmt.Exec(details.Responseslameterpercentage, details.Resolutionslameterpercentage, details.Businessaging, details.Calendaraging, details.Actualeffort, details.Slaidletime, details.Respoverduetime, details.Respoverdueperc, details.Resooverduetime, details.Resooverdueperc, details.ClientID, details.Mstorgnhirarchyid, details.RecordID)
// 	if err != nil {
// 		logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>Updaterecordfulldetails Execute Statement  Error", err)
// 		return false, err
// 	}
// 	return true, nil
// }*/

// func (dbc DbConn) UpdateRecordFullFillDetails(details entities.SLARcordFullfillUpdate) (bool, error) {
// 	// fmt.Println("UpdateRecordFullFillDetails >>>>>>>>>>>>>> ", details)
// 	// // p := logger.Log.Println
// 	//logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>Updaterecordfulldetails >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")

// 	var sql = `UPDATE recordfulldetails SET responseslameterpercentage = ?, resolutionslameterpercentage = ?, businessaging = ?-reopentimetaken, calendaraging = ?-reopentimetakencalendar, actualeffort = ?, slaidletime = ?, respoverduetime = ?, respoverdueperc = ?, resooverduetime = ?, resooverdueperc = ?,respslabreachstatus=?,resolslabreachstatus=?,worknotenotupdated=?  WHERE clientid = ? AND mstorgnhirarchyid = ? AND recordid = ?`
// 	// // fmt.Println(sql)
// 	// // fmt.Println(details.Responseslameterpercentage, details.Resolutionslameterpercentage, details.Businessaging, details.Calendaraging, details.Actualeffort, details.Slaidletime, details.Respoverduetime, details.Respoverdueperc, details.Resooverduetime, details.Resooverdueperc, details.ClientID, details.Mstorgnhirarchyid)
// 	// rows, err := dbc.DB.Query(sql, details.Responseslameterpercentage, details.Resolutionslameterpercentage, details.Businessaging, details.Calendaraging, details.Actualeffort, details.Slaidletime, details.Respoverduetime, details.Respoverdueperc, details.Resooverduetime, details.Resooverdueperc, details.ClientID, details.Mstorgnhirarchyid, details.RecordID)
// 	// defer rows.Close()
// 	// // p(rows)
// 	// if err != nil {
// 	//  logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>recordfulldetails Update Statement Prepare Error", err)
// 	//  return false, err
// 	// }
// 	// // p("dao 2222")
// 	stmt, err := dbc.DB.Prepare(sql)
// 	if err != nil {
// 		logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>Updaterecordfulldetails Prepare Statement  Error", err)
// 		return false, err
// 	}
// 	defer stmt.Close()
// 	_, err = stmt.Exec(details.Responseslameterpercentage, details.Resolutionslameterpercentage, details.Businessaging, details.Calendaraging, details.Actualeffort, details.Slaidletime, details.Respoverduetime, details.Respoverdueperc, details.Resooverduetime, details.Resooverdueperc, details.IsRespBreached, details.IsResoBreached, details.WorknoteNotUpdatedSince, details.ClientID, details.Mstorgnhirarchyid, details.RecordID)
// 	if err != nil {
// 		logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>Updaterecordfulldetails Execute Statement  Error", err)
// 		return false, err
// 	}
// 	return true, nil
// }

// func (dbc DbConn) FetchCurrentGrpID(ID int64) (int64, error) {
// 	// logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>In side FetchCurrentGrpID")
// 	var sql = "SELECT a.mstgroupid FROM mstrequest a,maprequestorecord b WHERE b.recordid=? AND a.id=b.mstrequestid AND b.activeflg=1 AND b.deleteflg=0 AND a.activeflg=1 AND a.deleteflg=0 order by b.id desc limit 1"
// 	var grpID int64
// 	rows, err := dbc.DB.Query(sql, ID)
// 	defer rows.Close()
// 	if err != nil {
// 		logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>FetchCurrentGrpID Get Statement Prepare Error", err)
// 		return grpID, err
// 	}
// 	for rows.Next() {
// 		err = rows.Scan(&grpID)
// 		if err != nil {
// 			logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>Scan error in FetchCurrentGrpID", err)
// 		}
// 	}
// 	return grpID, nil
// }

// var getrecordFromsladueold = `SELECT b.id, b.clientid, b.mstorgnhirarchyid, b.therecordid FROM recordfulldetails a, mstsladue b where a.status='Resolved' AND a.recordid = b.therecordid;`

// var getResolveDateSql = `SELECT recorddatetime FROM trnslaentityhistory WHERE therecordid = ? AND clientid = ? AND mstorgnhirarchyid = ? AND slastartstopindicator = 4 ORDER BY recorddatetime DESC LIMIT 1`

// var getResponseEndDateSql = `SELECT recorddatetime FROM trnslaentityhistory WHERE therecordid = ? AND clientid = ? AND mstorgnhirarchyid = ? AND slastartstopindicator = 3 ORDER BY recorddatetime ASC LIMIT 1`

// func (dbc DbConn) GetrecordFromsladueForOld() ([]entities.SLASchedulerEntity, error) {
// 	value := entities.SLASchedulerEntity{}
// 	allValue := []entities.SLASchedulerEntity{}
// 	// fmt.Println(">>>>>>>>>>>>>>>>>", getrecordFromsladue)
// 	rows, err := dbc.DB.Query(getrecordFromsladueold)
// 	defer rows.Close()
// 	if err != nil {
// 		logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>getrecordFromsladueold Get Statement Prepare Error", err)
// 		return allValue, err
// 	}
// 	// p("dao 2222")

// 	for rows.Next() {
// 		err := rows.Scan(&value.ID, &value.ClientID, &value.Mstorgnhirarchyid, &value.RecordID)
// 		if err != nil {

// 			logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>GetCategoryIds Get Statement Prepare Error", err)
// 		}

// 		allValue = append(allValue, entities.SLASchedulerEntity{ID: value.ID, ClientID: value.ClientID, Mstorgnhirarchyid: value.Mstorgnhirarchyid, RecordID: value.RecordID})
// 		// p("dao 4444")
// 	}
// 	return allValue, nil
// }

// func (dbc DbConn) GetSLAResolveDateSql(RecordId int64, ClientID int64, Mstorgnhirarchyid int64) (string, error) {
// 	var recordDateTime string
// 	rows, err := dbc.DB.Query(getResolveDateSql, RecordId, ClientID, Mstorgnhirarchyid)
// 	defer rows.Close()
// 	if err != nil {
// 		return recordDateTime, err
// 	}
// 	for rows.Next() {
// 		err = rows.Scan(&recordDateTime)
// 		if err != nil {
// 			logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>Scan error in FetchCurrentGrpID", err)
// 		}
// 	}
// 	return recordDateTime, nil
// }

// func (dbc DbConn) GetResponseEndDateSql(RecordId int64, ClientID int64, Mstorgnhirarchyid int64) (string, error) {
// 	var recordDateTime string
// 	rows, err := dbc.DB.Query(getResponseEndDateSql, RecordId, ClientID, Mstorgnhirarchyid)
// 	defer rows.Close()
// 	if err != nil {
// 		return recordDateTime, err
// 	}
// 	for rows.Next() {
// 		err = rows.Scan(&recordDateTime)
// 		if err != nil {
// 			logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>Scan error in FetchCurrentGrpID", err)
// 		}
// 	}
// 	return recordDateTime, nil
// }

// var getMstsladueFirst = `SELECT startdatetimeresolution,pushtime FROM mstsladue WHERE therecordid = ? AND clientid = ? AND mstorgnhirarchyid = ? AND activeflg = 1 AND deleteflg = 0 ORDER BY id ASC LIMIT 1`

// func (dbc DbConn) GetFirstStartTimeResolution(clientid int64, mstorgnhirarchyid int64, therecordid int64) (string, int64, error) {
// 	logger.Log.Println("Parameter :", therecordid, clientid, mstorgnhirarchyid)

// 	rows, err := dbc.DB.Query(getMstsladueFirst, therecordid, clientid, mstorgnhirarchyid)
// 	var Startdatetimeresolution = ""
// 	var pushtime int64
// 	defer rows.Close()
// 	if err != nil {
// 		logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>GetMstsladue Get Statement Prepare Error", err)
// 		return Startdatetimeresolution, pushtime, err
// 	}
// 	// p("dao 2222")

// 	for rows.Next() {
// 		err := rows.Scan(&Startdatetimeresolution, &pushtime)
// 		if err != nil {

// 			logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>GetMstsladue Get Statement Prepare Error", err)
// 		}
// 	}
// 	logger.Log.Println("PRV push time :", pushtime)

// 	return Startdatetimeresolution, pushtime, nil
// }
