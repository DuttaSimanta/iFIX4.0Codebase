package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func AddDashboardQueryCopy(tz *entities.DashboardQueryCopyEntity) (int64, bool, error, string) {
	logger.Log.Println("In side DashboardQueryCopy model")
	// db, err := config.ConnectMySqlDb()
	// defer db.Close()
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	tx, err := db.Begin()
	if err != nil {
		logger.Log.Println("database transaction connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	dataAccess1 := dao.TxConn{TX: tx}
	k := 0
	var id int64
	for i := 0; i < len(tz.Tilesids); i++ {
		check, err := dataAccess.Checktiles(tz, i)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}

		if check.Total == 0 {
			tilesname, err := dataAccess.Gettilesname(tz, i)
			if err != nil {
				return 0, false, err, "Something Went Wrong"
			}
			return 0, false, err, "'" + tilesname.TilesName + "' is not mapped, Please check it "

		}
		values, err := dataAccess.GetDashboardQueryCopy(tz, i)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		values.ToClientid = tz.ToClientid
		values.ToMstorgnhirarchyid = tz.ToMstorgnhirarchyid
		values.ToRecordDiffid = tz.ToRecordDiffid
		values.Tilesid = tz.Tilesids[i]
		values.QueryType = tz.QueryType
		values.Ismanagerialview = tz.Ismanagerialview
		count, err := dataAccess.CheckDuplicateDashboardQueryCopy(&values)
		if err != nil {
			err = tx.Rollback()
			if err != nil {
				// log.Print("MoveWorkflow  Statement Commit error", err)
				logger.Log.Print("DashboardQueryCopy  Statement Rollback error", err)
				return 0, false, err, ""
			}
			return 0, false, err, "Something Went Wrong"
		}
		if count.Total == 0 {
			id, err = dataAccess1.AddDashboardQueryCopy(&values)
			if err != nil {
				err = tx.Rollback()
				if err != nil {
					// log.Print("MoveWorkflow  Statement Commit error", err)
					logger.Log.Print("DashboardQueryCopy  Statement Rollback error", err)
					return 0, false, err, ""
				}
				return 0, false, err, "Something Went Wrong"
			}

			//return id, true, err, ""
		} else {
			//return 0, false, nil, "Data Already Exist."
			k++
		}
	}
	if k < len(tz.Tilesids) {
		err = tx.Commit()
		if err != nil {
			//log.Print("MoveWorkflow  Statement Commit error", err)
			logger.Log.Print("DashboardQueryCopy  Statement Commit error", err)
			return 0, false, err, ""
		}
		return id, true, err, ""
	} else {
		err = tx.Rollback()
		if err != nil {
			// log.Print("MoveWorkflow  Statement Commit error", err)
			logger.Log.Print("DashboardQueryCopy  Statement Commit error", err)
			return 0, false, err, ""
		}
		return 0, false, nil, "Data Already Exist."
	}
}

func GetAllDashboardQueryCopy(page *entities.DashboardQueryCopyEntity) (entities.DashboardQueryCopyEntities, bool, error, string) {
	logger.Log.Println("In side DashboardQueryCopy model")
	t := entities.DashboardQueryCopyEntities{}
	// db, err := config.ConnectMySqlDb()
	// defer db.Close()
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
	values, err1 := dataAccess.GetAllDashboardQueryCopy(page, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}

	if page.Offset == 0 {
		total, err1 := dataAccess.GetDashboardQueryCopyCount(page, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

func DeleteDashboardQueryCopy(tz *entities.DashboardQueryCopyEntity) (bool, error, string) {
	logger.Log.Println("In side DashboardQueryCopy model")
	// db, err := config.ConnectMySqlDb()
	// defer db.Close()
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteDashboardQueryCopy(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

// func UpdateDashboardQueryCopy(tz *entities.DashboardQueryCopyEntity) (bool, error, string) {
// 	logger.Log.Println("In side DashboardQueryCopy model")
// 	db, err := config.ConnectMySqlDb()
// 	defer db.Close()
// 	if err != nil {
// 		logger.Log.Println("database connection failure", err)
// 		return false, err, "Something Went Wrong"
// 	}
// 	dataAccess := dao.DbConn{DB: db}
// 	count, err := dataAccess.CheckDuplicateDashboardQueryCopy(tz)
// 	if err != nil {
// 		return false, err, "Something Went Wrong"
// 	}
// 	if count.Total == 0 {
// 		err := dataAccess.UpdateDashboardQueryCopy(tz)
// 		if err != nil {
// 			return false, err, "Something Went Wrong"
// 		}
// 		return true, err, ""
// 	} else {
// 		return false, nil, "Data Already Exist."
// 	}
// }
