package main

import (
	Logger "src/logger"
	"src/model"

	"github.com/jasonlvhit/gocron"
)

// func task() {
// 	logger.Log.Println("Task is being performed.")
// 	model.Lucenereindex()
// }

// func Closuretask() {
// 	logger.Log.Println("Closuretask is being performed.")
// 	model.Autoclosure()
// }
// func EmailTaskForAging() {
// 	Logger.Log.Println("EmailTaskForAging is being performed.")
// 	model.EmailNotificationForAging()
// 	model.EmailNotificationForCustomerWorkNote()
// }
// func EmailTaskForCustomerWorkNote() {
// 	Logger.Log.Println("EmailNotificationForCustomerWorkNote is being performed.")
// 	model.EmailNotificationForCustomerWorkNote()
// }
func AutoResponse() {
	Logger.Log.Println("===========================Auto Response Started===============")
	model.AutoResponse()
}

func main() {
	Logger.Log.Println("===========================Scheduler Started===============")

	s := gocron.NewScheduler()
	//s.Every(1).Minutes().Do(task)
	//s.Every(2).Minutes().Do(Closuretask)
	//s.Every(1).Day().At("10:30").Do(EmailTask)
	//s.Every(2).Minutes().Do(EmailTaskForAging)
	//s.Every(1).Day().At("05:35").Do(TicketDispatcherBot)

	s.Every(120).Seconds().Do(AutoResponse)
	//s.Every(1).Minutes().Do(EmailTaskForCustomerWorkNote)
	<-s.Start()
}
