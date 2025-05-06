package models


type TaxRule struct {
	ID      uint    `gorm:"primaryKey"`
	Country string
	Region  string
	Rate    float64
	Active  bool
}