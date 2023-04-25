package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertDashboarddtls = "INSERT INTO mstdashboarddtls (clientid, mstorgnhirarchyid, mstrecorddifferentiationid, mapfunctionalityid, querytype, query, queryparam) VALUES (?,?,?,?,?,?,?)"
var duplicateDashboarddtls = "SELECT count(id) total FROM  mstdashboarddtls WHERE clientid = ? AND mstorgnhirarchyid = ? AND mstrecorddifferentiationid = ? AND mapfunctionalityid = ? AND querytype=? AND deleteflg = 0 AND activeflg=1"

// var getDashboarddtls = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.mstrecorddifferentiationid as Mstrecorddifferentiationid, a.mapfunctionalityid as Mapfunctionalityid, a.querytype as Querytype, a.query as Query, a.queryparam as Queryparam, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.name as Recorddifferentiationname,e.description as Mapfunctionalityname FROM mstdashboarddtls a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiation d,mapfunctionality e WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.mstrecorddifferentiationid=d.id and a.mapfunctionalityid=e.funcdescid AND e.clientid = ? AND e.mstorgnhirarchyid = ? AND e.funcid=1 ORDER BY a.id DESC LIMIT ?,?"
// var getDashboarddtlscount = "SELECT count(a.id) as total FROM mstdashboarddtls a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiation d,mapfunctionality e WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.mstrecorddifferentiationid=d.id and a.mapfunctionalityid=e.id  AND e.funcid=1"
var updateDashboarddtls = "UPDATE mstdashboarddtls SET mstorgnhirarchyid = ?, mstrecordtypeid = ?, mapfunctionalityid = ?, querytype = ?, query = ?, queryparam = ? WHERE id = ? "
var deleteDashboarddtls = "UPDATE mstdashboarddtls SET deleteflg = '1' WHERE id = ? "

func (dbc DbConn) CheckDuplicateDashboarddtls(tz *entities.DashboarddtlsEntity) (entities.DashboarddtlsEntities, error) {
	logger.Log.Println("In side CheckDuplicateDashboarddtls")
	value := entities.DashboarddtlsEntities{}
	err := dbc.DB.QueryRow(duplicateDashboarddtls, tz.Clientid, tz.Mstorgnhirarchyid, tz.Mstrecorddifferentiationid, tz.Mapfunctionalityid, tz.Querytype).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateDashboarddtls Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertDashboarddtls(tz *entities.DashboarddtlsEntity) (int64, error) {
	logger.Log.Println("In side InsertDashboarddtls")
	logger.Log.Println("Query -->", insertDashboarddtls)
	stmt, err := dbc.DB.Prepare(insertDashboarddtls)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertDashboarddtls Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Mstrecorddifferentiationid, tz.Mapfunctionalityid, tz.Querytype, tz.Query, tz.Queryparam)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Mstrecorddifferentiationid, tz.Mapfunctionalityid, tz.Querytype, tz.Query, tz.Queryparam)
	if err != nil {
		logger.Log.Println("InsertDashboarddtls Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllDashboarddtls(tz *entities.DashboarddtlsEntity, OrgnType int64) ([]entities.DashboarddtlsEntity, error) {
	logger.Log.Println("In side GelAllDashboarddtls")
	values := []entities.DashboarddtlsEntity{}
	var getDashboarddtls string
	var params []interface{}
	if OrgnType == 1 {
		getDashboarddtls = "SELECT distinct a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.mstrecorddifferentiationid as Mstrecorddifferentiationid, a.mapfunctionalityid as Mapfunctionalityid, a.querytype as Querytype, a.query as Query, a.queryparam as Queryparam, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.name as Recorddifferentiationname,e.description as Mapfunctionalityname FROM mstdashboarddtls a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiation d,mapfunctionality e WHERE a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.mstrecorddifferentiationid=d.id and a.mapfunctionalityid=e.funcdescid AND e.funcid=1 and a.clientid=e.clientid and a.mstorgnhirarchyid=e.mstorgnhirarchyid ORDER BY a.id DESC LIMIT ?,?;"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		getDashboarddtls = "SELECT distinct a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.mstrecorddifferentiationid as Mstrecorddifferentiationid, a.mapfunctionalityid as Mapfunctionalityid, a.querytype as Querytype, a.query as Query, a.queryparam as Queryparam, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.name as Recorddifferentiationname,e.description as Mapfunctionalityname FROM mstdashboarddtls a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiation d,mapfunctionality e WHERE a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.mstrecorddifferentiationid=d.id and a.mapfunctionalityid=e.funcdescid AND e.funcid=1 and a.clientid=e.clientid and a.mstorgnhirarchyid=e.mstorgnhirarchyid and a.clientid=? ORDER BY a.id DESC LIMIT ?,?;"
		params = append(params, tz.Clientid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		getDashboarddtls = "SELECT distinct a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.mstrecorddifferentiationid as Mstrecorddifferentiationid, a.mapfunctionalityid as Mapfunctionalityid, a.querytype as Querytype, a.query as Query, a.queryparam as Queryparam, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.name as Recorddifferentiationname,e.description as Mapfunctionalityname FROM mstdashboarddtls a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiation d,mapfunctionality e WHERE a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.mstrecorddifferentiationid=d.id and a.mapfunctionalityid=e.funcdescid AND e.funcid=1 and a.clientid=e.clientid and a.mstorgnhirarchyid=e.mstorgnhirarchyid and a.clientid=? and a.mstorgnhirarchyid=? ORDER BY a.id DESC LIMIT ?,?;"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}

	rows, err := dbc.DB.Query(getDashboarddtls, params...)

	// rows, err := dbc.DB.Query(getDashboarddtls, page.Clientid, page.Mstorgnhirarchyid, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllDashboarddtls Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.DashboarddtlsEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Mstrecorddifferentiationid, &value.Mapfunctionalityid, &value.Querytype, &value.Query, &value.Queryparam, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Recorddifferentiationname, &value.Mapfunctionalityname)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateDashboarddtls(tz *entities.DashboarddtlsEntity) error {
	logger.Log.Println("In side UpdateDashboarddtls")
	stmt, err := dbc.DB.Prepare(updateDashboarddtls)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateDashboarddtls Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Mstrecorddifferentiationid, tz.Mapfunctionalityid, tz.Querytype, tz.Query, tz.Queryparam, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateDashboarddtls Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteDashboarddtls(tz *entities.DashboarddtlsEntity) error {
	logger.Log.Println("In side DeleteDashboarddtls")
	stmt, err := dbc.DB.Prepare(deleteDashboarddtls)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteDashboarddtls Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteDashboarddtls Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetDashboarddtlsCount(tz *entities.DashboarddtlsEntity, OrgnTypeID int64) (entities.DashboarddtlsEntities, error) {
	logger.Log.Println("In side GetDashboarddtlsCount")
	value := entities.DashboarddtlsEntities{}
	var getDashboarddtlscount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getDashboarddtlscount = " SELECT count(distinct a.id) as total FROM  mstdashboarddtls a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiation d,mapfunctionality e WHERE a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.mstrecorddifferentiationid=d.id and a.mapfunctionalityid=e.funcdescid AND e.funcid=1 and a.clientid=e.clientid and a.mstorgnhirarchyid=e.mstorgnhirarchyid"
	} else if OrgnTypeID == 2 {
		logger.Log.Println("inside client")
		getDashboarddtlscount = " SELECT count(distinct a.id) as total FROM  mstdashboarddtls a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiation d,mapfunctionality e WHERE a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.mstrecorddifferentiationid=d.id and a.mapfunctionalityid=e.funcdescid AND e.funcid=1 and a.clientid=e.clientid and a.mstorgnhirarchyid=e.mstorgnhirarchyid and a.clientid=?"
		params = append(params, tz.Clientid)
	} else {
		getDashboarddtlscount = " SELECT count(distinct a.id) as total FROM  mstdashboarddtls a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiation d,mapfunctionality e WHERE a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.mstrecorddifferentiationid=d.id and a.mapfunctionalityid=e.funcdescid AND e.funcid=1 and a.clientid=e.clientid and a.mstorgnhirarchyid=e.mstorgnhirarchyid and a.clientid=? and a.mstorgnhirarchyid=?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getDashboarddtlscount, params...).Scan(&value.Total)

	// err := dbc.DB.QueryRow(getDashboarddtlscount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetDashboarddtlsCount Get Statement Prepare Error", err)
		return value, err
	}
}
