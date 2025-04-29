package models

// LinkwardenStatsResponse minimalized response from Linkwarden Collections and Tags API
type LinkwardenStatsResponse struct {
	Response []struct {
		Count struct {
			Links int `json:"links"`
		} `json:"_count"`
	} `json:"response"`
}
