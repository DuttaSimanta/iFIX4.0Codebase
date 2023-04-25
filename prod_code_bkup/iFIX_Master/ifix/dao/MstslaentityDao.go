package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertMstslaentity = "INSERT INTO mstslaentity (clientid, mstorgnhirarchyid, slaid, associatedmstclienttableid, associatedmstclienttablefieldid) VALUES (?,?,?,?,?)"
var duplicateMstslaentity = "SELECT count(id) total FROM  mstslaentity WHERE clientid = ? AND mstorgnhirarchyid = ? AND slaid = ? AND associatedmstclienttableid = ? AND associatedmstclienttablefieldid = ? AND deleteflg = 0"
var getMstslaentity = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.slaid as Slaid, a.associatedmstclienttableid as Associatedmstclienttableid, a.associatedmstclienttablefieldid as Associatedmstclienttablefieldid, a.activeflg as Activeflg,(select name from mstclient where id = a.clientid ) as Clientname,(select name from mstorgnhierarchy where id = a.mstorgnhirarchyid ) as Mstorgnhirarchyname,(select columnname from mstdatadictionaryfield where id = a.associatedmstclienttableid) as Fieldname,(select tablename from mstdatadictionarytable where id =a.associatedmstclienttableid ) as Tablename,(select mstdatadictionarydbid from mstdatadictionarytable where id =a.associatedmstclienttableid ) as Dbid,(select slaname from mstclientsla where id=a.slaid and deleteflg =0 and activeflg=1) as Slaname FROM mstslaentity a WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 ORDER BY a.id DESC LIMIT ?,?"
var getMstslaentitycount = "SELECT count(a.id) as total FROM mstslaentity a,mstclient b,mstorgnhierarchy c,mstdatadictionaryfield d,mstdatadictionarytable e WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.deleteflg =0 and a.activeflg=1 and b.id = a.clientid and c.id = a.mstorgnhirarchyid and d.id = a.associatedmstclienttablefieldid and e.id =a.associatedmstclienttableid"
var updateMstslaentity = "UPDATE mstslaentity SET mstorgnhirarchyid = ?, slaid = ?, associatedmstclienttableid = ?, associatedmstclienttablefieldid = ? WHERE id = ? "
var deleteMstslaentity = "UPDATE mstslaentity SET deleteflg = '1' WHERE id = ? "

func (dbc DbConn) CheckDuplicateMstslaentity(tz *entities.MstslaentityEntity) (entities.MstslaentityEntities, error) {
	logger.Log.Println("In side CheckDuplicateMstslaentity")
	value := entities.MstslaentityEntities{}
	err := dbc.DB.QueryRow(duplicateMstslaentity, tz.Clientid, tz.Mstorgnhirarchyid, tz.Slaid, tz.Associatedmstclienttableid, tz.Associatedmstclienttablefieldid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMstslaentity Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertMstslaentity(tz *entities.MstslaentityEntity) (int64, error) {
	logger.Log.Println("In side InsertMstslaentity")
	logger.Log.Println("Query -->", insertMstslaentity)
	stmt, err := dbc.DB.Prepare(insertMstslaentity)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMstslaentity Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Slaid, tz.Associatedmstclienttableid, tz.Associatedmstclienttablefieldid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Slaid, tz.Associatedmstclienttableid, tz.Associatedmstclienttablefieldid)
	if err != nil {
		logger.Log.Println("InsertMstslaentity Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllMstslaentity(page *entities.MstslaentityEntity) ([]entities.MstslaentityEntity, error) {
	logger.Log.Println("In side GelAllMstslaentity")
	values := []entities.MstslaentityEntity{}
	rows, err := dbc.DB.Query(getMstslaentity, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstslaentity Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstslaentityEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Slaid, &value.Associatedmstclienttableid, &value.Associatedmstclienttablefieldid, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Fieldname, &value.Tablename, &value.Dbid, &value.Slaname)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateMstslaentity(tz *entities.MstslaentityEntity) error {
	logger.Log.Println("In side UpdateMstslaentity")
	stmt, err := dbc.DB.Prepare(updateMstslaentity)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMstslaentity Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Slaid, tz.Associatedmstclienttableid, tz.Associatedmstclienttablefieldid, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMstslaentity Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMstslaentity(tz *entities.MstslaentityEntity) error {
	logger.Log.Println("In side DeleteMstslaentity")
	stmt, err := dbc.DB.Prepare(deleteMstslaentity)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMstslaentity Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMstslaentity Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMstslaentityCount(tz *entities.MstslaentityEntity) (entities.MstslaentityEntities, error) {
	logger.Log.Println("In side GetMstslaentityCount")
	value := entities.MstslaentityEntities{}
	err := dbc.DB.QueryRow(getMstslaentitycount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMstslaentityCount Get Statement Prepare Error", err)
		return value, err
	}
}
