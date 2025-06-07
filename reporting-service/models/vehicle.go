package models

import "gorm.io/gorm"

type Parking struct {
	gorm.Model
	Name    string
	City    string
	Address string
}

type Vehicle struct {
	gorm.Model
	LicensePlate string
	ParkingID    uint
	Exited       bool
	Subscription bool
}

type Payment struct {
	gorm.Model
	VehicleID uint
	Amount    float64
}
