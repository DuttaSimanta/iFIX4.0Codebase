package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var duplicateRoleProperty = "SELECT count(id) total FROM  mapuserroleproperty WHERE clientid = ? AND mstorgnhirarchyid = ? AND roleid = ? AND propertyid = ? AND deleteflg = 0 AND activeflg=1"
var insertRoleProperty = "INSERT INTO mapuserroleproperty (clientid, mstorgnhirarchyid, roleid, propertyid) VALUES (?,?,?,?)"
var updateRoleProperty = "UPDATE mapuserroleproperty SET clientid = ?,mstorgnhirarchyid = ?, roleid = ?, propertyid = ? WHERE id = ? "
var deleteUserPropertyName = "UPDATE mapuserroleproperty SET deleteflg = '1' WHERE id = ? "

func (dbc DbConn) CheckDuplicateRoleProperty(tz *entities.MapUserRolePropertyEntity, i int) (entities.MapUserRolePropertyEntities, error) {
	logger.Log.Println("In side CheckDuplicateRoleProperty")
	value := entities.MapUserRolePropertyEntities{}
	logger.Log.Println(tz.Clientid, tz.Mstorgnhirarchyid, tz.Roleid[i], tz.Propertyid)
	err := dbc.DB.QueryRow(duplicateRoleProperty, tz.Clientid, tz.Mstorgnhirarchyid, tz.Roleid[i], tz.Propertyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateRoleProperty Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc TxConn) InsertRoleProperty(tz *entities.MapUserRolePropertyEntity, i int) (int64, error) {
	logger.Log.Println("In side InsertRoleProperty")
	logger.Log.Println("Query -->", insertRoleProperty)
	stmt, err := dbc.TX.Prepare(insertRoleProperty)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertRoleProperty Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Roleid[i], tz.Propertyid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Roleid[i], tz.Propertyid)
	if err != nil {
		logger.Log.Println("InsertBInsertRolePropertyanner Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllUserRoleProperty(tz *entities.MapUserRolePropertyEntity, OrgnType int64) ([]entities.MapUserRolePropertyEntity, error) {
	logger.Log.Println("In side dao GetAllUserRoleProperty")
	values := []entities.MapUserRolePropertyEntity{}
	var getUserRoleProperty string
	var params []interface{}
	if OrgnType == 1 {
		getUserRoleProperty = "SELECT a.id, a.clientid, a.mstorgnhirarchyid, a.roleid, a.propertyid,b.name as mstorgnhirarchyname,c.rolename as rolename, d.propertyname as propertyname FROM mapuserroleproperty a ,mstorgnhierarchy b,mstclientuserrole c, mstuserproperty d WHERE a.mstorgnhirarchyid = b.id and a.roleid = c.id and a.propertyid = d.id and a.activeflg=1 and a.deleteflg =0  and c.activeflg=1 and c.deleteflg=0 ORDER BY a.id DESC LIMIT ?,?;"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		getUserRoleProperty = "SELECT a.id, a.clientid, a.mstorgnhirarchyid, a.roleid, a.propertyid, b.name as mstorgnhirarchyname, c.rolename as rolename,d.propertyname as propertyname FROM mapuserroleproperty a, mstorgnhierarchy b, mstclientuserrole c, mstuserproperty d WHERE a.clientid=? and a.mstorgnhirarchyid = b.id and a.roleid = c.id and a.propertyid = d.id and a.activeflg=1 and a.deleteflg =0 ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		getUserRoleProperty = "SELECT a.id, a.clientid, a.mstorgnhirarchyid, a.roleid, a.propertyid, b.name as mstorgnhirarchyname, c.rolename as rolename,d.propertyname as propertyname FROM mapuserroleproperty a, mstorgnhierarchy b, mstclientuserrole c, mstuserproperty d WHERE a.clientid=? and a.mstorgnhirarchyid=? and a.mstorgnhirarchyid = b.id and a.roleid = c.id and a.propertyid = d.id and a.activeflg=1 and a.deleteflg =0  and c.activeflg=1 and c.deleteflg=0 ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}
	rows, err := dbc.DB.Query(getUserRoleProperty, params...)
	// rows, err := dbc.DB.Query(getBanner, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllUserRoleProperty Get Statement Prepare Error", err)
		return values, err
	}
	var roleid int64
	for rows.Next() {
		value := entities.MapUserRolePropertyEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &roleid, &value.Propertyid, &value.Mstorgnhirarchyname, &value.Rolename, &value.Propertyname)
		logger.Log.Println(roleid)
		value.Roleid = append(value.Roleid, roleid)
		logger.Log.Println(value.Roleid)
		values = append(values, value)
	}
	logger.Log.Println(values)
	return values, nil
}

func (dbc DbConn) GetUserRolePropertyCount(tz *entities.MapUserRolePropertyEntity, OrgnTypeID int64) (entities.MapUserRolePropertyEntities, error) {
	logger.Log.Println("In side GetUserRoleProperty")
	value := entities.MapUserRolePropertyEntities{}
	var getUserRoleProperty string
	var params []interface{}
	if OrgnTypeID == 1 {
		getUserRoleProperty = "SELECT count(a.id) as total FROM mapuserroleproperty a ,mstorgnhierarchy b,mstclientuserrole c, mstuserproperty d WHERE  a.mstorgnhirarchyid = b.id and a.roleid = c.id and a.propertyid = d.id and a.activeflg=1 and a.deleteflg =0  and c.activeflg=1 and c.deleteflg=0"
	} else if OrgnTypeID == 2 {
		getUserRoleProperty = "SELECT count(a.id) as total FROM mapuserroleproperty a ,mstorgnhierarchy b,mstclientuserrole c, mstuserproperty d  WHERE a.clientid=? and a.mstorgnhirarchyid = b.id and a.roleid = c.id and a.propertyid = d.id and a.activeflg=1 and a.deleteflg =0  and c.activeflg=1 and c.deleteflg=0;"
		params = append(params, tz.Clientid)
	} else {
		getUserRoleProperty = "SELECT count(a.id) as total FROM mapuserroleproperty a ,mstorgnhierarchy b,mstclientuserrole c, mstuserproperty d WHERE a.clientid=? and a.mstorgnhirarchyid=? and a.mstorgnhirarchyid = b.id and a.roleid = c.id and a.propertyid = d.id and a.activeflg=1 and a.deleteflg =0  and c.activeflg=1 and c.deleteflg=0;"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getUserRoleProperty, params...).Scan(&value.Total)
	// err := dbc.DB.QueryRow(getUserRoleProperty, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetUserRoleProperty Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) GetUserPropertyName(page *entities.MapUserRolePropertyEntity) ([]entities.GetUserPropertyNameEntity, error) {
	logger.Log.Println("In side GetUserPropertyName")
	// logger.Log.Println(getRecorddifferentiation)
	values := []entities.GetUserPropertyNameEntity{}
	var Getallpropertyname = "select id, propertyname from mstuserproperty where activeflg = 1  and deleteflg=0"
	rows, err := dbc.DB.Query(Getallpropertyname)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetUserPropertyName Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.GetUserPropertyNameEntity{}
		rows.Scan(&value.Id, &value.Propertyname)
		values = append(values, value)
		logger.Log.Println(values)
	}
	return values, nil
}

func (dbc DbConn) UpdateUserPropertyName(tz *entities.MapUserRolePropertyEntity) error {
	logger.Log.Println("In side UpdateUserPropertyName")
	stmt, err := dbc.DB.Prepare(updateRoleProperty)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateUserPropertyName Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Roleid[0], tz.Propertyid, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateUserPropertyName Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteUserPropertyName(tz *entities.MapUserRolePropertyEntity) error {
	logger.Log.Println("In side DeleteUserPropertyName")
	stmt, err := dbc.DB.Prepare(deleteUserPropertyName)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteUserPropertyName Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteUserPropertyName Execute Statement  Error", err)
		return err
	}
	return nil
}
