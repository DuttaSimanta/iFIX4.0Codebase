package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertMstExcelTemplate = "INSERT INTO mstexceltemplate ( clientid, mstorgnhirarchyid, recorddifftypeid, headername, seqno, templatetypeid, recorddiffid) VALUES (?,?,?,?,?,?,?)"
var duplicateMstExcelTemplate = "SELECT count(id) total FROM  mstexceltemplate WHERE clientid = ? AND mstorgnhirarchyid = ?  AND recorddifftypeid=? AND headername=? AND seqno=? AND templatetypeid=? AND recorddiffid=? AND activeflg =1 AND deleteflg = 0 "

// var getMstExcelTemplate = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.recorddifftypeid as RecordDiffTypeid, a.headername as HeaderName, a.seqno as SeqNo,a.templatetypeid as TemplateTypeid,a.recorddiffid as RecordDiffid,a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname ,d.typename as RecordDiffTypeName,e.name as RecordDiffName,f.typename as TemplateTypeName FROM mstexceltemplate a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstexceltemplatetype f WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.recorddifftypeid=d.id AND a.recorddiffid=e.id AND a.templatetypeid=f.id  AND d.activeflg=1 AND d.deleteflg=0 AND e.activeflg=1 AND e.deleteflg=0  ORDER BY a.id DESC LIMIT ?,?"
// var getMstExcelTemplatecount = "SELECT count(a.id) as total FROM mstexceltemplate a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstexceltemplatetype f WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.recorddifftypeid=d.id AND a.recorddiffid=e.id AND a.templatetypeid=f.id  AND d.activeflg=1 AND d.deleteflg=0 AND e.activeflg=1 AND e.deleteflg=0 "
var updateMstExcelTemplate = "UPDATE mstexceltemplate SET clientid=?,mstorgnhirarchyid = ?, recorddifftypeid = ?, headername = ?, seqno = ?,templatetypeid=?,recorddiffid=? WHERE id = ? "
var deleteMstExcelTemplate = "UPDATE mstexceltemplate SET deleteflg ='1' WHERE id = ? "

func (dbc DbConn) CheckDuplicateMstExcelTemplate(tz *entities.MstExcelTemplateEntity) (entities.MstExcelTemplateEntities, error) {
	logger.Log.Println("In side CheckDuplicateMstExcelTemplate ")
	value := entities.MstExcelTemplateEntities{}
	err := dbc.DB.QueryRow(duplicateMstExcelTemplate, tz.Clientid, tz.Mstorgnhirarchyid, tz.RecordDiffTypeid, tz.HeaderName, tz.SeqNo, tz.TemplateTypeid, tz.RecordDiffid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMstExcelTemplate Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) AddMstExcelTemplate(tz *entities.MstExcelTemplateEntity) (int64, error) {
	logger.Log.Println("In side AddMstExcelTemplate")
	logger.Log.Println("Query -->", insertMstExcelTemplate)
	stmt, err := dbc.DB.Prepare(insertMstExcelTemplate)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("AddMstMstExcelTemplate Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.RecordDiffTypeid, tz.HeaderName, tz.SeqNo, tz.TemplateTypeid, tz.RecordDiffid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.RecordDiffTypeid, tz.HeaderName, tz.SeqNo, tz.TemplateTypeid, tz.RecordDiffid)
	if err != nil {
		logger.Log.Println("AddMstMstExcelTemplate Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllMstExcelTemplate(page *entities.MstExcelTemplateEntity, OrgnType int64) ([]entities.MstExcelTemplateEntity, error) {
	logger.Log.Println("In side GetAllMstExcelTemplate")
	values := []entities.MstExcelTemplateEntity{}
	var getMstExcelTemplate string
	var params []interface{}
	if OrgnType == 1 {
		getMstExcelTemplate = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.recorddifftypeid as RecordDiffTypeid, a.headername as HeaderName, a.seqno as SeqNo,a.templatetypeid as TemplateTypeid,a.recorddiffid as RecordDiffid,a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname ,d.typename as RecordDiffTypeName,e.name as RecordDiffName,f.typename as TemplateTypeName FROM mstexceltemplate a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstexceltemplatetype f WHERE a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.recorddifftypeid=d.id AND a.recorddiffid=e.id AND a.templatetypeid=f.id  AND d.activeflg=1 AND d.deleteflg=0 AND e.activeflg=1 AND e.deleteflg=0  ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else if OrgnType == 2 {
		getMstExcelTemplate = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.recorddifftypeid as RecordDiffTypeid, a.headername as HeaderName, a.seqno as SeqNo,a.templatetypeid as TemplateTypeid,a.recorddiffid as RecordDiffid,a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname ,d.typename as RecordDiffTypeName,e.name as RecordDiffName,f.typename as TemplateTypeName FROM mstexceltemplate a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstexceltemplatetype f WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.recorddifftypeid=d.id AND a.recorddiffid=e.id AND a.templatetypeid=f.id  AND d.activeflg=1 AND d.deleteflg=0 AND e.activeflg=1 AND e.deleteflg=0  ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else {
		getMstExcelTemplate = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.recorddifftypeid as RecordDiffTypeid, a.headername as HeaderName, a.seqno as SeqNo,a.templatetypeid as TemplateTypeid,a.recorddiffid as RecordDiffid,a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname ,d.typename as RecordDiffTypeName,e.name as RecordDiffName,f.typename as TemplateTypeName FROM mstexceltemplate a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstexceltemplatetype f WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.recorddifftypeid=d.id AND a.recorddiffid=e.id AND a.templatetypeid=f.id  AND d.activeflg=1 AND d.deleteflg=0 AND e.activeflg=1 AND e.deleteflg=0  ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Mstorgnhirarchyid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	}
	rows, err := dbc.DB.Query(getMstExcelTemplate, params...)

	// rows, err := dbc.DB.Query(getMstExcelTemplate, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstExcelTemplate Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstExcelTemplateEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.RecordDiffTypeid, &value.HeaderName, &value.SeqNo, &value.TemplateTypeid, &value.RecordDiffid, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.RecordDiffTypeName, &value.RecordDiffName, &value.TemplateTypeName)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateMstExcelTemplate(tz *entities.MstExcelTemplateEntity) error {
	logger.Log.Println("In side UpdateMstExcelTemplate")
	stmt, err := dbc.DB.Prepare(updateMstExcelTemplate)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMstExcelTemplate Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.RecordDiffTypeid, tz.HeaderName, tz.SeqNo, tz.TemplateTypeid, tz.RecordDiffid, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMstExcelTemplate Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMstExcelTemplate(tz *entities.MstExcelTemplateEntity) error {
	logger.Log.Println("In side DeleteMstExcelTemplate", tz)
	stmt, err := dbc.DB.Prepare(deleteMstExcelTemplate)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMstExcelTemplate Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMstExcelTemplate Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMstExcelTemplateCount(tz *entities.MstExcelTemplateEntity, OrgnTypeID int64) (entities.MstExcelTemplateEntities, error) {
	logger.Log.Println("In side GetMstExcelTemplateCount")
	value := entities.MstExcelTemplateEntities{}

	var getMstExcelTemplatecount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getMstExcelTemplatecount = "SELECT count(a.id) as total FROM mstexceltemplate a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstexceltemplatetype f WHERE a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.recorddifftypeid=d.id AND a.recorddiffid=e.id AND a.templatetypeid=f.id  AND d.activeflg=1 AND d.deleteflg=0 AND e.activeflg=1 AND e.deleteflg=0 "
	} else if OrgnTypeID == 2 {
		getMstExcelTemplatecount = "SELECT count(a.id) as total FROM mstexceltemplate a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstexceltemplatetype f WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.recorddifftypeid=d.id AND a.recorddiffid=e.id AND a.templatetypeid=f.id  AND d.activeflg=1 AND d.deleteflg=0 AND e.activeflg=1 AND e.deleteflg=0 "
		params = append(params, tz.Clientid)
	} else {
		getMstExcelTemplatecount = "SELECT count(a.id) as total FROM mstexceltemplate a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstexceltemplatetype f WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.recorddifftypeid=d.id AND a.recorddiffid=e.id AND a.templatetypeid=f.id  AND d.activeflg=1 AND d.deleteflg=0 AND e.activeflg=1 AND e.deleteflg=0 "
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getMstExcelTemplatecount, params...).Scan(&value.Total)

	// err := dbc.DB.QueryRow(getMstExcelTemplatecount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMstExcelTemplateCount Get Statement Prepare Error", err)
		return value, err
	}
}
