package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertMstslastartendcriteria = "INSERT INTO mstslastartendcriteria (clientid, mstorgnhirarchyid, workflowid, stateid, slaid, startorend) VALUES (?,?,?,?,?,?)"
var duplicateMstslastartendcriteria = "SELECT count(id) total FROM  mstslastartendcriteria WHERE clientid = ? AND mstorgnhirarchyid = ? AND workflowid = ? AND stateid = ? AND slaid = ? AND startorend = ? AND deleteflg = 0"
var getMstslastartendcriteria = "SELECT id as Id, clientid as Clientid, mstorgnhirarchyid as Mstorgnhirarchyid, workflowid as Workflowid, stateid as Stateid, slaid as Slaid, startorend as Startorend, activeflg as Activeflg,(SELECT name FROM mstclient WHERE clientid=id) as Clientname,(SELECT name FROM mstorgnhierarchy WHERE mstorgnhirarchyid=id) as Mstorgnhirarchyname,(SELECT statename FROM mststate WHERE id=stateid and deleteflg =0 and activeflg=1) as Statename,(SELECT statetypeid FROM mststate WHERE id=stateid and deleteflg =0 and activeflg=1) as statetypeid ,(SELECT slaname FROM mstclientsla WHERE id=slaid and deleteflg =0 and activeflg=1) as Slaname,(SELECT processname FROM mstprocess WHERE id=workflowid and deleteflg =0 and activeflg=1) as Workflowname FROM mstslastartendcriteria WHERE clientid = ? AND mstorgnhirarchyid = ?  AND deleteflg =0 and activeflg=1 ORDER BY id DESC LIMIT ?,?"
var getMstslastartendcriteriacount = "SELECT count(a.id) total FROM  mstslastartendcriteria a,mstclient b,mstorgnhierarchy c,mstprocess d,mstclientsla e,mststate f WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND  a.deleteflg =0 and a.activeflg=1 and a.clientid=b.id and a.mstorgnhirarchyid=c.id and a.workflowid=d.id and d.deleteflg =0 and d.activeflg=1 and a.slaid=e.id and e.deleteflg =0 and e.activeflg=1 and a.stateid=f.id and f.deleteflg =0 and f.activeflg=1"
var updateMstslastartendcriteria = "UPDATE mstslastartendcriteria SET mstorgnhirarchyid = ?, workflowid = ?, stateid = ?, slaid = ?, startorend = ? WHERE id = ? "
var deleteMstslastartendcriteria = "UPDATE mstslastartendcriteria SET deleteflg = '1' WHERE id = ? "
var getslanameagainstworkflowid = "select c.id,c.slaname from mstprocessrecordmap a,mstslafullfillmentcriteria b,mstclientsla c where a.mstprocessid=? and a.clientid=? and a.mstorgnhirarchyid=? and a.deleteflg=0 and a.activeflg=1 and a.recorddiffid=b.mstrecorddifferentiationworkingcatid and b.deleteflg=0 and b.activeflg=1 and b.slaid=c.id and c.deleteflg=0 and c.activeflg=1"

func (dbc DbConn) CheckDuplicateMstslastartendcriteria(tz *entities.MstslastartendcriteriaEntity) (entities.MstslastartendcriteriaEntities, error) {
	logger.Log.Println("In side CheckDuplicateMstslastartendcriteria")
	value := entities.MstslastartendcriteriaEntities{}
	err := dbc.DB.QueryRow(duplicateMstslastartendcriteria, tz.Clientid, tz.Mstorgnhirarchyid, tz.Workflowid, tz.Stateid, tz.Slaid, tz.Startorend).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMstslastartendcriteria Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertMstslastartendcriteria(tz *entities.MstslastartendcriteriaEntity) (int64, error) {
	logger.Log.Println("In side InsertMstslastartendcriteria")
	logger.Log.Println("Query -->", insertMstslastartendcriteria)
	stmt, err := dbc.DB.Prepare(insertMstslastartendcriteria)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMstslastartendcriteria Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Workflowid, tz.Stateid, tz.Slaid, tz.Startorend)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Workflowid, tz.Stateid, tz.Slaid, tz.Startorend)
	if err != nil {
		logger.Log.Println("InsertMstslastartendcriteria Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllMstslastartendcriteria(page *entities.MstslastartendcriteriaEntity) ([]entities.MstslastartendcriteriaEntity, error) {
	logger.Log.Println("In side GelAllMstslastartendcriteria")
	values := []entities.MstslastartendcriteriaEntity{}
	rows, err := dbc.DB.Query(getMstslastartendcriteria, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstslastartendcriteria Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstslastartendcriteriaEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Workflowid, &value.Stateid, &value.Slaid, &value.Startorend, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Statename,&value.Statetypeid, &value.Slaname, &value.Workflowname)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateMstslastartendcriteria(tz *entities.MstslastartendcriteriaEntity) error {
	logger.Log.Println("In side UpdateMstslastartendcriteria")
	stmt, err := dbc.DB.Prepare(updateMstslastartendcriteria)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMstslastartendcriteria Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Workflowid, tz.Stateid, tz.Slaid, tz.Startorend, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMstslastartendcriteria Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMstslastartendcriteria(tz *entities.MstslastartendcriteriaEntity) error {
	logger.Log.Println("In side DeleteMstslastartendcriteria")
	stmt, err := dbc.DB.Prepare(deleteMstslastartendcriteria)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMstslastartendcriteria Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMstslastartendcriteria Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMstslastartendcriteriaCount(tz *entities.MstslastartendcriteriaEntity) (entities.MstslastartendcriteriaEntities, error) {
	logger.Log.Println("In side GetMstslastartendcriteriaCount")
	value := entities.MstslastartendcriteriaEntities{}
	err := dbc.DB.QueryRow(getMstslastartendcriteriacount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMstslastartendcriteriaCount Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) GetSlanameagainstworkflowid(page *entities.MstslastartendcriteriaEntity) ([]entities.MstslanameagaistworkflowEntity, error) {
	logger.Log.Println("In side GetSlanameagainstworkflowid")
	values := []entities.MstslanameagaistworkflowEntity{}
	rows, err := dbc.DB.Query(getslanameagainstworkflowid, page.Workflowid, page.Clientid, page.Mstorgnhirarchyid)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetSlanameagainstworkflowid Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstslanameagaistworkflowEntity{}
		rows.Scan(&value.Id, &value.Slaname)
		values = append(values, value)
	}
	return values, nil
}
