package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetTrackInfo(track string,token string) (*SpotifyTrackResponse, error) {
	// Set up the request
	apiURL := fmt.Sprintf("https://api.spotify.com/v1/tracks/%s", track)
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil,err
	}

	
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil,err
	}
	defer resp.Body.Close()

	// Read and print the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil,err
	}

	var trackInfo SpotifyTrackResponse
	if err := json.Unmarshal(body, &trackInfo);err!=nil{
		fmt.Println("Error unmarshalling response body:", err)
		return nil,err
	}
	return &trackInfo, nil
}
