package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertMstslafcrecorddiff = "INSERT INTO mstslafcrecorddiff (clientid, mstorgnhirarchyid, mstslaid, recorddifftypeidtype, recorddiffidtype, recorddifftypeidstatus, recorddiffidstatus, startstopindicator,slametertypeid) VALUES (?,?,?,?,?,?,?,?,?)"
var duplicateMstslafcrecorddiff = "SELECT count(id) total FROM  mstslafcrecorddiff WHERE clientid = ? AND mstorgnhirarchyid = ? AND mstslaid = ? AND recorddifftypeidtype = ? AND recorddiffidtype = ? AND recorddifftypeidstatus = ? AND recorddiffidstatus = ? AND startstopindicator = ? AND slametertypeid=? AND deleteflg = 0 and activeflg=1"
var getMstslafcrecorddiff = "SELECT id as Id, clientid as Clientid, mstorgnhirarchyid as Mstorgnhirarchyid, mstslaid as Mstslaid, recorddifftypeidtype as Recorddifftypeidtype, recorddiffidtype as Recorddiffidtype, recorddifftypeidstatus as Recorddifftypeidstatus, recorddiffidstatus as Recorddiffidstatus, startstopindicator as Startstopindicator, activeflg as Activeflg,(select name from mstclient where id=clientid) as Clientname,(select name from mstorgnhierarchy where id=mstorgnhirarchyid) as Mstorgnhirarchyname,(select typename from mstrecorddifferentiationtype where id=recorddifftypeidtype and deleteflg =0 and activeflg=1) AS Recorddifftypeidtypenm,(select name from mstrecorddifferentiation where id=recorddiffidtype and deleteflg =0 and activeflg=1) as Recorddiffidtypenm,(select typename from mstrecorddifferentiationtype where id=recorddifftypeidstatus and deleteflg =0 and activeflg=1) AS Recorddifftypeidstatusnm,(select name from mstrecorddifferentiation where id=recorddiffidstatus and deleteflg =0 and activeflg=1) as Recorddiffidstatusnm,(select slaname from mstclientsla where id=mstslaid and deleteflg =0 and activeflg=1) as Slaname,COALESCE((select parentid from mstrecorddifferentiationtype where id=recorddifftypeidtype and deleteflg =0 and activeflg=1),'0') as Recorddifftypetypeparent,COALESCE((select parentid from mstrecorddifferentiationtype where id=recorddifftypeidstatus and deleteflg =0 and activeflg=1),'0') as Recorddifftypeidstatusparent,slametertypeid as SLAmetertypeID,(SELECT name FROM mstslametertype WHERE id=slametertypeid) as SLAmetertypename,(SELECT name FROM mstslaindicatorterm WHERE id=startstopindicator AND activeflg=1 AND deleteflg=0) as Startstopindicatorname  FROM mstslafcrecorddiff WHERE clientid = ? AND mstorgnhirarchyid = ?  AND deleteflg =0 and activeflg=1 ORDER BY id DESC LIMIT ?,?"
//var getMstslafcrecorddiffcount = "SELECT count(a.id) total FROM  mstslafcrecorddiff a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstrecorddifferentiationtype f,mstrecorddifferentiation g,mstclientsla h WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND  a.deleteflg =0 and a.activeflg=1 and b.id=a.clientid and c.id=a.mstorgnhirarchyid and d.id=a.recorddifftypeidtype and d.deleteflg =0 and d.activeflg=1 and e.id=a.recorddiffidtype and e.deleteflg =0 and e.activeflg=1 and f.id=a.recorddifftypeidstatus and f.deleteflg =0 and f.activeflg=1 and g.id=a.recorddiffidstatus and g.deleteflg =0 and g.activeflg=1 and h.id=a.mstslaid and h.deleteflg =0 and h.activeflg=1"
var updateMstslafcrecorddiff = "UPDATE mstslafcrecorddiff SET mstorgnhirarchyid = ?, mstslaid = ?, recorddifftypeidtype = ?, recorddiffidtype = ?, recorddifftypeidstatus = ?, recorddiffidstatus = ?, startstopindicator = ?,slametertypeid=? WHERE id = ? "
var deleteMstslafcrecorddiff = "UPDATE mstslafcrecorddiff SET deleteflg = '1' WHERE id = ? "
var slametertype = "SELECT id,name FROM mstslametertype"
var slatermname = "SELECT id,name FROM mstslaindicatorterm WHERE clientid=? AND mstorgnhirarchyid=? AND metertypeid=? AND activeflg=1 AND deleteflg=0"

func (dbc DbConn) CheckDuplicateMstslafcrecorddiff(tz *entities.MstslafcrecorddiffEntity) (entities.MstslafcrecorddiffEntities, error) {
	logger.Log.Println("In side CheckDuplicateMstslafcrecorddiff")
	value := entities.MstslafcrecorddiffEntities{}
	err := dbc.DB.QueryRow(duplicateMstslafcrecorddiff, tz.Clientid, tz.Mstorgnhirarchyid, tz.Mstslaid, tz.Recorddifftypeidtype, tz.Recorddiffidtype, tz.Recorddifftypeidstatus, tz.Recorddiffidstatus, tz.Startstopindicator, tz.SLAmetertypeID).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMstslafcrecorddiff Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertMstslafcrecorddiff(tz *entities.MstslafcrecorddiffEntity) (int64, error) {
	logger.Log.Println("In side InsertMstslafcrecorddiff")
	logger.Log.Println("Query -->", insertMstslafcrecorddiff)
	stmt, err := dbc.DB.Prepare(insertMstslafcrecorddiff)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMstslafcrecorddiff Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Mstslaid, tz.Recorddifftypeidtype, tz.Recorddiffidtype, tz.Recorddifftypeidstatus, tz.Recorddiffidstatus, tz.Startstopindicator)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Mstslaid, tz.Recorddifftypeidtype, tz.Recorddiffidtype, tz.Recorddifftypeidstatus, tz.Recorddiffidstatus, tz.Startstopindicator, tz.SLAmetertypeID)
	if err != nil {
		logger.Log.Println("InsertMstslafcrecorddiff Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllMstslafcrecorddiff(tz *entities.MstslafcrecorddiffEntity, OrgnType int64) ([]entities.MstslafcrecorddiffEntity, error) {
	logger.Log.Println("In side GelAllMstslafcrecorddiff")
	values := []entities.MstslafcrecorddiffEntity{}
	var getMstslafcrecorddiff string
	var params []interface{}
	if OrgnType == 1 {
		getMstslafcrecorddiff = "SELECT id as Id, clientid as Clientid, mstorgnhirarchyid as Mstorgnhirarchyid, mstslaid as Mstslaid, recorddifftypeidtype as Recorddifftypeidtype, recorddiffidtype as Recorddiffidtype, recorddifftypeidstatus as Recorddifftypeidstatus, recorddiffidstatus as Recorddiffidstatus, startstopindicator as Startstopindicator, activeflg as Activeflg,(select name from mstclient where id=clientid) as Clientname,(select name from mstorgnhierarchy where id=mstorgnhirarchyid) as Mstorgnhirarchyname,(select typename from mstrecorddifferentiationtype where id=recorddifftypeidtype and deleteflg =0 and activeflg=1) AS Recorddifftypeidtypenm,(select name from mstrecorddifferentiation where id=recorddiffidtype and deleteflg =0 and activeflg=1) as Recorddiffidtypenm,(select typename from mstrecorddifferentiationtype where id=recorddifftypeidstatus and deleteflg =0 and activeflg=1) AS Recorddifftypeidstatusnm,(select name from mstrecorddifferentiation where id=recorddiffidstatus and deleteflg =0 and activeflg=1) as Recorddiffidstatusnm,(select slaname from mstclientsla where id=mstslaid and deleteflg =0 and activeflg=1) as Slaname,COALESCE((select parentid from mstrecorddifferentiationtype where id=recorddifftypeidtype and deleteflg =0 and activeflg=1),'0') as Recorddifftypetypeparent,COALESCE((select parentid from mstrecorddifferentiationtype where id=recorddifftypeidstatus and deleteflg =0 and activeflg=1),'0') as Recorddifftypeidstatusparent,slametertypeid as SLAmetertypeID,(SELECT name FROM mstslametertype WHERE id=slametertypeid) as SLAmetertypename,(SELECT name FROM mstslaindicatorterm WHERE id=startstopindicator AND activeflg=1 AND deleteflg=0) as Startstopindicatorname  FROM mstslafcrecorddiff WHERE  deleteflg =0 and activeflg=1 ORDER BY id DESC LIMIT ?,?"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		getMstslafcrecorddiff = "SELECT id as Id, clientid as Clientid, mstorgnhirarchyid as Mstorgnhirarchyid, mstslaid as Mstslaid, recorddifftypeidtype as Recorddifftypeidtype, recorddiffidtype as Recorddiffidtype, recorddifftypeidstatus as Recorddifftypeidstatus, recorddiffidstatus as Recorddiffidstatus, startstopindicator as Startstopindicator, activeflg as Activeflg,(select name from mstclient where id=clientid) as Clientname,(select name from mstorgnhierarchy where id=mstorgnhirarchyid) as Mstorgnhirarchyname,(select typename from mstrecorddifferentiationtype where id=recorddifftypeidtype and deleteflg =0 and activeflg=1) AS Recorddifftypeidtypenm,(select name from mstrecorddifferentiation where id=recorddiffidtype and deleteflg =0 and activeflg=1) as Recorddiffidtypenm,(select typename from mstrecorddifferentiationtype where id=recorddifftypeidstatus and deleteflg =0 and activeflg=1) AS Recorddifftypeidstatusnm,(select name from mstrecorddifferentiation where id=recorddiffidstatus and deleteflg =0 and activeflg=1) as Recorddiffidstatusnm,(select slaname from mstclientsla where id=mstslaid and deleteflg =0 and activeflg=1) as Slaname,COALESCE((select parentid from mstrecorddifferentiationtype where id=recorddifftypeidtype and deleteflg =0 and activeflg=1),'0') as Recorddifftypetypeparent,COALESCE((select parentid from mstrecorddifferentiationtype where id=recorddifftypeidstatus and deleteflg =0 and activeflg=1),'0') as Recorddifftypeidstatusparent,slametertypeid as SLAmetertypeID,(SELECT name FROM mstslametertype WHERE id=slametertypeid) as SLAmetertypename,(SELECT name FROM mstslaindicatorterm WHERE id=startstopindicator AND activeflg=1 AND deleteflg=0) as Startstopindicatorname  FROM mstslafcrecorddiff WHERE clientid = ? AND  deleteflg =0 and activeflg=1 ORDER BY id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		getMstslafcrecorddiff = "SELECT id as Id, clientid as Clientid, mstorgnhirarchyid as Mstorgnhirarchyid, mstslaid as Mstslaid, recorddifftypeidtype as Recorddifftypeidtype, recorddiffidtype as Recorddiffidtype, recorddifftypeidstatus as Recorddifftypeidstatus, recorddiffidstatus as Recorddiffidstatus, startstopindicator as Startstopindicator, activeflg as Activeflg,(select name from mstclient where id=clientid) as Clientname,(select name from mstorgnhierarchy where id=mstorgnhirarchyid) as Mstorgnhirarchyname,(select typename from mstrecorddifferentiationtype where id=recorddifftypeidtype and deleteflg =0 and activeflg=1) AS Recorddifftypeidtypenm,(select name from mstrecorddifferentiation where id=recorddiffidtype and deleteflg =0 and activeflg=1) as Recorddiffidtypenm,(select typename from mstrecorddifferentiationtype where id=recorddifftypeidstatus and deleteflg =0 and activeflg=1) AS Recorddifftypeidstatusnm,(select name from mstrecorddifferentiation where id=recorddiffidstatus and deleteflg =0 and activeflg=1) as Recorddiffidstatusnm,(select slaname from mstclientsla where id=mstslaid and deleteflg =0 and activeflg=1) as Slaname,COALESCE((select parentid from mstrecorddifferentiationtype where id=recorddifftypeidtype and deleteflg =0 and activeflg=1),'0') as Recorddifftypetypeparent,COALESCE((select parentid from mstrecorddifferentiationtype where id=recorddifftypeidstatus and deleteflg =0 and activeflg=1),'0') as Recorddifftypeidstatusparent,slametertypeid as SLAmetertypeID,(SELECT name FROM mstslametertype WHERE id=slametertypeid) as SLAmetertypename,(SELECT name FROM mstslaindicatorterm WHERE id=startstopindicator AND activeflg=1 AND deleteflg=0) as Startstopindicatorname  FROM mstslafcrecorddiff WHERE clientid = ? AND mstorgnhirarchyid = ?  AND deleteflg =0 and activeflg=1 ORDER BY id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}
	rows, err := dbc.DB.Query(getMstslafcrecorddiff, params...)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstslafcrecorddiff Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstslafcrecorddiffEntity{}
		err = rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Mstslaid, &value.Recorddifftypeidtype, &value.Recorddiffidtype, &value.Recorddifftypeidstatus, &value.Recorddiffidstatus, &value.Startstopindicator, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Recorddifftypeidtypenm, &value.Recorddiffidtypenm, &value.Recorddifftypeidstatusnm, &value.Recorddiffidstatusnm, &value.Slaname, &value.Recorddifftypetypeparent, &value.Recorddifftypeidstatusparent, &value.SLAmetertypeID, &value.SLAmetertypename, &value.Startstopindicatorname)
		logger.Log.Println(err)

		values = append(values, value)
	}
	return values, nil
}
func (dbc DbConn) UpdateMstslafcrecorddiff(tz *entities.MstslafcrecorddiffEntity) error {
	logger.Log.Println("In side UpdateMstslafcrecorddiff")
	stmt, err := dbc.DB.Prepare(updateMstslafcrecorddiff)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMstslafcrecorddiff Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Mstslaid, tz.Recorddifftypeidtype, tz.Recorddiffidtype, tz.Recorddifftypeidstatus, tz.Recorddiffidstatus, tz.Startstopindicator, tz.SLAmetertypeID, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMstslafcrecorddiff Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMstslafcrecorddiff(tz *entities.MstslafcrecorddiffEntity) error {
	logger.Log.Println("In side DeleteMstslafcrecorddiff")
	stmt, err := dbc.DB.Prepare(deleteMstslafcrecorddiff)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMstslafcrecorddiff Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMstslafcrecorddiff Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMstslafcrecorddiffCount(tz *entities.MstslafcrecorddiffEntity, OrgnTypeID int64) (entities.MstslafcrecorddiffEntities, error) {
	logger.Log.Println("In side GetMstslafcrecorddiffCount")
	value := entities.MstslafcrecorddiffEntities{}
	var getMstslafcrecorddiffcount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getMstslafcrecorddiffcount = "SELECT count(a.id) total FROM  mstslafcrecorddiff a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstrecorddifferentiationtype f,mstrecorddifferentiation g,mstclientsla h WHERE a.deleteflg =0 and a.activeflg=1 and b.id=a.clientid and c.id=a.mstorgnhirarchyid and d.id=a.recorddifftypeidtype and d.deleteflg =0 and d.activeflg=1 and e.id=a.recorddiffidtype and e.deleteflg =0 and e.activeflg=1 and f.id=a.recorddifftypeidstatus and f.deleteflg =0 and f.activeflg=1 and g.id=a.recorddiffidstatus and g.deleteflg =0 and g.activeflg=1 and h.id=a.mstslaid and h.deleteflg =0 and h.activeflg=1"
	} else if OrgnTypeID == 2 {
		getMstslafcrecorddiffcount = "SELECT count(a.id) total FROM  mstslafcrecorddiff a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstrecorddifferentiationtype f,mstrecorddifferentiation g,mstclientsla h WHERE a.clientid = ? AND  a.deleteflg =0 and a.activeflg=1 and b.id=a.clientid and c.id=a.mstorgnhirarchyid and d.id=a.recorddifftypeidtype and d.deleteflg =0 and d.activeflg=1 and e.id=a.recorddiffidtype and e.deleteflg =0 and e.activeflg=1 and f.id=a.recorddifftypeidstatus and f.deleteflg =0 and f.activeflg=1 and g.id=a.recorddiffidstatus and g.deleteflg =0 and g.activeflg=1 and h.id=a.mstslaid and h.deleteflg =0 and h.activeflg=1"
		params = append(params, tz.Clientid)
	} else {
		getMstslafcrecorddiffcount = "SELECT count(a.id) total FROM  mstslafcrecorddiff a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstrecorddifferentiationtype f,mstrecorddifferentiation g,mstclientsla h WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND  a.deleteflg =0 and a.activeflg=1 and b.id=a.clientid and c.id=a.mstorgnhirarchyid and d.id=a.recorddifftypeidtype and d.deleteflg =0 and d.activeflg=1 and e.id=a.recorddiffidtype and e.deleteflg =0 and e.activeflg=1 and f.id=a.recorddifftypeidstatus and f.deleteflg =0 and f.activeflg=1 and g.id=a.recorddiffidstatus and g.deleteflg =0 and g.activeflg=1 and h.id=a.mstslaid and h.deleteflg =0 and h.activeflg=1"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getMstslafcrecorddiffcount, params...).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMstslafcrecorddiffCount Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) GetSLAmetertype() ([]entities.MstslametertypeEntity, error) {
	logger.Log.Println("In side GetSLAmetertype")
	values := []entities.MstslametertypeEntity{}
	rows, err := dbc.DB.Query(slametertype)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetSLAmetertype Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstslametertypeEntity{}
		rows.Scan(&value.ID, &value.Name)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) GetSLAtermnames(page *entities.MstslafcrecorddiffEntity) ([]entities.MstslaindicatortermEntity, error) {
	logger.Log.Println("In side GetSLAtermnames")
	values := []entities.MstslaindicatortermEntity{}
	rows, err := dbc.DB.Query(slatermname, page.Clientid, page.Mstorgnhirarchyid, page.SLAmetertypeID)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetSLAtermnames Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstslaindicatortermEntity{}
		rows.Scan(&value.ID, &value.Name)
		values = append(values, value)
	}
	return values, nil
}
