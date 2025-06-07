package routes

import (
	"github.com/ragsharan/vehicle-parking-app/reporting-service/handlers"

	"github.com/gorilla/mux"
)

func RegisterRoutes(rc *handlers.ReportingHandler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/report", rc.GetReport).Methods("GET")
	return router
}
