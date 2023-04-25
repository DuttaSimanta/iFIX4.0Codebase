package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func AddDashboardQuery(tz *entities.DashboardQueryEntity) (int64, bool, error, string) {
	logger.Log.Println("In side DashboardQuery model")
	// db, err := config.ConnectMySqlDb()
	// defer db.Close()
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateDashboardQuery(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		id, err := dataAccess.AddDashboardQuery(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		return id, true, err, ""
	} else {
		return 0, false, nil, "Data Already Exist."
	}
}

// func GetAllDashboardQuery(page *entities.DashboardQueryEntity) (entities.DashboardQueryEntities, bool, error, string) {
//     logger.Log.Println("In side DashboardQuery model")
//     t := entities.DashboardQueryEntities{}
//     db, err := config.ConnectMySqlDb()
//     defer db.Close()
//     if err != nil {
//         logger.Log.Println("database connection failure", err)
//          return t, false, err, "Something Went Wrong"
//     }
//     dataAccess := dao.DbConn{DB: db}
//     values, err1 := dataAccess.GetAllDashboardQuery(page)
//     if err1 != nil {
//         return t, false, err1, "Something Went Wrong"
//     }
//     if page.Offset == 0{
//         total, err1 := dataAccess.GetDashboardQueryCount(page)
//     if err1 != nil {
//         return t, false, err1, "Something Went Wrong"
//     }
//     t.Total=total.Total
//     t.Values=values
//     }
//     t.Values=values
//     return t, true, err, ""
// }

// func DeleteDashboardQuery(tz *entities.DashboardQueryEntity) (bool,error, string) {
//     logger.Log.Println("In side DashboardQuery model")
//     db, err := config.ConnectMySqlDb()
//     defer db.Close()
//     if err != nil {
//         logger.Log.Println("database connection failure", err)
//         return false,err, "Something Went Wrong"
//     }
//     dataAccess := dao.DbConn{DB: db}
//     err1 := dataAccess.DeleteDashboardQuery(tz)
//     if err1 != nil {
//         return false,err1, "Something Went Wrong"
//     }
//     return true,nil, ""
// }

func UpdateDashboardQuery(tz *entities.DashboardQueryEntity) (bool, error, string) {
	logger.Log.Println("In side DashboardQuery model")
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
	count, err := dataAccess.CheckDuplicateDashboardQuery(tz)
	if err != nil {
		return false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		err := dataAccess.UpdateDashboardQuery(tz)
		if err != nil {
			return false, err, "Something Went Wrong"
		}
		return true, err, ""
	} else {
		return false, nil, "Data Already Exist."
	}
}
