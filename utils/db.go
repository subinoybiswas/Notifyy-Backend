package utils

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func DBConnection() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("libsql", url)

	if err != nil {
		log.Fatal(err)
	}
	return db
}
