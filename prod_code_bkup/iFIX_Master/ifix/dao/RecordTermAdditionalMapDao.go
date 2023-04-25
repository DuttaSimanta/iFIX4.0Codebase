package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertRecordTermAdditionalMap = "INSERT INTO mstrecordfield (clientid, mstorgnhirarchyid, mstrecordfieldtype, recordtermid,displayseq) VALUES (?,?,?,?,?)"
var duplicateRecordTermAdditionalMap = "SELECT count(id) total FROM  mstrecordfield WHERE clientid = ? AND mstorgnhirarchyid = ? AND mstrecordfieldtype = ? AND recordtermid = ? AND activeflg = 1 AND deleteflg = 0"

//var getRecordTermAdditionalMap = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.mstrecordfieldtype as RecordfieldtypeName, a.recordtermid as Recordtermid, b.termname as RecordtermName ,a.activeflg as Activeflg,c.name as Clientname,d.name as Mstorgnhirarchyid FROM mstrecordfield a,mstrecordterms b,mstclient c,mstorgnhierarchy d WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 and a.recordtermid=b.id and a.clientid=c.id and a.mstorgnhirarchyid=d.id and b.activeflg=1 and b.deleteflg=0 ORDER BY id DESC LIMIT ?,?"
//var getRecordTermAdditionalMapcount = "SELECT count(a.id) total FROM  mstrecordfield a,mstrecordterms b,mstclient c,mstorgnhierarchy d WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 and a.recordtermid=b.id and a.clientid=c.id and a.mstorgnhirarchyid=d.id and b.activeflg=1 and b.deleteflg=0"

//var updateRecordTermAdditionalMap = "UPDATE mstrecordfield SET mstorgnhirarchyid = ?, mstrecordfieldtype = ?, recordtermid = ? WHERE id = ? "
var deleteRecordTermAdditionalMap = "UPDATE mstrecordfield SET deleteflg = '1' WHERE id = ? "
var getRecordTermAdditionalMapMaxSeq = "select max(displayseq) from mstrecordfield where clientid=? and mstorgnhirarchyid=? and mstrecordfieldtype=? and activeflg=1 and deleteflg=0"

func (dbc DbConn) CheckDuplicateRecordTermAdditionalMap(tz *entities.RecordTermAdditionalMapEntity) (entities.RecordTermAdditionalMapEntities, error) {
	logger.Log.Println("In side CheckDuplicateRecordTermAdditionalMap")
	value := entities.RecordTermAdditionalMapEntities{}
	err := dbc.DB.QueryRow(duplicateRecordTermAdditionalMap, tz.Clientid, tz.Mstorgnhirarchyid, tz.RecordfieldtypeName, tz.Recordtermid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateRecordTermAdditionalMap Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertRecordTermAdditionalMap(tz *entities.RecordTermAdditionalMapEntity) (int64, error) {
	logger.Log.Println("In side InsertRecordTermAdditionalMap")
	logger.Log.Println("Query -->", insertRecordTermAdditionalMap)
	stmt, err := dbc.DB.Prepare(insertRecordTermAdditionalMap)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertRecordTermAdditionalMap Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.RecordfieldtypeName, tz.Recordtermid, tz.DisplaySeq)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.RecordfieldtypeName, tz.Recordtermid, tz.DisplaySeq)
	if err != nil {
		logger.Log.Println("InsertRecordTermAdditionalMap Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllRecordTermAdditionalMap(page *entities.RecordTermAdditionalMapEntity, OrgnType int64) ([]entities.RecordTermAdditionalMapEntity, error) {
	logger.Log.Println("In side GelAllRecordTermAdditionalMap")
	values := []entities.RecordTermAdditionalMapEntity{}
	var params []interface{}
	var getRecordTermAdditionalMap string

	if OrgnType == 1 {
		getRecordTermAdditionalMap = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.mstrecordfieldtype as RecordfieldtypeName, a.recordtermid as Recordtermid, b.termname as RecordtermName ,a.activeflg as Activeflg,c.name as Clientname,d.name as Mstorgnhirarchyid FROM mstrecordfield a,mstrecordterms b,mstclient c,mstorgnhierarchy d WHERE  a.deleteflg =0 and a.activeflg=1 and a.recordtermid=b.id and a.clientid=c.id and a.mstorgnhirarchyid=d.id and b.activeflg=1 and b.deleteflg=0 ORDER BY id DESC LIMIT ?,?"
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else if OrgnType == 2 {
		getRecordTermAdditionalMap = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.mstrecordfieldtype as RecordfieldtypeName, a.recordtermid as Recordtermid, b.termname as RecordtermName ,a.activeflg as Activeflg,c.name as Clientname,d.name as Mstorgnhirarchyid FROM mstrecordfield a,mstrecordterms b,mstclient c,mstorgnhierarchy d WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 and a.recordtermid=b.id and a.clientid=c.id and a.mstorgnhirarchyid=d.id and b.activeflg=1 and b.deleteflg=0 ORDER BY id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else {
		getRecordTermAdditionalMap = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.mstrecordfieldtype as RecordfieldtypeName, a.recordtermid as Recordtermid, b.termname as RecordtermName ,a.activeflg as Activeflg,c.name as Clientname,d.name as Mstorgnhirarchyid FROM mstrecordfield a,mstrecordterms b,mstclient c,mstorgnhierarchy d WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 and a.recordtermid=b.id and a.clientid=c.id and a.mstorgnhirarchyid=d.id and b.activeflg=1 and b.deleteflg=0 ORDER BY id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Mstorgnhirarchyid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	}
	rows, err := dbc.DB.Query(getRecordTermAdditionalMap, params...)
	//rows, err := dbc.DB.Query(getRecordTermAdditionalMap, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllRecordTermAdditionalMap Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.RecordTermAdditionalMapEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.RecordfieldtypeName, &value.Recordtermid, &value.RecordtermName, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname)
		values = append(values, value)
	}
	return values, nil
}

// func (dbc DbConn) UpdateRecordTermAdditionalMap(tz *entities.RecordTermAdditionalMapEntity) error {
// 	logger.Log.Println("In side UpdateRecordTermAdditionalMap")
// 	stmt, err := dbc.DB.Prepare(updateRecordTermAdditionalMap)
// 	defer stmt.Close()
// 	if err != nil {
// 		logger.Log.Println("UpdateRecordTermAdditionalMap Prepare Statement  Error", err)
// 		return err
// 	}
// 	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.RecordfieldtypeName, tz.Recordtermid, tz.Id)
// 	if err != nil {
// 		logger.Log.Println("UpdateRecordTermAdditionalMap Execute Statement  Error", err)
// 		return err
// 	}
// 	return nil
// }

func (dbc DbConn) DeleteRecordTermAdditionalMap(tz *entities.RecordTermAdditionalMapEntity) error {
	logger.Log.Println("In side DeleteRecordTermAdditionalMap")
	stmt, err := dbc.DB.Prepare(deleteRecordTermAdditionalMap)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteRecordTermAdditionalMap Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteRecordTermAdditionalMap Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetRecordTermAdditionalMapCount(page *entities.RecordTermAdditionalMapEntity, OrgnType int64) (entities.RecordTermAdditionalMapEntities, error) {
	logger.Log.Println("In side GetRecordTermAdditionalMapCount")
	value := entities.RecordTermAdditionalMapEntities{}
	var params []interface{}
	var getRecordTermAdditionalMapcount string
	logger.Log.Println("orgntype:", OrgnType)
	if OrgnType == 1 {
		getRecordTermAdditionalMapcount = "SELECT count(a.id) total FROM  mstrecordfield a,mstrecordterms b,mstclient c,mstorgnhierarchy d WHERE  a.deleteflg =0 and a.activeflg=1 and a.recordtermid=b.id and a.clientid=c.id and a.mstorgnhirarchyid=d.id and b.activeflg=1 and b.deleteflg=0"

	} else if OrgnType == 2 {
		getRecordTermAdditionalMapcount = "SELECT count(a.id) total FROM  mstrecordfield a,mstrecordterms b,mstclient c,mstorgnhierarchy d WHERE a.clientid = ?   AND a.deleteflg =0 and a.activeflg=1 and a.recordtermid=b.id and a.clientid=c.id and a.mstorgnhirarchyid=d.id and b.activeflg=1 and b.deleteflg=0"
		params = append(params, page.Clientid)

	} else {
		getRecordTermAdditionalMapcount = "SELECT count(a.id) total FROM  mstrecordfield a,mstrecordterms b,mstclient c,mstorgnhierarchy d WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 and a.recordtermid=b.id and a.clientid=c.id and a.mstorgnhirarchyid=d.id and b.activeflg=1 and b.deleteflg=0"
		params = append(params, page.Clientid)
		params = append(params, page.Mstorgnhirarchyid)

	}

	err := dbc.DB.QueryRow(getRecordTermAdditionalMapcount, params...).Scan(&value.Total)
	//err := dbc.DB.QueryRow(getRecordTermAdditionalMapcount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetRecordTermAdditionalMapCount Get Statement Prepare Error", err)
		return value, err
	}
}
func (dbc DbConn) GetRecordTermAdditionalMapMaxSeq(tz *entities.RecordTermAdditionalMapEntity) (int64, error) {
	logger.Log.Println("In side GetRecordTermAdditionalMapmaxseq")
	//value := entities.RecordTermAdditionalMapEntities{}
	var x int64
	var seq interface{}
	err := dbc.DB.QueryRow(getRecordTermAdditionalMapMaxSeq, tz.Clientid, tz.Mstorgnhirarchyid, tz.RecordfieldtypeName).Scan(&seq)
	switch err {
	case sql.ErrNoRows:
		x = 0
		return x, nil
	case nil:
		if seq != nil {
			return seq.(int64), nil
		}
		return x, nil

	default:
		logger.Log.Println("GetRecordTermAdditionalMapmaxseq Get Statement Prepare Error", err)
		return x, err
	}
}
func (dbc DbConn) GetAdditionalTab() ([]entities.AdditionalTabEntity, error) {
	logger.Log.Println("In side GelAllRecordTermAdditionalMap")
	values := []entities.AdditionalTabEntity{}
	rows, err := dbc.DB.Query("SELECT id,tabname from mstadditionaltabsname")
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllRecordTermAdditionalMap Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.AdditionalTabEntity{}
		rows.Scan(&value.Id, &value.TabName)
		values = append(values, value)
	}
	return values, nil
}
