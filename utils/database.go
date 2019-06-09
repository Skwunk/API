package utils

import (
	"log"
	"database/sql"
	_ "github.com/lib/pq"
)

var Database *sql.DB

func databaseConnection(c *Config) (*sql.DB) {
	connstr := "user=" + c.User + " password=" + c.Password + " dbname=" + c.DBName + " sslmode= " + c.SSLMode
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Database connected")
	}
	return db
}

func NewDatabaseConnection(c *Config) {
	Database = databaseConnection(c)
}