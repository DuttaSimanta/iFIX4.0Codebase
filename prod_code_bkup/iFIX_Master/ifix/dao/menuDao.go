package dao

import (
	"database/sql"
	"errors"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
)

var menuinsert = "INSERT INTO dtlclientmenuinfo(clientid,mstorgnhirarchyid,moduleid,parentmenuid,menudesc,sequence_no,leafnode) VALUES (?,?,?,?,?,?,?)"
var menuduplicate = "SELECT count(id) total from dtlclientmenuinfo WHERE clientid=? AND mstorgnhirarchyid=? AND moduleid=? AND parentmenuid=? AND menudesc=?  AND deleteflg=0 AND activeflg=1"
var parentmenu = "SELECT id as  ID ,menudesc as Menudesc FROM dtlclientmenuinfo WHERE clientid=? AND mstorgnhirarchyid=? and moduleid =?  and leafnode=? and activeflg=1 AND deleteflg = 0 "
var menudetails = "select a.id,a.clientid,a.mstorgnhirarchyid,a.parentmenuid,a.menudesc,a.sequence_no,a.moduleid,a.activeflg as Activeflg,COALESCE((select menudesc from dtlclientmenuinfo e where a.parentmenuid=e.id),'Parent Menu') Parentmenu,b.name as Clientname,c.name as Orgnname,d.modulename as Modulename from dtlclientmenuinfo a,mstclient b,mstorgnhierarchy c,mstmodule d where a.clientid = b.id  and a.mstorgnhirarchyid=c.id and a.moduleid=d.id and a.activeflg=1 and a.deleteflg=0 and d.deleteflg=0 ORDER BY a.id DESC LIMIT ?,?;"
var menucount = "SELECT count(a.id) as total FROM dtlclientmenuinfo a,mstclient b,mstorgnhierarchy c,mstmodule d where a.clientid = b.id  and a.mstorgnhirarchyid=c.id and a.moduleid=d.id and a.activeflg=1 and a.deleteflg=0 and d.deleteflg=0"

var updatemenu = "UPDATE dtlclientmenuinfo SET menudesc=? where id=?"
var deletemenu = "UPDATE dtlclientmenuinfo SET deleteflg=1 where id=?"
var addurltomenu = "UPDATE dtlclientmenuinfo SET urlid=? ,leafnode=0 where id=?"
var getmenuurl = "select a.id,a.clientid,a.mstorgnhirarchyid,a.parentmenuid,a.menudesc,a.sequence_no,a.moduleid,a.activeflg as Activeflg,a.urlid as Urlid,b.name as Clientname,c.name as Orgnname,d.modulename as Modulename,e.url as Url from dtlclientmenuinfo a,mstclient b,mstorgnhierarchy c,mstmodule d,msturl e where a.clientid = b.id  and a.mstorgnhirarchyid=c.id and a.moduleid=d.id and a.urlid=e.id AND a.leafnode=0 and a.activeflg=1 and a.deleteflg=0 and d.deleteflg=0 and e.deleteflg=0 ORDER BY a.id DESC LIMIT ?,?;"
var urlmenucount = "SELECT count(a.id) as total FROM dtlclientmenuinfo a,mstclient b,mstorgnhierarchy c,mstmodule d,msturl e where a.clientid = b.id  and a.mstorgnhirarchyid=c.id and a.moduleid=d.id and a.urlid=e.id AND a.leafnode=0 and a.activeflg=1 and a.deleteflg=0 and d.deleteflg=0"
var menubymodule = "SELECT id as  ID ,menudesc as Menudesc FROM dtlclientmenuinfo WHERE clientid=? AND mstorgnhirarchyid=? and moduleid =?  AND activeflg =1 AND deleteflg = 0 "
var sqlMenuByUserID = "SELECT  distinct dtlclientmenuinfo.id,  dtlclientmenuinfo.menudesc label, dtlclientmenuinfo.parentmenuid parent, IF(msturl.url IS NULL, '', msturl.url) path FROM mapclientuserroleuser,  mstclientmoduleurlroleuser JOIN dtlclientmenuinfo ON mstclientmoduleurlroleuser.menuid = dtlclientmenuinfo.id LEFT JOIN msturl ON dtlclientmenuinfo.urlid = msturl.id  AND msturl.activeflg = 1 AND msturl.deleteflg = 0 WHERE mstclientmoduleurlroleuser.clientid = ? AND mstclientmoduleurlroleuser.mstorgnhirarchyid = ? AND mstclientmoduleurlroleuser.deleteflg = 0  AND mstclientmoduleurlroleuser.activeflg = 1 AND mstclientmoduleurlroleuser.userid = ? AND mapclientuserroleuser.roleid = mstclientmoduleurlroleuser.roleid AND mapclientuserroleuser.userid = mstclientmoduleurlroleuser.userid AND mapclientuserroleuser.clientid = ? AND mapclientuserroleuser.mstorgnhirarchyid = ? AND mapclientuserroleuser.deleteflg = 0 AND mapclientuserroleuser.activeflg = 1 AND dtlclientmenuinfo.clientid = ?  AND dtlclientmenuinfo.mstorgnhirarchyid = ? AND dtlclientmenuinfo.deleteflg = 0 AND dtlclientmenuinfo.activeflg = 1 ORDER BY dtlclientmenuinfo.parentmenuid ASC , dtlclientmenuinfo.sequence_no ASC , dtlclientmenuinfo.leafnode DESC"
var sqlMenuByUIDNRole = "SELECT `dtlclientmenuinfo`.id,`dtlclientmenuinfo`.`menudesc` label,`dtlclientmenuinfo`.`parentmenuid` parent,IF(`msturl`.`url` IS NULL,'',`msturl`.`url`) path FROM `mapclientuserroleuser` JOIN mstmodulerolemap ON mapclientuserroleuser.roleid=mstmodulerolemap.roleid JOIN `dtlclientmenuinfo` ON `mstmodulerolemap`.`menuid`=`dtlclientmenuinfo`.`id` LEFT  JOIN `msturl` ON dtlclientmenuinfo.urlid=msturl.id AND  `msturl`.`activeflg` = 1 AND `msturl`.`deleteflg` = 0  WHERE `mapclientuserroleuser`.`clientid` = ? AND `mapclientuserroleuser`.`mstorgnhirarchyid` = ? AND `mapclientuserroleuser`.`deleteflg` = 0 AND `mapclientuserroleuser`.`activeflg` = 1 AND `mapclientuserroleuser`.`userid`=? AND `dtlclientmenuinfo`.`clientid` = ? AND dtlclientmenuinfo.`mstorgnhirarchyid` = ? AND dtlclientmenuinfo.`deleteflg` = 0 AND dtlclientmenuinfo.`activeflg` = 1  ORDER BY `dtlclientmenuinfo`.`parentmenuid` ASC,`dtlclientmenuinfo`.`sequence_no` ASC,`dtlclientmenuinfo`.`leafnode` DESC"
var deleteurlfrommenu = "UPDATE dtlclientmenuinfo SET urlid=NULL ,leafnode=1 where id=?"
var searchmenubyuser = "SELECT `dtlclientmenuinfo`.`menudesc` label,`msturl`.`url` path  FROM `mstclientmoduleurlroleuser` JOIN `dtlclientmenuinfo` ON `mstclientmoduleurlroleuser`.`menuid`=`dtlclientmenuinfo`.`id` LEFT  JOIN `msturl` ON dtlclientmenuinfo.urlid=msturl.id AND  `msturl`.`activeflg` = 1 AND `msturl`.`deleteflg` = 0  WHERE `mstclientmoduleurlroleuser`.`clientid` = ? AND `mstclientmoduleurlroleuser`.`mstorgnhirarchyid` = ? AND `mstclientmoduleurlroleuser`.`deleteflg` = 0 AND `mstclientmoduleurlroleuser`.`activeflg` = 1 AND `mstclientmoduleurlroleuser`.`userid`= ? AND `dtlclientmenuinfo`.`clientid` = ? AND dtlclientmenuinfo.`mstorgnhirarchyid` = ? AND dtlclientmenuinfo.`leafnode`=0 and dtlclientmenuinfo.`menudesc` like ? AND dtlclientmenuinfo.`deleteflg` = 0 AND dtlclientmenuinfo.`activeflg` = 1 ;"
var searchmenubyrole = "SELECT `dtlclientmenuinfo`.`menudesc` label,`msturl`.`url` path FROM `mapclientuserroleuser` JOIN mstmodulerolemap ON mapclientuserroleuser.roleid=mstmodulerolemap.roleid JOIN `dtlclientmenuinfo` ON `mstmodulerolemap`.`menuid`=`dtlclientmenuinfo`.`id` LEFT  JOIN `msturl` ON dtlclientmenuinfo.urlid=msturl.id AND  `msturl`.`activeflg` = 1 AND `msturl`.`deleteflg` = 0  WHERE `mapclientuserroleuser`.`clientid` = ? AND `mapclientuserroleuser`.`mstorgnhirarchyid` = ? AND `mapclientuserroleuser`.`deleteflg` = 0 AND `mapclientuserroleuser`.`activeflg` = 1 AND `mapclientuserroleuser`.`userid`=? AND `dtlclientmenuinfo`.`clientid` = ? AND dtlclientmenuinfo.`mstorgnhirarchyid` = ? AND dtlclientmenuinfo.`leafnode`=0 and dtlclientmenuinfo.`menudesc` like ? AND dtlclientmenuinfo.`deleteflg` = 0 AND dtlclientmenuinfo.`activeflg` = 1 "
var menseqincrement = "update dtlclientmenuinfo set sequence_no=sequence_no+1 where clientid=? and mstorgnhirarchyid=? and parentmenuid=? and moduleid=? and sequence_no>=?"

/**
Delete url from a menu field
*/
func (mdao DbConn) DeleteUrlFromMenu(tz *entities.MenuEntity) error {
	logger.Log.Println("DeleteMenu -->", deleteurlfrommenu)
	logger.Log.Println("parameters -->", tz)
	stmt, err := mdao.DB.Prepare(deleteurlfrommenu)
	defer stmt.Close()
	if err != nil {
		logger.Log.Print("DeleteUrlFromMenu Prepare Statement  Error", err)
		log.Print("DeleteUrlFromMenu Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.ID)
	if err != nil {
		logger.Log.Print("DeleteUrlFromMenu Execute Statement  Error", err)
		log.Print("DeleteUrlFromMenu Execute Statement  Error", err)
		return err
	}
	return nil
}
func (mdao DbConn) GetMenuByUserID(request *entities.MenuByUserRequest) ([]entities.MenuHierarchyEntity, error) {
	logger.Log.Println("GetMenuByUserID -->", sqlMenuByUserID)
	//logger.Log.Println("parameters -->", request)
	values := []entities.MenuHierarchyEntity{}
	rows, err := mdao.DB.Query(sqlMenuByUserID, request.ClientID, request.MstorgnhirarchyID, request.UserID, request.ClientID, request.MstorgnhirarchyID, request.ClientID, request.MstorgnhirarchyID)
	if err != nil {
		log.Print("GetMenuByUserID Get Statement Prepare Error", err)
		logger.Log.Print("GetMenuByUserID Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		value := entities.MenuHierarchyEntity{}
		err = rows.Scan(&value.ID, &value.Label, &value.Parent, &value.Path)
		if err != nil {
			log.Print("GetMenuByUserID Scan Error", err)
			logger.Log.Print("GetMenuByUserID Scan Error", err)
			return values, err
		}
		values = append(values, value)
	}
	return values, nil
}
func (mdao DbConn) SearchMenuByUserID(request *entities.MenuByUserRequest) ([]entities.MenuHierarchyEntity, error) {
	//logger.Log.Println("SearchMenuByUserID -->", sqlMenuByUserID)
	//logger.Log.Println("parameters -->", request)
	values := []entities.MenuHierarchyEntity{}
	rows, err := mdao.DB.Query(searchmenubyuser, request.ClientID, request.MstorgnhirarchyID, request.UserID, request.ClientID, request.MstorgnhirarchyID, "%"+request.Menu+"%")
	if err != nil {
		log.Print("SearchMenuByUserID Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		value := entities.MenuHierarchyEntity{}
		err = rows.Scan(&value.Label, &value.Path)
		if err != nil {
			log.Print("SearchMenuByUserID Scan Error", err)
			return values, err
		}
		values = append(values, value)
	}
	return values, nil
}

//func (mdao DbConn) Getsinglemenubyuserid(request *entities.MenuByUserRequest) ([]entities.MenuHierarchyEntity, error) {
//	logger.Log.Println("Getsinglemenubyuserid -->", sqlMenuByUserID)
//	logger.Log.Println("parameters -->", request)
//	values := []entities.MenuHierarchyEntity{}
//	singlemenubyuser:="select a.id from mstclientmoduleurlroleuser a,dtlclientmenuinfo b,mapclientuserroleuser c,mstmodule d ,mstclientuser e,mstclientuserrole f where a.clientid=? and a.mstorgnhirarchyid=? and a.userid =? and a.menuid=b.id and a.moduleid=d.id and b.urlid =(SELECT id from msturl where url=? and deleteflg=0 and activeflg=1) and a.roleid =c.roleid and a.userid=e.id and a.roleid=f.id and c.userid=? and c.activeflg=1 and c.deleteflg=0 and d.activeflg=1 and d.deleteflg=0 and e.activeflag=1 and e.deleteflag=0 and f.activeflg=1 and f.deleteflg=0"
//	rows, err := mdao.DB.Query(singlemenubyuser, request.ClientID, request.MstorgnhirarchyID, request.UserID, request.Menu, request.UserID)
//	if err != nil {
//		log.Print("Getsinglemenubyuserid Get Statement Prepare Error", err)
//		return values, err
//	}
//	defer rows.Close()
//	for rows.Next() {
//		value := entities.MenuHierarchyEntity{}
//		err = rows.Scan( &value.Label, &value.Path)
//		if err != nil {
//			log.Print("Getsinglemenubyuserid Scan Error", err)
//			return values, err
//		}
//		values = append(values, value)
//	}
//	return values, nil
//}

func (mdao DbConn) GetMenuByUserIDNRole(request *entities.MenuByUserRequest) ([]entities.MenuHierarchyEntity, error) {
	//logger.Log.Println("GetMenuByUserIDNRole -->", sqlMenuByUIDNRole)
	//logger.Log.Println("parameters -->", request)
	values := []entities.MenuHierarchyEntity{}
	rows, err := mdao.DB.Query(sqlMenuByUIDNRole, request.ClientID, request.MstorgnhirarchyID, request.UserID, request.ClientID, request.MstorgnhirarchyID)
	defer rows.Close()
	if err != nil {
		log.Print("GetMenuByUserIDNRole Get Statement Prepare Error", err)
		logger.Log.Print("GetMenuByUserIDNRole Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MenuHierarchyEntity{}
		err = rows.Scan(&value.ID, &value.Label, &value.Parent, &value.Path)
		if err != nil {
			log.Print("GetMenuByUserIDNRole Scan Error", err)
			logger.Log.Print("GetMenuByUserIDNRole Scan Error", err)
			return values, err
		}
		values = append(values, value)
	}
	return values, nil
}
func (mdao DbConn) SearchMenuByUserIDNRole(request *entities.MenuByUserRequest) ([]entities.MenuHierarchyEntity, error) {
	//logger.Log.Println("SearchMenuByUserIDNRole -->", sqlMenuByUIDNRole)
	//logger.Log.Println("parameters -->", request)
	values := []entities.MenuHierarchyEntity{}
	rows, err := mdao.DB.Query(searchmenubyrole, request.ClientID, request.MstorgnhirarchyID, request.UserID, request.ClientID, request.MstorgnhirarchyID, "%"+request.Menu+"%")
	defer rows.Close()
	if err != nil {
		log.Print("SearchMenuByUserIDNRole Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MenuHierarchyEntity{}
		err = rows.Scan(&value.Label, &value.Path)
		if err != nil {
			log.Print("SearchMenuByUserIDNRole Scan Error", err)
			return values, err
		}
		values = append(values, value)
	}
	return values, nil
}

//func (mdao DbConn) Getsinglemenubyuseridrole(request *entities.MenuByUserRequest) ([]entities.MenuHierarchyEntity, error) {
//	//logger.Log.Println("SearchMenuByUserIDNRole -->", sqlMenuByUIDNRole)
//	logger.Log.Println("parameters -->", request)
//	values := []entities.MenuHierarchyEntity{}
//	rows, err := mdao.DB.Query(searchmenubyrole, request.ClientID, request.MstorgnhirarchyID, request.UserID, request.ClientID, request.MstorgnhirarchyID,"%"+request.Menu+"%")
//	defer rows.Close()
//	if err != nil {
//		log.Print("Getsinglemenubyuseridrole Get Statement Prepare Error", err)
//		return values, err
//	}
//	for rows.Next() {
//		value := entities.MenuHierarchyEntity{}
//		err = rows.Scan( &value.Label, &value.Path)
//		if err != nil {
//			log.Print("SearchMenuByUserIDNRole Scan Error", err)
//			return values, err
//		}
//		values = append(values, value)
//	}
//	return values, nil
//}

//InsertMenu function insert menu details
func (mdao TxConn) InsertMenu(tz *entities.MenuEntity) (int64, error) {
	logger.Log.Println("InsertMenu query -->", menuinsert)
	logger.Log.Println("parameters -->", tz)
	stmt, err := mdao.TX.Prepare(menuinsert)
	defer stmt.Close()
	if err != nil {
		log.Print("InsertMenu Prepare Statement  Error", err)
		logger.Log.Print("InsertMenu Prepare Statement  Error", err)
		return 0, err
	}
	res, err := stmt.Exec(tz.ClientID, tz.MstorgnhirarchyID, tz.Moduleid, tz.Parentmenuid, tz.Menudesc, tz.Sequence_no, 1)
	if err != nil {
		log.Print("InsertMenu Execute Statement  Error", err)
		logger.Log.Print("InsertMenu Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedID, err := res.LastInsertId()
	if err != nil {
		logger.Log.Println(err)
		return 0, errors.New("SQL Last ID fetch Error")
	}
	return lastInsertedID, nil
}

// will increment each seq by 1 where seq>tz.seq
func (mdao TxConn) IncrementMenuseq(tz *entities.MenuEntity) (int64, error) {
	logger.Log.Println("InsertMenu query -->", menuinsert)
	logger.Log.Println("parameters -->", tz)
	stmt, err := mdao.TX.Prepare(menseqincrement)
	defer stmt.Close()
	if err != nil {
		log.Print("InsertMenu Prepare Statement  Error", err)
		logger.Log.Print("InsertMenu Prepare Statement  Error", err)
		return 0, err
	}
	res, err := stmt.Exec(tz.ClientID, tz.MstorgnhirarchyID, tz.Parentmenuid, tz.Moduleid, tz.Sequence_no)
	if err != nil {
		log.Print("InsertMenu Execute Statement  Error", err)
		logger.Log.Print("InsertMenu Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedID, err := res.LastInsertId()
	if err != nil {
		logger.Log.Println(err)
		return 0, errors.New("SQL Last ID fetch Error")
	}
	return lastInsertedID, nil
}

//CheckDuplicateMenu check duplicate menu
func (mdao DbConn) Checkduplicatemenu(tz *entities.MenuEntity) (entities.MenuEntities, error) {
	logger.Log.Println("CheckDuplicateMenu Query -->", menuduplicate)
	logger.Log.Println("parameters -->", tz)
	value := entities.MenuEntities{}
	err := mdao.DB.QueryRow(menuduplicate, tz.ClientID, tz.MstorgnhirarchyID, tz.Moduleid, tz.Parentmenuid, tz.Menudesc).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		log.Println(err)
		logger.Log.Print("Checkduplicatemenu Get Statement Prepare Error", err)
		return value, err
	}
}

/**
return menu with no child menu
*/
func (mdao DbConn) Getparentmenu(tz *entities.MenuEntity) ([]entities.MenuSingleEntity, error) {
	logger.Log.Println("Getparentmenu -->", parentmenu)
	logger.Log.Println("parameters -->", tz)
	values := []entities.MenuSingleEntity{}
	rows, err := mdao.DB.Query(parentmenu, tz.ClientID, tz.MstorgnhirarchyID, tz.Moduleid, tz.Leafnode)
	defer rows.Close()
	if err != nil {
		log.Print("Getparentmenu Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MenuSingleEntity{}
		err = rows.Scan(&value.ID, &value.Menudesc)
		if err != nil {
			log.Print("Getparentmenu Scan Error", err)
			return values, err
		}
		values = append(values, value)
	}
	return values, nil
}

/**
return all menu by module id
*/
func (mdao DbConn) Getmenubymodule(tz *entities.MenuEntity) ([]entities.MenuSingleEntity, error) {
	logger.Log.Println("Getmenubymodule -->", parentmenu)
	logger.Log.Println("parameters -->", tz)
	values := []entities.MenuSingleEntity{}
	rows, err := mdao.DB.Query(menubymodule, tz.ClientID, tz.MstorgnhirarchyID, tz.Moduleid)
	defer rows.Close()
	if err != nil {
		log.Print("Getmenubymodule Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MenuSingleEntity{}
		err = rows.Scan(&value.ID, &value.Menudesc)
		if err != nil {
			log.Print("Getmenubymodule Scan Error", err)
			return values, err
		}
		values = append(values, value)
	}
	return values, nil
}

/**
return all menu details
*/
func (mdao DbConn) Getmenudetails(page *entities.PaginationEntity) ([]entities.MenuEntityResp, error) {
	logger.Log.Println("Getmenudetails -->", menudetails)
	logger.Log.Println("parameters -->", page)
	values := []entities.MenuEntityResp{}
	rows, err := mdao.DB.Query(menudetails, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		log.Print("Getmenudetails Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MenuEntityResp{}
		err = rows.Scan(&value.ID, &value.ClientID, &value.MstorgnhirarchyID, &value.Parentmenuid, &value.Menudesc, &value.Sequence_no, &value.Moduleid, &value.Activeflg, &value.Parentmenu, &value.Clientname, &value.Orgnname, &value.Modulename)
		if err != nil {
			log.Print("Getmenudetails Scan Error", err)
			return values, err
		}
		values = append(values, value)
	}
	return values, nil
}

/**
return all menu count
*/
func (dbc DbConn) Getmenucount() (entities.MenuEntities, error) {
	logger.Log.Println("In side GetMenuCount")
	value := entities.MenuEntities{}
	err := dbc.DB.QueryRow(menucount).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		log.Print("GetMenuCount Get Statement Prepare Error", err)
		return value, err
	}
}

/**
update menu field
*/
//changes done by josim on 16 mar 2022
//update query changes ...only menudesc can be updated
func (mdao DbConn) UpdateMenu(tz *entities.MenuEntity) error {
	logger.Log.Println("UpdateMenu -->", updatemenu)
	logger.Log.Println("parameters -->", tz)
	stmt, err := mdao.DB.Prepare(updatemenu)
	defer stmt.Close()
	if err != nil {
		log.Print("UpdateMenu Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Menudesc, tz.ID)
	if err != nil {
		log.Print("UpdateMenu Execute Statement  Error", err)
		return err
	}
	return nil
}

/**
add url to menu
*/
func (mdao DbConn) Addurltomenu(tz *entities.MenuEntity) error {
	logger.Log.Println("Addurltomenu -->", addurltomenu)
	logger.Log.Println("parameters -->", tz)
	stmt, err := mdao.DB.Prepare(addurltomenu)
	defer stmt.Close()
	if err != nil {
		log.Print("Addurltomenu Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Urlid, tz.ID)
	if err != nil {
		log.Print("Addurltomenu Execute Statement  Error", err)
		return err
	}
	return nil
}

/**
Delete menu field
*/
func (mdao DbConn) DeleteMenu(tz *entities.MenuEntity) error {
	logger.Log.Println("DeleteMenu -->", deletemenu)
	logger.Log.Println("parameters -->", tz)
	stmt, err := mdao.DB.Prepare(deletemenu)
	defer stmt.Close()
	if err != nil {
		logger.Log.Print("DeleteMenu Prepare Statement  Error", err)
		log.Print("DeleteMenu Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.ID)
	if err != nil {
		logger.Log.Print("DeleteMenu Execute Statement  Error", err)
		log.Print("DeleteMenu Execute Statement  Error", err)
		return err
	}
	return nil
}

/**
return all menu details which are the last menu/has the url with them
*/
func (mdao DbConn) Geturlmenudetails(page *entities.PaginationEntity) ([]entities.MenuEntityResp, error) {
	logger.Log.Println("Geturlmenudetails -->", menudetails)
	logger.Log.Println("parameters -->", page)
	values := []entities.MenuEntityResp{}
	rows, err := mdao.DB.Query(getmenuurl, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		log.Print("Geturlmenudetails Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MenuEntityResp{}
		err = rows.Scan(&value.ID, &value.ClientID, &value.MstorgnhirarchyID, &value.Parentmenuid, &value.Menudesc, &value.Sequence_no, &value.Moduleid, &value.Activeflg, &value.Urlid, &value.Clientname, &value.Orgnname, &value.Modulename, &value.Url)
		if err != nil {
			log.Print("Geturlmenudetails Scan Error", err)
			return values, err
		}
		values = append(values, value)
	}
	return values, nil
}

/**
return all menu (which is not leaf node) count
*/
func (dbc DbConn) Geturlmenucount() (entities.MenuEntities, error) {
	logger.Log.Println("In side Geturlmenucount")
	value := entities.MenuEntities{}
	err := dbc.DB.QueryRow(urlmenucount).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		log.Print("Geturlmenucount Get Statement Prepare Error", err)
		return value, err
	}
}
