package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
)

func (mdao DbConn) FetchAssignedTicketByUser(tz *entities.MstClientUserEntity) ([]entities.Workflowentity, error) {
	var fetchassignedticket = "SELECT id from mstrequest where mstuserid=? and activeflg=1 and deleteflg=0"
	log.Println("Userid:->",tz.ID)
	values := []entities.Workflowentity{}
	rows, err := mdao.DB.Query(fetchassignedticket,tz.ID)

	if err != nil {
		log.Print("FetchAssignedTicketByUser Prepare Statement  Error", err)
		logger.Log.Print("FetchAssignedTicketByUser Prepare Statement  Error", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		value := entities.Workflowentity{}
		rows.Scan(&value.Mstrequestid)
		values = append(values, value)
	}
	return values,nil
}
func (mdao DbConn) Gethistorydetails(id int64) ([]entities.Workflowentity, error) {
	var fetchassignedticket = "SELECT  a.clientid ,a.mstorgnhirarchyid,a.processid,a.userid,a.currentstateid,a.transitionid,a.manualstateselection,a.mainrequestid,a.title,a.mstgroupid,a.mstuserid,b.supportgrouplevelid from mstrequesthistory a,mstclientsupportgroup b where a.mainrequestid=? and a.activeflg=1 and a.deleteflg=0 and a.mstgroupid =b.grpid order by a.id desc limit 2"

	values := []entities.Workflowentity{}
	rows, err := mdao.DB.Query(fetchassignedticket,id)

	if err != nil {
		log.Print("Gethistorydetails Prepare Statement  Error", err)
		logger.Log.Print("Gethistorydetails Prepare Statement  Error", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		value := entities.Workflowentity{}
		rows.Scan(&value.Clientid,&value.Mstorgnhirarchyid,&value.Processid,&value.Userid,&value.Currentstateid,&value.Transitionid,&value.Manualstateselection,&value.Mstrequestid,&value.Recordtitle,&value.Mstgroupid,&value.Mstuserid,&value.Levelid)
		values = append(values, value)
	}
	return values,nil
}
func InsertProcessHistory(tz *entities.Workflowentity, tx *sql.Tx, latestTime int64,isAttached string) error {
	histStmt, histErr := tx.Prepare(insertRequestHistory)

	if histErr != nil {
		log.Print("UpsertProcessDetails insertRequestHistory Prepare Statement Prepare Error", histErr)
		return histErr
	}
	defer histStmt.Close()
	_, histErr = histStmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Processid, tz.Mstrequestid, tz.Recordtitle, tz.Userid, latestTime, tz.Currentstateid, tz.Transitionid, latestTime, tz.Manualstateselection, tz.Mstgroupid, tz.Mstuserid,isAttached)
	if histErr != nil {
		log.Print("UpsertProcessDetails insertRequestHistory Save Statement Execution Error", histErr)
		return histErr
	}
	return nil
}
func UpdateAssignedUser(groupid int64,userid int64, tx *sql.Tx,  requestId int64) error {
	var updateRequest="UPDATE mstrequest set mstuserid=?,mstgroupid=? where id=? "
	reqStmt, err := tx.Prepare(updateRequest)
	if err != nil {
		log.Print("UpdateAssignedUser Prepare Statement Prepare Error", err)
		logger.Log.Print("UpdateAssignedUser Statement Prepare Error", err)
		return err
	}
	defer reqStmt.Close()
	logger.Log.Print(updateRequest)
	_, err = reqStmt.Exec( userid, groupid, requestId)
	if err != nil {
		log.Print("UpdateAssignedUser Statement Execution Error", err)
		logger.Log.Print("UpdateAssignedUser Save Statement Execution Error", err)
		return err
	}
	return nil
}
