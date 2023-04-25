package dao

import (
	"database/sql"
	"errors"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
)

var moduleclientinsert = "INSERT INTO mstmoduleclient(clientid,mstorgnhirarchyid,moduleid) VALUES (?,?,?)"
var moduleclientdelete = "UPDATE mstmoduleclient set deleteflg=1 WHERE id=?"
var moduleclientduplicate = "SELECT count(id) total from mstmoduleclient WHERE clientid=? AND mstorgnhirarchyid=? AND moduleid=? AND deleteflg=0 AND activeflg=1"
var getmoduleclient = "SELECT a.id as ID,date_format(a.fromdate,'%d-%M-%Y')  as Fromdate,b.name as Clientname, d.name as Mstorgnhirarchyname, c.moduleName as Modulename,a.moduleid as Moduleid FROM mstmoduleclient a, mstclient b, mstmodule c,mstorgnhierarchy d WHERE  a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.moduleid = c.id AND a.deleteflg = 0 ORDER BY id DESC LIMIT ?,?"
var moduleclientcount = "SELECT count(a.id) total FROM mstmoduleclient a, mstclient b, mstmodule c,mstorgnhierarchy d WHERE  a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.moduleid = c.id AND a.deleteflg = 0"
var moduleclientupdate = "UPDATE mstmoduleclient SET clientid=?,mstorgnhirarchyid=?,moduleid=? WHERE id=?"
var modulebyorgid="SELECT a.modulename as Modulename,a.id as ID from mstmodule a,mstmoduleclient b where a.id=b.moduleid and b.clientid=? and b.mstorgnhirarchyid=? and b.activeflg=1 and b.deleteflg=0 order by a.modulename"
var lastseqbymodule ="SELECT max(sequence_no)  from dtlclientmenuinfo where clientid=? and mstorgnhirarchyid=? and parentmenuid=0  and activeflg=1 and deleteflg=0"

//CheckDuplicateModuleCient check duplicate record
func (mdao DbConn) CheckDuplicateModuleCient(tz *entities.MstModuleClientEntity) (entities.MstModuleClientEntities, error) {
	logger.Log.Println("moduleclientduplicate Query -->", moduleclientduplicate)
	logger.Log.Println("parameters -->", tz.ClientID, tz.Mstorgnhirarchyid, tz.ModuleID)
	value := entities.MstModuleClientEntities{}
	err := mdao.DB.QueryRow(moduleclientduplicate, tz.ClientID, tz.Mstorgnhirarchyid, tz.ModuleID).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		log.Println(err)
		logger.Log.Print("moduleclientduplicate Get Statement Prepare Error", err)
		return value, err
	}
}

//InsertModuleClientData data insertd in mstclientuser table
func (mdao DbConn) InsertModuleClientData(tz *entities.MstModuleClientEntity) (int64, error) {
	logger.Log.Println("moduleclientinsert query -->", moduleclientinsert)
	logger.Log.Println("parameters -->", tz.ClientID, tz.Mstorgnhirarchyid, tz.ModuleID)
	stmt, err := mdao.DB.Prepare(moduleclientinsert)
	defer stmt.Close()
	if err != nil {
		logger.Log.Print("InsertModuleClientData Prepare Statement  Error", err)
		return 0, err
	}
	res, err := stmt.Exec(tz.ClientID, tz.Mstorgnhirarchyid, tz.ModuleID)
	if err != nil {
		logger.Log.Print("InsertModuleClientData Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedID, err := res.LastInsertId()
	if err != nil {
		logger.Log.Println(err)
		return 0, errors.New("SQL Last ID fetch Error")
	}
	return lastInsertedID, nil
}

//UpdateClientModuleData update mstclientuser table
func (mdao DbConn) UpdateClientModuleData(data *entities.MstModuleClientEntity) error {
	logger.Log.Println("moduleclientupdate Query -->", moduleclientupdate)
	logger.Log.Println("parameters -->", data.ClientID, data.Mstorgnhirarchyid, data.ModuleID, data.Deleteflag, data.Activeflag, data.ID)
	stmt, err := mdao.DB.Prepare(moduleclientupdate)
	defer stmt.Close()
	if err != nil {
		logger.Log.Print("Update Client Module Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(data.ClientID, data.Mstorgnhirarchyid, data.ModuleID,  data.ID)
	if err != nil {
		logger.Log.Print("Update Client Module Execute Statement  Error", err)
		return err
	}
	return nil
}

//DeleteClientModuleData update mstclientuser table
func (mdao DbConn) DeleteClientModuleData(tz *entities.MstModuleClientEntity) error {
	logger.Log.Println("moduleclientdelete Query -->", moduleclientdelete)
	logger.Log.Println("parameters -->", tz.ID)

	stmt, err := mdao.DB.Prepare(moduleclientdelete)
	defer stmt.Close()
	if err != nil {
		logger.Log.Print("Delete Module Client Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.ID)
	if err != nil {
		logger.Log.Print("Delete Module Client Execute Statement  Error", err)
		return err
	}
	return nil
}

//GetClientModuleCount get user count with condition
func (mdao DbConn) GetClientModuleCount(tz *entities.MstModuleClientEntity) (entities.MstModuleClientEntities, error) {
	logger.Log.Println("moduleclientcount query -->", moduleclientcount)
	logger.Log.Println("parameters -->", tz.ClientID, tz.Mstorgnhirarchyid, tz.ModuleID)
	value := entities.MstModuleClientEntities{}
	err := mdao.DB.QueryRow(moduleclientcount).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Print("Get Module Client Count Get Statement Prepare Error", err)
		return value, err
	}
}

//GetAllModuleClients get user count with condition
func (mdao DbConn) GetAllModuleClients(page *entities.MstModuleClientEntity) ([]entities.MstModuleClientEntityResp, error) {
	logger.Log.Println("getmoduleclient Query -->", getmoduleclient)
	logger.Log.Println("parameters -->", page.ClientID, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	values := []entities.MstModuleClientEntityResp{}
	rows, err := mdao.DB.Query(getmoduleclient,  page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllClients Get Statement Prepare Error", err)
		log.Print("GetAllClients Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstModuleClientEntityResp{}
		rows.Scan(&value.ID, &value.Fromdate, &value.Clientname, &value.Mstorgnhirarchyname, &value.Modulename,&value.ModuleID)
		logger.Log.Println("value -->", value)
		values = append(values, value)
	}
	logger.Log.Println("values -->", values)
	return values, nil
}

//GetModuleByOrgId get modules by clientid and orgid
func (mdao DbConn) GetModuleByOrgId(page *entities.MstModuleClientEntity) ([]entities.MstModuleByClientEntity, error) {
	logger.Log.Println("GetModuleByOrgId Query -->", modulebyorgid)
	logger.Log.Println("GetModuleByOrgId -->", page.ClientID, page.Mstorgnhirarchyid)
	values := []entities.MstModuleByClientEntity{}
	rows, err := mdao.DB.Query(modulebyorgid, page.ClientID, page.Mstorgnhirarchyid)
	defer rows.Close()
	if err != nil {
		log.Print("GetModuleByOrgId Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstModuleByClientEntity{}
		rows.Scan( &value.Modulename,&value.ID)
		logger.Log.Println("value -->", value)
		values = append(values, value)
	}
	logger.Log.Println("values -->", values)
	return values, nil
}
// GetLastSeqFromMenu retrieve max seq when parentid is 0
func (mdao DbConn) GetLastSeqFromMenu(tz *entities.MstModuleClientEntity) ([]entities.MenuEntity, error) {
	log.Println("In side dao")
	values := []entities.MenuEntity{}
	rows, err := mdao.DB.Query(lastseqbymodule,tz.ClientID,tz.Mstorgnhirarchyid)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetLastSeqFromMenu Get Statement Prepare Error", err)
		log.Print("GetLastSeqFromMenu Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MenuEntity{}
		rows.Scan(&value.Sequence_no)
		values = append(values, value)
	}
	return values, nil
}
