package fileutils

import (
	"errors"
	"os"
	"src/entities"
	Logger "src/logger"
	"strconv"
	"strings"

	"gopkg.in/gomail.v2"
	//"html"
	//ReadProperties "src/fileutils"
)

func SendMail(clientID int64, orgID int64, subject string, mailBody string, toEmailAddress string,
	ccEmailAddress string, attachments []entities.Attachment, smtpHostForMessaging string,
	smtpPort string, emailUserName string, emailPassword string) error {
	Logger.Log.Println("SendMail Starting...")
	wd, err := os.Getwd()
	if err != nil {
		Logger.Log.Println(err)
		return errors.New("ERROR: Unable to Open Directory")
	}
	//log.Println(wd)
	contextPath := strings.ReplaceAll(wd, "\\", "/")
	//log.Println(contextPath)
	// props, err := ReadPropertiesFile(contextPath + "/resource/application.properties")
	// if err != nil {
	// 	Logger.Log.Println(err)
	// 	return errors.New("ERROR: Unable to Read Properties File")
	// }
	Port, err := strconv.Atoi(smtpPort)
	if err != nil {
		Logger.Log.Println(err)
		return errors.New("ERROR: Unable to Read read smtpPort")
	}
	dialer := gomail.NewDialer(smtpHostForMessaging, Port, emailUserName, emailPassword)
	message := gomail.NewMessage()
	message.SetHeader("From", emailUserName)
	message.SetHeader("To", toEmailAddress)
	//var AttachedFiles []string
	if len(attachments) > 0 {
		for i := 0; i < len(attachments); i++ {
			filePath := contextPath + "/resource/downloads/" + attachments[i].OriginalFilenames

			fileDownloadErr := DownloadFileFromUrl(clientID, orgID, attachments[i].OriginalFilenames, attachments[i].UploadedFileName, filePath)
			if fileDownloadErr != nil {
				Logger.Log.Println(fileDownloadErr)
				return fileDownloadErr
			}
			Logger.Log.Println("File DownLoaded Successfull...Path==>", filePath)
			//AttachedFiles = append(AttachedFiles, filePath)
			message.Attach(filePath)

		}
		//message.Attach()
	}
	if !strings.EqualFold(ccEmailAddress, "") {
		var ccSlice []string = strings.Split(ccEmailAddress, ",")
		message.SetHeader("Cc", ccSlice...)
	}
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", mailBody)
	Logger.Log.Println(toEmailAddress)
	if err := dialer.DialAndSend(message); err != nil {

		Logger.Log.Println(err)
		return errors.New("ERROR: Unable to Send Mail")
	}
	Logger.Log.Println("Email Sent Successfully")
	return nil
}
