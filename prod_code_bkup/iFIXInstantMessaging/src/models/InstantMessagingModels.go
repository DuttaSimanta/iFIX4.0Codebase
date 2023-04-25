package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"src/config"
	"src/dao"
	"src/entities"
	FileUtils "src/fileutils"
	Logger "src/logger"
	"strings"
	//ReadProperties "src"
	//"database/sql"
)

func RecordInstantMessaging(instantMessagingEntity *entities.InstantMessagingEntity) error {

	clientID := instantMessagingEntity.ClientID
	orgID := instantMessagingEntity.OrgID
	recordID := instantMessagingEntity.RecordID
	emailTo := instantMessagingEntity.EmailTo
	emailCc := instantMessagingEntity.EmailCc
	emailSub := instantMessagingEntity.EmailSub
	emailBody := instantMessagingEntity.EmailBody
	createdID := instantMessagingEntity.CreatedID
	createdGrpID := instantMessagingEntity.CreatedGrpID
	//attachmentsUrl := requestData["attachmentsurl"].([]string)
	wd, err := os.Getwd() // to get working directory
	if err != nil {
		Logger.Log.Println(err)
	}
	//log.Println(wd)
	contextPath := strings.ReplaceAll(wd, "\\", "/") // replacing backslash by  forwardslash
	//log.Println(contextPath)
	props, err := FileUtils.ReadPropertiesFile(contextPath + "/resource/application.properties")
	if err != nil {
		log.Println(err)
	}

	var activityName string = props["ActivityName"]
	//var activitySeqNo int64
	db, dBerr := config.GetDB()
	if dBerr != nil {
		Logger.Log.Println(dBerr)
		return errors.New("ERROR: Unable to connect DB")
	}
	ticketTypeSequenceNo, recordDiffTypeID, recordDiffID, recordeStagedID, getTicketTypeSequenceNoErr := dao.GetTicketTypeSequenceNo(db, clientID, orgID, recordID)
	if getTicketTypeSequenceNoErr != nil {
		Logger.Log.Println(getTicketTypeSequenceNoErr)
		return errors.New("ERROR: Unable to fetch ticketType Seq")
	}
	Logger.Log.Println("ticketTypeSequenceNo===>", ticketTypeSequenceNo)

	if ticketTypeSequenceNo == 3 {
		Logger.Log.Println("ticketTypeSequenceNo===>", ticketTypeSequenceNo)
		sRTicketID, sRTicketIDErr := dao.GetSRIDFromStask(db, clientID, orgID, recordID)
		if sRTicketIDErr != nil {
			Logger.Log.Println(getTicketTypeSequenceNoErr)
			return errors.New("ERROR: Unable to fetch SR Ticket ID from Stask")
		}

		Logger.Log.Println("SR ID===>", sRTicketID)
		ticketTypeSequenceNo, recordDiffTypeID, recordDiffID, recordeStagedID, getTicketTypeSequenceNoErr = dao.GetTicketTypeSequenceNo(db, clientID, orgID, sRTicketID)
		if getTicketTypeSequenceNoErr != nil {
			Logger.Log.Println(getTicketTypeSequenceNoErr)
			return errors.New("ERROR: Unable to fetch ticketType Seq")
		}
		recordID = sRTicketID
		//ticketTypeSequenceNo=2

	}
	if ticketTypeSequenceNo == 1 {
		var recordcommonEntity entities.RecordcommonEntity
		recordcommonEntity.ClientID = clientID
		recordcommonEntity.Mstorgnhirarchyid = orgID
		recordcommonEntity.ForuserID = createdID
		recordcommonEntity.Usergroupid = createdGrpID
		recordcommonEntity.UserID = createdID
		recordcommonEntity.Recorddifftypeid = recordDiffTypeID
		recordcommonEntity.Recorddiffid = recordDiffID
		recordcommonEntity.RecordstageID = recordeStagedID
		recordcommonEntity.RecordID = recordID
		recordcommonEntity.Termseq = 1
		var attachmentsName string

		for i := 0; i < len(instantMessagingEntity.Attachments); i++ {
			recordcommonEntity.Termvalue = instantMessagingEntity.Attachments[i].OriginalFilenames
			recordcommonEntity.Termdescription = instantMessagingEntity.Attachments[i].UploadedFileName
			attachmentsName = attachmentsName + "," + instantMessagingEntity.Attachments[i].OriginalFilenames

			sendData, err := json.Marshal(recordcommonEntity)
			if err != nil {
				Logger.Log.Println(err)
				return errors.New("Unable to marshal data")
			}
			Logger.Log.Println("Add Comment Json Json=======>   ", string(sendData))

			resp, err := http.Post(props["URLAddComment"], "application/json", bytes.NewBuffer(sendData))

			var result map[string]interface{}
			respBody, err1 := ioutil.ReadAll(resp.Body)
			if err1 != nil {
				Logger.Log.Println(err1)
				return errors.New("Unable to read response data")
			}
			err2 := json.Unmarshal(respBody, &result)
			if err2 != nil {
				Logger.Log.Println(err2)
				return errors.New("Unable to Unmarshal data")
			}

			//  json.NewDecoder(resp.Body).Decode(&res)
			if result["success"].(bool) == false {
				Logger.Log.Println("getting False for Attachment added")
				return errors.New("Ticket Comment Updation failed intermittently. Please try again later")
			} else {

				//Logger.Log.Println("Attachment Added Successfull")
				Logger.Log.Println("Attachment Added Successfull in TicketID======>", recordID)
			}
		}

		activitySeqNo, activitySeqNoResultSetErr := dao.GetActivitySeqNo(db, clientID, orgID, activityName)
		if activitySeqNoResultSetErr != nil {
			Logger.Log.Println(activitySeqNoResultSetErr)
			return errors.New("ERROR: Unable to Fetch Activity Seq No")
		}
		attachmentsName = strings.Trim(attachmentsName, ",")
		var logValue = "\nEmail To : " + emailTo + "\nCC	: " + emailCc + "\nEmail Subject : " + emailSub + "\nAttachments :" + attachmentsName + "\nEmail Body : " + emailBody + "\n"
		//var logValue = "<br/><b>Email To&emsp;:&emsp;</b>" + emailTo + "<br/><b>CC&emsp;:</b>&emsp;" + emailCc + "<br/><b>Email Subject&emsp;:&emsp;</b>" + emailSub + "<br/><b>Email Body&emsp;:&emsp;</b>" + emailBody + "<br/><br/>"
		tx, txErr := db.Begin()
		if txErr != nil {
			Logger.Log.Println(txErr)
			return errors.New("ERROR: Databse Error")

		}
		insertMstInstantMessagingErr := dao.InsertMstInstantMessagingRecord(db, tx, clientID, orgID, recordID, emailTo, emailCc, emailSub, emailBody)
		//insert record to mstrecordinstantmesseging
		if insertMstInstantMessagingErr != nil {
			Logger.Log.Println(insertMstInstantMessagingErr)
			return errors.New("ERROR: Error in Inserting Record for mstrecordinstantmesseging")
		}
		insertActivityLogError := dao.InsertMstRecordActivityLogs(db, tx, clientID, orgID, recordID, activitySeqNo, logValue, createdID, createdGrpID)
		if insertActivityLogError != nil {
			Logger.Log.Println(insertActivityLogError)
			return errors.New("ERROR: Error in Inserting Record for mstrecordactivitylogs")
		}

		if strings.Contains(emailBody, "\n") {
			emailBody = strings.ReplaceAll(emailBody, "\n", "<br/>")
		}
		smtpHostForNotification, smtpPort, emailUserName, emailPassword, smtpError := dao.GetInstantMessagingSMTPDetails(db, clientID, orgID)
		if smtpError != nil {
			Logger.Log.Println(smtpError)
			return errors.New("ERROR: SMTP Configuration Not Found!!!")
		}
		sendMailErr := FileUtils.SendMail(clientID, orgID, emailSub, emailBody, emailTo, emailCc, instantMessagingEntity.Attachments, smtpHostForNotification, smtpPort, emailUserName, emailPassword)
		if sendMailErr != nil {
			Logger.Log.Println(sendMailErr)
			tx.Rollback()
			return errors.New("ERROR: Unable to Send Mail")
		}
		commitErr := tx.Commit()
		if commitErr != nil {
			Logger.Log.Println(err)
			tx.Rollback()
			return errors.New("ERROR: Commit Error")
		}
	}
	//defer db.Close()
	if ticketTypeSequenceNo == 2 {
		sTaskIDList, sTaskIDListErr := dao.GetSTaskIDListOfSR(db, clientID, orgID, recordID)
		if sTaskIDListErr != nil {
			Logger.Log.Println(sTaskIDListErr)
			return errors.New("ERROR: Unable to Get STask")
		}
		var recordcommonEntity entities.RecordcommonEntity
		recordcommonEntity.ClientID = clientID
		recordcommonEntity.Mstorgnhirarchyid = orgID
		recordcommonEntity.ForuserID = createdID
		recordcommonEntity.Usergroupid = createdGrpID
		recordcommonEntity.UserID = createdID
		recordcommonEntity.Recorddifftypeid = recordDiffTypeID
		recordcommonEntity.Recorddiffid = recordDiffID
		recordcommonEntity.RecordstageID = recordeStagedID
		recordcommonEntity.RecordID = recordID
		recordcommonEntity.Termseq = 1
		var attachmentsName string

		for i := 0; i < len(instantMessagingEntity.Attachments); i++ {
			recordcommonEntity.Termvalue = instantMessagingEntity.Attachments[i].OriginalFilenames
			recordcommonEntity.Termdescription = instantMessagingEntity.Attachments[i].UploadedFileName
			attachmentsName = attachmentsName + "," + instantMessagingEntity.Attachments[i].OriginalFilenames

			sendData, err := json.Marshal(recordcommonEntity)
			if err != nil {
				Logger.Log.Println(err)
				return errors.New("Unable to marshal data")
			}
			Logger.Log.Println("Add Comment Json Json=======>   ", string(sendData))

			resp, err := http.Post(props["URLAddComment"], "application/json", bytes.NewBuffer(sendData))

			var result map[string]interface{}
			respBody, err1 := ioutil.ReadAll(resp.Body)
			if err1 != nil {
				Logger.Log.Println(err1)
				return errors.New("Unable to read response data")
			}
			err2 := json.Unmarshal(respBody, &result)
			if err2 != nil {
				Logger.Log.Println(err2)
				return errors.New("Unable to Unmarshal data")
			}

			//  json.NewDecoder(resp.Body).Decode(&res)
			if result["success"].(bool) == false {
				Logger.Log.Println("getting False for Attachment added")
				return errors.New("Ticket Comment Updation failed intermittently. Please try again later")
			} else {

				//Logger.Log.Println("Attachment Added Successfull")
				Logger.Log.Println("Attachment Added Successfull in TicketID======>", recordID)
			}
		}
		activitySeqNo, activitySeqNoResultSetErr := dao.GetActivitySeqNo(db, clientID, orgID, activityName)
		if activitySeqNoResultSetErr != nil {
			Logger.Log.Println(activitySeqNoResultSetErr)
			return errors.New("ERROR: Unable to Fetch Activity Seq No")
		}
		attachmentsName = strings.Trim(attachmentsName, ",")
		var logValue = "\nEmail To : " + emailTo + "\nCC	: " + emailCc + "\nEmail Subject : " + emailSub + "\nAttachments :" + attachmentsName + "\nEmail Body : " + emailBody + "\n"
		sTaskIDList = append(sTaskIDList, recordID)
		tx, txErr := db.Begin()
		if txErr != nil {
			Logger.Log.Println(txErr)
			return errors.New("ERROR: Databse Error")

		}

		for i := 0; i < len(sTaskIDList); i++ {
			insertMstInstantMessagingErr := dao.InsertMstInstantMessagingRecord(db, tx, clientID, orgID, sTaskIDList[i], emailTo, emailCc, emailSub, emailBody)
			//insert record to mstrecordinstantmesseging
			if insertMstInstantMessagingErr != nil {
				Logger.Log.Println(insertMstInstantMessagingErr)
				return errors.New("ERROR: Error in Inserting Record for mstrecordinstantmesseging")
			}
			insertActivityLogError := dao.InsertMstRecordActivityLogs(db, tx, clientID, orgID, sTaskIDList[i], activitySeqNo, logValue, createdID, createdGrpID)
			if insertActivityLogError != nil {
				Logger.Log.Println(insertActivityLogError)
				return errors.New("ERROR: Error in Inserting Record for mstrecordactivitylogs")
			}
		}
		if strings.Contains(emailBody, "\n") {
			emailBody = strings.ReplaceAll(emailBody, "\n", "<br/>")
		}

		smtpHostForNotification, smtpPort, emailUserName, emailPassword, smtpError := dao.GetInstantMessagingSMTPDetails(db, clientID, orgID)
		if smtpError != nil {
			Logger.Log.Println(smtpError)
			return errors.New("ERROR: SMTP Configuration Not Found!!!")
		}
		sendMailErr := FileUtils.SendMail(clientID, orgID, emailSub, emailBody, emailTo, emailCc, instantMessagingEntity.Attachments, smtpHostForNotification, smtpPort, emailUserName, emailPassword)
		if sendMailErr != nil {
			Logger.Log.Println(sendMailErr)
			tx.Rollback()
			return errors.New("ERROR: Unable to Send Mail")
		}
		commitErr := tx.Commit()
		if commitErr != nil {
			Logger.Log.Println(err)
			tx.Rollback()
			return errors.New("ERROR: Commit Error")
		}

	}
	return nil
}
