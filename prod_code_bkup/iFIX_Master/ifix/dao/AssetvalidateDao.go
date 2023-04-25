package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertAssetvalidate = "INSERT INTO mapassetvalidate (clientid, mstorgnhirarchyid, mstdifferentiationtypeid, mstdifferentiationid, validationrule) VALUES (?,?,?,?,?)"
var duplicateAssetvalidate = "SELECT count(id) total FROM  mapassetvalidate WHERE clientid = ? AND mstorgnhirarchyid = ? AND mstdifferentiationtypeid = ? AND mstdifferentiationid = ? AND validationrule = ? AND deleteflg = 0 AND activeflg=1"

// var getAssetvalidate = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.mstdifferentiationtypeid as Mstdifferentiationtypeid, a.mstdifferentiationid as Mstdifferentiationid, a.validationrule as Validationrule, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Recorddifferentiationtypename,e.name as Recorddifferentiationname FROM mapassetvalidate a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.mstdifferentiationtypeid=d.id and a.mstdifferentiationid=e.id ORDER BY a.id DESC LIMIT ?,?"
// var getAssetvalidatecount = "SELECT count(a.id) as total FROM mapassetvalidate a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.mstdifferentiationtypeid=d.id and a.mstdifferentiationid=e.id"
var updateAssetvalidate = "UPDATE mapassetvalidate SET mstorgnhirarchyid = ?, mstdifferentiationtypeid = ?, mstdifferentiationid = ?, validationrule = ? WHERE id = ? "
var deleteAssetvalidate = "UPDATE mapassetvalidate SET deleteflg = '1' WHERE id = ? "

func (dbc DbConn) CheckDuplicateAssetvalidate(tz *entities.AssetvalidateEntity) (entities.AssetvalidateEntities, error) {
	logger.Log.Println("In side CheckDuplicateAssetvalidate")
	value := entities.AssetvalidateEntities{}
	err := dbc.DB.QueryRow(duplicateAssetvalidate, tz.Clientid, tz.Mstorgnhirarchyid, tz.Mstdifferentiationtypeid, tz.Mstdifferentiationid, tz.Validationrule).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateAssetvalidate Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertAssetvalidate(tz *entities.AssetvalidateEntity) (int64, error) {
	logger.Log.Println("In side InsertAssetvalidate")
	logger.Log.Println("Query -->", insertAssetvalidate)
	stmt, err := dbc.DB.Prepare(insertAssetvalidate)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertAssetvalidate Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Mstdifferentiationtypeid, tz.Mstdifferentiationid, tz.Validationrule)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Mstdifferentiationtypeid, tz.Mstdifferentiationid, tz.Validationrule)
	if err != nil {
		logger.Log.Println("InsertAssetvalidate Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllAssetvalidate(tz *entities.AssetvalidateEntity, OrgnType int64) ([]entities.AssetvalidateEntity, error) {
	logger.Log.Println("In side GelAllAssetvalidate")
	values := []entities.AssetvalidateEntity{}
	var getAssetvalidate string
	var params []interface{}
	if OrgnType == 1 {
		getAssetvalidate = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.mstdifferentiationtypeid as Mstdifferentiationtypeid, a.mstdifferentiationid as Mstdifferentiationid, a.validationrule as Validationrule, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Recorddifferentiationtypename,e.name as Recorddifferentiationname FROM mapassetvalidate a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e WHERE a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.mstdifferentiationtypeid=d.id and a.mstdifferentiationid=e.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		getAssetvalidate = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.mstdifferentiationtypeid as Mstdifferentiationtypeid, a.mstdifferentiationid as Mstdifferentiationid, a.validationrule as Validationrule, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Recorddifferentiationtypename,e.name as Recorddifferentiationname FROM mapassetvalidate a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.mstdifferentiationtypeid=d.id and a.mstdifferentiationid=e.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		getAssetvalidate = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.mstdifferentiationtypeid as Mstdifferentiationtypeid, a.mstdifferentiationid as Mstdifferentiationid, a.validationrule as Validationrule, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Recorddifferentiationtypename,e.name as Recorddifferentiationname FROM mapassetvalidate a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.mstdifferentiationtypeid=d.id and a.mstdifferentiationid=e.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}

	rows, err := dbc.DB.Query(getAssetvalidate, params...)

	// rows, err := dbc.DB.Query(getAssetvalidate, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllAssetvalidate Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.AssetvalidateEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Mstdifferentiationtypeid, &value.Mstdifferentiationid, &value.Validationrule, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Recorddifferentiationtypename, &value.Recorddifferentiationname)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateAssetvalidate(tz *entities.AssetvalidateEntity) error {
	logger.Log.Println("In side UpdateAssetvalidate")
	stmt, err := dbc.DB.Prepare(updateAssetvalidate)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateAssetvalidate Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Mstdifferentiationtypeid, tz.Mstdifferentiationid, tz.Validationrule, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateAssetvalidate Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteAssetvalidate(tz *entities.AssetvalidateEntity) error {
	logger.Log.Println("In side DeleteAssetvalidate")
	stmt, err := dbc.DB.Prepare(deleteAssetvalidate)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteAssetvalidate Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteAssetvalidate Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetAssetvalidateCount(tz *entities.AssetvalidateEntity, OrgnTypeID int64) (entities.AssetvalidateEntities, error) {
	logger.Log.Println("In side GetAssetvalidateCount")
	value := entities.AssetvalidateEntities{}
	var getAssetvalidatecount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getAssetvalidatecount = "SELECT count(a.id) as total FROM mapassetvalidate a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e WHERE a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.mstdifferentiationtypeid=d.id and a.mstdifferentiationid=e.id"
	} else if OrgnTypeID == 2 {
		getAssetvalidatecount = "SELECT count(a.id) as total FROM mapassetvalidate a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.mstdifferentiationtypeid=d.id and a.mstdifferentiationid=e.id"
		params = append(params, tz.Clientid)
	} else {
		getAssetvalidatecount = "SELECT count(a.id) as total FROM mapassetvalidate a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.mstdifferentiationtypeid=d.id and a.mstdifferentiationid=e.id"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getAssetvalidatecount, params...).Scan(&value.Total)

	// err := dbc.DB.QueryRow(getAssetvalidatecount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetAssetvalidateCount Get Statement Prepare Error", err)
		return value, err
	}
}
