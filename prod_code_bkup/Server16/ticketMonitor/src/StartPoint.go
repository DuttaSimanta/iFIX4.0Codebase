package main

import (
	Logger "src/logger"
	"src/models"

	"github.com/jasonlvhit/gocron"
)

func UserRecordCodeDelete() {
	Logger.Log.Println("<================AutoUserRecordCodeDelete is Started===============>")
	models.AutoUserRecordCodeDelete()
	Logger.Log.Println("<===================AutoUserRecordCodeDelete is finished====================>")
}
func main() {
	Logger.Log.Println("===========================Scheduler Started===============")

	s := gocron.NewScheduler()

	s.Every(5).Minutes().Do(UserRecordCodeDelete)
	<-s.Start()
}
