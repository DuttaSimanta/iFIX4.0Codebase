package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertMstactivity = "INSERT INTO mstactivity (clientid, mstorgnhirarchyid, actiontypeid, processid, actionname, description) VALUES (?,?,?,?,?,?)"
var duplicateMstactivity = "SELECT count(id) total FROM  mstactivity WHERE clientid = ? AND mstorgnhirarchyid = ? AND actiontypeid = ? AND processid = ? AND actionname = ? AND description = ? AND deleteflg = 0"
var getMstactivity = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.actiontypeid as Actiontypeid, a.processid as Processid, a.actionname as Actionname, a.description as Description, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.processname as Processname,e.actiontypename as Actiontypename FROM mstactivity a,mstclient b,mstorgnhierarchy c,mstprocess d,mstactiontype e WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.processid=d.id AND d.deleteflg =0 and d.activeflg=1 AND a.actiontypeid=e.id AND e.deleteflg =0 and e.activeflg=1 ORDER BY a.id DESC LIMIT ?,?"
var getMstactivitycount = "SELECT count(a.id) as Total FROM mstactivity a,mstclient b,mstorgnhierarchy c,mstprocess d,mstactiontype e WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.processid=d.id AND d.deleteflg =0 and d.activeflg=1 AND a.actiontypeid=e.id AND e.deleteflg =0 and e.activeflg=1"
var updateMstactivity = "UPDATE mstactivity SET mstorgnhirarchyid = ?, actiontypeid = ?, processid = ?, actionname = ?, description = ? WHERE id = ? "
var deleteMstactivity = "UPDATE mstactivity SET deleteflg = '1' WHERE id = ? "
var getactiontypename = "SELECT id as Id,actiontypename as Actiontypename FROM mstactiontype WHERE deleteflg =0 AND activeflg=1"
var getMstactivity1 = "select id,actiontypeid,actionname from mstactivity where clientid=? and mstorgnhirarchyid=? and processid=? and activeflg=1 and deleteflg=0"

func (dbc DbConn) Getactivitywithtype(page *entities.MstactivityEntity) ([]entities.MstactivitySingleEntity, error) {
	logger.Log.Println("In side GelAllMstactivity")
	values := []entities.MstactivitySingleEntity{}
	rows, err := dbc.DB.Query(getMstactivity1, page.Clientid, page.Mstorgnhirarchyid,page.Processid)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("Getactivitywithtype Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstactivitySingleEntity{}
		rows.Scan(&value.Id, &value.Actiontypeid,  &value.Actionname)
		values = append(values, value)
	}
	return values, nil
}
func (dbc DbConn) CheckDuplicateMstactivity(tz *entities.MstactivityEntity) (entities.MstactivityEntities, error) {
	logger.Log.Println("In side CheckDuplicateMstactivity")
	value := entities.MstactivityEntities{}
	err := dbc.DB.QueryRow(duplicateMstactivity, tz.Clientid, tz.Mstorgnhirarchyid, tz.Actiontypeid, tz.Processid, tz.Actionname, tz.Description).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMstactivity Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertMstactivity(tz *entities.MstactivityEntity) (int64, error) {
	logger.Log.Println("In side InsertMstactivity")
	logger.Log.Println("Query -->", insertMstactivity)
	stmt, err := dbc.DB.Prepare(insertMstactivity)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMstactivity Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Actiontypeid, tz.Processid, tz.Actionname, tz.Description)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Actiontypeid, tz.Processid, tz.Actionname, tz.Description)
	if err != nil {
		logger.Log.Println("InsertMstactivity Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllMstactivity(page *entities.MstactivityEntity) ([]entities.MstactivityEntity, error) {
	logger.Log.Println("In side GelAllMstactivity")
	values := []entities.MstactivityEntity{}
	rows, err := dbc.DB.Query(getMstactivity, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstactivity Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstactivityEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Actiontypeid, &value.Processid, &value.Actionname, &value.Description, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Processname, &value.Actiontypename)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateMstactivity(tz *entities.MstactivityEntity) error {
	logger.Log.Println("In side UpdateMstactivity")
	stmt, err := dbc.DB.Prepare(updateMstactivity)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMstactivity Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Actiontypeid, tz.Processid, tz.Actionname, tz.Description, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMstactivity Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMstactivity(tz *entities.MstactivityEntity) error {
	logger.Log.Println("In side DeleteMstactivity")
	stmt, err := dbc.DB.Prepare(deleteMstactivity)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMstactivity Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMstactivity Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMstactivityCount(tz *entities.MstactivityEntity) (entities.MstactivityEntities, error) {
	logger.Log.Println("In side GetMstactivityCount")
	value := entities.MstactivityEntities{}
	err := dbc.DB.QueryRow(getMstactivitycount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMstactivityCount Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) GetActiontypenames(page *entities.MstactivityEntity) ([]entities.MstactiontypeEntity, error) {
	logger.Log.Println("In side GetActiontypenames")
	values := []entities.MstactiontypeEntity{}
	rows, err := dbc.DB.Query(getactiontypename)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstactivity Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstactiontypeEntity{}
		rows.Scan(&value.Id, &value.Actiontypename)
		values = append(values, value)
	}
	return values, nil
}
