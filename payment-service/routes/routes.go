package routes

import (
	handler "github.com/ragsharan/vehicle-parking-app/payment-service/handlers"

	"github.com/gorilla/mux"
)

func RegisterRoutes(pc *handler.PaymentHandler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/payments", pc.CreatePayment).Methods("POST")
	router.HandleFunc("/payments/{id}", pc.GetPayment).Methods("GET")
	router.HandleFunc("/payments", pc.ListPayments).Methods("GET")
	return router
}
