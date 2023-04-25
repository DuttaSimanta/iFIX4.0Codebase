package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

//id, clientid, mstorgnhirarchyid, name, metertypeid, seqno
var insertSlaTermEntry = "INSERT INTO mstslaindicatorterm ( clientid, mstorgnhirarchyid, name, metertypeid, seqno) VALUES (?,?,?,?,?)"
var duplicateSlaTermEntry = "SELECT count(id) total FROM  mstslaindicatorterm WHERE clientid = ? AND mstorgnhirarchyid = ?  AND name=? AND metertypeid=? AND activeflg =1 AND deleteflg = 0 "
var getSlaTermEntry = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.name as metername, a.metertypeid as metertypeid, a.seqno as seqno,a.activeflg as Activeflg,b.name  as clientname,c.name as mstorgnhirarchyname,d.name as metertypename FROM mstslaindicatorterm a,mstclient b,mstorgnhierarchy c,mstslametertype d WHERE  a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.metertypeid=d.id ORDER BY a.id DESC LIMIT ?,?"
var getSlaTermEntrycount = "SELECT count(a.id) as total FROM mstslaindicatorterm a,mstclient b,mstorgnhierarchy c,mstslametertype d WHERE  a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.metertypeid=d.id"

//var updateSlaTermEntry = "UPDATE mstexceltemplate SET clientid=?,mstorgnhirarchyid = ?, recorddifftypeid = ?, headername = ?, seqno = ?,templatetypeid=?,recorddiffid=? WHERE id = ? "
var deleteSlaTermEntry = "UPDATE mstslaindicatorterm SET deleteflg ='1' WHERE id = ? "
var getrow = "SELECT  a.seqno as seqno from mstslaindicatorterm a where a.clientid=? AND a.mstorgnhirarchyid=? AND a.name=? AND a.metertypeid=? "

func (dbc DbConn) CheckDuplicateSlaTermEntry(tz *entities.SlaTermEntryEntity, i int) (entities.SlaTermEntryEntities, error) {
	logger.Log.Println("In side CheckDuplicateSlaTermEntity ")
	value := entities.SlaTermEntryEntities{}
	err := dbc.DB.QueryRow(duplicateSlaTermEntry, tz.ToClientid, tz.ToMstorgnhirarchyid, tz.MeterNames[i], tz.MetertTypeid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateSlaTermEntity Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc TxConn) AddSlaTermEntry(tz *entities.SlaTermEntryEntity, i int) (int64, error) {
	logger.Log.Println("In side AddSlaTermEntry")
	logger.Log.Println("Query -->", insertSlaTermEntry)
	stmt, err := dbc.TX.Prepare(insertSlaTermEntry)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("AddSlaTermEntity Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.ToClientid, tz.ToMstorgnhirarchyid, tz.MeterName, tz.MetertTypeid, tz.Seqno)
	res, err := stmt.Exec(tz.ToClientid, tz.ToMstorgnhirarchyid, tz.MeterNames[i], tz.MetertTypeid, tz.Seqno)
	if err != nil {
		logger.Log.Println("AddSlaTermEntity Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllSlaTermEntry(page *entities.SlaTermEntryEntity) ([]entities.SlaTermEntryEntity, error) {
	logger.Log.Println("In side GetAllSlaTermEntry")
	values := []entities.SlaTermEntryEntity{}

	rows, err := dbc.DB.Query(getSlaTermEntry, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllSlaTermEntity Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.SlaTermEntryEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.MeterName, &value.MetertTypeid, &value.Seqno, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.MetertypeName)
		values = append(values, value)
	}
	return values, nil
}

// func (dbc DbConn) UpdateSlaTermEntity(tz *entities.SlaTermEntryEntity) error {
// 	logger.Log.Println("In side UpdateSlaTermEntity")
// 	stmt, err := dbc.DB.Prepare(SlaTermEntry)
// 	defer stmt.Close()
// 	if err != nil {
// 		logger.Log.Println("UpdateSlaTermEntity Prepare Statement  Error", err)
// 		return err
// 	}
// 	_, err = stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.RecordDiffTypeid, tz.HeaderName, tz.SeqNo, tz.TemplateTypeid, tz.RecordDiffid, tz.Id)
// 	if err != nil {
// 		logger.Log.Println("UpdateSlaTermEntity Execute Statement  Error", err)
// 		return err
// 	}
// 	return nil
// }

func (dbc DbConn) DeleteSlaTermEntry(tz *entities.SlaTermEntryEntity) error {
	logger.Log.Println("In side DeleteSlaTermEntity", tz)
	stmt, err := dbc.DB.Prepare(deleteSlaTermEntry)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteSlaTermEntry Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteSlaTermEntry Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetSlaTermEntryCount(tz *entities.SlaTermEntryEntity) (entities.SlaTermEntryEntities, error) {
	logger.Log.Println("In side GetSlaTermEntityCount")
	value := entities.SlaTermEntryEntities{}
	err := dbc.DB.QueryRow(getSlaTermEntrycount).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetSlaTermEntityCount Get Statement Prepare Error", err)
		return value, err
	}
}
func (dbc DbConn) GetRow(page *entities.SlaTermEntryEntity, i int) (entities.SlaTermEntryEntity, error) {
	logger.Log.Println("In side Getrow")
	value := entities.SlaTermEntryEntity{}

	rows, err := dbc.DB.Query(getrow, page.Clientid, page.Mstorgnhirarchyid, page.MeterNames[i], page.MetertTypeid)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("getrow Get Statement Prepare Error", err)
		return value, err
	}
	// logger.Log.Println("getrow Get Statement Prepare Error", err)
	for rows.Next() {
		//rows.Next()
		//value := entities.SlaTermEntryEntity{}
		rows.Scan(&value.Seqno)
		//values = append(values, value)
	}
	logger.Log.Println("getrow Get Statement Prepare Error", i)
	return value, nil
}
