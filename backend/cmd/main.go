package main

import (
	"billbridge/database"
	"billbridge/router"
	"log"
	"net/http"
	"billbridge/database"
	"billbridge/router"

	"github.com/rs/cors"
)

func main() {
	// Initialize database connection
	database.Connect()

	// Initialize router
	r := router.InitRouter()

	// Configurar CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"}, // ou "*" para todos, se estiveres em dev
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	// Start the server
	handler := c.Handler(r)
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}