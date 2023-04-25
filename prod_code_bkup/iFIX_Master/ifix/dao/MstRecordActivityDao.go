package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertMstRecordActivity = "INSERT INTO mstrecordactivitymst ( clientid, mstorgnhirarchyid, activitydesc, seqno) VALUES (?,?,?,?) "
var duplicateMstRecordActivity = "SELECT count(id) total FROM  mstrecordactivitymst WHERE clientid = ? AND mstorgnhirarchyid = ?  AND activitydesc=? AND activeflg =1 AND deleteflg = 0 "
var duplicateMstRecordActivitytable = "SELECT count(id) total FROM  mstrecordactivitymst WHERE activitydesc=? AND activeflg =1 AND deleteflg = 0 "

//var getMstRecordActivity = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid,a.activitydesc as activitydesc,a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname FROM mstrecordactivitymst a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id ORDER BY a.id DESC LIMIT ?,?"
var getMstRecordActivitycount = "SELECT count(a.id) as total FROM mstrecordactivitymst a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id ORDER BY a.id DESC LIMIT ?,?"
var updateMstRecordActivity = "UPDATE mstrecordactivitymst SET clientid=?,mstorgnhirarchyid = ?, activitydesc = ? WHERE id = ? "
var updateMstRecordActivitycopyseq = "UPDATE mstrecordactivitymst SET clientid=?,mstorgnhirarchyid = ?, activitydesc = ?,seqno=? WHERE id = ? "

var deleteMstRecordActivity = "UPDATE mstrecordactivitymst SET deleteflg ='1' WHERE id = ? "
var insertMstRecordActivitycopyseq = "INSERT INTO mstrecordactivitymst ( clientid, mstorgnhirarchyid, activitydesc, seqno) VALUES (?,?,?,(select a.seqno from mstrecordactivitymst a where a.activitydesc=? and a.activeflg=1 and a.deleteflg=0 limit 1))"
var getseq = "select max(seqno) as seq FROM mstrecordactivitymst where activeflg=1 and deleteflg=0"
var getseqcopy = "select distinct seqno as seq FROM mstrecordactivitymst where activitydesc=? and activeflg=1 and deleteflg=0"
var getOrgWiseActivitydesc = "select activitydesc from mstrecordactivitymst where clientid=? and mstorgnhirarchyid=? and activeflg=1 and deleteflg=0"

func (dbc DbConn) GetRows(tz *entities.MstRecordActivityEntity, str string) ([]entities.MstRecordActivityEntity, error) {
	logger.Log.Println("In side GetRowsDao", str)
	values := []entities.MstRecordActivityEntity{}
	var getrows = "SELECT a.activitydesc as activitydesc,a.seqno as sequence FROM mstrecordactivitymst a WHERE a.activitydesc in (" + str + ") and a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 "

	rows, err := dbc.DB.Query(getrows, tz.Clientid, tz.Mstorgnhirarchyid)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetRowsDao Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstRecordActivityEntity{}

		rows.Scan(&value.Activitydesc, &value.Sequence)
		//logger.Log.Println(values)
		values = append(values, value)
	}
	return values, nil
}
func (dbc DbConn) CheckDuplicateMstRecordActivityTable(tz *entities.MstRecordActivityEntity) (entities.MstRecordActivityEntities, error) {
	logger.Log.Println("In side CheckDuplicateMstRecordActivityTableDao ")
	value := entities.MstRecordActivityEntities{}
	err := dbc.DB.QueryRow(duplicateMstRecordActivitytable, tz.Activitydesc).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMstRecordActivityTableDao Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) CheckDuplicateMstRecordActivity(tz *entities.MstRecordActivityEntity) (*entities.MstRecordActivityEntities, error) {
	logger.Log.Println("In side CheckDuplicateMstRecordActivityDao ")
	value := entities.MstRecordActivityEntities{}
	err := dbc.DB.QueryRow(duplicateMstRecordActivity, tz.Clientid, tz.Mstorgnhirarchyid, tz.Activitydesc).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return &value, nil
	case nil:
		return &value, nil
	default:
		logger.Log.Println("CheckDuplicateMstRecordActivityDao Get Statement Prepare Error", err)
		return &value, err
	}
}

func (dbc TxConn) AddMstRecordActivityCopy(tz *entities.MstRecordActivityEntity) (int64, error) {
	logger.Log.Println("In side AddMstRecordActivityCopyDao")
	logger.Log.Println("Query -->", insertMstRecordActivity)
	stmt, err := dbc.TX.Prepare(insertMstRecordActivity)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("AddMstRecordActivityCopyDao Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Activitydesc, tz.Sequence)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Activitydesc, tz.Sequence)
	if err != nil {
		logger.Log.Println("AddMstRecordActivityCopyDao Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}
func (dbc DbConn) AddMstRecordActivitycopyseq(tz *entities.MstRecordActivityEntity) (int64, error) {
	logger.Log.Println("In side AddMstRecordActivitycopyseqDao")
	logger.Log.Println("Query -->", insertMstRecordActivitycopyseq)
	stmt, err := dbc.DB.Prepare(insertMstRecordActivitycopyseq)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("AddMstRecordActivitycopyseqDao Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Activitydesc)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Activitydesc, tz.Activitydesc)
	if err != nil {
		logger.Log.Println("AddMstRecordActivitycopyseqDao Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}
func (dbc DbConn) AddMstRecordActivity(tz *entities.MstRecordActivityEntity) (int64, error) {
	logger.Log.Println("In side AddMstRecordActivityDao")
	logger.Log.Println("Query -->", insertMstRecordActivity)
	stmt, err := dbc.DB.Prepare(insertMstRecordActivity)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("AddMstRecordActivityDao Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Activitydesc, tz.Sequence)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Activitydesc, tz.Sequence)
	if err != nil {
		logger.Log.Println("AddMstRecordActivityDao Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllMstRecordActivity(tz *entities.MstRecordActivityEntity, OrgnType int64) ([]entities.MstRecordActivityEntity, error) {
	logger.Log.Println("In side GetAllMstRecordActivityDao")
	values := []entities.MstRecordActivityEntity{}

	var getMstRecordActivity string
	var params []interface{}
	if OrgnType == 1 {
		getMstRecordActivity = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid,a.activitydesc as activitydesc,a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname FROM mstrecordactivitymst a,mstclient b,mstorgnhierarchy c WHERE    a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.deleteflg =0 and a.activeflg=1  ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		getMstRecordActivity = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid,a.activitydesc as activitydesc,a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname FROM mstrecordactivitymst a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ?   AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.deleteflg =0 and a.activeflg=1 ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		getMstRecordActivity = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid,a.activitydesc as activitydesc,a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname FROM mstrecordactivitymst a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?    AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.deleteflg =0 and a.activeflg=1 ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}
	rows, err := dbc.DB.Query(getMstRecordActivity, params...)

	//rows, err := dbc.DB.Query(getMstRecordActivity, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstRecordActivityDao Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstRecordActivityEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Activitydesc, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateMstRecordActivity(tz *entities.MstRecordActivityEntity) error {
	logger.Log.Println("In side UpdateMstRecordActivityDao")
	stmt, err := dbc.DB.Prepare(updateMstRecordActivity)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMstRecordActivityDao Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Activitydesc, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMstRecordActivityDao Execute Statement  Error", err)
		return err
	}
	return nil
}
func (dbc DbConn) UpdateMstRecordActivitycopyseq(tz *entities.MstRecordActivityEntity) error {
	logger.Log.Println("In side UpdateMstRecordActivitycopyseqDao")
	stmt, err := dbc.DB.Prepare(updateMstRecordActivitycopyseq)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMstRecordActivitycopyseqDao Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Activitydesc, tz.Sequence, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMstRecordActivitycopyseqDao Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMstRecordActivity(tz *entities.MstRecordActivityEntity) error {
	logger.Log.Println("In side DeleteMstRecordActivityDao", tz)
	stmt, err := dbc.DB.Prepare(deleteMstRecordActivity)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMstRecordActivityDao Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMstRecordActivityDao Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMstRecordActivityCount(tz *entities.MstRecordActivityEntity, OrgnTypeID int64) (entities.MstRecordActivityEntities, error) {
	logger.Log.Println("In side GetMstRecordActivityCountdao")
	value := entities.MstRecordActivityEntities{}

	var getClientsupportgroupnewcount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getClientsupportgroupnewcount = "SELECT count(a.id) as total FROM mstrecordactivitymst a,mstclient b,mstorgnhierarchy c WHERE   a.clientid =b.id AND a.mstorgnhirarchyid = c.id  and a.deleteflg =0 and a.activeflg=1 "
	} else if OrgnTypeID == 2 {
		getClientsupportgroupnewcount = "SELECT count(a.id) as total FROM mstrecordactivitymst a,mstclient b,mstorgnhierarchy c WHERE  a.clientid = ? AND  a.clientid =b.id AND a.mstorgnhirarchyid = c.id and    a.deleteflg =0 and a.activeflg=1"
		params = append(params, tz.Clientid)
	} else {
		getClientsupportgroupnewcount = "SELECT count(a.id) as total FROM mstrecordactivitymst a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?   AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.deleteflg =0 and a.activeflg=1 "
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getClientsupportgroupnewcount, params...).Scan(&value.Total)

	//err := dbc.DB.QueryRow(getMstRecordActivitycount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMstRecordActivityCountdao Get Statement Prepare Error", err)
		return value, err
	}
}

// func (dbc DbConn) GetSeq(tz *entities.MstRecordActivityEntity) (interface{}, error) {
// 	logger.Log.Println("In side GetSeq")
// 	var seq interface{}
// 	// value := entities.CountryEntities{}
// 	err := dbc.DB.QueryRow("select max(seqno) as seq FROM mstrecordactivitymst where activeflg=1 and deleteflg=0").Scan(&seq)
// 	switch err {
// 	case sql.ErrNoRows:
// 		return seq, nil
// 	case nil:
// 		logger.Log.Println(seq)
// 		return seq, nil
// 	default:
// 		logger.Log.Println("Getseq Get Statement Prepare Error", err)
// 		return seq, err
// 	}
// }
func (mdao DbConn) GetSeq(tz *entities.MstRecordActivityEntity) ([]entities.MstRecordActivityEntity, error) {
	logger.Log.Println("In side GetSeqdao")
	values := []entities.MstRecordActivityEntity{}
	rows, err := mdao.DB.Query(getseq)

	if err != nil {
		logger.Log.Print("GetSeqdao Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		value := entities.MstRecordActivityEntity{}
		rows.Scan(&value.Sequence)
		values = append(values, value)
	}
	return values, nil
}
func (mdao DbConn) GetSeqcopy(tz *entities.MstRecordActivityEntity) ([]entities.MstRecordActivityEntity, error) {
	logger.Log.Println("In side GetSeqcopydao")
	values := []entities.MstRecordActivityEntity{}
	rows, err := mdao.DB.Query(getseqcopy, tz.Activitydesc)

	if err != nil {
		logger.Log.Print("GetSeqcopydao Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		value := entities.MstRecordActivityEntity{}
		rows.Scan(&value.Sequence)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) GetOrgWiseActivitydesc(tz *entities.MstRecordActivityEntity) ([]entities.Activitydesces, error) {
	logger.Log.Println("In side GetOrgWiseActivitydescDao")
	values := []entities.Activitydesces{}

	rows, err := dbc.DB.Query(getOrgWiseActivitydesc, tz.Clientid, tz.Mstorgnhirarchyid)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetOrgWiseActivitydescDao Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.Activitydesces{}

		rows.Scan(&value.Activitydesc)
		//logger.Log.Println(values)
		values = append(values, value)
	}
	return values, nil
}
