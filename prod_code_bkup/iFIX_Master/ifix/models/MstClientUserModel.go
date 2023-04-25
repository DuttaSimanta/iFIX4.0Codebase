package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/utility"
	"strconv"
)

//var lock = &sync.Mutex{}
//AddClientUsers for implements business logic
func AddClientUsers(tz *entities.MstClientUserEntity) (int64, bool, error, string) {
	logger.Log.Println("In side model")
	logger.Log.Println("------>",tz)
	//dbcon, err := config.ConnectMySqlDb()
	lock.Lock()
	defer lock.Unlock()
	dbcon, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("Database connection failure.", err)
		return 0, false, err, "Something Went Wrong"
	}

	tx, err := dbcon.Begin()
	if err != nil {
		logger.Log.Println("Transaction creation error.", err)
		return 0, false, err, "Something Went Wrong"
	}
	tz.Password = utility.HashAndSalt([]byte(tz.Password))
	count, err := dao.CheckDuplicateCientUser(tx, tz)
	if err != nil {
		tx.Rollback()
		return 0, false, err, "Something Went Wrong"
	}
	if count == 0 {
		id, err := dao.InsertClientUserData(tx, tz)
		if err != nil {
			tx.Rollback()
			return 0, false, err, "Something Went Wrong"
		}
		if id > 0 {
			count1, err := dao.CheckDuplicateMstUser(tx, tz)
			if err != nil {
				tx.Rollback()
				// dbcon.Close()
				return 0, false, err, "Something Went Wrong"
			}
			if count1 == 0 {
				mstuserid, err := dao.InsertMstUserData(tx, tz, id)
				if err != nil {
					tx.Rollback()
					// dbcon.Close()
					return 0, false, err, "Something Went Wrong"
				}

				if mstuserid > 0 {
					err = tx.Commit()
					if err != nil {
						logger.Log.Println(err)
						tx.Rollback()
						// dbcon.Close()
						return 0, false, err, "Data insertion failure."
					}
					// dbcon.Close()
					return id, true, err, ""
				} else {
					tx.Rollback()
					// dbcon.Close()
					return 0, false, nil, "User Already Exist."
				}
			} else {
				tx.Rollback()
				// dbcon.Close()
				return 0, false, nil, "User Already Exist."
			}
		} else {
			tx.Rollback()
			// dbcon.Close()
			return 0, false, nil, "User Already Exist."
		}

	} else {
		tx.Rollback()
		dbcon.Close()
		return 0, false, nil, "User Already Exist."
	}
}

//GetAllUsers for implements business logic
func GetAllUsers(tz *entities.MstClientUserEntity) (entities.MstClientUserEntities, bool, error, string) {
	logger.Log.Println("In side model")
	t := entities.MstClientUserEntities{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	orgntype, err1 := dataAccess.GetOrgnType(tz.ClientID, tz.MstorgnhirarchyID)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	values, err1 := dataAccess.GetAllUsers(tz, orgntype)
	if err1 != nil {
		logger.Log.Println(err1)
		return t, false, err1, "Something Went Wrong"
	}
	if tz.Offset == 0 {
		total, err1 := dataAccess.GetClientUserCount(tz, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

//DeleteUsers for implements business logic
/*func DeleteUsers(tz *entities.MstClientUserEntity) (bool, error, string) {
	//dbcon, err := config.ConnectMySqlDb()
	lock.Lock()
	defer lock.Unlock()
	dbcon, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("Database connection failure.", err)
		return false, err, "Something Went Wrong"
	}

	tx, err := dbcon.Begin()
	if err != nil {
		tx.Rollback()
		// dbcon.Close()
		logger.Log.Println("Transaction creation error.", err)
		return false, err, "Something Went Wrong"
	}

	err1 := dao.DeleteClientUserData(tx, tz)
	if err1 != nil {
		tx.Rollback()
		// dbcon.Close()
		logger.Log.Println("DeleteUsers table data delete error.", err1)
		return false, err1, "Something Went Wrong"
	}
	err2 := dao.DeleteMstUserData(tx, tz)
	if err2 != nil {
		tx.Rollback()
		// dbcon.Close()
		logger.Log.Println("mstrecordtype table data delete error.", err1)
		return false, err1, "Something Went Wrong"
	}
	err = tx.Commit()
	if err != nil {
		logger.Log.Println(err)
		tx.Rollback()
		// dbcon.Close()
		return false, err, "Data deletion failure."
	}
	RemoveWorkflowUsers(tz,dbcon)
	// dbcon.Close()
	return true, nil, ""

}*/
func Updateusercolor(tz *entities.MstClientUserEntity) (bool, error, string) {
	logger.Log.Println("In side Catalogmodel")
	//db, err := config.ConnectMySqlDb()
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	//defer db.Close()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.Updateusercolor(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

//UpdateUsers for implements business logic
func UpdateUsers(tz *entities.MstClientUserEntity) (bool, error, string) {
	logger.Log.Println("In side model")

	//dbcon, err := config.ConnectMySqlDb()
	lock.Lock()
	defer lock.Unlock()
	dbcon, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	tx, err := dbcon.Begin()
	if err != nil {
		// dbcon.Close()
		logger.Log.Println("Transaction creation error.", err)
		return false, err, "Something Went Wrong"
	}
	count, err := dao.CheckUpdateDuplicateCientUser(tx, tz)
	if err != nil {
		logger.Log.Println(err)
		tx.Rollback()
		dbcon.Close()
		return false, err, "Data updation failure."
	}
	if count == 0 {
		err = dao.UpdateClientUserData(tx, tz)
		if err != nil {
			logger.Log.Println(err)
			tx.Rollback()
			// dbcon.Close()
			return false, err, "Data updation failure."
		}
	}
	//count1, err1 := dao.CheckUpdateDuplicateMstUser(tx, tz)
	//if err1 != nil {
	//	logger.Log.Println(err)
	//	tx.Rollback()
	//	dbcon.Close()
	//	return false, err, "Data updation failure."
	//}
	//if count1 == 0 {
	err = dao.UpdateMstUserData(tx, tz)
	if err != nil {
		logger.Log.Println(err)
		tx.Rollback()
		// dbcon.Close()
		return false, err, "Data updation failure."
	}
	tx.Commit()
	return true, err, ""
	//} else {
	//	return false, nil, "Data Already Exist."
	//}

}

//SearchUser for implements business logic
func SearchUser(tz *entities.MstClientUserEntity) ([]entities.MstUserSearchEntity, bool, error, string) {
	logger.Log.Println("In side model")
	t := []entities.MstUserSearchEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.SearchUser(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}

//SearchUserByOrgnId for implements business logic
func SearchUserByOrgnId(tz *entities.MstClientUserEntity) ([]entities.MstUserSearchEntity, bool, error, string) {
	logger.Log.Println("In side model")
	t := []entities.MstUserSearchEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.SearchUserByOrgnId(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}
func Searchuserbyclientid(tz *entities.MstClientUserEntity) ([]entities.MstUserSearchEntity, bool, error, string) {
	logger.Log.Println("In side model")
	t := []entities.MstUserSearchEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Searchuserbyclientid(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}

//Recordwiseuserinfo for implements business logic
func Recordwiseuserinfo(tz *entities.MstGetUserByRecordidEntity) ([]entities.MstClientUserEntity, bool, error, string) {
	logger.Log.Println("In side model")
	t := []entities.MstClientUserEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Recordwiseuserinfo(tz)
	if err1 != nil {
		logger.Log.Println(err1)
		return t, false, err1, "Something Went Wrong"
	}

	return values, true, err, ""
}

//IDwiseuserinfo for implements business logic
func IDwiseuserinfo(tz *entities.MstGetUserByRecordidEntity) ([]entities.MstClientUserEntity, bool, error, string) {
	logger.Log.Println("In side IDwiseuserinfo model")
	t := []entities.MstClientUserEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.IDwiseuserinfo(tz)
	if err1 != nil {
		logger.Log.Println(err1)
		return t, false, err1, "Something Went Wrong"
	}

	return values, true, err, ""
}

func SearchLoginName(tz *entities.MstClientUserEntity) ([]entities.LoginNameSearchEntity, bool, error, string) {
	logger.Log.Println("In side model")
	t := []entities.LoginNameSearchEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.SearchLoginName(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}
func SearchName(tz *entities.MstClientUserEntity) ([]entities.NameSearchEntity, bool, error, string) {
	logger.Log.Println("In side model")
	t := []entities.NameSearchEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.SearchName(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}
func SearchBranch(tz *entities.MstClientUserEntity) ([]entities.BranchSearchEntity, bool, error, string) {
	logger.Log.Println("In side model")
	t := []entities.BranchSearchEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.SearchBranch(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}
func SearchLoginamebyGroupids(tz *entities.MstClientUserEntity) ([]entities.LoginnameAndNameEntity, bool, error, string) {
	logger.Log.Println("In side model")
	t := []entities.LoginnameAndNameEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}

	var ids string = ""
	for i, groupid := range tz.Groupids {
		if i > 0 {
			ids += ","
		}
		ids += strconv.Itoa(int(groupid))
	}

	values, err1 := dataAccess.SearchLoginamebyGroupids(tz, ids)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}

//func UpdateuserDefaultgrp(tz *entities.MstClientUserEntity) (bool, error, string) {
//	logger.Log.Println("In side UpdateuserDefaultgrpmodel")
//	lock.Lock()
//	defer lock.Unlock()
//	db, err := config.ConnectMySqlDbSingleton()
//	if err != nil {
//		logger.Log.Println("database connection failure", err)
//		return false, err, "Something Went Wrong"
//	}
//	dataAccess := dao.DbConn{DB: db}
//	err1 := dataAccess.UpdateuserDefaultgrp(tz)
//	if err1 != nil {
//		return false, err1, "Something Went Wrong"
//	}
//	return true, nil, ""
//}
func DeleteUsers(tz *entities.MstClientUserEntity) (bool, error, string) {
	//dbcon, err := config.ConnectMySqlDb()
	lock.Lock()
	defer lock.Unlock()
	dbcon, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("Database connection failure.", err)
		return false, err, "Something Went Wrong"
	}

	tx, err := dbcon.Begin()
	if err != nil {
		tx.Rollback()
		// dbcon.Close()
		logger.Log.Println("Transaction creation error.", err)
		return false, err, "Something Went Wrong"
	}

	err1 := dao.DeleteClientUserData(tx, tz)
	if err1 != nil {
		tx.Rollback()
		// dbcon.Close()
		logger.Log.Println("DeleteUsers table data delete error.", err1)
		return false, err1, "Something Went Wrong"
	}
	err2 := dao.DeleteMstUserData(tx, tz)
	if err2 != nil {
		tx.Rollback()
		// dbcon.Close()
		logger.Log.Println("mstrecordtype table data delete error.", err1)
		return false, err1, "Something Went Wrong"
	}

	// Here Start New Code 04.11.2022...
	err3 := dao.UpdateTrnreordDeleteUserID(tx, tz)
	if err3 != nil {
		tx.Rollback()
		// dbcon.Close()
		logger.Log.Println("Trnrecord table data delete error.", err1)
		return false, err1, "Something Went Wrong"
	}
	err4 := dao.UpdateTrnreordDeleteOriginalUserID(tx, tz)
	if err4 != nil {
		tx.Rollback()
		// dbcon.Close()
		logger.Log.Println("Trnrecord table data delete error.", err1)
		return false, err1, "Something Went Wrong"
	}
	err5 := dao.UpdateRecordfulldetailsDeleteRequesterid(tx, tz)
	if err5 != nil {
		tx.Rollback()
		// dbcon.Close()
		logger.Log.Println("Recordfulldetails table data delete error.", err1)
		return false, err1, "Something Went Wrong"
	}
	err6 := dao.UpdateRecordfulldetailsDeleteOriginalid(tx, tz)
	if err6 != nil {
		tx.Rollback()
		// dbcon.Close()
		logger.Log.Println("Recordfulldetails table data delete error.", err1)
		return false, err1, "Something Went Wrong"
	}
	// End Here...
	err = tx.Commit()
	if err != nil {
		logger.Log.Println(err)
		tx.Rollback()
		// dbcon.Close()
		return false, err, "Data deletion failure."
	}
	// dbcon.Close()
	return true, nil, ""
}
