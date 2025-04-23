package models

// UptimeKumaStatsResponse minimalized response from Uptime Kuma staus page API
type UptimeKumaStatsResponse struct {
	Incident *struct {
		CreatedDate string `json:"createdDate"`
	} `json:"incident,omitempty"`
}

// UptimeKumaStatsHeartbeatResponse minimalized response from Uptime Kuma staus page heartbeat API
type UptimeKumaStatsHeartbeatResponse struct {
	HeartbeatList map[string][]struct {
		Status int `json:"status"`
	} `json:"heartbeatList"`
	UptimeList map[string]float64 `json:"uptimeList"`
}
