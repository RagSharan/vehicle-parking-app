package routes

import (
	"github.com/ragsharan/vehicle-parking-app/vehicle-service/handlers"

	"github.com/gorilla/mux"
)

func RegisterRoutes(vc *handlers.VehicleHandler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/vehicles", vc.RegisterVehicle).Methods("POST")
	router.HandleFunc("/vehicles/{id}/exit", vc.MarkVehicleExit).Methods("PUT")
	router.HandleFunc("/vehicles", vc.ListVehicles).Methods("GET")
	return router
}
