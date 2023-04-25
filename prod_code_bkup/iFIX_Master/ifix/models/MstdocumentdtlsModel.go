package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func InsertMstdocumentdtls(tz *entities.MstdocumentdtlsEntity) (int64, bool, error, string) {
	logger.Log.Println("In side Mstdocumentdtlsmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	// count,err :=dataAccess.CheckDuplicateMstdocumentdtls(tz)
	// if err != nil {
	//     return 0, false, err, "Something Went Wrong"
	// }
	// if count.Total == 0 {
	//     id, err := dataAccess.InsertMstdocumentdtls(tz)
	//     if err != nil {
	//         return 0, false, err, "Something Went Wrong"
	//     }
	//     return id, true, err, ""
	// }else{
	//     return 0, false, nil, "Data Already Exist."
	// }

	var updatecount int
	for k := 0; k < len(tz.Groupid); k++ {
		count, err := dataAccess.CheckDuplicateMstdocumentdtls(tz, tz.Groupid[k])
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		if count.Total == 0 {
			_, err := dataAccess.InsertMstdocumentdtls(tz, tz.Groupid[k])
			if err != nil {
				return 0, false, err, "Something Went Wrong"
			}

		} else {
			updatecount++
		}
	}
	if len(tz.Groupid) == updatecount {
		return 0, false, nil, "Data Already Exist."
	}
	return 0, true, err, ""
}

func GetAllMstdocumentdtls(page *entities.MstdocumentdtlsEntity) (entities.MstdocumentdtlsEntities, bool, error, string) {
	logger.Log.Println("In side Mstdocumentdtlsmodel")
	t := entities.MstdocumentdtlsEntities{}
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
	values, err1 := dataAccess.GetAllMstdocumentdtls(page, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetMstdocumentdtlsCount(page, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

func DeleteMstdocumentdtls(tz *entities.MstdocumentdtlsEntity) (bool, error, string) {
	logger.Log.Println("In side Mstdocumentdtlsmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteMstdocumentdtls(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

func UpdateMstdocumentdtls(tz *entities.MstdocumentdtlsEntity) (bool, error, string) {
	logger.Log.Println("In side Mstdocumentdtlsmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	// count, err := dataAccess.CheckDuplicateMstdocumentdtls(tz)
	// if err != nil {
	// 	return false, err, "Something Went Wrong"
	// }
	// if count.Total == 0 {
	// 	err := dataAccess.UpdateMstdocumentdtls(tz)
	// 	if err != nil {
	// 		return false, err, "Something Went Wrong"
	// 	}
	// 	return true, err, ""
	// } else {
	// 	return false, nil, "Data Already Exist."
	// }

	var updatecount int
	for k := 0; k < len(tz.Groupid); k++ {
		count, err := dataAccess.CheckDuplicateMstdocumentdtls(tz, tz.Groupid[k])
		if err != nil {
			return false, err, "Something Went Wrong"
		}
		if count.Total == 0 {
			err := dataAccess.UpdateMstdocumentdtls(tz, tz.Groupid[k])
			if err != nil {
				return false, err, "Something Went Wrong"
			}

		} else {
			updatecount++
		}
	}
	if len(tz.Groupid) == updatecount {
		return false, nil, "Data Already Exist."
	}
	return true, err, ""
}
