package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertMapexternalattributes = "INSERT INTO mapexternalattributes (clientid, mstorgnhirarchyid, extsystemid, extattr, sysattr) VALUES (?,?,?,?,?)"

// var getmapexternalattributes = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid,a.extsystemid as Systemid,a.activeflg as Activeflg,b.name as Mstorgnhirarchyname,c.name as SystemName,a.extattr as Extattr,a.sysattr as Sysattr FROM mapexternalattributes a ,mstorgnhierarchy b,mstlogintype c WHERE a.clientid=? and a.mstorgnhirarchyid=? and a.mstorgnhirarchyid = b.id and a.extsystemid = c.id and a.activeflg=1 and a.deleteflg =0  ORDER BY a.id DESC LIMIT ?,?"
// var getmapexternalattributescount = "SELECT count(a.id) as total FROM mapexternalattributes a ,mstorgnhierarchy b,mstlogintype c WHERE a.clientid=? and a.mstorgnhirarchyid=? and a.mstorgnhirarchyid = b.id and a.extsystemid = c.id and a.activeflg=1 and a.deleteflg =0  "
var deletemapexternalattributes = "UPDATE mapexternalattributes SET deleteflg = '1' WHERE id = ? "
var mappedattributes = "SELECT extattr, sysattr from mapexternalattributes where clientid=? and mstorgnhirarchyid=? and extsystemid=? and activeflg=1 and deleteflg=0"

func (dbc DbConn) GetMappedattributes(page *entities.MapexternalattributesEntity) ([]entities.Attr, error) {
	logger.Log.Println("In side GetMappedattributes")
	values := []entities.Attr{}
	rows, err := dbc.DB.Query(mappedattributes, page.Clientid, page.Mstorgnhirarchyid, page.Systemid)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetMappedattributes Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.Attr{}
		rows.Scan(&value.Extattr, &value.Sysattr)
		values = append(values, value)
	}
	return values, nil
}

func (dbc TxConn) InsertMapexternalattributes(tz *entities.MapexternalattributesEntity, i int) (int64, error) {
	logger.Log.Println("In side InsertMapexternalattributes")
	logger.Log.Println("Query -->", insertMapexternalattributes)
	stmt, err := dbc.TX.Prepare(insertMapexternalattributes)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMapexternalattributes Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Map[i].Extattr, tz.Systemid, tz.Map[i].Sysattr)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Systemid, tz.Map[i].Extattr, tz.Map[i].Sysattr)
	if err != nil {
		logger.Log.Println("InsertMapexternalattributes Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllMapexternalattributes(page *entities.MapexternalattributesEntity, OrgnType int64) ([]entities.MapexternalattributesEntity, error) {
	logger.Log.Println("In side GelAllMapexternalattributes")
	values := []entities.MapexternalattributesEntity{}
	var getmapexternalattributes string
	var params []interface{}
	if OrgnType == 1 {
		getmapexternalattributes = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid,a.extsystemid as Systemid,a.activeflg as Activeflg,b.name as Mstorgnhirarchyname,c.name as SystemName,a.extattr as Extattr,a.sysattr as Sysattr FROM mapexternalattributes a ,mstorgnhierarchy b,mstlogintype c WHERE  a.mstorgnhirarchyid = b.id and a.extsystemid = c.id and a.activeflg=1 and a.deleteflg =0  ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else if OrgnType == 2 {
		getmapexternalattributes = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid,a.extsystemid as Systemid,a.activeflg as Activeflg,b.name as Mstorgnhirarchyname,c.name as SystemName,a.extattr as Extattr,a.sysattr as Sysattr FROM mapexternalattributes a ,mstorgnhierarchy b,mstlogintype c WHERE a.clientid=? and a.mstorgnhirarchyid = b.id and a.extsystemid = c.id and a.activeflg=1 and a.deleteflg =0  ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else {
		getmapexternalattributes = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid,a.extsystemid as Systemid,a.activeflg as Activeflg,b.name as Mstorgnhirarchyname,c.name as SystemName,a.extattr as Extattr,a.sysattr as Sysattr FROM mapexternalattributes a ,mstorgnhierarchy b,mstlogintype c WHERE a.clientid=? and a.mstorgnhirarchyid=? and a.mstorgnhirarchyid = b.id and a.extsystemid = c.id and a.activeflg=1 and a.deleteflg =0  ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Mstorgnhirarchyid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	}
	rows, err := dbc.DB.Query(getmapexternalattributes, params...)

	// rows, err := dbc.DB.Query(getmapexternalattributes, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMapexternalattributes Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MapexternalattributesEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Systemid, &value.Activeflg, &value.Mstorgnhirarchyname, &value.SystemName, &value.Extattr, &value.Sysattr)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) DeleteMapexternalattributes(tz *entities.MapexternalattributesEntity) error {
	logger.Log.Println("In side DeleteMapexternalattributes")
	stmt, err := dbc.DB.Prepare(deletemapexternalattributes)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMapexternalattributes Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMapexternalattributes Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMapexternalattributesCount(tz *entities.MapexternalattributesEntity, OrgnTypeID int64) (entities.MapexternalattributesEntities, error) {
	logger.Log.Println("In side GetMapexternalattributesCount")
	value := entities.MapexternalattributesEntities{}
	var getmapexternalattributescount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getmapexternalattributescount = "SELECT count(a.id) as total FROM mapexternalattributes a ,mstorgnhierarchy b,mstlogintype c WHERE  a.mstorgnhirarchyid = b.id and a.extsystemid = c.id and a.activeflg=1 and a.deleteflg =0  "
	} else if OrgnTypeID == 2 {
		getmapexternalattributescount = "SELECT count(a.id) as total FROM mapexternalattributes a ,mstorgnhierarchy b,mstlogintype c WHERE a.clientid=? and  a.mstorgnhirarchyid = b.id and a.extsystemid = c.id and a.activeflg=1 and a.deleteflg =0  "
		params = append(params, tz.Clientid)
	} else {
		getmapexternalattributescount = "SELECT count(a.id) as total FROM mapexternalattributes a ,mstorgnhierarchy b,mstlogintype c WHERE a.clientid=? and a.mstorgnhirarchyid=? and a.mstorgnhirarchyid = b.id and a.extsystemid = c.id and a.activeflg=1 and a.deleteflg =0  "
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}

	err := dbc.DB.QueryRow(getmapexternalattributescount, params...).Scan(&value.Total)
	// err := dbc.DB.QueryRow(getmapexternalattributescount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMapexternalattributesCount Get Statement Prepare Error", err)
		return value, err
	}
}
