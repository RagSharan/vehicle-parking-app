package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ragsharan/vehicle-parking-app/reporting-service/models"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type ReportingHandler struct {
	DB *gorm.DB
}

func (rc *ReportingHandler) GenerateReport() {
	var parkings []models.Parking
	rc.DB.Find(&parkings)

	report := "Daily Parking Report:\n\n"

	for _, p := range parkings {
		var vehicleCount int64
		var subscriptionCount int64
		var totalPayments float64

		rc.DB.Model(&models.Vehicle{}).Where("parking_id = ?", p.ID).Count(&vehicleCount)
		rc.DB.Model(&models.Vehicle{}).Where("parking_id = ? AND subscription = true", p.ID).Count(&subscriptionCount)
		rc.DB.Raw("SELECT COALESCE(SUM(amount), 0) FROM payments WHERE vehicle_id IN (SELECT id FROM vehicles WHERE parking_id = ?)", p.ID).Scan(&totalPayments)

		report += fmt.Sprintf("Parking: %s (%s)\n", p.Name, p.City)
		report += fmt.Sprintf("Address: %s\n", p.Address)
		report += fmt.Sprintf("Vehicles Parked: %d\n", vehicleCount)
		report += fmt.Sprintf("Subscription Vehicles: %d\n", subscriptionCount)
		report += fmt.Sprintf("Total Payments: $%.2f\n\n", totalPayments)
	}

	owners := viper.GetStringSlice("owners")
	for _, email := range owners {
		log.Printf("Sending report to %s:\n%s", email, report)
	}
}

func (rc *ReportingHandler) GetReport(w http.ResponseWriter, r *http.Request) {
	rc.GenerateReport()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Report generated and sent to owners."))
}
