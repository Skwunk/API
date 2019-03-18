package testapi

import (
	"log"
	"database/sql"
	_ "github.com/lib/pq"
)

var Database *sql.DB

func databaseConnection() (*sql.DB) {
	connstr := "user=web password=123456789 dbname=pqgotest sslmode=disable"
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Database connected")
	}
	return db
}

func NewDatabaseConnection() {
	Database = databaseConnection()
}