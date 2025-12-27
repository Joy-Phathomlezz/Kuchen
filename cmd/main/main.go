package main

import (
	"log"
	"net/http"

	"github.com/Joy-Phathomlezz/Kuchen/internals/database"
	"github.com/Joy-Phathomlezz/Kuchen/internals/handlers"
	"github.com/Joy-Phathomlezz/Kuchen/internals/models"
	"github.com/Joy-Phathomlezz/Kuchen/internals/routes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Connect to the database and get the GORM DB instance
	database, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, _ := database.DB()
	defer sqlDB.Close()

	log.Println("Database connected ")

	// Auto-migrate the Cake model to create the table if it doesn't exist
	database.AutoMigrate(&models.Cake{})

	// Create a CakeHandler with the DB instance for handling requests
	handler := handlers.NewCakeHandler(database)
	routes.CakeStoreResult(router, handler)

	err = http.ListenAndServe(":9010", router)
	if err != nil {
		log.Print("Server failed , ", err)
	}
}
