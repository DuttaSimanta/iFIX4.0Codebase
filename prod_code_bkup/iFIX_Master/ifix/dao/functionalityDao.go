package dao

import (
	"database/sql"
	"iFIX/ifix/logger"

	//"database/sql"
	"iFIX/ifix/entities"
	"log"
)

var funcmaster = "SELECT id ,name from mstfunctionality"
var insertmapping = "INSERT into mapfunctionality(clientid,mstorgnhirarchyid,funcid,funcdescid,description,seqno,colorcode,image,ismanegerialview,iscatalog) values(?,?,?,?,?,?,?,?,?,?)"
var lastfuncid = "SELECT max(funcdescid)  from mapfunctionality where clientid=? and mstorgnhirarchyid=? and funcid=?  and activeflg=1 and deleteflg=0"
var duplicatemapping = "SELECT count(id) total from mapfunctionality where clientid=? and mstorgnhirarchyid=? and funcid=?  and description=? AND funcdescid=? and iscatalog=?  and activeflg=1 and deleteflg=0"
var mappingbycatalogtype = "SELECT funcdescid,description,seqno,colorcode,image,ismanegerialview from mapfunctionality where clientid=? and mstorgnhirarchyid=? and funcid=? and iscatalog=? and activeflg=1 and deleteflg=0 order by description"
var mappingbytype = "SELECT funcdescid,description,seqno,colorcode,image,ismanegerialview,iscatalog from mapfunctionality where clientid=? and mstorgnhirarchyid=? and funcid=?  and activeflg=1 and deleteflg=0 order by description"
var mappingdetails = "SELECT a.id ,a.description,a.activeflg,a.image,a.ismanegerialview,a.iscatalog,a.funcdescid,b.name as Clientname, d.name as Orgname,c.name as funcname from mapfunctionality a,mstclient b, mstorgnhierarchy d,mstfunctionality c where a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.funcid=c.id and a.activeflg=1 and a.deleteflg=0 ORDER BY a.id DESC LIMIT ?,? "
var mappingdetailscount = "SELECT count(a.id) total from mapfunctionality a,mstclient b, mstorgnhierarchy d,mstfunctionality c where a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.funcid=c.id and a.activeflg=1 and a.deleteflg=0 "
var deletemapping = "UPDATE mapfunctionality set deleteflg=1 where id=?"

func (mdao DbConn) Getfunctionality() ([]entities.FuncmasterEntity, error) {
	log.Println("In side dao")
	values := []entities.FuncmasterEntity{}
	rows, err := mdao.DB.Query(funcmaster)
	defer rows.Close()
	if err != nil {
		logger.Log.Print("Getfunctionality Get Statement Prepare Error", err)
		log.Print("Getfunctionality Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.FuncmasterEntity{}
		rows.Scan(&value.ID, &value.Name)
		values = append(values, value)
	}
	return values, nil
}
func (mdao DbConn) Getfuncmappingbytype(tz *entities.FuncmappingEntity) ([]entities.FuncmappingsingleRespEntity, error) {
	log.Println("In side dao")
	values := []entities.FuncmappingsingleRespEntity{}
	rows, err := mdao.DB.Query(mappingbytype, tz.Clientid, tz.Mstorgnhirarchyid, tz.Funcid)
	defer rows.Close()
	if err != nil {
		logger.Log.Print("Getfuncmappingbytype Get Statement Prepare Error", err)
		log.Print("Getfuncmappingbytype Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.FuncmappingsingleRespEntity{}
		rows.Scan(&value.Funcdescid, &value.Description, &value.Seqno, &value.Colorcode, &value.Image, &value.Ismanegerialview, &value.Iscatalog)
		values = append(values, value)
	}
	return values, nil
}
func (mdao DbConn) Getfuncmappingbycatalogtype(tz *entities.FuncmappingEntity) ([]entities.FuncmappingsingleRespEntity, error) {
	log.Println("In side dao")
	values := []entities.FuncmappingsingleRespEntity{}
	rows, err := mdao.DB.Query(mappingbycatalogtype, tz.Clientid, tz.Mstorgnhirarchyid, tz.Funcid, tz.Iscatalog)
	defer rows.Close()
	if err != nil {
		logger.Log.Print("Getfuncmappingbytype Get Statement Prepare Error", err)
		log.Print("Getfuncmappingbytype Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.FuncmappingsingleRespEntity{}
		rows.Scan(&value.Funcdescid, &value.Description, &value.Seqno, &value.Colorcode, &value.Image, &value.Ismanegerialview)
		values = append(values, value)
	}
	return values, nil
}
func (mdao DbConn) Getfuncmappingbytypeforquery(tz *entities.FuncmappingEntity) ([]entities.FuncmappingsingleRespEntity, error) {
	log.Println("In side dao")
	values := []entities.FuncmappingsingleRespEntity{}
	var mappingbytypeforquery = "SELECT funcdescid,description,seqno,colorcode,image,ismanegerialview from mapfunctionality where clientid=? and mstorgnhirarchyid=? and funcid=? and ismanegerialview=? and activeflg=1 and deleteflg=0 order by description"
	rows, err := mdao.DB.Query(mappingbytypeforquery, tz.Clientid, tz.Mstorgnhirarchyid, tz.Funcid, tz.Ismanegerialview)
	defer rows.Close()
	if err != nil {
		logger.Log.Print("Getfuncmappingbytype Get Statement Prepare Error", err)
		log.Print("Getfuncmappingbytype Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.FuncmappingsingleRespEntity{}
		rows.Scan(&value.Funcdescid, &value.Description, &value.Seqno, &value.Colorcode, &value.Image, &value.Ismanegerialview)
		values = append(values, value)
	}
	return values, nil
}
func (mdao DbConn) Getlastfuncdescid(tz *entities.FuncmappingEntity) ([]entities.FuncmappingEntity, error) {
	log.Println("In side dao")
	values := []entities.FuncmappingEntity{}
	rows, err := mdao.DB.Query(lastfuncid, tz.Clientid, tz.Mstorgnhirarchyid, tz.Funcid)
	defer rows.Close()
	if err != nil {
		logger.Log.Print("Getlastfuncdescid Get Statement Prepare Error", err)
		log.Print("Getlastfuncdescid Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.FuncmappingEntity{}
		rows.Scan(&value.Funcdescid)
		values = append(values, value)
	}
	return values, nil
}
func (mdao DbConn) Checkduplicatefuncmapping(tz *entities.FuncmappingEntity) (entities.FuncmappingEntitities, error) {
	log.Println("In side dao")
	value := entities.FuncmappingEntitities{}
	err := mdao.DB.QueryRow(duplicatemapping, tz.Clientid, tz.Mstorgnhirarchyid, tz.Funcid, tz.Description, tz.Funcdescid, tz.Iscatalog).Scan(&value.Total)
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
func (mdao DbConn) Insertfuncmapping(tz *entities.FuncmappingEntity) (int64, error) {
	log.Println("In side dao")
	log.Println(tz.Clientid, tz.Mstorgnhirarchyid, tz.Funcid, tz.Funcdescid, tz.Description, tz.Seqno, tz.Colorcode, tz.Image, tz.Ismanegerialview, tz.Iscatalog)
	stmt, err := mdao.DB.Prepare(insertmapping)
	defer stmt.Close()
	if err != nil {
		logger.Log.Print("Insertfuncmapping Prepare Statement  Error", err)
		log.Print("Insertfuncmapping Prepare Statement  Error", err)
		return 0, err
	}
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Funcid, tz.Funcdescid, tz.Description, tz.Seqno, tz.Colorcode, tz.Image, tz.Ismanegerialview, tz.Iscatalog)
	if err != nil {
		logger.Log.Print("Insertfuncmapping Execute Statement  Error", err)
		log.Print("Insertfuncmapping Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}
func (mdao DbConn) Getfuncmappingdetails(tz *entities.FuncmappingEntity) ([]entities.FuncmappingRespEntity, error) {
	log.Println("In side dao")
	values := []entities.FuncmappingRespEntity{}
	rows, err := mdao.DB.Query(mappingdetails, tz.Offset, tz.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Print("Getfuncmappingdetails Get Statement Prepare Error", err)
		log.Print("Getfuncmappingdetails Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.FuncmappingRespEntity{}
		err = rows.Scan(&value.ID, &value.Description, &value.Activeflg, &value.Image, &value.Ismanegerialview, &value.Iscatalog,&value.Seqno, &value.Clientname, &value.Orgname, &value.Funcname)
		if err != nil {
			logger.Log.Print("Getfuncmappingdetails scan Error", err)
			log.Print("Getfuncmappingdetails scan Error", err)
		}
		values = append(values, value)
	}
	return values, nil
}
func (mdao DbConn) Getmappingdetailscount() (entities.FuncmappingEntitities, error) {
	log.Println("In side dao")
	value := entities.FuncmappingEntitities{}
	err := mdao.DB.QueryRow(mappingdetailscount).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Print("Getmappingdetailscount Get Statement Prepare Error", err)
		log.Print("Getmappingdetailscount Get Statement Prepare Error", err)
		return value, err
	}
}
func (mdao DbConn) Deletefunctionmapping(tz *entities.FuncmappingEntity) error {
	log.Println("In side dao")
	stmt, err := mdao.DB.Prepare(deletemapping)
	defer stmt.Close()
	if err != nil {
		log.Print("Deletefunctionmappingf Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.ID)
	if err != nil {
		log.Print("Deletefunctionmappingf Execute Statement  Error", err)
		return err
	}
	return nil
}
