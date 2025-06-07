package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ragsharan/vehicle-parking-app/parking-service/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type VehicleHandler struct {
	DB *gorm.DB
}

func (vc *VehicleHandler) RegisterVehicle(w http.ResponseWriter, r *http.Request) {
	var vehicle models.Vehicle
	if err := json.NewDecoder(r.Body).Decode(&vehicle); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := vc.DB.Create(&vehicle).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(vehicle)
}

func (vc *VehicleHandler) MarkVehicleExit(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var vehicle models.Vehicle
	if err := vc.DB.First(&vehicle, params["id"]).Error; err != nil {
		http.Error(w, "Vehicle not found", http.StatusNotFound)
		return
	}
	vehicle.Exited = true
	if err := vc.DB.Save(&vehicle).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(vehicle)
}

func (vc *VehicleHandler) ListVehicles(w http.ResponseWriter, r *http.Request) {
	var vehicles []models.Vehicle
	if err := vc.DB.Find(&vehicles).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(vehicles)
}

func (vc *VehicleHandler) ListParkings(w http.ResponseWriter, r *http.Request) {
	var parkings []models.Parking
	if err := vc.DB.Find(&parkings).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(parkings)
}
