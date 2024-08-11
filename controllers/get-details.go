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
	// body,_:=helpers.AuthorizeSpotify()
	// fmt.Printf("Body: %v\n", body.AccessToken)
	// track,_:=helpers.GetTrackInfo("11dFghVXANMlKmJXsNCbNl",body.AccessToken)
	// fmt.Printf("Track: %v\n", track.Name)
	c.JSON(http.StatusOK, response)
}