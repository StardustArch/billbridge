package models

import "time"

type Invoice struct {
	ID          uint      `gorm:"primaryKey"`
	JobID       uint
	ProviderID  uint
	ClientID    uint
	IssueDate   time.Time
	DueDate     time.Time
	TotalAmount float64
	TaxAmount   float64
	TaxRate     float64
	Currency    string
	Status      string // "Paid", "Pending", "Overdue"
	PDFPath     string
}

