package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
)
var getprocesstemplatedetails = "SELECT id, details,detailsjson,iscomplete from mstprocesstemplatedetails where clientid=? and mstorgnhirarchyid=? and processid=? and activeflg=1 and deleteflg=0"
var duplicatetemplatetransition = "SELECT id  FROM  mstprocesstemplatetransition WHERE clientid = ? AND mstorgnhirarchyid = ? AND processid=? AND currentstateid = ? AND previousstateid=? AND activeflg=1 AND deleteflg = 0"
var inserttemplatetransitionquery = "INSERT INTO mstprocesstemplatetransition (clientid, mstorgnhirarchyid, processid,currentstateid,previousstateid) VALUES (?,?,?,?,?)"
var alltemplatetstatedetails = "SELECT a.mstgroupid,a.mstuserid FROM mapprocesstemplategroup a where a.transitionid=? and a.activeflg=1 and a.deleteflg=0"
var inserttemplatetransitiongroupquery = "INSERT INTO mapprocesstemplategroup (clientid, mstorgnhirarchyid, processid,mstgroupid,mstuserid,transitionid) VALUES (?,?,?,?,?,?)"
var insertprocesstemplate = "INSERT into mstprocesstemplatedetails(clientid,mstorgnhirarchyid,processid,details,detailsjson,iscomplete) values(?,?,?,?,?,?)"
var updateprocesstemplatedetails = "UPDATE mstprocesstemplatedetails set details=?,detailsjson=?,iscomplete=? where id=?"
var templatestatedetails = "SELECT a.mstgroupid,a.mstuserid FROM mapprocesstemplategroup a where a.transitionid=? and a.activeflg=1 and a.deleteflg=0"
var templatetransitiongroup = "SELECT COALESCE(a.id,0) as Mstgroupid,COALESCE(a.name,'') as groupname,COALESCE(b.id,0) as Mstuserid,COALESCE(b.loginname,'') loginname from mapprocesstemplategroup c left join mstclientuser b on c.mstuserid=b.id and b.activeflag=1 and b.deleteflag=0 left join mstsupportgrp a on a.id=c.mstgroupid  and a.activeflg=1 and a.deleteflg=0  where  c.clientid=? and c.mstorgnhirarchyid=? and c.transitionid=?  and c.activeflg=1 and c.deleteflg=0"
var getprocesstemplate ="SELECT a.id,a.processname from mstprocesstemplatedetails b,mstprocesstemplate a where b.clientid=? and b.mstorgnhirarchyid=? and b.iscomplete=1 and b.processid=a.id and b.activeflg=1 and b.deleteflg=0 and a.activeflg=1 and a.deleteflg=0"

func (mdao DbConn) Getprocesstemplatedetails(tz *entities.Workflowentity) ([]entities.WorkflowResponseEntity, error) {
	values := []entities.WorkflowResponseEntity{}
	rows, err := mdao.DB.Query(getprocesstemplatedetails, tz.Clientid, tz.Mstorgnhirarchyid, tz.Processid)

	if err != nil {
		log.Print("Getprocesstemplatedetails Get Statement Prepare Error", err)
		logger.Log.Print("Getprocesstemplatedetails Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		value := entities.WorkflowResponseEntity{}
		rows.Scan(&value.Id, &value.Details, &value.Detailsjson,&value.Iscomplete)
		values = append(values, value)
	}
	return values, nil
}
func (mdao DbConn) Checkduplicatetemplatestate(tz *entities.Workflowentity) ([]entities.Workflowentity, error) {
	log.Println("In side dao")
	values := []entities.Workflowentity{}
	rows, err := mdao.DB.Query(duplicatetemplatetransition, tz.Clientid, tz.Mstorgnhirarchyid, tz.Processid, tz.Currentstateid, tz.Previousstateid)
	if err != nil {
		log.Print("duplicatetemplatetransition Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		value := entities.Workflowentity{}
		rows.Scan(&value.Id)
		values = append(values, value)
	}
	return values, nil
}
func Createtemplatetransition(tz *entities.Workflowentity, tx *sql.Tx) (int64, error) {
	stmt, err := tx.Prepare(inserttemplatetransitionquery)
	if err != nil {
		logger.Log.Print("Createtemplatetransition Prepare Statement Prepare Error", err)
		log.Print("Createtemplatetransition Prepare Statement Prepare Error", err)
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Processid, tz.Currentstateid, tz.Previousstateid)
	if err != nil {
		logger.Log.Print("Createtemplatetransition Save Statement Execution Error", err)
		log.Print("Createtemplatetransition Save Statement Execution Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}
func (mdao DbConn) Getalltemplatetransitionstatedetails(tz *entities.Workflowentity) ([]entities.Workflowentity, error) {
	log.Println("In side dao")
	values := []entities.Workflowentity{}
	rows, err := mdao.DB.Query(alltemplatetstatedetails, tz.Transitionid)

	if err != nil {
		log.Print("Getalltemplatetransitionstatedetails Get Statement Prepare Error", err)
		logger.Log.Print("Getalltemplatetransitionstatedetails Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		value := entities.Workflowentity{}
		err := rows.Scan(&value.Mstgroupid, &value.Mstuserid)
		if err != nil {
			log.Print("Getalltemplatetransitionstatedetails Sacn Error", err)
			logger.Log.Print("Getalltemplatetransitionstatedetails Sacn Error", err)
			return values, err
		}
		values = append(values, value)
	}
	return values, nil
}
func Inserttemplatetransitiondetails(tz *entities.Workflowentity, tx *sql.Tx) (int64, error) {
	grpstmt, grperr := tx.Prepare(inserttemplatetransitiongroupquery)

	if grperr != nil {
		log.Print("Inserttemplatetransitiondetails group Prepare Statement Prepare Error", grperr)
		return 0, grperr
	}
	defer grpstmt.Close()
	_, err := grpstmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Processid, tz.Mstgroupid, tz.Mstuserid, tz.Transitionid)
	if err != nil {
		log.Print("Inserttemplatetransitiondetails group Save Statement Execution Error", err)
		return 0, err
	}
	return 1, nil
}
func (mdao DbConn) Insertprocesstemplate(tz *entities.Workflowentity) (int64, error) {
	log.Println("In side dao")
	stmt, err := mdao.DB.Prepare(insertprocesstemplate)
	if err != nil {
		logger.Log.Print("Insertprocesstemplate Prepare Statement  Error", err)
		log.Print("Insertprocesstemplate Prepare Statement  Error", err)
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Processid, tz.Details, tz.Detailsjson, tz.Iscomplete)
	if err != nil {
		logger.Log.Print("Insertprocesstemplate Execute Statement  Error", err)
		log.Print("Insertprocesstemplate Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	//defer mdao.DB.Close()
	return lastInsertedId, nil
}
func (mdao DbConn) Updateprocesstemplatedetails(tz *entities.Workflowentity) error {
	log.Println("In side dao")
	stmt, err := mdao.DB.Prepare(updateprocesstemplatedetails)
	if err != nil {
		logger.Log.Print("Updateprocesstemplatedetails Prepare Statement  Error", err)
		log.Print("Updateprocesstemplatedetails Prepare Statement  Error", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(tz.Details, tz.Detailsjson, tz.Iscomplete, tz.Id)
	if err != nil {
		logger.Log.Print("Updateprocesstemplatedetails Execute Statement  Error", err)
		log.Print("Updateprocesstemplatedetails Execute Statement  Error", err)
		return err
	}
	return nil
}
func Deletetemplatetransitiondetails(tz *entities.Workflowentity, tx *sql.Tx, ids string) error {
	logger.Log.Print("Deletetransitiondetails ", tz, ids)
	log.Print("Deletetransitiondetails", tz, ids)
	var deletetransitiondetails = "DELETE FROM mapprocesstemplategroup  WHERE clientid=? and mstorgnhirarchyid=? and transitionid in (" + ids + ")"
	stmt, err := tx.Prepare(deletetransitiondetails)
	if err != nil {
		logger.Log.Print("Deletetemplatetransitiondetails Prepare Statement  Error", err)
		log.Print("Deletetemplatetransitiondetails Prepare Statement  Error", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid)
	if err != nil {
		logger.Log.Print("Deletetemplatetransitiondetails Execute Statement  Error", err)
		log.Print("Deletetemplatetransitiondetails Execute Statement  Error", err)
		return err
	}
	return nil
}
func Deletetemplatetransition(tz *entities.Workflowentity, tx *sql.Tx, ids string) error {
	logger.Log.Print("Deletetemplatetransition ", tz, ids)
	log.Print("Deletetemplatetransition", tz, ids)
	var deletetemplatetransition = "DELETE FROM mstprocesstemplatetransition  WHERE clientid=? and mstorgnhirarchyid=? and id in (" + ids + ")"
	stmt, err := tx.Prepare(deletetemplatetransition)
	if err != nil {
		logger.Log.Print("Deletetemplatetransition Prepare Statement  Error", err)
		log.Print("Deletetemplatetransition Prepare Statement  Error", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid)
	if err != nil {
		logger.Log.Print("Deletetemplatetransition Execute Statement  Error", err)
		log.Print("Deletetemplatetransition Execute Statement  Error", err)
		return err
	}
	return nil
}
func (mdao DbConn) Gettemplatetransitionstatedetails(tz *entities.Workflowentity) ([]entities.WorkflowResponseEntity, error) {
	log.Println("In side dao")
	values := []entities.WorkflowResponseEntity{}
	rows, err := mdao.DB.Query(templatestatedetails, tz.Transitionid)

	if err != nil {
		logger.Log.Print("Gettemplatetransitionstatedetails Get Statement Prepare Error", err)
		log.Print("Gettemplatetransitionstatedetails Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		value := entities.WorkflowResponseEntity{}
		err := rows.Scan(&value.Mstgroupid, &value.Mstuserid)
		if err != nil {
			logger.Log.Print("Gettemplatetransitionstatedetails Sacn Error", err)
			log.Print("Gettemplatetransitionstatedetails Sacn Error", err)
			return values, err
		}
		values = append(values, value)
	}
	return values, nil
}
func (mdao DbConn) Gettemplatetransitiongroup(tz *entities.Workflowentity) ([]entities.WorkflowResponseEntity, error) {
	log.Println("In side dao")
	log.Println(tz.Clientid, tz.Mstorgnhirarchyid, tz.Transitionid)
	values := []entities.WorkflowResponseEntity{}
	rows, err := mdao.DB.Query(templatetransitiongroup, tz.Clientid, tz.Mstorgnhirarchyid, tz.Transitionid)

	if err != nil {
		log.Print("Gettemplatetransitiongroup Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		value := entities.WorkflowResponseEntity{}
		err := rows.Scan(&value.Mstgroupid, &value.Groupname, &value.Mstuserid, &value.Loginname)
		if err != nil {
			log.Print("Gettemplatetransitiongroup Sacn Error", err)
			return values, err
		}
		values = append(values, value)
	}
	return values, nil
}
func (mdao DbConn) Getprocesstemplate(tz *entities.Workflowentity) ([]entities.WorkflowResponseEntity, error) {
	log.Println("In side dao")
	log.Println(tz.Clientid, tz.Mstorgnhirarchyid, tz.Transitionid)
	values := []entities.WorkflowResponseEntity{}
	rows, err := mdao.DB.Query(getprocesstemplate, tz.Clientid, tz.Mstorgnhirarchyid)
	if err != nil {
		log.Print("Getprocesstemplate Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		value := entities.WorkflowResponseEntity{}
		err := rows.Scan(&value.Id, &value.Processname)
		if err != nil {
			log.Print("Getprocesstemplate Sacn Error", err)
			return values, err
		}
		values = append(values, value)
	}
	return values, nil
}