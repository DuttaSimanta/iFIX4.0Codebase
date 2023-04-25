package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertMststateterm = "INSERT INTO mststateterm (clientid, mstorgnhirarchyid, recorddifftypeid, recorddiffid, recordtermid, recordtermvalue, iscompulsory) VALUES (?,?,?,?,?,?,?)"
var duplicateMststateterm = "SELECT count(id) total FROM  mststateterm WHERE clientid = ? AND mstorgnhirarchyid = ? AND recorddifftypeid = ? AND recorddiffid = ? AND recordtermid = ? AND deleteflg = 0"

// var getMststateterm = "SELECT mststateterm.id as Id, mststateterm.clientid as Clientid, mststateterm.mstorgnhirarchyid as Mstorgnhirarchyid, mststateterm.recorddifftypeid as Recorddifftypeid, mststateterm.recorddiffid as Recorddiffid, mststateterm.recordtermid as Recordtermid, mststateterm.recordtermvalue as Recordtermvalue, mststateterm.iscompulsory as Iscompulsory, mststateterm.activeflg as Activeflg, mstclient.name as Clientname,mstorgnhierarchy.name as Mstorgnhirarchyname,mstrecordterms.termname as Termname,mstrecorddifferentiationtype.typename AS mstrecorddifferentiationtypename,mstrecorddifferentiation.name mstrecorddifferentiationname FROM mststateterm JOIN mstclient ON  mststateterm.clientid=mstclient.id  JOIN mstorgnhierarchy ON  mststateterm.mstorgnhirarchyid=mstorgnhierarchy.id AND mstclient.id=mstorgnhierarchy.clientid JOIN mstrecordterms ON mststateterm.recordtermid=mstrecordterms.id AND mstrecordterms.deleteflg =0 AND mstrecordterms.activeflg=1 JOIN mstrecorddifferentiationtype ON mststateterm.recorddifftypeid=mstrecorddifferentiationtype.id JOIN mstrecorddifferentiation ON mststateterm.recorddiffid=mstrecorddifferentiation.id WHERE mststateterm.clientid = ? AND mststateterm.mstorgnhirarchyid = ?  AND mststateterm.deleteflg =0 and mststateterm.activeflg=1 ORDER BY mststateterm.id DESC LIMIT ?,?"
var getMststatetermcount = "SELECT count(id) total FROM  mststateterm WHERE clientid = ? AND mstorgnhirarchyid = ? AND  deleteflg =0 and activeflg=1"
var updateMststateterm = "UPDATE mststateterm SET recorddifftypeid = ?, recorddiffid = ?, recordtermid = ?, recordtermvalue = ?, iscompulsory = ? WHERE id = ? "
var deleteMststateterm = "UPDATE mststateterm SET deleteflg = '1' WHERE id = ? "

func (dbc DbConn) CheckDuplicateMststateterm(tz *entities.MststatetermEntity) (entities.MststatetermEntities, error) {
	logger.Log.Println("In side CheckDuplicateMststateterm")
	value := entities.MststatetermEntities{}
	err := dbc.DB.QueryRow(duplicateMststateterm, tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Recordtermid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMststateterm Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertMststateterm(tz *entities.MststatetermEntity) (int64, error) {
	logger.Log.Println("In side InsertMststateterm")
	logger.Log.Println("Query -->", insertMststateterm)
	stmt, err := dbc.DB.Prepare(insertMststateterm)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMststateterm Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Recordtermid, tz.Recordtermvalue, tz.Iscompulsory)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Recordtermid, tz.Recordtermvalue, tz.Iscompulsory)
	if err != nil {
		logger.Log.Println("InsertMststateterm Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllMststateterm(tz *entities.MststatetermEntity, OrgnType int64) ([]entities.MststatetermEntity, error) {
	logger.Log.Println("In side GelAllMststateterm")
	values := []entities.MststatetermEntity{}
	var getMststateterm string
	var params []interface{}
	if OrgnType == 1 {
		getMststateterm = "SELECT mststateterm.id as Id, mststateterm.clientid as Clientid, mststateterm.mstorgnhirarchyid as Mstorgnhirarchyid, mststateterm.recorddifftypeid as Recorddifftypeid, mststateterm.recorddiffid as Recorddiffid, mststateterm.recordtermid as Recordtermid, mststateterm.recordtermvalue as Recordtermvalue, mststateterm.iscompulsory as Iscompulsory, mststateterm.activeflg as Activeflg, mstclient.name as Clientname,mstorgnhierarchy.name as Mstorgnhirarchyname,mstrecordterms.termname as Termname,mstrecorddifferentiationtype.typename AS mstrecorddifferentiationtypename,mstrecorddifferentiation.name mstrecorddifferentiationname FROM mststateterm JOIN mstclient ON  mststateterm.clientid=mstclient.id  JOIN mstorgnhierarchy ON  mststateterm.mstorgnhirarchyid=mstorgnhierarchy.id AND mstclient.id=mstorgnhierarchy.clientid JOIN mstrecordterms ON mststateterm.recordtermid=mstrecordterms.id AND mstrecordterms.deleteflg =0 AND mstrecordterms.activeflg=1 JOIN mstrecorddifferentiationtype ON mststateterm.recorddifftypeid=mstrecorddifferentiationtype.id JOIN mstrecorddifferentiation ON mststateterm.recorddiffid=mstrecorddifferentiation.id WHERE mststateterm.deleteflg =0 and mststateterm.activeflg=1 ORDER BY mststateterm.id DESC LIMIT ?,?"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		getMststateterm = "SELECT mststateterm.id as Id, mststateterm.clientid as Clientid, mststateterm.mstorgnhirarchyid as Mstorgnhirarchyid, mststateterm.recorddifftypeid as Recorddifftypeid, mststateterm.recorddiffid as Recorddiffid, mststateterm.recordtermid as Recordtermid, mststateterm.recordtermvalue as Recordtermvalue, mststateterm.iscompulsory as Iscompulsory, mststateterm.activeflg as Activeflg, mstclient.name as Clientname,mstorgnhierarchy.name as Mstorgnhirarchyname,mstrecordterms.termname as Termname,mstrecorddifferentiationtype.typename AS mstrecorddifferentiationtypename,mstrecorddifferentiation.name mstrecorddifferentiationname FROM mststateterm JOIN mstclient ON  mststateterm.clientid=mstclient.id  JOIN mstorgnhierarchy ON  mststateterm.mstorgnhirarchyid=mstorgnhierarchy.id AND mstclient.id=mstorgnhierarchy.clientid JOIN mstrecordterms ON mststateterm.recordtermid=mstrecordterms.id AND mstrecordterms.deleteflg =0 AND mstrecordterms.activeflg=1 JOIN mstrecorddifferentiationtype ON mststateterm.recorddifftypeid=mstrecorddifferentiationtype.id JOIN mstrecorddifferentiation ON mststateterm.recorddiffid=mstrecorddifferentiation.id WHERE mststateterm.clientid = ? AND mststateterm.deleteflg =0 and mststateterm.activeflg=1 ORDER BY mststateterm.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		getMststateterm = "SELECT mststateterm.id as Id, mststateterm.clientid as Clientid, mststateterm.mstorgnhirarchyid as Mstorgnhirarchyid, mststateterm.recorddifftypeid as Recorddifftypeid, mststateterm.recorddiffid as Recorddiffid, mststateterm.recordtermid as Recordtermid, mststateterm.recordtermvalue as Recordtermvalue, mststateterm.iscompulsory as Iscompulsory, mststateterm.activeflg as Activeflg, mstclient.name as Clientname,mstorgnhierarchy.name as Mstorgnhirarchyname,mstrecordterms.termname as Termname,mstrecorddifferentiationtype.typename AS mstrecorddifferentiationtypename,mstrecorddifferentiation.name mstrecorddifferentiationname FROM mststateterm JOIN mstclient ON  mststateterm.clientid=mstclient.id  JOIN mstorgnhierarchy ON  mststateterm.mstorgnhirarchyid=mstorgnhierarchy.id AND mstclient.id=mstorgnhierarchy.clientid JOIN mstrecordterms ON mststateterm.recordtermid=mstrecordterms.id AND mstrecordterms.deleteflg =0 AND mstrecordterms.activeflg=1 JOIN mstrecorddifferentiationtype ON mststateterm.recorddifftypeid=mstrecorddifferentiationtype.id JOIN mstrecorddifferentiation ON mststateterm.recorddiffid=mstrecorddifferentiation.id WHERE mststateterm.clientid = ? AND mststateterm.mstorgnhirarchyid = ?  AND mststateterm.deleteflg =0 and mststateterm.activeflg=1 ORDER BY mststateterm.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}

	rows, err := dbc.DB.Query(getMststateterm, params...)
	// rows, err := dbc.DB.Query(getMststateterm, page.Clientid,page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMststateterm Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MststatetermEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Recorddifftypeid, &value.Recorddiffid, &value.Recordtermid, &value.Recordtermvalue, &value.Iscompulsory, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Termname, &value.Mstrecorddifferentiationtypename, &value.Mstrecorddifferentiationname)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateMststateterm(tz *entities.MststatetermEntity) error {
	logger.Log.Println("In side UpdateMststateterm")
	stmt, err := dbc.DB.Prepare(updateMststateterm)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMststateterm Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Recorddifftypeid, tz.Recorddiffid, tz.Recordtermid, tz.Recordtermvalue, tz.Iscompulsory, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMststateterm Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMststateterm(tz *entities.MststatetermEntity) error {
	logger.Log.Println("In side DeleteMststateterm")
	stmt, err := dbc.DB.Prepare(deleteMststateterm)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMststateterm Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMststateterm Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMststatetermCount(tz *entities.MststatetermEntity, OrgnTypeID int64) (entities.MststatetermEntities, error) {
	logger.Log.Println("In side GetMststatetermCount")
	value := entities.MststatetermEntities{}
	var getMststatetermcount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getMststatetermcount = "SELECT count(id) total FROM  mststateterm WHERE  deleteflg =0 and activeflg=1"
	} else if OrgnTypeID == 2 {
		getMststatetermcount = "SELECT count(id) total FROM  mststateterm WHERE clientid = ? AND  deleteflg =0 and activeflg=1"
		params = append(params, tz.Clientid)
	} else {
		getMststatetermcount = "SELECT count(id) total FROM  mststateterm WHERE clientid = ? AND mstorgnhirarchyid = ? AND  deleteflg =0 and activeflg=1"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getMststatetermcount, params...).Scan(&value.Total)
	// err := dbc.DB.QueryRow(getMststatetermcount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMststatetermCount Get Statement Prepare Error", err)
		return value, err
	}
}
