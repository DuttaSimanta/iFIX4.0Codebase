package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
)

var insertMstrecordterms = "INSERT INTO mstrecordterms (clientid, mstorgnhirarchyid, termname, termtypeid, termvalue,seq) VALUES (?,?,?,?,?,?)"
var duplicateMstrecordterms = "SELECT count(id) total FROM  mstrecordterms WHERE clientid = ? AND mstorgnhirarchyid = ? AND termname like binary ? AND termtypeid=? AND deleteflg = 0"
var duplicateMsttermvalue = "SELECT count(id) total FROM  mstrecordterms WHERE clientid = ? AND mstorgnhirarchyid = ? AND termname like binary ? AND termtypeid=? AND termvalue = ?  AND deleteflg = 0"

//var getMstrecordterms = "SELECT mstrecordterms.id as Id, mstrecordterms.clientid as Clientid, mstrecordterms.mstorgnhirarchyid as Mstorgnhirarchyid, mstrecordterms.termname as Termname, mstrecordterms.termtypeid as Termtypeid, mstrecordterms.termvalue as Termvalue, mstrecordterms.activeflg as Activeflg, mstclient.name as Clientname,mstorgnhierarchy.name as Mstorgnhirarchyname, msttermtype.termtypename AS Termtypename FROM mstrecordterms JOIN msttermtype ON  mstrecordterms.termtypeid=msttermtype.id JOIN mstclient ON  mstrecordterms.clientid=mstclient.id  JOIN mstorgnhierarchy ON  mstrecordterms.mstorgnhirarchyid=mstorgnhierarchy.id AND mstclient.id=mstorgnhierarchy.clientid   WHERE mstrecordterms.deleteflg =0 AND mstrecordterms.activeflg=1 AND mstrecordterms.clientid=? AND mstrecordterms.mstorgnhirarchyid=? ORDER BY mstrecordterms.id DESC LIMIT ?,?"
var getMstrecordtermsList = "SELECT a.id as Id, a.termname as Termname, a.termtypeid as Termtypeid,a.termvalue as Termvalue,a.activeflg as Activeflg,a.seq as Sequance,b.termtypename as Termtypename FROM mstrecordterms a,msttermtype b WHERE a.deleteflg =0 AND a.activeflg=1 AND a.clientid = ? AND a.mstorgnhirarchyid =? AND a.termtypeid=b.id ORDER BY a.id ASC  "

//var getMstrecordtermscount = "SELECT count(id) total FROM  mstrecordterms WHERE deleteflg =0 and activeflg=1 AND mstrecordterms.clientid=? AND mstrecordterms.mstorgnhirarchyid=?"
var updateMstrecordterms = "UPDATE mstrecordterms SET termname = ?, termtypeid = ?, termvalue = ? WHERE id = ? "
var deleteMstrecordterms = "UPDATE mstrecordterms SET deleteflg = '1' WHERE id = ? "
var lastseqbyterms = "SELECT max(seq)  from mstrecordterms where clientid=? and mstorgnhirarchyid = ?  and activeflg=1 and deleteflg=0"

var duplicatetermeseq = "SELECT count(id) total from mstrecordterms where clientid=? and mstorgnhirarchyid = ? and  seq=? and deleteflg=0"

func (mdao DbConn) CheckDuplicatetermseq(tz *entities.MstrecordtermsEntity) (entities.MstrecordtermsEntities, error) {
	log.Println("In side dao")
	value := entities.MstrecordtermsEntities{}
	err := mdao.DB.QueryRow(duplicatetermeseq, tz.Clientid, tz.Mstorgnhirarchyid, tz.Termseq).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		log.Print("CheckDuplicatetermseq Get Statement Prepare Error", err)
		return value, err
	}
}

func (mdao DbConn) GetLastSeqFromterms(tz *entities.MstrecordtermsEntity) ([]entities.TermsEntity, error) {
	log.Println("In side dao")
	values := []entities.TermsEntity{}
	rows, err := mdao.DB.Query(lastseqbyterms, tz.Clientid, tz.Mstorgnhirarchyid)

	if err != nil {
		log.Print("GetLastSeqFromterms Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		value := entities.TermsEntity{}
		rows.Scan(&value.Termseq)
		values = append(values, value)
	}
	return values, nil
}
func (dbc DbConn) CheckDuplicateMstrecordterms(tz *entities.MstrecordtermsEntity) (entities.MstrecordtermsEntities, error) {
	logger.Log.Println("In side CheckDuplicateMstrecordterms")
	value := entities.MstrecordtermsEntities{}
	err := dbc.DB.QueryRow(duplicateMstrecordterms, tz.Clientid, tz.Mstorgnhirarchyid, tz.Termname, tz.Termtypeid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMstrecordterms Get Statement Prepare Error", err)
		return value, err
	}
}

func (mdao DbConn) CheckDuplicatetermval(tz *entities.MstrecordtermsEntity) (entities.MstrecordtermsEntities, error) {
	log.Println("In side CheckDuplicatetermval")
	value := entities.MstrecordtermsEntities{}
	logger.Log.Println(tz.Clientid, tz.Mstorgnhirarchyid, tz.Termname, tz.Termtypeid, tz.Termvalue)
	err := mdao.DB.QueryRow(duplicateMsttermvalue, tz.Clientid, tz.Mstorgnhirarchyid, tz.Termname, tz.Termtypeid, tz.Termvalue).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		log.Print("CheckDuplicatetermval Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertMstrecordterms(tz *entities.MstrecordtermsEntity) (int64, error) {
	logger.Log.Println("In side InsertMstrecordterms")
	logger.Log.Println("Query -->", insertMstrecordterms)
	stmt, err := dbc.DB.Prepare(insertMstrecordterms)

	if err != nil {
		logger.Log.Println("InsertMstrecordterms Prepare Statement  Error", err)
		return 0, err
	}
	defer stmt.Close()
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Termname, tz.Termtypeid, tz.Termvalue)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Termname, tz.Termtypeid, tz.Termvalue, tz.Termseq)
	if err != nil {
		logger.Log.Println("InsertMstrecordterms Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllMstrecordterms(tz *entities.MstrecordtermsEntity, OrgnType int64) ([]entities.MstrecordtermsEntity, error) {
	logger.Log.Println("In side GelAllMstrecordterms")
	values := []entities.MstrecordtermsEntity{}
	var getMstrecordterms string
	var params []interface{}
	if OrgnType == 1 {
		getMstrecordterms = "SELECT mstrecordterms.id as Id, mstrecordterms.clientid as Clientid, mstrecordterms.mstorgnhirarchyid as Mstorgnhirarchyid, mstrecordterms.termname as Termname, mstrecordterms.termtypeid as Termtypeid, mstrecordterms.termvalue as Termvalue, mstrecordterms.activeflg as Activeflg, mstclient.name as Clientname,mstorgnhierarchy.name as Mstorgnhirarchyname, msttermtype.termtypename AS Termtypename FROM mstrecordterms JOIN msttermtype ON  mstrecordterms.termtypeid=msttermtype.id JOIN mstclient ON  mstrecordterms.clientid=mstclient.id  JOIN mstorgnhierarchy ON  mstrecordterms.mstorgnhirarchyid=mstorgnhierarchy.id AND mstclient.id=mstorgnhierarchy.clientid   WHERE mstrecordterms.deleteflg =0 AND mstrecordterms.activeflg=1 ORDER BY mstrecordterms.id DESC LIMIT ?,?"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		getMstrecordterms = "SELECT mstrecordterms.id as Id, mstrecordterms.clientid as Clientid, mstrecordterms.mstorgnhirarchyid as Mstorgnhirarchyid, mstrecordterms.termname as Termname, mstrecordterms.termtypeid as Termtypeid, mstrecordterms.termvalue as Termvalue, mstrecordterms.activeflg as Activeflg, mstclient.name as Clientname,mstorgnhierarchy.name as Mstorgnhirarchyname, msttermtype.termtypename AS Termtypename FROM mstrecordterms JOIN msttermtype ON  mstrecordterms.termtypeid=msttermtype.id JOIN mstclient ON  mstrecordterms.clientid=mstclient.id  JOIN mstorgnhierarchy ON  mstrecordterms.mstorgnhirarchyid=mstorgnhierarchy.id AND mstclient.id=mstorgnhierarchy.clientid   WHERE mstrecordterms.deleteflg =0 AND mstrecordterms.activeflg=1 AND mstrecordterms.clientid=? ORDER BY mstrecordterms.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		getMstrecordterms = "SELECT mstrecordterms.id as Id, mstrecordterms.clientid as Clientid, mstrecordterms.mstorgnhirarchyid as Mstorgnhirarchyid, mstrecordterms.termname as Termname, mstrecordterms.termtypeid as Termtypeid, mstrecordterms.termvalue as Termvalue, mstrecordterms.activeflg as Activeflg, mstclient.name as Clientname,mstorgnhierarchy.name as Mstorgnhirarchyname, msttermtype.termtypename AS Termtypename FROM mstrecordterms JOIN msttermtype ON  mstrecordterms.termtypeid=msttermtype.id JOIN mstclient ON  mstrecordterms.clientid=mstclient.id  JOIN mstorgnhierarchy ON  mstrecordterms.mstorgnhirarchyid=mstorgnhierarchy.id AND mstclient.id=mstorgnhierarchy.clientid   WHERE mstrecordterms.deleteflg =0 AND mstrecordterms.activeflg=1 AND mstrecordterms.clientid=? AND mstrecordterms.mstorgnhirarchyid=? ORDER BY mstrecordterms.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}
	rows, err := dbc.DB.Query(getMstrecordterms, params...)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstrecordterms Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstrecordtermsEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Termname, &value.Termtypeid, &value.Termvalue, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Termtypename)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) GetListMstrecordterms(tz *entities.MstrecordtermsEntity) ([]entities.TermsEntity, error) {
	logger.Log.Println("In side GetListMstrecordterms")
	values := []entities.TermsEntity{}
	rows, err := dbc.DB.Query(getMstrecordtermsList, tz.Clientid, tz.Mstorgnhirarchyid)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetListMstrecordterms Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.TermsEntity{}

		rows.Scan(&value.Id, &value.Termname, &value.Termtypeid, &value.Termvalue, &value.Activeflg, &value.Termseq, &value.Termtypename)
		//logger.Log.Println(values)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateMstrecordterms(tz *entities.MstrecordtermsEntity) error {
	logger.Log.Println("In side UpdateMstrecordterms")
	stmt, err := dbc.DB.Prepare(updateMstrecordterms)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMstrecordterms Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Termname, tz.Termtypeid, tz.Termvalue, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMstrecordterms Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMstrecordterms(tz *entities.MstrecordtermsEntity) error {
	logger.Log.Println("In side DeleteMstrecordterms")
	stmt, err := dbc.DB.Prepare(deleteMstrecordterms)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMstrecordterms Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMstrecordterms Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMstrecordtermsCount(tz *entities.MstrecordtermsEntity, OrgnTypeID int64) (entities.MstrecordtermsEntities, error) {
	logger.Log.Println("In side GetMstrecordtermsCount")
	value := entities.MstrecordtermsEntities{}
	var getMstrecordtermscount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getMstrecordtermscount = "SELECT count(id) total FROM  mstrecordterms WHERE deleteflg =0 and activeflg=1"
	} else if OrgnTypeID == 2 {
		getMstrecordtermscount = "SELECT count(id) total FROM  mstrecordterms WHERE deleteflg =0 and activeflg=1 AND mstrecordterms.clientid=?"
		params = append(params, tz.Clientid)
	} else {
		getMstrecordtermscount = "SELECT count(id) total FROM  mstrecordterms WHERE deleteflg =0 and activeflg=1 AND mstrecordterms.clientid=? AND mstrecordterms.mstorgnhirarchyid=?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getMstrecordtermscount, params...).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMstrecordtermsCount Get Statement Prepare Error", err)
		return value, err
	}
}
