package models

import "time"

type Invoice struct {
	ID           uint   `gorm:"primaryKey"`
	JobName      string 
	ProviderName string 
	ClientName   string 
	Country      string 
	Region       string 
	IssueDate    time.Time
	DueDate      time.Time
	TotalAmount  float64
	TaxAmount    float64
	TaxRate      float64
	Currency     string
	Status       string 
}
