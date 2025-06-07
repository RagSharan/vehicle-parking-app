package models

import "gorm.io/gorm"

type Vehicle struct {
    gorm.Model
    LicensePlate string `json:"license_plate"`
    Exited       bool   `json:"exited"`
}
