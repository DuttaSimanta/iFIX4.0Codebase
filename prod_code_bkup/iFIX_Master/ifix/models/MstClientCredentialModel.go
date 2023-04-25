package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func InsertMstClientCredential(tz *entities.MstClientCredentialEntity) (int64, bool, error, string) {
	logger.Log.Println("In side MstClientCredentialmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	if tz.DefaultConfig == 1 {
		values, errstring, err := dataAccess.GetData(tz)
		if err != nil {
			if errstring == "NOROWS" {
				return 0, false, err, "No client is there for this data"

			}
			return 0, false, err, "Something Went Wrong"

		}
		logger.Log.Println(values)
		tz.Credentialtypeid = values.Credentialtypeid
		tz.CredentialAccount = values.CredentialAccount
		tz.CredentialPassword = values.CredentialPassword
		tz.CredentialKey = values.CredentialKey
		tz.CredentialEndPoint = values.CredentialEndPoint

	}
	count, err := dataAccess.CheckDuplicateMstClientCredential(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		id, err := dataAccess.InsertMstClientCredential(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		return id, true, err, ""
	} else {
		return 0, false, nil, "Data Already Exist."
	}
}

func GetAllMstClientCredential(page *entities.MstClientCredentialEntity) (entities.MstClientCredentialEntities, bool, error, string) {
	logger.Log.Println("In side MstClientCredentialmodel")
	t := entities.MstClientCredentialEntities{}
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
	values, err1 := dataAccess.GetAllMstClientCredential(page, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetMstClientCredentialCount(page, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

func DeleteMstClientCredential(tz *entities.MstClientCredentialEntity) (bool, error, string) {
	logger.Log.Println("In side MstClientCredentialmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteMstClientCredential(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

func UpdateMstClientCredential(tz *entities.MstClientCredentialEntity) (bool, error, string) {
	logger.Log.Println("In side MstClientCredentialmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateMstClientCredential(tz)
	if err != nil {
		return false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		err := dataAccess.UpdateMstClientCredential(tz)
		if err != nil {
			return false, err, "Something Went Wrong"
		}
		return true, err, ""
	} else {
		return false, nil, "Data Already Exist."
	}
}
