package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func InsertUserWithGroupAndCategory(tz *entities.UserWithGroupAndCategoryEntity) (int64, bool, error, string) {
	logger.Log.Println("In side UserWithGroupAndCategorymodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	var response int64
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return response, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	var errCounter int64
	// var updatedID int64
	for _, v := range tz.WorkingCategories {
		tz.Categoryid = v
		count, err := dataAccess.CheckDuplicateUserWithGroupAndCategory(tz)
		if err != nil {
			return response, false, err, "Something Went Wrong"
		}
		errCounter = errCounter + count.Total
		// updatedID = count.Total
	}
	// logger.Log.Println("\n\n\n  updatedID ====    ", updatedID)
	if errCounter == 0 {
		for _, v := range tz.WorkingCategories {
			tz.Categoryid = v
			id, err := dataAccess.InsertUserWithGroupAndCategory(tz)
			response = id
			if err != nil {
				return response, false, err, "Something Went Wrong"
			}
		}
		return response, true, err, ""
	} else {
		return response, false, nil, "Data Already Exist."
	}

}

func GetAllUserWithGroupAndCategory(page *entities.UserWithGroupAndCategoryEntity) (entities.UserWithGroupAndCategoryEntities, bool, error, string) {
	logger.Log.Println("In side UserWithGroupAndCategorymodel")
	t := entities.UserWithGroupAndCategoryEntities{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	orgntype, err1 := dataAccess.GetOrgnType(page.Clientid, page.Mstorgnhirarchyid)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	values, err1 := dataAccess.GetAllUserWithGroupAndCategory(page, orgntype)
	logger.Log.Println("\n Values ==  ", values)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetUserWithGroupAndCategoryCount(page, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values

	return t, true, err, ""
}

func DeleteUserWithGroupAndCategory(tz *entities.UserWithGroupAndCategoryEntity) (bool, error, string) {
	logger.Log.Println("In side UserWithGroupAndCategorymodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteUserWithGroupAndCategory(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

func UpdateUserWithGroupAndCategory(tz *entities.UserWithGroupAndCategoryEntity) (bool, error, string) {
	logger.Log.Println("In side UserWithGroupAndCategorymodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateUserWithGroupAndCategory(tz)
	if err != nil {
		return false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		err := dataAccess.UpdateUserWithGroupAndCategory(tz)
		if err != nil {
			return false, err, "Something Went Wrong"
		}
		return true, err, ""
	} else {
		return false, nil, "Data Already Exist."
	}

}
