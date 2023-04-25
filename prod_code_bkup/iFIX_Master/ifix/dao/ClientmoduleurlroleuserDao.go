package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
)

var insertClientmoduleurlroleuser = "INSERT INTO mstclientmoduleurlroleuser (clientid, mstorgnhirarchyid, moduleid, roleid, menuid, userid) VALUES (?,?,?,?,?,?)"
var duplicateClientmoduleurlroleuser = "SELECT count(id) total FROM  mstclientmoduleurlroleuser WHERE clientid = ? AND mstorgnhirarchyid = ? AND moduleid = ? AND roleid = ? AND menuid = ? AND userid = ? AND deleteflg = 0 and activeflg=1"

//var getClientmoduleurlroleuser = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.moduleid as Moduleid, a.roleid as Roleid, a.menuid as Menuid,a.userid as Userid,a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.rolename as Rolename,e.modulename as Modulename,f.menudesc as Menuname,g.name as Username FROM mstclientmoduleurlroleuser a,mstclient b,mstorgnhierarchy c,mstclientuserrole d,mstmodule e,dtlclientmenuinfo f,mstclientuser g WHERE a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.roleid=d.id AND a.moduleid=e.id AND a.menuid=f.id AND a.userid=g.id AND a.deleteflg =0 and a.activeflg=1 ORDER BY id DESC LIMIT ?,?"

//var getClientmoduleurlroleusercount = "SELECT count(id) total FROM  mstclientmoduleurlroleuser WH deleteflg =0 and activeflg=1"
var updateClientmoduleurlroleuser = "UPDATE mstclientmoduleurlroleuser SET clientid = ?, mstorgnhirarchyid = ?, moduleid = ?, roleid = ?, menuid = ?, userid = ? WHERE id = ? "
var deleteClientmoduleurlroleuser = "UPDATE mstclientmoduleurlroleuser SET deleteflg = 1 WHERE id = ? "

//var getClientmoduleurlroleusercount = "SELECT count(a.id) as total FROM mstclientmoduleurlroleuser a,mstclient b,mstorgnhierarchy c,mstclientuserrole d,mstmodule e,dtlclientmenuinfo f,mstclientuser g WHERE a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.roleid=d.id AND a.moduleid=e.id AND a.menuid=f.id AND a.userid=g.id AND a.deleteflg =0 and a.activeflg=1"

func (dbc DbConn) CheckDuplicateClientmoduleurlroleuser(tz *entities.ClientmoduleurlroleuserEntity, i int) (entities.ClientmoduleurlroleuserEntities, error) {
	logger.Log.Println("In side CheckDuplicateClientmoduleurlroleuser")
	value := entities.ClientmoduleurlroleuserEntities{}
	err := dbc.DB.QueryRow(duplicateClientmoduleurlroleuser, tz.Clientid, tz.Mstorgnhirarchyid, tz.Moduleid, tz.Roleid, tz.Menuids[i], tz.Refuserid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateClientmoduleurlroleuser Get Statement Prepare Error", err)
		return value, err
	}
}
func (dbc DbConn) CheckDuplicateClientmoduleurlroleuserforupdate(tz *entities.ClientmoduleurlroleuserEntity) (entities.ClientmoduleurlroleuserEntities, error) {
	logger.Log.Println("In side CheckDuplicateClientmoduleurlroleuser")
	value := entities.ClientmoduleurlroleuserEntities{}
	err := dbc.DB.QueryRow(duplicateClientmoduleurlroleuser, tz.Clientid, tz.Mstorgnhirarchyid, tz.Moduleid, tz.Roleid, tz.Menuid, tz.Refuserid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateClientmoduleurlroleuser Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc TxConn) InsertClientmoduleurlroleuser(tz *entities.ClientmoduleurlroleuserEntity, i int) (int64, error) {
	logger.Log.Println("In side InsertClientmoduleurlroleuser")
	logger.Log.Println("Query -->", insertClientmoduleurlroleuser)
	stmt, err := dbc.TX.Prepare(insertClientmoduleurlroleuser)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertClientmoduleurlroleuser Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Moduleid, tz.Roleid, tz.Menuids[i], tz.Refuserid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Moduleid, tz.Roleid, tz.Menuids[i], tz.Refuserid)
	if err != nil {
		logger.Log.Println("InsertClientmoduleurlroleuser Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllClientmoduleurlroleuser(tz *entities.ClientmoduleurlroleuserEntity, OrgnType int64) ([]entities.ClientmoduleurlroleuserEntity, error) {
	logger.Log.Println("In side GelAllClientmoduleurlroleuser")
	var getClientmoduleurlroleuser string
	var params []interface{}
	values := []entities.ClientmoduleurlroleuserEntity{}
	if OrgnType == 1 {
		getClientmoduleurlroleuser = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.moduleid as Moduleid, a.roleid as Roleid, a.menuid as Menuid,a.userid as Userid,a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.rolename as Rolename,e.modulename as Modulename,f.menudesc as Menuname,g.name as Username FROM mstclientmoduleurlroleuser a,mstclient b,mstorgnhierarchy c,mstclientuserrole d,mstmodule e,dtlclientmenuinfo f,mstclientuser g WHERE a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.roleid=d.id AND a.moduleid=e.id AND a.menuid=f.id AND a.userid=g.id AND a.deleteflg =0 and a.activeflg=1 ORDER BY id DESC LIMIT ?,?"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		getClientmoduleurlroleuser = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.moduleid as Moduleid, a.roleid as Roleid, a.menuid as Menuid,a.userid as Userid,a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.rolename as Rolename,e.modulename as Modulename,f.menudesc as Menuname,g.name as Username FROM mstclientmoduleurlroleuser a,mstclient b,mstorgnhierarchy c,mstclientuserrole d,mstmodule e,dtlclientmenuinfo f,mstclientuser g WHERE a.clientid =? AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.roleid=d.id AND a.moduleid=e.id AND a.menuid=f.id AND a.userid=g.id AND a.deleteflg =0 and a.activeflg=1 ORDER BY id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		getClientmoduleurlroleuser = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.moduleid as Moduleid, a.roleid as Roleid, a.menuid as Menuid,a.userid as Userid,a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.rolename as Rolename,e.modulename as Modulename,f.menudesc as Menuname,g.name as Username FROM mstclientmoduleurlroleuser a,mstclient b,mstorgnhierarchy c,mstclientuserrole d,mstmodule e,dtlclientmenuinfo f,mstclientuser g WHERE a.clientid =? AND a.mstorgnhirarchyid =? AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.roleid=d.id AND a.moduleid=e.id AND a.menuid=f.id AND a.userid=g.id AND a.deleteflg =0 and a.activeflg=1 ORDER BY id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}
	rows, err := dbc.DB.Query(getClientmoduleurlroleuser, params...)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllClientmoduleurlroleuser Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.ClientmoduleurlroleuserEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Moduleid, &value.Roleid, &value.Menuid, &value.Refuserid, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Rolename, &value.Modulename, &value.Menuname, &value.Refusername)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateClientmoduleurlroleuser(tz *entities.ClientmoduleurlroleuserEntity) error {
	logger.Log.Println("In side UpdateClientmoduleurlroleuser")
	stmt, err := dbc.DB.Prepare(updateClientmoduleurlroleuser)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateClientmoduleurlroleuser Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Moduleid, tz.Roleid, tz.Menuid, tz.Refuserid, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateClientmoduleurlroleuser Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteClientmoduleurlroleuser(tz *entities.ClientmoduleurlroleuserEntity) error {
	logger.Log.Println("In side DeleteClientmoduleurlroleuser")
	stmt, err := dbc.DB.Prepare(deleteClientmoduleurlroleuser)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteClientmoduleurlroleuser Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteClientmoduleurlroleuser Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetClientmoduleurlroleuserCount(tz *entities.ClientmoduleurlroleuserEntity, OrgnTypeID int64) (entities.ClientmoduleurlroleuserEntities, error) {
	logger.Log.Println("In side GetClientmoduleurlroleuserCount")
	value := entities.ClientmoduleurlroleuserEntities{}
	var getClientmoduleurlroleusercount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getClientmoduleurlroleusercount = "SELECT count(a.id) as total FROM mstclientmoduleurlroleuser a,mstclient b,mstorgnhierarchy c,mstclientuserrole d,mstmodule e,dtlclientmenuinfo f,mstclientuser g WHERE a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.roleid=d.id AND a.moduleid=e.id AND a.menuid=f.id AND a.userid=g.id AND a.deleteflg =0 and a.activeflg=1"
	} else if OrgnTypeID == 2 {
		getClientmoduleurlroleusercount = "SELECT count(a.id) as total FROM mstclientmoduleurlroleuser a,mstclient b,mstorgnhierarchy c,mstclientuserrole d,mstmodule e,dtlclientmenuinfo f,mstclientuser g WHERE a.clientid =? AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.roleid=d.id AND a.moduleid=e.id AND a.menuid=f.id AND a.userid=g.id AND a.deleteflg =0 and a.activeflg=1"
		params = append(params, tz.Clientid)
	} else {
		getClientmoduleurlroleusercount = "SELECT count(a.id) as total FROM mstclientmoduleurlroleuser a,mstclient b,mstorgnhierarchy c,mstclientuserrole d,mstmodule e,dtlclientmenuinfo f,mstclientuser g WHERE a.clientid =? AND a.mstorgnhirarchyid=?  AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.roleid=d.id AND a.moduleid=e.id AND a.menuid=f.id AND a.userid=g.id AND a.deleteflg =0 and a.activeflg=1"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getClientmoduleurlroleusercount, params...).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		log.Print("GetClientmoduleurlroleuserCount Get Statement Prepare Error", err)
		return value, err
	}
}
