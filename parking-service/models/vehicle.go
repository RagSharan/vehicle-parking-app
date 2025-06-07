package models

import "gorm.io/gorm"

type Vehicle struct {
	gorm.Model
	LicensePlate string `json:"license_plate"`
	ParkingID    uint   `json:"parking_id"`
	Exited       bool   `json:"exited"`
}
