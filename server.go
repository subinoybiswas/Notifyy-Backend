package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	controllers "notifyy.app/backend/controllers"
)

type User struct {
	ID      sql.NullInt64
	Name    sql.NullString
	Email   sql.NullString
	Credits sql.NullInt64
}

type NotifyUsers struct {
	UserID sql.NullInt64
	Name  sql.NullString
	Email sql.NullString
	PreferredTime sql.NullString
	Surprise sql.NullString
	NotificationID sql.NullInt64
}

func queryUsers(db *sql.DB) {
	rows, err := db.Query("SELECT id, name, email, credits FROM users")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query: %v\n", err)
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User

		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Credits); err != nil {
			fmt.Println("Error scanning row:", err)
			log.Fatal(err)
		}

		users = append(users, user)
		fmt.Printf("ID: %v, Name: %v, Email: %v, Credits: %v\n",
			user.ID.Int64, user.Name.String, user.Email.String, user.Credits.Int64)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error during rows iteration:", err)
		log.Fatal(err)
	}
}


func main() {
	// Connect to the database
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("libsql", url)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a new Gin router
	r := gin.New()

	// Define a route to query users
	r.GET("/details", controllers.GetDetails)
	r.GET("/users", func(c *gin.Context) {
		queryUsers(db)
	})
	r.POST("/configuration", controllers.UpdateConfiguration)
	// Start the server
	r.Run(":8080")
}
