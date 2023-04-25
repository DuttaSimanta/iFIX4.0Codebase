package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertUidGen = "INSERT INTO uidgen (clientid, mstorgnhirarchyid, difftypeid, code, uid) VALUES (?,?,(Select id from mstrecorddifferentiationtype where seqno=?),?,?)"
var duplicateUidGen = "SELECT count(id) total FROM  uidgen WHERE clientid = ? AND mstorgnhirarchyid = ? AND difftypeid IN (Select id from mstrecorddifferentiationtype where seqno= ?)  AND deleteflg = 0"

//var getUidGen = "SELECT a.id as Id, coalesce(a.clientid,0) as Clientid, coalesce(a.mstorgnhirarchyid,0) as Mstorgnhirarchyid, coalesce(a.difftypeid,0) as difftypeid, coalesce(a.code,'') as code, a.uid as uid,a.activeflg as activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Difftypename FROM uidgen a left join mstclient b on a.clientid=b.id left join mstorgnhierarchy c on a.mstorgnhirarchyid=c.id ,  mstrecorddifferentiationtype d WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.difftypeid=d.id and d.activeflg=1 and d.deleteflg=0  AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id  ORDER BY a.id DESC LIMIT ?,?"
//var getUidGencount = "SELECT count(a.id) as total FROM uidgen a left join mstclient b on a.clientid=b.id left join mstorgnhierarchy c on a.mstorgnhirarchyid=c.id ,  mstrecorddifferentiationtype d WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.difftypeid=d.id and d.activeflg=1 and d.deleteflg=0  AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id "
var updateUidGen = "UPDATE uidgen SET mstorgnhirarchyid = ?, difftypeid = ?, code = ?, uid = ? WHERE id = ? "
var updateUidGenforemail = "UPDATE uidgen SET uid = ? WHERE difftypeid = ? "

var deleteUidGen = "UPDATE uidgen SET deleteflg = '1' WHERE id = ? "

func (dbc DbConn) CheckDuplicateUidGen(tz *entities.UidGenEntity) (entities.UidGenEntities, error) {
	logger.Log.Println("In side CheckDuplicateUidGen")
	value := entities.UidGenEntities{}
	err := dbc.DB.QueryRow(duplicateUidGen, tz.Clientid, tz.Mstorgnhirarchyid, tz.Difftypeseq).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateUidGen Get Statement Prepare Error", err)
		return value, err
	}
}

func (mdao DbConn) GetAllOrganizationwithOrgtype(page *entities.MstorgnhierarchywithOrgtypeEntity, OrgnTypeID int64) ([]entities.MstorgnhierarchywithOrgtypeEntityResp, error) {
	values := []entities.MstorgnhierarchywithOrgtypeEntityResp{}
	var organizationselectclientwise string
	var params []interface{}
	if OrgnTypeID == 1 {
		organizationselectclientwise = "SELECT id as ID,name as Organizationname, mstorgnhierarchytypeid as MstorgnhierarchytypeID  FROM mstorgnhierarchy ORDER BY name"
	} else if OrgnTypeID == 2 {
		organizationselectclientwise = "SELECT id as ID,name as Organizationname, mstorgnhierarchytypeid as MstorgnhierarchytypeID  FROM mstorgnhierarchy  WHERE clientid =?  ORDER BY name"
		params = append(params, page.ClientID)
	} else {
		organizationselectclientwise = "SELECT id as ID,name as Organizationname,mstorgnhierarchytypeid as MstorgnhierarchytypeID  FROM mstorgnhierarchy WHERE clientid =? AND id=?  ORDER BY  name"
		params = append(params, page.ClientID)
		params = append(params, page.Mstorgnhirarchyid)
	}
	rows, err := mdao.DB.Query(organizationselectclientwise, params...)
	defer rows.Close()
	if err != nil {
		logger.Log.Print("GetAllClients Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstorgnhierarchywithOrgtypeEntityResp{}
		rows.Scan(&value.ID, &value.Organizationname, &value.MstorgnhierarchytypeID)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) InsertUidGen(tz *entities.UidGenEntity) (int64, error) {
	logger.Log.Println("In side InsertUidGen")
	logger.Log.Println("Query -->", insertUidGen)
	stmt, err := dbc.DB.Prepare(insertUidGen)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertUidGen Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Difftypeseq, tz.Code, tz.Uid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Difftypeseq, tz.Code, tz.Uid)
	if err != nil {
		logger.Log.Println("InsertUidGen Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllUidGen(tz *entities.UidGenEntity, OrgnType int64) ([]entities.UidGenEntity, error) {
	logger.Log.Println("In side GelAllUidGen")
	values := []entities.UidGenEntity{}

	var getUidGen string
	var params []interface{}
	if OrgnType == 1 {
		getUidGen = "SELECT a.id as Id, coalesce(a.clientid,0) as Clientid, coalesce(a.mstorgnhirarchyid,0) as Mstorgnhirarchyid, coalesce(a.difftypeid,0) as difftypeid, coalesce(a.code,'') as code, a.uid as uid,a.activeflg as activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Difftypename FROM uidgen a left join mstclient b on a.clientid=b.id left join mstorgnhierarchy c on a.mstorgnhirarchyid=c.id ,  mstrecorddifferentiationtype d WHERE a.difftypeid=d.id and d.activeflg=1 and d.deleteflg=0  AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.activeflg=1 AND a.deleteflg=0 ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		getUidGen = "SELECT a.id as Id, coalesce(a.clientid,0) as Clientid, coalesce(a.mstorgnhirarchyid,0) as Mstorgnhirarchyid, coalesce(a.difftypeid,0) as difftypeid, coalesce(a.code,'') as code, a.uid as uid,a.activeflg as activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Difftypename FROM uidgen a left join mstclient b on a.clientid=b.id left join mstorgnhierarchy c on a.mstorgnhirarchyid=c.id ,  mstrecorddifferentiationtype d WHERE a.clientid = ? AND a.difftypeid=d.id and d.activeflg=1 and d.deleteflg=0  AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.activeflg=1 AND a.deleteflg=0  ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		getUidGen = "SELECT a.id as Id, coalesce(a.clientid,0) as Clientid, coalesce(a.mstorgnhirarchyid,0) as Mstorgnhirarchyid, coalesce(a.difftypeid,0) as difftypeid, coalesce(a.code,'') as code, a.uid as uid,a.activeflg as activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Difftypename FROM uidgen a left join mstclient b on a.clientid=b.id left join mstorgnhierarchy c on a.mstorgnhirarchyid=c.id ,  mstrecorddifferentiationtype d WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.difftypeid=d.id and d.activeflg=1 and d.deleteflg=0  AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.activeflg=1 AND a.deleteflg=0  ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}
	rows, err := dbc.DB.Query(getUidGen, params...)

	//rows, err := dbc.DB.Query(getUidGen, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllUidGen Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.UidGenEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Difftypeseq, &value.Code, &value.Uid, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Difftypename)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateUidGen(tz *entities.UidGenEntity) error {
	logger.Log.Println("In side UpdateUidGen")
	stmt, err := dbc.DB.Prepare(updateUidGen)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateUidGen Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Difftypeseq, tz.Code, tz.Uid, tz.Id)
	logger.Log.Println("Data", tz.Mstorgnhirarchyid, tz.Difftypeseq, tz.Code, tz.Uid, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateUidGen Execute Statement  Error", err)
		return err
	}
	return nil
}
func (dbc DbConn) UpdateUidGenforemail(tz *entities.UidGenEntity) error {
	logger.Log.Println("In side UpdateUidGen")
	stmt, err := dbc.DB.Prepare(updateUidGenforemail)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateUidGen Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Uid, tz.Difftypeseq)
	if err != nil {
		logger.Log.Println("UpdateUidGen Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteUidGen(tz *entities.UidGenEntity) error {
	logger.Log.Println("In side DeleteUidGen")
	stmt, err := dbc.DB.Prepare(deleteUidGen)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteUidGen Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteUidGen Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetUidGenCount(tz *entities.UidGenEntity, OrgnTypeID int64) (entities.UidGenEntities, error) {
	logger.Log.Println("In side GetUidGenCount")
	value := entities.UidGenEntities{}
	var getUidGencount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getUidGencount = "SELECT count(a.id) as total FROM uidgen a left join mstclient b on a.clientid=b.id left join mstorgnhierarchy c on a.mstorgnhirarchyid=c.id ,  mstrecorddifferentiationtype d WHERE a.difftypeid=d.id and d.activeflg=1 and d.deleteflg=0  AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.activeflg=1 AND a.deleteflg=0"
	} else if OrgnTypeID == 2 {
		getUidGencount = "SELECT count(a.id) as total FROM uidgen a left join mstclient b on a.clientid=b.id left join mstorgnhierarchy c on a.mstorgnhirarchyid=c.id ,  mstrecorddifferentiationtype d WHERE a.clientid = ? AND a.difftypeid=d.id and d.activeflg=1 and d.deleteflg=0  AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.activeflg=1 AND a.deleteflg=0"
		params = append(params, tz.Clientid)
	} else {
		getUidGencount = "SELECT count(a.id) as total FROM uidgen a left join mstclient b on a.clientid=b.id left join mstorgnhierarchy c on a.mstorgnhirarchyid=c.id ,  mstrecorddifferentiationtype d WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.difftypeid=d.id and d.activeflg=1 and d.deleteflg=0  AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.activeflg=1 AND a.deleteflg=0"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getUidGencount, params...).Scan(&value.Total)

	//err := dbc.DB.QueryRow(getUidGencount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetUidGenCount Get Statement Prepare Error", err)
		return value, err
	}
}
