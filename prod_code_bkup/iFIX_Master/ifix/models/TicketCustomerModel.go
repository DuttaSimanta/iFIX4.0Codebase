package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func GetAllTicketCustomer(page *entities.TicketCustomerEntity) (entities.TicketCustomerEntities, bool, error, string) {
	logger.Log.Println("In side model")
	t := entities.TicketCustomerEntities{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetAllTicketCustomer(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	//if page.Offset == 0{
	//total, err1 := dataAccess.GetTicketCustomerCount(page)
	//if err1 != nil {
	//	return t, false, err1, "Something Went Wrong"
	//}
	//t.Total = total.Total
	//t.Values = values
	//}
	t.Values = values
	return t, true, err1, ""
}
