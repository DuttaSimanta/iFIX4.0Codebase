package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertAssetdifferentiation = "INSERT INTO mapassetdifferentiation (clientid, mstorgnhirarchyid, trnassetid, mstdifferentiationtypeid, mstdifferentiationid, value) VALUES (?,?,?,?,?,?)"
var duplicateAssetdifferentiation = "SELECT count(id) total FROM  mapassetdifferentiation WHERE clientid = ? AND mstorgnhirarchyid = ? AND trnassetid = ? AND mstdifferentiationtypeid = ? AND mstdifferentiationid = ? AND deleteflg = 0 and activeflg=1"

// var getAssetdifferentiation = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.trnassetid as Trnassetid, a.mstdifferentiationtypeid as Mstdifferentiationtypeid, a.mstdifferentiationid as Mstdifferentiationid, a.value as Value, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Recorddifferentiationtypename,e.name as Recorddifferentiationname,f.assetid as Assetid FROM mapassetdifferentiation a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,trnasset f WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.mstdifferentiationtypeid=d.id and a.mstdifferentiationid=e.id and a.trnassetid=f.id ORDER BY a.id DESC LIMIT ?,?"
// var getAssetdifferentiationcount = "SELECT count(a.id) as total FROM mapassetdifferentiation a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,trnasset f WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.mstdifferentiationtypeid=d.id and a.mstdifferentiationid=e.id and a.trnassetid=f.id"
var updateAssetdifferentiation = "UPDATE mapassetdifferentiation SET mstorgnhirarchyid = ?, trnassetid = ?, mstdifferentiationtypeid = ?, mstdifferentiationid = ?, value = ? WHERE id = ? "
var deleteAssetdifferentiation = "UPDATE mapassetdifferentiation SET deleteflg = '1' WHERE id = ? "

func (dbc DbConn) CheckDuplicateAssetdifferentiation(tz *entities.AssetdifferentiationEntity) (entities.AssetdifferentiationEntities, error) {
	logger.Log.Println("In side CheckDuplicateAssetdifferentiation")
	value := entities.AssetdifferentiationEntities{}
	err := dbc.DB.QueryRow(duplicateAssetdifferentiation, tz.Clientid, tz.Mstorgnhirarchyid, tz.Trnassetid, tz.Mstdifferentiationtypeid, tz.Mstdifferentiationid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateAssetdifferentiation Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertAssetdifferentiation(tz *entities.AssetdifferentiationEntity) (int64, error) {
	logger.Log.Println("In side InsertAssetdifferentiation")
	logger.Log.Println("Query -->", insertAssetdifferentiation)
	stmt, err := dbc.DB.Prepare(insertAssetdifferentiation)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertAssetdifferentiation Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Trnassetid, tz.Mstdifferentiationtypeid, tz.Mstdifferentiationid, tz.Value)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Trnassetid, tz.Mstdifferentiationtypeid, tz.Mstdifferentiationid, tz.Value)
	if err != nil {
		logger.Log.Println("InsertAssetdifferentiation Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllAssetdifferentiation(tz *entities.AssetdifferentiationEntity, OrgnType int64) ([]entities.AssetdifferentiationEntity, error) {
	logger.Log.Println("In side GelAllAssetdifferentiation")
	values := []entities.AssetdifferentiationEntity{}
	var getAssetdifferentiation string
	var params []interface{}
	if OrgnType == 1 {
		getAssetdifferentiation = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.trnassetid as Trnassetid, a.mstdifferentiationtypeid as Mstdifferentiationtypeid, a.mstdifferentiationid as Mstdifferentiationid, a.value as Value, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Recorddifferentiationtypename,e.name as Recorddifferentiationname,f.assetid as Assetid FROM mapassetdifferentiation a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,trnasset f WHERE a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.mstdifferentiationtypeid=d.id and a.mstdifferentiationid=e.id and a.trnassetid=f.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		getAssetdifferentiation = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.trnassetid as Trnassetid, a.mstdifferentiationtypeid as Mstdifferentiationtypeid, a.mstdifferentiationid as Mstdifferentiationid, a.value as Value, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Recorddifferentiationtypename,e.name as Recorddifferentiationname,f.assetid as Assetid FROM mapassetdifferentiation a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,trnasset f WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.mstdifferentiationtypeid=d.id and a.mstdifferentiationid=e.id and a.trnassetid=f.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		getAssetdifferentiation = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.trnassetid as Trnassetid, a.mstdifferentiationtypeid as Mstdifferentiationtypeid, a.mstdifferentiationid as Mstdifferentiationid, a.value as Value, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Recorddifferentiationtypename,e.name as Recorddifferentiationname,f.assetid as Assetid FROM mapassetdifferentiation a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,trnasset f WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.mstdifferentiationtypeid=d.id and a.mstdifferentiationid=e.id and a.trnassetid=f.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}

	rows, err := dbc.DB.Query(getAssetdifferentiation, params...)

	// rows, err := dbc.DB.Query(getAssetdifferentiation, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllAssetdifferentiation Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.AssetdifferentiationEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Trnassetid, &value.Mstdifferentiationtypeid, &value.Mstdifferentiationid, &value.Value, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Recorddifftypename, &value.Recorddiffname, &value.Assetid)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateAssetdifferentiation(tz *entities.AssetdifferentiationEntity) error {
	logger.Log.Println("In side UpdateAssetdifferentiation")
	stmt, err := dbc.DB.Prepare(updateAssetdifferentiation)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateAssetdifferentiation Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Trnassetid, tz.Mstdifferentiationtypeid, tz.Mstdifferentiationid, tz.Value, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateAssetdifferentiation Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteAssetdifferentiation(tz *entities.AssetdifferentiationEntity) error {
	logger.Log.Println("In side DeleteAssetdifferentiation")
	stmt, err := dbc.DB.Prepare(deleteAssetdifferentiation)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteAssetdifferentiation Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteAssetdifferentiation Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetAssetdifferentiationCount(tz *entities.AssetdifferentiationEntity, OrgnTypeID int64) (entities.AssetdifferentiationEntities, error) {
	logger.Log.Println("In side GetAssetdifferentiationCount")
	value := entities.AssetdifferentiationEntities{}
	var getAssetdifferentiationcount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getAssetdifferentiationcount = "SELECT count(a.id) as total FROM mapassetdifferentiation a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,trnasset f WHERE a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.mstdifferentiationtypeid=d.id and a.mstdifferentiationid=e.id and a.trnassetid=f.id"
	} else if OrgnTypeID == 2 {
		getAssetdifferentiationcount = "SELECT count(a.id) as total FROM mapassetdifferentiation a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,trnasset f WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.mstdifferentiationtypeid=d.id and a.mstdifferentiationid=e.id and a.trnassetid=f.id"
		params = append(params, tz.Clientid)
	} else {
		getAssetdifferentiationcount = "SELECT count(a.id) as total FROM mapassetdifferentiation a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,trnasset f WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.mstdifferentiationtypeid=d.id and a.mstdifferentiationid=e.id and a.trnassetid=f.id"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getAssetdifferentiationcount, params...).Scan(&value.Total)

	// err := dbc.DB.QueryRow(getAssetdifferentiationcount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetAssetdifferentiationCount Get Statement Prepare Error", err)
		return value, err
	}
}
