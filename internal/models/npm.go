package models

import "time"

// NPMAuthRequest authentication request
type NPMAuthRequest struct {
	Identity string `json:"identity"`
	Secret   string `json:"secret"`
}

// NPMAuthResponse authentication response
type NPMAuthResponse struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}

// NPMStatsResponse minimalized response from NPM API
type NPMStatsResponse struct {
	Enabled bool `json:"enabled"`
}
