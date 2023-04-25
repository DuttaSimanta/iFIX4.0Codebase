package dao

import (
	"database/sql"
	"errors"
	"fmt"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
)

var userinsert = "INSERT INTO mstclientuser(clientid,mstorgnhirarchyid,loginname,name,useremail,usermobileno,password,secondaryno,division,brand,city,designation,branch,vipuser,usertype,firstname,lastname,relmanagerid,usermfa,createtype) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
var mstuserinsert = "INSERT INTO mstuser(clientid,mstorgnhirarchyid,loginname,name,useremail,usermobileno,password,secondaryno,division,brand,city,designation,branch,vipuser,usertype,externaluserid,firstname,lastname) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

var userdelete = "UPDATE mstclientuser SET deleteflag=1 WHERE id=?"
var mstuserdelete = "UPDATE mstuser SET deleteflg=1 WHERE externaluserid=?"
var userupdate = "UPDATE mstclientuser SET clientid=?,mstorgnhirarchyid=?,firstname=? ,lastname=?,name=?,useremail=?,usermobileno=?,secondaryno=?,division=?,brand=?,city=?,designation=?,branch=?,vipuser=?,usertype=?,relmanagerid=?,usermfa=?,createtype=? WHERE id=?"
var usercolorupdate = "UPDATE mstclientuser SET color=? WHERE id=? "
var mstuserupdate = "UPDATE mstuser SET clientid=?,mstorgnhirarchyid=?,firstname=? ,lastname=?,name=?,useremail=?,usermobileno=?,secondaryno=?,division=?,brand=?,city=?,designation=?,branch=?,vipuser=?,usertype=? WHERE externaluserid=?"

var duplicateuser = "SELECT count(id) total FROM mstclientuser WHERE clientid=? AND mstorgnhirarchyid=? AND loginname=? AND deleteflag=0 AND activeflag=1"
var duplicatemstuser = "SELECT count(id) total FROM mstuser WHERE clientid=? AND mstorgnhirarchyid=?  AND loginname=? AND deleteflg=0 AND activeflg=1"

//var usergetcount = "SELECT count(a.id) total FROM mstclientuser a,mstclient b, mstorgnhierarchy d where a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.clientid=? AND a.mstorgnhirarchyid=? AND a.deleteflag =0 AND activeflag=1"

//var getuser = "SELECT a.id as ID,a.clientid as ClientID,a.mstorgnhirarchyid as MstorgnhirarchyID,a.loginname as Loginname,COALESCE(a.name,'NA') as Name,COALESCE(a.useremail,'NA') as Useremail,COALESCE(a.usermobileno,'NA') as Usermobileno,b.name as Clientname, d.name as Orgname,COALESCE(a.secondaryno,'NA') as Secondaryno,COALESCE(a.division,'NA') as Division,COALESCE(a.brand,'NA') as Brand,COALESCE(a.city,'NA') as City,COALESCE(a.designation,'NA') as Designation,COALESCE(a.branch,'NA') as Branch,COALESCE(a.vipuser,'NA') as Vipuser,COALESCE(a.usertype,'NA') as Usertype,COALESCE(a.firstname,'NA') as Firstname,COALESCE(a.lastname,'NA') as Lastname from mstclientuser a,mstclient b, mstorgnhierarchy d where a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.clientid=? AND a.mstorgnhirarchyid=? AND a.deleteflag =0 AND a.activeflag=1 ORDER BY id DESC LIMIT ?,?"
var usersearch = "SELECT a.id as ID, a.name as Name ,a.useremail as Useremail,a.usermobileno as Usermobileno, a.loginname as Loginname FROM mstclientuser a,mapclientuserroleuser b where a.id=b.userid and b.clientid=? and b.mstorgnhirarchyid=?  and (a.loginname like ? or a.name like ?) and b.roleid=? and a.deleteflag=0 and b.deleteflg=0 limit 15"

var usersearchbyorg1 = "SELECT id as ID, name as Name ,useremail as Useremail,usermobileno as Usermobileno, loginname as Loginname,COALESCE(vipuser,'NA') vipuser,COALESCE(branch,'NA') branch,coalesce(firstname,'NA') firstname,COALESCE(lastname,'NA') lastname  FROM mstclientuser where clientid=? and mstorgnhirarchyid=? and (loginname like ? or NAME like ?) and deleteflag=0 limit 15;"
var usersearchbyorgemail1 = "SELECT id as ID, name as Name ,useremail as Useremail,usermobileno as Usermobileno, loginname as Loginname,COALESCE(vipuser,'NA') vipuser,COALESCE(branch,'NA') branch,coalesce(firstname,'NA') firstname,COALESCE(lastname,'NA') lastname  FROM mstclientuser where clientid=? and mstorgnhirarchyid=? and (useremail like ? or NAME like ?) and deleteflag=0 limit 15;"

var usersearchbyorg = "SELECT distinct a.id as ID, a.name as Name ,a.useremail as Useremail,a.usermobileno as Usermobileno, a.loginname as Loginname,COALESCE(a.vipuser,'NA') vipuser,COALESCE(a.branch,'NA') branch,coalesce(a.firstname,'NA') firstname,COALESCE(a.lastname,'NA') lastname  FROM mstgroupmember b,mstclientuser a where b.clientid=? and b.mstorgnhirarchyid=? and b.userid=a.id and (a.loginname like ? or a.NAME like ?) and b.activeflg=1 and b.deleteflg=0 and a.activeflag=1 and a.deleteflag=0 limit 15;"
var usersearchbyorgemail = "SELECT distinct a.id as ID, a.name as Name ,a.useremail as Useremail,a.usermobileno as Usermobileno, a.loginname as Loginname,COALESCE(a.vipuser,'NA') vipuser,COALESCE(a.branch,'NA') branch,coalesce(a.firstname,'NA') firstname,COALESCE(a.lastname,'NA') lastname  FROM mstgroupmember b,mstclientuser a where b.clientid=? and b.mstorgnhirarchyid=? and b.userid=a.id and (a.useremail like ? or a.NAME like ?) and b.activeflg=1 and b.deleteflg=0 and a.activeflag=1 and a.deleteflag=0 limit 15;"

var updateduplicateuser = "SELECT count(id) total FROM mstclientuser WHERE clientid=? AND mstorgnhirarchyid=? and loginname=? AND id !=? AND deleteflag=0 AND activeflag=1 "
var updateduplicatemstuser = "SELECT count(id) total FROM mstuser WHERE clientid=? AND mstorgnhirarchyid=? AND loginname=? AND deleteflg=0 AND activeflg=1 "

var recordwiseuserinfo = "SELECT a.id as ID,a.clientid as ClientID,a.mstorgnhirarchyid as MstorgnhirarchyID,a.loginname as Loginname,COALESCE(a.name,'NA') as Name,COALESCE(a.useremail,'NA') as Useremail,COALESCE(a.usermobileno,'NA') as Usermobileno,COALESCE(a.secondaryno,'NA') as Secondaryno,COALESCE(a.division,'NA') as Division,COALESCE(a.brand,'NA') as Brand,COALESCE(a.city,'NA') as City,COALESCE(a.designation,'NA') as Designation,COALESCE(a.branch,'NA') as Branch,COALESCE(a.vipuser,'NA') as Vipuser,COALESCE(a.usertype,'NA') as Usertype,COALESCE(a.firstname,'NA') as Firstname,COALESCE(a.lastname,'NA') as Lastname from mstclientuser a,trnrecord b where a.clientid=? AND a.mstorgnhirarchyid=? AND a.deleteflag =0 AND a.activeflag=1 AND b.id=? AND a.id=b.userid"
var idwiseuserinfo = "SELECT a.id as ID,a.clientid as ClientID,a.mstorgnhirarchyid as MstorgnhirarchyID,a.loginname as Loginname,COALESCE(a.name,'NA') as Name,COALESCE(a.useremail,'NA') as Useremail,COALESCE(a.usermobileno,'NA') as Usermobileno,COALESCE(a.secondaryno,'NA') as Secondaryno,COALESCE(a.division,'NA') as Division,COALESCE(a.brand,'NA') as Brand,COALESCE(a.city,'NA') as City,COALESCE(a.designation,'NA') as Designation,COALESCE(a.branch,'NA') as Branch,COALESCE(a.vipuser,'NA') as Vipuser,COALESCE(a.usertype,'NA') as Usertype,COALESCE(a.firstname,'NA') as Firstname,COALESCE(a.lastname,'NA') as Lastname from mstclientuser a where a.clientid=? AND a.mstorgnhirarchyid=? AND a.deleteflag =0 AND a.activeflag=1 AND a.id=?"
var searchloginname = "SELECT a.loginname as Loginname from mstclientuser a where a.clientid=? AND a.mstorgnhirarchyid=? AND a.loginname LIKE ? AND a.activeflag=1 AND a.deleteflag=0 LIMIT 15"
var searchname = "SELECT distinct a.name as Name from mstclientuser a,mstgroupmember b where b.clientid=? AND b.mstorgnhirarchyid=? AND b.userid=a.id AND a.name LIKE ? AND a.activeflag=1 AND a.deleteflag=0 and b.activeflg=1 and b.deleteflg=0 LIMIT 15"
var searchbranch = "SELECT DISTINCT a.branch as Branch from mstclientuser a where a.clientid=? AND a.mstorgnhirarchyid=? AND a.branch LIKE ? AND a.activeflag=1 AND a.deleteflag=0 LIMIT 15"

//var userdefaultgrpupdate = "UPDATE mstclientuser SET groupid=? WHERE id=?"

//var searchloginbygroupids="select a.loginname as Loginname ,a.name as Name from mstclientuser a where a.clientid=? and a.mstorgnhirarchyid=?  and a.id in (select b.userid from mstgroupmember b where b.groupid ? and b.activeflg=1 and b.deleteflg=0) and a.activeflag=1 and a.deleteflag=0 and a.loginname like ?"
//CheckDuplicateCientUser check duplicate record
func (mdao DbConn) CheckDuplicateCientUser(tz *entities.MstClientUserEntity) (entities.MstClientUserEntities, error) {
	logger.Log.Println("duplicateuser Query -->", duplicateuser)
	logger.Log.Println("parameters -->", tz.ClientID, tz.MstorgnhirarchyID, tz.Loginname)
	value := entities.MstClientUserEntities{}
	err := mdao.DB.QueryRow(duplicateuser, tz.ClientID, tz.MstorgnhirarchyID, tz.Loginname).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Print("CheckDuplicateCientUser Get Statement Prepare Error", err)
		return value, err
	}
}

func CheckDuplicateCientUser(tx *sql.Tx, tz *entities.MstClientUserEntity) (int64, error) {
	logger.Log.Println("duplicateuser Query -->", duplicateuser)
	logger.Log.Println("parameters -->", tz.Loginname)
	var total int64
	stmt, err := tx.Prepare(duplicateuser)
	if err != nil {
		logger.Log.Println("Exception in CheckDuplicateCientUser ---->", err)
		return total, err
	}

	rows, err := stmt.Query(tz.ClientID, tz.MstorgnhirarchyID, tz.Loginname)
	if err != nil {
		logger.Log.Println("Exception in CheckDuplicateCientUser Query Statement..", err)
		return total, err
	}

	for rows.Next() {
		if err := rows.Scan(&total); err != nil {
			logger.Log.Println("Error in fetching data")
		}

	}
	log.Println("User count value is :", total)
	return total, nil
}

func CheckDuplicateMstUser(tx *sql.Tx, tz *entities.MstClientUserEntity) (int64, error) {
	logger.Log.Println("duplicatemstuser Query -->", duplicatemstuser)
	logger.Log.Println("parameters -->", tz.ClientID, tz.MstorgnhirarchyID, tz.Loginname)
	var total int64
	stmt, err := tx.Prepare(duplicatemstuser)
	if err != nil {
		logger.Log.Println("Exception in CheckDuplicateMstUser Prepare Statement..", err)
		return total, err
	}

	rows, err := stmt.Query(tz.ClientID,tz.MstorgnhirarchyID, tz.Loginname)
	if err != nil {
		logger.Log.Println("Exception in CheckDuplicateMstUser Query Statement..", err)
		return total, err
	}

	for rows.Next() {
		if err := rows.Scan(&total); err != nil {
			logger.Log.Println("Error in fetching data")
		}

	}
	fmt.Println("User count value is :", total)
	return total, nil
}

func CheckUpdateDuplicateCientUser(tx *sql.Tx, tz *entities.MstClientUserEntity) (int64, error) {
	logger.Log.Println("duplicateuser Query -->", updateduplicateuser)
	logger.Log.Println("parameters -->", tz.ID, tz.Loginname)
	var total int64
	stmt, err := tx.Prepare(updateduplicateuser)
	if err != nil {
		logger.Log.Println("Exception in CheckDuplicateCientUser ---->", err)
		return total, err
	}

	rows, err := stmt.Query(tz.ClientID,tz.MstorgnhirarchyID, tz.Loginname, tz.ID)
	if err != nil {
		logger.Log.Println("Exception in CheckDuplicateCientUser Query Statement..", err)
		return total, err
	}

	for rows.Next() {
		if err := rows.Scan(&total); err != nil {
			logger.Log.Println("Error in fetching data")
		}

	}
	fmt.Println("User count value is :", total)
	return total, nil
}

func CheckUpdateDuplicateMstUser(tx *sql.Tx, tz *entities.MstClientUserEntity) (int64, error) {
	logger.Log.Println("duplicateuser Query -->", updateduplicatemstuser)
	logger.Log.Println("parameters -->", tz.ClientID, tz.MstorgnhirarchyID, tz.Loginname)
	var total int64
	stmt, err := tx.Prepare(updateduplicatemstuser)
	if err != nil {
		logger.Log.Println("Exception in CheckDuplicateCientUser ---->", err)
		return total, err
	}

	rows, err := stmt.Query(tz.ClientID, tz.MstorgnhirarchyID, tz.Loginname)
	if err != nil {
		logger.Log.Println("Exception in CheckDuplicateCientUser Query Statement..", err)
		return total, err
	}

	for rows.Next() {
		if err := rows.Scan(&total); err != nil {
			logger.Log.Println("Error in fetching data")
		}

	}
	fmt.Println("User count value is :", total)
	return total, nil
}

//InsertClientUserData data insertd in mstclientuser table
func (mdao DbConn) InsertClientUserData(data *entities.MstClientUserEntity) (int64, error) {
	logger.Log.Println("userinsert query -->", userinsert)
	logger.Log.Println("parameters -->", data.ClientID, data.MstorgnhirarchyID, data.Loginname, data.Name, data.Useremail, data.Usermobileno, data.Password)
	stmt, err := mdao.DB.Prepare(userinsert)
	defer stmt.Close()
	if err != nil {
		logger.Log.Print("InsertClientUserData Prepare Statement  Error", err)
		return 0, err
	}
	res, err := stmt.Exec(data.ClientID, data.MstorgnhirarchyID, data.Loginname, data.Name, data.Useremail, data.Usermobileno, data.Password)
	if err != nil {
		logger.Log.Print("InsertClientUserData Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedID, err := res.LastInsertId()
	if err != nil {
		logger.Log.Println(err)
		return 0, errors.New("SQL Last ID fetch Error")
	}
	return lastInsertedID, nil
}

//InsertClientUserData data insertd in mstclientuser table
func InsertClientUserData(tx *sql.Tx, data *entities.MstClientUserEntity) (int64, error) {
	logger.Log.Println("userinsert query -->", userinsert)
	logger.Log.Println("parameters -->", data.ClientID, data.MstorgnhirarchyID, data.Loginname, data.Name, data.Useremail, data.Usermobileno, data.Password, data.Secondaryno, data.Division, data.Brand, data.City, data.Designation, data.Branch, data.Vipuser, data.Usertype)
	stmt, err := tx.Prepare(userinsert)
	defer stmt.Close()
	if err != nil {
		logger.Log.Print("InsertClientUserData Prepare Statement  Error", err)
		return 0, err
	}
	res, err := stmt.Exec(data.ClientID, data.MstorgnhirarchyID, data.Loginname, data.Firstname+" "+data.Lastname, data.Useremail, data.Usermobileno, data.Password, data.Secondaryno, data.Division, data.Brand, data.City, data.Designation, data.Branch, data.Vipuser, data.Usertype, data.Firstname, data.Lastname, data.Relmanagerid, data.Mfa, data.Createtype)
	if err != nil {
		logger.Log.Print("InsertClientUserData Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedID, err := res.LastInsertId()
	if err != nil {
		logger.Log.Println(err)
		return 0, errors.New("SQL Last ID fetch Error")
	}
	return lastInsertedID, nil
}

func InsertMstUserData(tx *sql.Tx, data *entities.MstClientUserEntity, lastinsertedID int64) (int64, error) {
	logger.Log.Println("userinsert query -->", userinsert)
	logger.Log.Println("parameters -->", data.ClientID, data.MstorgnhirarchyID, data.Loginname, data.Name, data.Useremail, data.Usermobileno, data.Password, data.Secondaryno, data.Division, data.Brand, data.City, data.Designation, data.Branch, data.Vipuser, data.Usertype, lastinsertedID)
	stmt, err := tx.Prepare(mstuserinsert)
	defer stmt.Close()
	if err != nil {
		logger.Log.Print("InsertMstUserData Prepare Statement  Error", err)
		return 0, err
	}
	res, err := stmt.Exec(data.ClientID, data.MstorgnhirarchyID, data.Loginname, data.Firstname+" "+data.Lastname, data.Useremail, data.Usermobileno, data.Password, data.Secondaryno, data.Division, data.Brand, data.City, data.Designation, data.Branch, data.Vipuser, data.Usertype, lastinsertedID, data.Firstname, data.Lastname)
	if err != nil {
		logger.Log.Print("InsertMstUserData Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedID, err := res.LastInsertId()
	if err != nil {
		logger.Log.Println(err)
		return 0, errors.New("SQL Last ID fetch Error")
	}
	return lastInsertedID, nil
}

//UpdateClientUserData update mstclientuser table
func (mdao DbConn) UpdateClientUserData(data *entities.MstClientUserEntity) error {
	logger.Log.Println("userupdate Query -->", userupdate)
	logger.Log.Println("parameters -->", data.ClientID, data.MstorgnhirarchyID, data.Loginname, data.Name, data.Useremail, data.Usermobileno, data.Password, data.ID)
	stmt, err := mdao.DB.Prepare(userupdate)
	defer stmt.Close()
	if err != nil {
		logger.Log.Print("Update Client User Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(data.ClientID, data.MstorgnhirarchyID, data.Loginname, data.Name, data.Useremail, data.Usermobileno, data.ID)
	if err != nil {
		logger.Log.Print("Update Client User Execute Statement  Error", err)
		return err
	}
	return nil
}
func (mdao DbConn) Updateusercolor(data *entities.MstClientUserEntity) error {
	logger.Log.Println("usercolorupdate Query -->", usercolorupdate)
	logger.Log.Println("parameters -->", data.Color, data.Userid)
	stmt, err := mdao.DB.Prepare(usercolorupdate)
	defer stmt.Close()
	if err != nil {
		logger.Log.Print("Update Client User Color Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(data.Color, data.Userid)
	if err != nil {
		logger.Log.Print("Update Client User color Execute Statement  Error", err)
		return err
	}
	return nil
}

//secondaryno=?,division=?,brand=?,city=?,designation=?,branch=?,vipuser=?,usertype=?
func UpdateClientUserData(tx *sql.Tx, data *entities.MstClientUserEntity) error {
	logger.Log.Println("userupdate Query -->", userupdate)
	logger.Log.Println("parameters -->", data.ClientID, data.MstorgnhirarchyID, data.Loginname, data.Name, data.Useremail, data.Usermobileno, data.Password, data.ID)
	stmt, err := tx.Prepare(userupdate)

	if err != nil {
		logger.Log.Print("Update Client User Prepare Statement  Error", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(data.ClientID, data.MstorgnhirarchyID, data.Firstname, data.Lastname, data.Firstname+" "+data.Lastname, data.Useremail, data.Usermobileno, data.Secondaryno, data.Division, data.Brand, data.City, data.Designation, data.Branch, data.Vipuser, data.Usertype, data.Relmanagerid, data.Mfa, data.Createtype, data.ID)
	if err != nil {
		logger.Log.Print("Update Client User Execute Statement  Error", err)
		return err
	}
	return nil
}
func Updateuserpasswordtransaction(tx *sql.Tx, password string, id int64) error {
	logger.Log.Println("Updateuserpasswordtransaction Query -->", passworuserdupdate)
	stmt, err := tx.Prepare(passworuserdupdate)

	if err != nil {
		logger.Log.Print("Updateuserpasswordtransaction Prepare Statement  Error", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(password, id)
	if err != nil {
		logger.Log.Print("Updateuserpasswordtransaction Execute Statement  Error", err)
		return err
	}
	return nil
}
func UpdateMstUserData(tx *sql.Tx, data *entities.MstClientUserEntity) error {
	logger.Log.Println("userupdate Query -->", userupdate)
	logger.Log.Println("parameters -->", data.ClientID, data.MstorgnhirarchyID, data.Loginname, data.Name, data.Useremail, data.Usermobileno, data.Password, data.ID)
	stmt, err := tx.Prepare(mstuserupdate)
	defer stmt.Close()
	if err != nil {
		logger.Log.Print("Update Client User Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(data.ClientID, data.MstorgnhirarchyID, data.Firstname, data.Lastname, data.Firstname+" "+data.Lastname, data.Useremail, data.Usermobileno, data.Secondaryno, data.Division, data.Brand, data.City, data.Designation, data.Branch, data.Vipuser, data.Usertype, data.ID)
	if err != nil {
		logger.Log.Print("Update Client User Execute Statement  Error", err)
		return err
	}
	return nil
}

//DeleteClientUserData update mstclientuser table
func (mdao DbConn) DeleteClientUserData(tz *entities.MstClientUserEntity) error {
	logger.Log.Println("userdelete Query -->", userdelete)
	logger.Log.Println("parameters -->", tz.ID)

	stmt, err := mdao.DB.Prepare(userdelete)
	defer stmt.Close()
	if err != nil {
		logger.Log.Print("Delete Client User Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.ID)
	if err != nil {
		logger.Log.Print("Delete Client User Execute Statement  Error", err)
		return err
	}
	return nil
}

//DeleteClientUserData update mstclientuser table
func DeleteClientUserData(tx *sql.Tx, tz *entities.MstClientUserEntity) error {
	logger.Log.Println("userdelete Query -->", userdelete)
	logger.Log.Println("parameters -->", tz.ID)

	stmt, err := tx.Prepare(userdelete)
	defer stmt.Close()
	if err != nil {
		logger.Log.Print("Delete Client User Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.ID)
	if err != nil {
		logger.Log.Print("Delete Client User Execute Statement  Error", err)
		return err
	}
	return nil
}

func DeleteMstUserData(tx *sql.Tx, tz *entities.MstClientUserEntity) error {
	logger.Log.Println("userdelete Query -->", userdelete)
	logger.Log.Println("parameters -->", tz.ID)

	stmt, err := tx.Prepare(mstuserdelete)
	defer stmt.Close()
	if err != nil {
		logger.Log.Print("Delete Client User Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.ID)
	if err != nil {
		logger.Log.Print("Delete Client User Execute Statement  Error", err)
		return err
	}
	return nil
}

//GetClientUserCount get user count with condition
func (mdao DbConn) GetClientUserCount(tz *entities.MstClientUserEntity, OrgnTypeID int64) (entities.MstClientUserEntities, error) {
	logger.Log.Println("parameters -->", tz.ClientID, tz.MstorgnhirarchyID)
	value := entities.MstClientUserEntities{}
	var params []interface{}
	var usergetcount string
	if OrgnTypeID == 1 {
		usergetcount = "SELECT count(a.id) total FROM mstclientuser a,mstclient b, mstorgnhierarchy d where a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.deleteflag =0 AND activeflag=1"
	} else if OrgnTypeID == 2 {
		usergetcount = "SELECT count(a.id) total FROM mstclientuser a,mstclient b, mstorgnhierarchy d where a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.clientid=? AND a.deleteflag =0 AND activeflag=1"
		params = append(params, tz.ClientID)
	} else {
		usergetcount = "SELECT count(a.id) total FROM mstclientuser a,mstclient b, mstorgnhierarchy d where a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.clientid=? AND a.mstorgnhirarchyid=? AND a.deleteflag =0 AND activeflag=1"
		params = append(params, tz.ClientID)
		params = append(params, tz.MstorgnhirarchyID)
	}
	err := mdao.DB.QueryRow(usergetcount, params...).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		log.Print("GetUserCount Get Statement Prepare Error", err)
		return value, err
	}
}

//GetAllUsers get user count with condition
func (mdao DbConn) GetAllUsers(page *entities.MstClientUserEntity, OrgnType int64) ([]entities.MstClientUserEntity, error) {
	logger.Log.Println("In side dao")
	values := []entities.MstClientUserEntity{}
	var params []interface{}
	var getuser string

	if OrgnType == 1 {
		getuser = "SELECT a.relmanagerid,coalesce(c.name,'') relmanager,a.id as ID,a.clientid as ClientID,a.mstorgnhirarchyid as MstorgnhirarchyID,a.loginname as Loginname,COALESCE(a.name,'NA') as Name,COALESCE(a.useremail,'NA') as Useremail,COALESCE(a.usermobileno,'NA') as Usermobileno,b.name as Clientname, d.name as Orgname,COALESCE(a.secondaryno,'NA') as Secondaryno,COALESCE(a.division,'NA') as Division,COALESCE(a.brand,'NA') as Brand,COALESCE(a.city,'NA') as City,COALESCE(a.designation,'NA') as Designation,COALESCE(a.branch,'NA') as Branch,COALESCE(a.vipuser,'NA') as Vipuser,COALESCE(a.usertype,'NA') as Usertype,COALESCE(a.firstname,'NA') as Firstname,COALESCE(a.lastname,'NA') as Lastname,a.usermfa as mfa,IF(a.usermfa=1,'ENABLE','DISABLE') from mstclientuser a left join mstclientuser c on a.relmanagerid =c.id,mstclient b, mstorgnhierarchy d where a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.deleteflag =0 AND a.activeflag=1 ORDER BY id DESC LIMIT ?,?"
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else if OrgnType == 2 {
		getuser = "SELECT a.relmanagerid,coalesce(c.name,'') relmanager,a.id as ID,a.clientid as ClientID,a.mstorgnhirarchyid as MstorgnhirarchyID,a.loginname as Loginname,COALESCE(a.name,'NA') as Name,COALESCE(a.useremail,'NA') as Useremail,COALESCE(a.usermobileno,'NA') as Usermobileno,b.name as Clientname, d.name as Orgname,COALESCE(a.secondaryno,'NA') as Secondaryno,COALESCE(a.division,'NA') as Division,COALESCE(a.brand,'NA') as Brand,COALESCE(a.city,'NA') as City,COALESCE(a.designation,'NA') as Designation,COALESCE(a.branch,'NA') as Branch,COALESCE(a.vipuser,'NA') as Vipuser,COALESCE(a.usertype,'NA') as Usertype,COALESCE(a.firstname,'NA') as Firstname,COALESCE(a.lastname,'NA') as Lastname,a.usermfa as mfa,IF(a.usermfa=1,'ENABLE','DISABLE') from mstclientuser a left join mstclientuser c on a.relmanagerid =c.id,mstclient b, mstorgnhierarchy d where a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.clientid=? AND a.deleteflag =0 AND a.activeflag=1 ORDER BY id DESC LIMIT ?,?"
		params = append(params, page.ClientID)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else {
		getuser = "SELECT a.relmanagerid,coalesce(c.name,'') relmanager,a.id as ID,a.clientid as ClientID,a.mstorgnhirarchyid as MstorgnhirarchyID,a.loginname as Loginname,COALESCE(a.name,'NA') as Name,COALESCE(a.useremail,'NA') as Useremail,COALESCE(a.usermobileno,'NA') as Usermobileno,b.name as Clientname, d.name as Orgname,COALESCE(a.secondaryno,'NA') as Secondaryno,COALESCE(a.division,'NA') as Division,COALESCE(a.brand,'NA') as Brand,COALESCE(a.city,'NA') as City,COALESCE(a.designation,'NA') as Designation,COALESCE(a.branch,'NA') as Branch,COALESCE(a.vipuser,'NA') as Vipuser,COALESCE(a.usertype,'NA') as Usertype,COALESCE(a.firstname,'NA') as Firstname,COALESCE(a.lastname,'NA') as Lastname,a.usermfa as mfa,IF(a.usermfa=1,'ENABLE','DISABLE') from mstclientuser a left join mstclientuser c on a.relmanagerid =c.id,mstclient b, mstorgnhierarchy d where a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.clientid=? AND a.mstorgnhirarchyid=? AND a.deleteflag =0 AND a.activeflag=1 ORDER BY id DESC LIMIT ?,?"
		params = append(params, page.ClientID)
		params = append(params, page.MstorgnhirarchyID)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	}

	rows, err := mdao.DB.Query(getuser, params...)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllModules Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstClientUserEntity{}
		rows.Scan(&value.Relmanagerid, &value.Relmanager, &value.ID, &value.ClientID, &value.MstorgnhirarchyID, &value.Loginname, &value.Name, &value.Useremail, &value.Usermobileno, &value.Clientname, &value.Orgname, &value.Secondaryno, &value.Division, &value.Brand, &value.City, &value.Designation, &value.Branch, &value.Vipuser, &value.Usertype, &value.Firstname, &value.Lastname, &value.Mfa, &value.MfaName)
		values = append(values, value)
	}
	return values, nil
}

//SearchUser search a specific user using loginname ,roleid ,clientid,orgid
func (mdao DbConn) SearchUser(tz *entities.MstClientUserEntity) ([]entities.MstUserSearchEntity, error) {
	log.Println("In side dao")
	values := []entities.MstUserSearchEntity{}
	rows, err := mdao.DB.Query(usersearch, tz.ClientID, tz.MstorgnhirarchyID, "%"+tz.Loginname+"%", "%"+tz.Loginname+"%", tz.Roleid)
	defer rows.Close()
	if err != nil {
		log.Print("SearchUser Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstUserSearchEntity{}
		rows.Scan(&value.ID, &value.Name, &value.Useremail, &value.Usermobileno, &value.Loginname)
		values = append(values, value)
	}
	return values, nil
}

// search a specific user using loginname ,clientid and orgnid
func (mdao DbConn) SearchUserByOrgnId(tz *entities.MstClientUserEntity) ([]entities.MstUserSearchEntity, error) {
	log.Println("In side dao")
	values := []entities.MstUserSearchEntity{}
	var query string
	if tz.Type == "email" {
		query = usersearchbyorgemail
	} else {
		query = usersearchbyorg
	}
	rows, err := mdao.DB.Query(query, tz.ClientID, tz.MstorgnhirarchyID, "%"+tz.Loginname+"%", "%"+tz.Loginname+"%")
	defer rows.Close()
	if err != nil {
		log.Print("SearchUser Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstUserSearchEntity{}
		rows.Scan(&value.ID, &value.Name, &value.Useremail, &value.Usermobileno, &value.Loginname, &value.Vipuser, &value.Branch, &value.Firstname, &value.Lastname)
		values = append(values, value)
	}
	return values, nil
}
func (mdao DbConn) Searchuserbyclientid(tz *entities.MstClientUserEntity) ([]entities.MstUserSearchEntity, error) {
	log.Println("In side dao")
	values := []entities.MstUserSearchEntity{}
	var query string
	if tz.Type == "email" {
		query = usersearchbyorgemail1
	} else {
		query = usersearchbyorg1
	}
	rows, err := mdao.DB.Query(query, tz.ClientID, tz.MstorgnhirarchyID, "%"+tz.Loginname+"%", "%"+tz.Loginname+"%")
	defer rows.Close()
	if err != nil {
		log.Print("SearchUser Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstUserSearchEntity{}
		rows.Scan(&value.ID, &value.Name, &value.Useremail, &value.Usermobileno, &value.Loginname, &value.Vipuser, &value.Branch, &value.Firstname, &value.Lastname)
		values = append(values, value)
	}
	return values, nil
}

//Recordwiseuserinfo get user information
func (mdao DbConn) Recordwiseuserinfo(page *entities.MstGetUserByRecordidEntity) ([]entities.MstClientUserEntity, error) {
	logger.Log.Println("In side dao")
	values := []entities.MstClientUserEntity{}
	rows, err := mdao.DB.Query(recordwiseuserinfo, page.ClientID, page.MstorgnhirarchyID, page.RecordID)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllModules Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstClientUserEntity{}
		rows.Scan(&value.ID, &value.ClientID, &value.MstorgnhirarchyID, &value.Loginname, &value.Name, &value.Useremail, &value.Usermobileno, &value.Secondaryno, &value.Division, &value.Brand, &value.City, &value.Designation, &value.Branch, &value.Vipuser, &value.Usertype, &value.Firstname, &value.Lastname)
		values = append(values, value)
	}
	return values, nil
}

//Recordwiseuserinfo get user information
func (mdao DbConn) IDwiseuserinfo(page *entities.MstGetUserByRecordidEntity) ([]entities.MstClientUserEntity, error) {
	logger.Log.Println("In side IDwiseuserinfo dao")
	values := []entities.MstClientUserEntity{}
	rows, err := mdao.DB.Query(idwiseuserinfo, page.ClientID, page.MstorgnhirarchyID, page.ID)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("IDwiseuserinfo Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstClientUserEntity{}
		rows.Scan(&value.ID, &value.ClientID, &value.MstorgnhirarchyID, &value.Loginname, &value.Name, &value.Useremail, &value.Usermobileno, &value.Secondaryno, &value.Division, &value.Brand, &value.City, &value.Designation, &value.Branch, &value.Vipuser, &value.Usertype, &value.Firstname, &value.Lastname)
		values = append(values, value)
	}
	return values, nil
}

func (mdao DbConn) SearchLoginName(tz *entities.MstClientUserEntity) ([]entities.LoginNameSearchEntity, error) {
	log.Println("In side dao")
	values := []entities.LoginNameSearchEntity{}
	rows, err := mdao.DB.Query(searchloginname, tz.ClientID, tz.MstorgnhirarchyID, "%"+tz.Loginname+"%")
	defer rows.Close()
	if err != nil {
		log.Print("SearchUser Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.LoginNameSearchEntity{}
		rows.Scan(&value.Loginname)
		values = append(values, value)
	}
	return values, nil
}
func (mdao DbConn) SearchName(tz *entities.MstClientUserEntity) ([]entities.NameSearchEntity, error) {
	log.Println("In side dao")
	values := []entities.NameSearchEntity{}
	rows, err := mdao.DB.Query(searchname, tz.ClientID, tz.MstorgnhirarchyID, "%"+tz.Name+"%")
	defer rows.Close()
	if err != nil {
		log.Print("SearchUser Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.NameSearchEntity{}
		rows.Scan(&value.Name)
		values = append(values, value)
	}
	return values, nil
}
func (mdao DbConn) SearchBranch(tz *entities.MstClientUserEntity) ([]entities.BranchSearchEntity, error) {
	log.Println("In side dao")
	values := []entities.BranchSearchEntity{}
	rows, err := mdao.DB.Query(searchbranch, tz.ClientID, tz.MstorgnhirarchyID, "%"+tz.Branch+"%")
	defer rows.Close()
	if err != nil {
		log.Print("SearchUser Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.BranchSearchEntity{}
		rows.Scan(&value.Branch)
		values = append(values, value)
	}
	return values, nil
}
func (mdao DbConn) SearchLoginamebyGroupids(tz *entities.MstClientUserEntity, ids string) ([]entities.LoginnameAndNameEntity, error) {
	log.Println("In side dao")
	//var searchloginbygroupids="select a.id as Id,d.groupid as Groupid,c.supportgroupname as Groupname, a.loginname as Loginname ,a.name as Name from mstclientuser a,mstclientsupportgroup c,mstgroupmember d where a.clientid=? and a.mstorgnhirarchyid=?  and a.id in (select distinct b.userid from mstgroupmember b where b.groupid in ("+ids+") and b.activeflg=1 and b.deleteflg=0) and a.activeflag=1 and a.deleteflag=0 and c.id=d.groupid and a.loginname like ?"
	var searchloginbygroupids = "select distinct b.id as Id,a.groupid as Groupid,c.supportgroupname as Groupname, b.loginname as Loginname ,b.name as Name from mstclientuser b,mstclientsupportgroup c,mstgroupmember a where a.clientid=? and a.mstorgnhirarchyid=?  and a.groupid in (" + ids + ") and a.userid=b.id and a.groupid=c.id and b.activeflag=1 and b.deleteflag=0 and a.activeflg=1 and a.deleteflg=0 and c.activeflg=1 and c.deleteflg=0 and b.loginname like ?"

	values := []entities.LoginnameAndNameEntity{}
	rows, err := mdao.DB.Query(searchloginbygroupids, tz.ClientID, tz.MstorgnhirarchyID, "%"+tz.Loginname+"%")
	defer rows.Close()
	if err != nil {
		log.Print("SearchUser Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.LoginnameAndNameEntity{}
		rows.Scan(&value.Id, &value.Groupid, &value.Groupname, &value.Loginname, &value.Name)
		values = append(values, value)
	}
	return values, nil
}
func UpdateTrnreordDeleteUserID(tx *sql.Tx, tz *entities.MstClientUserEntity) error {
	logger.Log.Println("UpdateTrnreordDeleteUserID parameters -->", tz.ID)
	var query = "UPDATE trnrecord SET isuseriddelete='Y' WHERE userid=?"
	stmt, err := tx.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Log.Print("UpdateTrnreordDeleteUserID Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.ID)
	if err != nil {
		logger.Log.Print("UpdateTrnreordDeleteUserID Execute Statement  Error", err)
		return err
	}
	return nil
}

func UpdateTrnreordDeleteOriginalUserID(tx *sql.Tx, tz *entities.MstClientUserEntity) error {
	logger.Log.Println("UpdateTrnreordDeleteOriginalUserID parameters -->", tz.ID)
	var query = "UPDATE trnrecord SET isorginaluseriddelete='Y' WHERE originaluserid=?"
	stmt, err := tx.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Log.Print("UpdateTrnreordDeleteOriginalUserID Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.ID)
	if err != nil {
		logger.Log.Print("UpdateTrnreordDeleteOriginalUserID Execute Statement  Error", err)
		return err
	}
	return nil
}

func UpdateRecordfulldetailsDeleteRequesterid(tx *sql.Tx, tz *entities.MstClientUserEntity) error {
	logger.Log.Println("UpdateRecordfulldetailsDeleteRequesterid parameters -->", tz.ID)
	var query = "UPDATE recordfulldetails SET isrequestoriddelete='Y' WHERE requestorid=?"
	stmt, err := tx.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Log.Print("UpdateRecordfulldetailsDeleteRequesterid Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.ID)
	if err != nil {
		logger.Log.Print("UpdateRecordfulldetailsDeleteRequesterid Execute Statement  Error", err)
		return err
	}
	return nil
}

func UpdateRecordfulldetailsDeleteOriginalid(tx *sql.Tx, tz *entities.MstClientUserEntity) error {
	logger.Log.Println("UpdateRecordfulldetailsDeleteOriginalid parameters -->", tz.ID)
	var query = "UPDATE recordfulldetails SET isorgcreatoriddelete='Y' WHERE orgcreatorid=?"
	stmt, err := tx.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Log.Print("UpdateRecordfulldetailsDeleteOriginalid Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.ID)
	if err != nil {
		logger.Log.Print("UpdateRecordfulldetailsDeleteOriginalid Execute Statement  Error", err)
		return err
	}
	return nil
}

//func (mdao DbConn) UpdateuserDefaultgrp(data *entities.MstClientUserEntity) error {
//	logger.Log.Println("userDefaultGrpupdate Query -->", userdefaultgrpupdate)
//	logger.Log.Println("parameters -->", data.Groupid)
//	stmt, err := mdao.DB.Prepare(userdefaultgrpupdate)
//	defer stmt.Close()
//	if err != nil {
//		logger.Log.Print("Update Client User DefaultGrp Prepare Statement  Error", err)
//		return err
//	}
//	_, err = stmt.Exec(data.Groupid, data.Userid)
//	if err != nil {
//		logger.Log.Print("Update Client User DefaultGrp Execute Statement  Error", err)
//		return err
//	}
//	return nil
//}
