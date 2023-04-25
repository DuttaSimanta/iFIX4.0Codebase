package dao

import (
	"database/sql"
	"fmt"
	"math"
	"src/entities"
	"src/fileutils"
	"src/logger"
	"strings"
	"time"
	// "strconv"
)

func (dbc DbConn) RecordGridResultOnly(req map[string]interface{}) ([]map[string]interface{}, error) {
	// logger.log.Ptimein=time.Now()
	currentTime := time.Now()

	logger.Log.Println("YYYY.MM.DD 1111111111111 : ", currentTime.Format("2006.01.02 15:04:05 .999"))
	// logger.Log.Println("In side RecordGridResultOnly", time.Now)
	values := []map[string]interface{}{}
	pint := []interface{}{}

	var queryStr = "SELECT"

	if req["headers"] != nil {
		for i, v := range req["headers"].([]interface{}) {
			if i > 0 {
				queryStr = queryStr + " , "
			}
			queryStr = queryStr + " recordfulldetails." + v.(string) + " "
		}
		if len(req["headers"].([]interface{})) == 0 {
			queryStr = queryStr + " recordfulldetails.* "
		}
	} else {
		queryStr = queryStr + " recordfulldetails.* "
	}
	queryStr = queryStr + " FROM recordfulldetails  WHERE recordfulldetails.activeflg=1 AND recordfulldetails.deleteflg=0 "
	var whereMap []interface{}

	if req["where"] != nil {
		whereMap = req["where"].([]interface{})
		qstr, prm := GenerateCustomWhereQuery(whereMap, dbc.DB)
		queryStr = queryStr + " " + qstr
		pint = append(pint, prm...)
	}
	if req["cat"] != nil {
		catMap := req["cat"].([]interface{})
		if len(catMap) > 0 {
			qstr, prm := GenerateCatQuery(catMap)
			queryStr = queryStr + " AND " + qstr
			pint = append(pint, prm...)
		}
	}

	if req["order"] != nil {
		orderMap := req["order"].([]interface{})
		qstr := GenerateCustomOrderQuery(orderMap)
		queryStr = queryStr + " " + qstr
	}
	if req["limit"] != nil && req["offset"] != nil {
		queryStr = queryStr + " LIMIT ?,?"
		pint = append(pint, req["offset"])
		pint = append(pint, req["limit"])
	}
	logger.Log.Println("Query:===>", queryStr)
	logger.Log.Println("Params:===>", pint)
	rows, err := dbc.DB.Query(queryStr, pint...)
	if err != nil {
		logger.Log.Println("GetQueryResuls Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()

	for rows.Next() {
		fb := fileutils.NewDbFieldBind()
		err = fb.PutFields(rows)
		if err != nil {
			return values, err
		}
		rows.Scan(fb.GetFieldPtrArr()...)
		res := fb.GetFieldArr()
		// record := res["recordid"].(*fileutils.NullInt64)
		// recordid, _ := record.Value()
		// tickettypeid := res["tickettypeid"].(*fileutils.NullInt64)
		// diffid, _ := tickettypeid.Value()
		// client := res["clientid"].(*fileutils.NullInt64)
		// clientid, _ := client.Value()
		// org := res["mstorgnhirarchyid"].(*fileutils.NullInt64)
		// orgid, _ := org.Value()
		// categories, err := GetCategoryNames(dbc, clientid.(int64), orgid.(int64), 2, diffid.(int64), recordid.(int64))
		// if err != nil {
		// 	return values, err
		// }
		// res["categories"] = categories
		// statereson, err := GetStateReson(dbc, clientid.(int64), orgid.(int64), 2, diffid.(int64), recordid.(int64))
		// if err != nil {
		// 	return values, err
		// }
		// res["statusreson"] = statereson
		// visiblecomment, err := Getvisiblecomment(dbc, clientid.(int64), orgid.(int64), 2, diffid.(int64), recordid.(int64))
		// if err != nil {
		// 	return values, err
		// }
		// res["visiblecomment"] = visiblecomment
		// slatime, err := GetSlatime(dbc, clientid.(int64), orgid.(int64), 2, diffid.(int64), recordid.(int64))
		// if err != nil {
		// 	return values, err
		// }
		// if len(slatime) != 0 {
		// 	res["startdatetimeresponse"] = slatime[0]["startdatetimeresponse"]
		// 	res["startdatetimeresolution"] = slatime[0]["startdatetimeresolution"]

		// } else {
		// 	res["startdatetimeresponse"] = ""
		// 	res["startdatetimeresolution"] = ""

		// }

		// c := whereMap[2].(map[string]interface{})
		// if c["val"].(string) == "STask" || c["val"].(string) == "CTask" {

		// 	parentticket, err := GetParentTicketofCtaskAndStask(dbc, clientid.(int64), orgid.(int64), recordid.(int64))
		// 	if err != nil {
		// 		return values, err
		// 	}
		// 	res["parentticket"] = parentticket[0]["parentticket"]
		// } else {
		// 	res["parentticket"] = ""

		// }

		//		changetime := []string{"lastupdateddatetime", "latestresodatetime"}
		// changetime := []string{"lastupdateddatetime", "latestresodatetime", "firstresponsedatetime"}

		// for j := 0; j < len(changetime); j++ {

		// 	_, found := res[changetime[j]] // v == 3.14  found == true
		// 	if found == true {
		// 		date := res[changetime[j]].(*fileutils.NullString)
		// 		dateTime, _ := date.Value()
		// 		if clientid != 0 && orgid != 0 && dateTime != "" && dateTime != nil {

		// 			res[changetime[j]], _ = Getexacttime(clientid, orgid, dateTime.(string), dbc.DB)
		// 			// if err != nil {
		// 			// 	logger.Log.Println("time change error", err)
		// 			// 	return "", "", false, errors.New("ERROR: Time chanege error"), "Something Went Error"

		// 			// }
		// 		}
		// 	}
		// }
		// secondtoday := []string{"calendaraging"}
		// for j := 0; j < len(secondtoday); j++ {

		// 	_, found := res[secondtoday[j]] // v == 3.14  found == true
		// 	if found == true {
		// 		time := res[secondtoday[j]].(*fileutils.NullInt64)
		// 		timesec, _ := time.Value()
		// 		if clientid != 0 && orgid != 0 && timesec != "" && timesec != nil {

		// 			res[secondtoday[j]] = GetSecondToDay((math.Abs(float64(timesec.(int64)))))
		// 		}
		// 	}
		// }
		// makeabsoluteofint := []string{"worknotenotupdated"}
		// for j := 0; j < len(makeabsoluteofint); j++ {

		// 	_, found := res[makeabsoluteofint[j]] // v == 3.14  found == true
		// 	if found == true {
		// 		time := res[makeabsoluteofint[j]].(*fileutils.NullInt64)
		// 		timesec, _ := time.Value()
		// 		if clientid != 0 && orgid != 0 && timesec != "" && timesec != nil {

		// 			res[makeabsoluteofint[j]] = int64(math.Abs(math.Abs(float64(timesec.(int64)))))

		// 		}
		// 	}
		// }
		// makeabsoluteoffloat := []string{"resooverdueperc", "respoverdueperc", "resolutionslameterpercentage", "responseslameterpercentage"}
		// for j := 0; j < len(makeabsoluteoffloat); j++ {

		// 	_, found := res[makeabsoluteoffloat[j]] // v == 3.14  found == true
		// 	if found == true {
		// 		time := res[makeabsoluteoffloat[j]].(*fileutils.NullFloat64)
		// 		timesec, _ := time.Value()
		// 		if clientid != 0 && orgid != 0 && timesec != "" && timesec != nil {

		// 			//					res[makeabsoluteoffloat[j]] = int64(math.Abs(float64(timesec.(float64)
		// 			res[makeabsoluteoffloat[j]] = int64(math.Round(math.Abs(float64(timesec.(float64)))))

		// 		}
		// 	}
		// }
		// secondtohourmin := []string{"businessaging", "actualeffort"}
		// for j := 0; j < len(secondtohourmin); j++ {

		// 	_, found := res[secondtohourmin[j]] // v == 3.14  found == true
		// 	if found == true {
		// 		time := res[secondtohourmin[j]].(*fileutils.NullInt64)
		// 		timesec, _ := time.Value()
		// 		if clientid != 0 && orgid != 0 && timesec != "" && timesec != nil {

		// 			res[secondtohourmin[j]] = GetSecondToHourMin(int64(math.Abs(float64(timesec.(int64)))))

		// 		}
		// 	}
		// }
		// secondtomin := []string{"responsetime", "respoverduetime", "resooverduetime", "userreplytimetaken", "slaidletime", "resotimeexcludeidletime", "followuptimetaken"}
		// for j := 0; j < len(secondtomin); j++ {

		// 	_, found := res[secondtomin[j]] // v == 3.14  found == true
		// 	if found == true {
		// 		time := res[secondtomin[j]].(*fileutils.NullInt64)
		// 		timesec, _ := time.Value()
		// 		if clientid != 0 && orgid != 0 && timesec != "" && timesec != nil {

		// 			res[secondtomin[j]] = GetSecondToMin(int64(math.Abs(float64(timesec.(int64)))))

		// 		}
		// 	}
		// }
		values = append(values, res)
	}

	aftertime := time.Now()
	logger.Log.Println("YYYY.MM.DD 1111111111111 : ", aftertime.Format("2006.01.02 15:04:05 .999"))

	return values, nil
}
func GetSlatime(mdao DbConn, ClientID int64, Mstorgnhirarchyid int64, RecorddifftypeID int64, RecorddifID int64, RecordID int64) ([]map[string]interface{}, error) {
	var sql = "SELECT startdatetimeresponse,startdatetimeresolution  FROM mstsladue where  therecordid=?"
	values := []map[string]interface{}{}
	logger.Log.Println(sql)
	logger.Log.Println("PARAMETER", RecordID)
	rows, err := mdao.DB.Query(sql, RecordID)
	if err != nil {
		logger.Log.Println("GetLaststagevalue Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	// logger.Log.Println(rows)
	for rows.Next() {
		fb := fileutils.NewDbFieldBind()
		err = fb.PutFields(rows)
		if err != nil {
			return values, err
		}
		rows.Scan(fb.GetFieldPtrArr()...)
		res := fb.GetFieldArr()
		values = append(values, res)

	}
	return values, nil
}
/*func GenerateCatQuery(req []interface{}) (string, []interface{}) {
	var queryStr string
	pint := []interface{}{}
	queryStr = `recordfulldetails.recordid IN (SELECT
			DISTINCT f.recordid
		FROM
			recordfulldetails f,
			mstrecordtype a,
			mstrecorddifferentiationtype b,
			maprecordtorecorddifferentiation c,
			mstrecorddifferentiation d
		WHERE
			a.torecorddifftypeid = b.id
				AND a.clientid = f.clientid
				AND a.mstorgnhirarchyid = recordfulldetails.mstorgnhirarchyid`
	// AND a.deleteflg = 0
	// AND a.activeflg = 1
	// AND a.fromrecorddifftypeid = 2
	// AND a.fromrecorddiffid = f.tickettypeid
	// AND b.parentid = 1
	// AND c.recordid = f.recordid
	// AND c.islatest = 1
	// AND a.torecorddifftypeid = c.recorddifftypeid
	// AND c.recorddiffid = d.id
	if len(req) == 1 {
		queryStr = queryStr + " AND d.name LIKE '%"

		b := req[0]
		c := b.(map[string]interface{})
		logger.Log.Println("cate:", c)
		queryStr = queryStr + c["val"].(string) + "%'" + " AND d.seqno=? "
		pint = append(pint, int64(c["seq"].(float64)))
		// int64(requestData["clientid"].(float64))
	} else {
		queryStr = queryStr + " AND d.parentcategorynames LIKE '%"
		for _, v := range req {
			c := v.(map[string]interface{})

			queryStr = queryStr + c["val"].(string) + "%"
			// }
		}
		queryStr = queryStr + "'"
	}
	queryStr = queryStr + ` AND a.deleteflg = 0
                            AND a.activeflg = 1
                            AND a.fromrecorddifftypeid = 2
                            AND a.fromrecorddiffid = f.tickettypeid
                            AND b.parentid = 1
                            AND c.recordid = f.recordid
                            AND c.islatest = 1
                            AND a.torecorddifftypeid = c.recorddifftypeid
                            AND c.recorddiffid = d.id)`
	return queryStr, pint
}*/
func GenerateCatQuery(req []interface{}) (string, []interface{}) {
	var queryStr string
	pint := []interface{}{}
	queryStr = `recordfulldetails.recordid IN  (SELECT DISTINCT
		c.recordid
	FROM
		mstrecordtype a,
		mstrecorddifferentiationtype b,
		maprecordtorecorddifferentiation c,
		mstrecorddifferentiation d
	WHERE
		d.clientid = recordfulldetails.clientid
			AND d.mstorgnhirarchyid =recordfulldetails.mstorgnhirarchyid`
	// AND a.deleteflg = 0
	// AND a.activeflg = 1
	// AND a.fromrecorddifftypeid = 2
	// AND a.fromrecorddiffid = f.tickettypeid
	// AND b.parentid = 1
	// AND c.recordid = f.recordid
	// AND c.islatest = 1
	// AND a.torecorddifftypeid = c.recorddifftypeid
	// AND c.recorddiffid = d.id
	if len(req) == 1 {
		queryStr = queryStr + " AND d.name LIKE '%"

		b := req[0]
		c := b.(map[string]interface{})
		logger.Log.Println("cate:", c)
		queryStr = queryStr + c["val"].(string) + "%'" + " AND d.seqno=? "
		pint = append(pint, int64(c["seq"].(float64)))
		// int64(requestData["clientid"].(float64))
	} else {
		queryStr = queryStr + " AND d.parentcategorynames LIKE '%"
		for _, v := range req {
			c := v.(map[string]interface{})

			queryStr = queryStr + c["val"].(string) + "%"
			// }
		}
		queryStr = queryStr + "'"
	}
	queryStr = queryStr + ` AND d.id = c.recorddiffid
	and c.recordid=recordfulldetails.recordid
	AND c.islatest = 1
	AND c.recorddifftypeid = a.torecorddifftypeid
	AND a.deleteflg = 0
	AND a.activeflg = 1
	AND a.fromrecorddifftypeid = 2
	AND a.torecorddifftypeid = b.id
	AND a.fromrecorddiffid=recordfulldetails.tickettypeid
	AND b.parentid = 1)`
	return queryStr, pint
}

func GetStateReson(mdao DbConn, ClientID int64, Mstorgnhirarchyid int64, RecorddifftypeID int64, RecorddifID int64, RecordID int64) ([]map[string]interface{}, error) {
	var sql = "SELECT distinct c.recorddiffid,c.recordtermid,d.id,d.termname,e.recordtrackvalue FROM recordfulldetails a,mststateterm b,mststateterm c,mstrecordterms d,trnreordtracking e where a.recordid=? AND a.tickettypeid=b.recorddiffid AND b.recorddifftypeid=2 AND c.recorddifftypeid=3 AND a.statusid=c.recorddiffid AND c.recordtermid=d.id AND a.recordid=e.recordid AND d.id=e.recordtermid"
	values := []map[string]interface{}{}
	// logger.Log.Println(sql)
	// logger.Log.Println("PARAMETER", RecordID)
	rows, err := mdao.DB.Query(sql, RecordID)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetLaststagevalue Get Statement Prepare Error", err)
		return values, err
	}
	// logger.Log.Println(rows)
	for rows.Next() {
		fb := fileutils.NewDbFieldBind()
		err = fb.PutFields(rows)
		if err != nil {
			return values, err
		}
		rows.Scan(fb.GetFieldPtrArr()...)
		res := fb.GetFieldArr()
		values = append(values, res)

	}
	return values, nil
}
func Getvisiblecomment(mdao DbConn, ClientID int64, Mstorgnhirarchyid int64, RecorddifftypeID int64, RecorddifID int64, RecordID int64) ([]map[string]interface{}, error) {
	var sql = "SELECT distinct a.recordtrackvalue AS Comment,FROM_UNIXTIME(a.createddate) AS Createdate FROM trnreordtracking a,mstrecordterms b WHERE a.recordid=? AND a.recordtermid =b.id AND b.seq=11 order by a.createddate desc limit 1"
	values := []map[string]interface{}{}
	// logger.Log.Println(sql)
	// logger.Log.Println("PARAMETER", RecordID)
	rows, err := mdao.DB.Query(sql, RecordID)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetLaststagevalue Get Statement Prepare Error", err)
		return values, err
	}
	// logger.Log.Println(rows)
	for rows.Next() {
		fb := fileutils.NewDbFieldBind()
		err = fb.PutFields(rows)
		if err != nil {
			return values, err
		}
		rows.Scan(fb.GetFieldPtrArr()...)
		res := fb.GetFieldArr()
		values = append(values, res)

	}
	return values, nil
}
func GetCategoryNames(mdao DbConn, ClientID int64, Mstorgnhirarchyid int64, RecorddifftypeID int64, RecorddifID int64, RecordID int64) ([]map[string]interface{}, error) {
	var sql = "SELECT distinct (SELECT typename from mstrecorddifferentiationtype WHERE id=a.torecorddifftypeid) lable,d.name,a.torecorddifftypeid as lebelid FROM mstrecordtype a, mstrecorddifferentiationtype b,maprecordtorecorddifferentiation c,mstrecorddifferentiation d WHERE a.torecorddifftypeid = b.id AND a.clientid=? AND a.mstorgnhirarchyid=? AND a.deleteflg=0 AND a.activeflg=1 AND a.fromrecorddifftypeid=? AND a.fromrecorddiffid=? AND b.parentid=1 AND c.recordid=? AND c.islatest=1 AND a.torecorddifftypeid=c.recorddifftypeid and c.recorddiffid=d.id"
	values := []map[string]interface{}{}
	// logger.Log.Println(sql)
	// logger.Log.Println("PARAMETER", ClientID, Mstorgnhirarchyid, RecorddifftypeID, RecorddifID, RecordID)
	rows, err := mdao.DB.Query(sql, ClientID, Mstorgnhirarchyid, RecorddifftypeID, RecorddifID, RecordID)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetLaststagevalue Get Statement Prepare Error", err)
		return values, err
	}
	// logger.Log.Println(rows)
	for rows.Next() {
		fb := fileutils.NewDbFieldBind()
		err = fb.PutFields(rows)
		if err != nil {
			return values, err
		}
		rows.Scan(fb.GetFieldPtrArr()...)
		res := fb.GetFieldArr()
		values = append(values, res)

	}
	return values, nil
}
func (dbc DbConn) RecordGridCountOnly(req map[string]interface{}) (int64, error) {
	logger.Log.Println("In side RecordGridCountOnly")
	var values int64
	pint := []interface{}{}
	var queryStr = "SELECT COUNT(DISTINCT recordfulldetails.recordid) total FROM recordfulldetails "
	queryStr = queryStr + " WHERE recordfulldetails.activeflg=1 AND recordfulldetails.deleteflg=0 "

	if req["where"] != nil {
		whereMap := req["where"].([]interface{})
		qstr, prm := GenerateCustomWhereQuery(whereMap, dbc.DB)
		queryStr = queryStr + " " + qstr
		pint = append(pint, prm...)
	}
	if req["cat"] != nil {
		catMap := req["cat"].([]interface{})
		if len(catMap) > 0 {
			qstr, prm := GenerateCatQuery(catMap)
			queryStr = queryStr + " AND " + qstr
			pint = append(pint, prm...)
		}
	}
	logger.Log.Println("Query:===>", queryStr)
	logger.Log.Println("Params:===>", pint)

	err := dbc.DB.QueryRow(queryStr, pint...).Scan(&values)
	switch err {
	case sql.ErrNoRows:
		values = 0
		return values, nil
	case nil:
		return values, nil
	default:
		logger.Log.Println("RecordGridCountOnly Get Statement Prepare Error", err)
		return values, err

	}
}
func GenerateCustomWhereOrQuery(req []interface{}, db *sql.DB) (string, []interface{}) {
	var queryStr string
	pint := []interface{}{}
	queryStr = ""
	for i, v := range req {
		c := v.(map[string]interface{})
		x := c["op"].(string)
		if i > 0 {
			queryStr = queryStr + " OR "
		}
		switch x {
		case ">", "<", ">=", "<=", "!=", "=":
			pint = append(pint, c["val"].(string))
			queryStr = queryStr + " recordfulldetails." + c["field"].(string) + " " + c["op"].(string) + " " + "? "
		case "like":
			queryStr = queryStr + " recordfulldetails." + c["field"].(string) + " LIKE '%" + c["val"].(string) + "%' "
		case "in":
			rawval := c["val"].(string)
			var splitval = strings.Split(rawval, ",")
			var joinedval = " ('" + strings.Join(splitval, "','") + "') "
			queryStr = queryStr + " recordfulldetails." + c["field"].(string) + " IN " + joinedval
		case "notin":
			rawval := c["val"].(string)
			var splitval = strings.Split(rawval, ",")
			var joinedval = " ('" + strings.Join(splitval, "','") + "') "
			queryStr = queryStr + " recordfulldetails." + c["field"].(string) + " NOT IN " + joinedval
		case "between":
			rawval := c["val"].(string)
			var splitval = strings.Split(rawval, ",")
			if len(splitval) == 2 {
				queryStr = queryStr + " recordfulldetails." + c["field"].(string) + " BETWEEN  ? AND ? "
				if strings.EqualFold(c["field"].(string), "latestresodatetime") {
					splitval[0], _ = GetexacttimeIstToUtc(2, 2, splitval[0], db)
					splitval[1], _ = GetexacttimeIstToUtc(2, 2, splitval[1], db)
				}
				pint = append(pint, splitval[0])
				pint = append(pint, splitval[1])
			}
		}

	}

	return queryStr, pint
}
func GetexacttimeIstToUtc(clientid interface{}, mstorgnhierarchyid interface{}, datetime string, db *sql.DB) (string, error) {
	// lock.Lock()
	// defer lock.Unlock()
	// db, err := config.ConnectMySqlDbSingleton()
	// if err != nil {
	logger.Log.Println("TIME", clientid, mstorgnhierarchyid, datetime)
	// 	return "Something Went Wrong", err
	// }
	// dataAccess := dao.DbConn{DB: db}
	// tz := entities.UtilityEntity{}
	// tz.Clientid = clientid
	// tz.Mstorgnhirarchyid = mstorgnhierarchyid
	err1, util := Gettimediffbyinterface(clientid, mstorgnhierarchyid, db)
	if err1 != nil {
		return "Something Went Wrong", err1
	}
	// t, err1 := dataAccess.Gettimediff(clientID, orgnID)
	// if err1 != nil {
	// 	return t, false, err1, "Something Went Wrong"
	// }
	layout := "2006-01-02 15:04:05"
	parsetime, err := time.Parse(layout, datetime)
	if err != nil {
		// logger.Log.Println("parsetime error:", err, datetime)
		return "", err

	}
	// logger.Log.Println("Parsetime is", parsetime)
	// unixtime := parsetime.Unix()
	logger.Log.Println("addtime:", util[0].Timediff)
	// time := dao.Convertdate(int64(parsetime.Unix()), Timediff)
	// logger.Log.Println("Time before:" + datetime + "   Time Now:" + time)
	unixTime := int64(parsetime.Unix()) - util[0].Timediff
	// logger.Log.Println("unixtime:>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>", unixTime)

	t := time.Unix(unixTime, 0)
	// return t.Format("02-Jan-2006 15:04:05"), nil
	return t.Format(layout), nil
}
func GenerateCustomWhereQuery(req []interface{}, db *sql.DB) (string, []interface{}) {
	var queryStr string
	pint := []interface{}{}
	queryStr = ""
	for _, v := range req {
		c := v.(map[string]interface{})
		x := c["op"].(string)
		switch x {
		case ">", "<", ">=", "<=", "!=", "=":
			pint = append(pint, c["val"].(string))
			queryStr = queryStr + " AND recordfulldetails." + c["field"].(string) + " " + c["op"].(string) + " " + "? "
		case "like":
			queryStr = queryStr + " AND recordfulldetails." + c["field"].(string) + " LIKE '%" + c["val"].(string) + "%' "
		case "in":
			rawval := c["val"].(string)
			var splitval = strings.Split(rawval, ",")
			var joinedval = " ('" + strings.Join(splitval, "','") + "') "
			queryStr = queryStr + " AND recordfulldetails." + c["field"].(string) + " IN " + joinedval
		case "notin":
			rawval := c["val"].(string)
			var splitval = strings.Split(rawval, ",")
			var joinedval = " ('" + strings.Join(splitval, "','") + "') "
			queryStr = queryStr + " AND recordfulldetails." + c["field"].(string) + " NOT IN " + joinedval
		case "between":
			rawval := c["val"].(string)
			var splitval = strings.Split(rawval, ",")
			if len(splitval) == 2 {
				queryStr = queryStr + " AND recordfulldetails." + c["field"].(string) + " BETWEEN  ? AND ? "
				if strings.EqualFold(c["field"].(string), "latestresodatetime") {
					splitval[0], _ = GetexacttimeIstToUtc(2, 2, splitval[0], db)
					splitval[1], _ = GetexacttimeIstToUtc(2, 2, splitval[1], db)
				}
				pint = append(pint, splitval[0])
				pint = append(pint, splitval[1])
			}
		case "or":
			rawval := c["val"].([]interface{})
			resQuery, resParams := GenerateCustomWhereOrQuery(rawval, db)
			pint = append(pint, resParams...)
			queryStr = queryStr + " AND ( " + resQuery + " ) "
		}

	}

	return queryStr, pint
}

func GenerateCustomOrderQuery(req []interface{}) string {
	var queryStr string
	for i, v := range req {
		if i == 0 {
			queryStr = queryStr + " ORDER BY "
		} else {
			queryStr = queryStr + " , "
		}
		c := v.(map[string]interface{})
		queryStr = queryStr + " recordfulldetails." + c["field"].(string) + " " + c["dir"].(string)
	}
	logger.Log.Println(" WHERE Query:===>", queryStr)
	return queryStr
}
func Gettimediffbyinterface(Clientid interface{}, Mstorgnhirarchyid interface{}, db *sql.DB) (error, []entities.UtilityEntity) {
	requestIds := []entities.UtilityEntity{}
	stmt, err := db.Prepare("SELECT a.utcdiff as timediff,c.utcdiff as reporttimediff,coalesce(d.example,'') Timeformat,coalesce(e.example,'') Reporttimeformat from zone a,mstorgnhierarchy b,zone c,mstdatetimeformat d,mstdatetimeformat e where a.zone_id=b.timezoneid and b.reporttimezoneid=c.zone_id and b.clientid=? and b.id=? and b.timeformatid=d.id and b.reporttimeformatid=e.id")
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
func Getexacttime(clientid interface{}, mstorgnhierarchyid interface{}, datetime string, db *sql.DB) (string, error) {
	// lock.Lock()
	// defer lock.Unlock()
	// db, err := config.ConnectMySqlDbSingleton()
	// if err != nil {
	logger.Log.Println("TIME", clientid, mstorgnhierarchyid, datetime)
	// 	return "Something Went Wrong", err
	// }
	// dataAccess := dao.DbConn{DB: db}
	// tz := entities.UtilityEntity{}
	// tz.Clientid = clientid
	// tz.Mstorgnhirarchyid = mstorgnhierarchyid
	err1, util := Gettimediffbyinterface(clientid, mstorgnhierarchyid, db)
	if err1 != nil {
		return "Something Went Wrong", err1
	}
	// t, err1 := dataAccess.Gettimediff(clientID, orgnID)
	// if err1 != nil {
	// 	return t, false, err1, "Something Went Wrong"
	// }
	layout := "2006-01-02 15:04:05"
	parsetime, err := time.Parse(layout, datetime)
	if err != nil {
		// logger.Log.Println("parsetime error:", err, datetime)
		return "", err

	}
	// logger.Log.Println("Parsetime is", parsetime)
	// unixtime := parsetime.Unix()
	logger.Log.Println("addtime:", util[0].Timediff)
	// time := dao.Convertdate(int64(parsetime.Unix()), Timediff)
	// logger.Log.Println("Time before:" + datetime + "   Time Now:" + time)
	unixTime := int64(parsetime.Unix()) + util[0].Timediff
	// logger.Log.Println("unixtime:>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>", unixTime)

	t := time.Unix(unixTime, 0)
	// return t.Format("02-Jan-2006 15:04:05"), nil
	return t.Format(layout), nil

}
func Getclientadnorgbyuser(userid int64, db *sql.DB) (error, []entities.ClientOrgEntity) {
	requestIds := []entities.ClientOrgEntity{}
	stmt, err := db.Prepare("SELECT a.clientid as clientid,a.mstorgnhirarchyid as mstorgnhirarchyid from mstclientuser a  where a.id =? and a.activeflag=1 and a.deleteflag=0")
	if err != nil {
		logger.Log.Print("Gettimediff Statement Prepare Error", err)
		// log.Print("Gettimediff Statement Prepare Error", err)
		return err, requestIds
	}
	defer stmt.Close()
	rows, err := stmt.Query(userid)
	if err != nil {
		logger.Log.Print("Gettimediff Statement Execution Error", err)
		// log.Print("Gettimediff Statement Execution Error", err)
		return err, requestIds
	}
	for rows.Next() {
		value := entities.ClientOrgEntity{}
		rows.Scan(&value.Clientid, &value.Mstorgnhirarchyid)
		requestIds = append(requestIds, value)
	}
	return nil, requestIds
}
func GetSecondToDay(sec float64) string {
	fmt.Println("Input:", sec)
	day := math.Round(sec / 86400)
	// fmt.Println("output:", fmt.Sprintf("%.2f", day))

	return fmt.Sprintf("%v", day)
}
func GetSecondToHourMin(sec int64) string {
	// fmt.Println("Input:", sec)
	day := sec / 3600
	modsec := sec % 3600
	// var y float64 = float64(modsec)
	min := modsec / 60
	// hourmin:=""+day+min
	if min < 10 {

		// fmt.Println("output:", fmt.Sprintf("%v:0%v", day, min))
		return fmt.Sprintf("%v:0%v", day, min)

	}
	// fmt.Println("output:", fmt.Sprintf("%v:%v", day, min))

	return fmt.Sprintf("%v:%v", day, min)
}
func GetSecondToMin(sec float64) string {
	// fmt.Println("Input:", sec)
	min := math.Round(sec / 60)
	return fmt.Sprintf("%v", min)
}
func GetParentTicketofCtaskAndStask(mdao DbConn, ClientID int64, Mstorgnhirarchyid int64, RecordID int64) ([]map[string]interface{}, error) {
	var sql = "SELECT b.code as parentticket FROM mstparentchildmap a,trnrecord b where a.clientid=? and a.mstorgnhirarchyid=? and a.childrecordid=? and a.parentrecordid=b.id and a.deleteflg=0 and b.deleteflg=0;"
	values := []map[string]interface{}{}
	// logger.Log.Println(sql)
	// logger.Log.Println("PARAMETER", RecordID)
	rows, err := mdao.DB.Query(sql, ClientID, Mstorgnhirarchyid, RecordID)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetLaststagevalue Get Statement Prepare Error", err)
		return values, err
	}
	// logger.Log.Println(rows)
	for rows.Next() {
		fb := fileutils.NewDbFieldBind()
		err = fb.PutFields(rows)
		if err != nil {
			return values, err
		}
		rows.Scan(fb.GetFieldPtrArr()...)
		res := fb.GetFieldArr()
		values = append(values, res)

	}
	return values, nil
}

// func Getabsolute(sec int64) string {
// 	fmt.Println("Input:", sec)

// 	day := float64(sec) / 86400
// 	// fmt.Println("output:", fmt.Sprintf("%.2f", day))

// 	return fmt.Sprintf("%.2f", day)
// }
