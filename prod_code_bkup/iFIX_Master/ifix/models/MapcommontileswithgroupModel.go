package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func InsertMapcommontileswithgroup(tz *entities.MapcommontileswithgroupEntity) (int64, bool, error, string) {
	logger.Log.Println("In side Mapcommontileswithgroupmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	// count,err :=dataAccess.CheckDuplicateMapcommontileswithgroup(tz)
	// if err != nil {
	//     return 0, false, err, "Something Went Wrong"
	// }
	// if count.Total == 0 {
	//     id, err := dataAccess.InsertMapcommontileswithgroup(tz)
	//     if err != nil {
	//         return 0, false, err, "Something Went Wrong"
	//     }
	//     return id, true, err, ""
	// }else{
	//     return 0, false, nil, "Data Already Exist."
	// }
	var updatecount int
	for k := 0; k < len(tz.Groupid); k++ {
		count, err := dataAccess.CheckDuplicateMapcommontileswithgroup(tz, tz.Groupid[k])
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		if count.Total == 0 {
			_, err := dataAccess.InsertMapcommontileswithgroup(tz, tz.Groupid[k])
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

func GetAllMapcommontileswithgroup(page *entities.MapcommontileswithgroupEntity) (entities.MapcommontileswithgroupEntities, bool, error, string) {
	logger.Log.Println("In side Mapcommontileswithgroupmodel")
	t := entities.MapcommontileswithgroupEntities{}
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
	values, err1 := dataAccess.GetAllMapcommontileswithgroup(page, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetMapcommontileswithgroupCount(page, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

func DeleteMapcommontileswithgroup(tz *entities.MapcommontileswithgroupEntity) (bool, error, string) {
	logger.Log.Println("In side Mapcommontileswithgroupmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteMapcommontileswithgroup(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

func UpdateMapcommontileswithgroup(tz *entities.MapcommontileswithgroupEntity) (bool, error, string) {
	logger.Log.Println("In side Mapcommontileswithgroupmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	// count, err := dataAccess.CheckDuplicateMapcommontileswithgroup(tz)
	// if err != nil {
	// 	return false, err, "Something Went Wrong"
	// }
	// if count.Total == 0 {
	// 	err := dataAccess.UpdateMapcommontileswithgroup(tz)
	// 	if err != nil {
	// 		return false, err, "Something Went Wrong"
	// 	}
	// 	return true, err, ""
	// } else {
	// 	return false, nil, "Data Already Exist."
	// }

	var updatecount int
	for k := 0; k < len(tz.Groupid); k++ {
		count, err := dataAccess.CheckDuplicateMapcommontileswithgroup(tz, tz.Groupid[k])
		if err != nil {
			return false, err, "Something Went Wrong"
		}
		if count.Total == 0 {
			err := dataAccess.UpdateMapcommontileswithgroup(tz, tz.Groupid[k])
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
