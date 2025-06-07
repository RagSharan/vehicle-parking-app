package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ragsharan/vehicle-parking-app/customer-service/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type CustomerHandler struct {
	DB *gorm.DB
}

func (cc *CustomerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := cc.DB.Create(&customer).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(customer)
}

func (cc *CustomerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var customer models.Customer
	if err := cc.DB.First(&customer, params["id"]).Error; err != nil {
		http.Error(w, "Customer not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(customer)
}

func (cc *CustomerHandler) ListCustomers(w http.ResponseWriter, r *http.Request) {
	var customers []models.Customer
	if err := cc.DB.Find(&customers).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(customers)
}

func (cc *CustomerHandler) AddVehicle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var vehicle models.Vehicle
	if err := json.NewDecoder(r.Body).Decode(&vehicle); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	vehicle.CustomerID = params["id"]
	if err := cc.DB.Create(&vehicle).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(vehicle)
}

func (cc *CustomerHandler) ListVehicles(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var vehicles []models.Vehicle
	if err := cc.DB.Where("customer_id = ?", params["id"]).Find(&vehicles).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(vehicles)
}
