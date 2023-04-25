package dao

import (
	"ifixRecord/ifix/entities"
	"ifixRecord/ifix/logger"
)

var updaterecordfulticketiddeleteflag = "UPDATE recordfulldetails SET ticketid = ?, deleteflg = ? WHERE recordid =? "

var updatetrnrecordcodedeleteflgbyid = "UPDATE trnrecord SET code = ?, deleteflg =? WHERE id = ?"

func (mdao DbConn) UpdateRecordfulTicketIdDeleteFlag(tz *entities.RecordfulTicketIdDeleteFlagUpdateEntity) error {
	logger.Log.Println("Inside RecordfulTicketIdDeleteFlagUpdateEntity Doa")
	stmt, err := mdao.DB.Prepare(updaterecordfulticketiddeleteflag)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateRecordfulTicketIdDeleteFlag Prepare Statement Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Ticketid, tz.Deleteflg, tz.Recordid)
	if err != nil {
		logger.Log.Println("UpdateRecordfulTicketIdDeleteFlag Execute Statement Error", err)
		return err
	}
	return nil
}

func (mdao DbConn) UpdateTrnRecordCodeDeleteFlgById(tz *entities.TrnRecordCodeDeleteFlgUpdateByIdEntity) error {
	logger.Log.Println("Inside TrnRecordCodeDeleteFlgByIdUpdateEntity Dao")
	stmt, err := mdao.DB.Prepare(updatetrnrecordcodedeleteflgbyid)

	defer stmt.Close()
	if err != nil {
		logger.Log.Println(" UpdateTrnRecordCodeDeleteFlgById Prepare Statement Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Code, tz.Deleteflg, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateTrnRecordCodeDeleteFlgById Execute Statement Error", err)
		return err
	}
	return nil
}
