package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	// "fmt"
)

func InsertMstsupportgrptermmap(tz *entities.MstsupportgrptermmapEntity) (int64, bool, error, string) {
	logger.Log.Println("In side Mstsupportgrptermmapmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	var count int64
	k := len(tz.Termid)
	a := make([]int64, k)
	dataAccess := dao.DbConn{DB: db}
	for i, _ := range tz.Termid {

		value, err := dataAccess.CheckDuplicateMstsupportgrptermmap(tz, i)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		if value.Total == 0 {
			count = count + 1
			a[i] = 1
		}
		if value.Total == 1 {
			a[i] = 0
		}
	}
	logger.Log.Println("In side Mstsupportgrptermmapmodel")
	var idd int64
	if count > 0 {
		for i, _ := range tz.Termid {
			if a[i] == 1 {
				id, err := dataAccess.InsertMstsupportgrptermmap(tz, i)
				idd = id
				if err != nil {
					return 0, false, err, "Something Went Wrong"
				}
			}
		}
		return idd, true, err, ""
	} else {
		return 0, false, nil, "Data Already Exist."
	}
	//fmt.Println("inside models")

}

func GetAllMstsupportgrptermmap(page *entities.MstsupportgrptermmapEntity) (entities.MstsupportgrptermmapEntities, bool, error, string) {
	logger.Log.Println("In side Mstsupportgrptermmapmodel")
	t := entities.MstsupportgrptermmapEntities{}
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
	values, err1 := dataAccess.GetAllMstsupportgrptermmap(page, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetMstsupportgrptermmapCount(page, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

func DeleteMstsupportgrptermmap(tz *entities.MstsupportgrptermmapEntity) (bool, error, string) {
	logger.Log.Println("In side Mstsupportgrptermmapmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteMstsupportgrptermmap(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

func UpdateMstsupportgrptermmap(tz *entities.MstsupportgrptermmapEntity) (bool, error, string) {
	logger.Log.Println("In side Mstsupportgrptermmapmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, _ := dataAccess.CheckDuplicateMstsupportgrptermmap(tz, 0)
	/*if err != nil {
	    return false, err, "Something Went Wrong"
	}*/
	if count.Total == 0 {
		err := dataAccess.UpdateMstsupportgrptermmap(tz)
		if err != nil {
			return false, err, "Something Went Wrong"
		}
		return true, err, ""
	} else {
		return false, nil, "Data Already Exist."
	}
}
