package database

import (
	"database/sql"
	"log"
	"os"
)

var Database *sql.DB

func EstablishConnection() {
	databaseConStr := os.Getenv("DATABASE")
	var err error

	Database, err = sql.Open("postgres", databaseConStr)
	if err != nil {
		log.Fatalf("The following error occured. Err: %s", err)
	}
}
