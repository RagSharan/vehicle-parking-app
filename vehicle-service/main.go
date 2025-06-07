package main

import (
	"log"
	"net/http"

	"github.com/ragsharan/vehicle-parking-app/vehicle-service/handlers"
	"github.com/ragsharan/vehicle-parking-app/vehicle-service/models"
	"github.com/ragsharan/vehicle-parking-app/vehicle-service/routes"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initConfig() {
	viper.SetConfigName("environment-override")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
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

	db.AutoMigrate(&models.Vehicle{})

	vc := &handlers.VehicleHandler{DB: db}
	router := routes.RegisterRoutes(vc)

	log.Printf("Starting server on port %s", viper.GetString("server.port"))
	log.Fatal(http.ListenAndServe(":"+viper.GetString("server.port"), router))
}
