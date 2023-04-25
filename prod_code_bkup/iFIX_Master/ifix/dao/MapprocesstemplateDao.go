package dao

import (
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
)
var gettemplatename="SELECT processname from mstprocesstemplate where id=? and clientid=? and mstorgnhirarchyid=? and  activeflg=1 and deleteflg=0"
var gettemplateentity="SELECT  mstdatadictionaryfieldid from mapprocesstemplatetoentity where clientid=? and mstorgnhirarchyid=? and mstprocessid=? and  activeflg=1 and deleteflg=0"
var gettemplatetransaction="SELECT a.id,a.currentstateid,a.previousstateid,b.seqno currseq,coalesce(c.seqno,0) prevseq from mstprocesstemplatetransition a left join mststate c on a.previousstateid=c.id and c.activeflg=1 and c.deleteflg=0 ,mststate b where a.currentstateid=b.id and a.clientid=? and a.mstorgnhirarchyid=? and a.processid=? and a.activeflg=1 and a.deleteflg=0 and b.activeflg=1 and b.deleteflg=0"
var gettemplategroup ="SELECT mstgroupid,mstuserid from mapprocesstemplategroup where clientid=? and mstorgnhirarchyid=? and processid=? and transitionid=? and  activeflg=1 and deleteflg=0"
var getstateid="SELECT id from mststate where clientid=? and mstorgnhirarchyid=? and seqno=? and activeflg=1 and deleteflg=0"
var getmapprocesstemplatestate = "SELECT b.seqno from mapprocesstemplatestate a,mststate b where a.statetid=b.id and a.clientid=? and a.mstorgnhirarchyid=? and a.processid=? and  a.activeflg=1 and a.deleteflg=0 and b.activeflg=1 and b.deleteflg=0"
var getdiffname="SELECT parentcategorynames from mstrecorddifferentiation where id=?"

func (mdao DbConn) Gettemplatename(tz *entities.MapprocesstemplateEntity) ([]entities.WorkflowResponseEntity, error) {
	values := []entities.WorkflowResponseEntity{}
	stmt, err := mdao.DB.Prepare(gettemplatename)
	if err != nil {
		logger.Log.Print("Getprocessname Statement Prepare Error", err)
		log.Print("Getprocessname Statement Prepare Error", err)
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(tz.Processid,tz.Clientid, tz.Loggedinmstorgnhirarchyid )
	if err != nil {
		logger.Log.Print("Getprocessname Statement Execution Error", err)
		log.Print("Getprocessname Statement Execution Error", err)
		return nil, err
	}
	for rows.Next() {
		value := entities.WorkflowResponseEntity{}
		rows.Scan(&value.Processname)
		values = append(values, value)
	}
	return values, nil
}
func (mdao DbConn) Gettemplateentity(tz *entities.MapprocesstemplateEntity) ([]entities.MstprocessEntity, error) {
	values := []entities.MstprocessEntity{}
	stmt, err := mdao.DB.Prepare(gettemplateentity)
	if err != nil {
		logger.Log.Print("Gettemplateentity Statement Prepare Error", err)
		log.Print("Gettemplateentity Statement Prepare Error", err)
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(tz.Clientid, tz.Loggedinmstorgnhirarchyid, tz.Processid)
	if err != nil {
		logger.Log.Print("Gettemplateentity Statement Execution Error", err)
		log.Print("Gettemplateentity Statement Execution Error", err)
		return nil, err
	}
	for rows.Next() {
		value := entities.MstprocessEntity{}
		rows.Scan(&value.Mstdatadictionaryfieldid)
		values = append(values, value)
	}
	return values, nil
}
func (mdao DbConn) Gettemplatetransition(tz *entities.MapprocesstemplateEntity) ([]entities.WorkflowStateEntity, error) {
	values := []entities.WorkflowStateEntity{}
	stmt, err := mdao.DB.Prepare(gettemplatetransaction)
	if err != nil {
		logger.Log.Print("Gettemplatetransition Statement Prepare Error", err)
		log.Print("Gettemplatetransition Statement Prepare Error", err)
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(tz.Clientid, tz.Loggedinmstorgnhirarchyid, tz.Processid)
	if err != nil {
		logger.Log.Print("Gettemplatetransition Statement Execution Error", err)
		log.Print("Gettemplatetransition Statement Execution Error", err)
		return nil, err
	}
	for rows.Next() {
		value := entities.WorkflowStateEntity{}
		rows.Scan(&value.Templatetransitionid,&value.Currentstateid,&value.Previousstateid,&value.Currentseq,&value.Previousseq)
		values = append(values, value)
	}
	return values, nil
}
func (mdao DbConn) 	Gettemplategroupbytemplatetransition(tz *entities.MapprocesstemplateEntity) ([]entities.Workflowentity, error) {
	values := []entities.Workflowentity{}
	stmt, err := mdao.DB.Prepare(gettemplategroup)
	if err != nil {
		logger.Log.Print("Gettemplategroupbytemplatetransition Statement Prepare Error", err)
		log.Print("Gettemplategroupbytemplatetransition Statement Prepare Error", err)
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(tz.Clientid, tz.Loggedinmstorgnhirarchyid, tz.Processid,tz.Templatetransitionid)
	if err != nil {
		logger.Log.Print("Gettemplategroupbytemplatetransition Statement Execution Error", err)
		log.Print("Gettemplategroupbytemplatetransition Statement Execution Error", err)
		return nil, err
	}
	for rows.Next() {
		value := entities.Workflowentity{}
		rows.Scan(&value.Mstgroupid,&value.Mstuserid)
		values = append(values, value)
	}
	return values, nil
}
func (mdao DbConn) Getstateidbyseq(tz *entities.MapprocesstemplateEntity,seqno int64) ([]entities.Workflowentity, error) {
	values := []entities.Workflowentity{}
	stmt, err := mdao.DB.Prepare(getstateid)
	if err != nil {
		logger.Log.Print("Getstateidbyseq Statement Prepare Error", err)
		log.Print("Getstateidbyseq Statement Prepare Error", err)
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(tz.Clientid, tz.Mstorgnhirarchyid, seqno)
	if err != nil {
		logger.Log.Print("Getstateidbyseq Statement Execution Error", err)
		log.Print("Getstateidbyseq Statement Execution Error", err)
		return nil, err
	}
	for rows.Next() {
		value := entities.Workflowentity{}
		rows.Scan(&value.Id)
		values = append(values, value)
	}
	return values, nil
}
func (mdao DbConn) Gettemplatestate(tz *entities.MapprocesstemplateEntity) ([]entities.Workflowentity, error) {
	values := []entities.Workflowentity{}
	stmt, err := mdao.DB.Prepare(getmapprocesstemplatestate)
	if err != nil {
		logger.Log.Print("Gettemplatestate Statement Prepare Error", err)
		log.Print("Gettemplatestate Statement Prepare Error", err)
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(tz.Clientid, tz.Loggedinmstorgnhirarchyid, tz.Processid)
	if err != nil {
		logger.Log.Print("Gettemplatestate Statement Execution Error", err)
		log.Print("Gettemplatestate Statement Execution Error", err)
		return nil, err
	}
	for rows.Next() {
		value := entities.Workflowentity{}
		rows.Scan(&value.Id)
		values = append(values, value)
	}
	return values, nil
}
func (mdao DbConn) Getprocessstatebyseq(tz *entities.MapprocesstemplateEntity,ids string) ([]int64, error) {
	var values []int64
	var processsatebyseq="SELECT id from mststate where clientid=? and  mstorgnhirarchyid=? and seqno in ("+ids+") and activeflg=1 and deleteflg=0"
	stmt, err := mdao.DB.Prepare(processsatebyseq)
	if err != nil {
		logger.Log.Print("Getprocessstatebyseq Statement Prepare Error", err)
		log.Print("Getprocessstatebyseq Statement Prepare Error", err)
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(tz.Clientid, tz.Mstorgnhirarchyid)
	if err != nil {
		logger.Log.Print("Getprocessstatebyseq Statement Execution Error", err)
		log.Print("Getprocessstatebyseq Statement Execution Error", err)
		return nil, err
	}
	for rows.Next() {
		var value int64
		rows.Scan(&value)
		values = append(values, value)
	}
	return values, nil
}
func (mdao DbConn) Getdiffnamebyid(id int64) ([]string, error) {
	var values []string
	stmt, err := mdao.DB.Prepare(getdiffname)
	if err != nil {
		logger.Log.Print("Getdiffnamebyid Statement Prepare Error", err)
		log.Print("Getdiffnamebyid Statement Prepare Error", err)
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		logger.Log.Print("Getdiffnamebyid Statement Execution Error", err)
		log.Print("Getdiffnamebyid Statement Execution Error", err)
		return nil, err
	}
	for rows.Next() {
		var value string
		rows.Scan(&value)
		values = append(values, value)
	}
	return values, nil
}
