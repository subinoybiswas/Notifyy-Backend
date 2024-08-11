package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"notifyy.app/backend/helpers"
);

func GetDetails(c *gin.Context){
	response:=map[string]interface{}{
		"fact": "Music is a universal language.",
		"article": map[string]interface{}{
			"link":   "https://example.com/article",
			"heading": "The Evolution of Musicsss",
		}}
	body,_:=helpers.AuthorizeSpotify()
	fmt.Printf("Body: %v\n", body.AccessToken)
	track,_:=helpers.GetTrackInfo("11dFghVXANMlKmJXsNCbNl",body.AccessToken)
	fmt.Printf("Track: %v\n", track.Name)

response["spotify"]=map[string]interface{}{
		"name": track.Name,
		"artists": track.Artists,
		"album": track.Album,
		"image": track.Album.Images[0].URL,
		"release_date": track.Album.ReleaseDate,
}
	c.JSON(http.StatusOK, response)
}