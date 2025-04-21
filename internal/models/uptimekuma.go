package models

// UptimeKumaStatusPageStats response from Uptime Kuma status page
type UptimeKumaStatusPageStats struct {
	Incident *struct {
		CreatedDate string `json:"createdDate"`
	} `json:"incident"`
}

// UptimeKumaHeartbeatStats response from Uptime Kuma heartbeat page
type UptimeKumaHeartbeatStats struct {
	HeartbeatList map[string][]struct {
		Status int `json:"status"`
	} `json:"heartbeatList"`
	UptimeList map[string]float64 `json:"uptimeList"`
}

// UptimeKumaResponse minimalized response from Uptime Kuma
type UptimeKumaResponse struct {
	SitesUp      int     `json:"sitesUp"`
	SitesDown    int     `json:"sitesDown"`
	Uptime       float64 `json:"uptime"`
	IncidentTime *int    `json:"incidentTime,omitempty"`
}
