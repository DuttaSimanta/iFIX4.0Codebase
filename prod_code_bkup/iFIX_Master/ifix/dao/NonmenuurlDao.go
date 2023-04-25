package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
)

var insertNonmenuurl = "INSERT INTO mstnonmenuurl (clientid, mstorgnhirarchyid, urlid, url) VALUES (?,?,?,?)"
var duplicateNonmenuurl = "SELECT count(id) total FROM  mstnonmenuurl WHERE clientid = ? AND mstorgnhirarchyid = ? AND urlid = ? AND url = ? AND deleteflg = 0 and activeflg=1"
var getNonmenuurl = "SELECT a.id as Id ,a.clientid as Clientid,a.mstorgnhirarchyid as Mstorgnhirarchyid,a.urlid as UrlId,a.url as Url,a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.name as Urlname FROM mstnonmenuurl a,mstclient b,mstorgnhierarchy c,msturlkey d WHERE  a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.urlid=d.id ORDER BY a.id DESC LIMIT ?,?"
var getNonmenuurlcount = "SELECT count(a.id) total FROM mstnonmenuurl a,mstclient b,mstorgnhierarchy c,msturlkey d WHERE  a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.urlid=d.id"
var updateNonmenuurl = "UPDATE mstnonmenuurl SET clientid = ?, mstorgnhirarchyid = ?, urlid = ?, url = ? WHERE id = ? "
var deleteNonmenuurl = "UPDATE mstnonmenuurl SET deleteflg = '1' WHERE id = ? "
var geturlkey = "SELECT id as ID,name as Urlkeyname FROM msturlkey WHERE deleteflg = 0 and activeflg=1"
var urlbykey = "SELECT url from mstnonmenuurl where clientid=? and mstorgnhirarchyid =? and urlid in (select id from msturlkey where name =? and activeflg=1 and deleteflg=0) and activeflg=1 and deleteflg=0"

//var getNonmenuurl = "SELECT FROM mstnonmenuurl WHERE  AND deleteflg =0 and activeflg=1 ORDER BY id DESC LIMIT ?,?"
func (dbc DbConn) CheckDuplicateNonmenuurl(tz *entities.NonmenuurlEntity) (entities.NonmenuurlEntities, error) {
	logger.Log.Println("In side CheckDuplicateNonmenuurl")
	value := entities.NonmenuurlEntities{}
	err := dbc.DB.QueryRow(duplicateNonmenuurl, tz.Clientid, tz.Mstorgnhirarchyid, tz.UrlId, tz.Url).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateNonmenuurl Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertNonmenuurl(tz *entities.NonmenuurlEntity) (int64, error) {
	logger.Log.Println("In side InsertNonmenuurl")
	logger.Log.Println("Query -->", insertNonmenuurl)
	stmt, err := dbc.DB.Prepare(insertNonmenuurl)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertNonmenuurl Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.UrlId, tz.Url)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.UrlId, tz.Url)
	if err != nil {
		logger.Log.Println("InsertNonmenuurl Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllNonmenuurl(page *entities.NonmenuurlEntity) ([]entities.NonmenuurlEntity, error) {
	logger.Log.Println("In side GelAllNonmenuurl")
	logger.Log.Println("Query --->", getNonmenuurl)

	values := []entities.NonmenuurlEntity{}
	rows, err := dbc.DB.Query(getNonmenuurl, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllNonmenuurl Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.NonmenuurlEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.UrlId, &value.Url, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Urlname)
		values = append(values, value)
	}
	return values, nil
}
func (dbc DbConn) Geturlbykey(page *entities.NonmenuurlEntity) ([]entities.NonmenuurlEntity, error) {
	var query = ""
	values := []entities.NonmenuurlEntity{}
	var params []interface{}
	if page.Type == "code" {
		query = "SELECT url from mstnonmenuurl where clientid=(select clientid from mstorgnhierarchy where code =? and activeflg=1 and deleteflg =0) and mstorgnhirarchyid = (select id from mstorgnhierarchy where code =? and activeflg=1 and deleteflg =0) and urlid in (select id from msturlkey where name =? and activeflg=1 and deleteflg=0) and activeflg=1 and deleteflg=0"
		params=append(params,page.Code)
		params=append(params,page.Code)
		params=append(params,page.Urlname)
	} else {
		query = "SELECT url from mstnonmenuurl where clientid=? and mstorgnhirarchyid =? and urlid in (select id from msturlkey where name =? and activeflg=1 and deleteflg=0) and activeflg=1 and deleteflg=0"
		params=append(params,page.Clientid)
		params=append(params,page.Mstorgnhirarchyid)
		params=append(params,page.Urlname)
	}
	rows, err := dbc.DB.Query(query, params...)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("Geturlbykey Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.NonmenuurlEntity{}
		rows.Scan(&value.Url)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateNonmenuurl(tz *entities.NonmenuurlEntity) error {
	logger.Log.Println("In side UpdateNonmenuurl")
	stmt, err := dbc.DB.Prepare(updateNonmenuurl)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateNonmenuurl Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.UrlId, tz.Url, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateNonmenuurl Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteNonmenuurl(tz *entities.NonmenuurlEntity) error {
	logger.Log.Println("In side DeleteNonmenuurl")
	stmt, err := dbc.DB.Prepare(deleteNonmenuurl)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteNonmenuurl Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteNonmenuurl Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetNonmenuurlCount(tz *entities.NonmenuurlEntity) (entities.NonmenuurlEntities, error) {
	logger.Log.Println("In side GetNonmenuurlCount")
	value := entities.NonmenuurlEntities{}
	err := dbc.DB.QueryRow(getNonmenuurlcount).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		log.Print("GetNonmenuurlCount Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) GetAllUrlkey(page *entities.MsturlkeyInputEntity) ([]entities.MsturlkeyEntity, error) {
	logger.Log.Println("In side GetAllUrlkey")
	logger.Log.Println("Query --->", geturlkey)

	values := []entities.MsturlkeyEntity{}
	rows, err := dbc.DB.Query(geturlkey)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllUrlkey Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MsturlkeyEntity{}
		err = rows.Scan(&value.Id, &value.Urlkeyname)
		//logger.Log.Println(err)
		values = append(values, value)
	}
	//logger.Log.Println(values)
	return values, nil
}
