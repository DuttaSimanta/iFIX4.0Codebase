package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
)

var insertErrormesg = "INSERT INTO msterrormesg (errorcode, errormsg) VALUES (?,?)"
var duplicateErrormesg = "SELECT count(id) total FROM  msterrormesg WHERE deleteflg = 0"
var getErrormesg = "SELECT errorcode as Errorcode, errormsg as Errormsg FROM msterrormesg WHERE deleteflg =0 ORDER BY id DESC LIMIT ?,?"
var getErrormesgcount = "SELECT count(id) total FROM  msterrormesg WHERE deleteflg =0 "
var updateErrormesg = "UPDATE msterrormesg SET errormsg = ? WHERE id = ? "
var deleteErrormesg = "UPDATE msterrormesg SET deleteflg = '1' WHERE id = ?"

func (dbc DbConn) CheckDuplicateErrormesg(tz *entities.ErrormesgEntity) (entities.ErrormesgEntities, error) {
	logger.Log.Println("In side CheckDuplicateErrormesg")
	value := entities.ErrormesgEntities{}
	err := dbc.DB.QueryRow(duplicateErrormesg).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateErrormesg Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertErrormesg(tz *entities.ErrormesgEntity) (int64, error) {
	logger.Log.Println("In side InsertErrormesg")
	logger.Log.Println("Query -->", insertErrormesg)
	stmt, err := dbc.DB.Prepare(insertErrormesg)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertErrormesg Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Errorcode, tz.Errormsg)
	res, err := stmt.Exec(tz.Errorcode, tz.Errormsg)
	if err != nil {
		logger.Log.Println("InsertErrormesg Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllErrormesg(page *entities.ErrormesgEntity) ([]entities.ErrormesgEntity, error) {
	logger.Log.Println("In side GelAllErrormesg")
	values := []entities.ErrormesgEntity{}
	rows, err := dbc.DB.Query(getErrormesg, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllErrormesg Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.ErrormesgEntity{}
		rows.Scan(&value.Errorcode, &value.Errormsg)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateErrormesg(tz *entities.ErrormesgEntity) error {
	logger.Log.Println("In side UpdateErrormesg")
	stmt, err := dbc.DB.Prepare(updateErrormesg)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateErrormesg Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Errormsg, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateErrormesg Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteErrormesg(tz *entities.ErrormesgEntity) error {
	logger.Log.Println("In side DeleteErrormesg")
	stmt, err := dbc.DB.Prepare(deleteErrormesg)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteErrormesg Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteErrormesg Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetErrormesgCount(tz *entities.ErrormesgEntity) (entities.ErrormesgEntities, error) {
	logger.Log.Println("In side GetErrormesgCount")
	value := entities.ErrormesgEntities{}
	err := dbc.DB.QueryRow(getErrormesgcount).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		log.Print("GetErrormesgCount Get Statement Prepare Error", err)
		return value, err
	}
}
