package utils

import (
	"database/sql"
	"log"
	"os"
)

func DBConnection() *sql.DB {
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("libsql", url)

	if err != nil {
		log.Fatal(err)
	}
	return db
}
