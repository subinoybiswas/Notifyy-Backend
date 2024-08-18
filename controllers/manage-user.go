package controllers

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	types "notifyy.app/backend/dbtype"
	utils "notifyy.app/backend/utils"
)
type UserData struct {
    ID                    string              `json:"id"`
    Username              string              `json:"username"`
    FirstName             string              `json:"first_name"`
    LastName              string              `json:"last_name"`
    ProfileImageURL       string              `json:"profile_image_url"`
    EmailAddress          string			  `json:"email_address"`
	FCMID				  string              `json:"fcm_id"`
}
func 	ManageUser(c *gin.Context) {
	db := utils.DBConnection()
	var user UserData

	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		fmt.Printf("Error: %v", err)
		return
	}
	var default_time="20:00:00"
	var default_surprise="0"
	// Check if the user exists
	var existingUser types.NotifyUsers
	err := db.QueryRow("SELECT UserID, NAME, EMAIL FROM NotifyUsers WHERE UserID = ?", user.ID).Scan(&existingUser.UserID, &existingUser.Name, &existingUser.Email)
	if err != nil {
		// If the user does not exist, create a new one
		if err == sql.ErrNoRows {
			_, err := db.Exec("INSERT INTO NotifyUsers (userid, name, email, fcmid, preferredtime,surprises) VALUES (?, ?, ?, ?, ?, ?)", user.ID, user.FirstName, user.EmailAddress,user.FCMID,default_time,default_surprise)
			if err != nil {
				fmt.Printf("Error: %v", err)

				c.JSON(500, gin.H{
					"error": "Failed to create user",
				})
				return
			}
			c.JSON(201, gin.H{
				"message": "User created",
			})
			return
		}
		fmt.Printf("Error: %v", err)
		c.JSON(500, gin.H{
			"error": "Failed to retrieve user",
		})

		return
	}

	// If the user exists, update their information
	_, err = db.Exec("UPDATE NotifyUsers SET name = ?, email = ?,fcmid=? WHERE userid = ?", user.FirstName, user.EmailAddress,user.FCMID, user.ID)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to update user",
		})
		fmt.Printf("Error: %v", err)

		return
	}

	c.JSON(200, gin.H{
		"message": "User updated",
	})
}