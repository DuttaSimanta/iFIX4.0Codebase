package models


import (
  "iFIX/ifix/config"
  "iFIX/ifix/entities"
  "iFIX/ifix/logger"
  "iFIX/ifix/dao"
  )

func InsertRecordconfig(tz *entities.RecordconfigEntity) (int64, bool, error, string) {
    logger.Log.Println("In side Recordconfigmodel")
    lock.Lock()
    defer lock.Unlock()
    db, err := config.ConnectMySqlDbSingleton()
    if err != nil {
        logger.Log.Println("database connection failure", err)
        return 0, false, err, "Something Went Wrong"
    }
    dataAccess := dao.DbConn{DB: db}
    count,err :=dataAccess.CheckDuplicateRecordconfig(tz)
    if err != nil {
        return 0, false, err, "Something Went Wrong"
    }
    if count.Total == 0 {
        id, err := dataAccess.InsertRecordconfig(tz)
        if err != nil {
            return 0, false, err, "Something Went Wrong"
        }
        return id, true, err, ""
    }else{
        return 0, false, nil, "Module Name Already Exist."
    }
}


func GetAllRecordconfig(page *entities.RecordconfigEntity) (entities.RecordconfigEntities, bool, error, string) {
    logger.Log.Println("In side Recordconfigmodel")
    t := entities.RecordconfigEntities{}
    lock.Lock()
    defer lock.Unlock()
    db, err := config.ConnectMySqlDbSingleton()
    if err != nil {
        logger.Log.Println("database connection failure", err)
         return t, false, err, "Something Went Wrong"
    }
    dataAccess := dao.DbConn{DB: db}
    values, err1 := dataAccess.GetAllRecordconfig(page)
    if err1 != nil {
        return t, false, err1, "Something Went Wrong"
    }
    if page.Offset == 0{
        total, err1 := dataAccess.GetRecordconfigCount(page)
    if err1 != nil {
        return t, false, err1, "Something Went Wrong"
    }
    t.Total=total.Total
    t.Values=values
    }
    t.Values=values
    return t, true, err, ""
}


func DeleteRecordconfig(tz *entities.RecordconfigEntity) (bool,error, string) {
    logger.Log.Println("In side Recordconfigmodel")
    lock.Lock()
    defer lock.Unlock()
    db, err := config.ConnectMySqlDbSingleton()
    if err != nil {
        logger.Log.Println("database connection failure", err)
        return false,err, "Something Went Wrong"
    }
    dataAccess := dao.DbConn{DB: db}
    err1 := dataAccess.DeleteRecordconfig(tz)
    if err1 != nil {
        return false,err1, "Something Went Wrong"
    }
    return true,nil, ""
}


func UpdateRecordconfig(tz *entities.RecordconfigEntity) (bool,error, string) {
    logger.Log.Println("In side Recordconfigmodel")
    lock.Lock()
    defer lock.Unlock()
    db, err := config.ConnectMySqlDbSingleton()
    if err != nil {
        logger.Log.Println("database connection failure", err)
        return false,err, "Something Went Wrong"
    }
    dataAccess := dao.DbConn{DB: db}
    err1 := dataAccess.UpdateRecordconfig(tz)
    if err1 != nil {
        return false,err1, "Something Went Wrong"
    }
    return true,nil, ""
}


