package adhoc

import (
	"log"
	"notifyy.app/backend/utils"
)

type Notification struct {
	Title string
	Message  string
}

func AddNotifications() error {
	db := utils.DBConnection()
	stmt, err := db.Prepare("INSERT INTO notifications (title, message) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, notification := range Notifications {
		_, err := stmt.Exec(notification.Title, notification.Message)
		if err != nil {
			log.Printf("Error inserting notification: %v", err)
		}
	}
	return nil
}
