package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
)

var insertUserWithGroupAndCategory = "INSERT INTO mapuserwithgroupandcategory (clientid, mstorgnhirarchyid, recorddifftypeid, recorddiffid, categoryid, groupid, userid) VALUES (?,?,?,?,?,?,?)"
var duplicateUserWithGroupAndCategory = "SELECT count(id) total FROM  mapuserwithgroupandcategory WHERE clientid = ? AND mstorgnhirarchyid = ? AND recorddifftypeid = ? AND recorddiffid = ? AND categoryid = ? AND userid = ? AND deleteflg = 0 and activeflg=1"

// var getUserUpdatedID = "SELECT id FROM  mapuserwithgroupandcategory WHERE clientid = ? AND mstorgnhirarchyid = ? AND recorddifftypeid = ? AND recorddiffid = ? AND categoryid = ? AND userid = ? AND deleteflg = 0 and activeflg=1"

//var getUserWithGroupAndCategory = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.moduleid as Moduleid, a.roleid as Roleid, a.menuid as Menuid,a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.rolename as Rolename,e.modulename as Modulename,f.menudesc as Menuname FROM mapuserwithgroupandcategory a,mstclient b,mstorgnhierarchy c,mstclientuserrole d,mstmodule e,dtlclientmenuinfo f WHERE a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.roleid=d.id AND a.moduleid=e.id AND a.menuid=f.id AND a.clientid=? AND a.mstorgnhirarchyid=? AND a.deleteflg =0 and a.activeflg=1 AND d.deleteflg =0 AND e.deleteflg =0 AND f.deleteflg=0 ORDER BY id DESC LIMIT ?,?"
//var getUserWithGroupAndCategorycount = "SELECT count(a.id) total FROM mapuserwithgroupandcategory a,mstclient b,mstorgnhierarchy c,mstclientuserrole d,mstmodule e,dtlclientmenuinfo f WHERE a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.roleid=d.id AND a.moduleid=e.id AND a.menuid=f.id AND a.clientid=? AND a.mstorgnhirarchyid=? AND a.deleteflg =0 and a.activeflg=1 AND d.deleteflg =0 AND e.deleteflg =0 AND f.deleteflg=0"
var updateUserWithGroupAndCategory = "UPDATE mapuserwithgroupandcategory SET clientid = ?, mstorgnhirarchyid = ?, recorddifftypeid = ?, recorddiffid = ?, categoryid = ?, groupid = ?, userid = ? WHERE id = ? "
var deleteUserWithGroupAndCategory = "UPDATE mapuserwithgroupandcategory SET deleteflg = '1' WHERE id = ? "

//var getUserWithGroupAndCategory = "SELECT id as Id, clientid as Clientid, mstorgnhirarchyid as Mstorgnhirarchyid, moduleid as Moduleid, roleid as Roleid, menuid as Menuid FROM mapuserwithgroupandcategory WHERE clientid = ? AND mstorgnhirarchyid = ? AND deleteflg =0 and activeflg=1 ORDER BY id DESC LIMIT ?,?"
func (dbc DbConn) CheckDuplicateUserWithGroupAndCategory(tz *entities.UserWithGroupAndCategoryEntity) (entities.UserWithGroupAndCategoryEntities, error) {
	logger.Log.Println("In side CheckDuplicateUserWithGroupAndCategory")
	logger.Log.Println("Query -->", duplicateUserWithGroupAndCategory)
	logger.Log.Println("Parameter ---->", tz.Clientid, tz.Mstorgnhirarchyid)
	value := entities.UserWithGroupAndCategoryEntities{}
	// var Ids = []int64{}
	err := dbc.DB.QueryRow(duplicateUserWithGroupAndCategory, tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Categoryid, tz.Userid).Scan(&value.Total)
	logger.Log.Println("\n value ", value)
	switch err {
	case sql.ErrNoRows:
		logger.Log.Println("\n Inside CASE 1....")
		value.Total = 0
		return value, nil
	case nil:
		logger.Log.Println("\n Inside CASE 2....")
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateUserWithGroupAndCategory Get Statement Prepare Error", err)
		return value, err
	}
	// if err != nil {
	// 	value.Total = 0
	// 	return value, err
	// }
	// if len(Ids) == 0 {
	// 	value.Total = 0
	// 	logger.Log.Println("\n Inside IFFFFFFFF.........", value.Total)
	// 	return value, nil
	// } else {
	// 	value.Total = Ids[0]
	// 	logger.Log.Println("\n ELSEEEEEEEEEE.........", value.Total)
	// 	return value, nil
	// }
}

func (dbc DbConn) InsertUserWithGroupAndCategory(tz *entities.UserWithGroupAndCategoryEntity) (int64, error) {
	logger.Log.Println("In side InsertUserWithGroupAndCategory")
	logger.Log.Println("Query -->", insertUserWithGroupAndCategory)
	stmt, err := dbc.DB.Prepare(insertUserWithGroupAndCategory)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertUserWithGroupAndCategory Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Categoryid, tz.Groupid, tz.Userid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Categoryid, tz.Groupid, tz.Userid)
	if err != nil {
		logger.Log.Println("InsertUserWithGroupAndCategory Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllUserWithGroupAndCategory(page *entities.UserWithGroupAndCategoryEntity, OrgnType int64) ([]entities.UserWithGroupAndCategoryEntity, error) {
	values := []entities.UserWithGroupAndCategoryEntity{}
	var getUserWithGroupAndCategory string
	var params []interface{}
	if OrgnType == 1 {
		getUserWithGroupAndCategory = "Select a.id as Id, a.clientid as Clientid, (select name from mstclient where id=a.clientid) as Clientname, a.mstorgnhirarchyid as Mstorgnhirarchyid, (select name from mstorgnhierarchy where id=a.mstorgnhirarchyid) as Mstorgnhirarchyname, a.recorddifftypeid as Recorddifftypeid,(select typename from mstrecorddifferentiationtype where id=a.recorddifftypeid and deleteflg=0 and activeflg=1) as Recorddifftypename, a.recorddiffid as Recorddiffid,(select name from mstrecorddifferentiation where id=a.recorddiffid and deleteflg=0 and activeflg=1) as Recorddiffname, a.categoryid as Categoryid,(select name from mstrecorddifferentiation where id=a.categoryid and deleteflg=0 and activeflg=1) as Categoryname, a.groupid as Groupid,coalesce((select name from mstsupportgrp where id=a.groupid and deleteflg=0 and activeflg=1),'') as Groupname, a.userid as Userid , (select name from mstclientuser where id=a.userid and activeflag=1 and deleteflag=0) as Username, (select loginname from mstclientuser where id=a.userid and activeflag=1 and deleteflag=0) as Userloginname from mapuserwithgroupandcategory a where a.activeflg=1 and a.deleteflg=0 order by id desc limit ?,?"
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else if OrgnType == 2 {
		getUserWithGroupAndCategory = "Select a.id as Id, a.clientid as Clientid, (select name from mstclient where id=a.clientid) as Clientname, a.mstorgnhirarchyid as Mstorgnhirarchyid, (select name from mstorgnhierarchy where id=a.mstorgnhirarchyid) as Mstorgnhirarchyname, a.recorddifftypeid as Recorddifftypeid,(select typename from mstrecorddifferentiationtype where id=a.recorddifftypeid and deleteflg=0 and activeflg=1) as Recorddifftypename, a.recorddiffid as Recorddiffid,(select name from mstrecorddifferentiation where id=a.recorddiffid and deleteflg=0 and activeflg=1) as Recorddiffname, a.categoryid as Categoryid,(select name from mstrecorddifferentiation where id=a.categoryid and deleteflg=0 and activeflg=1) as Categoryname, a.groupid as Groupid,coalesce((select name from mstsupportgrp where id=a.groupid and deleteflg=0 and activeflg=1),'') as Groupname, a.userid as Userid , (select name from mstclientuser where id=a.userid and activeflag=1 and deleteflag=0) as Username, (select loginname from mstclientuser where id=a.userid and activeflag=1 and deleteflag=0) as Userloginname from mapuserwithgroupandcategory a where a.clientid=? and a.activeflg=1 and a.deleteflg=0 order by id desc limit ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else {
		getUserWithGroupAndCategory = "Select a.id as Id, a.clientid as Clientid, (select name from mstclient where id=a.clientid) as Clientname, a.mstorgnhirarchyid as Mstorgnhirarchyid, (select name from mstorgnhierarchy where id=a.mstorgnhirarchyid) as Mstorgnhirarchyname, a.recorddifftypeid as Recorddifftypeid,(select typename from mstrecorddifferentiationtype where id=a.recorddifftypeid and deleteflg=0 and activeflg=1) as Recorddifftypename, a.recorddiffid as Recorddiffid,(select name from mstrecorddifferentiation where id=a.recorddiffid and deleteflg=0 and activeflg=1) as Recorddiffname, a.categoryid as Categoryid,(select name from mstrecorddifferentiation where id=a.categoryid and deleteflg=0 and activeflg=1) as Categoryname, a.groupid as Groupid,coalesce((select name from mstsupportgrp where id=a.groupid and deleteflg=0 and activeflg=1),'') as Groupname, a.userid as Userid , (select name from mstclientuser where id=a.userid and activeflag=1 and deleteflag=0) as Username, (select loginname from mstclientuser where id=a.userid and activeflag=1 and deleteflag=0) as Userloginname from mapuserwithgroupandcategory a where a.clientid=? and a.mstorgnhirarchyid=? and a.activeflg=1 and a.deleteflg=0 order by id desc limit ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Mstorgnhirarchyid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	}
	rows, err := dbc.DB.Query(getUserWithGroupAndCategory, params...)
	defer rows.Close()
	logger.Log.Println("\n params :::   ", params)
	logger.Log.Println("\n rows :::   ", rows)
	if err != nil {
		logger.Log.Println("GetAllUserWithGroupAndCategory Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.UserWithGroupAndCategoryEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Clientname, &value.Mstorgnhirarchyid, &value.Mstorgnhirarchyname, &value.Recorddifftypeid, &value.Recorddifftypename, &value.Recorddiffid, &value.Recorddiffname, &value.Categoryid, &value.Categoryname, &value.Groupid, &value.Groupname, &value.Userid, &value.Username, &value.Userloginname)
		logger.Log.Println("\n Values :::   ", value)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateUserWithGroupAndCategory(tz *entities.UserWithGroupAndCategoryEntity) error {
	logger.Log.Println("In side UpdateUserWithGroupAndCategory")
	logger.Log.Println("Query -->", updateUserWithGroupAndCategory)
	logger.Log.Println("Parameter ---->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Categoryid, tz.Groupid, tz.Userid, tz.Id)

	stmt, err := dbc.DB.Prepare(updateUserWithGroupAndCategory)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateUserWithGroupAndCategory Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Categoryid, tz.Groupid, tz.Userid, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateUserWithGroupAndCategory Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteUserWithGroupAndCategory(tz *entities.UserWithGroupAndCategoryEntity) error {
	logger.Log.Println("In side DeleteUserWithGroupAndCategory")
	logger.Log.Println("Query -->", deleteUserWithGroupAndCategory)
	logger.Log.Println("Parameter ---->", tz.Id)
	stmt, err := dbc.DB.Prepare(deleteUserWithGroupAndCategory)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteUserWithGroupAndCategory Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteUserWithGroupAndCategory Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetUserWithGroupAndCategoryCount(tz *entities.UserWithGroupAndCategoryEntity, OrgnTypeID int64) (entities.UserWithGroupAndCategoryEntities, error) {
	value := entities.UserWithGroupAndCategoryEntities{}
	var getUserWithGroupAndCategorycount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getUserWithGroupAndCategorycount = "SELECT count(a.id) total FROM mapuserwithgroupandcategory a WHERE a.deleteflg =0 and a.activeflg=1"
	} else if OrgnTypeID == 2 {
		getUserWithGroupAndCategorycount = "SELECT count(a.id) total FROM mapuserwithgroupandcategory a WHERE a.clientid=? AND a.deleteflg =0 and a.activeflg=1"
		params = append(params, tz.Clientid)
	} else {
		getUserWithGroupAndCategorycount = "SELECT count(a.id) total FROM mapuserwithgroupandcategory a WHERE a.clientid=? AND a.mstorgnhirarchyid=? AND a.deleteflg =0 and a.activeflg=1"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getUserWithGroupAndCategorycount, params...).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		log.Print("GetUserWithGroupAndCategoryCount Get Statement Prepare Error", err)
		return value, err
	}
}

// func (dbc DbConn) GetUserUpdatedID(tz *entities.UserWithGroupAndCategoryForBulkUploadEntity) (int64, error) {
// 	logger.Log.Println("In side GetUpdatedID")
// 	var value int64
// 	var values []int64
// 	err := dbc.DB.QueryRow(getUserUpdatedID, tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Categoryid, tz.Userid).Scan(&value)
// 	values = append(values, value)
// 	if err != nil {
// 		return 0, err
// 	}
// 	if len(values) == 0 {
// 		return 0, nil
// 	} else {
// 		return values[0], nil
// 	}
// }
