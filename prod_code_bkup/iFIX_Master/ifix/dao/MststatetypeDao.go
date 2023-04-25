package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertMststatetype = "INSERT INTO mststatetype (clientid, mstorgnhirarchyid, statetypename) VALUES (?,?,?)"
var duplicateMststatetype = "SELECT count(id) total FROM  mststatetype WHERE clientid = ? AND mstorgnhirarchyid = ? AND statetypename = ? AND deleteflg = 0"
var getMststatetype = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.statetypename as Statetypename, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname FROM mststatetype a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id ORDER BY a.id DESC LIMIT ?,?"
//var getMststatetypecount = "SELECT count(a.id) as total FROM mststatetype a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id"
var updateMststatetype = "UPDATE mststatetype SET clientid = ?, mstorgnhirarchyid = ?, statetypename = ? WHERE id = ? "
var deleteMststatetype = "UPDATE mststatetype SET deleteflg = '1' WHERE id = ? "

func (dbc DbConn) CheckDuplicateMststatetype(tz *entities.MststatetypeEntity) (entities.MststatetypeEntities, error) {
	logger.Log.Println("In side CheckDuplicateMststatetype")
	value := entities.MststatetypeEntities{}
	err := dbc.DB.QueryRow(duplicateMststatetype, tz.Clientid, tz.Mstorgnhirarchyid, tz.Statetypename).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMststatetype Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertMststatetype(tz *entities.MststatetypeEntity) (int64, error) {
	logger.Log.Println("In side InsertMststatetype")
	logger.Log.Println("Query -->", insertMststatetype)
	stmt, err := dbc.DB.Prepare(insertMststatetype)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMststatetype Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Statetypename)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Statetypename)
	if err != nil {
		logger.Log.Println("InsertMststatetype Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

//func (dbc DbConn) GetAllMststatetype(page *entities.MststatetypeEntity) ([]entities.MststatetypeEntity, error) {
//	logger.Log.Println("In side GelAllMststatetype")
//	values := []entities.MststatetypeEntity{}
//	rows, err := dbc.DB.Query(getMststatetype, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
//	defer rows.Close()
//	if err != nil {
//		logger.Log.Println("GetAllMststatetype Get Statement Prepare Error", err)
//		return values, err
//	}
//	for rows.Next() {
//		value := entities.MststatetypeEntity{}
//		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Statetypename, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname)
//		values = append(values, value)
//	}
//	return values, nil
//}

func (dbc DbConn) UpdateMststatetype(tz *entities.MststatetypeEntity) error {
	logger.Log.Println("In side UpdateMststatetype")
	stmt, err := dbc.DB.Prepare(updateMststatetype)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMststatetype Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Statetypename, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMststatetype Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMststatetype(tz *entities.MststatetypeEntity) error {
	logger.Log.Println("In side DeleteMststatetype")
	stmt, err := dbc.DB.Prepare(deleteMststatetype)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMststatetype Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMststatetype Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMststatetypeCount(tz *entities.MststatetypeEntity, OrgnTypeID int64) (entities.MststatetypeEntities, error) {
	logger.Log.Println("In side GetMststatetypeCount")
	value := entities.MststatetypeEntities{}
	var getMststatetypecount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getMststatetypecount = "SELECT count(a.id) as total FROM mststatetype a,mstclient b,mstorgnhierarchy c WHERE a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id"
	} else if OrgnTypeID == 2 {
		getMststatetypecount = "SELECT count(a.id) as total FROM mststatetype a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id"
		params = append(params, tz.Clientid)
	} else {
		getMststatetypecount = "SELECT count(a.id) as total FROM mststatetype a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getMststatetypecount, params...).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMststatetypeCount Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) GetAllMststatetype(tz *entities.MststatetypeEntity, OrgnType int64) ([]entities.MststatetypeEntity, error) {
	logger.Log.Println("In side GelAllMststatetype")
	values := []entities.MststatetypeEntity{}
	var getMststatetype string
	var params []interface{}
	if OrgnType == 1 {
		getMststatetype = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.statetypename as Statetypename, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname FROM mststatetype a,mstclient b,mstorgnhierarchy c WHERE a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		getMststatetype = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.statetypename as Statetypename, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname FROM mststatetype a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		getMststatetype = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.statetypename as Statetypename, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname FROM mststatetype a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}
	rows, err := dbc.DB.Query(getMststatetype, params...)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMststatetype Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MststatetypeEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Statetypename, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname)
		values = append(values, value)
	}
	return values, nil
}
