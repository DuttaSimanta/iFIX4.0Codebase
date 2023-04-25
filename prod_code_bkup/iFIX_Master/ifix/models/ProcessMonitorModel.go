package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
)

func Getallserver() ([]entities.ProcessMonitorServerEntity, bool, error, string) {

	var t = []entities.ProcessMonitorServerEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		log.Println("database connection failure", err)
		return nil, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Getallserver()
	if err1 != nil {
		return nil, false, err1, "Something Went Wrong"
	}
	for _, value := range values {
		details, err1 := dataAccess.Getprocessbyserver(value.Serverno)
		if err1 != nil {
			return nil, false, err1, "Something Went Wrong"
		}
		var t1 = entities.ProcessMonitorServerEntity{}
		t1.Serverno = value.Serverno
		t1.Processlist = details
		t = append(t, t1)
	}
	return t, true, err, ""
}
func Getprocessbyserver(tz *entities.ProcessMonitorEntity) ([]entities.ProcessMonitorEntity, bool, error, string) {

	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		log.Println("database connection failure", err)
		return nil, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Getprocessbyserver(tz.Serverno)
	if err1 != nil {
		return nil, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}
func Getstatusbyprocess(tz *entities.ProcessMonitorEntity) ([]entities.ProcessMonitorEntity, bool, error, string) {

	//lock.Lock()
	//defer lock.Unlock()
	//db, err := config.ConnectMySqlDbSingleton()
	//if err != nil {
	//	logger.Log.Println("database connection failure", err)
	//	log.Println("database connection failure", err)
	//	return nil, false, err, "Something Went Wrong"
	//}
	var values = []entities.ProcessMonitorEntity{}
	resChanel := make(chan entities.ProcessMonitorEntity)
	errChannel := make(chan error)
	for _, process := range tz.Processlist {
		go getlateststatusbyprocess(process, resChanel, errChannel)
		//err, val := <-errChannel, <-resChanel
		//log.Println(err)
		//if err != nil {
		//	return nil, false, err, "Something Went Wrong"
		//}
		values = append(values, <-resChanel)
	}
	//log.Println(values)
	return values, true, nil, ""
}
func getlateststatusbyprocess(id int64, resChanel chan entities.ProcessMonitorEntity, errChannel chan error) {
	t := entities.ProcessMonitorEntity{}
	log.Println("---->", id)
	db, err := config.ConnectMySqlDb()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		log.Println("database connection failure", err)
		errChannel <- err
		//return t,  err
	}
	defer db.Close()
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Getstatusbyprocess(id)
	if err1 != nil {
		//return t,  err
		errChannel <- err
	}
	var isStop = true
	for _, value := range values {
		if value.Status == "true" {
			isStop = false
			break
		}
	}
	t.Id=id
	if isStop{
		t.Status="false"
	}else{
		t.Status="true"
	}
	log.Println("sending value for : ", id)
	resChanel <- t
}
