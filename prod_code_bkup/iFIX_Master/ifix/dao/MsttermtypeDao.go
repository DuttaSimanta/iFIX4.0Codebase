package dao


import (
  "iFIX/ifix/entities"
  "iFIX/ifix/logger"
  "database/sql"
  )

var insertMsttermtype = "INSERT INTO msttermtype (termtypename) VALUES (?)"
var duplicateMsttermtype = "SELECT count(id) total FROM  msttermtype WHERE termtypename = ? AND deleteflg = 0 AND activeflg=1"
var getMsttermtype = "SELECT id as Id, termtypename as Termtypename FROM msttermtype ORDER BY id ASC LIMIT ?,?"
var getMsttermtypecount = "SELECT count(id) total FROM  msttermtype "
var updateMsttermtype = "UPDATE msttermtype SET termtypename = ? WHERE id = ? "
var deleteMsttermtype = "UPDATE msttermtype SET deleteflg = '1' WHERE id = ? "


func (dbc DbConn) CheckDuplicateMsttermtype(tz *entities.MsttermtypeEntity) (entities.MsttermtypeEntities, error) {
    logger.Log.Println("In side CheckDuplicateMsttermtype")
    value := entities.MsttermtypeEntities{}
    err := dbc.DB.QueryRow(duplicateMsttermtype, tz.Termtypename).Scan(&value.Total)
    switch err {
        case sql.ErrNoRows:
            value.Total = 0
            return value, nil
        case nil:
            return value, nil
        default:
            logger.Log.Println("CheckDuplicateMsttermtype Get Statement Prepare Error", err)
            return value, err
    }
}


func (dbc DbConn) InsertMsttermtype(tz *entities.MsttermtypeEntity) (int64, error) {
    logger.Log.Println("In side InsertMsttermtype")
    logger.Log.Println("Query -->",insertMsttermtype)
    stmt, err := dbc.DB.Prepare(insertMsttermtype)
    defer stmt.Close()
    if err != nil {
        logger.Log.Println("InsertMsttermtype Prepare Statement  Error", err)
        return 0, err
    }
    logger.Log.Println("Parameter -->",tz.Termtypename)
    res, err := stmt.Exec(tz.Termtypename)
    if err != nil {
        logger.Log.Println("InsertMsttermtype Execute Statement  Error", err)
        return 0, err
    }
    lastInsertedId, err := res.LastInsertId()
    return lastInsertedId, nil
}


func (dbc DbConn) GetAllMsttermtype(page *entities.MsttermtypeEntity) ([]entities.MsttermtypeEntity, error) {
    logger.Log.Println("In side GelAllMsttermtype")
    values := []entities.MsttermtypeEntity{}
    rows, err := dbc.DB.Query(getMsttermtype, page.Offset, page.Limit)
    defer rows.Close()
    if err != nil {
        logger.Log.Println("GetAllMsttermtype Get Statement Prepare Error", err)
        return values, err
    }
    for rows.Next() {
        value := entities.MsttermtypeEntity{}
        rows.Scan(&value.Id,&value.Termtypename)
        values = append(values, value)
    }
    return values, nil
}


func (dbc DbConn) UpdateMsttermtype(tz *entities.MsttermtypeEntity) error {
    logger.Log.Println("In side UpdateMsttermtype")
    stmt, err := dbc.DB.Prepare(updateMsttermtype)
    defer stmt.Close()
    if err != nil {
        logger.Log.Println("UpdateMsttermtype Prepare Statement  Error", err)
        return err
    }
    _, err = stmt.Exec(tz.Termtypename, tz.Id)
    if err != nil {
        logger.Log.Println("UpdateMsttermtype Execute Statement  Error", err)
        return err
    }
    return nil
}


func (dbc DbConn) DeleteMsttermtype(tz *entities.MsttermtypeEntity) error {
    logger.Log.Println("In side DeleteMsttermtype")
    stmt, err := dbc.DB.Prepare(deleteMsttermtype)
    defer stmt.Close()
    if err != nil {
        logger.Log.Println("DeleteMsttermtype Prepare Statement  Error", err)
        return err
    }
    _, err = stmt.Exec(tz.Id)
    if err != nil {
        logger.Log.Println("DeleteMsttermtype Execute Statement  Error", err)
        return err
    }
    return nil
}


func (dbc DbConn) GetMsttermtypeCount(tz *entities.MsttermtypeEntity) (entities.MsttermtypeEntities, error) {
    logger.Log.Println("In side GetMsttermtypeCount")
    value := entities.MsttermtypeEntities{}
    err := dbc.DB.QueryRow(getMsttermtypecount).Scan(&value.Total)
    switch err {
        case sql.ErrNoRows:
            value.Total = 0
            return value, nil
        case nil:
            return value, nil
        default:
            logger.Log.Println("GetMsttermtypeCount Get Statement Prepare Error", err)
            return value, err
    }
}


