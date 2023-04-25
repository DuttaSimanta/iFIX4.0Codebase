package dao

import (
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
)

func (dbc DbConn) Getallserver()([]entities.ProcessMonitorEntity,error){
	values :=[]entities.ProcessMonitorEntity{}

	var getallserver="SELECT distinct serverno from mstservermonitorurl where deleteflg=0"
	rows, err := dbc.DB.Query(getallserver)

	if err != nil {
		logger.Log.Println("Getallserver Get Statement Prepare Error", err)
		log.Println("Getallserver Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		value := entities.ProcessMonitorEntity{}
		rows.Scan(&value.Serverno)
		values = append(values, value)
	}
	return values, nil
}
func (dbc DbConn) Getprocessbyserver(serverno string)([]entities.ProcessMonitorEntity,error){
	values :=[]entities.ProcessMonitorEntity{}

	var processbyserver="SELECT id,process from mstservermonitorurl where serverno=? and deleteflg=0"
	rows, err := dbc.DB.Query(processbyserver,serverno)

	if err != nil {
		logger.Log.Println("Getprocessbyserver Get Statement Prepare Error", err)
		log.Println("Getprocessbyserver Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		value := entities.ProcessMonitorEntity{}
		rows.Scan(&value.Id,&value.Process)
		values = append(values, value)
	}
	return values, nil
}
func (dbc DbConn) Getstatusbyprocess(id int64)([]entities.ProcessMonitorEntity,error){
	log.Println("inside")
	values :=[]entities.ProcessMonitorEntity{}

	var getallserver="SELECT status,processid from mstservermonitorstatus  where  processid =?  and deleteflg=0 order by id desc limit 5"
	rows, err := dbc.DB.Query(getallserver,id)

	if err != nil {
		logger.Log.Println("Getstatusbyprocess Get Statement Prepare Error", err)
		log.Println("Getstatusbyprocess Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		value := entities.ProcessMonitorEntity{}
		rows.Scan(&value.Status,&value.Id)
		values = append(values, value)
	}
	return values, nil
}
