package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertMstslafullfillmentcriteria = "INSERT INTO mstslafullfillmentcriteria (clientid, mstorgnhirarchyid, slaid, mstrecorddifferentiationtickettypeid, mstrecorddifferentiationpriorityid, mstrecorddifferentiationworkingcatid, responsetimeinhr, responsetimeinmin, responsetimeinsec, resolutiontimeinhr, resolutiontimeinmin, resolutiontimeinsec, responsecompliance, resolutioncompliance, supportgroupspecific) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
var duplicateMstslafullfillmentcriteria = "SELECT count(id) total FROM  mstslafullfillmentcriteria WHERE clientid = ? AND mstorgnhirarchyid = ? AND slaid = ? AND mstrecorddifferentiationtickettypeid = ? AND mstrecorddifferentiationpriorityid = ? AND mstrecorddifferentiationworkingcatid = ? AND responsetimeinhr = ? AND responsetimeinmin = ? AND responsetimeinsec = ? AND resolutiontimeinhr = ? AND resolutiontimeinmin = ? AND resolutiontimeinsec = ? AND responsecompliance = ? AND resolutioncompliance = ? AND supportgroupspecific = ? AND deleteflg = 0"

// var getMstslafullfillmentcriteria = "SELECT id as Id, clientid as Clientid, mstorgnhirarchyid as Mstorgnhirarchyid, slaid as Slaid, mstrecorddifferentiationtickettypeid as Mstrecorddifferentiationtickettypeid, mstrecorddifferentiationpriorityid as Mstrecorddifferentiationpriorityid, mstrecorddifferentiationworkingcatid as Mstrecorddifferentiationworkingcatid, responsetimeinhr as Responsetimeinhr, responsetimeinmin as Responsetimeinmin, responsetimeinsec as Responsetimeinsec, resolutiontimeinhr as Resolutiontimeinhr, resolutiontimeinmin as Resolutiontimeinmin, resolutiontimeinsec as Resolutiontimeinsec, supportgroupspecific as Supportgroupspecific, activeflg as Activeflg,(select name from mstclient where id = clientid ) as Clientname,(select name from mstorgnhierarchy where id = mstorgnhirarchyid ) as Mstorgnhirarchyname,(select name from mstrecorddifferentiation where id=mstrecorddifferentiationtickettypeid and deleteflg =0 and activeflg=1) as Tickettypename,(select name from mstrecorddifferentiation where id=mstrecorddifferentiationpriorityid and deleteflg =0 and activeflg=1) as Priorityname,COALESCE((select name from mstrecorddifferentiation where id=mstrecorddifferentiationworkingcatid and deleteflg =0 and activeflg=1),'') as Workingcatname,(select slaname from mstclientsla where id=slaid and deleteflg =0 and activeflg=1) as Slaname FROM mstslafullfillmentcriteria WHERE clientid = ? AND mstorgnhirarchyid = ?  AND deleteflg =0 and activeflg=1 ORDER BY id DESC LIMIT ?,?"
// var getMstslafullfillmentcriteriacount = "SELECT count(a.id) as total FROM mstslafullfillmentcriteria a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiation d,mstrecorddifferentiation e,mstclientsla g WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.deleteflg =0 and a.activeflg=1 and b.id = a.clientid and c.id = a.mstorgnhirarchyid and d.id = a.mstrecorddifferentiationtickettypeid and e.id =a.mstrecorddifferentiationpriorityid and g.id=a.slaid and d.deleteflg =0 and d.activeflg=1 and e.deleteflg =0 and e.activeflg=1 and g.deleteflg =0 and g.activeflg=1"
var updateMstslafullfillmentcriteria = "UPDATE mstslafullfillmentcriteria SET mstorgnhirarchyid = ?, slaid = ?, mstrecorddifferentiationtickettypeid = ?, mstrecorddifferentiationpriorityid = ?, mstrecorddifferentiationworkingcatid = ?, responsetimeinhr = ?, responsetimeinmin = ?, responsetimeinsec = ?, resolutiontimeinhr = ?, resolutiontimeinmin = ?, resolutiontimeinsec = ?, responsecompliance = ?, resolutioncompliance = ?, supportgroupspecific = ? WHERE id = ? "
var deleteMstslafullfillmentcriteria = "UPDATE mstslafullfillmentcriteria SET deleteflg = '1' WHERE id = ? "
var duplicateMstslafullfillmentcriteriawithoutcat = "SELECT count(id) total FROM  mstslafullfillmentcriteria WHERE clientid = ? AND mstorgnhirarchyid = ? AND slaid = ? AND mstrecorddifferentiationtickettypeid = ? AND mstrecorddifferentiationpriorityid = ? AND responsetimeinhr = ? AND responsetimeinmin = ? AND responsetimeinsec = ? AND resolutiontimeinhr = ? AND resolutiontimeinmin = ? AND resolutiontimeinsec = ? AND responsecompliance = ? AND resolutioncompliance = ? AND supportgroupspecific = ? AND deleteflg = 0"

func (dbc DbConn) CheckDuplicateMstslafullfillmentcriteria(tz *entities.MstslafullfillmentcriteriaEntity) (entities.MstslafullfillmentcriteriaEntities, error) {
	logger.Log.Println("In side CheckDuplicateMstslafullfillmentcriteria")
	value := entities.MstslafullfillmentcriteriaEntities{}
	err := dbc.DB.QueryRow(duplicateMstslafullfillmentcriteria, tz.Clientid, tz.Mstorgnhirarchyid, tz.Slaid, tz.Mstrecorddifferentiationtickettypeid, tz.Mstrecorddifferentiationpriorityid, tz.Mstrecorddifferentiationworkingcatid, tz.Responsetimeinhr, tz.Responsetimeinmin, tz.Responsetimeinsec, tz.Resolutiontimeinhr, tz.Resolutiontimeinmin, tz.Resolutiontimeinsec, tz.Responsecompliance, tz.Resolutioncompliance, tz.Supportgroupspecific).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMstslafullfillmentcriteria Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) CheckDuplicateMstslafullfillmentcriteriawithoutcat(tz *entities.MstslafullfillmentcriteriaEntity) (entities.MstslafullfillmentcriteriaEntities, error) {
	logger.Log.Println("In side CheckDuplicateMstslafullfillmentcriteria")
	value := entities.MstslafullfillmentcriteriaEntities{}
	err := dbc.DB.QueryRow(duplicateMstslafullfillmentcriteriawithoutcat, tz.Clientid, tz.Mstorgnhirarchyid, tz.Slaid, tz.Mstrecorddifferentiationtickettypeid, tz.Mstrecorddifferentiationpriorityid, tz.Responsetimeinhr, tz.Responsetimeinmin, tz.Responsetimeinsec, tz.Resolutiontimeinhr, tz.Resolutiontimeinmin, tz.Resolutiontimeinsec, tz.Responsecompliance, tz.Resolutioncompliance, tz.Supportgroupspecific).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMstslafullfillmentcriteria Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertMstslafullfillmentcriteria(tz *entities.MstslafullfillmentcriteriaEntity) (int64, error) {
	logger.Log.Println("In side InsertMstslafullfillmentcriteria")
	logger.Log.Println("Query -->", insertMstslafullfillmentcriteria)
	stmt, err := dbc.DB.Prepare(insertMstslafullfillmentcriteria)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMstslafullfillmentcriteria Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Slaid, tz.Mstrecorddifferentiationtickettypeid, tz.Mstrecorddifferentiationpriorityid, tz.Mstrecorddifferentiationworkingcatid, tz.Responsetimeinhr, tz.Responsetimeinmin, tz.Responsetimeinsec, tz.Resolutiontimeinhr, tz.Resolutiontimeinmin, tz.Resolutiontimeinsec, tz.Responsecompliance, tz.Resolutioncompliance, tz.Supportgroupspecific)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Slaid, tz.Mstrecorddifferentiationtickettypeid, tz.Mstrecorddifferentiationpriorityid, tz.Mstrecorddifferentiationworkingcatid, tz.Responsetimeinhr, tz.Responsetimeinmin, tz.Responsetimeinsec, tz.Resolutiontimeinhr, tz.Resolutiontimeinmin, tz.Resolutiontimeinsec, tz.Responsecompliance, tz.Resolutioncompliance, tz.Supportgroupspecific)
	if err != nil {
		logger.Log.Println("InsertMstslafullfillmentcriteria Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllMstslafullfillmentcriteria(page *entities.MstslafullfillmentcriteriaEntity, OrgnType int64) ([]entities.MstslafullfillmentcriteriaEntity, error) {
	logger.Log.Println("In side GelAllMstslafullfillmentcriteria")
	values := []entities.MstslafullfillmentcriteriaEntity{}
	var getMstslafullfillmentcriteria string
	var params []interface{}
	if OrgnType == 1 {
		getMstslafullfillmentcriteria = "SELECT id as Id, clientid as Clientid, mstorgnhirarchyid as Mstorgnhirarchyid, slaid as Slaid, mstrecorddifferentiationtickettypeid as Mstrecorddifferentiationtickettypeid, mstrecorddifferentiationpriorityid as Mstrecorddifferentiationpriorityid, mstrecorddifferentiationworkingcatid as Mstrecorddifferentiationworkingcatid, responsetimeinhr as Responsetimeinhr, responsetimeinmin as Responsetimeinmin, responsetimeinsec as Responsetimeinsec, resolutiontimeinhr as Resolutiontimeinhr, resolutiontimeinmin as Resolutiontimeinmin, resolutiontimeinsec as Resolutiontimeinsec, COALESCE(responsecompliance,0) as Responsecompliance, COALESCE(resolutioncompliance,0) as Resolutioncompliance, supportgroupspecific as Supportgroupspecific, activeflg as Activeflg,(select name from mstclient where id = clientid ) as Clientname,(select name from mstorgnhierarchy where id = mstorgnhirarchyid ) as Mstorgnhirarchyname,(select name from mstrecorddifferentiation where id=mstrecorddifferentiationtickettypeid and deleteflg =0 and activeflg=1) as Tickettypename,(select name from mstrecorddifferentiation where id=mstrecorddifferentiationpriorityid and deleteflg =0 and activeflg=1) as Priorityname,COALESCE((select name from mstrecorddifferentiation where id=mstrecorddifferentiationworkingcatid and deleteflg =0 and activeflg=1),'') as Workingcatname,(select slaname from mstclientsla where id=slaid and deleteflg =0 and activeflg=1) as Slaname FROM mstslafullfillmentcriteria WHERE deleteflg =0 and activeflg=1 ORDER BY id DESC LIMIT ?,?"
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else if OrgnType == 2 {
		getMstslafullfillmentcriteria = "SELECT id as Id, clientid as Clientid, mstorgnhirarchyid as Mstorgnhirarchyid, slaid as Slaid, mstrecorddifferentiationtickettypeid as Mstrecorddifferentiationtickettypeid, mstrecorddifferentiationpriorityid as Mstrecorddifferentiationpriorityid, mstrecorddifferentiationworkingcatid as Mstrecorddifferentiationworkingcatid, responsetimeinhr as Responsetimeinhr, responsetimeinmin as Responsetimeinmin, responsetimeinsec as Responsetimeinsec, resolutiontimeinhr as Resolutiontimeinhr, resolutiontimeinmin as Resolutiontimeinmin, resolutiontimeinsec as Resolutiontimeinsec, COALESCE(responsecompliance,0) as Responsecompliance, COALESCE(resolutioncompliance,0) as Resolutioncompliance, supportgroupspecific as Supportgroupspecific, activeflg as Activeflg,(select name from mstclient where id = clientid ) as Clientname,(select name from mstorgnhierarchy where id = mstorgnhirarchyid ) as Mstorgnhirarchyname,(select name from mstrecorddifferentiation where id=mstrecorddifferentiationtickettypeid and deleteflg =0 and activeflg=1) as Tickettypename,(select name from mstrecorddifferentiation where id=mstrecorddifferentiationpriorityid and deleteflg =0 and activeflg=1) as Priorityname,COALESCE((select name from mstrecorddifferentiation where id=mstrecorddifferentiationworkingcatid and deleteflg =0 and activeflg=1),'') as Workingcatname,(select slaname from mstclientsla where id=slaid and deleteflg =0 and activeflg=1) as Slaname FROM mstslafullfillmentcriteria WHERE clientid = ? AND deleteflg =0 and activeflg=1 ORDER BY id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else {
		getMstslafullfillmentcriteria = "SELECT id as Id, clientid as Clientid, mstorgnhirarchyid as Mstorgnhirarchyid, slaid as Slaid, mstrecorddifferentiationtickettypeid as Mstrecorddifferentiationtickettypeid, mstrecorddifferentiationpriorityid as Mstrecorddifferentiationpriorityid, mstrecorddifferentiationworkingcatid as Mstrecorddifferentiationworkingcatid, responsetimeinhr as Responsetimeinhr, responsetimeinmin as Responsetimeinmin, responsetimeinsec as Responsetimeinsec, resolutiontimeinhr as Resolutiontimeinhr, resolutiontimeinmin as Resolutiontimeinmin, resolutiontimeinsec as Resolutiontimeinsec, COALESCE(responsecompliance,0) as Responsecompliance, COALESCE(resolutioncompliance,0) as Resolutioncompliance, supportgroupspecific as Supportgroupspecific, activeflg as Activeflg,(select name from mstclient where id = clientid ) as Clientname,(select name from mstorgnhierarchy where id = mstorgnhirarchyid ) as Mstorgnhirarchyname,(select name from mstrecorddifferentiation where id=mstrecorddifferentiationtickettypeid and deleteflg =0 and activeflg=1) as Tickettypename,(select name from mstrecorddifferentiation where id=mstrecorddifferentiationpriorityid and deleteflg =0 and activeflg=1) as Priorityname,COALESCE((select name from mstrecorddifferentiation where id=mstrecorddifferentiationworkingcatid and deleteflg =0 and activeflg=1),'') as Workingcatname,(select slaname from mstclientsla where id=slaid and deleteflg =0 and activeflg=1) as Slaname FROM mstslafullfillmentcriteria WHERE clientid = ? AND mstorgnhirarchyid = ?  AND deleteflg =0 and activeflg=1 ORDER BY id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Mstorgnhirarchyid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	}
	rows, err := dbc.DB.Query(getMstslafullfillmentcriteria, params...)

	// rows, err := dbc.DB.Query(getMstslafullfillmentcriteria, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstslafullfillmentcriteria Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstslafullfillmentcriteriaEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Slaid, &value.Mstrecorddifferentiationtickettypeid, &value.Mstrecorddifferentiationpriorityid, &value.Mstrecorddifferentiationworkingcatid, &value.Responsetimeinhr, &value.Responsetimeinmin, &value.Responsetimeinsec, &value.Resolutiontimeinhr, &value.Resolutiontimeinmin, &value.Resolutiontimeinsec, &value.Responsecompliance, &value.Resolutioncompliance, &value.Supportgroupspecific, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Tickettypename, &value.Priorityname, &value.Workingcatname, &value.Slaname)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateMstslafullfillmentcriteria(tz *entities.MstslafullfillmentcriteriaEntity) error {
	logger.Log.Println("In side UpdateMstslafullfillmentcriteria")
	stmt, err := dbc.DB.Prepare(updateMstslafullfillmentcriteria)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMstslafullfillmentcriteria Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Slaid, tz.Mstrecorddifferentiationtickettypeid, tz.Mstrecorddifferentiationpriorityid, tz.Mstrecorddifferentiationworkingcatid, tz.Responsetimeinhr, tz.Responsetimeinmin, tz.Responsetimeinsec, tz.Resolutiontimeinhr, tz.Resolutiontimeinmin, tz.Resolutiontimeinsec, tz.Responsecompliance, tz.Resolutioncompliance, tz.Supportgroupspecific, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMstslafullfillmentcriteria Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMstslafullfillmentcriteria(tz *entities.MstslafullfillmentcriteriaEntity) error {
	logger.Log.Println("In side DeleteMstslafullfillmentcriteria")
	stmt, err := dbc.DB.Prepare(deleteMstslafullfillmentcriteria)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMstslafullfillmentcriteria Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMstslafullfillmentcriteria Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMstslafullfillmentcriteriaCount(tz *entities.MstslafullfillmentcriteriaEntity, OrgnTypeID int64) (entities.MstslafullfillmentcriteriaEntities, error) {
	logger.Log.Println("In side GetMstslafullfillmentcriteriaCount")
	value := entities.MstslafullfillmentcriteriaEntities{}
	var getMstslafullfillmentcriteriacount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getMstslafullfillmentcriteriacount = "SELECT count(a.id) as total FROM mstslafullfillmentcriteria a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiation d,mstrecorddifferentiation e,mstclientsla g WHERE  a.deleteflg =0 and a.activeflg=1 and b.id = a.clientid and c.id = a.mstorgnhirarchyid and d.id = a.mstrecorddifferentiationtickettypeid and e.id =a.mstrecorddifferentiationpriorityid and g.id=a.slaid and d.deleteflg =0 and d.activeflg=1 and e.deleteflg =0 and e.activeflg=1 and g.deleteflg =0 and g.activeflg=1"
	} else if OrgnTypeID == 2 {
		getMstslafullfillmentcriteriacount = "SELECT count(a.id) as total FROM mstslafullfillmentcriteria a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiation d,mstrecorddifferentiation e,mstclientsla g WHERE a.clientid = ? AND  a.deleteflg =0 and a.activeflg=1 and b.id = a.clientid and c.id = a.mstorgnhirarchyid and d.id = a.mstrecorddifferentiationtickettypeid and e.id =a.mstrecorddifferentiationpriorityid and g.id=a.slaid and d.deleteflg =0 and d.activeflg=1 and e.deleteflg =0 and e.activeflg=1 and g.deleteflg =0 and g.activeflg=1"
		params = append(params, tz.Clientid)
	} else {
		getMstslafullfillmentcriteriacount = "SELECT count(a.id) as total FROM mstslafullfillmentcriteria a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiation d,mstrecorddifferentiation e,mstclientsla g WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.deleteflg =0 and a.activeflg=1 and b.id = a.clientid and c.id = a.mstorgnhirarchyid and d.id = a.mstrecorddifferentiationtickettypeid and e.id =a.mstrecorddifferentiationpriorityid and g.id=a.slaid and d.deleteflg =0 and d.activeflg=1 and e.deleteflg =0 and e.activeflg=1 and g.deleteflg =0 and g.activeflg=1"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getMstslafullfillmentcriteriacount, params...).Scan(&value.Total)

	// err := dbc.DB.QueryRow(getMstslafullfillmentcriteriacount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMstslafullfillmentcriteriaCount Get Statement Prepare Error", err)
		return value, err
	}
}
