package models

import "time"

type NpmAuthRequest struct {
	Identity string `json:"identity"`
	Secret   string `json:"secret"`
}

type NpmAuthResponse struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}
type NPMProxyHostsStats struct {
	ID          int      `json:"id"`
	Enabled     bool     `json:"enabled"`
	DomainNames []string `json:"domain_names"`
}

// NPMResponse holds the statistics about proxy hosts
type NPMResponse struct {
	Total    int `json:"total"`
	Enabled  int `json:"enabled"`
	Disabled int `json:"disabled"`
}
