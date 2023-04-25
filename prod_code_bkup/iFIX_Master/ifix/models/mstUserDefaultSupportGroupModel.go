package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func InsertMstUserDefaultSupportGroup(tz *entities.MstUserDefaultSupportGroupEntity) (int64, bool, error, string) {
	logger.Log.Println("In side MstUserDefaultSupportGroupmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateMstUserDefaultSupportGroup(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		id, err := dataAccess.InsertMstUserDefaultSupportGroup(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		return id, true, err, ""
	} else {
		return 0, false, nil, "Data Already Exist."
	}
}

func GetAllMstUserDefaultSupportGroup(page *entities.MstUserDefaultSupportGroupEntity) (entities.MstUserDefaultSupportGroupEntities, bool, error, string) {
	logger.Log.Println("In side MstUserDefaultSupportGroupmodel")
	t := entities.MstUserDefaultSupportGroupEntities{}
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
	values, err1 := dataAccess.GetAllMstUserDefaultSupportGroup(page, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetMstUserDefaultSupportGroupCount(page, orgntype)

		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}
func DeleteMstUserDefaultSupportGroup(tz *entities.MstUserDefaultSupportGroupEntity) (bool, error, string) {
	logger.Log.Println("In side MstUserDefaultSupportGroupmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteMstUserDefaultSupportGroup(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

func UpdateMstUserDefaultSupportGroup(tz *entities.MstUserDefaultSupportGroupEntity) (bool, error, string) {
	logger.Log.Println("In side MstUserDefaultSupportGroupmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateMstUserDefaultSupportGroupUpdate(tz)
	if err != nil {
		return false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		err := dataAccess.UpdateMstUserDefaultSupportGroup(tz)
		if err != nil {
			return false, err, "Something Went Wrong"
		}
		return true, err, ""
	} else {
		return false, nil, "Data Already Exist."
	}
}




func MstUserSupportGroupChange(tz *entities.MstUserDefaultSupportGroupEntity) (bool, error, string) {
	logger.Log.Println("In side MstUserSupportGroupChangeModel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	ids, err := dataAccess.CheckDuplicateMstUserSupportGroupChange(tz)
	if err != nil {
		return false, err, "Something Went Wrong"
	}
	// log.Print("\n count ===  ",ids);
	// log.Print("\n len(ids) ===  ",len(ids));
	if len(ids) == 0 {
		// log.Print("\n Inside IFFFFFF.........  ");
		_, err := dataAccess.InsertMstUserDefaultSupportGroup(tz)
		if err != nil {
			return false, err, "Something Went Wrong"
		}
		return true, err, ""
	} else {
		// log.Print("\n ELSEEEEEEEEE.........  ");
		err := dataAccess.UpdateMstUserSupportGroupChange(tz, ids)
		if err != nil {
			return false, err, "Something Went Wrong"
		}
		return true, err, ""
	}
}


