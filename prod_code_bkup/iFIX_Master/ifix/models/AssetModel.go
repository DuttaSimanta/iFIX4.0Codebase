package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func InsertAsset(tz *entities.AssetEntity) (int64, bool, error, string) {
	logger.Log.Println("In side Assetmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateAsset(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		id, err := dataAccess.InsertAsset(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		return id, true, err, ""
	} else {
		return 0, false, nil, "Data Already Exist."
	}
}

func GetAllAsset(page *entities.AssetEntity) (entities.AssetEntities, bool, error, string) {
	logger.Log.Println("In side Assetmodel")
	t := entities.AssetEntities{}
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
	values, err1 := dataAccess.GetAllAsset(page, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetAssetCount(page, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

func DeleteAsset(tz *entities.AssetEntity) (bool, error, string) {
	logger.Log.Println("In side Assetmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteAsset(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

func UpdateAsset(tz *entities.AssetEntity) (bool, error, string) {
	logger.Log.Println("In side Assetmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateAsset(tz)
	if err != nil {
		return false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		err := dataAccess.UpdateAsset(tz)
		if err != nil {
			return false, err, "Something Went Wrong"
		}
		return true, err, ""
	} else {
		return false, nil, "Data Already Exist."
	}
}

func GetAssetBYType(tz *entities.AssetEntity) (entities.AssetEntitiesByType, bool, error, string) {
	logger.Log.Println("In side Assetmodel")
	t := entities.AssetEntitiesByType{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetAssetBYType(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}

	t.Values = values
	return t, true, err, ""
}

func GetAssetDiffVal(tz *entities.AssetEntity) (entities.AssetEntityDiffVals, bool, error, string) {
	logger.Log.Println("In side Assetmodel")
	t := entities.AssetEntityDiffVals{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetAssetDiffVal(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}

	t.Values = values
	return t, true, err, ""
}

func UpdateAssetDiffVal(tz *entities.AssetEntityDiffValUpdate) (bool, error, string) {
	logger.Log.Println("In side Assetmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err = dataAccess.DelAssetDiff(tz)
	if err != nil {
		return false, err, "Something Went Wrong"
	}

	err = dataAccess.InsertAssetDiff(tz)
	if err != nil {
		return false, err, "Something Went Wrong"
	}
	return true, err, ""

}

// GetClietWiseAsset function is used to get mappting with type
func GetClietWiseAsset(tz *entities.AssetEntity) (entities.AssetMapWithRecordTypes, bool, error, string) {
	logger.Log.Println("In side Assetmodel")
	t := entities.AssetMapWithRecordTypes{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	orgntype, err1 := dataAccess.GetOrgnType(tz.Clientid, tz.Mstorgnhirarchyid)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	values, err1 := dataAccess.GetClietWiseAsset(tz, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if tz.Offset == 0 {
		total, err1 := dataAccess.CountClietWiseAsset(tz, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

// GetAssetTypes function is used to get asset types
func GetAssetTypes(tz *entities.AssetEntity) ([]entities.Assettype, bool, error, string) {
	logger.Log.Println("In side Assetmodel")
	t := []entities.Assettype{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetAssetTypes(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}

// GetAssetAttributes function is used to get asset attributes
func GetAssetAttributes(tz *entities.AssetEntity) ([]entities.Assettype, bool, error, string) {
	logger.Log.Println("In side Assetmodel")
	t := []entities.Assettype{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetAssetAttributes(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}

// GetAssetByTypeNAtrrValue function is used to get assets with attributes and its value
func GetAssetByTypeNAtrrValue(tz *entities.AssetSearchEntity) (entities.AssetSearchResEntity, bool, error, string) {
	logger.Log.Println("In side Assetmodel")
	t := entities.AssetSearchResEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetAssetAttributesbyTypeId(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	t.AssetAttributes = values
	values1, err1 := dataAccess.GetAssetByTypeNAtrrValue(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	for i, v := range values1 {
		astValues, err1 := dataAccess.GetAssetDiffValByID(tz, v.ID)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		values1[i].Attributes = astValues
	}
	t.AssetValues = values1
	return t, true, err, ""
}
