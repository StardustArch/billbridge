package models

import "time"

type Invoice struct {
	ID           uint   `gorm:"primaryKey"`
	JobName      string // Nome do trabalho
	ProviderName string // Nome do provedor
	ClientName   string // Nome do cliente
	Country      string // País do invoice
	Region       string // Região do invoice
	IssueDate    time.Time
	DueDate      time.Time
	TotalAmount  float64
	TaxAmount    float64
	TaxRate      float64
	Currency     string
	Status       string // "Paid", "Pending", "Overdue"
}
