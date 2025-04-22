package models

// UptimeKumaHeartbeatStats response from Uptime Kuma heartbeat page
type UptimeKumaHeartbeatStats struct {
	HeartbeatList map[string][]struct {
		Status int `json:"status"`
	} `json:"heartbeatList"`
	UptimeList map[string]float64 `json:"uptimeList"`
}

// UptimeKumaResponse minimalized response from Uptime Kuma
type UptimeKumaResponse struct {
	SitesUp   int     `json:"sites-up"`
	SitesDown int     `json:"sites-down"`
	Uptime    float64 `json:"uptime"`
}
