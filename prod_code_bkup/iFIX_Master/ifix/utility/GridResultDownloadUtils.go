package utility

import (
	"errors"
	"fmt"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	Logger "iFIX/ifix/logger"
	"os"
	"strings"

	"github.com/tealeg/xlsx"
)

func GridResultToExcelConverter(tz entities.JsonToExcelGridResponse, page *entities.ResultGridRequestEntity) (string, string, error) {
	//filePath := "/home/mantech/Downloads/JsonToExcelDownload/details.xlsx"
	//keys := make([]string, len(tz.Details.RequestResultsetData[0]))
	contextPath, contextPatherr := os.Getwd() //getContextPath()
	logger.Log.Print("contextpath->", contextPath)
	if contextPatherr != nil {
		Logger.Log.Println(contextPatherr)
		return "", "", contextPatherr
	}
	filePath := contextPath + "/ifix/resource/downloads/iFIXGridData.xlsx"
	keys := page.Headers
	//logger.Log.Print(len(keys))
	// for k, _ := range tz.Details.RequestResultsetData[0] {

	// 	keys = append(keys, k)
	// }
	//logger.Log.Println(keys)
	file := xlsx.NewFile()
	sheet, sheetErr := file.AddSheet("Sheet")
	if sheetErr != nil {
		logger.Log.Print(sheetErr)

		//fmt.Printf(err.Error())
		return "", "", errors.New("ERROR: sheet adding error")
	}

	for i := 0; i <= len(tz.Details.RequestResultGridData); i++ {
		row := sheet.AddRow()
		if i == 0 {
			for j := 0; j < len(page.HeadersDisplay); j++ {
				cell := row.AddCell()
				cell.Value = page.HeadersDisplay[j]
			}
			// cat := tz.Details.RequestResultGridData[i]["categories"].([]interface{})
			// logger.Log.Println("categories:", cat)
			// for j := 0; j < len(cat); j++ {
			// 	catval := cat[j].(map[string]interface{})
			// 	cell := row.AddCell()
			// 	// cell.Value = page.HeadersDisplay[j]
			// 	cell.Value = fmt.Sprintf("%v", catval["leble"]) //tz.Details.RequestResultsetData[i-1][keys[j]].(string) //v.(string) // fmt.Sprint("%v", v)

			// }
		} else {
			//splittedParentCatagories := strings.Split(parentCategoryNames[i-1], "->") //(i-1) because for i=0 headernames is added
			for j := 5; j < len(keys); j++ {
				cell := row.AddCell()

				value := fmt.Sprintf("%v", tz.Details.RequestResultGridData[i-1][keys[j]]) //tz.Details.RequestResultsetData[i-1][keys[j]].(string) //v.(string) // fmt.Sprint("%v", v)
				if strings.EqualFold(value, "<nil>") {
					value = ""
				}
				cell.Value = value
			}
			cat := tz.Details.RequestResultGridData[i-1]["categories"].([]interface{})
			logger.Log.Println("categories:", cat)

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
