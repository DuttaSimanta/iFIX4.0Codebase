package dao

import (
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var getCountry = "SELECT id as Id, countryname as Countryname FROM mstcountry "

// func (dbc DbConn) CheckDuplicateCountry(tz *entities.CountryEntity) (entities.CountryEntities, error) {
//     logger.Log.Println("In side CheckDuplicateCountry")
//     value := entities.CountryEntities{}
//     err := dbc.DB.QueryRow(duplicateCountry, ).Scan(&value.Total)
//     switch err {
//         case sql.ErrNoRows:
//             value.Total = 0
//             return value, nil
//         case nil:
//             return value, nil
//         default:
//             logger.Log.Println("CheckDuplicateCountry Get Statement Prepare Error", err)
//             return value, err
//     }
// }

// func (dbc DbConn) InsertCountry(tz *entities.CountryEntity) (int64, error) {
//     logger.Log.Println("In side InsertCountry")
//     logger.Log.Println("Query -->",insertCountry)
//     stmt, err := dbc.DB.Prepare(insertCountry)
//     defer stmt.Close()
//     if err != nil {
//         logger.Log.Println("InsertCountry Prepare Statement  Error", err)
//         return 0, err
//     }
//     logger.Log.Println("Parameter -->",)
//     res, err := stmt.Exec()
//     if err != nil {
//         logger.Log.Println("InsertCountry Execute Statement  Error", err)
//         return 0, err
//     }
//     lastInsertedId, err := res.LastInsertId()
//     return lastInsertedId, nil
// }

func (dbc DbConn) GetAllCountry() ([]entities.CountryEntity, error) {
	logger.Log.Println("In side GelAllCountry")
	values := []entities.CountryEntity{}
	rows, err := dbc.DB.Query(getCountry)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllCountry Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.CountryEntity{}
		rows.Scan(&value.Id, &value.Countryname)
		values = append(values, value)
	}
	return values, nil
}

// func (dbc DbConn) UpdateCountry(tz *entities.CountryEntity) error {
//     logger.Log.Println("In side UpdateCountry")
//     stmt, err := dbc.DB.Prepare(updateCountry)
//     defer stmt.Close()
//     if err != nil {
//         logger.Log.Println("UpdateCountry Prepare Statement  Error", err)
//         return err
//     }
//     _, err = stmt.Exec( tz.Id)
//     if err != nil {
//         logger.Log.Println("UpdateCountry Execute Statement  Error", err)
//         return err
//     }
//     return nil
// }

// func (dbc DbConn) DeleteCountry(tz *entities.CountryEntity) error {
//     logger.Log.Println("In side DeleteCountry")
//     stmt, err := dbc.DB.Prepare(deleteCountry)
//     defer stmt.Close()
//     if err != nil {
//         logger.Log.Println("DeleteCountry Prepare Statement  Error", err)
//         return err
//     }
//     _, err = stmt.Exec(tz.Id)
//     if err != nil {
//         logger.Log.Println("DeleteCountry Execute Statement  Error", err)
//         return err
//     }
//     return nil
// }

// func (dbc DbConn) GetCountryCount(tz *entities.CountryEntity) (entities.CountryEntities, error) {
//     logger.Log.Println("In side GetCountryCount")
//     value := entities.CountryEntities{}
//     err := dbc.DB.QueryRow(getCountrycount, ).Scan(&value.Total)
//     switch err {
//         case sql.ErrNoRows:
//             value.Total = 0
//             return value, nil
//         case nil:
//             return value, nil
//         default:
//             logger.Log.Println("GetCountryCount Get Statement Prepare Error", err)
//             return value, err
//     }
// }
