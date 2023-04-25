package dao

import (
	"database/sql"
	"errors"
	"log"
	Logger "src/logger"

	// "strconv"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pwd []byte) string {

	// Use GenerateFromPassword to hash & salt pwd
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}
func GetHeaderName(db *sql.DB, clientID int64, orgID int64) ([]string, error) {
	var headerNames []string
	Logger.Log.Println("GetHeaderName Method")
	/* db, dBerr := config.GetDB()
	if dBerr != nil{
		Logger.Log.Println(dBerr)
		return  headerIds,headerNames, errors.New("ERROR: Unable to connect DB")
	} */
	//defer db.Close()
	var selectHeaderQuery string = "select headername from mstexceltemplate where clientid=? and mstorgnhirarchyid=? and templatetypeid=1 and activeflg=1 and deleteflg=0 order by seqno asc"
	//fetching category header Details and storing into slice
	HeadeResultSet, err := db.Query(selectHeaderQuery, clientID, orgID)
	if err != nil {
		Logger.Log.Println(err)
		return headerNames, errors.New("ERROR: Unable to Fetch Header Details from DB")
	}
	defer HeadeResultSet.Close()
	for HeadeResultSet.Next() {

		var header string
		err = HeadeResultSet.Scan(&header)
		if err != nil {
			Logger.Log.Println(err)
			return headerNames, errors.New("ERROR: Unable to Scan Header Details")
		}
		headerNames = append(headerNames, header)

	}

	return headerNames, nil
}
func UserUploadWithRoleMapAndGroupMap(db *sql.DB, clientID int64, orgID int64, coloumn []string, grpID int64, roleID int64) error {

	FirstName := coloumn[0]
	LastName := coloumn[1]
	loginName := coloumn[2]
	Name := coloumn[0] + " " + coloumn[1]
	UserEmail := coloumn[3]
	UserMobileNo := coloumn[4]
	Password := "12345"
	Password = HashAndSalt([]byte(Password))
	//Passwordactivatedate =
	SecondaryNo := coloumn[5]
	Division := coloumn[6]
	Brand := coloumn[7]
	Designation := coloumn[8]
	City := coloumn[9]
	Branch := coloumn[10]
	VipUser := coloumn[11]
	UserType := coloumn[12]
	var Activeflg int64 = 1
	var Deleteflg int64 = 0
	//Audittransactionid int64
	//Password = HashAndSalt([]byte(Password))
	// db, dBerr := config.GetDB()
	// if dBerr != nil {
	// 	Logger.Log.Println(dBerr)
	// 	return errors.New("ERROR: Unable to connect DB")
	// }
	// defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		Logger.Log.Println(err)
		return err
	}
	Logger.Log.Println("User Details===>", coloumn)
	var loginIDExistCount int64 = 0
	duplicateLoginIDCheckQuery := "select count(loginname) as count from mstclientuser where clientid=? and mstorgnhirarchyid=? and loginname=? and activeflag=1 and deleteflag=0"
	duplicateLoginIDCheckErr := db.QueryRow(duplicateLoginIDCheckQuery, clientID, orgID, loginName).Scan(&loginIDExistCount)

	if duplicateLoginIDCheckErr != nil {
		Logger.Log.Println(duplicateLoginIDCheckErr)
		//return errors.New("ERROR: Unable to Scan Duplicate LoginID")
	}
	var userID int64
	getUserIDQuery := "select id from mstclientuser where clientid=? and mstorgnhirarchyid=? and loginname=? and activeflag=1 and deleteflag=0"
	getUserIDErr := db.QueryRow(getUserIDQuery, clientID, orgID, loginName).Scan(&userID)

	if getUserIDErr != nil {
		Logger.Log.Println(getUserIDErr)
		//return errors.New("ERROR: Unable to Scan Duplicate LoginID")
	}

	var mstuserID int64
	getmstUserIDQuery := "select id from mstuser where clientid=? and mstorgnhirarchyid=? and externaluserid=? and activeflg=1 and deleteflg=0"
	getmstUserIDErr := db.QueryRow(getmstUserIDQuery, clientID, orgID, userID).Scan(&mstuserID)

	if getmstUserIDErr != nil {
		Logger.Log.Println(getmstUserIDErr)
		//return errors.New("ERROR: Unable to Scan Duplicate LoginID")
	}
	if loginIDExistCount > 0 {
		Logger.Log.Println("Duplicate LoginID======================> Updattion is goin on")
		UpdateMSTClientUserQuesry := "UPDATE mstclientuser" +
			" SET firstname = ?,lastname = ?,name = ?,useremail = ?,usermobileno = ?,secondaryno = ?,division = ?,brand = ?," +
			" city = ?,designation = ?,branch = ?,vipuser = ?,usertype = ? WHERE id = ?"

		stmtUpdateMstClientUser, stmtErr := tx.Prepare(UpdateMSTClientUserQuesry)
		if stmtErr != nil {
			Logger.Log.Println(stmtErr)
			return stmtErr
		}
		defer stmtUpdateMstClientUser.Close()
		updateIntoMSTClientUserResultSet, updateErr := stmtUpdateMstClientUser.Exec(FirstName, LastName, Name, UserEmail, UserMobileNo,
			SecondaryNo, Division, Brand, City, Designation, Branch, VipUser, UserType, userID)
		if updateErr != nil {

			Logger.Log.Println(updateErr)
			tx.Rollback()
			return updateErr
		}
		_, lastIdError := updateIntoMSTClientUserResultSet.RowsAffected()
		if lastIdError != nil {

			Logger.Log.Println(lastIdError)
			tx.Rollback()
			return lastIdError
		}
		updateIntoMSTtUserQuesry := "UPDATE mstuser" +
			" SET firstname = ?,lastname = ?,name = ?,useremail = ?,usermobileno = ?,secondaryno = ?,division = ?,brand = ?," +
			" city = ?,designation = ?,branch = ?,vipuser = ?,usertype = ? WHERE id=?"
		stmtupdatetMstUser, stmtErr := tx.Prepare(updateIntoMSTtUserQuesry)
		if stmtErr != nil {
			Logger.Log.Println(stmtErr)
			tx.Rollback()
			return stmtErr
		}
		defer stmtupdatetMstUser.Close()
		_, updatemstuserErr := stmtupdatetMstUser.Exec(FirstName, LastName, Name, UserEmail, UserMobileNo, SecondaryNo, Division, Brand, City, Designation, Branch, VipUser, UserType, mstuserID)
		if updatemstuserErr != nil {

			Logger.Log.Println(updatemstuserErr)
			tx.Rollback()
			return updatemstuserErr
		}
		_, lastIdError = updateIntoMSTClientUserResultSet.RowsAffected()
		if lastIdError != nil {

			Logger.Log.Println(lastIdError)
			tx.Rollback()
			return lastIdError
		}
		//return errors.New("Duplicate LoginID")
	} else {
		insertIntoMSTClientUserQuesry := "INSERT INTO `mstclientuser` (`clientid`,`mstorgnhirarchyid`,`loginname`,`firstname`,`lastname`,`name`,`useremail`,`usermobileno`,`password`,`secondaryno`,`division`,`brand`,`city`,`designation`,`branch`,`vipuser`,`usertype`,`activeflag`,`deleteflag`)VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

		stmtInsertMstClientUser, stmtErr := tx.Prepare(insertIntoMSTClientUserQuesry)
		if stmtErr != nil {
			Logger.Log.Println(stmtErr)
			return stmtErr
		}
		defer stmtInsertMstClientUser.Close()
		insertIntoMSTClientUserResultSet, insertErr := stmtInsertMstClientUser.Exec(clientID, orgID, loginName, FirstName, LastName, Name, UserEmail, UserMobileNo, Password, SecondaryNo, Division, Brand, City, Designation, Branch, VipUser, UserType, Activeflg, Deleteflg)
		if insertErr != nil {

			Logger.Log.Println(insertErr)
			tx.Rollback()
			return insertErr
		}
		lastInsertedMSTClientUser, lastIdError := insertIntoMSTClientUserResultSet.LastInsertId()
		if lastIdError != nil {

			Logger.Log.Println(lastIdError)
			tx.Rollback()
			return lastIdError
		}
		insertIntoMSTtUserQuesry := "INSERT INTO `mstuser` (`clientid`,`mstorgnhirarchyid`,`externaluserid`,`loginname`,`firstname`,`lastname`,`name`,`useremail`,`usermobileno`,`password`,`secondaryno`,`division`,`brand`,`city`,`designation`,`branch`,`vipuser`,`usertype`,`activeflg`,`deleteflg`)VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
		stmtInsertMstUser, stmtErr := tx.Prepare(insertIntoMSTtUserQuesry)
		if stmtErr != nil {
			Logger.Log.Println(stmtErr)
			tx.Rollback()
			return stmtErr
		}
		defer stmtInsertMstUser.Close()
		_, insertmstuserErr := stmtInsertMstUser.Exec(clientID, orgID, lastInsertedMSTClientUser, loginName, FirstName, LastName, Name, UserEmail, UserMobileNo, Password, SecondaryNo, Division, Brand, City, Designation, Branch, VipUser, UserType, Activeflg, Deleteflg)
		if insertmstuserErr != nil {

			Logger.Log.Println(insertmstuserErr)
			tx.Rollback()
			return insertmstuserErr
		}
		insertMapClientUserRoleUserQuery := "INSERT INTO `mapclientuserroleuser`(`clientid`,`mstorgnhirarchyid`,`roleid`,`userid`,`activeflg`,`deleteflg`)VALUES(?,?,?,?,?,?)"

		stmtInsertMapClientUserRoleUser, stmtErr := tx.Prepare(insertMapClientUserRoleUserQuery)
		if stmtErr != nil {
			Logger.Log.Println(stmtErr)
			tx.Rollback()
			return stmtErr
		}
		defer stmtInsertMapClientUserRoleUser.Close()
		_, insertMapClientUserRoleUserErr := stmtInsertMapClientUserRoleUser.Exec(clientID, orgID, roleID, lastInsertedMSTClientUser, Activeflg, Deleteflg)
		if insertMapClientUserRoleUserErr != nil {

			Logger.Log.Println(insertMapClientUserRoleUserErr)
			tx.Rollback()
			return insertMapClientUserRoleUserErr
		}
		insertMstGrpMemberQuery := "INSERT INTO `mstgroupmember`(`clientid`,`mstorgnhirarchyid`,`groupid`,`userid`,`deleteflg`,`activeflg`)VALUES(?,?,?,?,?,?)	"

		stmtInsertMstGrpMember, stmtErr := tx.Prepare(insertMstGrpMemberQuery)
		if stmtErr != nil {
			Logger.Log.Println(stmtErr)
			tx.Rollback()
			return stmtErr
		}
		defer stmtInsertMstGrpMember.Close()
		_, insertMstGrpMemberErr := stmtInsertMstGrpMember.Exec(clientID, orgID, grpID, lastInsertedMSTClientUser, Deleteflg, Activeflg)
		if insertMstGrpMemberErr != nil {

			Logger.Log.Println(insertMstGrpMemberErr)
			tx.Rollback()
			return insertMstGrpMemberErr
		}

	}
	commitErr := tx.Commit()
	if commitErr != nil {
		Logger.Log.Println(commitErr)
		return errors.New("ERROR: Unable to commit  User")

	}

	return nil
}

func GetOrgName(db *sql.DB, clientID int64, mstOrgnHirarchyId int64) (string, error) {
	var orgName string
	var OrgNameQuery string = "SELECT a.name FROM mstorgnhierarchy a  where a.clientid = ? and a.id = ?  and a.activeflg=1 and a.deleteflg=0"
	OrgNameScanErr := db.QueryRow(OrgNameQuery, clientID, mstOrgnHirarchyId).Scan(&orgName)
	if OrgNameScanErr != nil {
		Logger.Log.Println(OrgNameScanErr)
		return orgName, errors.New("ERROR: Scan Error For GetOrgName")
	}
	return orgName, nil
}

// func GetUserDetails(db *sql.DB, clientID int64, mstOrgnHirarchyId int64, grpID []int64) ([]string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, error) {
// 	var firstnames []string
// 	var lastnames []string
// 	var loginnames []string
// 	var useremails []string
// 	var pmobilenos []string
// 	var smobilenos []string
// 	var divisions []string
// 	var brands []string
// 	var designations []string
// 	var citys []string
// 	var branchs []string
// 	var vipusers []string
// 	var usertypes []string

// 	var ids string = ""
// 	for i, groupid := range grpID {
// 		if i > 0 {
// 			ids += ","
// 		}
// 		ids += strconv.Itoa(int(groupid))
// 	}
// 	var selectUserQuery = "select distinct a.firstname, a.lastname, a.loginname, a.useremail, a.usermobileno, a.secondaryno, a.division, a.brand, a.designation, a.city, a.branch, a.vipuser, a.usertype FROM mstclientuser a, mapclientuserroleuser b, mstgroupmember c where a.clientid = b.clientid and a.mstorgnhirarchyid= b.mstorgnhirarchyid and a.id = b.userid and  a.deleteflag = 0 and b.deleteflg = 0 and a.clientid = c.clientid and a.mstorgnhirarchyid = c.mstorgnhirarchyid and a.id = c.userid and c.groupid in(" + ids + ")  and c.deleteflg = 0 and a.clientid = ? and a.mstorgnhirarchyid = ? ;"
// 	userResult, err := db.Query(selectUserQuery, clientID, mstOrgnHirarchyId)
// 	if err != nil {
// 		Logger.Log.Println(err)

// 		return firstnames, lastnames, loginnames, useremails, pmobilenos, smobilenos, divisions, brands, designations, citys, branchs, vipusers, usertypes, err
// 	}
// 	defer userResult.Close()
// 	for userResult.Next() {
// 		var firstName string
// 		var lastName string
// 		var loginName string
// 		var userEmail string
// 		var pMobileNo string
// 		var sMobileNo string
// 		var division string
// 		var brand string
// 		var designation string
// 		var city string
// 		var branch string
// 		var vipuser string
// 		var usertype string

// 		//	var  diffTypeId int64
// 		err = userResult.Scan(&firstName, &lastName, &loginName, &userEmail, &pMobileNo, &sMobileNo, &division, &brand, &designation, &city, &branch, &vipuser, &usertype)
// 		if err != nil {
// 			Logger.Log.Println(err)

// 			return firstnames, lastnames, loginnames, useremails, pmobilenos, smobilenos, divisions, brands, designations, citys, branchs, vipusers, usertypes, err
// 		}
// 		//headerName = append(headerName, header)
// 		firstnames = append(firstnames, firstName)
// 		lastnames = append(lastnames, lastName)
// 		loginnames = append(loginnames, loginName)
// 		useremails = append(useremails, userEmail)
// 		pmobilenos = append(pmobilenos, pMobileNo)
// 		smobilenos = append(smobilenos, sMobileNo)
// 		divisions = append(divisions, division)
// 		brands = append(brands, brand)
// 		designations = append(designations, designation)
// 		citys = append(citys, city)
// 		branchs = append(branchs, branch)
// 		vipusers = append(vipusers, vipuser)
// 		usertypes = append(usertypes, usertype)

// 	}
// 	return firstnames, lastnames, loginnames, useremails, pmobilenos, smobilenos, divisions, brands, designations, citys, branchs, vipusers, usertypes, err
// }
