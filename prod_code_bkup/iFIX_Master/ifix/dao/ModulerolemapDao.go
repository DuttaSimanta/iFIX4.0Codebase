package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
)

var insertModulerolemap = "INSERT INTO mstmodulerolemap (clientid, mstorgnhirarchyid, moduleid, roleid, menuid) VALUES (?,?,?,?,?)"
var duplicateModulerolemap = "SELECT count(id) total FROM  mstmodulerolemap WHERE clientid = ? AND mstorgnhirarchyid = ? AND moduleid=? AND roleid=? and menuid=?  AND deleteflg = 0 and activeflg=1"

//var getModulerolemap = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.moduleid as Moduleid, a.roleid as Roleid, a.menuid as Menuid,a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.rolename as Rolename,e.modulename as Modulename,f.menudesc as Menuname FROM mstmodulerolemap a,mstclient b,mstorgnhierarchy c,mstclientuserrole d,mstmodule e,dtlclientmenuinfo f WHERE a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.roleid=d.id AND a.moduleid=e.id AND a.menuid=f.id AND a.clientid=? AND a.mstorgnhirarchyid=? AND a.deleteflg =0 and a.activeflg=1 AND d.deleteflg =0 AND e.deleteflg =0 AND f.deleteflg=0 ORDER BY id DESC LIMIT ?,?"
//var getModulerolemapcount = "SELECT count(a.id) total FROM mstmodulerolemap a,mstclient b,mstorgnhierarchy c,mstclientuserrole d,mstmodule e,dtlclientmenuinfo f WHERE a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.roleid=d.id AND a.moduleid=e.id AND a.menuid=f.id AND a.clientid=? AND a.mstorgnhirarchyid=? AND a.deleteflg =0 and a.activeflg=1 AND d.deleteflg =0 AND e.deleteflg =0 AND f.deleteflg=0"
var updateModulerolemap = "UPDATE mstmodulerolemap SET clientid = ?, mstorgnhirarchyid = ?, moduleid = ?, roleid = ?, menuid = ? WHERE id = ? "
var deleteModulerolemap = "UPDATE mstmodulerolemap SET deleteflg = '1' WHERE id = ? "

//var getModulerolemap = "SELECT id as Id, clientid as Clientid, mstorgnhirarchyid as Mstorgnhirarchyid, moduleid as Moduleid, roleid as Roleid, menuid as Menuid FROM mstmodulerolemap WHERE clientid = ? AND mstorgnhirarchyid = ? AND deleteflg =0 and activeflg=1 ORDER BY id DESC LIMIT ?,?"
func (dbc DbConn) CheckDuplicateModulerolemap(tz *entities.ModulerolemapEntity) (entities.ModulerolemapEntities, error) {
	logger.Log.Println("In side CheckDuplicateModulerolemap")
	logger.Log.Println("Query -->", duplicateModulerolemap)
	logger.Log.Println("Parameter ---->", tz.Clientid, tz.Mstorgnhirarchyid)
	value := entities.ModulerolemapEntities{}
	err := dbc.DB.QueryRow(duplicateModulerolemap, tz.Clientid, tz.Mstorgnhirarchyid, tz.Moduleid, tz.Roleid, tz.Menuid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateModulerolemap Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertModulerolemap(tz *entities.ModulerolemapEntity) (int64, error) {
	logger.Log.Println("In side InsertModulerolemap")
	logger.Log.Println("Query -->", insertModulerolemap)
	stmt, err := dbc.DB.Prepare(insertModulerolemap)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertModulerolemap Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Moduleid, tz.Roleid, tz.Menuid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Moduleid, tz.Roleid, tz.Menuid)
	if err != nil {
		logger.Log.Println("InsertModulerolemap Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllModulerolemap(page *entities.ModulerolemapEntity, OrgnType int64) ([]entities.ModulerolemapEntity, error) {
	values := []entities.ModulerolemapEntity{}
	var getModulerolemap string
	var params []interface{}
	if OrgnType == 1 {
		getModulerolemap = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.moduleid as Moduleid, a.roleid as Roleid, a.menuid as Menuid,a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.rolename as Rolename,e.modulename as Modulename,f.menudesc as Menuname FROM mstmodulerolemap a,mstclient b,mstorgnhierarchy c,mstclientuserrole d,mstmodule e,dtlclientmenuinfo f WHERE a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.roleid=d.id AND a.moduleid=e.id AND a.menuid=f.id AND a.deleteflg =0 and a.activeflg=1 AND d.deleteflg =0 AND e.deleteflg =0 AND f.deleteflg=0 ORDER BY id DESC LIMIT ?,?"
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else if OrgnType == 2 {
		getModulerolemap = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.moduleid as Moduleid, a.roleid as Roleid, a.menuid as Menuid,a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.rolename as Rolename,e.modulename as Modulename,f.menudesc as Menuname FROM mstmodulerolemap a,mstclient b,mstorgnhierarchy c,mstclientuserrole d,mstmodule e,dtlclientmenuinfo f WHERE a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.roleid=d.id AND a.moduleid=e.id AND a.menuid=f.id AND a.clientid=? AND a.deleteflg =0 and a.activeflg=1 AND d.deleteflg =0 AND e.deleteflg =0 AND f.deleteflg=0 ORDER BY id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else {
		getModulerolemap = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.moduleid as Moduleid, a.roleid as Roleid, a.menuid as Menuid,a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.rolename as Rolename,e.modulename as Modulename,f.menudesc as Menuname FROM mstmodulerolemap a,mstclient b,mstorgnhierarchy c,mstclientuserrole d,mstmodule e,dtlclientmenuinfo f WHERE a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.roleid=d.id AND a.moduleid=e.id AND a.menuid=f.id AND a.clientid=? AND a.mstorgnhirarchyid=? AND a.deleteflg =0 and a.activeflg=1 AND d.deleteflg =0 AND e.deleteflg =0 AND f.deleteflg=0 ORDER BY id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Mstorgnhirarchyid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	}
	rows, err := dbc.DB.Query(getModulerolemap, params...)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllModulerolemap Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.ModulerolemapEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Moduleid, &value.Roleid, &value.Menuid, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Rolename, &value.Modulename, &value.Menuname)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateModulerolemap(tz *entities.ModulerolemapEntity) error {
	logger.Log.Println("In side UpdateModulerolemap")
	logger.Log.Println("Query -->", updateModulerolemap)
	logger.Log.Println("Parameter ---->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Moduleid, tz.Roleid, tz.Menuid, tz.Id)

	stmt, err := dbc.DB.Prepare(updateModulerolemap)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateModulerolemap Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Moduleid, tz.Roleid, tz.Menuid, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateModulerolemap Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteModulerolemap(tz *entities.ModulerolemapEntity) error {
	logger.Log.Println("In side DeleteModulerolemap")
	logger.Log.Println("Query -->", deleteModulerolemap)
	logger.Log.Println("Parameter ---->", tz.Id)
	stmt, err := dbc.DB.Prepare(deleteModulerolemap)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteModulerolemap Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteModulerolemap Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetModulerolemapCount(tz *entities.ModulerolemapEntity, OrgnTypeID int64) (entities.ModulerolemapEntities, error) {
	value := entities.ModulerolemapEntities{}
	var getModulerolemapcount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getModulerolemapcount = "SELECT count(a.id) total FROM mstmodulerolemap a,mstclient b,mstorgnhierarchy c,mstclientuserrole d,mstmodule e,dtlclientmenuinfo f WHERE a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.roleid=d.id AND a.moduleid=e.id AND a.menuid=f.id AND a.deleteflg =0 and a.activeflg=1 AND d.deleteflg =0 AND e.deleteflg =0 AND f.deleteflg=0"
	} else if OrgnTypeID == 2 {
		getModulerolemapcount = "SELECT count(a.id) total FROM mstmodulerolemap a,mstclient b,mstorgnhierarchy c,mstclientuserrole d,mstmodule e,dtlclientmenuinfo f WHERE a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.roleid=d.id AND a.moduleid=e.id AND a.menuid=f.id AND a.clientid=? AND a.deleteflg =0 and a.activeflg=1 AND d.deleteflg =0 AND e.deleteflg =0 AND f.deleteflg=0"
		params = append(params, tz.Clientid)
	} else {
		getModulerolemapcount = "SELECT count(a.id) total FROM mstmodulerolemap a,mstclient b,mstorgnhierarchy c,mstclientuserrole d,mstmodule e,dtlclientmenuinfo f WHERE a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.roleid=d.id AND a.moduleid=e.id AND a.menuid=f.id AND a.clientid=? AND a.mstorgnhirarchyid=? AND a.deleteflg =0 and a.activeflg=1 AND d.deleteflg =0 AND e.deleteflg =0 AND f.deleteflg=0"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getModulerolemapcount, params...).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		log.Print("GetModulerolemapCount Get Statement Prepare Error", err)
		return value, err
	}
}
