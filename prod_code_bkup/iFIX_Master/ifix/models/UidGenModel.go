package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func GetAllOrganizationwithOrgtype(tz *entities.MstorgnhierarchywithOrgtypeEntity) ([]entities.MstorgnhierarchywithOrgtypeEntityResp, bool, error, string) {
	logger.Log.Println("In side model")
	t := []entities.MstorgnhierarchywithOrgtypeEntityResp{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	orgntype, err1 := dataAccess.GetOrgnType(tz.ClientID, tz.Mstorgnhirarchyid)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	values, err1 := dataAccess.GetAllOrganizationwithOrgtype(tz, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}

	return values, true, err, ""
}

func InsertUidGen(tz *entities.UidGenEntity) (int64, bool, error, string) {
	logger.Log.Println("In side UidGenmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	orgntype, err1 := dataAccess.GetOrgnType(tz.Clientid, tz.Mstorgnhirarchyid)
	if err1 != nil {
		return 0, false, err1, "Something Went Wrong"
	}
	logger.Log.Println("DIfftype ", tz.Difftypeseq)
	logger.Log.Println("Orgntype", orgntype)
	if tz.Difftypeseq == 11 && orgntype > 2 {
		return 0, false, err, "Adding UID is not allowed for this Organization type"
	}

	count, err := dataAccess.CheckDuplicateUidGen(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		id, err := dataAccess.InsertUidGen(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		return id, true, err, ""
	} else {
		return 0, false, nil, "Data Already Exist."
	}
}

func GetAllUidGen(page *entities.UidGenEntity) (entities.UidGenEntities, bool, error, string) {
	logger.Log.Println("In side UidGenmodel")
	t := entities.UidGenEntities{}
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
	values, err1 := dataAccess.GetAllUidGen(page, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetUidGenCount(page, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

func DeleteUidGen(tz *entities.UidGenEntity) (bool, error, string) {
	logger.Log.Println("In side UidGenmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	if tz.Difftypeseq == 11 {
		return false, nil, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteUidGen(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

func UpdateUidGen(tz *entities.UidGenEntity) (bool, error, string) {
	logger.Log.Println("In side UidGenmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	// if tz.Difftypeseq == 11 {
	// 	err := dataAccess.UpdateUidGenforemail(tz)
	// 	if err != nil {
	// 		return false, err, "Something Went Wrong"
	// 	}
	// 	return true, err, ""
	// }
	// count, err := dataAccess.CheckDuplicateUidGen(tz)
	// if err != nil {
	// 	return false, err, "Something Went Wrong"
	// }
	// if count.Total == 0 {
	err1 := dataAccess.UpdateUidGen(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, err1, ""
	// } else {
	// 	return false, nil, "Data Already Exist."
	// }
}
