package models

// WireGuardStatsResponse response for the WireGuard API
type WireGuardStatsResponse struct {
	Total     int `json:"total"`
	Connected int `json:"connected"`
}
