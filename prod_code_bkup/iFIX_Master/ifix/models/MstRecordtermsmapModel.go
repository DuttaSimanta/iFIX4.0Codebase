package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func InsertRecordtermsMap(tz *entities.MstRecordtermsmapEntity) (bool, error, string) {
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	logger.Log.Println(db)
	dataAccess := dao.DbConn{DB: db}
	tx, err := db.Begin()
	if err != nil {
		logger.Log.Println("Transaction creation error in AddRecordModelAction", err)
		return false, err, "Something Went Wrong"
	}
	//defer db.Close()
	hashmap, err := dataAccess.Termsequance(tz.FromclientID, tz.FromorgnID)
	if err != nil {
		//db.Close()
		return false, err, "Something Went Wrong"
	}
	if len(hashmap) == 0 {
		//db.Close()
		return false, err, "Something Went Wrong"
	}

	for i := 0; i < len(tz.ToorgnID); i++ {
		for k := 0; k < len(tz.TermsSeq); k++ {
			checkduplicate, err := dataAccess.GetduplicateCheck(tz.ToclinentID, tz.ToorgnID[i], tz.TermsSeq[k])
			if err != nil {
				//db.Close()
				return false, err, "Something Went Wrong"
			}
			if checkduplicate == 0 {
				termdtls, err := dataAccess.GetTermdtls(tz.FromclientID, tz.FromorgnID, tz.TermsSeq[k])
				if err != nil {
					//db.Close()
					return false, err, "Something Went Wrong"
				}

				if termdtls.ID > 0 {
					id, err := dao.InsertRecordterms(tx, tz.ToclinentID, tz.ToorgnID[i], termdtls.Termname, termdtls.Termvalue, termdtls.Termtype, tz.TermsSeq[k])
					if err != nil {
						//db.Close()
						return false, err, "Something Went Wrong"
					}
					if id > 0 {
						values, err := dataAccess.Getstatetems(tz.FromclientID, tz.FromorgnID, hashmap[tz.TermsSeq[k]])
						if err != nil {
							//db.Close()
							return false, err, "Something Went Wrong"
						}
						for p := 0; p < len(values); p++ {
							diffID, err := dataAccess.GetDiffID(tz.FromclientID, tz.FromorgnID, tz.ToclinentID, tz.ToorgnID[i], values[p].RecorddifftypeID, values[p].RecorddiffID)
							if err != nil {
								//db.Close()
								return false, err, "Something Went Wrong"
							}

							_, err = dao.InsertRecordstateTerms(tx, tz.ToclinentID, tz.ToorgnID[i], values[p].RecorddifftypeID, diffID, id, values[p].Recordtermvalue, values[p].Iscompulsery)
							if err != nil {
								//db.Close()
								return false, err, "Something Went Wrong"
							}
						}

						grpvalues, err := dataAccess.Getgrpids(tz.FromclientID, tz.FromorgnID, hashmap[tz.TermsSeq[k]])
						if err != nil {
							//db.Close()
							return false, err, "Something Went Wrong"
						}

						for j := 0; j < len(grpvalues); j++ {
							_, err = dao.InsertSupportgrpTerms(tx, tz.ToclinentID, tz.ToorgnID[i], grpvalues[j], id)
							if err != nil {
								//db.Close()
								return false, err, "Something Went Wrong"
							}

						}
					}

				}

			}

		} // termseq for loop end here...
	} // Orgn for loop end here...

	err = tx.Commit()
	if err != nil {
		logger.Log.Println(err)
		tx.Rollback()
		//db.Close()
		return false, err, "Something Went Wrong"
	}
	return true, nil, "Data Replicated Successfully Done.."
}
