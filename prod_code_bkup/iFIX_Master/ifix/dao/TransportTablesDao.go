package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
)

var insertTransporttable = "INSERT INTO msttransporttable (msttablename, tabletype, typedescription) VALUES (?,?,?)"
var duplicateTransporttable = "SELECT count(id) total FROM  msttransporttable WHERE msttablename=? AND deleteflg = 0"
var duplicateTransporttableupdate = "SELECT count(id) total FROM  msttransporttable WHERE msttablename=? AND deleteflg = 0 and id<>?"

var getTransporttable = "SELECT a.id as Id, a.msttablename as msttablename, a.tabletype as tabletype, a.typedescription as typedescription FROM msttransporttable a WHERE  a.deleteflg =0 and a.activeflg=1 ORDER BY a.id DESC LIMIT ?,?"
var getTransporttabletype = "SELECT a.tabletype as tabletype, a.typedescription as typedescription FROM  msttransporttable a WHERE typedescription=? AND deleteflg = 0"
var getTransporttablecount = "SELECT count(a.id) as total FROM msttransporttable a WHERE  a.deleteflg =0 and a.activeflg=1 "
var updateTransporttable = "UPDATE msttransporttable SET msttablename = ?, tabletype = ?, typedescription = ? WHERE id = ? "
var deleteTransporttable = "UPDATE msttransporttable SET deleteflg = '1' WHERE id = ? "
var getTransporttablemaxtype = "select coalesce(max(tabletype),-1) as tabletype from msttransporttable where deleteflg=0"

func (dbc DbConn) Getmaxtype() (int64, error) {
	logger.Log.Println("In side GelAllTransporttable")
	var value int64
	// var values interface{}
	err := dbc.DB.QueryRow(getTransporttablemaxtype).Scan(&value)
	switch err {
	case sql.ErrNoRows:
		value = -1
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetTransporttableCount Get Statement Prepare Error", err)
		return value, err
	}
}
func (dbc DbConn) Gettype(tz *entities.TransporttableEntity) ([]entities.TransporttableEntity, error) {
	logger.Log.Println("In side Gettype")
	values := []entities.TransporttableEntity{}
	logger.Log.Println(tz.Limit, tz.Offset)
	rows, err := dbc.DB.Query(getTransporttabletype, tz.Tabletypedescription)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("Gettype Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.TransporttableEntity{}
		rows.Scan(&value.Tabletype, &value.Tabletypedescription)
		values = append(values, value)
	}
	return values, nil
}
func (dbc DbConn) CheckDuplicateTransporttable(tz *entities.TransporttableEntity) (entities.TransporttableEntities, error) {
	logger.Log.Println("In side CheckDuplicateTransporttable")
	value := entities.TransporttableEntities{}
	err := dbc.DB.QueryRow(duplicateTransporttable, tz.Msttablename).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateTransporttable Get Statement Prepare Error", err)
		return value, err
	}
}
func (dbc DbConn) CheckDuplicateTransporttableupdate(tz *entities.TransporttableEntity) (entities.TransporttableEntities, error) {
	logger.Log.Println("In side CheckDuplicateTransporttable")
	value := entities.TransporttableEntities{}
	err := dbc.DB.QueryRow(duplicateTransporttableupdate, tz.Msttablename, tz.Id).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateTransporttable Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertTransporttable(tz *entities.TransporttableEntity) (int64, error) {
	logger.Log.Println("In side InsertTransporttable")
	logger.Log.Println("Query -->", insertTransporttable)
	stmt, err := dbc.DB.Prepare(insertTransporttable)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertTransporttable Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Msttablename, tz.Tabletype, tz.Tabletypedescription)
	res, err := stmt.Exec(tz.Msttablename, tz.Tabletype, tz.Tabletypedescription)
	if err != nil {
		logger.Log.Println("InsertTransporttable Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

//func (dbc DbConn) GetAllTransporttable(page *entities.TransporttableEntity) ([]entities.TransporttableEntity, error) {
//	logger.Log.Println("In side GelAllTransporttable")
//	values := []entities.TransporttableEntity{}
//	rows, err := dbc.DB.Query(getTransporttable, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
//	defer rows.Close()
//	if err != nil {
//		logger.Log.Println("GetAllTransporttable Get Statement Prepare Error", err)
//		return values, err
//	}
//	for rows.Next() {
//		value := entities.TransporttableEntity{}
//		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Statetypename, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname)
//		values = append(values, value)
//	}
//	return values, nil
//}

func (dbc DbConn) UpdateTransporttable(tz *entities.TransporttableEntity) error {
	logger.Log.Println("In side UpdateTransporttable")
	stmt, err := dbc.DB.Prepare(updateTransporttable)
	logger.Log.Println(tz)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateTransporttable Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Msttablename, tz.Tabletype, tz.Tabletypedescription, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateTransporttable Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteTransporttable(tz *entities.TransporttableEntity) error {
	logger.Log.Println("In side DeleteTransporttable")
	stmt, err := dbc.DB.Prepare(deleteTransporttable)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteTransporttable Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteTransporttable Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetTransporttableCount(tz *entities.TransporttableEntity) (entities.TransporttableEntities, error) {
	logger.Log.Println("In side GetTransporttableCount")
	value := entities.TransporttableEntities{}

	err := dbc.DB.QueryRow(getTransporttablecount).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetTransporttableCount Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) GetAllTransporttable(tz *entities.TransporttableEntity) ([]entities.TransporttableEntity, error) {
	logger.Log.Println("In side GelAllTransporttable")
	values := []entities.TransporttableEntity{}
	logger.Log.Println(tz.Limit, tz.Offset)
	rows, err := dbc.DB.Query(getTransporttable, tz.Offset, tz.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllTransporttable Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.TransporttableEntity{}
		rows.Scan(&value.Id, &value.Msttablename, &value.Tabletype, &value.Tabletypedescription)
		values = append(values, value)
	}
	return values, nil
}
func (mdao DbConn) Gettypedescription(tz *entities.TransporttableEntity) ([]entities.GettableEntity, error) {
	logger.Log.Println("In side Gettypedescriptiondao")
	values := []entities.GettableEntity{}
	var gettable = "SELECT distinct a.tabletype as tabletype, a.typedescription as typedescription FROM  msttransporttable a WHERE typedescription like ? AND deleteflg = 0"

	rows, err := mdao.DB.Query(gettable, "%"+tz.Tabletypedescription+"%")
	defer rows.Close()
	if err != nil {
		log.Print("Gettypedescriptiondao Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.GettableEntity{}
		rows.Scan(&value.Tabletype, &value.Tabletypedescription)
		values = append(values, value)
	}
	return values, nil
}
func (mdao DbConn) Gettable(tz *entities.TransporttableEntity) ([]entities.TableEntity, error) {
	logger.Log.Println("In side Gettabledao")
	values := []entities.TableEntity{}
	var gettable = "SELECT a.msttablename,a.tabletype FROM  msttransporttable a WHERE deleteflg = 0 AND tabletype=?"

	rows, err := mdao.DB.Query(gettable, tz.Tabletype)
	defer rows.Close()
	if err != nil {
		log.Print("Gettabledao Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.TableEntity{}
		rows.Scan(&value.Tablename, &value.Tabletype)
		values = append(values, value)
	}
	return values, nil
}
func (mdao DbConn) Gettypefortransport(tz *entities.TransporttableEntity) ([]entities.GettableEntity, error) {
	logger.Log.Println("In side Gettypefortransportdao")
	values := []entities.GettableEntity{}
	var gettable = "SELECT distinct a.tabletype as tabletype, a.typedescription as typedescription FROM  msttransporttable a WHERE deleteflg = 0"

	rows, err := mdao.DB.Query(gettable)
	defer rows.Close()
	if err != nil {
		log.Print("Gettypefortransportdao Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.GettableEntity{}
		rows.Scan(&value.Tabletype, &value.Tabletypedescription)
		values = append(values, value)
	}
	return values, nil
}
