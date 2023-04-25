package dao

import (
	"database/sql"
	"src/entities"
	"src/logger"
)

var getDownloadList = "SELECT id as Id, userid as Refuserid, originalfilename as Originalfilename, uploadedfilename as Uploadedfilename FROM mstreportgeneratedlog WHERE userid = ? AND deleteflg =0 and activeflg=1 ORDER BY reportgendate DESC LIMIT ?,?"

func (dbc DbConn) GetDownloadList(tz *entities.ReportDownloadListEntity) ([]entities.ReportDownloadListEntity, error) {
	logger.Log.Println("In side dao GetDownloadList")
	values := []entities.ReportDownloadListEntity{}

	var params []interface{}

	params = append(params, tz.Refuserid)
	params = append(params, tz.Offset)
	params = append(params, tz.Limit)
	rows, err := dbc.DB.Query(getDownloadList, params...)

	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetDownloadList Get Statement Prepare Error", err)
		return values, err
	}

	for rows.Next() {
		value := entities.ReportDownloadListEntity{}
		rows.Scan(&value.Id, &value.Refuserid, &value.Originalfilename, &value.Uploadedfilename)
		values = append(values, value)
	}

	return values, nil
}

func (dbc DbConn) GetDownloadListCount(tz *entities.ReportDownloadListEntity) (entities.ReportDownloadListEntities, error) {
	logger.Log.Println("In side GetDownloadListCount")
	value := entities.ReportDownloadListEntities{}
	var getDownloadListcount string
	var params []interface{}

	getDownloadListcount = "SELECT count(a.id) as total FROM mstreportgeneratedlog a WHERE a.userid = ?  AND a.deleteflg =0 and a.activeflg=1 "
	params = append(params, tz.Refuserid)

	err := dbc.DB.QueryRow(getDownloadListcount, params...).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetDownloadListCount Get Statement Prepare Error", err)
		return value, err
	}
}
