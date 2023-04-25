package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertMstbusinessmatrix = "INSERT INTO mstbusinessmatrix (clientid, mstorgnhirarchyid, mstrecorddifferentiationtickettypeid, mstrecorddifferentiationcatid, mstrecorddifferentiationimpactid, mstrecorddifferentiationurgencyid, mstrecorddifferentiationpriorityid) VALUES (?,?,?,?,?,?,?)"
var duplicateMstbusinessmatrixurgecy = "SELECT count(id) total FROM  mstbusinessmatrix WHERE clientid = ? AND mstorgnhirarchyid = ? AND mstrecorddifferentiationtickettypeid = ? AND mstrecorddifferentiationcatid = ? AND mstrecorddifferentiationimpactid = ? AND mstrecorddifferentiationurgencyid = ? AND mstrecorddifferentiationpriorityid = ? AND deleteflg = 0 and activeflg=1"
var duplicateMstbusinessmatrixcat = "SELECT count(id) total FROM  mstbusinessmatrix WHERE clientid = ? AND mstorgnhirarchyid = ? AND mstrecorddifferentiationtickettypeid = ? AND mstrecorddifferentiationcatid = ? AND mstrecorddifferentiationpriorityid = ? AND deleteflg = 0 and activeflg=1"

// var getMstbusinessmatrix = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.mstrecorddifferentiationtickettypeid as Mstrecorddifferentiationtickettypeid, a.mstrecorddifferentiationcatid as Mstrecorddifferentiationcatid, a.mstrecorddifferentiationimpactid as Mstrecorddifferentiationimpactid, a.mstrecorddifferentiationurgencyid as Mstrecorddifferentiationurgencyid, a.mstrecorddifferentiationpriorityid as Mstrecorddifferentiationpriorityid, a.activeflg as Activeflg,(select name from mstclient where id=a.clientid) as Clientname,(select name from mstorgnhierarchy where id=a.mstorgnhirarchyid) as Mstorgnhirarchyname,(select name from mstrecorddifferentiation where id=a.mstrecorddifferentiationtickettypeid and deleteflg =0 and activeflg=1) as Tickettype,(select recorddifftypeid from mstrecorddifferentiation where id=a.mstrecorddifferentiationtickettypeid and deleteflg =0 and activeflg=1) as Mstrecordtickettypedifftypeid,COALESCE((select name from mstrecorddifferentiation where id=a.mstrecorddifferentiationcatid and deleteflg =0 and activeflg=1),'NA') as Categoryname,COALESCE((select recorddifftypeid from mstrecorddifferentiation where id=a.mstrecorddifferentiationcatid and deleteflg =0 and activeflg=1),'0') as Mstrecordcatlabelid,COALESCE((select name from mstrecorddifferentiation where id=a.mstrecorddifferentiationimpactid and deleteflg =0 and activeflg=1),'NA') as Impactname,COALESCE((select recorddifftypeid from mstrecorddifferentiation where id=a.mstrecorddifferentiationimpactid and deleteflg =0 and activeflg=1),'0') as Mstrecordimpactdifftypeid,COALESCE((select name from mstrecorddifferentiation where id=a.mstrecorddifferentiationurgencyid and deleteflg =0 and activeflg=1),'NA') as Urgencyname,COALESCE((select recorddifftypeid from mstrecorddifferentiation where id=a.mstrecorddifferentiationurgencyid and deleteflg =0 and activeflg=1),'0') as Mstrecordurgencydifftypeid,COALESCE((select name from mstrecorddifferentiation where id=a.mstrecorddifferentiationpriorityid and deleteflg =0 and activeflg=1),'NA') as Priorityname,COALESCE((select recorddifftypeid from mstrecorddifferentiation where id=a.mstrecorddifferentiationpriorityid and deleteflg =0 and activeflg=1),'0') as Mstrecordprioritydifftypeid,(select parentcategorynames from mstrecorddifferentiation where id=a.mstrecorddifferentiationcatid and deleteflg =0 and activeflg=1) as Parentcatnames,e.estimatedtime,e.efficiency  FROM mstbusinessmatrix a LEFT JOIN mapcategorywithestimatetime e ON mstrecorddifferentiationcatid=e.recorddiffid AND e.deleteflg =0 and e.activeflg=1 WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 ORDER BY a.id DESC LIMIT ?,?"
// var getMstbusinessmatrixcount = "SELECT count(id) total FROM  mstbusinessmatrix WHERE clientid = ? AND mstorgnhirarchyid = ? AND  deleteflg =0 and activeflg=1"
var updateMstbusinessmatrix = "UPDATE mstbusinessmatrix SET mstorgnhirarchyid = ?, mstrecorddifferentiationtickettypeid = ?, mstrecorddifferentiationcatid = ?, mstrecorddifferentiationimpactid = ?, mstrecorddifferentiationurgencyid = ?, mstrecorddifferentiationpriorityid = ? WHERE id = ? "
var deleteMstbusinessmatrix = "UPDATE mstbusinessmatrix SET deleteflg = '1' WHERE id = ? "
var checkmatrixconfig = "SELECT direction from mstbusinessdirection where clientid = ? AND mstorgnhirarchyid = ? and mstrecorddifferentiationtypeid=? and mstrecorddifferentiationid=? and activeflg=1 and deleteflg=0"

//var lastlevelcategory = "SELECT   mstrecorddifferentiation.id as Id, mstrecorddifferentiation.name as Name, mstrecorddifferentiation.recorddifftypeid as Lastcategorylevelid,mstrecorddifferentiation.parentcategorynames as Parentcatnames FROM mstrecordtype, mstrecorddifferentiationtype, mstrecorddifferentiation WHERE mstrecordtype.clientid = ? AND mstrecordtype.mstorgnhirarchyid = ? AND mstrecordtype.activeflg = 1 AND mstrecordtype.deleteflg = 0 AND mstrecordtype.fromrecorddifftypeid = ? AND mstrecordtype.fromrecorddiffid = ? AND mstrecorddifferentiationtype.activeflg = 1 AND mstrecorddifferentiationtype.deleteflg = 0 AND mstrecorddifferentiationtype.id = mstrecordtype.torecorddifftypeid AND mstrecorddifferentiationtype.parentid = 1 AND mstrecorddifferentiation.activeflg = 1 AND mstrecorddifferentiation.deleteflg = 0 AND mstrecorddifferentiation.id = mstrecordtype.torecorddiffid AND mstrecorddifferentiation.recorddifftypeid IN (SELECT  mstrecorddifferentiationtype.id FROM mstrecorddifferentiationtype WHERE mstrecorddifferentiationtype.clientid = ? AND mstrecorddifferentiationtype.mstorgnhirarchyid = ? AND mstrecorddifferentiationtype.activeflg = 1 AND mstrecorddifferentiationtype.deleteflg = 0 AND mstrecorddifferentiationtype.parentid = 1 AND mstrecorddifferentiationtype.seqno = (SELECT  MAX(mstrecorddifferentiationtype.seqno) FROM mstrecorddifferentiationtype WHERE mstrecorddifferentiationtype.clientid = ? AND mstrecorddifferentiationtype.mstorgnhirarchyid = ? AND mstrecorddifferentiationtype.activeflg = 1 AND mstrecorddifferentiationtype.deleteflg = 0 AND mstrecorddifferentiationtype.parentid = 1)) ORDER BY mstrecorddifferentiationtype.seqno ASC , mstrecorddifferentiation.seqno ASC"
var lastlevelcategory = "SELECT  mstrecorddifferentiation.id AS Id,mstrecorddifferentiation.name AS Name,mstrecorddifferentiation.recorddifftypeid AS Lastcategorylevelid,mstrecorddifferentiation.parentcategorynames AS Parentcatnames FROM mstrecorddifferentiation,mstrecordtype WHERE mstrecorddifferentiation.recorddifftypeid=(SELECT MAX(mstrecorddifferentiationtype.id) AS ID FROM mstrecordtype ,mstrecorddifferentiationtype,mstrecorddifferentiation WHERE mstrecordtype.clientid = ? AND mstrecordtype.mstorgnhirarchyid = ? AND mstrecordtype.activeflg = 1 AND mstrecordtype.deleteflg = 0 AND mstrecordtype.fromrecorddifftypeid = ? AND mstrecordtype.fromrecorddiffid = ? AND mstrecordtype.torecorddifftypeid = mstrecorddifferentiationtype.id AND mstrecordtype.torecorddiffid = 0 AND mstrecorddifferentiationtype.parentid=1 AND mstrecorddifferentiationtype.id = mstrecorddifferentiation.recorddifftypeid) AND mstrecorddifferentiation.activeflg=1 AND mstrecorddifferentiation.deleteflg=0 AND mstrecorddifferentiation.id = mstrecordtype.torecorddiffid AND mstrecordtype.clientid = ? AND mstrecordtype.mstorgnhirarchyid = ? AND mstrecordtype.fromrecorddifftypeid = ? AND mstrecordtype.fromrecorddiffid = ? AND mstrecordtype.activeflg=1 AND mstrecordtype.deleteflg=0"

func (dbc DbConn) CheckDuplicateMstbusinessmatrixurgencywise(tz *entities.MstbusinessmatrixEntity) (entities.MstbusinessmatrixEntities, error) {
	logger.Log.Println("In side CheckDuplicateMstbusinessmatrix")
	value := entities.MstbusinessmatrixEntities{}
	err := dbc.DB.QueryRow(duplicateMstbusinessmatrixurgecy, tz.Clientid, tz.Mstorgnhirarchyid, tz.Mstrecorddifferentiationtickettypeid, tz.Mstrecorddifferentiationcatid, tz.Mstrecorddifferentiationimpactid, tz.Mstrecorddifferentiationurgencyid, tz.Mstrecorddifferentiationpriorityid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMstbusinessmatrix Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) CheckDuplicateMstbusinessmatrixcatwise(tz *entities.MstbusinessmatrixEntity) (entities.MstbusinessmatrixEntities, error) {
	logger.Log.Println("In side CheckDuplicateMstbusinessmatrix")
	value := entities.MstbusinessmatrixEntities{}
	err := dbc.DB.QueryRow(duplicateMstbusinessmatrixcat, tz.Clientid, tz.Mstorgnhirarchyid, tz.Mstrecorddifferentiationtickettypeid, tz.Mstrecorddifferentiationcatid, tz.Mstrecorddifferentiationpriorityid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMstbusinessmatrix Get Statement Prepare Error", err)
		return value, err
	}
}

func (mdao DbConn) CheckDuplicateEstimatedTimeeffort(tz *entities.MstbusinessmatrixEntity) (int64, error) {
	var sql = "SELECT COUNT(id) as count FROM mapcategorywithestimatetime WHERE clientid=? AND mstorgnhirarchyid=? AND recorddiffid=? AND deleteflg=0 AND activeflg=1"
	var count int64
	rows, err := mdao.DB.Query(sql, tz.Clientid, tz.Mstorgnhirarchyid, tz.Mstrecorddifferentiationcatid)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetTermID Get Statement Prepare Error", err)
		return count, err
	}
	for rows.Next() {
		err = rows.Scan(&count)
		logger.Log.Println("GetTermID rows.next() Error", err)
	}
	return count, nil
}

func (dbc DbConn) InsertMstbusinessmatrix(tz *entities.MstbusinessmatrixEntity) (int64, error) {
	logger.Log.Println("In side InsertMstbusinessmatrix")
	logger.Log.Println("Query -->", insertMstbusinessmatrix)
	stmt, err := dbc.DB.Prepare(insertMstbusinessmatrix)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMstbusinessmatrix Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Mstrecorddifferentiationtickettypeid, tz.Mstrecorddifferentiationcatid, tz.Mstrecorddifferentiationimpactid, tz.Mstrecorddifferentiationurgencyid, tz.Mstrecorddifferentiationpriorityid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Mstrecorddifferentiationtickettypeid, tz.Mstrecorddifferentiationcatid, tz.Mstrecorddifferentiationimpactid, tz.Mstrecorddifferentiationurgencyid, tz.Mstrecorddifferentiationpriorityid)
	if err != nil {
		logger.Log.Println("InsertMstbusinessmatrix Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) InsertEstimatedTimeeffort(tz *entities.MstbusinessmatrixEntity) (int64, error) {
	var sql = "INSERT INTO mapcategorywithestimatetime(clientid,mstorgnhirarchyid,recorddiffid,estimatedtime,efficiency,changetype) values (?,?,?,?,?,?)"
	stmt, err := dbc.DB.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMstbusinessmatrix Prepare Statement  Error", err)
		return 0, err
	}
	//logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Mstrecorddifferentiationtickettypeid, tz.Mstrecorddifferentiationcatid, tz.Mstrecorddifferentiationimpactid, tz.Mstrecorddifferentiationurgencyid, tz.Mstrecorddifferentiationpriorityid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Mstrecorddifferentiationcatid, tz.Estimatedeffort, tz.Slacompliance, tz.ChangeType)
	if err != nil {
		logger.Log.Println("InsertMstbusinessmatrix Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllMstbusinessmatrix(page *entities.MstbusinessmatrixEntity, OrgnType int64) ([]entities.MstbusinessmatrixEntity, error) {
	logger.Log.Println("In side GelAllMstbusinessmatrix")
	values := []entities.MstbusinessmatrixEntity{}
	var getMstbusinessmatrix string
	var params []interface{}
	if OrgnType == 1 {
		getMstbusinessmatrix = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.mstrecorddifferentiationtickettypeid as Mstrecorddifferentiationtickettypeid, a.mstrecorddifferentiationcatid as Mstrecorddifferentiationcatid, a.mstrecorddifferentiationimpactid as Mstrecorddifferentiationimpactid, a.mstrecorddifferentiationurgencyid as Mstrecorddifferentiationurgencyid, a.mstrecorddifferentiationpriorityid as Mstrecorddifferentiationpriorityid, a.activeflg as Activeflg,(select name from mstclient where id=a.clientid) as Clientname,(select name from mstorgnhierarchy where id=a.mstorgnhirarchyid) as Mstorgnhirarchyname,(select name from mstrecorddifferentiation where id=a.mstrecorddifferentiationtickettypeid and deleteflg =0 and activeflg=1) as Tickettype,(select recorddifftypeid from mstrecorddifferentiation where id=a.mstrecorddifferentiationtickettypeid and deleteflg =0 and activeflg=1) as Mstrecordtickettypedifftypeid,COALESCE((select name from mstrecorddifferentiation where id=a.mstrecorddifferentiationcatid and deleteflg =0 and activeflg=1),'NA') as Categoryname,COALESCE((select recorddifftypeid from mstrecorddifferentiation where id=a.mstrecorddifferentiationcatid and deleteflg =0 and activeflg=1),'0') as Mstrecordcatlabelid,COALESCE((select name from mstrecorddifferentiation where id=a.mstrecorddifferentiationimpactid and deleteflg =0 and activeflg=1),'NA') as Impactname,COALESCE((select recorddifftypeid from mstrecorddifferentiation where id=a.mstrecorddifferentiationimpactid and deleteflg =0 and activeflg=1),'0') as Mstrecordimpactdifftypeid,COALESCE((select name from mstrecorddifferentiation where id=a.mstrecorddifferentiationurgencyid and deleteflg =0 and activeflg=1),'NA') as Urgencyname,COALESCE((select recorddifftypeid from mstrecorddifferentiation where id=a.mstrecorddifferentiationurgencyid and deleteflg =0 and activeflg=1),'0') as Mstrecordurgencydifftypeid,COALESCE((select name from mstrecorddifferentiation where id=a.mstrecorddifferentiationpriorityid and deleteflg =0 and activeflg=1),'NA') as Priorityname,COALESCE((select recorddifftypeid from mstrecorddifferentiation where id=a.mstrecorddifferentiationpriorityid and deleteflg =0 and activeflg=1),'0') as Mstrecordprioritydifftypeid,coalesce((select parentcategorynames from mstrecorddifferentiation where id=a.mstrecorddifferentiationcatid and deleteflg =0 and activeflg=1),'N/A') as Parentcatnames,coalesce(e.estimatedtime,'N/A') as estimatedtime,coalesce(e.efficiency,'N/A') as efficiency,coalesce(e.changetype,'N/A') as changetype FROM mstbusinessmatrix a LEFT JOIN mapcategorywithestimatetime e ON a.mstrecorddifferentiationcatid=e.recorddiffid AND a.clientid=e.clientid AND a.mstorgnhirarchyid=e.mstorgnhirarchyid AND e.deleteflg =0 and e.activeflg=1  WHERE  a.deleteflg =0 and a.activeflg=1 ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else if OrgnType == 2 {
		getMstbusinessmatrix = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.mstrecorddifferentiationtickettypeid as Mstrecorddifferentiationtickettypeid, a.mstrecorddifferentiationcatid as Mstrecorddifferentiationcatid, a.mstrecorddifferentiationimpactid as Mstrecorddifferentiationimpactid, a.mstrecorddifferentiationurgencyid as Mstrecorddifferentiationurgencyid, a.mstrecorddifferentiationpriorityid as Mstrecorddifferentiationpriorityid, a.activeflg as Activeflg,(select name from mstclient where id=a.clientid) as Clientname,(select name from mstorgnhierarchy where id=a.mstorgnhirarchyid) as Mstorgnhirarchyname,(select name from mstrecorddifferentiation where id=a.mstrecorddifferentiationtickettypeid and deleteflg =0 and activeflg=1) as Tickettype,(select recorddifftypeid from mstrecorddifferentiation where id=a.mstrecorddifferentiationtickettypeid and deleteflg =0 and activeflg=1) as Mstrecordtickettypedifftypeid,COALESCE((select name from mstrecorddifferentiation where id=a.mstrecorddifferentiationcatid and deleteflg =0 and activeflg=1),'NA') as Categoryname,COALESCE((select recorddifftypeid from mstrecorddifferentiation where id=a.mstrecorddifferentiationcatid and deleteflg =0 and activeflg=1),'0') as Mstrecordcatlabelid,COALESCE((select name from mstrecorddifferentiation where id=a.mstrecorddifferentiationimpactid and deleteflg =0 and activeflg=1),'NA') as Impactname,COALESCE((select recorddifftypeid from mstrecorddifferentiation where id=a.mstrecorddifferentiationimpactid and deleteflg =0 and activeflg=1),'0') as Mstrecordimpactdifftypeid,COALESCE((select name from mstrecorddifferentiation where id=a.mstrecorddifferentiationurgencyid and deleteflg =0 and activeflg=1),'NA') as Urgencyname,COALESCE((select recorddifftypeid from mstrecorddifferentiation where id=a.mstrecorddifferentiationurgencyid and deleteflg =0 and activeflg=1),'0') as Mstrecordurgencydifftypeid,COALESCE((select name from mstrecorddifferentiation where id=a.mstrecorddifferentiationpriorityid and deleteflg =0 and activeflg=1),'NA') as Priorityname,COALESCE((select recorddifftypeid from mstrecorddifferentiation where id=a.mstrecorddifferentiationpriorityid and deleteflg =0 and activeflg=1),'0') as Mstrecordprioritydifftypeid,coalesce((select parentcategorynames from mstrecorddifferentiation where id=a.mstrecorddifferentiationcatid and deleteflg =0 and activeflg=1),'N/A') as Parentcatnames,coalesce(e.estimatedtime,'N/A') as estimatedtime,coalesce(e.efficiency,'N/A') as efficiency,coalesce(e.changetype,'N/A') as changetype  FROM mstbusinessmatrix a LEFT JOIN mapcategorywithestimatetime e ON a.mstrecorddifferentiationcatid=e.recorddiffid AND a.clientid=e.clientid AND a.mstorgnhirarchyid=e.mstorgnhirarchyid AND e.deleteflg =0 and e.activeflg=1 WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else {
		getMstbusinessmatrix = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.mstrecorddifferentiationtickettypeid as Mstrecorddifferentiationtickettypeid, a.mstrecorddifferentiationcatid as Mstrecorddifferentiationcatid, a.mstrecorddifferentiationimpactid as Mstrecorddifferentiationimpactid, a.mstrecorddifferentiationurgencyid as Mstrecorddifferentiationurgencyid, a.mstrecorddifferentiationpriorityid as Mstrecorddifferentiationpriorityid, a.activeflg as Activeflg,(select name from mstclient where id=a.clientid) as Clientname,(select name from mstorgnhierarchy where id=a.mstorgnhirarchyid) as Mstorgnhirarchyname,(select name from mstrecorddifferentiation where id=a.mstrecorddifferentiationtickettypeid and deleteflg =0 and activeflg=1) as Tickettype,(select recorddifftypeid from mstrecorddifferentiation where id=a.mstrecorddifferentiationtickettypeid and deleteflg =0 and activeflg=1) as Mstrecordtickettypedifftypeid,COALESCE((select name from mstrecorddifferentiation where id=a.mstrecorddifferentiationcatid and deleteflg =0 and activeflg=1),'NA') as Categoryname,COALESCE((select recorddifftypeid from mstrecorddifferentiation where id=a.mstrecorddifferentiationcatid and deleteflg =0 and activeflg=1),'0') as Mstrecordcatlabelid,COALESCE((select name from mstrecorddifferentiation where id=a.mstrecorddifferentiationimpactid and deleteflg =0 and activeflg=1),'NA') as Impactname,COALESCE((select recorddifftypeid from mstrecorddifferentiation where id=a.mstrecorddifferentiationimpactid and deleteflg =0 and activeflg=1),'0') as Mstrecordimpactdifftypeid,COALESCE((select name from mstrecorddifferentiation where id=a.mstrecorddifferentiationurgencyid and deleteflg =0 and activeflg=1),'NA') as Urgencyname,COALESCE((select recorddifftypeid from mstrecorddifferentiation where id=a.mstrecorddifferentiationurgencyid and deleteflg =0 and activeflg=1),'0') as Mstrecordurgencydifftypeid,COALESCE((select name from mstrecorddifferentiation where id=a.mstrecorddifferentiationpriorityid and deleteflg =0 and activeflg=1),'NA') as Priorityname,COALESCE((select recorddifftypeid from mstrecorddifferentiation where id=a.mstrecorddifferentiationpriorityid and deleteflg =0 and activeflg=1),'0') as Mstrecordprioritydifftypeid,coalesce((select parentcategorynames from mstrecorddifferentiation where id=a.mstrecorddifferentiationcatid and deleteflg =0 and activeflg=1),'N/A') as Parentcatnames,coalesce(e.estimatedtime,'N/A') as estimatedtime,coalesce(e.efficiency,'N/A') as efficiency,coalesce(e.changetype,'N/A') as changetype  FROM mstbusinessmatrix a LEFT JOIN mapcategorywithestimatetime e ON a.mstrecorddifferentiationcatid=e.recorddiffid AND a.clientid=e.clientid AND a.mstorgnhirarchyid=e.mstorgnhirarchyid AND e.deleteflg =0 and e.activeflg=1 WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Mstorgnhirarchyid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	}
	rows, err := dbc.DB.Query(getMstbusinessmatrix, params...)
	// fmt.Println(rows)

	// rows, err := dbc.DB.Query(getMstbusinessmatrix, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstbusinessmatrix Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstbusinessmatrixEntity{}
		err = rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Mstrecorddifferentiationtickettypeid, &value.Mstrecorddifferentiationcatid, &value.Mstrecorddifferentiationimpactid, &value.Mstrecorddifferentiationurgencyid, &value.Mstrecorddifferentiationpriorityid, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Tickettype, &value.Mstrecordtickettypedifftypeid, &value.Catname, &value.Mstrecordcatlabelid, &value.Impactname, &value.Mstrecordimpactdifftypeid, &value.Urgencyname, &value.Mstrecordurgencydifftypeid, &value.Priorityname, &value.Mstrecordprioritydifftypeid, &value.Parentcatname, &value.Estimatedeffort, &value.Slacompliance, &value.ChangeType)
		// fmt.Println(err)
		if len(value.Parentcatname) > 0 {
			value.Categoryname = value.Catname + "(" + value.Parentcatname + ")"
		} else {
			value.Categoryname = value.Catname
		}
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateMstbusinessmatrix(tz *entities.MstbusinessmatrixEntity) error {
	logger.Log.Println("In side UpdateMstbusinessmatrix")
	stmt, err := dbc.DB.Prepare(updateMstbusinessmatrix)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMstbusinessmatrix Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Mstrecorddifferentiationtickettypeid, tz.Mstrecorddifferentiationcatid, tz.Mstrecorddifferentiationimpactid, tz.Mstrecorddifferentiationurgencyid, tz.Mstrecorddifferentiationpriorityid, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMstbusinessmatrix Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) UpdateEstimatedTime(tz *entities.MstbusinessmatrixEntity) error {
	logger.Log.Println("In side UpdateEstimatedTime")
	var sql = "UPDATE mapcategorywithestimatetime SET estimatedtime=?,efficiency=? WHERE recorddiffid=? AND clientid=? AND mstorgnhirarchyid=?"
	stmt, err := dbc.DB.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMstbusinessmatrix Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Estimatedeffort, tz.Slacompliance, tz.Mstrecorddifferentiationcatid, tz.Clientid, tz.Mstorgnhirarchyid)
	if err != nil {
		logger.Log.Println("UpdateMstbusinessmatrix Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMstbusinessmatrix(tz *entities.MstbusinessmatrixEntity) error {
	logger.Log.Println("In side DeleteMstbusinessmatrix")
	stmt, err := dbc.DB.Prepare(deleteMstbusinessmatrix)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMstbusinessmatrix Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMstbusinessmatrix Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteEstimatedTimeeffort(tz *entities.MstbusinessmatrixEntity) error {
	logger.Log.Println("In side DeleteMstbusinessmatrix")
	var sql = "UPDATE mapcategorywithestimatetime SET deleteflg=1 WHERE recorddiffid  in (SELECT mstrecorddifferentiationcatid FROM mstbusinessmatrix WHERE id=?) AND id>0"
	stmt, err := dbc.DB.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMstbusinessmatrix Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMstbusinessmatrix Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMstbusinessmatrixCount(tz *entities.MstbusinessmatrixEntity, OrgnTypeID int64) (entities.MstbusinessmatrixEntities, error) {
	logger.Log.Println("In side GetMstbusinessmatrixCount")
	value := entities.MstbusinessmatrixEntities{}
	var getMstbusinessmatrixcount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getMstbusinessmatrixcount = "SELECT count(id) total FROM  mstbusinessmatrix WHERE  deleteflg =0 and activeflg=1"
	} else if OrgnTypeID == 2 {
		getMstbusinessmatrixcount = "SELECT count(id) total FROM  mstbusinessmatrix WHERE clientid = ? AND   deleteflg =0 and activeflg=1"
		params = append(params, tz.Clientid)
	} else {
		getMstbusinessmatrixcount = "SELECT count(id) total FROM  mstbusinessmatrix WHERE clientid = ? AND mstorgnhirarchyid = ? AND  deleteflg =0 and activeflg=1"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getMstbusinessmatrixcount, params...).Scan(&value.Total)

	// err := dbc.DB.QueryRow(getMstbusinessmatrixcount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMstbusinessmatrixCount Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) Checkmatrixconfig(tz *entities.MstbusinessmatrixEntity) (int64, error) {
	logger.Log.Println("In side Checkmatrixconfig")
	stmt, err := dbc.DB.Prepare(checkmatrixconfig)
	defer stmt.Close()
	var direction int64
	if err != nil {
		logger.Log.Println("Checkmatrixconfig Prepare Statement  Error", err)
		return 0, err
	}
	rows, err := stmt.Query(tz.Clientid, tz.Mstorgnhirarchyid, tz.Mstrecorddifferentickettypeid, tz.Mstrecorddifferentiationtickettypeid)
	if err != nil {
		logger.Log.Println("Checkmatrixconfig Execute Statement  Error", err)
		return 0, err
	}
	for rows.Next() {
		if err := rows.Scan(&direction); err != nil {
			logger.Log.Println("Error in fetching data")
		}

	}
	return direction, nil
}

func (dbc DbConn) Getlastlevelcategoryname(page *entities.MstbusinessmatrixEntity) ([]entities.MstlastlevelEntity, error) {
	logger.Log.Println("In side GelAllMstbusinessmatrix")
	values := []entities.MstlastlevelEntity{}
	//rows, err := dbc.DB.Query(lastlevelcategory, page.Clientid, page.Mstorgnhirarchyid, page.Mstrecorddifferentickettypeid, page.Mstrecorddifferentiationtickettypeid, page.Clientid, page.Mstorgnhirarchyid, page.Clientid, page.Mstorgnhirarchyid)
	rows, err := dbc.DB.Query(lastlevelcategory, page.Clientid, page.Mstorgnhirarchyid, page.Mstrecorddifferentickettypeid, page.Mstrecorddifferentiationtickettypeid, page.Clientid, page.Mstorgnhirarchyid, page.Mstrecorddifferentickettypeid, page.Mstrecorddifferentiationtickettypeid)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstbusinessmatrix Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstlastlevelEntity{}
		rows.Scan(&value.Id, &value.Catname, &value.Lastcategorylevelid, &value.Parentcatname)
		if len(value.Parentcatname) > 0 {
			value.Name = value.Catname + " (" + value.Parentcatname + ")"
		} else {
			value.Name = value.Catname
		}
		values = append(values, value)
	}
	return values, nil
}
