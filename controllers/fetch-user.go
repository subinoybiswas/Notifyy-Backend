package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	dbt "notifyy.app/backend/dbtype"
)

func FetchUserDetails(db *sql.DB, userID int) {
	rows, err := db.Query(`
    SELECT 
        NotifyUsers.UserID, 
        NotifyUsers.Name, 
        NotifyUsers.Email, 
        NotifyUsers.PreferredTime, 
        NotifyUsers.Surprises, 
        Notifications.NotificationID 
    FROM 
        NotifyUsers 
    LEFT JOIN 
        Notifications 
    ON 
        NotifyUsers.NotificationID = Notifications.NotificationID
    WHERE 
        NotifyUsers.UserID = ?
`, userID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query: %v\n", err)
		log.Fatal(err)
	}
	defer rows.Close()

	// Create a slice to hold user data
	var users []dbt.NotifyUsers

	for rows.Next() {
		var user dbt.NotifyUsers

		if err := rows.Scan(&user.UserID, &user.Name, &user.Email, &user.PreferredTime, &user.Surprise, &user.NotificationID); err != nil {
			fmt.Println("Error scanning row:", err)
			log.Fatal(err)
		}

		users = append(users, user)
		fmt.Printf("ID: %v, Name: %v, Email: %v, Preferred Time: %v, Surprise: %v, Notification ID: %v\n",
			user.UserID, user.Name, user.Email, user.PreferredTime, user.Surprise, user.NotificationID)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error during rows iteration:", err)
		log.Fatal(err)
	}
}
