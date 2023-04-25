package models

import Logger "src/logger"

func SendNotifications(requestData map[string]interface{}) error {
	var channelType float64 = 1
	requestData["channeltype"] = channelType
	Logger.Log.Println("Channel TYpe For EMAIL===>", requestData["channeltype"])
	notificationEmailErr := MailBodyFormationFromTemplate(requestData)
	if notificationEmailErr != nil {
		Logger.Log.Println(notificationEmailErr)
	}
	channelType = 2
	requestData["channeltype"] = channelType
	Logger.Log.Println("Channel TYpe For SMS===>", requestData["channeltype"])
	notificationSMSErr := SMSFormationFromTemplate(requestData)

	if notificationSMSErr != nil {
		Logger.Log.Println(notificationSMSErr)
	}
	return nil
}
