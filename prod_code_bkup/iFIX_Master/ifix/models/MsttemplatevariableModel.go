package models


import (
  "iFIX/ifix/config"
  "iFIX/ifix/entities"
  "iFIX/ifix/logger"
  "iFIX/ifix/dao"
  )

func InsertMsttemplatevariable(tz *entities.MsttemplatevariableEntity) (int64, bool, error, string) {
    logger.Log.Println("In side Msttemplatevariablemodel")
    lock.Lock()
    defer lock.Unlock()
    db, err := config.ConnectMySqlDbSingleton()
    if err != nil {
        logger.Log.Println("database connection failure", err)
        return 0, false, err, "Something Went Wrong"
    }
    dataAccess := dao.DbConn{DB: db}
    count,err :=dataAccess.CheckDuplicateMsttemplatevariable(tz)
    if err != nil {
        return 0, false, err, "Something Went Wrong"
    }
    if count.Total == 0 {
        id, err := dataAccess.InsertMsttemplatevariable(tz)
        if err != nil {
            return 0, false, err, "Something Went Wrong"
        }
        return id, true, err, ""
    }else{
        return 0, false, nil, "Data Already Exist."
    }
}


func GetAllMsttemplatevariable(page *entities.MsttemplatevariableEntity) (entities.MsttemplatevariableEntities, bool, error, string) {
    logger.Log.Println("In side Msttemplatevariablemodel")
    t := entities.MsttemplatevariableEntities{}
    lock.Lock()
    defer lock.Unlock()
    db, err := config.ConnectMySqlDbSingleton()
    if err != nil {
        logger.Log.Println("database connection failure", err)
         return t, false, err, "Something Went Wrong"
    }
    dataAccess := dao.DbConn{DB: db}
    values, err1 := dataAccess.GetAllMsttemplatevariable(page)
    if err1 != nil {
        return t, false, err1, "Something Went Wrong"
    }
    if page.Offset == 0{
        total, err1 := dataAccess.GetMsttemplatevariableCount(page)
    if err1 != nil {
        return t, false, err1, "Something Went Wrong"
    }
    t.Total=total.Total
    t.Values=values
    }
    t.Values=values
    return t, true, err, ""
}


func DeleteMsttemplatevariable(tz *entities.MsttemplatevariableEntity) (bool,error, string) {
    logger.Log.Println("In side Msttemplatevariablemodel")
    lock.Lock()
    defer lock.Unlock()
    db, err := config.ConnectMySqlDbSingleton()
    if err != nil {
        logger.Log.Println("database connection failure", err)
        return false,err, "Something Went Wrong"
    }
    dataAccess := dao.DbConn{DB: db}
    err1 := dataAccess.DeleteMsttemplatevariable(tz)
    if err1 != nil {
        return false,err1, "Something Went Wrong"
    }
    return true,nil, ""
}


func UpdateMsttemplatevariable(tz *entities.MsttemplatevariableEntity) (bool,error, string) {
    logger.Log.Println("In side Msttemplatevariablemodel")
    lock.Lock()
    defer lock.Unlock()
    db, err := config.ConnectMySqlDbSingleton()
    if err != nil {
        logger.Log.Println("database connection failure", err)
        return false,err, "Something Went Wrong"
    }
    dataAccess := dao.DbConn{DB: db}
    count,err :=dataAccess.CheckDuplicateMsttemplatevariable(tz)
    if err != nil {
        return false, err, "Something Went Wrong"
    }
    if count.Total == 0 {
         err := dataAccess.UpdateMsttemplatevariable(tz)
        if err != nil {
            return false, err, "Something Went Wrong"
        }
        return true, err, ""
    }else{
        return false, nil, "Data Already Exist."
    }
}


