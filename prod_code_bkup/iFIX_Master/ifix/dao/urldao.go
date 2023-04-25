package dao

import (
	"database/sql"
	"iFIX/ifix/logger"
	"log"
	"iFIX/ifix/entities"
)

var inserturl = "INSERT into msturl(urlkey,url,urldescription) values(?,?,?)"
var insertmodurl="INSERT into mstclientmoduleurl(clientid,mstorgnhirarchyid,moduleid,urlid) values(?,?,?,?)"
var duplicateUrl="SELECT id as Id from msturl where urlkey=? and url=? and deleteflg =0"
var duplicatemodurl ="SELECT count(id) total from mstclientmoduleurl where clientid=? and mstorgnhirarchyid=? and moduleid=? and urlid=? and deleteflg =0"
var getUrl="SELECT id as Id,urlkey as Urlkey , url as Url,urldescription as Urldescription FROM msturl WHERE deleteflg =0 ORDER BY id DESC LIMIT ?,? "
var getmoduleurl = "SELECT cu.id as Id ,u.url as Url,u.urlkey as Urlkey,m.modulename as Modulename,c.name as Clientname,o.name Orgname from mstclientmoduleurl cu ,msturl u,mstmodule m,mstclient c,mstorgnhierarchy o where cu.clientid=c.id and cu.mstorgnhirarchyid=o.id and cu.moduleid=m.id and cu.urlid=u.id and cu.deleteflg=0 and m.deleteflg=0 and u.deleteflg=0 ORDER BY id DESC LIMIT ?,?"
var updateurl="UPDATE msturl SET urlkey=?,url=?,urldescription=? where id=?"
var deleteurl= "UPDATE msturl SET deleteflg=1 where id=?"
var deletemoduleurl= "UPDATE mstclientmoduleurl SET deleteflg=1 where id=?"
var updatemoduleurl="UPDATE mstclientmoduleurl SET urlid=?  where clientid=? AND mstorgnhirarchyid=? AND moduleid=? and  urlid=? and deleteflg = 0 "
var geturlcount = "SELECT count(id) total from msturl where deleteflg =0"
var getmodurlcount = "SELECT count(cu.id) total from mstclientmoduleurl cu ,msturl u,mstmodule m,mstclient c,mstorgnhierarchy o where cu.clientid=c.id and cu.mstorgnhirarchyid=o.id and cu.moduleid=m.id and cu.urlid=u.id and cu.deleteflg=0 and m.deleteflg=0 and u.deleteflg=0"
var distincturl="SELECT u.id as Id, u.urlkey as `Urlkey`,u.url as `Url`,u.urldescription as `Urldescription` FROM mstclientmoduleurl m,msturl u where m.urlid=u.id and m.clientid=? and m.mstorgnhirarchyid=? and m.moduleid=? and  m.deleteflg=0 and u.deleteflg=0;"
var remainingurl="SELECT b.id AS Id,b.url Url,b.urlkey Urlkey,b.urldescription Urldescription FROM mstclientmoduleurl a,msturl b WHERE a.clientid = 1 AND a.mstorgnhirarchyid = 1 AND a.moduleid = ? AND a.deleteflg = '0' AND a.urlid = b.id AND b.deleteflg = '0' AND a.urlid NOT IN (SELECT urlid FROM dtlclientmenuinfo WHERE clientid = ? AND mstorgnhirarchyid = ? AND moduleid = ? AND deleteflg = '0' and urlid is not null) "
var urlmodulewise="SELECT urlid as Id from mstclientmoduleurl where clientid=? and mstorgnhirarchyid=? and moduleid=? and activeflg=1 and deleteflg=0"

func (mdao DbConn) CheckDuplicateUrl(tz *entities.UrlEntity) ([]entities.UrlEntity, error) {
	log.Println("In side dao")
	values := []entities.UrlEntity{}
	rows, err := mdao.DB.Query(duplicateUrl, tz.Urlkey,tz.Url)
	defer rows.Close()
	if err != nil {
		log.Print("CheckDuplicateUrl Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.UrlEntity{}
		rows.Scan(&value.Id)
		values = append(values, value)
	}
	return values, nil
}
func  CheckDuplicateModuleUrl(tz *entities.UrlEntity, tx *sql.Tx) (entities.ModuleUrlEntities, error) {
	log.Println("In side dao")
	value := entities.ModuleUrlEntities{}
	err := tx.QueryRow(duplicatemodurl, tz.Clientid,tz.Mstorgnhirarchyid,tz.Moduleid,tz.Id).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		log.Print("CheckDuplicateModuleUrl Get Statement Prepare Error", err)
		return value, err
	}
}

func (mdao DbConn) InsertUrl(tz *entities.UrlEntity) (int64, error) {
	log.Println("In side dao")
	tx, err := mdao.DB.Begin()
	if err != nil {
		log.Print("InsertUrl Begin Transaction create error")
		return 0, err
	}
	stmt, err := tx.Prepare(inserturl)
	defer stmt.Close()
	if err != nil {
		log.Print("InsertUrl Prepare Statement Prepare Error", err)
		return 0, err
	}
	res, err := stmt.Exec(tz.Urlkey, tz.Url, tz.Urldescription)
	if err != nil {
		log.Print("InsertUrl Save Statement Execution Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	tz.Id=lastInsertedId
	count,err:=CheckDuplicateModuleUrl(tz,tx)
	if err != nil {
		return 0, err
	}
	if count.Total ==0{
		if tz.Type=="update"{
			err:=updateModuleUrlFUnc(tz,tx)
			if err != nil {
				return 0, err
			}
		}else{
			err:=insertModuleUrl(tz,tx)
			if err != nil {
				return 0, err
			}
		}
	}else{
		return -1,nil
	}
	err = tx.Commit()
	if err != nil {
		log.Print("InsertUrl group Statement Commit error", err)
		tx.Rollback()
		return 0, err
	}
	return lastInsertedId, nil
}
func updateModuleUrlFUnc(tz *entities.UrlEntity,tx *sql.Tx) (error) {
	reqStmt, err := tx.Prepare(updatemoduleurl)
	defer reqStmt.Close()
	if err != nil {
		log.Print("updateModuleUrl Prepare Statement Prepare Error", err)
		return err
	}
	_, err = reqStmt.Exec(tz.Id, tz.Clientid, tz.Mstorgnhirarchyid, tz.Moduleid,tz.OldUrl)
	if err != nil {
		log.Print("updateModuleUrl Save Statement Execution Error", err)
		return err
	}
	return nil
}
func insertModuleUrl(tz *entities.UrlEntity,tx *sql.Tx) (error){
	stmt, err := tx.Prepare(insertmodurl)
	defer stmt.Close()
	if err != nil {
		log.Print("insertModuleUrl Prepare Statement Prepare Error", err)
		return err
	}
	_, err = stmt.Exec( tz.Clientid, tz.Mstorgnhirarchyid, tz.Moduleid,tz.Id)
	if err != nil {
		log.Print("insertModuleUrl Save Statement Execution Error", err)
		return err
	}
	return nil
}
func (mdao DbConn) InsertIntoModuleUrl(tz *entities.UrlEntity) (int64, error) {
	tx, err := mdao.DB.Begin()
	if err != nil {
		log.Print("InsertIntoModuleUrl Begin Transaction create error")
		return 0, err
	}
	count,err:=CheckDuplicateModuleUrl(tz,tx)
	if err != nil {
		return 0, err
	}
	if count.Total ==0{
		if tz.Type=="update"{
			err:=updateModuleUrlFUnc(tz,tx)
			if err != nil {
				return 0, err
			}
		}else{
			err:=insertModuleUrl(tz,tx)
			if err != nil {
				return 0, err
			}
		}
	}else{
		return -1,nil
	}
	err = tx.Commit()
	if err != nil {
		log.Print("InsertIntoModuleUrl group Statement Commit error", err)
		tx.Rollback()
		return 0, err
	}
	return tz.Id, nil
}

func (mdao DbConn) GetAllUrls(page *entities.PaginationEntity) ([]entities.UrlRespEntity, error) {
	log.Println("In side dao")
	values := []entities.UrlRespEntity{}
	rows, err := mdao.DB.Query(getUrl, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		log.Print("GetAllUrls Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.UrlRespEntity{}
		rows.Scan(&value.Id, &value.Urlkey, &value.Url,&value.Urldescription)
		values = append(values, value)
	}
	return values, nil
}

func (mdao DbConn) GetUrlsCount() (entities.UrlEntities, error) {
	log.Println("In side dao")
	value := entities.UrlEntities{}
	err := mdao.DB.QueryRow(geturlcount).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		log.Print("GetModuleCount Get Statement Prepare Error", err)
		return value, err
	}
}

func (mdao DbConn) GetAllModuleUrls(page *entities.PaginationEntity) ([]entities.ModuleUrlEntity, error) {
	log.Println("In side dao")
	values := []entities.ModuleUrlEntity{}
	rows, err := mdao.DB.Query(getmoduleurl, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		log.Print("GetAllUrls Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.ModuleUrlEntity{}
		rows.Scan(&value.Id, &value.Url, &value.Urlkey,&value.Modulename,&value.Clientname,&value.Orgname)
		values = append(values, value)
	}
	return values, nil
}

func (mdao DbConn) GetModuleUrlsCount() (entities.ModuleUrlEntities, error) {
	log.Println("In side dao")
	value := entities.ModuleUrlEntities{}
	err := mdao.DB.QueryRow(getmodurlcount).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		log.Print("GetModuleCount Get Statement Prepare Error", err)
		return value, err
	}
}

func (mdao DbConn) UpdateUrl(tz *entities.UrlEntity) error {
	log.Println("In side dao")
	stmt, err := mdao.DB.Prepare(updateurl)
	defer stmt.Close()
	if err != nil {
		log.Print("UpdateUrl Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Urlkey, tz.Url,tz.Urldescription, tz.Id)
	if err != nil {
		log.Print("UpdateUrl Execute Statement  Error", err)
		return err
	}
	return nil
}
func (mdao DbConn) DeleteUrl(tz *entities.UrlEntity) error {
	log.Println("In side dao")
	stmt, err := mdao.DB.Prepare(deleteurl)
	defer stmt.Close()
	if err != nil {
		log.Print("DeleteUrl Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		log.Print("DeleteUrl Execute Statement  Error", err)
		return err
	}
	return nil
}

func (mdao DbConn) DeleteModUrl(tz *entities.ModuleUrlEntity) error {
	log.Println("In side dao")
	stmt, err := mdao.DB.Prepare(deletemoduleurl)
	defer stmt.Close()
	if err != nil {
		log.Print("DeleteModUrl Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		log.Print("DeleteModUrl Execute Statement  Error", err)
		return err
	}
	return nil
}

func (mdao DbConn) GetDistinctUrl(tz *entities.UrlEntity) ([]entities.UrlRespEntity, error) {
	log.Println("In side dao")
	values := []entities.UrlRespEntity{}
	rows, err := mdao.DB.Query(distincturl, tz.Clientid, tz.Mstorgnhirarchyid,tz.Moduleid)
	defer rows.Close()
	if err != nil {
		log.Print("GetDistinctUrl Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.UrlRespEntity{}
		rows.Scan(&value.Id, &value.Urlkey, &value.Url,&value.Urldescription)
		values = append(values, value)
	}
	return values, nil
}
func (mdao DbConn) GetRemainingUrl(tz *entities.UrlEntity) ([]entities.UrlRespEntity, error) {
	log.Println("In side dao")
	values := []entities.UrlRespEntity{}
	rows, err := mdao.DB.Query(remainingurl, tz.Moduleid, tz.Clientid,tz.Mstorgnhirarchyid,tz.Moduleid)
	defer rows.Close()
	if err != nil {
		log.Print("GetRemainingUrl Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.UrlRespEntity{}
		rows.Scan(&value.Id,&value.Url, &value.Urlkey, &value.Urldescription)
		values = append(values, value)
	}
	return values, nil
}
/**
return urlid using clientid,moduleid,orgnid
 */
func (mdao DbConn) Geturlmodulewise(tz *entities.UrlEntity) ([]entities.UrlRespEntity, error) {
	log.Println("In side dao")
	values := []entities.UrlRespEntity{}
	rows, err := mdao.DB.Query(urlmodulewise, tz.Clientid, tz.Mstorgnhirarchyid,tz.Moduleid)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("Geturlmodulewise Get Statement Prepare Error", err)
		log.Print("Geturlmodulewise Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.UrlRespEntity{}
		rows.Scan(&value.Id)
		values = append(values, value)
	}
	return values, nil
}

