package dao

import (
	"iFIX/ifix/entities"
	"database/sql"
	"log"
)

var insertmodule = "INSERT INTO mstmodule (modulename, moduledescription) VALUES (?,?)"
var duplicateModule = "SELECT count(id) total from mstmodule where modulename=? and deleteflg =0"
var getmodule = "SELECT id as Id,modulename as Modulename ,moduledescription as Moduledescription from mstmodule where deleteflg =0 ORDER BY id DESC LIMIT ?,? "
var getcount = "SELECT count(id) total from mstmodule where deleteflg =0"
var updatemodule = "UPDATE mstmodule set modulename=?,moduledescription=? where id=?"
var deletemodule = "UPDATE mstmodule set deleteflg=1 where id=?"
var modulename="SELECT modulename as Modulename from mstmodule where id=? and deleteflg =0"

func (mdao DbConn) CheckDuplicateModule(tz *entities.ModuleEntity) (entities.ModuleEntities, error) {
	log.Println("In side dao")
	value := entities.ModuleEntities{}
	err := mdao.DB.QueryRow(duplicateModule, tz.Modulename).Scan(&value.Total)
	switch err {
		case sql.ErrNoRows:
			value.Total = 0
			return value, nil
		case nil:
			return value, nil
		default:
			log.Print("checkDuplicateModule Get Statement Prepare Error", err)
			return value, err
	}
}
func (mdao DbConn) InsertModule(tz *entities.ModuleEntity) (int64, error) {
	log.Println("In side dao")
	stmt, err := mdao.DB.Prepare(insertmodule)
	defer stmt.Close()
	if err != nil {
		log.Print("InsertModule Prepare Statement  Error", err)
		return 0, err
	}
	res, err := stmt.Exec(tz.Modulename, tz.Moduledescription)
	if err != nil {
		log.Print("InsertModule Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}
func (mdao DbConn) GetAllModules(page *entities.PaginationEntity) ([]entities.ModuleEntity, error) {
	log.Println("In side dao")
	values := []entities.ModuleEntity{}
	rows, err := mdao.DB.Query(getmodule, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		log.Print("GetAllModules Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.ModuleEntity{}
		rows.Scan(&value.Id, &value.Modulename, &value.Moduledescription)
		values = append(values, value)
	}
	return values, nil
}
func (mdao DbConn) UpdateModule(tz *entities.ModuleEntity) error {
	log.Println("In side dao")
	stmt, err := mdao.DB.Prepare(updatemodule)
	defer stmt.Close()
	if err != nil {
		log.Print("UpdateModule Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Modulename, tz.Moduledescription, tz.Id)
	if err != nil {
		log.Print("UpdateModule Execute Statement  Error", err)
		return err
	}
	return nil
}
func (mdao DbConn) DeleteModule(tz *entities.ModuleEntity) error {
	log.Println("In side dao")
	stmt, err := mdao.DB.Prepare(deletemodule)
	defer stmt.Close()
	if err != nil {
		log.Print("DeleteModule Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		log.Print("DeleteModule Execute Statement  Error", err)
		return err
	}
	return nil
}
func (mdao DbConn) GetModuleCount() (entities.ModuleEntities, error) {
	log.Println("In side dao")
	value := entities.ModuleEntities{}
	err := mdao.DB.QueryRow(getcount).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		log.Print("GetModuleCount Get Statement Prepare Error", err)
		return value, err
	}
}
func (mdao DbConn) Getmodulename(tz *entities.ModuleEntity) ([]entities.ModuleEntity, error) {
	log.Println("In side dao")
	values := []entities.ModuleEntity{}
	rows, err := mdao.DB.Query(modulename, tz.Id)
	defer rows.Close()
	if err != nil {
		log.Print("Getmodulename Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.ModuleEntity{}
		rows.Scan(&value.Modulename)
		values = append(values, value)
	}
	return values, nil
}
