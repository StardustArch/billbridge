package main

import (
	"log"
	"net/http"
	"billbridge/database"
	"billbridge/router"
)

func main() {
	// Initialize database connection
	database.Connect()

	// Initialize router
	r := router.InitRouter()

	// Start the server
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}