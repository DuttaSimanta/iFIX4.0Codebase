package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
)

var insertMststate = "INSERT INTO mststate (clientid, mstorgnhirarchyid, statetypeid, statename, description,seqno) VALUES (?,?,?,?,?,?)"
var duplicateMststate = "SELECT count(id) total FROM  mststate WHERE clientid = ? AND mstorgnhirarchyid = ? AND statetypeid = ? AND statename = ? AND description = ? AND deleteflg = 0"
var getMststate = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.statetypeid as Statetypeid, a.statename as Statename, a.description as Description, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.statetypename as Statetypename FROM mststate a,mstclient b,mstorgnhierarchy c,mststatetype d WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.statetypeid=d.id and d.activeflg=1 and d.deleteflg=0 ORDER BY a.id DESC LIMIT ?,?"
var getMststatecount = "SELECT count(a.id) as total FROM mststate a,mstclient b,mstorgnhierarchy c,mststatetype d WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.statetypeid=d.id"
var updateMststate = "UPDATE mststate SET mstorgnhirarchyid = ?, statetypeid = ?, statename = ?, description = ? WHERE id = ? "
var deleteMststate = "UPDATE mststate SET deleteflg = '1' WHERE id = ? "
var lastseqbystate = "SELECT max(seqno)  from mststate where clientid=? and mstorgnhirarchyid = ?  and activeflg=1 and deleteflg=0"
var duplicatestateseq = "SELECT count(id) total from mststate where clientid=? and mstorgnhirarchyid = ? and  seqno=? and deleteflg=0"

func (mdao DbConn) CheckDuplicatestateseq(tz *entities.MststateEntity) (entities.MststateEntities, error) {
	log.Println("In side dao")
	value := entities.MststateEntities{}
	err := mdao.DB.QueryRow(duplicatestateseq, tz.Clientid,tz.Mstorgnhirarchyid,tz.Seqno).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		log.Print("CheckDuplicatestateseq Get Statement Prepare Error", err)
		return value, err
	}
}
func (mdao DbConn) GetLastSeqFromstate(tz *entities.MststateEntity) ([]entities.MststateEntity, error) {
	log.Println("In side dao")
	values := []entities.MststateEntity{}
	rows, err := mdao.DB.Query(lastseqbystate, tz.Clientid,tz.Mstorgnhirarchyid)

	if err != nil {
		log.Print("GetLastSeqFromstate Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		value := entities.MststateEntity{}
		rows.Scan(&value.Seqno)
		values = append(values, value)
	}
	return values, nil
}
func (dbc DbConn) CheckDuplicateMststate(tz *entities.MststateEntity) (entities.MststateEntities, error) {
	logger.Log.Println("In side CheckDuplicateMststate")
	value := entities.MststateEntities{}
	err := dbc.DB.QueryRow(duplicateMststate, tz.Clientid, tz.Mstorgnhirarchyid, tz.Statetypeid, tz.Statename, tz.Description).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMststate Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertMststate(tz *entities.MststateEntity) (int64, error) {
	logger.Log.Println("In side InsertMststate")
	logger.Log.Println("Query -->", insertMststate)
	stmt, err := dbc.DB.Prepare(insertMststate)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMststate Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Statetypeid, tz.Statename, tz.Description)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Statetypeid, tz.Statename, tz.Description,tz.Seqno)
	if err != nil {
		logger.Log.Println("InsertMststate Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllMststate(tz *entities.MststateEntity, OrgnType int64) ([]entities.MststateEntity, error) {
	logger.Log.Println("In side GelAllMststate")
	values := []entities.MststateEntity{}
	var getMststate string
	var params []interface{}
	if OrgnType == 1 {
		getMststate = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.statetypeid as Statetypeid, a.statename as Statename, a.description as Description, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.statetypename as Statetypename FROM mststate a,mstclient b,mstorgnhierarchy c,mststatetype d WHERE a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.statetypeid=d.id and d.activeflg=1 and d.deleteflg=0 ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		getMststate = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.statetypeid as Statetypeid, a.statename as Statename, a.description as Description, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.statetypename as Statetypename FROM mststate a,mstclient b,mstorgnhierarchy c,mststatetype d WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.statetypeid=d.id and d.activeflg=1 and d.deleteflg=0 ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		getMststate = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.statetypeid as Statetypeid, a.statename as Statename, a.description as Description, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.statetypename as Statetypename FROM mststate a,mstclient b,mstorgnhierarchy c,mststatetype d WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.statetypeid=d.id and d.activeflg=1 and d.deleteflg=0 ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}

	rows, err := dbc.DB.Query(getMststate, params...)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMststate Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MststateEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Statetypeid, &value.Statename, &value.Description, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Statetypename)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateMststate(tz *entities.MststateEntity) error {
	logger.Log.Println("In side UpdateMststate")
	stmt, err := dbc.DB.Prepare(updateMststate)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMststate Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Statetypeid, tz.Statename, tz.Description, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMststate Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMststate(tz *entities.MststateEntity) error {
	logger.Log.Println("In side DeleteMststate")
	stmt, err := dbc.DB.Prepare(deleteMststate)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMststate Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMststate Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMststateCount(tz *entities.MststateEntity, OrgnTypeID int64) (entities.MststateEntities, error) {
	logger.Log.Println("In side GetMststateCount")
	value := entities.MststateEntities{}
	var getMststatecount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getMststatecount = "SELECT count(a.id) as total FROM mststate a,mstclient b,mstorgnhierarchy c,mststatetype d WHERE a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.statetypeid=d.id"
	} else if OrgnTypeID == 2 {
		getMststatecount = "SELECT count(a.id) as total FROM mststate a,mstclient b,mstorgnhierarchy c,mststatetype d WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.statetypeid=d.id"
		params = append(params, tz.Clientid)
	} else {
		getMststatecount = "SELECT count(a.id) as total FROM mststate a,mstclient b,mstorgnhierarchy c,mststatetype d WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.statetypeid=d.id"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getMststatecount, params...).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMststateCount Get Statement Prepare Error", err)
		return value, err
	}
}
