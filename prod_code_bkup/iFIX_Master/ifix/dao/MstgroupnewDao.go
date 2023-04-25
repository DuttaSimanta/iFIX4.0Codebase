package dao

import (
	"database/sql"
	"errors"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertMstgroupnew = "INSERT INTO mstgroup (clientid, mstorgnhirarchyid, grpid, externalgrpid, supportgrouplevelid, mstclienttimezoneid, reporttimezoneid, email, isworkflow, hascatalog) VALUES (?,?,?,?,?,?,?,?,?,?)"
var duplicateMstgroupupdatenew = "SELECT count(id) total FROM  mstgroup WHERE mstorgnhirarchyid = ? AND grpid = ? AND supportgrouplevelid = ? AND mstclienttimezoneid=? AND reporttimezoneid=? AND email=? AND isworkflow=? AND hascatalog=? AND deleteflg = 0 AND activeflg=1"
var duplicateMstgroupnew="SELECT count(id) total FROM  mstgroup WHERE mstorgnhirarchyid = ? AND grpid=? AND deleteflg = 0 AND activeflg=1"

//var getMstgroup = "SELECT id as Id, clientid as Clientid, mstorgnhirarchyid as Mstorgnhirarchyid, groupname as Groupname, activeflg as Activeflg FROM mstgroup WHERE clientid = ? AND mstorgnhirarchyid = ?  AND deleteflg =0 and activeflg=1 ORDER BY id DESC LIMIT ?,?"
//var getMstgroupcount = "SELECT count(id) total FROM  mstgroup WHERE clientid = ? AND mstorgnhirarchyid = ? AND  deleteflg =0 and activeflg=1"
var updateMstgroupnew = "UPDATE mstgroup SET mstorgnhirarchyid = ?, grpid = ?,supportgrouplevelid=?,mstclienttimezoneid=?,reporttimezoneid=?, email=?, hascatalog=? WHERE externalgrpid = ? "
var deleteMstgroupnew = "UPDATE mstgroup SET deleteflg = '1' WHERE externalgrpid = ? "

 
// All methods defination with transaction

func CheckDuplicateMstgroupupdatenewwithtransaction(tx *sql.Tx, tz *entities.ClientsupportgroupnewEntity) (entities.MstgroupnewEntities, error) {
	logger.Log.Println("In side CheckDuplicateMstgroupnewwithtransaction")
	value := entities.MstgroupnewEntities{}
	err := tx.QueryRow(duplicateMstgroupupdatenew,tz.Mstorgnhirarchyid, tz.Supportgroupid, tz.Supportgrouplevelid, tz.Mstclienttimezoneid, tz.Reporttimezoneid, tz.Email,tz.Isworkflow,tz.Hascatalog).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMstgroupnewwithtransaction Get Statement Prepare Error", err)
		return value, err
	}
}
func CheckDuplicateMstgroupnewwithtransaction(tx *sql.Tx, tz *entities.ClientsupportgroupnewEntity) (entities.MstgroupnewEntities, error) {
	logger.Log.Println("In side CheckDuplicateMstgroupnewwithtransaction")
	value := entities.MstgroupnewEntities{}
	err := tx.QueryRow(duplicateMstgroupnew,tz.Mstorgnhirarchyid, tz.Supportgroupid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMstgroupnewwithtransaction Get Statement Prepare Error", err)
		return value, err
	}
}

func InsertMstgroupnewwithtransaction(tx *sql.Tx, tz *entities.ClientsupportgroupnewEntity, lastinsertedID int64) (int64, error) {
	stmt, err := tx.Prepare(insertMstgroupnew)
	logger.Log.Println("insertMstgroupnew is ---->", insertMstgroupnew)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println(err)
		return 0, errors.New("SQL Prepare Error")
	}

	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Supportgroupid,lastinsertedID, tz.Supportgrouplevelid, tz.Mstclienttimezoneid, tz.Reporttimezoneid, tz.Email,tz.Isworkflow,tz.Hascatalog)
	if err != nil {
		logger.Log.Println(err)
		return 0, errors.New("SQL Execution Error")
	}
	lastInsertedID, err := res.LastInsertId()
	if err != nil {
		logger.Log.Println(err)
		return 0, errors.New("SQL Last ID fetch Error")
	}

	return lastInsertedID, nil
}

/*func GetAllMstgroupwithtransaction(tx *sql.Tx, page *entities.MstgroupEntity) ([]entities.MstgroupEntity, error) {
	logger.Log.Println("In side GelAllMstgroup")
	values := []entities.MstgroupEntity{}
	rows, err := tx.Query(getMstgroup, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstgroup Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstgroupEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Groupname, &value.Activeflg)
		values = append(values, value)
	}
	return values, nil
}*/

func DeleteMstgroupnewwithtransaction(tx *sql.Tx, tz *entities.ClientsupportgroupnewEntity) error {
	logger.Log.Println("In side DeleteMstgroupnew")
	stmt, err := tx.Prepare(deleteMstgroupnew)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMstgroupnew Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMstgroupnew Execute Statement  Error", err)
		return err
	}
	return nil
}

func UpdateMstgroupnewwithtransaction(tx *sql.Tx, tz *entities.ClientsupportgroupnewEntity) error {
	logger.Log.Println("In side UpdateMstgroupnewwithtransaction")
	stmt, err := tx.Prepare(updateMstgroupnew)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println(err)
		return errors.New("SQL Prepare Error")
	}

	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Supportgroupid, tz.Supportgrouplevelid, tz.Mstclienttimezoneid, tz.Reporttimezoneid, tz.Email,tz.Hascatalog, tz.Id)
	if err != nil {
		logger.Log.Println(err)
		return errors.New("SQL Execution Error")
	}

	return nil
}