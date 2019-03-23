package testapi

import (
	"log"
	"database/sql"
	_ "github.com/lib/pq"
)

var Database *sql.DB

func databaseConnection(c *Config) (*sql.DB) {
	log.Printf(c.User, c.Password, c.DBName)
	//connstr := "user=web password=123456789 dbname=pqgotest sslmode=disable"
	connstr := "user=" + c.User + " password=" + c.Password + " dbname=" + c.DBName + " sslmode= disable"
	log.Printf(connstr)
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