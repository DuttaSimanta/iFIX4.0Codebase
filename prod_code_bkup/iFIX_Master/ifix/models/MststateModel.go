package models


import (
  "iFIX/ifix/config"
  "iFIX/ifix/entities"
  "iFIX/ifix/logger"
  "iFIX/ifix/dao"
  )

func InsertMststate(tz *entities.MststateEntity) (int64, bool, error, string) {
    logger.Log.Println("In side Mststatemodel")
    lock.Lock()
    defer lock.Unlock()
    db, err := config.ConnectMySqlDbSingleton()
    if err != nil {
        logger.Log.Println("database connection failure", err)
        return 0, false, err, "Something Went Wrong"
    }
    dataAccess := dao.DbConn{DB: db}
    count,err :=dataAccess.CheckDuplicateMststate(tz)
    if err != nil {
        return 0, false, err, "Something Went Wrong"
    }
    if count.Total == 0 {
        if tz.Seqno == 0 {
            seq, err := dataAccess.GetLastSeqFromstate(tz)
            if err != nil {
                return 0, false, err, "Something Went Wrong"
            }
            if len(seq) == 0 {
                tz.Seqno = 1
            } else {
                tz.Seqno = seq[0].Seqno + 1
            }
        }else{
            count1, err := dataAccess.CheckDuplicatestateseq(tz)
            if err != nil {
                return 0, false, err, "Something Went Wrong"
            }
            if count1.Total >0{
                return 0, false, nil, "Cannot create state with same sequence"
            }
        }
        id, err := dataAccess.InsertMststate(tz)
        if err != nil {
            return 0, false, err, "Something Went Wrong"
        }
        return id, true, err, ""
    }else{
        return 0, false, nil, "Data Already Exist."
    }
}


func GetAllMststate(page *entities.MststateEntity) (entities.MststateEntities, bool, error, string) {
    logger.Log.Println("In side Mststatemodel")
    t := entities.MststateEntities{}
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
    values, err1 := dataAccess.GetAllMststate(page, orgntype)
    if err1 != nil {
        return t, false, err1, "Something Went Wrong"
    }
    if page.Offset == 0 {
        total, err1 := dataAccess.GetMststateCount(page, orgntype)
        if err1 != nil {
            return t, false, err1, "Something Went Wrong"
        }
        t.Total = total.Total
        t.Values = values
    }
    t.Values = values
    return t, true, err, ""
}



func DeleteMststate(tz *entities.MststateEntity) (bool,error, string) {
    logger.Log.Println("In side Mststatemodel")
    lock.Lock()
    defer lock.Unlock()
    db, err := config.ConnectMySqlDbSingleton()
    if err != nil {
        logger.Log.Println("database connection failure", err)
        return false,err, "Something Went Wrong"
    }
    dataAccess := dao.DbConn{DB: db}
    err1 := dataAccess.DeleteMststate(tz)
    if err1 != nil {
        return false,err1, "Something Went Wrong"
    }
    return true,nil, ""
}


func UpdateMststate(tz *entities.MststateEntity) (bool,error, string) {
    logger.Log.Println("In side Mststatemodel")
    lock.Lock()
    defer lock.Unlock()
    db, err := config.ConnectMySqlDbSingleton()
    if err != nil {
        logger.Log.Println("database connection failure", err)
        return false,err, "Something Went Wrong"
    }
    dataAccess := dao.DbConn{DB: db}
    count,err :=dataAccess.CheckDuplicateMststate(tz)
    if err != nil {
        return false, err, "Something Went Wrong"
    }
    if count.Total == 0 {
         err := dataAccess.UpdateMststate(tz)
        if err != nil {
            return false, err, "Something Went Wrong"
        }
        return true, err, ""
    }else{
        return false, nil, "Data Already Exist."
    }
}


