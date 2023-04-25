package dao
import (
	//"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
) 

var getMstExcelTemplateType= "SELECT a.id as Id, a.typename as TypeName FROM mstexceltemplatetype a ORDER BY a.id DESC"
 


func (dbc DbConn) GetAllMstExcelTemplateType() ([]entities.MstExcelTemplateTypeEntity, error) {
	logger.Log.Println("In side GetAllMstExcelTemplateType")
	values := []entities.MstExcelTemplateTypeEntity{}

	rows, err := dbc.DB.Query(getMstExcelTemplateType)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstExcelTemplateType Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstExcelTemplateTypeEntity{}
		rows.Scan(&value.Id, &value.TypeName)
		values = append(values, value)
	}
	return values, nil
}

 