package main

import (
	"log"
	"net/http"

	"car_inventory/config"
	"car_inventory/handlers"

	"github.com/gorilla/mux"
)

func main() {
	config.ConnectDB()

	r := mux.NewRouter()

	r.HandleFunc("/cars", handlers.GetAllCars).Methods("GET")
	r.HandleFunc("/cars/{id}", handlers.GetCar).Methods("GET")
	r.HandleFunc("/cars", handlers.CreateCar).Methods("POST")
	r.HandleFunc("/cars/{id}", handlers.UpdateCar).Methods("PUT")
	r.HandleFunc("/cars/{id}", handlers.DeleteCar).Methods("DELETE")

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
