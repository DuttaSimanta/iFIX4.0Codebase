package dao

import (
	"database/sql"
	"log"
	"src/logger"
)

func GetTickettype(db *sql.DB) []int64 {
	// logger.Log.Println("In side dao")
	values := []int64{}
	var sql = "SELECT recorddiffid FROM mstrecordconfig WHERE clientid=2 AND mstorgnhirarchyid=2 AND recorddifftypeid=2 AND deleteflg=0 AND activeflg=1"

	rows, err := db.Query(sql)

	if err != nil {
		logger.Log.Print("GetRecordDiffType Get Statement Prepare Error", err)
		// log.Print("GetRecordDiffType Get Statement Prepare Error", err)
		return values
	}
	defer rows.Close()
	for rows.Next() {
		var value int64
		rows.Scan(&value)
		values = append(values, value)
	}
	return values
}
func GetTicketCount(db *sql.DB, recorddiffID int64) int64 {
	// logger.Log.Println("In side dao")
	var value int64
	var sql = "SELECT count(id) FROM mstrecordcode WHERE clientid=2 AND mstorgnhirarchyid=2 AND recorddifftypeid=2 AND recorddiffid=? AND deleteflg=0 AND activeflg=1 and isuse=0"

	rows, err := db.Query(sql, recorddiffID)

	if err != nil {
		logger.Log.Print("GetRecordDiffType Get Statement Prepare Error", err)
		// log.Print("GetRecordDiffType Get Statement Prepare Error", err)
		return value
	}
	defer rows.Close()
	for rows.Next() {
		// var value string
		rows.Scan(&value)
		// values = append(values, value)
	}
	return value
}
func Getmstrecordconfig(db *sql.DB, clientid int64, mstorgnhirarchyid int64, recorddifftypeid int64, recorddiffID int64) []string {
	// logger.Log.Println("In side dao")
	var prefixarr []string
	var mstrecordconfig = "SELECT prefix,year,month,day,configurezero FROM mstrecordconfig WHERE clientid=? AND mstorgnhirarchyid=? AND recorddifftypeid=? AND recorddiffid=? AND deleteflg=0 AND activeflg=1 "

	rows, err := db.Query(mstrecordconfig, clientid, mstorgnhirarchyid, recorddifftypeid, recorddiffID)

	if err != nil {
		logger.Log.Print("GetRecordDiffType Get Statement Prepare Error", err)
		// log.Print("GetRecordDiffType Get Statement Prepare Error", err)
		return prefixarr
	}
	defer rows.Close()
	for rows.Next() {
		var prefix string
		var year string
		var month string
		var day string
		var configurezero string

		if err := rows.Scan(&prefix, &year, &month, &day, &configurezero); err != nil {
			logger.Log.Println("Error in fetching data")
		}
		prefixarr = append(prefixarr, prefix)
		prefixarr = append(prefixarr, year)
		prefixarr = append(prefixarr, month)
		prefixarr = append(prefixarr, day)
		prefixarr = append(prefixarr, configurezero)
	}
	return prefixarr
}
func Getnumber(db *sql.DB, clientid int64, mstorgnhirarchyid int64, recorddifftypeid int64, recorddiffID int64) int {
	// logger.Log.Println("In side dao")
	var value int
	var getnumber = "SELECT number FROM mstrecordautoincreament WHERE clientid=? AND mstorgnhirarchyid=? AND recorddifftypeid=? AND recorddiffid=? AND activeflg=1 AND deleteflg=0"

	rows, err := db.Query(getnumber, clientid, mstorgnhirarchyid, recorddifftypeid, recorddiffID)

	if err != nil {
		logger.Log.Print("GetRecordDiffType Get Statement Prepare Error", err)
		// log.Print("GetRecordDiffType Get Statement Prepare Error", err)
		return value
	}
	defer rows.Close()
	for rows.Next() {
		// var value string
		rows.Scan(&value)
		// values = append(values, value)
	}
	return value
}
func Updatemstrecordautoincreament(db *sql.DB, recorddiffID int64) error {
	// logger.Log.Println("In side dao")
	var updatenumber = "UPDATE mstrecordautoincreament SET number = (number+1) WHERE clientid=? AND mstorgnhirarchyid=? AND recorddifftypeid=? AND recorddiffid=? AND activeflg=1 AND deleteflg=0"

	stmt, err := db.Prepare(updatenumber)
	defer stmt.Close()
	if err != nil {
		logger.Log.Print("UpdateRecordDiff Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(2, 2, 2, recorddiffID)
	if err != nil {
		logger.Log.Print("UpdateRecordDiff Execute Statement  Error", err)
		return err
	}
	return nil
}
func InsertRecordcode(db *sql.DB, recorddiffID interface{}, ticketID string) (int64, error) {
	// log.Println("In side dao")
	stmt, err := db.Prepare("INSERT INTO mstrecordcode(clientid,mstorgnhirarchyid,recorddifftypeid,recorddiffid,code) VALUES (?,?,?,?,?)")
	defer stmt.Close()
	if err != nil {
		log.Print("InsertRecordDiff Prepare Statement  Error", err)
		return 0, err
	}
	res, err := stmt.Exec(2, 2, 2, recorddiffID, ticketID)
	if err != nil {
		log.Print("InsertRecordDiff Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}
func Deleterecordcode(db *sql.DB) {

	stmt, err4 := db.Prepare("delete from mstrecordcode where isuse=1")
	if err4 != nil {
		logger.Log.Println(err4.Error())
	}
	_, err := stmt.Exec()
	if err != nil {
		logger.Log.Println(err.Error())
	}
	stmt.Close()
}
