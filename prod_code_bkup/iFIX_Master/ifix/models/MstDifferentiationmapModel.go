package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func InsertDifferentiationModelsMap(tz *entities.MstDifferentiationmapEntity) (bool, error, string) {
	//db, err := config.ConnectMySqlDb()
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	//defer db.Close()
	dataAccess := dao.DbConn{DB: db}
	tx, err := db.Begin()
	if err != nil {
		logger.Log.Println("Transaction Failure...", err)
		// db.Close()
		return false, err, "Something Went Wrong"
	}
	var duplicatemsg string
	// Insert into first Differentiation Table...
	for i := 0; i < len(tz.ToorgnID); i++ {
		for k := 0; k < len(tz.DiffererntiationID); k++ {
			name, seq, err := dataAccess.GetDifferentiationname(tz.DiffererntiationID[k])
			if err != nil {
				logger.Log.Println("Transaction Failure...", err)
				tx.Rollback()
				// db.Close()
				return false, err, "Something Went Wrong"
			}

			count, err := dataAccess.DuplicateChecking(tz.ToclinentID, tz.ToorgnID[i], tz.DifferentiationtypeID, name)
			if err != nil {
				logger.Log.Println("Transaction Failure...", err)
				tx.Rollback()
				// db.Close()
				return false, err, "Something Went Wrong"
			}
			if count == 0 {
				insertedID, err := dao.InsertDifferentiationTBL(tx, tz.ToclinentID, tz.ToorgnID[i], tz.DifferentiationtypeID, name, seq)
				if err != nil {
					logger.Log.Println("Transaction Failure...", err)
					tx.Rollback()
					// db.Close()
					return false, err, "Something Went Wrong"
				}
				if insertedID > 0 {
					err := dao.InsertDifferentiationMAPTBL(tx, tz.FromclientID, tz.FromorgnID, tz.DifferentiationtypeID, tz.DiffererntiationID[k], tz.ToclinentID, tz.ToorgnID[i], tz.DifferentiationtypeID, insertedID)
					if err != nil {
						logger.Log.Println("Transaction Failure...", err)
						tx.Rollback()
						// db.Close()
						return false, err, "Something Went Wrong"
					}
				}
			} else {
				duplicatemsg = duplicatemsg + "{Status name :" + name + "}"
			}

		} // end 2nd for loop

	} // end 1st for loop

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		// db.Close()
		return false, err, "Something Went Wrong"
	}
	// db.Close()
	if len(duplicatemsg) > 0 {
		return true, nil, "Data Already exist. " + duplicatemsg
	} else {
		return true, nil, "Data Replicated Successfully Done.."
	}

}

func DeleteDifferentiationModelsMap(tz *entities.MstDifferentiationmapEntity) (bool, error, string) {
	//db, err := config.ConnectMySqlDb()
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	tx, err := db.Begin()
	if err != nil {
		logger.Log.Println("Transaction Failure...", err)
		// db.Close()
		return false, err, "Something Went Wrong"
	}
	err = dao.DeleteFromDifferentiation(tx, tz.ID)
	if err != nil {
		tx.Rollback()
		// db.Close()
		return false, err, "Something Went Wrong"
	}
	err = dao.DeleteFromDifferentiationMap(tx, tz.MapID)
	if err != nil {
		tx.Rollback()
		// db.Close()
		return false, err, "Something Went Wrong"
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		// db.Close()
		return false, err, "Something Went Wrong"
	}
	// db.Close()
	return true, nil, "Data Deleted Successfully Done.."
}

//GetAllDifferentiationMapDtls for implements business logic
func GetAllDifferentiationMapDtls(tz *entities.MstDifferentiationmapEntity) (entities.MstDifferentiationmaEntities, bool, error, string) {
	logger.Log.Println("In side model")
	t := entities.MstDifferentiationmaEntities{}
	//db, err := config.ConnectMySqlDb()
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	//defer db.Close()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetAllDifferentiationMapDtls(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if tz.Offset == 0 {
		total, err1 := dataAccess.GetAllDifferentiationMapDtlsCount(tz)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}
