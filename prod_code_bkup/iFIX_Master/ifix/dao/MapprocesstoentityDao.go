package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertMapprocesstoentity = "INSERT INTO mapprocesstoentity (clientid, mstorgnhirarchyid, mstprocessid, mstdatadictionaryfieldid) VALUES (?,?,?,?)"
var duplicateMapprocesstoentity = "SELECT count(id) total FROM  mapprocesstoentity WHERE clientid = ? AND mstorgnhirarchyid = ? AND mstprocessid = ? AND mstdatadictionaryfieldid = ? AND deleteflg = 0 and activeflg=1"
var getMapprocesstoentity = "SELECT id as Id, clientid as Clientid, mstorgnhirarchyid as Mstorgnhirarchyid, mstprocessid as Mstprocessid, mstdatadictionaryfieldid as Mstdatadictionaryfieldid, activeflg as Activeflg FROM mapprocesstoentity WHERE clientid = ? AND mstorgnhirarchyid = ?  AND deleteflg =0 and activeflg=1 ORDER BY id DESC LIMIT ?,?"
var getMapprocesstoentitycount = "SELECT count(id) total FROM  mapprocesstoentity WHERE clientid = ? AND mstorgnhirarchyid = ? AND  deleteflg =0 and activeflg=1"
var updateMapprocesstoentity = "UPDATE mapprocesstoentity SET mstorgnhirarchyid = ?, mstdatadictionaryfieldid = ? WHERE id = ? "
var deleteMapprocesstoentity = "UPDATE mapprocesstoentity SET deleteflg = '1' WHERE id = ? "

func (dbc DbConn) CheckDuplicateMapprocesstoentity(tz *entities.MstprocessEntity,processid int64) (entities.MapprocesstoentityEntities, error) {
	logger.Log.Println("In side CheckDuplicateMapprocesstoentity")
	value := entities.MapprocesstoentityEntities{}
	err := dbc.DB.QueryRow(duplicateMapprocesstoentity, tz.Clientid, tz.Mstorgnhirarchyid, processid, tz.Mstdatadictionaryfieldid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMapprocesstoentity Get Statement Prepare Error", err)
		return value, err
	}
}

func CheckDuplicateMapprocesstoentitywithtransaction(tx *sql.Tx, tz *entities.MstprocessEntity, processid int64) (entities.MapprocesstoentityEntities, error) {
	logger.Log.Println("In side CheckDuplicateMapprocesstoentity")
	value := entities.MapprocesstoentityEntities{}
	err := tx.QueryRow(duplicateMapprocesstoentity, tz.Clientid, tz.Mstorgnhirarchyid, processid, tz.Mstdatadictionaryfieldid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMapprocesstoentity Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertMapprocesstoentity(tz *entities.MapprocesstoentityEntity) (int64, error) {
	logger.Log.Println("In side InsertMapprocesstoentity")
	logger.Log.Println("Query -->", insertMapprocesstoentity)
	stmt, err := dbc.DB.Prepare(insertMapprocesstoentity)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMapprocesstoentity Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Mstprocessid, tz.Mstdatadictionaryfieldid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Mstprocessid, tz.Mstdatadictionaryfieldid)
	if err != nil {
		logger.Log.Println("InsertMapprocesstoentity Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func InsertMapprocesstoentitywithtransaction(tx *sql.Tx, tz *entities.MstprocessEntity, processid int64) (int64, error) {
	logger.Log.Println("In side InsertMapprocesstoentity")
	logger.Log.Println("Query -->", insertMapprocesstoentity)
	stmt, err := tx.Prepare(insertMapprocesstoentity)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMapprocesstoentity Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, processid, tz.Mstdatadictionaryfieldid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, processid, tz.Mstdatadictionaryfieldid)
	if err != nil {
		logger.Log.Println("InsertMapprocesstoentity Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllMapprocesstoentity(page *entities.MapprocesstoentityEntity) ([]entities.MapprocesstoentityEntity, error) {
	logger.Log.Println("In side GelAllMapprocesstoentity")
	values := []entities.MapprocesstoentityEntity{}
	rows, err := dbc.DB.Query(getMapprocesstoentity, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMapprocesstoentity Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MapprocesstoentityEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Mstprocessid, &value.Mstdatadictionaryfieldid, &value.Activeflg)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateMapprocesstoentity(tz *entities.MapprocesstoentityEntity) error {
	logger.Log.Println("In side UpdateMapprocesstoentity")
	stmt, err := dbc.DB.Prepare(updateMapprocesstoentity)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMapprocesstoentity Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Mstprocessid, tz.Mstdatadictionaryfieldid, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMapprocesstoentity Execute Statement  Error", err)
		return err
	}
	return nil
}

func UpdateMapprocesstoentitywithtransaction(tx *sql.Tx, tz *entities.MstprocessEntity) error {
	logger.Log.Println("In side UpdateMapprocesstoentity")
	stmt, err := tx.Prepare(updateMapprocesstoentity)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMapprocesstoentity Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Mstdatadictionaryfieldid, tz.Mstprocesstoentityid)
	if err != nil {
		logger.Log.Println("UpdateMapprocesstoentity Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMapprocesstoentity(tz *entities.MapprocesstoentityEntity) error {
	logger.Log.Println("In side DeleteMapprocesstoentity")
	stmt, err := dbc.DB.Prepare(deleteMapprocesstoentity)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMapprocesstoentity Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMapprocesstoentity Execute Statement  Error", err)
		return err
	}
	return nil
}

func DeleteMapprocesstoentitywithtransaction(tx *sql.Tx, tz *entities.MstprocessEntity) error {
	logger.Log.Println("In side DeleteMapprocesstoentity")
	stmt, err := tx.Prepare(deleteMapprocesstoentity)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMapprocesstoentity Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstprocesstoentityid)
	if err != nil {
		logger.Log.Println("DeleteMapprocesstoentity Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMapprocesstoentityCount(tz *entities.MapprocesstoentityEntity) (entities.MapprocesstoentityEntities, error) {
	logger.Log.Println("In side GetMapprocesstoentityCount")
	value := entities.MapprocesstoentityEntities{}
	err := dbc.DB.QueryRow(getMapprocesstoentitycount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMapprocesstoentityCount Get Statement Prepare Error", err)
		return value, err
	}
}
