package models

import "gorm.io/gorm"

type Parking struct {
	gorm.Model
	Name    string `json:"name"`
	City    string `json:"city"`
	Address string `json:"address"`
}
