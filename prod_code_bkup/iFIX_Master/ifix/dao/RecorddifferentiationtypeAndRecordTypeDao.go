package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertRecorddifftype = "INSERT INTO mstrecorddifferentiationtype (clientid, mstorgnhirarchyid, typename, seqno, parentid) VALUES (?,?,?,?,?)"
var duplicateRecorddifftype = "SELECT count(id) total FROM  mstrecorddifferentiationtype WHERE clientid = ? AND mstorgnhirarchyid = ? AND typename = ? AND seqno = ? AND deleteflg = 0 and activeflg=1"
var updateRecorddifftype = "UPDATE mstrecorddifferentiationtype SET clientid=?, mstorgnhirarchyid = ?, typename = ?, seqno = ?,parentid=? WHERE id = ? "
var deleteRecorddifftype = "UPDATE mstrecorddifferentiationtype SET deleteflg = '1' WHERE id = ? "

var insertRecrdType = "INSERT INTO mstrecordtype (clientid, mstorgnhirarchyid, fromrecorddifftypeid, fromrecorddiffid, torecorddifftypeid, torecorddiffid) VALUES (?,?,?,?,?,?)"
var duplicateRecrdType = "SELECT count(id) total FROM  mstrecordtype WHERE clientid = ? AND mstorgnhirarchyid = ? AND fromrecorddifftypeid = ? AND fromrecorddiffid = ? AND torecorddifftypeid = ? AND torecorddiffid=? AND deleteflg = 0"
var updateRecrdType = "UPDATE mstrecordtype SET clientid=?,mstorgnhirarchyid = ?, fromrecorddifftypeid = ?, fromrecorddiffid = ?, torecorddifftypeid = ?, torecorddiffid = ? WHERE torecorddifftypeid = ? "

// var deleteRecrdType = "UPDATE mstrecordtype SET deleteflg = '1' WHERE torecorddiffid=? and torecorddifftypeid = ? "
var deleteRecrdType = "UPDATE mstrecordtype SET deleteflg = '1' WHERE id= ? "

func (dbc TxConn) CheckDuplicateRecorddifftype(tz *entities.RecorddifftypeAndRecordTypeEntity) (entities.RecorddifftypeAndRecordTypeEntities, error) {
	logger.Log.Println("In side CheckDuplicateRecorddifferentiationtype")
	value := entities.RecorddifftypeAndRecordTypeEntities{}
	err := dbc.TX.QueryRow(duplicateRecorddifftype, tz.Clientid, tz.Mstorgnhirarchyid, tz.Typename, tz.Seqno).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateRecorddifferentiationtype Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc TxConn) InsertRecorddifftype(tz *entities.RecorddifftypeAndRecordTypeEntity) (int64, error) {
	logger.Log.Println("In side InsertRecorddifferentiationtype")
	logger.Log.Println("Query -->", insertRecorddifferentiationtype)
	stmt, err := dbc.TX.Prepare(insertRecorddifftype)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertRecorddifferentiationtype Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Typename, tz.Seqno, tz.Parentid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Typename, tz.Seqno, tz.Parentid)
	if err != nil {
		logger.Log.Println("InsertRecorddifferentiationtype Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}
func (dbc TxConn) UpdateRecorddifftype(tz *entities.RecorddifftypeAndRecordTypeEntity) error {
	logger.Log.Println("In side UpdateRecorddifferentiationtype")
	stmt, err := dbc.TX.Prepare(updateRecorddifftype)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateRecorddifferentiationtype Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Typename, tz.Seqno, tz.Parentid, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateRecorddifferentiationtype Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc TxConn) DeleteRecorddifftype(tz *entities.RecorddifftypeAndRecordTypeEntity) error {
	logger.Log.Println("In side DeleteRecorddifferentiationtype")
	stmt, err := dbc.TX.Prepare(deleteRecorddifftype)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteRecorddifferentiationtype Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteRecorddifferentiationtype Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc TxConn) CheckDuplicateRecrdType(tz *entities.RecorddifftypeAndRecordTypeEntity) (entities.RecorddifftypeAndRecordTypeEntities, error) {
	logger.Log.Println("In side CheckDuplicateRecordtype")
	value := entities.RecorddifftypeAndRecordTypeEntities{}
	err := dbc.TX.QueryRow(duplicateRecrdType, tz.Clientid, tz.Mstorgnhirarchyid, tz.Fromrecorddifftypeid, tz.Fromrecorddiffid, tz.Torecorddifftypeid, tz.Torecorddiffid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateRecordtype Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc TxConn) InsertRecrdType(tz *entities.RecorddifftypeAndRecordTypeEntity) (int64, error) {
	logger.Log.Println("In side InsertRecordtype")
	logger.Log.Println("Query -->", insertRecrdType)
	stmt, err := dbc.TX.Prepare(insertRecrdType)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertRecordtype Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Fromrecorddifftypeid, tz.Fromrecorddiffid, tz.Torecorddifftypeid, tz.Torecorddiffid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Fromrecorddifftypeid, tz.Fromrecorddiffid, tz.Torecorddifftypeid, tz.Torecorddiffid)
	if err != nil {
		logger.Log.Println("InsertRecordtype Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}
func (dbc TxConn) UpdateRecrdType(tz *entities.RecorddifftypeAndRecordTypeEntity) error {
	logger.Log.Println("In side UpdateRecordtype")
	stmt, err := dbc.TX.Prepare(updateRecrdType)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateRecordtype Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Fromrecorddifftypeid, tz.Fromrecorddiffid, tz.Torecorddifftypeid, tz.Torecorddiffid, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateRecordtype Execute Statement  Error", err)
		return err
	}
	return nil
}

// func (dbc DbConn) DeleteRecrdType(tz *entities.RecorddifftypeAndRecordTypeEntity) error {
// 	logger.Log.Println("In side DeleteRecordtype")
// 	stmt, err := dbc.DB.Prepare(deleteRecrdType)
// 	defer stmt.Close()
// 	if err != nil {
// 		logger.Log.Println("DeleteRecordtype Prepare Statement  Error", err)
// 		return err
// 	}
// 	_, err = stmt.Exec(tz.Torecorddiffid, tz.Id)
// 	if err != nil {
// 		logger.Log.Println("DeleteRecordtype Execute Statement  Error", err)
// 		return err
// 	}
// 	return nil
// }
func (dbc DbConn) DeleteRecrdType(tz *entities.RecorddifftypeAndRecordTypeEntity) error {
	logger.Log.Println("In side DeleteRecordtype")
	stmt, err := dbc.DB.Prepare(deleteRecrdType)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteRecordtype Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteRecordtype Execute Statement  Error", err)
		return err
	}
	return nil
}
