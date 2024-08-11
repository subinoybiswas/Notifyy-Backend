package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"notifyy.app/backend/helpers"
	utils "notifyy.app/backend/utils"
);

type Article struct {
	Link 		string `json:"link"`
	Heading 	string `json:"heading"`
}

type Details struct {
	Fact 		string `json:"fact"`
	Article 	Article `json:"article"`
	Spotify 	*helpers.SpotifyTrackResponse `json:"spotify"`

}


var DetailsReturn Details



func FetchDetails() error {
    db := utils.DBConnection()
    defer db.Close()

    // Start a transaction to ensure atomicity
    tx, err := db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback() // Rollback in case of an error

    // Query to fetch the first unchecked item
    query := `SELECT id, fact, article_link, article_heading, spotify_link 
              FROM details 
              WHERE checked = 0 
              LIMIT 1`

    row := tx.QueryRow(query)

    var id int
    var fact, articleLink, articleHeading, spotifyLink string

    // Scan the result into variables
    err = row.Scan(&id, &fact, &articleLink, &articleHeading, &spotifyLink)
    if err == sql.ErrNoRows {
        fmt.Println("No unchecked details found.")
        
        // If no unchecked item found, set default values
        body, _ := helpers.AuthorizeSpotify()
        DetailsReturn.Spotify, _ = helpers.GetTrackInfo("7MdxDjBTpf7OdwJDnttkc0", body.AccessToken)
        DetailsReturn.Fact = "Music is a universal language."
        DetailsReturn.Article.Link = "https://example.com/article"
        DetailsReturn.Article.Heading = "The Evolution of Music"

        return nil
    } else if err != nil {
        return err
    }

    // Update the `checked` field to 1
    updateStmt := `UPDATE details 
                   SET checked = 1 
                   WHERE id = ?`

    _, err = tx.Exec(updateStmt, id)
    if err != nil {
        return err
    }

    // Commit the transaction
    if err := tx.Commit(); err != nil {
        return err
    }

    // If `spotifyLink` is not empty, fetch additional Spotify track info
    if spotifyLink != "" {
		link,err:=helpers.ExtractTrackID(spotifyLink)
		if err!=nil{
			return err
		}
        body, _ := helpers.AuthorizeSpotify()
        track, _ := helpers.GetTrackInfo(link, body.AccessToken)
        DetailsReturn.Spotify = track
    }

    // Set the fetched details to `DetailsReturn`
    DetailsReturn.Fact = fact
    DetailsReturn.Article.Link = articleLink
    DetailsReturn.Article.Heading = articleHeading

    return nil
}



func GetDetails(c *gin.Context){
	
	// response:=map[string]interface{}{
	// 	"fact": "Music is a universal language.",
	// 	"article": map[string]interface{}{
	// 		"link":   "https://example.com/article",
	// 		"heading": "The Evolution of Musicsss",
	// 	}}
	// body,_:=helpers.AuthorizeSpotify()
	// track,_:=helpers.GetTrackInfo("11dFghVXANMlKmJXsNCbNl",body.AccessToken)
	// fmt.Printf("Track: %v\n", track.Name)

	// response["spotify"]=map[string]interface{}{
	// 	"name": track.Name,
	// 	"artists": track.Artists,
	// 	"album": track.Album,
	// 	"image": track.Album.Images[0].URL,
	// 	"release_date": track.Album.ReleaseDate,
	// }
	if err:=FetchDetails();err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		fmt.Printf("Error: %v", err)
		return
	}
	c.JSON(http.StatusOK, DetailsReturn)
}