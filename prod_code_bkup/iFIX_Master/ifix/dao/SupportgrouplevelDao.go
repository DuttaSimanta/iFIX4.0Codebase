package dao

import (
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

//var insertPortgrouplevel = "INSERT INTO supportgrouplevel (name) VALUES (?)"
//var duplicatePortgrouplevel = "SELECT count(id) total FROM  supportgrouplevel WHERE deleteflg = 0"
var getPortgrouplevel = "SELECT id as Id, name as Name FROM supportgrouplevel ORDER BY id "

//var getPortgrouplevelcount = "SELECT count(id) total FROM  supportgrouplevel WH deleteflg =0 and activeflg=1"
//var updatePortgrouplevel = "UPDATE supportgrouplevel SET name = ? WHERE "
//var deletePortgrouplevel = "UPDATE supportgrouplevel SET deleteflg = '1' WH"

// func (dbc DbConn) CheckDuplicatePortgrouplevel(tz *entities.PortgrouplevelEntity) (entities.PortgrouplevelEntities, error) {
//     logger.Log.Println("In side CheckDuplicatePortgrouplevel")
//     value := entities.PortgrouplevelEntities{}
//     err := dbc.DB.QueryRow(duplicatePortgrouplevel, ).Scan(&value.Total)
//     switch err {
//         case sql.ErrNoRows:
//             value.Total = 0
//             return value, nil
//         case nil:
//             return value, nil
//         default:
//             logger.Log.Println("CheckDuplicatePortgrouplevel Get Statement Prepare Error", err)
//             return value, err
//     }
// }

// func (dbc DbConn) InsertPortgrouplevel(tz *entities.PortgrouplevelEntity) (int64, error) {
//     logger.Log.Println("In side InsertPortgrouplevel")
//     logger.Log.Println("Query -->",insertPortgrouplevel)
//     stmt, err := dbc.DB.Prepare(insertPortgrouplevel)
//     defer stmt.Close()
//     if err != nil {
//         logger.Log.Println("InsertPortgrouplevel Prepare Statement  Error", err)
//         return 0, err
//     }
//     logger.Log.Println("Parameter -->",tz.Name)
//     res, err := stmt.Exec(tz.Name)
//     if err != nil {
//         logger.Log.Println("InsertPortgrouplevel Execute Statement  Error", err)
//         return 0, err
//     }
//     lastInsertedId, err := res.LastInsertId()
//     return lastInsertedId, nil
// }

func (dbc DbConn) GetAllPortgrouplevel() ([]entities.SupportgrouplevelEntity, error) {
	logger.Log.Println("In side GelAllPortgrouplevel")
	values := []entities.SupportgrouplevelEntity{}
	rows, err := dbc.DB.Query(getPortgrouplevel)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllPortgrouplevel Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.SupportgrouplevelEntity{}
		rows.Scan(&value.Id, &value.Name)
		values = append(values, value)
	}
	return values, nil
}

// func (dbc DbConn) UpdatePortgrouplevel(tz *entities.PortgrouplevelEntity) error {
//     logger.Log.Println("In side UpdatePortgrouplevel")
//     stmt, err := dbc.DB.Prepare(updatePortgrouplevel)
//     defer stmt.Close()
//     if err != nil {
//         logger.Log.Println("UpdatePortgrouplevel Prepare Statement  Error", err)
//         return err
//     }
//     _, err = stmt.Exec(tz.Name, tz.Id)
//     if err != nil {
//         logger.Log.Println("UpdatePortgrouplevel Execute Statement  Error", err)
//         return err
//     }
//     return nil
// }

// func (dbc DbConn) DeletePortgrouplevel(tz *entities.PortgrouplevelEntity) error {
//     logger.Log.Println("In side DeletePortgrouplevel")
//     stmt, err := dbc.DB.Prepare(deletePortgrouplevel)
//     defer stmt.Close()
//     if err != nil {
//         logger.Log.Println("DeletePortgrouplevel Prepare Statement  Error", err)
//         return err
//     }
//     _, err = stmt.Exec(tz.Id)
//     if err != nil {
//         logger.Log.Println("DeletePortgrouplevel Execute Statement  Error", err)
//         return err
//     }
//     return nil
// }

// func (dbc DbConn) GetPortgrouplevelCount(tz *entities.PortgrouplevelEntity) (entities.PortgrouplevelEntities, error) {
//     logger.Log.Println("In side GetPortgrouplevelCount")
//     value := entities.PortgrouplevelEntities{}
//     err := dbc.DB.QueryRow(getPortgrouplevelcount, ).Scan(&value.Total)
//     switch err {
//         case sql.ErrNoRows:
//             value.Total = 0
//             return value, nil
//         case nil:
//             return value, nil
//         default:
//             logger.Log.Println("GetPortgrouplevelCount Get Statement Prepare Error", err)
//             return value, err
//     }
// }
