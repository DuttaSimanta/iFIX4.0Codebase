package dao
import (
	//"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
) 

var getMstClientCredentialType="SELECT id as id,typename as typename from mstclientcredentialtype"
func (dbc DbConn) GetAllMstClientCredentialType() ([]entities.MstClientCredentialTypeEntity, error) {
	logger.Log.Println("In side GetAllMstExcelTemplateType")
	values := []entities.MstClientCredentialTypeEntity{}

	rows, err := dbc.DB.Query(getMstClientCredentialType)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstClientCredentialType Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstClientCredentialTypeEntity{}
		rows.Scan(&value.Id, &value.TypeName)
		values = append(values, value)
	}
	return values, nil
}