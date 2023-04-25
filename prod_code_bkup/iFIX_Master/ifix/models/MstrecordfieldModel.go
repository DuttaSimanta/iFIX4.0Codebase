package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func InsertMstrecordfield(tz *entities.MstrecordfieldEntity) (int64, bool, error, string) {
	logger.Log.Println("In side Mstrecordfieldmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateMstrecordfield(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		id, err := dataAccess.InsertMstrecordfield(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		for _, s := range tz.MstrecordfielddiffEntities {
			_, err := dataAccess.InsertMstrecordfielddiff(tz, &s, id)
			if err != nil {
				return 0, false, err, "Something Went Wrong"
			}
		}
		return id, true, err, ""
	} else {
		return 0, false, nil, "Data Already Exist."
	}
}

func GetAllMstrecordfield(page *entities.MstrecordfieldEntity) (entities.MstrecordfieldEntities, bool, error, string) {
	logger.Log.Println("In side Mstrecordfieldmodel")
	t := entities.MstrecordfieldEntities{}
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
	values1, err1 := dataAccess.GetAllMstrecordfield(page, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	for i, dt := range values1 {
		valdiff, _ := dataAccess.GetAllMstrecordfielddiff(&dt)
		values1[i].MstrecordfielddiffEntities = valdiff
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetMstrecordfieldCount(page, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values1
	}
	t.Values = values1
	return t, true, err, ""
}

func DeleteMstrecordfield(tz *entities.MstrecordfieldEntity) (bool, error, string) {
	logger.Log.Println("In side Mstrecordfieldmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteMstrecordfield(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	err2 := dataAccess.DeleteMstrecordfielddiff(tz)
	if err2 != nil {
		return false, err2, "Something Went Wrong"
	}
	return true, nil, ""
}

func UpdateMstrecordfield(tz *entities.MstrecordfieldEntity) (bool, error, string) {
	logger.Log.Println("In side Mstrecordfieldmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateMstrecordfield(tz)
	if err != nil {
		return false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		err := dataAccess.UpdateMstrecordfield(tz)
		if err != nil {
			return false, err, "Something Went Wrong"
		}
		err1 := dataAccess.DeleteMstrecordfielddiff(tz)
		if err1 != nil {
			return false, err1, "Something Went Wrong"
		}
		for _, s := range tz.MstrecordfielddiffEntities {
			dataAccess.InsertMstrecordfielddiff(tz, &s, tz.Id)
		}
		return true, err, ""
	} else {
		return false, nil, "Data Already Exist."
	}
}
