package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
);

func GetDetails(c *gin.Context){
	response:=map[string]interface{}{
		"fact": "Music is a universal language.",
		"article": map[string]interface{}{
			"link":   "https://example.com/article",
			"heading": "The Evolution of Musicsss",
		}}
	c.JSON(http.StatusOK, response)
}