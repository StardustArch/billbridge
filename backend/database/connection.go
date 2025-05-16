package database

import (
	"billbridge/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Se necessário, você poderia tentar carregar o .env somente para ambientes locais.
	// Mas em produção (via docker-compose), as variáveis já estarão disponíveis.
	// err := godotenv.Load("../.env")
	// if err != nil {
	//      log.Printf("Warning: .env file not found: %v", err)
	// }

	// Configuração do banco de dados
	config := struct {
		host     string
		user     string
		password string
		dbname   string
		port     string
	}{
		host:     getEnvWithDefault("DB_HOST", "db"),
		user:     getEnvWithDefault("DB_USER", "postgres"),
		password: getEnvWithDefault("DB_PASSWORD", "postgres"),
		dbname:   getEnvWithDefault("DB_NAME", "invoicedb"),
		port:     getEnvWithDefault("DB_PORT", "5432"),
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.host, config.user, config.password, config.dbname, config.port)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Auto-migrate dos modelos
	if err := DB.AutoMigrate(&models.Invoice{}, &models.TaxRule{}); err != nil {
		log.Fatalf("Error during database migration: %v", err)
	}

	log.Println("Successfully connected to database")
}

// getEnvWithDefault retorna o valor da variável de ambiente ou um valor padrão
func getEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
