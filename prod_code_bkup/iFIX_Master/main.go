//***************************//
// Package Name: Main
// Date Of Creation: 15/12/2020
// Authour Name: Moloy Mondal
// History: N/A
// Synopsis: This is Entry Point of API.
// Functions: main
// Global Variable: N/A
// Version: 1.0.0
//***************************//
package main

import (
	"iFIX/ifix/logger"
	"iFIX/ifix/router"
	"net/http"
)

// main function is uded to call the API function and to start the http server for API
func main() {

	router.NewRouter()
	logger.Log.Println("Server started at 8082 port")
	logger.Log.Fatal(http.ListenAndServe(":8082", nil))
}
