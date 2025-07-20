package config

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	connStr := "user=manage password=Golu@1234 dbname=cards_db sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to DB:", err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal("Ping error:", err)
	}
	log.Println("Connected to the database")
}