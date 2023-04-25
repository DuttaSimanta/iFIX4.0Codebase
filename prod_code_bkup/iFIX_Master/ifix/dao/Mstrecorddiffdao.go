package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
)

var getdifftype = "SELECT id as ID ,typename as Typename,seqno as Seqno,istextfield from mstrecorddifferentiationtype where seqno < 99 and activeflg=1 and deleteflg=0 order by typename"
var insertdiff = "INSERT into mstrecorddifferentiation(clientid,mstorgnhirarchyid,recorddifftypeid,parentid,seqno,name) values(?,?,?,?,?,?)"
var lastseqbydifftype = "SELECT max(seqno)  from mstrecorddifferentiation where clientid=? and mstorgnhirarchyid=? and recorddifftypeid=?  and activeflg=1 and deleteflg=0"
var duplicaterecord = "SELECT count(id) total from mstrecorddifferentiation where clientid=? and mstorgnhirarchyid=? and recorddifftypeid=? and parentid=?  and name=? and deleteflg=0"
var recorddiffcount = "SELECT count(id) total from mstrecorddifferentiation where clientid=? and mstorgnhirarchyid=? and recorddifftypeid=? and activeflg=1 and deleteflg=0"
var getrecord = "SELECT id as ID,name as Name,activeflg as Activeflg,seqno as Seqno from mstrecorddifferentiation where clientid=? and mstorgnhirarchyid=? and recorddifftypeid=? and activeflg=1 and deleteflg=0 ORDER BY id DESC LIMIT ?,?"
var updaterecorddiff = "UPDATE mstrecorddifferentiation set name=? where id=?"
var updateassetrecorddiff = "UPDATE mstrecorddifferentiation set clientid=?, mstorgnhirarchyid=?, name=?, recorddifftypeid=?, parentid = ? where id=?"
var deleterecorddiff = "UPDATE mstrecorddifferentiation set deleteflg=1 where id=?"
var recordbytype = "SELECT id as ID,name as Typename,seqno,recorddifftypeid,coalesce(parentcategorynames,'') from mstrecorddifferentiation WHERE clientid=? and mstorgnhirarchyid=? and recorddifftypeid in (SELECT id from mstrecorddifferentiationtype where seqno=? and deleteflg=0 and activeflg=1) and activeflg=1 and deleteflg=0 order by name"
var categorylabel = "SELECT id as ID ,typename as Typename,seqno as Seqno from mstrecorddifferentiationtype where clientid=? and mstorgnhirarchyid=? and parentid in (select id from mstrecorddifferentiationtype where seqno=? and activeflg and deleteflg=0 ) and activeflg=1 and deleteflg=0 order by typename"
var getrecordall = " SELECT a.id as ID, a.name as Name,a.activeflg as Activeflg,b.name as Clientname, d.name as Orgname,c.typename as type from mstrecorddifferentiation a,mstclient b, mstorgnhierarchy d,mstrecorddifferentiationtype c  where a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.recorddifftypeid =c.id and a.activeflg=1 and a.deleteflg=0  and c.deleteflg=0 ORDER BY a.id DESC LIMIT ?,?"
var getrecordbyorg = " SELECT a.id as ID, a.name as Name,a.activeflg as Activeflg,b.name as Clientname, d.name as Orgname,c.typename as type from mstrecorddifferentiation a,mstclient b, mstorgnhierarchy d,mstrecorddifferentiationtype c  where a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.recorddifftypeid =c.id and a.clientid =? AND a.mstorgnhirarchyid =? and a.activeflg=1 and a.deleteflg=0  and c.deleteflg=0 ORDER BY a.id DESC LIMIT ?,?"
var getrecordallcount = "SELECT count(a.id) total from mstrecorddifferentiation a,mstclient b, mstorgnhierarchy d,mstrecorddifferentiationtype c  where a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.recorddifftypeid =c.id  and a.activeflg=1 and a.deleteflg=0  and c.deleteflg=0"
var getrecordcountbyorg = "SELECT count(a.id) total from mstrecorddifferentiation a,mstclient b, mstorgnhierarchy d,mstrecorddifferentiationtype c  where a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.recorddifftypeid =c.id  and a.clientid =? AND a.mstorgnhirarchyid =? AND a.activeflg=1 and a.deleteflg=0  and c.deleteflg=0"
var allcategorylabel = "SELECT id as ID ,typename as Typename,seqno as Seqno from mstrecorddifferentiationtype where clientid=? and mstorgnhirarchyid=? AND seqno > 99 and activeflg=1 and deleteflg=0"
var categorylabels = "SELECT a.id as ID ,a.typename as Typename,a.seqno as Seqno from mstrecorddifferentiationtype a where a.id in(select b.torecorddifftypeid from mstrecordtype b where  b.fromrecorddiffid=? and b.fromrecorddifftypeid=? and b.torecorddiffid=0 and  b.torecorddifftypeid in( SELECT c.id from mstrecorddifferentiationtype c where c.clientid=? and c.mstorgnhirarchyid=? and c.parentid in (select d.id from mstrecorddifferentiationtype d where d.seqno=? and d.activeflg=1 and d.deleteflg=0 ) and c.activeflg=1 and c.deleteflg=0) and b.activeflg=1 and b.deleteflg=0) and a.activeflg=1 and a.deleteflg=0"

func (mdao DbConn) GetRecordDiffType() ([]entities.MstrecorddifftypeEntity, error) {
	log.Println("In side dao")
	values := []entities.MstrecorddifftypeEntity{}
	rows, err := mdao.DB.Query(getdifftype)
	defer rows.Close()
	if err != nil {
		log.Print("GetRecordDiffType Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstrecorddifftypeEntity{}
		rows.Scan(&value.ID, &value.Typename, &value.Seqno, &value.Istextfield)
		values = append(values, value)
	}
	return values, nil
}
func (mdao DbConn) GetRecordByDiffType(tz *entities.RecordDiffEntity) ([]entities.MstrecorddifftypeEntity, error) {
	log.Println("In side dao")
	values := []entities.MstrecorddifftypeEntity{}
	rows, err := mdao.DB.Query(recordbytype, tz.Clientid, tz.Mstorgnhirarchyid, tz.Seqno)

	if err != nil {
		logger.Log.Print("GetRecordDiffType Get Statement Prepare Error", err)
		log.Print("GetRecordDiffType Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		value := entities.MstrecorddifftypeEntity{}
		rows.Scan(&value.ID, &value.Typename, &value.Seqno, &value.Recorddifftypeid, &value.Parentname)
		values = append(values, value)
	}
	return values, nil
}
func (mdao DbConn) GetAllCategoryLevel(tz *entities.RecordDiffEntity) ([]entities.MstrecorddifftypeEntity, error) {
	log.Println("In side dao")
	values := []entities.MstrecorddifftypeEntity{}
	rows, err := mdao.DB.Query(allcategorylabel, tz.Clientid, tz.Mstorgnhirarchyid)
	defer rows.Close()
	if err != nil {
		logger.Log.Print("GetAllCategoryLevel Get Statement Prepare Error", err)
		log.Print("GetAllCategoryLevel Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstrecorddifftypeEntity{}
		rows.Scan(&value.ID, &value.Typename, &value.Seqno)
		values = append(values, value)
	}
	return values, nil
}
func (mdao DbConn) GetCategoryLevel(tz *entities.RecordDiffEntity) ([]entities.MstrecorddifftypeEntity, error) {
	log.Println("In side dao")
	values := []entities.MstrecorddifftypeEntity{}
	rows, err := mdao.DB.Query(categorylabel, tz.Clientid, tz.Mstorgnhirarchyid, tz.Seqno)
	defer rows.Close()
	if err != nil {
		log.Print("GetCategoryLevel Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstrecorddifftypeEntity{}
		rows.Scan(&value.ID, &value.Typename, &value.Seqno)
		values = append(values, value)
	}
	return values, nil
}
func (mdao DbConn) GetCategoriesLevel(tz *entities.RecordDiffEntity) ([]entities.MstrecorddifftypeEntity, error) {
	log.Println("In side dao")
	values := []entities.MstrecorddifftypeEntity{}
	rows, err := mdao.DB.Query(categorylabels, tz.Recorddiffid, tz.Recorddifftypeid, tz.Clientid, tz.Mstorgnhirarchyid, tz.Seqno)
	defer rows.Close()
	if err != nil {
		log.Print("GetCategoriesLevel Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstrecorddifftypeEntity{}
		rows.Scan(&value.ID, &value.Typename, &value.Seqno)
		values = append(values, value)
	}
	return values, nil
}
func (mdao DbConn) GetLastSeqFromRecordDiff(tz *entities.RecordDiffEntity) ([]entities.RecordDiffEntity, error) {
	log.Println("In side dao")
	values := []entities.RecordDiffEntity{}
	rows, err := mdao.DB.Query(lastseqbydifftype, tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid)
	defer rows.Close()
	if err != nil {
		log.Print("GetLastSeqFromRecordDiff Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.RecordDiffEntity{}
		rows.Scan(&value.Seqno)
		values = append(values, value)
	}
	return values, nil
}
func (mdao DbConn) CheckDuplicateRecordDiff(tz *entities.RecordDiffEntity) (entities.RecordDiffEntities, error) {
	log.Println("In side dao")
	value := entities.RecordDiffEntities{}
	err := mdao.DB.QueryRow(duplicaterecord, tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Parentid, tz.Name).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		log.Print("CheckDuplicateRecordDiff Get Statement Prepare Error", err)
		return value, err
	}
}
func (mdao DbConn) InsertRecordDiff(tz *entities.RecordDiffEntity) (int64, error) {
	log.Println("In side dao")
	stmt, err := mdao.DB.Prepare(insertdiff)
	defer stmt.Close()
	if err != nil {
		log.Print("InsertRecordDiff Prepare Statement  Error", err)
		return 0, err
	}
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Parentid, tz.Seqno, tz.Name)
	if err != nil {
		log.Print("InsertRecordDiff Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}
func (mdao DbConn) GetRecordDiffCount(tz *entities.RecordDiffEntity) (entities.RecordDiffEntities, error) {
	log.Println("In side dao")
	value := entities.RecordDiffEntities{}
	err := mdao.DB.QueryRow(recorddiffcount, tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		log.Print("GetRecordDiffCount Get Statement Prepare Error", err)
		return value, err
	}
}
func (mdao DbConn) GetAllRecordDiff(tz *entities.RecordDiffEntity) ([]entities.RecordDiffEntityResp, error) {
	log.Println("In side dao")
	values := []entities.RecordDiffEntityResp{}
	rows, err := mdao.DB.Query(getrecord, tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Offset, tz.Limit)
	defer rows.Close()
	if err != nil {
		log.Print("GetAllRecordDiff Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.RecordDiffEntityResp{}
		rows.Scan(&value.ID, &value.Name, &value.Activeflg, &value.Seqno)
		values = append(values, value)
	}
	return values, nil
}
func (mdao DbConn) GetAllRecordDiffCount() (entities.RecordDiffEntities, error) {
	log.Println("In side dao")
	value := entities.RecordDiffEntities{}
	err := mdao.DB.QueryRow(getrecordallcount).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		log.Print("GetRecordDiffCount Get Statement Prepare Error", err)
		return value, err
	}
}
func (mdao DbConn) GetRecordDiffCountByOrg(tz *entities.RecordDiffEntity, OrgnTypeID int64) (entities.RecordDiffEntities, error) {
	log.Println("In side dao")
	value := entities.RecordDiffEntities{}
	var getrecordcountbyorg string
	var params []interface{}
	if OrgnTypeID == 1 {
		getrecordcountbyorg = "SELECT count(a.id) total from mstrecorddifferentiation a,mstclient b, mstorgnhierarchy d,mstrecorddifferentiationtype c  where a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.recorddifftypeid =c.id  AND a.activeflg=1 and a.deleteflg=0  and c.deleteflg=0"
	} else if OrgnTypeID == 2 {
		getrecordcountbyorg = "SELECT count(a.id) total from mstrecorddifferentiation a,mstclient b, mstorgnhierarchy d,mstrecorddifferentiationtype c  where a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.recorddifftypeid =c.id  and a.clientid =? AND a.activeflg=1 and a.deleteflg=0  and c.deleteflg=0"
		params = append(params, tz.Clientid)
	} else {
		getrecordcountbyorg = "SELECT count(a.id) total from mstrecorddifferentiation a,mstclient b, mstorgnhierarchy d,mstrecorddifferentiationtype c  where a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.recorddifftypeid =c.id  and a.clientid =? AND a.mstorgnhirarchyid =? AND a.activeflg=1 and a.deleteflg=0  and c.deleteflg=0"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := mdao.DB.QueryRow(getrecordcountbyorg, params...).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		log.Print("GetRecordDiffCountByOrg Get Statement Prepare Error", err)
		return value, err
	}
}
func (mdao DbConn) GetRecordDiff(tz *entities.RecordDiffEntity) ([]entities.RecordDiffEntityResp, error) {
	log.Println("In side dao")
	values := []entities.RecordDiffEntityResp{}
	rows, err := mdao.DB.Query(getrecordall, tz.Offset, tz.Limit)
	defer rows.Close()
	if err != nil {
		log.Print("GetAllRecordDiff Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.RecordDiffEntityResp{}
		err := rows.Scan(&value.ID, &value.Name, &value.Activeflg, &value.Clientname, &value.Orgname, &value.Type)
		if err != nil {
			log.Print("GetAllRecordDiff Get Statement Prepare Error", err)
		}
		values = append(values, value)
	}
	log.Println(values)
	return values, nil
}
func (mdao DbConn) GetRecordDiffByOrg(tz *entities.RecordDiffEntity, OrgnType int64) ([]entities.RecordDiffEntityResp, error) {
	log.Println("In side dao")
	values := []entities.RecordDiffEntityResp{}
	var getrecordbyorg string
	var params []interface{}
	if OrgnType == 1 {
		getrecordbyorg = " SELECT a.id as ID,a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.recorddifftypeid as Recorddifftypeid, a.name as Name,a.activeflg as Activeflg,b.name as Clientname, d.name as Orgname,c.typename as type from mstrecorddifferentiation a,mstclient b, mstorgnhierarchy d,mstrecorddifferentiationtype c  where a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.recorddifftypeid =c.id and a.activeflg=1 and a.deleteflg=0  and c.deleteflg=0 ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		getrecordbyorg = " SELECT a.id as ID, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.recorddifftypeid as Recorddifftypeid,  a.name as Name,a.activeflg as Activeflg,b.name as Clientname, d.name as Orgname,c.typename as type from mstrecorddifferentiation a,mstclient b, mstorgnhierarchy d,mstrecorddifferentiationtype c  where a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.recorddifftypeid =c.id and a.clientid =? and a.activeflg=1 and a.deleteflg=0  and c.deleteflg=0  ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		getrecordbyorg = " SELECT a.id as ID, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.recorddifftypeid as Recorddifftypeid,  a.name as Name,a.activeflg as Activeflg, b.name as Clientname, d.name as Orgname,c.typename as type from mstrecorddifferentiation a,mstclient b, mstorgnhierarchy d,mstrecorddifferentiationtype c  where a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.recorddifftypeid =c.id and a.clientid =? AND a.mstorgnhirarchyid =? and a.activeflg=1 and a.deleteflg=0  and c.deleteflg=0 ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}
	rows, err := mdao.DB.Query(getrecordbyorg, params...)
	defer rows.Close()
	if err != nil {
		log.Print("GetRecordDiffByOrg Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.RecordDiffEntityResp{}
		err := rows.Scan(&value.ID, &value.Clientid, &value.Mstorgnhirarchyid, &value.Recorddifftypeid, &value.Name, &value.Activeflg, &value.Clientname, &value.Orgname, &value.Type)
		if err != nil {
			log.Print("GetRecordDiffByOrg Get Statement Prepare Error", err)
		}
		values = append(values, value)
	}
	log.Println(values)
	return values, nil
}

func (mdao DbConn) GetAssetRecordDiffByOrg(tz *entities.RecordDiffEntity, OrgnType int64) ([]entities.RecordDiffEntityResp, error) {
	log.Println("In side dao")
	values := []entities.RecordDiffEntityResp{}
	var getrecordbyorg string
	var params []interface{}
	if OrgnType == 1 {
		getrecordbyorg = " SELECT a.id as ID,a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.recorddifftypeid as Recorddifftypeid, a.name as Name,a.activeflg as Activeflg,b.name as Clientname, d.name as Orgname,c.typename as Type from mstrecorddifferentiation a,mstclient b, mstorgnhierarchy d,mstrecorddifferentiationtype c  where a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.recorddifftypeid =c.id and a.activeflg=1 and a.deleteflg=0  and c.deleteflg=0 and c.parentid=6 ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		getrecordbyorg = " SELECT a.id as ID, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.recorddifftypeid as Recorddifftypeid,  a.name as Name,a.activeflg as Activeflg,b.name as Clientname, d.name as Orgname,c.typename as Type from mstrecorddifferentiation a,mstclient b, mstorgnhierarchy d,mstrecorddifferentiationtype c  where a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.recorddifftypeid =c.id and a.clientid =? and a.activeflg=1 and a.deleteflg=0  and c.deleteflg=0 and c.parentid=6 ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		getrecordbyorg = " SELECT a.id as ID, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.recorddifftypeid as Recorddifftypeid,  a.name as Name,a.activeflg as Activeflg, b.name as Clientname, d.name as Orgname,c.typename as Type from mstrecorddifferentiation a,mstclient b, mstorgnhierarchy d,mstrecorddifferentiationtype c  where a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.recorddifftypeid =c.id and a.clientid =? AND a.mstorgnhirarchyid =? and a.activeflg=1 and a.deleteflg=0  and c.deleteflg=0 and c.parentid=6 ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}
	rows, err := mdao.DB.Query(getrecordbyorg, params...)
	defer rows.Close()
	if err != nil {
		log.Print("GetRecordDiffByOrg Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.RecordDiffEntityResp{}
		err := rows.Scan(&value.ID, &value.Clientid, &value.Mstorgnhirarchyid, &value.Recorddifftypeid, &value.Name, &value.Activeflg, &value.Clientname, &value.Orgname, &value.Type)
		if err != nil {
			log.Print("GetRecordDiffByOrg Get Statement Prepare Error", err)
		}
		values = append(values, value)
	}
	log.Println(values)
	return values, nil
}

func (mdao DbConn) UpdateRecordDiff(tz *entities.RecordDiffEntity) error {
	log.Println("In side dao")
	stmt, err := mdao.DB.Prepare(updaterecorddiff)
	defer stmt.Close()
	if err != nil {
		log.Print("UpdateRecordDiff Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Name, tz.ID)
	if err != nil {
		log.Print("UpdateRecordDiff Execute Statement  Error", err)
		return err
	}
	return nil
}

func (mdao DbConn) UpdateAssetRecordDiff(tz *entities.RecordDiffEntity) error {
	log.Println("In side dao")
	stmt, err := mdao.DB.Prepare(updateassetrecorddiff)
	defer stmt.Close()
	if err != nil {
		log.Print("updateassetrecorddiff Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Name, tz.Recorddifftypeid, tz.Parentid, tz.ID)
	log.Print(tz)
	if err != nil {
		log.Print("updateassetrecorddiff Execute Statement  Error", err)
		return err
	}
	return nil
}

func (mdao DbConn) DeleteRecordDiff(tz *entities.RecordDiffEntity) error {
	log.Println("In side dao")
	stmt, err := mdao.DB.Prepare(deleterecorddiff)
	defer stmt.Close()
	if err != nil {
		log.Print("DeleteRecordDiff Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.ID)
	if err != nil {
		log.Print("DeleteRecordDiff Execute Statement  Error", err)
		return err
	}
	return nil
}
