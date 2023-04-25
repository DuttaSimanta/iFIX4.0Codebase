package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertMstslastate = "INSERT INTO mstslastate (clientid, mstorgnhirarchyid, statename) VALUES (?,?,?)"
var duplicateMstslastate = "SELECT count(id) total FROM  mstslastate WHERE clientid = ? AND mstorgnhirarchyid = ? AND statename = ? AND deleteflg = 0"

// var getMstslastate = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.statename as Statename, a.activeflg as Activeflg,(select name from mstclient where id = a.clientid ) as Clientname,(select name from mstorgnhierarchy where id = a.mstorgnhirarchyid ) as Mstorgnhirarchyname FROM mstslastate a WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 ORDER BY a.id DESC LIMIT ?,?"
// var getMstslastatecount = "SELECT count(a.id) total FROM  mstslastate a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND  a.deleteflg =0 and a.activeflg=1 and a.clientid=b.id and a.mstorgnhirarchyid =c.id"
var updateMstslastate = "UPDATE mstslastate SET mstorgnhirarchyid = ?, statename = ? WHERE id = ? "
var deleteMstslastate = "UPDATE mstslastate SET deleteflg = '1' WHERE id = ? "

func (dbc DbConn) CheckDuplicateMstslastate(tz *entities.MstslastateEntity) (entities.MstslastateEntities, error) {
	logger.Log.Println("In side CheckDuplicateMstslastate")
	value := entities.MstslastateEntities{}
	err := dbc.DB.QueryRow(duplicateMstslastate, tz.Clientid, tz.Mstorgnhirarchyid, tz.Statename).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMstslastate Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertMstslastate(tz *entities.MstslastateEntity) (int64, error) {
	logger.Log.Println("In side InsertMstslastate")
	logger.Log.Println("Query -->", insertMstslastate)
	stmt, err := dbc.DB.Prepare(insertMstslastate)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMstslastate Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Statename)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Statename)
	if err != nil {
		logger.Log.Println("InsertMstslastate Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllMstslastate(page *entities.MstslastateEntity, OrgnType int64) ([]entities.MstslastateEntity, error) {
	logger.Log.Println("In side GelAllMstslastate")
	values := []entities.MstslastateEntity{}
	var getMstslastate string
	var params []interface{}
	if OrgnType == 1 {
		getMstslastate = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.statename as Statename, a.activeflg as Activeflg,(select name from mstclient where id = a.clientid ) as Clientname,(select name from mstorgnhierarchy where id = a.mstorgnhirarchyid ) as Mstorgnhirarchyname FROM mstslastate a WHERE a.deleteflg =0 and a.activeflg=1 ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else if OrgnType == 2 {
		getMstslastate = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.statename as Statename, a.activeflg as Activeflg,(select name from mstclient where id = a.clientid ) as Clientname,(select name from mstorgnhierarchy where id = a.mstorgnhirarchyid ) as Mstorgnhirarchyname FROM mstslastate a WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else {
		getMstslastate = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.statename as Statename, a.activeflg as Activeflg,(select name from mstclient where id = a.clientid ) as Clientname,(select name from mstorgnhierarchy where id = a.mstorgnhirarchyid ) as Mstorgnhirarchyname FROM mstslastate a WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Mstorgnhirarchyid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	}
	rows, err := dbc.DB.Query(getMstslastate, params...)

	//rows, err := dbc.DB.Query(getMstslastate, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstslastate Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstslastateEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Statename, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateMstslastate(tz *entities.MstslastateEntity) error {
	logger.Log.Println("In side UpdateMstslastate")
	stmt, err := dbc.DB.Prepare(updateMstslastate)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMstslastate Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Statename, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMstslastate Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMstslastate(tz *entities.MstslastateEntity) error {
	logger.Log.Println("In side DeleteMstslastate")
	stmt, err := dbc.DB.Prepare(deleteMstslastate)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMstslastate Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMstslastate Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMstslastateCount(tz *entities.MstslastateEntity, OrgnTypeID int64) (entities.MstslastateEntities, error) {
	logger.Log.Println("In side GetMstslastateCount")
	value := entities.MstslastateEntities{}
	var getMstslastatecount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getMstslastatecount = "SELECT count(a.id) total FROM  mstslastate a,mstclient b,mstorgnhierarchy c WHERE  a.deleteflg =0 and a.activeflg=1 and a.clientid=b.id and a.mstorgnhirarchyid =c.id"
	} else if OrgnTypeID == 2 {
		getMstslastatecount = "SELECT count(a.id) total FROM  mstslastate a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND  a.deleteflg =0 and a.activeflg=1 and a.clientid=b.id and a.mstorgnhirarchyid =c.id"
		params = append(params, tz.Clientid)
	} else {
		getMstslastatecount = "SELECT count(a.id) total FROM  mstslastate a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND  a.deleteflg =0 and a.activeflg=1 and a.clientid=b.id and a.mstorgnhirarchyid =c.id"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getMstslastatecount, params...).Scan(&value.Total)

	// err := dbc.DB.QueryRow(getMstslastatecount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMstslastateCount Get Statement Prepare Error", err)
		return value, err
	}
}
