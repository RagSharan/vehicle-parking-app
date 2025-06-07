package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ragsharan/vehicle-parking-app/payment-service/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type PaymentHandler struct {
	DB *gorm.DB
}

func (pc *PaymentHandler) CreatePayment(w http.ResponseWriter, r *http.Request) {
	var payment models.Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := pc.DB.Create(&payment).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(payment)
}

func (pc *PaymentHandler) GetPayment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var payment models.Payment
	if err := pc.DB.First(&payment, params["id"]).Error; err != nil {
		http.Error(w, "Payment not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(payment)
}

func (pc *PaymentHandler) ListPayments(w http.ResponseWriter, r *http.Request) {
	var payments []models.Payment
	if err := pc.DB.Find(&payments).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(payments)
}
