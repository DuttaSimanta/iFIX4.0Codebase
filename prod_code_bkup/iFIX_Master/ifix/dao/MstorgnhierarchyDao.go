package dao

import (
	"database/sql"
	"errors"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
)

var organizationinsert = "INSERT INTO mstorgnhierarchy(clientid,parentid,mstorgnhierarchytypeid,name,cityid,countryid,code,location,timezoneid,reporttimezoneid ,logintypeid,timeformatid,activationdate,reporttimeformatid,islocallogin,orgmfa,notification,originalbgimage,uploadedbgimage,originallogoimage,uploadedlogoimage ) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,round(UNIX_TIMESTAMP(?)),?,?,?,?,?,?,?,?)"
var organinsert = "INSERT INTO mstorgnhierarchy(clientid,mstorgnhierarchytypeid,name,code,activationdate,parentid,timeformatid,logintypeid,reporttimeformatid,islocallogin) VALUES(?,?,?,?,round(UNIX_TIMESTAMP(now())),0,?,?,?,?)"
var organizationupdate = "UPDATE mstorgnhierarchy SET clientid=?,parentid=?,mstorgnhierarchytypeid=?,name=?,cityid=?,countryid=?,code=?,location=?,timezoneid=?,reporttimezoneid=?,activationdate=round(UNIX_TIMESTAMP(?)),timeformatid=?,logintypeid=?,reporttimeformatid=?,islocallogin=?,orgmfa=?, notification=?,originalbgimage=?,uploadedbgimage=?,originallogoimage=?,uploadedlogoimage=? WHERE id=?"
var organizationduplicate = "SELECT COUNT(id) total FROM mstorgnhierarchy WHERE clientid=? AND mstorgnhierarchytypeid=? AND name=? AND code=? AND timeformatid=? AND logintypeid=? And reporttimeformatid=?"
var organizationcountclientwise = "SELECT COUNT(id) total FROM mstorgnhierarchy WHERE clientid=?"

var organizationselectclientwise = "SELECT a.id as ID,a.name as Organizationname  FROM mstorgnhierarchy a WHERE a.clientid =?  ORDER BY id DESC"

//var organizationselect = "SELECT a.id as ID,b.name as Clientname,a.clientid as ClientID,a.parentid as ParentID,COALESCE((Select b.name from mstorgnhierarchy b where a.parentid= b.id),'Base Organization') Parentname,a.mstorgnhierarchytypeid as MstorgnhierarchytypeID, d.orgntype as Mstorgnhierarchytypename,a.name as Organizationname,a.Code,a.location as Location  FROM mstorgnhierarchy a, mstclient b, mstorgnhierarchytype d WHERE a.clientid = b.id AND a.mstorgnhierarchytypeid = d.id ORDER BY id DESC LIMIT ?,?"
var organizationselect = "SELECT a.id as ID,a.clientid as ClientID,a.parentid as ParentID,a.mstorgnhierarchytypeid as MstorgnhierarchytypeID,a.name as Organizationname, coalesce(a.cityid,0) CityID,coalesce(a.countryid,0) CountryID,a.code as Code,coalesce(a.location,'')Location,coalesce(a.timezoneid,0) TimezoneID,coalesce(a.reporttimezoneid,0) ReporttimezoneID, FROM_UNIXTIME(a.activationdate) as Activationdate,b.name as Clientname,COALESCE((Select b.name from mstorgnhierarchy b where a.parentid= b.id),'Base Organization') Parentname, c.orgntype as Mstorgnhierarchytypename,COALESCE(d.cityname,'') as Cityname ,COALESCE(e.countryname,'') as Countryname ,COALESCE(f.zone_name,'') as Timezonename,COALESCE(g.zone_name,'') as Reporttimezonename,coalesce(a.islocallogin,'1') islocallogin,coalesce(h.datetime,'') timeformat,a.logintypeid as LogintypeID,coalesce(i.name,'') as LogintypeName,a.reporttimeformatid as ReportTimeformatid,a.timeformatid as Timeformatid,coalesce(j.datetime,'') as ReportTimeformat,a.orgmfa as mfa,IF(a.orgmfa=1,'ENABLE','DISABLE') as mfaname, a.notification,IF(a.notification=1,'ENABLE','DISABLE') as notificationname,coalesce(a.originalbgimage,'') as originalbgimage,coalesce(a.uploadedbgimage,'') as uploadedbgimage,coalesce(a.originallogoimage,'') as originallogoimage,coalesce(a.uploadedlogoimage,'') as uploadedlogoimage  FROM mstclient b,mstorgnhierarchytype c,mstorgnhierarchy a left join mstcity d on a.cityid=d.id left join mstcountry e on a.countryid=e.id left join zone f on a.timezoneid=f.zone_id left join zone g on a.reporttimezoneid=g.zone_id left join mstdatetimeformat h on a.timeformatid=h.id left join mstlogintype i on a.logintypeid=i.id left join mstdatetimeformat j on a.reporttimeformatid=j.id WHERE a.clientid = b.id AND a.mstorgnhierarchytypeid = c.id ORDER BY a.id DESC LIMIT ?,?"

//var organizationcount = "SELECT COUNT(a.id) total FROM mstorgnhierarchy a, mstclient b, mstorgnhierarchytype d WHERE a.clientid = b.id AND a.mstorgnhierarchytypeid = d.id"
var organizationcount = "SELECT count(a.id) as total  FROM mstclient b,mstorgnhierarchytype c,mstorgnhierarchy a left join mstcity d on a.cityid=d.id left join mstcountry e on a.countryid=e.id left join zone f on a.timezoneid=f.zone_id left join zone g on a.reporttimezoneid=g.zone_id left join mstdatetimeformat h on a.timeformatid=h.id left join mstlogintype i on a.logintypeid=i.id left join mstdatetimeformat j on a.reporttimeformatid=j.id WHERE a.clientid = b.id AND a.mstorgnhierarchytypeid = c.id"
var baseid = "SELECT id as ID,clientid as ClientID from mstorgnhierarchy where mstorgnhierarchytypeid=1"
var gettimeformat = "SELECT id ,datetime timeformat from mstdatetimeformat"
var logintype = "SELECT a.id as ID,a.name as Name from mstlogintype a"

func (mdao DbConn) UpdateOrganizationField(data *entities.MstorgnhierarchyFieldEntity) error {
	logger.Log.Println("Update Query -->", organizationupdate)
	logger.Log.Println("parameters -->", data.ClientID, data.Organizationname, data.ID)
	var organizationupdate = "UPDATE mstorgnhierarchy SET orgmfa=?, notification=?,autoclose=?,autoclosefromdt=?,autoclosetodt=? WHERE id=?"
	stmt, err := mdao.DB.Prepare(organizationupdate)
	defer stmt.Close()
	if err != nil {
		logger.Log.Print("Update Organization Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(data.Mfa, data.Notification, data.Autoclose, data.Autoclosefromdt, data.Autoclosetodt, data.ID)
	if err != nil {
		logger.Log.Print("Update Organization Execute Statement  Error", err)
		return err
	}
	return nil
}

func (mdao DbConn) GetAllOrganizationField(page *entities.MstorgnhierarchyFieldEntity, OrgnType int64) ([]entities.MstorgnhierarchyFieldEntity, error) {
	// logger.Log.Println("organizationselect Query -->", organizationselect)
	logger.Log.Println("parameters -->", page.Offset, page.Limit)
	values := []entities.MstorgnhierarchyFieldEntity{}

	var organizationfieldselect string
	var params []interface{}
	if OrgnType == 1 {
		organizationfieldselect = "SELECT a.id as organizationid,a.clientid as clientid,a.name as organizationid,b.name as clientname,a.orgmfa as mfa,IF(a.orgmfa=1,'ENABLE','DISABLE') as mfaname, a.notification,IF(a.notification=1,'ENABLE','DISABLE') as notificationname,a.autoclose,coalesce(a.autoclosefromdt,'') as autoclosefromdt,coalesce(a.autoclosetodt,'') as autoclosetodt,IF(a.autoclose=1,'ENABLE','DISABLE') as autoclosestatus  FROM mstclient b,mstorgnhierarchy a  WHERE a.clientid = b.id and a.deleteflg=0 ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else if OrgnType == 2 {
		organizationfieldselect = "SELECT a.id as organizationid,a.clientid as clientid,a.name as organizationid,b.name as clientname,a.orgmfa as mfa,IF(a.orgmfa=1,'ENABLE','DISABLE') as mfaname, a.notification,IF(a.notification=1,'ENABLE','DISABLE') as notificationname,a.autoclose,coalesce(a.autoclosefromdt,'') as autoclosefromdt,coalesce(a.autoclosetodt,'') as autoclosetodt,IF(a.autoclose=1,'ENABLE','DISABLE') as autoclosestatus  FROM mstclient b,mstorgnhierarchy a  WHERE a.clientid = b.id and a.deleteflg=0 and a.clientid=? ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.ClientID)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else {
		organizationfieldselect = "SELECT a.id as organizationid,a.clientid as clientid,a.name as organizationid,b.name as clientname,a.orgmfa as mfa,IF(a.orgmfa=1,'ENABLE','DISABLE') as mfaname, a.notification,IF(a.notification=1,'ENABLE','DISABLE') as notificationname,a.autoclose,coalesce(a.autoclosefromdt,'') as autoclosefromdt,coalesce(a.autoclosetodt,'') as autoclosetodt,IF(a.autoclose=1,'ENABLE','DISABLE') as autoclosestatus  FROM mstclient b,mstorgnhierarchy a  WHERE a.clientid = b.id and a.deleteflg=0 and a.clientid=? and a.id=? ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.ClientID)
		params = append(params, page.ID)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	}
	rows, err := mdao.DB.Query(organizationfieldselect, params...)

	// var organizationfieldselect = "SELECT a.id as organizationid,a.clientid as clientid,a.name as organizationid,b.name as clientname,a.orgmfa as mfa,IF(a.orgmfa=1,'ENABLE','DISABLE') as mfaname, a.notification,IF(a.notification=1,'ENABLE','DISABLE') as notificationname,a.autoclose,coalesce(a.autoclosefromdt,'') as autoclosefromdt,coalesce(a.autoclosetodt,'') as autoclosetodt,IF(a.autoclose=1,'ENABLE','DISABLE') as autoclosestatus  FROM mstclient b,mstorgnhierarchy a  WHERE a.clientid = b.id and a.deleteflg=0 ORDER BY a.id DESC LIMIT ?,?"
	// rows, err := mdao.DB.Query(organizationfieldselect, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Print("GetAllClients Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstorgnhierarchyFieldEntity{}
		err = rows.Scan(&value.ID, &value.ClientID, &value.Organizationname, &value.Clientname, &value.Mfa, &value.MfaName, &value.Notification, &value.NotificationName, &value.Autoclose, &value.Autoclosefromdt, &value.Autoclosetodt, &value.AutocloseStatus)
		logger.Log.Println("value -->", err)
		values = append(values, value)
	}
	logger.Log.Println("values -->", values)

	return values, nil
}
func (mdao DbConn) GetOrganizationFieldCount(tz *entities.MstorgnhierarchyFieldEntity, OrgnTypeID int64) (entities.MstorgnhierarchyFieldEntities, error) {
	// logger.Log.Println("Count Query -->", organizationcount)
	//logger.Log.Println("parameters -->", tz.ClientID)
	value := entities.MstorgnhierarchyFieldEntities{}

	var organizationfieldselect string
	var params []interface{}
	if OrgnTypeID == 1 {
		organizationfieldselect = "SELECT count(a.id) as total  FROM mstclient b,mstorgnhierarchy a  WHERE a.clientid = b.id and a.deleteflg=0"
	} else if OrgnTypeID == 2 {
		organizationfieldselect = "SELECT count(a.id) as total  FROM mstclient b,mstorgnhierarchy a  WHERE a.clientid = b.id and a.clientid =? and a.deleteflg=0"
		params = append(params, tz.ClientID)
	} else {
		organizationfieldselect = "SELECT count(a.id) as total  FROM mstclient b,mstorgnhierarchy a  WHERE a.clientid = b.id and a.clientid =? and a.id=? and a.deleteflg=0"
		params = append(params, tz.ClientID)
		params = append(params, tz.ID)
	}
	err := mdao.DB.QueryRow(organizationfieldselect, params...).Scan(&value.Total)

	// var organizationfieldselect = "SELECT count(a.id) as total  FROM mstclient b,mstorgnhierarchy a  WHERE a.clientid = b.id and a.deleteflg=0"
	// err := mdao.DB.QueryRow(organizationfieldselect).Scan(&value.Total)

	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		log.Print("Organization Get Statement Prepare Error", err)
		return value, err
	}
}


//mstorgnhierarchy a,mstclient b,mstorgnhierarchytype c,mstcity d,mstcountry e,zone f,zone g,mstlogintype h WHERE a.clientid = b.id AND a.mstorgnhierarchytypeid = c.id and a.cityid=d.id and a.countryid=e.id and a.timezoneid=f.zone_id and a.reporttimezoneid=g.zone_id and a.logintypeid=h.id"
//CheckDuplicateOrganization check duplicate record
func (mdao DbConn) CheckDuplicateOrganization(tz *entities.MstorgnhierarchyEntity) (entities.MstorgnhierarchyEntities, error) {
	logger.Log.Println("Check Duplicate Query -->", organizationduplicate)
	logger.Log.Println("parameters -->", tz.ClientID, tz.MstorgnhierarchytypeID, tz.Organizationname)
	value := entities.MstorgnhierarchyEntities{}
	err := mdao.DB.QueryRow(organizationduplicate, tz.ClientID, tz.MstorgnhierarchytypeID, tz.Organizationname, tz.Code, tz.Timeformatid, tz.LogintypeID, tz.ReportTimeformatid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Print("CheckDuplicateOrganization Get Statement Prepare Error", err)
		return value, err
	}
}

//CheckDuplicateOrganization check duplicate record
func CheckDuplicateOrganizationwithTX(tx *sql.Tx, ClientID int64, MstorgnhierarchytypeID int64, Organizationname string, Code string, Timeformatid int64, Reporttimeformatid int64, Logintypeid int64) (entities.MstorgnhierarchyEntities, error) {
	logger.Log.Println("Update Query -->", organizationduplicate)
	logger.Log.Println("parameters -->", ClientID, MstorgnhierarchytypeID, Organizationname)
	value := entities.MstorgnhierarchyEntities{}
	err := tx.QueryRow(organizationduplicate, ClientID, MstorgnhierarchytypeID, Organizationname, Code, Timeformatid, Logintypeid, Reporttimeformatid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Print("CheckDuplicateOrganization Get Statement Prepare Error", err)
		return value, err
	}
}

//InsertOrganization data insertd in mstclientuserrole table
func (mdao DbConn) InsertOrganization(data *entities.MstorgnhierarchyEntity) (int64, error) {
	logger.Log.Println("Insert query -->", organizationinsert)
	logger.Log.Println("parameters -->", data.ClientID, data.ParentID, data.MstorgnhierarchytypeID, data.Organizationname, data.CityID, data.CountryID, data.Code, data.Location, data.TimezoneID, data.ReporttimezoneID, data.LogintypeID)
	stmt, err := mdao.DB.Prepare(organizationinsert)
	defer stmt.Close()
	if err != nil {
		logger.Log.Print("Organization Prepare Statement  Error", err)
		return 0, err
	}
	res, err := stmt.Exec(data.ClientID, data.ParentID, data.MstorgnhierarchytypeID, data.Organizationname, data.CityID, data.CountryID, data.Code, data.Location, data.TimezoneID, data.ReporttimezoneID, data.LogintypeID, data.Timeformatid, data.Activationdate, data.ReportTimeformatid, data.Islocallogin, data.Mfa, data.Notification, data.Originalbgimage, data.Uploadedbgimage, data.Originallogoimage, data.Uploadedlogoimage)
	if err != nil {
		logger.Log.Print("Organization Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedID, err := res.LastInsertId()
	if err != nil {
		logger.Log.Println(err)
		return 0, errors.New("SQL Last ID fetch Error")
	}
	return lastInsertedID, nil
}

//InsertOrganization data insertd in mstclientuserrole table
func InsertOrganizationwithTX(tx *sql.Tx, ClientID int64, MstorgnhierarchytypeID int64, Organizationname string, Code string, Timeformatid int64, Reporttimeformatid int64, Logintypeid int64, Islocallogin int64) (int64, error) {
	logger.Log.Println("Insert query -->", organinsert)
	logger.Log.Println("parameters -->", ClientID, MstorgnhierarchytypeID, Organizationname, Code)
	stmt, err := tx.Prepare(organinsert)
	defer stmt.Close()
	if err != nil {
		logger.Log.Print("Organization Prepare Statement  Error", err)
		return 0, err
	}
	res, err := stmt.Exec(ClientID, MstorgnhierarchytypeID, Organizationname, Code, Timeformatid, Logintypeid, Reporttimeformatid, Islocallogin)
	if err != nil {
		logger.Log.Print("Organization Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedID, err := res.LastInsertId()
	if err != nil {
		logger.Log.Println(err)
		return 0, errors.New("SQL Last ID fetch Error")
	}
	return lastInsertedID, nil
}

//UpdateOrganization update record number table
func (mdao DbConn) UpdateOrganization(data *entities.MstorgnhierarchyEntity) error {
	logger.Log.Println("Update Query -->", organizationupdate)
	logger.Log.Println("parameters -->", data.ClientID, data.ParentID, data.MstorgnhierarchytypeID, data.Organizationname, data.CityID, data.CountryID, data.Code, data.Location, data.TimezoneID, data.ReporttimezoneID, data.ID)
	stmt, err := mdao.DB.Prepare(organizationupdate)
	defer stmt.Close()
	if err != nil {
		logger.Log.Print("Update Organization Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(data.ClientID, data.ParentID, data.MstorgnhierarchytypeID, data.Organizationname, data.CityID, data.CountryID, data.Code, data.Location, data.TimezoneID, data.ReporttimezoneID, data.Activationdate, data.Timeformatid, data.LogintypeID, data.ReportTimeformatid, data.Islocallogin, data.Mfa, data.Notification, data.Originalbgimage, data.Uploadedbgimage, data.Originallogoimage, data.Uploadedlogoimage, data.ID)
	if err != nil {
		logger.Log.Print("Update Organization Execute Statement  Error", err)
		return err
	}
	return nil
}

//GetOrganizationCountClientWise get role count with condition
func (mdao DbConn) GetOrganizationCountClientWise(tz *entities.MstorgnhierarchyEntity) (entities.MstorgnhierarchyEntities, error) {
	logger.Log.Println("Count Query -->", organizationcountclientwise)
	logger.Log.Println("parameters -->", tz.ClientID)
	value := entities.MstorgnhierarchyEntities{}
	err := mdao.DB.QueryRow(organizationcountclientwise, tz.ClientID).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		log.Print("Organization Get Statement Prepare Error", err)
		return value, err
	}
}

//GetOrganizationCount get role count with condition
func (mdao DbConn) GetOrganizationCount() (entities.MstorgnhierarchyEntities, error) {
	logger.Log.Println("Count Query -->", organizationcount)
	//logger.Log.Println("parameters -->", tz.ClientID)
	value := entities.MstorgnhierarchyEntities{}
	err := mdao.DB.QueryRow(organizationcount).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		log.Print("Organization Get Statement Prepare Error", err)
		return value, err
	}
}

//GetAllOrganizationClientWise get user count with condition
func (mdao DbConn) GetAllOrganizationClientWise(page *entities.MstorgnhierarchyEntity) ([]entities.MstorgnhierarchyEntityResp, error) {
	logger.Log.Println("organizationselect Query -->", organizationselectclientwise)
	logger.Log.Println("parameters -->", page.ClientID)
	values := []entities.MstorgnhierarchyEntityResp{}
	rows, err := mdao.DB.Query(organizationselectclientwise, page.ClientID)
	defer rows.Close()
	if err != nil {
		logger.Log.Print("GetAllClients Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstorgnhierarchyEntityResp{}
		rows.Scan(&value.ID, &value.Organizationname)
		values = append(values, value)
	}
	return values, nil
}

//GetAllOrganizationClientWise get user count with condition
func (mdao DbConn) GetAllOrganizationClientWisenew(page *entities.MstorgnhierarchyEntity, OrgnTypeID int64) ([]entities.MstorgnhierarchyEntityResp, error) {
	values := []entities.MstorgnhierarchyEntityResp{}
	var organizationselectclientwise string
	var params []interface{}
	if OrgnTypeID == 1 {
		organizationselectclientwise = "SELECT id as ID,name as Organizationname  FROM mstorgnhierarchy ORDER BY name"
	} else if OrgnTypeID == 2 {
		organizationselectclientwise = "SELECT id as ID,name as Organizationname  FROM mstorgnhierarchy  WHERE clientid =?  ORDER BY name"
		params = append(params, page.ClientID)
	} else {
		organizationselectclientwise = "SELECT id as ID,name as Organizationname  FROM mstorgnhierarchy WHERE clientid =? AND id=?  ORDER BY  name"
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
		value := entities.MstorgnhierarchyEntityResp{}
		rows.Scan(&value.ID, &value.Organizationname)
		values = append(values, value)
	}
	return values, nil
}

//GetAllOrganizationClientWise get user count with condition
func (mdao DbConn) Gettimeformat() ([]entities.MstorgnhierarchyEntityResp, error) {
	values := []entities.MstorgnhierarchyEntityResp{}
	rows, err := mdao.DB.Query(gettimeformat)

	if err != nil {
		logger.Log.Print("gettimeformat Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		value := entities.MstorgnhierarchyEntityResp{}
		rows.Scan(&value.ID, &value.Timeformat)
		values = append(values, value)
	}
	return values, nil
}

//GetAllOrganization get user count with condition
func (mdao DbConn) GetAllOrganization(page *entities.MstorgnhierarchyEntity) ([]entities.MstorgnhierarchyEntity, error) {
	logger.Log.Println("organizationselect Query -->", organizationselect)
	logger.Log.Println("parameters -->", page.Offset, page.Limit)
	values := []entities.MstorgnhierarchyEntity{}
	rows, err := mdao.DB.Query(organizationselect, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Print("GetAllClients Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstorgnhierarchyEntity{}
		err = rows.Scan(&value.ID, &value.ClientID, &value.ParentID, &value.MstorgnhierarchytypeID, &value.Organizationname, &value.CityID, &value.CountryID, &value.Code, &value.Location, &value.TimezoneID, &value.ReporttimezoneID, &value.Activationdate, &value.Clientname, &value.Parentname, &value.Mstorgnhierarchytypename, &value.Cityname, &value.Countryname, &value.Timezonename, &value.Reporttimezonename, &value.Islocallogin, &value.Timeformat, &value.LogintypeID, &value.LogintypeName, &value.ReportTimeformatid, &value.Timeformatid, &value.ReportTimeformat, &value.Mfa, &value.MfaName, &value.Notification, &value.NotificationName, &value.Originalbgimage, &value.Uploadedbgimage, &value.Originallogoimage, &value.Uploadedlogoimage)
		logger.Log.Println("value -->", err)
		values = append(values, value)
	}
	logger.Log.Println("values -->", values)

	return values, nil
}

//GetBaseOrgDetails returns base client id and org id
func (mdao DbConn) GetBaseOrgDetails() ([]entities.MstorgnhierarchyEntity, error) {
	logger.Log.Println("organizationselect Query -->", baseid)
	values := []entities.MstorgnhierarchyEntity{}
	rows, err := mdao.DB.Query(baseid)
	defer rows.Close()
	if err != nil {
		logger.Log.Print("GetBaseOrgDetails Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstorgnhierarchyEntity{}
		err = rows.Scan(&value.ID, &value.ClientID)
		logger.Log.Println("value -->", err)
		values = append(values, value)
	}
	logger.Log.Println("values -->", values)
	return values, nil
}

func (mdao DbConn) GetLogintype() ([]entities.LogintypeEntity, error) {
	logger.Log.Println("organizationselect Query -->", logintype)
	logger.Log.Println("parameters -->")
	values := []entities.LogintypeEntity{}
	rows, err := mdao.DB.Query(logintype)
	defer rows.Close()
	if err != nil {
		logger.Log.Print("GetLogintype Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.LogintypeEntity{}
		err = rows.Scan(&value.ID, &value.Name)
		logger.Log.Println("value -->", err)
		values = append(values, value)
	}
	logger.Log.Println("values -->", values)
	return values, nil
}
