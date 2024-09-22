package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	utils "notifyy.app/backend/utils"
)

type Notification struct {
	ID 		string `json:"id"`
	Title 	string `json:"title"`
	Body 	string `json:"body"`
}


var notifications []Notification

func FetchNotifications(db *sql.DB) {
	rows, err := db.Query(`
   SELECT 
   	NOTIFICATIONID, 
  	MESSAGE, 
  	COALESCE(TITLE, '') AS TITLE
   FROM Notifications WHERE CHECKED=1`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query: %v\n", err)
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var notification Notification

		if err := rows.Scan(&notification.ID,&notification.Body,&notification.Title); err != nil {
			fmt.Println("Error scanning row:", err)
			log.Fatal(err)
		}

		notifications = append(notifications, notification)
		fmt.Printf("ID: %v, Title: %v, Body: %v\n", notification.ID, notification.Title, notification.Body)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error during rows iteration:", err)
		log.Fatal(err)
	}
}

func GetNotifications(c *gin.Context) {
	db := utils.DBConnection()
	defer db.Close()
	FetchNotifications(db)
	c.JSON(200, gin.H{
		"notifications": notifications,
	})

}