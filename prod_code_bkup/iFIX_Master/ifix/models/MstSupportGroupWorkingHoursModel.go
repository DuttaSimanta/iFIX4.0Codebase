package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func InsertMstSupportGroupWorkingHours(tz *entities.MstSupportGroupWorkingHoursEntity) (int64, bool, error, string) {
	logger.Log.Println("In side MstSupportGroupWorkingHoursmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}

	details := tz.Details
	for k := 0; k < len(details); k++ {
		count, err := dataAccess.CheckDuplicateMstSupportGroupWorkingHours(tz, &details[k])
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		if count.Total == 0 {
			_, err := dataAccess.InsertMstSupportGroupWorkingHours(tz, &details[k])
			if err != nil {
				return 0, false, err, "Something Went Wrong"
			}
		} else {
			return 0, false, nil, "Data Already Exist."
		}
	}
	return 0, true, err, ""
}

func GetAllMstSupportGroupWorkingHours(page *entities.MstSupportGroupWorkingHoursEntity) (entities.MstSupportGroupWorkingHoursEntities, bool, error, string) {
	logger.Log.Println("In side MstSupportGroupWorkingHoursmodel")
	t := entities.MstSupportGroupWorkingHoursEntities{}
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
	values, err1 := dataAccess.GetAllMstSupportGroupWorkingHours(page, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetMstSupportGroupWorkingHoursCount(page, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

func DeleteMstSupportGroupWorkingHours(tz *entities.MstSupportGroupWorkingHoursEntity) (bool, error, string) {
	logger.Log.Println("In side MstSupportGroupWorkingHoursmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteMstSupportGroupWorkingHours(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

func UpdateMstSupportGroupWorkingHours(tz *entities.MstSupportGroupWorkingHoursUpdateEntity) (bool, error, string) {
	logger.Log.Println("In side MstSupportGroupWorkingHoursmodel")
	db, err := config.ConnectMySqlDb()
	defer db.Close()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.UpdateMstSupportGroupWorkingHours(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, err, ""
}

func GetSupportGroupWiseWorkingHours(page *entities.MstSupportGroupWorkingHoursEntity) ([]entities.MstSupportGroupWorkingHoursresponseEntity, bool, error, string) {
	logger.Log.Println("In side MstSupportGroupWorkingHoursmodel")
	t := []entities.MstSupportGroupWorkingHoursresponseEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetSupportGroupWiseWorkingHours(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}

	return values, true, err, ""
}
