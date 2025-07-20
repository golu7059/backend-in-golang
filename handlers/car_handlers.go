package handlers

import (
	"car_inventory/config"
	"car_inventory/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllCars(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT * FROM cars")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var cars []models.Car
	for rows.Next() {
		var car models.Car
		rows.Scan(&car.ID, &car.Name, &car.Model, &car.Brand, &car.Year, &car.Price)
		cars = append(cars, car)
	}
	json.NewEncoder(w).Encode(cars)
}

func GetCar(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	row := config.DB.QueryRow("SELECT * FROM cars WHERE id=$1", id)

	var car models.Car
	err := row.Scan(&car.ID, &car.Name, &car.Model, &car.Brand, &car.Year, &car.Price)
	if err != nil {
		http.Error(w, "Car not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(car)
}

func CreateCar(w http.ResponseWriter, r *http.Request) {
	var car models.Car
	json.NewDecoder(r.Body).Decode(&car)

	err := config.DB.QueryRow(
		"INSERT INTO cars(name, model, brand, year, price) VALUES($1, $2, $3, $4, $5) RETURNING id",
		car.Name, car.Model, car.Brand, car.Year, car.Price).Scan(&car.ID)
	if err != nil {
		http.Error(w, "Insert failed", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(car)
}

func UpdateCar(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var car models.Car
	json.NewDecoder(r.Body).Decode(&car)

	_, err := config.DB.Exec("UPDATE cars SET name=$1, model=$2, brand=$3, year=$4, price=$5 WHERE id=$6",
		car.Name, car.Model, car.Brand, car.Year, car.Price, id)
	if err != nil {
		http.Error(w, "Update failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteCar(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	_, err := config.DB.Exec("DELETE FROM cars WHERE id=$1", id)
	if err != nil {
		http.Error(w, "Delete failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
