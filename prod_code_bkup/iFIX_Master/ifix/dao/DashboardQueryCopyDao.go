package dao

import (
	"database/sql"
	"errors"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
)

var insertDashboardQueryCopy = "INSERT INTO mstdashboarddtls (clientid, mstorgnhirarchyid, mstrecorddifferentiationid, mapfunctionalityid, querytype, query, queryparam, joinquery) VALUES (?,?,?,?,?,?,?,?)"
var duplicateDashboardQueryCopy = "SELECT count(id) total FROM  mstdashboarddtls WHERE clientid = ? AND mstorgnhirarchyid = ?  AND mstrecorddifferentiationid=? AND mapfunctionalityid=? AND querytype=? AND activeflg =1 AND deleteflg = 0 "
var checktile = "SELECT count(id) total FROM  mapfunctionality WHERE clientid = ? AND mstorgnhirarchyid = ?  AND funcid=1 AND  funcdescid=? AND activeflg =1 AND deleteflg = 0 "
var gettilesname = "SELECT description  FROM  mapfunctionality WHERE clientid = ? AND mstorgnhirarchyid = ?  AND funcid=1 AND  funcdescid=? AND activeflg =1 AND deleteflg = 0 "

var getDashboardQueryCopy = "SELECT a.id as id,a.clientid as clientid,a.mstorgnhirarchyid as mstorgnhirarchyid,a.mstrecorddifferentiationid as recorddiffid,a.mapfunctionalityid as tilesid,coalesce(a.querytype,'N/A') as querytype,coalesce(a.query,'N/A') as query,coalesce(a.queryparam,'N/A') as queryparam,coalesce(a.joinquery,'N/A') as joinquery,b.name as clientname,c.name as mstorgnhirarchyname,coalesce(d.name,'N/A') as recorddiffname,coalesce(e.description,'N/A') as tilesname,IF(querytype=1,'Count','details') querytypename,a.ismanegerialview,IF(a.ismanegerialview=1,'My Workspace','Team Workspace') FROM mstdashboarddtls a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiation d,mapfunctionality e  WHERE a.clientid = b.id AND a.mstorgnhirarchyid = c.id AND a.mstrecorddifferentiationid = d.id AND a.mapfunctionalityid = e.funcdescid AND a.activeflg = 1 AND a.deleteflg = 0 AND d.activeflg = 1 AND d.deleteflg = 0  AND e.funcid=1 AND e.clientid=1 AND e.mstorgnhirarchyid=1 AND e.activeflg = 1 AND e.deleteflg = 0 ORDER BY a.id DESC LIMIT ?,?;"
var getDashboardQueryCopycount = "SELECT count(a.id) as total FROM mstdashboarddtls a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiation d,mapfunctionality e  WHERE a.clientid = b.id AND a.mstorgnhirarchyid = c.id AND a.mstrecorddifferentiationid = d.id AND a.mapfunctionalityid = e.funcdescid AND a.activeflg = 1 AND a.deleteflg = 0 AND d.activeflg = 1 AND d.deleteflg = 0  AND e.funcid=1 AND e.clientid=1 AND e.mstorgnhirarchyid=1 AND e.activeflg = 1 AND e.deleteflg = 0 ;"

var getDashboardValuesCopy = "SELECT a.query as Query,a.queryparam as QueryParam ,a.joinquery as JoinQuery,b.description tilesname FROM mstdashboarddtls a,mapfunctionality b where a.clientid=? AND a.mstorgnhirarchyid=? AND a.mstrecorddifferentiationid=? AND a.mapfunctionalityid=? AND a.querytype=? AND a.activeflg=1 AND a.deleteflg=0 and a.mapfunctionalityid = b.funcdescid and b.clientid=1 AND b.mstorgnhirarchyid=1 AND b.activeflg = 1 AND b.deleteflg = 0"

//var updateDashboardQueryCopy = "UPDATE mstexceltemplate SET clientid=?,mstorgnhirarchyid = ?, recorddifftypeid = ?, headername = ?, seqno = ?,templatetypeid=?,recorddiffid=? WHERE id = ? "
var deleteDashboardQueryCopy = "UPDATE mstdashboarddtls SET deleteflg ='1' WHERE id = ? "

func (dbc DbConn) Gettilesname(page *entities.DashboardQueryCopyEntity, i int) (entities.DashboardQueryCopyEntity, error) {
	logger.Log.Println("In side GetDashboardQueryCopy")
	values := entities.DashboardQueryCopyEntity{}
	err := dbc.DB.QueryRow(gettilesname, page.Clientid, page.Mstorgnhirarchyid, page.Tilesids[i]).Scan(&values.TilesName)

	//rows, err := dbc.DB.Query(getDashboardValuesCopy, page.Clientid, page.Mstorgnhirarchyid, page.RecordDiffid, page.Tilesid, page.QueryType)
	//defer rows.Close()
	switch err {
	case sql.ErrNoRows:
		logger.Log.Println(" No Rows with this information", err)
		return values, errors.New("ERROR: No Rows With This Information")
	case nil:
		return values, nil
	default:
		logger.Log.Println("GetDashboardQueryCopy Get Statement Prepare Error", err)
		return values, err
	}
}
func (dbc DbConn) Checktiles(tz *entities.DashboardQueryCopyEntity, i int) (entities.DashboardQueryCopyEntities, error) {
	logger.Log.Println("In side CheckDuplicateDashboardQueryCopy ")
	value := entities.DashboardQueryCopyEntities{}
	err := dbc.DB.QueryRow(checktile, tz.ToClientid, tz.ToMstorgnhirarchyid, tz.Tilesids[i]).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateDashboardQueryCopy Get Statement Prepare Error", err)
		return value, err
	}
}
func (dbc DbConn) CheckDuplicateDashboardQueryCopy(tz *entities.DashboardQueryCopyEntity) (entities.DashboardQueryCopyEntities, error) {
	logger.Log.Println("In side CheckDuplicateDashboardQueryCopy ")
	value := entities.DashboardQueryCopyEntities{}
	err := dbc.DB.QueryRow(duplicateDashboardQueryCopy, tz.ToClientid, tz.ToMstorgnhirarchyid, tz.ToRecordDiffid, tz.Tilesid, tz.QueryType).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateDashboardQueryCopy Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc TxConn) AddDashboardQueryCopy(tz *entities.DashboardQueryCopyEntity) (int64, error) {
	logger.Log.Println("In side AddDashboardQueryCopy")
	logger.Log.Println("Query -->", insertDashboardQueryCopy)
	stmt, err := dbc.TX.Prepare(insertDashboardQueryCopy)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("AddDashboardQueryCopy Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.ToClientid, tz.ToMstorgnhirarchyid, tz.ToRecordDiffid, tz.Tilesid, tz.QueryType, tz.Query, tz.QueryParam, tz.JoinQuery) //, tz.Query.(string), tz.QueryParam.(string), tz.JoinQuery.(string))
	res, err := stmt.Exec(tz.ToClientid, tz.ToMstorgnhirarchyid, tz.ToRecordDiffid, tz.Tilesid, tz.QueryType, tz.Query, tz.QueryParam, tz.JoinQuery)
	if err != nil {
		logger.Log.Println("AddDashboardQueryCopy Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetDashboardQueryCopy(page *entities.DashboardQueryCopyEntity, i int) (entities.DashboardQueryCopyEntity, error) {
	logger.Log.Println("In side GetDashboardQueryCopy")
	values := entities.DashboardQueryCopyEntity{}
	err := dbc.DB.QueryRow(getDashboardValuesCopy, page.Clientid, page.Mstorgnhirarchyid, page.RecordDiffid, page.Tilesids[i], page.QueryType).Scan(&values.Query, &values.QueryParam, &values.JoinQuery, &values.TilesName)

	//rows, err := dbc.DB.Query(getDashboardValuesCopy, page.Clientid, page.Mstorgnhirarchyid, page.RecordDiffid, page.Tilesid, page.QueryType)
	//defer rows.Close()
	switch err {
	case sql.ErrNoRows:
		logger.Log.Println(" No Rows with this information", err)
		return values, errors.New("ERROR: No Rows With This Information")
	case nil:
		return values, nil
	default:
		logger.Log.Println("GetDashboardQueryCopy Get Statement Prepare Error", err)
		return values, err
	}

	// if err == sql.ErrNoRows {
	// 	logger.Log.Println(" No Rows with this information", err)
	// 	return values, errors.New("ERROR: No Rows With This Information")
	// }
	// if err != nil {
	// 	logger.Log.Println("GetDashboardQueryCopy Get Statement Prepare Error", err)
	// 	return values, err
	// }

	// for rows.Next() {
	// 	//value := entities.DashboardQueryCopyEntity{}
	// 	rows.Scan(&values.Query, &values.QueryParam, &values.JoinQuery)
	// 	//values = append(values, value)
	// }
	// return values, nil
}

func (dbc DbConn) GetAllDashboardQueryCopy(tz *entities.DashboardQueryCopyEntity, OrgnType int64) ([]entities.DashboardQueryCopyEntity, error) {
	logger.Log.Println("In side GetAllDashboardQueryCopy")
	values := []entities.DashboardQueryCopyEntity{}
	var getDashboardQueryCopy string
	var params []interface{}
	if OrgnType == 1 {
		getDashboardQueryCopy = "SELECT a.id as id,a.clientid as clientid,a.mstorgnhirarchyid as mstorgnhirarchyid,coalesce(a.mstrecorddifferentiationid) AS recorddiffid,a.mapfunctionalityid as tilesid,coalesce(a.querytype,'N/A') as querytype,coalesce(a.query,'N/A') as query,coalesce(a.queryparam,'N/A') as queryparam,coalesce(a.joinquery,'N/A') as joinquery,b.name as clientname,c.name as mstorgnhirarchyname,coalesce(d.name,'N/A') as recorddiffname,coalesce(e.description,'N/A') as tilesname,IF(querytype=1,'Count','details') querytypename,e.ismanegerialview,IF(e.ismanegerialview=1,'My Workspace','Team Workspace') FROM mstdashboarddtls a LEFT JOIN mstrecorddifferentiation d ON a.mstrecorddifferentiationid = d.id AND d.activeflg = 1 AND d.deleteflg = 0,mstclient b,mstorgnhierarchy c,mapfunctionality e  WHERE   a.clientid = b.id AND a.mstorgnhirarchyid = c.id AND  a.mapfunctionalityid = e.funcdescid AND a.activeflg = 1 AND a.deleteflg = 0 AND e.funcid=1 AND e.clientid=a.clientid AND e.mstorgnhirarchyid=a.mstorgnhirarchyid AND e.activeflg = 1 AND e.deleteflg = 0 ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		getDashboardQueryCopy = "SELECT a.id as id,a.clientid as clientid,a.mstorgnhirarchyid as mstorgnhirarchyid,coalesce(a.mstrecorddifferentiationid) AS recorddiffid,a.mapfunctionalityid as tilesid,coalesce(a.querytype,'N/A') as querytype,coalesce(a.query,'N/A') as query,coalesce(a.queryparam,'N/A') as queryparam,coalesce(a.joinquery,'N/A') as joinquery,b.name as clientname,c.name as mstorgnhirarchyname,coalesce(d.name,'N/A') as recorddiffname,coalesce(e.description,'N/A') as tilesname,IF(querytype=1,'Count','details') querytypename,e.ismanegerialview,IF(e.ismanegerialview=1,'My Workspace','Team Workspace') FROM mstdashboarddtls a LEFT JOIN mstrecorddifferentiation d ON a.mstrecorddifferentiationid = d.id AND d.activeflg = 1 AND d.deleteflg = 0,mstclient b,mstorgnhierarchy c,mapfunctionality e  WHERE a.clientid=?  and a.clientid = b.id AND a.mstorgnhirarchyid = c.id  AND a.mapfunctionalityid = e.funcdescid AND a.activeflg = 1 AND a.deleteflg = 0 AND e.funcid=1 AND e.clientid=a.clientid AND e.mstorgnhirarchyid=a.mstorgnhirarchyid AND e.activeflg = 1 AND e.deleteflg = 0 ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		getDashboardQueryCopy = "SELECT a.id as id,a.clientid as clientid,a.mstorgnhirarchyid as mstorgnhirarchyid,coalesce(a.mstrecorddifferentiationid) AS recorddiffid,a.mapfunctionalityid as tilesid,coalesce(a.querytype,'N/A') as querytype,coalesce(a.query,'N/A') as query,coalesce(a.queryparam,'N/A') as queryparam,coalesce(a.joinquery,'N/A') as joinquery,b.name as clientname,c.name as mstorgnhirarchyname,coalesce(d.name,'N/A') as recorddiffname,coalesce(e.description,'N/A') as tilesname,IF(querytype=1,'Count','details') querytypename,e.ismanegerialview,IF(e.ismanegerialview=1,'My Workspace','Team Workspace') FROM mstdashboarddtls a LEFT JOIN mstrecorddifferentiation d ON a.mstrecorddifferentiationid = d.id AND d.activeflg = 1 AND d.deleteflg = 0,mstclient b,mstorgnhierarchy c,mapfunctionality e  WHERE a.clientid=? and a.mstorgnhirarchyid=? and a.clientid = b.id AND a.mstorgnhirarchyid = c.id AND a.mapfunctionalityid = e.funcdescid AND a.activeflg = 1 AND a.deleteflg = 0 AND e.funcid=1 AND e.clientid=a.clientid AND e.mstorgnhirarchyid=a.mstorgnhirarchyid AND e.activeflg = 1 AND e.deleteflg = 0 ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}
	rows, err := dbc.DB.Query(getDashboardQueryCopy, params...)
	// rows, err := dbc.DB.Query(getDashboardQueryCopy, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllDashboardQueryCopy Get Statement Prepare Error", err)
		log.Println("GetAllDashboardQueryCopy Get Statement Prepare Error", err)
		return values, err
	}

	for rows.Next() {
		value := entities.DashboardQueryCopyEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.RecordDiffid, &value.Tilesid, &value.QueryType, &value.Query, &value.QueryParam, &value.JoinQuery, &value.Clientname, &value.Mstorgnhirarchyname, &value.RecordDiffName, &value.TilesName, &value.QueryTypename, &value.Ismanagerialview, &value.IsmanagerialviewName)
		values = append(values, value)
	}
	//fmt.Println(values)
	return values, nil
}

// func (dbc DbConn) UpdateDashboardQueryCopy(tz *entities.DashboardQueryCopyEntity) error {
// 	logger.Log.Println("In side UpdateDashboardQueryCopy")
// 	stmt, err := dbc.DB.Prepare(updateDashboardQueryCopy)
// 	defer stmt.Close()
// 	if err != nil {
// 		logger.Log.Println("UpdateDashboardQueryCopy Prepare Statement  Error", err)
// 		return err
// 	}
// 	_, err = stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.RecordDiffTypeid, tz.HeaderName, tz.SeqNo, tz.TemplateTypeid, tz.RecordDiffid, tz.Id)
// 	if err != nil {
// 		logger.Log.Println("UpdateDashboardQueryCopy Execute Statement  Error", err)
// 		return err
// 	}
// 	return nil
// }

func (dbc DbConn) DeleteDashboardQueryCopy(tz *entities.DashboardQueryCopyEntity) error {
	logger.Log.Println("In side DeleteDashboardQueryCopy", tz)
	stmt, err := dbc.DB.Prepare(deleteDashboardQueryCopy)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteDashboardQueryCopy Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteDashboardQueryCopy Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetDashboardQueryCopyCount(tz *entities.DashboardQueryCopyEntity, OrgnTypeID int64) (entities.DashboardQueryCopyEntities, error) {
	logger.Log.Println("In side GetDashboardQueryCopyCount")
	value := entities.DashboardQueryCopyEntities{}
	var getDashboardQueryCopycount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getDashboardQueryCopycount = "SELECT count(a.id) as total FROM mstdashboarddtls a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiation d,mapfunctionality e  WHERE a.clientid = b.id AND a.mstorgnhirarchyid = c.id AND a.mstrecorddifferentiationid = d.id AND a.mapfunctionalityid = e.funcdescid AND a.activeflg = 1 AND a.deleteflg = 0 AND d.activeflg = 1 AND d.deleteflg = 0  AND e.funcid=1 AND e.clientid=a.clientid AND e.mstorgnhirarchyid=a.mstorgnhirarchyid AND e.activeflg = 1 AND e.deleteflg = 0"
	} else if OrgnTypeID == 2 {
		getDashboardQueryCopycount = "SELECT count(a.id) as total FROM mstdashboarddtls a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiation d,mapfunctionality e  WHERE a.clientid=? and a.clientid = b.id AND a.mstorgnhirarchyid = c.id AND a.mstrecorddifferentiationid = d.id AND a.mapfunctionalityid = e.funcdescid AND a.activeflg = 1 AND a.deleteflg = 0 AND d.activeflg = 1 AND d.deleteflg = 0  AND e.funcid=1 AND e.clientid=a.clientid AND e.mstorgnhirarchyid=a.mstorgnhirarchyid AND e.activeflg = 1 AND e.deleteflg = 0 "
		params = append(params, tz.Clientid)
	} else {
		getDashboardQueryCopycount = "SELECT count(a.id) as total FROM mstdashboarddtls a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiation d,mapfunctionality e  WHERE a.clientid=? and a.mstorgnhirarchyid=? and a.clientid = b.id AND a.mstorgnhirarchyid = c.id AND a.mstrecorddifferentiationid = d.id AND a.mapfunctionalityid = e.funcdescid AND a.activeflg = 1 AND a.deleteflg = 0 AND d.activeflg = 1 AND d.deleteflg = 0  AND e.funcid=1 AND e.clientid=a.clientid AND e.mstorgnhirarchyid=a.mstorgnhirarchyid AND e.activeflg = 1 AND e.deleteflg = 0"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getDashboardQueryCopycount, params...).Scan(&value.Total)
	// err := dbc.DB.QueryRow(getDashboardQueryCopycount).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetDashboardQueryCopyCount Get Statement Prepare Error", err)
		return value, err
	}
}
