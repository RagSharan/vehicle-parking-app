package routes

import (
	"github.com/gorilla/mux"
	"github.com/ragsharan/vehicle-parking-app/customer-service/handlers"
)

func RegisterRoutes(cc *handlers.CustomerHandler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/customers", cc.CreateCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", cc.GetCustomer).Methods("GET")
	router.HandleFunc("/customers", cc.ListCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}/vehicles", cc.AddVehicle).Methods("POST")
	router.HandleFunc("/customers/{id}/vehicles", cc.ListVehicles).Methods("GET")
	return router
}
