package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func InsertMstrecordterms(tz *entities.MstrecordtermsEntity) (int64, bool, error, string) {
	logger.Log.Println("In side Mstrecordtermsmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateMstrecordterms(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}

	if count.Total == 0 {
		if tz.Termseq == 0 {
			seq, err := dataAccess.GetLastSeqFromterms(tz)
			if err != nil {
				return 0, false, err, "Something Went Wrong"
			}
			if len(seq) == 0 {
				tz.Termseq = 1
			} else {
				tz.Termseq = seq[0].Termseq + 1
			}
		} else {
			count1, err := dataAccess.CheckDuplicatetermseq(tz)
			if err != nil {
				return 0, false, err, "Something Went Wrong"
			}
			if count1.Total > 0 {
				return 0, false, nil, "Cannot create state with same sequence"
			}
		}
		id, err := dataAccess.InsertMstrecordterms(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		return id, true, err, ""
	} else {
		return 0, false, nil, "Data Already Exist."
	}
}

func GetAllMstrecordterms(page *entities.MstrecordtermsEntity) (entities.MstrecordtermsEntities, bool, error, string) {
	logger.Log.Println("In side Mstrecordtermsmodel")
	t := entities.MstrecordtermsEntities{}
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
	values, err1 := dataAccess.GetAllMstrecordterms(page, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetMstrecordtermsCount(page, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

func GetListMstrecordterms(page *entities.MstrecordtermsEntity) ([]entities.TermsEntity, bool, error, string) {
	logger.Log.Println("In side Mstrecordtermsmodel")
	t := []entities.TermsEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	t, err1 := dataAccess.GetListMstrecordterms(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return t, true, err, ""
}

func DeleteMstrecordterms(tz *entities.MstrecordtermsEntity) (bool, error, string) {
	logger.Log.Println("In side Mstrecordtermsmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteMstrecordterms(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

func UpdateMstrecordterms(tz *entities.MstrecordtermsEntity) (bool, error, string) {
	logger.Log.Println("In side Mstrecordtermsmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateMstrecordterms(tz)
	if err != nil {
		return false, err, "Something Went Wrong"
	}
	termsvalcount, termsErr := dataAccess.CheckDuplicatetermval(tz)
	if termsErr != nil {
		return false, termsErr, "Something Went Wrong"
	}
	if count.Total == 0 || (count.Total == 1 && termsvalcount.Total == 0) {
		err := dataAccess.UpdateMstrecordterms(tz)
		if err != nil {
			return false, err, "Something Went Wrong"
		}
		return true, err, ""
	} else {
		return false, nil, "Data Already Exist."
	}
}
