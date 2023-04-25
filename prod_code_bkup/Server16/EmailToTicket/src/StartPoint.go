package main

import (
	"src/emailticket"
	Logger "src/logger"
)

func main() {
	Logger.Log.Println("Main Started")

	emailticket.CreateTicketByEmail()
}
