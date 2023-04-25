package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var deleteOpenTicket = "DELETE FROM mstuserrecord  WHERE id = ? "
var getOpenTicket = "SELECT a.id as Id, a.groupid as Groupid, c.name as Groupname, a.recordid as Recordid, d.code as Reccordcode, a.userid as Userid, b.name as Username  FROM mstuserrecord a ,mstclientuser b,mstsupportgrp c, trnrecord d WHERE a.groupid = c.id and a.userid = b.id and a.recordid = d.id and a.opendate >?  and a.activeflg=1 and a.deleteflg =0 ORDER BY a.id DESC LIMIT ?,?"

func (dbc DbConn) GetOpenTicketCount(tz *entities.OpenTicketEntity) (entities.OpenTicketEntities, error) {
	logger.Log.Println("In side GetBannerCount")
	value := entities.OpenTicketEntities{}
	var getOpentTicketcount string
	var params []interface{}

	getOpentTicketcount = "SELECT count(a.id) as total FROM mstuserrecord a ,mstclientuser b,mstsupportgrp c, trnrecord d WHERE  a.groupid = c.id and a.userid = b.id and a.recordid = d.id and a.opendate > ?  and a.activeflg=1 and a.deleteflg =0;"
	params = append(params, tz.Opendate)

	err := dbc.DB.QueryRow(getOpentTicketcount, params...).Scan(&value.Total)
	// err := dbc.DB.QueryRow(getBannercount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetBannerCount Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) GetOpenTicket(tz *entities.OpenTicketEntity) ([]entities.OpenTicketEntity, error) {
	logger.Log.Println("In side dao GetAllOpenTicket")
	values := []entities.OpenTicketEntity{}

	var params []interface{}

	params = append(params, tz.Opendate)
	params = append(params, tz.Offset)
	params = append(params, tz.Limit)
	rows, err := dbc.DB.Query(getOpenTicket, params...)
	// rows, err := dbc.DB.Query(getBanner, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllOpenTicket Get Statement Prepare Error", err)
		return values, err
	}

	for rows.Next() {
		value := entities.OpenTicketEntity{}
		rows.Scan(&value.Id, &value.Groupid, &value.Groupname, &value.Recordid, &value.Reccordcode, &value.Userid, &value.Username)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) DeleteOpenTicket(tz *entities.OpenTicketEntity) error {
	logger.Log.Println("In side DeleteOpenTicket")
	stmt, err := dbc.DB.Prepare(deleteOpenTicket)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteOpenTicket Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteOpenTicket Execute Statement  Error", err)
		return err
	}
	return nil
}
