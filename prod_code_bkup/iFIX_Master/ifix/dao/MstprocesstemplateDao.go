package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertMstprocesstemplate = "INSERT INTO mstprocesstemplate (clientid, mstorgnhirarchyid, processname) VALUES (?,?,?)"
var duplicateMstprocesstemplate = "SELECT count(id) total FROM  mstprocesstemplate WHERE clientid = ? AND mstorgnhirarchyid = ? AND processname = ? and activeflg=1 and deleteflg=0"
var getMstprocesstemplate = "select a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid,  c.id as Mstprocesstoentityid, c.mstdatadictionaryfieldid as Mstdatadictionaryfieldid, a.activeflg as Activeflg, a.processname as Processname , (select name from mstorgnhierarchy where id = a.mstorgnhirarchyid) as Mstorgnhirarchyname , (select columnname from mstdatadictionaryfield where id = c.mstdatadictionaryfieldid) as Mstdatadictionaryfieldname , (select tablename from mstdatadictionarytable where id = (select tableid from mstdatadictionaryfield where id = c.mstdatadictionaryfieldid)) as Mstdatadictionarytablename,(select tableid from mstdatadictionaryfield where id = c.mstdatadictionaryfieldid) as Tableid,(select mstdatadictionarydbid from mstdatadictionarytable where id = (select tableid from mstdatadictionaryfield where id = c.mstdatadictionaryfieldid)) as Mstdatadictionarydbid from mstprocesstemplate a, mapprocesstemplatetoentity c where a.clientid = ? and a.mstorgnhirarchyid = ? and a.activeflg = '1' and a.deleteflg=0 and c.activeflg = '1' and c.deleteflg=0 and  a.id = c.mstprocessid ORDER BY a.id DESC LIMIT ?,?"
var getMstprocesstemplatecount = "select count(distinct a.id) as total from mstprocesstemplate a,  mapprocesstemplatetoentity c where  a.id = c.mstprocessid and a.clientid = ? and a.mstorgnhirarchyid = ? and a.activeflg = 1 and a.deleteflg=0 and c.activeflg = '1' and c.deleteflg=0"
var updateMstprocesstemplate = "UPDATE mstprocesstemplate SET mstorgnhirarchyid = ?, processname = ? WHERE id = ? "
var deleteMstprocesstemplate = "UPDATE mstprocesstemplate SET deleteflg = '1' WHERE id=?"

func (dbc DbConn) CheckDuplicateMstprocesstemplate(tz *entities.MstprocessEntity) (entities.MstprocessEntities, error) {
	logger.Log.Println("In side CheckDuplicateMstprocess")
	value := entities.MstprocessEntities{}
	err := dbc.DB.QueryRow(duplicateMstprocesstemplate, tz.Clientid, tz.Mstorgnhirarchyid, tz.Processname).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMstprocess Get Statement Prepare Error", err)
		return value, err
	}
}

/*func CheckDuplicateMstprocesswithtransaction(tx *sql.Tx, tz *entities.MstprocessEntity) (entities.MstprocessEntities, error) {
	logger.Log.Println("In side CheckDuplicateMstprocess")
	value := entities.MstprocessEntities{}
	err := tx.QueryRow(duplicateMstprocess, tz.Clientid, tz.Mstorgnhirarchyid, tz.Processname).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMstprocess Get Statement Prepare Error", err)
		return value, err
	}
}
*/
/*func (dbc DbConn) InsertMstprocess(tz *entities.MstprocessEntity) (int64, error) {
	logger.Log.Println("In side InsertMstprocess")
	logger.Log.Println("Query -->", insertMstprocess)
	stmt, err := dbc.DB.Prepare(insertMstprocess)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMstprocess Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Processname)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Processname)
	if err != nil {
		logger.Log.Println("InsertMstprocess Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}
*/
func InsertMstprocesstemplatewithtransaction(tx *sql.Tx, tz *entities.MstprocessEntity) (int64, error) {
	logger.Log.Println("In side InsertMstprocess")
	//logger.Log.Println("Query -->", insertMstprocess)
	stmt, err := tx.Prepare(insertMstprocesstemplate)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMstprocesstemplatewithtransaction Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Processname)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Processname)
	if err != nil {
		logger.Log.Println("InsertMstprocesstemplatewithtransaction Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllMstprocesstemplate(page *entities.MstprocessEntity) ([]entities.MstprocessEntity, error) {
	logger.Log.Println("In side GelAllMstprocess")
	values := []entities.MstprocessEntity{}
	rows, err := dbc.DB.Query(getMstprocesstemplate, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstprocess Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstprocessEntity{}
		err = rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid,  &value.Mstprocesstoentityid, &value.Mstdatadictionaryfieldid, &value.Activeflg, &value.Processname,  &value.Mstorgnhirarchyname,  &value.Mstdatadictionaryfieldname, &value.Mstdatadictionarytablename, &value.Tableid, &value.Mstdatadictionarydbid)

		values = append(values, value)
	}
	return values, nil
}
//func (dbc DbConn) UpdateMstprocess(tz *entities.MstprocessEntity) error {
//	logger.Log.Println("In side UpdateMstprocess")
//	stmt, err := dbc.DB.Prepare(updateMstprocess)
//	defer stmt.Close()
//	if err != nil {
//		logger.Log.Println("UpdateMstprocess Prepare Statement  Error", err)
//		return err
//	}
//	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Processname, tz.Id)
//	if err != nil {
//		logger.Log.Println("UpdateMstprocess Execute Statement  Error", err)
//		return err
//	}
//	return nil
//}

func UpdateMstprocesstemplatewithtransaction(tx *sql.Tx, tz *entities.MstprocessEntity) error {
	logger.Log.Println("In side UpdateMstprocess")
	stmt, err := tx.Prepare(updateMstprocesstemplate)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMstprocess Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Processname, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMstprocess Execute Statement  Error", err)
		return err
	}
	return nil
}
/*
func (dbc DbConn) DeleteMstprocess(tz *entities.MstprocessEntity) error {
	logger.Log.Println("In side DeleteMstprocess")
	stmt, err := dbc.DB.Prepare(deleteMstprocess)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMstprocess Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMstprocess Execute Statement  Error", err)
		return err
	}
	return nil
}
*/
func DeleteMstprocesstemplatewithtransaction(tx *sql.Tx, tz *entities.MstprocessEntity) error {
	logger.Log.Println("In side DeleteMstprocess")
	stmt, err := tx.Prepare(deleteMstprocesstemplate)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMstprocesstemplatewithtransaction Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMstprocesstemplatewithtransaction Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMstprocesstemplateCount(tz *entities.MstprocessEntity) (entities.MstprocessEntities, error) {
	logger.Log.Println("In side GetMstprocessCount")
	value := entities.MstprocessEntities{}
	err := dbc.DB.QueryRow(getMstprocesstemplatecount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMstprocesstemplateCount Get Statement Prepare Error", err)
		return value, err
	}
}
