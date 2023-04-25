package dao

import "src/logger"

func (mdao DbConn) GetRecordInStartStatus() ([]int64, []int64, []int64, error) {
	logger.Log.Println("In side GetClientID")
	var records []int64
	var orgnIDs []int64
	var clientIDs []int64
	var sql = "select a.recordid,a.clientid,a.mstorgnhirarchyid from recordfulldetails a,trnrecord e where a.recordid not in (select b.recordid from maprequestorecord b,mstrequest c  where a.recordid=b.recordid and b.mstrequestid=c.id) and a.recordid=e.id and a.tickettype='Incident' and e.deleteflg=0 and a.deleteflg=0 order by a.id desc limit 20;"

	rows, err := mdao.DB.Query(sql)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetClientID Get Statement Prepare Error", err)
		return nil, nil, nil, err
	}
	for rows.Next() {
		var record int64
		var orgnID int64
		var clientID int64
		err = rows.Scan(&record, &clientID, &orgnID)
		if err != nil {

			logger.Log.Println("GetReopencount rows.next() Error", err)
		}
		records = append(records, record)
		clientIDs = append(clientIDs, clientID)
		orgnIDs = append(orgnIDs, orgnID)

	}
	return clientIDs, orgnIDs, records, nil
}
func (mdao DbConn) GetRecordTypeID(ClientID int64, OrgnIDTypeID int64, recordID int64) (int64, error) {
	logger.Log.Println("In side GetRecordTypeID")
	logger.Log.Println("In side GetRecordTypeID----->", recordID)
	var typeID int64
	var sql = "SELECT recorddiffid FROM maprecordtorecorddifferentiation WHERE clientid=? AND mstorgnhirarchyid=? AND recordid=? AND recorddifftypeid=2 AND islatest=1"
	rows, err := mdao.DB.Query(sql, ClientID, OrgnIDTypeID, recordID)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetRecordTypeID Get Statement Prepare Error", err)
		return typeID, err
	}
	for rows.Next() {
		err = rows.Scan(&typeID)
		logger.Log.Println("GetRecordTypeID rows.next() Error", err)
	}
	return typeID, nil
}
func (mdao DbConn) GetRecordCurrentGrpID(ClientID int64, OrgnIDTypeID int64, recordID int64) (int64, int64, error) {
	logger.Log.Println("In side getStageID")
	logger.Log.Println("In side getStageID----->", recordID)
	var grpID int64
	var userID int64
	var sql = "select usergroupid,userid from trnrecord where clientid=? and mstorgnhirarchyid=? and id=? and deleteflg=0 limit 1;"
	rows, err := mdao.DB.Query(sql, ClientID, OrgnIDTypeID, recordID)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("getStageID Get Statement Prepare Error", err)
		return grpID, userID, err
	}
	for rows.Next() {
		err = rows.Scan(&grpID, &userID)
		logger.Log.Println("getStageID rows.next() Error", err)
	}
	return grpID, userID, nil
}
func (mdao DbConn) GetWokinglabel(ClientID int64, OrgnIDTypeID int64, RecordID int64) (int64, int64, error) {
	logger.Log.Println("In side GetWokinglabel")
	logger.Log.Println("In side GetWokinglabel----->", RecordID)
	var workingtypeID int64
	var workingcatID int64
	var sql = "SELECT recorddifftypeid,recorddiffid FROM maprecordtorecorddifferentiation WHERE clientid=? AND mstorgnhirarchyid=? AND recordid =? AND isworking=1 order by id desc limit 1"
	rows, err := mdao.DB.Query(sql, ClientID, OrgnIDTypeID, RecordID)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetWokinglabel Get Statement Prepare Error", err)
		return workingtypeID, workingcatID, err
	}
	for rows.Next() {
		err = rows.Scan(&workingtypeID, &workingcatID)
		logger.Log.Println("GetWokinglabel rows.next() Error", err)
	}
	return workingtypeID, workingcatID, nil
}
