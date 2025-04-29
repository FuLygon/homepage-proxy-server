package models

// GotifyMessageStats messages stats from Gotify API
type GotifyMessageStats struct {
	Paging struct {
		Size  int `json:"size"`
		Since int `json:"since"`
		Limit int `json:"limit"`
	} `json:"paging"`
}
