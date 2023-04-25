package fileutils

import (
	"errors"
	"fmt"
	"os"
	"src/entities"
	"src/logger"
	Logger "src/logger"
	"strings"

	"github.com/tealeg/xlsx"
)

func JsonToExcelConverter(tz entities.JsonToExcelResponse, page *entities.ResultSetRequestEntity, datetime string) (string, string, error) {
	//filePath := "/home/mantech/Downloads/JsonToExcelDownload/details.xlsx"
	//keys := make([]string, len(tz.Details.RequestResultsetData[0]))
	contextPath, contextPatherr := os.Getwd() //getContextPath()
	logger.Log.Print("contextpath->", contextPath)
	if contextPatherr != nil {
		Logger.Log.Println(contextPatherr)
		return "", "", contextPatherr
	}

	filePath := contextPath + "/resource/downloads/iFIXReport_" + datetime + ".xlsx"
	logger.Log.Println("FILEPATH:", filePath)
	keys := page.Headers
	logger.Log.Print(len(keys))
	// for k, _ := range tz.Details.RequestResultsetData[0] {

	// 	keys = append(keys, k)
	// }
	logger.Log.Println(keys)
	file := xlsx.NewFile()
	sheet, sheetErr := file.AddSheet("Sheet")
	if sheetErr != nil {
		logger.Log.Print(sheetErr)

		//fmt.Printf(err.Error())
		return "", "", errors.New("ERROR: sheet adding error")
	}
	// var tempindex int
	// tempheader := []string{} //{"Response SLA Start Time", "Resolution SLA Start Time"}
	// var count int = 0
	for i := 0; i <= len(tz.Details.RequestResultsetData); i++ {
		row := sheet.AddRow()
		if i == 0 {
			for j := 0; j < len(page.HeadersDisplay); j++ {
				// if strings.EqualFold(page.HeadersDisplay[j], "Response SLA Start Date Time") || strings.EqualFold(page.HeadersDisplay[j], "Resolution SLA Start Date Time") {
				// 	tempheader = append(tempheader, page.HeadersDisplay[j])
				// 	continue
				// } else if strings.EqualFold(page.HeadersDisplay[j], "Status Reason") {
				// 	tempindex = j
				// 	break
				// }
				cell := row.AddCell()
				cell.Value = page.HeadersDisplay[j]
			}
			// for j := 0; j < len(tempheader); j++ {
			// 	cell := row.AddCell()
			// 	cell.Value = tempheader[j]
			// }
			// for j := tempindex; j < len(page.HeadersDisplay); j++ {
			// 	cell := row.AddCell()
			// 	cell.Value = page.HeadersDisplay[j]
			// }

		} else {
			//splittedParentCatagories := strings.Split(parentCategoryNames[i-1], "->") //(i-1) because for i=0 headernames is added
			for j := 5; j < len(keys); j++ {
				cell := row.AddCell()
				// cell.Value = fmt.Sprintf("%v", tz.Details.RequestResultsetData[i-1][keys[j]]) //tz.Details.RequestResultsetData[i-1][keys[j]].(string) //v.(string) // fmt.Sprint("%v", v)
				value := fmt.Sprintf("%v", tz.Details.RequestResultsetData[i-1][keys[j]]) //tz.Details.RequestResultsetData[i-1][keys[j]].(string) //v.(string) // fmt.Sprint("%v", v)
				if strings.EqualFold(value, "<nil>") {
					value = ""
				}
				cell.Value = value
			}
			// makeabsoluteofint := []string{"startdatetimeresponse"}
			// for j := 0; j < len(tempheader); j++ {

			// 	// _, found := tz.Details.RequestResultsetData[i-1][keys[j]] // v == 3.14  found == true
			// 	// if found == true {

			// 	cell := row.AddCell()
			// 	// cell.Value = fmt.Sprintf("%v", tz.Details.RequestResultsetData[i-1][keys[j]]) //tz.Details.RequestResultsetData[i-1][keys[j]].(string) //v.(string) // fmt.Sprint("%v", v)
			// 	value := fmt.Sprintf("%v", tz.Details.RequestResultsetData[i-1][makeabsoluteofint[0]]) //tz.Details.RequestResultsetData[i-1][keys[j]].(string) //v.(string) // fmt.Sprint("%v", v)
			// 	if strings.EqualFold(value, "<nil>") {
			// 		value = ""
			// 	}
			// 	cell.Value = value
			// 	// }
			// }
cell4 := row.AddCell()
			parentticket := ""
			parentticket = tz.Details.RequestResultsetData[i-1]["parentticket"].(string)
			if strings.EqualFold(parentticket, "<nil>") {
				parentticket = ""
			}
			cell4.Value = parentticket
			cell := row.AddCell()
			date := ""
			date = tz.Details.RequestResultsetData[i-1]["startdatetimeresponse"].(string)
			if strings.EqualFold(date, "<nil>") {
				date = ""
			}
			cell.Value = date
			cell1 := row.AddCell()
			date = tz.Details.RequestResultsetData[i-1]["startdatetimeresolution"].(string)
			if strings.EqualFold(date, "<nil>") {
				date = ""
			}
			cell1.Value = date
			statusreson := tz.Details.RequestResultsetData[i-1]["statusreson"].([]interface{})
			x := ""
			for i, v := range statusreson {
				c := v.(map[string]interface{})
				if i == 0 {
					x = x + c["termname"].(string) + ":" + c["recordtrackvalue"].(string)
				}
				x = x + "," + c["termname"].(string) + ":" + c["recordtrackvalue"].(string)
			}
			cell2 := row.AddCell()
			cell2.Value = fmt.Sprintf("%v", x) //tz.Details.RequestResultsetData[i-1][keys[j]].(string) //v.(string) // fmt.Sprint("%v", v)
			visiblecomment := tz.Details.RequestResultsetData[i-1]["visiblecomment"].([]interface{})
			comment := ""
			for i, v := range visiblecomment {
				c := v.(map[string]interface{})
				if i == 0 {
					comment = comment + c["Comment"].(string) + ":" + c["Createdate"].(string)
				}
				comment = comment + "," + c["Comment"].(string) + ":" + c["Createdate"].(string)
			}
			cell3 := row.AddCell()
			cell3.Value = fmt.Sprintf("%v", comment) //tz.Details.RequestResultsetData[i-1][keys[j]].(string) //v.(string) // fmt.Sprint("%v", v)

			cat := tz.Details.RequestResultsetData[i-1]["categories"].([]interface{})
			// logger.Log.Println("categories:", cat)

			for j := 0; j < len(cat); j++ {
				catval := cat[j].(map[string]interface{})
				cell := row.AddCell()
				// cell.Value = page.HeadersDisplay[j]
				cell.Value = fmt.Sprintf("%v", catval["name"]) //tz.Details.RequestResultsetData[i-1][keys[j]].(string) //v.(string) // fmt.Sprint("%v", v)
			}

		}
	}
	//logger.Log.Println("hiiiiiiii")
	saveErr := file.Save(filePath)
	if saveErr != nil {
		logger.Log.Print(saveErr)

		//fmt.Printf(err.Error())
		return "", "", errors.New("ERROR: File saving error")
	}
	return contextPath, filePath, nil
}
