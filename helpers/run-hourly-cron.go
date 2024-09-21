package helpers

import (
	"fmt"
	"time"

	types "notifyy.app/backend/dbtype"
	utils "notifyy.app/backend/utils"
)

type User types.NotifyUsers

func fetchConfiguration() ([]User, error) {
	currentTime := time.Now()
	timeOnly := currentTime.Format("15:04:05")
	timeAfterOneHour := currentTime.Add(time.Hour).Format("15:04:05")
	fmt.Printf("Current Time: %s, Time After One Hour: %s\n", timeOnly, timeAfterOneHour)

	db := utils.DBConnection()
	defer db.Close()

	var users []User

	rows, err := db.Query(`
	SELECT UserID, NAME, EMAIL, PREFERREDTIME, SURPRISES 
		FROM NotifyUsers 
		WHERE PREFERREDTIME >= ? AND PREFERREDTIME < ?`, timeOnly, timeAfterOneHour)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.UserID, &user.Name, &user.Email, &user.PreferredTime, &user.Surprise); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func HourlyCron() {
	users, err := fetchConfiguration()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for _, user := range users {
		fmt.Printf("UserID: %s, Name: %s, Email: %s, Preferred Time: %s, Surprises: %d\n",
			user.UserID, user.Name, user.Email, user.PreferredTime, user.Surprise)
	}
}
