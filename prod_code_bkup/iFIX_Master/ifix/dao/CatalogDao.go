package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
)

var insertCatalog = "INSERT INTO mstcatalog (clientid, mstorgnhirarchyid, catalogname) VALUES (?,?,?)"
var duplicateCatalog = "SELECT count(id) total FROM  mstcatalog WHERE clientid = ? AND mstorgnhirarchyid = ? AND catalogname = ? AND deleteflg = 0"

// var getCatalog = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.catalogname as Catalogname, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname FROM mstcatalog a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ?  AND a.mstorgnhirarchyid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id ORDER BY a.id DESC LIMIT ?,?"
// var getCatalogcount = "SELECT count(a.id) as total FROM mstcatalog a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ?  AND a.mstorgnhirarchyid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id"
var updateCatalog = "UPDATE mstcatalog SET mstorgnhirarchyid = ?, catalogname = ? WHERE id = ? "
var deleteCatalog = "UPDATE mstcatalog SET deleteflg = '1' WHERE id = ? "
var getparentcategory = "SELECT parentcategoryids as Id,parentcategorynames as NAME FROM mstrecorddifferentiation where id=?;"

//var getparentcategory ="SELECT parentcategoryids as Id  FROM mstrecorddifferentiation where parentcategorynames=?;"
var getttype = "SELECT  fromrecorddifftypeid as TypeId, fromrecorddiffid as ID FROM mstrecordtype where clientid=? and  mstorgnhirarchyid=? and torecorddiffid=? and activeflg=1 and deleteflg=0;"
var getcatalogid = "SELECT forrecorddiffid FROM mapcatalogwithcategory  where torecorddiffid=? and activeflg=1 and deleteflg=0 ;"

func (dbc DbConn) Getmappedcatalogid(page *entities.CatalogEntity) ([]entities.ParentCategoryEntity, error) {
	values := []entities.ParentCategoryEntity{}

	rows, err := dbc.DB.Query(getcatalogid, page.Id)
	if err != nil {
		logger.Log.Println("Getmappedcatalogid Get Statement Prepare Error", err)
		log.Println("Getmappedcatalogid Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		value := entities.ParentCategoryEntity{}
		rows.Scan(&value.ID)
		values = append(values, value)
	}
	return values, nil
}
func (dbc DbConn) Getcatalogdetails(page *entities.CatalogEntity, ids string) ([]entities.ParentCategoryEntity, error) {
	values := []entities.ParentCategoryEntity{}
	log.Print(page.Clientid, page.Mstorgnhirarchyid, page.Fromrecorddifftypeid, page.Fromrecorddiffid)
	logger.Log.Print(page.Clientid, page.Mstorgnhirarchyid, page.Fromrecorddifftypeid, page.Fromrecorddiffid)
	var getcatalog = "SELECT a.id,a.catalogname name,b.torecorddiffid FROM mapcatalogwithcategory b,mstcatalog a where b.clientid=? and b.mstorgnhirarchyid=? and b.fromrecorddifftypeid=? and b.fromrecorddiffid=? and b.torecorddiffid in (" + ids + ") and b.catalogid=a.id and b.activeflg=1 and b.deleteflg=0 and a.activeflg=1 and a.deleteflg=0;"
	log.Print(getcatalog)
	rows, err := dbc.DB.Query(getcatalog, page.Clientid, page.Mstorgnhirarchyid, page.Fromrecorddifftypeid, page.Fromrecorddiffid)
	if err != nil {
		logger.Log.Println("Getcatalogdetails Get Statement Prepare Error", err)
		log.Println("Getcatalogdetails Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		value := entities.ParentCategoryEntity{}
		rows.Scan(&value.ID, &value.NAME, &value.Torecorddiffid)
		values = append(values, value)
	}
	return values, nil
}
func (dbc DbConn) GetAllParentCategoryDetails(page *entities.CatalogEntity) (entities.ParentCategoryEntity, error) {
	logger.Log.Println("Parameter -->", page.Id)
	values := entities.ParentCategoryEntity{}
	rows, err := dbc.DB.Query(getparentcategory, page.Id)
	if err != nil {
		logger.Log.Println("GetAllParentCategoryDetails Get Statement Prepare Error", err)
		log.Println("GetAllParentCategoryDetails Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&values.ID, &values.NAME)
	}
	return values, nil
}
func (dbc DbConn) GetCatalogTickettype(page *entities.CatalogEntity) ([]entities.CatalogTicketTypeEntity, error) {
	logger.Log.Println("In side GetCatalogTickettype==>", getttype)
	logger.Log.Println("Parameter -->", page.Id)
	values := []entities.CatalogTicketTypeEntity{}
	rows, err := dbc.DB.Query(getttype, page.Clientid, page.Mstorgnhirarchyid, page.Id)
	if err != nil {
		logger.Log.Println("GetCatalogTickettype Get Statement Prepare Error", err)
		log.Println("GetCatalogTickettype Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		value := entities.CatalogTicketTypeEntity{}
		rows.Scan(&value.TypeId, &value.Id)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) CheckDuplicateCatalog(tz *entities.CatalogEntity) (entities.CatalogEntities, error) {
	logger.Log.Println("In side CheckDuplicateCatalog")
	value := entities.CatalogEntities{}
	err := dbc.DB.QueryRow(duplicateCatalog, tz.Clientid, tz.Mstorgnhirarchyid, tz.Catalogname).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateCatalog Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertCatalog(tz *entities.CatalogEntity) (int64, error) {
	logger.Log.Println("In side InsertCatalog")
	logger.Log.Println("Query -->", insertCatalog)
	stmt, err := dbc.DB.Prepare(insertCatalog)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertCatalog Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Catalogname)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Catalogname)
	if err != nil {
		logger.Log.Println("InsertCatalog Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllCatalog(tz *entities.CatalogEntity, OrgnType int64) ([]entities.CatalogEntity, error) {
	logger.Log.Println("In side GelAllCatalog")
	values := []entities.CatalogEntity{}
	var getCatalog string
	var params []interface{}
	if OrgnType == 1 {
		getCatalog = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.catalogname as Catalogname, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname FROM mstcatalog a,mstclient b,mstorgnhierarchy c WHERE a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		getCatalog = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.catalogname as Catalogname, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname FROM mstcatalog a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		getCatalog = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.catalogname as Catalogname, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname FROM mstcatalog a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ?  AND a.mstorgnhirarchyid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}

	rows, err := dbc.DB.Query(getCatalog, params...)

	// rows, err := dbc.DB.Query(getCatalog, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllCatalog Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.CatalogEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Catalogname, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname)
		values = append(values, value)
	}
	return values, nil
}
func (dbc DbConn) Getcatalogorgwise(tz *entities.CatalogEntity) ([]entities.CatalogEntitySingle, error) {
	logger.Log.Println("In side GelAllCatalog")
	values := []entities.CatalogEntitySingle{}
	var getCatalog string
	var params []interface{}
	getCatalog = "SELECT a.id as Id, a.catalogname as Catalogname FROM mstcatalog a WHERE a.clientid = ?  AND a.mstorgnhirarchyid = ? AND a.deleteflg =0 and a.activeflg=1 "
	params = append(params, tz.Clientid)
	params = append(params, tz.Mstorgnhirarchyid)

	rows, err := dbc.DB.Query(getCatalog, params...)

	defer rows.Close()
	if err != nil {
		logger.Log.Println("Getcatalogorgwise Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.CatalogEntitySingle{}
		rows.Scan(&value.Id,  &value.Catalogname)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateCatalog(tz *entities.CatalogEntity) error {
	logger.Log.Println("In side UpdateCatalog")
	stmt, err := dbc.DB.Prepare(updateCatalog)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateCatalog Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Catalogname, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateCatalog Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteCatalog(tz *entities.CatalogEntity) error {
	logger.Log.Println("In side DeleteCatalog")
	stmt, err := dbc.DB.Prepare(deleteCatalog)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteCatalog Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteCatalog Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetCatalogCount(tz *entities.CatalogEntity, OrgnTypeID int64) (entities.CatalogEntities, error) {
	logger.Log.Println("In side GetCatalogCount")
	value := entities.CatalogEntities{}
	var getCatalogcount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getCatalogcount = "SELECT count(a.id) as total FROM mstcatalog a,mstclient b,mstorgnhierarchy c WHERE a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id"
	} else if OrgnTypeID == 2 {
		getCatalogcount = "SELECT count(a.id) as total FROM mstcatalog a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id"
		params = append(params, tz.Clientid)
	} else {
		getCatalogcount = "SELECT count(a.id) as total FROM mstcatalog a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ?  AND a.mstorgnhirarchyid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getCatalogcount, params...).Scan(&value.Total)

	// err := dbc.DB.QueryRow(getCatalogcount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetCatalogCount Get Statement Prepare Error", err)
		return value, err
	}
}
