package controllers

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)


type UpdateBody  struct{
	Alarm string `json:"alarm"`
	Surprise bool `json:"surprise"`
}

func UpdatePreferredTime(db *sql.DB, userID int, newPreferredTime string) error {
    query := `
        UPDATE NotifyUsers 
        SET PreferredTime = ? 
        WHERE UserID = ?
    `
    _, err := db.Exec(query, newPreferredTime, userID)
    if err != nil {
        return err
    }
    return nil
}

func UpdateConfiguration(c *gin.Context) {
	var requestBody UpdateBody
	
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		fmt.Printf("Error: %v", err)
		return
	}
	c.JSON(200, gin.H{
		"message": "UpdateConfiguration",
	})
	fmt.Printf("Alarm: %v, Surprise: %v\n", requestBody.Alarm, requestBody.Surprise)
}