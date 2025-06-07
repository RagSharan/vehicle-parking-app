package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	VehicleID uint    `json:"vehicle_id"`
	Amount    float64 `json:"amount"`
	Paid      bool    `json:"paid"`
}
