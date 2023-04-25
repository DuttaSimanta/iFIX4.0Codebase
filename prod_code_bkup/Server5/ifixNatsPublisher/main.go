package main

import (
	"encoding/json"
	"ifixNatsPublisher/ifix/logger"
	"ifixNatsPublisher/ifix/models"

	"math/rand"
	"time"

	"github.com/jasonlvhit/gocron"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
	// "github.com/jasonlvhit/gocron"
	// "github.com/jasonlvhit/gocron"
)

var (
	rg = rand.New(rand.NewSource(time.Now().Unix()))
)

// func main() {
// 	url := "nats://localhost:4222"
// 	nc, err := nats.Connect(url)
// 	if err != nil {
// 		logrus.Fatal(err)
// 	}

// 	defer nc.Close()

// 	for i := 0; i < 1e5; i++ {
// 		s := fmt.Sprintf("Message %v: data: %v", i, rg.Intn(10000))

// 		nc.Publish("events.old", []byte(s))
// 	}

// }
func SchedulerTask() {
	// logger.Log.Println("SLASchedulerTask is being performed.")
	// models.ExecuteRemainingUpdate()
	//models.ExecuteOldRecords()
	bugtickets := models.GetBugTickets()
	url := "nats://localhost:4222"
	nc, err := nats.Connect(url)
	if err != nil {
		logrus.Fatal(err)
	}

        logger.Log.Println("Connected to " + url)

//	defer nc.Close()

	for i := 0; i < len(bugtickets); i++ {
		s, _ := json.Marshal(bugtickets[i])
		// s := fmt.Sprintf("Message %v: data: %v", i, rg.Intn(10000))

		err:=nc.Publish("events.incident", s)
                if err == nil {
                      logger.Log.Println("Incident:Message published",string(s))
                }

	}
defer nc.Close()
}
func SRSchedulerTask() {
	// logger.Log.Println("SLASchedulerTask is being performed.")
	// models.ExecuteRemainingUpdate()
	//models.ExecuteOldRecords()
	bugtickets := models.GetSRBugTickets()
	url := "nats://localhost:4222"
	nc, err := nats.Connect(url)
	if err != nil {
		logrus.Fatal(err)
	}

	logger.Log.Println("Connected to " + url)

	//	defer nc.Close()

	for i := 0; i < len(bugtickets); i++ {
		s, _ := json.Marshal(bugtickets[i])
		// s := fmt.Sprintf("Message %v: data: %v", i, rg.Intn(10000))

		err := nc.Publish("events.sr", s)
		if err == nil {
			logger.Log.Println("SR:Message published", string(s))
		}

	}
	defer nc.Close()
}
func DeleteDefectTkt(){
models.DeleteDefectTkt()
}
func main() {
	logger.Log.Println("===========================Scheduler Started===============")

	s := gocron.NewScheduler()

	s.Every(2).Minutes().Do(SchedulerTask)
	s.Every(5).Minutes().Do(SRSchedulerTask)
	s.Every(33).Minutes().Do(DeleteDefectTkt)
	<-s.Start()
	// models.ExecuteRemainingUpdate()
	// SLASchedulerTask()

}
