package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

// func InsertCountry(tz *entities.CountryEntity) (int64, bool, error, string) {
//     logger.Log.Println("In side Countrymodel")
//     db, err := config.ConnectMySqlDb()
//     defer db.Close()
//     if err != nil {
//         logger.Log.Println("database connection failure", err)
//         return 0, false, err, "Something Went Wrong"
//     }
//     dataAccess := dao.DbConn{DB: db}
//     count,err :=dataAccess.CheckDuplicateCountry(tz)
//     if err != nil {
//         return 0, false, err, "Something Went Wrong"
//     }
//     if count.Total == 0 {
//         id, err := dataAccess.InsertCountry(tz)
//         if err != nil {
//             return 0, false, err, "Something Went Wrong"
//         }
//         return id, true, err, ""
//     }else{
//         return 0, false, nil, "Data Already Exist."
//     }
// }

func GetAllCountry() (entities.CountryEntities, bool, error, string) {
	logger.Log.Println("In side Countrymodel")
	t := entities.CountryEntities{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetAllCountry()
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	// if page.Offset == 0{
	//     total, err1 := dataAccess.GetCountryCount(page)
	// if err1 != nil {
	//     return t, false, err1, "Something Went Wrong"
	// }
	// t.Total=total.Total
	// t.Values=values
	// }
	t.Values = values
	return t, true, err, ""
}

// func DeleteCountry(tz *entities.CountryEntity) (bool,error, string) {
//     logger.Log.Println("In side Countrymodel")
//     db, err := config.ConnectMySqlDb()
//     defer db.Close()
//     if err != nil {
//         logger.Log.Println("database connection failure", err)
//         return false,err, "Something Went Wrong"
//     }
//     dataAccess := dao.DbConn{DB: db}
//     err1 := dataAccess.DeleteCountry(tz)
//     if err1 != nil {
//         return false,err1, "Something Went Wrong"
//     }
//     return true,nil, ""
// }

// func UpdateCountry(tz *entities.CountryEntity) (bool,error, string) {
//     logger.Log.Println("In side Countrymodel")
//     db, err := config.ConnectMySqlDb()
//     defer db.Close()
//     if err != nil {
//         logger.Log.Println("database connection failure", err)
//         return false,err, "Something Went Wrong"
//     }
//     dataAccess := dao.DbConn{DB: db}
//     count,err :=dataAccess.CheckDuplicateCountry(tz)
//     if err != nil {
//         return false, err, "Something Went Wrong"
//     }
//     if count.Total == 0 {
//          err := dataAccess.UpdateCountry(tz)
//         if err != nil {
//             return false, err, "Something Went Wrong"
//         }
//         return true, err, ""
//     }else{
//         return false, nil, "Data Already Exist."
//     }
// }
