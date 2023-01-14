package main

import (
	"config"
	"log"
	"net/http"
	"routes"
)

func main() {
	// Connect to database
	config.DBConfig{URL: "mongodb://localhost:27017", Database: "taskdb"}.Connect()
	defer config.DB.Close()

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))

}
