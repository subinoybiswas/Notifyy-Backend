package helpers


// SpotifyTrackResponse represents the structure of the JSON response from Spotify API for a track
type SpotifyTrackResponse struct {
	Album          Album    `json:"album"`
	Artists        []Artist `json:"artists"`
	AvailableMarkets []string `json:"available_markets"`
	DiscNumber     int      `json:"disc_number"`
	DurationMs     int      `json:"duration_ms"`
	Explicit       bool     `json:"explicit"`
	ExternalIDs    ExternalIDs `json:"external_ids"`
	ExternalURLs   ExternalURLs `json:"external_urls"`
	Href           string   `json:"href"`
	ID             string   `json:"id"`
	IsLocal        bool     `json:"is_local"`
	Name           string   `json:"name"`
	Popularity     int      `json:"popularity"`
	PreviewURL     string   `json:"preview_url"`
	TrackNumber    int      `json:"track_number"`
	Type           string   `json:"type"`
	URI            string   `json:"uri"`
}

// Album represents the album details in the Spotify track response
type Album struct {
	AlbumType        string   `json:"album_type"`
	Artists          []Artist `json:"artists"`
	AvailableMarkets []string `json:"available_markets"`
	ExternalURLs     ExternalURLs `json:"external_urls"`
	Href             string   `json:"href"`
	ID               string   `json:"id"`
	Images           []Image  `json:"images"`
	Name             string   `json:"name"`
	ReleaseDate      string   `json:"release_date"`
	ReleaseDatePrecision string `json:"release_date_precision"`
	TotalTracks      int      `json:"total_tracks"`
	Type             string   `json:"type"`
	URI              string   `json:"uri"`
}

// Artist represents the artist details in the Spotify track and album response
type Artist struct {
	ExternalURLs ExternalURLs `json:"external_urls"`
	Href         string   `json:"href"`
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Type         string   `json:"type"`
	URI          string   `json:"uri"`
}

// ExternalURLs represents the external URLs in the Spotify track response
type ExternalURLs struct {
	Spotify string `json:"spotify"`
}

// ExternalIDs represents the external IDs in the Spotify track response
type ExternalIDs struct {
	ISRC string `json:"isrc"`
}

// Image represents an image in the album's images array in the Spotify track response
type Image struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
