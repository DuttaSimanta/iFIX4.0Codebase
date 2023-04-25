package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

//AddModuleClients for implements business logic
func AddModuleClients(tz *entities.MstModuleClientEntity) (int64, bool, error, string) {
	logger.Log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong.."
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateModuleCient(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong...."
	}
	if count.Total > 0 {
		return 0, false, nil, "Module Already Mapped."
	}

	id, err := dataAccess.InsertModuleClientData(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong....."
	}
	seq, err := dataAccess.GetLastSeqFromMenu(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong....."
	}
	modent := entities.ModuleEntity{}
	modent.Id = tz.ModuleID
	mod, err := dataAccess.Getmodulename(&modent)
	if err != nil {
		return 0, false, err, "Something Went Wrong....."
	}

	menu := entities.MenuEntity{}
	menu.ClientID = tz.ClientID
	menu.MstorgnhirarchyID = tz.Mstorgnhirarchyid
	menu.Moduleid = tz.ModuleID
	menu.Menudesc = mod[0].Modulename
	menu.Parentmenuid = 0
	if len(seq) == 0 {
		menu.Sequence_no = 1
	} else {
		menu.Sequence_no = seq[0].Sequence_no + 1
	}
	_, _, err, _ = InsertMenu(&menu,db)
	if err != nil {
		return 0, false, err, "Something Went Wrong....."
	}
	urlentity := entities.UrlEntity{}
	basedetails, err := dataAccess.GetBaseOrgDetails()
	if err != nil {
		return 0, false, err, "Something Went Wrong....."
	}

	urlentity.Clientid = basedetails[0].ClientID
	urlentity.Mstorgnhirarchyid = basedetails[0].ID
	urlentity.Moduleid = tz.ModuleID
	urls, err := dataAccess.Geturlmodulewise(&urlentity)
	if len(urls) == 0 {
		return id, true, err, "No Url is mapped with module.You can manually map url "
	} else {
		urlentity.Clientid = tz.ClientID
		urlentity.Mstorgnhirarchyid = tz.Mstorgnhirarchyid
		for i := 0; i < len(urls); i++ {
			urlentity.Id = urls[i].Id
			dataAccess.InsertIntoModuleUrl(&urlentity)
		}
		return id, true, err, ""
	}

}

//DeleteModuleClients for implements business logic
func DeleteModuleClients(tz *entities.MstModuleClientEntity) (bool, error, string) {
	logger.Log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteClientModuleData(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

//UpdateModuleClients for implements business logic
func UpdateModuleClients(tz *entities.MstModuleClientEntity) (bool, error, string) {
	logger.Log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.UpdateClientModuleData(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

//GetAllModuleClients for implements business logic
func GetAllModuleClients(tz *entities.MstModuleClientEntity) (entities.MstModuleClientEntities, bool, error, string) {
	logger.Log.Println("In side model")
	t := entities.MstModuleClientEntities{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetAllModuleClients(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if tz.Offset == 0 {
		total, err1 := dataAccess.GetClientModuleCount(tz)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

//GetModuleByOrgId for implements business logic
func GetModuleByOrgId(tz *entities.MstModuleClientEntity) ([]entities.MstModuleByClientEntity, bool, error, string) {
	logger.Log.Println("In side model")
	t := []entities.MstModuleByClientEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetModuleByOrgId(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}

	return values, true, err, ""
}
