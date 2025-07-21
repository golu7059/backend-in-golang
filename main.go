package main

import (
	"log"
	"net/http"

	"car_inventory/config"
	"car_inventory/handlers"
	"car_inventory/middlewares"

	"github.com/gorilla/mux"
)

func main() {
	config.ConnectDB()

	// Using gorilla mux instead of standard mux
	router := mux.NewRouter()

	// Routes with proper HTTP method handling
	router.HandleFunc("/cars", handlers.GetAllCars).Methods("GET")
	router.HandleFunc("/cars/{id}", handlers.GetCar).Methods("GET")
	router.HandleFunc("/cars", handlers.CreateCar).Methods("POST")
	router.HandleFunc("/cars/{id}", handlers.UpdateCar).Methods("PUT")
	router.HandleFunc("/cars/{id}", handlers.DeleteCar).Methods("DELETE")

	// Apply middlewares to the router
	loggedRouter := middlewares.Logger(router)
	secureRouter := middlewares.Security(loggedRouter)

	// start server
	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", secureRouter))
}
