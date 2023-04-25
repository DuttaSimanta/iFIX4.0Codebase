package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertMstorgcode = "INSERT INTO maporgcodewithtools (clientid, mstorgnhirarchyid, toolcode, orgcode) VALUES (?,?,?,?)"
var checkDuplicateMstorgcode = "SELECT count(id) total FROM  maporgcodewithtools WHERE clientid = ? AND mstorgnhirarchyid = ? AND toolcode = ? AND orgcode = ? AND deleteflg = 0"
var getAllMstorgcode = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.toolcode as Toolcode, a.orgcode as Orgcode, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname FROM maporgcodewithtools a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id ORDER BY a.id DESC LIMIT ?,?"

//var getMstorgcodecount = "SELECT count(a.id) as total FROM mststatetype a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id"
var updateMstorgcode = "UPDATE maporgcodewithtools SET clientid = ?, mstorgnhirarchyid = ?, toolcode = ?, orgcode = ? WHERE id = ? "
var deleteMstorgcode = "UPDATE maporgcodewithtools SET deleteflg = '1' WHERE id = ? "

func (dbc DbConn) CheckDuplicateMstorgcode(tz *entities.MstorgcodeEntity) (entities.MstorgcodeEntities, error) {
	logger.Log.Println("In side CheckDuplicateMstorgcode")
	value := entities.MstorgcodeEntities{}
	err := dbc.DB.QueryRow(checkDuplicateMstorgcode, tz.Clientid, tz.Mstorgnhirarchyid, tz.Toolcode, tz.Orgcode).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMstorgcode Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertMstorgcode(tz *entities.MstorgcodeEntity) (int64, error) {
	logger.Log.Println("In side InsertMstorgcode")
	// logger.Log.Println("Query -->", insertMstorgcode)
	stmt, err := dbc.DB.Prepare(insertMstorgcode)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMstorgcode Prepare Statement  Error", err)
		return 0, err
	}
	// logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Toolcode, tz.Orgcode)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Toolcode, tz.Orgcode)
	if err != nil {
		logger.Log.Println("InsertMstorgcode Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

//func (dbc DbConn) GetAllMststatetype(page *entities.MstorgcodeEntities) ([]entities.MstorgcodeEntities, error) {
//	logger.Log.Println("In side GelAllMststatetype")
//	values := []entities.MstorgcodeEntities{}
//	rows, err := dbc.DB.Query(getMststatetype, page.Clientid, page.mstorgnhirarchyid, page.Offset, page.Limit)
//	defer rows.Close()
//	if err != nil {
//		logger.Log.Println("GetAllMststatetype Get Statement Prepare Error", err)
//		return values, err
//	}
//	for rows.Next() {
//		value := entities.MstorgcodeEntities{}
//		rows.Scan(&value.Id, &value.Clientid, &value.mstorgnhirarchyid, &value.Arcosorgcode, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname)
//		values = append(values, value)
//	}
//	return values, nil
//}

func (dbc DbConn) UpdateMstorgcode(tz *entities.MstorgcodeEntity) error {
	logger.Log.Println("In side UpdateMstorgcode")
	stmt, err := dbc.DB.Prepare(updateMstorgcode)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMstorgcode Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Toolcode, tz.Orgcode, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMstorgcode Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMstorgcode(tz *entities.MstorgcodeEntity) error {
	logger.Log.Println("In side DeleteMstorgcode")
	stmt, err := dbc.DB.Prepare(deleteMstorgcode)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMstorgcode Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMstorgcode Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMstogrcodeCount(tz *entities.MstorgcodeEntity, OrgnTypeID int64) (entities.MstorgcodeEntities, error) {
	logger.Log.Println("In side GetMstogrcodeCount")
	value := entities.MstorgcodeEntities{}
	var getMstorgcodecount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getMstorgcodecount = "SELECT count(a.id) as total FROM maporgcodewithtools a,mstclient b,mstorgnhierarchy c WHERE a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id"
	} else if OrgnTypeID == 2 {
		getMstorgcodecount = "SELECT count(a.id) as total FROM maporgcodewithtools a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id"
		params = append(params, tz.Clientid)
	} else {
		getMstorgcodecount = "SELECT count(a.id) as total FROM maporgcodewithtools a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getMstorgcodecount, params...).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMstogrcodeCount Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) GetAllMstorgcode(tz *entities.MstorgcodeEntity, OrgnType int64) ([]entities.MstorgcodeEntity, error) {
	logger.Log.Println("In side GetAllMstorgcode")
	values := []entities.MstorgcodeEntity{}
	var getMstorgcode string
	var params []interface{}
	if OrgnType == 1 {
		getMstorgcode = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.toolcode as Toolcode, a.orgcode as Orgcode, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname FROM maporgcodewithtools a,mstclient b,mstorgnhierarchy c WHERE a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		getMstorgcode = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.toolcode as Toolcode, a.orgcode as Orgcode, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname FROM maporgcodewithtools a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		getMstorgcode = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.toolcode as Toolcode, a.orgcode as Orgcode, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname FROM maporgcodewithtools a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}
	rows, err := dbc.DB.Query(getMstorgcode, params...)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstorgcode Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstorgcodeEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Toolcode, &value.Orgcode, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) GetAlltoolvalue(page *entities.MstorgcodeEntity) ([]entities.Gettoolscode, error) {
	logger.Log.Println("In side GelAllRecorddifferentiation")
	// logger.Log.Println(getRecorddifferentiation)
	values := []entities.Gettoolscode{}
	var Getalltoolcode = "select distinct toolcode from maporgcodewithtools where clientid=? and  toolcode like ? and activeflg = 1  and deleteflg=0"
	rows, err := dbc.DB.Query(Getalltoolcode, page.Clientid, "%"+page.Toolcode+"%")
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllRecorddifferentiation Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.Gettoolscode{}
		rows.Scan(&value.Toolcode)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) GetAllorgvalue(page *entities.MstorgcodeEntity) ([]entities.Getorgcode, error) {
	logger.Log.Println("In side GelAllRecorddifferentiation")
	// logger.Log.Println(getRecorddifferentiation)
	values := []entities.Getorgcode{}
	var Getalltoolcode = "select distinct orgcode  from maporgcodewithtools where clientid=? and  orgcode like ? and activeflg = 1  and deleteflg=0"
	rows, err := dbc.DB.Query(Getalltoolcode, page.Clientid, "%"+page.Orgcode+"%")
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllRecorddifferentiation Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.Getorgcode{}
		rows.Scan(&value.Orgcode)
		values = append(values, value)
	}
	return values, nil
}
