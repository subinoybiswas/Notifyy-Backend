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

func fetchConfiguration() ([]User, error) {
	location, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		fmt.Println("Error loading location:", err)
	}

	currentTime := time.Now().In(location)
	timeOnly := currentTime.Format("15:04:05")
	timeAfterOneHour := currentTime.Add(time.Hour).Format("15:04:05")
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
	msg := &messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title: "Price drop",
			Body:  "5% off all electronics",
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
	sendFCMNotification(tokens)
}
