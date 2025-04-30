package models

// YourSpotifySongsResponse songs listened stats from Your Spotify API
type YourSpotifySongsResponse struct {
	Count int64 `json:"count"`
}

// YourSpotifyTimeResponse time listened stats from Your Spotify API
type YourSpotifyTimeResponse YourSpotifySongsResponse

// YourSpotifyArtistsResponse artist listened stats from Your Spotify API
type YourSpotifyArtistsResponse struct {
	Artists []interface{} `json:"artists"`
}

// YourSpotifyResponse processed response from Your Spotify API
type YourSpotifyResponse struct {
	SongsListened   int64 `json:"songs_listened"`
	TimeListened    int64 `json:"time_listened"`
	ArtistsListened int   `json:"artists_listened"`
}
