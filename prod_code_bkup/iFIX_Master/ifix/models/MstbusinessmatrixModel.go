package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func InsertMstbusinessmatrix(tz *entities.MstbusinessmatrixEntity) (int64, bool, error, string) {
	logger.Log.Println("In side Mstbusinessmatrixmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	direction, err := dataAccess.Checkmatrixconfig(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if direction == 1 {
		count, err := dataAccess.CheckDuplicateMstbusinessmatrixurgencywise(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		if count.Total == 0 {
			id, err := dataAccess.InsertMstbusinessmatrix(tz)
			if err != nil {
				return 0, false, err, "Something Went Wrong"
			}

			return id, true, err, ""

		} else {
			return 0, false, nil, "Data Already Exist."
		}
	} else {
		count, err := dataAccess.CheckDuplicateMstbusinessmatrixcatwise(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		count1, err := dataAccess.CheckDuplicateEstimatedTimeeffort(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}

		if count.Total == 0 && count1 == 0 {
			id, err := dataAccess.InsertMstbusinessmatrix(tz)
			if err != nil {
				return 0, false, err, "Something Went Wrong"
			}
			_, err1 := dataAccess.InsertEstimatedTimeeffort(tz)
			if err1 != nil {
				return 0, false, err1, "Something Went Wrong"
			}
			return id, true, err, ""
		} else {
			return 0, false, nil, "Data Already Exist."
		}

	}

}

func GetAllMstbusinessmatrix(page *entities.MstbusinessmatrixEntity) (entities.MstbusinessmatrixEntities, bool, error, string) {
	logger.Log.Println("In side Mstbusinessmatrixmodel")
	t := entities.MstbusinessmatrixEntities{}
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
	values, err1 := dataAccess.GetAllMstbusinessmatrix(page, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetMstbusinessmatrixCount(page, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

func DeleteMstbusinessmatrix(tz *entities.MstbusinessmatrixEntity) (bool, error, string) {
	logger.Log.Println("In side Mstbusinessmatrixmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteMstbusinessmatrix(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	err2 := dataAccess.DeleteEstimatedTimeeffort(tz)
	if err2 != nil {
		return false, err2, "Something Went Wrong"
	}
	return true, nil, ""
}

func UpdateMstbusinessmatrix(tz *entities.MstbusinessmatrixEntity) (bool, error, string) {
	logger.Log.Println("In side Mstbusinessmatrixmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	direction, err := dataAccess.Checkmatrixconfig(tz)
	if err != nil {
		return false, err, "Something Went Wrong"
	}
	if direction == 1 {
		count, err := dataAccess.CheckDuplicateMstbusinessmatrixurgencywise(tz)
		if err != nil {
			return false, err, "Something Went Wrong"
		}
		if count.Total == 0 {
			err := dataAccess.UpdateMstbusinessmatrix(tz)
			if err != nil {
				return false, err, "Something Went Wrong"
			}
			return true, err, ""
		} else {
			return false, nil, "Data Already Exist."
		}
	} else {
		count, err := dataAccess.CheckDuplicateMstbusinessmatrixcatwise(tz)
		if err != nil {
			return false, err, "Something Went Wrong"
		}
		if count.Total == 0 {
			err := dataAccess.UpdateMstbusinessmatrix(tz)
			if err != nil {
				return false, err, "Something Went Wrong"
			}
			err1 := dataAccess.UpdateEstimatedTime(tz)
			if err1 != nil {
				return false, err1, "Something Went Wrong"
			}
			return true, err, ""
		} else {
			return false, nil, "Data Already Exist."
		}
	}
}

func Checkbusinessmatrixconfig(tz *entities.MstbusinessmatrixEntity) (int64, bool, error, string) {
	logger.Log.Println("In side Mstbusinessmatrixmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	data, err1 := dataAccess.Checkmatrixconfig(tz)
	if err1 != nil {
		return 0, false, err1, "Something Went Wrong"
	}
	return data, true, nil, ""
}

func Getlastlevelcategoryname(page *entities.MstbusinessmatrixEntity) ([]entities.MstlastlevelEntity, bool, error, string) {
	logger.Log.Println("In side Mstbusinessmatrixmodel")
	t := []entities.MstlastlevelEntity{}
	db, err := config.ConnectMySqlDb()
	defer db.Close()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Getlastlevelcategoryname(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}

	return values, true, err, ""
}
