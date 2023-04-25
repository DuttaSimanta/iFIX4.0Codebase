package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"log"
)

func GetRecordDiffType() ([]entities.MstrecorddifftypeEntity, bool, error, string) {
	log.Println("In side model")
	t := []entities.MstrecorddifftypeEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetRecordDiffType()
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	//t=append(values[:0],values[1:]...)
	return values, true, err, ""
}
func GetRecordByDiffType(tz *entities.RecordDiffEntity) ([]entities.MstrecorddifftypeEntity, bool, error, string) {
	log.Println("In side model")
	t := []entities.MstrecorddifftypeEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetRecordByDiffType(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}

	return values, true, err, ""
}
func InsertRecordDiff(tz *entities.RecordDiffEntity) (int64, bool, error, string) {
	log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateRecordDiff(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		if tz.Seqno == -1 {
			seq, err := dataAccess.GetLastSeqFromRecordDiff(tz)
			if err != nil {
				return 0, false, err, "Something Went Wrong"
			}
			if len(seq) == 0 {
				tz.Seqno = 1
			} else {
				tz.Seqno = seq[0].Seqno + 1
			}
		}
		id, err := dataAccess.InsertRecordDiff(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}

		return id, true, err, ""
	} else {
		return 0, false, nil, "Name Already Exist."
	}
}

/**
Get ticket type/status for a certain client
*/
func GetAllRecordDiff(tz *entities.RecordDiffEntity) (entities.RecordDiffEntities, bool, error, string) {
	log.Println("In side model")
	t := entities.RecordDiffEntities{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetAllRecordDiff(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if tz.Offset == 0 {
		total, err1 := dataAccess.GetRecordDiffCount(tz)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

/**
Get ticket type/status for a all client
*/
func GetRecordDiff(tz *entities.RecordDiffEntity) (entities.RecordDiffEntities, bool, error, string) {
	log.Println("In side model")
	t := entities.RecordDiffEntities{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetRecordDiff(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if tz.Offset == 0 {
		total, err1 := dataAccess.GetAllRecordDiffCount()
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

/**
Get ticket type/status for a certain client and org
*/
func GetRecordDiffByOrg(tz *entities.RecordDiffEntity) (entities.RecordDiffEntities, bool, error, string) {
	log.Println("In side model")
	t := entities.RecordDiffEntities{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	orgntype, err1 := dataAccess.GetOrgnType(tz.Clientid, tz.Mstorgnhirarchyid)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	values, err1 := dataAccess.GetRecordDiffByOrg(tz, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if tz.Offset == 0 {
		total, err1 := dataAccess.GetRecordDiffCountByOrg(tz, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""

}

func GetAssetRecordDiffByOrg(tz *entities.RecordDiffEntity) (entities.RecordDiffEntities, bool, error, string) {
	log.Println("In side model")
	t := entities.RecordDiffEntities{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	orgntype, err1 := dataAccess.GetOrgnType(tz.Clientid, tz.Mstorgnhirarchyid)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	values, err1 := dataAccess.GetAssetRecordDiffByOrg(tz, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if tz.Offset == 0 {
		total, err1 := dataAccess.GetRecordDiffCountByOrg(tz, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""

}

/**
Update a diif record for a certain client
*/
func UpdateRecordDiff(tz *entities.RecordDiffEntity) (bool, error, string) {
	log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.UpdateRecordDiff(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

func UpdateAssetRecordDiff(tz *entities.RecordDiffEntity) (bool, error, string) {
	log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateRecordDiff(tz)
	if err != nil {
		return false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		err1 := dataAccess.UpdateAssetRecordDiff(tz)
		if err1 != nil {
			return false, err1, "Something Went Wrong"
		}
		return true, nil, ""
	} else {
		return false, nil, "Name Already Exist."
	}

}

/**
Delete a diff record for a certain client
*/
func DeleteRecordDiff(tz *entities.RecordDiffEntity) (bool, error, string) {
	log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteRecordDiff(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

/**
Get category level by client and orgid
*/
func GetCategoryLevel(tz *entities.RecordDiffEntity) ([]entities.MstrecorddifftypeEntity, bool, error, string) {
	log.Println("In side model")
	t := []entities.MstrecorddifftypeEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetCategoryLevel(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}

	return values, true, err, ""
}
func GetCategoriesLevel(tz *entities.RecordDiffEntity) ([]entities.MstrecorddifftypeEntity, bool, error, string) {
	log.Println("In side model")
	t := []entities.MstrecorddifftypeEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetCategoriesLevel(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}

	return values, true, err, ""
}

/**
Get All category level by client and orgid
*/
func GetAllCategoryLevel(tz *entities.RecordDiffEntity) ([]entities.MstrecorddifftypeEntity, bool, error, string) {
	log.Println("In side model")
	t := []entities.MstrecorddifftypeEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetAllCategoryLevel(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}

	return values, true, err, ""
}

/**
Get diff type with category level by client and orgid
*/
func GetAllRecordDiffTypeByClient(tz *entities.RecordDiffEntity) ([]entities.MstrecorddifftypeEntity, bool, error, string) {
	log.Println("In side model")
	t := []entities.MstrecorddifftypeEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	label, err1 := dataAccess.GetCategoryLevel(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	values, err1 := dataAccess.GetRecordDiffType()
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}

	log.Println("len:", len(label))
	for i := 0; i < len(label); i++ {
		log.Println("label:", label[i].Typename)
		t = append(t, label[i])
	}
	for i := 1; i < len(values); i++ {
		t = append(t, values[i])
	}

	return t, true, err, ""
}
