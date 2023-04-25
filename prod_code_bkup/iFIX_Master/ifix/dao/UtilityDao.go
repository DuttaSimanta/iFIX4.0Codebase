package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
	"time"
)

var gettimediff = "SELECT a.utcdiff as timediff,c.utcdiff as reporttimediff,coalesce(d.example,'') Timeformat,coalesce(e.example,'') Reporttimeformat from zone a,mstorgnhierarchy b,zone c,mstdatetimeformat d,mstdatetimeformat e where a.zone_id=b.timezoneid and b.reporttimezoneid=c.zone_id and b.clientid=? and b.id=? and b.timeformatid=d.id and b.reporttimeformatid=e.id"
var deletetoken = "DELETE from msttoken WHERE userid=? and deleteflg=0"
var inserttoken = "INSERT into msttoken(token,userid) values(?,?)"
var gettoken = " SELECT token from msttoken where userid=? and deleteflg=0"

func Deletetoken(tx *sql.Tx, id int64) error {
	log.Println("In side dao")
	stmt, err := tx.Prepare(deletetoken)

	if err != nil {
		logger.Log.Print("deletetoken Prepare Statement  Error: ", err)
		log.Print("deletetoken Prepare Statement  Error: ", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		logger.Log.Print("deletetoken Execute Statement  Error: ", err)
		log.Print("deletetoken Execute Statement  Error: ", err)
		return err
	}
	return nil
}
func Inserttoken(tx *sql.Tx, id int64, token string) error {
	log.Println("In side dao")
	stmt, err := tx.Prepare(inserttoken)

	if err != nil {
		logger.Log.Print("inserttoken Prepare Statement  Error", err)
		log.Print("inserttoken Prepare Statement  Error", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(token, id)
	if err != nil {
		logger.Log.Print("inserttoken Execute Statement  Error", err)
		log.Print("inserttoken Execute Statement  Error", err)
		return err
	}
	return nil
}
func (mdao DbConn) Gettoken(userid int64) (error, []string) {
	var tokens []string
	stmt, err := mdao.DB.Prepare(gettoken)
	if err != nil {
		logger.Log.Print("Gettoken Statement Prepare Error", err)
		log.Print("Gettoken Statement Prepare Error", err)
		return err, tokens
	}
	defer stmt.Close()
	rows, err := stmt.Query(userid)
	if err != nil {
		logger.Log.Print("Gettoken Statement Execution Error", err)
		log.Print("Gettoken Statement Execution Error", err)
		return err, tokens
	}
	for rows.Next() {
		var token string
		rows.Scan(&token)
		tokens = append(tokens, token)
	}
	return nil, tokens
}

func (mdao DbConn) Gettimediff(tz *entities.UtilityEntity) (error, []entities.UtilityEntity) {
	requestIds := []entities.UtilityEntity{}
	stmt, err := mdao.DB.Prepare(gettimediff)
	if err != nil {
		logger.Log.Print("Gettimediff Statement Prepare Error", err)
		log.Print("Gettimediff Statement Prepare Error", err)
		return err, requestIds
	}
	defer stmt.Close()
	rows, err := stmt.Query(tz.Clientid, tz.Mstorgnhirarchyid)
	if err != nil {
		logger.Log.Print("Gettimediff Statement Execution Error", err)
		log.Print("Gettimediff Statement Execution Error", err)
		return err, requestIds
	}
	for rows.Next() {
		value := entities.UtilityEntity{}
		rows.Scan(&value.Timediff, &value.Reporttimediff,&value.Timeformat,&value.Reporttimeformat)
		requestIds = append(requestIds, value)
	}
	return nil, requestIds
}

func Convertdate(date int64, timediff int64,format string) string {
	var unixTime = date + timediff
	t := time.Unix(unixTime, 0)
	//return t.Format("02-Jan-2006 15:04:05")
	return t.Format(format)
}
