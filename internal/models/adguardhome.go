package models

// AdGuardHomeResponse Adguard Home API after processing
type AdGuardHomeResponse struct {
	Queries  int     `json:"queries"`
	Blocked  int     `json:"blocked"`
	Filtered int     `json:"filtered"`
	Latency  float64 `json:"latency"`
}

// AdGuardHomeStats minimalized response from AdGuard Home API
type AdGuardHomeStats struct {
	NumDNSQueries           int     `json:"num_dns_queries"`
	NumBlockedFiltering     int     `json:"num_blocked_filtering"`
	NumReplacedSafeBrowsing int     `json:"num_replaced_safebrowsing"`
	NumReplacedSafeSearch   int     `json:"num_replaced_safesearch"`
	NumReplacedParental     int     `json:"num_replaced_parental"`
	AvgProcessingTime       float64 `json:"avg_processing_time"`
}
