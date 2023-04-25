package main

// import (
// 	"StateMismatchUpdateSchedular/ifix/logger"
// 	"StateMismatchUpdateSchedular/ifix/models"

// 	"github.com/jasonlvhit/gocron"
// 	// "github.com/jasonlvhit/gocron"
// 	// "github.com/jasonlvhit/gocron"
// )

// func SchedulerTask() {
// 	logger.Log.Println("SLASchedulerTask is being performed.")
// 	models.ExecuteRemainingUpdate()
// 	//models.ExecuteOldRecords()
// }

// func main() {
// 	logger.Log.Println("===========================Scheduler Started===============")

// 	s := gocron.NewScheduler()

// 	s.Every(5).Minutes().Do(SchedulerTask)

// 	<-s.Start()
// 	// models.ExecuteRemainingUpdate()
// 	// SLASchedulerTask()

// }

import (
	"ifixNatsSubscriber/ifix/logger"

	"encoding/json"
	"ifixNatsSubscriber/ifix/entities"
	"ifixNatsSubscriber/ifix/models"

	//	"time"
	"runtime"

	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

func main() {
	url := "nats://localhost:4222"
	nc, err := nats.Connect(url)
	if err != nil {
		logrus.Fatal(err)
	}
	logger.Log.Println("Connected to " + url)

	//	defer nc.Close()

	nc.Subscribe("events.incident", func(msg *nats.Msg) {
		logger.Log.Println("Incident:message recieved on subject:, data:", msg.Subject, string(msg.Data))

		tz := entities.RecordDetailsEntity{}

		err := json.Unmarshal(msg.Data, &tz)
		if err != nil {
			return
			// logger.Log.Println("ERRROR TO UNMARSHAL")
		}
		models.ExecuteRemainingUpdate(tz)
	})

	nc.Subscribe("events.sr", func(msg *nats.Msg) {
		logger.Log.Println("SR:message recieved on subject:, data:", msg.Subject, string(msg.Data))

		tz := entities.RecordDetailsEntity{}

		err := json.Unmarshal(msg.Data, &tz)
		if err != nil {
			return
			// logger.Log.Println("ERRROR TO UNMARSHAL")
		}
		models.ExecuteSRRemainingUpdate(tz)
	})

	runtime.Goexit()

}
