//SearchUser search a specific user using loginname ,roleid ,clientid,orgid

package dao

import (
	"database/sql"
	"log"
	"src/logger"

	"src/entities"
)

var insertLocation = "INSERT INTO mstlocationwisedifferentiationmap (clientid, mstorgnhirarchyid, location, fromrecorddifftypeid, fromrecorddiffid, torecorddifftypeid, torecorddiffid) VALUES (?,?,?,?,?,?,?)"
var checkduplicateLocation = "SELECT count(id) total FROM  mstlocationwisedifferentiationmap WHERE clientid = ? AND mstorgnhirarchyid = ? AND location=? AND fromrecorddifftypeid = ? AND fromrecorddiffid=? AND deleteflg = 0"
var checkduplicateLocationforupdate = "SELECT count(id) total FROM  mstlocationwisedifferentiationmap WHERE clientid = ? AND mstorgnhirarchyid = ? AND location=? AND fromrecorddifftypeid = ? AND fromrecorddiffid=? AND id<>? AND deleteflg = 0"

var locationsearch = "select id,location from mstlocationwisedifferentiationmap where clientid=? and mstorgnhirarchyid=? and fromrecorddiffid=? and deleteflg=0 and location like ?"
var locationselect = "select a.id as id,a.location as location,b.name as priority,b.id as priorityid,a.torecorddifftypeid,c.typename from mstlocationwisedifferentiationmap a,mstrecorddifferentiation b,mstrecorddifferentiationtype c where a.clientid=? and a.mstorgnhirarchyid=? and a.clientid=b.clientid and a.mstorgnhirarchyid=b.mstorgnhirarchyid and a.fromrecorddiffid=? and a.torecorddiffid=b.id and a.id=? and a.torecorddifftypeid=c.id and a.deleteflg=0"
var updateLocation = "UPDATE mstlocationwisedifferentiationmap SET clientid=?,mstorgnhirarchyid = ?, location = ?, fromrecorddifftypeid = ?, fromrecorddiffid = ?, torecorddifftypeid = ?, torecorddiffid = ? WHERE id = ? "
var deleteLocation = "UPDATE mstlocationwisedifferentiationmap SET deleteflg = '1' WHERE id = ? "

func SearchLocation(tz *entities.LocationPriorityEntity, db *sql.DB) ([]entities.LocationSearchEntity, error) {
	logger.Log.Println("In side dao", tz)
	values := []entities.LocationSearchEntity{}
	rows, err := db.Query(locationsearch, tz.ClientID, tz.MstorgnhirarchyID, tz.Recorddiffid, "%"+tz.Location+"%")
	defer rows.Close()
	if err != nil {
		log.Print("SearchUser Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.LocationSearchEntity{}
		rows.Scan(&value.ID, &value.Location)
		values = append(values, value)
	}
	return values, nil
}
func SelectLocation(tz *entities.LocationPriorityEntity, db *sql.DB) ([]entities.LocationSelectEntity, error) {
	logger.Log.Println("In side dao", tz)
	values := []entities.LocationSelectEntity{}
	rows, err := db.Query(locationselect, tz.ClientID, tz.MstorgnhirarchyID, tz.Recorddiffid, tz.ID)
	defer rows.Close()
	if err != nil {
		log.Print("SearchUser Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.LocationSelectEntity{}
		rows.Scan(&value.ID, &value.Location, &value.Priority, &value.Priorityid, &value.Recorddifftypeid, &value.RecorddifftypeName)
		values = append(values, value)
	}
	return values, nil
}
func CheckDuplicateLocation(tz *entities.LocationPriorityEntity, db *sql.DB) (entities.LocationEntities, error) {
	logger.Log.Println("In side CheckDuplicateAsset")
	value := entities.LocationEntities{}
	err := db.QueryRow(checkduplicateLocation, tz.ClientID, tz.MstorgnhirarchyID, tz.Location, tz.Recorddifftypeid, tz.Recorddiffid).Scan(&value.Total)
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
func AddLocation(tz *entities.LocationPriorityEntity, db *sql.DB) (int64, error) {
	logger.Log.Println("In side InsertAsset")
	logger.Log.Println("Query -->", insertLocation)
	stmt, err := db.Prepare(insertLocation)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertAsset Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.ClientID, tz.MstorgnhirarchyID, tz.Location, tz.Recorddifftypeid, tz.Recorddiffid, tz.ToRecorddifftypeid, tz.ToRecorddiffid)
	res, err := stmt.Exec(tz.ClientID, tz.MstorgnhirarchyID, tz.Location, tz.Recorddifftypeid, tz.Recorddiffid, tz.ToRecorddifftypeid, tz.ToRecorddiffid)
	if err != nil {
		logger.Log.Println("InsertAsset Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func GetAllLocation(tz *entities.LocationPriorityEntity, OrgnType int64, db *sql.DB) ([]entities.LocationPriorityEntity, error) {
	// logger.Log.Println("Parameter -->", page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	values := []entities.LocationPriorityEntity{}
	var getAsset string
	var params []interface{}
	if OrgnType == 1 {
		getAsset = "select a.id as id,a.clientid as clientid,a.mstorgnhirarchyid as mstorgnhirarchyid,a.location as location,a.fromrecorddifftypeid as fromrecorddifftypeid,a.fromrecorddiffid as fromrecorddiffid,a.torecorddifftypeid as torecorddifftypeid,a.torecorddiffid as torecorddiffid,b.name as recorddiffname,c.name as torecorddiffname,d.name as Clientname,e.name as Mstorgnhirarchyname,f.typename,g.typename  from mstlocationwisedifferentiationmap a,mstrecorddifferentiation b,mstrecorddifferentiation c,mstclient d,mstorgnhierarchy e,mstrecorddifferentiationtype f,mstrecorddifferentiationtype g  where  a.fromrecorddiffid=b.id and a.torecorddiffid=c.id and a.clientid=b.clientid and a.mstorgnhirarchyid=b.mstorgnhirarchyid and a.clientid=c.clientid and a.mstorgnhirarchyid=c.mstorgnhirarchyid and a.clientid=d.id and a.mstorgnhirarchyid=e.id and a.activeflg=1 and a.deleteflg=0 and a.fromrecorddifftypeid=f.id and a.torecorddifftypeid=g.id  ORDER BY a.id DESC LIMIT ?,?;"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		getAsset = "select a.id as id,a.clientid as clientid,a.mstorgnhirarchyid as mstorgnhirarchyid,a.location as location,a.fromrecorddifftypeid as fromrecorddifftypeid,a.fromrecorddiffid as fromrecorddiffid,a.torecorddifftypeid as torecorddifftypeid,a.torecorddiffid as torecorddiffid,b.name as recorddiffname,c.name as torecorddiffname,d.name as Clientname,e.name as Mstorgnhirarchyname,f.typename,g.typename  from mstlocationwisedifferentiationmap a,mstrecorddifferentiation b,mstrecorddifferentiation c,mstclient d,mstorgnhierarchy e,mstrecorddifferentiationtype f,mstrecorddifferentiationtype g  where a.clientid=? and a.fromrecorddiffid=b.id and a.torecorddiffid=c.id and a.clientid=b.clientid and a.mstorgnhirarchyid=b.mstorgnhirarchyid and a.clientid=c.clientid and a.mstorgnhirarchyid=c.mstorgnhirarchyid and a.clientid=d.id and a.mstorgnhirarchyid=e.id and a.activeflg=1 and a.deleteflg=0 and a.fromrecorddifftypeid=f.id and a.torecorddifftypeid=g.id  ORDER BY a.id DESC LIMIT ?,?;"
		params = append(params, tz.ClientID)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		getAsset = "select a.id as id,a.clientid as clientid,a.mstorgnhirarchyid as mstorgnhirarchyid,a.location as location,a.fromrecorddifftypeid as fromrecorddifftypeid,a.fromrecorddiffid as fromrecorddiffid,a.torecorddifftypeid as torecorddifftypeid,a.torecorddiffid as torecorddiffid,b.name as recorddiffname,c.name as torecorddiffname,d.name as Clientname,e.name as Mstorgnhirarchyname,f.typename,g.typename  from mstlocationwisedifferentiationmap a,mstrecorddifferentiation b,mstrecorddifferentiation c,mstclient d,mstorgnhierarchy e,mstrecorddifferentiationtype f,mstrecorddifferentiationtype g  where a.clientid=? and a.mstorgnhirarchyid=? and a.fromrecorddiffid=b.id and a.torecorddiffid=c.id and a.clientid=b.clientid and a.mstorgnhirarchyid=b.mstorgnhirarchyid and a.clientid=c.clientid and a.mstorgnhirarchyid=c.mstorgnhirarchyid and a.clientid=d.id and a.mstorgnhirarchyid=e.id and a.activeflg=1 and a.deleteflg=0 and a.fromrecorddifftypeid=f.id and a.torecorddifftypeid=g.id  ORDER BY a.id DESC LIMIT ?,?;"
		params = append(params, tz.ClientID)
		params = append(params, tz.MstorgnhirarchyID)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}

	logger.Log.Println("In side GelAllAsset==>", getAsset)
	rows, err := db.Query(getAsset, params...)

	// rows, err := dbc.DB.Query(getAsset, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllAsset Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.LocationPriorityEntity{}
		rows.Scan(&value.ID, &value.ClientID, &value.MstorgnhirarchyID, &value.Location, &value.Recorddifftypeid, &value.Recorddiffid, &value.ToRecorddifftypeid, &value.ToRecorddiffid, &value.ReccorddiffName, &value.ToReccorddiffName, &value.ClientName, &value.MstorgnhirarchyName, &value.RecorddifftypeName, &value.ToRecorddifftypeName)
		values = append(values, value)
	}
	return values, nil
}
func GetLocationCount(tz *entities.LocationPriorityEntity, OrgnTypeID int64, db *sql.DB) (entities.LocationEntities, error) {
	logger.Log.Println("In side GetAssetCount")
	value := entities.LocationEntities{}
	var getAssetcount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getAssetcount = "SELECT count(a.id) total FROM mstlocationwisedifferentiationmap a,mstrecorddifferentiation b,mstrecorddifferentiation c,mstclient d,mstorgnhierarchy e,mstrecorddifferentiationtype f,mstrecorddifferentiationtype g  where a.fromrecorddiffid=b.id and a.torecorddiffid=c.id and a.clientid=b.clientid and a.mstorgnhirarchyid=b.mstorgnhirarchyid and a.clientid=c.clientid and a.mstorgnhirarchyid=c.mstorgnhirarchyid and a.clientid=d.id and a.mstorgnhirarchyid=e.id and a.activeflg=1 and a.deleteflg=0 and a.fromrecorddifftypeid=f.id and a.torecorddifftypeid=g.id "
	} else if OrgnTypeID == 2 {
		getAssetcount = "SELECT count(a.id) total FROM  mstlocationwisedifferentiationmap a,mstrecorddifferentiation b,mstrecorddifferentiation c,mstclient d,mstorgnhierarchy e,mstrecorddifferentiationtype f,mstrecorddifferentiationtype g  where a.clientid=? and a.fromrecorddiffid=b.id and a.torecorddiffid=c.id and a.clientid=b.clientid and a.mstorgnhirarchyid=b.mstorgnhirarchyid and a.clientid=c.clientid and a.mstorgnhirarchyid=c.mstorgnhirarchyid and a.clientid=d.id and a.mstorgnhirarchyid=e.id and a.activeflg=1 and a.deleteflg=0 and a.fromrecorddifftypeid=f.id and a.torecorddifftypeid=g.id "
		params = append(params, tz.ClientID)
	} else {
		getAssetcount = "SELECT count(a.id) total FROM mstlocationwisedifferentiationmap a,mstrecorddifferentiation b,mstrecorddifferentiation c,mstclient d,mstorgnhierarchy e,mstrecorddifferentiationtype f,mstrecorddifferentiationtype g  where a.clientid=? and a.mstorgnhirarchyid=? and a.fromrecorddiffid=b.id and a.torecorddiffid=c.id and a.clientid=b.clientid and a.mstorgnhirarchyid=b.mstorgnhirarchyid and a.clientid=c.clientid and a.mstorgnhirarchyid=c.mstorgnhirarchyid and a.clientid=d.id and a.mstorgnhirarchyid=e.id and a.activeflg=1 and a.deleteflg=0 and a.fromrecorddifftypeid=f.id and a.torecorddifftypeid=g.id "
		params = append(params, tz.ClientID)
		params = append(params, tz.MstorgnhirarchyID)
	}
	err := db.QueryRow(getAssetcount, params...).Scan(&value.Total)

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
func GetOrgnType(ClientID int64, OrgnID int64, db *sql.DB) (int64, error) {
	logger.Log.Println("In side GetOrgnType")
	var OrgnTypeID int64
	var sql = "SELECT mstorgnhierarchytypeid FROM mstorgnhierarchy WHERE clientid=? AND id=?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		logger.Log.Println(err)
		return OrgnTypeID, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(ClientID, OrgnID)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetOrgnType Get Statement Prepare Error", err)
		return OrgnTypeID, err
	}
	for rows.Next() {
		err := rows.Scan(&OrgnTypeID)
		logger.Log.Println("Error is >>>>>>>", err)
	}
	return OrgnTypeID, nil
}

func UpdateLocation(tz *entities.LocationPriorityEntity, db *sql.DB) error {
	logger.Log.Println("In side UpdateLocation")
	stmt, err := db.Prepare(updateLocation)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateLocation Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.ClientID, tz.MstorgnhirarchyID, tz.Location, tz.Recorddifftypeid, tz.Recorddiffid, tz.ToRecorddifftypeid, tz.ToRecorddiffid, tz.ID)
	if err != nil {
		logger.Log.Println("UpdateLocation Execute Statement  Error", err)
		return err
	}
	return nil
}

func DeleteLocation(tz *entities.LocationPriorityEntity, db *sql.DB) error {
	logger.Log.Println("In side DeleteLocation")
	stmt, err := db.Prepare(deleteLocation)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteLocation Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.ID)
	if err != nil {
		logger.Log.Println("DeleteLocation Execute Statement  Error", err)
		return err
	}
	return nil
}
func CheckDuplicateLocationforupdate(tz *entities.LocationPriorityEntity, db *sql.DB) (entities.LocationEntities, error) {
	logger.Log.Println("In side CheckDuplicateAsset")
	value := entities.LocationEntities{}
	err := db.QueryRow(checkduplicateLocationforupdate, tz.ClientID, tz.MstorgnhirarchyID, tz.Location, tz.Recorddifftypeid, tz.Recorddiffid, tz.ID).Scan(&value.Total)
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
