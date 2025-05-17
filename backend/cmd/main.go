package main

import (
	"billbridge/database"
	"billbridge/router"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	// Initialize database connection
	database.Connect()

	// Initialize router
	r := router.InitRouter()

		// Configuração do CORS
		corsMiddleware := cors.New(cors.Options{
			AllowedOrigins: []string{"http://localhost:8081"}, // Permitir requisições do frontend
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"}, // Métodos permitidos
			AllowedHeaders: []string{"Content-Type", "Authorization"}, // Cabeçalhos permitidos
			AllowCredentials: true, // Permitir cookies e credenciais
		})
	
		// Aplica o middleware CORS
		handler := corsMiddleware.Handler(r)

	// Start the server
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}