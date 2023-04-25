package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

// func InsertCity(tz *entities.CityEntity) (int64, bool, error, string) {
//     logger.Log.Println("In side Citymodel")
//     db, err := config.ConnectMySqlDb()
//     defer db.Close()
//     if err != nil {
//         logger.Log.Println("database connection failure", err)
//         return 0, false, err, "Something Went Wrong"
//     }
//     dataAccess := dao.DbConn{DB: db}
//     count,err :=dataAccess.CheckDuplicateCity(tz)
//     if err != nil {
//         return 0, false, err, "Something Went Wrong"
//     }
//     if count.Total == 0 {
//         id, err := dataAccess.InsertCity(tz)
//         if err != nil {
//             return 0, false, err, "Something Went Wrong"
//         }
//         return id, true, err, ""
//     }else{
//         return 0, false, nil, "Data Already Exist."
//     }
// }

func GetAllCity() (entities.CityEntities, bool, error, string) {
	logger.Log.Println("In side Citymodel")
	t := entities.CityEntities{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetAllCity()
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	t.Values = values
	return t, true, err, ""
}

// func DeleteCity(tz *entities.CityEntity) (bool,error, string) {
//     logger.Log.Println("In side Citymodel")
//     db, err := config.ConnectMySqlDb()
//     defer db.Close()
//     if err != nil {
//         logger.Log.Println("database connection failure", err)
//         return false,err, "Something Went Wrong"
//     }
//     dataAccess := dao.DbConn{DB: db}
//     err1 := dataAccess.DeleteCity(tz)
//     if err1 != nil {
//         return false,err1, "Something Went Wrong"
//     }
//     return true,nil, ""
// }

// func UpdateCity(tz *entities.CityEntity) (bool,error, string) {
//     logger.Log.Println("In side Citymodel")
//     db, err := config.ConnectMySqlDb()
//     defer db.Close()
//     if err != nil {
//         logger.Log.Println("database connection failure", err)
//         return false,err, "Something Went Wrong"
//     }
//     dataAccess := dao.DbConn{DB: db}
//     count,err :=dataAccess.CheckDuplicateCity(tz)
//     if err != nil {
//         return false, err, "Something Went Wrong"
//     }
//     if count.Total == 0 {
//          err := dataAccess.UpdateCity(tz)
//         if err != nil {
//             return false, err, "Something Went Wrong"
//         }
//         return true, err, ""
//     }else{
//         return false, nil, "Data Already Exist."
//     }
// }
