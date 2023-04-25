package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertDashboardQuery = "INSERT INTO mstdashboarddtls (clientid, mstorgnhirarchyid, mstrecorddifferentiationid, mapfunctionalityid, querytype, query, queryparam, joinquery) VALUES (?,?,?,?,?,?,?,?)"
var duplicateDashboardQuery = "SELECT count(id) total FROM  mstdashboarddtls WHERE clientid = ? AND mstorgnhirarchyid = ?  AND mstrecorddifferentiationid=? AND mapfunctionalityid=? AND querytype=?  AND activeflg =1 AND deleteflg = 0 "

var updateDashboardQuery = "UPDATE mstdashboarddtls SET mstorgnhirarchyid = ?, mstrecorddifferentiationid = ?, mapfunctionalityid = ?, querytype = ?,query=?,queryparam=?,joinquery=? WHERE id = ? "

func (dbc DbConn) CheckDuplicateDashboardQuery(tz *entities.DashboardQueryEntity) (entities.DashboardQueryEntities, error) {
	logger.Log.Println("In side CheckDuplicateDashboardQuery ")
	value := entities.DashboardQueryEntities{}
	err := dbc.DB.QueryRow(duplicateDashboardQuery, tz.Clientid, tz.Mstorgnhirarchyid, tz.RecordDiffid, tz.Tilesid, tz.QueryType).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		logger.Log.Println(value.Total)
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateDashboardQuery Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) AddDashboardQuery(tz *entities.DashboardQueryEntity) (int64, error) {
	logger.Log.Println("In side AddMstExcelTemplate")
	logger.Log.Println("Query -->", insertDashboardQuery)
	stmt, err := dbc.DB.Prepare(insertDashboardQuery)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("AddDashboardQueryCopy Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.RecordDiffid, tz.Tilesid, tz.QueryType, tz.Query, tz.QueryParam, tz.JoinQuery)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.RecordDiffid, tz.Tilesid, tz.QueryType, tz.Query, tz.QueryParam, tz.JoinQuery)
	if err != nil {
		logger.Log.Println("AddDashboardQueryCopy Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) UpdateDashboardQuery(tz *entities.DashboardQueryEntity) error {
	logger.Log.Println("In side UpdateDashboardQuery")
	stmt, err := dbc.DB.Prepare(updateDashboardQuery)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateDashboardQuery Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.RecordDiffid, tz.Tilesid, tz.QueryType, tz.Query, tz.QueryParam, tz.JoinQuery, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateDashboardQuery Execute Statement  Error", err)
		return err
	}
	return nil
}
