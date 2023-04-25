package dao

import (
	"iFIX/ifix/logger"
	"log"
	"iFIX/ifix/entities"
)

var zonedetails ="SELECT zone_id as Id,zone_name as Zonename from zone where upper(zone_name) like ? LIMIT 10; "

func (mdao DbConn) Searchzone(tz *entities.ZoneEntity) ([]entities.ZoneEntity, error) {
	log.Println("In side dao Searchzone")
	values := []entities.ZoneEntity{}
	rows, err := mdao.DB.Query(zonedetails,"%"+tz.Zonename+"%")
	defer rows.Close()
	if err != nil {
		logger.Log.Println("Searchzone Get Statement Prepare Error", err)
		log.Print("Searchzone Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.ZoneEntity{}
		rows.Scan(&value.Id, &value.Zonename)
		values = append(values, value)
	}
	return values, nil
}
