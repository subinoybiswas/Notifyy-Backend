package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func DBConnection() *sql.DB {
	url := os.Getenv("DATABASE_URL")
	fmt.Printf("URL: %v\n", url)
	db, err := sql.Open("libsql", url)

	if err != nil {
		log.Fatal(err)
	}
	return db
}
