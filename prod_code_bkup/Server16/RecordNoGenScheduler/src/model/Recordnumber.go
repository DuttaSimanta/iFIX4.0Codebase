package model

import (
	"src/dao"
	"src/logger"
	"strconv"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var lockk = &sync.Mutex{}
var TicketNO = 25000

// var lockk2 = &sync.Mutex{}

func Recordnumbergeneration() {
	lockk.Lock()
	defer lockk.Unlock()
	if db == nil {
		dbcon, err := GetDB()
		if err != nil {
			logger.Log.Println("Error in DBConnection")
			return
		}
		db = dbcon
	}
	logger.Log.Println("*********************RecordNoGen STarted ************************")
	recorddiffID := dao.GetTickettype(db)
	logger.Log.Println("ICCM Ticket Type ID Array Values ::::::", recorddiffID)

	for i, _ := range recorddiffID {
		count := dao.GetTicketCount(db, recorddiffID[i])
		logger.Log.Println("count is", count)
		logger.Log.Println("ICCM Ticket Type ID :::::", recorddiffID[i])
		var prefixarr []string
		prefixarr = dao.Getmstrecordconfig(db, 2, 2, 2, recorddiffID[i])
		var k int = 0
		for k = 0; k < TicketNO-int(count); k++ {
			var ticketID string
			var number int
			number = dao.Getnumber(db, 2, 2, 2, recorddiffID[i])
			currentTime := time.Now()
			var aa = currentTime.Format("20060102")
			if len(prefixarr[0]) > 0 {
				ticketID = ticketID + prefixarr[0]
			}
			if len(prefixarr[2]) > 0 {
				if prefixarr[2] != "NA" {
					var mm = string(aa[4:6])
					ticketID = ticketID + mm
				}
			}
			if len(prefixarr[3]) > 0 {
				if prefixarr[3] != "NA" {
					var day = string(aa[6:8])
					ticketID = ticketID + day
				}
			}
			if len(prefixarr[1]) > 0 {
				if prefixarr[1] != "NA" {
					if len(prefixarr[1]) == 2 {
						var yy = string(aa[2:4])
						ticketID = ticketID + yy
					} else {
						var yy = string(aa[0:4])
						ticketID = ticketID + yy
					}

				}
			}
			if len(prefixarr[4]) > 0 {
				var s = prefixarr[4]
				var numberlength = len(strconv.Itoa(number))
				var safeSubstring = string(s[0 : len(s)-numberlength])
				ticketID = ticketID + safeSubstring + strconv.Itoa(number)
			}

			err := dao.Updatemstrecordautoincreament(db, recorddiffID[i])
			if err != nil {
				logger.Log.Println("Updatemstrecordautoincreament ERROR:", err)
			}
			_, err1 := dao.InsertRecordcode(db, recorddiffID[i], ticketID)
			if err1 != nil {
				logger.Log.Println("InsertRecordcode ERROR:", err1)
			}
			// Insert Ticket code in new table end here ......
		} // End 300000 for loop here for each ticket type...
		logger.Log.Println("Record inserted:", k, recorddiffID[i])
	} // End recorddiffID for loop here.....
	logger.Log.Println("*********************RecordNoGen Ended************************")

}

// func main() {
// 	Recordnumbergeneration()
// }
func Deleterecordcode() {
	lockk.Lock()
	defer lockk.Unlock()
	if db == nil {
		dbcon, err := GetDB()
		if err != nil {
			logger.Log.Println("Error in DBConnection")
			return
		}
		db = dbcon
	}
	logger.Log.Println("*********************STARTED DELETING************************")

	dao.Deleterecordcode(db)
	logger.Log.Println("*********************ENDED DELETING************************")

}
