package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertMapCategoryWithKeyword = "INSERT INTO mapcategorywithkeyword (clientid, mstorgnhirarchyid, keyword, categoryvalue) VALUES (?,?,?,?)"
var duplicateMapCategoryWithKeyword = "SELECT count(id) total FROM  mapcategorywithkeyword WHERE clientid = ? AND mstorgnhirarchyid = ? AND keyword = ? AND categoryvalue = ? AND deleteflg = 0"

//var getMapCategoryWithKeyword = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.keyword as keyword, a.categoryvalue as categoryvalue, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname FROM mapcategorywithkeyword a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid=b.id AND a.mstorgnhirarchyid=c.id ORDER BY a.id DESC LIMIT ?,?"
//var getMapCategoryWithKeywordcount = "SELECT count(a.id) as total FROM mapcategorywithkeyword a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid=b.id AND a.mstorgnhirarchyid=c.id "
var updateMapCategoryWithKeyword = "UPDATE mapcategorywithkeyword SET mstorgnhirarchyid = ?, keyword = ?, categoryvalue = ? WHERE id = ? "
var deleteMapCategoryWithKeyword = "UPDATE mapcategorywithkeyword SET deleteflg = '1' WHERE id = ? "

func (dbc DbConn) CheckDuplicateMapCategoryWithKeyword(tz *entities.MapCategoryWithKeywordEntity) (entities.MapCategoryWithKeywordEntities, error) {
	logger.Log.Println("In side CheckDuplicateMapCategoryWithKeyword")
	value := entities.MapCategoryWithKeywordEntities{}
	err := dbc.DB.QueryRow(duplicateMapCategoryWithKeyword, tz.Clientid, tz.Mstorgnhirarchyid, tz.Keyword, tz.Categoryvalue).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMapCategoryWithKeyword Get Statement Prepare Error", err)
		return value, err
	}
}
func (dbc DbConn) CheckDuplicateMapCategoryWithKeywordforstupa(tz *entities.MapCategoryWithKeywordEntity) (entities.MapCategoryWithKeywordEntities, error) {
	logger.Log.Println("In side CheckDuplicateMapCategoryWithKeyword")
	value := entities.MapCategoryWithKeywordEntities{}
	err := dbc.DB.QueryRow(duplicateMapCategoryWithKeyword, 1, 1, tz.Keyword, tz.Categoryvalue).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMapCategoryWithKeyword Get Statement Prepare Error", err)
		return value, err
	}
}
func (dbc TxConn) InsertMapCategoryWithKeyword(tz *entities.MapCategoryWithKeywordEntity) (int64, error) {
	logger.Log.Println("In side InsertMapCategoryWithKeyword")
	logger.Log.Println("Query -->", insertMapCategoryWithKeyword)
	stmt, err := dbc.TX.Prepare(insertMapCategoryWithKeyword)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMapCategoryWithKeyword Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Keyword, tz.Categoryvalue)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Keyword, tz.Categoryvalue)
	if err != nil {
		logger.Log.Println("InsertMapCategoryWithKeyword Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllMapCategoryWithKeyword(tz *entities.MapCategoryWithKeywordEntity, OrgnType int64) ([]entities.MapCategoryWithKeywordEntity, error) {
	logger.Log.Println("In side GelAllMapCategoryWithKeyword")
	values := []entities.MapCategoryWithKeywordEntity{}

	var getMapCategoryWithKeyword string
	var params []interface{}
	if OrgnType == 1 {
		getMapCategoryWithKeyword = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.keyword as keyword, a.categoryvalue as categoryvalue, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname FROM mapcategorywithkeyword a,mstclient b,mstorgnhierarchy c WHERE  a.deleteflg =0 and a.activeflg=1 AND a.clientid=b.id AND a.mstorgnhirarchyid=c.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		getMapCategoryWithKeyword = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.keyword as keyword, a.categoryvalue as categoryvalue, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname FROM mapcategorywithkeyword a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid=b.id AND a.mstorgnhirarchyid=c.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		getMapCategoryWithKeyword = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.keyword as keyword, a.categoryvalue as categoryvalue, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname FROM mapcategorywithkeyword a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid=b.id AND a.mstorgnhirarchyid=c.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}
	rows, err := dbc.DB.Query(getMapCategoryWithKeyword, params...)

	//rows, err := dbc.DB.Query(getMapCategoryWithKeyword, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMapCategoryWithKeyword Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MapCategoryWithKeywordEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Keyword, &value.Categoryvalue, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname)
		values = append(values, value)
	}
	return values, nil
}

func (dbc TxConn) UpdateMapCategoryWithKeyword(tz *entities.MapCategoryWithKeywordEntity) error {
	logger.Log.Println("In side UpdateMapCategoryWithKeyword")
	stmt, err := dbc.TX.Prepare(updateMapCategoryWithKeyword)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMapCategoryWithKeyword Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Keyword, tz.Categoryvalue, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMapCategoryWithKeyword Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMapCategoryWithKeyword(tz *entities.MapCategoryWithKeywordEntity) error {
	logger.Log.Println("In side DeleteMapCategoryWithKeyword")
	stmt, err := dbc.DB.Prepare(deleteMapCategoryWithKeyword)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMapCategoryWithKeyword Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMapCategoryWithKeyword Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMapCategoryWithKeywordCount(tz *entities.MapCategoryWithKeywordEntity, OrgnTypeID int64) (entities.MapCategoryWithKeywordEntities, error) {
	logger.Log.Println("In side GetMapCategoryWithKeywordCount")
	value := entities.MapCategoryWithKeywordEntities{}
	var getMapCategoryWithKeywordcount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getMapCategoryWithKeywordcount = "SELECT count(a.id) as total FROM mapcategorywithkeyword a,mstclient b,mstorgnhierarchy c WHERE  a.deleteflg =0 and a.activeflg=1 AND a.clientid=b.id AND a.mstorgnhirarchyid=c.id "
	} else if OrgnTypeID == 2 {
		getMapCategoryWithKeywordcount = "SELECT count(a.id) as total FROM mapcategorywithkeyword a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid=b.id AND a.mstorgnhirarchyid=c.id "
		params = append(params, tz.Clientid)
	} else {
		getMapCategoryWithKeywordcount = "SELECT count(a.id) as total FROM mapcategorywithkeyword a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid=b.id AND a.mstorgnhirarchyid=c.id "
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getMapCategoryWithKeywordcount, params...).Scan(&value.Total)

	//err = dbc.DB.QueryRow(getMapCategoryWithKeywordcount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMapCategoryWithKeywordCount Get Statement Prepare Error", err)
		return value, err
	}
}

// func (dbc DbConn) GetAllkeyword(page *entities.MapCategoryWithKeywordEntity) ([]entities.Getkeyword, error) {
// 	logger.Log.Println("In side GelAllRecorddifferentiation")
// 	logger.Log.Println(getRecorddifferentiation)
// 	values := []entities.Getkeyword{}
// 	var Getallkeyword = "select id ,keyword from mapcategorywithkeyword where clientid=? and mstorgnhirarchyid=? and deleteflg=0"
// 	rows, err := dbc.DB.Query(Getallkeyword, page.Clientid, page.Mstorgnhirarchyid)
// 	defer rows.Close()
// 	if err != nil {
// 		logger.Log.Println("GetAllRecorddifferentiation Get Statement Prepare Error", err)
// 		return values, err
// 	}
// 	for rows.Next() {
// 		value := entities.Getkeyword{}
// 		rows.Scan(&value.Id, &value.Keyword)
// 		values = append(values, value)
// 	}
// 	return values, nil
// }
func (dbc DbConn) GetAllCategoryvalue(page *entities.MapCategoryWithKeywordEntity) ([]entities.Getkeyword, error) {
	logger.Log.Println("In side GelAllRecorddifferentiation")
	logger.Log.Println(getRecorddifferentiation)
	values := []entities.Getkeyword{}
	var Getallkeyword = "select id ,keyword,categoryvalue from mapcategorywithkeyword where clientid=? and mstorgnhirarchyid=? and keyword like ? and deleteflg=0"
	rows, err := dbc.DB.Query(Getallkeyword, page.Clientid, page.Mstorgnhirarchyid, "%"+page.Keyword+"%")
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllRecorddifferentiation Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.Getkeyword{}
		rows.Scan(&value.Id, &value.Keyword, &value.Categoryvalue)
		values = append(values, value)
	}
	return values, nil
}
