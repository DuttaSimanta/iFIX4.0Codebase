package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"

	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/tealeg/xlsx"
)

func Getrecordfulldetailss(tz *entities.RecordfulldetailsRequestEntity) (entities.ResultEntity, bool, error, string) {
	t := entities.ResultEntity{}

	return t, false, nil, "Something Went Wrong"

}
func Getidsinstring(arrids []int64) string {
	var ids string = ""
	for i, id := range arrids {
		if i > 0 {
			ids += ","
		}
		ids += strconv.Itoa(int(id))
	}
	return ids
}

// func Getrecordfulldetails(tz *entities.RecordfulldetailsRequestEntity) (entities.ResultEntity, bool, error, string) {
// 	logger.Log.Println("In side Getrecordfulldetails")
// 	t := entities.ResultEntity{}
// 	lock.Lock()
// 	defer lock.Unlock()
// 	db, err := config.ConnectMySqlDbSingleton()
// 	if err != nil {
// 		logger.Log.Println("database connection failure", err)
// 		return t, false, err, "Something Went Wrong"
// 	}
// 	dataAccess := dao.DbConn{DB: db}
// 	var orgids string = Getidsinstring(tz.Mstorgnhirarchyids)
// 	var statusseqnos string = Getidsinstring(tz.Diffstatusseqnos)
// 	var tickettypeid = strconv.Itoa(int(tz.Tickettypeseq))
// 	tz.Seqno = 3
// 	statusidsarr, err1 := dataAccess.GetStatusIdsOfMultiOrg(tz, orgids, statusseqnos, tickettypeid)
// 	if err1 != nil {
// 		logger.Log.Println("Fail to get schema:", err1)
// 		return t, false, err1, "Something Went Wrong"
// 	}
// 	tz.Seqno = 2
// 	tickettypearr, err1 := dataAccess.GetStatusIdsOfMultiOrg(tz, orgids, statusseqnos, tickettypeid)
// 	if err1 != nil {
// 		logger.Log.Println("Fail to get schema:", err1)
// 		return t, false, err1, "Something Went Wrong"
// 	}
// 	var statusids = Getidsinstring(statusidsarr)
// 	var tickettypeids = Getidsinstring(tickettypearr)

// 	values, err1 := dataAccess.GetSchema()
// 	if err1 != nil {
// 		logger.Log.Println("Fail to get schema:", err1)
// 		return t, false, err1, "Something Went Wrong"
// 	}
// 	var value []string

// 	for i := 0; i < len(values); i++ {
// 		if (!strings.HasSuffix(values[i], "id") || (values[i] == "ticketid" || values[i] == "clientid" || values[i] == "mstorgnhirarchyid")) && !(values[i] == "activeflg" || values[i] == "deleteflg") {
// 			value = append(value, values[i])
// 		}
// 	}

// 	result, err1 := dataAccess.GetTableData(value, orgids, tz, statusids, tickettypeids)
// 	if err1 != nil {
// 		logger.Log.Println("fail to get table data", err1)
// 		return t, false, err1, "Something Went Wrong"
// 	}
// 	t.Recordfullresult = result
// 	t.Order = value
// 	return t, true, err, ""
// }
func Getrecordfulldetails(tz *entities.RecordfulldetailsRequestEntity, db *sql.DB) (entities.ResultEntity, bool, error, string) {
	logger.Log.Println("In side Getrecordfulldetails")
	t := entities.ResultEntity{}
	// lock.Lock()
	// defer lock.Unlock()
	// db, err := config.ConnectMySqlDbSingleton()
	// if err != nil {
	// 	logger.Log.Println("database connection failure", err)
	// 	return t, false, err, "Something Went Wrong"
	// }
	dataAccess := dao.DbConn{DB: db}
	var orgids string = Getidsinstring(tz.Mstorgnhirarchyids)
	var statusseqnos string = Getidsinstring(tz.Diffstatusseqnos)
	var tickettypeid = strconv.Itoa(int(tz.Tickettypeseq))
	tz.Seqno = 3
	statusidsarr, err1 := dataAccess.GetStatusIdsOfMultiOrg(tz, orgids, statusseqnos, tickettypeid)
	if err1 != nil {
		logger.Log.Println("Fail to get schema:", err1)
		return t, false, err1, "Something Went Wrong"
	}
	tz.Seqno = 2
	tickettypearr, err1 := dataAccess.GetStatusIdsOfMultiOrg(tz, orgids, statusseqnos, tickettypeid)
	if err1 != nil {
		logger.Log.Println("Fail to get schema:", err1)
		return t, false, err1, "Something Went Wrong"
	}
	var statusids = Getidsinstring(statusidsarr)
	var tickettypeids = Getidsinstring(tickettypearr)

	headerNames, headerErr := dao.GetTemplateHeaderNamesForValidation(db, tz.Clientid, tz.Mstorgnhirarchyid, 0)
	if headerErr != nil {
		fmt.Println(headerErr)
		logger.Log.Println(headerErr)
		return t, false, err1, "Something Went Wrong"
	}
	// var value []string

	// for i := 0; i < len(values); i++ {
	// 	if (!strings.HasSuffix(values[i], "id") || (values[i] == "ticketid" || values[i] == "clientid" || values[i] == "mstorgnhirarchyid")) && !(values[i] == "activeflg" || values[i] == "deleteflg") {
	// 		value = append(value, values[i])
	// 	}
	// }

	result, err1 := dataAccess.GetTableData(orgids, tz, statusids, tickettypeids)
	if err1 != nil {
		logger.Log.Println("fail to get table data", err1)
		return t, false, err1, "Something Went Wrong"
	}
	var categories []entities.Categorydetails
	var err error
	for i := 0; i < len(result); i++ {
		categories, err = dataAccess.GetExternalcategorynames(result[i].Clientid, result[i].Mstorgnhirarchyid, 2, result[i].Tickettypeid, result[i].Recordid)
		if err != nil {
			return t, false, err, "Something Went Wrong"
		}
		result[i].Category = categories

	}
	// result[i].Category = categories
	logger.Log.Println("RESULT:", result)
	t.Recordfullresult = result
	t.Order = headerNames
	return t, true, nil, ""
}

func RecordfulldetailsToExcelConverter(tz *entities.RecordfulldetailsRequestEntity) (string, string, bool, error, string) {
	logger.Log.Println("In side RecordfulldetailsToExcelConverter")

	contextPath, contextPatherr := os.Getwd()
	// props, err := utility.ReadPropertiesFile(contextPath + "/ifix/resource/application.properties")
	// if err != nil {
	// 	logger.Log.Println("Unable to Get URL From utility.ReadPropertiesFile", err)
	// 	return "", "", false, err, "Something Went Wrong"
	// }
	if contextPatherr != nil {
		logger.Log.Println(contextPatherr)
		return "", "", false, contextPatherr, "Contextpath ERROR"
	}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	result, _, err1, _ := Getrecordfulldetails(tz, db)
	if err1 != nil {
		logger.Log.Println("fail to get table data", err1)
		return "", "", false, err1, "No Records Found"
	}

	if len(result.Recordfullresult) == 0 {
		logger.Log.Println("No Data Is There")
		return "", "", false, errors.New("NO record found"), "No Records Found"
	}

	if err != nil {
		return "", "", false, err, "Something Went Wrong"
	}
	// changetime := []string{"lastupdateddate", "userreplieddatetime", "pendinguserdatetime", "latestresodatetime", "firstresodatetime", "latestresponsedatetime", "firstresponsedatetime", "followuprespdatetime", "followupdatetime", "reopendatetime", "lastupdateddatetime"}
	for i := 0; i < len(result.Recordfullresult); i++ {
		// for j := 0; j < len(changetime); j++ {
		datetime := fmt.Sprintf("%v", result.Recordfullresult[i].Latestresodatetime)
		clientid := result.Recordfullresult[i].Clientid
		orgid := result.Recordfullresult[i].Mstorgnhirarchyid
		if clientid != 0 && orgid != 0 && datetime != "" || datetime != "nil" {
			newdatetime, err := Getexacttime(clientid, orgid, datetime, db)
			if err != nil {
				logger.Log.Println("time change error", err)
				return "", "", false, errors.New("ERROR: Time chanege error"), "Something Went Error"

			}
			result.Recordfullresult[i].Latestresodatetime = newdatetime
		}
		// }
	}
	filePath := contextPath + "/ifix/resource/downloads/iFIXRecordfulldetails.xlsx"
	// keys := result.Order
	var fromdate = strings.Split(tz.Fromdate, " ")
	var todate = strings.Split(tz.Todate, " ")

	keys := []string{"Ticket Logged Date From", "Ticket Logged date To", "Customer Name", "Ticket No", "Ticket Title", "Status", "Resolved Date", "Vendor Name", "Vendor Ticket No"}
	logger.Log.Print(len(keys))
	logger.Log.Println(keys)
	file := xlsx.NewFile()
	sheet, sheetErr := file.AddSheet("Sheet")
	if sheetErr != nil {
		logger.Log.Print(sheetErr)
		return "", "", false, errors.New("ERROR: sheet adding error"), "ERROR: sheet adding error"
	}

	for i := 0; i <= len(result.Recordfullresult); i++ {
		row := sheet.AddRow()
		if i == 0 {
			for j := 0; j < len(keys); j++ {
				cell := row.AddCell()
				cell.Value = keys[j]
			}
			for k := 0; k < len(result.Recordfullresult[i].Category); k++ {
				cell := row.AddCell()
				cell.Value = fmt.Sprintf("%v", result.Recordfullresult[i].Category[k].Label) //tz.Details.RequestResultsetData[i-1][keys[j]].(string) //v.(string) // fmt.Sprint("%v", v)
			}
		} else {
			for j := 4; j < len(keys)+4; j++ {
				if j == 4 {
					cell := row.AddCell()
					cell.Value = fmt.Sprintf("%v", fromdate[0]) //tz.Details.RequestResultsetData[i-1][keys[j]].(string) //v.(string) // fmt.Sprint("%v", v)
				}
				if j == 5 {
					cell := row.AddCell()
					cell.Value = fmt.Sprintf("%v", todate[0]) //tz.Details.RequestResultsetData[i-1][keys[j]].(string) //v.(string) // fmt.Sprint("%v", v)
				}
				if j == 6 {
					cell := row.AddCell()
					cell.Value = fmt.Sprintf("%v", result.Recordfullresult[i-1].CustomerName) //tz.Details.RequestResultsetData[i-1][keys[j]].(string) //v.(string) // fmt.Sprint("%v", v)
				}
				if j == 7 {
					cell := row.AddCell()
					cell.Value = fmt.Sprintf("%v", result.Recordfullresult[i-1].TicketNo) //tz.Details.RequestResultsetData[i-1][keys[j]].(string) //v.(string) // fmt.Sprint("%v", v)
				}
				if j == 8 {
					cell := row.AddCell()
					cell.Value = fmt.Sprintf("%v", result.Recordfullresult[i-1].Shortdescription) //tz.Details.RequestResultsetData[i-1][keys[j]].(string) //v.(string) // fmt.Sprint("%v", v)
				}
				if j == 9 {
					cell := row.AddCell()
					cell.Value = fmt.Sprintf("%v", result.Recordfullresult[i-1].Status) //tz.Details.RequestResultsetData[i-1][keys[j]].(string) //v.(string) // fmt.Sprint("%v", v)
				}
				if j == 10 {
					cell := row.AddCell()
					cell.Value = fmt.Sprintf("%v", result.Recordfullresult[i-1].Latestresodatetime) //tz.Details.RequestResultsetData[i-1][keys[j]].(string) //v.(string) // fmt.Sprint("%v", v)
				}
				if j == 11 {
					cell := row.AddCell()
					cell.Value = fmt.Sprintf("%v", result.Recordfullresult[i-1].VendorName) //tz.Details.RequestResultsetData[i-1][keys[j]].(string) //v.(string) // fmt.Sprint("%v", v)
				}
				if j == 12 {
					cell := row.AddCell()
					cell.Value = fmt.Sprintf("%v", result.Recordfullresult[i-1].VendorTicketid) //tz.Details.RequestResultsetData[i-1][keys[j]].(string) //v.(string) // fmt.Sprint("%v", v)
				} // else {

				// }

			}
			for k := 0; k < len(result.Recordfullresult[i-1].Category); k++ {
				cell := row.AddCell()
				cell.Value = fmt.Sprintf("%v", result.Recordfullresult[i-1].Category[k].Categoryname) //tz.Details.RequestResultsetData[i-1][keys[j]].(string) //v.(string) // fmt.Sprint("%v", v)
			}
		}
	}
	saveErr := file.Save(filePath)
	if saveErr != nil {
		logger.Log.Print(saveErr)
		return "", "", false, errors.New("ERROR: File saving error"), "ERROR: File saving error"
	}
	// OriginalFileName, UploadedFileName, err := utility.FileUploadAPICall(1, 1, config.FileUploadUrl, filePath)
	// if err != nil {
	// 	logger.Log.Println("Error while downloading", "-", err)
	// }
	buf, _ := ioutil.ReadFile(filePath)
	typee := http.DetectContentType(buf)
	// OriginalFileName, UploadedFileName, err := utility.FileUploadAPICall(1, 1, config.FileUploadUrl, filePath)
	// if err != nil {
	// 	logger.Log.Println("Error while downloading", "-", err)
	// }
	//fmt.Println(t.OriginalFileName, t.UploadedFileName)
	var data = entities.FileuploadEntity{}
	data.Clientid = 1
	data.Mstorgnhirarchyid = 1
	// imgbytes := buf.Bytes()
	// OriginalFileName, UploadedFileName, err := utility.FileUploadAPICall(values[0].Clientid, values[0].Mstorgnhirarchyid, props["fileUploadUrl"], filePath)
	dataDetails, success, err, msg := UploadFileWithConn(&data, buf, filePath, typee, db) //utility.FileUploadAPICall(values[0].Clientid, values[0].Mstorgnhirarchyid, props["fileUploadUrl"], filePath)
	if err != nil {
		logger.Log.Println("Error while downloading", "-", err)
		return "", "", false, err, "Something Went Wrong"
	}
	logger.Log.Println("===========FileUploadMessage==============", msg)
	logger.Log.Println("===========FileUploadSuccss==============", success)

	OriginalFileName := filepath.Base(dataDetails.Originalfile)
	UploadedFileName := dataDetails.Filename

	logger.Log.Println("===========OriginalFileName==============", OriginalFileName)
	logger.Log.Println("===========UploadedFileName==============", UploadedFileName)
	return OriginalFileName, UploadedFileName, true, nil, ""
}

// 	return "OriginalFileName", "UploadedFileName", true, nil, ""

// }

func Getexacttime(clientid interface{}, mstorgnhierarchyid interface{}, datetime string, db *sql.DB) (string, error) {
	// lock.Lock()
	// defer lock.Unlock()
	// db, err := config.ConnectMySqlDbSingleton()
	// if err != nil {
	// 	logger.Log.Println("database connection failure", err)
	// 	return "Something Went Wrong", err
	// }
	dataAccess := dao.DbConn{DB: db}
	// tz := entities.UtilityEntity{}
	// tz.Clientid = clientid
	// tz.Mstorgnhirarchyid = mstorgnhierarchyid
	err1, util := dataAccess.Gettimediffbyinterface(clientid, mstorgnhierarchyid)
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
		return "", nil

	}
	// logger.Log.Println("Parsetime is", parsetime)
	// unixtime := parsetime.Unix()
	// logger.Log.Println("unixtime:", unixtime, Timediff)
	// time := dao.Convertdate(int64(parsetime.Unix()), Timediff)
	// logger.Log.Println("Time before:" + datetime + "   Time Now:" + time)
	unixTime := int64(parsetime.Unix()) + util[0].Timediff
	// logger.Log.Println("unixtime:>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>", unixTime)

	t := time.Unix(unixTime, 0)
	// return t.Format("02-Jan-2006 15:04:05"), nil
	return t.Format(layout), nil

}
func contains(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func GetRecordByDiffTypeOfMultiOrg(tz *entities.RecordfulldetailsRequestEntity) ([]entities.RecordDiffOfMultiOrgEntity, bool, error, string) {
	log.Println("In side model")
	t := []entities.RecordDiffOfMultiOrgEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	var orgids string = ""
	for i, orgid := range tz.Mstorgnhirarchyids {
		if i > 0 {
			orgids += ","
		}
		orgids += strconv.Itoa(int(orgid))
	}
	values, err1 := dataAccess.GetRecordByDiffTypeOfMultiOrg(tz, orgids)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}

	return values, true, err, ""
}
