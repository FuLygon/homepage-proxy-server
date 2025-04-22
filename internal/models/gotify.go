package models

// GotifyMessageStats messages stats from Gotify API
type GotifyMessageStats struct {
	Paging struct {
		Size  int `json:"size"`
		Since int `json:"since"`
		Limit int `json:"limit"`
	} `json:"paging"`
}

// GotifyResponse minimalized response from Gotify API
type GotifyResponse struct {
	Applications int `json:"applications"`
	Clients      int `json:"clients"`
	Messages     int `json:"messages"`
}
