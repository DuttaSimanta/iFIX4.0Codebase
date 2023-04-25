package dao

import (
	"database/sql"
	// "errors"
	// "log"
	Logger "src/logger"
	"strconv"
	// "golang.org/x/crypto/bcrypt"
)

func GetUserDetails(db *sql.DB, clientID int64, mstOrgnHirarchyId int64, grpID []int64) ([]string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, error) {
	var firstnames []string
	var lastnames []string
	var loginnames []string
	var useremails []string
	var pmobilenos []string
	var smobilenos []string
	var divisions []string
	var brands []string
	var designations []string
	var citys []string
	var branchs []string
	var vipusers []string
	var usertypes []string

	var ids string = ""
	for i, groupid := range grpID {
		if i > 0 {
			ids += ","
		}
		ids += strconv.Itoa(int(groupid))
	}
	var selectUserQuery = "select distinct a.firstname, a.lastname, a.loginname, a.useremail, a.usermobileno, a.secondaryno, a.division, a.brand, a.designation, a.city, a.branch, a.vipuser, a.usertype FROM mstclientuser a, mapclientuserroleuser b, mstgroupmember c where a.clientid = b.clientid and a.mstorgnhirarchyid= b.mstorgnhirarchyid and a.id = b.userid and  a.deleteflag = 0 and b.deleteflg = 0 and a.clientid = c.clientid and a.mstorgnhirarchyid = c.mstorgnhirarchyid and a.id = c.userid and c.groupid in(" + ids + ")  and c.deleteflg = 0 and a.clientid = ? and a.mstorgnhirarchyid = ? ;"
	userResult, err := db.Query(selectUserQuery, clientID, mstOrgnHirarchyId)
	if err != nil {
		Logger.Log.Println(err)

		return firstnames, lastnames, loginnames, useremails, pmobilenos, smobilenos, divisions, brands, designations, citys, branchs, vipusers, usertypes, err
	}
	defer userResult.Close()
	for userResult.Next() {
		var firstName string
		var lastName string
		var loginName string
		var userEmail string
		var pMobileNo string
		var sMobileNo string
		var division string
		var brand string
		var designation string
		var city string
		var branch string
		var vipuser string
		var usertype string

		//	var  diffTypeId int64
		err = userResult.Scan(&firstName, &lastName, &loginName, &userEmail, &pMobileNo, &sMobileNo, &division, &brand, &designation, &city, &branch, &vipuser, &usertype)
		if err != nil {
			Logger.Log.Println(err)

			return firstnames, lastnames, loginnames, useremails, pmobilenos, smobilenos, divisions, brands, designations, citys, branchs, vipusers, usertypes, err
		}
		//headerName = append(headerName, header)
		firstnames = append(firstnames, firstName)
		lastnames = append(lastnames, lastName)
		loginnames = append(loginnames, loginName)
		useremails = append(useremails, userEmail)
		pmobilenos = append(pmobilenos, pMobileNo)
		smobilenos = append(smobilenos, sMobileNo)
		divisions = append(divisions, division)
		brands = append(brands, brand)
		designations = append(designations, designation)
		citys = append(citys, city)
		branchs = append(branchs, branch)
		vipusers = append(vipusers, vipuser)
		usertypes = append(usertypes, usertype)

	}
	return firstnames, lastnames, loginnames, useremails, pmobilenos, smobilenos, divisions, brands, designations, citys, branchs, vipusers, usertypes, err
}
