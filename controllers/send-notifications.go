package controllers

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"notifyy.app/backend/cron"
)

func SendNotification(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	auth := os.Getenv("AUTH_HEADER")
	if authHeader == "" && authHeader != auth {
		c.JSON(401, gin.H{
			"error": "Invalid Authorization header",
		})
		return
	}
	cron.StartCron()
	c.JSON(200, gin.H{
		"message": "Sending Notifications",
	})
}
