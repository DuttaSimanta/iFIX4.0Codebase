package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertMstClientCredential = "INSERT INTO mstclientcredential (clientid,mstorgnhirarchyid,credentialtypeid,credentialaccount,credentialpassword,credentialkey,credentialendpoint,defaultconfig) VALUES (?,?,?,?,?,?,?,?)"
var duplicateMstClientCredential = "SELECT count(id) total FROM  mstclientcredential WHERE clientid = ? AND mstorgnhirarchyid = ? AND credentialtypeid = ? AND credentialaccount = ? AND credentialpassword = ? AND credentialkey=? AND credentialendpoint=? AND defaultconfig=? AND deleteflg = 0"

//var getMstClientCredential = "SELECT a.id as Id, a.clientid as clientid, a.mstorgnhirarchyid as mstorgnhirarchyid, a.credentialtypeid as Credentialtypeid,a.credentialaccount as credentialaccount ,a.credentialpassword as credentialpassword, a.credentialkey as credentialkey,coalesce(a.credentialendpoint,'') as credentialendpoint , a.defaultconfig as defaultconfig, a.activeflg as activeflg,b.name as clientname,c.name as mstorgnhirarchyname,d.typename as credentialtypename ,IF(defaultconfig=1,'Default Configuration','New Configuration') defaultconfigname FROM mstclientcredential a,mstclient b,mstorgnhierarchy c,mstclientcredentialtype d WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.credentialtypeid=d.id ORDER BY a.id DESC LIMIT ?,?"
//var getMstClientCredentialcount = "SELECT count(a.id) as total FROM mstclientcredential a,mstclient b,mstorgnhierarchy c,mstclientcredentialtype d WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.credentialtypeid=d.id"
var updateMstClientCredential = "UPDATE mstclientcredential SET clientid = ?, mstorgnhirarchyid = ?, credentialtypeid = ?, credentialaccount = ?,credentialpassword=?,credentialkey=?,credentialendpoint=?,defaultconfig=? WHERE id = ? "
var deleteMstClientCredential = "UPDATE mstclientcredential SET deleteflg = '1' WHERE id = ? "
var getdata = "SELECT a.credentialtypeid as credentialtypeid, a.credentialaccount as credentialaccount,a.credentialpassword as credentialpassword,a.credentialkey as credentialkey,coalesce(a.credentialendpoint,'') as credentialendpoint from mstclientcredential a where a.clientid=? AND a.mstorgnhirarchyid=? AND a.credentialtypeid=? AND activeflg=1 AND deleteflg=0"

func (dbc DbConn) CheckDuplicateMstClientCredential(tz *entities.MstClientCredentialEntity) (entities.MstClientCredentialEntities, error) {
	logger.Log.Println("In side CheckDuplicateMstClientCredential")
	value := entities.MstClientCredentialEntities{}
	err := dbc.DB.QueryRow(duplicateMstClientCredential, tz.Clientid, tz.Mstorgnhirarchyid, tz.Credentialtypeid, tz.CredentialAccount, tz.CredentialPassword, tz.CredentialKey, tz.CredentialEndPoint, tz.DefaultConfig).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMstClientCredential Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertMstClientCredential(tz *entities.MstClientCredentialEntity) (int64, error) {
	logger.Log.Println("In side InsertMstClientCredential")
	logger.Log.Println("Query -->", insertMstClientCredential)
	stmt, err := dbc.DB.Prepare(insertMstClientCredential)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMstClientCredential Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Credentialtypeid, tz.CredentialAccount, tz.CredentialPassword, tz.CredentialKey, tz.CredentialEndPoint, tz.DefaultConfig)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Credentialtypeid, tz.CredentialAccount, tz.CredentialPassword, tz.CredentialKey, tz.CredentialEndPoint, tz.DefaultConfig)
	if err != nil {
		logger.Log.Println("InsertMstClientCredential Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllMstClientCredential(tz *entities.MstClientCredentialEntity, OrgnType int64) ([]entities.MstClientCredentialEntity, error) {
	logger.Log.Println("In side GelAllMstClientCredential")
	values := []entities.MstClientCredentialEntity{}

	var getMstClientCredential string
	var params []interface{}
	if OrgnType == 1 {
		getMstClientCredential = "SELECT a.id as Id, a.clientid as clientid, a.mstorgnhirarchyid as mstorgnhirarchyid, a.credentialtypeid as Credentialtypeid,a.credentialaccount as credentialaccount ,a.credentialpassword as credentialpassword, a.credentialkey as credentialkey,coalesce(a.credentialendpoint,'') as credentialendpoint , a.defaultconfig as defaultconfig, a.activeflg as activeflg,b.name as clientname,c.name as mstorgnhirarchyname,d.typename as credentialtypename ,IF(defaultconfig=1,'Default Configuration','New Configuration') defaultconfigname FROM mstclientcredential a,mstclient b,mstorgnhierarchy c,mstclientcredentialtype d WHERE  a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.credentialtypeid=d.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		getMstClientCredential = "SELECT a.id as Id, a.clientid as clientid, a.mstorgnhirarchyid as mstorgnhirarchyid, a.credentialtypeid as Credentialtypeid,a.credentialaccount as credentialaccount ,a.credentialpassword as credentialpassword, a.credentialkey as credentialkey,coalesce(a.credentialendpoint,'') as credentialendpoint , a.defaultconfig as defaultconfig, a.activeflg as activeflg,b.name as clientname,c.name as mstorgnhirarchyname,d.typename as credentialtypename ,IF(defaultconfig=1,'Default Configuration','New Configuration') defaultconfigname FROM mstclientcredential a,mstclient b,mstorgnhierarchy c,mstclientcredentialtype d WHERE a.clientid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.credentialtypeid=d.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		getMstClientCredential = "SELECT a.id as Id, a.clientid as clientid, a.mstorgnhirarchyid as mstorgnhirarchyid, a.credentialtypeid as Credentialtypeid,a.credentialaccount as credentialaccount ,a.credentialpassword as credentialpassword, a.credentialkey as credentialkey,coalesce(a.credentialendpoint,'') as credentialendpoint , a.defaultconfig as defaultconfig, a.activeflg as activeflg,b.name as clientname,c.name as mstorgnhirarchyname,d.typename as credentialtypename ,IF(defaultconfig=1,'Default Configuration','New Configuration') defaultconfigname FROM mstclientcredential a,mstclient b,mstorgnhierarchy c,mstclientcredentialtype d WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.credentialtypeid=d.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}
	rows, err := dbc.DB.Query(getMstClientCredential, params...)

	//rows, err := dbc.DB.Query(getMstClientCredential, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstClientCredential Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstClientCredentialEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Credentialtypeid, &value.CredentialAccount, &value.CredentialPassword, &value.CredentialKey, &value.CredentialEndPoint, &value.DefaultConfig, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Credentialtypename, &value.DefaultConfigName)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateMstClientCredential(tz *entities.MstClientCredentialEntity) error {
	logger.Log.Println("In side UpdateMstClientCredential")
	stmt, err := dbc.DB.Prepare(updateMstClientCredential)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMstClientCredential Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Credentialtypeid, tz.CredentialAccount, tz.CredentialPassword, tz.CredentialKey, tz.CredentialEndPoint, tz.DefaultConfig, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMstClientCredential Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMstClientCredential(tz *entities.MstClientCredentialEntity) error {
	logger.Log.Println("In side DeleteMstClientCredential")
	stmt, err := dbc.DB.Prepare(deleteMstClientCredential)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMstClientCredential Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMstClientCredential Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMstClientCredentialCount(tz *entities.MstClientCredentialEntity, OrgnType int64) (entities.MstClientCredentialEntities, error) {
	logger.Log.Println("In side GetMstClientCredentialCount")
	value := entities.MstClientCredentialEntities{}

	var getMstClientCredentialcount string
	var params []interface{}
	if OrgnType == 1 {
		getMstClientCredentialcount = "SELECT count(a.id) as total FROM mstclientcredential a,mstclient b,mstorgnhierarchy c,mstclientcredentialtype d WHERE a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.credentialtypeid=d.id"
	} else if OrgnType == 2 {
		getMstClientCredentialcount = "SELECT count(a.id) as total FROM mstclientcredential a,mstclient b,mstorgnhierarchy c,mstclientcredentialtype d WHERE a.clientid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.credentialtypeid=d.id"
		params = append(params, tz.Clientid)
	} else {
		getMstClientCredentialcount = "SELECT count(a.id) as total FROM mstclientcredential a,mstclient b,mstorgnhierarchy c,mstclientcredentialtype d WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.credentialtypeid=d.id"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getMstClientCredentialcount, params...).Scan(&value.Total)

	//err = dbc.DB.QueryRow(getMstClientCredentialcount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMstClientCredentialCount Get Statement Prepare Error", err)
		return value, err
	}
}
func (dbc DbConn) GetData(tz *entities.MstClientCredentialEntity) (entities.MstClientCredentialEntity, string, error) {
	/*logger.Log.Println("In side GelAllMstClientCredential")
	values := entities.MstClientCredentialEntity{}
	rows, err := dbc.DB.Query(getdata, page.Clientid, page.Mstorgnhirarchyid,page.Credentialtypeid)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstClientCredential Get Statement Prepare Error", err)
		return values, err
	}
	//for rows.Next() {
		err=rows.Next()
		switch err {
	    case sql.ErrNoRows:
		//value.Total = 0
		return value, nil
	}
		/*if err=sql.ErrNoRows{
			return value,"NO ROWS"
		}*
		value := entities.MstClientCredentialEntity{}
		rows.Scan(&value.Credentialtypeid,&value.CredentialAccount,&value.CredentialPassword,&value.CredentialKey,&value.CredentialEndPoint)
		//values = append(values, value)
	//}
	return value, nil*/
	logger.Log.Println("In side GetMstClientCredential")
	value := entities.MstClientCredentialEntity{}
	err := dbc.DB.QueryRow(getdata, 1, 1, tz.Credentialtypeid).Scan(&value.Credentialtypeid, &value.CredentialAccount, &value.CredentialPassword, &value.CredentialKey, &value.CredentialEndPoint)
	switch err {
	case sql.ErrNoRows:
		//value.Total = 0
		return value, "NOROWS", err
	case nil:
		return value, "", nil
	default:
		logger.Log.Println("GetMstClientCredential Get Statement Prepare Error", err)
		return value, "", err
	}
}
