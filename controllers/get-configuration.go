package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	types "notifyy.app/backend/dbtype"
	utils "notifyy.app/backend/utils"
)


type GetBody  struct{
	ID string `json:"id"`	
}

var ExistingUser types.NotifyUsers

func FetchConfiguration(userID string) error {
	db := utils.DBConnection()
	err := db.QueryRow("SELECT UserID, NAME, EMAIL, PREFERREDTIME, SURPRISES FROM NotifyUsers WHERE USERID = ?", userID).Scan(&ExistingUser.UserID, &ExistingUser.Name, &ExistingUser.Email,&ExistingUser.PreferredTime,&ExistingUser.Surprise)
    if err != nil {
        return err
    }
    return nil
}

func GetConfiguration(c *gin.Context) {
	slug := c.Param("slug")
	// var requestBody GetBody
	
	// if err := c.BindJSON(&requestBody); err != nil {
	// 	c.JSON(400, gin.H{
	// 		"error": "Invalid request body",
	// 	})
	// 	fmt.Printf("Error: %v", err)
	// 	return
	// }
	if err:=FetchConfiguration(slug);err!=nil{
		c.JSON(500, gin.H{
			"error": "Internal Server Error",
		})
		fmt.Printf("Error: %v", err)
		return
	}
fmt.Printf("User: %v, Surprise: %v\n", ExistingUser.Email, ExistingUser.Surprise)
	c.JSON(200, gin.H{
		"alarm": ExistingUser.PreferredTime,
		"surprises": ExistingUser.Surprise,
	})

}