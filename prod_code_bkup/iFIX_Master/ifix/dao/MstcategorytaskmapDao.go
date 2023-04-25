package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertMstcategorytaskmap = "INSERT INTO mstcategorytaskmap (clientid, mstorgnhirarchyid, fromtickettypedifftypeid, fromtickettypediffid, fromcatdifftypeid, fromcatlabelid, fromcatdiffid, totickettypedifftypeid, totickettypediffid, tocatdifftypeid, tocatlabelid, tocatdiffid) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)"
var duplicateMstcategorytaskmap = "SELECT count(id) total FROM  mstcategorytaskmap WHERE clientid = ? AND mstorgnhirarchyid = ? AND fromtickettypedifftypeid = ? AND fromtickettypediffid = ? AND fromcatdifftypeid = ? AND fromcatlabelid = ? AND fromcatdiffid = ? AND totickettypedifftypeid = ? AND totickettypediffid = ? AND tocatdifftypeid = ? AND tocatlabelid = ? AND tocatdiffid = ? AND deleteflg = 0"

// var getMstcategorytaskmap = "SELECT id as Id, clientid as Clientid, mstorgnhirarchyid as Mstorgnhirarchyid, fromtickettypedifftypeid as Fromtickettypedifftypeid, fromtickettypediffid as Fromtickettypediffid, fromcatdifftypeid as Fromcatdifftypeid, fromcatlabelid as Fromcatlabelid, fromcatdiffid as Fromcatdiffid, totickettypedifftypeid as Totickettypedifftypeid, totickettypediffid as Totickettypediffid, tocatdifftypeid as Tocatdifftypeid, tocatlabelid as Tocatlabelid, tocatdiffid as Tocatdiffid, activeflg as Activeflg,(select name from mstclient where id=clientid) as Clientname,(select name from mstorgnhierarchy where id=mstorgnhirarchyid) as Mstorgnhirarchyname,(select typename from mstrecorddifferentiationtype where id=fromtickettypedifftypeid and deleteflg =0 and activeflg=1) as Fromtickettypedifftypename,(select name from mstrecorddifferentiation where id=fromtickettypediffid and deleteflg =0 and activeflg=1) as Fromtickettypediffname,(select typename from mstrecorddifferentiationtype where id=fromcatdifftypeid and deleteflg =0 and activeflg=1) as Fromcatdifftypename,(select typename from mstrecorddifferentiationtype where id=fromcatlabelid and deleteflg =0 and activeflg=1) as Fromcatlabelname,(select name from mstrecorddifferentiation where id=fromcatdiffid and deleteflg =0 and activeflg=1) as Fromcatdiffname,(select typename from mstrecorddifferentiationtype where id=totickettypedifftypeid and deleteflg =0 and activeflg=1) as Totickettypedifftypename,(select name from mstrecorddifferentiation where id=totickettypediffid and deleteflg =0 and activeflg=1) as Totickettypediffname,(select typename from mstrecorddifferentiationtype where id=tocatdifftypeid and deleteflg =0 and activeflg=1) as Tocatdifftypename,(select typename from mstrecorddifferentiationtype where id=tocatlabelid and deleteflg =0 and activeflg=1) as Tocatlabelname,(select name from mstrecorddifferentiation where id=tocatdiffid and deleteflg =0 and activeflg=1) as Tocatdiffnam FROM mstcategorytaskmap WHERE clientid = ? AND mstorgnhirarchyid = ?  AND deleteflg =0 and activeflg=1 ORDER BY id DESC LIMIT ?,?"
// var getMstcategorytaskmapcount = "SELECT count(a.id) total FROM  mstcategorytaskmap a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstrecorddifferentiationtype f, mstrecorddifferentiationtype g, mstrecorddifferentiation h, mstrecorddifferentiationtype i, mstrecorddifferentiation j, mstrecorddifferentiationtype k, mstrecorddifferentiationtype l, mstrecorddifferentiation m WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND  a.deleteflg =0 and a.activeflg=1 and b.id=a.clientid and c.id=a.mstorgnhirarchyid and d.id=a.fromtickettypedifftypeid and d.deleteflg =0 and d.activeflg=1 and e.id=a.fromtickettypediffid and e.deleteflg =0 and e.activeflg=1 and f.id=a.fromcatdifftypeid and f.deleteflg =0 and f.activeflg=1 and g.id=a.fromcatlabelid and g.deleteflg =0 and g.activeflg=1 and h.id=a.fromcatdiffid and h.deleteflg =0 and h.activeflg=1 and i.id=a.totickettypedifftypeid and i.deleteflg =0 and i.activeflg=1 and j.id=a.totickettypediffid and j.deleteflg =0 and j.activeflg=1 and k.id=a.tocatdifftypeid and k.deleteflg =0 and k.activeflg=1 and l.id=a.tocatlabelid and l.deleteflg =0 and l.activeflg=1 and m.id=a.tocatdiffid and m.deleteflg =0 and m.activeflg=1"
var updateMstcategorytaskmap = "UPDATE mstcategorytaskmap SET mstorgnhirarchyid = ?, fromtickettypedifftypeid = ?, fromtickettypediffid = ?, fromcatdifftypeid = ?, fromcatlabelid = ?, fromcatdiffid = ?, totickettypedifftypeid = ?, totickettypediffid = ?, tocatdifftypeid = ?, tocatlabelid = ?, tocatdiffid = ? WHERE id = ? "
var deleteMstcategorytaskmap = "UPDATE mstcategorytaskmap SET deleteflg = '1' WHERE id = ? "

func (dbc DbConn) CheckDuplicateMstcategorytaskmap(tz *entities.MstcategorytaskmapEntity) (entities.MstcategorytaskmapEntities, error) {
	logger.Log.Println("In side CheckDuplicateMstcategorytaskmap")
	value := entities.MstcategorytaskmapEntities{}
	err := dbc.DB.QueryRow(duplicateMstcategorytaskmap, tz.Clientid, tz.Mstorgnhirarchyid, tz.Fromtickettypedifftypeid, tz.Fromtickettypediffid, tz.Fromcatdifftypeid, tz.Fromcatlabelid, tz.Fromcatdiffid, tz.Totickettypedifftypeid, tz.Totickettypediffid, tz.Tocatdifftypeid, tz.Tocatlabelid, tz.Tocatdiffid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMstcategorytaskmap Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertMstcategorytaskmap(tz *entities.MstcategorytaskmapEntity) (int64, error) {
	logger.Log.Println("In side InsertMstcategorytaskmap")
	logger.Log.Println("Query -->", insertMstcategorytaskmap)
	stmt, err := dbc.DB.Prepare(insertMstcategorytaskmap)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMstcategorytaskmap Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Fromtickettypedifftypeid, tz.Fromtickettypediffid, tz.Fromcatdifftypeid, tz.Fromcatlabelid, tz.Fromcatdiffid, tz.Totickettypedifftypeid, tz.Totickettypediffid, tz.Tocatdifftypeid, tz.Tocatlabelid, tz.Tocatdiffid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Fromtickettypedifftypeid, tz.Fromtickettypediffid, tz.Fromcatdifftypeid, tz.Fromcatlabelid, tz.Fromcatdiffid, tz.Totickettypedifftypeid, tz.Totickettypediffid, tz.Tocatdifftypeid, tz.Tocatlabelid, tz.Tocatdiffid)
	if err != nil {
		logger.Log.Println("InsertMstcategorytaskmap Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllMstcategorytaskmap(page *entities.MstcategorytaskmapEntity, OrgnType int64) ([]entities.MstcategorytaskmapEntity, error) {
	logger.Log.Println("In side GelAllMstcategorytaskmap")
	values := []entities.MstcategorytaskmapEntity{}
	var getMstcategorytaskmap string
	var params []interface{}
	if OrgnType == 1 {
		getMstcategorytaskmap = "SELECT id as Id, clientid as Clientid, mstorgnhirarchyid as Mstorgnhirarchyid, fromtickettypedifftypeid as Fromtickettypedifftypeid, fromtickettypediffid as Fromtickettypediffid, fromcatdifftypeid as Fromcatdifftypeid, fromcatlabelid as Fromcatlabelid, fromcatdiffid as Fromcatdiffid, totickettypedifftypeid as Totickettypedifftypeid, totickettypediffid as Totickettypediffid, tocatdifftypeid as Tocatdifftypeid, tocatlabelid as Tocatlabelid, tocatdiffid as Tocatdiffid, activeflg as Activeflg,(select name from mstclient where id=clientid) as Clientname,(select name from mstorgnhierarchy where id=mstorgnhirarchyid) as Mstorgnhirarchyname,(select typename from mstrecorddifferentiationtype where id=fromtickettypedifftypeid and deleteflg =0 and activeflg=1) as Fromtickettypedifftypename,(select name from mstrecorddifferentiation where id=fromtickettypediffid and deleteflg =0 and activeflg=1) as Fromtickettypediffname,(select typename from mstrecorddifferentiationtype where id=fromcatdifftypeid and deleteflg =0 and activeflg=1) as Fromcatdifftypename,(select typename from mstrecorddifferentiationtype where id=fromcatlabelid and deleteflg =0 and activeflg=1) as Fromcatlabelname,(select name from mstrecorddifferentiation where id=fromcatdiffid and deleteflg =0 and activeflg=1) as Fromcatdiffname,(select typename from mstrecorddifferentiationtype where id=totickettypedifftypeid and deleteflg =0 and activeflg=1) as Totickettypedifftypename,(select name from mstrecorddifferentiation where id=totickettypediffid and deleteflg =0 and activeflg=1) as Totickettypediffname,(select typename from mstrecorddifferentiationtype where id=tocatdifftypeid and deleteflg =0 and activeflg=1) as Tocatdifftypename,(select typename from mstrecorddifferentiationtype where id=tocatlabelid and deleteflg =0 and activeflg=1) as Tocatlabelname,(select name from mstrecorddifferentiation where id=tocatdiffid and deleteflg =0 and activeflg=1) as Tocatdiffnam FROM mstcategorytaskmap WHERE deleteflg =0 and activeflg=1 ORDER BY id DESC LIMIT ?,?"
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else if OrgnType == 2 {
		getMstcategorytaskmap = "SELECT id as Id, clientid as Clientid, mstorgnhirarchyid as Mstorgnhirarchyid, fromtickettypedifftypeid as Fromtickettypedifftypeid, fromtickettypediffid as Fromtickettypediffid, fromcatdifftypeid as Fromcatdifftypeid, fromcatlabelid as Fromcatlabelid, fromcatdiffid as Fromcatdiffid, totickettypedifftypeid as Totickettypedifftypeid, totickettypediffid as Totickettypediffid, tocatdifftypeid as Tocatdifftypeid, tocatlabelid as Tocatlabelid, tocatdiffid as Tocatdiffid, activeflg as Activeflg,(select name from mstclient where id=clientid) as Clientname,(select name from mstorgnhierarchy where id=mstorgnhirarchyid) as Mstorgnhirarchyname,(select typename from mstrecorddifferentiationtype where id=fromtickettypedifftypeid and deleteflg =0 and activeflg=1) as Fromtickettypedifftypename,(select name from mstrecorddifferentiation where id=fromtickettypediffid and deleteflg =0 and activeflg=1) as Fromtickettypediffname,(select typename from mstrecorddifferentiationtype where id=fromcatdifftypeid and deleteflg =0 and activeflg=1) as Fromcatdifftypename,(select typename from mstrecorddifferentiationtype where id=fromcatlabelid and deleteflg =0 and activeflg=1) as Fromcatlabelname,(select name from mstrecorddifferentiation where id=fromcatdiffid and deleteflg =0 and activeflg=1) as Fromcatdiffname,(select typename from mstrecorddifferentiationtype where id=totickettypedifftypeid and deleteflg =0 and activeflg=1) as Totickettypedifftypename,(select name from mstrecorddifferentiation where id=totickettypediffid and deleteflg =0 and activeflg=1) as Totickettypediffname,(select typename from mstrecorddifferentiationtype where id=tocatdifftypeid and deleteflg =0 and activeflg=1) as Tocatdifftypename,(select typename from mstrecorddifferentiationtype where id=tocatlabelid and deleteflg =0 and activeflg=1) as Tocatlabelname,(select name from mstrecorddifferentiation where id=tocatdiffid and deleteflg =0 and activeflg=1) as Tocatdiffnam FROM mstcategorytaskmap WHERE clientid = ? AND deleteflg =0 and activeflg=1 ORDER BY id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else {
		getMstcategorytaskmap = "SELECT id as Id, clientid as Clientid, mstorgnhirarchyid as Mstorgnhirarchyid, fromtickettypedifftypeid as Fromtickettypedifftypeid, fromtickettypediffid as Fromtickettypediffid, fromcatdifftypeid as Fromcatdifftypeid, fromcatlabelid as Fromcatlabelid, fromcatdiffid as Fromcatdiffid, totickettypedifftypeid as Totickettypedifftypeid, totickettypediffid as Totickettypediffid, tocatdifftypeid as Tocatdifftypeid, tocatlabelid as Tocatlabelid, tocatdiffid as Tocatdiffid, activeflg as Activeflg,(select name from mstclient where id=clientid) as Clientname,(select name from mstorgnhierarchy where id=mstorgnhirarchyid) as Mstorgnhirarchyname,(select typename from mstrecorddifferentiationtype where id=fromtickettypedifftypeid and deleteflg =0 and activeflg=1) as Fromtickettypedifftypename,(select name from mstrecorddifferentiation where id=fromtickettypediffid and deleteflg =0 and activeflg=1) as Fromtickettypediffname,(select typename from mstrecorddifferentiationtype where id=fromcatdifftypeid and deleteflg =0 and activeflg=1) as Fromcatdifftypename,(select typename from mstrecorddifferentiationtype where id=fromcatlabelid and deleteflg =0 and activeflg=1) as Fromcatlabelname,(select name from mstrecorddifferentiation where id=fromcatdiffid and deleteflg =0 and activeflg=1) as Fromcatdiffname,(select typename from mstrecorddifferentiationtype where id=totickettypedifftypeid and deleteflg =0 and activeflg=1) as Totickettypedifftypename,(select name from mstrecorddifferentiation where id=totickettypediffid and deleteflg =0 and activeflg=1) as Totickettypediffname,(select typename from mstrecorddifferentiationtype where id=tocatdifftypeid and deleteflg =0 and activeflg=1) as Tocatdifftypename,(select typename from mstrecorddifferentiationtype where id=tocatlabelid and deleteflg =0 and activeflg=1) as Tocatlabelname,(select name from mstrecorddifferentiation where id=tocatdiffid and deleteflg =0 and activeflg=1) as Tocatdiffnam FROM mstcategorytaskmap WHERE clientid = ? AND mstorgnhirarchyid = ?  AND deleteflg =0 and activeflg=1 ORDER BY id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Mstorgnhirarchyid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	}
	rows, err := dbc.DB.Query(getMstcategorytaskmap, params...)

	// rows, err := dbc.DB.Query(getMstcategorytaskmap, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstcategorytaskmap Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstcategorytaskmapEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Fromtickettypedifftypeid, &value.Fromtickettypediffid, &value.Fromcatdifftypeid, &value.Fromcatlabelid, &value.Fromcatdiffid, &value.Totickettypedifftypeid, &value.Totickettypediffid, &value.Tocatdifftypeid, &value.Tocatlabelid, &value.Tocatdiffid, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Fromtickettypedifftypename, &value.Fromtickettypediffname, &value.Fromcatdifftypename, &value.Fromcatlabelname, &value.Fromcatdiffname, &value.Totickettypedifftypename, &value.Totickettypediffname, &value.Tocatdifftypename, &value.Tocatlabelname, &value.Tocatdiffnam)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateMstcategorytaskmap(tz *entities.MstcategorytaskmapEntity) error {
	logger.Log.Println("In side UpdateMstcategorytaskmap")
	stmt, err := dbc.DB.Prepare(updateMstcategorytaskmap)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMstcategorytaskmap Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Fromtickettypedifftypeid, tz.Fromtickettypediffid, tz.Fromcatdifftypeid, tz.Fromcatlabelid, tz.Fromcatdiffid, tz.Totickettypedifftypeid, tz.Totickettypediffid, tz.Tocatdifftypeid, tz.Tocatlabelid, tz.Tocatdiffid, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMstcategorytaskmap Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMstcategorytaskmap(tz *entities.MstcategorytaskmapEntity) error {
	logger.Log.Println("In side DeleteMstcategorytaskmap")
	stmt, err := dbc.DB.Prepare(deleteMstcategorytaskmap)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMstcategorytaskmap Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMstcategorytaskmap Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMstcategorytaskmapCount(tz *entities.MstcategorytaskmapEntity, OrgnTypeID int64) (entities.MstcategorytaskmapEntities, error) {
	logger.Log.Println("In side GetMstcategorytaskmapCount")
	value := entities.MstcategorytaskmapEntities{}

	var getMstcategorytaskmapcount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getMstcategorytaskmapcount = "SELECT count(a.id) total FROM  mstcategorytaskmap a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstrecorddifferentiationtype f, mstrecorddifferentiationtype g, mstrecorddifferentiation h, mstrecorddifferentiationtype i, mstrecorddifferentiation j, mstrecorddifferentiationtype k, mstrecorddifferentiationtype l, mstrecorddifferentiation m WHERE   a.deleteflg =0 and a.activeflg=1 and b.id=a.clientid and c.id=a.mstorgnhirarchyid and d.id=a.fromtickettypedifftypeid and d.deleteflg =0 and d.activeflg=1 and e.id=a.fromtickettypediffid and e.deleteflg =0 and e.activeflg=1 and f.id=a.fromcatdifftypeid and f.deleteflg =0 and f.activeflg=1 and g.id=a.fromcatlabelid and g.deleteflg =0 and g.activeflg=1 and h.id=a.fromcatdiffid and h.deleteflg =0 and h.activeflg=1 and i.id=a.totickettypedifftypeid and i.deleteflg =0 and i.activeflg=1 and j.id=a.totickettypediffid and j.deleteflg =0 and j.activeflg=1 and k.id=a.tocatdifftypeid and k.deleteflg =0 and k.activeflg=1 and l.id=a.tocatlabelid and l.deleteflg =0 and l.activeflg=1 and m.id=a.tocatdiffid and m.deleteflg =0 and m.activeflg=1"
	} else if OrgnTypeID == 2 {
		getMstcategorytaskmapcount = "SELECT count(a.id) total FROM  mstcategorytaskmap a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstrecorddifferentiationtype f, mstrecorddifferentiationtype g, mstrecorddifferentiation h, mstrecorddifferentiationtype i, mstrecorddifferentiation j, mstrecorddifferentiationtype k, mstrecorddifferentiationtype l, mstrecorddifferentiation m WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 and b.id=a.clientid and c.id=a.mstorgnhirarchyid and d.id=a.fromtickettypedifftypeid and d.deleteflg =0 and d.activeflg=1 and e.id=a.fromtickettypediffid and e.deleteflg =0 and e.activeflg=1 and f.id=a.fromcatdifftypeid and f.deleteflg =0 and f.activeflg=1 and g.id=a.fromcatlabelid and g.deleteflg =0 and g.activeflg=1 and h.id=a.fromcatdiffid and h.deleteflg =0 and h.activeflg=1 and i.id=a.totickettypedifftypeid and i.deleteflg =0 and i.activeflg=1 and j.id=a.totickettypediffid and j.deleteflg =0 and j.activeflg=1 and k.id=a.tocatdifftypeid and k.deleteflg =0 and k.activeflg=1 and l.id=a.tocatlabelid and l.deleteflg =0 and l.activeflg=1 and m.id=a.tocatdiffid and m.deleteflg =0 and m.activeflg=1"
		params = append(params, tz.Clientid)
	} else {
		getMstcategorytaskmapcount = "SELECT count(a.id) total FROM  mstcategorytaskmap a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstrecorddifferentiationtype f, mstrecorddifferentiationtype g, mstrecorddifferentiation h, mstrecorddifferentiationtype i, mstrecorddifferentiation j, mstrecorddifferentiationtype k, mstrecorddifferentiationtype l, mstrecorddifferentiation m WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND  a.deleteflg =0 and a.activeflg=1 and b.id=a.clientid and c.id=a.mstorgnhirarchyid and d.id=a.fromtickettypedifftypeid and d.deleteflg =0 and d.activeflg=1 and e.id=a.fromtickettypediffid and e.deleteflg =0 and e.activeflg=1 and f.id=a.fromcatdifftypeid and f.deleteflg =0 and f.activeflg=1 and g.id=a.fromcatlabelid and g.deleteflg =0 and g.activeflg=1 and h.id=a.fromcatdiffid and h.deleteflg =0 and h.activeflg=1 and i.id=a.totickettypedifftypeid and i.deleteflg =0 and i.activeflg=1 and j.id=a.totickettypediffid and j.deleteflg =0 and j.activeflg=1 and k.id=a.tocatdifftypeid and k.deleteflg =0 and k.activeflg=1 and l.id=a.tocatlabelid and l.deleteflg =0 and l.activeflg=1 and m.id=a.tocatdiffid and m.deleteflg =0 and m.activeflg=1"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getMstcategorytaskmapcount, params...).Scan(&value.Total)

	// err := dbc.DB.QueryRow(getMstcategorytaskmapcount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMstcategorytaskmapCount Get Statement Prepare Error", err)
		return value, err
	}
}
