package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
)

var insertGroupmember = "INSERT INTO mstgroupmember (clientid, mstorgnhirarchyid, groupid, userid) VALUES (?,?,?,?)"
var duplicateGroupmember = "SELECT count(id) total FROM  mstgroupmember WHERE clientid = ? AND mstorgnhirarchyid = ? AND groupid = ? AND userid = ? AND deleteflg = 0 AND activeflg=1"

//var getGroupmember = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.groupid as Groupid, a.userid as Refuserid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,e.supportgroupname as Supportgroupname,f.name as Username,f.loginname FROM mstgroupmember a,mstclient b,mstorgnhierarchy c,mstclientsupportgroup e,mstclientuser f WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.groupid=e.id AND a.userid=f.id AND e.deleteflg =0 and e.activeflg=1  AND f.deleteflag =0 and f.activeflag=1 ORDER BY a.id DESC LIMIT ?,?"
var getGroupmember = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.groupid as Groupid, a.userid as Refuserid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,e.name as Supportgroupname,f.name as Username,f.loginname FROM mstgroupmember a,mstclient b,mstorgnhierarchy c,mstsupportgrp e,mstclientuser f WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.groupid=e.id AND a.userid=f.id AND e.deleteflg =0 and e.activeflg=1  AND f.deleteflag =0 and f.activeflag=1 ORDER BY a.id DESC LIMIT ?,?"
var getGroupmembercount = "SELECT count(a.id) as total FROM mstgroupmember a,mstclient b,mstorgnhierarchy c,mstsupportgrp e,mstclientuser f WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.groupid=e.id AND a.userid=f.id AND e.deleteflg =0 and e.activeflg=1  AND f.deleteflag =0 and f.activeflag=1"
var updateGroupmember = "UPDATE mstgroupmember SET mstorgnhirarchyid = ?, groupid = ?, userid = ? WHERE id = ? "
var deleteGroupmember = "UPDATE mstgroupmember SET deleteflg = '1' WHERE id = ? "
var usersearchbygroup = "SELECT a.id as ID, a.name as Name,COALESCE(a.vipuser,'NA') vipuser ,a.useremail as Useremail,a.usermobileno as Usermobileno, a.loginname as Loginname,COALESCE(a.branch,'NA') branch,coalesce(a.firstname,'NA') firstname,COALESCE(a.lastname,'NA') lastname FROM mstgroupmember b,mstclientuser a where b.userid=a.id and b.clientid=? and b.mstorgnhirarchyid=?   and b.groupid=?  and b.activeflg=1 and b.deleteflg=0 and (a.loginname like ? or a.name like ?) and a.activeflag=1 and a.deleteflag=0 limit 15;"
var usersearchbygroupemail = "SELECT a.id as ID, a.name as Name,COALESCE(a.vipuser,'NA') vipuser ,a.useremail as Useremail,a.usermobileno as Usermobileno, a.loginname as Loginname,COALESCE(a.branch,'NA') branch,coalesce(a.firstname,'NA') firstname,COALESCE(a.lastname,'NA') lastname FROM mstgroupmember b,mstclientuser a where b.userid=a.id and b.clientid=? and b.mstorgnhirarchyid=?   and b.groupid=?  and b.activeflg=1 and b.deleteflg=0 and (a.useremail like ? or a.name like ?) and a.activeflag=1 and a.deleteflag=0 limit 15;"
var duplicateclientuser = "SELECT count(id) total FROM  mstclientsupportgroup WHERE clientid = ? AND mstorgnhirarchyid = ? AND grpid = ? AND deleteflg = 0 AND activeflg=1"
var usersearchbygroupid = "SELECT  a.name as Name,a.loginname as Loginname FROM mstgroupmember b,mstclientuser a where b.userid=a.id and b.clientid=? and b.mstorgnhirarchyid=?   and b.groupid=?  and b.activeflg=1 and b.deleteflg=0  and a.activeflag=1 and a.deleteflag=0 ;"

// search a specific user using loginname ,clientid , orgnid and groupid
var getGrpmember = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.groupid as Groupid, a.userid as Refuserid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,e.name as Supportgroupname,f.name as Username,f.loginname FROM mstgroupmember a,mstclient b,mstorgnhierarchy c,mstsupportgrp e,mstclientuser f WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.groupid=e.id AND a.userid=f.id AND e.deleteflg =0 and e.activeflg=1  AND f.deleteflag =0 and f.activeflag=1 ORDER BY a.id DESC LIMIT ?,?"

var getGrpmembercount = "SELECT count(a.id) as total FROM mstgroupmember a,mstclient b,mstorgnhierarchy c,mstsupportgrp e,mstclientuser f WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.groupid=e.id AND a.userid=f.id AND e.deleteflg =0 and e.activeflg=1  AND f.deleteflag =0 and f.activeflag=1 "
var searchanalyst = "SELECT distinct a.id as ID, a.name as Name ,a.useremail as Useremail  FROM mstgroupmember b,mstclientuser a,mstclientsupportgroup c where b.clientid=? and b.mstorgnhirarchyid=? and b.userid=a.id and b.groupid=c.grpid and a.useremail like ? and c.supportgrouplevelid >1 and b.activeflg=1 and b.deleteflg=0 and a.activeflag=1 and a.deleteflag=0 and c.activeflg=1 and c.deleteflg=0 limit 15;"

func (mdao DbConn) SearchAnalystOrgWise(tz *entities.GroupmemberEntity) ([]entities.MstUserSearchEntity, error) {
	log.Println("In side dao")
	values := []entities.MstUserSearchEntity{}
	rows, err := mdao.DB.Query(searchanalyst, tz.Clientid, tz.Mstorgnhirarchyid, "%"+tz.Email+"%")
	defer rows.Close()
	if err != nil {
		log.Print("SearchUser Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstUserSearchEntity{}
		rows.Scan(&value.ID, &value.Name, &value.Useremail)
		values = append(values, value)
	}
	return values, nil
}
func (mdao DbConn) SearchUserByGroupId(tz *entities.GroupmemberEntity) ([]entities.MstUserSearchEntity, error) {
	log.Println("In side dao")
	values := []entities.MstUserSearchEntity{}
	var query string
	if tz.Type == "email" {
		query = usersearchbygroupemail
	} else {
		query = usersearchbygroup
	}
	rows, err := mdao.DB.Query(query, tz.Clientid, tz.Mstorgnhirarchyid, tz.Groupid, "%"+tz.Loginname+"%", "%"+tz.Loginname+"%")
	defer rows.Close()
	if err != nil {
		log.Print("SearchUser Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstUserSearchEntity{}
		rows.Scan(&value.ID, &value.Name, &value.Vipuser, &value.Useremail, &value.Usermobileno, &value.Loginname, &value.Branch, &value.Firstname, &value.Lastname)
		values = append(values, value)
	}
	return values, nil
}
func (mdao DbConn) Searchuserdetailsbygroupid(tz *entities.GroupmemberEntity) ([]entities.MstUserSearchEntity, error) {
	log.Println("In side dao")
	values := []entities.MstUserSearchEntity{}

	rows, err := mdao.DB.Query(usersearchbygroupid, tz.Clientid, tz.Mstorgnhirarchyid, tz.Groupid)
	defer rows.Close()
	if err != nil {
		log.Print("Searchuserdetailsbygroupdd Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstUserSearchEntity{}
		rows.Scan(&value.Name, &value.Loginname)
		values = append(values, value)
	}
	return values, nil
}
func (mdao DbConn) Groupbyuserwise(tz *entities.GroupmemberEntity) ([]entities.ClientsupportgroupsingleEntity, error) {
	log.Println("In side dao")
	values := []entities.ClientsupportgroupsingleEntity{}
	//var query = "SELECT distinct a.groupid,b.supportgrouplevelid FROM mstgroupmember a,mstclientsupportgroup b where a.clientid=? and a.mstorgnhirarchyid=? and a.userid=? and a.groupid=b.grpid and a.activeflg=1 and a.deleteflg=0 and b.activeflg=1 and b.deleteflg=0 order by b.supportgrouplevelid,a.groupid desc;"
	var query = "SELECT distinct a.groupid,b.supportgrouplevelid,c.name FROM mstgroupmember a,mstclientsupportgroup b,mstsupportgrp c where a.clientid=? and a.mstorgnhirarchyid=? and a.userid=? and a.groupid=b.grpid and a.groupid=c.id and a.activeflg=1 and a.deleteflg=0 and b.activeflg=1 and b.deleteflg=0 and c.activeflg=1 and c.deleteflg=0 order by b.supportgrouplevelid desc;"
	rows, err := mdao.DB.Query(query, tz.Clientid, tz.Mstorgnhirarchyid, tz.Refuserid)
	defer rows.Close()
	if err != nil {
		log.Print("Searchuserdetailsbygroupdd Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.ClientsupportgroupsingleEntity{}
		rows.Scan(&value.Id, &value.Levelid, &value.Groupname)
		values = append(values, value)
	}
	return values, nil
}
func (mdao DbConn) Workflowgroupbyuserwise(tz *entities.GroupmemberEntity) ([]entities.ClientsupportgroupsingleEntity, error) {
	log.Println("In side dao")
	values := []entities.ClientsupportgroupsingleEntity{}
	//var query="SELECT a.groupid,b.supportgrouplevelid FROM mstgroupmember a,mstclientsupportgroup b where a.clientid=? and a.mstorgnhirarchyid=? and a.userid=? and a.groupid=b.grpid and a.activeflg=1 and a.deleteflg=0 and b.activeflg=1 and b.deleteflg=0 order by b.supportgrouplevelid desc;"
	var query = "SELECT distinct a.groupid,b.supportgrouplevelid,c.name FROM mstgroupmember a,mstclientsupportgroup b,mstsupportgrp c where a.clientid=? and a.mstorgnhirarchyid in(" + tz.Mstorgnhirarchyids + ") and a.userid=? and a.groupid=b.grpid and a.groupid=c.id and a.activeflg=1 and a.deleteflg=0 and b.isworkflow='Y' and b.activeflg=1 and b.deleteflg=0 and c.activeflg=1 and c.deleteflg=0 order by b.supportgrouplevelid desc;"
	rows, err := mdao.DB.Query(query, tz.Clientid, tz.Refuserid)
	defer rows.Close()
	if err != nil {
		log.Print("Searchuserdetailsbygroupdd Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.ClientsupportgroupsingleEntity{}
		rows.Scan(&value.Id, &value.Levelid, &value.Groupname)
		values = append(values, value)
	}
	return values, nil
}
func (dbc DbConn) CheckDuplicateGroupmember(tz *entities.GroupmemberEntity) (entities.GroupmemberEntities, error) {
	logger.Log.Println("In side CheckDuplicateGroupmember")
	value := entities.GroupmemberEntities{}
	err := dbc.DB.QueryRow(duplicateGroupmember, tz.Clientid, tz.Mstorgnhirarchyid, tz.Groupid, tz.Refuserid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateGroupmember Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertGroupmember(tz *entities.GroupmemberEntity) (int64, error) {
	logger.Log.Println("In side InsertGroupmember")
	logger.Log.Println("Query -->", insertGroupmember)
	stmt, err := dbc.DB.Prepare(insertGroupmember)

	if err != nil {
		logger.Log.Println("InsertGroupmember Prepare Statement  Error", err)
		return 0, err
	}
	defer stmt.Close()
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Groupid, tz.Refuserid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Groupid, tz.Refuserid)
	if err != nil {
		logger.Log.Println("InsertGroupmember Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}
func (dbc DbConn) InsertGroupmembertransaction(tz *entities.GroupmemberEntity, tx *sql.Tx) (int64, error) {
	logger.Log.Println("In side InsertGroupmember")
	logger.Log.Println("Query -->", insertGroupmember)
	stmt, err := tx.Prepare(insertGroupmember)

	if err != nil {
		logger.Log.Println("InsertGroupmembertransaction Prepare Statement  Error", err)
		return 0, err
	}
	defer stmt.Close()
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Groupid, tz.Refuserid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Groupid, tz.Refuserid)
	if err != nil {
		logger.Log.Println("InsertGroupmembertransaction Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) UpdateGroupmember(tz *entities.GroupmemberEntity) error {
	logger.Log.Println("In side UpdateGroupmember")
	stmt, err := dbc.DB.Prepare(updateGroupmember)

	if err != nil {
		logger.Log.Println("UpdateGroupmember Prepare Statement  Error", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Groupid, tz.Refuserid, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateGroupmember Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteGroupmember(tz *entities.GroupmemberEntity) error {
	logger.Log.Println("In side DeleteGroupmember")
	stmt, err := dbc.DB.Prepare(deleteGroupmember)

	if err != nil {
		logger.Log.Println("DeleteGroupmember Prepare Statement  Error", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteGroupmember Execute Statement  Error", err)
		return err
	}
	return nil
}
func (dbc DbConn) CheckClientuser(tz *entities.GroupmemberEntity, i int) (entities.GroupmemberEntities, error) {
	logger.Log.Println("In side CheckDuplicateGroupmember")
	value := entities.GroupmemberEntities{}
	err := dbc.DB.QueryRow(duplicateclientuser, tz.ToClientid, tz.ToMstorgnhirarchyid[i], tz.Groupid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateGroupmember Get Statement Prepare Error", err)
		return value, err
	}
}
func (dbc DbConn) CheckDuplicateGrpmember(tz *entities.GroupmemberEntity, i int, j int) (entities.GroupmemberEntities, error) {
	logger.Log.Println("In side CheckDuplicateGroupmember")
	value := entities.GroupmemberEntities{}
	err := dbc.DB.QueryRow(duplicateGroupmember, tz.ToClientid, tz.ToMstorgnhirarchyid[i], tz.Groupid, tz.Userids[j]).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateGroupmember Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc TxConn) InsertGrpmember(tz *entities.GroupmemberEntity, i int, j int) (int64, error) {
	logger.Log.Println("In side InsertGroupmember")
	logger.Log.Println("Query -->", insertGroupmember)
	stmt, err := dbc.TX.Prepare(insertGroupmember)

	if err != nil {
		logger.Log.Println("InsertGroupmember Prepare Statement  Error", err)
		return 0, err
	}
	defer stmt.Close()
	logger.Log.Println("Parameter -->", tz.ToClientid, tz.ToMstorgnhirarchyid[i], tz.Groupid, tz.Userids[j])
	res, err := stmt.Exec(tz.ToClientid, tz.ToMstorgnhirarchyid[i], tz.Groupid, tz.Userids[j])
	if err != nil {
		logger.Log.Println("InsertGroupmember Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}
func (dbc DbConn) GetAllGrpmember(page *entities.GroupmemberEntity) ([]entities.GroupmemberEntity, error) {
	logger.Log.Println("In side GelAllGroupmember")
	values := []entities.GroupmemberEntity{}
	rows, err := dbc.DB.Query(getGrpmember, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)

	if err != nil {
		logger.Log.Println("GetAllGroupmember Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		value := entities.GroupmemberEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Groupid, &value.Refuserid, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Supportgroupname, &value.Username, &value.Loginname)
		values = append(values, value)
	}
	return values, nil
}
func (dbc DbConn) GetGrpmemberCount(tz *entities.GroupmemberEntity) (entities.GroupmemberEntities, error) {
	logger.Log.Println("In side GetGroupmemberCount")
	value := entities.GroupmemberEntities{}
	err := dbc.DB.QueryRow(getGrpmembercount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetGroupmemberCount Get Statement Prepare Error", err)
		return value, err
	}
}
func (mdao DbConn) GetUserByGroupId(tz *entities.GroupmemberEntity) ([]entities.MstUserSearchEntity, error) {
	log.Println("In side dao")
	values := []entities.MstUserSearchEntity{}
	var query = "SELECT a.id as ID, a.name as Name FROM mstgroupmember b,mstclientuser a where b.userid=a.id and b.clientid=? and b.mstorgnhirarchyid=? and b.groupid=?  and b.activeflg=1 and b.deleteflg=0 and a.activeflag=1 and a.deleteflag=0"
	rows, err := mdao.DB.Query(query, tz.Clientid, tz.Mstorgnhirarchyid, tz.Groupid)
	defer rows.Close()
	if err != nil {
		log.Print("SearchUser Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstUserSearchEntity{}
		rows.Scan(&value.ID, &value.Name)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) GetAllGroupmember(page *entities.GroupmemberEntity, OrgnType int64) ([]entities.GroupmemberEntity, error) {
	logger.Log.Println("In side GelAllGroupmember")
	values := []entities.GroupmemberEntity{}
	var getGroupmember string
	var params []interface{}
	if OrgnType == 1 {
		getGroupmember = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.groupid as Groupid, a.userid as Refuserid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,e.name as Supportgroupname,f.name as Username,f.loginname FROM mstgroupmember a,mstclient b,mstorgnhierarchy c,mstsupportgrp e,mstclientuser f WHERE a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.groupid=e.id AND a.userid=f.id AND e.deleteflg =0 and e.activeflg=1  AND f.deleteflag =0 and f.activeflag=1 ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else if OrgnType == 2 {
		getGroupmember = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.groupid as Groupid, a.userid as Refuserid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,e.name as Supportgroupname,f.name as Username,f.loginname FROM mstgroupmember a,mstclient b,mstorgnhierarchy c,mstsupportgrp e,mstclientuser f WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.groupid=e.id AND a.userid=f.id AND e.deleteflg =0 and e.activeflg=1  AND f.deleteflag =0 and f.activeflag=1 ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else {
		getGroupmember = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.groupid as Groupid, a.userid as Refuserid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,e.name as Supportgroupname,f.name as Username,f.loginname FROM mstgroupmember a,mstclient b,mstorgnhierarchy c,mstsupportgrp e,mstclientuser f WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.groupid=e.id AND a.userid=f.id AND e.deleteflg =0 and e.activeflg=1  AND f.deleteflag =0 and f.activeflag=1 ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Mstorgnhirarchyid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	}
	rows, err := dbc.DB.Query(getGroupmember, params...)

	if err != nil {
		logger.Log.Println("GetAllGroupmember Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		value := entities.GroupmemberEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Groupid, &value.Refuserid, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Supportgroupname, &value.Username, &value.Loginname)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) GetGroupmemberCount(tz *entities.GroupmemberEntity, OrgnTypeID int64) (entities.GroupmemberEntities, error) {
	logger.Log.Println("In side GetGroupmemberCount")
	value := entities.GroupmemberEntities{}
	var getGroupmembercount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getGroupmembercount = "SELECT count(a.id) as total FROM mstgroupmember a,mstclient b,mstorgnhierarchy c,mstsupportgrp e,mstclientuser f WHERE a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.groupid=e.id AND a.userid=f.id AND e.deleteflg =0 and e.activeflg=1  AND f.deleteflag =0 and f.activeflag=1"
	} else if OrgnTypeID == 2 {
		getGroupmembercount = "SELECT count(a.id) as total FROM mstgroupmember a,mstclient b,mstorgnhierarchy c,mstsupportgrp e,mstclientuser f WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.groupid=e.id AND a.userid=f.id AND e.deleteflg =0 and e.activeflg=1  AND f.deleteflag =0 and f.activeflag=1"
		params = append(params, tz.Clientid)
	} else {
		getGroupmembercount = "SELECT count(a.id) as total FROM mstgroupmember a,mstclient b,mstorgnhierarchy c,mstsupportgrp e,mstclientuser f WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.groupid=e.id AND a.userid=f.id AND e.deleteflg =0 and e.activeflg=1  AND f.deleteflag =0 and f.activeflag=1"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getGroupmembercount, params...).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetGroupmemberCount Get Statement Prepare Error", err)
		return value, err
	}
}
