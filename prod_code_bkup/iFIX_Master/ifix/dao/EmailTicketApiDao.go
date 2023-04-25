package dao

import (
	"database/sql"
	"errors"
	"iFIX/ifix/entities"
	Logger "iFIX/ifix/logger"
	"log"
)

func CheckDuplicateEmailBaseConfig(db *sql.DB, tz *entities.EmailBaseConfig) (int64, error) {

	var count int64
	countQuery := "SELECT count(id) as count from mstrecorddifferentiationtype where parentid=? and clientid=? and mstorgnhirarchyid=? and deleteflg=0 and activeflg=1"
	resulsetErr := db.QueryRow(countQuery, tz.DiffTypeSeq, tz.ClientID, tz.OrgID).Scan(&count)
	if resulsetErr != nil {
		Logger.Log.Println(resulsetErr)
		return count, errors.New("Unable to fetch resultset")
	}

	return count, nil
}
func AddEmailBaseConfig(db *sql.DB, tz *entities.EmailBaseConfig, count int64) error {
	tx, txerr := db.Begin()
	if txerr != nil {
		Logger.Log.Println(txerr)
		return errors.New("ERROR: Unable to create Transaction")
	}
	if count == 0 {
		sql := "INSERT INTO mstrecorddifferentiationtype(`clientid`,`mstorgnhirarchyid`,`typename`,`seqno`,`deleteflg`,`activeflg`,`parentid`)VALUES(?,?,?,201,0,1,?);"
		statementTX, stmnterr := tx.Prepare(sql)
		if stmnterr != nil {
			Logger.Log.Println(stmnterr)
			return errors.New("Unable to Prepare resultset")
		}
		defer statementTX.Close()
		resultSet, execError := statementTX.Exec(tz.ClientID, tz.OrgID, tz.DelimiterHeader, tz.DiffTypeSeq)
		if execError != nil {
			Logger.Log.Println(execError)
			return errors.New("Unable to Exec Query")
		}
		recordDiffTypeID, fetchErr := resultSet.LastInsertId()
		if fetchErr != nil {
			tx.Rollback()
			Logger.Log.Println(fetchErr)
			return errors.New("Unable to Fetch  last inserted ID")
		}

		sql1 := "INSERT INTO mstrecorddifferentiation(`clientid`,`mstorgnhirarchyid`,`recorddifftypeid`,`parentid`,`name`,`seqno`,`deleteflg`,`activeflg`)VALUES(?,?,?,0,?,1,0,1)"
		statementTX1, stmnterr1 := tx.Prepare(sql1)
		if stmnterr1 != nil {
			tx.Rollback()
			Logger.Log.Println(stmnterr1)
			return errors.New("Unable to Prepare resultset")
		}
		defer statementTX1.Close()
		resultSet1, execError1 := statementTX1.Exec(tz.ClientID, tz.OrgID, recordDiffTypeID, tz.DelimiterVal)
		if execError1 != nil {
			tx.Rollback()
			Logger.Log.Println(execError1)
			return errors.New("Unable to Exec Query")
		}
		_, err := resultSet1.RowsAffected()
		if err != nil {
			tx.Rollback()
			Logger.Log.Println(err)
			return errors.New("No rows affected")
		}
	} else {
		var recordDiffTypeID int64
		var count int64

		getRecordDifTypeidQuery := "SELECT id  as Id from mstrecorddifferentiationtype where parentid=? and clientid=? and mstorgnhirarchyid=? and deleteflg=0 and activeflg=1"
		resulsetErr := db.QueryRow(getRecordDifTypeidQuery, tz.DiffTypeSeq, tz.ClientID, tz.OrgID).Scan(&recordDiffTypeID)
		if resulsetErr != nil {
			Logger.Log.Println(resulsetErr)
			return errors.New("Unable to fetch resultset")
		}

		countQuery := "SELECT count(id)  as count from mstrecorddifferentiation where clientid=? and mstorgnhirarchyid=? and deleteflg=0 and activeflg=1 and recorddifftypeid=? and name=?"
		resulsetErr1 := db.QueryRow(countQuery, tz.ClientID, tz.OrgID, recordDiffTypeID, tz.DelimiterVal).Scan(&count)
		if resulsetErr1 != nil {
			Logger.Log.Println(resulsetErr1)
			return errors.New("Unable to fetch resultset")
		}

		if count == 0 {
			var seqno int64
			getSeqnoQuery := "SELECT seqno  as seqno from mstrecorddifferentiation where clientid=? and mstorgnhirarchyid=? and deleteflg=0 and activeflg=1 and recorddifftypeid=?  order by seqno desc"
			resulsetErr1 := db.QueryRow(getSeqnoQuery, tz.ClientID, tz.OrgID, recordDiffTypeID).Scan(&seqno)
			if resulsetErr1 != nil {
				Logger.Log.Println(resulsetErr1)
				// return errors.New("Unable to fetch resultset")
			}

			seqno = seqno + 1
			Logger.Log.Print("seqno", seqno)
			sql1 := "INSERT INTO mstrecorddifferentiation(`clientid`,`mstorgnhirarchyid`,`recorddifftypeid`,`parentid`,`name`,`seqno`,`deleteflg`,`activeflg`)VALUES(?,?,?,0,?,?,0,1)"
			statementTX1, stmnterr1 := tx.Prepare(sql1)
			if stmnterr1 != nil {

				Logger.Log.Println(stmnterr1)
				return errors.New("Unable to Prepare resultset")
			}
			defer statementTX1.Close()
			resultSet1, execError1 := statementTX1.Exec(tz.ClientID, tz.OrgID, recordDiffTypeID, tz.DelimiterVal, seqno)
			if execError1 != nil {

				Logger.Log.Println(execError1)
				return errors.New("Unable to Exec Query")
			}
			_, err := resultSet1.RowsAffected()
			if err != nil {

				Logger.Log.Println(err)
				return errors.New("No rows affected")
			}
		} else {
			return errors.New("Data already exist")
		}

	}

	commiterr := tx.Commit()
	if commiterr != nil {
		tx.Rollback()
		Logger.Log.Println(commiterr)
		return errors.New("unable to commit")
	}

	return nil
}

func GetDelimiter(db *sql.DB, clientID int64, orgID int64) ([]string, error) {
	var delimeterList []string
	getDelimetreQuery := "select name from mstrecorddifferentiation a where a.clientid=? and  a.recorddifftypeid in(select id from mstrecorddifferentiationtype where clientid=a.clientid and seqno=201 and parentid=11 and deleteflg = 0 and activeflg = 1) and a.deleteflg = 0 and a.activeflg = 1"
	resutlset, resulsetErr := db.Query(getDelimetreQuery, clientID)
	if resulsetErr != nil {
		Logger.Log.Println(resulsetErr)
		return delimeterList, errors.New("Delemetre is not Present")
	}
	defer resutlset.Close()
	for resutlset.Next() {
		var delimeter string
		scanErr := resutlset.Scan(&delimeter)
		if scanErr != nil {
			Logger.Log.Println(scanErr)
			return delimeterList, errors.New("Service User Scan Error")
		}
		delimeterList = append(delimeterList, delimeter)
	}
	return delimeterList, nil
}

func GetSeviceUserList(db *sql.DB, clientID int64, orgID int64) (entities.ServiceUserEntities, error) {
	var usersList entities.ServiceUserEntities
	// getServiceUserQuery := "select m.id, m.loginname, m.name , (select groupid from mstgroupmember where userid=m.id and clientid= m.clientid and mstorgnhirarchyid=m.mstorgnhirarchyid and groupid in (select grpid from mstclientsupportgroup where clientid=m.clientid and mstorgnhirarchyid=m.mstorgnhirarchyid and supportgrouplevelid=1 and activeflag=1 and deleteflg=0)) serviceusergroupid from mstclientuser m where m.clientid=? and m.mstorgnhirarchyid=? and m.usertype='Service' and m.activeflag=1 and m.deleteflag=0"
	getServiceUserQuery := "select m.id, m.loginname, m.name,n.groupid from mstclientuser m,mstgroupmember n where n.userid=m.id and n.clientid= m.clientid and n.mstorgnhirarchyid=m.mstorgnhirarchyid and n.groupid in (select grpid from mstclientsupportgroup where clientid=m.clientid and mstorgnhirarchyid=m.mstorgnhirarchyid and supportgrouplevelid=1 and activeflag=1 and deleteflg=0) and m.clientid=? and m.mstorgnhirarchyid=? and m.usertype='Service' and m.activeflag=1 and m.deleteflag=0 and n.activeflg=1 and n.deleteflg=0"
	//The above getServiceUserQuery commented by me josim bcs of its error "No service User Found"
	serviceUserResultset, resulsetErr := db.Query(getServiceUserQuery, clientID, orgID)
	if resulsetErr != nil {
		Logger.Log.Println(resulsetErr)
		return usersList, errors.New("No service User Found")
	}
	defer serviceUserResultset.Close()
	for serviceUserResultset.Next() {
		var user entities.ServiceUserEntity
		scanErr := serviceUserResultset.Scan(&user.ID, &user.LoginName, &user.Name, &user.GroupID)
		if scanErr != nil {
			Logger.Log.Println(scanErr)
			return usersList, errors.New("Service User Scan Error")
		}

		usersList.Values = append(usersList.Values, user)
	}

	return usersList, nil
}

func GetCategoryList(db *sql.DB, clientID int64, orgID int64, categoryLavelID int64) (entities.CategoryList, error) {
	categoryList := entities.CategoryList{}
	category := entities.Category{}

	getCategoryLevel5ListQuery := "select id,name,parentcategoryids,parentcategorynames from mstrecorddifferentiation where clientid=? and mstorgnhirarchyid=? and  recorddifftypeid =? and  activeflg=1 and deleteflg=0"
	getCategoryLevel5ListResultSet, resulsetErr := db.Query(getCategoryLevel5ListQuery, clientID, orgID, categoryLavelID)
	if resulsetErr != nil {
		Logger.Log.Println(resulsetErr)
		return categoryList, errors.New("No Categories found Found")
	}
	defer getCategoryLevel5ListResultSet.Close()
	for getCategoryLevel5ListResultSet.Next() {
		var ID int64
		var categoryName string
		var categoryParentIds string
		var categoryParentNames string
		scanErr := getCategoryLevel5ListResultSet.Scan(&ID, &categoryName, &categoryParentIds, &categoryParentNames)

		if scanErr != nil {
			Logger.Log.Println(scanErr)
			return categoryList, errors.New("List categories Scan Error")
		}

		category.ID = ID
		category.CategoryName = categoryName
		category.CategoryParentIds = categoryParentIds
		category.CategoryParentNames = categoryParentNames
		category.CategoryNameWithPath = categoryName + " ( " + categoryParentNames + " )"

		categoryList.Values = append(categoryList.Values, category)

	}

	return categoryList, nil
}

func InsertMstEmailTicketConfig(db *sql.DB, emailTicketObj entities.MstEmailTicket) error {

	getPriorityQuery := "select mstrecorddifferentiationpriorityid from mstbusinessmatrix where clientid=? and mstorgnhirarchyid=? and mstrecorddifferentiationcatid=? and activeflg=1 and deleteflg=0"
	resulsetErr := db.QueryRow(getPriorityQuery, emailTicketObj.ClientID, emailTicketObj.OrgID, emailTicketObj.LastCategoryID).Scan(&emailTicketObj.PriorityID)
	if resulsetErr != nil {
		Logger.Log.Println(resulsetErr)
		return errors.New("Unable to fetch Priority")
	}

	// getSenderTypeIDQuery := "select id from mstrecorddifferentiation where recorddifftypeid in (select id from mstrecorddifferentiationtype where clientid = ? and mstorgnhirarchyid = ? and activeflg=1 and deleteflg=0 and seqno=202 and parentid in (select id from mstrecorddifferentiationtype where seqno=11)) and seqno=?"
	// getSenderTypeIDResulsetErr := db.QueryRow(getSenderTypeIDQuery, emailTicketObj.ClientID, emailTicketObj.OrgID, emailTicketObj.SenderTypeSeq).Scan(&emailTicketObj.SenderTypeID)
	// if getSenderTypeIDResulsetErr != nil {
	// 	Logger.Log.Println(getSenderTypeIDResulsetErr)
	// 	return errors.New("Unable to SenderTypeID")
	// }

	var rowCount int64

	if emailTicketObj.SenderTypeSeq == 1 {
		if emailTicketObj.DefaultSeq == 0 {
			checkForDuplicateEntryQuery := "select count(*) as count from mstemailticket where clientid = ? and mstorgnhirarchyid=? and lastcategoryid=? and emailsubkeyword=? and senderemail=? and defaultseq=? and activeflg=1 and deleteflg=0 "
			resulsetErr := db.QueryRow(checkForDuplicateEntryQuery, emailTicketObj.ClientID, emailTicketObj.OrgID, emailTicketObj.LastCategoryID, emailTicketObj.EmailSubKeyword, emailTicketObj.SenderEmail, emailTicketObj.DefaultSeq).Scan(&rowCount)
			if resulsetErr != nil {
				Logger.Log.Println(resulsetErr)
				//return errors.New("Unable to fetch Priority")
			}
			if rowCount > 0 {
				Logger.Log.Println("rowCount===>", rowCount)
				return errors.New("Duplicate Entry")
			}
			log.Println("default=0,sendertype=1")
			insertMstEmailTicketQuery := "INSERT INTO `mstemailticket`(`clientid`,`mstorgnhirarchyid`,`mstrecorddifftypeid`,`mstrecorddiffid`,`categorydifftypeid`,`categorylevelid`,`lastcategoryid`,`lastcategoryname`,`categoryidlist`,`categorynamelist`,`categorywithpath`,`serviceuserid`,`serviceusergroupid`,`senderemail`,`emailsubkeyword`,`priorityid`,`sendertypeseq`,`delimiter`,`createdbyid`,`createddate`)VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,round(UNIX_TIMESTAMP(now())))"
			Resultset, insertError := db.Query(insertMstEmailTicketQuery, emailTicketObj.ClientID, emailTicketObj.OrgID, emailTicketObj.TicketDiffTypeID, emailTicketObj.TicketDiffID, emailTicketObj.CategoryDiffTypeID, emailTicketObj.CategoryLevelID, emailTicketObj.LastCategoryID, emailTicketObj.LastCategoryName, emailTicketObj.CategoryIDList, emailTicketObj.CategoryNameList, emailTicketObj.CategoryWithPath, emailTicketObj.ServiceUserID, emailTicketObj.ServiceUserGroupID, emailTicketObj.SenderEmail, emailTicketObj.EmailSubKeyword, emailTicketObj.PriorityID, emailTicketObj.SenderTypeSeq, emailTicketObj.Delimiter, emailTicketObj.CreatedByID)
			if insertError != nil {
				Logger.Log.Println(insertError)
				return errors.New("unable to insert")
			}
			Resultset.Close()

		} else {
			//log.Println("default=1,sendertype=1")
			checkForDuplicateEntryQuery := "select count(*) as count from mstemailticket where clientid = ? and mstorgnhirarchyid=? and lastcategoryid=? and senderemail=? and defaultseq=? and activeflg=1 and deleteflg=0"
			resulsetErr := db.QueryRow(checkForDuplicateEntryQuery, emailTicketObj.ClientID, emailTicketObj.OrgID, emailTicketObj.LastCategoryID, emailTicketObj.SenderEmail, emailTicketObj.DefaultSeq).Scan(&rowCount)
			if resulsetErr != nil {
				Logger.Log.Println(resulsetErr)
				//return errors.New("Unable to fetch Priority")
			}
			if rowCount > 0 {
				Logger.Log.Println("rowCount===>", rowCount)
				return errors.New("Duplicate Entry")
			}
			log.Println("default=1,sendertype=1")
			insertMstEmailTicketQuery := "INSERT INTO `mstemailticket`(`clientid`,`mstorgnhirarchyid`,`mstrecorddifftypeid`,`mstrecorddiffid`,`categorydifftypeid`,`categorylevelid`,`lastcategoryid`,`lastcategoryname`,`categoryidlist`,`categorynamelist`,`categorywithpath`,`serviceuserid`,`serviceusergroupid`,`senderemail`,`priorityid`,`sendertypeseq`,`createdbyid`,`createddate`,`defaultseq`)VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,round(UNIX_TIMESTAMP(now())),?)"
			Resultset, insertError := db.Query(insertMstEmailTicketQuery, emailTicketObj.ClientID, emailTicketObj.OrgID, emailTicketObj.TicketDiffTypeID, emailTicketObj.TicketDiffID, emailTicketObj.CategoryDiffTypeID, emailTicketObj.CategoryLevelID, emailTicketObj.LastCategoryID, emailTicketObj.LastCategoryName, emailTicketObj.CategoryIDList, emailTicketObj.CategoryNameList, emailTicketObj.CategoryWithPath, emailTicketObj.ServiceUserID, emailTicketObj.ServiceUserGroupID, emailTicketObj.SenderEmail, emailTicketObj.PriorityID, emailTicketObj.SenderTypeSeq, emailTicketObj.CreatedByID, emailTicketObj.DefaultSeq)
			if insertError != nil {
				Logger.Log.Println(insertError)
				return errors.New("unable to insert")
			}
			Resultset.Close()

		}

	} else if emailTicketObj.SenderTypeSeq == 2 {
		if emailTicketObj.DefaultSeq == 0 {
			checkForDuplicateEntryQuery := "select count(*) as count from mstemailticket where clientid = ? and mstorgnhirarchyid=? and lastcategoryid=? and emailsubkeyword=? and senderdomain=? and defaultseq=? and activeflg=1 and deleteflg=0"
			resulsetErr := db.QueryRow(checkForDuplicateEntryQuery, emailTicketObj.ClientID, emailTicketObj.OrgID, emailTicketObj.LastCategoryID, emailTicketObj.EmailSubKeyword, emailTicketObj.SenderDomain, emailTicketObj.DefaultSeq).Scan(&rowCount)
			if resulsetErr != nil {
				Logger.Log.Println(resulsetErr)
				return errors.New("Unable to fetch data")
			}
			if rowCount > 0 {
				Logger.Log.Println("rowCount===>", rowCount)
				return errors.New("Duplicate Entry")
			}
			log.Println("default=0,sendertype=2")
			insertMstEmailTicketQuery := "INSERT INTO `mstemailticket`(`clientid`,`mstorgnhirarchyid`,`mstrecorddifftypeid`,`mstrecorddiffid`,`categorydifftypeid`,`categorylevelid`,`lastcategoryid`,`lastcategoryname`,`categoryidlist`,`categorynamelist`,`categorywithpath`,`serviceuserid`,`serviceusergroupid`,`senderdomain`,`emailsubkeyword`,`priorityid`,`sendertypeseq`,`delimiter`,`createdbyid`,`createddate`)VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,round(UNIX_TIMESTAMP(now())))"
			Resultset, insertError := db.Query(insertMstEmailTicketQuery, emailTicketObj.ClientID, emailTicketObj.OrgID, emailTicketObj.TicketDiffTypeID, emailTicketObj.TicketDiffID, emailTicketObj.CategoryDiffTypeID, emailTicketObj.CategoryLevelID, emailTicketObj.LastCategoryID, emailTicketObj.LastCategoryName, emailTicketObj.CategoryIDList, emailTicketObj.CategoryNameList, emailTicketObj.CategoryWithPath, emailTicketObj.ServiceUserID, emailTicketObj.ServiceUserGroupID, emailTicketObj.SenderDomain, emailTicketObj.EmailSubKeyword, emailTicketObj.PriorityID, emailTicketObj.SenderTypeSeq, emailTicketObj.Delimiter, emailTicketObj.CreatedByID)
			if insertError != nil {
				Logger.Log.Println(insertError)
				return errors.New("unable to insert")
			}
			Resultset.Close()

		} else {
			checkForDuplicateEntryQuery := "select count(*) as count from mstemailticket where clientid = ? and mstorgnhirarchyid=? and lastcategoryid=? and senderdomain=? and defaultseq = ? and activeflg=1 and deleteflg=0"
			resulsetErr := db.QueryRow(checkForDuplicateEntryQuery, emailTicketObj.ClientID, emailTicketObj.OrgID, emailTicketObj.LastCategoryID, emailTicketObj.SenderDomain, emailTicketObj.DefaultSeq).Scan(&rowCount)
			if resulsetErr != nil {
				Logger.Log.Println(resulsetErr)
				return errors.New("Unable to fetch Priority")
			}
			if rowCount > 0 {
				Logger.Log.Println("rowCount===>", rowCount)
				return errors.New("Duplicate Entry")
			}
			log.Println("default=1,sendertype=2")
			insertMstEmailTicketQuery := "INSERT INTO `mstemailticket`(`clientid`,`mstorgnhirarchyid`,`mstrecorddifftypeid`,`mstrecorddiffid`,`categorydifftypeid`,`categorylevelid`,`lastcategoryid`,`lastcategoryname`,`categoryidlist`,`categorynamelist`,`categorywithpath`,`serviceuserid`,`serviceusergroupid`,`senderdomain`,`priorityid`,`sendertypeseq`,`createdbyid`,`createddate`,`defaultseq`)VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,round(UNIX_TIMESTAMP(now())),?)"
			Resultset, insertError := db.Query(insertMstEmailTicketQuery, emailTicketObj.ClientID, emailTicketObj.OrgID, emailTicketObj.TicketDiffTypeID, emailTicketObj.TicketDiffID, emailTicketObj.CategoryDiffTypeID, emailTicketObj.CategoryLevelID, emailTicketObj.LastCategoryID, emailTicketObj.LastCategoryName, emailTicketObj.CategoryIDList, emailTicketObj.CategoryNameList, emailTicketObj.CategoryWithPath, emailTicketObj.ServiceUserID, emailTicketObj.ServiceUserGroupID, emailTicketObj.SenderDomain, emailTicketObj.PriorityID, emailTicketObj.SenderTypeSeq, emailTicketObj.CreatedByID, emailTicketObj.DefaultSeq)
			if insertError != nil {
				Logger.Log.Println(insertError)
				return errors.New("unable to insert")
			}
			Resultset.Close()
		}

	}
	return nil
}
func UpdateMstEmailTicket(db *sql.DB, emailTicketObj entities.MstEmailTicket) error {

	getPriorityQuery := "select mstrecorddifferentiationpriorityid from mstbusinessmatrix where clientid=? and mstorgnhirarchyid=? and mstrecorddifferentiationcatid=? and activeflg=1 and deleteflg=0"
	resulsetErr := db.QueryRow(getPriorityQuery, emailTicketObj.ClientID, emailTicketObj.OrgID, emailTicketObj.LastCategoryID).Scan(&emailTicketObj.PriorityID)
	if resulsetErr != nil {
		Logger.Log.Println(resulsetErr)
		return errors.New("Unable to fetch Priority")
	}

	// getSenderTypeIDQuery := "select id from mstrecorddifferentiation where recorddifftypeid in (select id from mstrecorddifferentiationtype where clientid = ? and mstorgnhirarchyid = ? and activeflg=1 and deleteflg=0 and seqno=202 and parentid in (select id from mstrecorddifferentiationtype where seqno=11)) and seqno=?"
	// getSenderTypeIDResulsetErr := db.QueryRow(getSenderTypeIDQuery, emailTicketObj.ClientID, emailTicketObj.OrgID, emailTicketObj.SenderTypeSeq).Scan(&emailTicketObj.SenderTypeID)
	// if getSenderTypeIDResulsetErr != nil {
	// 	Logger.Log.Println(getSenderTypeIDResulsetErr)
	// 	return errors.New("Unable to SenderTypeID")
	// }

	var rowCount int64

	if emailTicketObj.SenderTypeSeq == 1 {
		if emailTicketObj.DefaultSeq == 0 {
			checkForDuplicateEntryQuery := "select count(*) as count from mstemailticket where clientid = ? and mstorgnhirarchyid=? and lastcategoryid=? and emailsubkeyword=? and senderemail=? and defaultseq=? and activeflg=1 and deleteflg=0"
			resulsetErr := db.QueryRow(checkForDuplicateEntryQuery, emailTicketObj.ClientID, emailTicketObj.OrgID, emailTicketObj.LastCategoryID, emailTicketObj.EmailSubKeyword, emailTicketObj.SenderEmail, emailTicketObj.DefaultSeq).Scan(&rowCount)
			if resulsetErr != nil {
				Logger.Log.Println(resulsetErr)
				return errors.New("Unable to fetch Priority")
			}
			if rowCount > 0 {
				Logger.Log.Println("rowCount===>", rowCount)
				return errors.New("Duplicate Entry")
			}
			log.Println("default=0,sendertype=1")
			updateMstEmailTicketQuery := "UPDATE mstemailticket SET clientid = ?, mstorgnhirarchyid = ?,mstrecorddifftypeid = ?, mstrecorddiffid = ?,categorydifftypeid = ?, " +
				"categorylevelid = ?, lastcategoryid = ?, lastcategoryname = ?, categoryidlist = ?, categorynamelist = ?, categorywithpath = ?, serviceuserid = ?," +
				"serviceusergroupid = ?,senderemail = ?,emailsubkeyword = ?,priorityid = ?,sendertypeseq = ?,delimiter = ?,createdbyid = ?, defaultseq = ? WHERE id = ?"
			Resultset, insertError := db.Query(updateMstEmailTicketQuery, emailTicketObj.ClientID, emailTicketObj.OrgID, emailTicketObj.TicketDiffTypeID, emailTicketObj.TicketDiffID, emailTicketObj.CategoryDiffTypeID, emailTicketObj.CategoryLevelID, emailTicketObj.LastCategoryID, emailTicketObj.LastCategoryName, emailTicketObj.CategoryIDList, emailTicketObj.CategoryNameList, emailTicketObj.CategoryWithPath, emailTicketObj.ServiceUserID, emailTicketObj.ServiceUserGroupID, emailTicketObj.SenderEmail, emailTicketObj.EmailSubKeyword, emailTicketObj.PriorityID, emailTicketObj.SenderTypeSeq, emailTicketObj.Delimiter, emailTicketObj.CreatedByID, emailTicketObj.DefaultSeq, emailTicketObj.ID)
			if insertError != nil {
				Logger.Log.Println(insertError)
				return errors.New("unable to update")
			}
			Resultset.Close()
		} else {
			//log.Println("default=1,sendertype=1")
			checkForDuplicateEntryQuery := "select count(*) as count from mstemailticket where clientid = ? and mstorgnhirarchyid=? and lastcategoryid=? and senderemail=? and defaultseq=? and  activeflg=1 and deleteflg=0"
			resulsetErr := db.QueryRow(checkForDuplicateEntryQuery, emailTicketObj.ClientID, emailTicketObj.OrgID, emailTicketObj.LastCategoryID, emailTicketObj.SenderEmail, emailTicketObj.DefaultSeq).Scan(&rowCount)
			if resulsetErr != nil {
				Logger.Log.Println(resulsetErr)
				return errors.New("Unable to fetch Priority")
			}
			if rowCount > 0 {
				Logger.Log.Println("rowCount===>", rowCount)
				return errors.New("Duplicate Entry")
			}
			updateMstEmailTicketQuery := "UPDATE mstemailticket SET clientid = ?, mstorgnhirarchyid = ?,mstrecorddifftypeid = ?, mstrecorddiffid = ?,categorydifftypeid = ?, " +
				"categorylevelid = ?, lastcategoryid = ?, lastcategoryname = ?, categoryidlist = ?, categorynamelist = ?, categorywithpath = ?, serviceuserid = ?," +
				"serviceusergroupid = ?,senderemail = ?,priorityid = ?,sendertypeseq = ?,createdbyid = ?, defaultseq = ? , emailsubkeyword = ?,delimiter = ? WHERE id = ?"
			Resultset, insertError := db.Query(updateMstEmailTicketQuery, emailTicketObj.ClientID, emailTicketObj.OrgID, emailTicketObj.TicketDiffTypeID, emailTicketObj.TicketDiffID, emailTicketObj.CategoryDiffTypeID, emailTicketObj.CategoryLevelID, emailTicketObj.LastCategoryID, emailTicketObj.LastCategoryName, emailTicketObj.CategoryIDList, emailTicketObj.CategoryNameList, emailTicketObj.CategoryWithPath, emailTicketObj.ServiceUserID, emailTicketObj.ServiceUserGroupID, emailTicketObj.SenderEmail, emailTicketObj.PriorityID, emailTicketObj.SenderTypeSeq, emailTicketObj.CreatedByID, emailTicketObj.DefaultSeq, emailTicketObj.EmailSubKeyword, emailTicketObj.Delimiter, emailTicketObj.ID)
			if insertError != nil {
				Logger.Log.Println(insertError)
				return errors.New("unable to update")
			}
			Resultset.Close()
		}

	} else if emailTicketObj.SenderTypeSeq == 2 {
		if emailTicketObj.DefaultSeq == 0 {
			checkForDuplicateEntryQuery := "select count(*) as count from mstemailticket where clientid = ? and mstorgnhirarchyid=? and lastcategoryid=? and emailsubkeyword=? and senderdomain=? and defaultseq=? and activeflg=1 and deleteflg=0"
			resulsetErr := db.QueryRow(checkForDuplicateEntryQuery, emailTicketObj.ClientID, emailTicketObj.OrgID, emailTicketObj.LastCategoryID, emailTicketObj.EmailSubKeyword, emailTicketObj.SenderDomain, emailTicketObj.DefaultSeq).Scan(&rowCount)
			if resulsetErr != nil {
				Logger.Log.Println(resulsetErr)
				return errors.New("Unable to fetch Priority")
			}
			if rowCount > 0 {
				Logger.Log.Println("rowCount===>", rowCount)
				return errors.New("Duplicate Entry")
			}
			log.Println("default=0,sendertype=2")
			updateMstEmailTicketQuery := "UPDATE mstemailticket SET clientid = ?, mstorgnhirarchyid = ?,mstrecorddifftypeid = ?, mstrecorddiffid = ?,categorydifftypeid = ?, " +
				"categorylevelid = ?, lastcategoryid = ?, lastcategoryname = ?, categoryidlist = ?, categorynamelist = ?, categorywithpath = ?, serviceuserid = ?," +
				"serviceusergroupid = ?,senderdomain = ?,emailsubkeyword = ?,priorityid = ?,sendertypeseq = ?,delimiter = ?,createdbyid = ?, defaultseq = ? WHERE id = ?"
			Resultset, insertError := db.Query(updateMstEmailTicketQuery, emailTicketObj.ClientID, emailTicketObj.OrgID, emailTicketObj.TicketDiffTypeID, emailTicketObj.TicketDiffID, emailTicketObj.CategoryDiffTypeID, emailTicketObj.CategoryLevelID, emailTicketObj.LastCategoryID, emailTicketObj.LastCategoryName, emailTicketObj.CategoryIDList, emailTicketObj.CategoryNameList, emailTicketObj.CategoryWithPath, emailTicketObj.ServiceUserID, emailTicketObj.ServiceUserGroupID, emailTicketObj.SenderDomain, emailTicketObj.EmailSubKeyword, emailTicketObj.PriorityID, emailTicketObj.SenderTypeSeq, emailTicketObj.Delimiter, emailTicketObj.CreatedByID, emailTicketObj.DefaultSeq, emailTicketObj.ID)
			if insertError != nil {
				Logger.Log.Println(insertError)
				return errors.New("unable to update")
			}
			Resultset.Close()
		} else {
			checkForDuplicateEntryQuery := "select count(*) as count from mstemailticket where clientid = ? and mstorgnhirarchyid=? and lastcategoryid=? and senderdomain=? and defaultseq = ? and activeflg=1 and deleteflg=0"
			resulsetErr := db.QueryRow(checkForDuplicateEntryQuery, emailTicketObj.ClientID, emailTicketObj.OrgID, emailTicketObj.LastCategoryID, emailTicketObj.SenderDomain, emailTicketObj.DefaultSeq).Scan(&rowCount)
			if resulsetErr != nil {
				Logger.Log.Println(resulsetErr)
				return errors.New("Unable to fetch Priority")
			}
			if rowCount > 0 {
				Logger.Log.Println("rowCount===>", rowCount)
				return errors.New("Duplicate Entry")
			}
			log.Println("default=1,sendertype=2")
			updateMstEmailTicketQuery := "UPDATE mstemailticket SET clientid = ?, mstorgnhirarchyid = ?,mstrecorddifftypeid = ?, mstrecorddiffid = ?,categorydifftypeid = ?, " +
				"categorylevelid = ?, lastcategoryid = ?, lastcategoryname = ?, categoryidlist = ?, categorynamelist = ?, categorywithpath = ?, serviceuserid = ?," +
				"serviceusergroupid = ?,senderdomain = ?,priorityid = ?,sendertypeseq = ?,createdbyid = ?, defaultseq = ? , emailsubkeyword = ?,delimiter = ? WHERE id = ?"
			Resultset, insertError := db.Query(updateMstEmailTicketQuery, emailTicketObj.ClientID, emailTicketObj.OrgID, emailTicketObj.TicketDiffTypeID, emailTicketObj.TicketDiffID, emailTicketObj.CategoryDiffTypeID, emailTicketObj.CategoryLevelID, emailTicketObj.LastCategoryID, emailTicketObj.LastCategoryName, emailTicketObj.CategoryIDList, emailTicketObj.CategoryNameList, emailTicketObj.CategoryWithPath, emailTicketObj.ServiceUserID, emailTicketObj.ServiceUserGroupID, emailTicketObj.SenderDomain, emailTicketObj.PriorityID, emailTicketObj.SenderTypeSeq, emailTicketObj.CreatedByID, emailTicketObj.DefaultSeq, emailTicketObj.EmailSubKeyword, emailTicketObj.Delimiter, emailTicketObj.ID)
			if insertError != nil {
				Logger.Log.Println(insertError)
				return errors.New("unable to update")
			}
			Resultset.Close()
		}

	}

	return nil
}

func GetEmailTicketConfigurations(db *sql.DB, clientID int64, orgID int64, limit int64, offset int64, orgntype int64) (entities.MstEmailTicketConfigtList, error) {
	var emailTicketViewList entities.MstEmailTicketConfigtList

	//var serviceUsers map[int64]string
	var getEmailTicketViewListQuery string
	var params []interface{}
	if orgntype == 2 {
		getEmailTicketViewListQuery = "select m.id,m.clientid,m.mstorgnhirarchyid,COALESCE((select name from mstclient where id=m.clientid),'') clientname," +
			" COALESCE((select name from mstorgnhierarchy where id=m.mstorgnhirarchyid) ,'')orgname,COALESCE((SELECT name  FROM mstrecorddifferentiation where id=m.mstrecorddiffid),'') tickettype," +
			" COALESCE(m.categorywithpath,'') categorywithpath,COALESCE(m.emailsubkeyword,'') emailsubkeyword, COALESCE(m.delimiter,''),COALESCE(m.senderemail,''),COALESCE(m.senderdomain,'')," +
			" COALESCE((select name from mstclientuser where id=m.serviceuserid),'')serviceusername,COALESCE((select name from mstclientuser where id=m.createdbyid),'')createdbyname, " +
			" if(m.sendertypeseq=1, 'From Specific Email','From Specific Domain') sendertype,COALESCE(m.mstrecorddifftypeid,'') mstrecorddifftypeid,COALESCE(m.mstrecorddiffid,'') mstrecorddiffid," +
			" COALESCE(m.categorydifftypeid,'') categorydifftypeid,COALESCE(m.categorylevelid,'') categorylevelid, COALESCE(m.lastcategoryid,'') lastcategoryid, COALESCE(m.lastcategoryname,'') lastcategoryname," +
			" COALESCE(m.categoryidlist,'') categoryidlist,COALESCE(m.categorynamelist,'') categorynamelist,COALESCE(m.serviceuserid,'') serviceuserid, COALESCE(m.serviceusergroupid,'') serviceusergroupid ," +
			" COALESCE(m.defaultseq,'') defaultseq,m.sendertypeseq from mstemailticket m where m.clientid=?  and activeflg=1 and deleteflg=0 order by id DESC"
		params = append(params, clientID)
	} else {
		getEmailTicketViewListQuery = "select m.id,m.clientid,m.mstorgnhirarchyid,COALESCE((select name from mstclient where id=m.clientid),'') clientname," +
			" COALESCE((select name from mstorgnhierarchy where id=m.mstorgnhirarchyid) ,'')orgname,COALESCE((SELECT name  FROM mstrecorddifferentiation where id=m.mstrecorddiffid),'') tickettype," +
			" COALESCE(m.categorywithpath,'') categorywithpath,COALESCE(m.emailsubkeyword,'') emailsubkeyword, COALESCE(m.delimiter,''),COALESCE(m.senderemail,''),COALESCE(m.senderdomain,'')," +
			" COALESCE((select name from mstclientuser where id=m.serviceuserid),'')serviceusername,COALESCE((select name from mstclientuser where id=m.createdbyid),'')createdbyname, " +
			" if(m.sendertypeseq=1, 'From Specific Email','From Specific Domain') sendertype,COALESCE(m.mstrecorddifftypeid,'') mstrecorddifftypeid,COALESCE(m.mstrecorddiffid,'') mstrecorddiffid," +
			" COALESCE(m.categorydifftypeid,'') categorydifftypeid,COALESCE(m.categorylevelid,'') categorylevelid, COALESCE(m.lastcategoryid,'') lastcategoryid, COALESCE(m.lastcategoryname,'') lastcategoryname," +
			" COALESCE(m.categoryidlist,'') categoryidlist,COALESCE(m.categorynamelist,'') categorynamelist,COALESCE(m.serviceuserid,'') serviceuserid, COALESCE(m.serviceusergroupid,'') serviceusergroupid ," +
			" COALESCE(m.defaultseq,'') defaultseq,m.sendertypeseq from mstemailticket m where m.clientid=? and m.mstorgnhirarchyid=? and activeflg=1 and deleteflg=0 order by id DESC"
		params = append(params, clientID)
		params = append(params, orgID)
	}
	// params = append(params, limit)
	// params = append(params, offset)
	EmailTicketViewListResultset, resulsetErr := db.Query(getEmailTicketViewListQuery, params...)
	if resulsetErr != nil {
		Logger.Log.Println(resulsetErr)
		return emailTicketViewList, errors.New("No emailTicketView  Found")
	}
	defer EmailTicketViewListResultset.Close()
	for EmailTicketViewListResultset.Next() {
		var emailTicketView entities.MstEmailTicketConfig
		scanErr := EmailTicketViewListResultset.Scan(&emailTicketView.ID, &emailTicketView.ClientID, &emailTicketView.OrgID, &emailTicketView.ClientName,
			&emailTicketView.OrgName, &emailTicketView.TicketTypename, &emailTicketView.CategoryWithPath, &emailTicketView.EmailSubKeyword, &emailTicketView.Delimiter,
			&emailTicketView.SenderEmail, &emailTicketView.SenderDomain, &emailTicketView.ServiceUserName, &emailTicketView.CreatedBynName, &emailTicketView.SenderType, &emailTicketView.TicketDiffTypeID, &emailTicketView.TicketDiffID,
			&emailTicketView.CategoryDiffTypeID, &emailTicketView.CategoryLevelID, &emailTicketView.LastCategoryID, &emailTicketView.LastCategoryName,
			&emailTicketView.CategoryIDList, &emailTicketView.CategoryNameList, &emailTicketView.ServiceUserID, &emailTicketView.ServiceUserGroupID, &emailTicketView.DefaultSeq, &emailTicketView.SenderTypeSeq)
		if scanErr != nil {
			Logger.Log.Println(scanErr)
			return emailTicketViewList, errors.New("emailTicketView Scan Error")
		}

		emailTicketViewList.Values = append(emailTicketViewList.Values, emailTicketView)
	}
	emailTicketViewList.Total = int64(len(emailTicketViewList.Values))
	return emailTicketViewList, nil
}

func GetDelimiterForAllClient(db *sql.DB, clientID int64, orgID int64, limit int64, offset int64) (entities.EmailTicketBaseConfigtList, error) {
	var emailTicketBaseConfigViewList entities.EmailTicketBaseConfigtList

	//var serviceUsers map[int64]string
	var getEmailTicketBaseConfigViewListQuery string = "select b.id as ID,a.clientid as ClientID,COALESCE((select name from" +
		" mstclient where id=a.clientid),'') clientname,a.mstorgnhirarchyid as OrgID,COALESCE((select name from " +
		" mstorgnhierarchy where id=a.mstorgnhirarchyid) ,'')OrgName,a.typename as TypeName,b.name as Name from " +
		" mstrecorddifferentiationtype a, mstrecorddifferentiation b where a.clientid=b.clientid and " +
		" a.mstorgnhirarchyid=b.mstorgnhirarchyid and a.id=b.recorddifftypeid and a.parentid in(select id from " +
		" mstrecorddifferentiationtype where seqno=11) and a.activeflg=1 and b.activeflg=1 and a.deleteflg=0 and " +
		" b.deleteflg=0 order by id DESC limit ? offset ?"
	EmailTicketBaseConfigViewListResultset, resulsetErr := db.Query(getEmailTicketBaseConfigViewListQuery, limit, offset)
	if resulsetErr != nil {
		Logger.Log.Println(resulsetErr)
		return emailTicketBaseConfigViewList, errors.New("No emailTicketBaseConfigViewList  Found")
	}
	defer EmailTicketBaseConfigViewListResultset.Close()
	for EmailTicketBaseConfigViewListResultset.Next() {
		var emailTicketBaseConfigView entities.EmailTicketBaseConfig
		scanErr := EmailTicketBaseConfigViewListResultset.Scan(&emailTicketBaseConfigView.ID, &emailTicketBaseConfigView.ClientID, &emailTicketBaseConfigView.ClientName, &emailTicketBaseConfigView.OrgID,
			&emailTicketBaseConfigView.OrgName, &emailTicketBaseConfigView.TypeName, &emailTicketBaseConfigView.Name)
		if scanErr != nil {
			Logger.Log.Println(scanErr)
			return emailTicketBaseConfigViewList, errors.New("emailTicketBaseConfigViewList Scan Error")
		}

		emailTicketBaseConfigViewList.Values = append(emailTicketBaseConfigViewList.Values, emailTicketBaseConfigView)
	}
	emailTicketBaseConfigViewList.Total = int64(len(emailTicketBaseConfigViewList.Values))
	return emailTicketBaseConfigViewList, nil
}

func DeleteEmailTicketConfiguration(db *sql.DB, rowID int64) error {

	deleteEmailTicketRowQuery := "update mstemailticket set deleteflg = 1 where id=?"
	Resultset, deleteEmailTicketRowResulsetErr := db.Query(deleteEmailTicketRowQuery, rowID)

	if deleteEmailTicketRowResulsetErr != nil {
		Logger.Log.Println(deleteEmailTicketRowResulsetErr)
		return errors.New("Unable to delete")
	}
	Resultset.Close()
	return nil
}

func DeleteEmailTicketConfigu(db *sql.DB, rowID int64) error {

	deleteEmailTicketRowQuery := "update mstrecorddifferentiation set deleteflg = 1 where id=?"
	Resultset, deleteEmailTicketRowResulsetErr := db.Query(deleteEmailTicketRowQuery, rowID)

	if deleteEmailTicketRowResulsetErr != nil {
		Logger.Log.Println(deleteEmailTicketRowResulsetErr)
		return errors.New("Unable to delete")
	}
	Resultset.Close()
	return nil
}
