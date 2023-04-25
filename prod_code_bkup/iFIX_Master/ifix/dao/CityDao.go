package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var getCity = "SELECT id as Id, cityname as Cityname FROM mstcity"
var getCitycount = "SELECT count(a.id) as total FROM mstcity"

// func (dbc DbConn) CheckDuplicateCity(tz *entities.CityEntity) (entities.CityEntities, error) {
//     logger.Log.Println("In side CheckDuplicateCity")
//     value := entities.CityEntities{}
//     err := dbc.DB.QueryRow(duplicateCity, ).Scan(&value.Total)
//     switch err {
//         case sql.ErrNoRows:
//             value.Total = 0
//             return value, nil
//         case nil:
//             return value, nil
//         default:
//             logger.Log.Println("CheckDuplicateCity Get Statement Prepare Error", err)
//             return value, err
//     }
// }

// func (dbc DbConn) InsertCity(tz *entities.CityEntity) (int64, error) {
//     logger.Log.Println("In side InsertCity")
//     logger.Log.Println("Query -->",insertCity)
//     stmt, err := dbc.DB.Prepare(insertCity)
//     defer stmt.Close()
//     if err != nil {
//         logger.Log.Println("InsertCity Prepare Statement  Error", err)
//         return 0, err
//     }
//     logger.Log.Println("Parameter -->",)
//     res, err := stmt.Exec()
//     if err != nil {
//         logger.Log.Println("InsertCity Execute Statement  Error", err)
//         return 0, err
//     }
//     lastInsertedId, err := res.LastInsertId()
//     return lastInsertedId, nil
// }

func (dbc DbConn) GetAllCity() ([]entities.CityEntity, error) {
	logger.Log.Println("In side GelAllCity")
	values := []entities.CityEntity{}
	rows, err := dbc.DB.Query(getCity)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllCity Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.CityEntity{}
		rows.Scan(&value.Id, &value.Cityname)
		values = append(values, value)
	}
	return values, nil
}

// func (dbc DbConn) UpdateCity(tz *entities.CityEntity) error {
//     logger.Log.Println("In side UpdateCity")
//     stmt, err := dbc.DB.Prepare(updateCity)
//     defer stmt.Close()
//     if err != nil {
//         logger.Log.Println("UpdateCity Prepare Statement  Error", err)
//         return err
//     }
//     _, err = stmt.Exec( tz.Id)
//     if err != nil {
//         logger.Log.Println("UpdateCity Execute Statement  Error", err)
//         return err
//     }
//     return nil
// }

// func (dbc DbConn) DeleteCity(tz *entities.CityEntity) error {
//     logger.Log.Println("In side DeleteCity")
//     stmt, err := dbc.DB.Prepare(deleteCity)
//     defer stmt.Close()
//     if err != nil {
//         logger.Log.Println("DeleteCity Prepare Statement  Error", err)
//         return err
//     }
//     _, err = stmt.Exec(tz.Id)
//     if err != nil {
//         logger.Log.Println("DeleteCity Execute Statement  Error", err)
//         return err
//     }
//     return nil
// }

func (dbc DbConn) GetCityCount(tz *entities.CityEntity) (entities.CityEntities, error) {
	logger.Log.Println("In side GetCityCount")
	value := entities.CityEntities{}
	err := dbc.DB.QueryRow(getCitycount).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetCityCount Get Statement Prepare Error", err)
		return value, err
	}
}
