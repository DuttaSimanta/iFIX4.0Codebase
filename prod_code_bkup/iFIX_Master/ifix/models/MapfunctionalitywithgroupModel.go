package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func InsertMapfunctionalitywithgroup(tz *entities.MapfunctionalitywithgroupEntity) (int64, bool, error, string) {
	logger.Log.Println("In side Mapfunctionalitywithgroupmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	if len(tz.Diffid) > 0 && len(tz.Groupid) > 0 && len(tz.Recorddiffstatusid) > 0 {
		for a := 0; a < len(tz.Recorddiffstatusid); a++ {
			for b := 0; b < len(tz.Diffid); b++ {
				for c := 0; c < len(tz.Groupid); c++ {
					count, err := dataAccess.CheckDuplicateMapfunctionalitywithgroup(tz, tz.Diffid[b], tz.Groupid[c], tz.Recorddiffstatusid[a])
					if err != nil {
						return 0, false, err, "Something Went Wrong"
					}
					if count.Total == 0 {
						_, err := dataAccess.InsertMapfunctionalitywithgroup(tz, tz.Diffid[b], tz.Groupid[c], tz.Recorddiffstatusid[a])
						if err != nil {
							return 0, false, err, "Something Went Wrong"
						}
					}
				}
			}
		}
	} // here length checking ...
	if len(tz.Recorddiffstatusid) == 0 {
		for k := 0; k < len(tz.Diffid); k++ {
			for c := 0; c < len(tz.Groupid); c++ {
				count, err := dataAccess.CheckDuplicateMapfunctionalitywithgroupwithoutStatus(tz, tz.Diffid[k], tz.Groupid[c])
				if err != nil {
					return 0, false, err, "Something Went Wrong"
				}
				if count.Total == 0 {
					_, err := dataAccess.InsertMapfunctionalitywithgroupwithoutstatus(tz, tz.Diffid[k], tz.Groupid[c])
					if err != nil {
						return 0, false, err, "Something Went Wrong"
					}

				}
			}
		}
	}

	return 0, true, err, ""

}

func GetAllMapfunctionalitywithgroup(page *entities.MapfunctionalitywithgroupEntity) (entities.MapfunctionalitywithgroupEntities, bool, error, string) {
	logger.Log.Println("In side Mapfunctionalitywithgroupmodel")
	t := entities.MapfunctionalitywithgroupEntities{}
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

	values, err1 := dataAccess.GetAllMapfunctionalitywithgroup(page, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetMapfunctionalitywithgroupCount(page, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

func DeleteMapfunctionalitywithgroup(tz *entities.MapfunctionalitywithgroupEntity) (bool, error, string) {
	logger.Log.Println("In side Mapfunctionalitywithgroupmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteMapfunctionalitywithgroup(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

func UpdateMapfunctionalitywithgroup(tz *entities.MapfunctionalitywithgroupEntity) (bool, error, string) {
	logger.Log.Println("In side Mapfunctionalitywithgroupmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	var updatecount int
	var k int
	// if len(tz.Refuserid) == 0 {
	// 	for k := 0; k < len(tz.Diffid); k++ {
	// 		count, err := dataAccess.CheckDuplicateMapfunctionalitywithgroup(tz, tz.Diffid[k])
	// 		if err != nil {
	// 			return false, err, "Something Went Wrong"
	// 		}
	// 		if count.Total == 0 {
	// 			err := dataAccess.UpdateMapfunctionalitywithgroup(tz, tz.Diffid[k])
	// 			if err != nil {
	// 				return false, err, "Something Went Wrong"
	// 			}
	// 		} else {
	// 			updatecount++
	// 		}
	// 	}

	// } else {
	// 	for k := 0; k < len(tz.Diffid); k++ {
	// 		for p := 0; p < len(tz.Refuserid); p++ {
	// 			count, err := dataAccess.CheckDuplicateMapfunctionalitywithgroupwithuser(tz, tz.Diffid[k], tz.Refuserid[p])
	// 			if err != nil {
	// 				return false, err, "Something Went Wrong"
	// 			}
	// 			if count.Total == 0 {
	// 				_, err := dataAccess.InsertMapfunctionalitywithgroupwithuser(tz, tz.Diffid[k], tz.Refuserid[p])
	// 				if err != nil {
	// 					return false, err, "Something Went Wrong"
	// 				}

	// 			} else {
	// 				updatecount++
	// 			}
	// 		}
	// 	}

	// 	if len(tz.Refuserid) == updatecount {
	// 		return false, nil, "Data Already Exist."
	// 	}
	// }

	if len(tz.Diffid) > 0 && len(tz.Groupid) > 0 && len(tz.Recorddiffstatusid) > 0 {
		for a := 0; a < len(tz.Recorddiffstatusid); a++ {
			for b := 0; b < len(tz.Diffid); b++ {
				for c := 0; c < len(tz.Groupid); c++ {
					k++
					count, err := dataAccess.CheckDuplicateMapfunctionalitywithgroup(tz, tz.Diffid[b], tz.Groupid[c], tz.Recorddiffstatusid[a])
					if err != nil {
						return false, err, "Something Went Wrong"
					}
					if count.Total == 0 {
						err := dataAccess.UpdateMapfunctionalitywithgroup(tz, tz.Diffid[b], tz.Groupid[c], tz.Recorddiffstatusid[a])
						if err != nil {
							return false, err, "Something Went Wrong"
						}
					} else {
						updatecount++
					}
				}
			}
		}
	} // here length checking ...

	if k == updatecount {
		return false, nil, "Data Already Exist."
	}

	return true, err, ""
}

func GetAllOrganizationgrpnames(page *entities.MapfunctionalitywithgroupEntity) (entities.OrganizationgrpnameEntities, bool, error, string) {
	logger.Log.Println("In side Mapfunctionalitywithgroupmodel")
	t := entities.OrganizationgrpnameEntities{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetAllOrganizationgrpnames(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	t.Values = values
	return t, true, err, ""
}
