package helpers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"firebase.google.com/go/v4/messaging"
	fcm "github.com/appleboy/go-fcm"
	"github.com/joho/godotenv"
	types "notifyy.app/backend/dbtype"
	utils "notifyy.app/backend/utils"
)

type User types.NotifyUsers

type notification struct {
	Id      int
	Title   string
	Message string
	Checked bool
}

var location *time.Location
var timeOnly string
var timeAfterOneHour string

func init() {
	var err error
	location, err = time.LoadLocation("Asia/Kolkata")
	if err != nil {
		fmt.Println("Error loading location:", err)
	}

	currentTime := time.Now().In(location)
	timeOnly = currentTime.Format("15:04:05")
	timeAfterOneHour = currentTime.Add(time.Hour).Format("15:04:05")
}

func fetchConfiguration() ([]User, error) {

	fmt.Printf("Current Time: %s, Time After One Hour: %s\n", timeOnly, timeAfterOneHour)

	db := utils.DBConnection()
	defer db.Close()

	var users []User

	rows, err := db.Query(`
	SELECT UserID, NAME, EMAIL, FCMID
		FROM NotifyUsers 
		WHERE PREFERREDTIME >= ? AND PREFERREDTIME < ?`, timeOnly, timeAfterOneHour)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.UserID, &user.Name, &user.Email, &user.FCMID); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func fetchNotificationContent() (notification, error) {
	var currentNotification notification
	db := utils.DBConnection()
	defer db.Close()
	rows, err := db.Query("SELECT notificationID,title,message,checked FROM notifications WHERE checked = 0 LIMIT 1")
	if err != nil {
		return notification{}, err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&currentNotification.Id, &currentNotification.Title, &currentNotification.Message, &currentNotification.Checked)
		if err != nil {
			return notification{}, err
		}
		timeOnlyParsed, err := time.Parse("15:04:05", timeOnly)
		if err != nil {
			fmt.Println("Error parsing time:", err)
		}
		fmt.Printf("Current Time: %d, Time After One Hour: %s\n", timeOnlyParsed.Hour(), timeAfterOneHour)
		if timeOnlyParsed.Hour() >= 23 {
			_, err := db.Query("UPDATE notifications SET checked = 1 WHERE NotificationID= ? ", currentNotification.Id)
			if err != nil {
				fmt.Println("Error updating notifications:", err)
			}
		}
		return currentNotification, nil

	}
	return notification{

	}, nil

}
func sendFCMNotification(tokens []string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	path := os.Getenv("SECRET_LOCATION")
	ctx := context.Background()
	client, err := fcm.NewClient(
		ctx,
		fcm.WithCredentialsFile(path),
	)
	if err != nil {
		log.Fatal(err)
	}

	registrationTokens := tokens
	sendNotification, err := fetchNotificationContent()
	if err != nil {
		log.Fatal(err)
	}
	msg := &messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title: sendNotification.Title,
			Body:  sendNotification.Message,
		},
		Tokens: registrationTokens,
	}
	client.SendMulticast(
		ctx,
		msg,
	)
	if err != nil {
		log.Fatal(err)
	}

}
func HourlyCron() {
	users, err := fetchConfiguration()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	var tokens []string
	for _, user := range users {
		tokens = append(tokens, user.FCMID)
		fmt.Printf(user.FCMID)
	}
	if len(tokens) == 0 {
		return
	}
	sendFCMNotification(tokens)
}
