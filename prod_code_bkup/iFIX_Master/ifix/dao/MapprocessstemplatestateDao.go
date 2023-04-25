package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertMapprocesstemplatestate = "INSERT INTO mapprocesstemplatestate (clientid, mstorgnhirarchyid, statetid, processid) VALUES (?,?,?,?)"
var duplicateMapprocesstemplatestate = "SELECT count(id) total FROM  mapprocesstemplatestate WHERE clientid = ? AND mstorgnhirarchyid = ? AND statetid = ? AND processid = ? AND deleteflg = 0 AND activeflg=1"
var getMapprocesstemplatestate = "SELECT a.id as Id,e.seqno, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.statetid as Statetid,e.statetypeid , a.processid as Processid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.processname as Processname,e.statename as Statename FROM mapprocesstemplatestate a,mstclient b,mstorgnhierarchy c,mstprocesstemplate d,mststate e WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.processid=d.id and a.statetid=e.id  ORDER BY a.id DESC LIMIT ?,?"
var getMapprocesstemplatestatecount = "SELECT count(a.id) as total FROM mapprocesstemplatestate a,mstclient b,mstorgnhierarchy c,mstprocesstemplate d,mststate e WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.processid=d.id and a.statetid=e.id"
var updateMapprocesstemplatestate = "UPDATE mapprocesstemplatestate SET mstorgnhirarchyid = ?, statetid = ?, processid = ?,seqno=? WHERE id = ? "
var deleteMapprocesstemplatestate = "UPDATE mapprocesstemplatestate SET deleteflg = '1' WHERE id = ? "
var deletestatebytemplateid = "UPDATE mapprocesstemplatestate SET deleteflg = '1' WHERE processid = ? "

func (dbc DbConn) CheckDuplicateMapprocesstemplatestate(tz *entities.MapprocessstateEntity) (entities.MapprocessstateEntities, error) {
	logger.Log.Println("In side CheckDuplicateMapprocesstemplatestate")
	value := entities.MapprocessstateEntities{}
	err := dbc.DB.QueryRow(duplicateMapprocesstemplatestate, tz.Clientid, tz.Mstorgnhirarchyid, tz.Statetid, tz.Processid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMapprocesstemplatestate Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertMapprocesstemplatestate(tz *entities.MapprocessstateEntity) (int64, error) {
	logger.Log.Println("In side InsertMapprocesstemplatestate")
	logger.Log.Println("Query -->", insertMapprocessstate)
	stmt, err := dbc.DB.Prepare(insertMapprocesstemplatestate)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMapprocesstemplatestate Prepare Statement  Error", err)
		return 0, err
	}
	//logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Statetid, tz.Processid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Statetid, tz.Processid)
	if err != nil {
		logger.Log.Println("InsertMapprocesstemplatestate Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllMapprocesstemplatestate(page *entities.MapprocessstateEntity) ([]entities.MapprocessstateEntity, error) {
	logger.Log.Println("In side GelAllMapprocessstate")
	values := []entities.MapprocessstateEntity{}
	rows, err := dbc.DB.Query(getMapprocesstemplatestate, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMapprocesstemplatestate Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MapprocessstateEntity{}
		rows.Scan(&value.Id,&value.Seqno, &value.Clientid, &value.Mstorgnhirarchyid, &value.Statetid, &value.Statetypeid,&value.Processid, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Processname, &value.Statename)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateMapprocesstemplatestate(tz *entities.MapprocessstateEntity) error {
	logger.Log.Println("In side UpdateMapprocessstate")
	stmt, err := dbc.DB.Prepare(updateMapprocesstemplatestate)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMapprocesstemplatestate Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Statetid, tz.Processid,tz.Seqno, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMapprocesstemplatestate Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMapprocesstemplatestate(tz *entities.MapprocessstateEntity) error {
	logger.Log.Println("In side DeleteMapprocessstate")
	stmt, err := dbc.DB.Prepare(deleteMapprocesstemplatestate)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMapprocesstemplatestate Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMapprocesstemplatestate Execute Statement  Error", err)
		return err
	}
	return nil
}
func  Deletestatebytemplateid(tz *entities.MapprocessstateEntity,tx *sql.Tx) error {
	logger.Log.Println("In side Deletestatebytemplateid")
	stmt, err := tx.Prepare(deletestatebytemplateid)

	if err != nil {
		logger.Log.Println("Deletestatebytemplateid Prepare Statement  Error", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("Deletestatebytemplateid Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMapprocesstemplatestateCount(tz *entities.MapprocessstateEntity) (entities.MapprocessstateEntities, error) {
	logger.Log.Println("In side GetMapprocessstateCount")
	value := entities.MapprocessstateEntities{}
	err := dbc.DB.QueryRow(getMapprocesstemplatestatecount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMapprocesstemplatestateCount Get Statement Prepare Error", err)
		return value, err
	}
}
