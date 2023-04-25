package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertMapprocessstate = "INSERT INTO mapprocessstate (clientid, mstorgnhirarchyid, statetid, processid) VALUES (?,?,?,?)"
var duplicateMapprocessstate = "SELECT count(id) total FROM  mapprocessstate WHERE clientid = ? AND mstorgnhirarchyid = ? AND statetid = ? AND processid = ? AND deleteflg = 0 AND activeflg=1"

// var getMapprocessstate = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.statetid as Statetid,e.statetypeid , a.processid as Processid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.processname as Processname,e.statename as Statename FROM mapprocessstate a,mstclient b,mstorgnhierarchy c,mstprocess d,mststate e WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.processid=d.id and a.statetid=e.id  ORDER BY a.id DESC LIMIT ?,?"
// var getMapprocessstatecount = "SELECT count(a.id) as total FROM mapprocessstate a,mstclient b,mstorgnhierarchy c,mstprocess d,mststate e WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.processid=d.id and a.statetid=e.id"
var updateMapprocessstate = "UPDATE mapprocessstate SET mstorgnhirarchyid = ?, statetid = ?, processid = ? WHERE id = ? "
var deleteMapprocessstate = "UPDATE mapprocessstate SET deleteflg = '1' WHERE id = ? "
var deletestatebyprocessid = "UPDATE mapprocessstate SET deleteflg = '1' WHERE processid = ? "

func (dbc DbConn) CheckDuplicateMapprocessstate(tz *entities.MapprocessstateEntity) (entities.MapprocessstateEntities, error) {
	logger.Log.Println("In side CheckDuplicateMapprocessstate")
	value := entities.MapprocessstateEntities{}
	err := dbc.DB.QueryRow(duplicateMapprocessstate, tz.Clientid, tz.Mstorgnhirarchyid, tz.Statetid, tz.Processid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMapprocessstate Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertMapprocessstate(tz *entities.MapprocessstateEntity) (int64, error) {
	logger.Log.Println("In side InsertMapprocessstate")
	logger.Log.Println("Query -->", insertMapprocessstate)
	stmt, err := dbc.DB.Prepare(insertMapprocessstate)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMapprocessstate Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Statetid, tz.Processid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Statetid, tz.Processid)
	if err != nil {
		logger.Log.Println("InsertMapprocessstate Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}
func InsertMapprocessstatewithtransaction(tz *entities.MapprocessstateEntity, tx *sql.Tx) (int64, error) {
	logger.Log.Println("In side InsertMapprocessstate")
	logger.Log.Println("Query -->", insertMapprocessstate)
	stmt, err := tx.Prepare(insertMapprocessstate)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMapprocessstate Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Statetid, tz.Processid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Statetid, tz.Processid)
	if err != nil {
		logger.Log.Println("InsertMapprocessstate Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}
func (dbc DbConn) GetAllMapprocessstate(page *entities.MapprocessstateEntity, OrgnType int64) ([]entities.MapprocessstateEntity, error) {
	logger.Log.Println("In side GelAllMapprocessstate")
	values := []entities.MapprocessstateEntity{}
	var getMapprocessstate string
	var params []interface{}
	if OrgnType == 1 {
		getMapprocessstate = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.statetid as Statetid,e.statetypeid , a.processid as Processid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.processname as Processname,e.statename as Statename FROM mapprocessstate a,mstclient b,mstorgnhierarchy c,mstprocess d,mststate e WHERE a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.processid=d.id and a.statetid=e.id  ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else if OrgnType == 2 {
		getMapprocessstate = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.statetid as Statetid,e.statetypeid , a.processid as Processid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.processname as Processname,e.statename as Statename FROM mapprocessstate a,mstclient b,mstorgnhierarchy c,mstprocess d,mststate e WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.processid=d.id and a.statetid=e.id  ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else {
		getMapprocessstate = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.statetid as Statetid,e.statetypeid , a.processid as Processid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.processname as Processname,e.statename as Statename FROM mapprocessstate a,mstclient b,mstorgnhierarchy c,mstprocess d,mststate e WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.processid=d.id and a.statetid=e.id  ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Mstorgnhirarchyid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	}

	rows, err := dbc.DB.Query(getMapprocessstate, params...)

	// rows, err := dbc.DB.Query(getMapprocessstate, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMapprocessstate Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MapprocessstateEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Statetid, &value.Statetypeid, &value.Processid, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Processname, &value.Statename)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateMapprocessstate(tz *entities.MapprocessstateEntity) error {
	logger.Log.Println("In side UpdateMapprocessstate")
	stmt, err := dbc.DB.Prepare(updateMapprocessstate)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMapprocessstate Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Statetid, tz.Processid, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMapprocessstate Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMapprocessstate(tz *entities.MapprocessstateEntity) error {
	logger.Log.Println("In side DeleteMapprocessstate")
	stmt, err := dbc.DB.Prepare(deleteMapprocessstate)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMapprocessstate Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMapprocessstate Execute Statement  Error", err)
		return err
	}
	return nil
}
func Deletestatebyprocessid(tz *entities.MapprocessstateEntity, tx *sql.Tx) error {
	logger.Log.Println("In side Deletestatebyprocessid")
	stmt, err := tx.Prepare(deletestatebyprocessid)

	if err != nil {
		logger.Log.Println("Deletestatebyprocessid Prepare Statement  Error", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(tz.Processid)
	if err != nil {
		logger.Log.Println("Deletestatebyprocessid Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMapprocessstateCount(tz *entities.MapprocessstateEntity, OrgnTypeID int64) (entities.MapprocessstateEntities, error) {
	logger.Log.Println("In side GetMapprocessstateCount")
	value := entities.MapprocessstateEntities{}
	var getMapprocessstatecount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getMapprocessstatecount = "SELECT count(a.id) as total FROM mapprocessstate a,mstclient b,mstorgnhierarchy c,mstprocess d,mststate e WHERE a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.processid=d.id and a.statetid=e.id"
	} else if OrgnTypeID == 2 {
		getMapprocessstatecount = "SELECT count(a.id) as total FROM mapprocessstate a,mstclient b,mstorgnhierarchy c,mstprocess d,mststate e WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.processid=d.id and a.statetid=e.id"
		params = append(params, tz.Clientid)
	} else {
		getMapprocessstatecount = "SELECT count(a.id) as total FROM mapprocessstate a,mstclient b,mstorgnhierarchy c,mstprocess d,mststate e WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.processid=d.id and a.statetid=e.id"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getMapprocessstatecount, params...).Scan(&value.Total)

	// err := dbc.DB.QueryRow(getMapprocessstatecount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMapprocessstateCount Get Statement Prepare Error", err)
		return value, err
	}
}
