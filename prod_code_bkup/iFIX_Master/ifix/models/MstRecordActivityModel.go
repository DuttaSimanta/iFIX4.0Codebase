package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"strconv"
	"strings"
)

func AddMstRecordActivityCopy(tz *entities.MstRecordActivityEntity) (int64, bool, error, string) {
	logger.Log.Println("In side AddMstRecordActivityCopymodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	k := 0
	var id int64
	id = 0
	tx, err := db.Begin()
	if err != nil {
		logger.Log.Println("database transaction connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	dataAccess1 := dao.TxConn{TX: tx}
	for i := 0; i < len(tz.Activitydesces); i++ {
		tz.Activitydesces[i] = strconv.Quote(tz.Activitydesces[i])
	}
	activity := strings.Join(tz.Activitydesces, ",")
	logger.Log.Println(activity)

	rows, err := dataAccess.GetRows(tz, activity)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			logger.Log.Print("AddMstRecordActivityCopy  Statement Rollback error", err)
			return 0, false, err, ""
		}
		return 0, false, err, "Something Went Wrong"
	}
	logger.Log.Print(tz.ToClientid)
	for i := 0; i < len(tz.ToMstorgnhirarchyids); i++ {

		for j := 0; j < len(rows); j++ {
			rows[j].Clientid = tz.ToClientid
			rows[j].Mstorgnhirarchyid = tz.ToMstorgnhirarchyids[i]
			logger.Log.Print("row", rows[j], tz)

			count, err := dataAccess.CheckDuplicateMstRecordActivity(&rows[j])
			if err != nil {
				err = tx.Rollback()
				if err != nil {
					logger.Log.Print("AddMstRecordActivityCopy  Statement Rollback error", err)
					return 0, false, err, ""
				}
				return 0, false, err, "Something Went Wrong"
			}
			if count.Total > 0 {
				k++
				logger.Log.Print("False")

			} else {
				id, err = dataAccess1.AddMstRecordActivityCopy(&rows[j])
				if err != nil {
					err1 := tx.Rollback()
					if err != nil {
						logger.Log.Print("AddMstRecordActivityCopy  Statement Rollback error", err1)
						return 0, false, err1, ""
					}
					return 0, false, err, "Something Went Wrong"
				}
			}
		}
	}

	if k < len(rows)*len(tz.ToMstorgnhirarchyids) {
		err = tx.Commit()
		if err != nil {
			logger.Log.Print("Banner  Statement Commit error", err)
			return 0, false, err, ""
		}
		return id, true, err, ""
	} else {
		err = tx.Commit()
		if err != nil {
			logger.Log.Print("Banner  Statement Commit error", err)
			return 0, false, err, ""
		}
		return 0, false, nil, "Data Already Exist."
	}

}

func AddMstRecordActivity(tz *entities.MstRecordActivityEntity) (int64, bool, error, string) {
	logger.Log.Println("In side MstRecordActivity model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	tz.Activitydesc = strings.TrimSpace(tz.Activitydesc)
	count1, err := dataAccess.CheckDuplicateMstRecordActivityTable(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if count1.Total != 0 {
		count, err := dataAccess.CheckDuplicateMstRecordActivity(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		if count.Total == 0 {
			id, err := dataAccess.AddMstRecordActivitycopyseq(tz)
			if err != nil {
				return 0, false, err, "Something Went Wrong"
			}
			return id, true, err, ""
		} else {
			return 0, false, nil, "Data Already Exist."
		}
	} else {
		seq, err1 := dataAccess.GetSeq(tz)
		if err1 != nil {
			return 0, false, err1, "Something Went Wrong"
		}
		if len(seq) == 0 {
			tz.Sequence = 1
		} else {
			tz.Sequence = seq[0].Sequence + 1
		}
		id, err := dataAccess.AddMstRecordActivity(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		return id, true, err, ""
	}

}

func GetAllMstRecordActivity(page *entities.MstRecordActivityEntity) (entities.MstRecordActivityEntities, bool, error, string) {
	logger.Log.Println("In side MstRecordActivity model")
	t := entities.MstRecordActivityEntities{}
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
	values, err1 := dataAccess.GetAllMstRecordActivity(page, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetMstRecordActivityCount(page, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

func DeleteMstRecordActivity(tz *entities.MstRecordActivityEntity) (bool, error, string) {
	logger.Log.Println("In side MstRecordActivity model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteMstRecordActivity(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

func UpdateMstRecordActivity(tz *entities.MstRecordActivityEntity) (bool, error, string) {
	logger.Log.Println("In side MstRecordActivity model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	// estra
	tz.Activitydesc = strings.TrimSpace(tz.Activitydesc)
	count1, err := dataAccess.CheckDuplicateMstRecordActivityTable(tz)
	if err != nil {
		return false, err, "Something Went Wrong"
	}
	if count1.Total != 0 {
		count, err := dataAccess.CheckDuplicateMstRecordActivity(tz)
		if err != nil {
			return false, err, "Something Went Wrong"
		}
		if count.Total == 0 {
			seq, err1 := dataAccess.GetSeqcopy(tz)
			if err1 != nil {
				return false, err1, "Something Went Wrong"
			}
			if len(seq) == 0 {
				tz.Sequence = 0
			} else {
				tz.Sequence = seq[0].Sequence
			}
			err := dataAccess.UpdateMstRecordActivitycopyseq(tz)
			if err != nil {
				return false, err, "Something Went Wrong"
			}
			return true, err, ""
		} else {
			return false, nil, "Data Already Exist."
		}
	} else {

		count, err := dataAccess.CheckDuplicateMstRecordActivity(tz)
		if err != nil {
			return false, err, "Something Went Wrong"
		}
		if count.Total == 0 {
			err := dataAccess.UpdateMstRecordActivitycopyseq(tz)
			if err != nil {
				return false, err, "Something Went Wrong"
			}
			return true, err, ""
		} else {
			return false, nil, "Data Already Exist."
		}
	}
}
func GetOrgWiseActivitydesc(page *entities.MstRecordActivityEntity) ([]entities.Activitydesces, bool, error, string) {
	logger.Log.Println("In side GetOrgWiseActivitydescmodel")
	t := []entities.Activitydesces{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	t, err1 := dataAccess.GetOrgWiseActivitydesc(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return t, true, err, ""
}
