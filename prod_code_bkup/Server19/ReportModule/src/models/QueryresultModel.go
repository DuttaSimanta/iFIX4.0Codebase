package models

import (
	"math"
	"src/dao"
	"src/fileutils"

	// FileUtils "src/fileutils"
	"src/logger"
	// Excel "github.com/tealeg/xlsx"
)

func RecordGridResultOnly(page map[string]interface{}) (map[string]interface{}, bool, error, string) {
	var t map[string]interface{}
	var response map[string]interface{}
	// if mutexutility.MutexLocked(lock) == false {
	// 	lock.Lock()
	// 	defer lock.Unlock()
	// }
	// db, err := dbconfig.ConnectMySqlDb()
	// if err != nil {
	// 	logger.Log.Println("database connection failure", err)
	// 	return t, false, err, "Something Went Wrong"
	// }
	//defer db.Close()
	if db == nil {
		dbcon, err := ConnectMySqlDb()
		if err != nil {
			logger.Log.Println("Error in DBConnection")
			return t, false, err, "Something Went Wrong"
		}
		db = dbcon
	}
	dbc := dao.DbConn{DB: db}
	dataAccess := dao.DbConn{DB: db}
	response = make(map[string]interface{}, 2)
	result, err := dataAccess.RecordGridResultOnly(page)
	if err != nil {
		return t, false, err, "Something Went Wrong"
	}

	for i, _ := range result {
		record := result[i]["recordid"].(*fileutils.NullInt64)
		recordid, _ := record.Value()
		tickettypeid := result[i]["tickettypeid"].(*fileutils.NullInt64)
		diffid, _ := tickettypeid.Value()
		client := result[i]["clientid"].(*fileutils.NullInt64)
		clientid, _ := client.Value()
		org := result[i]["mstorgnhirarchyid"].(*fileutils.NullInt64)
		orgid, _ := org.Value()
		//To get Categories of a record=============================================
		categories, err := dao.GetCategoryNames(dbc, clientid.(int64), orgid.(int64), 2, diffid.(int64), recordid.(int64))
		if err != nil {
			return t, false, err, "Something Went Wrong"
		}
		result[i]["categories"] = categories
		//To get Status Reason==============================================
		statereson, err := dao.GetStateReson(dbc, clientid.(int64), orgid.(int64), 2, diffid.(int64), recordid.(int64))
		if err != nil {
			return t, false, err, "Something Went Wrong"
		}
		result[i]["statusreson"] = statereson
		//To get customervisible comment===========================================
		visiblecomment, err := dao.Getvisiblecomment(dbc, clientid.(int64), orgid.(int64), 2, diffid.(int64), recordid.(int64))
		if err != nil {
			return t, false, err, "Something Went Wrong"
		}
		result[i]["visiblecomment"] = visiblecomment
		//To get startdatetimeresponse , startdatetimeresolution==================================
		slatime, err := dao.GetSlatime(dbc, clientid.(int64), orgid.(int64), 2, diffid.(int64), recordid.(int64))
		if err != nil {
			return t, false, err, "Something Went Wrong"
		}
		if len(slatime) != 0 {
			result[i]["startdatetimeresponse"] = slatime[0]["startdatetimeresponse"]
			result[i]["startdatetimeresolution"] = slatime[0]["startdatetimeresolution"]

		} else {
			result[i]["startdatetimeresponse"] = ""
			result[i]["startdatetimeresolution"] = ""

		}
		// To get Parent Ticket of Stask And Ctask===================================
		var whereMap []interface{}
		if page["where"] != nil {
			whereMap = page["where"].([]interface{})
		}
		c := whereMap[2].(map[string]interface{})
		if c["val"].(string) == "STask" || c["val"].(string) == "CTask" {

			parentticket, err := dao.GetParentTicketofCtaskAndStask(dbc, clientid.(int64), orgid.(int64), recordid.(int64))
			if err != nil {
				return t, false, err, "Something Went Wrong"
			}
//			result[i]["parentticket"] = parentticket[0]["parentticket"]
                        if len(parentticket)>0{
				result[i]["parentticket"] = parentticket[0]["parentticket"]
			}else{
				result[i]["parentticket"] = ""
			}
		} else {
			result[i]["parentticket"] = ""

		}
		//To get utc to standard time========================================
		changetime := []string{"lastupdateddatetime", "latestresodatetime", "firstresponsedatetime"}

		for j := 0; j < len(changetime); j++ {

			_, found := result[i][changetime[j]] // v == 3.14  found == true
			if found {
				date := result[i][changetime[j]].(*fileutils.NullString)
				dateTime, _ := date.Value()
				if clientid != 0 && orgid != 0 && dateTime != "" && dateTime != nil {

					result[i][changetime[j]], _ = dao.Getexacttime(clientid, orgid, dateTime.(string), dbc.DB)
					// if err != nil {
					// 	logger.Log.Println("time change error", err)
					// 	return "", "", false, errors.New("ERROR: Time chanege error"), "Something Went Error"

					// }
				}
			}
		}
		//To get Convert Second to day==================================================
		secondtoday := []string{"calendaraging"}
		for j := 0; j < len(secondtoday); j++ {

			_, found := result[i][secondtoday[j]] // v == 3.14  found == true
			if found {
				time := result[i][secondtoday[j]].(*fileutils.NullInt64)
				timesec, _ := time.Value()
				if clientid != 0 && orgid != 0 && timesec != "" && timesec != nil {

					result[i][secondtoday[j]] = dao.GetSecondToDay((math.Abs(float64(timesec.(int64)))))
				}
			}
		}
		//To Make Absolute of Int value=========================================
		makeabsoluteofint := []string{"worknotenotupdated"}
		for j := 0; j < len(makeabsoluteofint); j++ {

			_, found := result[i][makeabsoluteofint[j]] // v == 3.14  found == true
			if found {
				time := result[i][makeabsoluteofint[j]].(*fileutils.NullInt64)
				timesec, _ := time.Value()
				if clientid != 0 && orgid != 0 && timesec != "" && timesec != nil {

					result[i][makeabsoluteofint[j]] = int64(math.Abs(math.Abs(float64(timesec.(int64)))))

				}
			}
		}
		// To make Absolute of Float value======================================================
		makeabsoluteoffloat := []string{"resooverdueperc", "respoverdueperc", "resolutionslameterpercentage", "responseslameterpercentage"}
		for j := 0; j < len(makeabsoluteoffloat); j++ {

			_, found := result[i][makeabsoluteoffloat[j]] // v == 3.14  found == true
			if found {
				time := result[i][makeabsoluteoffloat[j]].(*fileutils.NullFloat64)
				timesec, _ := time.Value()
				if clientid != 0 && orgid != 0 && timesec != "" && timesec != nil {

					//					result[i][makeabsoluteoffloat[j]] = int64(math.Abs(float64(timesec.(float64)
					result[i][makeabsoluteoffloat[j]] = int64(math.Round(math.Abs(float64(timesec.(float64)))))

				}
			}
		}
		//To Convert Second To Hour:Minute===================================
		secondtohourmin := []string{"actualeffort"}
		for j := 0; j < len(secondtohourmin); j++ {

			_, found := result[i][secondtohourmin[j]] // v == 3.14  found == true
			if found {
				time := result[i][secondtohourmin[j]].(*fileutils.NullInt64)
				timesec, _ := time.Value()
				if clientid != 0 && orgid != 0 && timesec != "" && timesec != nil {

					result[i][secondtohourmin[j]] = dao.GetSecondToHourMin(int64(math.Abs(float64(timesec.(int64)))))

				}
			}
		}
		//To Convert Second To Minute===================================
		secondtomin := []string{"responsetime", "respoverduetime", "resooverduetime", "userreplytimetaken", "slaidletime", "resotimeexcludeidletime", "followuptimetaken", "businessaging"}
		for j := 0; j < len(secondtomin); j++ {

			_, found := result[i][secondtomin[j]] // v == 3.14  found == true
			if found {
				time := result[i][secondtomin[j]].(*fileutils.NullInt64)
				timesec, _ := time.Value()
				if clientid != 0 && orgid != 0 && timesec != "" && timesec != nil {

					result[i][secondtomin[j]] = dao.GetSecondToMin((math.Abs(float64(timesec.(int64)))))

				}
			}
		}
	}

	total, err := dataAccess.RecordGridCountOnly(page)
	if err != nil {
		return t, false, err, "Something Went Wrong"
	}
	response["result"] = result

	response["total"] = total
	return response, true, nil, ""
}
