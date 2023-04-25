package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertRecordconfig = "INSERT INTO mstrecordconfig (clientid, mstorgnhirarchyid, recorddifftypeid, recorddiffid, prefix, year, month,day,configurezero,isclient) VALUES (?,?,?,?,?,?,?,?,?,?)"
var duplicateRecordconfig = "SELECT count(id) total FROM  mstrecordconfig WHERE clientid = ? AND mstorgnhirarchyid = ? AND recorddifftypeid = ? AND recorddiffid = ? AND isclient=? AND deleteflg = 0 AND activeflg=1"
var getRecordconfig = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.recorddifftypeid as Recorddifftypeid, a.recorddiffid as Recorddiffid, a.prefix as Prefix, a.year as Year, a.month as Month,a.day as Day,a.configurezero as Configurezero, a.isclient as IsClient, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.name as Recorddifferentiationname,e.typename as Recorddifferentiationtypename FROM mstrecordconfig a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiation d,mstrecorddifferentiationtype e WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.recorddiffid=d.id AND d.deleteflg =0 and d.activeflg=1 and a.recorddifftypeid=e.id AND e.deleteflg =0 and e.activeflg=1  ORDER BY a.id DESC LIMIT ?,?"
var getRecordconfigcount = "SELECT count(a.id) total FROM mstrecordconfig a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiation d,mstrecorddifferentiationtype e WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.recorddiffid=d.id AND d.deleteflg =0 and d.activeflg=1 and a.recorddifftypeid=e.id AND e.deleteflg =0 and e.activeflg=1"
var updateRecordconfig = "UPDATE mstrecordconfig SET mstorgnhirarchyid = ?, recorddifftypeid = ?, recorddiffid = ?, prefix = ?, year = ?, month = ?,day=?,configurezero=?,isclient=? WHERE id = ? "
var deleteRecordconfig = "UPDATE mstrecordconfig SET deleteflg = '1' WHERE id = ? "

func (dbc DbConn) CheckDuplicateRecordconfig(tz *entities.RecordconfigEntity) (entities.RecordconfigEntities, error) {
	logger.Log.Println("In side CheckDuplicateRecordconfig")
	value := entities.RecordconfigEntities{}
	err := dbc.DB.QueryRow(duplicateRecordconfig, tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid,tz.IsClient).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateRecordconfig Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertRecordconfig(tz *entities.RecordconfigEntity) (int64, error) {
	logger.Log.Println("In side InsertRecordconfig")
	logger.Log.Println("Query -->", insertRecordconfig)
	stmt, err := dbc.DB.Prepare(insertRecordconfig)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertRecordconfig Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Prefix, tz.Year, tz.Month, tz.Day, tz.Configurezero,tz.IsClient)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Prefix, tz.Year, tz.Month, tz.Day, tz.Configurezero,tz.IsClient)
	if err != nil {
		logger.Log.Println("InsertRecordconfig Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllRecordconfig(page *entities.RecordconfigEntity) ([]entities.RecordconfigEntity, error) {
	logger.Log.Println("In side GelAllRecordconfig")
	values := []entities.RecordconfigEntity{}
	rows, err := dbc.DB.Query(getRecordconfig, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllRecordconfig Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.RecordconfigEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Recorddifftypeid, &value.Recorddiffid, &value.Prefix, &value.Year, &value.Month, &value.Day, &value.Configurezero,&value.IsClient, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Recorddifferentiationname, &value.Recorddifferentiationtypename)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateRecordconfig(tz *entities.RecordconfigEntity) error {
	logger.Log.Println("In side UpdateRecordconfig")
	stmt, err := dbc.DB.Prepare(updateRecordconfig)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateRecordconfig Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Prefix, tz.Year, tz.Month, tz.Day, tz.Configurezero,tz.IsClient, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateRecordconfig Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteRecordconfig(tz *entities.RecordconfigEntity) error {
	logger.Log.Println("In side DeleteRecordconfig")
	stmt, err := dbc.DB.Prepare(deleteRecordconfig)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteRecordconfig Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteRecordconfig Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetRecordconfigCount(tz *entities.RecordconfigEntity) (entities.RecordconfigEntities, error) {
	logger.Log.Println("In side GetRecordconfigCount")
	value := entities.RecordconfigEntities{}
	err := dbc.DB.QueryRow(getRecordconfigcount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetRecordconfigCount Get Statement Prepare Error", err)
		return value, err
	}
}
