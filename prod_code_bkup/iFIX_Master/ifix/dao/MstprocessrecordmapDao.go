package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertMstprocessrecordmap = "INSERT INTO mstprocessrecordmap (clientid, mstorgnhirarchyid, recorddifftypeid, recorddiffid, mstprocessid) VALUES (?,?,?,?,?)"
var duplicateMstprocessrecordmap = "SELECT count(id) total FROM  mstprocessrecordmap WHERE clientid = ? AND mstorgnhirarchyid = ? AND recorddifftypeid = ? AND recorddiffid = ?  AND deleteflg = 0 and activeflg=1"
var getMstprocessrecordmap = "SELECT id as Id, clientid as Clientid, mstorgnhirarchyid as Mstorgnhirarchyid, recorddifftypeid as Recorddifftypeid, recorddiffid as Recorddiffid, mstprocessid as Mstprocessid, activeflg as Activeflg FROM mstprocessrecordmap WHERE clientid = ? AND mstorgnhirarchyid = ?  AND deleteflg =0 and activeflg=1 ORDER BY id DESC LIMIT ?,?"
var getMstprocessrecordmapcount = "SELECT count(id) total FROM  mstprocessrecordmap WHERE clientid = ? AND mstorgnhirarchyid = ? AND  deleteflg =0 and activeflg=1"
var updateMstprocessrecordmap = "UPDATE mstprocessrecordmap SET mstorgnhirarchyid = ?, recorddifftypeid = ?, recorddiffid = ? WHERE id = ? "
var deleteMstprocessrecordmap = "UPDATE mstprocessrecordmap SET deleteflg = '1' WHERE id = ? "

func (dbc DbConn) CheckDuplicateMstprocessrecordmap(tz *entities.MstprocessEntity) (entities.MstprocessrecordmapEntities, error) {
	logger.Log.Println("In side CheckDuplicateMstprocessrecordmap")
	value := entities.MstprocessrecordmapEntities{}
	err := dbc.DB.QueryRow(duplicateMstprocessrecordmap, tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMstprocessrecordmap Get Statement Prepare Error", err)
		return value, err
	}
}

func CheckDuplicateMstprocessrecordmapwithtransaction(tx *sql.Tx, tz *entities.MstprocessEntity, processid int64) (entities.MstprocessrecordmapEntities, error) {
	logger.Log.Println("In side CheckDuplicateMstprocessrecordmap")
	value := entities.MstprocessrecordmapEntities{}
	err := tx.QueryRow(duplicateMstprocessrecordmap, tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMstprocessrecordmap Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertMstprocessrecordmap(tz *entities.MstprocessrecordmapEntity) (int64, error) {
	logger.Log.Println("In side InsertMstprocessrecordmap")
	logger.Log.Println("Query -->", insertMstprocessrecordmap)
	stmt, err := dbc.DB.Prepare(insertMstprocessrecordmap)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMstprocessrecordmap Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Mstprocessid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Mstprocessid)
	if err != nil {
		logger.Log.Println("InsertMstprocessrecordmap Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func InsertMstprocessrecordmapwithtransaction(tx *sql.Tx, tz *entities.MstprocessEntity, processid int64) (int64, error) {
	logger.Log.Println("In side InsertMstprocessrecordmap")
	logger.Log.Println("Query -->", insertMstprocessrecordmap)
	stmt, err := tx.Prepare(insertMstprocessrecordmap)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMstprocessrecordmap Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, processid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, processid)
	if err != nil {
		logger.Log.Println("InsertMstprocessrecordmap Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllMstprocessrecordmap(page *entities.MstprocessrecordmapEntity) ([]entities.MstprocessrecordmapEntity, error) {
	logger.Log.Println("In side GelAllMstprocessrecordmap")
	values := []entities.MstprocessrecordmapEntity{}
	rows, err := dbc.DB.Query(getMstprocessrecordmap, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstprocessrecordmap Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstprocessrecordmapEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Recorddifftypeid, &value.Recorddiffid, &value.Mstprocessid, &value.Activeflg)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateMstprocessrecordmap(tz *entities.MstprocessrecordmapEntity) error {
	logger.Log.Println("In side UpdateMstprocessrecordmap")
	stmt, err := dbc.DB.Prepare(updateMstprocessrecordmap)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMstprocessrecordmap Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Mstprocessid, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMstprocessrecordmap Execute Statement  Error", err)
		return err
	}
	return nil
}

func UpdateMstprocessrecordmapwithtransaction(tx *sql.Tx, tz *entities.MstprocessEntity) error {
	logger.Log.Println("In side UpdateMstprocessrecordmap")
	stmt, err := tx.Prepare(updateMstprocessrecordmap)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMstprocessrecordmap Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Mstprocessrecordmapid)
	if err != nil {
		logger.Log.Println("UpdateMstprocessrecordmap Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMstprocessrecordmap(tz *entities.MstprocessrecordmapEntity) error {
	logger.Log.Println("In side DeleteMstprocessrecordmap")
	stmt, err := dbc.DB.Prepare(deleteMstprocessrecordmap)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMstprocessrecordmap Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMstprocessrecordmap Execute Statement  Error", err)
		return err
	}
	return nil
}

func DeleteMstprocessrecordmapwithtransaction(tx *sql.Tx, tz *entities.MstprocessEntity) error {
	logger.Log.Println("In side DeleteMstprocessrecordmap")
	stmt, err := tx.Prepare(deleteMstprocessrecordmap)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMstprocessrecordmap Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstprocessrecordmapid)
	if err != nil {
		logger.Log.Println("DeleteMstprocessrecordmap Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMstprocessrecordmapCount(tz *entities.MstprocessrecordmapEntity) (entities.MstprocessrecordmapEntities, error) {
	logger.Log.Println("In side GetMstprocessrecordmapCount")
	value := entities.MstprocessrecordmapEntities{}
	err := dbc.DB.QueryRow(getMstprocessrecordmapcount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMstprocessrecordmapCount Get Statement Prepare Error", err)
		return value, err
	}
}
