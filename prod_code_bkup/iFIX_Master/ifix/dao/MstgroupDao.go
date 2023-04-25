package dao

import (
	"database/sql"
	"errors"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertMstgroup = "INSERT INTO mstgroup (clientid, mstorgnhirarchyid, groupname,externalgrpid) VALUES (?,?,?,?)"
var duplicateMstgroup = "SELECT count(id) total FROM  mstgroup WHERE clientid = ? AND mstorgnhirarchyid = ? AND groupname = ? AND deleteflg = 0"
var getMstgroup = "SELECT id as Id, clientid as Clientid, mstorgnhirarchyid as Mstorgnhirarchyid, groupname as Groupname, activeflg as Activeflg FROM mstgroup WHERE clientid = ? AND mstorgnhirarchyid = ?  AND deleteflg =0 and activeflg=1 ORDER BY id DESC LIMIT ?,?"
var getMstgroupcount = "SELECT count(id) total FROM  mstgroup WHERE clientid = ? AND mstorgnhirarchyid = ? AND  deleteflg =0 and activeflg=1"
var updateMstgroup = "UPDATE mstgroup SET mstorgnhirarchyid = ?, groupname = ? WHERE externalgrpid = ? "
var deleteMstgroup = "UPDATE mstgroup SET deleteflg = '1' WHERE externalgrpid = ? "

func (dbc DbConn) CheckDuplicateMstgroup(tz *entities.MstgroupEntity) (entities.MstgroupEntities, error) {
	logger.Log.Println("In side CheckDuplicateMstgroup")
	value := entities.MstgroupEntities{}
	err := dbc.DB.QueryRow(duplicateMstgroup, tz.Clientid, tz.Mstorgnhirarchyid, tz.Groupname).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMstgroup Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertMstgroup(tz *entities.MstgroupEntity) (int64, error) {
	logger.Log.Println("In side InsertMstgroup")
	logger.Log.Println("Query -->", insertMstgroup)
	stmt, err := dbc.DB.Prepare(insertMstgroup)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMstgroup Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Groupname)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Groupname)
	if err != nil {
		logger.Log.Println("InsertMstgroup Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllMstgroup(page *entities.MstgroupEntity) ([]entities.MstgroupEntity, error) {
	logger.Log.Println("In side GelAllMstgroup")
	values := []entities.MstgroupEntity{}
	rows, err := dbc.DB.Query(getMstgroup, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstgroup Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstgroupEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Groupname, &value.Activeflg)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateMstgroup(tz *entities.MstgroupEntity) error {
	logger.Log.Println("In side UpdateMstgroup")
	stmt, err := dbc.DB.Prepare(updateMstgroup)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMstgroup Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Groupname, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMstgroup Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMstgroup(tz *entities.MstgroupEntity) error {
	logger.Log.Println("In side DeleteMstgroup")
	stmt, err := dbc.DB.Prepare(deleteMstgroup)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMstgroup Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMstgroup Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMstgroupCount(tz *entities.MstgroupEntity) (entities.MstgroupEntities, error) {
	logger.Log.Println("In side GetMstgroupCount")
	value := entities.MstgroupEntities{}
	err := dbc.DB.QueryRow(getMstgroupcount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMstgroupCount Get Statement Prepare Error", err)
		return value, err
	}
}

// All methods defination with transaction

func CheckDuplicateMstgroupwithtransaction(tx *sql.Tx, tz *entities.ClientsupportgroupEntity) (entities.MstgroupEntities, error) {
	logger.Log.Println("In side CheckDuplicateMstgroupwithtransaction")
	value := entities.MstgroupEntities{}
	err := tx.QueryRow(duplicateMstgroup, tz.Clientid, tz.Mstorgnhirarchyid, tz.Supportgroupname).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMstgroupwithtransaction Get Statement Prepare Error", err)
		return value, err
	}
}

func InsertMstgroupwithtransaction(tx *sql.Tx, tz *entities.ClientsupportgroupEntity, lastinsertedID int64) (int64, error) {
	stmt, err := tx.Prepare(insertMstgroup)
	logger.Log.Println("insertMstgroup is ---->", insertMstgroup)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println(err)
		return 0, errors.New("SQL Prepare Error")
	}

	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Supportgroupname, lastinsertedID)
	if err != nil {
		logger.Log.Println(err)
		return 0, errors.New("SQL Execution Error")
	}
	lastInsertedID, err := res.LastInsertId()
	if err != nil {
		logger.Log.Println(err)
		return 0, errors.New("SQL Last ID fetch Error")
	}

	return lastInsertedID, nil
}

func GetAllMstgroupwithtransaction(tx *sql.Tx, page *entities.MstgroupEntity) ([]entities.MstgroupEntity, error) {
	logger.Log.Println("In side GelAllMstgroup")
	values := []entities.MstgroupEntity{}
	rows, err := tx.Query(getMstgroup, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstgroup Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstgroupEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Groupname, &value.Activeflg)
		values = append(values, value)
	}
	return values, nil
}

func DeleteMstgroupwithtransaction(tx *sql.Tx, tz *entities.ClientsupportgroupEntity) error {
	logger.Log.Println("In side DeleteMstgroup")
	stmt, err := tx.Prepare(deleteMstgroup)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMstgroup Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMstgroup Execute Statement  Error", err)
		return err
	}
	return nil
}

func UpdateMstgroupwithtransaction(tx *sql.Tx, tz *entities.ClientsupportgroupEntity) error {
	logger.Log.Println("In side UpdateMstgroupwithtransaction")
	stmt, err := tx.Prepare(updateMstgroup)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println(err)
		return errors.New("SQL Prepare Error")
	}

	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Supportgroupname, tz.Id)
	if err != nil {
		logger.Log.Println(err)
		return errors.New("SQL Execution Error")
	}

	return nil
}
