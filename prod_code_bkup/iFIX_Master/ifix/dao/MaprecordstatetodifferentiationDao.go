package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertMaprecordstatetodifferentiation = "INSERT INTO maprecordstatetodifferentiation (clientid, mstorgnhirarchyid, recorddifftypeid, recorddiffid, mststatetypeid, mststateid) VALUES (?,?,?,?,?,?)"
var duplicateMaprecordstatetodifferentiation = "SELECT count(id) total FROM  maprecordstatetodifferentiation WHERE clientid = ? AND mstorgnhirarchyid = ? AND recorddifftypeid = ? AND recorddiffid = ?  AND deleteflg = 0 and activeflg=1"

// var getMaprecordstatetodifferentiation = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.recorddifftypeid as Recorddifftypeid, a.recorddiffid as Recorddiffid, a.mststatetypeid as Mststatetypeid, a.mststateid as Mststateid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Recorddifferentiationtypename,e.name as Recorddifferentiationname,f.statetypename as Statetypename,g.statename as Statename FROM maprecordstatetodifferentiation a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mststatetype f,mststate g WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.recorddifftypeid=d.id and a.recorddiffid=e.id and a.mststatetypeid=f.id  and a.mststateid=g.id  ORDER BY a.id DESC LIMIT ?,?"
// var getMaprecordstatetodifferentiationcount = "SELECT count(a.id) as total FROM maprecordstatetodifferentiation a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mststatetype f,mststate g WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.recorddifftypeid=d.id and a.recorddiffid=e.id and a.mststatetypeid=f.id  and a.mststateid=g.id"
var updateMaprecordstatetodifferentiation = "UPDATE maprecordstatetodifferentiation SET mstorgnhirarchyid = ?, recorddifftypeid = ?, recorddiffid = ?, mststatetypeid = ?, mststateid = ? WHERE id = ? "
var deleteMaprecordstatetodifferentiation = "UPDATE maprecordstatetodifferentiation SET deleteflg = '1' WHERE id = ? "

func (dbc DbConn) CheckDuplicateMaprecordstatetodifferentiation(tz *entities.MaprecordstatetodifferentiationEntity) (entities.MaprecordstatetodifferentiationEntities, error) {
	logger.Log.Println("In side CheckDuplicateMaprecordstatetodifferentiation")
	value := entities.MaprecordstatetodifferentiationEntities{}
	err := dbc.DB.QueryRow(duplicateMaprecordstatetodifferentiation, tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMaprecordstatetodifferentiation Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertMaprecordstatetodifferentiation(tz *entities.MaprecordstatetodifferentiationEntity) (int64, error) {
	logger.Log.Println("In side InsertMaprecordstatetodifferentiation")
	logger.Log.Println("Query -->", insertMaprecordstatetodifferentiation)
	stmt, err := dbc.DB.Prepare(insertMaprecordstatetodifferentiation)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMaprecordstatetodifferentiation Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Mststatetypeid, tz.Mststateid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Mststatetypeid, tz.Mststateid)
	if err != nil {
		logger.Log.Println("InsertMaprecordstatetodifferentiation Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllMaprecordstatetodifferentiation(page *entities.MaprecordstatetodifferentiationEntity, OrgnType int64) ([]entities.MaprecordstatetodifferentiationEntity, error) {
	logger.Log.Println("In side GelAllMaprecordstatetodifferentiation")
	values := []entities.MaprecordstatetodifferentiationEntity{}
	var getMaprecordstatetodifferentiation string
	var params []interface{}
	if OrgnType == 1 {
		getMaprecordstatetodifferentiation = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.recorddifftypeid as Recorddifftypeid, a.recorddiffid as Recorddiffid, a.mststatetypeid as Mststatetypeid, a.mststateid as Mststateid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Recorddifferentiationtypename,e.name as Recorddifferentiationname,f.statetypename as Statetypename,g.statename as Statename FROM maprecordstatetodifferentiation a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mststatetype f,mststate g WHERE  a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.recorddifftypeid=d.id and a.recorddiffid=e.id and a.mststatetypeid=f.id  and a.mststateid=g.id  ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else if OrgnType == 2 {
		getMaprecordstatetodifferentiation = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.recorddifftypeid as Recorddifftypeid, a.recorddiffid as Recorddiffid, a.mststatetypeid as Mststatetypeid, a.mststateid as Mststateid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Recorddifferentiationtypename,e.name as Recorddifferentiationname,f.statetypename as Statetypename,g.statename as Statename FROM maprecordstatetodifferentiation a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mststatetype f,mststate g WHERE a.clientid = ? AND  a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.recorddifftypeid=d.id and a.recorddiffid=e.id and a.mststatetypeid=f.id  and a.mststateid=g.id  ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else {
		getMaprecordstatetodifferentiation = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.recorddifftypeid as Recorddifftypeid, a.recorddiffid as Recorddiffid, a.mststatetypeid as Mststatetypeid, a.mststateid as Mststateid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Recorddifferentiationtypename,e.name as Recorddifferentiationname,f.statetypename as Statetypename,g.statename as Statename FROM maprecordstatetodifferentiation a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mststatetype f,mststate g WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.recorddifftypeid=d.id and a.recorddiffid=e.id and a.mststatetypeid=f.id  and a.mststateid=g.id  ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Mstorgnhirarchyid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	}
	rows, err := dbc.DB.Query(getMaprecordstatetodifferentiation, params...)

	// rows, err := dbc.DB.Query(getMaprecordstatetodifferentiation, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMaprecordstatetodifferentiation Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MaprecordstatetodifferentiationEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Recorddifftypeid, &value.Recorddiffid, &value.Mststatetypeid, &value.Mststateid, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Recorddifferentiationtypename, &value.Recorddifferentiationname, &value.Statetypename, &value.Statename)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateMaprecordstatetodifferentiation(tz *entities.MaprecordstatetodifferentiationEntity) error {
	logger.Log.Println("In side UpdateMaprecordstatetodifferentiation")
	stmt, err := dbc.DB.Prepare(updateMaprecordstatetodifferentiation)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMaprecordstatetodifferentiation Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Mststatetypeid, tz.Mststateid, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMaprecordstatetodifferentiation Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMaprecordstatetodifferentiation(tz *entities.MaprecordstatetodifferentiationEntity) error {
	logger.Log.Println("In side DeleteMaprecordstatetodifferentiation")
	stmt, err := dbc.DB.Prepare(deleteMaprecordstatetodifferentiation)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMaprecordstatetodifferentiation Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMaprecordstatetodifferentiation Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMaprecordstatetodifferentiationCount(tz *entities.MaprecordstatetodifferentiationEntity, OrgnTypeID int64) (entities.MaprecordstatetodifferentiationEntities, error) {
	logger.Log.Println("In side GetMaprecordstatetodifferentiationCount")
	value := entities.MaprecordstatetodifferentiationEntities{}
	var getMaprecordstatetodifferentiationcount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getMaprecordstatetodifferentiationcount = "SELECT count(a.id) as total FROM maprecordstatetodifferentiation a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mststatetype f,mststate g WHERE a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.recorddifftypeid=d.id and a.recorddiffid=e.id and a.mststatetypeid=f.id  and a.mststateid=g.id"
	} else if OrgnTypeID == 2 {
		getMaprecordstatetodifferentiationcount = "SELECT count(a.id) as total FROM maprecordstatetodifferentiation a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mststatetype f,mststate g WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.recorddifftypeid=d.id and a.recorddiffid=e.id and a.mststatetypeid=f.id  and a.mststateid=g.id"
		params = append(params, tz.Clientid)
	} else {
		getMaprecordstatetodifferentiationcount = "SELECT count(a.id) as total FROM maprecordstatetodifferentiation a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mststatetype f,mststate g WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.recorddifftypeid=d.id and a.recorddiffid=e.id and a.mststatetypeid=f.id  and a.mststateid=g.id"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getMaprecordstatetodifferentiationcount, params...).Scan(&value.Total)

	// err := dbc.DB.QueryRow(getMaprecordstatetodifferentiationcount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMaprecordstatetodifferentiationCount Get Statement Prepare Error", err)
		return value, err
	}
}
