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

    tx, err := db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()

    query := `SELECT id, fact, article_link, article_heading, spotify_link 
              FROM details 
              WHERE  DATE(date) = DATE('now') 
              LIMIT 1`
	query2 := `SELECT id, fact, article_link, article_heading, spotify_link 
    		FROM details 
    		WHERE  checked = 0
    		LIMIT 1`
    row := tx.QueryRow(query)

    var id int
    var fact, articleLink, articleHeading, spotifyLink string

    err = row.Scan(&id, &fact, &articleLink, &articleHeading, &spotifyLink)
    if err == sql.ErrNoRows {

        fmt.Println("No unchecked details found.")
        row = tx.QueryRow(query2)
		err = row.Scan(&id, &fact, &articleLink, &articleHeading, &spotifyLink)
		if err == sql.ErrNoRows {
			body, _ := helpers.AuthorizeSpotify()
			DetailsReturn.Spotify, _ = helpers.GetTrackInfo("7MdxDjBTpf7OdwJDnttkc0", body.AccessToken)
			DetailsReturn.Fact = "Music is a universal language."
			DetailsReturn.Article.Link = "https://example.com/article"
			DetailsReturn.Article.Heading = "The Evolution of Music"
			return nil
		} else if err != nil {
			return err
		}
    } else if err != nil {
        return err
    }

updateStmt := `UPDATE details 
               SET checked = 1, date = DATE('now') 
               WHERE id = ?`

    _, err = tx.Exec(updateStmt, id)
    if err != nil {
        return err
    }

    if err := tx.Commit(); err != nil {
        return err
    }

    if spotifyLink != "" {
		link,err:=helpers.ExtractTrackID(spotifyLink)
		if err!=nil{
			return err
		}
        body, _ := helpers.AuthorizeSpotify()
        track, _ := helpers.GetTrackInfo(link, body.AccessToken)
        DetailsReturn.Spotify = track
    }

    DetailsReturn.Fact = fact
    DetailsReturn.Article.Link = articleLink
    DetailsReturn.Article.Heading = articleHeading

    return nil
}



func GetDetails(c *gin.Context){
	
	if err:=FetchDetails();err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		fmt.Printf("Error: %v", err)
		return
	}
	c.JSON(http.StatusOK, DetailsReturn)
}