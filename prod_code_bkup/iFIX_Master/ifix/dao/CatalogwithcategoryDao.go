package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"strconv"
	"strings"
)

var insertCatalogwithcategory = "INSERT INTO mapcatalogwithcategory (clientid, mstorgnhirarchyid, fromrecorddifftypeid, fromrecorddiffid, torecorddifftypeid, torecorddiffid, catalogid,forrecorddiffid) VALUES (?,?,?,?,?,?,?,?)"
var duplicateCatalogwithcategory = "SELECT count(id) total FROM  mapcatalogwithcategory WHERE clientid = ? AND mstorgnhirarchyid = ? AND torecorddiffid = ? AND catalogid = ? AND deleteflg = 0"

// var getCatalogwithcategory = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.fromrecorddifftypeid as Fromrecorddifftypeid, a.fromrecorddiffid as Fromrecorddiffid, a.torecorddifftypeid as Torecorddifftypeid, a.torecorddiffid as Torecorddiffid,a.catalogid as Catalogid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Fromrecorddifftypename,e.name as Fromrecorddiffname,f.typename as Torecorddifftypename,g.name as Torecorddiffname,h.catalogname as Catalogname FROM mapcatalogwithcategory a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstrecorddifferentiationtype f,mstrecorddifferentiation g,mstcatalog h WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.forrecorddiffid =0 AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.fromrecorddifftypeid=d.id and a.fromrecorddiffid=e.id and a.torecorddifftypeid=f.id and a.torecorddiffid=g.id  and a.catalogid= h.id ORDER BY a.id DESC LIMIT ?,?"
// var getCatalogwithcategorycount = "SELECT count(a.id) as total FROM mapcatalogwithcategory a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstrecorddifferentiationtype f,mstrecorddifferentiation g,mstcatalog h WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.forrecorddiffid =0  AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.fromrecorddifftypeid=d.id and a.fromrecorddiffid=e.id and a.torecorddifftypeid=f.id and a.torecorddiffid=g.id  and a.catalogid= h.id"
var updateCatalogwithcategory = "UPDATE mapcatalogwithcategory SET mstorgnhirarchyid = ?, fromrecorddifftypeid = ?, fromrecorddiffid = ?, torecorddifftypeid = ?, torecorddiffid = ?, catalogid = ? WHERE id = ? "
var deleteCatalogwithcategory = "UPDATE mapcatalogwithcategory SET deleteflg = '1' WHERE id = ? "
var deletemappedcategory = "UPDATE mapcatalogwithcategory SET deleteflg = '1' WHERE forrecorddiffid = ? and deleteflg=0"

//var categorybycatalog = "SELECT a.torecorddifftypeid,a.torecorddiffid,b.name as torecorddiffname,a.fromrecorddifftypeid,a.fromrecorddiffid,c.seqno  from mapcatalogwithcategory a,mstrecorddifferentiation b,mstrecorddifferentiation c where a.clientid = ? AND a.mstorgnhirarchyid = ?  and a.catalogid=? and a.forrecorddiffid=0 and a.torecorddiffid=b.id and a.fromrecorddiffid=c.id and a.activeflg=1 AND a.deleteflg =0 and b.activeflg=1 AND b.deleteflg =0 and c.activeflg=1 AND c.deleteflg =0 "
var categorybycatalog = "SELECT distinct b.name as torecorddiffname,b.parentcategorynames from mapcatalogwithcategory a,mstrecorddifferentiation b where a.clientid = ? AND a.mstorgnhirarchyid = ?  and a.catalogid=? and a.forrecorddiffid=0 and a.torecorddiffid=b.id and a.activeflg=1 AND a.deleteflg =0 and b.activeflg=1 AND b.deleteflg =0  "
var gettorecorddiffids = "SELECT concat(parentcategoryids, ?,id) FROM iFIX.mstrecorddifferentiation where clientid=? AND mstorgnhirarchyid=? AND (parentcategoryids like ? or parentcategoryids like ? or parentcategoryids like ?)"
var categorybyparent = "select distinct name,parentcategorynames from mstrecorddifferentiation where parentid in ( SELECT a.id FROM mstrecorddifferentiation a,mapcatalogwithcategory b where a.clientid=? and a.mstorgnhirarchyid=? and a.parentcategorynames =? and a.id=b.torecorddiffid  and a.activeflg=1 and a.deleteflg=0 and b.activeflg=1 and b.deleteflg=0) and activeflg=1 and deleteflg=0;"
var fromtypebydiffname = "SELECT b.fromrecorddifftypeid,b.fromrecorddiffid,a.seqno,b.torecorddiffid FROM mstrecordtype b,mstrecorddifferentiation a where b.clientid=? and b.mstorgnhirarchyid=? and b.fromrecorddiffid=a.id and b.torecorddiffid in (select id from mstrecorddifferentiation where parentcategorynames=? and activeflg=1 and deleteflg=0)  and b.activeflg=1 and b.deleteflg=0 and a.activeflg=1 and a.deleteflg=0;"

func (dbc DbConn) Getalltorecorddiffids(page *entities.CatalogwithcategoryEntity) (entities.CatalogwithcategoryEntity, error) {
	logger.Log.Println("In side Getalltorecorddiffids")
	values := entities.CatalogwithcategoryEntity{}

	rows, err := dbc.DB.Query(gettorecorddiffids, "->", page.Clientid, page.Mstorgnhirarchyid, "%->"+strconv.FormatInt(page.Torecorddiffid, 10), "%->"+strconv.FormatInt(page.Torecorddiffid, 10)+"->%", strconv.FormatInt(page.Torecorddiffid, 10)+"->%")
	defer rows.Close()
	if err != nil {
		logger.Log.Println("Getalltorecorddiffids Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		var torecorddiffids string
		rows.Scan(&torecorddiffids)
		var ids = strings.Split(torecorddiffids, "->")
		for i := 0; i < len(ids); i++ {
			x, err := strconv.ParseInt(ids[i], 10, 64)
			if err != nil {
				return values, err
			}
			values.Torecorddiffids = append(values.Torecorddiffids, x)
		}

		//values.Torecorddiffids = append(values.Torecorddiffids, id)
	}
	return values, nil
}
func (dbc DbConn) Getcategorybycatalog(page *entities.CatalogwithcategoryEntity) ([]entities.CatalogwithsingleEntity, error) {
	logger.Log.Println("In side GelAllCatalogwithcategory")
	values := []entities.CatalogwithsingleEntity{}
	rows, err := dbc.DB.Query(categorybycatalog, page.Clientid, page.Mstorgnhirarchyid, page.Catalogid)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllCatalogwithcategory Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.CatalogwithsingleEntity{}
		rows.Scan(&value.Torecorddiffname, &value.Parentpath)
		values = append(values, value)
	}
	return values, nil
}
func (dbc DbConn) Getcategorybyparentname(page *entities.CatalogwithcategoryEntity) ([]entities.CatalogwithsingleEntity, error) {
	values := []entities.CatalogwithsingleEntity{}
	rows, err := dbc.DB.Query(categorybyparent, page.Clientid, page.Mstorgnhirarchyid, page.Parentname)

	if err != nil {
		logger.Log.Println("Getcategorybyparentname Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		value := entities.CatalogwithsingleEntity{}
		rows.Scan(&value.Torecorddiffname, &value.Parentpath)
		values = append(values, value)
	}
	return values, nil
}
func (dbc DbConn) Getfromtypebydiffname(page *entities.CatalogwithcategoryEntity) ([]entities.CatalogwithsingleEntity, error) {
	values := []entities.CatalogwithsingleEntity{}
	rows, err := dbc.DB.Query(fromtypebydiffname, page.Clientid, page.Mstorgnhirarchyid, page.Parentname)

	if err != nil {
		logger.Log.Println("Getfromtypebydiffname Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		value := entities.CatalogwithsingleEntity{}
		rows.Scan(&value.Fromrecorddifftypeid, &value.Fromrecorddiffid, &value.Seqno, &value.Torecorddiffid)
		values = append(values, value)
	}
	return values, nil
}
func (dbc DbConn) CheckDuplicateCatalogwithcategory(tz *entities.CatalogwithcategoryEntity, i int) (entities.CatalogwithcategoryEntities, error) {
	logger.Log.Println("In side CheckDuplicateCatalogwithcategory")
	value := entities.CatalogwithcategoryEntities{}
	err := dbc.DB.QueryRow(duplicateCatalogwithcategory, tz.Clientid, tz.Mstorgnhirarchyid, tz.Torecorddiffids[i], tz.Catalogid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateCatalogwithcategory Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc TxConn) InsertCatalogwithcategory(tz *entities.CatalogwithcategoryEntity, i int) (int64, error) {
	logger.Log.Println("In side InsertCatalogwithcategory")
	logger.Log.Println("Query -->", insertCatalogwithcategory)
	if tz.Torecorddiffids[i] == tz.Torecorddiffid {
		tz.Forrecorddiffid = 0
	}
	stmt, err := dbc.TX.Prepare(insertCatalogwithcategory)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertCatalogwithcategory Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Fromrecorddifftypeid, tz.Fromrecorddiffid, tz.Torecorddifftypeid, tz.Torecorddiffid, tz.Catalogid, tz.Torecorddiffid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Fromrecorddifftypeid, tz.Fromrecorddiffid, tz.Torecorddifftypeid, tz.Torecorddiffids[i], tz.Catalogid, tz.Forrecorddiffid)
	if err != nil {
		logger.Log.Println("InsertCatalogwithcategory Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllCatalogwithcategory(tz *entities.CatalogwithcategoryEntity, OrgnType int64) ([]entities.CatalogwithcategoryEntity, error) {
	logger.Log.Println("In side GelAllCatalogwithcategory")
	values := []entities.CatalogwithcategoryEntity{}
	var getCatalogwithcategory string
	var params []interface{}
	if OrgnType == 1 {
		getCatalogwithcategory = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.fromrecorddifftypeid as Fromrecorddifftypeid, a.fromrecorddiffid as Fromrecorddiffid, a.torecorddifftypeid as Torecorddifftypeid, a.torecorddiffid as Torecorddiffid,a.catalogid as Catalogid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Fromrecorddifftypename,e.name as Fromrecorddiffname,f.typename as Torecorddifftypename,g.name as Torecorddiffname,h.catalogname as Catalogname FROM mapcatalogwithcategory a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstrecorddifferentiationtype f,mstrecorddifferentiation g,mstcatalog h WHERE  a.forrecorddiffid =0 AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.fromrecorddifftypeid=d.id and a.fromrecorddiffid=e.id and a.torecorddifftypeid=f.id and a.torecorddiffid=g.id  and a.catalogid= h.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		getCatalogwithcategory = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.fromrecorddifftypeid as Fromrecorddifftypeid, a.fromrecorddiffid as Fromrecorddiffid, a.torecorddifftypeid as Torecorddifftypeid, a.torecorddiffid as Torecorddiffid,a.catalogid as Catalogid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Fromrecorddifftypename,e.name as Fromrecorddiffname,f.typename as Torecorddifftypename,g.name as Torecorddiffname,h.catalogname as Catalogname FROM mapcatalogwithcategory a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstrecorddifferentiationtype f,mstrecorddifferentiation g,mstcatalog h WHERE a.clientid = ? AND a.forrecorddiffid =0 AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.fromrecorddifftypeid=d.id and a.fromrecorddiffid=e.id and a.torecorddifftypeid=f.id and a.torecorddiffid=g.id  and a.catalogid= h.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		getCatalogwithcategory = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.fromrecorddifftypeid as Fromrecorddifftypeid, a.fromrecorddiffid as Fromrecorddiffid, a.torecorddifftypeid as Torecorddifftypeid, a.torecorddiffid as Torecorddiffid,a.catalogid as Catalogid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Fromrecorddifftypename,e.name as Fromrecorddiffname,f.typename as Torecorddifftypename,g.name as Torecorddiffname,h.catalogname as Catalogname FROM mapcatalogwithcategory a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstrecorddifferentiationtype f,mstrecorddifferentiation g,mstcatalog h WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.forrecorddiffid =0 AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.fromrecorddifftypeid=d.id and a.fromrecorddiffid=e.id and a.torecorddifftypeid=f.id and a.torecorddiffid=g.id  and a.catalogid= h.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}

	rows, err := dbc.DB.Query(getCatalogwithcategory, params...)

	// rows, err := dbc.DB.Query(getCatalogwithcategory, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllCatalogwithcategory Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.CatalogwithcategoryEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Fromrecorddifftypeid, &value.Fromrecorddiffid, &value.Torecorddifftypeid, &value.Torecorddiffid, &value.Catalogid, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Fromrecorddifftypename, &value.Fromrecorddiffname, &value.Torecorddifftypename, &value.Torecorddiffname, &value.Catalogname)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateCatalogwithcategory(tz *entities.CatalogwithcategoryEntity) error {
	logger.Log.Println("In side UpdateCatalogwithcategory")
	stmt, err := dbc.DB.Prepare(updateCatalogwithcategory)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateCatalogwithcategory Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Fromrecorddifftypeid, tz.Fromrecorddiffid, tz.Torecorddifftypeid, tz.Torecorddiffid, tz.Catalogid, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateCatalogwithcategory Execute Statement  Error", err)
		return err
	}
	return nil
}

func DeleteCatalogwithcategory(tx *sql.Tx, tz *entities.CatalogwithcategoryEntity) error {
	logger.Log.Println("In side DeleteCatalogwithcategory")
	stmt, err := tx.Prepare(deleteCatalogwithcategory)

	if err != nil {
		logger.Log.Println("DeleteCatalogwithcategory Prepare Statement  Error", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteCatalogwithcategory Execute Statement  Error", err)
		return err
	}
	return nil
}
func Deletemappedcategory(tx *sql.Tx, tz *entities.CatalogwithcategoryEntity) error {
	logger.Log.Println("In side DeleteCatalogwithcategory")
	stmt, err := tx.Prepare(deletemappedcategory)

	if err != nil {
		logger.Log.Println("Deletemappedcategory Prepare Statement  Error", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(tz.Forrecorddiffid)
	if err != nil {
		logger.Log.Println("Deletemappedcategory Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetCatalogwithcategoryCount(tz *entities.CatalogwithcategoryEntity, OrgnTypeID int64) (entities.CatalogwithcategoryEntities, error) {
	logger.Log.Println("In side GetCatalogwithcategoryCount")
	value := entities.CatalogwithcategoryEntities{}
	var getCatalogwithcategorycount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getCatalogwithcategorycount = "SELECT count(a.id) as total FROM mapcatalogwithcategory a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstrecorddifferentiationtype f,mstrecorddifferentiation g,mstcatalog h WHERE a.forrecorddiffid =0  AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.fromrecorddifftypeid=d.id and a.fromrecorddiffid=e.id and a.torecorddifftypeid=f.id and a.torecorddiffid=g.id  and a.catalogid= h.id"
	} else if OrgnTypeID == 2 {
		getCatalogwithcategorycount = "SELECT count(a.id) as total FROM mapcatalogwithcategory a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstrecorddifferentiationtype f,mstrecorddifferentiation g,mstcatalog h WHERE a.clientid = ? AND a.forrecorddiffid =0  AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.fromrecorddifftypeid=d.id and a.fromrecorddiffid=e.id and a.torecorddifftypeid=f.id and a.torecorddiffid=g.id  and a.catalogid= h.id"
		params = append(params, tz.Clientid)
	} else {
		getCatalogwithcategorycount = "SELECT count(a.id) as total FROM mapcatalogwithcategory a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstrecorddifferentiationtype f,mstrecorddifferentiation g,mstcatalog h WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.forrecorddiffid =0  AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.fromrecorddifftypeid=d.id and a.fromrecorddiffid=e.id and a.torecorddifftypeid=f.id and a.torecorddiffid=g.id  and a.catalogid= h.id"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getCatalogwithcategorycount, params...).Scan(&value.Total)

	// err := dbc.DB.QueryRow(getCatalogwithcategorycount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetCatalogwithcategoryCount Get Statement Prepare Error", err)
		return value, err
	}
}
