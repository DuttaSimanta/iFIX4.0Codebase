package main

import (
	Logger "src/logger"
	"src/models"

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
func Closuretask() {
	Logger.Log.Println("<================Closuretask is Started===============>")
	models.Autoclosure()
	Logger.Log.Println("<===================Closuretask is finished====================>")
}
func main() {
	Logger.Log.Println("===========================Scheduler Started===============")

	s := gocron.NewScheduler()
	//s.Every(1).Minutes().Do(task)
	//s.Every(2).Minutes().Do(Closuretask)
	//s.Every(1).Day().At("10:30").Do(EmailTask)
	//s.Every(2).Minutes().Do(EmailTaskForAging)
	//s.Every(1).Day().At("05:35").Do(TicketDispatcherBot)
	//	s.Every(1).Hour().Do(Closuretask)

	s.Every(2).Minutes().Do(Closuretask)
	//s.Every(1).Minutes().Do(EmailTaskForCustomerWorkNote)
	<-s.Start()
}
