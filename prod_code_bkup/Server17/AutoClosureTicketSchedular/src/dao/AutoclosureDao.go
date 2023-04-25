package dao

import (
	"database/sql"
	"errors"
	"log"
	"src/entities"
	Logger "src/logger"
)

func UpdateDefectiveFlag(db *sql.DB, ClientID int64, Mstorgnhirarchyid int64, RecordID int64) error {
	var sql = "UPDATE mstrecordautoclosure SET defective='Y' WHERE clientid=? AND mstorgnhirarchyid=? AND recordid=?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		Logger.Log.Println(err)
		return errors.New("SQL Prepare Error")
	}
	defer stmt.Close()
	_, err = stmt.Exec(ClientID, Mstorgnhirarchyid, RecordID)
	if err != nil {
		Logger.Log.Println(err)
		return errors.New("SQL Execution Error")
	}

	return nil
}
func GetAutoClosureConfigList(db *sql.DB) ([]entities.AutoClosureConfig, error) {
	Logger.Log.Println("In side GetResolvedRecordsInfo model function")
	t := []entities.AutoClosureConfig{}
	var Sql = "SELECT clientid,id as mstorgnid,autoclose,coalesce(autoclosefromdt,'') autoclosefromdt,coalesce(autoclosetodt,'') autoclosetodt FROM mstorgnhierarchy where autoclose=1 and  deleteflg=0"
	Logger.Log.Println(Sql)

	stmt, err := db.Prepare(Sql)
	if err != nil {
		log.Print("Error in GetResolvedRecordsInfo--->", err)
		return t, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Print("Error in GetResolvedRecordsInfo--->", err)
		return t, err
	}
	for rows.Next() {
		value := entities.AutoClosureConfig{}
		rows.Scan(&value.ClientID, &value.MstorgnhirarchyID, &value.IsAutoclose, &value.AutocloseFromdate, &value.AutocloseTodate)
		t = append(t, value)
	}
	return t, nil
}
func GetResolvedRecordsInfo(db *sql.DB, tz entities.AutoClosureConfig) ([]entities.RecordInfo, error) {
	Logger.Log.Println("In side GetResolvedRecordsInfo model function")
	t := []entities.RecordInfo{}
	//var Sql = "SELECT a.clientid,a.mstorgnhirarchyid,a.recordid,b.recordstageid,a.closuredate,b.recorddifftypeid,b.recorddiffid,c.usergroupid,c.userid,e.mststateid FROM mstrecordautoclosure a,maprecordtorecorddifferentiation b,trnrecord c,mstrecorddifferentiation d,maprecordstatetodifferentiation e WHERE a.closureflag='N' AND a.deleteflg=0 AND a.activeflg=1 AND a.recordid=b.recordid AND b.islatest=1 AND b.isworking=1 AND a.recordid=c.id AND a.recordseq=d.seqno AND d.clientid=c.clientid AND d.mstorgnhirarchyid=c.mstorgnhirarchyid AND d.recorddifftypeid=3 AND d.seqno=3 AND d.deleteflg=0 AND d.activeflg=1 AND d.id = e.recorddiffid AND e.deleteflg=0 AND e.activeflg=1 AND a.closuredate < round(UNIX_TIMESTAMP(now()))"
	//	var Sql = " select a.clientid,a.mstorgnhirarchyid,a.recordid,a.recordstageid,b.closuredate,d.recorddifftypeid,d.recorddiffid,c.usergroupid,c.userid,e.mststateid from maprecordtorecorddifferentiation a , mstrecordautoclosure b, trnrecord c,maprecordtorecorddifferentiation d,maprecordstatetodifferentiation e where  a.recorddiffid in (select id from mstrecorddifferentiation where recorddifftypeid=3 and seqno=3 and deleteflg=0 and activeflg=1 ) and b.islatest=1 and a.islatest=1 and d.isworking=1 and d.islatest=1 and d.recordid=a.recordid and a.recorddifftypeid=3 and a.recordid=b.recordid and b.closureflag='N' AND  a.recordid= c.id and a.recorddiffid=e.recorddiffid and a.clientid=? and a.mstorgnhirarchyid=? and  FROM_UNIXTIME(b.closuredate) between ? and ? limit 10"
	var Sql = `select a.clientid,a.mstorgnhirarchyid,a.recordid,a.recordstageid,b.closuredate,d.recorddifftypeid,d.recorddiffid,
c.usergroupid,c.userid,e.mststateid from maprecordtorecorddifferentiation a , mstrecordautoclosure b, trnrecord c,
maprecordtorecorddifferentiation d,maprecordstatetodifferentiation e,maprecordtorecorddifferentiation f where  a.recorddiffid in
 (select id from mstrecorddifferentiation where recorddifftypeid=3 and seqno=3 and deleteflg=0 and activeflg=1 ) and b.islatest=1 and
 a.islatest=1 and d.isworking=1 and d.islatest=1 and d.recordid=a.recordid and a.recorddifftypeid=3 and a.recordid=b.recordid and 
b.closureflag='N' AND defective='N' AND  a.recordid= c.id and a.recorddiffid=e.recorddiffid and a.clientid=? and a.mstorgnhirarchyid=? and 
 b.closuredt between ? and ?  and b.closuredt<=ADDTIME(now(), '5:30:0.0') and a.recordid=f.recordid and f.recorddifftypeid=2 and
 f.islatest=1 and f.recorddiffid in (select id from mstrecorddifferentiation where recorddifftypeid=2 and seqno in (1,2) and deleteflg=0 and
 activeflg=1 )  limit 200;`

	Logger.Log.Println(Sql, tz.ClientID, tz.MstorgnhirarchyID, tz.AutocloseFromdate, tz.AutocloseTodate)

	stmt, err := db.Prepare(Sql)
	if err != nil {
		log.Print("Error in GetResolvedRecordsInfo--->", err)
		return t, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(tz.ClientID, tz.MstorgnhirarchyID, tz.AutocloseFromdate, tz.AutocloseTodate)
	if err != nil {
		log.Print("Error in GetResolvedRecordsInfo--->", err)
		return t, err
	}
	for rows.Next() {
		value := entities.RecordInfo{}
		rows.Scan(&value.ClientID, &value.MstorgnhirarchyID, &value.RecordID, &value.RecordStageID, &value.Closuredate, &value.WorkingDifftypeID, &value.WorkingDiffID, &value.CreatedgrpID, &value.MstuserID, &value.PreviousStateID)
		t = append(t, value)
	}
	return t, nil
}

func GetNxtStateID(db *sql.DB, ClientID int64, MstorgnhirarchyID int64) (int64, error) {
	Logger.Log.Println("In side GetNxtStateID model function")
	var ID int64
	var Sql = "SELECT b.mststateid FROM mstrecorddifferentiation a,maprecordstatetodifferentiation b WHERE a.clientid=? AND a.mstorgnhirarchyid=? AND a.recorddifftypeid=3 AND a.seqno=8 AND a.deleteflg=0 AND a.activeflg=1 AND a.id = b.recorddiffid AND b.deleteflg=0 AND b.activeflg=1"
	stmt, err := db.Prepare(Sql)
	if err != nil {
		log.Print("Error in GetResolvedRecordsInfo--->", err)
		return ID, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(ClientID, MstorgnhirarchyID)
	if err != nil {
		log.Print("Error in GetResolvedRecordsInfo--->", err)
		return ID, err
	}
	for rows.Next() {
		rows.Scan(&ID)
	}
	return ID, nil
}
func GetDiffID(db *sql.DB, ClientID int64, MstorgnhirarchyID int64, RecordID int64) (int64, error) {
	Logger.Log.Println("In side GetNxtStateID model function")
	var ID int64
	var Sql = "SELECT b.mststateid FROM mstrecorddifferentiation a,maprecordstatetodifferentiation b WHERE a.clientid=? AND a.mstorgnhirarchyid=? AND a.recorddifftypeid=3 AND a.seqno=8 AND a.deleteflg=0 AND a.activeflg=1 AND a.id = b.recorddiffid AND b.deleteflg=0 AND b.activeflg=1"
	stmt, err := db.Prepare(Sql)
	if err != nil {
		log.Print("Error in GetResolvedRecordsInfo--->", err)
		return ID, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(ClientID, MstorgnhirarchyID)
	if err != nil {
		log.Print("Error in GetResolvedRecordsInfo--->", err)
		return ID, err
	}
	for rows.Next() {
		rows.Scan(&ID)
	}
	return ID, nil
}
func GetCurrentStatusSeq(db *sql.DB, ClientID int64, MstorgnhirarchyID int64, RecordID int64) (int64, error) {
	Logger.Log.Println("In side GetCurrentStatusSeq model function")
	var Seq int64
	var Sql = "SELECT b.seqno FROM  maprecordtorecorddifferentiation a,mstrecorddifferentiation b WHERE a.clientid=? AND a.mstorgnhirarchyid=? AND a.recordid=? AND a.recorddifftypeid=3 AND a.islatest=1 AND a.recorddiffid=b.id AND b.activeflg=1 AND b.deleteflg=0"
	stmt, err := db.Prepare(Sql)
	if err != nil {
		log.Print("Error in GetCurrentStatusSeq--->", err)
		return Seq, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(ClientID, MstorgnhirarchyID, RecordID)
	if err != nil {
		log.Print("Error in GetCurrentStatusSeq--->", err)
		return Seq, err
	}
	for rows.Next() {
		rows.Scan(&Seq)
	}
	return Seq, nil
}

func Termsequance(db *sql.DB, ClientID int64, Mstorgnhirarchyid int64) (map[int64]int64, error) {
	var termseq = "SELECT id,seq FROM mstrecordterms WHERE clientid=? AND mstorgnhirarchyid=? AND deleteflg=0 AND activeflg=1 AND seq is not null"
	var t = make(map[int64]int64)
	var ID int64
	var Seq int64
	rows, err := db.Query(termseq, ClientID, Mstorgnhirarchyid)
	defer rows.Close()
	if err != nil {
		Logger.Log.Println("Getfrequentissues Get Statement Prepare Error", err)
		return t, err
	}
	for rows.Next() {
		err = rows.Scan(&ID, &Seq)
		t[Seq] = ID
	}
	//Logger.Log.Println("Hashmap value is---------------->", t)
	return t, nil
}

func InsertClosureComment(db *sql.DB, ClientID int64, Mstorgnhirarchyid int64, RecordID int64, RecordStageID int64, TermID int64, TermValue string, UserID int64, GrpID int64) error {
	Logger.Log.Println("parameters -->", ClientID, Mstorgnhirarchyid, RecordID, RecordStageID, TermID, UserID)
	var Sql = "INSERT INTO trnreordtracking (clientid, mstorgnhirarchyid,recordid, recordstageid,recordtermid, recordtrackvalue,createdbyid,createddate,createdgrpid) VALUES (?,?,?,?,?,?,?,round(UNIX_TIMESTAMP(now())),?)"
	stmt, err := db.Prepare(Sql)

	if err != nil {
		Logger.Log.Println(err)
		return errors.New("SQL Prepare Error")
	}
	defer stmt.Close()
	_, err = stmt.Exec(ClientID, Mstorgnhirarchyid, RecordID, RecordStageID, TermID, TermValue, UserID, GrpID)
	if err != nil {
		Logger.Log.Println(err)
		return errors.New("SQL Execution Error")
	}

	return nil
}

func InsertNPSFeedback(db *sql.DB, ClientID int64, Mstorgnhirarchyid int64, RecordID int64, RecordStageID int64, TermID int64, TermValue string, UserID int64, GrpID int64) error {
	Logger.Log.Println("parameters -->", ClientID, Mstorgnhirarchyid, RecordID, RecordStageID, TermID, UserID)
	var Sql = "INSERT INTO trnreordtracking (clientid, mstorgnhirarchyid,recordid, recordstageid,recordtermid, recordtrackvalue,createdbyid,createddate,createdgrpid) VALUES (?,?,?,?,?,?,?,round(UNIX_TIMESTAMP(now())),?)"
	stmt, err := db.Prepare(Sql)

	if err != nil {
		Logger.Log.Println(err)
		return errors.New("SQL Prepare Error")
	}
	defer stmt.Close()
	_, err = stmt.Exec(ClientID, Mstorgnhirarchyid, RecordID, RecordStageID, TermID, TermValue, UserID, GrpID)
	if err != nil {
		Logger.Log.Println(err)
		return errors.New("SQL Execution Error")
	}

	return nil
}

func Gettermnamebyid(db *sql.DB, Termid int64, ClientID int64, Mstorgnhirarchyid int64) (string, error) {
	//Logger.Log.Println("In side Gettermnamebyid")
	var tername string
	var termnmbyid = "SELECT termname FROM mstrecordterms where clientid=? AND mstorgnhirarchyid=? AND id=?"
	rows, err := db.Query(termnmbyid, ClientID, Mstorgnhirarchyid, Termid)
	defer rows.Close()
	if err != nil {
		Logger.Log.Println("Gettermnamebyid Get Statement Prepare Error", err)
		return tername, err
	}
	for rows.Next() {
		err = rows.Scan(&tername)
		Logger.Log.Println("Gettermnamebyid rows.next() Error", err)
	}
	return tername, nil
}

func InsertActivityLogs(db *sql.DB, ClientID int64, Mstorgnhirarchyid int64, RecordID int64, ActivityID int64, Logvalue string, CreatedID int64, CreatedgrpID int64, TermID int64) error {
	//Logger.Log.Println("parameters -->", ClientID, Mstorgnhirarchyid, RecordID, ActivityID, Logvalue, CreatedID, CreatedgrpID, TermID)
	var insertlogswithgenericID = "INSERT INTO mstrecordactivitylogs(clientid,mstorgnhirarchyid,recordid,activityseqno,logValue,createdid,createddate,createdgrpid,genericid) VALUES (?,?,?,?,?,?,round(UNIX_TIMESTAMP(now())),?,?)"
	stmt, err := db.Prepare(insertlogswithgenericID)

	if err != nil {
		Logger.Log.Println(err)
		return errors.New("SQL Prepare Error")
	}
	defer stmt.Close()
	_, err = stmt.Exec(ClientID, Mstorgnhirarchyid, RecordID, ActivityID, Logvalue, CreatedID, CreatedgrpID, TermID)
	if err != nil {
		Logger.Log.Println(err)
		return errors.New("SQL Execution Error")
	}

	return nil
}

func UpdateclosureFlag(db *sql.DB, ClientID int64, Mstorgnhirarchyid int64, RecordID int64) error {
	var sql = "UPDATE mstrecordautoclosure SET closureflag='Y' WHERE clientid=? AND mstorgnhirarchyid=? AND recordid=?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		Logger.Log.Println(err)
		return errors.New("SQL Prepare Error")
	}
	defer stmt.Close()
	_, err = stmt.Exec(ClientID, Mstorgnhirarchyid, RecordID)
	if err != nil {
		Logger.Log.Println(err)
		return errors.New("SQL Execution Error")
	}

	return nil
}
