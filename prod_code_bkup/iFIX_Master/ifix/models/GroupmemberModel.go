package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func InsertGroupmember(tz *entities.GroupmemberEntity) (int64, bool, error, string) {
	logger.Log.Println("In side Groupmembermodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateGroupmember(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		id, err := dataAccess.InsertGroupmember(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		return id, true, err, ""
	} else {
		return 0, false, nil, "Data Already Exist."
	}
}

//
//func GetAllGroupmember(page *entities.GroupmemberEntity) (entities.GroupmemberEntities, bool, error, string) {
//    logger.Log.Println("In side Groupmembermodel")
//    t := entities.GroupmemberEntities{}
//    db, err := config.ConnectMySqlDb()
//    defer db.Close()
//    if err != nil {
//        logger.Log.Println("database connection failure", err)
//         return t, false, err, "Something Went Wrong"
//    }
//    dataAccess := dao.DbConn{DB: db}
//    values, err1 := dataAccess.GetAllGroupmember(page)
//    if err1 != nil {
//        return t, false, err1, "Something Went Wrong"
//    }
//    if page.Offset == 0{
//        total, err1 := dataAccess.GetGroupmemberCount(page)
//    if err1 != nil {
//        return t, false, err1, "Something Went Wrong"
//    }
//    t.Total=total.Total
//    t.Values=values
//    }
//    t.Values=values
//    return t, true, err, ""
//}

//SearchUserByGroupId for implements business logic
func SearchAnalystOrgWise(tz *entities.GroupmemberEntity) ([]entities.MstUserSearchEntity, bool, error, string) {
	logger.Log.Println("In side model")
	t := []entities.MstUserSearchEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.SearchAnalystOrgWise(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}
func SearchUserByGroupId(tz *entities.GroupmemberEntity) ([]entities.MstUserSearchEntity, bool, error, string) {
	logger.Log.Println("In side model")
	t := []entities.MstUserSearchEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.SearchUserByGroupId(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}
func Searchuserdetailsbygroupid(tz *entities.GroupmemberEntity) ([]entities.MstUserSearchEntity, bool, error, string) {
	logger.Log.Println("In side model")
	//t := []entities.MstUserSearchEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return nil, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Searchuserdetailsbygroupid(tz)
	if err1 != nil {
		return nil, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}
func Groupbyuserwise(tz *entities.GroupmemberEntity) ([]entities.ClientsupportgroupsingleEntity, bool, error, string) {
	logger.Log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return nil, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Groupbyuserwise(tz)
	if err1 != nil {
		return nil, false, err1, "Something Went Wrong"
	}
	if len(values) > 0 {
		user := entities.UserEntity{}
		user.Clientid = tz.Clientid
		user.Mstorgnhirarchyid = tz.Mstorgnhirarchyid
		defaultgrp, err3 := dataAccess.Getdefaultgroupbyid(&user, tz.Refuserid)
		if err3 != nil {
			return nil, false, err3, "Something Went Wrong"
		}
		if len(defaultgrp) > 0 {
			//for _, group := range values {
			for i := 0; i < len(values); i++ {
				//for _, defgrp := range defaultgrp {
				for j :=0;j<len(defaultgrp);j++{
					if values[i].Id == defaultgrp[j].Deafultgroup{
						values[i].Defaultgroup=1
						break
					}
				}
			}
		}
	}
	return values, true, err, ""
}
func Workflowgroupbyuserwise(tz *entities.GroupmemberEntity) ([]entities.ClientsupportgroupsingleEntity, bool, error, string) {
	logger.Log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return nil, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Workflowgroupbyuserwise(tz)
	if err1 != nil {
		return nil, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}

func DeleteGroupmember(tz *entities.GroupmemberEntity) (bool, error, string) {
	logger.Log.Println("In side Groupmembermodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteGroupmember(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

func UpdateGroupmember(tz *entities.GroupmemberEntity) (bool, error, string) {
	logger.Log.Println("In side Groupmembermodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateGroupmember(tz)
	if err != nil {
		return false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		err := dataAccess.UpdateGroupmember(tz)
		if err != nil {
			return false, err, "Something Went Wrong"
		}
		return true, err, ""
	} else {
		return false, nil, "Data Already Exist."
	}
}

/*func AddGroupmember(tz *entities.GroupmemberEntity) (int64, bool, error, string) {
    logger.Log.Println("In side Groupmembermodel")
    db, err := config.ConnectMySqlDb()
    defer db.Close()
    if err != nil {
        logger.Log.Println("database connection failure", err)
        return 0, false, err, "Something Went Wrong"
    }
    dataAccess := dao.DbConn{DB: db}
    count1,err :=dataAccess.CheckClientuser(tz)
    if err != nil {
        return 0, false, err, "Something Went Wrong"
    }
    if count1.Total>0{
        count,err :=dataAccess.CheckDuplicateGroupmember(tz)
        if err != nil {
            return 0, false, err, "Something Went Wrong"
        }
        if count.Total == 0 {
            id, err := dataAccess.InsertGroupmember(tz)
            if err != nil {
                return 0, false, err, "Something Went Wrong"
            }
            return id, true, err, ""
        }else{
            return 0, false, nil, "Data Already Exist."
        }
    }else{
        return 0, false, nil, "No clientuser matches with this data"

    }
}*/
func AddGroupmember(tz *entities.GroupmemberEntity) (int64, bool, error, string) {
	logger.Log.Println("In side Groupmembermodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	n := 0
	m := 0
	var id int64
	id = 0

	/* Starting Transaction*/
	//tx, err := db.Begin()
	if err != nil {
		logger.Log.Println("database transaction connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	// dataAccess1 := dao.TxConn{TX: tx}
	for i := 0; i < len(tz.ToMstorgnhirarchyid); i++ {
		tx, err := db.Begin()
		dataAccess1 := dao.TxConn{TX: tx}
		count, err := dataAccess.CheckClientuser(tz, i)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		if count.Total > 0 {
			for j := 0; j < len(tz.Userids); j++ {
				//dataAccess1 := dao.TxConn{TX: tx}
				count1, err := dataAccess.CheckDuplicateGrpmember(tz, i, j)
				if err != nil {
					return 0, false, err, "Something Went Wrong"
				}
				/*if count1.Total>0{
				   n++
				}*/
				if count1.Total == 0 {
					n++
					id, err = dataAccess1.InsertGrpmember(tz, i, j)
					if err != nil {
						/* err = tx.Commit()
						   if err != nil {
						       //log.Print("MoveWorkflow  Statement Commit error", err)
						       logger.Log.Print("Banner  Statement Commit error", err)
						       return 0, false, err, ""
						   }*/
						return 0, false, err, "Something Went Wrong"
					}
				}
			}
			/*err = tx.Commit()
			  if err != nil {
			      //log.Print("MoveWorkflow  Statement Commit error", err)
			      logger.Log.Print("Banner  Statement Commit error", err)
			      return 0, false, err, ""
			  }*/
		} else {
			m++
		}
		err = tx.Commit()
		if err != nil {
			//log.Print("MoveWorkflow  Statement Commit error", err)
			logger.Log.Print("Banner  Statement Commit error", err)
			return 0, false, err, ""
		}
	}
	if m == len(tz.ToMstorgnhirarchyid) && n == 0 {
		return 0, false, nil, "No clientuser matches with this data"
	} else if n == 0 { //len(tz.ToMstorgnhirarchyid)*len(tz.Userids){
		return 0, false, nil, "Data Already Exist."
	} else {
		return id, true, err, ""
	}

}

func GetAllGrpmember(page *entities.GroupmemberEntity) (entities.GroupmemberEntities, bool, error, string) {
	logger.Log.Println("In side Groupmembermodel")
	t := entities.GroupmemberEntities{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetAllGrpmember(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetGrpmemberCount(page)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}
func GetUserByGroupId(tz *entities.GroupmemberEntity) ([]entities.MstUserSearchEntity, bool, error, string) {
	logger.Log.Println("In side model")
	t := []entities.MstUserSearchEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetUserByGroupId(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}
func GetAllGroupmember(page *entities.GroupmemberEntity) (entities.GroupmemberEntities, bool, error, string) {
	logger.Log.Println("In side Groupmembermodel")
	t := entities.GroupmemberEntities{}
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
	values, err1 := dataAccess.GetAllGroupmember(page, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetGroupmemberCount(page, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, nil, ""
}
