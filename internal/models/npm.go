package models

import "time"

// NpmAuthRequest authentication request
type NpmAuthRequest struct {
	Identity string `json:"identity"`
	Secret   string `json:"secret"`
}

// NpmAuthResponse authentication response
type NpmAuthResponse struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}

// NPMProxyHostsStats lists of proxy hosts
type NPMProxyHostsStats struct {
	ID          int      `json:"id"`
	Enabled     bool     `json:"enabled"`
	DomainNames []string `json:"domain_names"`
}

// NPMResponse minimalized response from NPM API
type NPMResponse struct {
	Total    int `json:"total"`
	Enabled  int `json:"enabled"`
	Disabled int `json:"disabled"`
}
