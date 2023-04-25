package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertRecorddifferentiationtype = "INSERT INTO mstrecorddifferentiationtype (clientid, mstorgnhirarchyid, typename, seqno, parentid) VALUES (?,?,?,?,?)"
var duplicateRecorddifferentiationtype = "SELECT count(id) total FROM  mstrecorddifferentiationtype WHERE clientid = ? AND mstorgnhirarchyid = ? AND typename = ? AND seqno = ? AND deleteflg = 0 and activeflg=1"
var getRecorddifferentiationtype = "SELECT a.id as Id,a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.typename as Typename, a.seqno as Seqno, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname FROM mstrecorddifferentiationtype a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.parentid = ?  AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id ORDER BY a.id DESC LIMIT ?,?"

// var getAllRecorddifferentiationtype = "SELECT a.id as Id,a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid,(SELECT f.typename from mstrecorddifferentiationtype f where f.id=a.parentid and f.activeflg=1 and f.deleteflg=0 ) parentname, a.typename as Typename, a.seqno as Seqno, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname FROM mstrecorddifferentiationtype a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND  a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id ORDER BY a.id DESC LIMIT ?,?"
var getRecorddifferentiationtypecount = "SELECT count(a.id) as total FROM mstrecorddifferentiationtype a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.parentid = ?  AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id "
var getAllRecorddifferentiationtypecount = "SELECT count(a.id) as total FROM mstrecorddifferentiationtype a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id "
var updateRecorddifferentiationtype = "UPDATE mstrecorddifferentiationtype SET mstorgnhirarchyid = ?, typename = ?, seqno = ? WHERE id = ? "
var deleteRecorddifferentiationtype = "UPDATE mstrecorddifferentiationtype SET deleteflg = '1' WHERE id = ? "

func (dbc DbConn) CheckDuplicateRecorddifferentiationtype(tz *entities.RecorddifferentiationtypeEntity) (entities.RecorddifferentiationtypeEntities, error) {
	logger.Log.Println("In side CheckDuplicateRecorddifferentiationtype")
	value := entities.RecorddifferentiationtypeEntities{}
	err := dbc.DB.QueryRow(duplicateRecorddifferentiationtype, tz.Clientid, tz.Mstorgnhirarchyid, tz.Typename, tz.Seqno).Scan(&value.Total)
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

func (dbc DbConn) InsertRecorddifferentiationtype(tz *entities.RecorddifferentiationtypeEntity) (int64, error) {
	logger.Log.Println("In side InsertRecorddifferentiationtype")
	logger.Log.Println("Query -->", insertRecorddifferentiationtype)
	stmt, err := dbc.DB.Prepare(insertRecorddifferentiationtype)
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

func (dbc DbConn) GetAllRecorddifferentiationtype(page *entities.RecorddifferentiationtypeEntity) ([]entities.RecorddifferentiationtypeEntity, error) {
	logger.Log.Println("In side GelAllRecorddifferentiationtype")
	values := []entities.RecorddifferentiationtypeEntity{}
	rows, err := dbc.DB.Query(getRecorddifferentiationtype, page.Clientid, page.Mstorgnhirarchyid, page.Parentid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllRecorddifferentiationtype Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.RecorddifferentiationtypeEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Typename, &value.Seqno, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname)
		values = append(values, value)
	}
	return values, nil
}
func (dbc DbConn) GetRecorddifferentiationtype(tz *entities.RecorddifferentiationtypeEntity, OrgnType int64) ([]entities.RecorddifferentiationtypeEntity, error) {
	logger.Log.Println("In side GelAllRecorddifferentiationtype")
	values := []entities.RecorddifferentiationtypeEntity{}
	var getAllRecorddifferentiationtype string
	var params []interface{}
	if OrgnType == 1 {
		getAllRecorddifferentiationtype = "SELECT a.id as Id,a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid,(SELECT f.typename from mstrecorddifferentiationtype f where f.id=a.parentid and f.activeflg=1 and f.deleteflg=0 ) parentname, a.typename as Typename, a.seqno as Seqno, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname FROM mstrecorddifferentiationtype a,mstclient b,mstorgnhierarchy c WHERE  a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		getAllRecorddifferentiationtype = "SELECT a.id as Id,a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid,(SELECT f.typename from mstrecorddifferentiationtype f where f.id=a.parentid and f.activeflg=1 and f.deleteflg=0 ) parentname, a.typename as Typename, a.seqno as Seqno, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname FROM mstrecorddifferentiationtype a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		getAllRecorddifferentiationtype = "SELECT a.id as Id,a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid,(SELECT f.typename from mstrecorddifferentiationtype f where f.id=a.parentid and f.activeflg=1 and f.deleteflg=0 ) parentname, a.typename as Typename, a.seqno as Seqno, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname FROM mstrecorddifferentiationtype a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND  a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}

	rows, err := dbc.DB.Query(getAllRecorddifferentiationtype, params...)

	// rows, err := dbc.DB.Query(getAllRecorddifferentiationtype, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetRecorddifferentiationtype Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.RecorddifferentiationtypeEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Parentname, &value.Typename, &value.Seqno, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname)
		values = append(values, value)
	}
	return values, nil
}
func (dbc DbConn) UpdateRecorddifferentiationtype(tz *entities.RecorddifferentiationtypeEntity) error {
	logger.Log.Println("In side UpdateRecorddifferentiationtype")
	stmt, err := dbc.DB.Prepare(updateRecorddifferentiationtype)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateRecorddifferentiationtype Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Typename, tz.Seqno, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateRecorddifferentiationtype Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteRecorddifferentiationtype(tz *entities.RecorddifferentiationtypeEntity) error {
	logger.Log.Println("In side DeleteRecorddifferentiationtype")
	stmt, err := dbc.DB.Prepare(deleteRecorddifferentiationtype)
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

func (dbc DbConn) GetRecorddifferentiationtypeCount(tz *entities.RecorddifferentiationtypeEntity) (entities.RecorddifferentiationtypeEntities, error) {
	logger.Log.Println("In side GetRecorddifferentiationtypeCount")
	value := entities.RecorddifferentiationtypeEntities{}
	err := dbc.DB.QueryRow(getRecorddifferentiationtypecount, tz.Clientid, tz.Mstorgnhirarchyid, tz.Parentid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetRecorddifferentiationtypeCount Get Statement Prepare Error", err)
		return value, err
	}
}
func (dbc DbConn) GetAllRecorddifferentiationtypeCount(tz *entities.RecorddifferentiationtypeEntity, OrgnTypeID int64) (entities.RecorddifferentiationtypeEntities, error) {
	logger.Log.Println("In side GetAllRecorddifferentiationtypeCount")
	value := entities.RecorddifferentiationtypeEntities{}
	var getAllRecorddifferentiationtypecount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getAllRecorddifferentiationtypecount = "SELECT count(a.id) as total FROM mstrecorddifferentiationtype a,mstclient b,mstorgnhierarchy c WHERE a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id "
	} else if OrgnTypeID == 2 {
		getAllRecorddifferentiationtypecount = "SELECT count(a.id) as total FROM mstrecorddifferentiationtype a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id "
		params = append(params, tz.Clientid)
	} else {
		getAllRecorddifferentiationtypecount = "SELECT count(a.id) as total FROM mstrecorddifferentiationtype a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id "
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getAllRecorddifferentiationtypecount, params...).Scan(&value.Total)
	// err := dbc.DB.QueryRow(getAllRecorddifferentiationtypecount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetAllRecorddifferentiationtypeCount Get Statement Prepare Error", err)
		return value, err
	}
}
