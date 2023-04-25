package dao

import (
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
)

// var recordbytypebymultiorg=""
func (mdao DbConn) GetExternalcategorynames(ClientID int64, Mstorgnhirarchyid int64, RecorddifftypeID int64, RecorddifID int64, RecordID int64) ([]entities.Categorydetails, error) {
	var sql = "SELECT distinct (SELECT typename from mstrecorddifferentiationtype WHERE id=a.torecorddifftypeid) lable,d.name FROM mstrecordtype a, mstrecorddifferentiationtype b,maprecordtorecorddifferentiation c,mstrecorddifferentiation d WHERE a.torecorddifftypeid = b.id AND a.clientid=? AND a.mstorgnhirarchyid=? AND a.deleteflg=0 AND a.activeflg=1 AND a.fromrecorddifftypeid=? AND a.fromrecorddiffid=? AND b.parentid=1 AND c.recordid=? AND c.islatest=1 AND a.torecorddifftypeid=c.recorddifftypeid and c.recorddiffid=d.id"
	v := []entities.Categorydetails{}
	logger.Log.Println(sql)
	logger.Log.Println("PARAMETER", ClientID, Mstorgnhirarchyid, RecorddifftypeID, RecorddifID, RecordID)
	rows, err := mdao.DB.Query(sql, ClientID, Mstorgnhirarchyid, RecorddifftypeID, RecorddifID, RecordID)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetLaststagevalue Get Statement Prepare Error", err)
		return v, err
	}
	// logger.Log.Println(rows)
	for rows.Next() {
		values := entities.Categorydetails{}
		err = rows.Scan(&values.Label, &values.Categoryname)
		logger.Log.Println("GetLaststagevalue rows.next() Error", err)
		v = append(v, values)
	}
	logger.Log.Println(v)
	return v, nil
}
func (dbc DbConn) GetSchema() ([]string, error) {
	logger.Log.Println("In side GetSchema")
	getschema := "select COLUMN_NAME from INFORMATION_SCHEMA.COLUMNS where TABLE_Name='recordfulldetails' order by ORDINAL_POSITION"
	// logger.Log.Println(getschema)
	var values []string
	rows, err := dbc.DB.Query(getschema)
	logger.Log.Println(getschema)
	if err != nil {
		logger.Log.Println("GetSchema Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		var value string
		rows.Scan(&value)
		values = append(values, value)
	}
	return values, nil
}
func (dbc DbConn) GetTableData(orgids string, page *entities.RecordfulldetailsRequestEntity, statusids string, tickettypeids string) ([]entities.ExcelEntity, error) {
	logger.Log.Println("In side GetTableData")
	values := []entities.ExcelEntity{}
	str := "a.clientid,a.mstorgnhirarchyid,a.recordid,a.tickettypeid,coalesce(a.levelonecatename,''),coalesce(a.ticketid,''),coalesce(a.shortdescription,''),coalesce(a.status,''),coalesce(a.latestresodatetime,'') as resolveddate,coalesce(b.recordtrackvalue,''),coalesce(e.recordtrackvalue,'')"
	getschema := "select " + str + " FROM recordfulldetails a LEFT JOIN mstrecordterms c ON c.clientid=a.clientid AND c.mstorgnhirarchyid=a.mstorgnhirarchyid AND c.seq=4 AND c.activeflg=1 AND c.deleteflg=0 LEFT JOIN trnreordtracking b ON a.recordid = b.recordid AND c.id=b.recordtermid LEFT JOIN mstrecordterms d ON d.clientid=a.clientid AND d.mstorgnhirarchyid=a.mstorgnhirarchyid AND d.seq=5 AND d.activeflg=1 AND d.deleteflg=0 LEFT JOIN trnreordtracking e ON a.recordid = e.recordid AND d.id=e.recordtermid WHERE a.clientid=? AND a.mstorgnhirarchyid in (" + orgids + ") and a.createddatetime between ? and ? and a.statusid in (" + statusids + ") and a.tickettypeid in(" + tickettypeids + ")"
	logger.Log.Println(getschema)
	rows, err := dbc.DB.Query(getschema, page.Clientid, page.Fromdate, page.Todate)

	// rows, err := dbc.DB.Query(getAsset, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllAsset Get Statement Prepare Error", err)
		return values, err
	}
	// logger.Log.Println(rows)
	for rows.Next() {
		value := entities.ExcelEntity{}
		rows.Scan(&value.Clientid, &value.Mstorgnhirarchyid, &value.Recordid, &value.Tickettypeid, &value.CustomerName, &value.TicketNo, &value.Shortdescription, &value.Status, &value.Latestresodatetime, &value.VendorName, &value.VendorTicketid)
		values = append(values, value)
	}
	return values, nil

}

// func (dbc DbConn) GetTableData(column []string, orgids string, page *entities.RecordfulldetailsRequestEntity, statusids string, tickettypeids string) ([]map[string]interface{}, error) {
// 	logger.Log.Println("In side GetTableData")
// 	str := strings.Join(column, ",")
// 	// logger.Log.Pr,&value.ln()
// 	var getschema string
// 	var params []interface{}
// 	var values []map[string]interface{}

// 	// values := []map[,&value.],&value.erface{}{}
// 	if page.Fromdate == "" && page.Todate == "" {
// 		getschema = "select " + str + " from recordfulldetails where clientid=? and mstorgnhirarchyid in (" + orgids + ") and  statusid in (" + statusids + ") and tickettypeid in(" + tickettypeids + ")"
// 		params = append(params, page.Clientid)
// 	} else {

// 		getschema = "select " + str + " from recordfulldetails where clientid=? and mstorgnhirarchyid in (" + orgids + ") and createddatetime between ? and ? and statusid in (" + statusids + ") and tickettypeid in(" + tickettypeids + ")"
// 		params = append(params, page.Clientid)
// 		params = append(params, page.Fromdate)
// 		params = append(params, page.Todate)

// 	}
// 	logger.Log.Println("Query>>>>>>>", getschema)
// 	rows, err := dbc.DB.Query(getschema, params...)
// 	if err != nil {
// 		logger.Log.Println("GetTableData Get Statement Prepare Error", err)
// 		return values, err
// 	}
// 	for rows.Next() {
// 		fb := entities.NewDbFieldBind()
// 		err = fb.PutFields(rows)
// 		if err != nil {
// 			return values, err
// 		}
// 		err := rows.Scan(fb.GetFieldPtrArr()...)
// 		if err != nil {
// 			logger.Log.Println(err)
// 		}
// 		//logger.Log.Println("values", fb.GetFieldArr())
// 		values = append(values, fb.GetFieldArr())
// 		// fmt.Println(fb.GetFieldArr())
// 	}
// 	// logger.Log.Pr,&value.ln(len(values))
// 	return values, nil
// }

// type RecordfulldetailsRequestEntity struct {
// 	Clientid           int64                    `json:"clientid"`
// 	Mstorgnhirarchyid  int64                    `json:"mstorgnhirarchyid"`
// 	Mstorgnhirarchyids []int64                  `json:"mstorgnhirarchyids"`
// 	Fromdate           string                   `json:"fromdate"`
// 	Todate             string                   `json:"todate"`
// 	Result             []map[string]interface{} `json:"result"`
// }

// func (w *RecordfulldetailsRequestEntity) FromJSON(r io.Reader) error {
// 	e := json.NewDecoder(r)
// 	return e.Decode(w)
// }

// type RecordfulldetailsresultEntities struct {
// 	Id                           int
// 	Clientid                     int
// 	Mstorgnhirarchyid            int
// 	Recordid                     int
// 	Ticketid                     string
// 	Source                       string
// 	Requestorid                  int
// 	Requestorname                string
// 	Requestorlocation            string
// 	Requestorphone               string
// 	Requestoremail               string
// 	Requestorloginid             string
// 	Orgcreatorlocation           string
// 	Orgcreatorphone              string
// 	Orgcreatorid                 int
// 	Orgcreatorname               string
// 	Orgcreatorloginid            string
// 	Orgcreatoremail              string
// 	Tickettypeid                 int
// 	Tickettype                   string
// 	Priorityid                   int
// 	Priority                     string
// 	Statusid                     int
// 	Status                       string
// 	Vipticket                    string
// 	Urgencyid                    int
// 	Urgency                      string
// 	Impactid                     int
// 	Impact                       string
// 	Shortdescription             string
// 	Assigneduserloginid          string
// 	Createddatetime              string
// 	Lastupdateddatetime          string
// 	Assignedgroupid              int
// 	Assignedgroup                string
// 	Assigneduserid               int
// 	Assigneduser                 string
// 	Resogroupid                  int
// 	Resogroup                    string
// 	Resolveduserid               int
// 	Resolveduser                 string
// 	Lastuserid                   int
// 	Lastuser                     string
// 	Reassigncount                int
// 	Respslastatusid              int
// 	Respslabreachstatus          string
// 	Resostatusid                 int
// 	Resolslabreachstatus         string
// 	Reopencount                  int
// 	Reopendatetime               string
// 	Reopenedflag                 string
// 	Prioritycount                int
// 	Followupcount                int
// 	Followupdatetime             string
// 	Followuprespdatetime         string
// 	Followuptimetaken            int
// 	Outboundcount                int
// 	Isparent                     string
// 	Childcount                   int
// 	Parentticketid               int
// 	Responseslameterpercentage   float64
// 	Resolutionslameterpercentage float64
// 	Worknotenotupdated           int
// 	Respsladuedatetime           string
// 	Resosladuedatetime           string
// 	Categorychangecount          int
// 	Firstresponsedatetime        string
// 	Latestresponsedatetime       string
// 	Firstresodatetime            string
// 	Latestresodatetime           string
// 	Responsetime                 int
// 	Resolutiontime               int
// 	Respclockstatus              string
// 	Resoclockstatus              string
// 	Businessaging                int
// 	Calendaraging                int
// 	Actualeffort                 int
// 	Slaidletime                  int
// 	Respoverduetime              int
// 	Respoverdueperc              float64
// 	Resooverduetime              int
// 	Resooverdueperc              float64
// 	Pendinguserdatetime          string
// 	Userreplieddatetime          string
// 	Userreplytimetaken           int
// 	Pendingusercount             int
// 	Closedatetime                string
// 	Csatscore                    string
// 	Csatcomment                  string
// 	Activeflg                    int
// 	Deleteflg                    int
// 	Audittransactionid           int
// 	Ifixsysid                    string
// 	Lastupdatedby                int
// 	Lastupdateddate              string
// }
func (mdao DbConn) Gettimediffbyinterface(Clientid interface{}, Mstorgnhirarchyid interface{}) (error, []entities.UtilityEntity) {
	requestIds := []entities.UtilityEntity{}
	stmt, err := mdao.DB.Prepare("SELECT a.utcdiff as timediff,c.utcdiff as reporttimediff,coalesce(d.example,'') Timeformat,coalesce(e.example,'') Reporttimeformat from zone a,mstorgnhierarchy b,zone c,mstdatetimeformat d,mstdatetimeformat e where a.zone_id=b.timezoneid and b.reporttimezoneid=c.zone_id and b.clientid=? and b.id=? and b.timeformatid=d.id and b.reporttimeformatid=e.id")
	if err != nil {
		logger.Log.Print("Gettimediff Statement Prepare Error", err)
		// log.Print("Gettimediff Statement Prepare Error", err)
		return err, requestIds
	}
	defer stmt.Close()
	rows, err := stmt.Query(Clientid, Mstorgnhirarchyid)
	if err != nil {
		logger.Log.Print("Gettimediff Statement Execution Error", err)
		// log.Print("Gettimediff Statement Execution Error", err)
		return err, requestIds
	}
	for rows.Next() {
		value := entities.UtilityEntity{}
		rows.Scan(&value.Timediff, &value.Reporttimediff, &value.Timeformat, &value.Reporttimeformat)
		requestIds = append(requestIds, value)
	}
	return nil, requestIds
}
func (mdao DbConn) GetRecordByDiffTypeOfMultiOrg(tz *entities.RecordfulldetailsRequestEntity, ids string) ([]entities.RecordDiffOfMultiOrgEntity, error) {
	log.Println("In side dao")
	values := []entities.RecordDiffOfMultiOrgEntity{}
	var recordbytypebymultiorg string
	var param []interface{}
	// recordbytypebymultiorg := "SELECT distinct a.id as ID,concat(a.name,'(',b.name,')') as name,a.seqno from mstrecorddifferentiation a,mstorgnhierarchy b WHERE a.clientid=? and a.mstorgnhirarchyid in(" + ids + ") and a.mstorgnhirarchyid=b.id and a.recorddifftypeid in (SELECT c.id from mstrecorddifferentiationtype c where c.seqno=? and c.deleteflg=0 and c.activeflg=1) and a.activeflg=1 and a.deleteflg=0"
	if tz.Seqno == 1 {
		recordbytypebymultiorg = "select distinct seqno,name  from mstrecorddifferentiation where recorddifftypeid=2 and clientid=? and mstorgnhirarchyid in(" + ids + ") and deleteflg=0 and activeflg=1"
		param = append(param, tz.Clientid)
	} else {
		recordbytypebymultiorg = "select distinct a.seqno,a.name from mstrecorddifferentiation a ,mstrecordtype b where b.fromrecorddifftypeid=2 and b.fromrecorddiffid in (select id from mstrecorddifferentiation where seqno=? and mstorgnhirarchyid in (" + ids + ") and clientid=? and  deleteflg=0 and activeflg=1) and b.torecorddifftypeid=3 and a.id=b.torecorddiffid and b.activeflg=1 and b.deleteflg=0 and a.activeflg=1 and a.deleteflg=0;"
		param = append(param, tz.Tickettypeseq, tz.Clientid)
	}

	rows, err := mdao.DB.Query(recordbytypebymultiorg, param...)

	if err != nil {
		logger.Log.Print("GetRecordDiffType Get Statement Prepare Error", err)
		log.Print("GetRecordDiffType Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		value := entities.RecordDiffOfMultiOrgEntity{}
		rows.Scan(&value.Seqno, &value.Name)
		values = append(values, value)
	}
	return values, nil
}
func (mdao DbConn) GetStatusIdsOfMultiOrg(tz *entities.RecordfulldetailsRequestEntity, orgids string, statusids string, tickettypeids string) ([]int64, error) {
	log.Println("In side dao")
	values := []int64{}
	var recordbytypebymultiorg string

	var param []interface{}
	if tz.Seqno == 3 {
		recordbytypebymultiorg = "select id from mstrecorddifferentiation where clientid=2 and mstorgnhirarchyid in(" + orgids + ") and seqno in(" + statusids + ") and deleteflg=0 and recorddifftypeid=3"
	} else {
		recordbytypebymultiorg = "select id from mstrecorddifferentiation where clientid=2 and mstorgnhirarchyid in(" + orgids + ") and seqno in(" + tickettypeids + ") and deleteflg=0 and recorddifftypeid=2"
	}
	// if tz.Seqno == 1 {
	// 	recordbytypebymultiorg = "select distinct seqno,name  from mstrecorddifferentiation where recorddifftypeid=2 and clientid=? and mstorgnhirarchyid in(" + ids + ") and deleteflg=0 and activeflg=1"
	// 	param = append(param, tz.Clientid)
	// } else {
	// 	recordbytypebymultiorg = "select distinct a.seqno,a.name from mstrecorddifferentiation a ,mstrecordtype b where b.fromrecorddifftypeid=2 and b.fromrecorddiffid in (select id from mstrecorddifferentiation where seqno=? and mstorgnhirarchyid in (" + ids + ") and clientid=? and  deleteflg=0 and activeflg=1) and b.torecorddifftypeid=3 and a.id=b.torecorddiffid and b.activeflg=1 and b.deleteflg=0 and a.activeflg=1 and a.deleteflg=0;"
	// 	param = append(param, tz.Tickettypeseq, tz.Clientid)
	// }

	rows, err := mdao.DB.Query(recordbytypebymultiorg, param...)

	if err != nil {
		logger.Log.Print("GetRecordDiffType Get Statement Prepare Error", err)
		log.Print("GetRecordDiffType Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		var value int64
		rows.Scan(&value)
		values = append(values, value)
	}
	return values, nil
}
