package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
)

var getticketcustomer = "SELECT  distinct a.mstorgnhirarchyid as Mstorgnhirarchyid,b.name as Mstorgnhirarchyname FROM mstgroupmember a,mstorgnhierarchy b WHERE a.clientid=? AND a.userid=? AND a.activeflg=1 AND a.deleteflg=0 AND a.mstorgnhirarchyid=b.id;"
var getticketcustomercount = "SELECT count(distinct a.mstorgnhirarchyid) total from mstgroupmember a,mstorgnhierarchy b WHERE a.clientid=? AND a.userid=? AND a.activeflg=1 AND a.deleteflg=0 AND a.mstorgnhirarchyid=b.id"

func (mdao DbConn) GetAllTicketCustomer(page *entities.TicketCustomerEntity) ([]entities.TicketCustomerEntity, error) {
	logger.Log.Println("In side dao")
	values := []entities.TicketCustomerEntity{}
	rows, err := mdao.DB.Query(getticketcustomer, page.Clientid, page.Refuserid)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllticketcustomer Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.TicketCustomerEntity{}
		rows.Scan(&value.Mstorgnhirarchyid, &value.Mstorgnhirarchyname)

		values = append(values, value)
	}
	return values, nil
}

func (mdao DbConn) GetTicketCustomerCount(tz *entities.TicketCustomerEntity) (entities.ModuleEntities, error) {
	logger.Log.Println("In side dao")
	value := entities.ModuleEntities{}
	err := mdao.DB.QueryRow(getticketcustomercount, tz.Clientid, tz.Refuserid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		log.Print("GetticketcustomerCount Get Statement Prepare Error", err)
		return value, err
	}
}
