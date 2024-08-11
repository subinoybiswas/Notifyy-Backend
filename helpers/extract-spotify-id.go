package helpers

import (
	"fmt"
	"strings"
)

func ExtractTrackID(url string) (string, error) {
	const prefix = "https://open.spotify.com/track/"


	if !strings.HasPrefix(url, prefix) {
		return "", fmt.Errorf("invalid Spotify track URL")
	}

	trackID := strings.TrimPrefix(url, prefix)

	if idx := strings.Index(trackID, "?"); idx != -1 {
		trackID = trackID[:idx]
	}

	return trackID, nil
}