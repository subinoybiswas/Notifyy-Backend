package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	utils "notifyy.app/backend/utils"
)


type UpdateBody  struct{
	ID string `json:"id"`	
	Alarm string `json:"alarm"`
	Surprise bool `json:"surprise"`
}

func UpdateDetails(userID string, newPreferredTime string,newSurprises bool) error {
	var surprise int
	if newSurprises{
		surprise=1;
	}else{
		surprise=0;
	}
	db := utils.DBConnection()
    query := `
        UPDATE NotifyUsers 
        SET PreferredTime = ?, Surprises = ?
        WHERE UserID = ?
    `
    _, err := db.Exec(query, newPreferredTime,surprise, userID)
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
	if err:=UpdateDetails("1", requestBody.Alarm,requestBody.Surprise);err!=nil{
		c.JSON(500, gin.H{
			"error": "Internal Server Error",
		})
		fmt.Printf("Error: %v", err)
		return
	}
	c.JSON(200, gin.H{
		"message": "UpdateConfiguration",
	})
	fmt.Printf("Alarm: %v, Surprise: %v\n", requestBody.Alarm, requestBody.Surprise)
}