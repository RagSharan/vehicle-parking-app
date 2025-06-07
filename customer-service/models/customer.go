package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	CustomerID uint      `json:"customer_id"`
	Name       string    `json:"name"`
	Phone      string    `json:"phone"`
	Email      string    `json:"email"`
	Vehicles   []Vehicle `json:"vehicles" gorm:"foreignKey:CustomerID"`
}

type Vehicle struct {
	gorm.Model
	LicensePlate string `json:"license_plate"`
	CustomerID   uint   `json:"customer_id"`
}
