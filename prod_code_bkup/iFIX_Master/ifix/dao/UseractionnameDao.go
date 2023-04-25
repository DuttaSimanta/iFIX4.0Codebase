package dao

import (
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var getuserwiseroleaction = "select actionid from mstclientroleuseraction where clientid=? and mstorgnhirarchyid=? and roleid in (select roleid from mapclientuserroleuser where userid=? and activeflg=1 and deleteflg=0) and userid=? and activeflg=1 and deleteflg=0"
var getrolewiseaction = "select actionid from mstclientroleaction where clientid=? and mstorgnhirarchyid=? and roleid in (select roleid from mapclientuserroleuser where userid=? and activeflg=1 and deleteflg=0) and activeflg=1 and deleteflg=0"

func (dbc DbConn) GetUserRolewiseActionname(page *entities.UserroleactionnameEntity) ([]int64, error) {
	logger.Log.Println("In side GelAllRecorddifferentiation")
	logger.Log.Println(getuserwiseroleaction)
	var ids []int64
	rows, err := dbc.DB.Query(getuserwiseroleaction, page.Clientid, page.Mstorgnhirarchyid, page.UserID, page.UserID)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllRecorddifferentiation Get Statement Prepare Error", err)
		return ids, err
	}
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			logger.Log.Println("Error in fetching data")
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func (dbc DbConn) GetRolewiseActionname(page *entities.UserroleactionnameEntity) ([]int64, error) {
	logger.Log.Println("In side GelAllRecorddifferentiation")
	logger.Log.Println(getrolewiseaction)
	var ids []int64
	rows, err := dbc.DB.Query(getrolewiseaction, page.Clientid, page.Mstorgnhirarchyid, page.UserID)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllRecorddifferentiation Get Statement Prepare Error", err)
		return ids, err
	}
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			logger.Log.Println("Error in fetching data")
		}
		ids = append(ids, id)
	}
	return ids, nil
}
