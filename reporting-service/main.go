package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/ragsharan/vehicle-parking-app/reporting-service/handlers"
	"github.com/ragsharan/vehicle-parking-app/reporting-service/models"
	"github.com/ragsharan/vehicle-parking-app/reporting-service/routes"
)

func initConfig() {
	viper.SetConfigName("environment-override")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
}

func main() {
	initConfig()

	dsn := "host=" + viper.GetString("database.host") +
		" user=" + viper.GetString("database.user") +
		" password=" + viper.GetString("database.password") +
		" dbname=" + viper.GetString("database.name") +
		" port=" + viper.GetString("database.port") +
		" sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&models.Parking{}, &models.Vehicle{}, &models.Payment{})

	rc := &handlers.ReportingHandler{DB: db}
	router := routes.RegisterRoutes(rc)

	scheduler := gocron.NewScheduler(time.Local)
	scheduler.Every(1).Day().At("09:00").Do(rc.GenerateReport)
	scheduler.Every(1).Day().At("22:00").Do(rc.GenerateReport)
	scheduler.StartAsync()

	log.Printf("Starting reporting service on port %s", viper.GetString("server.port"))
	log.Fatal(http.ListenAndServe(":"+viper.GetString("server.port"), router))
}
