package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertAsset = "INSERT INTO trnasset (clientid, mstorgnhirarchyid, assetid,mstdifferentiationtypeid, additionalattr) VALUES (?,?,?,?,?)"
var duplicateAsset = "SELECT count(id) total FROM  trnasset WHERE clientid = ? AND mstorgnhirarchyid = ? AND mstdifferentiationtypeid=? AND assetid = ? AND deleteflg = 0"

// var getAsset = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.assetid as Assetid, a.additionalattr as Additionalattr, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,a.mstdifferentiationtypeid AS Mstdifferentiationtypeid,d.typename AS Mstdifferentiationtypename FROM trnasset a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstdifferentiationtypeid=d.id AND a.mstorgnhirarchyid = c.id ORDER BY a.id DESC LIMIT ?,?"
var getAssetcount = "SELECT count(a.id) total FROM trnasset a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.mstdifferentiationtypeid=d.id"
var updateAsset = "UPDATE trnasset SET mstorgnhirarchyid = ?, assetid = ? WHERE id = ? "
var deleteAsset = "UPDATE trnasset SET deleteflg = '1' WHERE id = ? "

// var getClietWiseAssetsql = "SELECT mstrecordtype.clientid, mstclient.name clientname, mstrecordtype.mstorgnhirarchyid, mstorgnhierarchy.name orgname, mstrecordtype.id, difftype.id typeid, difftype.typename, COALESCE(difftype.parentid, 0) parentid, mstrecorddifferentiation.id, mstrecorddifferentiation.name FROM mstclient, mstorgnhierarchy, mstrecordtype JOIN mstrecorddifferentiationtype ON mstrecorddifferentiationtype.activeflg = 1 AND mstrecorddifferentiationtype.deleteflg = 0 AND mstrecorddifferentiationtype.id = mstrecordtype.torecorddifftypeid AND mstrecorddifferentiationtype.seqno = 5 JOIN mstrecorddifferentiation ON mstrecordtype.fromrecorddiffid = mstrecorddifferentiation.id AND mstrecorddifferentiation.activeflg = 1 AND mstrecorddifferentiation.deleteflg = 0 AND mstrecordtype.clientid = mstrecorddifferentiation.clientid AND mstrecordtype.mstorgnhirarchyid = mstrecorddifferentiation.mstorgnhirarchyid AND mstrecordtype.fromrecorddifftypeid = mstrecorddifferentiation.recorddifftypeid JOIN mstrecorddifferentiationtype difftype ON difftype.activeflg = 1 AND difftype.deleteflg = 0 AND difftype.id = mstrecorddifferentiation.recorddifftypeid WHERE mstrecordtype.activeflg = 1 AND mstrecordtype.deleteflg = 0 AND mstclient.id = mstorgnhierarchy.clientid AND mstorgnhierarchy.clientid = mstrecordtype.clientid AND mstorgnhierarchy.id=mstrecordtype.mstorgnhirarchyid AND mstrecordtype.clientid = ? AND mstrecordtype.mstorgnhirarchyid = ? LIMIT ? , ?"
// var countClietWiseAssetsql = "SELECT COUNT(mstrecordtype.id) total FROM mstrecordtype JOIN  mstrecorddifferentiationtype ON mstrecorddifferentiationtype.activeflg=1 AND mstrecorddifferentiationtype.deleteflg=0 AND mstrecorddifferentiationtype.id=mstrecordtype.torecorddifftypeid AND mstrecorddifferentiationtype.seqno=5 JOIN mstrecorddifferentiation ON mstrecordtype.fromrecorddiffid=mstrecorddifferentiation.id AND mstrecorddifferentiation.activeflg=1 AND mstrecorddifferentiation.deleteflg=0 AND mstrecordtype.clientid=mstrecorddifferentiation.clientid AND mstrecordtype.mstorgnhirarchyid=mstrecorddifferentiation.mstorgnhirarchyid AND mstrecordtype.fromrecorddifftypeid=mstrecorddifferentiation.recorddifftypeid JOIN mstrecorddifferentiationtype difftype ON difftype.activeflg=1 AND difftype.deleteflg=0 AND difftype.id=mstrecorddifferentiation.recorddifftypeid WHERE mstrecordtype.activeflg=1 AND mstrecordtype.deleteflg=0 AND mstrecordtype.clientid=? AND mstrecordtype.mstorgnhirarchyid=?"

var getAssetByTypeSql = "SELECT trnasset.id AS Id, trnasset.assetid AS Assetid FROM trnasset WHERE trnasset.activeflg = 1 AND trnasset.deleteflg = 0 AND trnasset.clientid= ? AND trnasset.mstorgnhirarchyid= ? AND trnasset.mstdifferentiationtypeid= ? "

var getAssetDiffVal = "SELECT  mstrecorddifferentiationtype.id,mstrecorddifferentiationtype.typename,mstrecorddifferentiation.id dId,mstrecorddifferentiation.name,IF(mapassetdifferentiation.value IS NULL,'',mapassetdifferentiation.value) Value FROM mstrecorddifferentiationtype JOIN mstrecorddifferentiation ON mstrecorddifferentiationtype.id=mstrecorddifferentiation.recorddifftypeid AND mstrecorddifferentiationtype.clientid=mstrecorddifferentiation.clientid AND mstrecorddifferentiationtype.mstorgnhirarchyid=mstrecorddifferentiation.mstorgnhirarchyid  LEFT JOIN trnasset ON mstrecorddifferentiationtype.id = trnasset.mstdifferentiationtypeid AND trnasset.id=? AND trnasset.activeflg=1 AND trnasset.deleteflg=0 LEFT JOIN mapassetdifferentiation ON trnasset.id=mapassetdifferentiation.trnassetid AND mapassetdifferentiation.mstdifferentiationid=mstrecorddifferentiation.id AND mapassetdifferentiation.mstdifferentiationtypeid=mstrecorddifferentiationtype.id AND mapassetdifferentiation.activeflg=1 AND mapassetdifferentiation.deleteflg=0 WHERE mstrecorddifferentiationtype.activeflg =1 AND mstrecorddifferentiationtype.deleteflg=0 AND mstrecorddifferentiation.activeflg=1 AND mstrecorddifferentiation.deleteflg=0  AND mstrecorddifferentiation.clientid=? AND mstrecorddifferentiationtype.mstorgnhirarchyid=? AND mstrecorddifferentiationtype.id=? ORDER BY mstrecorddifferentiation.seqno ASC"

var delAssetDiffVal = "UPDATE mapassetdifferentiation SET deleteflg=1 WHERE clientid=? AND mstorgnhirarchyid=? AND mstdifferentiationtypeid=? AND trnassetid=? "
var insAssetDiffVal = "INSERT INTO mapassetdifferentiation(clientid, mstorgnhirarchyid, mstdifferentiationtypeid, mstdifferentiationid, trnassetid, value, deleteflg, activeflg, audittransactionid) VALUES (?,?,?,?,?,?,0,1,1) "

var getassettypessql = "SELECT mstrecorddifferentiationtype.id,mstrecorddifferentiationtype.typename,mstrecorddifferentiationtype.seqno,difftype.id FROM mstrecorddifferentiationtype JOIN mstrecorddifferentiationtype difftype ON mstrecorddifferentiationtype.parentid=difftype.id AND difftype.seqno=5 WHERE mstrecorddifferentiationtype.activeflg=1 AND mstrecorddifferentiationtype.deleteflg=0  AND mstrecorddifferentiationtype.clientid=? AND mstrecorddifferentiationtype.mstorgnhirarchyid=? ORDER BY mstrecorddifferentiationtype.seqno ASC"

var getassetattributessql = "SELECT mstrecorddifferentiation.id,mstrecorddifferentiation.name,mstrecorddifferentiation.seqno,mstrecorddifferentiation.recorddifftypeid FROM mstrecorddifferentiation JOIN mstrecorddifferentiationtype ON mstrecorddifferentiation.recorddifftypeid=mstrecorddifferentiationtype.id AND mstrecorddifferentiation.clientid=mstrecorddifferentiationtype.clientid AND mstrecorddifferentiation.mstorgnhirarchyid=mstrecorddifferentiationtype.mstorgnhirarchyid WHERE mstrecorddifferentiationtype.activeflg=1 AND mstrecorddifferentiationtype.deleteflg=0 AND mstrecorddifferentiation.activeflg=1 AND mstrecorddifferentiation.deleteflg=0 AND mstrecorddifferentiation.clientid=? AND mstrecorddifferentiation.mstorgnhirarchyid=? AND mstrecorddifferentiation.recorddifftypeid=? ORDER BY mstrecorddifferentiation.seqno "

var getassetsbytypesql = "SELECT trnasset.id,trnasset.assetid,coalesce(trnasset.history,'') as history FROM trnasset JOIN mapassetdifferentiation ON trnasset.id=mapassetdifferentiation.trnassetid AND trnasset.clientid=mapassetdifferentiation.clientid AND trnasset.mstorgnhirarchyid=mapassetdifferentiation.mstorgnhirarchyid AND mapassetdifferentiation.activeflg=1 AND mapassetdifferentiation.deleteflg=0 WHERE trnasset.activeflg=1 AND trnasset.deleteflg=0 AND trnasset.clientid=? AND trnasset.mstorgnhirarchyid=? AND mapassetdifferentiation.mstdifferentiationtypeid=? GROUP BY trnasset.id"
var getassetsbyvaluesql = "SELECT trnasset.id,trnasset.assetid,coalesce(trnasset.history,'') as history FROM trnasset JOIN mapassetdifferentiation ON trnasset.id=mapassetdifferentiation.trnassetid AND trnasset.clientid=mapassetdifferentiation.clientid AND trnasset.mstorgnhirarchyid=mapassetdifferentiation.mstorgnhirarchyid AND mapassetdifferentiation.activeflg=1 AND mapassetdifferentiation.deleteflg=0 WHERE trnasset.activeflg=1 AND trnasset.deleteflg=0 AND trnasset.clientid=? AND trnasset.mstorgnhirarchyid=? AND mapassetdifferentiation.mstdifferentiationtypeid=? AND mapassetdifferentiation.mstdifferentiationid=? AND mapassetdifferentiation.value LIKE ? GROUP BY trnasset.id"
var getassetsbyidsql = "SELECT trnasset.id,trnasset.assetid ,coalesce(trnasset.history,'') as history FROM trnasset JOIN mapassetdifferentiation ON trnasset.id=mapassetdifferentiation.trnassetid AND trnasset.clientid=mapassetdifferentiation.clientid AND trnasset.mstorgnhirarchyid=mapassetdifferentiation.mstorgnhirarchyid AND mapassetdifferentiation.activeflg=1 AND mapassetdifferentiation.deleteflg=0 WHERE trnasset.activeflg=1 AND trnasset.deleteflg=0 AND trnasset.clientid=? AND trnasset.mstorgnhirarchyid=? AND mapassetdifferentiation.mstdifferentiationtypeid=? AND  trnasset.assetid LIKE ? GROUP BY trnasset.id"

func (dbc DbConn) CheckDuplicateAsset(tz *entities.AssetEntity) (entities.AssetEntities, error) {
	logger.Log.Println("In side CheckDuplicateAsset")
	value := entities.AssetEntities{}
	err := dbc.DB.QueryRow(duplicateAsset, tz.Clientid, tz.Mstorgnhirarchyid, tz.Mstdifferentiationtypeid, tz.Assetid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateAsset Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertAsset(tz *entities.AssetEntity) (int64, error) {
	logger.Log.Println("In side InsertAsset")
	logger.Log.Println("Query -->", insertAsset)
	stmt, err := dbc.DB.Prepare(insertAsset)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertAsset Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Assetid, tz.Mstdifferentiationtypeid, tz.Additionalattr)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Assetid, tz.Mstdifferentiationtypeid, tz.Additionalattr)
	if err != nil {
		logger.Log.Println("InsertAsset Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllAsset(tz *entities.AssetEntity, OrgnType int64) ([]entities.AssetEntity, error) {
	// logger.Log.Println("Parameter -->", page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	values := []entities.AssetEntity{}
	var getAsset string
	var params []interface{}
	if OrgnType == 1 {
		getAsset = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.assetid as Assetid, a.additionalattr as Additionalattr, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,a.mstdifferentiationtypeid AS Mstdifferentiationtypeid,d.typename AS Mstdifferentiationtypename FROM trnasset a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d WHERE a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstdifferentiationtypeid=d.id AND a.mstorgnhirarchyid = c.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		getAsset = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.assetid as Assetid, a.additionalattr as Additionalattr, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,a.mstdifferentiationtypeid AS Mstdifferentiationtypeid,d.typename AS Mstdifferentiationtypename FROM trnasset a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstdifferentiationtypeid=d.id AND a.mstorgnhirarchyid = c.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		getAsset = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.assetid as Assetid, a.additionalattr as Additionalattr, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,a.mstdifferentiationtypeid AS Mstdifferentiationtypeid,d.typename AS Mstdifferentiationtypename FROM trnasset a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstdifferentiationtypeid=d.id AND a.mstorgnhirarchyid = c.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}

	logger.Log.Println("In side GelAllAsset==>", getAsset)
	rows, err := dbc.DB.Query(getAsset, params...)

	// rows, err := dbc.DB.Query(getAsset, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllAsset Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.AssetEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Assetid, &value.Additionalattr, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Mstdifferentiationtypeid, &value.Mstdifferentiationtypename)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateAsset(tz *entities.AssetEntity) error {
	logger.Log.Println("In side UpdateAsset")
	stmt, err := dbc.DB.Prepare(updateAsset)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateAsset Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Assetid, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateAsset Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteAsset(tz *entities.AssetEntity) error {
	logger.Log.Println("In side DeleteAsset")
	stmt, err := dbc.DB.Prepare(deleteAsset)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteAsset Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteAsset Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetAssetCount(tz *entities.AssetEntity, OrgnTypeID int64) (entities.AssetEntities, error) {
	logger.Log.Println("In side GetAssetCount")
	value := entities.AssetEntities{}
	var getAssetcount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getAssetcount = "SELECT count(a.id) total FROM trnasset a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d WHERE a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.mstdifferentiationtypeid=d.id"
	} else if OrgnTypeID == 2 {
		getAssetcount = "SELECT count(a.id) total FROM trnasset a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.mstdifferentiationtypeid=d.id"
		params = append(params, tz.Clientid)
	} else {
		getAssetcount = "SELECT count(a.id) total FROM trnasset a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.mstdifferentiationtypeid=d.id"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getAssetcount, params...).Scan(&value.Total)

	// err := dbc.DB.QueryRow(getAssetcount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetAssetCount Get Statement Prepareentities.AssetMapWithRecordType Error", err)
		return value, err
	}
}

func (dbc DbConn) GetAssetBYType(page *entities.AssetEntity) ([]entities.AssetEntityByType, error) {
	logger.Log.Println("In side GetAssetBYType==>", getAssetByTypeSql)
	logger.Log.Println("Parameter -->", page.Clientid, page.Mstorgnhirarchyid, page.Mstdifferentiationtypeid)
	values := []entities.AssetEntityByType{}
	rows, err := dbc.DB.Query(getAssetByTypeSql, page.Clientid, page.Mstorgnhirarchyid, page.Mstdifferentiationtypeid)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllAsset Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.AssetEntityByType{}
		rows.Scan(&value.Id, &value.Assetid)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) GetAssetDiffVal(page *entities.AssetEntity) ([]entities.AssetEntityDiffVal, error) {
	logger.Log.Println("In side GetAssetDiffVal==>", getAssetDiffVal)
	logger.Log.Println("Parameter -->", page.Id, page.Clientid, page.Mstorgnhirarchyid, page.Mstdifferentiationtypeid)
	values := []entities.AssetEntityDiffVal{}
	rows, err := dbc.DB.Query(getAssetDiffVal, page.Id, page.Clientid, page.Mstorgnhirarchyid, page.Mstdifferentiationtypeid)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllAsset Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.AssetEntityDiffVal{}
		rows.Scan(&value.TypeId, &value.TypeName, &value.AttrId, &value.AttrName, &value.Value)
		values = append(values, value)
	}
	return values, nil
}
func (dbc DbConn) GetAssetDiffValByID(page *entities.AssetSearchEntity, id int64) ([]entities.AssetEntityDiffVal, error) {
	logger.Log.Println("In side GetAssetDiffVal==>", getAssetDiffVal)
	logger.Log.Println("Parameter -->", id, page.Clientid, page.Mstorgnhirarchyid, page.Mstdifferentiationtypeid)
	values := []entities.AssetEntityDiffVal{}
	rows, err := dbc.DB.Query(getAssetDiffVal, id, page.Clientid, page.Mstorgnhirarchyid, page.Mstdifferentiationtypeid)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllAsset Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.AssetEntityDiffVal{}
		rows.Scan(&value.TypeId, &value.TypeName, &value.AttrId, &value.AttrName, &value.Value)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) DelAssetDiff(tz *entities.AssetEntityDiffValUpdate) error {
	logger.Log.Println("In side DelAssetDiff")
	stmt, err := dbc.DB.Prepare(delAssetDiffVal)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DelAssetDiff Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Mstdifferentiationtypeid, tz.Assetid)
	if err != nil {
		logger.Log.Println("DelAssetDiff Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) InsertAssetDiff(tz *entities.AssetEntityDiffValUpdate) error {
	logger.Log.Println("In side DelAssetDiff")
	stmt, err := dbc.DB.Prepare(insAssetDiffVal)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DelAssetDiff Prepare Statement  Error", err)
		return err
	}
	for _, s := range tz.Attributes {
		_, err = stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, s.TypeId, s.AttrId, tz.Assetid, s.Value)
		if err != nil {
			logger.Log.Println("DelAssetDiff Execute Statement  Error", err)
			return err
		}
	}

	return nil
}

// GetClietWiseAsset function is used to get asset mapping with Ticket Type wise
func (dbc DbConn) GetClietWiseAsset(page *entities.AssetEntity, OrgnType int64) ([]entities.AssetMapWithRecordType, error) {
	logger.Log.Println("Parameter -->", page.Clientid, page.Mstorgnhirarchyid)
	values := []entities.AssetMapWithRecordType{}
	var getClietWiseAssetsql string
	var params []interface{}
	if OrgnType == 1 {
		getClietWiseAssetsql = "SELECT mstrecordtype.clientid, mstclient.name clientname, mstrecordtype.mstorgnhirarchyid, mstorgnhierarchy.name orgname, mstrecordtype.id, difftype.id typeid, difftype.typename, COALESCE(difftype.parentid, 0) parentid, mstrecorddifferentiation.id, mstrecorddifferentiation.name FROM mstclient, mstorgnhierarchy, mstrecordtype JOIN mstrecorddifferentiationtype ON mstrecorddifferentiationtype.activeflg = 1 AND mstrecorddifferentiationtype.deleteflg = 0 AND mstrecorddifferentiationtype.id = mstrecordtype.torecorddifftypeid AND mstrecorddifferentiationtype.seqno = 5 JOIN mstrecorddifferentiation ON mstrecordtype.fromrecorddiffid = mstrecorddifferentiation.id AND mstrecorddifferentiation.activeflg = 1 AND mstrecorddifferentiation.deleteflg = 0 AND mstrecordtype.clientid = mstrecorddifferentiation.clientid AND mstrecordtype.mstorgnhirarchyid = mstrecorddifferentiation.mstorgnhirarchyid AND mstrecordtype.fromrecorddifftypeid = mstrecorddifferentiation.recorddifftypeid JOIN mstrecorddifferentiationtype difftype ON difftype.activeflg = 1 AND difftype.deleteflg = 0 AND difftype.id = mstrecorddifferentiation.recorddifftypeid WHERE mstrecordtype.activeflg = 1 AND mstrecordtype.deleteflg = 0 AND mstclient.id = mstorgnhierarchy.clientid AND mstorgnhierarchy.clientid = mstrecordtype.clientid AND mstorgnhierarchy.id=mstrecordtype.mstorgnhirarchyid LIMIT ? , ?"
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else if OrgnType == 2 {
		getClietWiseAssetsql = "SELECT mstrecordtype.clientid, mstclient.name clientname, mstrecordtype.mstorgnhirarchyid, mstorgnhierarchy.name orgname, mstrecordtype.id, difftype.id typeid, difftype.typename, COALESCE(difftype.parentid, 0) parentid, mstrecorddifferentiation.id, mstrecorddifferentiation.name FROM mstclient, mstorgnhierarchy, mstrecordtype JOIN mstrecorddifferentiationtype ON mstrecorddifferentiationtype.activeflg = 1 AND mstrecorddifferentiationtype.deleteflg = 0 AND mstrecorddifferentiationtype.id = mstrecordtype.torecorddifftypeid AND mstrecorddifferentiationtype.seqno = 5 JOIN mstrecorddifferentiation ON mstrecordtype.fromrecorddiffid = mstrecorddifferentiation.id AND mstrecorddifferentiation.activeflg = 1 AND mstrecorddifferentiation.deleteflg = 0 AND mstrecordtype.clientid = mstrecorddifferentiation.clientid AND mstrecordtype.mstorgnhirarchyid = mstrecorddifferentiation.mstorgnhirarchyid AND mstrecordtype.fromrecorddifftypeid = mstrecorddifferentiation.recorddifftypeid JOIN mstrecorddifferentiationtype difftype ON difftype.activeflg = 1 AND difftype.deleteflg = 0 AND difftype.id = mstrecorddifferentiation.recorddifftypeid WHERE mstrecordtype.activeflg = 1 AND mstrecordtype.deleteflg = 0 AND mstclient.id = mstorgnhierarchy.clientid AND mstorgnhierarchy.clientid = mstrecordtype.clientid AND mstorgnhierarchy.id=mstrecordtype.mstorgnhirarchyid AND mstrecordtype.clientid = ? LIMIT ? , ?"
		params = append(params, page.Clientid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else {
		getClietWiseAssetsql = "SELECT mstrecordtype.clientid, mstclient.name clientname, mstrecordtype.mstorgnhirarchyid, mstorgnhierarchy.name orgname, mstrecordtype.id, difftype.id typeid, difftype.typename, COALESCE(difftype.parentid, 0) parentid, mstrecorddifferentiation.id, mstrecorddifferentiation.name FROM mstclient, mstorgnhierarchy, mstrecordtype JOIN mstrecorddifferentiationtype ON mstrecorddifferentiationtype.activeflg = 1 AND mstrecorddifferentiationtype.deleteflg = 0 AND mstrecorddifferentiationtype.id = mstrecordtype.torecorddifftypeid AND mstrecorddifferentiationtype.seqno = 5 JOIN mstrecorddifferentiation ON mstrecordtype.fromrecorddiffid = mstrecorddifferentiation.id AND mstrecorddifferentiation.activeflg = 1 AND mstrecorddifferentiation.deleteflg = 0 AND mstrecordtype.clientid = mstrecorddifferentiation.clientid AND mstrecordtype.mstorgnhirarchyid = mstrecorddifferentiation.mstorgnhirarchyid AND mstrecordtype.fromrecorddifftypeid = mstrecorddifferentiation.recorddifftypeid JOIN mstrecorddifferentiationtype difftype ON difftype.activeflg = 1 AND difftype.deleteflg = 0 AND difftype.id = mstrecorddifferentiation.recorddifftypeid WHERE mstrecordtype.activeflg = 1 AND mstrecordtype.deleteflg = 0 AND mstclient.id = mstorgnhierarchy.clientid AND mstorgnhierarchy.clientid = mstrecordtype.clientid AND mstorgnhierarchy.id=mstrecordtype.mstorgnhirarchyid AND mstrecordtype.clientid = ? AND mstrecordtype.mstorgnhirarchyid = ? LIMIT ? , ?"
		params = append(params, page.Clientid)
		params = append(params, page.Mstorgnhirarchyid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	}
	logger.Log.Println("In side GetClietWiseAsset==>", getClietWiseAssetsql)

	rows, err := dbc.DB.Query(getClietWiseAssetsql, params...)

	// rows, err := dbc.DB.Query(getClietWiseAssetsql, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetClietWiseAsset Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.AssetMapWithRecordType{}
		rows.Scan(&value.Clientid, &value.Clientname, &value.Mstorgnhirarchyid, &value.Mstorgnhirarchyname, &value.ID, &value.DiffTypeID, &value.DiffTypeName, &value.DiffTypeParent, &value.DiffID, &value.DiffName)
		values = append(values, value)
	}
	return values, nil
}

// CountClietWiseAsset function is used to get asset mapping with Ticket Type wise
func (dbc DbConn) CountClietWiseAsset(page *entities.AssetEntity, OrgnTypeID int64) (entities.AssetEntities, error) {
	logger.Log.Println("In side CountClietWiseAsset")
	value := entities.AssetEntities{}
	var countClietWiseAssetsql string
	var params []interface{}
	if OrgnTypeID == 1 {
		countClietWiseAssetsql = "SELECT COUNT(mstrecordtype.id) total FROM mstrecordtype JOIN  mstrecorddifferentiationtype ON mstrecorddifferentiationtype.activeflg=1 AND mstrecorddifferentiationtype.deleteflg=0 AND mstrecorddifferentiationtype.id=mstrecordtype.torecorddifftypeid AND mstrecorddifferentiationtype.seqno=5 JOIN mstrecorddifferentiation ON mstrecordtype.fromrecorddiffid=mstrecorddifferentiation.id AND mstrecorddifferentiation.activeflg=1 AND mstrecorddifferentiation.deleteflg=0 AND mstrecordtype.clientid=mstrecorddifferentiation.clientid AND mstrecordtype.mstorgnhirarchyid=mstrecorddifferentiation.mstorgnhirarchyid AND mstrecordtype.fromrecorddifftypeid=mstrecorddifferentiation.recorddifftypeid JOIN mstrecorddifferentiationtype difftype ON difftype.activeflg=1 AND difftype.deleteflg=0 AND difftype.id=mstrecorddifferentiation.recorddifftypeid WHERE mstrecordtype.activeflg=1 AND mstrecordtype.deleteflg=0"
	} else if OrgnTypeID == 2 {
		countClietWiseAssetsql = "SELECT COUNT(mstrecordtype.id) total FROM mstrecordtype JOIN  mstrecorddifferentiationtype ON mstrecorddifferentiationtype.activeflg=1 AND mstrecorddifferentiationtype.deleteflg=0 AND mstrecorddifferentiationtype.id=mstrecordtype.torecorddifftypeid AND mstrecorddifferentiationtype.seqno=5 JOIN mstrecorddifferentiation ON mstrecordtype.fromrecorddiffid=mstrecorddifferentiation.id AND mstrecorddifferentiation.activeflg=1 AND mstrecorddifferentiation.deleteflg=0 AND mstrecordtype.clientid=mstrecorddifferentiation.clientid AND mstrecordtype.mstorgnhirarchyid=mstrecorddifferentiation.mstorgnhirarchyid AND mstrecordtype.fromrecorddifftypeid=mstrecorddifferentiation.recorddifftypeid JOIN mstrecorddifferentiationtype difftype ON difftype.activeflg=1 AND difftype.deleteflg=0 AND difftype.id=mstrecorddifferentiation.recorddifftypeid WHERE mstrecordtype.activeflg=1 AND mstrecordtype.deleteflg=0 AND mstrecordtype.clientid=? "
		params = append(params, page.Clientid)
	} else {
		countClietWiseAssetsql = "SELECT COUNT(mstrecordtype.id) total FROM mstrecordtype JOIN  mstrecorddifferentiationtype ON mstrecorddifferentiationtype.activeflg=1 AND mstrecorddifferentiationtype.deleteflg=0 AND mstrecorddifferentiationtype.id=mstrecordtype.torecorddifftypeid AND mstrecorddifferentiationtype.seqno=5 JOIN mstrecorddifferentiation ON mstrecordtype.fromrecorddiffid=mstrecorddifferentiation.id AND mstrecorddifferentiation.activeflg=1 AND mstrecorddifferentiation.deleteflg=0 AND mstrecordtype.clientid=mstrecorddifferentiation.clientid AND mstrecordtype.mstorgnhirarchyid=mstrecorddifferentiation.mstorgnhirarchyid AND mstrecordtype.fromrecorddifftypeid=mstrecorddifferentiation.recorddifftypeid JOIN mstrecorddifferentiationtype difftype ON difftype.activeflg=1 AND difftype.deleteflg=0 AND difftype.id=mstrecorddifferentiation.recorddifftypeid WHERE mstrecordtype.activeflg=1 AND mstrecordtype.deleteflg=0 AND mstrecordtype.clientid=? AND mstrecordtype.mstorgnhirarchyid=?"
		params = append(params, page.Clientid)
		params = append(params, page.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(countClietWiseAssetsql, params...).Scan(&value.Total)

	// err := dbc.DB.QueryRow(countClietWiseAssetsql, page.Clientid, page.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("countClietWiseAssetsql Get Statement Prepare entities.AssetEntities Error", err)
		return value, err
	}
}

// GetAssetTypes function is used to get asset mapping with Ticket Type wise
func (dbc DbConn) GetAssetTypes(page *entities.AssetEntity) ([]entities.Assettype, error) {
	logger.Log.Println("In side GetAssetTypes==>", getassettypessql)
	logger.Log.Println("Parameter -->", page.Clientid, page.Mstorgnhirarchyid)
	values := []entities.Assettype{}
	rows, err := dbc.DB.Query(getassettypessql, page.Clientid, page.Mstorgnhirarchyid)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAssetTypes Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.Assettype{}
		rows.Scan(&value.ID, &value.Name, &value.Seqno, &value.Parent)
		values = append(values, value)
	}
	return values, nil
}

// GetAssetAttributes function is used to get asset mapping with Ticket Type wise
func (dbc DbConn) GetAssetAttributes(page *entities.AssetEntity) ([]entities.Assettype, error) {
	logger.Log.Println("In side GetAssetAttributes==>", getassetattributessql)
	logger.Log.Println("Parameter -->", page.Clientid, page.Mstorgnhirarchyid, page.Mstdifferentiationtypeid)
	values := []entities.Assettype{}
	rows, err := dbc.DB.Query(getassetattributessql, page.Clientid, page.Mstorgnhirarchyid, page.Mstdifferentiationtypeid)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAssetAttributes Get Statement Prepare Error", err)
		return values, err
	}
	var value = entities.Assettype{ID: 0, Name: "Asset ID", Parent: 0, Seqno: 0}
	values = append(values, value)
	for rows.Next() {
		value := entities.Assettype{}
		rows.Scan(&value.ID, &value.Name, &value.Seqno, &value.Parent)
		values = append(values, value)
	}
	return values, nil
}

// GetAssetAttributes function is used to get asset mapping with Ticket Type wise
func (dbc DbConn) GetAssetAttributesbyTypeId(page *entities.AssetSearchEntity) ([]entities.Assettype, error) {
	logger.Log.Println("In side GetAssetAttributes==>", getassetattributessql)
	logger.Log.Println("Parameter -->", page.Clientid, page.Mstorgnhirarchyid, page.Mstdifferentiationtypeid)
	values := []entities.Assettype{}
	rows, err := dbc.DB.Query(getassetattributessql, page.Clientid, page.Mstorgnhirarchyid, page.Mstdifferentiationtypeid)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAssetAttributes Get Statement Prepare Error", err)
		return values, err
	}
	var value = entities.Assettype{ID: 0, Name: "Asset ID", Parent: 0, Seqno: 0}
	values = append(values, value)
	for rows.Next() {
		value := entities.Assettype{}
		rows.Scan(&value.ID, &value.Name, &value.Seqno, &value.Parent)
		values = append(values, value)
	}
	return values, nil
}

// GetAssetByTypeNAtrrValue function is used to get assets with attributes and its value
func (dbc DbConn) GetAssetByTypeNAtrrValue(page *entities.AssetSearchEntity) ([]entities.AssetIDEntity, error) {
	values := []entities.AssetIDEntity{}
	if page.Mstdifferentiationid == 0 && page.Value != "" {
		logger.Log.Println("In side GetAssetByTypeNAtrrValue==>", getassetsbyidsql)
		logger.Log.Println("Parameter -->", page.Clientid, page.Mstorgnhirarchyid, page.Mstdifferentiationtypeid, page.Value)
		rows, err := dbc.DB.Query(getassetsbyidsql, page.Clientid, page.Mstorgnhirarchyid, page.Mstdifferentiationtypeid, page.Value)
		defer rows.Close()
		if err != nil {
			logger.Log.Println("GetAssetAttributes Get Statement Prepare Error", err)
			return values, err
		}
		for rows.Next() {
			value := entities.AssetIDEntity{}
			rows.Scan(&value.ID, &value.Assetid, &value.History)
			values = append(values, value)
		}
		return values, nil
	} else if page.Mstdifferentiationid > 0 && page.Value != "" {
		logger.Log.Println("In side GetAssetByTypeNAtrrValue==>", getassetsbyvaluesql)
		logger.Log.Println("Parameter -->", page.Clientid, page.Mstorgnhirarchyid, page.Mstdifferentiationtypeid, page.Mstdifferentiationid, page.Value)
		rows, err := dbc.DB.Query(getassetsbyvaluesql, page.Clientid, page.Mstorgnhirarchyid, page.Mstdifferentiationtypeid, page.Mstdifferentiationid, page.Value)
		defer rows.Close()
		if err != nil {
			logger.Log.Println("GetAssetAttributes Get Statement Prepare Error", err)
			return values, err
		}
		for rows.Next() {
			value := entities.AssetIDEntity{}
			rows.Scan(&value.ID, &value.Assetid, &value.History)
			values = append(values, value)
		}
		return values, nil
	} else {
		logger.Log.Println("In side GetAssetByTypeNAtrrValue==>", getassetsbytypesql)
		logger.Log.Println("Parameter -->", page.Clientid, page.Mstorgnhirarchyid, page.Mstdifferentiationtypeid)
		rows, err := dbc.DB.Query(getassetsbytypesql, page.Clientid, page.Mstorgnhirarchyid, page.Mstdifferentiationtypeid)
		defer rows.Close()
		if err != nil {
			logger.Log.Println("GetAssetAttributes Get Statement Prepare Error", err)
			return values, err
		}
		for rows.Next() {
			value := entities.AssetIDEntity{}
			rows.Scan(&value.ID, &value.Assetid, &value.History)
			values = append(values, value)
		}
		return values, nil
	}

}
